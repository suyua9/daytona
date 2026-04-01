# MousePositionResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**x** | **int** |  | [optional] 
**y** | **int** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client.models.mouse_position_response import MousePositionResponse

# TODO update the JSON string below
json = "{}"
# create an instance of MousePositionResponse from a JSON string
mouse_position_response_instance = MousePositionResponse.from_json(json)
# print the JSON string representation of the object
print(MousePositionResponse.to_json())

# convert the object into a dict
mouse_position_response_dict = mouse_position_response_instance.to_dict()
# create an instance of MousePositionResponse from a dict
mouse_position_response_from_dict = MousePositionResponse.from_dict(mouse_position_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


