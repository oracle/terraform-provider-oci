// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListImpactedResourcesRequest wrapper for the ListImpactedResources operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListImpactedResources.go.html to see an example of how to use ListImpactedResourcesRequest.
type ListImpactedResourcesRequest struct {

	// OCId of the problem.
	ProblemId *string `mandatory:"true" contributesTo:"path" name:"problemId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListImpactedResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListImpactedResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListImpactedResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListImpactedResourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListImpactedResourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListImpactedResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListImpactedResourcesResponse wrapper for the ListImpactedResources operation
type ListImpactedResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ImpactedResourceCollection instances
	ImpactedResourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListImpactedResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListImpactedResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListImpactedResourcesSortOrderEnum Enum with underlying type: string
type ListImpactedResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListImpactedResourcesSortOrderEnum
const (
	ListImpactedResourcesSortOrderAsc  ListImpactedResourcesSortOrderEnum = "ASC"
	ListImpactedResourcesSortOrderDesc ListImpactedResourcesSortOrderEnum = "DESC"
)

var mappingListImpactedResourcesSortOrder = map[string]ListImpactedResourcesSortOrderEnum{
	"ASC":  ListImpactedResourcesSortOrderAsc,
	"DESC": ListImpactedResourcesSortOrderDesc,
}

// GetListImpactedResourcesSortOrderEnumValues Enumerates the set of values for ListImpactedResourcesSortOrderEnum
func GetListImpactedResourcesSortOrderEnumValues() []ListImpactedResourcesSortOrderEnum {
	values := make([]ListImpactedResourcesSortOrderEnum, 0)
	for _, v := range mappingListImpactedResourcesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListImpactedResourcesSortByEnum Enum with underlying type: string
type ListImpactedResourcesSortByEnum string

// Set of constants representing the allowable values for ListImpactedResourcesSortByEnum
const (
	ListImpactedResourcesSortByTimecreated ListImpactedResourcesSortByEnum = "timeCreated"
)

var mappingListImpactedResourcesSortBy = map[string]ListImpactedResourcesSortByEnum{
	"timeCreated": ListImpactedResourcesSortByTimecreated,
}

// GetListImpactedResourcesSortByEnumValues Enumerates the set of values for ListImpactedResourcesSortByEnum
func GetListImpactedResourcesSortByEnumValues() []ListImpactedResourcesSortByEnum {
	values := make([]ListImpactedResourcesSortByEnum, 0)
	for _, v := range mappingListImpactedResourcesSortBy {
		values = append(values, v)
	}
	return values
}
