# StopRecordingRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 

## Example

```python
from daytona_toolbox_api_client.models.stop_recording_request import StopRecordingRequest

# TODO update the JSON string below
json = "{}"
# create an instance of StopRecordingRequest from a JSON string
stop_recording_request_instance = StopRecordingRequest.from_json(json)
# print the JSON string representation of the object
print(StopRecordingRequest.to_json())

# convert the object into a dict
stop_recording_request_dict = stop_recording_request_instance.to_dict()
# create an instance of StopRecordingRequest from a dict
stop_recording_request_from_dict = StopRecordingRequest.from_dict(stop_recording_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


