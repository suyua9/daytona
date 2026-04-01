# daytona_toolbox_api_client.ComputerUseApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**click**](ComputerUseApi.md#click) | **POST** /computeruse/mouse/click | Click mouse button
[**delete_recording**](ComputerUseApi.md#delete_recording) | **DELETE** /computeruse/recordings/{id} | Delete a recording
[**download_recording**](ComputerUseApi.md#download_recording) | **GET** /computeruse/recordings/{id}/download | Download a recording
[**drag**](ComputerUseApi.md#drag) | **POST** /computeruse/mouse/drag | Drag mouse
[**get_computer_use_status**](ComputerUseApi.md#get_computer_use_status) | **GET** /computeruse/process-status | Get computer use process status
[**get_computer_use_system_status**](ComputerUseApi.md#get_computer_use_system_status) | **GET** /computeruse/status | Get computer use status
[**get_display_info**](ComputerUseApi.md#get_display_info) | **GET** /computeruse/display/info | Get display information
[**get_mouse_position**](ComputerUseApi.md#get_mouse_position) | **GET** /computeruse/mouse/position | Get mouse position
[**get_process_errors**](ComputerUseApi.md#get_process_errors) | **GET** /computeruse/process/{processName}/errors | Get process errors
[**get_process_logs**](ComputerUseApi.md#get_process_logs) | **GET** /computeruse/process/{processName}/logs | Get process logs
[**get_process_status**](ComputerUseApi.md#get_process_status) | **GET** /computeruse/process/{processName}/status | Get specific process status
[**get_recording**](ComputerUseApi.md#get_recording) | **GET** /computeruse/recordings/{id} | Get recording details
[**get_windows**](ComputerUseApi.md#get_windows) | **GET** /computeruse/display/windows | Get windows information
[**list_recordings**](ComputerUseApi.md#list_recordings) | **GET** /computeruse/recordings | List all recordings
[**move_mouse**](ComputerUseApi.md#move_mouse) | **POST** /computeruse/mouse/move | Move mouse cursor
[**press_hotkey**](ComputerUseApi.md#press_hotkey) | **POST** /computeruse/keyboard/hotkey | Press hotkey
[**press_key**](ComputerUseApi.md#press_key) | **POST** /computeruse/keyboard/key | Press key
[**restart_process**](ComputerUseApi.md#restart_process) | **POST** /computeruse/process/{processName}/restart | Restart specific process
[**scroll**](ComputerUseApi.md#scroll) | **POST** /computeruse/mouse/scroll | Scroll mouse wheel
[**start_computer_use**](ComputerUseApi.md#start_computer_use) | **POST** /computeruse/start | Start computer use processes
[**start_recording**](ComputerUseApi.md#start_recording) | **POST** /computeruse/recordings/start | Start a new recording
[**stop_computer_use**](ComputerUseApi.md#stop_computer_use) | **POST** /computeruse/stop | Stop computer use processes
[**stop_recording**](ComputerUseApi.md#stop_recording) | **POST** /computeruse/recordings/stop | Stop a recording
[**take_compressed_region_screenshot**](ComputerUseApi.md#take_compressed_region_screenshot) | **GET** /computeruse/screenshot/region/compressed | Take a compressed region screenshot
[**take_compressed_screenshot**](ComputerUseApi.md#take_compressed_screenshot) | **GET** /computeruse/screenshot/compressed | Take a compressed screenshot
[**take_region_screenshot**](ComputerUseApi.md#take_region_screenshot) | **GET** /computeruse/screenshot/region | Take a region screenshot
[**take_screenshot**](ComputerUseApi.md#take_screenshot) | **GET** /computeruse/screenshot | Take a screenshot
[**type_text**](ComputerUseApi.md#type_text) | **POST** /computeruse/keyboard/type | Type text


# **click**
> MouseClickResponse click(request)

Click mouse button

Click the mouse button at the specified coordinates

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.mouse_click_request import MouseClickRequest
from daytona_toolbox_api_client.models.mouse_click_response import MouseClickResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    request = daytona_toolbox_api_client.MouseClickRequest() # MouseClickRequest | Mouse click request

    try:
        # Click mouse button
        api_response = api_instance.click(request)
        print("The response of ComputerUseApi->click:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->click: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**MouseClickRequest**](MouseClickRequest.md)| Mouse click request | 

### Return type

[**MouseClickResponse**](MouseClickResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_recording**
> delete_recording(id)

Delete a recording

Delete a recording file by ID

### Example


```python
import daytona_toolbox_api_client
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    id = 'id_example' # str | Recording ID

    try:
        # Delete a recording
        api_instance.delete_recording(id)
    except Exception as e:
        print("Exception when calling ComputerUseApi->delete_recording: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| Recording ID | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | No Content |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **download_recording**
> bytearray download_recording(id)

Download a recording

Download a recording by providing its ID

### Example


```python
import daytona_toolbox_api_client
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    id = 'id_example' # str | Recording ID

    try:
        # Download a recording
        api_response = api_instance.download_recording(id)
        print("The response of ComputerUseApi->download_recording:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->download_recording: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| Recording ID | 

### Return type

**bytearray**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **drag**
> MouseDragResponse drag(request)

Drag mouse

Drag the mouse from start to end coordinates

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.mouse_drag_request import MouseDragRequest
from daytona_toolbox_api_client.models.mouse_drag_response import MouseDragResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    request = daytona_toolbox_api_client.MouseDragRequest() # MouseDragRequest | Mouse drag request

    try:
        # Drag mouse
        api_response = api_instance.drag(request)
        print("The response of ComputerUseApi->drag:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->drag: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**MouseDragRequest**](MouseDragRequest.md)| Mouse drag request | 

### Return type

[**MouseDragResponse**](MouseDragResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_computer_use_status**
> ComputerUseStatusResponse get_computer_use_status()

Get computer use process status

Get the status of all computer use processes

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.computer_use_status_response import ComputerUseStatusResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)

    try:
        # Get computer use process status
        api_response = api_instance.get_computer_use_status()
        print("The response of ComputerUseApi->get_computer_use_status:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->get_computer_use_status: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**ComputerUseStatusResponse**](ComputerUseStatusResponse.md)

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

# **get_computer_use_system_status**
> ComputerUseStatusResponse get_computer_use_system_status()

Get computer use status

Get the current status of the computer use system

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.computer_use_status_response import ComputerUseStatusResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)

    try:
        # Get computer use status
        api_response = api_instance.get_computer_use_system_status()
        print("The response of ComputerUseApi->get_computer_use_system_status:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->get_computer_use_system_status: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**ComputerUseStatusResponse**](ComputerUseStatusResponse.md)

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

# **get_display_info**
> DisplayInfoResponse get_display_info()

Get display information

Get information about all available displays

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.display_info_response import DisplayInfoResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)

    try:
        # Get display information
        api_response = api_instance.get_display_info()
        print("The response of ComputerUseApi->get_display_info:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->get_display_info: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**DisplayInfoResponse**](DisplayInfoResponse.md)

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

# **get_mouse_position**
> MousePositionResponse get_mouse_position()

Get mouse position

Get the current mouse cursor position

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.mouse_position_response import MousePositionResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)

    try:
        # Get mouse position
        api_response = api_instance.get_mouse_position()
        print("The response of ComputerUseApi->get_mouse_position:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->get_mouse_position: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**MousePositionResponse**](MousePositionResponse.md)

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

# **get_process_errors**
> ProcessErrorsResponse get_process_errors(process_name)

Get process errors

Get errors for a specific computer use process

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.process_errors_response import ProcessErrorsResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    process_name = 'process_name_example' # str | Process name to get errors for

    try:
        # Get process errors
        api_response = api_instance.get_process_errors(process_name)
        print("The response of ComputerUseApi->get_process_errors:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->get_process_errors: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **process_name** | **str**| Process name to get errors for | 

### Return type

[**ProcessErrorsResponse**](ProcessErrorsResponse.md)

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

# **get_process_logs**
> ProcessLogsResponse get_process_logs(process_name)

Get process logs

Get logs for a specific computer use process

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.process_logs_response import ProcessLogsResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    process_name = 'process_name_example' # str | Process name to get logs for

    try:
        # Get process logs
        api_response = api_instance.get_process_logs(process_name)
        print("The response of ComputerUseApi->get_process_logs:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->get_process_logs: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **process_name** | **str**| Process name to get logs for | 

### Return type

[**ProcessLogsResponse**](ProcessLogsResponse.md)

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

# **get_process_status**
> ProcessStatusResponse get_process_status(process_name)

Get specific process status

Check if a specific computer use process is running

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.process_status_response import ProcessStatusResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    process_name = 'process_name_example' # str | Process name to check

    try:
        # Get specific process status
        api_response = api_instance.get_process_status(process_name)
        print("The response of ComputerUseApi->get_process_status:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->get_process_status: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **process_name** | **str**| Process name to check | 

### Return type

[**ProcessStatusResponse**](ProcessStatusResponse.md)

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

# **get_recording**
> Recording get_recording(id)

Get recording details

Get details of a specific recording by ID

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.recording import Recording
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    id = 'id_example' # str | Recording ID

    try:
        # Get recording details
        api_response = api_instance.get_recording(id)
        print("The response of ComputerUseApi->get_recording:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->get_recording: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| Recording ID | 

### Return type

[**Recording**](Recording.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_windows**
> WindowsResponse get_windows()

Get windows information

Get information about all open windows

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.windows_response import WindowsResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)

    try:
        # Get windows information
        api_response = api_instance.get_windows()
        print("The response of ComputerUseApi->get_windows:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->get_windows: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**WindowsResponse**](WindowsResponse.md)

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

# **list_recordings**
> ListRecordingsResponse list_recordings()

List all recordings

Get a list of all recordings (active and completed)

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.list_recordings_response import ListRecordingsResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)

    try:
        # List all recordings
        api_response = api_instance.list_recordings()
        print("The response of ComputerUseApi->list_recordings:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->list_recordings: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**ListRecordingsResponse**](ListRecordingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **move_mouse**
> MousePositionResponse move_mouse(request)

Move mouse cursor

Move the mouse cursor to the specified coordinates

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.mouse_move_request import MouseMoveRequest
from daytona_toolbox_api_client.models.mouse_position_response import MousePositionResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    request = daytona_toolbox_api_client.MouseMoveRequest() # MouseMoveRequest | Mouse move request

    try:
        # Move mouse cursor
        api_response = api_instance.move_mouse(request)
        print("The response of ComputerUseApi->move_mouse:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->move_mouse: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**MouseMoveRequest**](MouseMoveRequest.md)| Mouse move request | 

### Return type

[**MousePositionResponse**](MousePositionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **press_hotkey**
> object press_hotkey(request)

Press hotkey

Press a hotkey combination (e.g., ctrl+c, cmd+v)

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.keyboard_hotkey_request import KeyboardHotkeyRequest
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    request = daytona_toolbox_api_client.KeyboardHotkeyRequest() # KeyboardHotkeyRequest | Hotkey press request

    try:
        # Press hotkey
        api_response = api_instance.press_hotkey(request)
        print("The response of ComputerUseApi->press_hotkey:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->press_hotkey: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**KeyboardHotkeyRequest**](KeyboardHotkeyRequest.md)| Hotkey press request | 

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **press_key**
> object press_key(request)

Press key

Press a key with optional modifiers

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.keyboard_press_request import KeyboardPressRequest
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    request = daytona_toolbox_api_client.KeyboardPressRequest() # KeyboardPressRequest | Key press request

    try:
        # Press key
        api_response = api_instance.press_key(request)
        print("The response of ComputerUseApi->press_key:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->press_key: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**KeyboardPressRequest**](KeyboardPressRequest.md)| Key press request | 

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **restart_process**
> ProcessRestartResponse restart_process(process_name)

Restart specific process

Restart a specific computer use process

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.process_restart_response import ProcessRestartResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    process_name = 'process_name_example' # str | Process name to restart

    try:
        # Restart specific process
        api_response = api_instance.restart_process(process_name)
        print("The response of ComputerUseApi->restart_process:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->restart_process: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **process_name** | **str**| Process name to restart | 

### Return type

[**ProcessRestartResponse**](ProcessRestartResponse.md)

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

# **scroll**
> ScrollResponse scroll(request)

Scroll mouse wheel

Scroll the mouse wheel at the specified coordinates

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.mouse_scroll_request import MouseScrollRequest
from daytona_toolbox_api_client.models.scroll_response import ScrollResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    request = daytona_toolbox_api_client.MouseScrollRequest() # MouseScrollRequest | Mouse scroll request

    try:
        # Scroll mouse wheel
        api_response = api_instance.scroll(request)
        print("The response of ComputerUseApi->scroll:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->scroll: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**MouseScrollRequest**](MouseScrollRequest.md)| Mouse scroll request | 

### Return type

[**ScrollResponse**](ScrollResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **start_computer_use**
> ComputerUseStartResponse start_computer_use()

Start computer use processes

Start all computer use processes and return their status

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.computer_use_start_response import ComputerUseStartResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)

    try:
        # Start computer use processes
        api_response = api_instance.start_computer_use()
        print("The response of ComputerUseApi->start_computer_use:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->start_computer_use: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**ComputerUseStartResponse**](ComputerUseStartResponse.md)

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

# **start_recording**
> Recording start_recording(request=request)

Start a new recording

Start a new screen recording session

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.recording import Recording
from daytona_toolbox_api_client.models.start_recording_request import StartRecordingRequest
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    request = daytona_toolbox_api_client.StartRecordingRequest() # StartRecordingRequest | Recording options (optional)

    try:
        # Start a new recording
        api_response = api_instance.start_recording(request=request)
        print("The response of ComputerUseApi->start_recording:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->start_recording: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**StartRecordingRequest**](StartRecordingRequest.md)| Recording options | [optional] 

### Return type

[**Recording**](Recording.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Created |  -  |
**400** | Bad Request |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **stop_computer_use**
> ComputerUseStopResponse stop_computer_use()

Stop computer use processes

Stop all computer use processes and return their status

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.computer_use_stop_response import ComputerUseStopResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)

    try:
        # Stop computer use processes
        api_response = api_instance.stop_computer_use()
        print("The response of ComputerUseApi->stop_computer_use:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->stop_computer_use: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**ComputerUseStopResponse**](ComputerUseStopResponse.md)

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

# **stop_recording**
> Recording stop_recording(request)

Stop a recording

Stop an active screen recording session

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.recording import Recording
from daytona_toolbox_api_client.models.stop_recording_request import StopRecordingRequest
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    request = daytona_toolbox_api_client.StopRecordingRequest() # StopRecordingRequest | Recording ID to stop

    try:
        # Stop a recording
        api_response = api_instance.stop_recording(request)
        print("The response of ComputerUseApi->stop_recording:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->stop_recording: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**StopRecordingRequest**](StopRecordingRequest.md)| Recording ID to stop | 

### Return type

[**Recording**](Recording.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **take_compressed_region_screenshot**
> ScreenshotResponse take_compressed_region_screenshot(x, y, width, height, show_cursor=show_cursor, format=format, quality=quality, scale=scale)

Take a compressed region screenshot

Take a compressed screenshot of a specific region of the screen

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.screenshot_response import ScreenshotResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    x = 56 # int | X coordinate of the region
    y = 56 # int | Y coordinate of the region
    width = 56 # int | Width of the region
    height = 56 # int | Height of the region
    show_cursor = True # bool | Whether to show cursor in screenshot (optional)
    format = 'format_example' # str | Image format (png or jpeg) (optional)
    quality = 56 # int | JPEG quality (1-100) (optional)
    scale = 3.4 # float | Scale factor (0.1-1.0) (optional)

    try:
        # Take a compressed region screenshot
        api_response = api_instance.take_compressed_region_screenshot(x, y, width, height, show_cursor=show_cursor, format=format, quality=quality, scale=scale)
        print("The response of ComputerUseApi->take_compressed_region_screenshot:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->take_compressed_region_screenshot: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x** | **int**| X coordinate of the region | 
 **y** | **int**| Y coordinate of the region | 
 **width** | **int**| Width of the region | 
 **height** | **int**| Height of the region | 
 **show_cursor** | **bool**| Whether to show cursor in screenshot | [optional] 
 **format** | **str**| Image format (png or jpeg) | [optional] 
 **quality** | **int**| JPEG quality (1-100) | [optional] 
 **scale** | **float**| Scale factor (0.1-1.0) | [optional] 

### Return type

[**ScreenshotResponse**](ScreenshotResponse.md)

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

# **take_compressed_screenshot**
> ScreenshotResponse take_compressed_screenshot(show_cursor=show_cursor, format=format, quality=quality, scale=scale)

Take a compressed screenshot

Take a compressed screenshot of the entire screen

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.screenshot_response import ScreenshotResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    show_cursor = True # bool | Whether to show cursor in screenshot (optional)
    format = 'format_example' # str | Image format (png or jpeg) (optional)
    quality = 56 # int | JPEG quality (1-100) (optional)
    scale = 3.4 # float | Scale factor (0.1-1.0) (optional)

    try:
        # Take a compressed screenshot
        api_response = api_instance.take_compressed_screenshot(show_cursor=show_cursor, format=format, quality=quality, scale=scale)
        print("The response of ComputerUseApi->take_compressed_screenshot:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->take_compressed_screenshot: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **show_cursor** | **bool**| Whether to show cursor in screenshot | [optional] 
 **format** | **str**| Image format (png or jpeg) | [optional] 
 **quality** | **int**| JPEG quality (1-100) | [optional] 
 **scale** | **float**| Scale factor (0.1-1.0) | [optional] 

### Return type

[**ScreenshotResponse**](ScreenshotResponse.md)

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

# **take_region_screenshot**
> ScreenshotResponse take_region_screenshot(x, y, width, height, show_cursor=show_cursor)

Take a region screenshot

Take a screenshot of a specific region of the screen

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.screenshot_response import ScreenshotResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    x = 56 # int | X coordinate of the region
    y = 56 # int | Y coordinate of the region
    width = 56 # int | Width of the region
    height = 56 # int | Height of the region
    show_cursor = True # bool | Whether to show cursor in screenshot (optional)

    try:
        # Take a region screenshot
        api_response = api_instance.take_region_screenshot(x, y, width, height, show_cursor=show_cursor)
        print("The response of ComputerUseApi->take_region_screenshot:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->take_region_screenshot: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x** | **int**| X coordinate of the region | 
 **y** | **int**| Y coordinate of the region | 
 **width** | **int**| Width of the region | 
 **height** | **int**| Height of the region | 
 **show_cursor** | **bool**| Whether to show cursor in screenshot | [optional] 

### Return type

[**ScreenshotResponse**](ScreenshotResponse.md)

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

# **take_screenshot**
> ScreenshotResponse take_screenshot(show_cursor=show_cursor)

Take a screenshot

Take a screenshot of the entire screen

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.screenshot_response import ScreenshotResponse
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    show_cursor = True # bool | Whether to show cursor in screenshot (optional)

    try:
        # Take a screenshot
        api_response = api_instance.take_screenshot(show_cursor=show_cursor)
        print("The response of ComputerUseApi->take_screenshot:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->take_screenshot: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **show_cursor** | **bool**| Whether to show cursor in screenshot | [optional] 

### Return type

[**ScreenshotResponse**](ScreenshotResponse.md)

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

# **type_text**
> object type_text(request)

Type text

Type text with optional delay between keystrokes

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.keyboard_type_request import KeyboardTypeRequest
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
    api_instance = daytona_toolbox_api_client.ComputerUseApi(api_client)
    request = daytona_toolbox_api_client.KeyboardTypeRequest() # KeyboardTypeRequest | Text typing request

    try:
        # Type text
        api_response = api_instance.type_text(request)
        print("The response of ComputerUseApi->type_text:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ComputerUseApi->type_text: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**KeyboardTypeRequest**](KeyboardTypeRequest.md)| Text typing request | 

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

