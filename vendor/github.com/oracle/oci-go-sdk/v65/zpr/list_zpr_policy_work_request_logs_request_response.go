// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package zpr

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListZprPolicyWorkRequestLogsRequest wrapper for the ListZprPolicyWorkRequestLogs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/ListZprPolicyWorkRequestLogs.go.html to see an example of how to use ListZprPolicyWorkRequestLogsRequest.
type ListZprPolicyWorkRequestLogsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the asynchronous work request.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for `timestamp` is descending.
	SortBy ListZprPolicyWorkRequestLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListZprPolicyWorkRequestLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListZprPolicyWorkRequestLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListZprPolicyWorkRequestLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListZprPolicyWorkRequestLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListZprPolicyWorkRequestLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListZprPolicyWorkRequestLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListZprPolicyWorkRequestLogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListZprPolicyWorkRequestLogsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListZprPolicyWorkRequestLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListZprPolicyWorkRequestLogsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListZprPolicyWorkRequestLogsResponse wrapper for the ListZprPolicyWorkRequestLogs operation
type ListZprPolicyWorkRequestLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestLogEntryCollection instances
	WorkRequestLogEntryCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListZprPolicyWorkRequestLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListZprPolicyWorkRequestLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListZprPolicyWorkRequestLogsSortByEnum Enum with underlying type: string
type ListZprPolicyWorkRequestLogsSortByEnum string

// Set of constants representing the allowable values for ListZprPolicyWorkRequestLogsSortByEnum
const (
	ListZprPolicyWorkRequestLogsSortByTimestamp ListZprPolicyWorkRequestLogsSortByEnum = "timestamp"
)

var mappingListZprPolicyWorkRequestLogsSortByEnum = map[string]ListZprPolicyWorkRequestLogsSortByEnum{
	"timestamp": ListZprPolicyWorkRequestLogsSortByTimestamp,
}

var mappingListZprPolicyWorkRequestLogsSortByEnumLowerCase = map[string]ListZprPolicyWorkRequestLogsSortByEnum{
	"timestamp": ListZprPolicyWorkRequestLogsSortByTimestamp,
}

// GetListZprPolicyWorkRequestLogsSortByEnumValues Enumerates the set of values for ListZprPolicyWorkRequestLogsSortByEnum
func GetListZprPolicyWorkRequestLogsSortByEnumValues() []ListZprPolicyWorkRequestLogsSortByEnum {
	values := make([]ListZprPolicyWorkRequestLogsSortByEnum, 0)
	for _, v := range mappingListZprPolicyWorkRequestLogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprPolicyWorkRequestLogsSortByEnumStringValues Enumerates the set of values in String for ListZprPolicyWorkRequestLogsSortByEnum
func GetListZprPolicyWorkRequestLogsSortByEnumStringValues() []string {
	return []string{
		"timestamp",
	}
}

// GetMappingListZprPolicyWorkRequestLogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprPolicyWorkRequestLogsSortByEnum(val string) (ListZprPolicyWorkRequestLogsSortByEnum, bool) {
	enum, ok := mappingListZprPolicyWorkRequestLogsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListZprPolicyWorkRequestLogsSortOrderEnum Enum with underlying type: string
type ListZprPolicyWorkRequestLogsSortOrderEnum string

// Set of constants representing the allowable values for ListZprPolicyWorkRequestLogsSortOrderEnum
const (
	ListZprPolicyWorkRequestLogsSortOrderAsc  ListZprPolicyWorkRequestLogsSortOrderEnum = "ASC"
	ListZprPolicyWorkRequestLogsSortOrderDesc ListZprPolicyWorkRequestLogsSortOrderEnum = "DESC"
)

var mappingListZprPolicyWorkRequestLogsSortOrderEnum = map[string]ListZprPolicyWorkRequestLogsSortOrderEnum{
	"ASC":  ListZprPolicyWorkRequestLogsSortOrderAsc,
	"DESC": ListZprPolicyWorkRequestLogsSortOrderDesc,
}

var mappingListZprPolicyWorkRequestLogsSortOrderEnumLowerCase = map[string]ListZprPolicyWorkRequestLogsSortOrderEnum{
	"asc":  ListZprPolicyWorkRequestLogsSortOrderAsc,
	"desc": ListZprPolicyWorkRequestLogsSortOrderDesc,
}

// GetListZprPolicyWorkRequestLogsSortOrderEnumValues Enumerates the set of values for ListZprPolicyWorkRequestLogsSortOrderEnum
func GetListZprPolicyWorkRequestLogsSortOrderEnumValues() []ListZprPolicyWorkRequestLogsSortOrderEnum {
	values := make([]ListZprPolicyWorkRequestLogsSortOrderEnum, 0)
	for _, v := range mappingListZprPolicyWorkRequestLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprPolicyWorkRequestLogsSortOrderEnumStringValues Enumerates the set of values in String for ListZprPolicyWorkRequestLogsSortOrderEnum
func GetListZprPolicyWorkRequestLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListZprPolicyWorkRequestLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprPolicyWorkRequestLogsSortOrderEnum(val string) (ListZprPolicyWorkRequestLogsSortOrderEnum, bool) {
	enum, ok := mappingListZprPolicyWorkRequestLogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
