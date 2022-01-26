// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListRefsRequest wrapper for the ListRefs operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListRefs.go.html to see an example of how to use ListRefsRequest.
type ListRefsRequest struct {

	// Unique repository identifier.
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// Reference type to distinguish between branch and tag. If it is not specified, all references are returned.
	RefType ListRefsRefTypeEnum `mandatory:"false" contributesTo:"query" name:"refType" omitEmpty:"true"`

	// Commit ID in a repository.
	CommitId *string `mandatory:"false" contributesTo:"query" name:"commitId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given reference name.
	RefName *string `mandatory:"false" contributesTo:"query" name:"refName"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListRefsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for reference name is ascending. Default order for reference type is ascending. If no value is specified reference name is default.
	SortBy ListRefsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRefsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRefsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRefsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRefsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListRefsResponse wrapper for the ListRefs operation
type ListRefsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RepositoryRefCollection instances
	RepositoryRefCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRefsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRefsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRefsRefTypeEnum Enum with underlying type: string
type ListRefsRefTypeEnum string

// Set of constants representing the allowable values for ListRefsRefTypeEnum
const (
	ListRefsRefTypeBranch ListRefsRefTypeEnum = "BRANCH"
	ListRefsRefTypeTag    ListRefsRefTypeEnum = "TAG"
)

var mappingListRefsRefType = map[string]ListRefsRefTypeEnum{
	"BRANCH": ListRefsRefTypeBranch,
	"TAG":    ListRefsRefTypeTag,
}

// GetListRefsRefTypeEnumValues Enumerates the set of values for ListRefsRefTypeEnum
func GetListRefsRefTypeEnumValues() []ListRefsRefTypeEnum {
	values := make([]ListRefsRefTypeEnum, 0)
	for _, v := range mappingListRefsRefType {
		values = append(values, v)
	}
	return values
}

// ListRefsSortOrderEnum Enum with underlying type: string
type ListRefsSortOrderEnum string

// Set of constants representing the allowable values for ListRefsSortOrderEnum
const (
	ListRefsSortOrderAsc  ListRefsSortOrderEnum = "ASC"
	ListRefsSortOrderDesc ListRefsSortOrderEnum = "DESC"
)

var mappingListRefsSortOrder = map[string]ListRefsSortOrderEnum{
	"ASC":  ListRefsSortOrderAsc,
	"DESC": ListRefsSortOrderDesc,
}

// GetListRefsSortOrderEnumValues Enumerates the set of values for ListRefsSortOrderEnum
func GetListRefsSortOrderEnumValues() []ListRefsSortOrderEnum {
	values := make([]ListRefsSortOrderEnum, 0)
	for _, v := range mappingListRefsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListRefsSortByEnum Enum with underlying type: string
type ListRefsSortByEnum string

// Set of constants representing the allowable values for ListRefsSortByEnum
const (
	ListRefsSortByReftype ListRefsSortByEnum = "refType"
	ListRefsSortByRefname ListRefsSortByEnum = "refName"
)

var mappingListRefsSortBy = map[string]ListRefsSortByEnum{
	"refType": ListRefsSortByReftype,
	"refName": ListRefsSortByRefname,
}

// GetListRefsSortByEnumValues Enumerates the set of values for ListRefsSortByEnum
func GetListRefsSortByEnumValues() []ListRefsSortByEnum {
	values := make([]ListRefsSortByEnum, 0)
	for _, v := range mappingListRefsSortBy {
		values = append(values, v)
	}
	return values
}
