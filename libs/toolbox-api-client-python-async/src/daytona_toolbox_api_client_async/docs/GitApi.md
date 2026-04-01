# daytona_toolbox_api_client_async.GitApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_files**](GitApi.md#add_files) | **POST** /git/add | Add files to Git staging
[**checkout_branch**](GitApi.md#checkout_branch) | **POST** /git/checkout | Checkout branch or commit
[**clone_repository**](GitApi.md#clone_repository) | **POST** /git/clone | Clone a Git repository
[**commit_changes**](GitApi.md#commit_changes) | **POST** /git/commit | Commit changes
[**create_branch**](GitApi.md#create_branch) | **POST** /git/branches | Create a new branch
[**delete_branch**](GitApi.md#delete_branch) | **DELETE** /git/branches | Delete a branch
[**get_commit_history**](GitApi.md#get_commit_history) | **GET** /git/history | Get commit history
[**get_status**](GitApi.md#get_status) | **GET** /git/status | Get Git status
[**list_branches**](GitApi.md#list_branches) | **GET** /git/branches | List branches
[**pull_changes**](GitApi.md#pull_changes) | **POST** /git/pull | Pull changes from remote
[**push_changes**](GitApi.md#push_changes) | **POST** /git/push | Push changes to remote


# **add_files**
> add_files(request)

Add files to Git staging

Add files to the Git staging area

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.git_add_request import GitAddRequest
from daytona_toolbox_api_client_async.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client_async.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
async with daytona_toolbox_api_client_async.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client_async.GitApi(api_client)
    request = daytona_toolbox_api_client_async.GitAddRequest() # GitAddRequest | Add files request

    try:
        # Add files to Git staging
        await api_instance.add_files(request)
    except Exception as e:
        print("Exception when calling GitApi->add_files: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**GitAddRequest**](GitAddRequest.md)| Add files request | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **checkout_branch**
> checkout_branch(request)

Checkout branch or commit

Switch to a different branch or commit in the Git repository

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.git_checkout_request import GitCheckoutRequest
from daytona_toolbox_api_client_async.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client_async.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
async with daytona_toolbox_api_client_async.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client_async.GitApi(api_client)
    request = daytona_toolbox_api_client_async.GitCheckoutRequest() # GitCheckoutRequest | Checkout request

    try:
        # Checkout branch or commit
        await api_instance.checkout_branch(request)
    except Exception as e:
        print("Exception when calling GitApi->checkout_branch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**GitCheckoutRequest**](GitCheckoutRequest.md)| Checkout request | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **clone_repository**
> clone_repository(request)

Clone a Git repository

Clone a Git repository to the specified path

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.git_clone_request import GitCloneRequest
from daytona_toolbox_api_client_async.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client_async.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
async with daytona_toolbox_api_client_async.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client_async.GitApi(api_client)
    request = daytona_toolbox_api_client_async.GitCloneRequest() # GitCloneRequest | Clone repository request

    try:
        # Clone a Git repository
        await api_instance.clone_repository(request)
    except Exception as e:
        print("Exception when calling GitApi->clone_repository: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**GitCloneRequest**](GitCloneRequest.md)| Clone repository request | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **commit_changes**
> GitCommitResponse commit_changes(request)

Commit changes

Commit staged changes to the Git repository

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.git_commit_request import GitCommitRequest
from daytona_toolbox_api_client_async.models.git_commit_response import GitCommitResponse
from daytona_toolbox_api_client_async.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client_async.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
async with daytona_toolbox_api_client_async.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client_async.GitApi(api_client)
    request = daytona_toolbox_api_client_async.GitCommitRequest() # GitCommitRequest | Commit request

    try:
        # Commit changes
        api_response = await api_instance.commit_changes(request)
        print("The response of GitApi->commit_changes:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GitApi->commit_changes: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**GitCommitRequest**](GitCommitRequest.md)| Commit request | 

### Return type

[**GitCommitResponse**](GitCommitResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_branch**
> create_branch(request)

Create a new branch

Create a new branch in the Git repository

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.git_branch_request import GitBranchRequest
from daytona_toolbox_api_client_async.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client_async.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
async with daytona_toolbox_api_client_async.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client_async.GitApi(api_client)
    request = daytona_toolbox_api_client_async.GitBranchRequest() # GitBranchRequest | Create branch request

    try:
        # Create a new branch
        await api_instance.create_branch(request)
    except Exception as e:
        print("Exception when calling GitApi->create_branch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**GitBranchRequest**](GitBranchRequest.md)| Create branch request | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Created |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_branch**
> delete_branch(request)

Delete a branch

Delete a branch from the Git repository

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.pkg_toolbox_git_git_delete_branch_request import PkgToolboxGitGitDeleteBranchRequest
from daytona_toolbox_api_client_async.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client_async.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
async with daytona_toolbox_api_client_async.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client_async.GitApi(api_client)
    request = daytona_toolbox_api_client_async.PkgToolboxGitGitDeleteBranchRequest() # PkgToolboxGitGitDeleteBranchRequest | Delete branch request

    try:
        # Delete a branch
        await api_instance.delete_branch(request)
    except Exception as e:
        print("Exception when calling GitApi->delete_branch: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PkgToolboxGitGitDeleteBranchRequest**](PkgToolboxGitGitDeleteBranchRequest.md)| Delete branch request | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | No Content |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_commit_history**
> List[GitCommitInfo] get_commit_history(path)

Get commit history

Get the commit history of the Git repository

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.git_commit_info import GitCommitInfo
from daytona_toolbox_api_client_async.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client_async.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
async with daytona_toolbox_api_client_async.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client_async.GitApi(api_client)
    path = 'path_example' # str | Repository path

    try:
        # Get commit history
        api_response = await api_instance.get_commit_history(path)
        print("The response of GitApi->get_commit_history:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GitApi->get_commit_history: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| Repository path | 

### Return type

[**List[GitCommitInfo]**](GitCommitInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_status**
> GitStatus get_status(path)

Get Git status

Get the Git status of the repository at the specified path

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.git_status import GitStatus
from daytona_toolbox_api_client_async.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client_async.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
async with daytona_toolbox_api_client_async.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client_async.GitApi(api_client)
    path = 'path_example' # str | Repository path

    try:
        # Get Git status
        api_response = await api_instance.get_status(path)
        print("The response of GitApi->get_status:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GitApi->get_status: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| Repository path | 

### Return type

[**GitStatus**](GitStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_branches**
> ListBranchResponse list_branches(path)

List branches

Get a list of all branches in the Git repository

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.list_branch_response import ListBranchResponse
from daytona_toolbox_api_client_async.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client_async.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
async with daytona_toolbox_api_client_async.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client_async.GitApi(api_client)
    path = 'path_example' # str | Repository path

    try:
        # List branches
        api_response = await api_instance.list_branches(path)
        print("The response of GitApi->list_branches:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GitApi->list_branches: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| Repository path | 

### Return type

[**ListBranchResponse**](ListBranchResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **pull_changes**
> pull_changes(request)

Pull changes from remote

Pull changes from the remote Git repository

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.git_repo_request import GitRepoRequest
from daytona_toolbox_api_client_async.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client_async.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
async with daytona_toolbox_api_client_async.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client_async.GitApi(api_client)
    request = daytona_toolbox_api_client_async.GitRepoRequest() # GitRepoRequest | Pull request

    try:
        # Pull changes from remote
        await api_instance.pull_changes(request)
    except Exception as e:
        print("Exception when calling GitApi->pull_changes: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**GitRepoRequest**](GitRepoRequest.md)| Pull request | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **push_changes**
> push_changes(request)

Push changes to remote

Push local changes to the remote Git repository

### Example


```python
import daytona_toolbox_api_client_async
from daytona_toolbox_api_client_async.models.git_repo_request import GitRepoRequest
from daytona_toolbox_api_client_async.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = daytona_toolbox_api_client_async.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
async with daytona_toolbox_api_client_async.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = daytona_toolbox_api_client_async.GitApi(api_client)
    request = daytona_toolbox_api_client_async.GitRepoRequest() # GitRepoRequest | Push request

    try:
        # Push changes to remote
        await api_instance.push_changes(request)
    except Exception as e:
        print("Exception when calling GitApi->push_changes: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**GitRepoRequest**](GitRepoRequest.md)| Push request | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

