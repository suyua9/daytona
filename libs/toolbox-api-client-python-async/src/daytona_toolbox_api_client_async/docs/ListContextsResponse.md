# ListContextsResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**contexts** | [**List[InterpreterContext]**](InterpreterContext.md) |  | 

## Example

```python
from daytona_toolbox_api_client_async.models.list_contexts_response import ListContextsResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ListContextsResponse from a JSON string
list_contexts_response_instance = ListContextsResponse.from_json(json)
# print the JSON string representation of the object
print(ListContextsResponse.to_json())

# convert the object into a dict
list_contexts_response_dict = list_contexts_response_instance.to_dict()
# create an instance of ListContextsResponse from a dict
list_contexts_response_from_dict = ListContextsResponse.from_dict(list_contexts_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


