# ProcessErrorsResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**errors** | **str** |  | [optional] 
**process_name** | **str** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client.models.process_errors_response import ProcessErrorsResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessErrorsResponse from a JSON string
process_errors_response_instance = ProcessErrorsResponse.from_json(json)
# print the JSON string representation of the object
print(ProcessErrorsResponse.to_json())

# convert the object into a dict
process_errors_response_dict = process_errors_response_instance.to_dict()
# create an instance of ProcessErrorsResponse from a dict
process_errors_response_from_dict = ProcessErrorsResponse.from_dict(process_errors_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


