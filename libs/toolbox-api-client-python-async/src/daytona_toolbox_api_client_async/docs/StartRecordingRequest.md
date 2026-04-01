# StartRecordingRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**label** | **str** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client_async.models.start_recording_request import StartRecordingRequest

# TODO update the JSON string below
json = "{}"
# create an instance of StartRecordingRequest from a JSON string
start_recording_request_instance = StartRecordingRequest.from_json(json)
# print the JSON string representation of the object
print(StartRecordingRequest.to_json())

# convert the object into a dict
start_recording_request_dict = start_recording_request_instance.to_dict()
# create an instance of StartRecordingRequest from a dict
start_recording_request_from_dict = StartRecordingRequest.from_dict(start_recording_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


