# MouseClickRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**button** | **str** | left, right, middle | [optional] 
**double** | **bool** |  | [optional] 
**x** | **int** |  | [optional] 
**y** | **int** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client_async.models.mouse_click_request import MouseClickRequest

# TODO update the JSON string below
json = "{}"
# create an instance of MouseClickRequest from a JSON string
mouse_click_request_instance = MouseClickRequest.from_json(json)
# print the JSON string representation of the object
print(MouseClickRequest.to_json())

# convert the object into a dict
mouse_click_request_dict = mouse_click_request_instance.to_dict()
# create an instance of MouseClickRequest from a dict
mouse_click_request_from_dict = MouseClickRequest.from_dict(mouse_click_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


