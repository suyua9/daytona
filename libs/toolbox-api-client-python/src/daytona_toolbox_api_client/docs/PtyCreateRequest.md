# PtyCreateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cols** | **int** |  | [optional] 
**cwd** | **str** |  | [optional] 
**envs** | **Dict[str, str]** |  | [optional] 
**id** | **str** |  | [optional] 
**lazy_start** | **bool** | Don&#39;t start PTY until first client connects | [optional] 
**rows** | **int** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client.models.pty_create_request import PtyCreateRequest

# TODO update the JSON string below
json = "{}"
# create an instance of PtyCreateRequest from a JSON string
pty_create_request_instance = PtyCreateRequest.from_json(json)
# print the JSON string representation of the object
print(PtyCreateRequest.to_json())

# convert the object into a dict
pty_create_request_dict = pty_create_request_instance.to_dict()
# create an instance of PtyCreateRequest from a dict
pty_create_request_from_dict = PtyCreateRequest.from_dict(pty_create_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


