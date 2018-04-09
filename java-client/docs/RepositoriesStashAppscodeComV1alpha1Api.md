# RepositoriesStashAppscodeComV1alpha1Api

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**deleteRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot**](RepositoriesStashAppscodeComV1alpha1Api.md#deleteRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot) | **DELETE** /apis/repositories.stash.appscode.com/v1alpha1/namespaces/{namespace}/snapshots/{name} | 
[**getRepositoriesStashAppscodeComV1alpha1APIResources**](RepositoriesStashAppscodeComV1alpha1Api.md#getRepositoriesStashAppscodeComV1alpha1APIResources) | **GET** /apis/repositories.stash.appscode.com/v1alpha1/ | 
[**listRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot**](RepositoriesStashAppscodeComV1alpha1Api.md#listRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot) | **GET** /apis/repositories.stash.appscode.com/v1alpha1/namespaces/{namespace}/snapshots | 
[**listRepositoriesStashAppscodeComV1alpha1SnapshotForAllNamespaces**](RepositoriesStashAppscodeComV1alpha1Api.md#listRepositoriesStashAppscodeComV1alpha1SnapshotForAllNamespaces) | **GET** /apis/repositories.stash.appscode.com/v1alpha1/snapshots | 
[**readRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot**](RepositoriesStashAppscodeComV1alpha1Api.md#readRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot) | **GET** /apis/repositories.stash.appscode.com/v1alpha1/namespaces/{namespace}/snapshots/{name} | 


<a name="deleteRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot"></a>
# **deleteRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot**
> IoK8sApimachineryPkgApisMetaV1Status deleteRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot(name, namespace, pretty)



delete a Snapshot

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.RepositoriesStashAppscodeComV1alpha1Api;


RepositoriesStashAppscodeComV1alpha1Api apiInstance = new RepositoriesStashAppscodeComV1alpha1Api();
String name = "name_example"; // String | name of the Snapshot
String namespace = "namespace_example"; // String | object name and auth scope, such as for teams and projects
String pretty = "pretty_example"; // String | If 'true', then the output is pretty printed.
try {
    IoK8sApimachineryPkgApisMetaV1Status result = apiInstance.deleteRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot(name, namespace, pretty);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling RepositoriesStashAppscodeComV1alpha1Api#deleteRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **String**| name of the Snapshot |
 **namespace** | **String**| object name and auth scope, such as for teams and projects |
 **pretty** | **String**| If &#39;true&#39;, then the output is pretty printed. | [optional]

### Return type

[**IoK8sApimachineryPkgApisMetaV1Status**](IoK8sApimachineryPkgApisMetaV1Status.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

<a name="getRepositoriesStashAppscodeComV1alpha1APIResources"></a>
# **getRepositoriesStashAppscodeComV1alpha1APIResources**
> IoK8sApimachineryPkgApisMetaV1APIResourceList getRepositoriesStashAppscodeComV1alpha1APIResources()



get available resources

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.RepositoriesStashAppscodeComV1alpha1Api;


RepositoriesStashAppscodeComV1alpha1Api apiInstance = new RepositoriesStashAppscodeComV1alpha1Api();
try {
    IoK8sApimachineryPkgApisMetaV1APIResourceList result = apiInstance.getRepositoriesStashAppscodeComV1alpha1APIResources();
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling RepositoriesStashAppscodeComV1alpha1Api#getRepositoriesStashAppscodeComV1alpha1APIResources");
    e.printStackTrace();
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**IoK8sApimachineryPkgApisMetaV1APIResourceList**](IoK8sApimachineryPkgApisMetaV1APIResourceList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, application/vnd.kubernetes.protobuf
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

<a name="listRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot"></a>
# **listRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot**
> ComGithubAppscodeStashApisRepositoriesV1alpha1SnapshotList listRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot(namespace, _continue, fieldSelector, includeUninitialized, labelSelector, limit, pretty, resourceVersion, timeoutSeconds, watch)



list objects of kind Snapshot

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.RepositoriesStashAppscodeComV1alpha1Api;


RepositoriesStashAppscodeComV1alpha1Api apiInstance = new RepositoriesStashAppscodeComV1alpha1Api();
String namespace = "namespace_example"; // String | object name and auth scope, such as for teams and projects
String _continue = "_continue_example"; // String | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.
String fieldSelector = "fieldSelector_example"; // String | A selector to restrict the list of returned objects by their fields. Defaults to everything.
Boolean includeUninitialized = true; // Boolean | If true, partially initialized resources are included in the response.
String labelSelector = "labelSelector_example"; // String | A selector to restrict the list of returned objects by their labels. Defaults to everything.
Integer limit = 56; // Integer | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.
String pretty = "pretty_example"; // String | If 'true', then the output is pretty printed.
String resourceVersion = "resourceVersion_example"; // String | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.
Integer timeoutSeconds = 56; // Integer | Timeout for the list/watch call.
Boolean watch = true; // Boolean | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.
try {
    ComGithubAppscodeStashApisRepositoriesV1alpha1SnapshotList result = apiInstance.listRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot(namespace, _continue, fieldSelector, includeUninitialized, labelSelector, limit, pretty, resourceVersion, timeoutSeconds, watch);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling RepositoriesStashAppscodeComV1alpha1Api#listRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **String**| object name and auth scope, such as for teams and projects |
 **_continue** | **String**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | [optional]
 **fieldSelector** | **String**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | [optional]
 **includeUninitialized** | **Boolean**| If true, partially initialized resources are included in the response. | [optional]
 **labelSelector** | **String**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | [optional]
 **limit** | **Integer**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | [optional]
 **pretty** | **String**| If &#39;true&#39;, then the output is pretty printed. | [optional]
 **resourceVersion** | **String**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | [optional]
 **timeoutSeconds** | **Integer**| Timeout for the list/watch call. | [optional]
 **watch** | **Boolean**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | [optional]

### Return type

[**ComGithubAppscodeStashApisRepositoriesV1alpha1SnapshotList**](ComGithubAppscodeStashApisRepositoriesV1alpha1SnapshotList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

<a name="listRepositoriesStashAppscodeComV1alpha1SnapshotForAllNamespaces"></a>
# **listRepositoriesStashAppscodeComV1alpha1SnapshotForAllNamespaces**
> ComGithubAppscodeStashApisRepositoriesV1alpha1SnapshotList listRepositoriesStashAppscodeComV1alpha1SnapshotForAllNamespaces(_continue, fieldSelector, includeUninitialized, labelSelector, limit, pretty, resourceVersion, timeoutSeconds, watch)



list objects of kind Snapshot

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.RepositoriesStashAppscodeComV1alpha1Api;


RepositoriesStashAppscodeComV1alpha1Api apiInstance = new RepositoriesStashAppscodeComV1alpha1Api();
String _continue = "_continue_example"; // String | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.
String fieldSelector = "fieldSelector_example"; // String | A selector to restrict the list of returned objects by their fields. Defaults to everything.
Boolean includeUninitialized = true; // Boolean | If true, partially initialized resources are included in the response.
String labelSelector = "labelSelector_example"; // String | A selector to restrict the list of returned objects by their labels. Defaults to everything.
Integer limit = 56; // Integer | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.
String pretty = "pretty_example"; // String | If 'true', then the output is pretty printed.
String resourceVersion = "resourceVersion_example"; // String | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.
Integer timeoutSeconds = 56; // Integer | Timeout for the list/watch call.
Boolean watch = true; // Boolean | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.
try {
    ComGithubAppscodeStashApisRepositoriesV1alpha1SnapshotList result = apiInstance.listRepositoriesStashAppscodeComV1alpha1SnapshotForAllNamespaces(_continue, fieldSelector, includeUninitialized, labelSelector, limit, pretty, resourceVersion, timeoutSeconds, watch);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling RepositoriesStashAppscodeComV1alpha1Api#listRepositoriesStashAppscodeComV1alpha1SnapshotForAllNamespaces");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **_continue** | **String**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | [optional]
 **fieldSelector** | **String**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | [optional]
 **includeUninitialized** | **Boolean**| If true, partially initialized resources are included in the response. | [optional]
 **labelSelector** | **String**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | [optional]
 **limit** | **Integer**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | [optional]
 **pretty** | **String**| If &#39;true&#39;, then the output is pretty printed. | [optional]
 **resourceVersion** | **String**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | [optional]
 **timeoutSeconds** | **Integer**| Timeout for the list/watch call. | [optional]
 **watch** | **Boolean**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | [optional]

### Return type

[**ComGithubAppscodeStashApisRepositoriesV1alpha1SnapshotList**](ComGithubAppscodeStashApisRepositoriesV1alpha1SnapshotList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

<a name="readRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot"></a>
# **readRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot**
> ComGithubAppscodeStashApisRepositoriesV1alpha1Snapshot readRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot(name, namespace, pretty)



read the specified Snapshot

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.RepositoriesStashAppscodeComV1alpha1Api;


RepositoriesStashAppscodeComV1alpha1Api apiInstance = new RepositoriesStashAppscodeComV1alpha1Api();
String name = "name_example"; // String | name of the Snapshot
String namespace = "namespace_example"; // String | object name and auth scope, such as for teams and projects
String pretty = "pretty_example"; // String | If 'true', then the output is pretty printed.
try {
    ComGithubAppscodeStashApisRepositoriesV1alpha1Snapshot result = apiInstance.readRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot(name, namespace, pretty);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling RepositoriesStashAppscodeComV1alpha1Api#readRepositoriesStashAppscodeComV1alpha1NamespacedSnapshot");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **String**| name of the Snapshot |
 **namespace** | **String**| object name and auth scope, such as for teams and projects |
 **pretty** | **String**| If &#39;true&#39;, then the output is pretty printed. | [optional]

### Return type

[**ComGithubAppscodeStashApisRepositoriesV1alpha1Snapshot**](ComGithubAppscodeStashApisRepositoriesV1alpha1Snapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

