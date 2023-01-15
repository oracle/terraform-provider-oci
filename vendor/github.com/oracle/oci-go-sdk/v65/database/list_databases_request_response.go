// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDatabasesRequest wrapper for the ListDatabases operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDatabases.go.html to see an example of how to use ListDatabasesRequest.
type ListDatabasesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A Database Home OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DbHomeId *string `mandatory:"false" contributesTo:"query" name:"dbHomeId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata DB system that you want to filter the database results by. Applies only to Exadata DB systems.
	SystemId *string `mandatory:"false" contributesTo:"query" name:"systemId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DBNAME is ascending. The DBNAME sort order is case sensitive.
	SortBy ListDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState DatabaseSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire database name given. The match is not case sensitive.
	DbName *string `mandatory:"false" contributesTo:"query" name:"dbName"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabasesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabasesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDatabaseSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabasesResponse wrapper for the ListDatabases operation
type ListDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DatabaseSummary instances
	Items []DatabaseSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabasesSortByEnum Enum with underlying type: string
type ListDatabasesSortByEnum string

// Set of constants representing the allowable values for ListDatabasesSortByEnum
const (
	ListDatabasesSortByDbname      ListDatabasesSortByEnum = "DBNAME"
	ListDatabasesSortByTimecreated ListDatabasesSortByEnum = "TIMECREATED"
)

var mappingListDatabasesSortByEnum = map[string]ListDatabasesSortByEnum{
	"DBNAME":      ListDatabasesSortByDbname,
	"TIMECREATED": ListDatabasesSortByTimecreated,
}

var mappingListDatabasesSortByEnumLowerCase = map[string]ListDatabasesSortByEnum{
	"dbname":      ListDatabasesSortByDbname,
	"timecreated": ListDatabasesSortByTimecreated,
}

// GetListDatabasesSortByEnumValues Enumerates the set of values for ListDatabasesSortByEnum
func GetListDatabasesSortByEnumValues() []ListDatabasesSortByEnum {
	values := make([]ListDatabasesSortByEnum, 0)
	for _, v := range mappingListDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabasesSortByEnumStringValues Enumerates the set of values in String for ListDatabasesSortByEnum
func GetListDatabasesSortByEnumStringValues() []string {
	return []string{
		"DBNAME",
		"TIMECREATED",
	}
}

// GetMappingListDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabasesSortByEnum(val string) (ListDatabasesSortByEnum, bool) {
	enum, ok := mappingListDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabasesSortOrderEnum Enum with underlying type: string
type ListDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListDatabasesSortOrderEnum
const (
	ListDatabasesSortOrderAsc  ListDatabasesSortOrderEnum = "ASC"
	ListDatabasesSortOrderDesc ListDatabasesSortOrderEnum = "DESC"
)

var mappingListDatabasesSortOrderEnum = map[string]ListDatabasesSortOrderEnum{
	"ASC":  ListDatabasesSortOrderAsc,
	"DESC": ListDatabasesSortOrderDesc,
}

var mappingListDatabasesSortOrderEnumLowerCase = map[string]ListDatabasesSortOrderEnum{
	"asc":  ListDatabasesSortOrderAsc,
	"desc": ListDatabasesSortOrderDesc,
}

// GetListDatabasesSortOrderEnumValues Enumerates the set of values for ListDatabasesSortOrderEnum
func GetListDatabasesSortOrderEnumValues() []ListDatabasesSortOrderEnum {
	values := make([]ListDatabasesSortOrderEnum, 0)
	for _, v := range mappingListDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListDatabasesSortOrderEnum
func GetListDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabasesSortOrderEnum(val string) (ListDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
