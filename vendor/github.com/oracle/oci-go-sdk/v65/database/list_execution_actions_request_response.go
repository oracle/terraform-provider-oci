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

// ListExecutionActionsRequest wrapper for the ListExecutionActions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExecutionActions.go.html to see an example of how to use ListExecutionActionsRequest.
type ListExecutionActionsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListExecutionActionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListExecutionActionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState ExecutionActionSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given execution wondow id.
	ExecutionWindowId *string `mandatory:"false" contributesTo:"query" name:"executionWindowId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExecutionActionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExecutionActionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExecutionActionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExecutionActionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExecutionActionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExecutionActionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExecutionActionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExecutionActionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExecutionActionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExecutionActionSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetExecutionActionSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExecutionActionsResponse wrapper for the ListExecutionActions operation
type ListExecutionActionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExecutionActionSummary instances
	Items []ExecutionActionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExecutionActionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExecutionActionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExecutionActionsSortByEnum Enum with underlying type: string
type ListExecutionActionsSortByEnum string

// Set of constants representing the allowable values for ListExecutionActionsSortByEnum
const (
	ListExecutionActionsSortByTimecreated ListExecutionActionsSortByEnum = "TIMECREATED"
	ListExecutionActionsSortByDisplayname ListExecutionActionsSortByEnum = "DISPLAYNAME"
)

var mappingListExecutionActionsSortByEnum = map[string]ListExecutionActionsSortByEnum{
	"TIMECREATED": ListExecutionActionsSortByTimecreated,
	"DISPLAYNAME": ListExecutionActionsSortByDisplayname,
}

var mappingListExecutionActionsSortByEnumLowerCase = map[string]ListExecutionActionsSortByEnum{
	"timecreated": ListExecutionActionsSortByTimecreated,
	"displayname": ListExecutionActionsSortByDisplayname,
}

// GetListExecutionActionsSortByEnumValues Enumerates the set of values for ListExecutionActionsSortByEnum
func GetListExecutionActionsSortByEnumValues() []ListExecutionActionsSortByEnum {
	values := make([]ListExecutionActionsSortByEnum, 0)
	for _, v := range mappingListExecutionActionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExecutionActionsSortByEnumStringValues Enumerates the set of values in String for ListExecutionActionsSortByEnum
func GetListExecutionActionsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListExecutionActionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExecutionActionsSortByEnum(val string) (ListExecutionActionsSortByEnum, bool) {
	enum, ok := mappingListExecutionActionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExecutionActionsSortOrderEnum Enum with underlying type: string
type ListExecutionActionsSortOrderEnum string

// Set of constants representing the allowable values for ListExecutionActionsSortOrderEnum
const (
	ListExecutionActionsSortOrderAsc  ListExecutionActionsSortOrderEnum = "ASC"
	ListExecutionActionsSortOrderDesc ListExecutionActionsSortOrderEnum = "DESC"
)

var mappingListExecutionActionsSortOrderEnum = map[string]ListExecutionActionsSortOrderEnum{
	"ASC":  ListExecutionActionsSortOrderAsc,
	"DESC": ListExecutionActionsSortOrderDesc,
}

var mappingListExecutionActionsSortOrderEnumLowerCase = map[string]ListExecutionActionsSortOrderEnum{
	"asc":  ListExecutionActionsSortOrderAsc,
	"desc": ListExecutionActionsSortOrderDesc,
}

// GetListExecutionActionsSortOrderEnumValues Enumerates the set of values for ListExecutionActionsSortOrderEnum
func GetListExecutionActionsSortOrderEnumValues() []ListExecutionActionsSortOrderEnum {
	values := make([]ListExecutionActionsSortOrderEnum, 0)
	for _, v := range mappingListExecutionActionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExecutionActionsSortOrderEnumStringValues Enumerates the set of values in String for ListExecutionActionsSortOrderEnum
func GetListExecutionActionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExecutionActionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExecutionActionsSortOrderEnum(val string) (ListExecutionActionsSortOrderEnum, bool) {
	enum, ok := mappingListExecutionActionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
