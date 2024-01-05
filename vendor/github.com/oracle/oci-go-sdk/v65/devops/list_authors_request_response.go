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

// ListAuthorsRequest wrapper for the ListAuthors operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListAuthors.go.html to see an example of how to use ListAuthorsRequest.
type ListAuthorsRequest struct {

	// Unique repository identifier.
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// A filter to return only resources that match the given reference name.
	RefName *string `mandatory:"false" contributesTo:"query" name:"refName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListAuthorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAuthorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAuthorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAuthorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAuthorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAuthorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAuthorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAuthorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAuthorsResponse wrapper for the ListAuthors operation
type ListAuthorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RepositoryAuthorCollection instances
	RepositoryAuthorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAuthorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAuthorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAuthorsSortOrderEnum Enum with underlying type: string
type ListAuthorsSortOrderEnum string

// Set of constants representing the allowable values for ListAuthorsSortOrderEnum
const (
	ListAuthorsSortOrderAsc  ListAuthorsSortOrderEnum = "ASC"
	ListAuthorsSortOrderDesc ListAuthorsSortOrderEnum = "DESC"
)

var mappingListAuthorsSortOrderEnum = map[string]ListAuthorsSortOrderEnum{
	"ASC":  ListAuthorsSortOrderAsc,
	"DESC": ListAuthorsSortOrderDesc,
}

var mappingListAuthorsSortOrderEnumLowerCase = map[string]ListAuthorsSortOrderEnum{
	"asc":  ListAuthorsSortOrderAsc,
	"desc": ListAuthorsSortOrderDesc,
}

// GetListAuthorsSortOrderEnumValues Enumerates the set of values for ListAuthorsSortOrderEnum
func GetListAuthorsSortOrderEnumValues() []ListAuthorsSortOrderEnum {
	values := make([]ListAuthorsSortOrderEnum, 0)
	for _, v := range mappingListAuthorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuthorsSortOrderEnumStringValues Enumerates the set of values in String for ListAuthorsSortOrderEnum
func GetListAuthorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAuthorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuthorsSortOrderEnum(val string) (ListAuthorsSortOrderEnum, bool) {
	enum, ok := mappingListAuthorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
