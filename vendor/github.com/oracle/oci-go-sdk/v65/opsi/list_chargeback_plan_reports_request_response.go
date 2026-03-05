// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListChargebackPlanReportsRequest wrapper for the ListChargebackPlanReports operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListChargebackPlanReports.go.html to see an example of how to use ListChargebackPlanReportsRequest.
type ListChargebackPlanReportsRequest struct {

	// Unique Ops insight identifier
	Id *string `mandatory:"true" contributesTo:"query" name:"id"`

	// Filter by resource type.
	// Supported values are EXADATA_INSIGHT , HOST_INSIGHT, DATABASE_INSIGHT.
	ResourceType *string `mandatory:"true" contributesTo:"query" name:"resourceType"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListChargebackPlanReportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort chargeback plan reports.
	SortBy ListChargebackPlanReportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListChargebackPlanReportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListChargebackPlanReportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListChargebackPlanReportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListChargebackPlanReportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListChargebackPlanReportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListChargebackPlanReportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListChargebackPlanReportsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListChargebackPlanReportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListChargebackPlanReportsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListChargebackPlanReportsResponse wrapper for the ListChargebackPlanReports operation
type ListChargebackPlanReportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ChargebackPlanReportCollection instances
	ChargebackPlanReportCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListChargebackPlanReportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListChargebackPlanReportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListChargebackPlanReportsSortOrderEnum Enum with underlying type: string
type ListChargebackPlanReportsSortOrderEnum string

// Set of constants representing the allowable values for ListChargebackPlanReportsSortOrderEnum
const (
	ListChargebackPlanReportsSortOrderAsc  ListChargebackPlanReportsSortOrderEnum = "ASC"
	ListChargebackPlanReportsSortOrderDesc ListChargebackPlanReportsSortOrderEnum = "DESC"
)

var mappingListChargebackPlanReportsSortOrderEnum = map[string]ListChargebackPlanReportsSortOrderEnum{
	"ASC":  ListChargebackPlanReportsSortOrderAsc,
	"DESC": ListChargebackPlanReportsSortOrderDesc,
}

var mappingListChargebackPlanReportsSortOrderEnumLowerCase = map[string]ListChargebackPlanReportsSortOrderEnum{
	"asc":  ListChargebackPlanReportsSortOrderAsc,
	"desc": ListChargebackPlanReportsSortOrderDesc,
}

// GetListChargebackPlanReportsSortOrderEnumValues Enumerates the set of values for ListChargebackPlanReportsSortOrderEnum
func GetListChargebackPlanReportsSortOrderEnumValues() []ListChargebackPlanReportsSortOrderEnum {
	values := make([]ListChargebackPlanReportsSortOrderEnum, 0)
	for _, v := range mappingListChargebackPlanReportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListChargebackPlanReportsSortOrderEnumStringValues Enumerates the set of values in String for ListChargebackPlanReportsSortOrderEnum
func GetListChargebackPlanReportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListChargebackPlanReportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChargebackPlanReportsSortOrderEnum(val string) (ListChargebackPlanReportsSortOrderEnum, bool) {
	enum, ok := mappingListChargebackPlanReportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListChargebackPlanReportsSortByEnum Enum with underlying type: string
type ListChargebackPlanReportsSortByEnum string

// Set of constants representing the allowable values for ListChargebackPlanReportsSortByEnum
const (
	ListChargebackPlanReportsSortByTimecreated ListChargebackPlanReportsSortByEnum = "timeCreated"
	ListChargebackPlanReportsSortById          ListChargebackPlanReportsSortByEnum = "id"
)

var mappingListChargebackPlanReportsSortByEnum = map[string]ListChargebackPlanReportsSortByEnum{
	"timeCreated": ListChargebackPlanReportsSortByTimecreated,
	"id":          ListChargebackPlanReportsSortById,
}

var mappingListChargebackPlanReportsSortByEnumLowerCase = map[string]ListChargebackPlanReportsSortByEnum{
	"timecreated": ListChargebackPlanReportsSortByTimecreated,
	"id":          ListChargebackPlanReportsSortById,
}

// GetListChargebackPlanReportsSortByEnumValues Enumerates the set of values for ListChargebackPlanReportsSortByEnum
func GetListChargebackPlanReportsSortByEnumValues() []ListChargebackPlanReportsSortByEnum {
	values := make([]ListChargebackPlanReportsSortByEnum, 0)
	for _, v := range mappingListChargebackPlanReportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListChargebackPlanReportsSortByEnumStringValues Enumerates the set of values in String for ListChargebackPlanReportsSortByEnum
func GetListChargebackPlanReportsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"id",
	}
}

// GetMappingListChargebackPlanReportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChargebackPlanReportsSortByEnum(val string) (ListChargebackPlanReportsSortByEnum, bool) {
	enum, ok := mappingListChargebackPlanReportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
