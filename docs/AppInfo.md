# AppInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiName** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**ContractVersions** | Pointer to **[]string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Commit** | Pointer to **string** |  | [optional] 
**DeployDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAppInfo

`func NewAppInfo() *AppInfo`

NewAppInfo instantiates a new AppInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInfoWithDefaults

`func NewAppInfoWithDefaults() *AppInfo`

NewAppInfoWithDefaults instantiates a new AppInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiName

`func (o *AppInfo) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *AppInfo) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *AppInfo) SetApiName(v string)`

SetApiName sets ApiName field to given value.

### HasApiName

`func (o *AppInfo) HasApiName() bool`

HasApiName returns a boolean if a field has been set.

### GetApiVersion

`func (o *AppInfo) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AppInfo) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AppInfo) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AppInfo) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetContractVersions

`func (o *AppInfo) GetContractVersions() []string`

GetContractVersions returns the ContractVersions field if non-nil, zero value otherwise.

### GetContractVersionsOk

`func (o *AppInfo) GetContractVersionsOk() (*[]string, bool)`

GetContractVersionsOk returns a tuple with the ContractVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractVersions

`func (o *AppInfo) SetContractVersions(v []string)`

SetContractVersions sets ContractVersions field to given value.

### HasContractVersions

`func (o *AppInfo) HasContractVersions() bool`

HasContractVersions returns a boolean if a field has been set.

### GetBranch

`func (o *AppInfo) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *AppInfo) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *AppInfo) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *AppInfo) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCommit

`func (o *AppInfo) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *AppInfo) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *AppInfo) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *AppInfo) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetDeployDate

`func (o *AppInfo) GetDeployDate() time.Time`

GetDeployDate returns the DeployDate field if non-nil, zero value otherwise.

### GetDeployDateOk

`func (o *AppInfo) GetDeployDateOk() (*time.Time, bool)`

GetDeployDateOk returns a tuple with the DeployDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployDate

`func (o *AppInfo) SetDeployDate(v time.Time)`

SetDeployDate sets DeployDate field to given value.

### HasDeployDate

`func (o *AppInfo) HasDeployDate() bool`

HasDeployDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


