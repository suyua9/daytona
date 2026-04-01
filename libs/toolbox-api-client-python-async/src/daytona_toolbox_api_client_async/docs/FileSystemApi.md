# daytona_toolbox_api_client_async.FileSystemApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_folder**](FileSystemApi.md#create_folder) | **POST** /files/folder | Create a folder
[**delete_file**](FileSystemApi.md#delete_file) | **DELETE** /files | Delete a file or directory
[**download_file**](FileSystemApi.md#download_file) | **GET** /files/download | Download a file
[**download_files**](FileSystemApi.md#download_files) | **POST** /files/bulk-download | Download multiple files
[**find_in_files**](FileSystemApi.md#find_in_files) | **GET** /files/find | Find text in files
[**get_file_info**](FileSystemApi.md#get_file_info) | **GET** /files/info | Get file information
[**list_files**](FileSystemApi.md#list_files) | **GET** /files | List files and directories
[**move_file**](FileSystemApi.md#move_file) | **POST** /files/move | Move or rename file/directory
[**replace_in_files**](FileSystemApi.md#replace_in_files) | **POST** /files/replace | Replace text in files
[**search_files**](FileSystemApi.md#search_files) | **GET** /files/search | Search files by pattern
[**set_file_permissions**](FileSystemApi.md#set_file_permissions) | **POST** /files/permissions | Set file permissions
[**upload_file**](FileSystemApi.md#upload_file) | **POST** /files/upload | Upload a file
[**upload_files**](FileSystemApi.md#upload_files) | **POST** /files/bulk-upload | Upload multiple files


# **create_folder**
> create_folder(path, mode)

Create a folder

Create a folder with the specified path and optional permissions

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
    api_instance = daytona_toolbox_api_client_async.FileSystemApi(api_client)
    path = 'path_example' # str | Folder path to create
    mode = 'mode_example' # str | Octal permission mode (default: 0755)

    try:
        # Create a folder
        await api_instance.create_folder(path, mode)
    except Exception as e:
        print("Exception when calling FileSystemApi->create_folder: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| Folder path to create | 
 **mode** | **str**| Octal permission mode (default: 0755) | 

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
**201** | Created |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_file**
> delete_file(path, recursive=recursive)

Delete a file or directory

Delete a file or directory at the specified path

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
    api_instance = daytona_toolbox_api_client_async.FileSystemApi(api_client)
    path = 'path_example' # str | File or directory path to delete
    recursive = True # bool | Enable recursive deletion for directories (optional)

    try:
        # Delete a file or directory
        await api_instance.delete_file(path, recursive=recursive)
    except Exception as e:
        print("Exception when calling FileSystemApi->delete_file: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| File or directory path to delete | 
 **recursive** | **bool**| Enable recursive deletion for directories | [optional] 

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

# **download_file**
> bytearray download_file(path)

Download a file

Download a file by providing its path

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
    api_instance = daytona_toolbox_api_client_async.FileSystemApi(api_client)
    path = 'path_example' # str | File path to download

    try:
        # Download a file
        api_response = await api_instance.download_file(path)
        print("The response of FileSystemApi->download_file:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FileSystemApi->download_file: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| File path to download | 

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **download_files**
> Dict[str, object] download_files(download_files)

Download multiple files

Download multiple files by providing their paths

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.files_download_request import FilesDownloadRequest
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
    api_instance = daytona_toolbox_api_client_async.FileSystemApi(api_client)
    download_files = daytona_toolbox_api_client_async.FilesDownloadRequest() # FilesDownloadRequest | Paths of files to download

    try:
        # Download multiple files
        api_response = await api_instance.download_files(download_files)
        print("The response of FileSystemApi->download_files:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FileSystemApi->download_files: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **download_files** | [**FilesDownloadRequest**](FilesDownloadRequest.md)| Paths of files to download | 

### Return type

**Dict[str, object]**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_in_files**
> List[Match] find_in_files(path, pattern)

Find text in files

Search for text pattern within files in a directory

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.match import Match
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
    api_instance = daytona_toolbox_api_client_async.FileSystemApi(api_client)
    path = 'path_example' # str | Directory path to search in
    pattern = 'pattern_example' # str | Text pattern to search for

    try:
        # Find text in files
        api_response = await api_instance.find_in_files(path, pattern)
        print("The response of FileSystemApi->find_in_files:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FileSystemApi->find_in_files: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| Directory path to search in | 
 **pattern** | **str**| Text pattern to search for | 

### Return type

[**List[Match]**](Match.md)

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

# **get_file_info**
> FileInfo get_file_info(path)

Get file information

Get detailed information about a file or directory

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.file_info import FileInfo
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
    api_instance = daytona_toolbox_api_client_async.FileSystemApi(api_client)
    path = 'path_example' # str | File or directory path

    try:
        # Get file information
        api_response = await api_instance.get_file_info(path)
        print("The response of FileSystemApi->get_file_info:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FileSystemApi->get_file_info: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| File or directory path | 

### Return type

[**FileInfo**](FileInfo.md)

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

# **list_files**
> List[FileInfo] list_files(path=path)

List files and directories

List files and directories in the specified path

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.file_info import FileInfo
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
    api_instance = daytona_toolbox_api_client_async.FileSystemApi(api_client)
    path = 'path_example' # str | Directory path to list (defaults to working directory) (optional)

    try:
        # List files and directories
        api_response = await api_instance.list_files(path=path)
        print("The response of FileSystemApi->list_files:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FileSystemApi->list_files: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| Directory path to list (defaults to working directory) | [optional] 

### Return type

[**List[FileInfo]**](FileInfo.md)

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

# **move_file**
> move_file(source, destination)

Move or rename file/directory

Move or rename a file or directory from source to destination

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
    api_instance = daytona_toolbox_api_client_async.FileSystemApi(api_client)
    source = 'source_example' # str | Source file or directory path
    destination = 'destination_example' # str | Destination file or directory path

    try:
        # Move or rename file/directory
        await api_instance.move_file(source, destination)
    except Exception as e:
        print("Exception when calling FileSystemApi->move_file: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **str**| Source file or directory path | 
 **destination** | **str**| Destination file or directory path | 

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
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replace_in_files**
> List[ReplaceResult] replace_in_files(request)

Replace text in files

Replace text pattern with new value in multiple files

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.replace_request import ReplaceRequest
from daytona_toolbox_api_client_async.models.replace_result import ReplaceResult
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
    api_instance = daytona_toolbox_api_client_async.FileSystemApi(api_client)
    request = daytona_toolbox_api_client_async.ReplaceRequest() # ReplaceRequest | Replace request

    try:
        # Replace text in files
        api_response = await api_instance.replace_in_files(request)
        print("The response of FileSystemApi->replace_in_files:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FileSystemApi->replace_in_files: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ReplaceRequest**](ReplaceRequest.md)| Replace request | 

### Return type

[**List[ReplaceResult]**](ReplaceResult.md)

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

# **search_files**
> SearchFilesResponse search_files(path, pattern)

Search files by pattern

Search for files matching a specific pattern in a directory

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.search_files_response import SearchFilesResponse
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
    api_instance = daytona_toolbox_api_client_async.FileSystemApi(api_client)
    path = 'path_example' # str | Directory path to search in
    pattern = 'pattern_example' # str | File pattern to match (e.g., *.txt, *.go)

    try:
        # Search files by pattern
        api_response = await api_instance.search_files(path, pattern)
        print("The response of FileSystemApi->search_files:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FileSystemApi->search_files: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| Directory path to search in | 
 **pattern** | **str**| File pattern to match (e.g., *.txt, *.go) | 

### Return type

[**SearchFilesResponse**](SearchFilesResponse.md)

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

# **set_file_permissions**
> set_file_permissions(path, owner=owner, group=group, mode=mode)

Set file permissions

Set file permissions, ownership, and group for a file or directory

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
    api_instance = daytona_toolbox_api_client_async.FileSystemApi(api_client)
    path = 'path_example' # str | File or directory path
    owner = 'owner_example' # str | Owner (username or UID) (optional)
    group = 'group_example' # str | Group (group name or GID) (optional)
    mode = 'mode_example' # str | File mode in octal format (e.g., 0755) (optional)

    try:
        # Set file permissions
        await api_instance.set_file_permissions(path, owner=owner, group=group, mode=mode)
    except Exception as e:
        print("Exception when calling FileSystemApi->set_file_permissions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| File or directory path | 
 **owner** | **str**| Owner (username or UID) | [optional] 
 **group** | **str**| Group (group name or GID) | [optional] 
 **mode** | **str**| File mode in octal format (e.g., 0755) | [optional] 

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
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **upload_file**
> Dict[str, object] upload_file(path, file)

Upload a file

Upload a file to the specified path

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
    api_instance = daytona_toolbox_api_client_async.FileSystemApi(api_client)
    path = 'path_example' # str | Destination path for the uploaded file
    file = None # bytearray | File to upload

    try:
        # Upload a file
        api_response = await api_instance.upload_file(path, file)
        print("The response of FileSystemApi->upload_file:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FileSystemApi->upload_file: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| Destination path for the uploaded file | 
 **file** | **bytearray**| File to upload | 

### Return type

**Dict[str, object]**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **upload_files**
> upload_files()

Upload multiple files

Upload multiple files with their destination paths

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
    api_instance = daytona_toolbox_api_client_async.FileSystemApi(api_client)

    try:
        # Upload multiple files
        await api_instance.upload_files()
    except Exception as e:
        print("Exception when calling FileSystemApi->upload_files: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

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
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

