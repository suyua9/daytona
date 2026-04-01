# ComputerUseStatusResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**status** | **str** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client_async.models.computer_use_status_response import ComputerUseStatusResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ComputerUseStatusResponse from a JSON string
computer_use_status_response_instance = ComputerUseStatusResponse.from_json(json)
# print the JSON string representation of the object
print(ComputerUseStatusResponse.to_json())

# convert the object into a dict
computer_use_status_response_dict = computer_use_status_response_instance.to_dict()
# create an instance of ComputerUseStatusResponse from a dict
computer_use_status_response_from_dict = ComputerUseStatusResponse.from_dict(computer_use_status_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


