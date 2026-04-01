# SessionExecuteResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cmd_id** | **str** |  | 
**exit_code** | **int** |  | [optional] 
**output** | **str** |  | [optional] 
**stderr** | **str** |  | [optional] 
**stdout** | **str** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client_async.models.session_execute_response import SessionExecuteResponse

# TODO update the JSON string below
json = "{}"
# create an instance of SessionExecuteResponse from a JSON string
session_execute_response_instance = SessionExecuteResponse.from_json(json)
# print the JSON string representation of the object
print(SessionExecuteResponse.to_json())

# convert the object into a dict
session_execute_response_dict = session_execute_response_instance.to_dict()
# create an instance of SessionExecuteResponse from a dict
session_execute_response_from_dict = SessionExecuteResponse.from_dict(session_execute_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


