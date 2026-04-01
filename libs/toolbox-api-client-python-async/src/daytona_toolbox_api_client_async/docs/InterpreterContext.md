# InterpreterContext


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**active** | **bool** |  | 
**created_at** | **str** |  | 
**cwd** | **str** |  | 
**id** | **str** |  | 
**language** | **str** |  | 

## Example

```python
from daytona_toolbox_api_client_async.models.interpreter_context import InterpreterContext

# TODO update the JSON string below
json = "{}"
# create an instance of InterpreterContext from a JSON string
interpreter_context_instance = InterpreterContext.from_json(json)
# print the JSON string representation of the object
print(InterpreterContext.to_json())

# convert the object into a dict
interpreter_context_dict = interpreter_context_instance.to_dict()
# create an instance of InterpreterContext from a dict
interpreter_context_from_dict = InterpreterContext.from_dict(interpreter_context_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


