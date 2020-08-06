// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListJobDefinitionsRequest wrapper for the ListJobDefinitions operation
type ListJobDefinitionsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu".
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// Job execution state.
	JobExecutionState ListJobDefinitionsJobExecutionStateEnum `mandatory:"false" contributesTo:"query" name:"jobExecutionState" omitEmpty:"true"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListJobDefinitionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Job type.
	JobType ListJobDefinitionsJobTypeEnum `mandatory:"false" contributesTo:"query" name:"jobType" omitEmpty:"true"`

	// Whether job definition is an incremental harvest (true) or a full harvest (false).
	IsIncremental *bool `mandatory:"false" contributesTo:"query" name:"isIncremental"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"false" contributesTo:"query" name:"dataAssetKey"`

	// Unique connection key.
	ConnectionKey *string `mandatory:"false" contributesTo:"query" name:"connectionKey"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// The sample data size in MB, specified as number of rows, for a metadata harvest.
	SampleDataSizeInMBs *string `mandatory:"false" contributesTo:"query" name:"sampleDataSizeInMBs"`

	// Specifies the fields to return in a job definition summary response.
	Fields []ListJobDefinitionsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. Default order for TIMELATESTEXECUTIONSTARTED is descending. If no value is specified TIMECREATED is default.
	SortBy ListJobDefinitionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListJobDefinitionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListJobDefinitionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJobDefinitionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobDefinitionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListJobDefinitionsResponse wrapper for the ListJobDefinitions operation
type ListJobDefinitionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JobDefinitionCollection instances
	JobDefinitionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJobDefinitionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJobDefinitionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJobDefinitionsJobExecutionStateEnum Enum with underlying type: string
type ListJobDefinitionsJobExecutionStateEnum string

// Set of constants representing the allowable values for ListJobDefinitionsJobExecutionStateEnum
const (
	ListJobDefinitionsJobExecutionStateCreated    ListJobDefinitionsJobExecutionStateEnum = "CREATED"
	ListJobDefinitionsJobExecutionStateInProgress ListJobDefinitionsJobExecutionStateEnum = "IN_PROGRESS"
	ListJobDefinitionsJobExecutionStateInactive   ListJobDefinitionsJobExecutionStateEnum = "INACTIVE"
	ListJobDefinitionsJobExecutionStateFailed     ListJobDefinitionsJobExecutionStateEnum = "FAILED"
	ListJobDefinitionsJobExecutionStateSucceeded  ListJobDefinitionsJobExecutionStateEnum = "SUCCEEDED"
	ListJobDefinitionsJobExecutionStateCanceled   ListJobDefinitionsJobExecutionStateEnum = "CANCELED"
)

var mappingListJobDefinitionsJobExecutionState = map[string]ListJobDefinitionsJobExecutionStateEnum{
	"CREATED":     ListJobDefinitionsJobExecutionStateCreated,
	"IN_PROGRESS": ListJobDefinitionsJobExecutionStateInProgress,
	"INACTIVE":    ListJobDefinitionsJobExecutionStateInactive,
	"FAILED":      ListJobDefinitionsJobExecutionStateFailed,
	"SUCCEEDED":   ListJobDefinitionsJobExecutionStateSucceeded,
	"CANCELED":    ListJobDefinitionsJobExecutionStateCanceled,
}

// GetListJobDefinitionsJobExecutionStateEnumValues Enumerates the set of values for ListJobDefinitionsJobExecutionStateEnum
func GetListJobDefinitionsJobExecutionStateEnumValues() []ListJobDefinitionsJobExecutionStateEnum {
	values := make([]ListJobDefinitionsJobExecutionStateEnum, 0)
	for _, v := range mappingListJobDefinitionsJobExecutionState {
		values = append(values, v)
	}
	return values
}

// ListJobDefinitionsLifecycleStateEnum Enum with underlying type: string
type ListJobDefinitionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListJobDefinitionsLifecycleStateEnum
const (
	ListJobDefinitionsLifecycleStateCreating ListJobDefinitionsLifecycleStateEnum = "CREATING"
	ListJobDefinitionsLifecycleStateActive   ListJobDefinitionsLifecycleStateEnum = "ACTIVE"
	ListJobDefinitionsLifecycleStateInactive ListJobDefinitionsLifecycleStateEnum = "INACTIVE"
	ListJobDefinitionsLifecycleStateUpdating ListJobDefinitionsLifecycleStateEnum = "UPDATING"
	ListJobDefinitionsLifecycleStateDeleting ListJobDefinitionsLifecycleStateEnum = "DELETING"
	ListJobDefinitionsLifecycleStateDeleted  ListJobDefinitionsLifecycleStateEnum = "DELETED"
	ListJobDefinitionsLifecycleStateFailed   ListJobDefinitionsLifecycleStateEnum = "FAILED"
	ListJobDefinitionsLifecycleStateMoving   ListJobDefinitionsLifecycleStateEnum = "MOVING"
)

var mappingListJobDefinitionsLifecycleState = map[string]ListJobDefinitionsLifecycleStateEnum{
	"CREATING": ListJobDefinitionsLifecycleStateCreating,
	"ACTIVE":   ListJobDefinitionsLifecycleStateActive,
	"INACTIVE": ListJobDefinitionsLifecycleStateInactive,
	"UPDATING": ListJobDefinitionsLifecycleStateUpdating,
	"DELETING": ListJobDefinitionsLifecycleStateDeleting,
	"DELETED":  ListJobDefinitionsLifecycleStateDeleted,
	"FAILED":   ListJobDefinitionsLifecycleStateFailed,
	"MOVING":   ListJobDefinitionsLifecycleStateMoving,
}

// GetListJobDefinitionsLifecycleStateEnumValues Enumerates the set of values for ListJobDefinitionsLifecycleStateEnum
func GetListJobDefinitionsLifecycleStateEnumValues() []ListJobDefinitionsLifecycleStateEnum {
	values := make([]ListJobDefinitionsLifecycleStateEnum, 0)
	for _, v := range mappingListJobDefinitionsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListJobDefinitionsJobTypeEnum Enum with underlying type: string
type ListJobDefinitionsJobTypeEnum string

// Set of constants representing the allowable values for ListJobDefinitionsJobTypeEnum
const (
	ListJobDefinitionsJobTypeHarvest                    ListJobDefinitionsJobTypeEnum = "HARVEST"
	ListJobDefinitionsJobTypeProfiling                  ListJobDefinitionsJobTypeEnum = "PROFILING"
	ListJobDefinitionsJobTypeSampling                   ListJobDefinitionsJobTypeEnum = "SAMPLING"
	ListJobDefinitionsJobTypePreview                    ListJobDefinitionsJobTypeEnum = "PREVIEW"
	ListJobDefinitionsJobTypeImport                     ListJobDefinitionsJobTypeEnum = "IMPORT"
	ListJobDefinitionsJobTypeExport                     ListJobDefinitionsJobTypeEnum = "EXPORT"
	ListJobDefinitionsJobTypeImportGlossary             ListJobDefinitionsJobTypeEnum = "IMPORT_GLOSSARY"
	ListJobDefinitionsJobTypeExportGlossary             ListJobDefinitionsJobTypeEnum = "EXPORT_GLOSSARY"
	ListJobDefinitionsJobTypeInternal                   ListJobDefinitionsJobTypeEnum = "INTERNAL"
	ListJobDefinitionsJobTypePurge                      ListJobDefinitionsJobTypeEnum = "PURGE"
	ListJobDefinitionsJobTypeImmediate                  ListJobDefinitionsJobTypeEnum = "IMMEDIATE"
	ListJobDefinitionsJobTypeScheduled                  ListJobDefinitionsJobTypeEnum = "SCHEDULED"
	ListJobDefinitionsJobTypeImmediateExecution         ListJobDefinitionsJobTypeEnum = "IMMEDIATE_EXECUTION"
	ListJobDefinitionsJobTypeScheduledExecution         ListJobDefinitionsJobTypeEnum = "SCHEDULED_EXECUTION"
	ListJobDefinitionsJobTypeScheduledExecutionInstance ListJobDefinitionsJobTypeEnum = "SCHEDULED_EXECUTION_INSTANCE"
)

var mappingListJobDefinitionsJobType = map[string]ListJobDefinitionsJobTypeEnum{
	"HARVEST":                      ListJobDefinitionsJobTypeHarvest,
	"PROFILING":                    ListJobDefinitionsJobTypeProfiling,
	"SAMPLING":                     ListJobDefinitionsJobTypeSampling,
	"PREVIEW":                      ListJobDefinitionsJobTypePreview,
	"IMPORT":                       ListJobDefinitionsJobTypeImport,
	"EXPORT":                       ListJobDefinitionsJobTypeExport,
	"IMPORT_GLOSSARY":              ListJobDefinitionsJobTypeImportGlossary,
	"EXPORT_GLOSSARY":              ListJobDefinitionsJobTypeExportGlossary,
	"INTERNAL":                     ListJobDefinitionsJobTypeInternal,
	"PURGE":                        ListJobDefinitionsJobTypePurge,
	"IMMEDIATE":                    ListJobDefinitionsJobTypeImmediate,
	"SCHEDULED":                    ListJobDefinitionsJobTypeScheduled,
	"IMMEDIATE_EXECUTION":          ListJobDefinitionsJobTypeImmediateExecution,
	"SCHEDULED_EXECUTION":          ListJobDefinitionsJobTypeScheduledExecution,
	"SCHEDULED_EXECUTION_INSTANCE": ListJobDefinitionsJobTypeScheduledExecutionInstance,
}

// GetListJobDefinitionsJobTypeEnumValues Enumerates the set of values for ListJobDefinitionsJobTypeEnum
func GetListJobDefinitionsJobTypeEnumValues() []ListJobDefinitionsJobTypeEnum {
	values := make([]ListJobDefinitionsJobTypeEnum, 0)
	for _, v := range mappingListJobDefinitionsJobType {
		values = append(values, v)
	}
	return values
}

// ListJobDefinitionsFieldsEnum Enum with underlying type: string
type ListJobDefinitionsFieldsEnum string

// Set of constants representing the allowable values for ListJobDefinitionsFieldsEnum
const (
	ListJobDefinitionsFieldsKey                        ListJobDefinitionsFieldsEnum = "key"
	ListJobDefinitionsFieldsDisplayname                ListJobDefinitionsFieldsEnum = "displayName"
	ListJobDefinitionsFieldsDescription                ListJobDefinitionsFieldsEnum = "description"
	ListJobDefinitionsFieldsCatalogid                  ListJobDefinitionsFieldsEnum = "catalogId"
	ListJobDefinitionsFieldsJobtype                    ListJobDefinitionsFieldsEnum = "jobType"
	ListJobDefinitionsFieldsConnectionkey              ListJobDefinitionsFieldsEnum = "connectionKey"
	ListJobDefinitionsFieldsLifecyclestate             ListJobDefinitionsFieldsEnum = "lifecycleState"
	ListJobDefinitionsFieldsTimecreated                ListJobDefinitionsFieldsEnum = "timeCreated"
	ListJobDefinitionsFieldsIssampledataextracted      ListJobDefinitionsFieldsEnum = "isSampleDataExtracted"
	ListJobDefinitionsFieldsUri                        ListJobDefinitionsFieldsEnum = "uri"
	ListJobDefinitionsFieldsTimelatestexecutionstarted ListJobDefinitionsFieldsEnum = "timeLatestExecutionStarted"
	ListJobDefinitionsFieldsTimelatestexecutionended   ListJobDefinitionsFieldsEnum = "timeLatestExecutionEnded"
	ListJobDefinitionsFieldsJobexecutionstate          ListJobDefinitionsFieldsEnum = "jobExecutionState"
	ListJobDefinitionsFieldsScheduletype               ListJobDefinitionsFieldsEnum = "scheduleType"
)

var mappingListJobDefinitionsFields = map[string]ListJobDefinitionsFieldsEnum{
	"key":                   ListJobDefinitionsFieldsKey,
	"displayName":           ListJobDefinitionsFieldsDisplayname,
	"description":           ListJobDefinitionsFieldsDescription,
	"catalogId":             ListJobDefinitionsFieldsCatalogid,
	"jobType":               ListJobDefinitionsFieldsJobtype,
	"connectionKey":         ListJobDefinitionsFieldsConnectionkey,
	"lifecycleState":        ListJobDefinitionsFieldsLifecyclestate,
	"timeCreated":           ListJobDefinitionsFieldsTimecreated,
	"isSampleDataExtracted": ListJobDefinitionsFieldsIssampledataextracted,
	"uri": ListJobDefinitionsFieldsUri,
	"timeLatestExecutionStarted": ListJobDefinitionsFieldsTimelatestexecutionstarted,
	"timeLatestExecutionEnded":   ListJobDefinitionsFieldsTimelatestexecutionended,
	"jobExecutionState":          ListJobDefinitionsFieldsJobexecutionstate,
	"scheduleType":               ListJobDefinitionsFieldsScheduletype,
}

// GetListJobDefinitionsFieldsEnumValues Enumerates the set of values for ListJobDefinitionsFieldsEnum
func GetListJobDefinitionsFieldsEnumValues() []ListJobDefinitionsFieldsEnum {
	values := make([]ListJobDefinitionsFieldsEnum, 0)
	for _, v := range mappingListJobDefinitionsFields {
		values = append(values, v)
	}
	return values
}

// ListJobDefinitionsSortByEnum Enum with underlying type: string
type ListJobDefinitionsSortByEnum string

// Set of constants representing the allowable values for ListJobDefinitionsSortByEnum
const (
	ListJobDefinitionsSortByTimecreated                ListJobDefinitionsSortByEnum = "TIMECREATED"
	ListJobDefinitionsSortByDisplayname                ListJobDefinitionsSortByEnum = "DISPLAYNAME"
	ListJobDefinitionsSortByTimelatestexecutionstarted ListJobDefinitionsSortByEnum = "TIMELATESTEXECUTIONSTARTED"
)

var mappingListJobDefinitionsSortBy = map[string]ListJobDefinitionsSortByEnum{
	"TIMECREATED":                ListJobDefinitionsSortByTimecreated,
	"DISPLAYNAME":                ListJobDefinitionsSortByDisplayname,
	"TIMELATESTEXECUTIONSTARTED": ListJobDefinitionsSortByTimelatestexecutionstarted,
}

// GetListJobDefinitionsSortByEnumValues Enumerates the set of values for ListJobDefinitionsSortByEnum
func GetListJobDefinitionsSortByEnumValues() []ListJobDefinitionsSortByEnum {
	values := make([]ListJobDefinitionsSortByEnum, 0)
	for _, v := range mappingListJobDefinitionsSortBy {
		values = append(values, v)
	}
	return values
}

// ListJobDefinitionsSortOrderEnum Enum with underlying type: string
type ListJobDefinitionsSortOrderEnum string

// Set of constants representing the allowable values for ListJobDefinitionsSortOrderEnum
const (
	ListJobDefinitionsSortOrderAsc  ListJobDefinitionsSortOrderEnum = "ASC"
	ListJobDefinitionsSortOrderDesc ListJobDefinitionsSortOrderEnum = "DESC"
)

var mappingListJobDefinitionsSortOrder = map[string]ListJobDefinitionsSortOrderEnum{
	"ASC":  ListJobDefinitionsSortOrderAsc,
	"DESC": ListJobDefinitionsSortOrderDesc,
}

// GetListJobDefinitionsSortOrderEnumValues Enumerates the set of values for ListJobDefinitionsSortOrderEnum
func GetListJobDefinitionsSortOrderEnumValues() []ListJobDefinitionsSortOrderEnum {
	values := make([]ListJobDefinitionsSortOrderEnum, 0)
	for _, v := range mappingListJobDefinitionsSortOrder {
		values = append(values, v)
	}
	return values
}
