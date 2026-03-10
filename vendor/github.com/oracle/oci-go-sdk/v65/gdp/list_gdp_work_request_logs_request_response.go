// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package gdp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListGdpWorkRequestLogsRequest wrapper for the ListGdpWorkRequestLogs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/gdp/ListGdpWorkRequestLogs.go.html to see an example of how to use ListGdpWorkRequestLogsRequest.
type ListGdpWorkRequestLogsRequest struct {

	// The ID of the asynchronous request.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timestamp is descending.
	SortBy ListGdpWorkRequestLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListGdpWorkRequestLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGdpWorkRequestLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGdpWorkRequestLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGdpWorkRequestLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGdpWorkRequestLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGdpWorkRequestLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGdpWorkRequestLogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGdpWorkRequestLogsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGdpWorkRequestLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGdpWorkRequestLogsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGdpWorkRequestLogsResponse wrapper for the ListGdpWorkRequestLogs operation
type ListGdpWorkRequestLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestLogEntryCollection instances
	WorkRequestLogEntryCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListGdpWorkRequestLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGdpWorkRequestLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGdpWorkRequestLogsSortByEnum Enum with underlying type: string
type ListGdpWorkRequestLogsSortByEnum string

// Set of constants representing the allowable values for ListGdpWorkRequestLogsSortByEnum
const (
	ListGdpWorkRequestLogsSortByTimestamp ListGdpWorkRequestLogsSortByEnum = "timestamp"
)

var mappingListGdpWorkRequestLogsSortByEnum = map[string]ListGdpWorkRequestLogsSortByEnum{
	"timestamp": ListGdpWorkRequestLogsSortByTimestamp,
}

var mappingListGdpWorkRequestLogsSortByEnumLowerCase = map[string]ListGdpWorkRequestLogsSortByEnum{
	"timestamp": ListGdpWorkRequestLogsSortByTimestamp,
}

// GetListGdpWorkRequestLogsSortByEnumValues Enumerates the set of values for ListGdpWorkRequestLogsSortByEnum
func GetListGdpWorkRequestLogsSortByEnumValues() []ListGdpWorkRequestLogsSortByEnum {
	values := make([]ListGdpWorkRequestLogsSortByEnum, 0)
	for _, v := range mappingListGdpWorkRequestLogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGdpWorkRequestLogsSortByEnumStringValues Enumerates the set of values in String for ListGdpWorkRequestLogsSortByEnum
func GetListGdpWorkRequestLogsSortByEnumStringValues() []string {
	return []string{
		"timestamp",
	}
}

// GetMappingListGdpWorkRequestLogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGdpWorkRequestLogsSortByEnum(val string) (ListGdpWorkRequestLogsSortByEnum, bool) {
	enum, ok := mappingListGdpWorkRequestLogsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGdpWorkRequestLogsSortOrderEnum Enum with underlying type: string
type ListGdpWorkRequestLogsSortOrderEnum string

// Set of constants representing the allowable values for ListGdpWorkRequestLogsSortOrderEnum
const (
	ListGdpWorkRequestLogsSortOrderAsc  ListGdpWorkRequestLogsSortOrderEnum = "ASC"
	ListGdpWorkRequestLogsSortOrderDesc ListGdpWorkRequestLogsSortOrderEnum = "DESC"
)

var mappingListGdpWorkRequestLogsSortOrderEnum = map[string]ListGdpWorkRequestLogsSortOrderEnum{
	"ASC":  ListGdpWorkRequestLogsSortOrderAsc,
	"DESC": ListGdpWorkRequestLogsSortOrderDesc,
}

var mappingListGdpWorkRequestLogsSortOrderEnumLowerCase = map[string]ListGdpWorkRequestLogsSortOrderEnum{
	"asc":  ListGdpWorkRequestLogsSortOrderAsc,
	"desc": ListGdpWorkRequestLogsSortOrderDesc,
}

// GetListGdpWorkRequestLogsSortOrderEnumValues Enumerates the set of values for ListGdpWorkRequestLogsSortOrderEnum
func GetListGdpWorkRequestLogsSortOrderEnumValues() []ListGdpWorkRequestLogsSortOrderEnum {
	values := make([]ListGdpWorkRequestLogsSortOrderEnum, 0)
	for _, v := range mappingListGdpWorkRequestLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGdpWorkRequestLogsSortOrderEnumStringValues Enumerates the set of values in String for ListGdpWorkRequestLogsSortOrderEnum
func GetListGdpWorkRequestLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGdpWorkRequestLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGdpWorkRequestLogsSortOrderEnum(val string) (ListGdpWorkRequestLogsSortOrderEnum, bool) {
	enum, ok := mappingListGdpWorkRequestLogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
