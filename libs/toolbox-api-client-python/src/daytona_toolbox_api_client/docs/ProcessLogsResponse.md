# ProcessLogsResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**logs** | **str** |  | [optional] 
**process_name** | **str** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client.models.process_logs_response import ProcessLogsResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessLogsResponse from a JSON string
process_logs_response_instance = ProcessLogsResponse.from_json(json)
# print the JSON string representation of the object
print(ProcessLogsResponse.to_json())

# convert the object into a dict
process_logs_response_dict = process_logs_response_instance.to_dict()
# create an instance of ProcessLogsResponse from a dict
process_logs_response_from_dict = ProcessLogsResponse.from_dict(process_logs_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


