// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

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
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListTemplates.go.html to see an example of how to use ListTemplatesRequest.
type ListTemplatesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The template type used for filtering. Only templates of the
	// specified type will be returned.
	Type *string `mandatory:"false" contributesTo:"query" name:"type"`

	// The template name used for filtering.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The template display text used for filtering. Only templates with the specified name or
	// description will be returned.
	TemplateDisplayText *string `mandatory:"false" contributesTo:"query" name:"templateDisplayText"`

	// The template lifecycle state used for filtering. Currently supported
	// values are ACTIVE and DELETED.
	LifecycleState ListTemplatesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// filter
	Filter *string `mandatory:"false" contributesTo:"query" name:"filter"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListTemplatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned templates
	SortBy ListTemplatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
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
	if _, ok := GetMappingListTemplatesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTemplatesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTemplatesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTemplatesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTemplatesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTemplatesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTemplatesResponse wrapper for the ListTemplates operation
type ListTemplatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsTemplateCollection instances
	LogAnalyticsTemplateCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListTemplatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTemplatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTemplatesLifecycleStateEnum Enum with underlying type: string
type ListTemplatesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTemplatesLifecycleStateEnum
const (
	ListTemplatesLifecycleStateActive  ListTemplatesLifecycleStateEnum = "ACTIVE"
	ListTemplatesLifecycleStateDeleted ListTemplatesLifecycleStateEnum = "DELETED"
)

var mappingListTemplatesLifecycleStateEnum = map[string]ListTemplatesLifecycleStateEnum{
	"ACTIVE":  ListTemplatesLifecycleStateActive,
	"DELETED": ListTemplatesLifecycleStateDeleted,
}

var mappingListTemplatesLifecycleStateEnumLowerCase = map[string]ListTemplatesLifecycleStateEnum{
	"active":  ListTemplatesLifecycleStateActive,
	"deleted": ListTemplatesLifecycleStateDeleted,
}

// GetListTemplatesLifecycleStateEnumValues Enumerates the set of values for ListTemplatesLifecycleStateEnum
func GetListTemplatesLifecycleStateEnumValues() []ListTemplatesLifecycleStateEnum {
	values := make([]ListTemplatesLifecycleStateEnum, 0)
	for _, v := range mappingListTemplatesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTemplatesLifecycleStateEnumStringValues Enumerates the set of values in String for ListTemplatesLifecycleStateEnum
func GetListTemplatesLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListTemplatesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTemplatesLifecycleStateEnum(val string) (ListTemplatesLifecycleStateEnum, bool) {
	enum, ok := mappingListTemplatesLifecycleStateEnumLowerCase[strings.ToLower(val)]
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

// ListTemplatesSortByEnum Enum with underlying type: string
type ListTemplatesSortByEnum string

// Set of constants representing the allowable values for ListTemplatesSortByEnum
const (
	ListTemplatesSortByDisplayname ListTemplatesSortByEnum = "displayName"
	ListTemplatesSortByTimecreated ListTemplatesSortByEnum = "timeCreated"
	ListTemplatesSortByTimeupdated ListTemplatesSortByEnum = "timeUpdated"
)

var mappingListTemplatesSortByEnum = map[string]ListTemplatesSortByEnum{
	"displayName": ListTemplatesSortByDisplayname,
	"timeCreated": ListTemplatesSortByTimecreated,
	"timeUpdated": ListTemplatesSortByTimeupdated,
}

var mappingListTemplatesSortByEnumLowerCase = map[string]ListTemplatesSortByEnum{
	"displayname": ListTemplatesSortByDisplayname,
	"timecreated": ListTemplatesSortByTimecreated,
	"timeupdated": ListTemplatesSortByTimeupdated,
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
		"displayName",
		"timeCreated",
		"timeUpdated",
	}
}

// GetMappingListTemplatesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTemplatesSortByEnum(val string) (ListTemplatesSortByEnum, bool) {
	enum, ok := mappingListTemplatesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
