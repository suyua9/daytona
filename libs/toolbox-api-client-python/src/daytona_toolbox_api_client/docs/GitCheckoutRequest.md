# GitCheckoutRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**branch** | **str** |  | 
**path** | **str** |  | 

## Example

```python
from daytona_toolbox_api_client.models.git_checkout_request import GitCheckoutRequest

# TODO update the JSON string below
json = "{}"
# create an instance of GitCheckoutRequest from a JSON string
git_checkout_request_instance = GitCheckoutRequest.from_json(json)
# print the JSON string representation of the object
print(GitCheckoutRequest.to_json())

# convert the object into a dict
git_checkout_request_dict = git_checkout_request_instance.to_dict()
# create an instance of GitCheckoutRequest from a dict
git_checkout_request_from_dict = GitCheckoutRequest.from_dict(git_checkout_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


