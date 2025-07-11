# Go API client for resources

Chains API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Generator version: 7.14.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import resources "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `resources.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), resources.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `resources.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), resources.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `resources.ContextOperationServerIndices` and `resources.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), resources.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), resources.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------


## Documentation For Models

 - [AccessToken](docs/AccessToken.md)
 - [AccessTokenData](docs/AccessTokenData.md)
 - [AccessTokenDataAttributes](docs/AccessTokenDataAttributes.md)
 - [Biography](docs/Biography.md)
 - [BiographyAttributes](docs/BiographyAttributes.md)
 - [BiographyData](docs/BiographyData.md)
 - [Cabinet](docs/Cabinet.md)
 - [CabinetAttributes](docs/CabinetAttributes.md)
 - [CabinetData](docs/CabinetData.md)
 - [CreateOwnProfile](docs/CreateOwnProfile.md)
 - [CreateOwnProfileData](docs/CreateOwnProfileData.md)
 - [CreateOwnProfileDataAttributes](docs/CreateOwnProfileDataAttributes.md)
 - [CreateUserByAdmin](docs/CreateUserByAdmin.md)
 - [CreateUserByAdminData](docs/CreateUserByAdminData.md)
 - [CreateUserByAdminDataAttributes](docs/CreateUserByAdminDataAttributes.md)
 - [Errors](docs/Errors.md)
 - [ErrorsErrorsInner](docs/ErrorsErrorsInner.md)
 - [Profile](docs/Profile.md)
 - [ProfileAttributes](docs/ProfileAttributes.md)
 - [ProfileData](docs/ProfileData.md)
 - [RefreshToken](docs/RefreshToken.md)
 - [RefreshTokenData](docs/RefreshTokenData.md)
 - [RefreshTokenDataAttributes](docs/RefreshTokenDataAttributes.md)
 - [Session](docs/Session.md)
 - [SessionAttributes](docs/SessionAttributes.md)
 - [SessionData](docs/SessionData.md)
 - [SessionsCollection](docs/SessionsCollection.md)
 - [SessionsCollectionData](docs/SessionsCollectionData.md)
 - [SessionsCollectionDataAttributes](docs/SessionsCollectionDataAttributes.md)
 - [TokensPair](docs/TokensPair.md)
 - [TokensPairData](docs/TokensPairData.md)
 - [TokensPairDataAttributes](docs/TokensPairDataAttributes.md)
 - [UpdateOwnProfile](docs/UpdateOwnProfile.md)
 - [UpdateOwnProfileData](docs/UpdateOwnProfileData.md)
 - [UpdateOwnProfileDataAttributes](docs/UpdateOwnProfileDataAttributes.md)
 - [UpdateProfileByAdmin](docs/UpdateProfileByAdmin.md)
 - [UpdateProfileByAdminData](docs/UpdateProfileByAdminData.md)
 - [UpdateProfileByAdminDataAttributes](docs/UpdateProfileByAdminDataAttributes.md)
 - [User](docs/User.md)
 - [UserData](docs/UserData.md)
 - [UserDataAttributes](docs/UserDataAttributes.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


