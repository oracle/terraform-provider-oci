// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSourcesRequest wrapper for the ListSources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListSources.go.html to see an example of how to use ListSourcesRequest.
type ListSourcesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only sources associated with entities of the specified type.
	// The match is case-insensitive.
	EntityType *string `mandatory:"false" contributesTo:"query" name:"entityType"`

	// The source display text used for filtering.  Only sources with the specified name
	// or description will be returned.
	SourceDisplayText *string `mandatory:"false" contributesTo:"query" name:"sourceDisplayText"`

	// The system value used for filtering.  Only items with the specified system value
	// will be returned.  Valid values are built in, custom (for user defined items), or
	// all (for all items, regardless of system value).
	IsSystem ListSourcesIsSystemEnum `mandatory:"false" contributesTo:"query" name:"isSystem" omitEmpty:"true"`

	// An auto-associate flag used for filtering.  Only sources which are marked for automatic
	// association will be returned.
	IsAutoAssociated *bool `mandatory:"false" contributesTo:"query" name:"isAutoAssociated"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned sources
	SortBy ListSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only log analytics entities whose name matches the entire name given. The match
	// is case-insensitive.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A comma-separated list of categories used for filtering
	Categories *string `mandatory:"false" contributesTo:"query" name:"categories"`

	// A flag specifying whether or not to return all source information, or a subset of the
	// information about each source.  A value of true will return only the source unique
	// identifier and the source name.  A value of false will return all source information
	// (such as author, updated date, system flag, etc.)
	IsSimplified *bool `mandatory:"false" contributesTo:"query" name:"isSimplified"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSourcesIsSystemEnum(string(request.IsSystem)); !ok && request.IsSystem != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsSystem: %s. Supported values are: %s.", request.IsSystem, strings.Join(GetListSourcesIsSystemEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSourcesResponse wrapper for the ListSources operation
type ListSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsSourceCollection instances
	LogAnalyticsSourceCollection `presentIn:"body"`

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

func (response ListSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSourcesIsSystemEnum Enum with underlying type: string
type ListSourcesIsSystemEnum string

// Set of constants representing the allowable values for ListSourcesIsSystemEnum
const (
	ListSourcesIsSystemAll     ListSourcesIsSystemEnum = "ALL"
	ListSourcesIsSystemCustom  ListSourcesIsSystemEnum = "CUSTOM"
	ListSourcesIsSystemBuiltIn ListSourcesIsSystemEnum = "BUILT_IN"
)

var mappingListSourcesIsSystemEnum = map[string]ListSourcesIsSystemEnum{
	"ALL":      ListSourcesIsSystemAll,
	"CUSTOM":   ListSourcesIsSystemCustom,
	"BUILT_IN": ListSourcesIsSystemBuiltIn,
}

var mappingListSourcesIsSystemEnumLowerCase = map[string]ListSourcesIsSystemEnum{
	"all":      ListSourcesIsSystemAll,
	"custom":   ListSourcesIsSystemCustom,
	"built_in": ListSourcesIsSystemBuiltIn,
}

// GetListSourcesIsSystemEnumValues Enumerates the set of values for ListSourcesIsSystemEnum
func GetListSourcesIsSystemEnumValues() []ListSourcesIsSystemEnum {
	values := make([]ListSourcesIsSystemEnum, 0)
	for _, v := range mappingListSourcesIsSystemEnum {
		values = append(values, v)
	}
	return values
}

// GetListSourcesIsSystemEnumStringValues Enumerates the set of values in String for ListSourcesIsSystemEnum
func GetListSourcesIsSystemEnumStringValues() []string {
	return []string{
		"ALL",
		"CUSTOM",
		"BUILT_IN",
	}
}

// GetMappingListSourcesIsSystemEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSourcesIsSystemEnum(val string) (ListSourcesIsSystemEnum, bool) {
	enum, ok := mappingListSourcesIsSystemEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSourcesSortOrderEnum Enum with underlying type: string
type ListSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListSourcesSortOrderEnum
const (
	ListSourcesSortOrderAsc  ListSourcesSortOrderEnum = "ASC"
	ListSourcesSortOrderDesc ListSourcesSortOrderEnum = "DESC"
)

var mappingListSourcesSortOrderEnum = map[string]ListSourcesSortOrderEnum{
	"ASC":  ListSourcesSortOrderAsc,
	"DESC": ListSourcesSortOrderDesc,
}

var mappingListSourcesSortOrderEnumLowerCase = map[string]ListSourcesSortOrderEnum{
	"asc":  ListSourcesSortOrderAsc,
	"desc": ListSourcesSortOrderDesc,
}

// GetListSourcesSortOrderEnumValues Enumerates the set of values for ListSourcesSortOrderEnum
func GetListSourcesSortOrderEnumValues() []ListSourcesSortOrderEnum {
	values := make([]ListSourcesSortOrderEnum, 0)
	for _, v := range mappingListSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListSourcesSortOrderEnum
func GetListSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSourcesSortOrderEnum(val string) (ListSourcesSortOrderEnum, bool) {
	enum, ok := mappingListSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSourcesSortByEnum Enum with underlying type: string
type ListSourcesSortByEnum string

// Set of constants representing the allowable values for ListSourcesSortByEnum
const (
	ListSourcesSortByName             ListSourcesSortByEnum = "name"
	ListSourcesSortByTimeupdated      ListSourcesSortByEnum = "timeUpdated"
	ListSourcesSortByAssociationcount ListSourcesSortByEnum = "associationCount"
	ListSourcesSortBySourcetype       ListSourcesSortByEnum = "sourceType"
)

var mappingListSourcesSortByEnum = map[string]ListSourcesSortByEnum{
	"name":             ListSourcesSortByName,
	"timeUpdated":      ListSourcesSortByTimeupdated,
	"associationCount": ListSourcesSortByAssociationcount,
	"sourceType":       ListSourcesSortBySourcetype,
}

var mappingListSourcesSortByEnumLowerCase = map[string]ListSourcesSortByEnum{
	"name":             ListSourcesSortByName,
	"timeupdated":      ListSourcesSortByTimeupdated,
	"associationcount": ListSourcesSortByAssociationcount,
	"sourcetype":       ListSourcesSortBySourcetype,
}

// GetListSourcesSortByEnumValues Enumerates the set of values for ListSourcesSortByEnum
func GetListSourcesSortByEnumValues() []ListSourcesSortByEnum {
	values := make([]ListSourcesSortByEnum, 0)
	for _, v := range mappingListSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSourcesSortByEnumStringValues Enumerates the set of values in String for ListSourcesSortByEnum
func GetListSourcesSortByEnumStringValues() []string {
	return []string{
		"name",
		"timeUpdated",
		"associationCount",
		"sourceType",
	}
}

// GetMappingListSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSourcesSortByEnum(val string) (ListSourcesSortByEnum, bool) {
	enum, ok := mappingListSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
