# DisplayInfoResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**displays** | [**List[DisplayInfo]**](DisplayInfo.md) |  | [optional] 

## Example

```python
from daytona_toolbox_api_client_async.models.display_info_response import DisplayInfoResponse

# TODO update the JSON string below
json = "{}"
# create an instance of DisplayInfoResponse from a JSON string
display_info_response_instance = DisplayInfoResponse.from_json(json)
# print the JSON string representation of the object
print(DisplayInfoResponse.to_json())

# convert the object into a dict
display_info_response_dict = display_info_response_instance.to_dict()
# create an instance of DisplayInfoResponse from a dict
display_info_response_from_dict = DisplayInfoResponse.from_dict(display_info_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


