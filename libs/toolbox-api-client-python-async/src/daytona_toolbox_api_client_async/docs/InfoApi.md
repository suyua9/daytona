# daytona_toolbox_api_client_async.InfoApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_user_home_dir**](InfoApi.md#get_user_home_dir) | **GET** /user-home-dir | Get user home directory
[**get_version**](InfoApi.md#get_version) | **GET** /version | Get version
[**get_work_dir**](InfoApi.md#get_work_dir) | **GET** /work-dir | Get working directory


# **get_user_home_dir**
> UserHomeDirResponse get_user_home_dir()

Get user home directory

Get the current user home directory path.

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.user_home_dir_response import UserHomeDirResponse
from daytona_toolbox_api_client_async.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client_async.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
async with daytona_toolbox_api_client_async.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client_async.InfoApi(api_client)

    try:
        # Get user home directory
        api_response = await api_instance.get_user_home_dir()
        print("The response of InfoApi->get_user_home_dir:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InfoApi->get_user_home_dir: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**UserHomeDirResponse**](UserHomeDirResponse.md)

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

# **get_version**
> Dict[str, str] get_version()

Get version

Get the current daemon version

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client_async.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
async with daytona_toolbox_api_client_async.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client_async.InfoApi(api_client)

    try:
        # Get version
        api_response = await api_instance.get_version()
        print("The response of InfoApi->get_version:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InfoApi->get_version: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

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

# **get_work_dir**
> WorkDirResponse get_work_dir()

Get working directory

Get the current working directory path. This is default directory used for running commands.

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.work_dir_response import WorkDirResponse
from daytona_toolbox_api_client_async.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client_async.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
async with daytona_toolbox_api_client_async.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client_async.InfoApi(api_client)

    try:
        # Get working directory
        api_response = await api_instance.get_work_dir()
        print("The response of InfoApi->get_work_dir:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InfoApi->get_work_dir: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**WorkDirResponse**](WorkDirResponse.md)

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

