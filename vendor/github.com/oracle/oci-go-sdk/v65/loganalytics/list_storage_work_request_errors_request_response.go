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

// ListStorageWorkRequestErrorsRequest wrapper for the ListStorageWorkRequestErrors operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListStorageWorkRequestErrors.go.html to see an example of how to use ListStorageWorkRequestErrorsRequest.
type ListStorageWorkRequestErrorsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Work Request Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the asynchronous request.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListStorageWorkRequestErrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListStorageWorkRequestErrorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStorageWorkRequestErrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStorageWorkRequestErrorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListStorageWorkRequestErrorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStorageWorkRequestErrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListStorageWorkRequestErrorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListStorageWorkRequestErrorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListStorageWorkRequestErrorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStorageWorkRequestErrorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListStorageWorkRequestErrorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListStorageWorkRequestErrorsResponse wrapper for the ListStorageWorkRequestErrors operation
type ListStorageWorkRequestErrorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestErrorCollection instances
	WorkRequestErrorCollection `presentIn:"body"`

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

func (response ListStorageWorkRequestErrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStorageWorkRequestErrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStorageWorkRequestErrorsSortOrderEnum Enum with underlying type: string
type ListStorageWorkRequestErrorsSortOrderEnum string

// Set of constants representing the allowable values for ListStorageWorkRequestErrorsSortOrderEnum
const (
	ListStorageWorkRequestErrorsSortOrderAsc  ListStorageWorkRequestErrorsSortOrderEnum = "ASC"
	ListStorageWorkRequestErrorsSortOrderDesc ListStorageWorkRequestErrorsSortOrderEnum = "DESC"
)

var mappingListStorageWorkRequestErrorsSortOrderEnum = map[string]ListStorageWorkRequestErrorsSortOrderEnum{
	"ASC":  ListStorageWorkRequestErrorsSortOrderAsc,
	"DESC": ListStorageWorkRequestErrorsSortOrderDesc,
}

var mappingListStorageWorkRequestErrorsSortOrderEnumLowerCase = map[string]ListStorageWorkRequestErrorsSortOrderEnum{
	"asc":  ListStorageWorkRequestErrorsSortOrderAsc,
	"desc": ListStorageWorkRequestErrorsSortOrderDesc,
}

// GetListStorageWorkRequestErrorsSortOrderEnumValues Enumerates the set of values for ListStorageWorkRequestErrorsSortOrderEnum
func GetListStorageWorkRequestErrorsSortOrderEnumValues() []ListStorageWorkRequestErrorsSortOrderEnum {
	values := make([]ListStorageWorkRequestErrorsSortOrderEnum, 0)
	for _, v := range mappingListStorageWorkRequestErrorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListStorageWorkRequestErrorsSortOrderEnumStringValues Enumerates the set of values in String for ListStorageWorkRequestErrorsSortOrderEnum
func GetListStorageWorkRequestErrorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListStorageWorkRequestErrorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStorageWorkRequestErrorsSortOrderEnum(val string) (ListStorageWorkRequestErrorsSortOrderEnum, bool) {
	enum, ok := mappingListStorageWorkRequestErrorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStorageWorkRequestErrorsSortByEnum Enum with underlying type: string
type ListStorageWorkRequestErrorsSortByEnum string

// Set of constants representing the allowable values for ListStorageWorkRequestErrorsSortByEnum
const (
	ListStorageWorkRequestErrorsSortByTimecreated ListStorageWorkRequestErrorsSortByEnum = "timeCreated"
)

var mappingListStorageWorkRequestErrorsSortByEnum = map[string]ListStorageWorkRequestErrorsSortByEnum{
	"timeCreated": ListStorageWorkRequestErrorsSortByTimecreated,
}

var mappingListStorageWorkRequestErrorsSortByEnumLowerCase = map[string]ListStorageWorkRequestErrorsSortByEnum{
	"timecreated": ListStorageWorkRequestErrorsSortByTimecreated,
}

// GetListStorageWorkRequestErrorsSortByEnumValues Enumerates the set of values for ListStorageWorkRequestErrorsSortByEnum
func GetListStorageWorkRequestErrorsSortByEnumValues() []ListStorageWorkRequestErrorsSortByEnum {
	values := make([]ListStorageWorkRequestErrorsSortByEnum, 0)
	for _, v := range mappingListStorageWorkRequestErrorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListStorageWorkRequestErrorsSortByEnumStringValues Enumerates the set of values in String for ListStorageWorkRequestErrorsSortByEnum
func GetListStorageWorkRequestErrorsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListStorageWorkRequestErrorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStorageWorkRequestErrorsSortByEnum(val string) (ListStorageWorkRequestErrorsSortByEnum, bool) {
	enum, ok := mappingListStorageWorkRequestErrorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
