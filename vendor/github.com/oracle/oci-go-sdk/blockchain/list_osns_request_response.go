// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package blockchain

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListOsnsRequest wrapper for the ListOsns operation
type ListOsnsRequest struct {

	// Unique service identifier.
	BlockchainPlatformId *string `mandatory:"true" contributesTo:"path" name:"blockchainPlatformId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOsnsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListOsnsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The page at which to start retrieving results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOsnsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOsnsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOsnsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListOsnsResponse wrapper for the ListOsns operation
type ListOsnsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OsnCollection instances
	OsnCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOsnsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOsnsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOsnsSortOrderEnum Enum with underlying type: string
type ListOsnsSortOrderEnum string

// Set of constants representing the allowable values for ListOsnsSortOrderEnum
const (
	ListOsnsSortOrderAsc  ListOsnsSortOrderEnum = "ASC"
	ListOsnsSortOrderDesc ListOsnsSortOrderEnum = "DESC"
)

var mappingListOsnsSortOrder = map[string]ListOsnsSortOrderEnum{
	"ASC":  ListOsnsSortOrderAsc,
	"DESC": ListOsnsSortOrderDesc,
}

// GetListOsnsSortOrderEnumValues Enumerates the set of values for ListOsnsSortOrderEnum
func GetListOsnsSortOrderEnumValues() []ListOsnsSortOrderEnum {
	values := make([]ListOsnsSortOrderEnum, 0)
	for _, v := range mappingListOsnsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListOsnsSortByEnum Enum with underlying type: string
type ListOsnsSortByEnum string

// Set of constants representing the allowable values for ListOsnsSortByEnum
const (
	ListOsnsSortByTimecreated ListOsnsSortByEnum = "timeCreated"
	ListOsnsSortByDisplayname ListOsnsSortByEnum = "displayName"
)

var mappingListOsnsSortBy = map[string]ListOsnsSortByEnum{
	"timeCreated": ListOsnsSortByTimecreated,
	"displayName": ListOsnsSortByDisplayname,
}

// GetListOsnsSortByEnumValues Enumerates the set of values for ListOsnsSortByEnum
func GetListOsnsSortByEnumValues() []ListOsnsSortByEnum {
	values := make([]ListOsnsSortByEnum, 0)
	for _, v := range mappingListOsnsSortBy {
		values = append(values, v)
	}
	return values
}
