# ListRecordingsResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**recordings** | [**List[Recording]**](Recording.md) |  | 

## Example

```python
from daytona_toolbox_api_client_async.models.list_recordings_response import ListRecordingsResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ListRecordingsResponse from a JSON string
list_recordings_response_instance = ListRecordingsResponse.from_json(json)
# print the JSON string representation of the object
print(ListRecordingsResponse.to_json())

# convert the object into a dict
list_recordings_response_dict = list_recordings_response_instance.to_dict()
# create an instance of ListRecordingsResponse from a dict
list_recordings_response_from_dict = ListRecordingsResponse.from_dict(list_recordings_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


