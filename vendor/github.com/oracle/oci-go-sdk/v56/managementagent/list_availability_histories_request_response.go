// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListAvailabilityHistoriesRequest wrapper for the ListAvailabilityHistories operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementagent/ListAvailabilityHistories.go.html to see an example of how to use ListAvailabilityHistoriesRequest.
type ListAvailabilityHistoriesRequest struct {

	// Unique Management Agent identifier
	ManagementAgentId *string `mandatory:"true" contributesTo:"path" name:"managementAgentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Filter to limit the availability history results to that of time after the input time including the boundary record.
	// Defaulted to current date minus one year.
	// The date and time to be given as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeAvailabilityStatusEndedGreaterThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeAvailabilityStatusEndedGreaterThan"`

	// Filter to limit the availability history results to that of time before the input time including the boundary record
	// Defaulted to current date.
	// The date and time to be given as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeAvailabilityStatusStartedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeAvailabilityStatusStartedLessThan"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAvailabilityHistoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Default order for timeAvailabilityStatusStarted is descending.
	SortBy ListAvailabilityHistoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailabilityHistoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailabilityHistoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAvailabilityHistoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailabilityHistoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAvailabilityHistoriesResponse wrapper for the ListAvailabilityHistories operation
type ListAvailabilityHistoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AvailabilityHistorySummary instances
	Items []AvailabilityHistorySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAvailabilityHistoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailabilityHistoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailabilityHistoriesSortOrderEnum Enum with underlying type: string
type ListAvailabilityHistoriesSortOrderEnum string

// Set of constants representing the allowable values for ListAvailabilityHistoriesSortOrderEnum
const (
	ListAvailabilityHistoriesSortOrderAsc  ListAvailabilityHistoriesSortOrderEnum = "ASC"
	ListAvailabilityHistoriesSortOrderDesc ListAvailabilityHistoriesSortOrderEnum = "DESC"
)

var mappingListAvailabilityHistoriesSortOrder = map[string]ListAvailabilityHistoriesSortOrderEnum{
	"ASC":  ListAvailabilityHistoriesSortOrderAsc,
	"DESC": ListAvailabilityHistoriesSortOrderDesc,
}

// GetListAvailabilityHistoriesSortOrderEnumValues Enumerates the set of values for ListAvailabilityHistoriesSortOrderEnum
func GetListAvailabilityHistoriesSortOrderEnumValues() []ListAvailabilityHistoriesSortOrderEnum {
	values := make([]ListAvailabilityHistoriesSortOrderEnum, 0)
	for _, v := range mappingListAvailabilityHistoriesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListAvailabilityHistoriesSortByEnum Enum with underlying type: string
type ListAvailabilityHistoriesSortByEnum string

// Set of constants representing the allowable values for ListAvailabilityHistoriesSortByEnum
const (
	ListAvailabilityHistoriesSortByTimeavailabilitystatusstarted ListAvailabilityHistoriesSortByEnum = "timeAvailabilityStatusStarted"
)

var mappingListAvailabilityHistoriesSortBy = map[string]ListAvailabilityHistoriesSortByEnum{
	"timeAvailabilityStatusStarted": ListAvailabilityHistoriesSortByTimeavailabilitystatusstarted,
}

// GetListAvailabilityHistoriesSortByEnumValues Enumerates the set of values for ListAvailabilityHistoriesSortByEnum
func GetListAvailabilityHistoriesSortByEnumValues() []ListAvailabilityHistoriesSortByEnum {
	values := make([]ListAvailabilityHistoriesSortByEnum, 0)
	for _, v := range mappingListAvailabilityHistoriesSortBy {
		values = append(values, v)
	}
	return values
}
