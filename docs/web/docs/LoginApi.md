# \LoginApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReNewsChainsAuthV1OwnCallbackPost**](LoginApi.md#ReNewsChainsAuthV1OwnCallbackPost) | **Post** /re-news/chains/auth/v1/own/callback | 
[**ReNewsChainsAuthV1OwnLoginPost**](LoginApi.md#ReNewsChainsAuthV1OwnLoginPost) | **Post** /re-news/chains/auth/v1/own/login | 
[**ReNewsChainsAuthV1OwnLogoutPost**](LoginApi.md#ReNewsChainsAuthV1OwnLogoutPost) | **Post** /re-news/chains/auth/v1/own/logout | 



## ReNewsChainsAuthV1OwnCallbackPost

> TokensPair ReNewsChainsAuthV1OwnCallbackPost(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.ReNewsChainsAuthV1OwnCallbackPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.ReNewsChainsAuthV1OwnCallbackPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReNewsChainsAuthV1OwnCallbackPost`: TokensPair
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.ReNewsChainsAuthV1OwnCallbackPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReNewsChainsAuthV1OwnCallbackPostRequest struct via the builder pattern


### Return type

[**TokensPair**](TokensPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReNewsChainsAuthV1OwnLoginPost

> ReNewsChainsAuthV1OwnLoginPost(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LoginApi.ReNewsChainsAuthV1OwnLoginPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.ReNewsChainsAuthV1OwnLoginPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReNewsChainsAuthV1OwnLoginPostRequest struct via the builder pattern


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


## ReNewsChainsAuthV1OwnLogoutPost

> TokensPair ReNewsChainsAuthV1OwnLogoutPost(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.ReNewsChainsAuthV1OwnLogoutPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.ReNewsChainsAuthV1OwnLogoutPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReNewsChainsAuthV1OwnLogoutPost`: TokensPair
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.ReNewsChainsAuthV1OwnLogoutPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReNewsChainsAuthV1OwnLogoutPostRequest struct via the builder pattern


### Return type

[**TokensPair**](TokensPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

