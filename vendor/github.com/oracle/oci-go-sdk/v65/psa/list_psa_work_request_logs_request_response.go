// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package psa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPsaWorkRequestLogsRequest wrapper for the ListPsaWorkRequestLogs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/ListPsaWorkRequestLogs.go.html to see an example of how to use ListPsaWorkRequestLogsRequest.
type ListPsaWorkRequestLogsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the asynchronous work request.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for `timeCreated` is descending.
	SortBy ListPsaWorkRequestLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListPsaWorkRequestLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPsaWorkRequestLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPsaWorkRequestLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPsaWorkRequestLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPsaWorkRequestLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPsaWorkRequestLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPsaWorkRequestLogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPsaWorkRequestLogsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPsaWorkRequestLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPsaWorkRequestLogsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPsaWorkRequestLogsResponse wrapper for the ListPsaWorkRequestLogs operation
type ListPsaWorkRequestLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestLogEntryCollection instances
	WorkRequestLogEntryCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListPsaWorkRequestLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPsaWorkRequestLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPsaWorkRequestLogsSortByEnum Enum with underlying type: string
type ListPsaWorkRequestLogsSortByEnum string

// Set of constants representing the allowable values for ListPsaWorkRequestLogsSortByEnum
const (
	ListPsaWorkRequestLogsSortByTimecreated ListPsaWorkRequestLogsSortByEnum = "timeCreated"
)

var mappingListPsaWorkRequestLogsSortByEnum = map[string]ListPsaWorkRequestLogsSortByEnum{
	"timeCreated": ListPsaWorkRequestLogsSortByTimecreated,
}

var mappingListPsaWorkRequestLogsSortByEnumLowerCase = map[string]ListPsaWorkRequestLogsSortByEnum{
	"timecreated": ListPsaWorkRequestLogsSortByTimecreated,
}

// GetListPsaWorkRequestLogsSortByEnumValues Enumerates the set of values for ListPsaWorkRequestLogsSortByEnum
func GetListPsaWorkRequestLogsSortByEnumValues() []ListPsaWorkRequestLogsSortByEnum {
	values := make([]ListPsaWorkRequestLogsSortByEnum, 0)
	for _, v := range mappingListPsaWorkRequestLogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPsaWorkRequestLogsSortByEnumStringValues Enumerates the set of values in String for ListPsaWorkRequestLogsSortByEnum
func GetListPsaWorkRequestLogsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListPsaWorkRequestLogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPsaWorkRequestLogsSortByEnum(val string) (ListPsaWorkRequestLogsSortByEnum, bool) {
	enum, ok := mappingListPsaWorkRequestLogsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPsaWorkRequestLogsSortOrderEnum Enum with underlying type: string
type ListPsaWorkRequestLogsSortOrderEnum string

// Set of constants representing the allowable values for ListPsaWorkRequestLogsSortOrderEnum
const (
	ListPsaWorkRequestLogsSortOrderAsc  ListPsaWorkRequestLogsSortOrderEnum = "ASC"
	ListPsaWorkRequestLogsSortOrderDesc ListPsaWorkRequestLogsSortOrderEnum = "DESC"
)

var mappingListPsaWorkRequestLogsSortOrderEnum = map[string]ListPsaWorkRequestLogsSortOrderEnum{
	"ASC":  ListPsaWorkRequestLogsSortOrderAsc,
	"DESC": ListPsaWorkRequestLogsSortOrderDesc,
}

var mappingListPsaWorkRequestLogsSortOrderEnumLowerCase = map[string]ListPsaWorkRequestLogsSortOrderEnum{
	"asc":  ListPsaWorkRequestLogsSortOrderAsc,
	"desc": ListPsaWorkRequestLogsSortOrderDesc,
}

// GetListPsaWorkRequestLogsSortOrderEnumValues Enumerates the set of values for ListPsaWorkRequestLogsSortOrderEnum
func GetListPsaWorkRequestLogsSortOrderEnumValues() []ListPsaWorkRequestLogsSortOrderEnum {
	values := make([]ListPsaWorkRequestLogsSortOrderEnum, 0)
	for _, v := range mappingListPsaWorkRequestLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPsaWorkRequestLogsSortOrderEnumStringValues Enumerates the set of values in String for ListPsaWorkRequestLogsSortOrderEnum
func GetListPsaWorkRequestLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPsaWorkRequestLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPsaWorkRequestLogsSortOrderEnum(val string) (ListPsaWorkRequestLogsSortOrderEnum, bool) {
	enum, ok := mappingListPsaWorkRequestLogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
