// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOneoffPatchesRequest wrapper for the ListOneoffPatches operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListOneoffPatches.go.html to see an example of how to use ListOneoffPatchesRequest.
type ListOneoffPatchesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListOneoffPatchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListOneoffPatchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly
	LifecycleState OneoffPatchSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOneoffPatchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOneoffPatchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOneoffPatchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOneoffPatchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOneoffPatchesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOneoffPatchesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOneoffPatchesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOneoffPatchesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOneoffPatchesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOneoffPatchSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOneoffPatchSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOneoffPatchesResponse wrapper for the ListOneoffPatches operation
type ListOneoffPatchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []OneoffPatchSummary instances
	Items []OneoffPatchSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOneoffPatchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOneoffPatchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOneoffPatchesSortByEnum Enum with underlying type: string
type ListOneoffPatchesSortByEnum string

// Set of constants representing the allowable values for ListOneoffPatchesSortByEnum
const (
	ListOneoffPatchesSortByTimecreated ListOneoffPatchesSortByEnum = "TIMECREATED"
	ListOneoffPatchesSortByDisplayname ListOneoffPatchesSortByEnum = "DISPLAYNAME"
)

var mappingListOneoffPatchesSortByEnum = map[string]ListOneoffPatchesSortByEnum{
	"TIMECREATED": ListOneoffPatchesSortByTimecreated,
	"DISPLAYNAME": ListOneoffPatchesSortByDisplayname,
}

var mappingListOneoffPatchesSortByEnumLowerCase = map[string]ListOneoffPatchesSortByEnum{
	"timecreated": ListOneoffPatchesSortByTimecreated,
	"displayname": ListOneoffPatchesSortByDisplayname,
}

// GetListOneoffPatchesSortByEnumValues Enumerates the set of values for ListOneoffPatchesSortByEnum
func GetListOneoffPatchesSortByEnumValues() []ListOneoffPatchesSortByEnum {
	values := make([]ListOneoffPatchesSortByEnum, 0)
	for _, v := range mappingListOneoffPatchesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOneoffPatchesSortByEnumStringValues Enumerates the set of values in String for ListOneoffPatchesSortByEnum
func GetListOneoffPatchesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListOneoffPatchesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOneoffPatchesSortByEnum(val string) (ListOneoffPatchesSortByEnum, bool) {
	enum, ok := mappingListOneoffPatchesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOneoffPatchesSortOrderEnum Enum with underlying type: string
type ListOneoffPatchesSortOrderEnum string

// Set of constants representing the allowable values for ListOneoffPatchesSortOrderEnum
const (
	ListOneoffPatchesSortOrderAsc  ListOneoffPatchesSortOrderEnum = "ASC"
	ListOneoffPatchesSortOrderDesc ListOneoffPatchesSortOrderEnum = "DESC"
)

var mappingListOneoffPatchesSortOrderEnum = map[string]ListOneoffPatchesSortOrderEnum{
	"ASC":  ListOneoffPatchesSortOrderAsc,
	"DESC": ListOneoffPatchesSortOrderDesc,
}

var mappingListOneoffPatchesSortOrderEnumLowerCase = map[string]ListOneoffPatchesSortOrderEnum{
	"asc":  ListOneoffPatchesSortOrderAsc,
	"desc": ListOneoffPatchesSortOrderDesc,
}

// GetListOneoffPatchesSortOrderEnumValues Enumerates the set of values for ListOneoffPatchesSortOrderEnum
func GetListOneoffPatchesSortOrderEnumValues() []ListOneoffPatchesSortOrderEnum {
	values := make([]ListOneoffPatchesSortOrderEnum, 0)
	for _, v := range mappingListOneoffPatchesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOneoffPatchesSortOrderEnumStringValues Enumerates the set of values in String for ListOneoffPatchesSortOrderEnum
func GetListOneoffPatchesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOneoffPatchesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOneoffPatchesSortOrderEnum(val string) (ListOneoffPatchesSortOrderEnum, bool) {
	enum, ok := mappingListOneoffPatchesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
