# PtySessionInfo


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**active** | **bool** |  | 
**cols** | **int** |  | 
**created_at** | **str** |  | 
**cwd** | **str** |  | 
**envs** | **Dict[str, str]** |  | 
**id** | **str** |  | 
**lazy_start** | **bool** | Whether this session uses lazy start | 
**rows** | **int** |  | 

## Example

```python
from daytona_toolbox_api_client.models.pty_session_info import PtySessionInfo

# TODO update the JSON string below
json = "{}"
# create an instance of PtySessionInfo from a JSON string
pty_session_info_instance = PtySessionInfo.from_json(json)
# print the JSON string representation of the object
print(PtySessionInfo.to_json())

# convert the object into a dict
pty_session_info_dict = pty_session_info_instance.to_dict()
# create an instance of PtySessionInfo from a dict
pty_session_info_from_dict = PtySessionInfo.from_dict(pty_session_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


