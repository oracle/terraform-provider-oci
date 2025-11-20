// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMulticloudSubscriptionsRequest wrapper for the ListMulticloudSubscriptions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListMulticloudSubscriptions.go.html to see an example of how to use ListMulticloudSubscriptionsRequest.
type ListMulticloudSubscriptionsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListMulticloudSubscriptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMulticloudSubscriptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMulticloudSubscriptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMulticloudSubscriptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMulticloudSubscriptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMulticloudSubscriptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMulticloudSubscriptionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMulticloudSubscriptionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMulticloudSubscriptionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMulticloudSubscriptionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMulticloudSubscriptionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMulticloudSubscriptionsResponse wrapper for the ListMulticloudSubscriptions operation
type ListMulticloudSubscriptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MulticloudSubscriptionCollection instances
	MulticloudSubscriptionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMulticloudSubscriptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMulticloudSubscriptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMulticloudSubscriptionsSortByEnum Enum with underlying type: string
type ListMulticloudSubscriptionsSortByEnum string

// Set of constants representing the allowable values for ListMulticloudSubscriptionsSortByEnum
const (
	ListMulticloudSubscriptionsSortByTimecreated ListMulticloudSubscriptionsSortByEnum = "timeCreated"
	ListMulticloudSubscriptionsSortByDisplayname ListMulticloudSubscriptionsSortByEnum = "displayName"
)

var mappingListMulticloudSubscriptionsSortByEnum = map[string]ListMulticloudSubscriptionsSortByEnum{
	"timeCreated": ListMulticloudSubscriptionsSortByTimecreated,
	"displayName": ListMulticloudSubscriptionsSortByDisplayname,
}

var mappingListMulticloudSubscriptionsSortByEnumLowerCase = map[string]ListMulticloudSubscriptionsSortByEnum{
	"timecreated": ListMulticloudSubscriptionsSortByTimecreated,
	"displayname": ListMulticloudSubscriptionsSortByDisplayname,
}

// GetListMulticloudSubscriptionsSortByEnumValues Enumerates the set of values for ListMulticloudSubscriptionsSortByEnum
func GetListMulticloudSubscriptionsSortByEnumValues() []ListMulticloudSubscriptionsSortByEnum {
	values := make([]ListMulticloudSubscriptionsSortByEnum, 0)
	for _, v := range mappingListMulticloudSubscriptionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMulticloudSubscriptionsSortByEnumStringValues Enumerates the set of values in String for ListMulticloudSubscriptionsSortByEnum
func GetListMulticloudSubscriptionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMulticloudSubscriptionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMulticloudSubscriptionsSortByEnum(val string) (ListMulticloudSubscriptionsSortByEnum, bool) {
	enum, ok := mappingListMulticloudSubscriptionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMulticloudSubscriptionsSortOrderEnum Enum with underlying type: string
type ListMulticloudSubscriptionsSortOrderEnum string

// Set of constants representing the allowable values for ListMulticloudSubscriptionsSortOrderEnum
const (
	ListMulticloudSubscriptionsSortOrderAsc  ListMulticloudSubscriptionsSortOrderEnum = "ASC"
	ListMulticloudSubscriptionsSortOrderDesc ListMulticloudSubscriptionsSortOrderEnum = "DESC"
)

var mappingListMulticloudSubscriptionsSortOrderEnum = map[string]ListMulticloudSubscriptionsSortOrderEnum{
	"ASC":  ListMulticloudSubscriptionsSortOrderAsc,
	"DESC": ListMulticloudSubscriptionsSortOrderDesc,
}

var mappingListMulticloudSubscriptionsSortOrderEnumLowerCase = map[string]ListMulticloudSubscriptionsSortOrderEnum{
	"asc":  ListMulticloudSubscriptionsSortOrderAsc,
	"desc": ListMulticloudSubscriptionsSortOrderDesc,
}

// GetListMulticloudSubscriptionsSortOrderEnumValues Enumerates the set of values for ListMulticloudSubscriptionsSortOrderEnum
func GetListMulticloudSubscriptionsSortOrderEnumValues() []ListMulticloudSubscriptionsSortOrderEnum {
	values := make([]ListMulticloudSubscriptionsSortOrderEnum, 0)
	for _, v := range mappingListMulticloudSubscriptionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMulticloudSubscriptionsSortOrderEnumStringValues Enumerates the set of values in String for ListMulticloudSubscriptionsSortOrderEnum
func GetListMulticloudSubscriptionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMulticloudSubscriptionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMulticloudSubscriptionsSortOrderEnum(val string) (ListMulticloudSubscriptionsSortOrderEnum, bool) {
	enum, ok := mappingListMulticloudSubscriptionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
