# SessionSendInputRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data** | **str** |  | 

## Example

```python
from daytona_toolbox_api_client_async.models.session_send_input_request import SessionSendInputRequest

# TODO update the JSON string below
json = "{}"
# create an instance of SessionSendInputRequest from a JSON string
session_send_input_request_instance = SessionSendInputRequest.from_json(json)
# print the JSON string representation of the object
print(SessionSendInputRequest.to_json())

# convert the object into a dict
session_send_input_request_dict = session_send_input_request_instance.to_dict()
# create an instance of SessionSendInputRequest from a dict
session_send_input_request_from_dict = SessionSendInputRequest.from_dict(session_send_input_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


