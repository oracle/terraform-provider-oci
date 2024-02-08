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

// ListSourceLabelOperatorsRequest wrapper for the ListSourceLabelOperators operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListSourceLabelOperators.go.html to see an example of how to use ListSourceLabelOperatorsRequest.
type ListSourceLabelOperatorsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The attribute used to sort the returned items
	SortBy ListSourceLabelOperatorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSourceLabelOperatorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSourceLabelOperatorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSourceLabelOperatorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSourceLabelOperatorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSourceLabelOperatorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSourceLabelOperatorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSourceLabelOperatorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSourceLabelOperatorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSourceLabelOperatorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSourceLabelOperatorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSourceLabelOperatorsResponse wrapper for the ListSourceLabelOperators operation
type ListSourceLabelOperatorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsLabelOperatorCollection instances
	LogAnalyticsLabelOperatorCollection `presentIn:"body"`

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

func (response ListSourceLabelOperatorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSourceLabelOperatorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSourceLabelOperatorsSortByEnum Enum with underlying type: string
type ListSourceLabelOperatorsSortByEnum string

// Set of constants representing the allowable values for ListSourceLabelOperatorsSortByEnum
const (
	ListSourceLabelOperatorsSortByName ListSourceLabelOperatorsSortByEnum = "name"
)

var mappingListSourceLabelOperatorsSortByEnum = map[string]ListSourceLabelOperatorsSortByEnum{
	"name": ListSourceLabelOperatorsSortByName,
}

var mappingListSourceLabelOperatorsSortByEnumLowerCase = map[string]ListSourceLabelOperatorsSortByEnum{
	"name": ListSourceLabelOperatorsSortByName,
}

// GetListSourceLabelOperatorsSortByEnumValues Enumerates the set of values for ListSourceLabelOperatorsSortByEnum
func GetListSourceLabelOperatorsSortByEnumValues() []ListSourceLabelOperatorsSortByEnum {
	values := make([]ListSourceLabelOperatorsSortByEnum, 0)
	for _, v := range mappingListSourceLabelOperatorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSourceLabelOperatorsSortByEnumStringValues Enumerates the set of values in String for ListSourceLabelOperatorsSortByEnum
func GetListSourceLabelOperatorsSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListSourceLabelOperatorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSourceLabelOperatorsSortByEnum(val string) (ListSourceLabelOperatorsSortByEnum, bool) {
	enum, ok := mappingListSourceLabelOperatorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSourceLabelOperatorsSortOrderEnum Enum with underlying type: string
type ListSourceLabelOperatorsSortOrderEnum string

// Set of constants representing the allowable values for ListSourceLabelOperatorsSortOrderEnum
const (
	ListSourceLabelOperatorsSortOrderAsc  ListSourceLabelOperatorsSortOrderEnum = "ASC"
	ListSourceLabelOperatorsSortOrderDesc ListSourceLabelOperatorsSortOrderEnum = "DESC"
)

var mappingListSourceLabelOperatorsSortOrderEnum = map[string]ListSourceLabelOperatorsSortOrderEnum{
	"ASC":  ListSourceLabelOperatorsSortOrderAsc,
	"DESC": ListSourceLabelOperatorsSortOrderDesc,
}

var mappingListSourceLabelOperatorsSortOrderEnumLowerCase = map[string]ListSourceLabelOperatorsSortOrderEnum{
	"asc":  ListSourceLabelOperatorsSortOrderAsc,
	"desc": ListSourceLabelOperatorsSortOrderDesc,
}

// GetListSourceLabelOperatorsSortOrderEnumValues Enumerates the set of values for ListSourceLabelOperatorsSortOrderEnum
func GetListSourceLabelOperatorsSortOrderEnumValues() []ListSourceLabelOperatorsSortOrderEnum {
	values := make([]ListSourceLabelOperatorsSortOrderEnum, 0)
	for _, v := range mappingListSourceLabelOperatorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSourceLabelOperatorsSortOrderEnumStringValues Enumerates the set of values in String for ListSourceLabelOperatorsSortOrderEnum
func GetListSourceLabelOperatorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSourceLabelOperatorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSourceLabelOperatorsSortOrderEnum(val string) (ListSourceLabelOperatorsSortOrderEnum, bool) {
	enum, ok := mappingListSourceLabelOperatorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
