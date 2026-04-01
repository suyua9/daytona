# DaytonaToolboxApiClient::PtyCreateRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cols** | **Integer** |  | [optional] |
| **cwd** | **String** |  | [optional] |
| **envs** | **Hash&lt;String, String&gt;** |  | [optional] |
| **id** | **String** |  | [optional] |
| **lazy_start** | **Boolean** | Don&#39;t start PTY until first client connects | [optional] |
| **rows** | **Integer** |  | [optional] |

## Example

```ruby
require 'daytona_toolbox_api_client'

instance = DaytonaToolboxApiClient::PtyCreateRequest.new(
  cols: null,
  cwd: null,
  envs: null,
  id: null,
  lazy_start: null,
  rows: null
)
```

