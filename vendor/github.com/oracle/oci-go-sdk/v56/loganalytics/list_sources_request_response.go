// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListSourcesRequest wrapper for the ListSources operation
//
// See also
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

var mappingListSourcesIsSystem = map[string]ListSourcesIsSystemEnum{
	"ALL":      ListSourcesIsSystemAll,
	"CUSTOM":   ListSourcesIsSystemCustom,
	"BUILT_IN": ListSourcesIsSystemBuiltIn,
}

// GetListSourcesIsSystemEnumValues Enumerates the set of values for ListSourcesIsSystemEnum
func GetListSourcesIsSystemEnumValues() []ListSourcesIsSystemEnum {
	values := make([]ListSourcesIsSystemEnum, 0)
	for _, v := range mappingListSourcesIsSystem {
		values = append(values, v)
	}
	return values
}

// ListSourcesSortOrderEnum Enum with underlying type: string
type ListSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListSourcesSortOrderEnum
const (
	ListSourcesSortOrderAsc  ListSourcesSortOrderEnum = "ASC"
	ListSourcesSortOrderDesc ListSourcesSortOrderEnum = "DESC"
)

var mappingListSourcesSortOrder = map[string]ListSourcesSortOrderEnum{
	"ASC":  ListSourcesSortOrderAsc,
	"DESC": ListSourcesSortOrderDesc,
}

// GetListSourcesSortOrderEnumValues Enumerates the set of values for ListSourcesSortOrderEnum
func GetListSourcesSortOrderEnumValues() []ListSourcesSortOrderEnum {
	values := make([]ListSourcesSortOrderEnum, 0)
	for _, v := range mappingListSourcesSortOrder {
		values = append(values, v)
	}
	return values
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

var mappingListSourcesSortBy = map[string]ListSourcesSortByEnum{
	"name":             ListSourcesSortByName,
	"timeUpdated":      ListSourcesSortByTimeupdated,
	"associationCount": ListSourcesSortByAssociationcount,
	"sourceType":       ListSourcesSortBySourcetype,
}

// GetListSourcesSortByEnumValues Enumerates the set of values for ListSourcesSortByEnum
func GetListSourcesSortByEnumValues() []ListSourcesSortByEnum {
	values := make([]ListSourcesSortByEnum, 0)
	for _, v := range mappingListSourcesSortBy {
		values = append(values, v)
	}
	return values
}
