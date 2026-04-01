# KeyboardTypeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**delay** | **int** | milliseconds between keystrokes | [optional] 
**text** | **str** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client.models.keyboard_type_request import KeyboardTypeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of KeyboardTypeRequest from a JSON string
keyboard_type_request_instance = KeyboardTypeRequest.from_json(json)
# print the JSON string representation of the object
print(KeyboardTypeRequest.to_json())

# convert the object into a dict
keyboard_type_request_dict = keyboard_type_request_instance.to_dict()
# create an instance of KeyboardTypeRequest from a dict
keyboard_type_request_from_dict = KeyboardTypeRequest.from_dict(keyboard_type_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


