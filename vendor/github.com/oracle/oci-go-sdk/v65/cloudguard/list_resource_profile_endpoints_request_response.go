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

// ListResourceProfileEndpointsRequest wrapper for the ListResourceProfileEndpoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResourceProfileEndpoints.go.html to see an example of how to use ListResourceProfileEndpointsRequest.
type ListResourceProfileEndpointsRequest struct {

	// OCID of the resource profile.
	ResourceProfileId *string `mandatory:"true" contributesTo:"path" name:"resourceProfileId"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListResourceProfileEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListResourceProfileEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourceProfileEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourceProfileEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourceProfileEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourceProfileEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResourceProfileEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResourceProfileEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResourceProfileEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceProfileEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResourceProfileEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResourceProfileEndpointsResponse wrapper for the ListResourceProfileEndpoints operation
type ListResourceProfileEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourceProfileEndpointCollection instances
	ResourceProfileEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResourceProfileEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourceProfileEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourceProfileEndpointsSortOrderEnum Enum with underlying type: string
type ListResourceProfileEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListResourceProfileEndpointsSortOrderEnum
const (
	ListResourceProfileEndpointsSortOrderAsc  ListResourceProfileEndpointsSortOrderEnum = "ASC"
	ListResourceProfileEndpointsSortOrderDesc ListResourceProfileEndpointsSortOrderEnum = "DESC"
)

var mappingListResourceProfileEndpointsSortOrderEnum = map[string]ListResourceProfileEndpointsSortOrderEnum{
	"ASC":  ListResourceProfileEndpointsSortOrderAsc,
	"DESC": ListResourceProfileEndpointsSortOrderDesc,
}

var mappingListResourceProfileEndpointsSortOrderEnumLowerCase = map[string]ListResourceProfileEndpointsSortOrderEnum{
	"asc":  ListResourceProfileEndpointsSortOrderAsc,
	"desc": ListResourceProfileEndpointsSortOrderDesc,
}

// GetListResourceProfileEndpointsSortOrderEnumValues Enumerates the set of values for ListResourceProfileEndpointsSortOrderEnum
func GetListResourceProfileEndpointsSortOrderEnumValues() []ListResourceProfileEndpointsSortOrderEnum {
	values := make([]ListResourceProfileEndpointsSortOrderEnum, 0)
	for _, v := range mappingListResourceProfileEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceProfileEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListResourceProfileEndpointsSortOrderEnum
func GetListResourceProfileEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResourceProfileEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceProfileEndpointsSortOrderEnum(val string) (ListResourceProfileEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListResourceProfileEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceProfileEndpointsSortByEnum Enum with underlying type: string
type ListResourceProfileEndpointsSortByEnum string

// Set of constants representing the allowable values for ListResourceProfileEndpointsSortByEnum
const (
	ListResourceProfileEndpointsSortByTimecreated ListResourceProfileEndpointsSortByEnum = "timeCreated"
)

var mappingListResourceProfileEndpointsSortByEnum = map[string]ListResourceProfileEndpointsSortByEnum{
	"timeCreated": ListResourceProfileEndpointsSortByTimecreated,
}

var mappingListResourceProfileEndpointsSortByEnumLowerCase = map[string]ListResourceProfileEndpointsSortByEnum{
	"timecreated": ListResourceProfileEndpointsSortByTimecreated,
}

// GetListResourceProfileEndpointsSortByEnumValues Enumerates the set of values for ListResourceProfileEndpointsSortByEnum
func GetListResourceProfileEndpointsSortByEnumValues() []ListResourceProfileEndpointsSortByEnum {
	values := make([]ListResourceProfileEndpointsSortByEnum, 0)
	for _, v := range mappingListResourceProfileEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceProfileEndpointsSortByEnumStringValues Enumerates the set of values in String for ListResourceProfileEndpointsSortByEnum
func GetListResourceProfileEndpointsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListResourceProfileEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceProfileEndpointsSortByEnum(val string) (ListResourceProfileEndpointsSortByEnum, bool) {
	enum, ok := mappingListResourceProfileEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
