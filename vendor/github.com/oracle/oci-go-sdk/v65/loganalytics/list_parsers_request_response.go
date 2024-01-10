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

// ListParsersRequest wrapper for the ListParsers operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListParsers.go.html to see an example of how to use ListParsersRequest.
type ListParsersRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// A flag indicating how to handle filtering when multiple filter criteria are specified.
	// A value of true will always result in the most expansive list of items being returned.
	// For example, if two field lists are supplies as filter criteria, a value of true will
	// result in any item matching a field in either list being returned, while a value of
	// false will result in a list of items which only have fields contained in both input lists.
	IsMatchAll *bool `mandatory:"false" contributesTo:"query" name:"isMatchAll"`

	// The source type used for filtering.  Only parsers associated with a source of the
	// specified type will be returned.
	SourceType ListParsersSourceTypeEnum `mandatory:"false" contributesTo:"query" name:"sourceType" omitEmpty:"true"`

	// The parser name used for filtering.
	ParserName *string `mandatory:"false" contributesTo:"query" name:"parserName"`

	// The parser display text used for filtering.  Only parsers with the specified name
	// or description will be returned.
	ParserDisplayText *string `mandatory:"false" contributesTo:"query" name:"parserDisplayText"`

	// The parser type used for filtering.  Only items with, or associated with, parsers
	// of the specified type will be returned.
	ParserType ListParsersParserTypeEnum `mandatory:"false" contributesTo:"query" name:"parserType" omitEmpty:"true"`

	// A comma-separated list of categories used for filtering
	Categories *string `mandatory:"false" contributesTo:"query" name:"categories"`

	// The system value used for filtering.  Only items with the specified system value
	// will be returned.  Valid values are built in, custom (for user defined items), or
	// all (for all items, regardless of system value).
	IsSystem ListParsersIsSystemEnum `mandatory:"false" contributesTo:"query" name:"isSystem" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListParsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned parsers
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
func (request ListParsersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListParsersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListParsersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListParsersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListParsersSourceTypeEnum(string(request.SourceType)); !ok && request.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", request.SourceType, strings.Join(GetListParsersSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListParsersParserTypeEnum(string(request.ParserType)); !ok && request.ParserType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ParserType: %s. Supported values are: %s.", request.ParserType, strings.Join(GetListParsersParserTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListParsersIsSystemEnum(string(request.IsSystem)); !ok && request.IsSystem != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsSystem: %s. Supported values are: %s.", request.IsSystem, strings.Join(GetListParsersIsSystemEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListParsersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListParsersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListParsersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListParsersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListParsersSourceTypeEnum = map[string]ListParsersSourceTypeEnum{
	"OS_FILE":        ListParsersSourceTypeOsFile,
	"SYSLOG":         ListParsersSourceTypeSyslog,
	"ODL":            ListParsersSourceTypeOdl,
	"OS_WINDOWS_SYS": ListParsersSourceTypeOsWindowsSys,
}

var mappingListParsersSourceTypeEnumLowerCase = map[string]ListParsersSourceTypeEnum{
	"os_file":        ListParsersSourceTypeOsFile,
	"syslog":         ListParsersSourceTypeSyslog,
	"odl":            ListParsersSourceTypeOdl,
	"os_windows_sys": ListParsersSourceTypeOsWindowsSys,
}

// GetListParsersSourceTypeEnumValues Enumerates the set of values for ListParsersSourceTypeEnum
func GetListParsersSourceTypeEnumValues() []ListParsersSourceTypeEnum {
	values := make([]ListParsersSourceTypeEnum, 0)
	for _, v := range mappingListParsersSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListParsersSourceTypeEnumStringValues Enumerates the set of values in String for ListParsersSourceTypeEnum
func GetListParsersSourceTypeEnumStringValues() []string {
	return []string{
		"OS_FILE",
		"SYSLOG",
		"ODL",
		"OS_WINDOWS_SYS",
	}
}

// GetMappingListParsersSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListParsersSourceTypeEnum(val string) (ListParsersSourceTypeEnum, bool) {
	enum, ok := mappingListParsersSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListParsersParserTypeEnum Enum with underlying type: string
type ListParsersParserTypeEnum string

// Set of constants representing the allowable values for ListParsersParserTypeEnum
const (
	ListParsersParserTypeAll       ListParsersParserTypeEnum = "ALL"
	ListParsersParserTypeRegex     ListParsersParserTypeEnum = "REGEX"
	ListParsersParserTypeXml       ListParsersParserTypeEnum = "XML"
	ListParsersParserTypeJson      ListParsersParserTypeEnum = "JSON"
	ListParsersParserTypeOdl       ListParsersParserTypeEnum = "ODL"
	ListParsersParserTypeDelimited ListParsersParserTypeEnum = "DELIMITED"
)

var mappingListParsersParserTypeEnum = map[string]ListParsersParserTypeEnum{
	"ALL":       ListParsersParserTypeAll,
	"REGEX":     ListParsersParserTypeRegex,
	"XML":       ListParsersParserTypeXml,
	"JSON":      ListParsersParserTypeJson,
	"ODL":       ListParsersParserTypeOdl,
	"DELIMITED": ListParsersParserTypeDelimited,
}

var mappingListParsersParserTypeEnumLowerCase = map[string]ListParsersParserTypeEnum{
	"all":       ListParsersParserTypeAll,
	"regex":     ListParsersParserTypeRegex,
	"xml":       ListParsersParserTypeXml,
	"json":      ListParsersParserTypeJson,
	"odl":       ListParsersParserTypeOdl,
	"delimited": ListParsersParserTypeDelimited,
}

// GetListParsersParserTypeEnumValues Enumerates the set of values for ListParsersParserTypeEnum
func GetListParsersParserTypeEnumValues() []ListParsersParserTypeEnum {
	values := make([]ListParsersParserTypeEnum, 0)
	for _, v := range mappingListParsersParserTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListParsersParserTypeEnumStringValues Enumerates the set of values in String for ListParsersParserTypeEnum
func GetListParsersParserTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"REGEX",
		"XML",
		"JSON",
		"ODL",
		"DELIMITED",
	}
}

// GetMappingListParsersParserTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListParsersParserTypeEnum(val string) (ListParsersParserTypeEnum, bool) {
	enum, ok := mappingListParsersParserTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListParsersIsSystemEnum Enum with underlying type: string
type ListParsersIsSystemEnum string

// Set of constants representing the allowable values for ListParsersIsSystemEnum
const (
	ListParsersIsSystemAll     ListParsersIsSystemEnum = "ALL"
	ListParsersIsSystemCustom  ListParsersIsSystemEnum = "CUSTOM"
	ListParsersIsSystemBuiltIn ListParsersIsSystemEnum = "BUILT_IN"
)

var mappingListParsersIsSystemEnum = map[string]ListParsersIsSystemEnum{
	"ALL":      ListParsersIsSystemAll,
	"CUSTOM":   ListParsersIsSystemCustom,
	"BUILT_IN": ListParsersIsSystemBuiltIn,
}

var mappingListParsersIsSystemEnumLowerCase = map[string]ListParsersIsSystemEnum{
	"all":      ListParsersIsSystemAll,
	"custom":   ListParsersIsSystemCustom,
	"built_in": ListParsersIsSystemBuiltIn,
}

// GetListParsersIsSystemEnumValues Enumerates the set of values for ListParsersIsSystemEnum
func GetListParsersIsSystemEnumValues() []ListParsersIsSystemEnum {
	values := make([]ListParsersIsSystemEnum, 0)
	for _, v := range mappingListParsersIsSystemEnum {
		values = append(values, v)
	}
	return values
}

// GetListParsersIsSystemEnumStringValues Enumerates the set of values in String for ListParsersIsSystemEnum
func GetListParsersIsSystemEnumStringValues() []string {
	return []string{
		"ALL",
		"CUSTOM",
		"BUILT_IN",
	}
}

// GetMappingListParsersIsSystemEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListParsersIsSystemEnum(val string) (ListParsersIsSystemEnum, bool) {
	enum, ok := mappingListParsersIsSystemEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListParsersSortOrderEnum Enum with underlying type: string
type ListParsersSortOrderEnum string

// Set of constants representing the allowable values for ListParsersSortOrderEnum
const (
	ListParsersSortOrderAsc  ListParsersSortOrderEnum = "ASC"
	ListParsersSortOrderDesc ListParsersSortOrderEnum = "DESC"
)

var mappingListParsersSortOrderEnum = map[string]ListParsersSortOrderEnum{
	"ASC":  ListParsersSortOrderAsc,
	"DESC": ListParsersSortOrderDesc,
}

var mappingListParsersSortOrderEnumLowerCase = map[string]ListParsersSortOrderEnum{
	"asc":  ListParsersSortOrderAsc,
	"desc": ListParsersSortOrderDesc,
}

// GetListParsersSortOrderEnumValues Enumerates the set of values for ListParsersSortOrderEnum
func GetListParsersSortOrderEnumValues() []ListParsersSortOrderEnum {
	values := make([]ListParsersSortOrderEnum, 0)
	for _, v := range mappingListParsersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListParsersSortOrderEnumStringValues Enumerates the set of values in String for ListParsersSortOrderEnum
func GetListParsersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListParsersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListParsersSortOrderEnum(val string) (ListParsersSortOrderEnum, bool) {
	enum, ok := mappingListParsersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingListParsersSortByEnum = map[string]ListParsersSortByEnum{
	"name":         ListParsersSortByName,
	"type":         ListParsersSortByType,
	"sourcesCount": ListParsersSortBySourcescount,
	"timeUpdated":  ListParsersSortByTimeupdated,
}

var mappingListParsersSortByEnumLowerCase = map[string]ListParsersSortByEnum{
	"name":         ListParsersSortByName,
	"type":         ListParsersSortByType,
	"sourcescount": ListParsersSortBySourcescount,
	"timeupdated":  ListParsersSortByTimeupdated,
}

// GetListParsersSortByEnumValues Enumerates the set of values for ListParsersSortByEnum
func GetListParsersSortByEnumValues() []ListParsersSortByEnum {
	values := make([]ListParsersSortByEnum, 0)
	for _, v := range mappingListParsersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListParsersSortByEnumStringValues Enumerates the set of values in String for ListParsersSortByEnum
func GetListParsersSortByEnumStringValues() []string {
	return []string{
		"name",
		"type",
		"sourcesCount",
		"timeUpdated",
	}
}

// GetMappingListParsersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListParsersSortByEnum(val string) (ListParsersSortByEnum, bool) {
	enum, ok := mappingListParsersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
