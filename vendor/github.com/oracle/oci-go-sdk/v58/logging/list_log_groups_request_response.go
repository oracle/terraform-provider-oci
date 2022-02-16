// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListLogGroupsRequest wrapper for the ListLogGroups operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/logging/ListLogGroups.go.html to see an example of how to use ListLogGroupsRequest.
type ListLogGroupsRequest struct {

	// Compartment OCID to list resources in. See compartmentIdInSubtree
	//      for nested compartments traversal.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Specifies whether or not nested compartments should be traversed. Defaults to false.
	IsCompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"isCompartmentIdInSubtree"`

	// Resource name
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The value of the `opc-next-page` or `opc-previous-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by (one column only). Default sort order is
	// ascending exception of `timeCreated` and `timeLastModified` columns (descending).
	SortBy ListLogGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListLogGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLogGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLogGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLogGroupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLogGroupsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLogGroupsResponse wrapper for the ListLogGroups operation
type ListLogGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []LogGroupSummary instances
	Items []LogGroupSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages
	// of results exist. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLogGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogGroupsSortByEnum Enum with underlying type: string
type ListLogGroupsSortByEnum string

// Set of constants representing the allowable values for ListLogGroupsSortByEnum
const (
	ListLogGroupsSortByTimecreated ListLogGroupsSortByEnum = "timeCreated"
	ListLogGroupsSortByDisplayname ListLogGroupsSortByEnum = "displayName"
)

var mappingListLogGroupsSortByEnum = map[string]ListLogGroupsSortByEnum{
	"timeCreated": ListLogGroupsSortByTimecreated,
	"displayName": ListLogGroupsSortByDisplayname,
}

// GetListLogGroupsSortByEnumValues Enumerates the set of values for ListLogGroupsSortByEnum
func GetListLogGroupsSortByEnumValues() []ListLogGroupsSortByEnum {
	values := make([]ListLogGroupsSortByEnum, 0)
	for _, v := range mappingListLogGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogGroupsSortByEnumStringValues Enumerates the set of values in String for ListLogGroupsSortByEnum
func GetListLogGroupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListLogGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogGroupsSortByEnum(val string) (ListLogGroupsSortByEnum, bool) {
	mappingListLogGroupsSortByEnumIgnoreCase := make(map[string]ListLogGroupsSortByEnum)
	for k, v := range mappingListLogGroupsSortByEnum {
		mappingListLogGroupsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListLogGroupsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogGroupsSortOrderEnum Enum with underlying type: string
type ListLogGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListLogGroupsSortOrderEnum
const (
	ListLogGroupsSortOrderAsc  ListLogGroupsSortOrderEnum = "ASC"
	ListLogGroupsSortOrderDesc ListLogGroupsSortOrderEnum = "DESC"
)

var mappingListLogGroupsSortOrderEnum = map[string]ListLogGroupsSortOrderEnum{
	"ASC":  ListLogGroupsSortOrderAsc,
	"DESC": ListLogGroupsSortOrderDesc,
}

// GetListLogGroupsSortOrderEnumValues Enumerates the set of values for ListLogGroupsSortOrderEnum
func GetListLogGroupsSortOrderEnumValues() []ListLogGroupsSortOrderEnum {
	values := make([]ListLogGroupsSortOrderEnum, 0)
	for _, v := range mappingListLogGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListLogGroupsSortOrderEnum
func GetListLogGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLogGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogGroupsSortOrderEnum(val string) (ListLogGroupsSortOrderEnum, bool) {
	mappingListLogGroupsSortOrderEnumIgnoreCase := make(map[string]ListLogGroupsSortOrderEnum)
	for k, v := range mappingListLogGroupsSortOrderEnum {
		mappingListLogGroupsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListLogGroupsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
