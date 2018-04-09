# StashAppscodeComApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getStashAppscodeComAPIGroup**](StashAppscodeComApi.md#getStashAppscodeComAPIGroup) | **GET** /apis/stash.appscode.com/ | 


<a name="getStashAppscodeComAPIGroup"></a>
# **getStashAppscodeComAPIGroup**
> IoK8sApimachineryPkgApisMetaV1APIGroup getStashAppscodeComAPIGroup()



get information of a group

### Example
```java
// Import classes:
//import io.swagger.client.ApiException;
//import io.swagger.client.api.StashAppscodeComApi;


StashAppscodeComApi apiInstance = new StashAppscodeComApi();
try {
    IoK8sApimachineryPkgApisMetaV1APIGroup result = apiInstance.getStashAppscodeComAPIGroup();
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling StashAppscodeComApi#getStashAppscodeComAPIGroup");
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

