# MouseDragResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**x** | **int** |  | [optional] 
**y** | **int** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client.models.mouse_drag_response import MouseDragResponse

# TODO update the JSON string below
json = "{}"
# create an instance of MouseDragResponse from a JSON string
mouse_drag_response_instance = MouseDragResponse.from_json(json)
# print the JSON string representation of the object
print(MouseDragResponse.to_json())

# convert the object into a dict
mouse_drag_response_dict = mouse_drag_response_instance.to_dict()
# create an instance of MouseDragResponse from a dict
mouse_drag_response_from_dict = MouseDragResponse.from_dict(mouse_drag_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


