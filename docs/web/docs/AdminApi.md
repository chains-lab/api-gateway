# \AdminApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReNewsChainsAuthV1AdminUserIdGet**](AdminApi.md#ReNewsChainsAuthV1AdminUserIdGet) | **Get** /re-news/chains/auth/v1/admin/{user_id} | admin get user
[**ReNewsChainsAuthV1AdminUserIdRolePost**](AdminApi.md#ReNewsChainsAuthV1AdminUserIdRolePost) | **Post** /re-news/chains/auth/v1/admin/{user_id}/{role} | admin role update
[**ReNewsChainsAuthV1AdminUserIdSessionsDelete**](AdminApi.md#ReNewsChainsAuthV1AdminUserIdSessionsDelete) | **Delete** /re-news/chains/auth/v1/admin/{user_id}/sessions | admin delete users sessions
[**ReNewsChainsAuthV1AdminUserIdSessionsGet**](AdminApi.md#ReNewsChainsAuthV1AdminUserIdSessionsGet) | **Get** /re-news/chains/auth/v1/admin/{user_id}/sessions | admin get sessions
[**ReNewsChainsAuthV1AdminUserIdSessionsSessionIdDelete**](AdminApi.md#ReNewsChainsAuthV1AdminUserIdSessionsSessionIdDelete) | **Delete** /re-news/chains/auth/v1/admin/{user_id}/sessions/{session_id} | admin delete user session
[**ReNewsChainsAuthV1AdminUserIdSessionsSessionIdGet**](AdminApi.md#ReNewsChainsAuthV1AdminUserIdSessionsSessionIdGet) | **Get** /re-news/chains/auth/v1/admin/{user_id}/sessions/{session_id} | admin get session



## ReNewsChainsAuthV1AdminUserIdGet

> User ReNewsChainsAuthV1AdminUserIdGet(ctx, userId).Execute()

admin get user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    userId := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.ReNewsChainsAuthV1AdminUserIdGet(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.ReNewsChainsAuthV1AdminUserIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReNewsChainsAuthV1AdminUserIdGet`: User
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.ReNewsChainsAuthV1AdminUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReNewsChainsAuthV1AdminUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReNewsChainsAuthV1AdminUserIdRolePost

> User ReNewsChainsAuthV1AdminUserIdRolePost(ctx, userId, role).Execute()

admin role update



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    userId := TODO // interface{} | 
    role := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.ReNewsChainsAuthV1AdminUserIdRolePost(context.Background(), userId, role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.ReNewsChainsAuthV1AdminUserIdRolePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReNewsChainsAuthV1AdminUserIdRolePost`: User
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.ReNewsChainsAuthV1AdminUserIdRolePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**interface{}**](.md) |  | 
**role** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReNewsChainsAuthV1AdminUserIdRolePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReNewsChainsAuthV1AdminUserIdSessionsDelete

> ReNewsChainsAuthV1AdminUserIdSessionsDelete(ctx, userId).Execute()

admin delete users sessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    userId := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.ReNewsChainsAuthV1AdminUserIdSessionsDelete(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.ReNewsChainsAuthV1AdminUserIdSessionsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReNewsChainsAuthV1AdminUserIdSessionsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReNewsChainsAuthV1AdminUserIdSessionsGet

> SessionsCollection ReNewsChainsAuthV1AdminUserIdSessionsGet(ctx, userId).Execute()

admin get sessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    userId := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.ReNewsChainsAuthV1AdminUserIdSessionsGet(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.ReNewsChainsAuthV1AdminUserIdSessionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReNewsChainsAuthV1AdminUserIdSessionsGet`: SessionsCollection
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.ReNewsChainsAuthV1AdminUserIdSessionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReNewsChainsAuthV1AdminUserIdSessionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SessionsCollection**](SessionsCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReNewsChainsAuthV1AdminUserIdSessionsSessionIdDelete

> ReNewsChainsAuthV1AdminUserIdSessionsSessionIdDelete(ctx, userId, sessionId).Execute()

admin delete user session



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    userId := TODO // interface{} | 
    sessionId := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminApi.ReNewsChainsAuthV1AdminUserIdSessionsSessionIdDelete(context.Background(), userId, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.ReNewsChainsAuthV1AdminUserIdSessionsSessionIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**interface{}**](.md) |  | 
**sessionId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReNewsChainsAuthV1AdminUserIdSessionsSessionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReNewsChainsAuthV1AdminUserIdSessionsSessionIdGet

> Session ReNewsChainsAuthV1AdminUserIdSessionsSessionIdGet(ctx, userId, sessionId).Execute()

admin get session



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    userId := TODO // interface{} | 
    sessionId := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.ReNewsChainsAuthV1AdminUserIdSessionsSessionIdGet(context.Background(), userId, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.ReNewsChainsAuthV1AdminUserIdSessionsSessionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReNewsChainsAuthV1AdminUserIdSessionsSessionIdGet`: Session
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.ReNewsChainsAuthV1AdminUserIdSessionsSessionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**interface{}**](.md) |  | 
**sessionId** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReNewsChainsAuthV1AdminUserIdSessionsSessionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

