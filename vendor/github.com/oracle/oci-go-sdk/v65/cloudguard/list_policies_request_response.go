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

// ListPoliciesRequest wrapper for the ListPolicies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListPolicies.go.html to see an example of how to use ListPoliciesRequest.
type ListPoliciesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPoliciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPoliciesResponse wrapper for the ListPolicies operation
type ListPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PolicyCollection instances
	PolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPoliciesSortOrderEnum Enum with underlying type: string
type ListPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListPoliciesSortOrderEnum
const (
	ListPoliciesSortOrderAsc  ListPoliciesSortOrderEnum = "ASC"
	ListPoliciesSortOrderDesc ListPoliciesSortOrderEnum = "DESC"
)

var mappingListPoliciesSortOrderEnum = map[string]ListPoliciesSortOrderEnum{
	"ASC":  ListPoliciesSortOrderAsc,
	"DESC": ListPoliciesSortOrderDesc,
}

var mappingListPoliciesSortOrderEnumLowerCase = map[string]ListPoliciesSortOrderEnum{
	"asc":  ListPoliciesSortOrderAsc,
	"desc": ListPoliciesSortOrderDesc,
}

// GetListPoliciesSortOrderEnumValues Enumerates the set of values for ListPoliciesSortOrderEnum
func GetListPoliciesSortOrderEnumValues() []ListPoliciesSortOrderEnum {
	values := make([]ListPoliciesSortOrderEnum, 0)
	for _, v := range mappingListPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListPoliciesSortOrderEnum
func GetListPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPoliciesSortOrderEnum(val string) (ListPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPoliciesSortByEnum Enum with underlying type: string
type ListPoliciesSortByEnum string

// Set of constants representing the allowable values for ListPoliciesSortByEnum
const (
	ListPoliciesSortByTimecreated ListPoliciesSortByEnum = "timeCreated"
	ListPoliciesSortByDisplayname ListPoliciesSortByEnum = "displayName"
)

var mappingListPoliciesSortByEnum = map[string]ListPoliciesSortByEnum{
	"timeCreated": ListPoliciesSortByTimecreated,
	"displayName": ListPoliciesSortByDisplayname,
}

var mappingListPoliciesSortByEnumLowerCase = map[string]ListPoliciesSortByEnum{
	"timecreated": ListPoliciesSortByTimecreated,
	"displayname": ListPoliciesSortByDisplayname,
}

// GetListPoliciesSortByEnumValues Enumerates the set of values for ListPoliciesSortByEnum
func GetListPoliciesSortByEnumValues() []ListPoliciesSortByEnum {
	values := make([]ListPoliciesSortByEnum, 0)
	for _, v := range mappingListPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPoliciesSortByEnumStringValues Enumerates the set of values in String for ListPoliciesSortByEnum
func GetListPoliciesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPoliciesSortByEnum(val string) (ListPoliciesSortByEnum, bool) {
	enum, ok := mappingListPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
