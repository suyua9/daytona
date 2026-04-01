# DaytonaToolboxApiClient::ProcessApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**code_run**](ProcessApi.md#code_run) | **POST** /process/code-run | Execute code |
| [**connect_pty_session**](ProcessApi.md#connect_pty_session) | **GET** /process/pty/{sessionId}/connect | Connect to PTY session via WebSocket |
| [**create_pty_session**](ProcessApi.md#create_pty_session) | **POST** /process/pty | Create a new PTY session |
| [**create_session**](ProcessApi.md#create_session) | **POST** /process/session | Create a new session |
| [**delete_pty_session**](ProcessApi.md#delete_pty_session) | **DELETE** /process/pty/{sessionId} | Delete a PTY session |
| [**delete_session**](ProcessApi.md#delete_session) | **DELETE** /process/session/{sessionId} | Delete a session |
| [**execute_command**](ProcessApi.md#execute_command) | **POST** /process/execute | Execute a command |
| [**get_entrypoint_logs**](ProcessApi.md#get_entrypoint_logs) | **GET** /process/session/entrypoint/logs | Get entrypoint logs |
| [**get_entrypoint_session**](ProcessApi.md#get_entrypoint_session) | **GET** /process/session/entrypoint | Get entrypoint session details |
| [**get_pty_session**](ProcessApi.md#get_pty_session) | **GET** /process/pty/{sessionId} | Get PTY session information |
| [**get_session**](ProcessApi.md#get_session) | **GET** /process/session/{sessionId} | Get session details |
| [**get_session_command**](ProcessApi.md#get_session_command) | **GET** /process/session/{sessionId}/command/{commandId} | Get session command details |
| [**get_session_command_logs**](ProcessApi.md#get_session_command_logs) | **GET** /process/session/{sessionId}/command/{commandId}/logs | Get session command logs |
| [**list_pty_sessions**](ProcessApi.md#list_pty_sessions) | **GET** /process/pty | List all PTY sessions |
| [**list_sessions**](ProcessApi.md#list_sessions) | **GET** /process/session | List all sessions |
| [**resize_pty_session**](ProcessApi.md#resize_pty_session) | **POST** /process/pty/{sessionId}/resize | Resize a PTY session |
| [**send_input**](ProcessApi.md#send_input) | **POST** /process/session/{sessionId}/command/{commandId}/input | Send input to command |
| [**session_execute_command**](ProcessApi.md#session_execute_command) | **POST** /process/session/{sessionId}/exec | Execute command in session |


## code_run

> <CodeRunResponse> code_run(request)

Execute code

Execute Python, JavaScript, or TypeScript code and return output, exit code, and artifacts

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
request = DaytonaToolboxApiClient::CodeRunRequest.new({code: 'code_example', language: 'language_example'}) # CodeRunRequest | Code execution request

begin
  # Execute code
  result = api_instance.code_run(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->code_run: #{e}"
end
```

#### Using the code_run_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CodeRunResponse>, Integer, Hash)> code_run_with_http_info(request)

```ruby
begin
  # Execute code
  data, status_code, headers = api_instance.code_run_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CodeRunResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->code_run_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**CodeRunRequest**](CodeRunRequest.md) | Code execution request |  |

### Return type

[**CodeRunResponse**](CodeRunResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## connect_pty_session

> connect_pty_session(session_id)

Connect to PTY session via WebSocket

Establish a WebSocket connection to interact with a pseudo-terminal session

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
session_id = 'session_id_example' # String | PTY session ID

begin
  # Connect to PTY session via WebSocket
  api_instance.connect_pty_session(session_id)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->connect_pty_session: #{e}"
end
```

#### Using the connect_pty_session_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> connect_pty_session_with_http_info(session_id)

```ruby
begin
  # Connect to PTY session via WebSocket
  data, status_code, headers = api_instance.connect_pty_session_with_http_info(session_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->connect_pty_session_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **session_id** | **String** | PTY session ID |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## create_pty_session

> <PtyCreateResponse> create_pty_session(request)

Create a new PTY session

Create a new pseudo-terminal session with specified configuration

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
request = DaytonaToolboxApiClient::PtyCreateRequest.new # PtyCreateRequest | PTY session creation request

begin
  # Create a new PTY session
  result = api_instance.create_pty_session(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->create_pty_session: #{e}"
end
```

#### Using the create_pty_session_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<PtyCreateResponse>, Integer, Hash)> create_pty_session_with_http_info(request)

```ruby
begin
  # Create a new PTY session
  data, status_code, headers = api_instance.create_pty_session_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <PtyCreateResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->create_pty_session_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**PtyCreateRequest**](PtyCreateRequest.md) | PTY session creation request |  |

### Return type

[**PtyCreateResponse**](PtyCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## create_session

> create_session(request)

Create a new session

Create a new shell session for command execution

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
request = DaytonaToolboxApiClient::CreateSessionRequest.new({session_id: 'session_id_example'}) # CreateSessionRequest | Session creation request

begin
  # Create a new session
  api_instance.create_session(request)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->create_session: #{e}"
end
```

#### Using the create_session_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> create_session_with_http_info(request)

```ruby
begin
  # Create a new session
  data, status_code, headers = api_instance.create_session_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->create_session_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**CreateSessionRequest**](CreateSessionRequest.md) | Session creation request |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## delete_pty_session

> Hash&lt;String, Object&gt; delete_pty_session(session_id)

Delete a PTY session

Delete a pseudo-terminal session and terminate its process

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
session_id = 'session_id_example' # String | PTY session ID

begin
  # Delete a PTY session
  result = api_instance.delete_pty_session(session_id)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->delete_pty_session: #{e}"
end
```

#### Using the delete_pty_session_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(Hash&lt;String, Object&gt;, Integer, Hash)> delete_pty_session_with_http_info(session_id)

```ruby
begin
  # Delete a PTY session
  data, status_code, headers = api_instance.delete_pty_session_with_http_info(session_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Hash&lt;String, Object&gt;
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->delete_pty_session_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **session_id** | **String** | PTY session ID |  |

### Return type

**Hash&lt;String, Object&gt;**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## delete_session

> delete_session(session_id)

Delete a session

Delete an existing shell session

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
session_id = 'session_id_example' # String | Session ID

begin
  # Delete a session
  api_instance.delete_session(session_id)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->delete_session: #{e}"
end
```

#### Using the delete_session_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> delete_session_with_http_info(session_id)

```ruby
begin
  # Delete a session
  data, status_code, headers = api_instance.delete_session_with_http_info(session_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->delete_session_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **session_id** | **String** | Session ID |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## execute_command

> <ExecuteResponse> execute_command(request)

Execute a command

Execute a shell command and return the output and exit code

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
request = DaytonaToolboxApiClient::ExecuteRequest.new({command: 'command_example'}) # ExecuteRequest | Command execution request

begin
  # Execute a command
  result = api_instance.execute_command(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->execute_command: #{e}"
end
```

#### Using the execute_command_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ExecuteResponse>, Integer, Hash)> execute_command_with_http_info(request)

```ruby
begin
  # Execute a command
  data, status_code, headers = api_instance.execute_command_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ExecuteResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->execute_command_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**ExecuteRequest**](ExecuteRequest.md) | Command execution request |  |

### Return type

[**ExecuteResponse**](ExecuteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## get_entrypoint_logs

> String get_entrypoint_logs(opts)

Get entrypoint logs

Get logs for a sandbox entrypoint session. Supports both HTTP and WebSocket streaming.

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
opts = {
  follow: true # Boolean | Follow logs in real-time (WebSocket only)
}

begin
  # Get entrypoint logs
  result = api_instance.get_entrypoint_logs(opts)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->get_entrypoint_logs: #{e}"
end
```

#### Using the get_entrypoint_logs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(String, Integer, Hash)> get_entrypoint_logs_with_http_info(opts)

```ruby
begin
  # Get entrypoint logs
  data, status_code, headers = api_instance.get_entrypoint_logs_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => String
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->get_entrypoint_logs_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **follow** | **Boolean** | Follow logs in real-time (WebSocket only) | [optional] |

### Return type

**String**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain


## get_entrypoint_session

> <Session> get_entrypoint_session

Get entrypoint session details

Get details of an entrypoint session including its commands

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new

begin
  # Get entrypoint session details
  result = api_instance.get_entrypoint_session
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->get_entrypoint_session: #{e}"
end
```

#### Using the get_entrypoint_session_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Session>, Integer, Hash)> get_entrypoint_session_with_http_info

```ruby
begin
  # Get entrypoint session details
  data, status_code, headers = api_instance.get_entrypoint_session_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Session>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->get_entrypoint_session_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_pty_session

> <PtySessionInfo> get_pty_session(session_id)

Get PTY session information

Get detailed information about a specific pseudo-terminal session

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
session_id = 'session_id_example' # String | PTY session ID

begin
  # Get PTY session information
  result = api_instance.get_pty_session(session_id)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->get_pty_session: #{e}"
end
```

#### Using the get_pty_session_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<PtySessionInfo>, Integer, Hash)> get_pty_session_with_http_info(session_id)

```ruby
begin
  # Get PTY session information
  data, status_code, headers = api_instance.get_pty_session_with_http_info(session_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <PtySessionInfo>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->get_pty_session_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **session_id** | **String** | PTY session ID |  |

### Return type

[**PtySessionInfo**](PtySessionInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_session

> <Session> get_session(session_id)

Get session details

Get details of a specific session including its commands

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
session_id = 'session_id_example' # String | Session ID

begin
  # Get session details
  result = api_instance.get_session(session_id)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->get_session: #{e}"
end
```

#### Using the get_session_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Session>, Integer, Hash)> get_session_with_http_info(session_id)

```ruby
begin
  # Get session details
  data, status_code, headers = api_instance.get_session_with_http_info(session_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Session>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->get_session_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **session_id** | **String** | Session ID |  |

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_session_command

> <Command> get_session_command(session_id, command_id)

Get session command details

Get details of a specific command within a session

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
session_id = 'session_id_example' # String | Session ID
command_id = 'command_id_example' # String | Command ID

begin
  # Get session command details
  result = api_instance.get_session_command(session_id, command_id)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->get_session_command: #{e}"
end
```

#### Using the get_session_command_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Command>, Integer, Hash)> get_session_command_with_http_info(session_id, command_id)

```ruby
begin
  # Get session command details
  data, status_code, headers = api_instance.get_session_command_with_http_info(session_id, command_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Command>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->get_session_command_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **session_id** | **String** | Session ID |  |
| **command_id** | **String** | Command ID |  |

### Return type

[**Command**](Command.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_session_command_logs

> String get_session_command_logs(session_id, command_id, opts)

Get session command logs

Get logs for a specific command within a session. Supports both HTTP and WebSocket streaming.

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
session_id = 'session_id_example' # String | Session ID
command_id = 'command_id_example' # String | Command ID
opts = {
  follow: true # Boolean | Follow logs in real-time (WebSocket only)
}

begin
  # Get session command logs
  result = api_instance.get_session_command_logs(session_id, command_id, opts)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->get_session_command_logs: #{e}"
end
```

#### Using the get_session_command_logs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(String, Integer, Hash)> get_session_command_logs_with_http_info(session_id, command_id, opts)

```ruby
begin
  # Get session command logs
  data, status_code, headers = api_instance.get_session_command_logs_with_http_info(session_id, command_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => String
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->get_session_command_logs_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **session_id** | **String** | Session ID |  |
| **command_id** | **String** | Command ID |  |
| **follow** | **Boolean** | Follow logs in real-time (WebSocket only) | [optional] |

### Return type

**String**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain


## list_pty_sessions

> <PtyListResponse> list_pty_sessions

List all PTY sessions

Get a list of all active pseudo-terminal sessions

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new

begin
  # List all PTY sessions
  result = api_instance.list_pty_sessions
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->list_pty_sessions: #{e}"
end
```

#### Using the list_pty_sessions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<PtyListResponse>, Integer, Hash)> list_pty_sessions_with_http_info

```ruby
begin
  # List all PTY sessions
  data, status_code, headers = api_instance.list_pty_sessions_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <PtyListResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->list_pty_sessions_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**PtyListResponse**](PtyListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## list_sessions

> <Array<Session>> list_sessions

List all sessions

Get a list of all active shell sessions

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new

begin
  # List all sessions
  result = api_instance.list_sessions
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->list_sessions: #{e}"
end
```

#### Using the list_sessions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Session>>, Integer, Hash)> list_sessions_with_http_info

```ruby
begin
  # List all sessions
  data, status_code, headers = api_instance.list_sessions_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Session>>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->list_sessions_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**Array&lt;Session&gt;**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## resize_pty_session

> <PtySessionInfo> resize_pty_session(session_id, request)

Resize a PTY session

Resize the terminal dimensions of a pseudo-terminal session

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
session_id = 'session_id_example' # String | PTY session ID
request = DaytonaToolboxApiClient::PtyResizeRequest.new({cols: 37, rows: 37}) # PtyResizeRequest | Resize request with new dimensions

begin
  # Resize a PTY session
  result = api_instance.resize_pty_session(session_id, request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->resize_pty_session: #{e}"
end
```

#### Using the resize_pty_session_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<PtySessionInfo>, Integer, Hash)> resize_pty_session_with_http_info(session_id, request)

```ruby
begin
  # Resize a PTY session
  data, status_code, headers = api_instance.resize_pty_session_with_http_info(session_id, request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <PtySessionInfo>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->resize_pty_session_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **session_id** | **String** | PTY session ID |  |
| **request** | [**PtyResizeRequest**](PtyResizeRequest.md) | Resize request with new dimensions |  |

### Return type

[**PtySessionInfo**](PtySessionInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## send_input

> send_input(session_id, command_id, request)

Send input to command

Send input data to a running command in a session for interactive execution

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
session_id = 'session_id_example' # String | Session ID
command_id = 'command_id_example' # String | Command ID
request = DaytonaToolboxApiClient::SessionSendInputRequest.new({data: 'data_example'}) # SessionSendInputRequest | Input send request

begin
  # Send input to command
  api_instance.send_input(session_id, command_id, request)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->send_input: #{e}"
end
```

#### Using the send_input_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> send_input_with_http_info(session_id, command_id, request)

```ruby
begin
  # Send input to command
  data, status_code, headers = api_instance.send_input_with_http_info(session_id, command_id, request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->send_input_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **session_id** | **String** | Session ID |  |
| **command_id** | **String** | Command ID |  |
| **request** | [**SessionSendInputRequest**](SessionSendInputRequest.md) | Input send request |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## session_execute_command

> <SessionExecuteResponse> session_execute_command(session_id, request)

Execute command in session

Execute a command within an existing shell session

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ProcessApi.new
session_id = 'session_id_example' # String | Session ID
request = DaytonaToolboxApiClient::SessionExecuteRequest.new({command: 'command_example'}) # SessionExecuteRequest | Command execution request

begin
  # Execute command in session
  result = api_instance.session_execute_command(session_id, request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->session_execute_command: #{e}"
end
```

#### Using the session_execute_command_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SessionExecuteResponse>, Integer, Hash)> session_execute_command_with_http_info(session_id, request)

```ruby
begin
  # Execute command in session
  data, status_code, headers = api_instance.session_execute_command_with_http_info(session_id, request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SessionExecuteResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ProcessApi->session_execute_command_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **session_id** | **String** | Session ID |  |
| **request** | [**SessionExecuteRequest**](SessionExecuteRequest.md) | Command execution request |  |

### Return type

[**SessionExecuteResponse**](SessionExecuteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

