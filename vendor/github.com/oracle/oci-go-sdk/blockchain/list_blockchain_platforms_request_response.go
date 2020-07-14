// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package blockchain

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListBlockchainPlatformsRequest wrapper for the ListBlockchainPlatforms operation
type ListBlockchainPlatformsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The page at which to start retrieving results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListBlockchainPlatformsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListBlockchainPlatformsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to only return resources that match the given lifecycle state.
	// The state value is case-insensitive.
	LifecycleState BlockchainPlatformLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBlockchainPlatformsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBlockchainPlatformsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBlockchainPlatformsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListBlockchainPlatformsResponse wrapper for the ListBlockchainPlatforms operation
type ListBlockchainPlatformsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BlockchainPlatformCollection instances
	BlockchainPlatformCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBlockchainPlatformsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBlockchainPlatformsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBlockchainPlatformsSortOrderEnum Enum with underlying type: string
type ListBlockchainPlatformsSortOrderEnum string

// Set of constants representing the allowable values for ListBlockchainPlatformsSortOrderEnum
const (
	ListBlockchainPlatformsSortOrderAsc  ListBlockchainPlatformsSortOrderEnum = "ASC"
	ListBlockchainPlatformsSortOrderDesc ListBlockchainPlatformsSortOrderEnum = "DESC"
)

var mappingListBlockchainPlatformsSortOrder = map[string]ListBlockchainPlatformsSortOrderEnum{
	"ASC":  ListBlockchainPlatformsSortOrderAsc,
	"DESC": ListBlockchainPlatformsSortOrderDesc,
}

// GetListBlockchainPlatformsSortOrderEnumValues Enumerates the set of values for ListBlockchainPlatformsSortOrderEnum
func GetListBlockchainPlatformsSortOrderEnumValues() []ListBlockchainPlatformsSortOrderEnum {
	values := make([]ListBlockchainPlatformsSortOrderEnum, 0)
	for _, v := range mappingListBlockchainPlatformsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListBlockchainPlatformsSortByEnum Enum with underlying type: string
type ListBlockchainPlatformsSortByEnum string

// Set of constants representing the allowable values for ListBlockchainPlatformsSortByEnum
const (
	ListBlockchainPlatformsSortByTimecreated ListBlockchainPlatformsSortByEnum = "timeCreated"
	ListBlockchainPlatformsSortByDisplayname ListBlockchainPlatformsSortByEnum = "displayName"
)

var mappingListBlockchainPlatformsSortBy = map[string]ListBlockchainPlatformsSortByEnum{
	"timeCreated": ListBlockchainPlatformsSortByTimecreated,
	"displayName": ListBlockchainPlatformsSortByDisplayname,
}

// GetListBlockchainPlatformsSortByEnumValues Enumerates the set of values for ListBlockchainPlatformsSortByEnum
func GetListBlockchainPlatformsSortByEnumValues() []ListBlockchainPlatformsSortByEnum {
	values := make([]ListBlockchainPlatformsSortByEnum, 0)
	for _, v := range mappingListBlockchainPlatformsSortBy {
		values = append(values, v)
	}
	return values
}
