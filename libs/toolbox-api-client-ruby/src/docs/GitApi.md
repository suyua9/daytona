# DaytonaToolboxApiClient::GitApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**add_files**](GitApi.md#add_files) | **POST** /git/add | Add files to Git staging |
| [**checkout_branch**](GitApi.md#checkout_branch) | **POST** /git/checkout | Checkout branch or commit |
| [**clone_repository**](GitApi.md#clone_repository) | **POST** /git/clone | Clone a Git repository |
| [**commit_changes**](GitApi.md#commit_changes) | **POST** /git/commit | Commit changes |
| [**create_branch**](GitApi.md#create_branch) | **POST** /git/branches | Create a new branch |
| [**delete_branch**](GitApi.md#delete_branch) | **DELETE** /git/branches | Delete a branch |
| [**get_commit_history**](GitApi.md#get_commit_history) | **GET** /git/history | Get commit history |
| [**get_status**](GitApi.md#get_status) | **GET** /git/status | Get Git status |
| [**list_branches**](GitApi.md#list_branches) | **GET** /git/branches | List branches |
| [**pull_changes**](GitApi.md#pull_changes) | **POST** /git/pull | Pull changes from remote |
| [**push_changes**](GitApi.md#push_changes) | **POST** /git/push | Push changes to remote |


## add_files

> add_files(request)

Add files to Git staging

Add files to the Git staging area

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::GitApi.new
request = DaytonaToolboxApiClient::GitAddRequest.new({files: ['files_example'], path: 'path_example'}) # GitAddRequest | Add files request

begin
  # Add files to Git staging
  api_instance.add_files(request)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->add_files: #{e}"
end
```

#### Using the add_files_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> add_files_with_http_info(request)

```ruby
begin
  # Add files to Git staging
  data, status_code, headers = api_instance.add_files_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->add_files_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**GitAddRequest**](GitAddRequest.md) | Add files request |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## checkout_branch

> checkout_branch(request)

Checkout branch or commit

Switch to a different branch or commit in the Git repository

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::GitApi.new
request = DaytonaToolboxApiClient::GitCheckoutRequest.new({branch: 'branch_example', path: 'path_example'}) # GitCheckoutRequest | Checkout request

begin
  # Checkout branch or commit
  api_instance.checkout_branch(request)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->checkout_branch: #{e}"
end
```

#### Using the checkout_branch_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> checkout_branch_with_http_info(request)

```ruby
begin
  # Checkout branch or commit
  data, status_code, headers = api_instance.checkout_branch_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->checkout_branch_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**GitCheckoutRequest**](GitCheckoutRequest.md) | Checkout request |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## clone_repository

> clone_repository(request)

Clone a Git repository

Clone a Git repository to the specified path

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::GitApi.new
request = DaytonaToolboxApiClient::GitCloneRequest.new({path: 'path_example', url: 'url_example'}) # GitCloneRequest | Clone repository request

begin
  # Clone a Git repository
  api_instance.clone_repository(request)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->clone_repository: #{e}"
end
```

#### Using the clone_repository_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> clone_repository_with_http_info(request)

```ruby
begin
  # Clone a Git repository
  data, status_code, headers = api_instance.clone_repository_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->clone_repository_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**GitCloneRequest**](GitCloneRequest.md) | Clone repository request |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## commit_changes

> <GitCommitResponse> commit_changes(request)

Commit changes

Commit staged changes to the Git repository

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::GitApi.new
request = DaytonaToolboxApiClient::GitCommitRequest.new({author: 'author_example', email: 'email_example', message: 'message_example', path: 'path_example'}) # GitCommitRequest | Commit request

begin
  # Commit changes
  result = api_instance.commit_changes(request)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->commit_changes: #{e}"
end
```

#### Using the commit_changes_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<GitCommitResponse>, Integer, Hash)> commit_changes_with_http_info(request)

```ruby
begin
  # Commit changes
  data, status_code, headers = api_instance.commit_changes_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <GitCommitResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->commit_changes_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**GitCommitRequest**](GitCommitRequest.md) | Commit request |  |

### Return type

[**GitCommitResponse**](GitCommitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## create_branch

> create_branch(request)

Create a new branch

Create a new branch in the Git repository

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::GitApi.new
request = DaytonaToolboxApiClient::GitBranchRequest.new({name: 'name_example', path: 'path_example'}) # GitBranchRequest | Create branch request

begin
  # Create a new branch
  api_instance.create_branch(request)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->create_branch: #{e}"
end
```

#### Using the create_branch_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> create_branch_with_http_info(request)

```ruby
begin
  # Create a new branch
  data, status_code, headers = api_instance.create_branch_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->create_branch_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**GitBranchRequest**](GitBranchRequest.md) | Create branch request |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## delete_branch

> delete_branch(request)

Delete a branch

Delete a branch from the Git repository

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::GitApi.new
request = DaytonaToolboxApiClient::PkgToolboxGitGitDeleteBranchRequest.new({name: 'name_example', path: 'path_example'}) # PkgToolboxGitGitDeleteBranchRequest | Delete branch request

begin
  # Delete a branch
  api_instance.delete_branch(request)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->delete_branch: #{e}"
end
```

#### Using the delete_branch_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> delete_branch_with_http_info(request)

```ruby
begin
  # Delete a branch
  data, status_code, headers = api_instance.delete_branch_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->delete_branch_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**PkgToolboxGitGitDeleteBranchRequest**](PkgToolboxGitGitDeleteBranchRequest.md) | Delete branch request |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## get_commit_history

> <Array<GitCommitInfo>> get_commit_history(path)

Get commit history

Get the commit history of the Git repository

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::GitApi.new
path = 'path_example' # String | Repository path

begin
  # Get commit history
  result = api_instance.get_commit_history(path)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->get_commit_history: #{e}"
end
```

#### Using the get_commit_history_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<GitCommitInfo>>, Integer, Hash)> get_commit_history_with_http_info(path)

```ruby
begin
  # Get commit history
  data, status_code, headers = api_instance.get_commit_history_with_http_info(path)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<GitCommitInfo>>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->get_commit_history_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | Repository path |  |

### Return type

[**Array&lt;GitCommitInfo&gt;**](GitCommitInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_status

> <GitStatus> get_status(path)

Get Git status

Get the Git status of the repository at the specified path

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::GitApi.new
path = 'path_example' # String | Repository path

begin
  # Get Git status
  result = api_instance.get_status(path)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->get_status: #{e}"
end
```

#### Using the get_status_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<GitStatus>, Integer, Hash)> get_status_with_http_info(path)

```ruby
begin
  # Get Git status
  data, status_code, headers = api_instance.get_status_with_http_info(path)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <GitStatus>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->get_status_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | Repository path |  |

### Return type

[**GitStatus**](GitStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## list_branches

> <ListBranchResponse> list_branches(path)

List branches

Get a list of all branches in the Git repository

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::GitApi.new
path = 'path_example' # String | Repository path

begin
  # List branches
  result = api_instance.list_branches(path)
  p result
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->list_branches: #{e}"
end
```

#### Using the list_branches_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ListBranchResponse>, Integer, Hash)> list_branches_with_http_info(path)

```ruby
begin
  # List branches
  data, status_code, headers = api_instance.list_branches_with_http_info(path)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ListBranchResponse>
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->list_branches_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | Repository path |  |

### Return type

[**ListBranchResponse**](ListBranchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## pull_changes

> pull_changes(request)

Pull changes from remote

Pull changes from the remote Git repository

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::GitApi.new
request = DaytonaToolboxApiClient::GitRepoRequest.new({path: 'path_example'}) # GitRepoRequest | Pull request

begin
  # Pull changes from remote
  api_instance.pull_changes(request)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->pull_changes: #{e}"
end
```

#### Using the pull_changes_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> pull_changes_with_http_info(request)

```ruby
begin
  # Pull changes from remote
  data, status_code, headers = api_instance.pull_changes_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->pull_changes_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**GitRepoRequest**](GitRepoRequest.md) | Pull request |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## push_changes

> push_changes(request)

Push changes to remote

Push local changes to the remote Git repository

### Examples

```ruby
require 'time'
require 'daytona_toolbox_api_client'

api_instance = DaytonaToolboxApiClient::GitApi.new
request = DaytonaToolboxApiClient::GitRepoRequest.new({path: 'path_example'}) # GitRepoRequest | Push request

begin
  # Push changes to remote
  api_instance.push_changes(request)
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->push_changes: #{e}"
end
```

#### Using the push_changes_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> push_changes_with_http_info(request)

```ruby
begin
  # Push changes to remote
  data, status_code, headers = api_instance.push_changes_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue DaytonaToolboxApiClient::ApiError => e
  puts "Error when calling GitApi->push_changes_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**GitRepoRequest**](GitRepoRequest.md) | Push request |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

