// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListPathsRequest wrapper for the ListPaths operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListPaths.go.html to see an example of how to use ListPathsRequest.
type ListPathsRequest struct {

	// Unique repository identifier.
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// The name of branch/tag or commit hash it points to. If names conflict, order of preference is commit > branch > tag.
	// You can disambiguate with "heads/foobar" and "tags/foobar". If left blank repository's default branch will be used.
	Ref *string `mandatory:"false" contributesTo:"query" name:"ref"`

	// Flag to determine if files must be retrived recursively. Flag is False by default.
	PathsInSubtree *bool `mandatory:"false" contributesTo:"query" name:"pathsInSubtree"`

	// The fully qualified path to the folder whose contents are returned, including the folder name. For example, /examples is a fully-qualified path to a folder named examples that was created off of the root directory (/) of a repository.
	FolderPath *string `mandatory:"false" contributesTo:"query" name:"folderPath"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListPathsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order is ascending. If no value is specified name is default.
	SortBy ListPathsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPathsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPathsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPathsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPathsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListPathsResponse wrapper for the ListPaths operation
type ListPathsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RepositoryPathCollection instances
	RepositoryPathCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPathsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPathsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPathsSortOrderEnum Enum with underlying type: string
type ListPathsSortOrderEnum string

// Set of constants representing the allowable values for ListPathsSortOrderEnum
const (
	ListPathsSortOrderAsc  ListPathsSortOrderEnum = "ASC"
	ListPathsSortOrderDesc ListPathsSortOrderEnum = "DESC"
)

var mappingListPathsSortOrder = map[string]ListPathsSortOrderEnum{
	"ASC":  ListPathsSortOrderAsc,
	"DESC": ListPathsSortOrderDesc,
}

// GetListPathsSortOrderEnumValues Enumerates the set of values for ListPathsSortOrderEnum
func GetListPathsSortOrderEnumValues() []ListPathsSortOrderEnum {
	values := make([]ListPathsSortOrderEnum, 0)
	for _, v := range mappingListPathsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListPathsSortByEnum Enum with underlying type: string
type ListPathsSortByEnum string

// Set of constants representing the allowable values for ListPathsSortByEnum
const (
	ListPathsSortByType        ListPathsSortByEnum = "type"
	ListPathsSortBySizeinbytes ListPathsSortByEnum = "sizeInBytes"
	ListPathsSortByName        ListPathsSortByEnum = "name"
)

var mappingListPathsSortBy = map[string]ListPathsSortByEnum{
	"type":        ListPathsSortByType,
	"sizeInBytes": ListPathsSortBySizeinbytes,
	"name":        ListPathsSortByName,
}

// GetListPathsSortByEnumValues Enumerates the set of values for ListPathsSortByEnum
func GetListPathsSortByEnumValues() []ListPathsSortByEnum {
	values := make([]ListPathsSortByEnum, 0)
	for _, v := range mappingListPathsSortBy {
		values = append(values, v)
	}
	return values
}
