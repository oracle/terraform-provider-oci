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

// ListParserActionsRequest wrapper for the ListParserActions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListParserActions.go.html to see an example of how to use ListParserActionsRequest.
type ListParserActionsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The parser action name used for filtering.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The parser action display text used for filtering.
	ActionDisplayText *string `mandatory:"false" contributesTo:"query" name:"actionDisplayText"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListParserActionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned parser actions
	SortBy ListParserActionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListParserActionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListParserActionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListParserActionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListParserActionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListParserActionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListParserActionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListParserActionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListParserActionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListParserActionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListParserActionsResponse wrapper for the ListParserActions operation
type ListParserActionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ParserActionSummaryCollection instances
	ParserActionSummaryCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListParserActionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListParserActionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListParserActionsSortOrderEnum Enum with underlying type: string
type ListParserActionsSortOrderEnum string

// Set of constants representing the allowable values for ListParserActionsSortOrderEnum
const (
	ListParserActionsSortOrderAsc  ListParserActionsSortOrderEnum = "ASC"
	ListParserActionsSortOrderDesc ListParserActionsSortOrderEnum = "DESC"
)

var mappingListParserActionsSortOrderEnum = map[string]ListParserActionsSortOrderEnum{
	"ASC":  ListParserActionsSortOrderAsc,
	"DESC": ListParserActionsSortOrderDesc,
}

var mappingListParserActionsSortOrderEnumLowerCase = map[string]ListParserActionsSortOrderEnum{
	"asc":  ListParserActionsSortOrderAsc,
	"desc": ListParserActionsSortOrderDesc,
}

// GetListParserActionsSortOrderEnumValues Enumerates the set of values for ListParserActionsSortOrderEnum
func GetListParserActionsSortOrderEnumValues() []ListParserActionsSortOrderEnum {
	values := make([]ListParserActionsSortOrderEnum, 0)
	for _, v := range mappingListParserActionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListParserActionsSortOrderEnumStringValues Enumerates the set of values in String for ListParserActionsSortOrderEnum
func GetListParserActionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListParserActionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListParserActionsSortOrderEnum(val string) (ListParserActionsSortOrderEnum, bool) {
	enum, ok := mappingListParserActionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListParserActionsSortByEnum Enum with underlying type: string
type ListParserActionsSortByEnum string

// Set of constants representing the allowable values for ListParserActionsSortByEnum
const (
	ListParserActionsSortByDisplayname ListParserActionsSortByEnum = "displayName"
)

var mappingListParserActionsSortByEnum = map[string]ListParserActionsSortByEnum{
	"displayName": ListParserActionsSortByDisplayname,
}

var mappingListParserActionsSortByEnumLowerCase = map[string]ListParserActionsSortByEnum{
	"displayname": ListParserActionsSortByDisplayname,
}

// GetListParserActionsSortByEnumValues Enumerates the set of values for ListParserActionsSortByEnum
func GetListParserActionsSortByEnumValues() []ListParserActionsSortByEnum {
	values := make([]ListParserActionsSortByEnum, 0)
	for _, v := range mappingListParserActionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListParserActionsSortByEnumStringValues Enumerates the set of values in String for ListParserActionsSortByEnum
func GetListParserActionsSortByEnumStringValues() []string {
	return []string{
		"displayName",
	}
}

// GetMappingListParserActionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListParserActionsSortByEnum(val string) (ListParserActionsSortByEnum, bool) {
	enum, ok := mappingListParserActionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
