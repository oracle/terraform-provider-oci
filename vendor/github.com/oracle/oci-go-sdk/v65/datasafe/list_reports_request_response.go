// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListReportsRequest wrapper for the ListReports operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListReports.go.html to see an example of how to use ListReportsRequest.
type ListReportsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListReportsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The name of the report definition to query.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListReportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeGenerated is descending. Default order for displayName is ascending. If no value is specified timeGenerated is default.
	SortBy ListReportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The ID of the report definition to filter the list of reports
	ReportDefinitionId *string `mandatory:"false" contributesTo:"query" name:"reportDefinitionId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// An optional filter to return only resources that match the specified lifecycle state.
	LifecycleState ListReportsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified type.
	Type ListReportsTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListReportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListReportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListReportsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListReportsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListReportsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListReportsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReportsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListReportsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReportsTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListReportsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListReportsResponse wrapper for the ListReports operation
type ListReportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ReportCollection instances
	ReportCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListReportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListReportsAccessLevelEnum Enum with underlying type: string
type ListReportsAccessLevelEnum string

// Set of constants representing the allowable values for ListReportsAccessLevelEnum
const (
	ListReportsAccessLevelRestricted ListReportsAccessLevelEnum = "RESTRICTED"
	ListReportsAccessLevelAccessible ListReportsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListReportsAccessLevelEnum = map[string]ListReportsAccessLevelEnum{
	"RESTRICTED": ListReportsAccessLevelRestricted,
	"ACCESSIBLE": ListReportsAccessLevelAccessible,
}

var mappingListReportsAccessLevelEnumLowerCase = map[string]ListReportsAccessLevelEnum{
	"restricted": ListReportsAccessLevelRestricted,
	"accessible": ListReportsAccessLevelAccessible,
}

// GetListReportsAccessLevelEnumValues Enumerates the set of values for ListReportsAccessLevelEnum
func GetListReportsAccessLevelEnumValues() []ListReportsAccessLevelEnum {
	values := make([]ListReportsAccessLevelEnum, 0)
	for _, v := range mappingListReportsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportsAccessLevelEnumStringValues Enumerates the set of values in String for ListReportsAccessLevelEnum
func GetListReportsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListReportsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportsAccessLevelEnum(val string) (ListReportsAccessLevelEnum, bool) {
	enum, ok := mappingListReportsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReportsSortOrderEnum Enum with underlying type: string
type ListReportsSortOrderEnum string

// Set of constants representing the allowable values for ListReportsSortOrderEnum
const (
	ListReportsSortOrderAsc  ListReportsSortOrderEnum = "ASC"
	ListReportsSortOrderDesc ListReportsSortOrderEnum = "DESC"
)

var mappingListReportsSortOrderEnum = map[string]ListReportsSortOrderEnum{
	"ASC":  ListReportsSortOrderAsc,
	"DESC": ListReportsSortOrderDesc,
}

var mappingListReportsSortOrderEnumLowerCase = map[string]ListReportsSortOrderEnum{
	"asc":  ListReportsSortOrderAsc,
	"desc": ListReportsSortOrderDesc,
}

// GetListReportsSortOrderEnumValues Enumerates the set of values for ListReportsSortOrderEnum
func GetListReportsSortOrderEnumValues() []ListReportsSortOrderEnum {
	values := make([]ListReportsSortOrderEnum, 0)
	for _, v := range mappingListReportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportsSortOrderEnumStringValues Enumerates the set of values in String for ListReportsSortOrderEnum
func GetListReportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListReportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportsSortOrderEnum(val string) (ListReportsSortOrderEnum, bool) {
	enum, ok := mappingListReportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReportsSortByEnum Enum with underlying type: string
type ListReportsSortByEnum string

// Set of constants representing the allowable values for ListReportsSortByEnum
const (
	ListReportsSortByTimegenerated ListReportsSortByEnum = "timeGenerated"
	ListReportsSortByDisplayname   ListReportsSortByEnum = "displayName"
)

var mappingListReportsSortByEnum = map[string]ListReportsSortByEnum{
	"timeGenerated": ListReportsSortByTimegenerated,
	"displayName":   ListReportsSortByDisplayname,
}

var mappingListReportsSortByEnumLowerCase = map[string]ListReportsSortByEnum{
	"timegenerated": ListReportsSortByTimegenerated,
	"displayname":   ListReportsSortByDisplayname,
}

// GetListReportsSortByEnumValues Enumerates the set of values for ListReportsSortByEnum
func GetListReportsSortByEnumValues() []ListReportsSortByEnum {
	values := make([]ListReportsSortByEnum, 0)
	for _, v := range mappingListReportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportsSortByEnumStringValues Enumerates the set of values in String for ListReportsSortByEnum
func GetListReportsSortByEnumStringValues() []string {
	return []string{
		"timeGenerated",
		"displayName",
	}
}

// GetMappingListReportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportsSortByEnum(val string) (ListReportsSortByEnum, bool) {
	enum, ok := mappingListReportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReportsLifecycleStateEnum Enum with underlying type: string
type ListReportsLifecycleStateEnum string

// Set of constants representing the allowable values for ListReportsLifecycleStateEnum
const (
	ListReportsLifecycleStateUpdating ListReportsLifecycleStateEnum = "UPDATING"
	ListReportsLifecycleStateActive   ListReportsLifecycleStateEnum = "ACTIVE"
)

var mappingListReportsLifecycleStateEnum = map[string]ListReportsLifecycleStateEnum{
	"UPDATING": ListReportsLifecycleStateUpdating,
	"ACTIVE":   ListReportsLifecycleStateActive,
}

var mappingListReportsLifecycleStateEnumLowerCase = map[string]ListReportsLifecycleStateEnum{
	"updating": ListReportsLifecycleStateUpdating,
	"active":   ListReportsLifecycleStateActive,
}

// GetListReportsLifecycleStateEnumValues Enumerates the set of values for ListReportsLifecycleStateEnum
func GetListReportsLifecycleStateEnumValues() []ListReportsLifecycleStateEnum {
	values := make([]ListReportsLifecycleStateEnum, 0)
	for _, v := range mappingListReportsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportsLifecycleStateEnumStringValues Enumerates the set of values in String for ListReportsLifecycleStateEnum
func GetListReportsLifecycleStateEnumStringValues() []string {
	return []string{
		"UPDATING",
		"ACTIVE",
	}
}

// GetMappingListReportsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportsLifecycleStateEnum(val string) (ListReportsLifecycleStateEnum, bool) {
	enum, ok := mappingListReportsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReportsTypeEnum Enum with underlying type: string
type ListReportsTypeEnum string

// Set of constants representing the allowable values for ListReportsTypeEnum
const (
	ListReportsTypeGenerated ListReportsTypeEnum = "GENERATED"
	ListReportsTypeScheduled ListReportsTypeEnum = "SCHEDULED"
)

var mappingListReportsTypeEnum = map[string]ListReportsTypeEnum{
	"GENERATED": ListReportsTypeGenerated,
	"SCHEDULED": ListReportsTypeScheduled,
}

var mappingListReportsTypeEnumLowerCase = map[string]ListReportsTypeEnum{
	"generated": ListReportsTypeGenerated,
	"scheduled": ListReportsTypeScheduled,
}

// GetListReportsTypeEnumValues Enumerates the set of values for ListReportsTypeEnum
func GetListReportsTypeEnumValues() []ListReportsTypeEnum {
	values := make([]ListReportsTypeEnum, 0)
	for _, v := range mappingListReportsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportsTypeEnumStringValues Enumerates the set of values in String for ListReportsTypeEnum
func GetListReportsTypeEnumStringValues() []string {
	return []string{
		"GENERATED",
		"SCHEDULED",
	}
}

// GetMappingListReportsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportsTypeEnum(val string) (ListReportsTypeEnum, bool) {
	enum, ok := mappingListReportsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
