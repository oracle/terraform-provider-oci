// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAvailableOverrideMaintenanceStartTimesRequest wrapper for the ListAvailableOverrideMaintenanceStartTimes operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/ListAvailableOverrideMaintenanceStartTimes.go.html to see an example of how to use ListAvailableOverrideMaintenanceStartTimesRequest.
type ListAvailableOverrideMaintenanceStartTimesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Lustre file system.
	Id *string `mandatory:"true" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide only one sort order.
	SortBy ListAvailableOverrideMaintenanceStartTimesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAvailableOverrideMaintenanceStartTimesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Date in format `YYYY-MM-DD`
	Date *string `mandatory:"false" contributesTo:"query" name:"date"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailableOverrideMaintenanceStartTimesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailableOverrideMaintenanceStartTimesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAvailableOverrideMaintenanceStartTimesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailableOverrideMaintenanceStartTimesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAvailableOverrideMaintenanceStartTimesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAvailableOverrideMaintenanceStartTimesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAvailableOverrideMaintenanceStartTimesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAvailableOverrideMaintenanceStartTimesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAvailableOverrideMaintenanceStartTimesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAvailableOverrideMaintenanceStartTimesResponse wrapper for the ListAvailableOverrideMaintenanceStartTimes operation
type ListAvailableOverrideMaintenanceStartTimesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AvailableOverrideMaintenanceStartTimeCollection instances
	AvailableOverrideMaintenanceStartTimeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAvailableOverrideMaintenanceStartTimesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailableOverrideMaintenanceStartTimesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailableOverrideMaintenanceStartTimesSortByEnum Enum with underlying type: string
type ListAvailableOverrideMaintenanceStartTimesSortByEnum string

// Set of constants representing the allowable values for ListAvailableOverrideMaintenanceStartTimesSortByEnum
const (
	ListAvailableOverrideMaintenanceStartTimesSortByDate ListAvailableOverrideMaintenanceStartTimesSortByEnum = "date"
)

var mappingListAvailableOverrideMaintenanceStartTimesSortByEnum = map[string]ListAvailableOverrideMaintenanceStartTimesSortByEnum{
	"date": ListAvailableOverrideMaintenanceStartTimesSortByDate,
}

var mappingListAvailableOverrideMaintenanceStartTimesSortByEnumLowerCase = map[string]ListAvailableOverrideMaintenanceStartTimesSortByEnum{
	"date": ListAvailableOverrideMaintenanceStartTimesSortByDate,
}

// GetListAvailableOverrideMaintenanceStartTimesSortByEnumValues Enumerates the set of values for ListAvailableOverrideMaintenanceStartTimesSortByEnum
func GetListAvailableOverrideMaintenanceStartTimesSortByEnumValues() []ListAvailableOverrideMaintenanceStartTimesSortByEnum {
	values := make([]ListAvailableOverrideMaintenanceStartTimesSortByEnum, 0)
	for _, v := range mappingListAvailableOverrideMaintenanceStartTimesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableOverrideMaintenanceStartTimesSortByEnumStringValues Enumerates the set of values in String for ListAvailableOverrideMaintenanceStartTimesSortByEnum
func GetListAvailableOverrideMaintenanceStartTimesSortByEnumStringValues() []string {
	return []string{
		"date",
	}
}

// GetMappingListAvailableOverrideMaintenanceStartTimesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableOverrideMaintenanceStartTimesSortByEnum(val string) (ListAvailableOverrideMaintenanceStartTimesSortByEnum, bool) {
	enum, ok := mappingListAvailableOverrideMaintenanceStartTimesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAvailableOverrideMaintenanceStartTimesSortOrderEnum Enum with underlying type: string
type ListAvailableOverrideMaintenanceStartTimesSortOrderEnum string

// Set of constants representing the allowable values for ListAvailableOverrideMaintenanceStartTimesSortOrderEnum
const (
	ListAvailableOverrideMaintenanceStartTimesSortOrderAsc  ListAvailableOverrideMaintenanceStartTimesSortOrderEnum = "ASC"
	ListAvailableOverrideMaintenanceStartTimesSortOrderDesc ListAvailableOverrideMaintenanceStartTimesSortOrderEnum = "DESC"
)

var mappingListAvailableOverrideMaintenanceStartTimesSortOrderEnum = map[string]ListAvailableOverrideMaintenanceStartTimesSortOrderEnum{
	"ASC":  ListAvailableOverrideMaintenanceStartTimesSortOrderAsc,
	"DESC": ListAvailableOverrideMaintenanceStartTimesSortOrderDesc,
}

var mappingListAvailableOverrideMaintenanceStartTimesSortOrderEnumLowerCase = map[string]ListAvailableOverrideMaintenanceStartTimesSortOrderEnum{
	"asc":  ListAvailableOverrideMaintenanceStartTimesSortOrderAsc,
	"desc": ListAvailableOverrideMaintenanceStartTimesSortOrderDesc,
}

// GetListAvailableOverrideMaintenanceStartTimesSortOrderEnumValues Enumerates the set of values for ListAvailableOverrideMaintenanceStartTimesSortOrderEnum
func GetListAvailableOverrideMaintenanceStartTimesSortOrderEnumValues() []ListAvailableOverrideMaintenanceStartTimesSortOrderEnum {
	values := make([]ListAvailableOverrideMaintenanceStartTimesSortOrderEnum, 0)
	for _, v := range mappingListAvailableOverrideMaintenanceStartTimesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableOverrideMaintenanceStartTimesSortOrderEnumStringValues Enumerates the set of values in String for ListAvailableOverrideMaintenanceStartTimesSortOrderEnum
func GetListAvailableOverrideMaintenanceStartTimesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAvailableOverrideMaintenanceStartTimesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableOverrideMaintenanceStartTimesSortOrderEnum(val string) (ListAvailableOverrideMaintenanceStartTimesSortOrderEnum, bool) {
	enum, ok := mappingListAvailableOverrideMaintenanceStartTimesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
