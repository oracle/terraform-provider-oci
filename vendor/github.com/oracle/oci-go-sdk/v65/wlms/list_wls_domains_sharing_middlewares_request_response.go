// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWlsDomainsSharingMiddlewaresRequest wrapper for the ListWlsDomainsSharingMiddlewares operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWlsDomainsSharingMiddlewares.go.html to see an example of how to use ListWlsDomainsSharingMiddlewaresRequest.
type ListWlsDomainsSharingMiddlewaresRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.
	WlsDomainId *string `mandatory:"true" contributesTo:"path" name:"wlsDomainId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token that represents the page at which to start retrieving results. The token is usually retrieved from a previous List call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order is either 'ASC' or 'DESC'.
	SortOrder ListWlsDomainsSharingMiddlewaresSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort the resource. Only one sort order may be provided.
	// Default order for _timeCreated_ is **descending**.
	// Default order for _displayName_ is **ascending**.
	// If no value is specified, _timeCreated_ is default.
	SortBy ListWlsDomainsSharingMiddlewaresSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWlsDomainsSharingMiddlewaresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWlsDomainsSharingMiddlewaresRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWlsDomainsSharingMiddlewaresRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWlsDomainsSharingMiddlewaresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWlsDomainsSharingMiddlewaresRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWlsDomainsSharingMiddlewaresSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWlsDomainsSharingMiddlewaresSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlsDomainsSharingMiddlewaresSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWlsDomainsSharingMiddlewaresSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWlsDomainsSharingMiddlewaresResponse wrapper for the ListWlsDomainsSharingMiddlewares operation
type ListWlsDomainsSharingMiddlewaresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WlsDomainCollection instances
	WlsDomainCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWlsDomainsSharingMiddlewaresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWlsDomainsSharingMiddlewaresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWlsDomainsSharingMiddlewaresSortOrderEnum Enum with underlying type: string
type ListWlsDomainsSharingMiddlewaresSortOrderEnum string

// Set of constants representing the allowable values for ListWlsDomainsSharingMiddlewaresSortOrderEnum
const (
	ListWlsDomainsSharingMiddlewaresSortOrderAsc  ListWlsDomainsSharingMiddlewaresSortOrderEnum = "ASC"
	ListWlsDomainsSharingMiddlewaresSortOrderDesc ListWlsDomainsSharingMiddlewaresSortOrderEnum = "DESC"
)

var mappingListWlsDomainsSharingMiddlewaresSortOrderEnum = map[string]ListWlsDomainsSharingMiddlewaresSortOrderEnum{
	"ASC":  ListWlsDomainsSharingMiddlewaresSortOrderAsc,
	"DESC": ListWlsDomainsSharingMiddlewaresSortOrderDesc,
}

var mappingListWlsDomainsSharingMiddlewaresSortOrderEnumLowerCase = map[string]ListWlsDomainsSharingMiddlewaresSortOrderEnum{
	"asc":  ListWlsDomainsSharingMiddlewaresSortOrderAsc,
	"desc": ListWlsDomainsSharingMiddlewaresSortOrderDesc,
}

// GetListWlsDomainsSharingMiddlewaresSortOrderEnumValues Enumerates the set of values for ListWlsDomainsSharingMiddlewaresSortOrderEnum
func GetListWlsDomainsSharingMiddlewaresSortOrderEnumValues() []ListWlsDomainsSharingMiddlewaresSortOrderEnum {
	values := make([]ListWlsDomainsSharingMiddlewaresSortOrderEnum, 0)
	for _, v := range mappingListWlsDomainsSharingMiddlewaresSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlsDomainsSharingMiddlewaresSortOrderEnumStringValues Enumerates the set of values in String for ListWlsDomainsSharingMiddlewaresSortOrderEnum
func GetListWlsDomainsSharingMiddlewaresSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWlsDomainsSharingMiddlewaresSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlsDomainsSharingMiddlewaresSortOrderEnum(val string) (ListWlsDomainsSharingMiddlewaresSortOrderEnum, bool) {
	enum, ok := mappingListWlsDomainsSharingMiddlewaresSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlsDomainsSharingMiddlewaresSortByEnum Enum with underlying type: string
type ListWlsDomainsSharingMiddlewaresSortByEnum string

// Set of constants representing the allowable values for ListWlsDomainsSharingMiddlewaresSortByEnum
const (
	ListWlsDomainsSharingMiddlewaresSortByTimecreated ListWlsDomainsSharingMiddlewaresSortByEnum = "timeCreated"
	ListWlsDomainsSharingMiddlewaresSortByDisplayname ListWlsDomainsSharingMiddlewaresSortByEnum = "displayName"
)

var mappingListWlsDomainsSharingMiddlewaresSortByEnum = map[string]ListWlsDomainsSharingMiddlewaresSortByEnum{
	"timeCreated": ListWlsDomainsSharingMiddlewaresSortByTimecreated,
	"displayName": ListWlsDomainsSharingMiddlewaresSortByDisplayname,
}

var mappingListWlsDomainsSharingMiddlewaresSortByEnumLowerCase = map[string]ListWlsDomainsSharingMiddlewaresSortByEnum{
	"timecreated": ListWlsDomainsSharingMiddlewaresSortByTimecreated,
	"displayname": ListWlsDomainsSharingMiddlewaresSortByDisplayname,
}

// GetListWlsDomainsSharingMiddlewaresSortByEnumValues Enumerates the set of values for ListWlsDomainsSharingMiddlewaresSortByEnum
func GetListWlsDomainsSharingMiddlewaresSortByEnumValues() []ListWlsDomainsSharingMiddlewaresSortByEnum {
	values := make([]ListWlsDomainsSharingMiddlewaresSortByEnum, 0)
	for _, v := range mappingListWlsDomainsSharingMiddlewaresSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlsDomainsSharingMiddlewaresSortByEnumStringValues Enumerates the set of values in String for ListWlsDomainsSharingMiddlewaresSortByEnum
func GetListWlsDomainsSharingMiddlewaresSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListWlsDomainsSharingMiddlewaresSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlsDomainsSharingMiddlewaresSortByEnum(val string) (ListWlsDomainsSharingMiddlewaresSortByEnum, bool) {
	enum, ok := mappingListWlsDomainsSharingMiddlewaresSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
