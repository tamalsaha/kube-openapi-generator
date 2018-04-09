package main

import (
	"fmt"

	logs "github.com/appscode/go/log/golog"
	repoinstall "github.com/appscode/stash/apis/repositories/install"
	repov1alpha1 "github.com/appscode/stash/apis/repositories/v1alpha1"
	stashinstall "github.com/appscode/stash/apis/stash/install"
	stashv1alpha1 "github.com/appscode/stash/apis/stash/v1alpha1"
	"github.com/go-openapi/spec"
	"github.com/golang/glog"
	"github.com/tamalsaha/kube-openapi-generator/lib"
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/kube-openapi/pkg/common"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	var (
		groupFactoryRegistry = make(announced.APIGroupFactoryRegistry)
		registry             = registered.NewOrDie("")
		Scheme               = runtime.NewScheme()
		Codecs               = serializer.NewCodecFactory(Scheme)
	)

	stashinstall.Install(groupFactoryRegistry, registry, Scheme)
	repoinstall.Install(groupFactoryRegistry, registry, Scheme)

	spec, err := lib.RenderOpenAPISpec(lib.Config{
		Registry: registry,
		Scheme:   Scheme,
		Codecs:   Codecs,
		Info: spec.InfoProps{
			Title:   "stash-server",
			Version: "v0",
		},
		OpenAPIDefinitions: []common.GetOpenAPIDefinitions{
			stashv1alpha1.GetOpenAPIDefinitions,
			repov1alpha1.GetOpenAPIDefinitions,
		},
		Resources: []schema.GroupVersionResource{
			stashv1alpha1.SchemeGroupVersion.WithResource(stashv1alpha1.ResourcePluralRestic),
			stashv1alpha1.SchemeGroupVersion.WithResource(stashv1alpha1.ResourcePluralRepository),
			stashv1alpha1.SchemeGroupVersion.WithResource(stashv1alpha1.ResourcePluralRecovery),
			repov1alpha1.SchemeGroupVersion.WithResource(repov1alpha1.ResourcePluralSnapshot),
		},
	})
	if err != nil {
		glog.Fatal(err)
	}
	fmt.Println(spec)
}
