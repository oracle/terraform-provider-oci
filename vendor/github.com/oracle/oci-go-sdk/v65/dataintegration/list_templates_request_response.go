// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTemplatesRequest wrapper for the ListTemplates operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListTemplates.go.html to see an example of how to use ListTemplatesRequest.
type ListTemplatesRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Used to filter by the identifier of the published object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListTemplatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListTemplatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTemplatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTemplatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTemplatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTemplatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTemplatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTemplatesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTemplatesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTemplatesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTemplatesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTemplatesResponse wrapper for the ListTemplates operation
type ListTemplatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TemplateSummaryCollection instances
	TemplateSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListTemplatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTemplatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTemplatesSortOrderEnum Enum with underlying type: string
type ListTemplatesSortOrderEnum string

// Set of constants representing the allowable values for ListTemplatesSortOrderEnum
const (
	ListTemplatesSortOrderAsc  ListTemplatesSortOrderEnum = "ASC"
	ListTemplatesSortOrderDesc ListTemplatesSortOrderEnum = "DESC"
)

var mappingListTemplatesSortOrderEnum = map[string]ListTemplatesSortOrderEnum{
	"ASC":  ListTemplatesSortOrderAsc,
	"DESC": ListTemplatesSortOrderDesc,
}

var mappingListTemplatesSortOrderEnumLowerCase = map[string]ListTemplatesSortOrderEnum{
	"asc":  ListTemplatesSortOrderAsc,
	"desc": ListTemplatesSortOrderDesc,
}

// GetListTemplatesSortOrderEnumValues Enumerates the set of values for ListTemplatesSortOrderEnum
func GetListTemplatesSortOrderEnumValues() []ListTemplatesSortOrderEnum {
	values := make([]ListTemplatesSortOrderEnum, 0)
	for _, v := range mappingListTemplatesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTemplatesSortOrderEnumStringValues Enumerates the set of values in String for ListTemplatesSortOrderEnum
func GetListTemplatesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTemplatesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTemplatesSortOrderEnum(val string) (ListTemplatesSortOrderEnum, bool) {
	enum, ok := mappingListTemplatesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTemplatesSortByEnum Enum with underlying type: string
type ListTemplatesSortByEnum string

// Set of constants representing the allowable values for ListTemplatesSortByEnum
const (
	ListTemplatesSortByTimeCreated ListTemplatesSortByEnum = "TIME_CREATED"
	ListTemplatesSortByDisplayName ListTemplatesSortByEnum = "DISPLAY_NAME"
	ListTemplatesSortByTimeUpdated ListTemplatesSortByEnum = "TIME_UPDATED"
)

var mappingListTemplatesSortByEnum = map[string]ListTemplatesSortByEnum{
	"TIME_CREATED": ListTemplatesSortByTimeCreated,
	"DISPLAY_NAME": ListTemplatesSortByDisplayName,
	"TIME_UPDATED": ListTemplatesSortByTimeUpdated,
}

var mappingListTemplatesSortByEnumLowerCase = map[string]ListTemplatesSortByEnum{
	"time_created": ListTemplatesSortByTimeCreated,
	"display_name": ListTemplatesSortByDisplayName,
	"time_updated": ListTemplatesSortByTimeUpdated,
}

// GetListTemplatesSortByEnumValues Enumerates the set of values for ListTemplatesSortByEnum
func GetListTemplatesSortByEnumValues() []ListTemplatesSortByEnum {
	values := make([]ListTemplatesSortByEnum, 0)
	for _, v := range mappingListTemplatesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTemplatesSortByEnumStringValues Enumerates the set of values in String for ListTemplatesSortByEnum
func GetListTemplatesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListTemplatesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTemplatesSortByEnum(val string) (ListTemplatesSortByEnum, bool) {
	enum, ok := mappingListTemplatesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
