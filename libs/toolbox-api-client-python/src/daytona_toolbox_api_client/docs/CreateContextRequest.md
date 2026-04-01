# CreateContextRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cwd** | **str** |  | [optional] 
**language** | **str** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client.models.create_context_request import CreateContextRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateContextRequest from a JSON string
create_context_request_instance = CreateContextRequest.from_json(json)
# print the JSON string representation of the object
print(CreateContextRequest.to_json())

# convert the object into a dict
create_context_request_dict = create_context_request_instance.to_dict()
# create an instance of CreateContextRequest from a dict
create_context_request_from_dict = CreateContextRequest.from_dict(create_context_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


