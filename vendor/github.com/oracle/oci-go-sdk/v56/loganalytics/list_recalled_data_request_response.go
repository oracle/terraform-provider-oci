// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListRecalledDataRequest wrapper for the ListRecalledData operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListRecalledData.go.html to see an example of how to use ListRecalledDataRequest.
type ListRecalledDataRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// This is the query parameter of which field to sort by. Only one sort order may be provided. Default order for timeDataStarted
	// is descending. If no value is specified timeDataStarted is default.
	SortBy ListRecalledDataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListRecalledDataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// This is the start of the time range for recalled data
	TimeDataStartedGreaterThanOrEqual *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeDataStartedGreaterThanOrEqual"`

	// This is the end of the time range for recalled data
	TimeDataEndedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeDataEndedLessThan"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRecalledDataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRecalledDataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRecalledDataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRecalledDataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListRecalledDataResponse wrapper for the ListRecalledData operation
type ListRecalledDataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecalledDataCollection instances
	RecalledDataCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListRecalledDataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRecalledDataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRecalledDataSortByEnum Enum with underlying type: string
type ListRecalledDataSortByEnum string

// Set of constants representing the allowable values for ListRecalledDataSortByEnum
const (
	ListRecalledDataSortByTimestarted     ListRecalledDataSortByEnum = "timeStarted"
	ListRecalledDataSortByTimedatastarted ListRecalledDataSortByEnum = "timeDataStarted"
)

var mappingListRecalledDataSortBy = map[string]ListRecalledDataSortByEnum{
	"timeStarted":     ListRecalledDataSortByTimestarted,
	"timeDataStarted": ListRecalledDataSortByTimedatastarted,
}

// GetListRecalledDataSortByEnumValues Enumerates the set of values for ListRecalledDataSortByEnum
func GetListRecalledDataSortByEnumValues() []ListRecalledDataSortByEnum {
	values := make([]ListRecalledDataSortByEnum, 0)
	for _, v := range mappingListRecalledDataSortBy {
		values = append(values, v)
	}
	return values
}

// ListRecalledDataSortOrderEnum Enum with underlying type: string
type ListRecalledDataSortOrderEnum string

// Set of constants representing the allowable values for ListRecalledDataSortOrderEnum
const (
	ListRecalledDataSortOrderAsc  ListRecalledDataSortOrderEnum = "ASC"
	ListRecalledDataSortOrderDesc ListRecalledDataSortOrderEnum = "DESC"
)

var mappingListRecalledDataSortOrder = map[string]ListRecalledDataSortOrderEnum{
	"ASC":  ListRecalledDataSortOrderAsc,
	"DESC": ListRecalledDataSortOrderDesc,
}

// GetListRecalledDataSortOrderEnumValues Enumerates the set of values for ListRecalledDataSortOrderEnum
func GetListRecalledDataSortOrderEnumValues() []ListRecalledDataSortOrderEnum {
	values := make([]ListRecalledDataSortOrderEnum, 0)
	for _, v := range mappingListRecalledDataSortOrder {
		values = append(values, v)
	}
	return values
}
