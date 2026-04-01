# WindowsResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**windows** | [**List[WindowInfo]**](WindowInfo.md) |  | [optional] 

## Example

```python
from daytona_toolbox_api_client_async.models.windows_response import WindowsResponse

# TODO update the JSON string below
json = "{}"
# create an instance of WindowsResponse from a JSON string
windows_response_instance = WindowsResponse.from_json(json)
# print the JSON string representation of the object
print(WindowsResponse.to_json())

# convert the object into a dict
windows_response_dict = windows_response_instance.to_dict()
# create an instance of WindowsResponse from a dict
windows_response_from_dict = WindowsResponse.from_dict(windows_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


