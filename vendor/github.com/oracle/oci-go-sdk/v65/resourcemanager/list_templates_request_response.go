// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourcemanager

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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListTemplates.go.html to see an example of how to use ListTemplatesRequest.
type ListTemplatesRequest struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that exist in the compartment, identified by OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the template category.
	// Possible values are `0` (Quickstarts), `1` (Service), `2` (Architecture), and `3` (Private).
	// Template category labels are displayed in the Console page listing templates.
	// Quickstarts, Service, and Architecture templates (categories 0, 1, and 2) are available in all compartments.
	// Each private template (category 3) is available in the compartment where it was created.
	TemplateCategoryId *string `mandatory:"false" contributesTo:"query" name:"templateCategoryId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the template.
	TemplateId *string `mandatory:"false" contributesTo:"query" name:"templateId"`

	// A filter to return only resources that match the given display name exactly.
	// Use this filter to list a resource by name.
	// Requires `sortBy` set to `DISPLAYNAME`.
	// Alternatively, when you know the resource OCID, use the related Get operation.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to use when sorting returned resources.
	// By default, `TIMECREATED` is ordered descending.
	// By default, `DISPLAYNAME` is ordered ascending. Note that you can sort only on one field.
	SortBy ListTemplatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use when sorting returned resources. Ascending (`ASC`) or descending (`DESC`).
	SortOrder ListTemplatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The number of items returned in a paginated `List` call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the preceding `List` call.
	// For information about pagination, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

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
	if _, ok := GetMappingListTemplatesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTemplatesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTemplatesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTemplatesSortOrderEnumStringValues(), ",")))
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

	// Unique identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of paginated list items. If the `opc-next-page`
	// header appears in the response, additional pages of results remain.
	// To receive the next page, include the header value in the `page` param.
	// If the `opc-next-page` header does not appear in the response, there
	// are no more list items to get. For more information about list pagination,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTemplatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTemplatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTemplatesSortByEnum Enum with underlying type: string
type ListTemplatesSortByEnum string

// Set of constants representing the allowable values for ListTemplatesSortByEnum
const (
	ListTemplatesSortByTimecreated ListTemplatesSortByEnum = "TIMECREATED"
	ListTemplatesSortByDisplayname ListTemplatesSortByEnum = "DISPLAYNAME"
)

var mappingListTemplatesSortByEnum = map[string]ListTemplatesSortByEnum{
	"TIMECREATED": ListTemplatesSortByTimecreated,
	"DISPLAYNAME": ListTemplatesSortByDisplayname,
}

var mappingListTemplatesSortByEnumLowerCase = map[string]ListTemplatesSortByEnum{
	"timecreated": ListTemplatesSortByTimecreated,
	"displayname": ListTemplatesSortByDisplayname,
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
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListTemplatesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTemplatesSortByEnum(val string) (ListTemplatesSortByEnum, bool) {
	enum, ok := mappingListTemplatesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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
