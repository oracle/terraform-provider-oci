// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListJobsRequest wrapper for the ListJobs operation
type ListJobsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu".
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// Job lifecycle state.
	LifecycleState ListJobsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Job type.
	JobType ListJobsJobTypeEnum `mandatory:"false" contributesTo:"query" name:"jobType" omitEmpty:"true"`

	// Unique job definition key.
	JobDefinitionKey *string `mandatory:"false" contributesTo:"query" name:"jobDefinitionKey"`

	// Schedule specified in the cron expression format that has seven fields for second, minute, hour, day-of-month, month, day-of-week, year.
	// It can also include special characters like * for all and ? for any. There are also pre-defined schedules that can be specified using
	// special strings. For example, @hourly will run the job every hour.
	ScheduleCronExpression *string `mandatory:"false" contributesTo:"query" name:"scheduleCronExpression"`

	// Date that the schedule should be operational. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeScheduleBegin *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScheduleBegin"`

	// Date that the schedule should end from being operational. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeScheduleEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScheduleEnd"`

	// Type of the job schedule.
	ScheduleType ListJobsScheduleTypeEnum `mandatory:"false" contributesTo:"query" name:"scheduleType" omitEmpty:"true"`

	// Unique connection key.
	ConnectionKey *string `mandatory:"false" contributesTo:"query" name:"connectionKey"`

	// Specifies the fields to return in a job summary response.
	Fields []ListJobsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The total number of executions for this job schedule.
	ExecutionCount *int `mandatory:"false" contributesTo:"query" name:"executionCount"`

	// The date and time the most recent execution for this job ,in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeOfLatestExecution *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeOfLatestExecution"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJobsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListJobsResponse wrapper for the ListJobs operation
type ListJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JobCollection instances
	JobCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJobsLifecycleStateEnum Enum with underlying type: string
type ListJobsLifecycleStateEnum string

// Set of constants representing the allowable values for ListJobsLifecycleStateEnum
const (
	ListJobsLifecycleStateActive   ListJobsLifecycleStateEnum = "ACTIVE"
	ListJobsLifecycleStateInactive ListJobsLifecycleStateEnum = "INACTIVE"
	ListJobsLifecycleStateExpired  ListJobsLifecycleStateEnum = "EXPIRED"
)

var mappingListJobsLifecycleState = map[string]ListJobsLifecycleStateEnum{
	"ACTIVE":   ListJobsLifecycleStateActive,
	"INACTIVE": ListJobsLifecycleStateInactive,
	"EXPIRED":  ListJobsLifecycleStateExpired,
}

// GetListJobsLifecycleStateEnumValues Enumerates the set of values for ListJobsLifecycleStateEnum
func GetListJobsLifecycleStateEnumValues() []ListJobsLifecycleStateEnum {
	values := make([]ListJobsLifecycleStateEnum, 0)
	for _, v := range mappingListJobsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListJobsJobTypeEnum Enum with underlying type: string
type ListJobsJobTypeEnum string

// Set of constants representing the allowable values for ListJobsJobTypeEnum
const (
	ListJobsJobTypeHarvest                    ListJobsJobTypeEnum = "HARVEST"
	ListJobsJobTypeProfiling                  ListJobsJobTypeEnum = "PROFILING"
	ListJobsJobTypeSampling                   ListJobsJobTypeEnum = "SAMPLING"
	ListJobsJobTypePreview                    ListJobsJobTypeEnum = "PREVIEW"
	ListJobsJobTypeImport                     ListJobsJobTypeEnum = "IMPORT"
	ListJobsJobTypeExport                     ListJobsJobTypeEnum = "EXPORT"
	ListJobsJobTypeImportGlossary             ListJobsJobTypeEnum = "IMPORT_GLOSSARY"
	ListJobsJobTypeExportGlossary             ListJobsJobTypeEnum = "EXPORT_GLOSSARY"
	ListJobsJobTypeInternal                   ListJobsJobTypeEnum = "INTERNAL"
	ListJobsJobTypePurge                      ListJobsJobTypeEnum = "PURGE"
	ListJobsJobTypeImmediate                  ListJobsJobTypeEnum = "IMMEDIATE"
	ListJobsJobTypeScheduled                  ListJobsJobTypeEnum = "SCHEDULED"
	ListJobsJobTypeImmediateExecution         ListJobsJobTypeEnum = "IMMEDIATE_EXECUTION"
	ListJobsJobTypeScheduledExecution         ListJobsJobTypeEnum = "SCHEDULED_EXECUTION"
	ListJobsJobTypeScheduledExecutionInstance ListJobsJobTypeEnum = "SCHEDULED_EXECUTION_INSTANCE"
)

var mappingListJobsJobType = map[string]ListJobsJobTypeEnum{
	"HARVEST":                      ListJobsJobTypeHarvest,
	"PROFILING":                    ListJobsJobTypeProfiling,
	"SAMPLING":                     ListJobsJobTypeSampling,
	"PREVIEW":                      ListJobsJobTypePreview,
	"IMPORT":                       ListJobsJobTypeImport,
	"EXPORT":                       ListJobsJobTypeExport,
	"IMPORT_GLOSSARY":              ListJobsJobTypeImportGlossary,
	"EXPORT_GLOSSARY":              ListJobsJobTypeExportGlossary,
	"INTERNAL":                     ListJobsJobTypeInternal,
	"PURGE":                        ListJobsJobTypePurge,
	"IMMEDIATE":                    ListJobsJobTypeImmediate,
	"SCHEDULED":                    ListJobsJobTypeScheduled,
	"IMMEDIATE_EXECUTION":          ListJobsJobTypeImmediateExecution,
	"SCHEDULED_EXECUTION":          ListJobsJobTypeScheduledExecution,
	"SCHEDULED_EXECUTION_INSTANCE": ListJobsJobTypeScheduledExecutionInstance,
}

// GetListJobsJobTypeEnumValues Enumerates the set of values for ListJobsJobTypeEnum
func GetListJobsJobTypeEnumValues() []ListJobsJobTypeEnum {
	values := make([]ListJobsJobTypeEnum, 0)
	for _, v := range mappingListJobsJobType {
		values = append(values, v)
	}
	return values
}

// ListJobsScheduleTypeEnum Enum with underlying type: string
type ListJobsScheduleTypeEnum string

// Set of constants representing the allowable values for ListJobsScheduleTypeEnum
const (
	ListJobsScheduleTypeScheduled ListJobsScheduleTypeEnum = "SCHEDULED"
	ListJobsScheduleTypeImmediate ListJobsScheduleTypeEnum = "IMMEDIATE"
)

var mappingListJobsScheduleType = map[string]ListJobsScheduleTypeEnum{
	"SCHEDULED": ListJobsScheduleTypeScheduled,
	"IMMEDIATE": ListJobsScheduleTypeImmediate,
}

// GetListJobsScheduleTypeEnumValues Enumerates the set of values for ListJobsScheduleTypeEnum
func GetListJobsScheduleTypeEnumValues() []ListJobsScheduleTypeEnum {
	values := make([]ListJobsScheduleTypeEnum, 0)
	for _, v := range mappingListJobsScheduleType {
		values = append(values, v)
	}
	return values
}

// ListJobsFieldsEnum Enum with underlying type: string
type ListJobsFieldsEnum string

// Set of constants representing the allowable values for ListJobsFieldsEnum
const (
	ListJobsFieldsKey                    ListJobsFieldsEnum = "key"
	ListJobsFieldsDisplayname            ListJobsFieldsEnum = "displayName"
	ListJobsFieldsDescription            ListJobsFieldsEnum = "description"
	ListJobsFieldsCatalogid              ListJobsFieldsEnum = "catalogId"
	ListJobsFieldsJobdefinitionkey       ListJobsFieldsEnum = "jobDefinitionKey"
	ListJobsFieldsLifecyclestate         ListJobsFieldsEnum = "lifecycleState"
	ListJobsFieldsTimecreated            ListJobsFieldsEnum = "timeCreated"
	ListJobsFieldsTimeupdated            ListJobsFieldsEnum = "timeUpdated"
	ListJobsFieldsCreatedbyid            ListJobsFieldsEnum = "createdById"
	ListJobsFieldsUpdatedbyid            ListJobsFieldsEnum = "updatedById"
	ListJobsFieldsJobtype                ListJobsFieldsEnum = "jobType"
	ListJobsFieldsSchedulecronexpression ListJobsFieldsEnum = "scheduleCronExpression"
	ListJobsFieldsTimeschedulebegin      ListJobsFieldsEnum = "timeScheduleBegin"
	ListJobsFieldsScheduletype           ListJobsFieldsEnum = "scheduleType"
	ListJobsFieldsExecutioncount         ListJobsFieldsEnum = "executionCount"
	ListJobsFieldsTimeoflatestexecution  ListJobsFieldsEnum = "timeOfLatestExecution"
	ListJobsFieldsExecutions             ListJobsFieldsEnum = "executions"
	ListJobsFieldsUri                    ListJobsFieldsEnum = "uri"
	ListJobsFieldsJobdefinitionname      ListJobsFieldsEnum = "jobDefinitionName"
	ListJobsFieldsErrorcode              ListJobsFieldsEnum = "errorCode"
	ListJobsFieldsErrormessage           ListJobsFieldsEnum = "errorMessage"
)

var mappingListJobsFields = map[string]ListJobsFieldsEnum{
	"key":                    ListJobsFieldsKey,
	"displayName":            ListJobsFieldsDisplayname,
	"description":            ListJobsFieldsDescription,
	"catalogId":              ListJobsFieldsCatalogid,
	"jobDefinitionKey":       ListJobsFieldsJobdefinitionkey,
	"lifecycleState":         ListJobsFieldsLifecyclestate,
	"timeCreated":            ListJobsFieldsTimecreated,
	"timeUpdated":            ListJobsFieldsTimeupdated,
	"createdById":            ListJobsFieldsCreatedbyid,
	"updatedById":            ListJobsFieldsUpdatedbyid,
	"jobType":                ListJobsFieldsJobtype,
	"scheduleCronExpression": ListJobsFieldsSchedulecronexpression,
	"timeScheduleBegin":      ListJobsFieldsTimeschedulebegin,
	"scheduleType":           ListJobsFieldsScheduletype,
	"executionCount":         ListJobsFieldsExecutioncount,
	"timeOfLatestExecution":  ListJobsFieldsTimeoflatestexecution,
	"executions":             ListJobsFieldsExecutions,
	"uri":                    ListJobsFieldsUri,
	"jobDefinitionName":      ListJobsFieldsJobdefinitionname,
	"errorCode":              ListJobsFieldsErrorcode,
	"errorMessage":           ListJobsFieldsErrormessage,
}

// GetListJobsFieldsEnumValues Enumerates the set of values for ListJobsFieldsEnum
func GetListJobsFieldsEnumValues() []ListJobsFieldsEnum {
	values := make([]ListJobsFieldsEnum, 0)
	for _, v := range mappingListJobsFields {
		values = append(values, v)
	}
	return values
}

// ListJobsSortByEnum Enum with underlying type: string
type ListJobsSortByEnum string

// Set of constants representing the allowable values for ListJobsSortByEnum
const (
	ListJobsSortByTimecreated ListJobsSortByEnum = "TIMECREATED"
	ListJobsSortByDisplayname ListJobsSortByEnum = "DISPLAYNAME"
)

var mappingListJobsSortBy = map[string]ListJobsSortByEnum{
	"TIMECREATED": ListJobsSortByTimecreated,
	"DISPLAYNAME": ListJobsSortByDisplayname,
}

// GetListJobsSortByEnumValues Enumerates the set of values for ListJobsSortByEnum
func GetListJobsSortByEnumValues() []ListJobsSortByEnum {
	values := make([]ListJobsSortByEnum, 0)
	for _, v := range mappingListJobsSortBy {
		values = append(values, v)
	}
	return values
}

// ListJobsSortOrderEnum Enum with underlying type: string
type ListJobsSortOrderEnum string

// Set of constants representing the allowable values for ListJobsSortOrderEnum
const (
	ListJobsSortOrderAsc  ListJobsSortOrderEnum = "ASC"
	ListJobsSortOrderDesc ListJobsSortOrderEnum = "DESC"
)

var mappingListJobsSortOrder = map[string]ListJobsSortOrderEnum{
	"ASC":  ListJobsSortOrderAsc,
	"DESC": ListJobsSortOrderDesc,
}

// GetListJobsSortOrderEnumValues Enumerates the set of values for ListJobsSortOrderEnum
func GetListJobsSortOrderEnumValues() []ListJobsSortOrderEnum {
	values := make([]ListJobsSortOrderEnum, 0)
	for _, v := range mappingListJobsSortOrder {
		values = append(values, v)
	}
	return values
}
