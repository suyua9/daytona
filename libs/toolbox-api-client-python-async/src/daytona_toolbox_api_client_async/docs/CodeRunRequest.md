# CodeRunRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**argv** | **List[str]** |  | [optional] 
**code** | **str** |  | 
**envs** | **Dict[str, str]** |  | [optional] 
**language** | **str** | python, javascript, typescript | 
**timeout** | **int** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client_async.models.code_run_request import CodeRunRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CodeRunRequest from a JSON string
code_run_request_instance = CodeRunRequest.from_json(json)
# print the JSON string representation of the object
print(CodeRunRequest.to_json())

# convert the object into a dict
code_run_request_dict = code_run_request_instance.to_dict()
# create an instance of CodeRunRequest from a dict
code_run_request_from_dict = CodeRunRequest.from_dict(code_run_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


