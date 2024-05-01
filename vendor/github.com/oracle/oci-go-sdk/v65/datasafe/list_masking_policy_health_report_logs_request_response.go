// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMaskingPolicyHealthReportLogsRequest wrapper for the ListMaskingPolicyHealthReportLogs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingPolicyHealthReportLogs.go.html to see an example of how to use ListMaskingPolicyHealthReportLogsRequest.
type ListMaskingPolicyHealthReportLogsRequest struct {

	// The OCID of the masking health report.
	MaskingPolicyHealthReportId *string `mandatory:"true" contributesTo:"path" name:"maskingPolicyHealthReportId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListMaskingPolicyHealthReportLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// sort by
	SortBy ListMaskingPolicyHealthReportLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only the resources that match the specified log message type.
	MessageType ListMaskingPolicyHealthReportLogsMessageTypeEnum `mandatory:"false" contributesTo:"query" name:"messageType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaskingPolicyHealthReportLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaskingPolicyHealthReportLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaskingPolicyHealthReportLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaskingPolicyHealthReportLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaskingPolicyHealthReportLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMaskingPolicyHealthReportLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaskingPolicyHealthReportLogsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingPolicyHealthReportLogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaskingPolicyHealthReportLogsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingPolicyHealthReportLogsMessageTypeEnum(string(request.MessageType)); !ok && request.MessageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MessageType: %s. Supported values are: %s.", request.MessageType, strings.Join(GetListMaskingPolicyHealthReportLogsMessageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaskingPolicyHealthReportLogsResponse wrapper for the ListMaskingPolicyHealthReportLogs operation
type ListMaskingPolicyHealthReportLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaskingPolicyHealthReportLogCollection instances
	MaskingPolicyHealthReportLogCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListMaskingPolicyHealthReportLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaskingPolicyHealthReportLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaskingPolicyHealthReportLogsSortOrderEnum Enum with underlying type: string
type ListMaskingPolicyHealthReportLogsSortOrderEnum string

// Set of constants representing the allowable values for ListMaskingPolicyHealthReportLogsSortOrderEnum
const (
	ListMaskingPolicyHealthReportLogsSortOrderAsc  ListMaskingPolicyHealthReportLogsSortOrderEnum = "ASC"
	ListMaskingPolicyHealthReportLogsSortOrderDesc ListMaskingPolicyHealthReportLogsSortOrderEnum = "DESC"
)

var mappingListMaskingPolicyHealthReportLogsSortOrderEnum = map[string]ListMaskingPolicyHealthReportLogsSortOrderEnum{
	"ASC":  ListMaskingPolicyHealthReportLogsSortOrderAsc,
	"DESC": ListMaskingPolicyHealthReportLogsSortOrderDesc,
}

var mappingListMaskingPolicyHealthReportLogsSortOrderEnumLowerCase = map[string]ListMaskingPolicyHealthReportLogsSortOrderEnum{
	"asc":  ListMaskingPolicyHealthReportLogsSortOrderAsc,
	"desc": ListMaskingPolicyHealthReportLogsSortOrderDesc,
}

// GetListMaskingPolicyHealthReportLogsSortOrderEnumValues Enumerates the set of values for ListMaskingPolicyHealthReportLogsSortOrderEnum
func GetListMaskingPolicyHealthReportLogsSortOrderEnumValues() []ListMaskingPolicyHealthReportLogsSortOrderEnum {
	values := make([]ListMaskingPolicyHealthReportLogsSortOrderEnum, 0)
	for _, v := range mappingListMaskingPolicyHealthReportLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingPolicyHealthReportLogsSortOrderEnumStringValues Enumerates the set of values in String for ListMaskingPolicyHealthReportLogsSortOrderEnum
func GetListMaskingPolicyHealthReportLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaskingPolicyHealthReportLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingPolicyHealthReportLogsSortOrderEnum(val string) (ListMaskingPolicyHealthReportLogsSortOrderEnum, bool) {
	enum, ok := mappingListMaskingPolicyHealthReportLogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingPolicyHealthReportLogsSortByEnum Enum with underlying type: string
type ListMaskingPolicyHealthReportLogsSortByEnum string

// Set of constants representing the allowable values for ListMaskingPolicyHealthReportLogsSortByEnum
const (
	ListMaskingPolicyHealthReportLogsSortByLogtype ListMaskingPolicyHealthReportLogsSortByEnum = "logType"
)

var mappingListMaskingPolicyHealthReportLogsSortByEnum = map[string]ListMaskingPolicyHealthReportLogsSortByEnum{
	"logType": ListMaskingPolicyHealthReportLogsSortByLogtype,
}

var mappingListMaskingPolicyHealthReportLogsSortByEnumLowerCase = map[string]ListMaskingPolicyHealthReportLogsSortByEnum{
	"logtype": ListMaskingPolicyHealthReportLogsSortByLogtype,
}

// GetListMaskingPolicyHealthReportLogsSortByEnumValues Enumerates the set of values for ListMaskingPolicyHealthReportLogsSortByEnum
func GetListMaskingPolicyHealthReportLogsSortByEnumValues() []ListMaskingPolicyHealthReportLogsSortByEnum {
	values := make([]ListMaskingPolicyHealthReportLogsSortByEnum, 0)
	for _, v := range mappingListMaskingPolicyHealthReportLogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingPolicyHealthReportLogsSortByEnumStringValues Enumerates the set of values in String for ListMaskingPolicyHealthReportLogsSortByEnum
func GetListMaskingPolicyHealthReportLogsSortByEnumStringValues() []string {
	return []string{
		"logType",
	}
}

// GetMappingListMaskingPolicyHealthReportLogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingPolicyHealthReportLogsSortByEnum(val string) (ListMaskingPolicyHealthReportLogsSortByEnum, bool) {
	enum, ok := mappingListMaskingPolicyHealthReportLogsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingPolicyHealthReportLogsMessageTypeEnum Enum with underlying type: string
type ListMaskingPolicyHealthReportLogsMessageTypeEnum string

// Set of constants representing the allowable values for ListMaskingPolicyHealthReportLogsMessageTypeEnum
const (
	ListMaskingPolicyHealthReportLogsMessageTypePass    ListMaskingPolicyHealthReportLogsMessageTypeEnum = "PASS"
	ListMaskingPolicyHealthReportLogsMessageTypeWarning ListMaskingPolicyHealthReportLogsMessageTypeEnum = "WARNING"
	ListMaskingPolicyHealthReportLogsMessageTypeError   ListMaskingPolicyHealthReportLogsMessageTypeEnum = "ERROR"
)

var mappingListMaskingPolicyHealthReportLogsMessageTypeEnum = map[string]ListMaskingPolicyHealthReportLogsMessageTypeEnum{
	"PASS":    ListMaskingPolicyHealthReportLogsMessageTypePass,
	"WARNING": ListMaskingPolicyHealthReportLogsMessageTypeWarning,
	"ERROR":   ListMaskingPolicyHealthReportLogsMessageTypeError,
}

var mappingListMaskingPolicyHealthReportLogsMessageTypeEnumLowerCase = map[string]ListMaskingPolicyHealthReportLogsMessageTypeEnum{
	"pass":    ListMaskingPolicyHealthReportLogsMessageTypePass,
	"warning": ListMaskingPolicyHealthReportLogsMessageTypeWarning,
	"error":   ListMaskingPolicyHealthReportLogsMessageTypeError,
}

// GetListMaskingPolicyHealthReportLogsMessageTypeEnumValues Enumerates the set of values for ListMaskingPolicyHealthReportLogsMessageTypeEnum
func GetListMaskingPolicyHealthReportLogsMessageTypeEnumValues() []ListMaskingPolicyHealthReportLogsMessageTypeEnum {
	values := make([]ListMaskingPolicyHealthReportLogsMessageTypeEnum, 0)
	for _, v := range mappingListMaskingPolicyHealthReportLogsMessageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingPolicyHealthReportLogsMessageTypeEnumStringValues Enumerates the set of values in String for ListMaskingPolicyHealthReportLogsMessageTypeEnum
func GetListMaskingPolicyHealthReportLogsMessageTypeEnumStringValues() []string {
	return []string{
		"PASS",
		"WARNING",
		"ERROR",
	}
}

// GetMappingListMaskingPolicyHealthReportLogsMessageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingPolicyHealthReportLogsMessageTypeEnum(val string) (ListMaskingPolicyHealthReportLogsMessageTypeEnum, bool) {
	enum, ok := mappingListMaskingPolicyHealthReportLogsMessageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
