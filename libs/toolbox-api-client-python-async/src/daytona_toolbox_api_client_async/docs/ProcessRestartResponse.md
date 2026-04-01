# ProcessRestartResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message** | **str** |  | [optional] 
**process_name** | **str** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client_async.models.process_restart_response import ProcessRestartResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessRestartResponse from a JSON string
process_restart_response_instance = ProcessRestartResponse.from_json(json)
# print the JSON string representation of the object
print(ProcessRestartResponse.to_json())

# convert the object into a dict
process_restart_response_dict = process_restart_response_instance.to_dict()
# create an instance of ProcessRestartResponse from a dict
process_restart_response_from_dict = ProcessRestartResponse.from_dict(process_restart_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


