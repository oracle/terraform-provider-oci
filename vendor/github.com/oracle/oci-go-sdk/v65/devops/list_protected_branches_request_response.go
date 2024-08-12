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

// ListProtectedBranchesRequest wrapper for the ListProtectedBranches operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListProtectedBranches.go.html to see an example of how to use ListProtectedBranchesRequest.
type ListProtectedBranchesRequest struct {

	// Unique repository identifier.
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// A filter to return only resources that match the given branch name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListProtectedBranchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for branch name is ascending. If no value is specified branch name is default.
	SortBy ListProtectedBranchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProtectedBranchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProtectedBranchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProtectedBranchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProtectedBranchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProtectedBranchesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProtectedBranchesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProtectedBranchesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProtectedBranchesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProtectedBranchesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProtectedBranchesResponse wrapper for the ListProtectedBranches operation
type ListProtectedBranchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProtectedBranchCollection instances
	ProtectedBranchCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProtectedBranchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProtectedBranchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProtectedBranchesSortOrderEnum Enum with underlying type: string
type ListProtectedBranchesSortOrderEnum string

// Set of constants representing the allowable values for ListProtectedBranchesSortOrderEnum
const (
	ListProtectedBranchesSortOrderAsc  ListProtectedBranchesSortOrderEnum = "ASC"
	ListProtectedBranchesSortOrderDesc ListProtectedBranchesSortOrderEnum = "DESC"
)

var mappingListProtectedBranchesSortOrderEnum = map[string]ListProtectedBranchesSortOrderEnum{
	"ASC":  ListProtectedBranchesSortOrderAsc,
	"DESC": ListProtectedBranchesSortOrderDesc,
}

var mappingListProtectedBranchesSortOrderEnumLowerCase = map[string]ListProtectedBranchesSortOrderEnum{
	"asc":  ListProtectedBranchesSortOrderAsc,
	"desc": ListProtectedBranchesSortOrderDesc,
}

// GetListProtectedBranchesSortOrderEnumValues Enumerates the set of values for ListProtectedBranchesSortOrderEnum
func GetListProtectedBranchesSortOrderEnumValues() []ListProtectedBranchesSortOrderEnum {
	values := make([]ListProtectedBranchesSortOrderEnum, 0)
	for _, v := range mappingListProtectedBranchesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProtectedBranchesSortOrderEnumStringValues Enumerates the set of values in String for ListProtectedBranchesSortOrderEnum
func GetListProtectedBranchesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProtectedBranchesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProtectedBranchesSortOrderEnum(val string) (ListProtectedBranchesSortOrderEnum, bool) {
	enum, ok := mappingListProtectedBranchesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProtectedBranchesSortByEnum Enum with underlying type: string
type ListProtectedBranchesSortByEnum string

// Set of constants representing the allowable values for ListProtectedBranchesSortByEnum
const (
	ListProtectedBranchesSortByBranchname ListProtectedBranchesSortByEnum = "branchName"
)

var mappingListProtectedBranchesSortByEnum = map[string]ListProtectedBranchesSortByEnum{
	"branchName": ListProtectedBranchesSortByBranchname,
}

var mappingListProtectedBranchesSortByEnumLowerCase = map[string]ListProtectedBranchesSortByEnum{
	"branchname": ListProtectedBranchesSortByBranchname,
}

// GetListProtectedBranchesSortByEnumValues Enumerates the set of values for ListProtectedBranchesSortByEnum
func GetListProtectedBranchesSortByEnumValues() []ListProtectedBranchesSortByEnum {
	values := make([]ListProtectedBranchesSortByEnum, 0)
	for _, v := range mappingListProtectedBranchesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProtectedBranchesSortByEnumStringValues Enumerates the set of values in String for ListProtectedBranchesSortByEnum
func GetListProtectedBranchesSortByEnumStringValues() []string {
	return []string{
		"branchName",
	}
}

// GetMappingListProtectedBranchesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProtectedBranchesSortByEnum(val string) (ListProtectedBranchesSortByEnum, bool) {
	enum, ok := mappingListProtectedBranchesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
