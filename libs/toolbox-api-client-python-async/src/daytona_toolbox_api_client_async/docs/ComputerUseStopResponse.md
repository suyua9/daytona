# ComputerUseStopResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message** | **str** |  | [optional] 
**status** | [**Dict[str, ProcessStatus]**](ProcessStatus.md) |  | [optional] 

## Example

```python
from daytona_toolbox_api_client_async.models.computer_use_stop_response import ComputerUseStopResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ComputerUseStopResponse from a JSON string
computer_use_stop_response_instance = ComputerUseStopResponse.from_json(json)
# print the JSON string representation of the object
print(ComputerUseStopResponse.to_json())

# convert the object into a dict
computer_use_stop_response_dict = computer_use_stop_response_instance.to_dict()
# create an instance of ComputerUseStopResponse from a dict
computer_use_stop_response_from_dict = ComputerUseStopResponse.from_dict(computer_use_stop_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


