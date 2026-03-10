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

// ListGdpWorkRequestErrorsRequest wrapper for the ListGdpWorkRequestErrors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/gdp/ListGdpWorkRequestErrors.go.html to see an example of how to use ListGdpWorkRequestErrorsRequest.
type ListGdpWorkRequestErrorsRequest struct {

	// The ID of the asynchronous request.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timestamp is descending.
	SortBy ListGdpWorkRequestErrorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListGdpWorkRequestErrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGdpWorkRequestErrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGdpWorkRequestErrorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGdpWorkRequestErrorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGdpWorkRequestErrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGdpWorkRequestErrorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGdpWorkRequestErrorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGdpWorkRequestErrorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGdpWorkRequestErrorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGdpWorkRequestErrorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGdpWorkRequestErrorsResponse wrapper for the ListGdpWorkRequestErrors operation
type ListGdpWorkRequestErrorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestErrorCollection instances
	WorkRequestErrorCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListGdpWorkRequestErrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGdpWorkRequestErrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGdpWorkRequestErrorsSortByEnum Enum with underlying type: string
type ListGdpWorkRequestErrorsSortByEnum string

// Set of constants representing the allowable values for ListGdpWorkRequestErrorsSortByEnum
const (
	ListGdpWorkRequestErrorsSortByTimestamp ListGdpWorkRequestErrorsSortByEnum = "timestamp"
)

var mappingListGdpWorkRequestErrorsSortByEnum = map[string]ListGdpWorkRequestErrorsSortByEnum{
	"timestamp": ListGdpWorkRequestErrorsSortByTimestamp,
}

var mappingListGdpWorkRequestErrorsSortByEnumLowerCase = map[string]ListGdpWorkRequestErrorsSortByEnum{
	"timestamp": ListGdpWorkRequestErrorsSortByTimestamp,
}

// GetListGdpWorkRequestErrorsSortByEnumValues Enumerates the set of values for ListGdpWorkRequestErrorsSortByEnum
func GetListGdpWorkRequestErrorsSortByEnumValues() []ListGdpWorkRequestErrorsSortByEnum {
	values := make([]ListGdpWorkRequestErrorsSortByEnum, 0)
	for _, v := range mappingListGdpWorkRequestErrorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGdpWorkRequestErrorsSortByEnumStringValues Enumerates the set of values in String for ListGdpWorkRequestErrorsSortByEnum
func GetListGdpWorkRequestErrorsSortByEnumStringValues() []string {
	return []string{
		"timestamp",
	}
}

// GetMappingListGdpWorkRequestErrorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGdpWorkRequestErrorsSortByEnum(val string) (ListGdpWorkRequestErrorsSortByEnum, bool) {
	enum, ok := mappingListGdpWorkRequestErrorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGdpWorkRequestErrorsSortOrderEnum Enum with underlying type: string
type ListGdpWorkRequestErrorsSortOrderEnum string

// Set of constants representing the allowable values for ListGdpWorkRequestErrorsSortOrderEnum
const (
	ListGdpWorkRequestErrorsSortOrderAsc  ListGdpWorkRequestErrorsSortOrderEnum = "ASC"
	ListGdpWorkRequestErrorsSortOrderDesc ListGdpWorkRequestErrorsSortOrderEnum = "DESC"
)

var mappingListGdpWorkRequestErrorsSortOrderEnum = map[string]ListGdpWorkRequestErrorsSortOrderEnum{
	"ASC":  ListGdpWorkRequestErrorsSortOrderAsc,
	"DESC": ListGdpWorkRequestErrorsSortOrderDesc,
}

var mappingListGdpWorkRequestErrorsSortOrderEnumLowerCase = map[string]ListGdpWorkRequestErrorsSortOrderEnum{
	"asc":  ListGdpWorkRequestErrorsSortOrderAsc,
	"desc": ListGdpWorkRequestErrorsSortOrderDesc,
}

// GetListGdpWorkRequestErrorsSortOrderEnumValues Enumerates the set of values for ListGdpWorkRequestErrorsSortOrderEnum
func GetListGdpWorkRequestErrorsSortOrderEnumValues() []ListGdpWorkRequestErrorsSortOrderEnum {
	values := make([]ListGdpWorkRequestErrorsSortOrderEnum, 0)
	for _, v := range mappingListGdpWorkRequestErrorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGdpWorkRequestErrorsSortOrderEnumStringValues Enumerates the set of values in String for ListGdpWorkRequestErrorsSortOrderEnum
func GetListGdpWorkRequestErrorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGdpWorkRequestErrorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGdpWorkRequestErrorsSortOrderEnum(val string) (ListGdpWorkRequestErrorsSortOrderEnum, bool) {
	enum, ok := mappingListGdpWorkRequestErrorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
