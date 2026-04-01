# ProcessStatus


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auto_restart** | **bool** |  | [optional] 
**pid** | **int** |  | [optional] 
**priority** | **int** |  | [optional] 
**running** | **bool** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client.models.process_status import ProcessStatus

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessStatus from a JSON string
process_status_instance = ProcessStatus.from_json(json)
# print the JSON string representation of the object
print(ProcessStatus.to_json())

# convert the object into a dict
process_status_dict = process_status_instance.to_dict()
# create an instance of ProcessStatus from a dict
process_status_from_dict = ProcessStatus.from_dict(process_status_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


