# MouseDragRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**button** | **str** |  | [optional] 
**end_x** | **int** |  | [optional] 
**end_y** | **int** |  | [optional] 
**start_x** | **int** |  | [optional] 
**start_y** | **int** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client.models.mouse_drag_request import MouseDragRequest

# TODO update the JSON string below
json = "{}"
# create an instance of MouseDragRequest from a JSON string
mouse_drag_request_instance = MouseDragRequest.from_json(json)
# print the JSON string representation of the object
print(MouseDragRequest.to_json())

# convert the object into a dict
mouse_drag_request_dict = mouse_drag_request_instance.to_dict()
# create an instance of MouseDragRequest from a dict
mouse_drag_request_from_dict = MouseDragRequest.from_dict(mouse_drag_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


