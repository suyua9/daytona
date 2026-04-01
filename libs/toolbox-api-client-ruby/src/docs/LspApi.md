# DaytonaToolboxApiClient::LspApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**completions**](LspApi.md#completions) | **POST** /lsp/completions | Get code completions |
| [**did_close**](LspApi.md#did_close) | **POST** /lsp/did-close | Notify document closed |
| [**did_open**](LspApi.md#did_open) | **POST** /lsp/did-open | Notify document opened |
| [**document_symbols**](LspApi.md#document_symbols) | **GET** /lsp/document-symbols | Get document symbols |
| [**start**](LspApi.md#start) | **POST** /lsp/start | Start LSP server |
| [**stop**](LspApi.md#stop) | **POST** /lsp/stop | Stop LSP server |
| [**workspace_symbols**](LspApi.md#workspace_symbols) | **GET** /lsp/workspacesymbols | Get workspace symbols |


## completions

> <CompletionList> completions(request)

Get code completions

Get code completion suggestions from the LSP server

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::LspApi.new
request = DaytonaToolboxApiClient::LspCompletionParams.new({language_id: 'language_id_example', path_to_project: 'path_to_project_example', position: DaytonaToolboxApiClient::LspPosition.new({character: 37, line: 37}), uri: 'uri_example'}) # LspCompletionParams | Completion request

begin
  # Get code completions
  result = api_instance.completions(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling LspApi->completions: #{e}"
end
```

#### Using the completions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CompletionList>, Integer, Hash)> completions_with_http_info(request)

```ruby
begin
  # Get code completions
  data, status_code, headers = api_instance.completions_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CompletionList>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling LspApi->completions_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**LspCompletionParams**](LspCompletionParams.md) | Completion request |  |

### Return type

[**CompletionList**](CompletionList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## did_close

> did_close(request)

Notify document closed

Notify the LSP server that a document has been closed

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::LspApi.new
request = DaytonaToolboxApiClient::LspDocumentRequest.new({language_id: 'language_id_example', path_to_project: 'path_to_project_example', uri: 'uri_example'}) # LspDocumentRequest | Document request

begin
  # Notify document closed
  api_instance.did_close(request)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling LspApi->did_close: #{e}"
end
```

#### Using the did_close_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> did_close_with_http_info(request)

```ruby
begin
  # Notify document closed
  data, status_code, headers = api_instance.did_close_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling LspApi->did_close_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**LspDocumentRequest**](LspDocumentRequest.md) | Document request |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## did_open

> did_open(request)

Notify document opened

Notify the LSP server that a document has been opened

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::LspApi.new
request = DaytonaToolboxApiClient::LspDocumentRequest.new({language_id: 'language_id_example', path_to_project: 'path_to_project_example', uri: 'uri_example'}) # LspDocumentRequest | Document request

begin
  # Notify document opened
  api_instance.did_open(request)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling LspApi->did_open: #{e}"
end
```

#### Using the did_open_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> did_open_with_http_info(request)

```ruby
begin
  # Notify document opened
  data, status_code, headers = api_instance.did_open_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling LspApi->did_open_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**LspDocumentRequest**](LspDocumentRequest.md) | Document request |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## document_symbols

> <Array<LspSymbol>> document_symbols(language_id, path_to_project, uri)

Get document symbols

Get symbols (functions, classes, etc.) from a document

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::LspApi.new
language_id = 'language_id_example' # String | Language ID (e.g., python, typescript)
path_to_project = 'path_to_project_example' # String | Path to project
uri = 'uri_example' # String | Document URI

begin
  # Get document symbols
  result = api_instance.document_symbols(language_id, path_to_project, uri)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling LspApi->document_symbols: #{e}"
end
```

#### Using the document_symbols_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<LspSymbol>>, Integer, Hash)> document_symbols_with_http_info(language_id, path_to_project, uri)

```ruby
begin
  # Get document symbols
  data, status_code, headers = api_instance.document_symbols_with_http_info(language_id, path_to_project, uri)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<LspSymbol>>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling LspApi->document_symbols_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **language_id** | **String** | Language ID (e.g., python, typescript) |  |
| **path_to_project** | **String** | Path to project |  |
| **uri** | **String** | Document URI |  |

### Return type

[**Array&lt;LspSymbol&gt;**](LspSymbol.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## start

> start(request)

Start LSP server

Start a Language Server Protocol server for the specified language

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::LspApi.new
request = DaytonaToolboxApiClient::LspServerRequest.new({language_id: 'language_id_example', path_to_project: 'path_to_project_example'}) # LspServerRequest | LSP server request

begin
  # Start LSP server
  api_instance.start(request)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling LspApi->start: #{e}"
end
```

#### Using the start_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> start_with_http_info(request)

```ruby
begin
  # Start LSP server
  data, status_code, headers = api_instance.start_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling LspApi->start_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**LspServerRequest**](LspServerRequest.md) | LSP server request |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## stop

> stop(request)

Stop LSP server

Stop a Language Server Protocol server

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::LspApi.new
request = DaytonaToolboxApiClient::LspServerRequest.new({language_id: 'language_id_example', path_to_project: 'path_to_project_example'}) # LspServerRequest | LSP server request

begin
  # Stop LSP server
  api_instance.stop(request)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling LspApi->stop: #{e}"
end
```

#### Using the stop_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> stop_with_http_info(request)

```ruby
begin
  # Stop LSP server
  data, status_code, headers = api_instance.stop_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling LspApi->stop_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**LspServerRequest**](LspServerRequest.md) | LSP server request |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## workspace_symbols

> <Array<LspSymbol>> workspace_symbols(query, language_id, path_to_project)

Get workspace symbols

Search for symbols across the entire workspace

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::LspApi.new
query = 'query_example' # String | Search query
language_id = 'language_id_example' # String | Language ID (e.g., python, typescript)
path_to_project = 'path_to_project_example' # String | Path to project

begin
  # Get workspace symbols
  result = api_instance.workspace_symbols(query, language_id, path_to_project)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling LspApi->workspace_symbols: #{e}"
end
```

#### Using the workspace_symbols_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<LspSymbol>>, Integer, Hash)> workspace_symbols_with_http_info(query, language_id, path_to_project)

```ruby
begin
  # Get workspace symbols
  data, status_code, headers = api_instance.workspace_symbols_with_http_info(query, language_id, path_to_project)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<LspSymbol>>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling LspApi->workspace_symbols_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **query** | **String** | Search query |  |
| **language_id** | **String** | Language ID (e.g., python, typescript) |  |
| **path_to_project** | **String** | Path to project |  |

### Return type

[**Array&lt;LspSymbol&gt;**](LspSymbol.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

