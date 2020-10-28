// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v28/common"
	"net/http"
)

// ListParsersRequest wrapper for the ListParsers operation
type ListParsersRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// isMatchAll
	IsMatchAll *bool `mandatory:"false" contributesTo:"query" name:"isMatchAll"`

	// source type
	SourceType ListParsersSourceTypeEnum `mandatory:"false" contributesTo:"query" name:"sourceType" omitEmpty:"true"`

	// parserName
	ParserName *string `mandatory:"false" contributesTo:"query" name:"parserName"`

	// search by parser display name or description
	ParserDisplayText *string `mandatory:"false" contributesTo:"query" name:"parserDisplayText"`

	// parserType
	ParserType ListParsersParserTypeEnum `mandatory:"false" contributesTo:"query" name:"parserType" omitEmpty:"true"`

	// Is system param of value (all, custom, sourceUsing)
	IsSystem ListParsersIsSystemEnum `mandatory:"false" contributesTo:"query" name:"isSystem" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListParsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// sort by parser
	SortBy ListParsersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListParsersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListParsersRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListParsersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListParsersResponse wrapper for the ListParsers operation
type ListParsersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsParserCollection instances
	LogAnalyticsParserCollection `presentIn:"body"`

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

func (response ListParsersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListParsersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListParsersSourceTypeEnum Enum with underlying type: string
type ListParsersSourceTypeEnum string

// Set of constants representing the allowable values for ListParsersSourceTypeEnum
const (
	ListParsersSourceTypeOsFile       ListParsersSourceTypeEnum = "OS_FILE"
	ListParsersSourceTypeSyslog       ListParsersSourceTypeEnum = "SYSLOG"
	ListParsersSourceTypeOdl          ListParsersSourceTypeEnum = "ODL"
	ListParsersSourceTypeOsWindowsSys ListParsersSourceTypeEnum = "OS_WINDOWS_SYS"
)

var mappingListParsersSourceType = map[string]ListParsersSourceTypeEnum{
	"OS_FILE":        ListParsersSourceTypeOsFile,
	"SYSLOG":         ListParsersSourceTypeSyslog,
	"ODL":            ListParsersSourceTypeOdl,
	"OS_WINDOWS_SYS": ListParsersSourceTypeOsWindowsSys,
}

// GetListParsersSourceTypeEnumValues Enumerates the set of values for ListParsersSourceTypeEnum
func GetListParsersSourceTypeEnumValues() []ListParsersSourceTypeEnum {
	values := make([]ListParsersSourceTypeEnum, 0)
	for _, v := range mappingListParsersSourceType {
		values = append(values, v)
	}
	return values
}

// ListParsersParserTypeEnum Enum with underlying type: string
type ListParsersParserTypeEnum string

// Set of constants representing the allowable values for ListParsersParserTypeEnum
const (
	ListParsersParserTypeAll   ListParsersParserTypeEnum = "ALL"
	ListParsersParserTypeRegex ListParsersParserTypeEnum = "REGEX"
	ListParsersParserTypeXml   ListParsersParserTypeEnum = "XML"
	ListParsersParserTypeJson  ListParsersParserTypeEnum = "JSON"
)

var mappingListParsersParserType = map[string]ListParsersParserTypeEnum{
	"ALL":   ListParsersParserTypeAll,
	"REGEX": ListParsersParserTypeRegex,
	"XML":   ListParsersParserTypeXml,
	"JSON":  ListParsersParserTypeJson,
}

// GetListParsersParserTypeEnumValues Enumerates the set of values for ListParsersParserTypeEnum
func GetListParsersParserTypeEnumValues() []ListParsersParserTypeEnum {
	values := make([]ListParsersParserTypeEnum, 0)
	for _, v := range mappingListParsersParserType {
		values = append(values, v)
	}
	return values
}

// ListParsersIsSystemEnum Enum with underlying type: string
type ListParsersIsSystemEnum string

// Set of constants representing the allowable values for ListParsersIsSystemEnum
const (
	ListParsersIsSystemAll     ListParsersIsSystemEnum = "ALL"
	ListParsersIsSystemCustom  ListParsersIsSystemEnum = "CUSTOM"
	ListParsersIsSystemBuiltIn ListParsersIsSystemEnum = "BUILT_IN"
)

var mappingListParsersIsSystem = map[string]ListParsersIsSystemEnum{
	"ALL":      ListParsersIsSystemAll,
	"CUSTOM":   ListParsersIsSystemCustom,
	"BUILT_IN": ListParsersIsSystemBuiltIn,
}

// GetListParsersIsSystemEnumValues Enumerates the set of values for ListParsersIsSystemEnum
func GetListParsersIsSystemEnumValues() []ListParsersIsSystemEnum {
	values := make([]ListParsersIsSystemEnum, 0)
	for _, v := range mappingListParsersIsSystem {
		values = append(values, v)
	}
	return values
}

// ListParsersSortOrderEnum Enum with underlying type: string
type ListParsersSortOrderEnum string

// Set of constants representing the allowable values for ListParsersSortOrderEnum
const (
	ListParsersSortOrderAsc  ListParsersSortOrderEnum = "ASC"
	ListParsersSortOrderDesc ListParsersSortOrderEnum = "DESC"
)

var mappingListParsersSortOrder = map[string]ListParsersSortOrderEnum{
	"ASC":  ListParsersSortOrderAsc,
	"DESC": ListParsersSortOrderDesc,
}

// GetListParsersSortOrderEnumValues Enumerates the set of values for ListParsersSortOrderEnum
func GetListParsersSortOrderEnumValues() []ListParsersSortOrderEnum {
	values := make([]ListParsersSortOrderEnum, 0)
	for _, v := range mappingListParsersSortOrder {
		values = append(values, v)
	}
	return values
}

// ListParsersSortByEnum Enum with underlying type: string
type ListParsersSortByEnum string

// Set of constants representing the allowable values for ListParsersSortByEnum
const (
	ListParsersSortByName         ListParsersSortByEnum = "name"
	ListParsersSortByType         ListParsersSortByEnum = "type"
	ListParsersSortBySourcescount ListParsersSortByEnum = "sourcesCount"
	ListParsersSortByTimeupdated  ListParsersSortByEnum = "timeUpdated"
)

var mappingListParsersSortBy = map[string]ListParsersSortByEnum{
	"name":         ListParsersSortByName,
	"type":         ListParsersSortByType,
	"sourcesCount": ListParsersSortBySourcescount,
	"timeUpdated":  ListParsersSortByTimeupdated,
}

// GetListParsersSortByEnumValues Enumerates the set of values for ListParsersSortByEnum
func GetListParsersSortByEnumValues() []ListParsersSortByEnum {
	values := make([]ListParsersSortByEnum, 0)
	for _, v := range mappingListParsersSortBy {
		values = append(values, v)
	}
	return values
}
