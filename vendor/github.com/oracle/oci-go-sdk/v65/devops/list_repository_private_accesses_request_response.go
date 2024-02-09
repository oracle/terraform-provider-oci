// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRepositoryPrivateAccessesRequest wrapper for the ListRepositoryPrivateAccesses operation
type ListRepositoryPrivateAccessesRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique identifier or OCID for listing a single resource by ID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return resources that match a given lifecycle state.
	LifecycleState RepositoryPrivateAccessLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListRepositoryPrivateAccessesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for display name is ascending. If no value is specified, then the default time created value is considered.
	SortBy ListRepositoryPrivateAccessesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRepositoryPrivateAccessesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRepositoryPrivateAccessesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRepositoryPrivateAccessesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRepositoryPrivateAccessesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRepositoryPrivateAccessesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRepositoryPrivateAccessLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRepositoryPrivateAccessLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRepositoryPrivateAccessesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRepositoryPrivateAccessesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRepositoryPrivateAccessesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRepositoryPrivateAccessesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRepositoryPrivateAccessesResponse wrapper for the ListRepositoryPrivateAccesses operation
type ListRepositoryPrivateAccessesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RepositoryPrivateAccessCollection instances
	RepositoryPrivateAccessCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRepositoryPrivateAccessesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRepositoryPrivateAccessesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRepositoryPrivateAccessesSortOrderEnum Enum with underlying type: string
type ListRepositoryPrivateAccessesSortOrderEnum string

// Set of constants representing the allowable values for ListRepositoryPrivateAccessesSortOrderEnum
const (
	ListRepositoryPrivateAccessesSortOrderAsc  ListRepositoryPrivateAccessesSortOrderEnum = "ASC"
	ListRepositoryPrivateAccessesSortOrderDesc ListRepositoryPrivateAccessesSortOrderEnum = "DESC"
)

var mappingListRepositoryPrivateAccessesSortOrderEnum = map[string]ListRepositoryPrivateAccessesSortOrderEnum{
	"ASC":  ListRepositoryPrivateAccessesSortOrderAsc,
	"DESC": ListRepositoryPrivateAccessesSortOrderDesc,
}

var mappingListRepositoryPrivateAccessesSortOrderEnumLowerCase = map[string]ListRepositoryPrivateAccessesSortOrderEnum{
	"asc":  ListRepositoryPrivateAccessesSortOrderAsc,
	"desc": ListRepositoryPrivateAccessesSortOrderDesc,
}

// GetListRepositoryPrivateAccessesSortOrderEnumValues Enumerates the set of values for ListRepositoryPrivateAccessesSortOrderEnum
func GetListRepositoryPrivateAccessesSortOrderEnumValues() []ListRepositoryPrivateAccessesSortOrderEnum {
	values := make([]ListRepositoryPrivateAccessesSortOrderEnum, 0)
	for _, v := range mappingListRepositoryPrivateAccessesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRepositoryPrivateAccessesSortOrderEnumStringValues Enumerates the set of values in String for ListRepositoryPrivateAccessesSortOrderEnum
func GetListRepositoryPrivateAccessesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRepositoryPrivateAccessesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRepositoryPrivateAccessesSortOrderEnum(val string) (ListRepositoryPrivateAccessesSortOrderEnum, bool) {
	enum, ok := mappingListRepositoryPrivateAccessesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRepositoryPrivateAccessesSortByEnum Enum with underlying type: string
type ListRepositoryPrivateAccessesSortByEnum string

// Set of constants representing the allowable values for ListRepositoryPrivateAccessesSortByEnum
const (
	ListRepositoryPrivateAccessesSortByTimecreated ListRepositoryPrivateAccessesSortByEnum = "timeCreated"
	ListRepositoryPrivateAccessesSortByDisplayname ListRepositoryPrivateAccessesSortByEnum = "displayName"
)

var mappingListRepositoryPrivateAccessesSortByEnum = map[string]ListRepositoryPrivateAccessesSortByEnum{
	"timeCreated": ListRepositoryPrivateAccessesSortByTimecreated,
	"displayName": ListRepositoryPrivateAccessesSortByDisplayname,
}

var mappingListRepositoryPrivateAccessesSortByEnumLowerCase = map[string]ListRepositoryPrivateAccessesSortByEnum{
	"timecreated": ListRepositoryPrivateAccessesSortByTimecreated,
	"displayname": ListRepositoryPrivateAccessesSortByDisplayname,
}

// GetListRepositoryPrivateAccessesSortByEnumValues Enumerates the set of values for ListRepositoryPrivateAccessesSortByEnum
func GetListRepositoryPrivateAccessesSortByEnumValues() []ListRepositoryPrivateAccessesSortByEnum {
	values := make([]ListRepositoryPrivateAccessesSortByEnum, 0)
	for _, v := range mappingListRepositoryPrivateAccessesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRepositoryPrivateAccessesSortByEnumStringValues Enumerates the set of values in String for ListRepositoryPrivateAccessesSortByEnum
func GetListRepositoryPrivateAccessesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListRepositoryPrivateAccessesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRepositoryPrivateAccessesSortByEnum(val string) (ListRepositoryPrivateAccessesSortByEnum, bool) {
	enum, ok := mappingListRepositoryPrivateAccessesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
