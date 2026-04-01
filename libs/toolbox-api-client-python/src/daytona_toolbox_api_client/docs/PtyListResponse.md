# PtyListResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sessions** | [**List[PtySessionInfo]**](PtySessionInfo.md) |  | 

## Example

```python
from daytona_toolbox_api_client.models.pty_list_response import PtyListResponse

# TODO update the JSON string below
json = "{}"
# create an instance of PtyListResponse from a JSON string
pty_list_response_instance = PtyListResponse.from_json(json)
# print the JSON string representation of the object
print(PtyListResponse.to_json())

# convert the object into a dict
pty_list_response_dict = pty_list_response_instance.to_dict()
# create an instance of PtyListResponse from a dict
pty_list_response_from_dict = PtyListResponse.from_dict(pty_list_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


