# DaytonaToolboxApiClient::FileSystemApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**create_folder**](FileSystemApi.md#create_folder) | **POST** /files/folder | Create a folder |
| [**delete_file**](FileSystemApi.md#delete_file) | **DELETE** /files | Delete a file or directory |
| [**download_file**](FileSystemApi.md#download_file) | **GET** /files/download | Download a file |
| [**download_files**](FileSystemApi.md#download_files) | **POST** /files/bulk-download | Download multiple files |
| [**find_in_files**](FileSystemApi.md#find_in_files) | **GET** /files/find | Find text in files |
| [**get_file_info**](FileSystemApi.md#get_file_info) | **GET** /files/info | Get file information |
| [**list_files**](FileSystemApi.md#list_files) | **GET** /files | List files and directories |
| [**move_file**](FileSystemApi.md#move_file) | **POST** /files/move | Move or rename file/directory |
| [**replace_in_files**](FileSystemApi.md#replace_in_files) | **POST** /files/replace | Replace text in files |
| [**search_files**](FileSystemApi.md#search_files) | **GET** /files/search | Search files by pattern |
| [**set_file_permissions**](FileSystemApi.md#set_file_permissions) | **POST** /files/permissions | Set file permissions |
| [**upload_file**](FileSystemApi.md#upload_file) | **POST** /files/upload | Upload a file |
| [**upload_files**](FileSystemApi.md#upload_files) | **POST** /files/bulk-upload | Upload multiple files |


## create_folder

> create_folder(path, mode)

Create a folder

Create a folder with the specified path and optional permissions

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::FileSystemApi.new
path = 'path_example' # String | Folder path to create
mode = 'mode_example' # String | Octal permission mode (default: 0755)

begin
  # Create a folder
  api_instance.create_folder(path, mode)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->create_folder: #{e}"
end
```

#### Using the create_folder_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> create_folder_with_http_info(path, mode)

```ruby
begin
  # Create a folder
  data, status_code, headers = api_instance.create_folder_with_http_info(path, mode)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->create_folder_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | Folder path to create |  |
| **mode** | **String** | Octal permission mode (default: 0755) |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## delete_file

> delete_file(path, opts)

Delete a file or directory

Delete a file or directory at the specified path

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::FileSystemApi.new
path = 'path_example' # String | File or directory path to delete
opts = {
  recursive: true # Boolean | Enable recursive deletion for directories
}

begin
  # Delete a file or directory
  api_instance.delete_file(path, opts)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->delete_file: #{e}"
end
```

#### Using the delete_file_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> delete_file_with_http_info(path, opts)

```ruby
begin
  # Delete a file or directory
  data, status_code, headers = api_instance.delete_file_with_http_info(path, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->delete_file_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | File or directory path to delete |  |
| **recursive** | **Boolean** | Enable recursive deletion for directories | [optional] |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## download_file

> File download_file(path)

Download a file

Download a file by providing its path

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::FileSystemApi.new
path = 'path_example' # String | File path to download

begin
  # Download a file
  result = api_instance.download_file(path)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->download_file: #{e}"
end
```

#### Using the download_file_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(File, Integer, Hash)> download_file_with_http_info(path)

```ruby
begin
  # Download a file
  data, status_code, headers = api_instance.download_file_with_http_info(path)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => File
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->download_file_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | File path to download |  |

### Return type

**File**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream


## download_files

> Hash&lt;String, Object&gt; download_files(download_files)

Download multiple files

Download multiple files by providing their paths

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::FileSystemApi.new
download_files = DaytonaToolboxApiClient::FilesDownloadRequest.new({paths: ['paths_example']}) # FilesDownloadRequest | Paths of files to download

begin
  # Download multiple files
  result = api_instance.download_files(download_files)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->download_files: #{e}"
end
```

#### Using the download_files_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(Hash&lt;String, Object&gt;, Integer, Hash)> download_files_with_http_info(download_files)

```ruby
begin
  # Download multiple files
  data, status_code, headers = api_instance.download_files_with_http_info(download_files)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Hash&lt;String, Object&gt;
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->download_files_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **download_files** | [**FilesDownloadRequest**](FilesDownloadRequest.md) | Paths of files to download |  |

### Return type

**Hash&lt;String, Object&gt;**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: multipart/form-data


## find_in_files

> <Array<Match>> find_in_files(path, pattern)

Find text in files

Search for text pattern within files in a directory

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::FileSystemApi.new
path = 'path_example' # String | Directory path to search in
pattern = 'pattern_example' # String | Text pattern to search for

begin
  # Find text in files
  result = api_instance.find_in_files(path, pattern)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->find_in_files: #{e}"
end
```

#### Using the find_in_files_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Match>>, Integer, Hash)> find_in_files_with_http_info(path, pattern)

```ruby
begin
  # Find text in files
  data, status_code, headers = api_instance.find_in_files_with_http_info(path, pattern)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Match>>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->find_in_files_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | Directory path to search in |  |
| **pattern** | **String** | Text pattern to search for |  |

### Return type

[**Array&lt;Match&gt;**](Match.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_file_info

> <FileInfo> get_file_info(path)

Get file information

Get detailed information about a file or directory

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::FileSystemApi.new
path = 'path_example' # String | File or directory path

begin
  # Get file information
  result = api_instance.get_file_info(path)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->get_file_info: #{e}"
end
```

#### Using the get_file_info_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<FileInfo>, Integer, Hash)> get_file_info_with_http_info(path)

```ruby
begin
  # Get file information
  data, status_code, headers = api_instance.get_file_info_with_http_info(path)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <FileInfo>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->get_file_info_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | File or directory path |  |

### Return type

[**FileInfo**](FileInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## list_files

> <Array<FileInfo>> list_files(opts)

List files and directories

List files and directories in the specified path

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::FileSystemApi.new
opts = {
  path: 'path_example' # String | Directory path to list (defaults to working directory)
}

begin
  # List files and directories
  result = api_instance.list_files(opts)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->list_files: #{e}"
end
```

#### Using the list_files_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<FileInfo>>, Integer, Hash)> list_files_with_http_info(opts)

```ruby
begin
  # List files and directories
  data, status_code, headers = api_instance.list_files_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<FileInfo>>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->list_files_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | Directory path to list (defaults to working directory) | [optional] |

### Return type

[**Array&lt;FileInfo&gt;**](FileInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## move_file

> move_file(source, destination)

Move or rename file/directory

Move or rename a file or directory from source to destination

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::FileSystemApi.new
source = 'source_example' # String | Source file or directory path
destination = 'destination_example' # String | Destination file or directory path

begin
  # Move or rename file/directory
  api_instance.move_file(source, destination)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->move_file: #{e}"
end
```

#### Using the move_file_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> move_file_with_http_info(source, destination)

```ruby
begin
  # Move or rename file/directory
  data, status_code, headers = api_instance.move_file_with_http_info(source, destination)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->move_file_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **source** | **String** | Source file or directory path |  |
| **destination** | **String** | Destination file or directory path |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## replace_in_files

> <Array<ReplaceResult>> replace_in_files(request)

Replace text in files

Replace text pattern with new value in multiple files

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::FileSystemApi.new
request = DaytonaToolboxApiClient::ReplaceRequest.new({files: ['files_example'], new_value: 'new_value_example', pattern: 'pattern_example'}) # ReplaceRequest | Replace request

begin
  # Replace text in files
  result = api_instance.replace_in_files(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->replace_in_files: #{e}"
end
```

#### Using the replace_in_files_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<ReplaceResult>>, Integer, Hash)> replace_in_files_with_http_info(request)

```ruby
begin
  # Replace text in files
  data, status_code, headers = api_instance.replace_in_files_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<ReplaceResult>>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->replace_in_files_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**ReplaceRequest**](ReplaceRequest.md) | Replace request |  |

### Return type

[**Array&lt;ReplaceResult&gt;**](ReplaceResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## search_files

> <SearchFilesResponse> search_files(path, pattern)

Search files by pattern

Search for files matching a specific pattern in a directory

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::FileSystemApi.new
path = 'path_example' # String | Directory path to search in
pattern = 'pattern_example' # String | File pattern to match (e.g., *.txt, *.go)

begin
  # Search files by pattern
  result = api_instance.search_files(path, pattern)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->search_files: #{e}"
end
```

#### Using the search_files_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SearchFilesResponse>, Integer, Hash)> search_files_with_http_info(path, pattern)

```ruby
begin
  # Search files by pattern
  data, status_code, headers = api_instance.search_files_with_http_info(path, pattern)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SearchFilesResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->search_files_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | Directory path to search in |  |
| **pattern** | **String** | File pattern to match (e.g., *.txt, *.go) |  |

### Return type

[**SearchFilesResponse**](SearchFilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## set_file_permissions

> set_file_permissions(path, opts)

Set file permissions

Set file permissions, ownership, and group for a file or directory

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::FileSystemApi.new
path = 'path_example' # String | File or directory path
opts = {
  owner: 'owner_example', # String | Owner (username or UID)
  group: 'group_example', # String | Group (group name or GID)
  mode: 'mode_example' # String | File mode in octal format (e.g., 0755)
}

begin
  # Set file permissions
  api_instance.set_file_permissions(path, opts)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->set_file_permissions: #{e}"
end
```

#### Using the set_file_permissions_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> set_file_permissions_with_http_info(path, opts)

```ruby
begin
  # Set file permissions
  data, status_code, headers = api_instance.set_file_permissions_with_http_info(path, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->set_file_permissions_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | File or directory path |  |
| **owner** | **String** | Owner (username or UID) | [optional] |
| **group** | **String** | Group (group name or GID) | [optional] |
| **mode** | **String** | File mode in octal format (e.g., 0755) | [optional] |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## upload_file

> Hash&lt;String, Object&gt; upload_file(path, file)

Upload a file

Upload a file to the specified path

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::FileSystemApi.new
path = 'path_example' # String | Destination path for the uploaded file
file = File.new('/path/to/some/file') # File | File to upload

begin
  # Upload a file
  result = api_instance.upload_file(path, file)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->upload_file: #{e}"
end
```

#### Using the upload_file_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(Hash&lt;String, Object&gt;, Integer, Hash)> upload_file_with_http_info(path, file)

```ruby
begin
  # Upload a file
  data, status_code, headers = api_instance.upload_file_with_http_info(path, file)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Hash&lt;String, Object&gt;
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->upload_file_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | Destination path for the uploaded file |  |
| **file** | **File** | File to upload |  |

### Return type

**Hash&lt;String, Object&gt;**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*


## upload_files

> upload_files

Upload multiple files

Upload multiple files with their destination paths

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::FileSystemApi.new

begin
  # Upload multiple files
  api_instance.upload_files
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->upload_files: #{e}"
end
```

#### Using the upload_files_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> upload_files_with_http_info

```ruby
begin
  # Upload multiple files
  data, status_code, headers = api_instance.upload_files_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling FileSystemApi->upload_files_with_http_info: #{e}"
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
- **Accept**: Not defined

