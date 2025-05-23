# Go API client for dnsclient

This is the Rackspace Cloud DNS v2 contract.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 1.0.0
- Generator version: 7.13.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import dnsclient "github.com/avinashtanniru/rackspace-dnsclient"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `dnsclient.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), dnsclient.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `dnsclient.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), dnsclient.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `dnsclient.ContextOperationServerIndices` and `dnsclient.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), dnsclient.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), dnsclient.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://dns.api.rackspacecloud.com/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuditsAPI* | [**GetAudits**](docs/AuditsAPI.md#getaudits) | **Get** /{tenantId}/zones/{zoneId}/audits | Get all audits for a zone.
*DefaultAPI* | [**InfoGet**](docs/DefaultAPI.md#infoget) | **Get** /info | Get metadata info
*JobsAPI* | [**GetJob**](docs/JobsAPI.md#getjob) | **Get** /{tenantId}/jobs/{jobId} | Get job by ID.
*JobsAPI* | [**GetJobs**](docs/JobsAPI.md#getjobs) | **Get** /{tenantId}/jobs | Get all jobs for a tenant.
*PtrAPI* | [**CreatePtrRecord**](docs/PtrAPI.md#createptrrecord) | **Post** /{tenantId}/ptr | Create a PTR record.
*PtrAPI* | [**DeletePtrRecord**](docs/PtrAPI.md#deleteptrrecord) | **Delete** /{tenantId}/ptr/{ptrId} | Delete a PTR record.
*PtrAPI* | [**GetPtrRecord**](docs/PtrAPI.md#getptrrecord) | **Get** /{tenantId}/ptr/{ptrId} | Get a PTR record by ID.
*PtrAPI* | [**GetPtrRecords**](docs/PtrAPI.md#getptrrecords) | **Get** /{tenantId}/ptr | Get all PTR records for a tenant.
*PtrAPI* | [**UpdatePtrRecord**](docs/PtrAPI.md#updateptrrecord) | **Patch** /{tenantId}/ptr/{ptrId} | Update a PTR record.
*RecordsAPI* | [**CreateRecord**](docs/RecordsAPI.md#createrecord) | **Post** /{tenantId}/zones/{zoneId}/records | Create a record.
*RecordsAPI* | [**DeleteRecord**](docs/RecordsAPI.md#deleterecord) | **Delete** /{tenantId}/zones/{zoneId}/records/{recordId} | Delete a record.
*RecordsAPI* | [**GetRecord**](docs/RecordsAPI.md#getrecord) | **Get** /{tenantId}/zones/{zoneId}/records/{recordId} | Get a record by ID.
*RecordsAPI* | [**GetRecords**](docs/RecordsAPI.md#getrecords) | **Get** /{tenantId}/zones/{zoneId}/records | Get all records for a zone.
*RecordsAPI* | [**UpdateRecord**](docs/RecordsAPI.md#updaterecord) | **Patch** /{tenantId}/zones/{zoneId}/records/{recordId} | Update a record.
*ZonesAPI* | [**CallImport**](docs/ZonesAPI.md#callimport) | **Post** /{tenantId}/zones/import | Import a zone file&#39;s contents.
*ZonesAPI* | [**CreateZone**](docs/ZonesAPI.md#createzone) | **Post** /{tenantId}/zones | Create a zone.
*ZonesAPI* | [**Export**](docs/ZonesAPI.md#export) | **Get** /{tenantId}/zones/{zoneId}/export | Exports a single zone
*ZonesAPI* | [**GetZone**](docs/ZonesAPI.md#getzone) | **Get** /{tenantId}/zones/{zoneId} | Get a zone by ID.
*ZonesAPI* | [**GetZones**](docs/ZonesAPI.md#getzones) | **Get** /{tenantId}/zones | Get all zones for a tenant.
*ZonesAPI* | [**Migrate**](docs/ZonesAPI.md#migrate) | **Post** /{tenantId}/zones/{zoneId}/migrate | Migrates a single zone to a different tenant
*ZonesAPI* | [**TenantIdZonesZoneIdDelete**](docs/ZonesAPI.md#tenantidzoneszoneiddelete) | **Delete** /{tenantId}/zones/{zoneId} | Delete a zone.
*ZonesAPI* | [**UpdateZone**](docs/ZonesAPI.md#updatezone) | **Patch** /{tenantId}/zones/{zoneId} | Update a zone.


## Documentation For Models

 - [AppInfo](docs/AppInfo.md)
 - [Audit](docs/Audit.md)
 - [AuditCollection](docs/AuditCollection.md)
 - [Job](docs/Job.md)
 - [JobCollection](docs/JobCollection.md)
 - [JobResponse](docs/JobResponse.md)
 - [PaginationLink](docs/PaginationLink.md)
 - [PtrCreate](docs/PtrCreate.md)
 - [PtrRecord](docs/PtrRecord.md)
 - [PtrRecordCollection](docs/PtrRecordCollection.md)
 - [PtrRecordSelfLink](docs/PtrRecordSelfLink.md)
 - [PtrUpdate](docs/PtrUpdate.md)
 - [Record](docs/Record.md)
 - [RecordCollection](docs/RecordCollection.md)
 - [RecordCreate](docs/RecordCreate.md)
 - [RecordSelfLink](docs/RecordSelfLink.md)
 - [RecordUpdate](docs/RecordUpdate.md)
 - [Zone](docs/Zone.md)
 - [ZoneCollection](docs/ZoneCollection.md)
 - [ZoneCollectionLinksInner](docs/ZoneCollectionLinksInner.md)
 - [ZoneCreate](docs/ZoneCreate.md)
 - [ZoneLinksInner](docs/ZoneLinksInner.md)
 - [ZoneMigrate](docs/ZoneMigrate.md)
 - [ZoneRecordLink](docs/ZoneRecordLink.md)
 - [ZoneSelfLink](docs/ZoneSelfLink.md)
 - [ZoneUpdate](docs/ZoneUpdate.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### AuthToken

- **Type**: API key
- **API key parameter name**: X-Auth-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: AuthToken and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		dnsclient.ContextAPIKeys,
		map[string]dnsclient.APIKey{
			"AuthToken": {Key: "API_KEY_STRING"},
		},
	)
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



