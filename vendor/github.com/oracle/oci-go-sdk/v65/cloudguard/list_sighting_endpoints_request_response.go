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

// ListSightingEndpointsRequest wrapper for the ListSightingEndpoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListSightingEndpoints.go.html to see an example of how to use ListSightingEndpointsRequest.
type ListSightingEndpointsRequest struct {

	// OCID of the sighting.
	SightingId *string `mandatory:"true" contributesTo:"path" name:"sightingId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListSightingEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListSightingEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSightingEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSightingEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSightingEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSightingEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSightingEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSightingEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSightingEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSightingEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSightingEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSightingEndpointsResponse wrapper for the ListSightingEndpoints operation
type ListSightingEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SightingEndpointCollection instances
	SightingEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSightingEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSightingEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSightingEndpointsSortOrderEnum Enum with underlying type: string
type ListSightingEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListSightingEndpointsSortOrderEnum
const (
	ListSightingEndpointsSortOrderAsc  ListSightingEndpointsSortOrderEnum = "ASC"
	ListSightingEndpointsSortOrderDesc ListSightingEndpointsSortOrderEnum = "DESC"
)

var mappingListSightingEndpointsSortOrderEnum = map[string]ListSightingEndpointsSortOrderEnum{
	"ASC":  ListSightingEndpointsSortOrderAsc,
	"DESC": ListSightingEndpointsSortOrderDesc,
}

var mappingListSightingEndpointsSortOrderEnumLowerCase = map[string]ListSightingEndpointsSortOrderEnum{
	"asc":  ListSightingEndpointsSortOrderAsc,
	"desc": ListSightingEndpointsSortOrderDesc,
}

// GetListSightingEndpointsSortOrderEnumValues Enumerates the set of values for ListSightingEndpointsSortOrderEnum
func GetListSightingEndpointsSortOrderEnumValues() []ListSightingEndpointsSortOrderEnum {
	values := make([]ListSightingEndpointsSortOrderEnum, 0)
	for _, v := range mappingListSightingEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSightingEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListSightingEndpointsSortOrderEnum
func GetListSightingEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSightingEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSightingEndpointsSortOrderEnum(val string) (ListSightingEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListSightingEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSightingEndpointsSortByEnum Enum with underlying type: string
type ListSightingEndpointsSortByEnum string

// Set of constants representing the allowable values for ListSightingEndpointsSortByEnum
const (
	ListSightingEndpointsSortByTimecreated ListSightingEndpointsSortByEnum = "timeCreated"
)

var mappingListSightingEndpointsSortByEnum = map[string]ListSightingEndpointsSortByEnum{
	"timeCreated": ListSightingEndpointsSortByTimecreated,
}

var mappingListSightingEndpointsSortByEnumLowerCase = map[string]ListSightingEndpointsSortByEnum{
	"timecreated": ListSightingEndpointsSortByTimecreated,
}

// GetListSightingEndpointsSortByEnumValues Enumerates the set of values for ListSightingEndpointsSortByEnum
func GetListSightingEndpointsSortByEnumValues() []ListSightingEndpointsSortByEnum {
	values := make([]ListSightingEndpointsSortByEnum, 0)
	for _, v := range mappingListSightingEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSightingEndpointsSortByEnumStringValues Enumerates the set of values in String for ListSightingEndpointsSortByEnum
func GetListSightingEndpointsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListSightingEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSightingEndpointsSortByEnum(val string) (ListSightingEndpointsSortByEnum, bool) {
	enum, ok := mappingListSightingEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
