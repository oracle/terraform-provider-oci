// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDiagnosticActionsRequest wrapper for the ListDiagnosticActions operation
type ListDiagnosticActionsRequest struct {

	// The ID of the compartment in which data is listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDiagnosticActionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for 'timeCreated' is descending.
	// Default order for 'displayName' and 'diagnosticActionType' is ascending.
	SortBy ListDiagnosticActionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDiagnosticActionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDiagnosticActionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDiagnosticActionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDiagnosticActionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDiagnosticActionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDiagnosticActionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDiagnosticActionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiagnosticActionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDiagnosticActionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDiagnosticActionsResponse wrapper for the ListDiagnosticActions operation
type ListDiagnosticActionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DiagnosticActionsCollection instances
	DiagnosticActionsCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDiagnosticActionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDiagnosticActionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDiagnosticActionsSortOrderEnum Enum with underlying type: string
type ListDiagnosticActionsSortOrderEnum string

// Set of constants representing the allowable values for ListDiagnosticActionsSortOrderEnum
const (
	ListDiagnosticActionsSortOrderAsc  ListDiagnosticActionsSortOrderEnum = "ASC"
	ListDiagnosticActionsSortOrderDesc ListDiagnosticActionsSortOrderEnum = "DESC"
)

var mappingListDiagnosticActionsSortOrderEnum = map[string]ListDiagnosticActionsSortOrderEnum{
	"ASC":  ListDiagnosticActionsSortOrderAsc,
	"DESC": ListDiagnosticActionsSortOrderDesc,
}

var mappingListDiagnosticActionsSortOrderEnumLowerCase = map[string]ListDiagnosticActionsSortOrderEnum{
	"asc":  ListDiagnosticActionsSortOrderAsc,
	"desc": ListDiagnosticActionsSortOrderDesc,
}

// GetListDiagnosticActionsSortOrderEnumValues Enumerates the set of values for ListDiagnosticActionsSortOrderEnum
func GetListDiagnosticActionsSortOrderEnumValues() []ListDiagnosticActionsSortOrderEnum {
	values := make([]ListDiagnosticActionsSortOrderEnum, 0)
	for _, v := range mappingListDiagnosticActionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiagnosticActionsSortOrderEnumStringValues Enumerates the set of values in String for ListDiagnosticActionsSortOrderEnum
func GetListDiagnosticActionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDiagnosticActionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiagnosticActionsSortOrderEnum(val string) (ListDiagnosticActionsSortOrderEnum, bool) {
	enum, ok := mappingListDiagnosticActionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDiagnosticActionsSortByEnum Enum with underlying type: string
type ListDiagnosticActionsSortByEnum string

// Set of constants representing the allowable values for ListDiagnosticActionsSortByEnum
const (
	ListDiagnosticActionsSortByTimecreated          ListDiagnosticActionsSortByEnum = "timeCreated"
	ListDiagnosticActionsSortByDiagnosticactiontype ListDiagnosticActionsSortByEnum = "diagnosticActionType"
	ListDiagnosticActionsSortByDisplayname          ListDiagnosticActionsSortByEnum = "displayName"
)

var mappingListDiagnosticActionsSortByEnum = map[string]ListDiagnosticActionsSortByEnum{
	"timeCreated":          ListDiagnosticActionsSortByTimecreated,
	"diagnosticActionType": ListDiagnosticActionsSortByDiagnosticactiontype,
	"displayName":          ListDiagnosticActionsSortByDisplayname,
}

var mappingListDiagnosticActionsSortByEnumLowerCase = map[string]ListDiagnosticActionsSortByEnum{
	"timecreated":          ListDiagnosticActionsSortByTimecreated,
	"diagnosticactiontype": ListDiagnosticActionsSortByDiagnosticactiontype,
	"displayname":          ListDiagnosticActionsSortByDisplayname,
}

// GetListDiagnosticActionsSortByEnumValues Enumerates the set of values for ListDiagnosticActionsSortByEnum
func GetListDiagnosticActionsSortByEnumValues() []ListDiagnosticActionsSortByEnum {
	values := make([]ListDiagnosticActionsSortByEnum, 0)
	for _, v := range mappingListDiagnosticActionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiagnosticActionsSortByEnumStringValues Enumerates the set of values in String for ListDiagnosticActionsSortByEnum
func GetListDiagnosticActionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"diagnosticActionType",
		"displayName",
	}
}

// GetMappingListDiagnosticActionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiagnosticActionsSortByEnum(val string) (ListDiagnosticActionsSortByEnum, bool) {
	enum, ok := mappingListDiagnosticActionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
