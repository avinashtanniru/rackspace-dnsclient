# JobCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**[]Job**](Job.md) |  | [optional] 
**Links** | Pointer to [**[]ZoneCollectionLinksInner**](ZoneCollectionLinksInner.md) |  | [optional] 

## Methods

### NewJobCollection

`func NewJobCollection() *JobCollection`

NewJobCollection instantiates a new JobCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobCollectionWithDefaults

`func NewJobCollectionWithDefaults() *JobCollection`

NewJobCollectionWithDefaults instantiates a new JobCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *JobCollection) GetJobs() []Job`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *JobCollection) GetJobsOk() (*[]Job, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *JobCollection) SetJobs(v []Job)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *JobCollection) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetLinks

`func (o *JobCollection) GetLinks() []ZoneCollectionLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *JobCollection) GetLinksOk() (*[]ZoneCollectionLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *JobCollection) SetLinks(v []ZoneCollectionLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *JobCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


