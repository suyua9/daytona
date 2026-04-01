# Recording


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**duration_seconds** | **float** |  | [optional] 
**end_time** | **str** |  | [optional] 
**file_name** | **str** |  | 
**file_path** | **str** |  | 
**id** | **str** |  | 
**size_bytes** | **int** |  | [optional] 
**start_time** | **str** |  | 
**status** | **str** |  | 

## Example

```python
from daytona_toolbox_api_client.models.recording import Recording

# TODO update the JSON string below
json = "{}"
# create an instance of Recording from a JSON string
recording_instance = Recording.from_json(json)
# print the JSON string representation of the object
print(Recording.to_json())

# convert the object into a dict
recording_dict = recording_instance.to_dict()
# create an instance of Recording from a dict
recording_from_dict = Recording.from_dict(recording_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


