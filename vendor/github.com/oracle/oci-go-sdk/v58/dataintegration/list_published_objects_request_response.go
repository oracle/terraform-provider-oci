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

// ListPublishedObjectsRequest wrapper for the ListPublishedObjects operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListPublishedObjects.go.html to see an example of how to use ListPublishedObjectsRequest.
type ListPublishedObjectsRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The application key.
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// This parameter can be used to filter objects by the names starting with the given value.
	NameStartsWith *string `mandatory:"false" contributesTo:"query" name:"nameStartsWith"`

	// This parameter can be used to filter objects by the names that match partially or fully with the given value.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// Used to filter by the identifier of the published object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// Used to filter by the object type of the object.
	// It can be suffixed with an optional filter operator InSubtree.
	// For Data Integration APIs, a filter based on type Task is used.
	Type []string `contributesTo:"query" name:"type" collectionFormat:"multi"`

	// Used in association with type parameter. If value is true,
	// then type all sub types of the given type parameter is considered.
	// If value is false, then sub types are not considered. Default is false.
	TypeInSubtree *string `mandatory:"false" contributesTo:"query" name:"typeInSubtree"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListPublishedObjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListPublishedObjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPublishedObjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPublishedObjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPublishedObjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPublishedObjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPublishedObjectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPublishedObjectsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPublishedObjectsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPublishedObjectsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPublishedObjectsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPublishedObjectsResponse wrapper for the ListPublishedObjects operation
type ListPublishedObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PublishedObjectSummaryCollection instances
	PublishedObjectSummaryCollection `presentIn:"body"`

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

func (response ListPublishedObjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPublishedObjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPublishedObjectsSortOrderEnum Enum with underlying type: string
type ListPublishedObjectsSortOrderEnum string

// Set of constants representing the allowable values for ListPublishedObjectsSortOrderEnum
const (
	ListPublishedObjectsSortOrderAsc  ListPublishedObjectsSortOrderEnum = "ASC"
	ListPublishedObjectsSortOrderDesc ListPublishedObjectsSortOrderEnum = "DESC"
)

var mappingListPublishedObjectsSortOrderEnum = map[string]ListPublishedObjectsSortOrderEnum{
	"ASC":  ListPublishedObjectsSortOrderAsc,
	"DESC": ListPublishedObjectsSortOrderDesc,
}

// GetListPublishedObjectsSortOrderEnumValues Enumerates the set of values for ListPublishedObjectsSortOrderEnum
func GetListPublishedObjectsSortOrderEnumValues() []ListPublishedObjectsSortOrderEnum {
	values := make([]ListPublishedObjectsSortOrderEnum, 0)
	for _, v := range mappingListPublishedObjectsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPublishedObjectsSortOrderEnumStringValues Enumerates the set of values in String for ListPublishedObjectsSortOrderEnum
func GetListPublishedObjectsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPublishedObjectsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPublishedObjectsSortOrderEnum(val string) (ListPublishedObjectsSortOrderEnum, bool) {
	mappingListPublishedObjectsSortOrderEnumIgnoreCase := make(map[string]ListPublishedObjectsSortOrderEnum)
	for k, v := range mappingListPublishedObjectsSortOrderEnum {
		mappingListPublishedObjectsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListPublishedObjectsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListPublishedObjectsSortByEnum Enum with underlying type: string
type ListPublishedObjectsSortByEnum string

// Set of constants representing the allowable values for ListPublishedObjectsSortByEnum
const (
	ListPublishedObjectsSortByTimeCreated ListPublishedObjectsSortByEnum = "TIME_CREATED"
	ListPublishedObjectsSortByDisplayName ListPublishedObjectsSortByEnum = "DISPLAY_NAME"
)

var mappingListPublishedObjectsSortByEnum = map[string]ListPublishedObjectsSortByEnum{
	"TIME_CREATED": ListPublishedObjectsSortByTimeCreated,
	"DISPLAY_NAME": ListPublishedObjectsSortByDisplayName,
}

// GetListPublishedObjectsSortByEnumValues Enumerates the set of values for ListPublishedObjectsSortByEnum
func GetListPublishedObjectsSortByEnumValues() []ListPublishedObjectsSortByEnum {
	values := make([]ListPublishedObjectsSortByEnum, 0)
	for _, v := range mappingListPublishedObjectsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPublishedObjectsSortByEnumStringValues Enumerates the set of values in String for ListPublishedObjectsSortByEnum
func GetListPublishedObjectsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListPublishedObjectsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPublishedObjectsSortByEnum(val string) (ListPublishedObjectsSortByEnum, bool) {
	mappingListPublishedObjectsSortByEnumIgnoreCase := make(map[string]ListPublishedObjectsSortByEnum)
	for k, v := range mappingListPublishedObjectsSortByEnum {
		mappingListPublishedObjectsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListPublishedObjectsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
