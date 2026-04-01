# DaytonaToolboxApiClient::InterpreterApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**create_interpreter_context**](InterpreterApi.md#create_interpreter_context) | **POST** /process/interpreter/context | Create a new interpreter context |
| [**delete_interpreter_context**](InterpreterApi.md#delete_interpreter_context) | **DELETE** /process/interpreter/context/{id} | Delete an interpreter context |
| [**execute_interpreter_code**](InterpreterApi.md#execute_interpreter_code) | **GET** /process/interpreter/execute | Execute code in an interpreter context |
| [**list_interpreter_contexts**](InterpreterApi.md#list_interpreter_contexts) | **GET** /process/interpreter/context | List all user-created interpreter contexts |


## create_interpreter_context

> <InterpreterContext> create_interpreter_context(request)

Create a new interpreter context

Creates a new isolated interpreter context with optional working directory and language

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::InterpreterApi.new
request = DaytonaToolboxApiClient::CreateContextRequest.new # CreateContextRequest | Context configuration

begin
  # Create a new interpreter context
  result = api_instance.create_interpreter_context(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling InterpreterApi->create_interpreter_context: #{e}"
end
```

#### Using the create_interpreter_context_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<InterpreterContext>, Integer, Hash)> create_interpreter_context_with_http_info(request)

```ruby
begin
  # Create a new interpreter context
  data, status_code, headers = api_instance.create_interpreter_context_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <InterpreterContext>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling InterpreterApi->create_interpreter_context_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**CreateContextRequest**](CreateContextRequest.md) | Context configuration |  |

### Return type

[**InterpreterContext**](InterpreterContext.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## delete_interpreter_context

> Hash&lt;String, String&gt; delete_interpreter_context(id)

Delete an interpreter context

Deletes an interpreter context and shuts down its worker process

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::InterpreterApi.new
id = 'id_example' # String | Context ID

begin
  # Delete an interpreter context
  result = api_instance.delete_interpreter_context(id)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling InterpreterApi->delete_interpreter_context: #{e}"
end
```

#### Using the delete_interpreter_context_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(Hash&lt;String, String&gt;, Integer, Hash)> delete_interpreter_context_with_http_info(id)

```ruby
begin
  # Delete an interpreter context
  data, status_code, headers = api_instance.delete_interpreter_context_with_http_info(id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Hash&lt;String, String&gt;
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling InterpreterApi->delete_interpreter_context_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** | Context ID |  |

### Return type

**Hash&lt;String, String&gt;**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## execute_interpreter_code

> execute_interpreter_code

Execute code in an interpreter context

Executes code in a specified context (or default context if not specified) via WebSocket streaming

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::InterpreterApi.new

begin
  # Execute code in an interpreter context
  api_instance.execute_interpreter_code
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling InterpreterApi->execute_interpreter_code: #{e}"
end
```

#### Using the execute_interpreter_code_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> execute_interpreter_code_with_http_info

```ruby
begin
  # Execute code in an interpreter context
  data, status_code, headers = api_instance.execute_interpreter_code_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling InterpreterApi->execute_interpreter_code_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## list_interpreter_contexts

> <ListContextsResponse> list_interpreter_contexts

List all user-created interpreter contexts

Returns information about all user-created interpreter contexts (excludes default context)

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::InterpreterApi.new

begin
  # List all user-created interpreter contexts
  result = api_instance.list_interpreter_contexts
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling InterpreterApi->list_interpreter_contexts: #{e}"
end
```

#### Using the list_interpreter_contexts_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ListContextsResponse>, Integer, Hash)> list_interpreter_contexts_with_http_info

```ruby
begin
  # List all user-created interpreter contexts
  data, status_code, headers = api_instance.list_interpreter_contexts_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ListContextsResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling InterpreterApi->list_interpreter_contexts_with_http_info: #{e}"
end
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

