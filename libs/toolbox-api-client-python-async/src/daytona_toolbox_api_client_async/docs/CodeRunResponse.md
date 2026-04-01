# CodeRunResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**artifacts** | [**CodeRunArtifacts**](CodeRunArtifacts.md) |  | [optional] 
**exit_code** | **int** |  | [optional] 
**result** | **str** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client_async.models.code_run_response import CodeRunResponse

# TODO update the JSON string below
json = "{}"
# create an instance of CodeRunResponse from a JSON string
code_run_response_instance = CodeRunResponse.from_json(json)
# print the JSON string representation of the object
print(CodeRunResponse.to_json())

# convert the object into a dict
code_run_response_dict = code_run_response_instance.to_dict()
# create an instance of CodeRunResponse from a dict
code_run_response_from_dict = CodeRunResponse.from_dict(code_run_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


