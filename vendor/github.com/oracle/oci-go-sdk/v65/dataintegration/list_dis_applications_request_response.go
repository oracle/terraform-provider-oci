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

// ListDisApplicationsRequest wrapper for the ListDisApplications operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListDisApplications.go.html to see an example of how to use ListDisApplicationsRequest.
type ListDisApplicationsRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// OCID of the compartment for which the list of DIS Applications is to be retrieved.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// This parameter can be used to filter objects by the names that match partially or fully with the given value.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// Used to filter by the identifier of the published object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListDisApplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListDisApplicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDisApplicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDisApplicationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDisApplicationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDisApplicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDisApplicationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDisApplicationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDisApplicationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDisApplicationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDisApplicationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDisApplicationsResponse wrapper for the ListDisApplications operation
type ListDisApplicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DisApplicationSummaryCollection instances
	DisApplicationSummaryCollection `presentIn:"body"`

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

func (response ListDisApplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDisApplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDisApplicationsSortOrderEnum Enum with underlying type: string
type ListDisApplicationsSortOrderEnum string

// Set of constants representing the allowable values for ListDisApplicationsSortOrderEnum
const (
	ListDisApplicationsSortOrderAsc  ListDisApplicationsSortOrderEnum = "ASC"
	ListDisApplicationsSortOrderDesc ListDisApplicationsSortOrderEnum = "DESC"
)

var mappingListDisApplicationsSortOrderEnum = map[string]ListDisApplicationsSortOrderEnum{
	"ASC":  ListDisApplicationsSortOrderAsc,
	"DESC": ListDisApplicationsSortOrderDesc,
}

var mappingListDisApplicationsSortOrderEnumLowerCase = map[string]ListDisApplicationsSortOrderEnum{
	"asc":  ListDisApplicationsSortOrderAsc,
	"desc": ListDisApplicationsSortOrderDesc,
}

// GetListDisApplicationsSortOrderEnumValues Enumerates the set of values for ListDisApplicationsSortOrderEnum
func GetListDisApplicationsSortOrderEnumValues() []ListDisApplicationsSortOrderEnum {
	values := make([]ListDisApplicationsSortOrderEnum, 0)
	for _, v := range mappingListDisApplicationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDisApplicationsSortOrderEnumStringValues Enumerates the set of values in String for ListDisApplicationsSortOrderEnum
func GetListDisApplicationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDisApplicationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDisApplicationsSortOrderEnum(val string) (ListDisApplicationsSortOrderEnum, bool) {
	enum, ok := mappingListDisApplicationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDisApplicationsSortByEnum Enum with underlying type: string
type ListDisApplicationsSortByEnum string

// Set of constants representing the allowable values for ListDisApplicationsSortByEnum
const (
	ListDisApplicationsSortByTimeCreated ListDisApplicationsSortByEnum = "TIME_CREATED"
	ListDisApplicationsSortByDisplayName ListDisApplicationsSortByEnum = "DISPLAY_NAME"
	ListDisApplicationsSortByTimeUpdated ListDisApplicationsSortByEnum = "TIME_UPDATED"
)

var mappingListDisApplicationsSortByEnum = map[string]ListDisApplicationsSortByEnum{
	"TIME_CREATED": ListDisApplicationsSortByTimeCreated,
	"DISPLAY_NAME": ListDisApplicationsSortByDisplayName,
	"TIME_UPDATED": ListDisApplicationsSortByTimeUpdated,
}

var mappingListDisApplicationsSortByEnumLowerCase = map[string]ListDisApplicationsSortByEnum{
	"time_created": ListDisApplicationsSortByTimeCreated,
	"display_name": ListDisApplicationsSortByDisplayName,
	"time_updated": ListDisApplicationsSortByTimeUpdated,
}

// GetListDisApplicationsSortByEnumValues Enumerates the set of values for ListDisApplicationsSortByEnum
func GetListDisApplicationsSortByEnumValues() []ListDisApplicationsSortByEnum {
	values := make([]ListDisApplicationsSortByEnum, 0)
	for _, v := range mappingListDisApplicationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDisApplicationsSortByEnumStringValues Enumerates the set of values in String for ListDisApplicationsSortByEnum
func GetListDisApplicationsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListDisApplicationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDisApplicationsSortByEnum(val string) (ListDisApplicationsSortByEnum, bool) {
	enum, ok := mappingListDisApplicationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
