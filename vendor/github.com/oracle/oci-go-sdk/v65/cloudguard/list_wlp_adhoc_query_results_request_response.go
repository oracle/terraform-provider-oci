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

// ListWlpAdhocQueryResultsRequest wrapper for the ListWlpAdhocQueryResults operation
type ListWlpAdhocQueryResultsRequest struct {

	// Adhoc query OCID.
	WlpAdhocQueryId *string `mandatory:"true" contributesTo:"path" name:"wlpAdhocQueryId"`

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListWlpAdhocQueryResultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListWlpAdhocQueryResultsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWlpAdhocQueryResultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWlpAdhocQueryResultsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWlpAdhocQueryResultsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWlpAdhocQueryResultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWlpAdhocQueryResultsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWlpAdhocQueryResultsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWlpAdhocQueryResultsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlpAdhocQueryResultsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWlpAdhocQueryResultsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWlpAdhocQueryResultsResponse wrapper for the ListWlpAdhocQueryResults operation
type ListWlpAdhocQueryResultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WlpAdhocQueryResultCollection instances
	WlpAdhocQueryResultCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWlpAdhocQueryResultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWlpAdhocQueryResultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWlpAdhocQueryResultsSortOrderEnum Enum with underlying type: string
type ListWlpAdhocQueryResultsSortOrderEnum string

// Set of constants representing the allowable values for ListWlpAdhocQueryResultsSortOrderEnum
const (
	ListWlpAdhocQueryResultsSortOrderAsc  ListWlpAdhocQueryResultsSortOrderEnum = "ASC"
	ListWlpAdhocQueryResultsSortOrderDesc ListWlpAdhocQueryResultsSortOrderEnum = "DESC"
)

var mappingListWlpAdhocQueryResultsSortOrderEnum = map[string]ListWlpAdhocQueryResultsSortOrderEnum{
	"ASC":  ListWlpAdhocQueryResultsSortOrderAsc,
	"DESC": ListWlpAdhocQueryResultsSortOrderDesc,
}

var mappingListWlpAdhocQueryResultsSortOrderEnumLowerCase = map[string]ListWlpAdhocQueryResultsSortOrderEnum{
	"asc":  ListWlpAdhocQueryResultsSortOrderAsc,
	"desc": ListWlpAdhocQueryResultsSortOrderDesc,
}

// GetListWlpAdhocQueryResultsSortOrderEnumValues Enumerates the set of values for ListWlpAdhocQueryResultsSortOrderEnum
func GetListWlpAdhocQueryResultsSortOrderEnumValues() []ListWlpAdhocQueryResultsSortOrderEnum {
	values := make([]ListWlpAdhocQueryResultsSortOrderEnum, 0)
	for _, v := range mappingListWlpAdhocQueryResultsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlpAdhocQueryResultsSortOrderEnumStringValues Enumerates the set of values in String for ListWlpAdhocQueryResultsSortOrderEnum
func GetListWlpAdhocQueryResultsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWlpAdhocQueryResultsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlpAdhocQueryResultsSortOrderEnum(val string) (ListWlpAdhocQueryResultsSortOrderEnum, bool) {
	enum, ok := mappingListWlpAdhocQueryResultsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlpAdhocQueryResultsSortByEnum Enum with underlying type: string
type ListWlpAdhocQueryResultsSortByEnum string

// Set of constants representing the allowable values for ListWlpAdhocQueryResultsSortByEnum
const (
	ListWlpAdhocQueryResultsSortByTimecreated ListWlpAdhocQueryResultsSortByEnum = "timeCreated"
	ListWlpAdhocQueryResultsSortByDisplayname ListWlpAdhocQueryResultsSortByEnum = "displayName"
)

var mappingListWlpAdhocQueryResultsSortByEnum = map[string]ListWlpAdhocQueryResultsSortByEnum{
	"timeCreated": ListWlpAdhocQueryResultsSortByTimecreated,
	"displayName": ListWlpAdhocQueryResultsSortByDisplayname,
}

var mappingListWlpAdhocQueryResultsSortByEnumLowerCase = map[string]ListWlpAdhocQueryResultsSortByEnum{
	"timecreated": ListWlpAdhocQueryResultsSortByTimecreated,
	"displayname": ListWlpAdhocQueryResultsSortByDisplayname,
}

// GetListWlpAdhocQueryResultsSortByEnumValues Enumerates the set of values for ListWlpAdhocQueryResultsSortByEnum
func GetListWlpAdhocQueryResultsSortByEnumValues() []ListWlpAdhocQueryResultsSortByEnum {
	values := make([]ListWlpAdhocQueryResultsSortByEnum, 0)
	for _, v := range mappingListWlpAdhocQueryResultsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlpAdhocQueryResultsSortByEnumStringValues Enumerates the set of values in String for ListWlpAdhocQueryResultsSortByEnum
func GetListWlpAdhocQueryResultsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListWlpAdhocQueryResultsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlpAdhocQueryResultsSortByEnum(val string) (ListWlpAdhocQueryResultsSortByEnum, bool) {
	enum, ok := mappingListWlpAdhocQueryResultsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
