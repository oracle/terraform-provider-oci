// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasetools

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDatabaseToolsSqlReportsRequest wrapper for the ListDatabaseToolsSqlReports operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsSqlReports.go.html to see an example of how to use ListDatabaseToolsSqlReportsRequest.
type ListDatabaseToolsSqlReportsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources only when their `databaseToolsSqlReportLifecycleState` matches the specified `databaseToolsSqlReportLifecycleState`.
	LifecycleState ListDatabaseToolsSqlReportsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatabaseToolsSqlReportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDatabaseToolsSqlReportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources with one of the specified type values.
	Type []DatabaseToolsSqlReportTypeEnum `contributesTo:"query" name:"type" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseToolsSqlReportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseToolsSqlReportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseToolsSqlReportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseToolsSqlReportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseToolsSqlReportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseToolsSqlReportsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDatabaseToolsSqlReportsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsSqlReportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseToolsSqlReportsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsSqlReportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseToolsSqlReportsSortByEnumStringValues(), ",")))
	}
	for _, val := range request.Type {
		if _, ok := GetMappingDatabaseToolsSqlReportTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", val, strings.Join(GetDatabaseToolsSqlReportTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseToolsSqlReportsResponse wrapper for the ListDatabaseToolsSqlReports operation
type ListDatabaseToolsSqlReportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseToolsSqlReportCollection instances
	DatabaseToolsSqlReportCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseToolsSqlReportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseToolsSqlReportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseToolsSqlReportsLifecycleStateEnum Enum with underlying type: string
type ListDatabaseToolsSqlReportsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDatabaseToolsSqlReportsLifecycleStateEnum
const (
	ListDatabaseToolsSqlReportsLifecycleStateActive  ListDatabaseToolsSqlReportsLifecycleStateEnum = "ACTIVE"
	ListDatabaseToolsSqlReportsLifecycleStateDeleted ListDatabaseToolsSqlReportsLifecycleStateEnum = "DELETED"
)

var mappingListDatabaseToolsSqlReportsLifecycleStateEnum = map[string]ListDatabaseToolsSqlReportsLifecycleStateEnum{
	"ACTIVE":  ListDatabaseToolsSqlReportsLifecycleStateActive,
	"DELETED": ListDatabaseToolsSqlReportsLifecycleStateDeleted,
}

var mappingListDatabaseToolsSqlReportsLifecycleStateEnumLowerCase = map[string]ListDatabaseToolsSqlReportsLifecycleStateEnum{
	"active":  ListDatabaseToolsSqlReportsLifecycleStateActive,
	"deleted": ListDatabaseToolsSqlReportsLifecycleStateDeleted,
}

// GetListDatabaseToolsSqlReportsLifecycleStateEnumValues Enumerates the set of values for ListDatabaseToolsSqlReportsLifecycleStateEnum
func GetListDatabaseToolsSqlReportsLifecycleStateEnumValues() []ListDatabaseToolsSqlReportsLifecycleStateEnum {
	values := make([]ListDatabaseToolsSqlReportsLifecycleStateEnum, 0)
	for _, v := range mappingListDatabaseToolsSqlReportsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsSqlReportsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDatabaseToolsSqlReportsLifecycleStateEnum
func GetListDatabaseToolsSqlReportsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListDatabaseToolsSqlReportsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsSqlReportsLifecycleStateEnum(val string) (ListDatabaseToolsSqlReportsLifecycleStateEnum, bool) {
	enum, ok := mappingListDatabaseToolsSqlReportsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsSqlReportsSortOrderEnum Enum with underlying type: string
type ListDatabaseToolsSqlReportsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseToolsSqlReportsSortOrderEnum
const (
	ListDatabaseToolsSqlReportsSortOrderAsc  ListDatabaseToolsSqlReportsSortOrderEnum = "ASC"
	ListDatabaseToolsSqlReportsSortOrderDesc ListDatabaseToolsSqlReportsSortOrderEnum = "DESC"
)

var mappingListDatabaseToolsSqlReportsSortOrderEnum = map[string]ListDatabaseToolsSqlReportsSortOrderEnum{
	"ASC":  ListDatabaseToolsSqlReportsSortOrderAsc,
	"DESC": ListDatabaseToolsSqlReportsSortOrderDesc,
}

var mappingListDatabaseToolsSqlReportsSortOrderEnumLowerCase = map[string]ListDatabaseToolsSqlReportsSortOrderEnum{
	"asc":  ListDatabaseToolsSqlReportsSortOrderAsc,
	"desc": ListDatabaseToolsSqlReportsSortOrderDesc,
}

// GetListDatabaseToolsSqlReportsSortOrderEnumValues Enumerates the set of values for ListDatabaseToolsSqlReportsSortOrderEnum
func GetListDatabaseToolsSqlReportsSortOrderEnumValues() []ListDatabaseToolsSqlReportsSortOrderEnum {
	values := make([]ListDatabaseToolsSqlReportsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseToolsSqlReportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsSqlReportsSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseToolsSqlReportsSortOrderEnum
func GetListDatabaseToolsSqlReportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseToolsSqlReportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsSqlReportsSortOrderEnum(val string) (ListDatabaseToolsSqlReportsSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseToolsSqlReportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsSqlReportsSortByEnum Enum with underlying type: string
type ListDatabaseToolsSqlReportsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseToolsSqlReportsSortByEnum
const (
	ListDatabaseToolsSqlReportsSortByTimecreated ListDatabaseToolsSqlReportsSortByEnum = "timeCreated"
	ListDatabaseToolsSqlReportsSortByDisplayname ListDatabaseToolsSqlReportsSortByEnum = "displayName"
)

var mappingListDatabaseToolsSqlReportsSortByEnum = map[string]ListDatabaseToolsSqlReportsSortByEnum{
	"timeCreated": ListDatabaseToolsSqlReportsSortByTimecreated,
	"displayName": ListDatabaseToolsSqlReportsSortByDisplayname,
}

var mappingListDatabaseToolsSqlReportsSortByEnumLowerCase = map[string]ListDatabaseToolsSqlReportsSortByEnum{
	"timecreated": ListDatabaseToolsSqlReportsSortByTimecreated,
	"displayname": ListDatabaseToolsSqlReportsSortByDisplayname,
}

// GetListDatabaseToolsSqlReportsSortByEnumValues Enumerates the set of values for ListDatabaseToolsSqlReportsSortByEnum
func GetListDatabaseToolsSqlReportsSortByEnumValues() []ListDatabaseToolsSqlReportsSortByEnum {
	values := make([]ListDatabaseToolsSqlReportsSortByEnum, 0)
	for _, v := range mappingListDatabaseToolsSqlReportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsSqlReportsSortByEnumStringValues Enumerates the set of values in String for ListDatabaseToolsSqlReportsSortByEnum
func GetListDatabaseToolsSqlReportsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatabaseToolsSqlReportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsSqlReportsSortByEnum(val string) (ListDatabaseToolsSqlReportsSortByEnum, bool) {
	enum, ok := mappingListDatabaseToolsSqlReportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
