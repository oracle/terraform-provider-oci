// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListJobExecutionsRequest wrapper for the ListJobExecutions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListJobExecutions.go.html to see an example of how to use ListJobExecutionsRequest.
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

	// The field to sort by. Only one sort order may be provided; the default is descending. Use sortOrder query param to specify order.
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
func (request ListJobExecutionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJobExecutionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobExecutionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJobExecutionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJobExecutionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListJobExecutionsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobExecutionsJobTypeEnum(string(request.JobType)); !ok && request.JobType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobType: %s. Supported values are: %s.", request.JobType, strings.Join(GetListJobExecutionsJobTypeEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListJobExecutionsFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListJobExecutionsFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListJobExecutionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJobExecutionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobExecutionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJobExecutionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
	ListJobExecutionsLifecycleStateCreated               ListJobExecutionsLifecycleStateEnum = "CREATED"
	ListJobExecutionsLifecycleStateInProgress            ListJobExecutionsLifecycleStateEnum = "IN_PROGRESS"
	ListJobExecutionsLifecycleStateInactive              ListJobExecutionsLifecycleStateEnum = "INACTIVE"
	ListJobExecutionsLifecycleStateFailed                ListJobExecutionsLifecycleStateEnum = "FAILED"
	ListJobExecutionsLifecycleStateSucceeded             ListJobExecutionsLifecycleStateEnum = "SUCCEEDED"
	ListJobExecutionsLifecycleStateCanceled              ListJobExecutionsLifecycleStateEnum = "CANCELED"
	ListJobExecutionsLifecycleStateSucceededWithWarnings ListJobExecutionsLifecycleStateEnum = "SUCCEEDED_WITH_WARNINGS"
)

var mappingListJobExecutionsLifecycleStateEnum = map[string]ListJobExecutionsLifecycleStateEnum{
	"CREATED":                 ListJobExecutionsLifecycleStateCreated,
	"IN_PROGRESS":             ListJobExecutionsLifecycleStateInProgress,
	"INACTIVE":                ListJobExecutionsLifecycleStateInactive,
	"FAILED":                  ListJobExecutionsLifecycleStateFailed,
	"SUCCEEDED":               ListJobExecutionsLifecycleStateSucceeded,
	"CANCELED":                ListJobExecutionsLifecycleStateCanceled,
	"SUCCEEDED_WITH_WARNINGS": ListJobExecutionsLifecycleStateSucceededWithWarnings,
}

// GetListJobExecutionsLifecycleStateEnumValues Enumerates the set of values for ListJobExecutionsLifecycleStateEnum
func GetListJobExecutionsLifecycleStateEnumValues() []ListJobExecutionsLifecycleStateEnum {
	values := make([]ListJobExecutionsLifecycleStateEnum, 0)
	for _, v := range mappingListJobExecutionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobExecutionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListJobExecutionsLifecycleStateEnum
func GetListJobExecutionsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATED",
		"IN_PROGRESS",
		"INACTIVE",
		"FAILED",
		"SUCCEEDED",
		"CANCELED",
		"SUCCEEDED_WITH_WARNINGS",
	}
}

// GetMappingListJobExecutionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobExecutionsLifecycleStateEnum(val string) (ListJobExecutionsLifecycleStateEnum, bool) {
	mappingListJobExecutionsLifecycleStateEnumIgnoreCase := make(map[string]ListJobExecutionsLifecycleStateEnum)
	for k, v := range mappingListJobExecutionsLifecycleStateEnum {
		mappingListJobExecutionsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListJobExecutionsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
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
	ListJobExecutionsJobTypeAsyncDelete                ListJobExecutionsJobTypeEnum = "ASYNC_DELETE"
	ListJobExecutionsJobTypeImportDataAsset            ListJobExecutionsJobTypeEnum = "IMPORT_DATA_ASSET"
)

var mappingListJobExecutionsJobTypeEnum = map[string]ListJobExecutionsJobTypeEnum{
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
	"ASYNC_DELETE":                 ListJobExecutionsJobTypeAsyncDelete,
	"IMPORT_DATA_ASSET":            ListJobExecutionsJobTypeImportDataAsset,
}

// GetListJobExecutionsJobTypeEnumValues Enumerates the set of values for ListJobExecutionsJobTypeEnum
func GetListJobExecutionsJobTypeEnumValues() []ListJobExecutionsJobTypeEnum {
	values := make([]ListJobExecutionsJobTypeEnum, 0)
	for _, v := range mappingListJobExecutionsJobTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobExecutionsJobTypeEnumStringValues Enumerates the set of values in String for ListJobExecutionsJobTypeEnum
func GetListJobExecutionsJobTypeEnumStringValues() []string {
	return []string{
		"HARVEST",
		"PROFILING",
		"SAMPLING",
		"PREVIEW",
		"IMPORT",
		"EXPORT",
		"IMPORT_GLOSSARY",
		"EXPORT_GLOSSARY",
		"INTERNAL",
		"PURGE",
		"IMMEDIATE",
		"SCHEDULED",
		"IMMEDIATE_EXECUTION",
		"SCHEDULED_EXECUTION",
		"SCHEDULED_EXECUTION_INSTANCE",
		"ASYNC_DELETE",
		"IMPORT_DATA_ASSET",
	}
}

// GetMappingListJobExecutionsJobTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobExecutionsJobTypeEnum(val string) (ListJobExecutionsJobTypeEnum, bool) {
	mappingListJobExecutionsJobTypeEnumIgnoreCase := make(map[string]ListJobExecutionsJobTypeEnum)
	for k, v := range mappingListJobExecutionsJobTypeEnum {
		mappingListJobExecutionsJobTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListJobExecutionsJobTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
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

var mappingListJobExecutionsFieldsEnum = map[string]ListJobExecutionsFieldsEnum{
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
	for _, v := range mappingListJobExecutionsFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobExecutionsFieldsEnumStringValues Enumerates the set of values in String for ListJobExecutionsFieldsEnum
func GetListJobExecutionsFieldsEnumStringValues() []string {
	return []string{
		"key",
		"jobKey",
		"jobType",
		"parentKey",
		"scheduleInstanceKey",
		"lifecycleState",
		"timeCreated",
		"timeStarted",
		"timeEnded",
		"uri",
	}
}

// GetMappingListJobExecutionsFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobExecutionsFieldsEnum(val string) (ListJobExecutionsFieldsEnum, bool) {
	mappingListJobExecutionsFieldsEnumIgnoreCase := make(map[string]ListJobExecutionsFieldsEnum)
	for k, v := range mappingListJobExecutionsFieldsEnum {
		mappingListJobExecutionsFieldsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListJobExecutionsFieldsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobExecutionsSortByEnum Enum with underlying type: string
type ListJobExecutionsSortByEnum string

// Set of constants representing the allowable values for ListJobExecutionsSortByEnum
const (
	ListJobExecutionsSortByTimecreated ListJobExecutionsSortByEnum = "TIMECREATED"
)

var mappingListJobExecutionsSortByEnum = map[string]ListJobExecutionsSortByEnum{
	"TIMECREATED": ListJobExecutionsSortByTimecreated,
}

// GetListJobExecutionsSortByEnumValues Enumerates the set of values for ListJobExecutionsSortByEnum
func GetListJobExecutionsSortByEnumValues() []ListJobExecutionsSortByEnum {
	values := make([]ListJobExecutionsSortByEnum, 0)
	for _, v := range mappingListJobExecutionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobExecutionsSortByEnumStringValues Enumerates the set of values in String for ListJobExecutionsSortByEnum
func GetListJobExecutionsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
	}
}

// GetMappingListJobExecutionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobExecutionsSortByEnum(val string) (ListJobExecutionsSortByEnum, bool) {
	mappingListJobExecutionsSortByEnumIgnoreCase := make(map[string]ListJobExecutionsSortByEnum)
	for k, v := range mappingListJobExecutionsSortByEnum {
		mappingListJobExecutionsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListJobExecutionsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobExecutionsSortOrderEnum Enum with underlying type: string
type ListJobExecutionsSortOrderEnum string

// Set of constants representing the allowable values for ListJobExecutionsSortOrderEnum
const (
	ListJobExecutionsSortOrderAsc  ListJobExecutionsSortOrderEnum = "ASC"
	ListJobExecutionsSortOrderDesc ListJobExecutionsSortOrderEnum = "DESC"
)

var mappingListJobExecutionsSortOrderEnum = map[string]ListJobExecutionsSortOrderEnum{
	"ASC":  ListJobExecutionsSortOrderAsc,
	"DESC": ListJobExecutionsSortOrderDesc,
}

// GetListJobExecutionsSortOrderEnumValues Enumerates the set of values for ListJobExecutionsSortOrderEnum
func GetListJobExecutionsSortOrderEnumValues() []ListJobExecutionsSortOrderEnum {
	values := make([]ListJobExecutionsSortOrderEnum, 0)
	for _, v := range mappingListJobExecutionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobExecutionsSortOrderEnumStringValues Enumerates the set of values in String for ListJobExecutionsSortOrderEnum
func GetListJobExecutionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJobExecutionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobExecutionsSortOrderEnum(val string) (ListJobExecutionsSortOrderEnum, bool) {
	mappingListJobExecutionsSortOrderEnumIgnoreCase := make(map[string]ListJobExecutionsSortOrderEnum)
	for k, v := range mappingListJobExecutionsSortOrderEnum {
		mappingListJobExecutionsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListJobExecutionsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
