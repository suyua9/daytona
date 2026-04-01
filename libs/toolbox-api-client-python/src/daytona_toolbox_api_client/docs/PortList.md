# PortList


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ports** | **List[int]** |  | [optional] 

## Example

```python
from daytona_toolbox_api_client.models.port_list import PortList

# TODO update the JSON string below
json = "{}"
# create an instance of PortList from a JSON string
port_list_instance = PortList.from_json(json)
# print the JSON string representation of the object
print(PortList.to_json())

# convert the object into a dict
port_list_dict = port_list_instance.to_dict()
# create an instance of PortList from a dict
port_list_from_dict = PortList.from_dict(port_list_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


