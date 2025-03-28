// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPatchChangesRequest wrapper for the ListPatchChanges operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListPatchChanges.go.html to see an example of how to use ListPatchChangesRequest.
type ListPatchChangesRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The application key.
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Specifies the patch key to query from.
	SincePatch *string `mandatory:"false" contributesTo:"query" name:"sincePatch"`

	// Specifies the patch key to query to.
	ToPatch *string `mandatory:"false" contributesTo:"query" name:"toPatch"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListPatchChangesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListPatchChangesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPatchChangesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPatchChangesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPatchChangesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPatchChangesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPatchChangesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPatchChangesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPatchChangesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatchChangesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPatchChangesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPatchChangesResponse wrapper for the ListPatchChanges operation
type ListPatchChangesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PatchChangeSummaryCollection instances
	PatchChangeSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListPatchChangesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPatchChangesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPatchChangesSortOrderEnum Enum with underlying type: string
type ListPatchChangesSortOrderEnum string

// Set of constants representing the allowable values for ListPatchChangesSortOrderEnum
const (
	ListPatchChangesSortOrderAsc  ListPatchChangesSortOrderEnum = "ASC"
	ListPatchChangesSortOrderDesc ListPatchChangesSortOrderEnum = "DESC"
)

var mappingListPatchChangesSortOrderEnum = map[string]ListPatchChangesSortOrderEnum{
	"ASC":  ListPatchChangesSortOrderAsc,
	"DESC": ListPatchChangesSortOrderDesc,
}

var mappingListPatchChangesSortOrderEnumLowerCase = map[string]ListPatchChangesSortOrderEnum{
	"asc":  ListPatchChangesSortOrderAsc,
	"desc": ListPatchChangesSortOrderDesc,
}

// GetListPatchChangesSortOrderEnumValues Enumerates the set of values for ListPatchChangesSortOrderEnum
func GetListPatchChangesSortOrderEnumValues() []ListPatchChangesSortOrderEnum {
	values := make([]ListPatchChangesSortOrderEnum, 0)
	for _, v := range mappingListPatchChangesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchChangesSortOrderEnumStringValues Enumerates the set of values in String for ListPatchChangesSortOrderEnum
func GetListPatchChangesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPatchChangesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchChangesSortOrderEnum(val string) (ListPatchChangesSortOrderEnum, bool) {
	enum, ok := mappingListPatchChangesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPatchChangesSortByEnum Enum with underlying type: string
type ListPatchChangesSortByEnum string

// Set of constants representing the allowable values for ListPatchChangesSortByEnum
const (
	ListPatchChangesSortByTimeCreated ListPatchChangesSortByEnum = "TIME_CREATED"
	ListPatchChangesSortByDisplayName ListPatchChangesSortByEnum = "DISPLAY_NAME"
	ListPatchChangesSortByTimeUpdated ListPatchChangesSortByEnum = "TIME_UPDATED"
)

var mappingListPatchChangesSortByEnum = map[string]ListPatchChangesSortByEnum{
	"TIME_CREATED": ListPatchChangesSortByTimeCreated,
	"DISPLAY_NAME": ListPatchChangesSortByDisplayName,
	"TIME_UPDATED": ListPatchChangesSortByTimeUpdated,
}

var mappingListPatchChangesSortByEnumLowerCase = map[string]ListPatchChangesSortByEnum{
	"time_created": ListPatchChangesSortByTimeCreated,
	"display_name": ListPatchChangesSortByDisplayName,
	"time_updated": ListPatchChangesSortByTimeUpdated,
}

// GetListPatchChangesSortByEnumValues Enumerates the set of values for ListPatchChangesSortByEnum
func GetListPatchChangesSortByEnumValues() []ListPatchChangesSortByEnum {
	values := make([]ListPatchChangesSortByEnum, 0)
	for _, v := range mappingListPatchChangesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchChangesSortByEnumStringValues Enumerates the set of values in String for ListPatchChangesSortByEnum
func GetListPatchChangesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListPatchChangesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchChangesSortByEnum(val string) (ListPatchChangesSortByEnum, bool) {
	enum, ok := mappingListPatchChangesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
