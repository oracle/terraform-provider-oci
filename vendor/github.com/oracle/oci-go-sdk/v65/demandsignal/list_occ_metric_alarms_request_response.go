// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package demandsignal

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOccMetricAlarmsRequest wrapper for the ListOccMetricAlarms operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/demandsignal/ListOccMetricAlarms.go.html to see an example of how to use ListOccMetricAlarmsRequest.
type ListOccMetricAlarmsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter to list only active or inactive alarms.
	IsActive *bool `mandatory:"false" contributesTo:"query" name:"isActive"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListOccMetricAlarmsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListOccMetricAlarmsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccMetricAlarmsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccMetricAlarmsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccMetricAlarmsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccMetricAlarmsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccMetricAlarmsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOccMetricAlarmsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccMetricAlarmsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccMetricAlarmsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccMetricAlarmsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccMetricAlarmsResponse wrapper for the ListOccMetricAlarms operation
type ListOccMetricAlarmsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OccMetricAlarmCollection instances
	OccMetricAlarmCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOccMetricAlarmsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccMetricAlarmsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccMetricAlarmsSortOrderEnum Enum with underlying type: string
type ListOccMetricAlarmsSortOrderEnum string

// Set of constants representing the allowable values for ListOccMetricAlarmsSortOrderEnum
const (
	ListOccMetricAlarmsSortOrderAsc  ListOccMetricAlarmsSortOrderEnum = "ASC"
	ListOccMetricAlarmsSortOrderDesc ListOccMetricAlarmsSortOrderEnum = "DESC"
)

var mappingListOccMetricAlarmsSortOrderEnum = map[string]ListOccMetricAlarmsSortOrderEnum{
	"ASC":  ListOccMetricAlarmsSortOrderAsc,
	"DESC": ListOccMetricAlarmsSortOrderDesc,
}

var mappingListOccMetricAlarmsSortOrderEnumLowerCase = map[string]ListOccMetricAlarmsSortOrderEnum{
	"asc":  ListOccMetricAlarmsSortOrderAsc,
	"desc": ListOccMetricAlarmsSortOrderDesc,
}

// GetListOccMetricAlarmsSortOrderEnumValues Enumerates the set of values for ListOccMetricAlarmsSortOrderEnum
func GetListOccMetricAlarmsSortOrderEnumValues() []ListOccMetricAlarmsSortOrderEnum {
	values := make([]ListOccMetricAlarmsSortOrderEnum, 0)
	for _, v := range mappingListOccMetricAlarmsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccMetricAlarmsSortOrderEnumStringValues Enumerates the set of values in String for ListOccMetricAlarmsSortOrderEnum
func GetListOccMetricAlarmsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccMetricAlarmsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccMetricAlarmsSortOrderEnum(val string) (ListOccMetricAlarmsSortOrderEnum, bool) {
	enum, ok := mappingListOccMetricAlarmsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccMetricAlarmsSortByEnum Enum with underlying type: string
type ListOccMetricAlarmsSortByEnum string

// Set of constants representing the allowable values for ListOccMetricAlarmsSortByEnum
const (
	ListOccMetricAlarmsSortByTimecreated ListOccMetricAlarmsSortByEnum = "timeCreated"
	ListOccMetricAlarmsSortByDisplayname ListOccMetricAlarmsSortByEnum = "displayName"
)

var mappingListOccMetricAlarmsSortByEnum = map[string]ListOccMetricAlarmsSortByEnum{
	"timeCreated": ListOccMetricAlarmsSortByTimecreated,
	"displayName": ListOccMetricAlarmsSortByDisplayname,
}

var mappingListOccMetricAlarmsSortByEnumLowerCase = map[string]ListOccMetricAlarmsSortByEnum{
	"timecreated": ListOccMetricAlarmsSortByTimecreated,
	"displayname": ListOccMetricAlarmsSortByDisplayname,
}

// GetListOccMetricAlarmsSortByEnumValues Enumerates the set of values for ListOccMetricAlarmsSortByEnum
func GetListOccMetricAlarmsSortByEnumValues() []ListOccMetricAlarmsSortByEnum {
	values := make([]ListOccMetricAlarmsSortByEnum, 0)
	for _, v := range mappingListOccMetricAlarmsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccMetricAlarmsSortByEnumStringValues Enumerates the set of values in String for ListOccMetricAlarmsSortByEnum
func GetListOccMetricAlarmsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOccMetricAlarmsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccMetricAlarmsSortByEnum(val string) (ListOccMetricAlarmsSortByEnum, bool) {
	enum, ok := mappingListOccMetricAlarmsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
