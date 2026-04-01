# FilesDownloadRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**paths** | **List[str]** |  | 

## Example

```python
from daytona_toolbox_api_client.models.files_download_request import FilesDownloadRequest

# TODO update the JSON string below
json = "{}"
# create an instance of FilesDownloadRequest from a JSON string
files_download_request_instance = FilesDownloadRequest.from_json(json)
# print the JSON string representation of the object
print(FilesDownloadRequest.to_json())

# convert the object into a dict
files_download_request_dict = files_download_request_instance.to_dict()
# create an instance of FilesDownloadRequest from a dict
files_download_request_from_dict = FilesDownloadRequest.from_dict(files_download_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


