// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package onesubscription

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOrganizationSubscriptionsRequest wrapper for the ListOrganizationSubscriptions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/onesubscription/ListOrganizationSubscriptions.go.html to see an example of how to use ListOrganizationSubscriptionsRequest.
type ListOrganizationSubscriptionsRequest struct {

	// The OCID of the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return in a paginated "List" call. Default: (`50`)
	// Example: '500'
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the 'opc-next-page' response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending ('ASC') or descending ('DESC').
	SortOrder ListOrganizationSubscriptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order ('sortOrder').
	SortBy ListOrganizationSubscriptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOrganizationSubscriptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOrganizationSubscriptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOrganizationSubscriptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOrganizationSubscriptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOrganizationSubscriptionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOrganizationSubscriptionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOrganizationSubscriptionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOrganizationSubscriptionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOrganizationSubscriptionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOrganizationSubscriptionsResponse wrapper for the ListOrganizationSubscriptions operation
type ListOrganizationSubscriptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []OrganizationSubscriptionSummary instances
	Items []OrganizationSubscriptionSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListOrganizationSubscriptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOrganizationSubscriptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOrganizationSubscriptionsSortOrderEnum Enum with underlying type: string
type ListOrganizationSubscriptionsSortOrderEnum string

// Set of constants representing the allowable values for ListOrganizationSubscriptionsSortOrderEnum
const (
	ListOrganizationSubscriptionsSortOrderAsc  ListOrganizationSubscriptionsSortOrderEnum = "ASC"
	ListOrganizationSubscriptionsSortOrderDesc ListOrganizationSubscriptionsSortOrderEnum = "DESC"
)

var mappingListOrganizationSubscriptionsSortOrderEnum = map[string]ListOrganizationSubscriptionsSortOrderEnum{
	"ASC":  ListOrganizationSubscriptionsSortOrderAsc,
	"DESC": ListOrganizationSubscriptionsSortOrderDesc,
}

var mappingListOrganizationSubscriptionsSortOrderEnumLowerCase = map[string]ListOrganizationSubscriptionsSortOrderEnum{
	"asc":  ListOrganizationSubscriptionsSortOrderAsc,
	"desc": ListOrganizationSubscriptionsSortOrderDesc,
}

// GetListOrganizationSubscriptionsSortOrderEnumValues Enumerates the set of values for ListOrganizationSubscriptionsSortOrderEnum
func GetListOrganizationSubscriptionsSortOrderEnumValues() []ListOrganizationSubscriptionsSortOrderEnum {
	values := make([]ListOrganizationSubscriptionsSortOrderEnum, 0)
	for _, v := range mappingListOrganizationSubscriptionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOrganizationSubscriptionsSortOrderEnumStringValues Enumerates the set of values in String for ListOrganizationSubscriptionsSortOrderEnum
func GetListOrganizationSubscriptionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOrganizationSubscriptionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOrganizationSubscriptionsSortOrderEnum(val string) (ListOrganizationSubscriptionsSortOrderEnum, bool) {
	enum, ok := mappingListOrganizationSubscriptionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOrganizationSubscriptionsSortByEnum Enum with underlying type: string
type ListOrganizationSubscriptionsSortByEnum string

// Set of constants representing the allowable values for ListOrganizationSubscriptionsSortByEnum
const (
	ListOrganizationSubscriptionsSortByOrdernumber   ListOrganizationSubscriptionsSortByEnum = "ORDERNUMBER"
	ListOrganizationSubscriptionsSortByTimeinvoicing ListOrganizationSubscriptionsSortByEnum = "TIMEINVOICING"
)

var mappingListOrganizationSubscriptionsSortByEnum = map[string]ListOrganizationSubscriptionsSortByEnum{
	"ORDERNUMBER":   ListOrganizationSubscriptionsSortByOrdernumber,
	"TIMEINVOICING": ListOrganizationSubscriptionsSortByTimeinvoicing,
}

var mappingListOrganizationSubscriptionsSortByEnumLowerCase = map[string]ListOrganizationSubscriptionsSortByEnum{
	"ordernumber":   ListOrganizationSubscriptionsSortByOrdernumber,
	"timeinvoicing": ListOrganizationSubscriptionsSortByTimeinvoicing,
}

// GetListOrganizationSubscriptionsSortByEnumValues Enumerates the set of values for ListOrganizationSubscriptionsSortByEnum
func GetListOrganizationSubscriptionsSortByEnumValues() []ListOrganizationSubscriptionsSortByEnum {
	values := make([]ListOrganizationSubscriptionsSortByEnum, 0)
	for _, v := range mappingListOrganizationSubscriptionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOrganizationSubscriptionsSortByEnumStringValues Enumerates the set of values in String for ListOrganizationSubscriptionsSortByEnum
func GetListOrganizationSubscriptionsSortByEnumStringValues() []string {
	return []string{
		"ORDERNUMBER",
		"TIMEINVOICING",
	}
}

// GetMappingListOrganizationSubscriptionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOrganizationSubscriptionsSortByEnum(val string) (ListOrganizationSubscriptionsSortByEnum, bool) {
	enum, ok := mappingListOrganizationSubscriptionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
