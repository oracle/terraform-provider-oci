// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListResourcePortsRequest wrapper for the ListResourcePorts operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResourcePorts.go.html to see an example of how to use ListResourcePortsRequest.
type ListResourcePortsRequest struct {

	// CloudGuard resource OCID
	ResourceId *string `mandatory:"true" contributesTo:"path" name:"resourceId"`

	// open port associated with the resource.
	OpenPort *string `mandatory:"false" contributesTo:"query" name:"openPort"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListResourcePortsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListResourcePortsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourcePortsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourcePortsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourcePortsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourcePortsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResourcePortsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResourcePortsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResourcePortsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourcePortsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResourcePortsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResourcePortsResponse wrapper for the ListResourcePorts operation
type ListResourcePortsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourcePortCollection instances
	ResourcePortCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResourcePortsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourcePortsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourcePortsSortOrderEnum Enum with underlying type: string
type ListResourcePortsSortOrderEnum string

// Set of constants representing the allowable values for ListResourcePortsSortOrderEnum
const (
	ListResourcePortsSortOrderAsc  ListResourcePortsSortOrderEnum = "ASC"
	ListResourcePortsSortOrderDesc ListResourcePortsSortOrderEnum = "DESC"
)

var mappingListResourcePortsSortOrderEnum = map[string]ListResourcePortsSortOrderEnum{
	"ASC":  ListResourcePortsSortOrderAsc,
	"DESC": ListResourcePortsSortOrderDesc,
}

var mappingListResourcePortsSortOrderEnumLowerCase = map[string]ListResourcePortsSortOrderEnum{
	"asc":  ListResourcePortsSortOrderAsc,
	"desc": ListResourcePortsSortOrderDesc,
}

// GetListResourcePortsSortOrderEnumValues Enumerates the set of values for ListResourcePortsSortOrderEnum
func GetListResourcePortsSortOrderEnumValues() []ListResourcePortsSortOrderEnum {
	values := make([]ListResourcePortsSortOrderEnum, 0)
	for _, v := range mappingListResourcePortsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourcePortsSortOrderEnumStringValues Enumerates the set of values in String for ListResourcePortsSortOrderEnum
func GetListResourcePortsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResourcePortsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourcePortsSortOrderEnum(val string) (ListResourcePortsSortOrderEnum, bool) {
	enum, ok := mappingListResourcePortsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourcePortsSortByEnum Enum with underlying type: string
type ListResourcePortsSortByEnum string

// Set of constants representing the allowable values for ListResourcePortsSortByEnum
const (
	ListResourcePortsSortByTimecreated ListResourcePortsSortByEnum = "timeCreated"
	ListResourcePortsSortByDisplayname ListResourcePortsSortByEnum = "displayName"
)

var mappingListResourcePortsSortByEnum = map[string]ListResourcePortsSortByEnum{
	"timeCreated": ListResourcePortsSortByTimecreated,
	"displayName": ListResourcePortsSortByDisplayname,
}

var mappingListResourcePortsSortByEnumLowerCase = map[string]ListResourcePortsSortByEnum{
	"timecreated": ListResourcePortsSortByTimecreated,
	"displayname": ListResourcePortsSortByDisplayname,
}

// GetListResourcePortsSortByEnumValues Enumerates the set of values for ListResourcePortsSortByEnum
func GetListResourcePortsSortByEnumValues() []ListResourcePortsSortByEnum {
	values := make([]ListResourcePortsSortByEnum, 0)
	for _, v := range mappingListResourcePortsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourcePortsSortByEnumStringValues Enumerates the set of values in String for ListResourcePortsSortByEnum
func GetListResourcePortsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListResourcePortsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourcePortsSortByEnum(val string) (ListResourcePortsSortByEnum, bool) {
	enum, ok := mappingListResourcePortsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
