# swagger_client.RepositoriesStashAppscodeComApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_repositories_stash_appscode_com_api_group**](RepositoriesStashAppscodeComApi.md#get_repositories_stash_appscode_com_api_group) | **GET** /apis/repositories.stash.appscode.com/ | 


# **get_repositories_stash_appscode_com_api_group**
> IoK8sApimachineryPkgApisMetaV1APIGroup get_repositories_stash_appscode_com_api_group()



get information of a group

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.RepositoriesStashAppscodeComApi()

try:
    api_response = api_instance.get_repositories_stash_appscode_com_api_group()
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RepositoriesStashAppscodeComApi->get_repositories_stash_appscode_com_api_group: %s\n" % e)
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

