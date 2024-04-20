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

// ListWlpAgentsRequest wrapper for the ListWlpAgents operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListWlpAgents.go.html to see an example of how to use ListWlpAgentsRequest.
type ListWlpAgentsRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListWlpAgentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListWlpAgentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWlpAgentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWlpAgentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWlpAgentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWlpAgentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWlpAgentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWlpAgentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWlpAgentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlpAgentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWlpAgentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWlpAgentsResponse wrapper for the ListWlpAgents operation
type ListWlpAgentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WlpAgentCollection instances
	WlpAgentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWlpAgentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWlpAgentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWlpAgentsSortOrderEnum Enum with underlying type: string
type ListWlpAgentsSortOrderEnum string

// Set of constants representing the allowable values for ListWlpAgentsSortOrderEnum
const (
	ListWlpAgentsSortOrderAsc  ListWlpAgentsSortOrderEnum = "ASC"
	ListWlpAgentsSortOrderDesc ListWlpAgentsSortOrderEnum = "DESC"
)

var mappingListWlpAgentsSortOrderEnum = map[string]ListWlpAgentsSortOrderEnum{
	"ASC":  ListWlpAgentsSortOrderAsc,
	"DESC": ListWlpAgentsSortOrderDesc,
}

var mappingListWlpAgentsSortOrderEnumLowerCase = map[string]ListWlpAgentsSortOrderEnum{
	"asc":  ListWlpAgentsSortOrderAsc,
	"desc": ListWlpAgentsSortOrderDesc,
}

// GetListWlpAgentsSortOrderEnumValues Enumerates the set of values for ListWlpAgentsSortOrderEnum
func GetListWlpAgentsSortOrderEnumValues() []ListWlpAgentsSortOrderEnum {
	values := make([]ListWlpAgentsSortOrderEnum, 0)
	for _, v := range mappingListWlpAgentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlpAgentsSortOrderEnumStringValues Enumerates the set of values in String for ListWlpAgentsSortOrderEnum
func GetListWlpAgentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWlpAgentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlpAgentsSortOrderEnum(val string) (ListWlpAgentsSortOrderEnum, bool) {
	enum, ok := mappingListWlpAgentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlpAgentsSortByEnum Enum with underlying type: string
type ListWlpAgentsSortByEnum string

// Set of constants representing the allowable values for ListWlpAgentsSortByEnum
const (
	ListWlpAgentsSortByTimecreated ListWlpAgentsSortByEnum = "timeCreated"
	ListWlpAgentsSortByDisplayname ListWlpAgentsSortByEnum = "displayName"
)

var mappingListWlpAgentsSortByEnum = map[string]ListWlpAgentsSortByEnum{
	"timeCreated": ListWlpAgentsSortByTimecreated,
	"displayName": ListWlpAgentsSortByDisplayname,
}

var mappingListWlpAgentsSortByEnumLowerCase = map[string]ListWlpAgentsSortByEnum{
	"timecreated": ListWlpAgentsSortByTimecreated,
	"displayname": ListWlpAgentsSortByDisplayname,
}

// GetListWlpAgentsSortByEnumValues Enumerates the set of values for ListWlpAgentsSortByEnum
func GetListWlpAgentsSortByEnumValues() []ListWlpAgentsSortByEnum {
	values := make([]ListWlpAgentsSortByEnum, 0)
	for _, v := range mappingListWlpAgentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlpAgentsSortByEnumStringValues Enumerates the set of values in String for ListWlpAgentsSortByEnum
func GetListWlpAgentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListWlpAgentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlpAgentsSortByEnum(val string) (ListWlpAgentsSortByEnum, bool) {
	enum, ok := mappingListWlpAgentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
