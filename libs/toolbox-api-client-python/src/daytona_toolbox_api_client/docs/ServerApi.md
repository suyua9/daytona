# daytona_toolbox_api_client.ServerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**initialize**](ServerApi.md#initialize) | **POST** /init | Initialize toolbox server


# **initialize**
> Dict[str, str] initialize(request)

Initialize toolbox server

Set the auth token and initialize telemetry for the toolbox server

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.initialize_request import InitializeRequest
from daytona_toolbox_api_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with daytona_toolbox_api_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client.ServerApi(api_client)
    request = daytona_toolbox_api_client.InitializeRequest() # InitializeRequest | Initialization request

    try:
        # Initialize toolbox server
        api_response = api_instance.initialize(request)
        print("The response of ServerApi->initialize:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ServerApi->initialize: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**InitializeRequest**](InitializeRequest.md)| Initialization request | 

### Return type

**Dict[str, str]**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

