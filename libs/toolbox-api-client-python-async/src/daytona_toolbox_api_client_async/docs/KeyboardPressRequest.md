# KeyboardPressRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**key** | **str** |  | [optional] 
**modifiers** | **List[str]** | ctrl, alt, shift, cmd | [optional] 

## Example

```python
from daytona_toolbox_api_client_async.models.keyboard_press_request import KeyboardPressRequest

# TODO update the JSON string below
json = "{}"
# create an instance of KeyboardPressRequest from a JSON string
keyboard_press_request_instance = KeyboardPressRequest.from_json(json)
# print the JSON string representation of the object
print(KeyboardPressRequest.to_json())

# convert the object into a dict
keyboard_press_request_dict = keyboard_press_request_instance.to_dict()
# create an instance of KeyboardPressRequest from a dict
keyboard_press_request_from_dict = KeyboardPressRequest.from_dict(keyboard_press_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


