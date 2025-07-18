# Go API client for threatinsight

OpenAPI specification for Infoblox NIOS WAPI THREATINSIGHT objects

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.13.6
- Package version: 1.0.0
- Generator version: 7.5.0
- Build package: com.infoblox.codegen.NiosGoClientCodegen
For more information, please visit [https://www.infoblox.com](https://www.infoblox.com)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import threatinsight "github.com/infobloxopen/infoblox-nios-go-client/threatinsight"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `threatinsight.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), threatinsight.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `threatinsight.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), threatinsight.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `threatinsight.ContextOperationServerIndices` and `threatinsight.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), threatinsight.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), threatinsight.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/wapi/v2.13.6*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ThreatinsightAllowlistAPI* | [**Create**](docs/ThreatinsightAllowlistAPI.md#create) | **Post** /threatinsight:allowlist | Create a threatinsight:allowlist object
*ThreatinsightAllowlistAPI* | [**Delete**](docs/ThreatinsightAllowlistAPI.md#delete) | **Delete** /threatinsight:allowlist/{reference} | Delete a threatinsight:allowlist object
*ThreatinsightAllowlistAPI* | [**List**](docs/ThreatinsightAllowlistAPI.md#list) | **Get** /threatinsight:allowlist | Retrieve threatinsight:allowlist objects
*ThreatinsightAllowlistAPI* | [**Read**](docs/ThreatinsightAllowlistAPI.md#read) | **Get** /threatinsight:allowlist/{reference} | Get a specific threatinsight:allowlist object
*ThreatinsightAllowlistAPI* | [**Update**](docs/ThreatinsightAllowlistAPI.md#update) | **Put** /threatinsight:allowlist/{reference} | Update a threatinsight:allowlist object
*ThreatinsightCloudclientAPI* | [**List**](docs/ThreatinsightCloudclientAPI.md#list) | **Get** /threatinsight:cloudclient | Retrieve threatinsight:cloudclient objects
*ThreatinsightCloudclientAPI* | [**Read**](docs/ThreatinsightCloudclientAPI.md#read) | **Get** /threatinsight:cloudclient/{reference} | Get a specific threatinsight:cloudclient object
*ThreatinsightCloudclientAPI* | [**Update**](docs/ThreatinsightCloudclientAPI.md#update) | **Put** /threatinsight:cloudclient/{reference} | Update a threatinsight:cloudclient object
*ThreatinsightInsightAllowlistAPI* | [**List**](docs/ThreatinsightInsightAllowlistAPI.md#list) | **Get** /threatinsight:insight_allowlist | Retrieve threatinsight:insight_allowlist objects
*ThreatinsightInsightAllowlistAPI* | [**Read**](docs/ThreatinsightInsightAllowlistAPI.md#read) | **Get** /threatinsight:insight_allowlist/{reference} | Get a specific threatinsight:insight_allowlist object
*ThreatinsightModulesetAPI* | [**List**](docs/ThreatinsightModulesetAPI.md#list) | **Get** /threatinsight:moduleset | Retrieve threatinsight:moduleset objects
*ThreatinsightModulesetAPI* | [**Read**](docs/ThreatinsightModulesetAPI.md#read) | **Get** /threatinsight:moduleset/{reference} | Get a specific threatinsight:moduleset object


## Documentation For Models

 - [CreateThreatinsightAllowlistResponse](docs/CreateThreatinsightAllowlistResponse.md)
 - [CreateThreatinsightAllowlistResponseAsObject](docs/CreateThreatinsightAllowlistResponseAsObject.md)
 - [ExtAttrs](docs/ExtAttrs.md)
 - [GetThreatinsightAllowlistResponse](docs/GetThreatinsightAllowlistResponse.md)
 - [GetThreatinsightAllowlistResponseObjectAsResult](docs/GetThreatinsightAllowlistResponseObjectAsResult.md)
 - [GetThreatinsightCloudclientResponse](docs/GetThreatinsightCloudclientResponse.md)
 - [GetThreatinsightCloudclientResponseObjectAsResult](docs/GetThreatinsightCloudclientResponseObjectAsResult.md)
 - [GetThreatinsightInsightAllowlistResponse](docs/GetThreatinsightInsightAllowlistResponse.md)
 - [GetThreatinsightInsightAllowlistResponseObjectAsResult](docs/GetThreatinsightInsightAllowlistResponseObjectAsResult.md)
 - [GetThreatinsightModulesetResponse](docs/GetThreatinsightModulesetResponse.md)
 - [GetThreatinsightModulesetResponseObjectAsResult](docs/GetThreatinsightModulesetResponseObjectAsResult.md)
 - [ListThreatinsightAllowlistResponse](docs/ListThreatinsightAllowlistResponse.md)
 - [ListThreatinsightAllowlistResponseObject](docs/ListThreatinsightAllowlistResponseObject.md)
 - [ListThreatinsightCloudclientResponse](docs/ListThreatinsightCloudclientResponse.md)
 - [ListThreatinsightCloudclientResponseObject](docs/ListThreatinsightCloudclientResponseObject.md)
 - [ListThreatinsightInsightAllowlistResponse](docs/ListThreatinsightInsightAllowlistResponse.md)
 - [ListThreatinsightInsightAllowlistResponseObject](docs/ListThreatinsightInsightAllowlistResponseObject.md)
 - [ListThreatinsightModulesetResponse](docs/ListThreatinsightModulesetResponse.md)
 - [ListThreatinsightModulesetResponseObject](docs/ListThreatinsightModulesetResponseObject.md)
 - [ThreatinsightAllowlist](docs/ThreatinsightAllowlist.md)
 - [ThreatinsightCloudclient](docs/ThreatinsightCloudclient.md)
 - [ThreatinsightInsightAllowlist](docs/ThreatinsightInsightAllowlist.md)
 - [ThreatinsightModuleset](docs/ThreatinsightModuleset.md)
 - [UpdateThreatinsightAllowlistResponse](docs/UpdateThreatinsightAllowlistResponse.md)
 - [UpdateThreatinsightAllowlistResponseAsObject](docs/UpdateThreatinsightAllowlistResponseAsObject.md)
 - [UpdateThreatinsightCloudclientResponse](docs/UpdateThreatinsightCloudclientResponse.md)
 - [UpdateThreatinsightCloudclientResponseAsObject](docs/UpdateThreatinsightCloudclientResponseAsObject.md)
 - [UpdateThreatinsightInsightAllowlistResponse](docs/UpdateThreatinsightInsightAllowlistResponse.md)
 - [UpdateThreatinsightInsightAllowlistResponseAsObject](docs/UpdateThreatinsightInsightAllowlistResponseAsObject.md)
 - [UpdateThreatinsightModulesetResponse](docs/UpdateThreatinsightModulesetResponse.md)
 - [UpdateThreatinsightModulesetResponseAsObject](docs/UpdateThreatinsightModulesetResponseAsObject.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### basicAuth

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), threatinsight.ContextBasicAuth, threatinsight.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


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



