// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListAutoAssociationsRequest wrapper for the ListAutoAssociations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListAutoAssociations.go.html to see an example of how to use ListAutoAssociationsRequest.
type ListAutoAssociationsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The source name.
	SourceName *string `mandatory:"true" contributesTo:"path" name:"sourceName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The attribute used to sort the returned auto association information.
	SortBy ListAutoAssociationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAutoAssociationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutoAssociationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutoAssociationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutoAssociationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutoAssociationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutoAssociationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAutoAssociationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAutoAssociationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutoAssociationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAutoAssociationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutoAssociationsResponse wrapper for the ListAutoAssociations operation
type ListAutoAssociationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AutoAssociationCollection instances
	AutoAssociationCollection `presentIn:"body"`

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

func (response ListAutoAssociationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutoAssociationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutoAssociationsSortByEnum Enum with underlying type: string
type ListAutoAssociationsSortByEnum string

// Set of constants representing the allowable values for ListAutoAssociationsSortByEnum
const (
	ListAutoAssociationsSortByIsenabled ListAutoAssociationsSortByEnum = "isEnabled"
)

var mappingListAutoAssociationsSortByEnum = map[string]ListAutoAssociationsSortByEnum{
	"isEnabled": ListAutoAssociationsSortByIsenabled,
}

// GetListAutoAssociationsSortByEnumValues Enumerates the set of values for ListAutoAssociationsSortByEnum
func GetListAutoAssociationsSortByEnumValues() []ListAutoAssociationsSortByEnum {
	values := make([]ListAutoAssociationsSortByEnum, 0)
	for _, v := range mappingListAutoAssociationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutoAssociationsSortByEnumStringValues Enumerates the set of values in String for ListAutoAssociationsSortByEnum
func GetListAutoAssociationsSortByEnumStringValues() []string {
	return []string{
		"isEnabled",
	}
}

// GetMappingListAutoAssociationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutoAssociationsSortByEnum(val string) (ListAutoAssociationsSortByEnum, bool) {
	mappingListAutoAssociationsSortByEnumIgnoreCase := make(map[string]ListAutoAssociationsSortByEnum)
	for k, v := range mappingListAutoAssociationsSortByEnum {
		mappingListAutoAssociationsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAutoAssociationsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutoAssociationsSortOrderEnum Enum with underlying type: string
type ListAutoAssociationsSortOrderEnum string

// Set of constants representing the allowable values for ListAutoAssociationsSortOrderEnum
const (
	ListAutoAssociationsSortOrderAsc  ListAutoAssociationsSortOrderEnum = "ASC"
	ListAutoAssociationsSortOrderDesc ListAutoAssociationsSortOrderEnum = "DESC"
)

var mappingListAutoAssociationsSortOrderEnum = map[string]ListAutoAssociationsSortOrderEnum{
	"ASC":  ListAutoAssociationsSortOrderAsc,
	"DESC": ListAutoAssociationsSortOrderDesc,
}

// GetListAutoAssociationsSortOrderEnumValues Enumerates the set of values for ListAutoAssociationsSortOrderEnum
func GetListAutoAssociationsSortOrderEnumValues() []ListAutoAssociationsSortOrderEnum {
	values := make([]ListAutoAssociationsSortOrderEnum, 0)
	for _, v := range mappingListAutoAssociationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutoAssociationsSortOrderEnumStringValues Enumerates the set of values in String for ListAutoAssociationsSortOrderEnum
func GetListAutoAssociationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAutoAssociationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutoAssociationsSortOrderEnum(val string) (ListAutoAssociationsSortOrderEnum, bool) {
	mappingListAutoAssociationsSortOrderEnumIgnoreCase := make(map[string]ListAutoAssociationsSortOrderEnum)
	for k, v := range mappingListAutoAssociationsSortOrderEnum {
		mappingListAutoAssociationsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAutoAssociationsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
