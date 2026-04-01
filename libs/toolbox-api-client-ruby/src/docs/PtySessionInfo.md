# DaytonaToolboxApiClient::PtySessionInfo

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **active** | **Boolean** |  |  |
| **cols** | **Integer** |  |  |
| **created_at** | **String** |  |  |
| **cwd** | **String** |  |  |
| **envs** | **Hash&lt;String, String&gt;** |  |  |
| **id** | **String** |  |  |
| **lazy_start** | **Boolean** | Whether this session uses lazy start |  |
| **rows** | **Integer** |  |  |

## Example

```ruby
require 'daytona_toolbox_api_client'

instance = DaytonaToolboxApiClient::PtySessionInfo.new(
  active: null,
  cols: null,
  created_at: null,
  cwd: null,
  envs: null,
  id: null,
  lazy_start: null,
  rows: null
)
```

