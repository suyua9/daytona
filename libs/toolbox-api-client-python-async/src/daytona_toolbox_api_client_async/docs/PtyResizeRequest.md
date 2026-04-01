# PtyResizeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cols** | **int** |  | 
**rows** | **int** |  | 

## Example

```python
from daytona_toolbox_api_client_async.models.pty_resize_request import PtyResizeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of PtyResizeRequest from a JSON string
pty_resize_request_instance = PtyResizeRequest.from_json(json)
# print the JSON string representation of the object
print(PtyResizeRequest.to_json())

# convert the object into a dict
pty_resize_request_dict = pty_resize_request_instance.to_dict()
# create an instance of PtyResizeRequest from a dict
pty_resize_request_from_dict = PtyResizeRequest.from_dict(pty_resize_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


