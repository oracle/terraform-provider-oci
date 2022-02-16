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

// ListDataAssetsRequest wrapper for the ListDataAssets operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListDataAssets.go.html to see an example of how to use ListDataAssetsRequest.
type ListDataAssetsRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Type of the object to filter the results with.
	Type *string `mandatory:"false" contributesTo:"query" name:"type"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListDataAssetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListDataAssetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataAssetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataAssetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataAssetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataAssetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataAssetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataAssetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataAssetsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataAssetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataAssetsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataAssetsResponse wrapper for the ListDataAssets operation
type ListDataAssetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataAssetSummaryCollection instances
	DataAssetSummaryCollection `presentIn:"body"`

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

func (response ListDataAssetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataAssetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataAssetsSortByEnum Enum with underlying type: string
type ListDataAssetsSortByEnum string

// Set of constants representing the allowable values for ListDataAssetsSortByEnum
const (
	ListDataAssetsSortByTimeCreated ListDataAssetsSortByEnum = "TIME_CREATED"
	ListDataAssetsSortByDisplayName ListDataAssetsSortByEnum = "DISPLAY_NAME"
)

var mappingListDataAssetsSortByEnum = map[string]ListDataAssetsSortByEnum{
	"TIME_CREATED": ListDataAssetsSortByTimeCreated,
	"DISPLAY_NAME": ListDataAssetsSortByDisplayName,
}

// GetListDataAssetsSortByEnumValues Enumerates the set of values for ListDataAssetsSortByEnum
func GetListDataAssetsSortByEnumValues() []ListDataAssetsSortByEnum {
	values := make([]ListDataAssetsSortByEnum, 0)
	for _, v := range mappingListDataAssetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataAssetsSortByEnumStringValues Enumerates the set of values in String for ListDataAssetsSortByEnum
func GetListDataAssetsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListDataAssetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataAssetsSortByEnum(val string) (ListDataAssetsSortByEnum, bool) {
	mappingListDataAssetsSortByEnumIgnoreCase := make(map[string]ListDataAssetsSortByEnum)
	for k, v := range mappingListDataAssetsSortByEnum {
		mappingListDataAssetsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDataAssetsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataAssetsSortOrderEnum Enum with underlying type: string
type ListDataAssetsSortOrderEnum string

// Set of constants representing the allowable values for ListDataAssetsSortOrderEnum
const (
	ListDataAssetsSortOrderAsc  ListDataAssetsSortOrderEnum = "ASC"
	ListDataAssetsSortOrderDesc ListDataAssetsSortOrderEnum = "DESC"
)

var mappingListDataAssetsSortOrderEnum = map[string]ListDataAssetsSortOrderEnum{
	"ASC":  ListDataAssetsSortOrderAsc,
	"DESC": ListDataAssetsSortOrderDesc,
}

// GetListDataAssetsSortOrderEnumValues Enumerates the set of values for ListDataAssetsSortOrderEnum
func GetListDataAssetsSortOrderEnumValues() []ListDataAssetsSortOrderEnum {
	values := make([]ListDataAssetsSortOrderEnum, 0)
	for _, v := range mappingListDataAssetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataAssetsSortOrderEnumStringValues Enumerates the set of values in String for ListDataAssetsSortOrderEnum
func GetListDataAssetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataAssetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataAssetsSortOrderEnum(val string) (ListDataAssetsSortOrderEnum, bool) {
	mappingListDataAssetsSortOrderEnumIgnoreCase := make(map[string]ListDataAssetsSortOrderEnum)
	for k, v := range mappingListDataAssetsSortOrderEnum {
		mappingListDataAssetsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDataAssetsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
