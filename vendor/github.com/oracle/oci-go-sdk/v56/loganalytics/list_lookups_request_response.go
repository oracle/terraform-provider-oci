// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListLookupsRequest wrapper for the ListLookups operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLookups.go.html to see an example of how to use ListLookupsRequest.
type ListLookupsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The lookup type.  Valid values are Lookup or Dictionary.
	Type ListLookupsTypeEnum `mandatory:"true" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The lookup text used for filtering.  Only lookups with the specified name
	// or description will be returned.
	LookupDisplayText *string `mandatory:"false" contributesTo:"query" name:"lookupDisplayText"`

	// The system value used for filtering.  Only items with the specified system value
	// will be returned.  Valid values are built in, custom (for user defined items), or
	// all (for all items, regardless of system value).
	IsSystem ListLookupsIsSystemEnum `mandatory:"false" contributesTo:"query" name:"isSystem" omitEmpty:"true"`

	// sort by field
	SortBy ListLookupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The lookup status used for filtering when fetching a list of lookups.
	Status ListLookupsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// A comma-separated list of categories used for filtering
	Categories *string `mandatory:"false" contributesTo:"query" name:"categories"`

	// A flag indicating whether or not to return OMC annotated or hidden lookups.
	IsHideSpecial *bool `mandatory:"false" contributesTo:"query" name:"isHideSpecial"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLookupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLookupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLookupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLookupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLookupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLookupsResponse wrapper for the ListLookups operation
type ListLookupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsLookupCollection instances
	LogAnalyticsLookupCollection `presentIn:"body"`

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

func (response ListLookupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLookupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLookupsTypeEnum Enum with underlying type: string
type ListLookupsTypeEnum string

// Set of constants representing the allowable values for ListLookupsTypeEnum
const (
	ListLookupsTypeLookup     ListLookupsTypeEnum = "Lookup"
	ListLookupsTypeDictionary ListLookupsTypeEnum = "Dictionary"
)

var mappingListLookupsType = map[string]ListLookupsTypeEnum{
	"Lookup":     ListLookupsTypeLookup,
	"Dictionary": ListLookupsTypeDictionary,
}

// GetListLookupsTypeEnumValues Enumerates the set of values for ListLookupsTypeEnum
func GetListLookupsTypeEnumValues() []ListLookupsTypeEnum {
	values := make([]ListLookupsTypeEnum, 0)
	for _, v := range mappingListLookupsType {
		values = append(values, v)
	}
	return values
}

// ListLookupsIsSystemEnum Enum with underlying type: string
type ListLookupsIsSystemEnum string

// Set of constants representing the allowable values for ListLookupsIsSystemEnum
const (
	ListLookupsIsSystemAll     ListLookupsIsSystemEnum = "ALL"
	ListLookupsIsSystemCustom  ListLookupsIsSystemEnum = "CUSTOM"
	ListLookupsIsSystemBuiltIn ListLookupsIsSystemEnum = "BUILT_IN"
)

var mappingListLookupsIsSystem = map[string]ListLookupsIsSystemEnum{
	"ALL":      ListLookupsIsSystemAll,
	"CUSTOM":   ListLookupsIsSystemCustom,
	"BUILT_IN": ListLookupsIsSystemBuiltIn,
}

// GetListLookupsIsSystemEnumValues Enumerates the set of values for ListLookupsIsSystemEnum
func GetListLookupsIsSystemEnumValues() []ListLookupsIsSystemEnum {
	values := make([]ListLookupsIsSystemEnum, 0)
	for _, v := range mappingListLookupsIsSystem {
		values = append(values, v)
	}
	return values
}

// ListLookupsSortByEnum Enum with underlying type: string
type ListLookupsSortByEnum string

// Set of constants representing the allowable values for ListLookupsSortByEnum
const (
	ListLookupsSortByDisplayname  ListLookupsSortByEnum = "displayName"
	ListLookupsSortByStatus       ListLookupsSortByEnum = "status"
	ListLookupsSortByType         ListLookupsSortByEnum = "type"
	ListLookupsSortByUpdatedtime  ListLookupsSortByEnum = "updatedTime"
	ListLookupsSortByCreationtype ListLookupsSortByEnum = "creationType"
)

var mappingListLookupsSortBy = map[string]ListLookupsSortByEnum{
	"displayName":  ListLookupsSortByDisplayname,
	"status":       ListLookupsSortByStatus,
	"type":         ListLookupsSortByType,
	"updatedTime":  ListLookupsSortByUpdatedtime,
	"creationType": ListLookupsSortByCreationtype,
}

// GetListLookupsSortByEnumValues Enumerates the set of values for ListLookupsSortByEnum
func GetListLookupsSortByEnumValues() []ListLookupsSortByEnum {
	values := make([]ListLookupsSortByEnum, 0)
	for _, v := range mappingListLookupsSortBy {
		values = append(values, v)
	}
	return values
}

// ListLookupsStatusEnum Enum with underlying type: string
type ListLookupsStatusEnum string

// Set of constants representing the allowable values for ListLookupsStatusEnum
const (
	ListLookupsStatusAll        ListLookupsStatusEnum = "ALL"
	ListLookupsStatusSuccessful ListLookupsStatusEnum = "SUCCESSFUL"
	ListLookupsStatusFailed     ListLookupsStatusEnum = "FAILED"
	ListLookupsStatusInprogress ListLookupsStatusEnum = "INPROGRESS"
)

var mappingListLookupsStatus = map[string]ListLookupsStatusEnum{
	"ALL":        ListLookupsStatusAll,
	"SUCCESSFUL": ListLookupsStatusSuccessful,
	"FAILED":     ListLookupsStatusFailed,
	"INPROGRESS": ListLookupsStatusInprogress,
}

// GetListLookupsStatusEnumValues Enumerates the set of values for ListLookupsStatusEnum
func GetListLookupsStatusEnumValues() []ListLookupsStatusEnum {
	values := make([]ListLookupsStatusEnum, 0)
	for _, v := range mappingListLookupsStatus {
		values = append(values, v)
	}
	return values
}

// ListLookupsSortOrderEnum Enum with underlying type: string
type ListLookupsSortOrderEnum string

// Set of constants representing the allowable values for ListLookupsSortOrderEnum
const (
	ListLookupsSortOrderAsc  ListLookupsSortOrderEnum = "ASC"
	ListLookupsSortOrderDesc ListLookupsSortOrderEnum = "DESC"
)

var mappingListLookupsSortOrder = map[string]ListLookupsSortOrderEnum{
	"ASC":  ListLookupsSortOrderAsc,
	"DESC": ListLookupsSortOrderDesc,
}

// GetListLookupsSortOrderEnumValues Enumerates the set of values for ListLookupsSortOrderEnum
func GetListLookupsSortOrderEnumValues() []ListLookupsSortOrderEnum {
	values := make([]ListLookupsSortOrderEnum, 0)
	for _, v := range mappingListLookupsSortOrder {
		values = append(values, v)
	}
	return values
}
