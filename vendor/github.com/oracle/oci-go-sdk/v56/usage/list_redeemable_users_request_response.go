// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package usage

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListRedeemableUsersRequest wrapper for the ListRedeemableUsers operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usage/ListRedeemableUsers.go.html to see an example of how to use ListRedeemableUsersRequest.
type ListRedeemableUsersRequest struct {

	// The OCID of the tenancy.
	TenancyId *string `mandatory:"true" contributesTo:"query" name:"tenancyId"`

	// The subscriptionId for which rewards information is requested for.
	SubscriptionId *string `mandatory:"true" contributesTo:"path" name:"subscriptionId"`

	// Unique, Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The value of the 'opc-next-page' response header from the previous call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, can be ascending (ASC) or descending (DESC).
	SortOrder ListRedeemableUsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by, supports one sort Order.
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

// ListRedeemableUsersResponse wrapper for the ListRedeemableUsers operation
type ListRedeemableUsersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RedeemableUserCollection instances
	RedeemableUserCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

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

var mappingListRedeemableUsersSortOrder = map[string]ListRedeemableUsersSortOrderEnum{
	"ASC":  ListRedeemableUsersSortOrderAsc,
	"DESC": ListRedeemableUsersSortOrderDesc,
}

// GetListRedeemableUsersSortOrderEnumValues Enumerates the set of values for ListRedeemableUsersSortOrderEnum
func GetListRedeemableUsersSortOrderEnumValues() []ListRedeemableUsersSortOrderEnum {
	values := make([]ListRedeemableUsersSortOrderEnum, 0)
	for _, v := range mappingListRedeemableUsersSortOrder {
		values = append(values, v)
	}
	return values
}

// ListRedeemableUsersSortByEnum Enum with underlying type: string
type ListRedeemableUsersSortByEnum string

// Set of constants representing the allowable values for ListRedeemableUsersSortByEnum
const (
	ListRedeemableUsersSortByTimecreated ListRedeemableUsersSortByEnum = "TIMECREATED"
	ListRedeemableUsersSortByTimestart   ListRedeemableUsersSortByEnum = "TIMESTART"
)

var mappingListRedeemableUsersSortBy = map[string]ListRedeemableUsersSortByEnum{
	"TIMECREATED": ListRedeemableUsersSortByTimecreated,
	"TIMESTART":   ListRedeemableUsersSortByTimestart,
}

// GetListRedeemableUsersSortByEnumValues Enumerates the set of values for ListRedeemableUsersSortByEnum
func GetListRedeemableUsersSortByEnumValues() []ListRedeemableUsersSortByEnum {
	values := make([]ListRedeemableUsersSortByEnum, 0)
	for _, v := range mappingListRedeemableUsersSortBy {
		values = append(values, v)
	}
	return values
}
