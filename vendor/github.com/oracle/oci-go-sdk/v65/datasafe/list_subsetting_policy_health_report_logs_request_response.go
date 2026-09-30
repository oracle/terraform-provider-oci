// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSubsettingPolicyHealthReportLogsRequest wrapper for the ListSubsettingPolicyHealthReportLogs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSubsettingPolicyHealthReportLogs.go.html to see an example of how to use ListSubsettingPolicyHealthReportLogsRequest.
type ListSubsettingPolicyHealthReportLogsRequest struct {

	// The OCID of the subsetting health report.
	SubsettingPolicyHealthReportId *string `mandatory:"true" contributesTo:"path" name:"subsettingPolicyHealthReportId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSubsettingPolicyHealthReportLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for messageType is ascending.
	SortBy ListSubsettingPolicyHealthReportLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only the resources that match the specified log message type.
	MessageType ListSubsettingPolicyHealthReportLogsMessageTypeEnum `mandatory:"false" contributesTo:"query" name:"messageType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubsettingPolicyHealthReportLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubsettingPolicyHealthReportLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubsettingPolicyHealthReportLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubsettingPolicyHealthReportLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubsettingPolicyHealthReportLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubsettingPolicyHealthReportLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubsettingPolicyHealthReportLogsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingPolicyHealthReportLogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubsettingPolicyHealthReportLogsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingPolicyHealthReportLogsMessageTypeEnum(string(request.MessageType)); !ok && request.MessageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MessageType: %s. Supported values are: %s.", request.MessageType, strings.Join(GetListSubsettingPolicyHealthReportLogsMessageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubsettingPolicyHealthReportLogsResponse wrapper for the ListSubsettingPolicyHealthReportLogs operation
type ListSubsettingPolicyHealthReportLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubsettingPolicyHealthReportLogCollection instances
	SubsettingPolicyHealthReportLogCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSubsettingPolicyHealthReportLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubsettingPolicyHealthReportLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubsettingPolicyHealthReportLogsSortOrderEnum Enum with underlying type: string
type ListSubsettingPolicyHealthReportLogsSortOrderEnum string

// Set of constants representing the allowable values for ListSubsettingPolicyHealthReportLogsSortOrderEnum
const (
	ListSubsettingPolicyHealthReportLogsSortOrderAsc  ListSubsettingPolicyHealthReportLogsSortOrderEnum = "ASC"
	ListSubsettingPolicyHealthReportLogsSortOrderDesc ListSubsettingPolicyHealthReportLogsSortOrderEnum = "DESC"
)

var mappingListSubsettingPolicyHealthReportLogsSortOrderEnum = map[string]ListSubsettingPolicyHealthReportLogsSortOrderEnum{
	"ASC":  ListSubsettingPolicyHealthReportLogsSortOrderAsc,
	"DESC": ListSubsettingPolicyHealthReportLogsSortOrderDesc,
}

var mappingListSubsettingPolicyHealthReportLogsSortOrderEnumLowerCase = map[string]ListSubsettingPolicyHealthReportLogsSortOrderEnum{
	"asc":  ListSubsettingPolicyHealthReportLogsSortOrderAsc,
	"desc": ListSubsettingPolicyHealthReportLogsSortOrderDesc,
}

// GetListSubsettingPolicyHealthReportLogsSortOrderEnumValues Enumerates the set of values for ListSubsettingPolicyHealthReportLogsSortOrderEnum
func GetListSubsettingPolicyHealthReportLogsSortOrderEnumValues() []ListSubsettingPolicyHealthReportLogsSortOrderEnum {
	values := make([]ListSubsettingPolicyHealthReportLogsSortOrderEnum, 0)
	for _, v := range mappingListSubsettingPolicyHealthReportLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingPolicyHealthReportLogsSortOrderEnumStringValues Enumerates the set of values in String for ListSubsettingPolicyHealthReportLogsSortOrderEnum
func GetListSubsettingPolicyHealthReportLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubsettingPolicyHealthReportLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingPolicyHealthReportLogsSortOrderEnum(val string) (ListSubsettingPolicyHealthReportLogsSortOrderEnum, bool) {
	enum, ok := mappingListSubsettingPolicyHealthReportLogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingPolicyHealthReportLogsSortByEnum Enum with underlying type: string
type ListSubsettingPolicyHealthReportLogsSortByEnum string

// Set of constants representing the allowable values for ListSubsettingPolicyHealthReportLogsSortByEnum
const (
	ListSubsettingPolicyHealthReportLogsSortByMessagetype ListSubsettingPolicyHealthReportLogsSortByEnum = "messageType"
)

var mappingListSubsettingPolicyHealthReportLogsSortByEnum = map[string]ListSubsettingPolicyHealthReportLogsSortByEnum{
	"messageType": ListSubsettingPolicyHealthReportLogsSortByMessagetype,
}

var mappingListSubsettingPolicyHealthReportLogsSortByEnumLowerCase = map[string]ListSubsettingPolicyHealthReportLogsSortByEnum{
	"messagetype": ListSubsettingPolicyHealthReportLogsSortByMessagetype,
}

// GetListSubsettingPolicyHealthReportLogsSortByEnumValues Enumerates the set of values for ListSubsettingPolicyHealthReportLogsSortByEnum
func GetListSubsettingPolicyHealthReportLogsSortByEnumValues() []ListSubsettingPolicyHealthReportLogsSortByEnum {
	values := make([]ListSubsettingPolicyHealthReportLogsSortByEnum, 0)
	for _, v := range mappingListSubsettingPolicyHealthReportLogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingPolicyHealthReportLogsSortByEnumStringValues Enumerates the set of values in String for ListSubsettingPolicyHealthReportLogsSortByEnum
func GetListSubsettingPolicyHealthReportLogsSortByEnumStringValues() []string {
	return []string{
		"messageType",
	}
}

// GetMappingListSubsettingPolicyHealthReportLogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingPolicyHealthReportLogsSortByEnum(val string) (ListSubsettingPolicyHealthReportLogsSortByEnum, bool) {
	enum, ok := mappingListSubsettingPolicyHealthReportLogsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingPolicyHealthReportLogsMessageTypeEnum Enum with underlying type: string
type ListSubsettingPolicyHealthReportLogsMessageTypeEnum string

// Set of constants representing the allowable values for ListSubsettingPolicyHealthReportLogsMessageTypeEnum
const (
	ListSubsettingPolicyHealthReportLogsMessageTypePass    ListSubsettingPolicyHealthReportLogsMessageTypeEnum = "PASS"
	ListSubsettingPolicyHealthReportLogsMessageTypeWarning ListSubsettingPolicyHealthReportLogsMessageTypeEnum = "WARNING"
	ListSubsettingPolicyHealthReportLogsMessageTypeError   ListSubsettingPolicyHealthReportLogsMessageTypeEnum = "ERROR"
)

var mappingListSubsettingPolicyHealthReportLogsMessageTypeEnum = map[string]ListSubsettingPolicyHealthReportLogsMessageTypeEnum{
	"PASS":    ListSubsettingPolicyHealthReportLogsMessageTypePass,
	"WARNING": ListSubsettingPolicyHealthReportLogsMessageTypeWarning,
	"ERROR":   ListSubsettingPolicyHealthReportLogsMessageTypeError,
}

var mappingListSubsettingPolicyHealthReportLogsMessageTypeEnumLowerCase = map[string]ListSubsettingPolicyHealthReportLogsMessageTypeEnum{
	"pass":    ListSubsettingPolicyHealthReportLogsMessageTypePass,
	"warning": ListSubsettingPolicyHealthReportLogsMessageTypeWarning,
	"error":   ListSubsettingPolicyHealthReportLogsMessageTypeError,
}

// GetListSubsettingPolicyHealthReportLogsMessageTypeEnumValues Enumerates the set of values for ListSubsettingPolicyHealthReportLogsMessageTypeEnum
func GetListSubsettingPolicyHealthReportLogsMessageTypeEnumValues() []ListSubsettingPolicyHealthReportLogsMessageTypeEnum {
	values := make([]ListSubsettingPolicyHealthReportLogsMessageTypeEnum, 0)
	for _, v := range mappingListSubsettingPolicyHealthReportLogsMessageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingPolicyHealthReportLogsMessageTypeEnumStringValues Enumerates the set of values in String for ListSubsettingPolicyHealthReportLogsMessageTypeEnum
func GetListSubsettingPolicyHealthReportLogsMessageTypeEnumStringValues() []string {
	return []string{
		"PASS",
		"WARNING",
		"ERROR",
	}
}

// GetMappingListSubsettingPolicyHealthReportLogsMessageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingPolicyHealthReportLogsMessageTypeEnum(val string) (ListSubsettingPolicyHealthReportLogsMessageTypeEnum, bool) {
	enum, ok := mappingListSubsettingPolicyHealthReportLogsMessageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
