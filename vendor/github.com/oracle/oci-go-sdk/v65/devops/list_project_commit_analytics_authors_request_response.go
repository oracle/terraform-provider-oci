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

// ListProjectCommitAnalyticsAuthorsRequest wrapper for the ListProjectCommitAnalyticsAuthors operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListProjectCommitAnalyticsAuthors.go.html to see an example of how to use ListProjectCommitAnalyticsAuthorsRequest.
type ListProjectCommitAnalyticsAuthorsRequest struct {

	// Unique project identifier.
	ProjectId *string `mandatory:"true" contributesTo:"path" name:"projectId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListProjectCommitAnalyticsAuthorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. Only one sort by value is supported for this parameter. Default order for author name is ascending.
	SortBy ListProjectCommitAnalyticsAuthorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProjectCommitAnalyticsAuthorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProjectCommitAnalyticsAuthorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProjectCommitAnalyticsAuthorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProjectCommitAnalyticsAuthorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProjectCommitAnalyticsAuthorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProjectCommitAnalyticsAuthorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProjectCommitAnalyticsAuthorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProjectCommitAnalyticsAuthorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProjectCommitAnalyticsAuthorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProjectCommitAnalyticsAuthorsResponse wrapper for the ListProjectCommitAnalyticsAuthors operation
type ListProjectCommitAnalyticsAuthorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CommitAnalyticsAuthorCollection instances
	CommitAnalyticsAuthorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProjectCommitAnalyticsAuthorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProjectCommitAnalyticsAuthorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProjectCommitAnalyticsAuthorsSortOrderEnum Enum with underlying type: string
type ListProjectCommitAnalyticsAuthorsSortOrderEnum string

// Set of constants representing the allowable values for ListProjectCommitAnalyticsAuthorsSortOrderEnum
const (
	ListProjectCommitAnalyticsAuthorsSortOrderAsc  ListProjectCommitAnalyticsAuthorsSortOrderEnum = "ASC"
	ListProjectCommitAnalyticsAuthorsSortOrderDesc ListProjectCommitAnalyticsAuthorsSortOrderEnum = "DESC"
)

var mappingListProjectCommitAnalyticsAuthorsSortOrderEnum = map[string]ListProjectCommitAnalyticsAuthorsSortOrderEnum{
	"ASC":  ListProjectCommitAnalyticsAuthorsSortOrderAsc,
	"DESC": ListProjectCommitAnalyticsAuthorsSortOrderDesc,
}

var mappingListProjectCommitAnalyticsAuthorsSortOrderEnumLowerCase = map[string]ListProjectCommitAnalyticsAuthorsSortOrderEnum{
	"asc":  ListProjectCommitAnalyticsAuthorsSortOrderAsc,
	"desc": ListProjectCommitAnalyticsAuthorsSortOrderDesc,
}

// GetListProjectCommitAnalyticsAuthorsSortOrderEnumValues Enumerates the set of values for ListProjectCommitAnalyticsAuthorsSortOrderEnum
func GetListProjectCommitAnalyticsAuthorsSortOrderEnumValues() []ListProjectCommitAnalyticsAuthorsSortOrderEnum {
	values := make([]ListProjectCommitAnalyticsAuthorsSortOrderEnum, 0)
	for _, v := range mappingListProjectCommitAnalyticsAuthorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProjectCommitAnalyticsAuthorsSortOrderEnumStringValues Enumerates the set of values in String for ListProjectCommitAnalyticsAuthorsSortOrderEnum
func GetListProjectCommitAnalyticsAuthorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProjectCommitAnalyticsAuthorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProjectCommitAnalyticsAuthorsSortOrderEnum(val string) (ListProjectCommitAnalyticsAuthorsSortOrderEnum, bool) {
	enum, ok := mappingListProjectCommitAnalyticsAuthorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProjectCommitAnalyticsAuthorsSortByEnum Enum with underlying type: string
type ListProjectCommitAnalyticsAuthorsSortByEnum string

// Set of constants representing the allowable values for ListProjectCommitAnalyticsAuthorsSortByEnum
const (
	ListProjectCommitAnalyticsAuthorsSortByAuthorname ListProjectCommitAnalyticsAuthorsSortByEnum = "authorName"
)

var mappingListProjectCommitAnalyticsAuthorsSortByEnum = map[string]ListProjectCommitAnalyticsAuthorsSortByEnum{
	"authorName": ListProjectCommitAnalyticsAuthorsSortByAuthorname,
}

var mappingListProjectCommitAnalyticsAuthorsSortByEnumLowerCase = map[string]ListProjectCommitAnalyticsAuthorsSortByEnum{
	"authorname": ListProjectCommitAnalyticsAuthorsSortByAuthorname,
}

// GetListProjectCommitAnalyticsAuthorsSortByEnumValues Enumerates the set of values for ListProjectCommitAnalyticsAuthorsSortByEnum
func GetListProjectCommitAnalyticsAuthorsSortByEnumValues() []ListProjectCommitAnalyticsAuthorsSortByEnum {
	values := make([]ListProjectCommitAnalyticsAuthorsSortByEnum, 0)
	for _, v := range mappingListProjectCommitAnalyticsAuthorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProjectCommitAnalyticsAuthorsSortByEnumStringValues Enumerates the set of values in String for ListProjectCommitAnalyticsAuthorsSortByEnum
func GetListProjectCommitAnalyticsAuthorsSortByEnumStringValues() []string {
	return []string{
		"authorName",
	}
}

// GetMappingListProjectCommitAnalyticsAuthorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProjectCommitAnalyticsAuthorsSortByEnum(val string) (ListProjectCommitAnalyticsAuthorsSortByEnum, bool) {
	enum, ok := mappingListProjectCommitAnalyticsAuthorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
