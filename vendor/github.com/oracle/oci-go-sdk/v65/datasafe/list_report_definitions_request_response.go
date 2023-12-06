// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListReportDefinitionsRequest wrapper for the ListReportDefinitions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListReportDefinitions.go.html to see an example of how to use ListReportDefinitionsRequest.
type ListReportDefinitionsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListReportDefinitionsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The name of the report definition to query.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListReportDefinitionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting parameter order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListReportDefinitionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A boolean flag indicating to list seeded report definitions. Set this parameter to get list of seeded report definitions.
	IsSeeded *bool `mandatory:"false" contributesTo:"query" name:"isSeeded"`

	// Specifies the name of a resource that provides data for the report. For example  alerts, events.
	DataSource ListReportDefinitionsDataSourceEnum `mandatory:"false" contributesTo:"query" name:"dataSource" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified lifecycle state.
	LifecycleState ListReportDefinitionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified category.
	Category ListReportDefinitionsCategoryEnum `mandatory:"false" contributesTo:"query" name:"category" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReportDefinitionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReportDefinitionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListReportDefinitionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReportDefinitionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListReportDefinitionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListReportDefinitionsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListReportDefinitionsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReportDefinitionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListReportDefinitionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReportDefinitionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListReportDefinitionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReportDefinitionsDataSourceEnum(string(request.DataSource)); !ok && request.DataSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataSource: %s. Supported values are: %s.", request.DataSource, strings.Join(GetListReportDefinitionsDataSourceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReportDefinitionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListReportDefinitionsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReportDefinitionsCategoryEnum(string(request.Category)); !ok && request.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", request.Category, strings.Join(GetListReportDefinitionsCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListReportDefinitionsResponse wrapper for the ListReportDefinitions operation
type ListReportDefinitionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ReportDefinitionCollection instances
	ReportDefinitionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListReportDefinitionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReportDefinitionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListReportDefinitionsAccessLevelEnum Enum with underlying type: string
type ListReportDefinitionsAccessLevelEnum string

// Set of constants representing the allowable values for ListReportDefinitionsAccessLevelEnum
const (
	ListReportDefinitionsAccessLevelRestricted ListReportDefinitionsAccessLevelEnum = "RESTRICTED"
	ListReportDefinitionsAccessLevelAccessible ListReportDefinitionsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListReportDefinitionsAccessLevelEnum = map[string]ListReportDefinitionsAccessLevelEnum{
	"RESTRICTED": ListReportDefinitionsAccessLevelRestricted,
	"ACCESSIBLE": ListReportDefinitionsAccessLevelAccessible,
}

var mappingListReportDefinitionsAccessLevelEnumLowerCase = map[string]ListReportDefinitionsAccessLevelEnum{
	"restricted": ListReportDefinitionsAccessLevelRestricted,
	"accessible": ListReportDefinitionsAccessLevelAccessible,
}

// GetListReportDefinitionsAccessLevelEnumValues Enumerates the set of values for ListReportDefinitionsAccessLevelEnum
func GetListReportDefinitionsAccessLevelEnumValues() []ListReportDefinitionsAccessLevelEnum {
	values := make([]ListReportDefinitionsAccessLevelEnum, 0)
	for _, v := range mappingListReportDefinitionsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportDefinitionsAccessLevelEnumStringValues Enumerates the set of values in String for ListReportDefinitionsAccessLevelEnum
func GetListReportDefinitionsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListReportDefinitionsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportDefinitionsAccessLevelEnum(val string) (ListReportDefinitionsAccessLevelEnum, bool) {
	enum, ok := mappingListReportDefinitionsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReportDefinitionsSortOrderEnum Enum with underlying type: string
type ListReportDefinitionsSortOrderEnum string

// Set of constants representing the allowable values for ListReportDefinitionsSortOrderEnum
const (
	ListReportDefinitionsSortOrderAsc  ListReportDefinitionsSortOrderEnum = "ASC"
	ListReportDefinitionsSortOrderDesc ListReportDefinitionsSortOrderEnum = "DESC"
)

var mappingListReportDefinitionsSortOrderEnum = map[string]ListReportDefinitionsSortOrderEnum{
	"ASC":  ListReportDefinitionsSortOrderAsc,
	"DESC": ListReportDefinitionsSortOrderDesc,
}

var mappingListReportDefinitionsSortOrderEnumLowerCase = map[string]ListReportDefinitionsSortOrderEnum{
	"asc":  ListReportDefinitionsSortOrderAsc,
	"desc": ListReportDefinitionsSortOrderDesc,
}

// GetListReportDefinitionsSortOrderEnumValues Enumerates the set of values for ListReportDefinitionsSortOrderEnum
func GetListReportDefinitionsSortOrderEnumValues() []ListReportDefinitionsSortOrderEnum {
	values := make([]ListReportDefinitionsSortOrderEnum, 0)
	for _, v := range mappingListReportDefinitionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportDefinitionsSortOrderEnumStringValues Enumerates the set of values in String for ListReportDefinitionsSortOrderEnum
func GetListReportDefinitionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListReportDefinitionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportDefinitionsSortOrderEnum(val string) (ListReportDefinitionsSortOrderEnum, bool) {
	enum, ok := mappingListReportDefinitionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReportDefinitionsSortByEnum Enum with underlying type: string
type ListReportDefinitionsSortByEnum string

// Set of constants representing the allowable values for ListReportDefinitionsSortByEnum
const (
	ListReportDefinitionsSortByTimecreated  ListReportDefinitionsSortByEnum = "TIMECREATED"
	ListReportDefinitionsSortByDisplayname  ListReportDefinitionsSortByEnum = "DISPLAYNAME"
	ListReportDefinitionsSortByDisplayorder ListReportDefinitionsSortByEnum = "DISPLAYORDER"
)

var mappingListReportDefinitionsSortByEnum = map[string]ListReportDefinitionsSortByEnum{
	"TIMECREATED":  ListReportDefinitionsSortByTimecreated,
	"DISPLAYNAME":  ListReportDefinitionsSortByDisplayname,
	"DISPLAYORDER": ListReportDefinitionsSortByDisplayorder,
}

var mappingListReportDefinitionsSortByEnumLowerCase = map[string]ListReportDefinitionsSortByEnum{
	"timecreated":  ListReportDefinitionsSortByTimecreated,
	"displayname":  ListReportDefinitionsSortByDisplayname,
	"displayorder": ListReportDefinitionsSortByDisplayorder,
}

// GetListReportDefinitionsSortByEnumValues Enumerates the set of values for ListReportDefinitionsSortByEnum
func GetListReportDefinitionsSortByEnumValues() []ListReportDefinitionsSortByEnum {
	values := make([]ListReportDefinitionsSortByEnum, 0)
	for _, v := range mappingListReportDefinitionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportDefinitionsSortByEnumStringValues Enumerates the set of values in String for ListReportDefinitionsSortByEnum
func GetListReportDefinitionsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
		"DISPLAYORDER",
	}
}

// GetMappingListReportDefinitionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportDefinitionsSortByEnum(val string) (ListReportDefinitionsSortByEnum, bool) {
	enum, ok := mappingListReportDefinitionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReportDefinitionsDataSourceEnum Enum with underlying type: string
type ListReportDefinitionsDataSourceEnum string

// Set of constants representing the allowable values for ListReportDefinitionsDataSourceEnum
const (
	ListReportDefinitionsDataSourceEvents     ListReportDefinitionsDataSourceEnum = "EVENTS"
	ListReportDefinitionsDataSourceAlerts     ListReportDefinitionsDataSourceEnum = "ALERTS"
	ListReportDefinitionsDataSourceViolations ListReportDefinitionsDataSourceEnum = "VIOLATIONS"
	ListReportDefinitionsDataSourceAllowedSql ListReportDefinitionsDataSourceEnum = "ALLOWED_SQL"
)

var mappingListReportDefinitionsDataSourceEnum = map[string]ListReportDefinitionsDataSourceEnum{
	"EVENTS":      ListReportDefinitionsDataSourceEvents,
	"ALERTS":      ListReportDefinitionsDataSourceAlerts,
	"VIOLATIONS":  ListReportDefinitionsDataSourceViolations,
	"ALLOWED_SQL": ListReportDefinitionsDataSourceAllowedSql,
}

var mappingListReportDefinitionsDataSourceEnumLowerCase = map[string]ListReportDefinitionsDataSourceEnum{
	"events":      ListReportDefinitionsDataSourceEvents,
	"alerts":      ListReportDefinitionsDataSourceAlerts,
	"violations":  ListReportDefinitionsDataSourceViolations,
	"allowed_sql": ListReportDefinitionsDataSourceAllowedSql,
}

// GetListReportDefinitionsDataSourceEnumValues Enumerates the set of values for ListReportDefinitionsDataSourceEnum
func GetListReportDefinitionsDataSourceEnumValues() []ListReportDefinitionsDataSourceEnum {
	values := make([]ListReportDefinitionsDataSourceEnum, 0)
	for _, v := range mappingListReportDefinitionsDataSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportDefinitionsDataSourceEnumStringValues Enumerates the set of values in String for ListReportDefinitionsDataSourceEnum
func GetListReportDefinitionsDataSourceEnumStringValues() []string {
	return []string{
		"EVENTS",
		"ALERTS",
		"VIOLATIONS",
		"ALLOWED_SQL",
	}
}

// GetMappingListReportDefinitionsDataSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportDefinitionsDataSourceEnum(val string) (ListReportDefinitionsDataSourceEnum, bool) {
	enum, ok := mappingListReportDefinitionsDataSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReportDefinitionsLifecycleStateEnum Enum with underlying type: string
type ListReportDefinitionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListReportDefinitionsLifecycleStateEnum
const (
	ListReportDefinitionsLifecycleStateCreating ListReportDefinitionsLifecycleStateEnum = "CREATING"
	ListReportDefinitionsLifecycleStateUpdating ListReportDefinitionsLifecycleStateEnum = "UPDATING"
	ListReportDefinitionsLifecycleStateActive   ListReportDefinitionsLifecycleStateEnum = "ACTIVE"
	ListReportDefinitionsLifecycleStateDeleting ListReportDefinitionsLifecycleStateEnum = "DELETING"
	ListReportDefinitionsLifecycleStateDeleted  ListReportDefinitionsLifecycleStateEnum = "DELETED"
)

var mappingListReportDefinitionsLifecycleStateEnum = map[string]ListReportDefinitionsLifecycleStateEnum{
	"CREATING": ListReportDefinitionsLifecycleStateCreating,
	"UPDATING": ListReportDefinitionsLifecycleStateUpdating,
	"ACTIVE":   ListReportDefinitionsLifecycleStateActive,
	"DELETING": ListReportDefinitionsLifecycleStateDeleting,
	"DELETED":  ListReportDefinitionsLifecycleStateDeleted,
}

var mappingListReportDefinitionsLifecycleStateEnumLowerCase = map[string]ListReportDefinitionsLifecycleStateEnum{
	"creating": ListReportDefinitionsLifecycleStateCreating,
	"updating": ListReportDefinitionsLifecycleStateUpdating,
	"active":   ListReportDefinitionsLifecycleStateActive,
	"deleting": ListReportDefinitionsLifecycleStateDeleting,
	"deleted":  ListReportDefinitionsLifecycleStateDeleted,
}

// GetListReportDefinitionsLifecycleStateEnumValues Enumerates the set of values for ListReportDefinitionsLifecycleStateEnum
func GetListReportDefinitionsLifecycleStateEnumValues() []ListReportDefinitionsLifecycleStateEnum {
	values := make([]ListReportDefinitionsLifecycleStateEnum, 0)
	for _, v := range mappingListReportDefinitionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportDefinitionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListReportDefinitionsLifecycleStateEnum
func GetListReportDefinitionsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingListReportDefinitionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportDefinitionsLifecycleStateEnum(val string) (ListReportDefinitionsLifecycleStateEnum, bool) {
	enum, ok := mappingListReportDefinitionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReportDefinitionsCategoryEnum Enum with underlying type: string
type ListReportDefinitionsCategoryEnum string

// Set of constants representing the allowable values for ListReportDefinitionsCategoryEnum
const (
	ListReportDefinitionsCategoryCustomReports    ListReportDefinitionsCategoryEnum = "CUSTOM_REPORTS"
	ListReportDefinitionsCategorySummary          ListReportDefinitionsCategoryEnum = "SUMMARY"
	ListReportDefinitionsCategoryActivityAuditing ListReportDefinitionsCategoryEnum = "ACTIVITY_AUDITING"
)

var mappingListReportDefinitionsCategoryEnum = map[string]ListReportDefinitionsCategoryEnum{
	"CUSTOM_REPORTS":    ListReportDefinitionsCategoryCustomReports,
	"SUMMARY":           ListReportDefinitionsCategorySummary,
	"ACTIVITY_AUDITING": ListReportDefinitionsCategoryActivityAuditing,
}

var mappingListReportDefinitionsCategoryEnumLowerCase = map[string]ListReportDefinitionsCategoryEnum{
	"custom_reports":    ListReportDefinitionsCategoryCustomReports,
	"summary":           ListReportDefinitionsCategorySummary,
	"activity_auditing": ListReportDefinitionsCategoryActivityAuditing,
}

// GetListReportDefinitionsCategoryEnumValues Enumerates the set of values for ListReportDefinitionsCategoryEnum
func GetListReportDefinitionsCategoryEnumValues() []ListReportDefinitionsCategoryEnum {
	values := make([]ListReportDefinitionsCategoryEnum, 0)
	for _, v := range mappingListReportDefinitionsCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportDefinitionsCategoryEnumStringValues Enumerates the set of values in String for ListReportDefinitionsCategoryEnum
func GetListReportDefinitionsCategoryEnumStringValues() []string {
	return []string{
		"CUSTOM_REPORTS",
		"SUMMARY",
		"ACTIVITY_AUDITING",
	}
}

// GetMappingListReportDefinitionsCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportDefinitionsCategoryEnum(val string) (ListReportDefinitionsCategoryEnum, bool) {
	enum, ok := mappingListReportDefinitionsCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
