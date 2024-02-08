// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOverlappingRecallsRequest wrapper for the ListOverlappingRecalls operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListOverlappingRecalls.go.html to see an example of how to use ListOverlappingRecallsRequest.
type ListOverlappingRecallsRequest struct {

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
	SortBy ListOverlappingRecallsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListOverlappingRecallsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// This is the start of the time range for recalled data
	TimeDataStarted *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeDataStarted"`

	// This is the end of the time range for recalled data
	TimeDataEnded *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeDataEnded"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOverlappingRecallsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOverlappingRecallsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOverlappingRecallsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOverlappingRecallsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOverlappingRecallsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOverlappingRecallsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOverlappingRecallsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOverlappingRecallsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOverlappingRecallsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOverlappingRecallsResponse wrapper for the ListOverlappingRecalls operation
type ListOverlappingRecallsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OverlappingRecallCollection instances
	OverlappingRecallCollection `presentIn:"body"`

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

func (response ListOverlappingRecallsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOverlappingRecallsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOverlappingRecallsSortByEnum Enum with underlying type: string
type ListOverlappingRecallsSortByEnum string

// Set of constants representing the allowable values for ListOverlappingRecallsSortByEnum
const (
	ListOverlappingRecallsSortByTimestarted     ListOverlappingRecallsSortByEnum = "timeStarted"
	ListOverlappingRecallsSortByTimedatastarted ListOverlappingRecallsSortByEnum = "timeDataStarted"
)

var mappingListOverlappingRecallsSortByEnum = map[string]ListOverlappingRecallsSortByEnum{
	"timeStarted":     ListOverlappingRecallsSortByTimestarted,
	"timeDataStarted": ListOverlappingRecallsSortByTimedatastarted,
}

var mappingListOverlappingRecallsSortByEnumLowerCase = map[string]ListOverlappingRecallsSortByEnum{
	"timestarted":     ListOverlappingRecallsSortByTimestarted,
	"timedatastarted": ListOverlappingRecallsSortByTimedatastarted,
}

// GetListOverlappingRecallsSortByEnumValues Enumerates the set of values for ListOverlappingRecallsSortByEnum
func GetListOverlappingRecallsSortByEnumValues() []ListOverlappingRecallsSortByEnum {
	values := make([]ListOverlappingRecallsSortByEnum, 0)
	for _, v := range mappingListOverlappingRecallsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOverlappingRecallsSortByEnumStringValues Enumerates the set of values in String for ListOverlappingRecallsSortByEnum
func GetListOverlappingRecallsSortByEnumStringValues() []string {
	return []string{
		"timeStarted",
		"timeDataStarted",
	}
}

// GetMappingListOverlappingRecallsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOverlappingRecallsSortByEnum(val string) (ListOverlappingRecallsSortByEnum, bool) {
	enum, ok := mappingListOverlappingRecallsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOverlappingRecallsSortOrderEnum Enum with underlying type: string
type ListOverlappingRecallsSortOrderEnum string

// Set of constants representing the allowable values for ListOverlappingRecallsSortOrderEnum
const (
	ListOverlappingRecallsSortOrderAsc  ListOverlappingRecallsSortOrderEnum = "ASC"
	ListOverlappingRecallsSortOrderDesc ListOverlappingRecallsSortOrderEnum = "DESC"
)

var mappingListOverlappingRecallsSortOrderEnum = map[string]ListOverlappingRecallsSortOrderEnum{
	"ASC":  ListOverlappingRecallsSortOrderAsc,
	"DESC": ListOverlappingRecallsSortOrderDesc,
}

var mappingListOverlappingRecallsSortOrderEnumLowerCase = map[string]ListOverlappingRecallsSortOrderEnum{
	"asc":  ListOverlappingRecallsSortOrderAsc,
	"desc": ListOverlappingRecallsSortOrderDesc,
}

// GetListOverlappingRecallsSortOrderEnumValues Enumerates the set of values for ListOverlappingRecallsSortOrderEnum
func GetListOverlappingRecallsSortOrderEnumValues() []ListOverlappingRecallsSortOrderEnum {
	values := make([]ListOverlappingRecallsSortOrderEnum, 0)
	for _, v := range mappingListOverlappingRecallsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOverlappingRecallsSortOrderEnumStringValues Enumerates the set of values in String for ListOverlappingRecallsSortOrderEnum
func GetListOverlappingRecallsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOverlappingRecallsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOverlappingRecallsSortOrderEnum(val string) (ListOverlappingRecallsSortOrderEnum, bool) {
	enum, ok := mappingListOverlappingRecallsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
