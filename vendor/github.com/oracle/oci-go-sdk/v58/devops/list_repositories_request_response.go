// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListRepositoriesRequest wrapper for the ListRepositories operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListRepositories.go.html to see an example of how to use ListRepositoriesRequest.
type ListRepositoriesRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// unique project identifier
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// Unique repository identifier.
	RepositoryId *string `mandatory:"false" contributesTo:"query" name:"repositoryId"`

	// A filter to return only resources whose lifecycle state matches the given lifecycle state.
	LifecycleState RepositoryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListRepositoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for name is ascending. If no value is specified time created is default.
	SortBy ListRepositoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRepositoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRepositoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRepositoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRepositoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRepositoriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRepositoryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRepositoryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRepositoriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRepositoriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRepositoriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRepositoriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRepositoriesResponse wrapper for the ListRepositories operation
type ListRepositoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RepositoryCollection instances
	RepositoryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRepositoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRepositoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRepositoriesSortOrderEnum Enum with underlying type: string
type ListRepositoriesSortOrderEnum string

// Set of constants representing the allowable values for ListRepositoriesSortOrderEnum
const (
	ListRepositoriesSortOrderAsc  ListRepositoriesSortOrderEnum = "ASC"
	ListRepositoriesSortOrderDesc ListRepositoriesSortOrderEnum = "DESC"
)

var mappingListRepositoriesSortOrderEnum = map[string]ListRepositoriesSortOrderEnum{
	"ASC":  ListRepositoriesSortOrderAsc,
	"DESC": ListRepositoriesSortOrderDesc,
}

// GetListRepositoriesSortOrderEnumValues Enumerates the set of values for ListRepositoriesSortOrderEnum
func GetListRepositoriesSortOrderEnumValues() []ListRepositoriesSortOrderEnum {
	values := make([]ListRepositoriesSortOrderEnum, 0)
	for _, v := range mappingListRepositoriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRepositoriesSortOrderEnumStringValues Enumerates the set of values in String for ListRepositoriesSortOrderEnum
func GetListRepositoriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRepositoriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRepositoriesSortOrderEnum(val string) (ListRepositoriesSortOrderEnum, bool) {
	mappingListRepositoriesSortOrderEnumIgnoreCase := make(map[string]ListRepositoriesSortOrderEnum)
	for k, v := range mappingListRepositoriesSortOrderEnum {
		mappingListRepositoriesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRepositoriesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListRepositoriesSortByEnum Enum with underlying type: string
type ListRepositoriesSortByEnum string

// Set of constants representing the allowable values for ListRepositoriesSortByEnum
const (
	ListRepositoriesSortByTimecreated ListRepositoriesSortByEnum = "timeCreated"
	ListRepositoriesSortByName        ListRepositoriesSortByEnum = "name"
)

var mappingListRepositoriesSortByEnum = map[string]ListRepositoriesSortByEnum{
	"timeCreated": ListRepositoriesSortByTimecreated,
	"name":        ListRepositoriesSortByName,
}

// GetListRepositoriesSortByEnumValues Enumerates the set of values for ListRepositoriesSortByEnum
func GetListRepositoriesSortByEnumValues() []ListRepositoriesSortByEnum {
	values := make([]ListRepositoriesSortByEnum, 0)
	for _, v := range mappingListRepositoriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRepositoriesSortByEnumStringValues Enumerates the set of values in String for ListRepositoriesSortByEnum
func GetListRepositoriesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListRepositoriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRepositoriesSortByEnum(val string) (ListRepositoriesSortByEnum, bool) {
	mappingListRepositoriesSortByEnumIgnoreCase := make(map[string]ListRepositoriesSortByEnum)
	for k, v := range mappingListRepositoriesSortByEnum {
		mappingListRepositoriesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRepositoriesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
