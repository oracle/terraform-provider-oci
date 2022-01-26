// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waf

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListNetworkAddressListsRequest wrapper for the ListNetworkAddressLists operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ListNetworkAddressLists.go.html to see an example of how to use ListNetworkAddressListsRequest.
type ListNetworkAddressListsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycleState.
	LifecycleState []NetworkAddressListLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only the NetworkAddressList with the given OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListNetworkAddressListsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending.
	// Default order for displayName is ascending.
	// If no value is specified timeCreated is default.
	SortBy ListNetworkAddressListsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNetworkAddressListsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNetworkAddressListsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNetworkAddressListsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNetworkAddressListsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListNetworkAddressListsResponse wrapper for the ListNetworkAddressLists operation
type ListNetworkAddressListsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of NetworkAddressListCollection instances
	NetworkAddressListCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListNetworkAddressListsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNetworkAddressListsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNetworkAddressListsSortOrderEnum Enum with underlying type: string
type ListNetworkAddressListsSortOrderEnum string

// Set of constants representing the allowable values for ListNetworkAddressListsSortOrderEnum
const (
	ListNetworkAddressListsSortOrderAsc  ListNetworkAddressListsSortOrderEnum = "ASC"
	ListNetworkAddressListsSortOrderDesc ListNetworkAddressListsSortOrderEnum = "DESC"
)

var mappingListNetworkAddressListsSortOrder = map[string]ListNetworkAddressListsSortOrderEnum{
	"ASC":  ListNetworkAddressListsSortOrderAsc,
	"DESC": ListNetworkAddressListsSortOrderDesc,
}

// GetListNetworkAddressListsSortOrderEnumValues Enumerates the set of values for ListNetworkAddressListsSortOrderEnum
func GetListNetworkAddressListsSortOrderEnumValues() []ListNetworkAddressListsSortOrderEnum {
	values := make([]ListNetworkAddressListsSortOrderEnum, 0)
	for _, v := range mappingListNetworkAddressListsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListNetworkAddressListsSortByEnum Enum with underlying type: string
type ListNetworkAddressListsSortByEnum string

// Set of constants representing the allowable values for ListNetworkAddressListsSortByEnum
const (
	ListNetworkAddressListsSortByTimecreated ListNetworkAddressListsSortByEnum = "timeCreated"
	ListNetworkAddressListsSortByDisplayname ListNetworkAddressListsSortByEnum = "displayName"
)

var mappingListNetworkAddressListsSortBy = map[string]ListNetworkAddressListsSortByEnum{
	"timeCreated": ListNetworkAddressListsSortByTimecreated,
	"displayName": ListNetworkAddressListsSortByDisplayname,
}

// GetListNetworkAddressListsSortByEnumValues Enumerates the set of values for ListNetworkAddressListsSortByEnum
func GetListNetworkAddressListsSortByEnumValues() []ListNetworkAddressListsSortByEnum {
	values := make([]ListNetworkAddressListsSortByEnum, 0)
	for _, v := range mappingListNetworkAddressListsSortBy {
		values = append(values, v)
	}
	return values
}
