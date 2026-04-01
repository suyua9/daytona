# daytona_toolbox_api_client.InterpreterApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_interpreter_context**](InterpreterApi.md#create_interpreter_context) | **POST** /process/interpreter/context | Create a new interpreter context
[**delete_interpreter_context**](InterpreterApi.md#delete_interpreter_context) | **DELETE** /process/interpreter/context/{id} | Delete an interpreter context
[**execute_interpreter_code**](InterpreterApi.md#execute_interpreter_code) | **GET** /process/interpreter/execute | Execute code in an interpreter context
[**list_interpreter_contexts**](InterpreterApi.md#list_interpreter_contexts) | **GET** /process/interpreter/context | List all user-created interpreter contexts


# **create_interpreter_context**
> InterpreterContext create_interpreter_context(request)

Create a new interpreter context

Creates a new isolated interpreter context with optional working directory and language

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.create_context_request import CreateContextRequest
from daytona_toolbox_api_client.models.interpreter_context import InterpreterContext
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
    api_instance = daytona_toolbox_api_client.InterpreterApi(api_client)
    request = daytona_toolbox_api_client.CreateContextRequest() # CreateContextRequest | Context configuration

    try:
        # Create a new interpreter context
        api_response = api_instance.create_interpreter_context(request)
        print("The response of InterpreterApi->create_interpreter_context:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InterpreterApi->create_interpreter_context: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**CreateContextRequest**](CreateContextRequest.md)| Context configuration | 

### Return type

[**InterpreterContext**](InterpreterContext.md)

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
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_interpreter_context**
> Dict[str, str] delete_interpreter_context(id)

Delete an interpreter context

Deletes an interpreter context and shuts down its worker process

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
    api_instance = daytona_toolbox_api_client.InterpreterApi(api_client)
    id = 'id_example' # str | Context ID

    try:
        # Delete an interpreter context
        api_response = api_instance.delete_interpreter_context(id)
        print("The response of InterpreterApi->delete_interpreter_context:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InterpreterApi->delete_interpreter_context: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| Context ID | 

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
**400** | Bad Request |  -  |
**404** | Not Found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **execute_interpreter_code**
> execute_interpreter_code()

Execute code in an interpreter context

Executes code in a specified context (or default context if not specified) via WebSocket streaming

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
    api_instance = daytona_toolbox_api_client.InterpreterApi(api_client)

    try:
        # Execute code in an interpreter context
        api_instance.execute_interpreter_code()
    except Exception as e:
        print("Exception when calling InterpreterApi->execute_interpreter_code: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**101** | Switching Protocols |  * Connection - Upgrade <br>  * Upgrade - websocket <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_interpreter_contexts**
> ListContextsResponse list_interpreter_contexts()

List all user-created interpreter contexts

Returns information about all user-created interpreter contexts (excludes default context)

### Example


```python
import daytona_toolbox_api_client
from daytona_toolbox_api_client.models.list_contexts_response import ListContextsResponse
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
    api_instance = daytona_toolbox_api_client.InterpreterApi(api_client)

    try:
        # List all user-created interpreter contexts
        api_response = api_instance.list_interpreter_contexts()
        print("The response of InterpreterApi->list_interpreter_contexts:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InterpreterApi->list_interpreter_contexts: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**ListContextsResponse**](ListContextsResponse.md)

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

