# DaytonaToolboxApiClient::ExecuteRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **command** | **String** |  |  |
| **cwd** | **String** | Current working directory | [optional] |
| **envs** | **Hash&lt;String, String&gt;** | Environment variables to set for the command | [optional] |
| **timeout** | **Integer** | Timeout in seconds, defaults to 10 seconds | [optional] |

## Example

```ruby
require 'daytona_toolbox_api_client'

instance = DaytonaToolboxApiClient::ExecuteRequest.new(
  command: null,
  cwd: null,
  envs: null,
  timeout: null
)
```

