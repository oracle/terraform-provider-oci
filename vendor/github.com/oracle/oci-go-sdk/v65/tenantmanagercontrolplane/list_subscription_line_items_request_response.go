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

// ListSubscriptionLineItemsRequest wrapper for the ListSubscriptionLineItems operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListSubscriptionLineItems.go.html to see an example of how to use ListSubscriptionLineItemsRequest.
type ListSubscriptionLineItemsRequest struct {

	// OCID of the subscription.
	SubscriptionId *string `mandatory:"true" contributesTo:"path" name:"subscriptionId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListSubscriptionLineItemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided.
	// * The default order for timeCreated is descending.
	// * The default order for displayName is ascending.
	// * If no value is specified, timeCreated is the default.
	SortBy ListSubscriptionLineItemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubscriptionLineItemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubscriptionLineItemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubscriptionLineItemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubscriptionLineItemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubscriptionLineItemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubscriptionLineItemsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubscriptionLineItemsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubscriptionLineItemsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubscriptionLineItemsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubscriptionLineItemsResponse wrapper for the ListSubscriptionLineItems operation
type ListSubscriptionLineItemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubscriptionLineItemCollection instances
	SubscriptionLineItemCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSubscriptionLineItemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubscriptionLineItemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubscriptionLineItemsSortOrderEnum Enum with underlying type: string
type ListSubscriptionLineItemsSortOrderEnum string

// Set of constants representing the allowable values for ListSubscriptionLineItemsSortOrderEnum
const (
	ListSubscriptionLineItemsSortOrderAsc  ListSubscriptionLineItemsSortOrderEnum = "ASC"
	ListSubscriptionLineItemsSortOrderDesc ListSubscriptionLineItemsSortOrderEnum = "DESC"
)

var mappingListSubscriptionLineItemsSortOrderEnum = map[string]ListSubscriptionLineItemsSortOrderEnum{
	"ASC":  ListSubscriptionLineItemsSortOrderAsc,
	"DESC": ListSubscriptionLineItemsSortOrderDesc,
}

var mappingListSubscriptionLineItemsSortOrderEnumLowerCase = map[string]ListSubscriptionLineItemsSortOrderEnum{
	"asc":  ListSubscriptionLineItemsSortOrderAsc,
	"desc": ListSubscriptionLineItemsSortOrderDesc,
}

// GetListSubscriptionLineItemsSortOrderEnumValues Enumerates the set of values for ListSubscriptionLineItemsSortOrderEnum
func GetListSubscriptionLineItemsSortOrderEnumValues() []ListSubscriptionLineItemsSortOrderEnum {
	values := make([]ListSubscriptionLineItemsSortOrderEnum, 0)
	for _, v := range mappingListSubscriptionLineItemsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubscriptionLineItemsSortOrderEnumStringValues Enumerates the set of values in String for ListSubscriptionLineItemsSortOrderEnum
func GetListSubscriptionLineItemsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubscriptionLineItemsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubscriptionLineItemsSortOrderEnum(val string) (ListSubscriptionLineItemsSortOrderEnum, bool) {
	enum, ok := mappingListSubscriptionLineItemsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubscriptionLineItemsSortByEnum Enum with underlying type: string
type ListSubscriptionLineItemsSortByEnum string

// Set of constants representing the allowable values for ListSubscriptionLineItemsSortByEnum
const (
	ListSubscriptionLineItemsSortByTimecreated ListSubscriptionLineItemsSortByEnum = "timeCreated"
	ListSubscriptionLineItemsSortByDisplayname ListSubscriptionLineItemsSortByEnum = "displayName"
)

var mappingListSubscriptionLineItemsSortByEnum = map[string]ListSubscriptionLineItemsSortByEnum{
	"timeCreated": ListSubscriptionLineItemsSortByTimecreated,
	"displayName": ListSubscriptionLineItemsSortByDisplayname,
}

var mappingListSubscriptionLineItemsSortByEnumLowerCase = map[string]ListSubscriptionLineItemsSortByEnum{
	"timecreated": ListSubscriptionLineItemsSortByTimecreated,
	"displayname": ListSubscriptionLineItemsSortByDisplayname,
}

// GetListSubscriptionLineItemsSortByEnumValues Enumerates the set of values for ListSubscriptionLineItemsSortByEnum
func GetListSubscriptionLineItemsSortByEnumValues() []ListSubscriptionLineItemsSortByEnum {
	values := make([]ListSubscriptionLineItemsSortByEnum, 0)
	for _, v := range mappingListSubscriptionLineItemsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubscriptionLineItemsSortByEnumStringValues Enumerates the set of values in String for ListSubscriptionLineItemsSortByEnum
func GetListSubscriptionLineItemsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSubscriptionLineItemsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubscriptionLineItemsSortByEnum(val string) (ListSubscriptionLineItemsSortByEnum, bool) {
	enum, ok := mappingListSubscriptionLineItemsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
