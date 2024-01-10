// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDiscoveryJobLogsRequest wrapper for the ListDiscoveryJobLogs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListDiscoveryJobLogs.go.html to see an example of how to use ListDiscoveryJobLogsRequest.
type ListDiscoveryJobLogsRequest struct {

	// The Discovery Job ID
	DiscoveryJobId *string `mandatory:"true" contributesTo:"path" name:"discoveryJobId"`

	// The log type like INFO, WARNING, ERROR, SUCCESS
	LogType ListDiscoveryJobLogsLogTypeEnum `mandatory:"false" contributesTo:"query" name:"logType" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDiscoveryJobLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for logType is ascending.
	SortBy ListDiscoveryJobLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDiscoveryJobLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDiscoveryJobLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDiscoveryJobLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDiscoveryJobLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDiscoveryJobLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDiscoveryJobLogsLogTypeEnum(string(request.LogType)); !ok && request.LogType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogType: %s. Supported values are: %s.", request.LogType, strings.Join(GetListDiscoveryJobLogsLogTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiscoveryJobLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDiscoveryJobLogsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiscoveryJobLogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDiscoveryJobLogsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDiscoveryJobLogsResponse wrapper for the ListDiscoveryJobLogs operation
type ListDiscoveryJobLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DiscoveryJobLogCollection instances
	DiscoveryJobLogCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDiscoveryJobLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDiscoveryJobLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDiscoveryJobLogsLogTypeEnum Enum with underlying type: string
type ListDiscoveryJobLogsLogTypeEnum string

// Set of constants representing the allowable values for ListDiscoveryJobLogsLogTypeEnum
const (
	ListDiscoveryJobLogsLogTypeInfo    ListDiscoveryJobLogsLogTypeEnum = "INFO"
	ListDiscoveryJobLogsLogTypeWarning ListDiscoveryJobLogsLogTypeEnum = "WARNING"
	ListDiscoveryJobLogsLogTypeError   ListDiscoveryJobLogsLogTypeEnum = "ERROR"
	ListDiscoveryJobLogsLogTypeSuccess ListDiscoveryJobLogsLogTypeEnum = "SUCCESS"
)

var mappingListDiscoveryJobLogsLogTypeEnum = map[string]ListDiscoveryJobLogsLogTypeEnum{
	"INFO":    ListDiscoveryJobLogsLogTypeInfo,
	"WARNING": ListDiscoveryJobLogsLogTypeWarning,
	"ERROR":   ListDiscoveryJobLogsLogTypeError,
	"SUCCESS": ListDiscoveryJobLogsLogTypeSuccess,
}

var mappingListDiscoveryJobLogsLogTypeEnumLowerCase = map[string]ListDiscoveryJobLogsLogTypeEnum{
	"info":    ListDiscoveryJobLogsLogTypeInfo,
	"warning": ListDiscoveryJobLogsLogTypeWarning,
	"error":   ListDiscoveryJobLogsLogTypeError,
	"success": ListDiscoveryJobLogsLogTypeSuccess,
}

// GetListDiscoveryJobLogsLogTypeEnumValues Enumerates the set of values for ListDiscoveryJobLogsLogTypeEnum
func GetListDiscoveryJobLogsLogTypeEnumValues() []ListDiscoveryJobLogsLogTypeEnum {
	values := make([]ListDiscoveryJobLogsLogTypeEnum, 0)
	for _, v := range mappingListDiscoveryJobLogsLogTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryJobLogsLogTypeEnumStringValues Enumerates the set of values in String for ListDiscoveryJobLogsLogTypeEnum
func GetListDiscoveryJobLogsLogTypeEnumStringValues() []string {
	return []string{
		"INFO",
		"WARNING",
		"ERROR",
		"SUCCESS",
	}
}

// GetMappingListDiscoveryJobLogsLogTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryJobLogsLogTypeEnum(val string) (ListDiscoveryJobLogsLogTypeEnum, bool) {
	enum, ok := mappingListDiscoveryJobLogsLogTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDiscoveryJobLogsSortOrderEnum Enum with underlying type: string
type ListDiscoveryJobLogsSortOrderEnum string

// Set of constants representing the allowable values for ListDiscoveryJobLogsSortOrderEnum
const (
	ListDiscoveryJobLogsSortOrderAsc  ListDiscoveryJobLogsSortOrderEnum = "ASC"
	ListDiscoveryJobLogsSortOrderDesc ListDiscoveryJobLogsSortOrderEnum = "DESC"
)

var mappingListDiscoveryJobLogsSortOrderEnum = map[string]ListDiscoveryJobLogsSortOrderEnum{
	"ASC":  ListDiscoveryJobLogsSortOrderAsc,
	"DESC": ListDiscoveryJobLogsSortOrderDesc,
}

var mappingListDiscoveryJobLogsSortOrderEnumLowerCase = map[string]ListDiscoveryJobLogsSortOrderEnum{
	"asc":  ListDiscoveryJobLogsSortOrderAsc,
	"desc": ListDiscoveryJobLogsSortOrderDesc,
}

// GetListDiscoveryJobLogsSortOrderEnumValues Enumerates the set of values for ListDiscoveryJobLogsSortOrderEnum
func GetListDiscoveryJobLogsSortOrderEnumValues() []ListDiscoveryJobLogsSortOrderEnum {
	values := make([]ListDiscoveryJobLogsSortOrderEnum, 0)
	for _, v := range mappingListDiscoveryJobLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryJobLogsSortOrderEnumStringValues Enumerates the set of values in String for ListDiscoveryJobLogsSortOrderEnum
func GetListDiscoveryJobLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDiscoveryJobLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryJobLogsSortOrderEnum(val string) (ListDiscoveryJobLogsSortOrderEnum, bool) {
	enum, ok := mappingListDiscoveryJobLogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDiscoveryJobLogsSortByEnum Enum with underlying type: string
type ListDiscoveryJobLogsSortByEnum string

// Set of constants representing the allowable values for ListDiscoveryJobLogsSortByEnum
const (
	ListDiscoveryJobLogsSortByTimecreated ListDiscoveryJobLogsSortByEnum = "timeCreated"
	ListDiscoveryJobLogsSortByLogtype     ListDiscoveryJobLogsSortByEnum = "logType"
)

var mappingListDiscoveryJobLogsSortByEnum = map[string]ListDiscoveryJobLogsSortByEnum{
	"timeCreated": ListDiscoveryJobLogsSortByTimecreated,
	"logType":     ListDiscoveryJobLogsSortByLogtype,
}

var mappingListDiscoveryJobLogsSortByEnumLowerCase = map[string]ListDiscoveryJobLogsSortByEnum{
	"timecreated": ListDiscoveryJobLogsSortByTimecreated,
	"logtype":     ListDiscoveryJobLogsSortByLogtype,
}

// GetListDiscoveryJobLogsSortByEnumValues Enumerates the set of values for ListDiscoveryJobLogsSortByEnum
func GetListDiscoveryJobLogsSortByEnumValues() []ListDiscoveryJobLogsSortByEnum {
	values := make([]ListDiscoveryJobLogsSortByEnum, 0)
	for _, v := range mappingListDiscoveryJobLogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryJobLogsSortByEnumStringValues Enumerates the set of values in String for ListDiscoveryJobLogsSortByEnum
func GetListDiscoveryJobLogsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"logType",
	}
}

// GetMappingListDiscoveryJobLogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryJobLogsSortByEnum(val string) (ListDiscoveryJobLogsSortByEnum, bool) {
	enum, ok := mappingListDiscoveryJobLogsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
