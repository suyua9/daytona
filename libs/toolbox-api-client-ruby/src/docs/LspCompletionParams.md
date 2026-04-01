# DaytonaToolboxApiClient::LspCompletionParams

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **context** | [**CompletionContext**](CompletionContext.md) |  | [optional] |
| **language_id** | **String** |  |  |
| **path_to_project** | **String** |  |  |
| **position** | [**LspPosition**](LspPosition.md) |  |  |
| **uri** | **String** |  |  |

## Example

```ruby
require 'daytona_toolbox_api_client'

instance = DaytonaToolboxApiClient::LspCompletionParams.new(
  context: null,
  language_id: null,
  path_to_project: null,
  position: null,
  uri: null
)
```

