# MouseMoveRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**x** | **int** |  | [optional] 
**y** | **int** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client.models.mouse_move_request import MouseMoveRequest

# TODO update the JSON string below
json = "{}"
# create an instance of MouseMoveRequest from a JSON string
mouse_move_request_instance = MouseMoveRequest.from_json(json)
# print the JSON string representation of the object
print(MouseMoveRequest.to_json())

# convert the object into a dict
mouse_move_request_dict = mouse_move_request_instance.to_dict()
# create an instance of MouseMoveRequest from a dict
mouse_move_request_from_dict = MouseMoveRequest.from_dict(mouse_move_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


