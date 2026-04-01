# WindowInfo


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**height** | **int** |  | [optional] 
**id** | **int** |  | [optional] 
**is_active** | **bool** |  | [optional] 
**title** | **str** |  | [optional] 
**width** | **int** |  | [optional] 
**x** | **int** |  | [optional] 
**y** | **int** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client.models.window_info import WindowInfo

# TODO update the JSON string below
json = "{}"
# create an instance of WindowInfo from a JSON string
window_info_instance = WindowInfo.from_json(json)
# print the JSON string representation of the object
print(WindowInfo.to_json())

# convert the object into a dict
window_info_dict = window_info_instance.to_dict()
# create an instance of WindowInfo from a dict
window_info_from_dict = WindowInfo.from_dict(window_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


