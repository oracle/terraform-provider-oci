// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package usage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRedeemableUsersRequest wrapper for the ListRedeemableUsers operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usage/ListRedeemableUsers.go.html to see an example of how to use ListRedeemableUsersRequest.
type ListRedeemableUsersRequest struct {

	// The OCID of the tenancy.
	TenancyId *string `mandatory:"true" contributesTo:"query" name:"tenancyId"`

	// The subscription ID for which rewards information is requested for.
	SubscriptionId *string `mandatory:"true" contributesTo:"path" name:"subscriptionId"`

	// Unique, Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The value of the 'opc-next-page' response header from the previous call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, which can be ascending (ASC) or descending (DESC).
	SortOrder ListRedeemableUsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Supports one sort order.
	SortBy ListRedeemableUsersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRedeemableUsersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRedeemableUsersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRedeemableUsersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRedeemableUsersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRedeemableUsersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRedeemableUsersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRedeemableUsersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRedeemableUsersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRedeemableUsersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRedeemableUsersResponse wrapper for the ListRedeemableUsers operation
type ListRedeemableUsersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RedeemableUserCollection instances
	RedeemableUserCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRedeemableUsersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRedeemableUsersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRedeemableUsersSortOrderEnum Enum with underlying type: string
type ListRedeemableUsersSortOrderEnum string

// Set of constants representing the allowable values for ListRedeemableUsersSortOrderEnum
const (
	ListRedeemableUsersSortOrderAsc  ListRedeemableUsersSortOrderEnum = "ASC"
	ListRedeemableUsersSortOrderDesc ListRedeemableUsersSortOrderEnum = "DESC"
)

var mappingListRedeemableUsersSortOrderEnum = map[string]ListRedeemableUsersSortOrderEnum{
	"ASC":  ListRedeemableUsersSortOrderAsc,
	"DESC": ListRedeemableUsersSortOrderDesc,
}

var mappingListRedeemableUsersSortOrderEnumLowerCase = map[string]ListRedeemableUsersSortOrderEnum{
	"asc":  ListRedeemableUsersSortOrderAsc,
	"desc": ListRedeemableUsersSortOrderDesc,
}

// GetListRedeemableUsersSortOrderEnumValues Enumerates the set of values for ListRedeemableUsersSortOrderEnum
func GetListRedeemableUsersSortOrderEnumValues() []ListRedeemableUsersSortOrderEnum {
	values := make([]ListRedeemableUsersSortOrderEnum, 0)
	for _, v := range mappingListRedeemableUsersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRedeemableUsersSortOrderEnumStringValues Enumerates the set of values in String for ListRedeemableUsersSortOrderEnum
func GetListRedeemableUsersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRedeemableUsersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRedeemableUsersSortOrderEnum(val string) (ListRedeemableUsersSortOrderEnum, bool) {
	enum, ok := mappingListRedeemableUsersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRedeemableUsersSortByEnum Enum with underlying type: string
type ListRedeemableUsersSortByEnum string

// Set of constants representing the allowable values for ListRedeemableUsersSortByEnum
const (
	ListRedeemableUsersSortByTimecreated ListRedeemableUsersSortByEnum = "TIMECREATED"
	ListRedeemableUsersSortByTimestart   ListRedeemableUsersSortByEnum = "TIMESTART"
)

var mappingListRedeemableUsersSortByEnum = map[string]ListRedeemableUsersSortByEnum{
	"TIMECREATED": ListRedeemableUsersSortByTimecreated,
	"TIMESTART":   ListRedeemableUsersSortByTimestart,
}

var mappingListRedeemableUsersSortByEnumLowerCase = map[string]ListRedeemableUsersSortByEnum{
	"timecreated": ListRedeemableUsersSortByTimecreated,
	"timestart":   ListRedeemableUsersSortByTimestart,
}

// GetListRedeemableUsersSortByEnumValues Enumerates the set of values for ListRedeemableUsersSortByEnum
func GetListRedeemableUsersSortByEnumValues() []ListRedeemableUsersSortByEnum {
	values := make([]ListRedeemableUsersSortByEnum, 0)
	for _, v := range mappingListRedeemableUsersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRedeemableUsersSortByEnumStringValues Enumerates the set of values in String for ListRedeemableUsersSortByEnum
func GetListRedeemableUsersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"TIMESTART",
	}
}

// GetMappingListRedeemableUsersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRedeemableUsersSortByEnum(val string) (ListRedeemableUsersSortByEnum, bool) {
	enum, ok := mappingListRedeemableUsersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
