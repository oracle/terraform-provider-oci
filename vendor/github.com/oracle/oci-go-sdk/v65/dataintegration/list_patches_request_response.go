// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPatchesRequest wrapper for the ListPatches operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListPatches.go.html to see an example of how to use ListPatchesRequest.
type ListPatchesRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The application key.
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Used to filter by the identifier of the published object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListPatchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListPatchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPatchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPatchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPatchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPatchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPatchesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPatchesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPatchesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatchesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPatchesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPatchesResponse wrapper for the ListPatches operation
type ListPatchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PatchSummaryCollection instances
	PatchSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListPatchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPatchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPatchesSortOrderEnum Enum with underlying type: string
type ListPatchesSortOrderEnum string

// Set of constants representing the allowable values for ListPatchesSortOrderEnum
const (
	ListPatchesSortOrderAsc  ListPatchesSortOrderEnum = "ASC"
	ListPatchesSortOrderDesc ListPatchesSortOrderEnum = "DESC"
)

var mappingListPatchesSortOrderEnum = map[string]ListPatchesSortOrderEnum{
	"ASC":  ListPatchesSortOrderAsc,
	"DESC": ListPatchesSortOrderDesc,
}

var mappingListPatchesSortOrderEnumLowerCase = map[string]ListPatchesSortOrderEnum{
	"asc":  ListPatchesSortOrderAsc,
	"desc": ListPatchesSortOrderDesc,
}

// GetListPatchesSortOrderEnumValues Enumerates the set of values for ListPatchesSortOrderEnum
func GetListPatchesSortOrderEnumValues() []ListPatchesSortOrderEnum {
	values := make([]ListPatchesSortOrderEnum, 0)
	for _, v := range mappingListPatchesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchesSortOrderEnumStringValues Enumerates the set of values in String for ListPatchesSortOrderEnum
func GetListPatchesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPatchesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchesSortOrderEnum(val string) (ListPatchesSortOrderEnum, bool) {
	enum, ok := mappingListPatchesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPatchesSortByEnum Enum with underlying type: string
type ListPatchesSortByEnum string

// Set of constants representing the allowable values for ListPatchesSortByEnum
const (
	ListPatchesSortByTimeCreated ListPatchesSortByEnum = "TIME_CREATED"
	ListPatchesSortByDisplayName ListPatchesSortByEnum = "DISPLAY_NAME"
	ListPatchesSortByTimeUpdated ListPatchesSortByEnum = "TIME_UPDATED"
)

var mappingListPatchesSortByEnum = map[string]ListPatchesSortByEnum{
	"TIME_CREATED": ListPatchesSortByTimeCreated,
	"DISPLAY_NAME": ListPatchesSortByDisplayName,
	"TIME_UPDATED": ListPatchesSortByTimeUpdated,
}

var mappingListPatchesSortByEnumLowerCase = map[string]ListPatchesSortByEnum{
	"time_created": ListPatchesSortByTimeCreated,
	"display_name": ListPatchesSortByDisplayName,
	"time_updated": ListPatchesSortByTimeUpdated,
}

// GetListPatchesSortByEnumValues Enumerates the set of values for ListPatchesSortByEnum
func GetListPatchesSortByEnumValues() []ListPatchesSortByEnum {
	values := make([]ListPatchesSortByEnum, 0)
	for _, v := range mappingListPatchesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchesSortByEnumStringValues Enumerates the set of values in String for ListPatchesSortByEnum
func GetListPatchesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListPatchesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchesSortByEnum(val string) (ListPatchesSortByEnum, bool) {
	enum, ok := mappingListPatchesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
