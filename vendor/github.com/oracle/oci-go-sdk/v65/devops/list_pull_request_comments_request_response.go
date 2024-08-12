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

// ListPullRequestCommentsRequest wrapper for the ListPullRequestComments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListPullRequestComments.go.html to see an example of how to use ListPullRequestCommentsRequest.
type ListPullRequestCommentsRequest struct {

	// unique PullRequest identifier
	PullRequestId *string `mandatory:"true" contributesTo:"path" name:"pullRequestId"`

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListPullRequestCommentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order is ascending. If no value is specified timeCreated is default.
	SortBy ListPullRequestCommentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// unique PullRequest Comment identifier
	CommentId *string `mandatory:"false" contributesTo:"query" name:"commentId"`

	// PullRequest Comment Commit SHA
	CommitId *string `mandatory:"false" contributesTo:"query" name:"commitId"`

	// PullRequest File Path
	FilePath *string `mandatory:"false" contributesTo:"query" name:"filePath"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPullRequestCommentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPullRequestCommentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPullRequestCommentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPullRequestCommentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPullRequestCommentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPullRequestCommentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPullRequestCommentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPullRequestCommentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPullRequestCommentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPullRequestCommentsResponse wrapper for the ListPullRequestComments operation
type ListPullRequestCommentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PullRequestCommentCollection instances
	PullRequestCommentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPullRequestCommentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPullRequestCommentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPullRequestCommentsSortOrderEnum Enum with underlying type: string
type ListPullRequestCommentsSortOrderEnum string

// Set of constants representing the allowable values for ListPullRequestCommentsSortOrderEnum
const (
	ListPullRequestCommentsSortOrderAsc  ListPullRequestCommentsSortOrderEnum = "ASC"
	ListPullRequestCommentsSortOrderDesc ListPullRequestCommentsSortOrderEnum = "DESC"
)

var mappingListPullRequestCommentsSortOrderEnum = map[string]ListPullRequestCommentsSortOrderEnum{
	"ASC":  ListPullRequestCommentsSortOrderAsc,
	"DESC": ListPullRequestCommentsSortOrderDesc,
}

var mappingListPullRequestCommentsSortOrderEnumLowerCase = map[string]ListPullRequestCommentsSortOrderEnum{
	"asc":  ListPullRequestCommentsSortOrderAsc,
	"desc": ListPullRequestCommentsSortOrderDesc,
}

// GetListPullRequestCommentsSortOrderEnumValues Enumerates the set of values for ListPullRequestCommentsSortOrderEnum
func GetListPullRequestCommentsSortOrderEnumValues() []ListPullRequestCommentsSortOrderEnum {
	values := make([]ListPullRequestCommentsSortOrderEnum, 0)
	for _, v := range mappingListPullRequestCommentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPullRequestCommentsSortOrderEnumStringValues Enumerates the set of values in String for ListPullRequestCommentsSortOrderEnum
func GetListPullRequestCommentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPullRequestCommentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPullRequestCommentsSortOrderEnum(val string) (ListPullRequestCommentsSortOrderEnum, bool) {
	enum, ok := mappingListPullRequestCommentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPullRequestCommentsSortByEnum Enum with underlying type: string
type ListPullRequestCommentsSortByEnum string

// Set of constants representing the allowable values for ListPullRequestCommentsSortByEnum
const (
	ListPullRequestCommentsSortByTimecreated ListPullRequestCommentsSortByEnum = "timeCreated"
	ListPullRequestCommentsSortByCreatedby   ListPullRequestCommentsSortByEnum = "createdBy"
)

var mappingListPullRequestCommentsSortByEnum = map[string]ListPullRequestCommentsSortByEnum{
	"timeCreated": ListPullRequestCommentsSortByTimecreated,
	"createdBy":   ListPullRequestCommentsSortByCreatedby,
}

var mappingListPullRequestCommentsSortByEnumLowerCase = map[string]ListPullRequestCommentsSortByEnum{
	"timecreated": ListPullRequestCommentsSortByTimecreated,
	"createdby":   ListPullRequestCommentsSortByCreatedby,
}

// GetListPullRequestCommentsSortByEnumValues Enumerates the set of values for ListPullRequestCommentsSortByEnum
func GetListPullRequestCommentsSortByEnumValues() []ListPullRequestCommentsSortByEnum {
	values := make([]ListPullRequestCommentsSortByEnum, 0)
	for _, v := range mappingListPullRequestCommentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPullRequestCommentsSortByEnumStringValues Enumerates the set of values in String for ListPullRequestCommentsSortByEnum
func GetListPullRequestCommentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"createdBy",
	}
}

// GetMappingListPullRequestCommentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPullRequestCommentsSortByEnum(val string) (ListPullRequestCommentsSortByEnum, bool) {
	enum, ok := mappingListPullRequestCommentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
