// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListFieldsRequest wrapper for the ListFields operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListFields.go.html to see an example of how to use ListFieldsRequest.
type ListFieldsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// A flag indicating how to handle filtering when multiple filter criteria are specified.
	// A value of true will always result in the most expansive list of items being returned.
	// For example, if two field lists are supplies as filter criteria, a value of true will
	// result in any item matching a field in either list being returned, while a value of
	// false will result in a list of items which only have fields contained in both input lists.
	IsMatchAll *bool `mandatory:"false" contributesTo:"query" name:"isMatchAll"`

	// A list of source IDs used for filtering.  Only fields used by the specified
	// sources will be returned.
	SourceIds *string `mandatory:"false" contributesTo:"query" name:"sourceIds"`

	// A list of source names used for filtering.  Only fields used by the specified
	// sources will be returned.
	SourceNames *string `mandatory:"false" contributesTo:"query" name:"sourceNames"`

	// The parser type used for filtering.  Only items with, or associated with, parsers
	// of the specified type will be returned.
	ParserType ListFieldsParserTypeEnum `mandatory:"false" contributesTo:"query" name:"parserType" omitEmpty:"true"`

	// A list of parser names used for filtering.  Only fields used by the specified
	// parsers will be returned.
	ParserIds *string `mandatory:"false" contributesTo:"query" name:"parserIds"`

	// A list of parser names used for filtering.  Only fields used by the specified
	// parsers will be returned.
	ParserNames *string `mandatory:"false" contributesTo:"query" name:"parserNames"`

	// isIncludeParser
	IsIncludeParser *bool `mandatory:"false" contributesTo:"query" name:"isIncludeParser"`

	// filter
	Filter *string `mandatory:"false" contributesTo:"query" name:"filter"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListFieldsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned fields
	SortBy ListFieldsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFieldsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFieldsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFieldsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFieldsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFieldsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFieldsParserTypeEnum(string(request.ParserType)); !ok && request.ParserType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ParserType: %s. Supported values are: %s.", request.ParserType, strings.Join(GetListFieldsParserTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFieldsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFieldsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFieldsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFieldsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFieldsResponse wrapper for the ListFields operation
type ListFieldsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsFieldCollection instances
	LogAnalyticsFieldCollection `presentIn:"body"`

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

func (response ListFieldsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFieldsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFieldsParserTypeEnum Enum with underlying type: string
type ListFieldsParserTypeEnum string

// Set of constants representing the allowable values for ListFieldsParserTypeEnum
const (
	ListFieldsParserTypeAll       ListFieldsParserTypeEnum = "ALL"
	ListFieldsParserTypeRegex     ListFieldsParserTypeEnum = "REGEX"
	ListFieldsParserTypeXml       ListFieldsParserTypeEnum = "XML"
	ListFieldsParserTypeJson      ListFieldsParserTypeEnum = "JSON"
	ListFieldsParserTypeOdl       ListFieldsParserTypeEnum = "ODL"
	ListFieldsParserTypeDelimited ListFieldsParserTypeEnum = "DELIMITED"
)

var mappingListFieldsParserTypeEnum = map[string]ListFieldsParserTypeEnum{
	"ALL":       ListFieldsParserTypeAll,
	"REGEX":     ListFieldsParserTypeRegex,
	"XML":       ListFieldsParserTypeXml,
	"JSON":      ListFieldsParserTypeJson,
	"ODL":       ListFieldsParserTypeOdl,
	"DELIMITED": ListFieldsParserTypeDelimited,
}

// GetListFieldsParserTypeEnumValues Enumerates the set of values for ListFieldsParserTypeEnum
func GetListFieldsParserTypeEnumValues() []ListFieldsParserTypeEnum {
	values := make([]ListFieldsParserTypeEnum, 0)
	for _, v := range mappingListFieldsParserTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListFieldsParserTypeEnumStringValues Enumerates the set of values in String for ListFieldsParserTypeEnum
func GetListFieldsParserTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"REGEX",
		"XML",
		"JSON",
		"ODL",
		"DELIMITED",
	}
}

// GetMappingListFieldsParserTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFieldsParserTypeEnum(val string) (ListFieldsParserTypeEnum, bool) {
	mappingListFieldsParserTypeEnumIgnoreCase := make(map[string]ListFieldsParserTypeEnum)
	for k, v := range mappingListFieldsParserTypeEnum {
		mappingListFieldsParserTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListFieldsParserTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListFieldsSortOrderEnum Enum with underlying type: string
type ListFieldsSortOrderEnum string

// Set of constants representing the allowable values for ListFieldsSortOrderEnum
const (
	ListFieldsSortOrderAsc  ListFieldsSortOrderEnum = "ASC"
	ListFieldsSortOrderDesc ListFieldsSortOrderEnum = "DESC"
)

var mappingListFieldsSortOrderEnum = map[string]ListFieldsSortOrderEnum{
	"ASC":  ListFieldsSortOrderAsc,
	"DESC": ListFieldsSortOrderDesc,
}

// GetListFieldsSortOrderEnumValues Enumerates the set of values for ListFieldsSortOrderEnum
func GetListFieldsSortOrderEnumValues() []ListFieldsSortOrderEnum {
	values := make([]ListFieldsSortOrderEnum, 0)
	for _, v := range mappingListFieldsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFieldsSortOrderEnumStringValues Enumerates the set of values in String for ListFieldsSortOrderEnum
func GetListFieldsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFieldsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFieldsSortOrderEnum(val string) (ListFieldsSortOrderEnum, bool) {
	mappingListFieldsSortOrderEnumIgnoreCase := make(map[string]ListFieldsSortOrderEnum)
	for k, v := range mappingListFieldsSortOrderEnum {
		mappingListFieldsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListFieldsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListFieldsSortByEnum Enum with underlying type: string
type ListFieldsSortByEnum string

// Set of constants representing the allowable values for ListFieldsSortByEnum
const (
	ListFieldsSortByName     ListFieldsSortByEnum = "name"
	ListFieldsSortByDatatype ListFieldsSortByEnum = "dataType"
)

var mappingListFieldsSortByEnum = map[string]ListFieldsSortByEnum{
	"name":     ListFieldsSortByName,
	"dataType": ListFieldsSortByDatatype,
}

// GetListFieldsSortByEnumValues Enumerates the set of values for ListFieldsSortByEnum
func GetListFieldsSortByEnumValues() []ListFieldsSortByEnum {
	values := make([]ListFieldsSortByEnum, 0)
	for _, v := range mappingListFieldsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFieldsSortByEnumStringValues Enumerates the set of values in String for ListFieldsSortByEnum
func GetListFieldsSortByEnumStringValues() []string {
	return []string{
		"name",
		"dataType",
	}
}

// GetMappingListFieldsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFieldsSortByEnum(val string) (ListFieldsSortByEnum, bool) {
	mappingListFieldsSortByEnumIgnoreCase := make(map[string]ListFieldsSortByEnum)
	for k, v := range mappingListFieldsSortByEnum {
		mappingListFieldsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListFieldsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
