# KeyboardHotkeyRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**keys** | **str** | e.g., \&quot;ctrl+c\&quot;, \&quot;cmd+v\&quot; | [optional] 

## Example

```python
from daytona_toolbox_api_client.models.keyboard_hotkey_request import KeyboardHotkeyRequest

# TODO update the JSON string below
json = "{}"
# create an instance of KeyboardHotkeyRequest from a JSON string
keyboard_hotkey_request_instance = KeyboardHotkeyRequest.from_json(json)
# print the JSON string representation of the object
print(KeyboardHotkeyRequest.to_json())

# convert the object into a dict
keyboard_hotkey_request_dict = keyboard_hotkey_request_instance.to_dict()
# create an instance of KeyboardHotkeyRequest from a dict
keyboard_hotkey_request_from_dict = KeyboardHotkeyRequest.from_dict(keyboard_hotkey_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


