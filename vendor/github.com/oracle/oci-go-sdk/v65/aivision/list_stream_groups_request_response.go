// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package aivision

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListStreamGroupsRequest wrapper for the ListStreamGroups operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aivision/ListStreamGroups.go.html to see an example of how to use ListStreamGroupsRequest.
type ListStreamGroupsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The filter to find the device with the given identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListStreamGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. The default order for timeCreated is descending. The default order for displayName is ascending.
	SortBy ListStreamGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStreamGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStreamGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListStreamGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStreamGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListStreamGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListStreamGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListStreamGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStreamGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListStreamGroupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListStreamGroupsResponse wrapper for the ListStreamGroups operation
type ListStreamGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of StreamGroupCollection instances
	StreamGroupCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListStreamGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStreamGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStreamGroupsSortOrderEnum Enum with underlying type: string
type ListStreamGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListStreamGroupsSortOrderEnum
const (
	ListStreamGroupsSortOrderAsc  ListStreamGroupsSortOrderEnum = "ASC"
	ListStreamGroupsSortOrderDesc ListStreamGroupsSortOrderEnum = "DESC"
)

var mappingListStreamGroupsSortOrderEnum = map[string]ListStreamGroupsSortOrderEnum{
	"ASC":  ListStreamGroupsSortOrderAsc,
	"DESC": ListStreamGroupsSortOrderDesc,
}

var mappingListStreamGroupsSortOrderEnumLowerCase = map[string]ListStreamGroupsSortOrderEnum{
	"asc":  ListStreamGroupsSortOrderAsc,
	"desc": ListStreamGroupsSortOrderDesc,
}

// GetListStreamGroupsSortOrderEnumValues Enumerates the set of values for ListStreamGroupsSortOrderEnum
func GetListStreamGroupsSortOrderEnumValues() []ListStreamGroupsSortOrderEnum {
	values := make([]ListStreamGroupsSortOrderEnum, 0)
	for _, v := range mappingListStreamGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListStreamGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListStreamGroupsSortOrderEnum
func GetListStreamGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListStreamGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStreamGroupsSortOrderEnum(val string) (ListStreamGroupsSortOrderEnum, bool) {
	enum, ok := mappingListStreamGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStreamGroupsSortByEnum Enum with underlying type: string
type ListStreamGroupsSortByEnum string

// Set of constants representing the allowable values for ListStreamGroupsSortByEnum
const (
	ListStreamGroupsSortByTimecreated ListStreamGroupsSortByEnum = "timeCreated"
	ListStreamGroupsSortByDisplayname ListStreamGroupsSortByEnum = "displayName"
)

var mappingListStreamGroupsSortByEnum = map[string]ListStreamGroupsSortByEnum{
	"timeCreated": ListStreamGroupsSortByTimecreated,
	"displayName": ListStreamGroupsSortByDisplayname,
}

var mappingListStreamGroupsSortByEnumLowerCase = map[string]ListStreamGroupsSortByEnum{
	"timecreated": ListStreamGroupsSortByTimecreated,
	"displayname": ListStreamGroupsSortByDisplayname,
}

// GetListStreamGroupsSortByEnumValues Enumerates the set of values for ListStreamGroupsSortByEnum
func GetListStreamGroupsSortByEnumValues() []ListStreamGroupsSortByEnum {
	values := make([]ListStreamGroupsSortByEnum, 0)
	for _, v := range mappingListStreamGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListStreamGroupsSortByEnumStringValues Enumerates the set of values in String for ListStreamGroupsSortByEnum
func GetListStreamGroupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListStreamGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStreamGroupsSortByEnum(val string) (ListStreamGroupsSortByEnum, bool) {
	enum, ok := mappingListStreamGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
