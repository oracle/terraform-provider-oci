// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListWorkRequestsRequest wrapper for the ListWorkRequests operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListWorkRequests.go.html to see an example of how to use ListWorkRequestsRequest.
type ListWorkRequestsRequest struct {

	// List the Digital Assistant instances that belong to this compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// List only the information for this Digital Assistant instance.
	OdaInstanceId *string `mandatory:"false" contributesTo:"query" name:"odaInstanceId"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. You can specify only one sort order. If no value is specified, then the default is `TIME_ACCEPTED`.
	// The default sort order for the time fields is descending. The default order for `DISPLAYNAME` and `STATUS` is ascending.default: TIME_ACCEPTED
	SortBy ListWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWorkRequestsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWorkRequestsResponse wrapper for the ListWorkRequests operation
type ListWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WorkRequestSummary instances
	Items []WorkRequestSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you contact Oracle
	// about this request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// When you are paging through a list, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the
	// `page` query parameter for the subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestsSortByEnum Enum with underlying type: string
type ListWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortByEnum
const (
	ListWorkRequestsSortByOperationType ListWorkRequestsSortByEnum = "OPERATION_TYPE"
	ListWorkRequestsSortByStatus        ListWorkRequestsSortByEnum = "STATUS"
	ListWorkRequestsSortByTimeAccepted  ListWorkRequestsSortByEnum = "TIME_ACCEPTED"
	ListWorkRequestsSortByTimeStarted   ListWorkRequestsSortByEnum = "TIME_STARTED"
	ListWorkRequestsSortByTimeFinished  ListWorkRequestsSortByEnum = "TIME_FINISHED"
)

var mappingListWorkRequestsSortByEnum = map[string]ListWorkRequestsSortByEnum{
	"OPERATION_TYPE": ListWorkRequestsSortByOperationType,
	"STATUS":         ListWorkRequestsSortByStatus,
	"TIME_ACCEPTED":  ListWorkRequestsSortByTimeAccepted,
	"TIME_STARTED":   ListWorkRequestsSortByTimeStarted,
	"TIME_FINISHED":  ListWorkRequestsSortByTimeFinished,
}

// GetListWorkRequestsSortByEnumValues Enumerates the set of values for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumValues() []ListWorkRequestsSortByEnum {
	values := make([]ListWorkRequestsSortByEnum, 0)
	for _, v := range mappingListWorkRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsSortByEnumStringValues Enumerates the set of values in String for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumStringValues() []string {
	return []string{
		"OPERATION_TYPE",
		"STATUS",
		"TIME_ACCEPTED",
		"TIME_STARTED",
		"TIME_FINISHED",
	}
}

// GetMappingListWorkRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsSortByEnum(val string) (ListWorkRequestsSortByEnum, bool) {
	mappingListWorkRequestsSortByEnumIgnoreCase := make(map[string]ListWorkRequestsSortByEnum)
	for k, v := range mappingListWorkRequestsSortByEnum {
		mappingListWorkRequestsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListWorkRequestsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestsSortOrderEnum Enum with underlying type: string
type ListWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortOrderEnum
const (
	ListWorkRequestsSortOrderAsc  ListWorkRequestsSortOrderEnum = "ASC"
	ListWorkRequestsSortOrderDesc ListWorkRequestsSortOrderEnum = "DESC"
)

var mappingListWorkRequestsSortOrderEnum = map[string]ListWorkRequestsSortOrderEnum{
	"ASC":  ListWorkRequestsSortOrderAsc,
	"DESC": ListWorkRequestsSortOrderDesc,
}

// GetListWorkRequestsSortOrderEnumValues Enumerates the set of values for ListWorkRequestsSortOrderEnum
func GetListWorkRequestsSortOrderEnumValues() []ListWorkRequestsSortOrderEnum {
	values := make([]ListWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListWorkRequestsSortOrderEnum
func GetListWorkRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWorkRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsSortOrderEnum(val string) (ListWorkRequestsSortOrderEnum, bool) {
	mappingListWorkRequestsSortOrderEnumIgnoreCase := make(map[string]ListWorkRequestsSortOrderEnum)
	for k, v := range mappingListWorkRequestsSortOrderEnum {
		mappingListWorkRequestsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListWorkRequestsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
