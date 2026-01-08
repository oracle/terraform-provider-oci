// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListReportVersionsRequest wrapper for the ListReportVersions operation
type ListReportVersionsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Report.
	ReportId *string `mandatory:"true" contributesTo:"path" name:"reportId"`

	// The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources that match the given user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListReportVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeUpdated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListReportVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReportVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReportVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListReportVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReportVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListReportVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListReportVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListReportVersionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReportVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListReportVersionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListReportVersionsResponse wrapper for the ListReportVersions operation
type ListReportVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of VersionCollection instances
	VersionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListReportVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReportVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListReportVersionsSortOrderEnum Enum with underlying type: string
type ListReportVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListReportVersionsSortOrderEnum
const (
	ListReportVersionsSortOrderAsc  ListReportVersionsSortOrderEnum = "ASC"
	ListReportVersionsSortOrderDesc ListReportVersionsSortOrderEnum = "DESC"
)

var mappingListReportVersionsSortOrderEnum = map[string]ListReportVersionsSortOrderEnum{
	"ASC":  ListReportVersionsSortOrderAsc,
	"DESC": ListReportVersionsSortOrderDesc,
}

var mappingListReportVersionsSortOrderEnumLowerCase = map[string]ListReportVersionsSortOrderEnum{
	"asc":  ListReportVersionsSortOrderAsc,
	"desc": ListReportVersionsSortOrderDesc,
}

// GetListReportVersionsSortOrderEnumValues Enumerates the set of values for ListReportVersionsSortOrderEnum
func GetListReportVersionsSortOrderEnumValues() []ListReportVersionsSortOrderEnum {
	values := make([]ListReportVersionsSortOrderEnum, 0)
	for _, v := range mappingListReportVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListReportVersionsSortOrderEnum
func GetListReportVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListReportVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportVersionsSortOrderEnum(val string) (ListReportVersionsSortOrderEnum, bool) {
	enum, ok := mappingListReportVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReportVersionsSortByEnum Enum with underlying type: string
type ListReportVersionsSortByEnum string

// Set of constants representing the allowable values for ListReportVersionsSortByEnum
const (
	ListReportVersionsSortByTimecreated ListReportVersionsSortByEnum = "timeCreated"
	ListReportVersionsSortByTimeupdated ListReportVersionsSortByEnum = "timeUpdated"
	ListReportVersionsSortByDisplayname ListReportVersionsSortByEnum = "displayName"
)

var mappingListReportVersionsSortByEnum = map[string]ListReportVersionsSortByEnum{
	"timeCreated": ListReportVersionsSortByTimecreated,
	"timeUpdated": ListReportVersionsSortByTimeupdated,
	"displayName": ListReportVersionsSortByDisplayname,
}

var mappingListReportVersionsSortByEnumLowerCase = map[string]ListReportVersionsSortByEnum{
	"timecreated": ListReportVersionsSortByTimecreated,
	"timeupdated": ListReportVersionsSortByTimeupdated,
	"displayname": ListReportVersionsSortByDisplayname,
}

// GetListReportVersionsSortByEnumValues Enumerates the set of values for ListReportVersionsSortByEnum
func GetListReportVersionsSortByEnumValues() []ListReportVersionsSortByEnum {
	values := make([]ListReportVersionsSortByEnum, 0)
	for _, v := range mappingListReportVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportVersionsSortByEnumStringValues Enumerates the set of values in String for ListReportVersionsSortByEnum
func GetListReportVersionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListReportVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportVersionsSortByEnum(val string) (ListReportVersionsSortByEnum, bool) {
	enum, ok := mappingListReportVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
