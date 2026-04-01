# DaytonaToolboxApiClient::InfoApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get_user_home_dir**](InfoApi.md#get_user_home_dir) | **GET** /user-home-dir | Get user home directory |
| [**get_version**](InfoApi.md#get_version) | **GET** /version | Get version |
| [**get_work_dir**](InfoApi.md#get_work_dir) | **GET** /work-dir | Get working directory |


## get_user_home_dir

> <UserHomeDirResponse> get_user_home_dir

Get user home directory

Get the current user home directory path.

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::InfoApi.new

begin
  # Get user home directory
  result = api_instance.get_user_home_dir
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling InfoApi->get_user_home_dir: #{e}"
end
```

#### Using the get_user_home_dir_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<UserHomeDirResponse>, Integer, Hash)> get_user_home_dir_with_http_info

```ruby
begin
  # Get user home directory
  data, status_code, headers = api_instance.get_user_home_dir_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <UserHomeDirResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling InfoApi->get_user_home_dir_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**UserHomeDirResponse**](UserHomeDirResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_version

> Hash&lt;String, String&gt; get_version

Get version

Get the current daemon version

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::InfoApi.new

begin
  # Get version
  result = api_instance.get_version
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling InfoApi->get_version: #{e}"
end
```

#### Using the get_version_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(Hash&lt;String, String&gt;, Integer, Hash)> get_version_with_http_info

```ruby
begin
  # Get version
  data, status_code, headers = api_instance.get_version_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Hash&lt;String, String&gt;
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling InfoApi->get_version_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

**Hash&lt;String, String&gt;**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_work_dir

> <WorkDirResponse> get_work_dir

Get working directory

Get the current working directory path. This is default directory used for running commands.

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::InfoApi.new

begin
  # Get working directory
  result = api_instance.get_work_dir
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling InfoApi->get_work_dir: #{e}"
end
```

#### Using the get_work_dir_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<WorkDirResponse>, Integer, Hash)> get_work_dir_with_http_info

```ruby
begin
  # Get working directory
  data, status_code, headers = api_instance.get_work_dir_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <WorkDirResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling InfoApi->get_work_dir_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**WorkDirResponse**](WorkDirResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

