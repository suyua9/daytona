# daytona_toolbox_api_client_async.LspApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**completions**](LspApi.md#completions) | **POST** /lsp/completions | Get code completions
[**did_close**](LspApi.md#did_close) | **POST** /lsp/did-close | Notify document closed
[**did_open**](LspApi.md#did_open) | **POST** /lsp/did-open | Notify document opened
[**document_symbols**](LspApi.md#document_symbols) | **GET** /lsp/document-symbols | Get document symbols
[**start**](LspApi.md#start) | **POST** /lsp/start | Start LSP server
[**stop**](LspApi.md#stop) | **POST** /lsp/stop | Stop LSP server
[**workspace_symbols**](LspApi.md#workspace_symbols) | **GET** /lsp/workspacesymbols | Get workspace symbols


# **completions**
> CompletionList completions(request)

Get code completions

Get code completion suggestions from the LSP server

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.completion_list import CompletionList
from daytona_toolbox_api_client_async.models.lsp_completion_params import LspCompletionParams
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
    api_instance = daytona_toolbox_api_client_async.LspApi(api_client)
    request = daytona_toolbox_api_client_async.LspCompletionParams() # LspCompletionParams | Completion request

    try:
        # Get code completions
        api_response = await api_instance.completions(request)
        print("The response of LspApi->completions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling LspApi->completions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**LspCompletionParams**](LspCompletionParams.md)| Completion request | 

### Return type

[**CompletionList**](CompletionList.md)

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

# **did_close**
> did_close(request)

Notify document closed

Notify the LSP server that a document has been closed

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.lsp_document_request import LspDocumentRequest
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
    api_instance = daytona_toolbox_api_client_async.LspApi(api_client)
    request = daytona_toolbox_api_client_async.LspDocumentRequest() # LspDocumentRequest | Document request

    try:
        # Notify document closed
        await api_instance.did_close(request)
    except Exception as e:
        print("Exception when calling LspApi->did_close: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**LspDocumentRequest**](LspDocumentRequest.md)| Document request | 

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
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **did_open**
> did_open(request)

Notify document opened

Notify the LSP server that a document has been opened

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.lsp_document_request import LspDocumentRequest
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
    api_instance = daytona_toolbox_api_client_async.LspApi(api_client)
    request = daytona_toolbox_api_client_async.LspDocumentRequest() # LspDocumentRequest | Document request

    try:
        # Notify document opened
        await api_instance.did_open(request)
    except Exception as e:
        print("Exception when calling LspApi->did_open: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**LspDocumentRequest**](LspDocumentRequest.md)| Document request | 

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
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **document_symbols**
> List[LspSymbol] document_symbols(language_id, path_to_project, uri)

Get document symbols

Get symbols (functions, classes, etc.) from a document

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.lsp_symbol import LspSymbol
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
    api_instance = daytona_toolbox_api_client_async.LspApi(api_client)
    language_id = 'language_id_example' # str | Language ID (e.g., python, typescript)
    path_to_project = 'path_to_project_example' # str | Path to project
    uri = 'uri_example' # str | Document URI

    try:
        # Get document symbols
        api_response = await api_instance.document_symbols(language_id, path_to_project, uri)
        print("The response of LspApi->document_symbols:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling LspApi->document_symbols: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language_id** | **str**| Language ID (e.g., python, typescript) | 
 **path_to_project** | **str**| Path to project | 
 **uri** | **str**| Document URI | 

### Return type

[**List[LspSymbol]**](LspSymbol.md)

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

# **start**
> start(request)

Start LSP server

Start a Language Server Protocol server for the specified language

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.lsp_server_request import LspServerRequest
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
    api_instance = daytona_toolbox_api_client_async.LspApi(api_client)
    request = daytona_toolbox_api_client_async.LspServerRequest() # LspServerRequest | LSP server request

    try:
        # Start LSP server
        await api_instance.start(request)
    except Exception as e:
        print("Exception when calling LspApi->start: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**LspServerRequest**](LspServerRequest.md)| LSP server request | 

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
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **stop**
> stop(request)

Stop LSP server

Stop a Language Server Protocol server

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.lsp_server_request import LspServerRequest
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
    api_instance = daytona_toolbox_api_client_async.LspApi(api_client)
    request = daytona_toolbox_api_client_async.LspServerRequest() # LspServerRequest | LSP server request

    try:
        # Stop LSP server
        await api_instance.stop(request)
    except Exception as e:
        print("Exception when calling LspApi->stop: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**LspServerRequest**](LspServerRequest.md)| LSP server request | 

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
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **workspace_symbols**
> List[LspSymbol] workspace_symbols(query, language_id, path_to_project)

Get workspace symbols

Search for symbols across the entire workspace

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.lsp_symbol import LspSymbol
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
    api_instance = daytona_toolbox_api_client_async.LspApi(api_client)
    query = 'query_example' # str | Search query
    language_id = 'language_id_example' # str | Language ID (e.g., python, typescript)
    path_to_project = 'path_to_project_example' # str | Path to project

    try:
        # Get workspace symbols
        api_response = await api_instance.workspace_symbols(query, language_id, path_to_project)
        print("The response of LspApi->workspace_symbols:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling LspApi->workspace_symbols: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **str**| Search query | 
 **language_id** | **str**| Language ID (e.g., python, typescript) | 
 **path_to_project** | **str**| Path to project | 

### Return type

[**List[LspSymbol]**](LspSymbol.md)

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

