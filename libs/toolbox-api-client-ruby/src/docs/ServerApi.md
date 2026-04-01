# DaytonaToolboxApiClient::ServerApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**initialize**](ServerApi.md#initialize) | **POST** /init | Initialize toolbox server |


## initialize

> Hash&lt;String, String&gt; initialize(request)

Initialize toolbox server

Set the auth token and initialize telemetry for the toolbox server

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::ServerApi.new
request = DaytonaToolboxApiClient::InitializeRequest.new({token: 'token_example'}) # InitializeRequest | Initialization request

begin
  # Initialize toolbox server
  result = api_instance.initialize(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ServerApi->initialize: #{e}"
end
```

#### Using the initialize_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(Hash&lt;String, String&gt;, Integer, Hash)> initialize_with_http_info(request)

```ruby
begin
  # Initialize toolbox server
  data, status_code, headers = api_instance.initialize_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Hash&lt;String, String&gt;
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling ServerApi->initialize_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**InitializeRequest**](InitializeRequest.md) | Initialization request |  |

### Return type

**Hash&lt;String, String&gt;**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

