# MouseClickResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**x** | **int** |  | [optional] 
**y** | **int** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client.models.mouse_click_response import MouseClickResponse

# TODO update the JSON string below
json = "{}"
# create an instance of MouseClickResponse from a JSON string
mouse_click_response_instance = MouseClickResponse.from_json(json)
# print the JSON string representation of the object
print(MouseClickResponse.to_json())

# convert the object into a dict
mouse_click_response_dict = mouse_click_response_instance.to_dict()
# create an instance of MouseClickResponse from a dict
mouse_click_response_from_dict = MouseClickResponse.from_dict(mouse_click_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


