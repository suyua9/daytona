# DaytonaToolboxApiClient::ComputerUseApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**click**](ComputerUseApi.md#click) | **POST** /computeruse/mouse/click | Click mouse button |
| [**delete_recording**](ComputerUseApi.md#delete_recording) | **DELETE** /computeruse/recordings/{id} | Delete a recording |
| [**download_recording**](ComputerUseApi.md#download_recording) | **GET** /computeruse/recordings/{id}/download | Download a recording |
| [**drag**](ComputerUseApi.md#drag) | **POST** /computeruse/mouse/drag | Drag mouse |
| [**get_computer_use_status**](ComputerUseApi.md#get_computer_use_status) | **GET** /computeruse/process-status | Get computer use process status |
| [**get_computer_use_system_status**](ComputerUseApi.md#get_computer_use_system_status) | **GET** /computeruse/status | Get computer use status |
| [**get_display_info**](ComputerUseApi.md#get_display_info) | **GET** /computeruse/display/info | Get display information |
| [**get_mouse_position**](ComputerUseApi.md#get_mouse_position) | **GET** /computeruse/mouse/position | Get mouse position |
| [**get_process_errors**](ComputerUseApi.md#get_process_errors) | **GET** /computeruse/process/{processName}/errors | Get process errors |
| [**get_process_logs**](ComputerUseApi.md#get_process_logs) | **GET** /computeruse/process/{processName}/logs | Get process logs |
| [**get_process_status**](ComputerUseApi.md#get_process_status) | **GET** /computeruse/process/{processName}/status | Get specific process status |
| [**get_recording**](ComputerUseApi.md#get_recording) | **GET** /computeruse/recordings/{id} | Get recording details |
| [**get_windows**](ComputerUseApi.md#get_windows) | **GET** /computeruse/display/windows | Get windows information |
| [**list_recordings**](ComputerUseApi.md#list_recordings) | **GET** /computeruse/recordings | List all recordings |
| [**move_mouse**](ComputerUseApi.md#move_mouse) | **POST** /computeruse/mouse/move | Move mouse cursor |
| [**press_hotkey**](ComputerUseApi.md#press_hotkey) | **POST** /computeruse/keyboard/hotkey | Press hotkey |
| [**press_key**](ComputerUseApi.md#press_key) | **POST** /computeruse/keyboard/key | Press key |
| [**restart_process**](ComputerUseApi.md#restart_process) | **POST** /computeruse/process/{processName}/restart | Restart specific process |
| [**scroll**](ComputerUseApi.md#scroll) | **POST** /computeruse/mouse/scroll | Scroll mouse wheel |
| [**start_computer_use**](ComputerUseApi.md#start_computer_use) | **POST** /computeruse/start | Start computer use processes |
| [**start_recording**](ComputerUseApi.md#start_recording) | **POST** /computeruse/recordings/start | Start a new recording |
| [**stop_computer_use**](ComputerUseApi.md#stop_computer_use) | **POST** /computeruse/stop | Stop computer use processes |
| [**stop_recording**](ComputerUseApi.md#stop_recording) | **POST** /computeruse/recordings/stop | Stop a recording |
| [**take_compressed_region_screenshot**](ComputerUseApi.md#take_compressed_region_screenshot) | **GET** /computeruse/screenshot/region/compressed | Take a compressed region screenshot |
| [**take_compressed_screenshot**](ComputerUseApi.md#take_compressed_screenshot) | **GET** /computeruse/screenshot/compressed | Take a compressed screenshot |
| [**take_region_screenshot**](ComputerUseApi.md#take_region_screenshot) | **GET** /computeruse/screenshot/region | Take a region screenshot |
| [**take_screenshot**](ComputerUseApi.md#take_screenshot) | **GET** /computeruse/screenshot | Take a screenshot |
| [**type_text**](ComputerUseApi.md#type_text) | **POST** /computeruse/keyboard/type | Type text |


## click

> <MouseClickResponse> click(request)

Click mouse button

Click the mouse button at the specified coordinates

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
request = DaytonaToolboxApiClient::MouseClickRequest.new # MouseClickRequest | Mouse click request

begin
  # Click mouse button
  result = api_instance.click(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->click: #{e}"
end
```

#### Using the click_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<MouseClickResponse>, Integer, Hash)> click_with_http_info(request)

```ruby
begin
  # Click mouse button
  data, status_code, headers = api_instance.click_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <MouseClickResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->click_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**MouseClickRequest**](MouseClickRequest.md) | Mouse click request |  |

### Return type

[**MouseClickResponse**](MouseClickResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## delete_recording

> delete_recording(id)

Delete a recording

Delete a recording file by ID

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
id = 'id_example' # String | Recording ID

begin
  # Delete a recording
  api_instance.delete_recording(id)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->delete_recording: #{e}"
end
```

#### Using the delete_recording_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> delete_recording_with_http_info(id)

```ruby
begin
  # Delete a recording
  data, status_code, headers = api_instance.delete_recording_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->delete_recording_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** | Recording ID |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*


## download_recording

> File download_recording(id)

Download a recording

Download a recording by providing its ID

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
id = 'id_example' # String | Recording ID

begin
  # Download a recording
  result = api_instance.download_recording(id)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->download_recording: #{e}"
end
```

#### Using the download_recording_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(File, Integer, Hash)> download_recording_with_http_info(id)

```ruby
begin
  # Download a recording
  data, status_code, headers = api_instance.download_recording_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => File
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->download_recording_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** | Recording ID |  |

### Return type

**File**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream


## drag

> <MouseDragResponse> drag(request)

Drag mouse

Drag the mouse from start to end coordinates

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
request = DaytonaToolboxApiClient::MouseDragRequest.new # MouseDragRequest | Mouse drag request

begin
  # Drag mouse
  result = api_instance.drag(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->drag: #{e}"
end
```

#### Using the drag_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<MouseDragResponse>, Integer, Hash)> drag_with_http_info(request)

```ruby
begin
  # Drag mouse
  data, status_code, headers = api_instance.drag_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <MouseDragResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->drag_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**MouseDragRequest**](MouseDragRequest.md) | Mouse drag request |  |

### Return type

[**MouseDragResponse**](MouseDragResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## get_computer_use_status

> <ComputerUseStatusResponse> get_computer_use_status

Get computer use process status

Get the status of all computer use processes

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new

begin
  # Get computer use process status
  result = api_instance.get_computer_use_status
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_computer_use_status: #{e}"
end
```

#### Using the get_computer_use_status_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ComputerUseStatusResponse>, Integer, Hash)> get_computer_use_status_with_http_info

```ruby
begin
  # Get computer use process status
  data, status_code, headers = api_instance.get_computer_use_status_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ComputerUseStatusResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_computer_use_status_with_http_info: #{e}"
end
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


## get_computer_use_system_status

> <ComputerUseStatusResponse> get_computer_use_system_status

Get computer use status

Get the current status of the computer use system

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new

begin
  # Get computer use status
  result = api_instance.get_computer_use_system_status
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_computer_use_system_status: #{e}"
end
```

#### Using the get_computer_use_system_status_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ComputerUseStatusResponse>, Integer, Hash)> get_computer_use_system_status_with_http_info

```ruby
begin
  # Get computer use status
  data, status_code, headers = api_instance.get_computer_use_system_status_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ComputerUseStatusResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_computer_use_system_status_with_http_info: #{e}"
end
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


## get_display_info

> <DisplayInfoResponse> get_display_info

Get display information

Get information about all available displays

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new

begin
  # Get display information
  result = api_instance.get_display_info
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_display_info: #{e}"
end
```

#### Using the get_display_info_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DisplayInfoResponse>, Integer, Hash)> get_display_info_with_http_info

```ruby
begin
  # Get display information
  data, status_code, headers = api_instance.get_display_info_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DisplayInfoResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_display_info_with_http_info: #{e}"
end
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


## get_mouse_position

> <MousePositionResponse> get_mouse_position

Get mouse position

Get the current mouse cursor position

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new

begin
  # Get mouse position
  result = api_instance.get_mouse_position
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_mouse_position: #{e}"
end
```

#### Using the get_mouse_position_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<MousePositionResponse>, Integer, Hash)> get_mouse_position_with_http_info

```ruby
begin
  # Get mouse position
  data, status_code, headers = api_instance.get_mouse_position_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <MousePositionResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_mouse_position_with_http_info: #{e}"
end
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


## get_process_errors

> <ProcessErrorsResponse> get_process_errors(process_name)

Get process errors

Get errors for a specific computer use process

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
process_name = 'process_name_example' # String | Process name to get errors for

begin
  # Get process errors
  result = api_instance.get_process_errors(process_name)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_process_errors: #{e}"
end
```

#### Using the get_process_errors_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ProcessErrorsResponse>, Integer, Hash)> get_process_errors_with_http_info(process_name)

```ruby
begin
  # Get process errors
  data, status_code, headers = api_instance.get_process_errors_with_http_info(process_name)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ProcessErrorsResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_process_errors_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **process_name** | **String** | Process name to get errors for |  |

### Return type

[**ProcessErrorsResponse**](ProcessErrorsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_process_logs

> <ProcessLogsResponse> get_process_logs(process_name)

Get process logs

Get logs for a specific computer use process

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
process_name = 'process_name_example' # String | Process name to get logs for

begin
  # Get process logs
  result = api_instance.get_process_logs(process_name)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_process_logs: #{e}"
end
```

#### Using the get_process_logs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ProcessLogsResponse>, Integer, Hash)> get_process_logs_with_http_info(process_name)

```ruby
begin
  # Get process logs
  data, status_code, headers = api_instance.get_process_logs_with_http_info(process_name)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ProcessLogsResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_process_logs_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **process_name** | **String** | Process name to get logs for |  |

### Return type

[**ProcessLogsResponse**](ProcessLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_process_status

> <ProcessStatusResponse> get_process_status(process_name)

Get specific process status

Check if a specific computer use process is running

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
process_name = 'process_name_example' # String | Process name to check

begin
  # Get specific process status
  result = api_instance.get_process_status(process_name)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_process_status: #{e}"
end
```

#### Using the get_process_status_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ProcessStatusResponse>, Integer, Hash)> get_process_status_with_http_info(process_name)

```ruby
begin
  # Get specific process status
  data, status_code, headers = api_instance.get_process_status_with_http_info(process_name)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ProcessStatusResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_process_status_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **process_name** | **String** | Process name to check |  |

### Return type

[**ProcessStatusResponse**](ProcessStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_recording

> <Recording> get_recording(id)

Get recording details

Get details of a specific recording by ID

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
id = 'id_example' # String | Recording ID

begin
  # Get recording details
  result = api_instance.get_recording(id)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_recording: #{e}"
end
```

#### Using the get_recording_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Recording>, Integer, Hash)> get_recording_with_http_info(id)

```ruby
begin
  # Get recording details
  data, status_code, headers = api_instance.get_recording_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Recording>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_recording_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** | Recording ID |  |

### Return type

[**Recording**](Recording.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_windows

> <WindowsResponse> get_windows

Get windows information

Get information about all open windows

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new

begin
  # Get windows information
  result = api_instance.get_windows
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_windows: #{e}"
end
```

#### Using the get_windows_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<WindowsResponse>, Integer, Hash)> get_windows_with_http_info

```ruby
begin
  # Get windows information
  data, status_code, headers = api_instance.get_windows_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <WindowsResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->get_windows_with_http_info: #{e}"
end
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


## list_recordings

> <ListRecordingsResponse> list_recordings

List all recordings

Get a list of all recordings (active and completed)

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new

begin
  # List all recordings
  result = api_instance.list_recordings
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->list_recordings: #{e}"
end
```

#### Using the list_recordings_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ListRecordingsResponse>, Integer, Hash)> list_recordings_with_http_info

```ruby
begin
  # List all recordings
  data, status_code, headers = api_instance.list_recordings_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ListRecordingsResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->list_recordings_with_http_info: #{e}"
end
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


## move_mouse

> <MousePositionResponse> move_mouse(request)

Move mouse cursor

Move the mouse cursor to the specified coordinates

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
request = DaytonaToolboxApiClient::MouseMoveRequest.new # MouseMoveRequest | Mouse move request

begin
  # Move mouse cursor
  result = api_instance.move_mouse(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->move_mouse: #{e}"
end
```

#### Using the move_mouse_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<MousePositionResponse>, Integer, Hash)> move_mouse_with_http_info(request)

```ruby
begin
  # Move mouse cursor
  data, status_code, headers = api_instance.move_mouse_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <MousePositionResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->move_mouse_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**MouseMoveRequest**](MouseMoveRequest.md) | Mouse move request |  |

### Return type

[**MousePositionResponse**](MousePositionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## press_hotkey

> Object press_hotkey(request)

Press hotkey

Press a hotkey combination (e.g., ctrl+c, cmd+v)

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
request = DaytonaToolboxApiClient::KeyboardHotkeyRequest.new # KeyboardHotkeyRequest | Hotkey press request

begin
  # Press hotkey
  result = api_instance.press_hotkey(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->press_hotkey: #{e}"
end
```

#### Using the press_hotkey_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(Object, Integer, Hash)> press_hotkey_with_http_info(request)

```ruby
begin
  # Press hotkey
  data, status_code, headers = api_instance.press_hotkey_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Object
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->press_hotkey_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**KeyboardHotkeyRequest**](KeyboardHotkeyRequest.md) | Hotkey press request |  |

### Return type

**Object**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## press_key

> Object press_key(request)

Press key

Press a key with optional modifiers

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
request = DaytonaToolboxApiClient::KeyboardPressRequest.new # KeyboardPressRequest | Key press request

begin
  # Press key
  result = api_instance.press_key(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->press_key: #{e}"
end
```

#### Using the press_key_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(Object, Integer, Hash)> press_key_with_http_info(request)

```ruby
begin
  # Press key
  data, status_code, headers = api_instance.press_key_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Object
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->press_key_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**KeyboardPressRequest**](KeyboardPressRequest.md) | Key press request |  |

### Return type

**Object**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## restart_process

> <ProcessRestartResponse> restart_process(process_name)

Restart specific process

Restart a specific computer use process

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
process_name = 'process_name_example' # String | Process name to restart

begin
  # Restart specific process
  result = api_instance.restart_process(process_name)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->restart_process: #{e}"
end
```

#### Using the restart_process_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ProcessRestartResponse>, Integer, Hash)> restart_process_with_http_info(process_name)

```ruby
begin
  # Restart specific process
  data, status_code, headers = api_instance.restart_process_with_http_info(process_name)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ProcessRestartResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->restart_process_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **process_name** | **String** | Process name to restart |  |

### Return type

[**ProcessRestartResponse**](ProcessRestartResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## scroll

> <ScrollResponse> scroll(request)

Scroll mouse wheel

Scroll the mouse wheel at the specified coordinates

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
request = DaytonaToolboxApiClient::MouseScrollRequest.new # MouseScrollRequest | Mouse scroll request

begin
  # Scroll mouse wheel
  result = api_instance.scroll(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->scroll: #{e}"
end
```

#### Using the scroll_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ScrollResponse>, Integer, Hash)> scroll_with_http_info(request)

```ruby
begin
  # Scroll mouse wheel
  data, status_code, headers = api_instance.scroll_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ScrollResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->scroll_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**MouseScrollRequest**](MouseScrollRequest.md) | Mouse scroll request |  |

### Return type

[**ScrollResponse**](ScrollResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## start_computer_use

> <ComputerUseStartResponse> start_computer_use

Start computer use processes

Start all computer use processes and return their status

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new

begin
  # Start computer use processes
  result = api_instance.start_computer_use
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->start_computer_use: #{e}"
end
```

#### Using the start_computer_use_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ComputerUseStartResponse>, Integer, Hash)> start_computer_use_with_http_info

```ruby
begin
  # Start computer use processes
  data, status_code, headers = api_instance.start_computer_use_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ComputerUseStartResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->start_computer_use_with_http_info: #{e}"
end
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


## start_recording

> <Recording> start_recording(opts)

Start a new recording

Start a new screen recording session

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
opts = {
  request: DaytonaToolboxApiClient::StartRecordingRequest.new # StartRecordingRequest | Recording options
}

begin
  # Start a new recording
  result = api_instance.start_recording(opts)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->start_recording: #{e}"
end
```

#### Using the start_recording_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Recording>, Integer, Hash)> start_recording_with_http_info(opts)

```ruby
begin
  # Start a new recording
  data, status_code, headers = api_instance.start_recording_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Recording>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->start_recording_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**StartRecordingRequest**](StartRecordingRequest.md) | Recording options | [optional] |

### Return type

[**Recording**](Recording.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## stop_computer_use

> <ComputerUseStopResponse> stop_computer_use

Stop computer use processes

Stop all computer use processes and return their status

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new

begin
  # Stop computer use processes
  result = api_instance.stop_computer_use
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->stop_computer_use: #{e}"
end
```

#### Using the stop_computer_use_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ComputerUseStopResponse>, Integer, Hash)> stop_computer_use_with_http_info

```ruby
begin
  # Stop computer use processes
  data, status_code, headers = api_instance.stop_computer_use_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ComputerUseStopResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->stop_computer_use_with_http_info: #{e}"
end
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


## stop_recording

> <Recording> stop_recording(request)

Stop a recording

Stop an active screen recording session

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
request = DaytonaToolboxApiClient::StopRecordingRequest.new({id: 'id_example'}) # StopRecordingRequest | Recording ID to stop

begin
  # Stop a recording
  result = api_instance.stop_recording(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->stop_recording: #{e}"
end
```

#### Using the stop_recording_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Recording>, Integer, Hash)> stop_recording_with_http_info(request)

```ruby
begin
  # Stop a recording
  data, status_code, headers = api_instance.stop_recording_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Recording>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->stop_recording_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**StopRecordingRequest**](StopRecordingRequest.md) | Recording ID to stop |  |

### Return type

[**Recording**](Recording.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## take_compressed_region_screenshot

> <ScreenshotResponse> take_compressed_region_screenshot(x, y, width, height, opts)

Take a compressed region screenshot

Take a compressed screenshot of a specific region of the screen

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
x = 56 # Integer | X coordinate of the region
y = 56 # Integer | Y coordinate of the region
width = 56 # Integer | Width of the region
height = 56 # Integer | Height of the region
opts = {
  show_cursor: true, # Boolean | Whether to show cursor in screenshot
  format: 'format_example', # String | Image format (png or jpeg)
  quality: 56, # Integer | JPEG quality (1-100)
  scale: 8.14 # Float | Scale factor (0.1-1.0)
}

begin
  # Take a compressed region screenshot
  result = api_instance.take_compressed_region_screenshot(x, y, width, height, opts)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->take_compressed_region_screenshot: #{e}"
end
```

#### Using the take_compressed_region_screenshot_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ScreenshotResponse>, Integer, Hash)> take_compressed_region_screenshot_with_http_info(x, y, width, height, opts)

```ruby
begin
  # Take a compressed region screenshot
  data, status_code, headers = api_instance.take_compressed_region_screenshot_with_http_info(x, y, width, height, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ScreenshotResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->take_compressed_region_screenshot_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **x** | **Integer** | X coordinate of the region |  |
| **y** | **Integer** | Y coordinate of the region |  |
| **width** | **Integer** | Width of the region |  |
| **height** | **Integer** | Height of the region |  |
| **show_cursor** | **Boolean** | Whether to show cursor in screenshot | [optional] |
| **format** | **String** | Image format (png or jpeg) | [optional] |
| **quality** | **Integer** | JPEG quality (1-100) | [optional] |
| **scale** | **Float** | Scale factor (0.1-1.0) | [optional] |

### Return type

[**ScreenshotResponse**](ScreenshotResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## take_compressed_screenshot

> <ScreenshotResponse> take_compressed_screenshot(opts)

Take a compressed screenshot

Take a compressed screenshot of the entire screen

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
opts = {
  show_cursor: true, # Boolean | Whether to show cursor in screenshot
  format: 'format_example', # String | Image format (png or jpeg)
  quality: 56, # Integer | JPEG quality (1-100)
  scale: 8.14 # Float | Scale factor (0.1-1.0)
}

begin
  # Take a compressed screenshot
  result = api_instance.take_compressed_screenshot(opts)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->take_compressed_screenshot: #{e}"
end
```

#### Using the take_compressed_screenshot_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ScreenshotResponse>, Integer, Hash)> take_compressed_screenshot_with_http_info(opts)

```ruby
begin
  # Take a compressed screenshot
  data, status_code, headers = api_instance.take_compressed_screenshot_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ScreenshotResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->take_compressed_screenshot_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **show_cursor** | **Boolean** | Whether to show cursor in screenshot | [optional] |
| **format** | **String** | Image format (png or jpeg) | [optional] |
| **quality** | **Integer** | JPEG quality (1-100) | [optional] |
| **scale** | **Float** | Scale factor (0.1-1.0) | [optional] |

### Return type

[**ScreenshotResponse**](ScreenshotResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## take_region_screenshot

> <ScreenshotResponse> take_region_screenshot(x, y, width, height, opts)

Take a region screenshot

Take a screenshot of a specific region of the screen

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
x = 56 # Integer | X coordinate of the region
y = 56 # Integer | Y coordinate of the region
width = 56 # Integer | Width of the region
height = 56 # Integer | Height of the region
opts = {
  show_cursor: true # Boolean | Whether to show cursor in screenshot
}

begin
  # Take a region screenshot
  result = api_instance.take_region_screenshot(x, y, width, height, opts)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->take_region_screenshot: #{e}"
end
```

#### Using the take_region_screenshot_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ScreenshotResponse>, Integer, Hash)> take_region_screenshot_with_http_info(x, y, width, height, opts)

```ruby
begin
  # Take a region screenshot
  data, status_code, headers = api_instance.take_region_screenshot_with_http_info(x, y, width, height, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ScreenshotResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->take_region_screenshot_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **x** | **Integer** | X coordinate of the region |  |
| **y** | **Integer** | Y coordinate of the region |  |
| **width** | **Integer** | Width of the region |  |
| **height** | **Integer** | Height of the region |  |
| **show_cursor** | **Boolean** | Whether to show cursor in screenshot | [optional] |

### Return type

[**ScreenshotResponse**](ScreenshotResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## take_screenshot

> <ScreenshotResponse> take_screenshot(opts)

Take a screenshot

Take a screenshot of the entire screen

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
opts = {
  show_cursor: true # Boolean | Whether to show cursor in screenshot
}

begin
  # Take a screenshot
  result = api_instance.take_screenshot(opts)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->take_screenshot: #{e}"
end
```

#### Using the take_screenshot_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ScreenshotResponse>, Integer, Hash)> take_screenshot_with_http_info(opts)

```ruby
begin
  # Take a screenshot
  data, status_code, headers = api_instance.take_screenshot_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ScreenshotResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->take_screenshot_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **show_cursor** | **Boolean** | Whether to show cursor in screenshot | [optional] |

### Return type

[**ScreenshotResponse**](ScreenshotResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## type_text

> Object type_text(request)

Type text

Type text with optional delay between keystrokes

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ComputerUseApi.new
request = DaytonaToolboxApiClient::KeyboardTypeRequest.new # KeyboardTypeRequest | Text typing request

begin
  # Type text
  result = api_instance.type_text(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->type_text: #{e}"
end
```

#### Using the type_text_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(Object, Integer, Hash)> type_text_with_http_info(request)

```ruby
begin
  # Type text
  data, status_code, headers = api_instance.type_text_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Object
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ComputerUseApi->type_text_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**KeyboardTypeRequest**](KeyboardTypeRequest.md) | Text typing request |  |

### Return type

**Object**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

