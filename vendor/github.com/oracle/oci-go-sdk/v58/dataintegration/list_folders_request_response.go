// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListFoldersRequest wrapper for the ListFolders operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListFolders.go.html to see an example of how to use ListFoldersRequest.
type ListFoldersRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Used to filter by the project or the folder object.
	AggregatorKey *string `mandatory:"false" contributesTo:"query" name:"aggregatorKey"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// This parameter can be used to filter objects by the names that match partially or fully with the given value.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// Used to filter by the identifier of the object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListFoldersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListFoldersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFoldersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFoldersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFoldersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFoldersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFoldersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFoldersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFoldersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFoldersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFoldersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFoldersResponse wrapper for the ListFolders operation
type ListFoldersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FolderSummaryCollection instances
	FolderSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListFoldersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFoldersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFoldersSortOrderEnum Enum with underlying type: string
type ListFoldersSortOrderEnum string

// Set of constants representing the allowable values for ListFoldersSortOrderEnum
const (
	ListFoldersSortOrderAsc  ListFoldersSortOrderEnum = "ASC"
	ListFoldersSortOrderDesc ListFoldersSortOrderEnum = "DESC"
)

var mappingListFoldersSortOrderEnum = map[string]ListFoldersSortOrderEnum{
	"ASC":  ListFoldersSortOrderAsc,
	"DESC": ListFoldersSortOrderDesc,
}

// GetListFoldersSortOrderEnumValues Enumerates the set of values for ListFoldersSortOrderEnum
func GetListFoldersSortOrderEnumValues() []ListFoldersSortOrderEnum {
	values := make([]ListFoldersSortOrderEnum, 0)
	for _, v := range mappingListFoldersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFoldersSortOrderEnumStringValues Enumerates the set of values in String for ListFoldersSortOrderEnum
func GetListFoldersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFoldersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFoldersSortOrderEnum(val string) (ListFoldersSortOrderEnum, bool) {
	mappingListFoldersSortOrderEnumIgnoreCase := make(map[string]ListFoldersSortOrderEnum)
	for k, v := range mappingListFoldersSortOrderEnum {
		mappingListFoldersSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListFoldersSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListFoldersSortByEnum Enum with underlying type: string
type ListFoldersSortByEnum string

// Set of constants representing the allowable values for ListFoldersSortByEnum
const (
	ListFoldersSortByTimeCreated ListFoldersSortByEnum = "TIME_CREATED"
	ListFoldersSortByDisplayName ListFoldersSortByEnum = "DISPLAY_NAME"
)

var mappingListFoldersSortByEnum = map[string]ListFoldersSortByEnum{
	"TIME_CREATED": ListFoldersSortByTimeCreated,
	"DISPLAY_NAME": ListFoldersSortByDisplayName,
}

// GetListFoldersSortByEnumValues Enumerates the set of values for ListFoldersSortByEnum
func GetListFoldersSortByEnumValues() []ListFoldersSortByEnum {
	values := make([]ListFoldersSortByEnum, 0)
	for _, v := range mappingListFoldersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFoldersSortByEnumStringValues Enumerates the set of values in String for ListFoldersSortByEnum
func GetListFoldersSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListFoldersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFoldersSortByEnum(val string) (ListFoldersSortByEnum, bool) {
	mappingListFoldersSortByEnumIgnoreCase := make(map[string]ListFoldersSortByEnum)
	for k, v := range mappingListFoldersSortByEnum {
		mappingListFoldersSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListFoldersSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
