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

// ListAdhocQueryResultsRequest wrapper for the ListAdhocQueryResults operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListAdhocQueryResults.go.html to see an example of how to use ListAdhocQueryResultsRequest.
type ListAdhocQueryResultsRequest struct {

	// Adhoc query OCID.
	AdhocQueryId *string `mandatory:"true" contributesTo:"path" name:"adhocQueryId"`

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListAdhocQueryResultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListAdhocQueryResultsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAdhocQueryResultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAdhocQueryResultsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAdhocQueryResultsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAdhocQueryResultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAdhocQueryResultsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAdhocQueryResultsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAdhocQueryResultsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAdhocQueryResultsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAdhocQueryResultsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAdhocQueryResultsResponse wrapper for the ListAdhocQueryResults operation
type ListAdhocQueryResultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AdhocQueryResultCollection instances
	AdhocQueryResultCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAdhocQueryResultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAdhocQueryResultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAdhocQueryResultsSortOrderEnum Enum with underlying type: string
type ListAdhocQueryResultsSortOrderEnum string

// Set of constants representing the allowable values for ListAdhocQueryResultsSortOrderEnum
const (
	ListAdhocQueryResultsSortOrderAsc  ListAdhocQueryResultsSortOrderEnum = "ASC"
	ListAdhocQueryResultsSortOrderDesc ListAdhocQueryResultsSortOrderEnum = "DESC"
)

var mappingListAdhocQueryResultsSortOrderEnum = map[string]ListAdhocQueryResultsSortOrderEnum{
	"ASC":  ListAdhocQueryResultsSortOrderAsc,
	"DESC": ListAdhocQueryResultsSortOrderDesc,
}

var mappingListAdhocQueryResultsSortOrderEnumLowerCase = map[string]ListAdhocQueryResultsSortOrderEnum{
	"asc":  ListAdhocQueryResultsSortOrderAsc,
	"desc": ListAdhocQueryResultsSortOrderDesc,
}

// GetListAdhocQueryResultsSortOrderEnumValues Enumerates the set of values for ListAdhocQueryResultsSortOrderEnum
func GetListAdhocQueryResultsSortOrderEnumValues() []ListAdhocQueryResultsSortOrderEnum {
	values := make([]ListAdhocQueryResultsSortOrderEnum, 0)
	for _, v := range mappingListAdhocQueryResultsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdhocQueryResultsSortOrderEnumStringValues Enumerates the set of values in String for ListAdhocQueryResultsSortOrderEnum
func GetListAdhocQueryResultsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAdhocQueryResultsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdhocQueryResultsSortOrderEnum(val string) (ListAdhocQueryResultsSortOrderEnum, bool) {
	enum, ok := mappingListAdhocQueryResultsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAdhocQueryResultsSortByEnum Enum with underlying type: string
type ListAdhocQueryResultsSortByEnum string

// Set of constants representing the allowable values for ListAdhocQueryResultsSortByEnum
const (
	ListAdhocQueryResultsSortByTimecreated ListAdhocQueryResultsSortByEnum = "timeCreated"
	ListAdhocQueryResultsSortByDisplayname ListAdhocQueryResultsSortByEnum = "displayName"
)

var mappingListAdhocQueryResultsSortByEnum = map[string]ListAdhocQueryResultsSortByEnum{
	"timeCreated": ListAdhocQueryResultsSortByTimecreated,
	"displayName": ListAdhocQueryResultsSortByDisplayname,
}

var mappingListAdhocQueryResultsSortByEnumLowerCase = map[string]ListAdhocQueryResultsSortByEnum{
	"timecreated": ListAdhocQueryResultsSortByTimecreated,
	"displayname": ListAdhocQueryResultsSortByDisplayname,
}

// GetListAdhocQueryResultsSortByEnumValues Enumerates the set of values for ListAdhocQueryResultsSortByEnum
func GetListAdhocQueryResultsSortByEnumValues() []ListAdhocQueryResultsSortByEnum {
	values := make([]ListAdhocQueryResultsSortByEnum, 0)
	for _, v := range mappingListAdhocQueryResultsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdhocQueryResultsSortByEnumStringValues Enumerates the set of values in String for ListAdhocQueryResultsSortByEnum
func GetListAdhocQueryResultsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAdhocQueryResultsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdhocQueryResultsSortByEnum(val string) (ListAdhocQueryResultsSortByEnum, bool) {
	enum, ok := mappingListAdhocQueryResultsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
