# DaytonaToolboxApiClient::PortApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get_ports**](PortApi.md#get_ports) | **GET** /port | Get active ports |
| [**is_port_in_use**](PortApi.md#is_port_in_use) | **GET** /port/{port}/in-use | Check if port is in use |


## get_ports

> <PortList> get_ports

Get active ports

Get a list of all currently active ports

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::PortApi.new

begin
  # Get active ports
  result = api_instance.get_ports
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling PortApi->get_ports: #{e}"
end
```

#### Using the get_ports_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<PortList>, Integer, Hash)> get_ports_with_http_info

```ruby
begin
  # Get active ports
  data, status_code, headers = api_instance.get_ports_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <PortList>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling PortApi->get_ports_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**PortList**](PortList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## is_port_in_use

> <IsPortInUseResponse> is_port_in_use(port)

Check if port is in use

Check if a specific port is currently in use

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::PortApi.new
port = 56 # Integer | Port number (3000-9999)

begin
  # Check if port is in use
  result = api_instance.is_port_in_use(port)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling PortApi->is_port_in_use: #{e}"
end
```

#### Using the is_port_in_use_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<IsPortInUseResponse>, Integer, Hash)> is_port_in_use_with_http_info(port)

```ruby
begin
  # Check if port is in use
  data, status_code, headers = api_instance.is_port_in_use_with_http_info(port)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <IsPortInUseResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling PortApi->is_port_in_use_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **port** | **Integer** | Port number (3000-9999) |  |

### Return type

[**IsPortInUseResponse**](IsPortInUseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

