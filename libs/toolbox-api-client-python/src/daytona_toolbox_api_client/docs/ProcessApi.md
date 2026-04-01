# daytona_toolbox_api_client.ProcessApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**code_run**](ProcessApi.md#code_run) | **POST** /process/code-run | Execute code
[**connect_pty_session**](ProcessApi.md#connect_pty_session) | **GET** /process/pty/{sessionId}/connect | Connect to PTY session via WebSocket
[**create_pty_session**](ProcessApi.md#create_pty_session) | **POST** /process/pty | Create a new PTY session
[**create_session**](ProcessApi.md#create_session) | **POST** /process/session | Create a new session
[**delete_pty_session**](ProcessApi.md#delete_pty_session) | **DELETE** /process/pty/{sessionId} | Delete a PTY session
[**delete_session**](ProcessApi.md#delete_session) | **DELETE** /process/session/{sessionId} | Delete a session
[**execute_command**](ProcessApi.md#execute_command) | **POST** /process/execute | Execute a command
[**get_entrypoint_logs**](ProcessApi.md#get_entrypoint_logs) | **GET** /process/session/entrypoint/logs | Get entrypoint logs
[**get_entrypoint_session**](ProcessApi.md#get_entrypoint_session) | **GET** /process/session/entrypoint | Get entrypoint session details
[**get_pty_session**](ProcessApi.md#get_pty_session) | **GET** /process/pty/{sessionId} | Get PTY session information
[**get_session**](ProcessApi.md#get_session) | **GET** /process/session/{sessionId} | Get session details
[**get_session_command**](ProcessApi.md#get_session_command) | **GET** /process/session/{sessionId}/command/{commandId} | Get session command details
[**get_session_command_logs**](ProcessApi.md#get_session_command_logs) | **GET** /process/session/{sessionId}/command/{commandId}/logs | Get session command logs
[**list_pty_sessions**](ProcessApi.md#list_pty_sessions) | **GET** /process/pty | List all PTY sessions
[**list_sessions**](ProcessApi.md#list_sessions) | **GET** /process/session | List all sessions
[**resize_pty_session**](ProcessApi.md#resize_pty_session) | **POST** /process/pty/{sessionId}/resize | Resize a PTY session
[**send_input**](ProcessApi.md#send_input) | **POST** /process/session/{sessionId}/command/{commandId}/input | Send input to command
[**session_execute_command**](ProcessApi.md#session_execute_command) | **POST** /process/session/{sessionId}/exec | Execute command in session


# **code_run**
> CodeRunResponse code_run(request)

Execute code

Execute Python, JavaScript, or TypeScript code and return output, exit code, and artifacts

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.code_run_request import CodeRunRequest
from daytona_toolbox_api_client.models.code_run_response import CodeRunResponse
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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    request = daytona_toolbox_api_client.CodeRunRequest() # CodeRunRequest | Code execution request

    try:
        # Execute code
        api_response = api_instance.code_run(request)
        print("The response of ProcessApi->code_run:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProcessApi->code_run: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**CodeRunRequest**](CodeRunRequest.md)| Code execution request | 

### Return type

[**CodeRunResponse**](CodeRunResponse.md)

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

# **connect_pty_session**
> connect_pty_session(session_id)

Connect to PTY session via WebSocket

Establish a WebSocket connection to interact with a pseudo-terminal session

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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    session_id = 'session_id_example' # str | PTY session ID

    try:
        # Connect to PTY session via WebSocket
        api_instance.connect_pty_session(session_id)
    except Exception as e:
        print("Exception when calling ProcessApi->connect_pty_session: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **session_id** | **str**| PTY session ID | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**101** | Switching Protocols - WebSocket connection established |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_pty_session**
> PtyCreateResponse create_pty_session(request)

Create a new PTY session

Create a new pseudo-terminal session with specified configuration

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.pty_create_request import PtyCreateRequest
from daytona_toolbox_api_client.models.pty_create_response import PtyCreateResponse
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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    request = daytona_toolbox_api_client.PtyCreateRequest() # PtyCreateRequest | PTY session creation request

    try:
        # Create a new PTY session
        api_response = api_instance.create_pty_session(request)
        print("The response of ProcessApi->create_pty_session:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProcessApi->create_pty_session: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PtyCreateRequest**](PtyCreateRequest.md)| PTY session creation request | 

### Return type

[**PtyCreateResponse**](PtyCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Created |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_session**
> create_session(request)

Create a new session

Create a new shell session for command execution

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.create_session_request import CreateSessionRequest
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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    request = daytona_toolbox_api_client.CreateSessionRequest() # CreateSessionRequest | Session creation request

    try:
        # Create a new session
        api_instance.create_session(request)
    except Exception as e:
        print("Exception when calling ProcessApi->create_session: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**CreateSessionRequest**](CreateSessionRequest.md)| Session creation request | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Created |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_pty_session**
> Dict[str, object] delete_pty_session(session_id)

Delete a PTY session

Delete a pseudo-terminal session and terminate its process

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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    session_id = 'session_id_example' # str | PTY session ID

    try:
        # Delete a PTY session
        api_response = api_instance.delete_pty_session(session_id)
        print("The response of ProcessApi->delete_pty_session:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProcessApi->delete_pty_session: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **session_id** | **str**| PTY session ID | 

### Return type

**Dict[str, object]**

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

# **delete_session**
> delete_session(session_id)

Delete a session

Delete an existing shell session

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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    session_id = 'session_id_example' # str | Session ID

    try:
        # Delete a session
        api_instance.delete_session(session_id)
    except Exception as e:
        print("Exception when calling ProcessApi->delete_session: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **session_id** | **str**| Session ID | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | No Content |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **execute_command**
> ExecuteResponse execute_command(request)

Execute a command

Execute a shell command and return the output and exit code

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.execute_request import ExecuteRequest
from daytona_toolbox_api_client.models.execute_response import ExecuteResponse
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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    request = daytona_toolbox_api_client.ExecuteRequest() # ExecuteRequest | Command execution request

    try:
        # Execute a command
        api_response = api_instance.execute_command(request)
        print("The response of ProcessApi->execute_command:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProcessApi->execute_command: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ExecuteRequest**](ExecuteRequest.md)| Command execution request | 

### Return type

[**ExecuteResponse**](ExecuteResponse.md)

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

# **get_entrypoint_logs**
> str get_entrypoint_logs(follow=follow)

Get entrypoint logs

Get logs for a sandbox entrypoint session. Supports both HTTP and WebSocket streaming.

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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    follow = True # bool | Follow logs in real-time (WebSocket only) (optional)

    try:
        # Get entrypoint logs
        api_response = api_instance.get_entrypoint_logs(follow=follow)
        print("The response of ProcessApi->get_entrypoint_logs:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProcessApi->get_entrypoint_logs: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **follow** | **bool**| Follow logs in real-time (WebSocket only) | [optional] 

### Return type

**str**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Entrypoint log content |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_entrypoint_session**
> Session get_entrypoint_session()

Get entrypoint session details

Get details of an entrypoint session including its commands

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.session import Session
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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)

    try:
        # Get entrypoint session details
        api_response = api_instance.get_entrypoint_session()
        print("The response of ProcessApi->get_entrypoint_session:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProcessApi->get_entrypoint_session: %s\n" % e)
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

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_pty_session**
> PtySessionInfo get_pty_session(session_id)

Get PTY session information

Get detailed information about a specific pseudo-terminal session

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.pty_session_info import PtySessionInfo
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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    session_id = 'session_id_example' # str | PTY session ID

    try:
        # Get PTY session information
        api_response = api_instance.get_pty_session(session_id)
        print("The response of ProcessApi->get_pty_session:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProcessApi->get_pty_session: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **session_id** | **str**| PTY session ID | 

### Return type

[**PtySessionInfo**](PtySessionInfo.md)

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

# **get_session**
> Session get_session(session_id)

Get session details

Get details of a specific session including its commands

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.session import Session
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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    session_id = 'session_id_example' # str | Session ID

    try:
        # Get session details
        api_response = api_instance.get_session(session_id)
        print("The response of ProcessApi->get_session:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProcessApi->get_session: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **session_id** | **str**| Session ID | 

### Return type

[**Session**](Session.md)

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

# **get_session_command**
> Command get_session_command(session_id, command_id)

Get session command details

Get details of a specific command within a session

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.command import Command
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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    session_id = 'session_id_example' # str | Session ID
    command_id = 'command_id_example' # str | Command ID

    try:
        # Get session command details
        api_response = api_instance.get_session_command(session_id, command_id)
        print("The response of ProcessApi->get_session_command:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProcessApi->get_session_command: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **session_id** | **str**| Session ID | 
 **command_id** | **str**| Command ID | 

### Return type

[**Command**](Command.md)

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

# **get_session_command_logs**
> str get_session_command_logs(session_id, command_id, follow=follow)

Get session command logs

Get logs for a specific command within a session. Supports both HTTP and WebSocket streaming.

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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    session_id = 'session_id_example' # str | Session ID
    command_id = 'command_id_example' # str | Command ID
    follow = True # bool | Follow logs in real-time (WebSocket only) (optional)

    try:
        # Get session command logs
        api_response = api_instance.get_session_command_logs(session_id, command_id, follow=follow)
        print("The response of ProcessApi->get_session_command_logs:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProcessApi->get_session_command_logs: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **session_id** | **str**| Session ID | 
 **command_id** | **str**| Command ID | 
 **follow** | **bool**| Follow logs in real-time (WebSocket only) | [optional] 

### Return type

**str**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Log content |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_pty_sessions**
> PtyListResponse list_pty_sessions()

List all PTY sessions

Get a list of all active pseudo-terminal sessions

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.pty_list_response import PtyListResponse
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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)

    try:
        # List all PTY sessions
        api_response = api_instance.list_pty_sessions()
        print("The response of ProcessApi->list_pty_sessions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProcessApi->list_pty_sessions: %s\n" % e)
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

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_sessions**
> List[Session] list_sessions()

List all sessions

Get a list of all active shell sessions

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.session import Session
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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)

    try:
        # List all sessions
        api_response = api_instance.list_sessions()
        print("The response of ProcessApi->list_sessions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProcessApi->list_sessions: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**List[Session]**](Session.md)

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

# **resize_pty_session**
> PtySessionInfo resize_pty_session(session_id, request)

Resize a PTY session

Resize the terminal dimensions of a pseudo-terminal session

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.pty_resize_request import PtyResizeRequest
from daytona_toolbox_api_client.models.pty_session_info import PtySessionInfo
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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    session_id = 'session_id_example' # str | PTY session ID
    request = daytona_toolbox_api_client.PtyResizeRequest() # PtyResizeRequest | Resize request with new dimensions

    try:
        # Resize a PTY session
        api_response = api_instance.resize_pty_session(session_id, request)
        print("The response of ProcessApi->resize_pty_session:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProcessApi->resize_pty_session: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **session_id** | **str**| PTY session ID | 
 **request** | [**PtyResizeRequest**](PtyResizeRequest.md)| Resize request with new dimensions | 

### Return type

[**PtySessionInfo**](PtySessionInfo.md)

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

# **send_input**
> send_input(session_id, command_id, request)

Send input to command

Send input data to a running command in a session for interactive execution

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.session_send_input_request import SessionSendInputRequest
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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    session_id = 'session_id_example' # str | Session ID
    command_id = 'command_id_example' # str | Command ID
    request = daytona_toolbox_api_client.SessionSendInputRequest() # SessionSendInputRequest | Input send request

    try:
        # Send input to command
        api_instance.send_input(session_id, command_id, request)
    except Exception as e:
        print("Exception when calling ProcessApi->send_input: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **session_id** | **str**| Session ID | 
 **command_id** | **str**| Command ID | 
 **request** | [**SessionSendInputRequest**](SessionSendInputRequest.md)| Input send request | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | No Content |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **session_execute_command**
> SessionExecuteResponse session_execute_command(session_id, request)

Execute command in session

Execute a command within an existing shell session

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.session_execute_request import SessionExecuteRequest
from daytona_toolbox_api_client.models.session_execute_response import SessionExecuteResponse
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
    api_instance = daytona_toolbox_api_client.ProcessApi(api_client)
    session_id = 'session_id_example' # str | Session ID
    request = daytona_toolbox_api_client.SessionExecuteRequest() # SessionExecuteRequest | Command execution request

    try:
        # Execute command in session
        api_response = api_instance.session_execute_command(session_id, request)
        print("The response of ProcessApi->session_execute_command:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProcessApi->session_execute_command: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **session_id** | **str**| Session ID | 
 **request** | [**SessionExecuteRequest**](SessionExecuteRequest.md)| Command execution request | 

### Return type

[**SessionExecuteResponse**](SessionExecuteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**202** | Accepted |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

