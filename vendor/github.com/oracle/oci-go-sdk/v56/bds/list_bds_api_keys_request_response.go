// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bds

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListBdsApiKeysRequest wrapper for the ListBdsApiKeys operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListBdsApiKeys.go.html to see an example of how to use ListBdsApiKeysRequest.
type ListBdsApiKeysRequest struct {

	// The OCID of the cluster.
	BdsInstanceId *string `mandatory:"true" contributesTo:"path" name:"bdsInstanceId"`

	// The state of the API key.
	LifecycleState BdsApiKeyLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The OCID of the user for whom the API key belongs.
	UserId *string `mandatory:"false" contributesTo:"query" name:"userId"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListBdsApiKeysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListBdsApiKeysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBdsApiKeysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBdsApiKeysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBdsApiKeysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBdsApiKeysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListBdsApiKeysResponse wrapper for the ListBdsApiKeys operation
type ListBdsApiKeysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []BdsApiKeySummary instances
	Items []BdsApiKeySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBdsApiKeysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBdsApiKeysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBdsApiKeysSortByEnum Enum with underlying type: string
type ListBdsApiKeysSortByEnum string

// Set of constants representing the allowable values for ListBdsApiKeysSortByEnum
const (
	ListBdsApiKeysSortByTimecreated ListBdsApiKeysSortByEnum = "timeCreated"
	ListBdsApiKeysSortByDisplayname ListBdsApiKeysSortByEnum = "displayName"
)

var mappingListBdsApiKeysSortBy = map[string]ListBdsApiKeysSortByEnum{
	"timeCreated": ListBdsApiKeysSortByTimecreated,
	"displayName": ListBdsApiKeysSortByDisplayname,
}

// GetListBdsApiKeysSortByEnumValues Enumerates the set of values for ListBdsApiKeysSortByEnum
func GetListBdsApiKeysSortByEnumValues() []ListBdsApiKeysSortByEnum {
	values := make([]ListBdsApiKeysSortByEnum, 0)
	for _, v := range mappingListBdsApiKeysSortBy {
		values = append(values, v)
	}
	return values
}

// ListBdsApiKeysSortOrderEnum Enum with underlying type: string
type ListBdsApiKeysSortOrderEnum string

// Set of constants representing the allowable values for ListBdsApiKeysSortOrderEnum
const (
	ListBdsApiKeysSortOrderAsc  ListBdsApiKeysSortOrderEnum = "ASC"
	ListBdsApiKeysSortOrderDesc ListBdsApiKeysSortOrderEnum = "DESC"
)

var mappingListBdsApiKeysSortOrder = map[string]ListBdsApiKeysSortOrderEnum{
	"ASC":  ListBdsApiKeysSortOrderAsc,
	"DESC": ListBdsApiKeysSortOrderDesc,
}

// GetListBdsApiKeysSortOrderEnumValues Enumerates the set of values for ListBdsApiKeysSortOrderEnum
func GetListBdsApiKeysSortOrderEnumValues() []ListBdsApiKeysSortOrderEnum {
	values := make([]ListBdsApiKeysSortOrderEnum, 0)
	for _, v := range mappingListBdsApiKeysSortOrder {
		values = append(values, v)
	}
	return values
}
