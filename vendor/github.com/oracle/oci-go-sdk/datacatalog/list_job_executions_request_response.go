// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListJobExecutionsRequest wrapper for the ListJobExecutions operation
type ListJobExecutionsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique job key.
	JobKey *string `mandatory:"true" contributesTo:"path" name:"jobKey"`

	// Job execution lifecycle state.
	LifecycleState ListJobExecutionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Job type.
	JobType ListJobExecutionsJobTypeEnum `mandatory:"false" contributesTo:"query" name:"jobType" omitEmpty:"true"`

	// Sub-type of this job execution.
	SubType *string `mandatory:"false" contributesTo:"query" name:"subType"`

	// The unique key of the parent execution or null if this job execution has no parent.
	ParentKey *string `mandatory:"false" contributesTo:"query" name:"parentKey"`

	// Time that the job execution was started or in the case of a future time, the time when the job will start.
	// An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// Time that the job execution ended or null if the job is still running or hasn't run yet.
	// An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// Error code returned from the job execution or null if job is still running or didn't return an error.
	ErrorCode *string `mandatory:"false" contributesTo:"query" name:"errorCode"`

	// Error message returned from the job execution or null if job is still running or didn't return an error.
	ErrorMessage *string `mandatory:"false" contributesTo:"query" name:"errorMessage"`

	// Process identifier related to the job execution.
	ProcessKey *string `mandatory:"false" contributesTo:"query" name:"processKey"`

	// The a URL of the job for accessing this resource and its status.
	ExternalUrl *string `mandatory:"false" contributesTo:"query" name:"externalUrl"`

	// Event that triggered the execution of this job or null.
	EventKey *string `mandatory:"false" contributesTo:"query" name:"eventKey"`

	// Unique entity key.
	DataEntityKey *string `mandatory:"false" contributesTo:"query" name:"dataEntityKey"`

	// Specifies the fields to return in a job execution summary response.
	Fields []ListJobExecutionsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListJobExecutionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListJobExecutionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJobExecutionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJobExecutionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobExecutionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListJobExecutionsResponse wrapper for the ListJobExecutions operation
type ListJobExecutionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JobExecutionCollection instances
	JobExecutionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJobExecutionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJobExecutionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJobExecutionsLifecycleStateEnum Enum with underlying type: string
type ListJobExecutionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListJobExecutionsLifecycleStateEnum
const (
	ListJobExecutionsLifecycleStateCreated    ListJobExecutionsLifecycleStateEnum = "CREATED"
	ListJobExecutionsLifecycleStateInProgress ListJobExecutionsLifecycleStateEnum = "IN_PROGRESS"
	ListJobExecutionsLifecycleStateInactive   ListJobExecutionsLifecycleStateEnum = "INACTIVE"
	ListJobExecutionsLifecycleStateFailed     ListJobExecutionsLifecycleStateEnum = "FAILED"
	ListJobExecutionsLifecycleStateSucceeded  ListJobExecutionsLifecycleStateEnum = "SUCCEEDED"
	ListJobExecutionsLifecycleStateCanceled   ListJobExecutionsLifecycleStateEnum = "CANCELED"
)

var mappingListJobExecutionsLifecycleState = map[string]ListJobExecutionsLifecycleStateEnum{
	"CREATED":     ListJobExecutionsLifecycleStateCreated,
	"IN_PROGRESS": ListJobExecutionsLifecycleStateInProgress,
	"INACTIVE":    ListJobExecutionsLifecycleStateInactive,
	"FAILED":      ListJobExecutionsLifecycleStateFailed,
	"SUCCEEDED":   ListJobExecutionsLifecycleStateSucceeded,
	"CANCELED":    ListJobExecutionsLifecycleStateCanceled,
}

// GetListJobExecutionsLifecycleStateEnumValues Enumerates the set of values for ListJobExecutionsLifecycleStateEnum
func GetListJobExecutionsLifecycleStateEnumValues() []ListJobExecutionsLifecycleStateEnum {
	values := make([]ListJobExecutionsLifecycleStateEnum, 0)
	for _, v := range mappingListJobExecutionsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListJobExecutionsJobTypeEnum Enum with underlying type: string
type ListJobExecutionsJobTypeEnum string

// Set of constants representing the allowable values for ListJobExecutionsJobTypeEnum
const (
	ListJobExecutionsJobTypeHarvest                    ListJobExecutionsJobTypeEnum = "HARVEST"
	ListJobExecutionsJobTypeProfiling                  ListJobExecutionsJobTypeEnum = "PROFILING"
	ListJobExecutionsJobTypeSampling                   ListJobExecutionsJobTypeEnum = "SAMPLING"
	ListJobExecutionsJobTypePreview                    ListJobExecutionsJobTypeEnum = "PREVIEW"
	ListJobExecutionsJobTypeImport                     ListJobExecutionsJobTypeEnum = "IMPORT"
	ListJobExecutionsJobTypeExport                     ListJobExecutionsJobTypeEnum = "EXPORT"
	ListJobExecutionsJobTypeImportGlossary             ListJobExecutionsJobTypeEnum = "IMPORT_GLOSSARY"
	ListJobExecutionsJobTypeExportGlossary             ListJobExecutionsJobTypeEnum = "EXPORT_GLOSSARY"
	ListJobExecutionsJobTypeInternal                   ListJobExecutionsJobTypeEnum = "INTERNAL"
	ListJobExecutionsJobTypePurge                      ListJobExecutionsJobTypeEnum = "PURGE"
	ListJobExecutionsJobTypeImmediate                  ListJobExecutionsJobTypeEnum = "IMMEDIATE"
	ListJobExecutionsJobTypeScheduled                  ListJobExecutionsJobTypeEnum = "SCHEDULED"
	ListJobExecutionsJobTypeImmediateExecution         ListJobExecutionsJobTypeEnum = "IMMEDIATE_EXECUTION"
	ListJobExecutionsJobTypeScheduledExecution         ListJobExecutionsJobTypeEnum = "SCHEDULED_EXECUTION"
	ListJobExecutionsJobTypeScheduledExecutionInstance ListJobExecutionsJobTypeEnum = "SCHEDULED_EXECUTION_INSTANCE"
)

var mappingListJobExecutionsJobType = map[string]ListJobExecutionsJobTypeEnum{
	"HARVEST":                      ListJobExecutionsJobTypeHarvest,
	"PROFILING":                    ListJobExecutionsJobTypeProfiling,
	"SAMPLING":                     ListJobExecutionsJobTypeSampling,
	"PREVIEW":                      ListJobExecutionsJobTypePreview,
	"IMPORT":                       ListJobExecutionsJobTypeImport,
	"EXPORT":                       ListJobExecutionsJobTypeExport,
	"IMPORT_GLOSSARY":              ListJobExecutionsJobTypeImportGlossary,
	"EXPORT_GLOSSARY":              ListJobExecutionsJobTypeExportGlossary,
	"INTERNAL":                     ListJobExecutionsJobTypeInternal,
	"PURGE":                        ListJobExecutionsJobTypePurge,
	"IMMEDIATE":                    ListJobExecutionsJobTypeImmediate,
	"SCHEDULED":                    ListJobExecutionsJobTypeScheduled,
	"IMMEDIATE_EXECUTION":          ListJobExecutionsJobTypeImmediateExecution,
	"SCHEDULED_EXECUTION":          ListJobExecutionsJobTypeScheduledExecution,
	"SCHEDULED_EXECUTION_INSTANCE": ListJobExecutionsJobTypeScheduledExecutionInstance,
}

// GetListJobExecutionsJobTypeEnumValues Enumerates the set of values for ListJobExecutionsJobTypeEnum
func GetListJobExecutionsJobTypeEnumValues() []ListJobExecutionsJobTypeEnum {
	values := make([]ListJobExecutionsJobTypeEnum, 0)
	for _, v := range mappingListJobExecutionsJobType {
		values = append(values, v)
	}
	return values
}

// ListJobExecutionsFieldsEnum Enum with underlying type: string
type ListJobExecutionsFieldsEnum string

// Set of constants representing the allowable values for ListJobExecutionsFieldsEnum
const (
	ListJobExecutionsFieldsKey                 ListJobExecutionsFieldsEnum = "key"
	ListJobExecutionsFieldsJobkey              ListJobExecutionsFieldsEnum = "jobKey"
	ListJobExecutionsFieldsJobtype             ListJobExecutionsFieldsEnum = "jobType"
	ListJobExecutionsFieldsParentkey           ListJobExecutionsFieldsEnum = "parentKey"
	ListJobExecutionsFieldsScheduleinstancekey ListJobExecutionsFieldsEnum = "scheduleInstanceKey"
	ListJobExecutionsFieldsLifecyclestate      ListJobExecutionsFieldsEnum = "lifecycleState"
	ListJobExecutionsFieldsTimecreated         ListJobExecutionsFieldsEnum = "timeCreated"
	ListJobExecutionsFieldsTimestarted         ListJobExecutionsFieldsEnum = "timeStarted"
	ListJobExecutionsFieldsTimeended           ListJobExecutionsFieldsEnum = "timeEnded"
	ListJobExecutionsFieldsUri                 ListJobExecutionsFieldsEnum = "uri"
)

var mappingListJobExecutionsFields = map[string]ListJobExecutionsFieldsEnum{
	"key":                 ListJobExecutionsFieldsKey,
	"jobKey":              ListJobExecutionsFieldsJobkey,
	"jobType":             ListJobExecutionsFieldsJobtype,
	"parentKey":           ListJobExecutionsFieldsParentkey,
	"scheduleInstanceKey": ListJobExecutionsFieldsScheduleinstancekey,
	"lifecycleState":      ListJobExecutionsFieldsLifecyclestate,
	"timeCreated":         ListJobExecutionsFieldsTimecreated,
	"timeStarted":         ListJobExecutionsFieldsTimestarted,
	"timeEnded":           ListJobExecutionsFieldsTimeended,
	"uri":                 ListJobExecutionsFieldsUri,
}

// GetListJobExecutionsFieldsEnumValues Enumerates the set of values for ListJobExecutionsFieldsEnum
func GetListJobExecutionsFieldsEnumValues() []ListJobExecutionsFieldsEnum {
	values := make([]ListJobExecutionsFieldsEnum, 0)
	for _, v := range mappingListJobExecutionsFields {
		values = append(values, v)
	}
	return values
}

// ListJobExecutionsSortByEnum Enum with underlying type: string
type ListJobExecutionsSortByEnum string

// Set of constants representing the allowable values for ListJobExecutionsSortByEnum
const (
	ListJobExecutionsSortByTimecreated ListJobExecutionsSortByEnum = "TIMECREATED"
	ListJobExecutionsSortByDisplayname ListJobExecutionsSortByEnum = "DISPLAYNAME"
)

var mappingListJobExecutionsSortBy = map[string]ListJobExecutionsSortByEnum{
	"TIMECREATED": ListJobExecutionsSortByTimecreated,
	"DISPLAYNAME": ListJobExecutionsSortByDisplayname,
}

// GetListJobExecutionsSortByEnumValues Enumerates the set of values for ListJobExecutionsSortByEnum
func GetListJobExecutionsSortByEnumValues() []ListJobExecutionsSortByEnum {
	values := make([]ListJobExecutionsSortByEnum, 0)
	for _, v := range mappingListJobExecutionsSortBy {
		values = append(values, v)
	}
	return values
}

// ListJobExecutionsSortOrderEnum Enum with underlying type: string
type ListJobExecutionsSortOrderEnum string

// Set of constants representing the allowable values for ListJobExecutionsSortOrderEnum
const (
	ListJobExecutionsSortOrderAsc  ListJobExecutionsSortOrderEnum = "ASC"
	ListJobExecutionsSortOrderDesc ListJobExecutionsSortOrderEnum = "DESC"
)

var mappingListJobExecutionsSortOrder = map[string]ListJobExecutionsSortOrderEnum{
	"ASC":  ListJobExecutionsSortOrderAsc,
	"DESC": ListJobExecutionsSortOrderDesc,
}

// GetListJobExecutionsSortOrderEnumValues Enumerates the set of values for ListJobExecutionsSortOrderEnum
func GetListJobExecutionsSortOrderEnumValues() []ListJobExecutionsSortOrderEnum {
	values := make([]ListJobExecutionsSortOrderEnum, 0)
	for _, v := range mappingListJobExecutionsSortOrder {
		values = append(values, v)
	}
	return values
}
