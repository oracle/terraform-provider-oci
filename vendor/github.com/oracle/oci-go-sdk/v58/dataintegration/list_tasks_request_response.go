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

// ListTasksRequest wrapper for the ListTasks operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListTasks.go.html to see an example of how to use ListTasksRequest.
type ListTasksRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Unique key of the folder.
	FolderId *string `mandatory:"false" contributesTo:"query" name:"folderId"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Used to filter by the key of the object.
	Key []string `contributesTo:"query" name:"key" collectionFormat:"multi"`

	// Used to filter by the identifier of the object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// Used to filter by the object type of the object. It can be suffixed with an optional filter operator InSubtree. If this operator is not specified, then exact match is considered. <br><br><B>Examples:</B><br> <ul> <li><B>?type=DATA_LOADER_TASK&typeInSubtree=false</B> returns all objects of type data loader task</li> <li><B>?type=DATA_LOADER_TASK</B> returns all objects of type data loader task</li> <li><B>?type=DATA_LOADER_TASK&typeInSubtree=true</B> returns all objects of type data loader task</li> </ul>
	Type []string `contributesTo:"query" name:"type" collectionFormat:"multi"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListTasksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListTasksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTasksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTasksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTasksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTasksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTasksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTasksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTasksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTasksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTasksSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTasksResponse wrapper for the ListTasks operation
type ListTasksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TaskSummaryCollection instances
	TaskSummaryCollection `presentIn:"body"`

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

func (response ListTasksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTasksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTasksSortOrderEnum Enum with underlying type: string
type ListTasksSortOrderEnum string

// Set of constants representing the allowable values for ListTasksSortOrderEnum
const (
	ListTasksSortOrderAsc  ListTasksSortOrderEnum = "ASC"
	ListTasksSortOrderDesc ListTasksSortOrderEnum = "DESC"
)

var mappingListTasksSortOrderEnum = map[string]ListTasksSortOrderEnum{
	"ASC":  ListTasksSortOrderAsc,
	"DESC": ListTasksSortOrderDesc,
}

// GetListTasksSortOrderEnumValues Enumerates the set of values for ListTasksSortOrderEnum
func GetListTasksSortOrderEnumValues() []ListTasksSortOrderEnum {
	values := make([]ListTasksSortOrderEnum, 0)
	for _, v := range mappingListTasksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTasksSortOrderEnumStringValues Enumerates the set of values in String for ListTasksSortOrderEnum
func GetListTasksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTasksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTasksSortOrderEnum(val string) (ListTasksSortOrderEnum, bool) {
	mappingListTasksSortOrderEnumIgnoreCase := make(map[string]ListTasksSortOrderEnum)
	for k, v := range mappingListTasksSortOrderEnum {
		mappingListTasksSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListTasksSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListTasksSortByEnum Enum with underlying type: string
type ListTasksSortByEnum string

// Set of constants representing the allowable values for ListTasksSortByEnum
const (
	ListTasksSortByTimeCreated ListTasksSortByEnum = "TIME_CREATED"
	ListTasksSortByDisplayName ListTasksSortByEnum = "DISPLAY_NAME"
)

var mappingListTasksSortByEnum = map[string]ListTasksSortByEnum{
	"TIME_CREATED": ListTasksSortByTimeCreated,
	"DISPLAY_NAME": ListTasksSortByDisplayName,
}

// GetListTasksSortByEnumValues Enumerates the set of values for ListTasksSortByEnum
func GetListTasksSortByEnumValues() []ListTasksSortByEnum {
	values := make([]ListTasksSortByEnum, 0)
	for _, v := range mappingListTasksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTasksSortByEnumStringValues Enumerates the set of values in String for ListTasksSortByEnum
func GetListTasksSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListTasksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTasksSortByEnum(val string) (ListTasksSortByEnum, bool) {
	mappingListTasksSortByEnumIgnoreCase := make(map[string]ListTasksSortByEnum)
	for k, v := range mappingListTasksSortByEnum {
		mappingListTasksSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListTasksSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
