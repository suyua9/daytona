# DaytonaToolboxApiClient::ProcessStatus

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **auto_restart** | **Boolean** |  | [optional] |
| **pid** | **Integer** |  | [optional] |
| **priority** | **Integer** |  | [optional] |
| **running** | **Boolean** |  | [optional] |

## Example

```ruby
require 'daytona_toolbox_api_client'

instance = DaytonaToolboxApiClient::ProcessStatus.new(
  auto_restart: null,
  pid: null,
  priority: null,
  running: null
)
```

