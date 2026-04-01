# DisplayInfo


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**height** | **int** |  | [optional] 
**id** | **int** |  | [optional] 
**is_active** | **bool** |  | [optional] 
**width** | **int** |  | [optional] 
**x** | **int** |  | [optional] 
**y** | **int** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client_async.models.display_info import DisplayInfo

# TODO update the JSON string below
json = "{}"
# create an instance of DisplayInfo from a JSON string
display_info_instance = DisplayInfo.from_json(json)
# print the JSON string representation of the object
print(DisplayInfo.to_json())

# convert the object into a dict
display_info_dict = display_info_instance.to_dict()
# create an instance of DisplayInfo from a dict
display_info_from_dict = DisplayInfo.from_dict(display_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


