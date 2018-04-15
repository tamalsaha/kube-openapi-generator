# ComGithubAppscodeStashApisStashV1alpha1ResticSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**backend** | [**ComGithubAppscodeStashApisStashV1alpha1Backend**](ComGithubAppscodeStashApisStashV1alpha1Backend.md) |  | [optional] 
**file_groups** | [**list[ComGithubAppscodeStashApisStashV1alpha1FileGroup]**](ComGithubAppscodeStashApisStashV1alpha1FileGroup.md) |  | [optional] 
**image_pull_secrets** | [**list[IoK8sApiCoreV1LocalObjectReference]**](IoK8sApiCoreV1LocalObjectReference.md) | ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod | [optional] 
**paused** | **bool** | Indicates that the Restic is paused from taking backup. Default value is &#39;false&#39; | [optional] 
**resources** | [**IoK8sApiCoreV1ResourceRequirements**](IoK8sApiCoreV1ResourceRequirements.md) | Compute Resources required by the sidecar container. | [optional] 
**retention_policies** | [**list[ComGithubAppscodeStashApisStashV1alpha1RetentionPolicy]**](ComGithubAppscodeStashApisStashV1alpha1RetentionPolicy.md) |  | [optional] 
**schedule** | **str** |  | [optional] 
**selector** | [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | [optional] 
**type** | **str** | https://github.com/appscode/stash/issues/225 | [optional] 
**volume_mounts** | [**list[IoK8sApiCoreV1VolumeMount]**](IoK8sApiCoreV1VolumeMount.md) | Pod volumes to mount into the sidecar container&#39;s filesystem. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


