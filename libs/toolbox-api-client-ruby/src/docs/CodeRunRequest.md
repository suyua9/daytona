# DaytonaToolboxApiClient::CodeRunRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **argv** | **Array&lt;String&gt;** |  | [optional] |
| **code** | **String** |  |  |
| **envs** | **Hash&lt;String, String&gt;** |  | [optional] |
| **language** | **String** | python, javascript, typescript |  |
| **timeout** | **Integer** |  | [optional] |

## Example

```ruby
require 'daytona_toolbox_api_client'

instance = DaytonaToolboxApiClient::CodeRunRequest.new(
  argv: null,
  code: null,
  envs: null,
  language: null,
  timeout: null
)
```

