package main

import (
	"encoding/json"
	"fmt"
	"net"

	logs "github.com/appscode/go/log/golog"
	"github.com/appscode/stash/apis/repositories"
	"github.com/appscode/stash/apis/repositories/install"
	repov1alpha1 "github.com/appscode/stash/apis/repositories/v1alpha1"
	"github.com/appscode/stash/apis/stash"
	stashv1alpha1 "github.com/appscode/stash/apis/stash/v1alpha1"
	"github.com/golang/glog"
	admissionv1beta1 "k8s.io/api/admission/v1beta1"
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/kube-openapi/pkg/builder"
	"k8s.io/kube-openapi/pkg/common"
)

var (
	groupFactoryRegistry = make(announced.APIGroupFactoryRegistry)
	registry             = registered.NewOrDie("")
	Scheme               = runtime.NewScheme()
	Codecs               = serializer.NewCodecFactory(Scheme)
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	Install(groupFactoryRegistry, registry, Scheme)
	install.Install(groupFactoryRegistry, registry, Scheme)

	// we need to add the options to empty v1
	// TODO fix the server code to avoid this
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})

	// TODO: keep the generic API server from wanting this
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)

	// --------------
	recommendedOptions := genericoptions.NewRecommendedOptions("/registry/foo.com", Codecs.LegacyCodec(admissionv1beta1.SchemeGroupVersion))
	// recommendedOptions.SecureServing = nil
	recommendedOptions.SecureServing.BindPort = 8443
	recommendedOptions.Etcd = nil
	recommendedOptions.Authentication = nil
	recommendedOptions.Authorization = nil
	recommendedOptions.CoreAPI = nil

	// TODO have a "real" external address
	if err := recommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		glog.Fatal(fmt.Errorf("error creating self-signed certificates: %v", err))
	}

	serverConfig := genericapiserver.NewRecommendedConfig(Codecs)
	if err := recommendedOptions.ApplyTo(serverConfig); err != nil {
		glog.Fatalln(err)
	}
	serverConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(func(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
		out := map[string]common.OpenAPIDefinition{}
		for k, v := range stashv1alpha1.GetOpenAPIDefinitions(ref) {
			out[k] = v
		}
		for k, v := range repov1alpha1.GetOpenAPIDefinitions(ref) {
			out[k] = v
		}
		return out
	}, Scheme)
	serverConfig.OpenAPIConfig.Info.Title = "stash-server"
	serverConfig.OpenAPIConfig.Info.Version = stashv1alpha1.SchemeGroupVersion.Version
	serverConfig.OpenAPIConfig.IgnorePrefixes = []string{
		"/swaggerapi",
		"/apis/admission.stash.appscode.com/v1alpha1/restics",
		"/apis/admission.stash.appscode.com/v1alpha1/recoveries",
		"/apis/admission.stash.appscode.com/v1alpha1/deployments",
		"/apis/admission.stash.appscode.com/v1alpha1/daemonsets",
		"/apis/admission.stash.appscode.com/v1alpha1/statefulsets",
		"/apis/admission.stash.appscode.com/v1alpha1/replicationcontrollers",
		"/apis/admission.stash.appscode.com/v1alpha1/replicasets",
	}

	genericServer, err := serverConfig.Complete().New("stash-server", genericapiserver.EmptyDelegate) // completion is done in Complete, no need for a second time
	if err != nil {
		glog.Fatalln(err)
	}

	// k8s.io/apiserver/pkg/server/routes/openapi.go
	{
		apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(stash.GroupName, registry, Scheme, metav1.ParameterCodec, Codecs)
		apiGroupInfo.GroupMeta.GroupVersion = stashv1alpha1.SchemeGroupVersion
		v1alpha1storage := map[string]rest.Storage{}
		v1alpha1storage[stashv1alpha1.ResourcePluralRestic] = NewREST(
			stashv1alpha1.SchemeGroupVersion.WithKind(stashv1alpha1.ResourceKindRestic),
			&stashv1alpha1.Restic{},
			&stashv1alpha1.ResticList{})
		v1alpha1storage[stashv1alpha1.ResourcePluralRecovery] = NewREST(
			stashv1alpha1.SchemeGroupVersion.WithKind(stashv1alpha1.ResourceKindRecovery),
			&stashv1alpha1.Recovery{},
			&stashv1alpha1.RecoveryList{})
		apiGroupInfo.VersionedResourcesStorageMap["v1alpha1"] = v1alpha1storage

		if err := genericServer.InstallAPIGroup(&apiGroupInfo); err != nil {
			glog.Fatal(err)
		}
	}

	{
		apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(repositories.GroupName, registry, Scheme, metav1.ParameterCodec, Codecs)
		apiGroupInfo.GroupMeta.GroupVersion = repov1alpha1.SchemeGroupVersion
		v1alpha1storage := map[string]rest.Storage{}
		v1alpha1storage[repov1alpha1.ResourcePluralSnapshot] = NewREST(
			repov1alpha1.SchemeGroupVersion.WithKind(repov1alpha1.ResourceKindSnapshot),
			&repov1alpha1.Snapshot{},
			&repov1alpha1.SnapshotList{})
		apiGroupInfo.VersionedResourcesStorageMap["v1alpha1"] = v1alpha1storage

		if err := genericServer.InstallAPIGroup(&apiGroupInfo); err != nil {
			glog.Fatal(err)
		}
	}

	spec, err := builder.BuildOpenAPISpec(genericServer.Handler.GoRestfulContainer.RegisteredWebServices(), serverConfig.OpenAPIConfig)
	if err != nil {
		glog.Fatal(err)
	}
	data, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		glog.Fatal(err)
	}
	fmt.Println(string(data))
}

// Install registers the API group and adds types to a scheme
func Install(groupFactoryRegistry announced.APIGroupFactoryRegistry, registry *registered.APIRegistrationManager, scheme *runtime.Scheme) {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:              stash.GroupName,
			RootScopedKinds:        sets.NewString(),
			VersionPreferenceOrder: []string{stashv1alpha1.SchemeGroupVersion.Version},
			// AddInternalObjectsToScheme: stash.AddToScheme,
		},
		announced.VersionToSchemeFunc{
			stashv1alpha1.SchemeGroupVersion.Version: stashv1alpha1.AddToScheme,
		},
	).Announce(groupFactoryRegistry).RegisterAndEnable(registry, scheme); err != nil {
		panic(err)
	}
}
