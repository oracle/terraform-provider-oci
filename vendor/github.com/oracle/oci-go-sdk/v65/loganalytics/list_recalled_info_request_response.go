// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRecalledInfoRequest wrapper for the ListRecalledInfo operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListRecalledInfo.go.html to see an example of how to use ListRecalledInfoRequest.
type ListRecalledInfoRequest struct {

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
	SortBy ListRecalledInfoSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListRecalledInfoSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// This is the start of the time range for recalled data
	TimeDataStartedGreaterThanOrEqual *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeDataStartedGreaterThanOrEqual"`

	// This is the end of the time range for recalled data
	TimeDataEndedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeDataEndedLessThan"`

	// This is the set of logsets to filter recalled collection by if any
	LogSets []string `contributesTo:"query" name:"logSets" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRecalledInfoRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRecalledInfoRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRecalledInfoRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRecalledInfoRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRecalledInfoRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRecalledInfoSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRecalledInfoSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecalledInfoSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRecalledInfoSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRecalledInfoResponse wrapper for the ListRecalledInfo operation
type ListRecalledInfoResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecalledInfoCollection instances
	RecalledInfoCollection `presentIn:"body"`

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

func (response ListRecalledInfoResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRecalledInfoResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRecalledInfoSortByEnum Enum with underlying type: string
type ListRecalledInfoSortByEnum string

// Set of constants representing the allowable values for ListRecalledInfoSortByEnum
const (
	ListRecalledInfoSortByTimestarted     ListRecalledInfoSortByEnum = "timeStarted"
	ListRecalledInfoSortByTimedatastarted ListRecalledInfoSortByEnum = "timeDataStarted"
)

var mappingListRecalledInfoSortByEnum = map[string]ListRecalledInfoSortByEnum{
	"timeStarted":     ListRecalledInfoSortByTimestarted,
	"timeDataStarted": ListRecalledInfoSortByTimedatastarted,
}

var mappingListRecalledInfoSortByEnumLowerCase = map[string]ListRecalledInfoSortByEnum{
	"timestarted":     ListRecalledInfoSortByTimestarted,
	"timedatastarted": ListRecalledInfoSortByTimedatastarted,
}

// GetListRecalledInfoSortByEnumValues Enumerates the set of values for ListRecalledInfoSortByEnum
func GetListRecalledInfoSortByEnumValues() []ListRecalledInfoSortByEnum {
	values := make([]ListRecalledInfoSortByEnum, 0)
	for _, v := range mappingListRecalledInfoSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecalledInfoSortByEnumStringValues Enumerates the set of values in String for ListRecalledInfoSortByEnum
func GetListRecalledInfoSortByEnumStringValues() []string {
	return []string{
		"timeStarted",
		"timeDataStarted",
	}
}

// GetMappingListRecalledInfoSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecalledInfoSortByEnum(val string) (ListRecalledInfoSortByEnum, bool) {
	enum, ok := mappingListRecalledInfoSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecalledInfoSortOrderEnum Enum with underlying type: string
type ListRecalledInfoSortOrderEnum string

// Set of constants representing the allowable values for ListRecalledInfoSortOrderEnum
const (
	ListRecalledInfoSortOrderAsc  ListRecalledInfoSortOrderEnum = "ASC"
	ListRecalledInfoSortOrderDesc ListRecalledInfoSortOrderEnum = "DESC"
)

var mappingListRecalledInfoSortOrderEnum = map[string]ListRecalledInfoSortOrderEnum{
	"ASC":  ListRecalledInfoSortOrderAsc,
	"DESC": ListRecalledInfoSortOrderDesc,
}

var mappingListRecalledInfoSortOrderEnumLowerCase = map[string]ListRecalledInfoSortOrderEnum{
	"asc":  ListRecalledInfoSortOrderAsc,
	"desc": ListRecalledInfoSortOrderDesc,
}

// GetListRecalledInfoSortOrderEnumValues Enumerates the set of values for ListRecalledInfoSortOrderEnum
func GetListRecalledInfoSortOrderEnumValues() []ListRecalledInfoSortOrderEnum {
	values := make([]ListRecalledInfoSortOrderEnum, 0)
	for _, v := range mappingListRecalledInfoSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecalledInfoSortOrderEnumStringValues Enumerates the set of values in String for ListRecalledInfoSortOrderEnum
func GetListRecalledInfoSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRecalledInfoSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecalledInfoSortOrderEnum(val string) (ListRecalledInfoSortOrderEnum, bool) {
	enum, ok := mappingListRecalledInfoSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
