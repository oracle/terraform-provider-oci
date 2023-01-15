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

// ListExternalNonContainerDatabasesRequest wrapper for the ListExternalNonContainerDatabases operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExternalNonContainerDatabases.go.html to see an example of how to use ListExternalNonContainerDatabasesRequest.
type ListExternalNonContainerDatabasesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for TIMECREATED is descending.
	// Default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListExternalNonContainerDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListExternalNonContainerDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the specified lifecycle state.
	LifecycleState ExternalDatabaseBaseLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalNonContainerDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalNonContainerDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalNonContainerDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalNonContainerDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalNonContainerDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalNonContainerDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalNonContainerDatabasesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalNonContainerDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalNonContainerDatabasesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExternalDatabaseBaseLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetExternalDatabaseBaseLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalNonContainerDatabasesResponse wrapper for the ListExternalNonContainerDatabases operation
type ListExternalNonContainerDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExternalNonContainerDatabaseSummary instances
	Items []ExternalNonContainerDatabaseSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalNonContainerDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalNonContainerDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalNonContainerDatabasesSortByEnum Enum with underlying type: string
type ListExternalNonContainerDatabasesSortByEnum string

// Set of constants representing the allowable values for ListExternalNonContainerDatabasesSortByEnum
const (
	ListExternalNonContainerDatabasesSortByDisplayname ListExternalNonContainerDatabasesSortByEnum = "DISPLAYNAME"
	ListExternalNonContainerDatabasesSortByTimecreated ListExternalNonContainerDatabasesSortByEnum = "TIMECREATED"
)

var mappingListExternalNonContainerDatabasesSortByEnum = map[string]ListExternalNonContainerDatabasesSortByEnum{
	"DISPLAYNAME": ListExternalNonContainerDatabasesSortByDisplayname,
	"TIMECREATED": ListExternalNonContainerDatabasesSortByTimecreated,
}

var mappingListExternalNonContainerDatabasesSortByEnumLowerCase = map[string]ListExternalNonContainerDatabasesSortByEnum{
	"displayname": ListExternalNonContainerDatabasesSortByDisplayname,
	"timecreated": ListExternalNonContainerDatabasesSortByTimecreated,
}

// GetListExternalNonContainerDatabasesSortByEnumValues Enumerates the set of values for ListExternalNonContainerDatabasesSortByEnum
func GetListExternalNonContainerDatabasesSortByEnumValues() []ListExternalNonContainerDatabasesSortByEnum {
	values := make([]ListExternalNonContainerDatabasesSortByEnum, 0)
	for _, v := range mappingListExternalNonContainerDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalNonContainerDatabasesSortByEnumStringValues Enumerates the set of values in String for ListExternalNonContainerDatabasesSortByEnum
func GetListExternalNonContainerDatabasesSortByEnumStringValues() []string {
	return []string{
		"DISPLAYNAME",
		"TIMECREATED",
	}
}

// GetMappingListExternalNonContainerDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalNonContainerDatabasesSortByEnum(val string) (ListExternalNonContainerDatabasesSortByEnum, bool) {
	enum, ok := mappingListExternalNonContainerDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalNonContainerDatabasesSortOrderEnum Enum with underlying type: string
type ListExternalNonContainerDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListExternalNonContainerDatabasesSortOrderEnum
const (
	ListExternalNonContainerDatabasesSortOrderAsc  ListExternalNonContainerDatabasesSortOrderEnum = "ASC"
	ListExternalNonContainerDatabasesSortOrderDesc ListExternalNonContainerDatabasesSortOrderEnum = "DESC"
)

var mappingListExternalNonContainerDatabasesSortOrderEnum = map[string]ListExternalNonContainerDatabasesSortOrderEnum{
	"ASC":  ListExternalNonContainerDatabasesSortOrderAsc,
	"DESC": ListExternalNonContainerDatabasesSortOrderDesc,
}

var mappingListExternalNonContainerDatabasesSortOrderEnumLowerCase = map[string]ListExternalNonContainerDatabasesSortOrderEnum{
	"asc":  ListExternalNonContainerDatabasesSortOrderAsc,
	"desc": ListExternalNonContainerDatabasesSortOrderDesc,
}

// GetListExternalNonContainerDatabasesSortOrderEnumValues Enumerates the set of values for ListExternalNonContainerDatabasesSortOrderEnum
func GetListExternalNonContainerDatabasesSortOrderEnumValues() []ListExternalNonContainerDatabasesSortOrderEnum {
	values := make([]ListExternalNonContainerDatabasesSortOrderEnum, 0)
	for _, v := range mappingListExternalNonContainerDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalNonContainerDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListExternalNonContainerDatabasesSortOrderEnum
func GetListExternalNonContainerDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalNonContainerDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalNonContainerDatabasesSortOrderEnum(val string) (ListExternalNonContainerDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListExternalNonContainerDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
