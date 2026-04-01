# daytona_toolbox_api_client.PortApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_ports**](PortApi.md#get_ports) | **GET** /port | Get active ports
[**is_port_in_use**](PortApi.md#is_port_in_use) | **GET** /port/{port}/in-use | Check if port is in use


# **get_ports**
> PortList get_ports()

Get active ports

Get a list of all currently active ports

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.port_list import PortList
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
    api_instance = daytona_toolbox_api_client.PortApi(api_client)

    try:
        # Get active ports
        api_response = api_instance.get_ports()
        print("The response of PortApi->get_ports:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PortApi->get_ports: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**PortList**](PortList.md)

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

# **is_port_in_use**
> IsPortInUseResponse is_port_in_use(port)

Check if port is in use

Check if a specific port is currently in use

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.is_port_in_use_response import IsPortInUseResponse
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
    api_instance = daytona_toolbox_api_client.PortApi(api_client)
    port = 56 # int | Port number (3000-9999)

    try:
        # Check if port is in use
        api_response = api_instance.is_port_in_use(port)
        print("The response of PortApi->is_port_in_use:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PortApi->is_port_in_use: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **port** | **int**| Port number (3000-9999) | 

### Return type

[**IsPortInUseResponse**](IsPortInUseResponse.md)

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

