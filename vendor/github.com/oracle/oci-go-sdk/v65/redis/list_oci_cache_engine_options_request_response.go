// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOciCacheEngineOptionsRequest wrapper for the ListOciCacheEngineOptions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListOciCacheEngineOptions.go.html to see an example of how to use ListOciCacheEngineOptionsRequest.
type ListOciCacheEngineOptionsRequest struct {

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOciCacheEngineOptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListOciCacheEngineOptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOciCacheEngineOptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOciCacheEngineOptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOciCacheEngineOptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOciCacheEngineOptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOciCacheEngineOptionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOciCacheEngineOptionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOciCacheEngineOptionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOciCacheEngineOptionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOciCacheEngineOptionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOciCacheEngineOptionsResponse wrapper for the ListOciCacheEngineOptions operation
type ListOciCacheEngineOptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OciCacheEngineOptionsCollection instances
	OciCacheEngineOptionsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOciCacheEngineOptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOciCacheEngineOptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOciCacheEngineOptionsSortOrderEnum Enum with underlying type: string
type ListOciCacheEngineOptionsSortOrderEnum string

// Set of constants representing the allowable values for ListOciCacheEngineOptionsSortOrderEnum
const (
	ListOciCacheEngineOptionsSortOrderAsc  ListOciCacheEngineOptionsSortOrderEnum = "ASC"
	ListOciCacheEngineOptionsSortOrderDesc ListOciCacheEngineOptionsSortOrderEnum = "DESC"
)

var mappingListOciCacheEngineOptionsSortOrderEnum = map[string]ListOciCacheEngineOptionsSortOrderEnum{
	"ASC":  ListOciCacheEngineOptionsSortOrderAsc,
	"DESC": ListOciCacheEngineOptionsSortOrderDesc,
}

var mappingListOciCacheEngineOptionsSortOrderEnumLowerCase = map[string]ListOciCacheEngineOptionsSortOrderEnum{
	"asc":  ListOciCacheEngineOptionsSortOrderAsc,
	"desc": ListOciCacheEngineOptionsSortOrderDesc,
}

// GetListOciCacheEngineOptionsSortOrderEnumValues Enumerates the set of values for ListOciCacheEngineOptionsSortOrderEnum
func GetListOciCacheEngineOptionsSortOrderEnumValues() []ListOciCacheEngineOptionsSortOrderEnum {
	values := make([]ListOciCacheEngineOptionsSortOrderEnum, 0)
	for _, v := range mappingListOciCacheEngineOptionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOciCacheEngineOptionsSortOrderEnumStringValues Enumerates the set of values in String for ListOciCacheEngineOptionsSortOrderEnum
func GetListOciCacheEngineOptionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOciCacheEngineOptionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOciCacheEngineOptionsSortOrderEnum(val string) (ListOciCacheEngineOptionsSortOrderEnum, bool) {
	enum, ok := mappingListOciCacheEngineOptionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOciCacheEngineOptionsSortByEnum Enum with underlying type: string
type ListOciCacheEngineOptionsSortByEnum string

// Set of constants representing the allowable values for ListOciCacheEngineOptionsSortByEnum
const (
	ListOciCacheEngineOptionsSortByTimecreated ListOciCacheEngineOptionsSortByEnum = "timeCreated"
	ListOciCacheEngineOptionsSortByDisplayname ListOciCacheEngineOptionsSortByEnum = "displayName"
)

var mappingListOciCacheEngineOptionsSortByEnum = map[string]ListOciCacheEngineOptionsSortByEnum{
	"timeCreated": ListOciCacheEngineOptionsSortByTimecreated,
	"displayName": ListOciCacheEngineOptionsSortByDisplayname,
}

var mappingListOciCacheEngineOptionsSortByEnumLowerCase = map[string]ListOciCacheEngineOptionsSortByEnum{
	"timecreated": ListOciCacheEngineOptionsSortByTimecreated,
	"displayname": ListOciCacheEngineOptionsSortByDisplayname,
}

// GetListOciCacheEngineOptionsSortByEnumValues Enumerates the set of values for ListOciCacheEngineOptionsSortByEnum
func GetListOciCacheEngineOptionsSortByEnumValues() []ListOciCacheEngineOptionsSortByEnum {
	values := make([]ListOciCacheEngineOptionsSortByEnum, 0)
	for _, v := range mappingListOciCacheEngineOptionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOciCacheEngineOptionsSortByEnumStringValues Enumerates the set of values in String for ListOciCacheEngineOptionsSortByEnum
func GetListOciCacheEngineOptionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOciCacheEngineOptionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOciCacheEngineOptionsSortByEnum(val string) (ListOciCacheEngineOptionsSortByEnum, bool) {
	enum, ok := mappingListOciCacheEngineOptionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
