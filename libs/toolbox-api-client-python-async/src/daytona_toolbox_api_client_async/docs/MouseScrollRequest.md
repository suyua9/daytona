# MouseScrollRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**amount** | **int** |  | [optional] 
**direction** | **str** | up, down | [optional] 
**x** | **int** |  | [optional] 
**y** | **int** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client_async.models.mouse_scroll_request import MouseScrollRequest

# TODO update the JSON string below
json = "{}"
# create an instance of MouseScrollRequest from a JSON string
mouse_scroll_request_instance = MouseScrollRequest.from_json(json)
# print the JSON string representation of the object
print(MouseScrollRequest.to_json())

# convert the object into a dict
mouse_scroll_request_dict = mouse_scroll_request_instance.to_dict()
# create an instance of MouseScrollRequest from a dict
mouse_scroll_request_from_dict = MouseScrollRequest.from_dict(mouse_scroll_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


