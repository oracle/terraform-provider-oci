// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAssignedSubscriptionLineItemsRequest wrapper for the ListAssignedSubscriptionLineItems operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListAssignedSubscriptionLineItems.go.html to see an example of how to use ListAssignedSubscriptionLineItemsRequest.
type ListAssignedSubscriptionLineItemsRequest struct {

	// OCID of the assigned Oracle Cloud Subscription.
	AssignedSubscriptionId *string `mandatory:"true" contributesTo:"path" name:"assignedSubscriptionId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListAssignedSubscriptionLineItemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided.
	// * The default order for timeCreated is descending.
	// * The default order for displayName is ascending.
	// * If no value is specified, timeCreated is the default.
	SortBy ListAssignedSubscriptionLineItemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssignedSubscriptionLineItemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssignedSubscriptionLineItemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssignedSubscriptionLineItemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssignedSubscriptionLineItemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAssignedSubscriptionLineItemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAssignedSubscriptionLineItemsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAssignedSubscriptionLineItemsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssignedSubscriptionLineItemsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAssignedSubscriptionLineItemsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAssignedSubscriptionLineItemsResponse wrapper for the ListAssignedSubscriptionLineItems operation
type ListAssignedSubscriptionLineItemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssignedSubscriptionLineItemCollection instances
	AssignedSubscriptionLineItemCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAssignedSubscriptionLineItemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssignedSubscriptionLineItemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssignedSubscriptionLineItemsSortOrderEnum Enum with underlying type: string
type ListAssignedSubscriptionLineItemsSortOrderEnum string

// Set of constants representing the allowable values for ListAssignedSubscriptionLineItemsSortOrderEnum
const (
	ListAssignedSubscriptionLineItemsSortOrderAsc  ListAssignedSubscriptionLineItemsSortOrderEnum = "ASC"
	ListAssignedSubscriptionLineItemsSortOrderDesc ListAssignedSubscriptionLineItemsSortOrderEnum = "DESC"
)

var mappingListAssignedSubscriptionLineItemsSortOrderEnum = map[string]ListAssignedSubscriptionLineItemsSortOrderEnum{
	"ASC":  ListAssignedSubscriptionLineItemsSortOrderAsc,
	"DESC": ListAssignedSubscriptionLineItemsSortOrderDesc,
}

var mappingListAssignedSubscriptionLineItemsSortOrderEnumLowerCase = map[string]ListAssignedSubscriptionLineItemsSortOrderEnum{
	"asc":  ListAssignedSubscriptionLineItemsSortOrderAsc,
	"desc": ListAssignedSubscriptionLineItemsSortOrderDesc,
}

// GetListAssignedSubscriptionLineItemsSortOrderEnumValues Enumerates the set of values for ListAssignedSubscriptionLineItemsSortOrderEnum
func GetListAssignedSubscriptionLineItemsSortOrderEnumValues() []ListAssignedSubscriptionLineItemsSortOrderEnum {
	values := make([]ListAssignedSubscriptionLineItemsSortOrderEnum, 0)
	for _, v := range mappingListAssignedSubscriptionLineItemsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssignedSubscriptionLineItemsSortOrderEnumStringValues Enumerates the set of values in String for ListAssignedSubscriptionLineItemsSortOrderEnum
func GetListAssignedSubscriptionLineItemsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAssignedSubscriptionLineItemsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssignedSubscriptionLineItemsSortOrderEnum(val string) (ListAssignedSubscriptionLineItemsSortOrderEnum, bool) {
	enum, ok := mappingListAssignedSubscriptionLineItemsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssignedSubscriptionLineItemsSortByEnum Enum with underlying type: string
type ListAssignedSubscriptionLineItemsSortByEnum string

// Set of constants representing the allowable values for ListAssignedSubscriptionLineItemsSortByEnum
const (
	ListAssignedSubscriptionLineItemsSortByTimecreated ListAssignedSubscriptionLineItemsSortByEnum = "timeCreated"
	ListAssignedSubscriptionLineItemsSortByDisplayname ListAssignedSubscriptionLineItemsSortByEnum = "displayName"
)

var mappingListAssignedSubscriptionLineItemsSortByEnum = map[string]ListAssignedSubscriptionLineItemsSortByEnum{
	"timeCreated": ListAssignedSubscriptionLineItemsSortByTimecreated,
	"displayName": ListAssignedSubscriptionLineItemsSortByDisplayname,
}

var mappingListAssignedSubscriptionLineItemsSortByEnumLowerCase = map[string]ListAssignedSubscriptionLineItemsSortByEnum{
	"timecreated": ListAssignedSubscriptionLineItemsSortByTimecreated,
	"displayname": ListAssignedSubscriptionLineItemsSortByDisplayname,
}

// GetListAssignedSubscriptionLineItemsSortByEnumValues Enumerates the set of values for ListAssignedSubscriptionLineItemsSortByEnum
func GetListAssignedSubscriptionLineItemsSortByEnumValues() []ListAssignedSubscriptionLineItemsSortByEnum {
	values := make([]ListAssignedSubscriptionLineItemsSortByEnum, 0)
	for _, v := range mappingListAssignedSubscriptionLineItemsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssignedSubscriptionLineItemsSortByEnumStringValues Enumerates the set of values in String for ListAssignedSubscriptionLineItemsSortByEnum
func GetListAssignedSubscriptionLineItemsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAssignedSubscriptionLineItemsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssignedSubscriptionLineItemsSortByEnum(val string) (ListAssignedSubscriptionLineItemsSortByEnum, bool) {
	enum, ok := mappingListAssignedSubscriptionLineItemsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
