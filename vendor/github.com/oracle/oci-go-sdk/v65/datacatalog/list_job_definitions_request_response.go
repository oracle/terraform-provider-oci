// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListJobDefinitionsRequest wrapper for the ListJobDefinitions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListJobDefinitions.go.html to see an example of how to use ListJobDefinitionsRequest.
type ListJobDefinitionsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
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

	// Unique glossary key.
	GlossaryKey *string `mandatory:"false" contributesTo:"query" name:"glossaryKey"`

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
func (request ListJobDefinitionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJobDefinitionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobDefinitionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJobDefinitionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJobDefinitionsJobExecutionStateEnum(string(request.JobExecutionState)); !ok && request.JobExecutionState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobExecutionState: %s. Supported values are: %s.", request.JobExecutionState, strings.Join(GetListJobDefinitionsJobExecutionStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobDefinitionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListJobDefinitionsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobDefinitionsJobTypeEnum(string(request.JobType)); !ok && request.JobType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobType: %s. Supported values are: %s.", request.JobType, strings.Join(GetListJobDefinitionsJobTypeEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListJobDefinitionsFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListJobDefinitionsFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListJobDefinitionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJobDefinitionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobDefinitionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJobDefinitionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
	ListJobDefinitionsJobExecutionStateCreated               ListJobDefinitionsJobExecutionStateEnum = "CREATED"
	ListJobDefinitionsJobExecutionStateInProgress            ListJobDefinitionsJobExecutionStateEnum = "IN_PROGRESS"
	ListJobDefinitionsJobExecutionStateInactive              ListJobDefinitionsJobExecutionStateEnum = "INACTIVE"
	ListJobDefinitionsJobExecutionStateFailed                ListJobDefinitionsJobExecutionStateEnum = "FAILED"
	ListJobDefinitionsJobExecutionStateSucceeded             ListJobDefinitionsJobExecutionStateEnum = "SUCCEEDED"
	ListJobDefinitionsJobExecutionStateCanceled              ListJobDefinitionsJobExecutionStateEnum = "CANCELED"
	ListJobDefinitionsJobExecutionStateSucceededWithWarnings ListJobDefinitionsJobExecutionStateEnum = "SUCCEEDED_WITH_WARNINGS"
)

var mappingListJobDefinitionsJobExecutionStateEnum = map[string]ListJobDefinitionsJobExecutionStateEnum{
	"CREATED":                 ListJobDefinitionsJobExecutionStateCreated,
	"IN_PROGRESS":             ListJobDefinitionsJobExecutionStateInProgress,
	"INACTIVE":                ListJobDefinitionsJobExecutionStateInactive,
	"FAILED":                  ListJobDefinitionsJobExecutionStateFailed,
	"SUCCEEDED":               ListJobDefinitionsJobExecutionStateSucceeded,
	"CANCELED":                ListJobDefinitionsJobExecutionStateCanceled,
	"SUCCEEDED_WITH_WARNINGS": ListJobDefinitionsJobExecutionStateSucceededWithWarnings,
}

var mappingListJobDefinitionsJobExecutionStateEnumLowerCase = map[string]ListJobDefinitionsJobExecutionStateEnum{
	"created":                 ListJobDefinitionsJobExecutionStateCreated,
	"in_progress":             ListJobDefinitionsJobExecutionStateInProgress,
	"inactive":                ListJobDefinitionsJobExecutionStateInactive,
	"failed":                  ListJobDefinitionsJobExecutionStateFailed,
	"succeeded":               ListJobDefinitionsJobExecutionStateSucceeded,
	"canceled":                ListJobDefinitionsJobExecutionStateCanceled,
	"succeeded_with_warnings": ListJobDefinitionsJobExecutionStateSucceededWithWarnings,
}

// GetListJobDefinitionsJobExecutionStateEnumValues Enumerates the set of values for ListJobDefinitionsJobExecutionStateEnum
func GetListJobDefinitionsJobExecutionStateEnumValues() []ListJobDefinitionsJobExecutionStateEnum {
	values := make([]ListJobDefinitionsJobExecutionStateEnum, 0)
	for _, v := range mappingListJobDefinitionsJobExecutionStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobDefinitionsJobExecutionStateEnumStringValues Enumerates the set of values in String for ListJobDefinitionsJobExecutionStateEnum
func GetListJobDefinitionsJobExecutionStateEnumStringValues() []string {
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

// GetMappingListJobDefinitionsJobExecutionStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobDefinitionsJobExecutionStateEnum(val string) (ListJobDefinitionsJobExecutionStateEnum, bool) {
	enum, ok := mappingListJobDefinitionsJobExecutionStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingListJobDefinitionsLifecycleStateEnum = map[string]ListJobDefinitionsLifecycleStateEnum{
	"CREATING": ListJobDefinitionsLifecycleStateCreating,
	"ACTIVE":   ListJobDefinitionsLifecycleStateActive,
	"INACTIVE": ListJobDefinitionsLifecycleStateInactive,
	"UPDATING": ListJobDefinitionsLifecycleStateUpdating,
	"DELETING": ListJobDefinitionsLifecycleStateDeleting,
	"DELETED":  ListJobDefinitionsLifecycleStateDeleted,
	"FAILED":   ListJobDefinitionsLifecycleStateFailed,
	"MOVING":   ListJobDefinitionsLifecycleStateMoving,
}

var mappingListJobDefinitionsLifecycleStateEnumLowerCase = map[string]ListJobDefinitionsLifecycleStateEnum{
	"creating": ListJobDefinitionsLifecycleStateCreating,
	"active":   ListJobDefinitionsLifecycleStateActive,
	"inactive": ListJobDefinitionsLifecycleStateInactive,
	"updating": ListJobDefinitionsLifecycleStateUpdating,
	"deleting": ListJobDefinitionsLifecycleStateDeleting,
	"deleted":  ListJobDefinitionsLifecycleStateDeleted,
	"failed":   ListJobDefinitionsLifecycleStateFailed,
	"moving":   ListJobDefinitionsLifecycleStateMoving,
}

// GetListJobDefinitionsLifecycleStateEnumValues Enumerates the set of values for ListJobDefinitionsLifecycleStateEnum
func GetListJobDefinitionsLifecycleStateEnumValues() []ListJobDefinitionsLifecycleStateEnum {
	values := make([]ListJobDefinitionsLifecycleStateEnum, 0)
	for _, v := range mappingListJobDefinitionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobDefinitionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListJobDefinitionsLifecycleStateEnum
func GetListJobDefinitionsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"MOVING",
	}
}

// GetMappingListJobDefinitionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobDefinitionsLifecycleStateEnum(val string) (ListJobDefinitionsLifecycleStateEnum, bool) {
	enum, ok := mappingListJobDefinitionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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
	ListJobDefinitionsJobTypeAsyncDelete                ListJobDefinitionsJobTypeEnum = "ASYNC_DELETE"
	ListJobDefinitionsJobTypeImportDataAsset            ListJobDefinitionsJobTypeEnum = "IMPORT_DATA_ASSET"
	ListJobDefinitionsJobTypeCreateScanProxy            ListJobDefinitionsJobTypeEnum = "CREATE_SCAN_PROXY"
	ListJobDefinitionsJobTypeAsyncExportGlossary        ListJobDefinitionsJobTypeEnum = "ASYNC_EXPORT_GLOSSARY"
)

var mappingListJobDefinitionsJobTypeEnum = map[string]ListJobDefinitionsJobTypeEnum{
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
	"ASYNC_DELETE":                 ListJobDefinitionsJobTypeAsyncDelete,
	"IMPORT_DATA_ASSET":            ListJobDefinitionsJobTypeImportDataAsset,
	"CREATE_SCAN_PROXY":            ListJobDefinitionsJobTypeCreateScanProxy,
	"ASYNC_EXPORT_GLOSSARY":        ListJobDefinitionsJobTypeAsyncExportGlossary,
}

var mappingListJobDefinitionsJobTypeEnumLowerCase = map[string]ListJobDefinitionsJobTypeEnum{
	"harvest":                      ListJobDefinitionsJobTypeHarvest,
	"profiling":                    ListJobDefinitionsJobTypeProfiling,
	"sampling":                     ListJobDefinitionsJobTypeSampling,
	"preview":                      ListJobDefinitionsJobTypePreview,
	"import":                       ListJobDefinitionsJobTypeImport,
	"export":                       ListJobDefinitionsJobTypeExport,
	"import_glossary":              ListJobDefinitionsJobTypeImportGlossary,
	"export_glossary":              ListJobDefinitionsJobTypeExportGlossary,
	"internal":                     ListJobDefinitionsJobTypeInternal,
	"purge":                        ListJobDefinitionsJobTypePurge,
	"immediate":                    ListJobDefinitionsJobTypeImmediate,
	"scheduled":                    ListJobDefinitionsJobTypeScheduled,
	"immediate_execution":          ListJobDefinitionsJobTypeImmediateExecution,
	"scheduled_execution":          ListJobDefinitionsJobTypeScheduledExecution,
	"scheduled_execution_instance": ListJobDefinitionsJobTypeScheduledExecutionInstance,
	"async_delete":                 ListJobDefinitionsJobTypeAsyncDelete,
	"import_data_asset":            ListJobDefinitionsJobTypeImportDataAsset,
	"create_scan_proxy":            ListJobDefinitionsJobTypeCreateScanProxy,
	"async_export_glossary":        ListJobDefinitionsJobTypeAsyncExportGlossary,
}

// GetListJobDefinitionsJobTypeEnumValues Enumerates the set of values for ListJobDefinitionsJobTypeEnum
func GetListJobDefinitionsJobTypeEnumValues() []ListJobDefinitionsJobTypeEnum {
	values := make([]ListJobDefinitionsJobTypeEnum, 0)
	for _, v := range mappingListJobDefinitionsJobTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobDefinitionsJobTypeEnumStringValues Enumerates the set of values in String for ListJobDefinitionsJobTypeEnum
func GetListJobDefinitionsJobTypeEnumStringValues() []string {
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
		"CREATE_SCAN_PROXY",
		"ASYNC_EXPORT_GLOSSARY",
	}
}

// GetMappingListJobDefinitionsJobTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobDefinitionsJobTypeEnum(val string) (ListJobDefinitionsJobTypeEnum, bool) {
	enum, ok := mappingListJobDefinitionsJobTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingListJobDefinitionsFieldsEnum = map[string]ListJobDefinitionsFieldsEnum{
	"key":                        ListJobDefinitionsFieldsKey,
	"displayName":                ListJobDefinitionsFieldsDisplayname,
	"description":                ListJobDefinitionsFieldsDescription,
	"catalogId":                  ListJobDefinitionsFieldsCatalogid,
	"jobType":                    ListJobDefinitionsFieldsJobtype,
	"connectionKey":              ListJobDefinitionsFieldsConnectionkey,
	"lifecycleState":             ListJobDefinitionsFieldsLifecyclestate,
	"timeCreated":                ListJobDefinitionsFieldsTimecreated,
	"isSampleDataExtracted":      ListJobDefinitionsFieldsIssampledataextracted,
	"uri":                        ListJobDefinitionsFieldsUri,
	"timeLatestExecutionStarted": ListJobDefinitionsFieldsTimelatestexecutionstarted,
	"timeLatestExecutionEnded":   ListJobDefinitionsFieldsTimelatestexecutionended,
	"jobExecutionState":          ListJobDefinitionsFieldsJobexecutionstate,
	"scheduleType":               ListJobDefinitionsFieldsScheduletype,
}

var mappingListJobDefinitionsFieldsEnumLowerCase = map[string]ListJobDefinitionsFieldsEnum{
	"key":                        ListJobDefinitionsFieldsKey,
	"displayname":                ListJobDefinitionsFieldsDisplayname,
	"description":                ListJobDefinitionsFieldsDescription,
	"catalogid":                  ListJobDefinitionsFieldsCatalogid,
	"jobtype":                    ListJobDefinitionsFieldsJobtype,
	"connectionkey":              ListJobDefinitionsFieldsConnectionkey,
	"lifecyclestate":             ListJobDefinitionsFieldsLifecyclestate,
	"timecreated":                ListJobDefinitionsFieldsTimecreated,
	"issampledataextracted":      ListJobDefinitionsFieldsIssampledataextracted,
	"uri":                        ListJobDefinitionsFieldsUri,
	"timelatestexecutionstarted": ListJobDefinitionsFieldsTimelatestexecutionstarted,
	"timelatestexecutionended":   ListJobDefinitionsFieldsTimelatestexecutionended,
	"jobexecutionstate":          ListJobDefinitionsFieldsJobexecutionstate,
	"scheduletype":               ListJobDefinitionsFieldsScheduletype,
}

// GetListJobDefinitionsFieldsEnumValues Enumerates the set of values for ListJobDefinitionsFieldsEnum
func GetListJobDefinitionsFieldsEnumValues() []ListJobDefinitionsFieldsEnum {
	values := make([]ListJobDefinitionsFieldsEnum, 0)
	for _, v := range mappingListJobDefinitionsFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobDefinitionsFieldsEnumStringValues Enumerates the set of values in String for ListJobDefinitionsFieldsEnum
func GetListJobDefinitionsFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"catalogId",
		"jobType",
		"connectionKey",
		"lifecycleState",
		"timeCreated",
		"isSampleDataExtracted",
		"uri",
		"timeLatestExecutionStarted",
		"timeLatestExecutionEnded",
		"jobExecutionState",
		"scheduleType",
	}
}

// GetMappingListJobDefinitionsFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobDefinitionsFieldsEnum(val string) (ListJobDefinitionsFieldsEnum, bool) {
	enum, ok := mappingListJobDefinitionsFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobDefinitionsSortByEnum Enum with underlying type: string
type ListJobDefinitionsSortByEnum string

// Set of constants representing the allowable values for ListJobDefinitionsSortByEnum
const (
	ListJobDefinitionsSortByTimecreated                ListJobDefinitionsSortByEnum = "TIMECREATED"
	ListJobDefinitionsSortByDisplayname                ListJobDefinitionsSortByEnum = "DISPLAYNAME"
	ListJobDefinitionsSortByTimelatestexecutionstarted ListJobDefinitionsSortByEnum = "TIMELATESTEXECUTIONSTARTED"
)

var mappingListJobDefinitionsSortByEnum = map[string]ListJobDefinitionsSortByEnum{
	"TIMECREATED":                ListJobDefinitionsSortByTimecreated,
	"DISPLAYNAME":                ListJobDefinitionsSortByDisplayname,
	"TIMELATESTEXECUTIONSTARTED": ListJobDefinitionsSortByTimelatestexecutionstarted,
}

var mappingListJobDefinitionsSortByEnumLowerCase = map[string]ListJobDefinitionsSortByEnum{
	"timecreated":                ListJobDefinitionsSortByTimecreated,
	"displayname":                ListJobDefinitionsSortByDisplayname,
	"timelatestexecutionstarted": ListJobDefinitionsSortByTimelatestexecutionstarted,
}

// GetListJobDefinitionsSortByEnumValues Enumerates the set of values for ListJobDefinitionsSortByEnum
func GetListJobDefinitionsSortByEnumValues() []ListJobDefinitionsSortByEnum {
	values := make([]ListJobDefinitionsSortByEnum, 0)
	for _, v := range mappingListJobDefinitionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobDefinitionsSortByEnumStringValues Enumerates the set of values in String for ListJobDefinitionsSortByEnum
func GetListJobDefinitionsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
		"TIMELATESTEXECUTIONSTARTED",
	}
}

// GetMappingListJobDefinitionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobDefinitionsSortByEnum(val string) (ListJobDefinitionsSortByEnum, bool) {
	enum, ok := mappingListJobDefinitionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobDefinitionsSortOrderEnum Enum with underlying type: string
type ListJobDefinitionsSortOrderEnum string

// Set of constants representing the allowable values for ListJobDefinitionsSortOrderEnum
const (
	ListJobDefinitionsSortOrderAsc  ListJobDefinitionsSortOrderEnum = "ASC"
	ListJobDefinitionsSortOrderDesc ListJobDefinitionsSortOrderEnum = "DESC"
)

var mappingListJobDefinitionsSortOrderEnum = map[string]ListJobDefinitionsSortOrderEnum{
	"ASC":  ListJobDefinitionsSortOrderAsc,
	"DESC": ListJobDefinitionsSortOrderDesc,
}

var mappingListJobDefinitionsSortOrderEnumLowerCase = map[string]ListJobDefinitionsSortOrderEnum{
	"asc":  ListJobDefinitionsSortOrderAsc,
	"desc": ListJobDefinitionsSortOrderDesc,
}

// GetListJobDefinitionsSortOrderEnumValues Enumerates the set of values for ListJobDefinitionsSortOrderEnum
func GetListJobDefinitionsSortOrderEnumValues() []ListJobDefinitionsSortOrderEnum {
	values := make([]ListJobDefinitionsSortOrderEnum, 0)
	for _, v := range mappingListJobDefinitionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobDefinitionsSortOrderEnumStringValues Enumerates the set of values in String for ListJobDefinitionsSortOrderEnum
func GetListJobDefinitionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJobDefinitionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobDefinitionsSortOrderEnum(val string) (ListJobDefinitionsSortOrderEnum, bool) {
	enum, ok := mappingListJobDefinitionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
