# DaytonaToolboxApiClient::GitStatus

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ahead** | **Integer** |  | [optional] |
| **behind** | **Integer** |  | [optional] |
| **branch_published** | **Boolean** |  | [optional] |
| **current_branch** | **String** |  |  |
| **file_status** | [**Array&lt;FileStatus&gt;**](FileStatus.md) |  |  |

## Example

```ruby
require 'daytona_toolbox_api_client'

instance = DaytonaToolboxApiClient::GitStatus.new(
  ahead: null,
  behind: null,
  branch_published: null,
  current_branch: null,
  file_status: null
)
```

