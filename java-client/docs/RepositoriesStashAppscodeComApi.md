# RepositoriesStashAppscodeComApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getRepositoriesStashAppscodeComAPIGroup**](RepositoriesStashAppscodeComApi.md#getRepositoriesStashAppscodeComAPIGroup) | **GET** /apis/repositories.stash.appscode.com/ | 


<a name="getRepositoriesStashAppscodeComAPIGroup"></a>
# **getRepositoriesStashAppscodeComAPIGroup**
> IoK8sApimachineryPkgApisMetaV1APIGroup getRepositoriesStashAppscodeComAPIGroup()



get information of a group

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.RepositoriesStashAppscodeComApi;


RepositoriesStashAppscodeComApi apiInstance = new RepositoriesStashAppscodeComApi();
try {
    IoK8sApimachineryPkgApisMetaV1APIGroup result = apiInstance.getRepositoriesStashAppscodeComAPIGroup();
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling RepositoriesStashAppscodeComApi#getRepositoriesStashAppscodeComAPIGroup");
    e.printStackTrace();
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**IoK8sApimachineryPkgApisMetaV1APIGroup**](IoK8sApimachineryPkgApisMetaV1APIGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, application/vnd.kubernetes.protobuf
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

