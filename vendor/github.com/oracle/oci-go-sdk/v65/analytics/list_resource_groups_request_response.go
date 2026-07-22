// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package analytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListResourceGroupsRequest wrapper for the ListResourceGroups operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/ListResourceGroups.go.html to see an example of how to use ListResourceGroupsRequest.
type ListResourceGroupsRequest struct {

	// The OCID of the Analytics instance.
	AnalyticsInstanceId *string `mandatory:"true" contributesTo:"path" name:"analyticsInstanceId"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by (one column only). Default sort order is
	// default group first, then sort by resource name, ascending.
	SortBy ListResourceGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListResourceGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourceGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourceGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourceGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourceGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResourceGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResourceGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResourceGroupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResourceGroupsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResourceGroupsResponse wrapper for the ListResourceGroups operation
type ListResourceGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []InstanceResourceGroupSummary instances
	Items []InstanceResourceGroupSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResourceGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourceGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourceGroupsSortByEnum Enum with underlying type: string
type ListResourceGroupsSortByEnum string

// Set of constants representing the allowable values for ListResourceGroupsSortByEnum
const (
	ListResourceGroupsSortById           ListResourceGroupsSortByEnum = "id"
	ListResourceGroupsSortByResourcename ListResourceGroupsSortByEnum = "resourceName"
	ListResourceGroupsSortByDisplayname  ListResourceGroupsSortByEnum = "displayName"
	ListResourceGroupsSortByDescription  ListResourceGroupsSortByEnum = "description"
	ListResourceGroupsSortByCapacity     ListResourceGroupsSortByEnum = "capacity"
	ListResourceGroupsSortByStatus       ListResourceGroupsSortByEnum = "status"
)

var mappingListResourceGroupsSortByEnum = map[string]ListResourceGroupsSortByEnum{
	"id":           ListResourceGroupsSortById,
	"resourceName": ListResourceGroupsSortByResourcename,
	"displayName":  ListResourceGroupsSortByDisplayname,
	"description":  ListResourceGroupsSortByDescription,
	"capacity":     ListResourceGroupsSortByCapacity,
	"status":       ListResourceGroupsSortByStatus,
}

var mappingListResourceGroupsSortByEnumLowerCase = map[string]ListResourceGroupsSortByEnum{
	"id":           ListResourceGroupsSortById,
	"resourcename": ListResourceGroupsSortByResourcename,
	"displayname":  ListResourceGroupsSortByDisplayname,
	"description":  ListResourceGroupsSortByDescription,
	"capacity":     ListResourceGroupsSortByCapacity,
	"status":       ListResourceGroupsSortByStatus,
}

// GetListResourceGroupsSortByEnumValues Enumerates the set of values for ListResourceGroupsSortByEnum
func GetListResourceGroupsSortByEnumValues() []ListResourceGroupsSortByEnum {
	values := make([]ListResourceGroupsSortByEnum, 0)
	for _, v := range mappingListResourceGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceGroupsSortByEnumStringValues Enumerates the set of values in String for ListResourceGroupsSortByEnum
func GetListResourceGroupsSortByEnumStringValues() []string {
	return []string{
		"id",
		"resourceName",
		"displayName",
		"description",
		"capacity",
		"status",
	}
}

// GetMappingListResourceGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceGroupsSortByEnum(val string) (ListResourceGroupsSortByEnum, bool) {
	enum, ok := mappingListResourceGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceGroupsSortOrderEnum Enum with underlying type: string
type ListResourceGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListResourceGroupsSortOrderEnum
const (
	ListResourceGroupsSortOrderAsc  ListResourceGroupsSortOrderEnum = "ASC"
	ListResourceGroupsSortOrderDesc ListResourceGroupsSortOrderEnum = "DESC"
)

var mappingListResourceGroupsSortOrderEnum = map[string]ListResourceGroupsSortOrderEnum{
	"ASC":  ListResourceGroupsSortOrderAsc,
	"DESC": ListResourceGroupsSortOrderDesc,
}

var mappingListResourceGroupsSortOrderEnumLowerCase = map[string]ListResourceGroupsSortOrderEnum{
	"asc":  ListResourceGroupsSortOrderAsc,
	"desc": ListResourceGroupsSortOrderDesc,
}

// GetListResourceGroupsSortOrderEnumValues Enumerates the set of values for ListResourceGroupsSortOrderEnum
func GetListResourceGroupsSortOrderEnumValues() []ListResourceGroupsSortOrderEnum {
	values := make([]ListResourceGroupsSortOrderEnum, 0)
	for _, v := range mappingListResourceGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListResourceGroupsSortOrderEnum
func GetListResourceGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResourceGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceGroupsSortOrderEnum(val string) (ListResourceGroupsSortOrderEnum, bool) {
	enum, ok := mappingListResourceGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
