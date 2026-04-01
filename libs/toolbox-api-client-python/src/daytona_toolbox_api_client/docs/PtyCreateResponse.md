# PtyCreateResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**session_id** | **str** |  | 

## Example

```python
from daytona_toolbox_api_client.models.pty_create_response import PtyCreateResponse

# TODO update the JSON string below
json = "{}"
# create an instance of PtyCreateResponse from a JSON string
pty_create_response_instance = PtyCreateResponse.from_json(json)
# print the JSON string representation of the object
print(PtyCreateResponse.to_json())

# convert the object into a dict
pty_create_response_dict = pty_create_response_instance.to_dict()
# create an instance of PtyCreateResponse from a dict
pty_create_response_from_dict = PtyCreateResponse.from_dict(pty_create_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


