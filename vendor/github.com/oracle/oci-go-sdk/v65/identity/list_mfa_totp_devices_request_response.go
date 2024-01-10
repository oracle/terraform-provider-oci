// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMfaTotpDevicesRequest wrapper for the ListMfaTotpDevices operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListMfaTotpDevices.go.html to see an example of how to use ListMfaTotpDevicesRequest.
type ListMfaTotpDevicesRequest struct {

	// The OCID of the user.
	UserId *string `mandatory:"true" contributesTo:"path" name:"userId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for NAME is ascending. The NAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListMfaTotpDevicesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListMfaTotpDevicesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMfaTotpDevicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMfaTotpDevicesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMfaTotpDevicesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMfaTotpDevicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMfaTotpDevicesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMfaTotpDevicesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMfaTotpDevicesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMfaTotpDevicesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMfaTotpDevicesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMfaTotpDevicesResponse wrapper for the ListMfaTotpDevices operation
type ListMfaTotpDevicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []MfaTotpDeviceSummary instances
	Items []MfaTotpDeviceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMfaTotpDevicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMfaTotpDevicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMfaTotpDevicesSortByEnum Enum with underlying type: string
type ListMfaTotpDevicesSortByEnum string

// Set of constants representing the allowable values for ListMfaTotpDevicesSortByEnum
const (
	ListMfaTotpDevicesSortByTimecreated ListMfaTotpDevicesSortByEnum = "TIMECREATED"
	ListMfaTotpDevicesSortByName        ListMfaTotpDevicesSortByEnum = "NAME"
)

var mappingListMfaTotpDevicesSortByEnum = map[string]ListMfaTotpDevicesSortByEnum{
	"TIMECREATED": ListMfaTotpDevicesSortByTimecreated,
	"NAME":        ListMfaTotpDevicesSortByName,
}

var mappingListMfaTotpDevicesSortByEnumLowerCase = map[string]ListMfaTotpDevicesSortByEnum{
	"timecreated": ListMfaTotpDevicesSortByTimecreated,
	"name":        ListMfaTotpDevicesSortByName,
}

// GetListMfaTotpDevicesSortByEnumValues Enumerates the set of values for ListMfaTotpDevicesSortByEnum
func GetListMfaTotpDevicesSortByEnumValues() []ListMfaTotpDevicesSortByEnum {
	values := make([]ListMfaTotpDevicesSortByEnum, 0)
	for _, v := range mappingListMfaTotpDevicesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMfaTotpDevicesSortByEnumStringValues Enumerates the set of values in String for ListMfaTotpDevicesSortByEnum
func GetListMfaTotpDevicesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListMfaTotpDevicesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMfaTotpDevicesSortByEnum(val string) (ListMfaTotpDevicesSortByEnum, bool) {
	enum, ok := mappingListMfaTotpDevicesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMfaTotpDevicesSortOrderEnum Enum with underlying type: string
type ListMfaTotpDevicesSortOrderEnum string

// Set of constants representing the allowable values for ListMfaTotpDevicesSortOrderEnum
const (
	ListMfaTotpDevicesSortOrderAsc  ListMfaTotpDevicesSortOrderEnum = "ASC"
	ListMfaTotpDevicesSortOrderDesc ListMfaTotpDevicesSortOrderEnum = "DESC"
)

var mappingListMfaTotpDevicesSortOrderEnum = map[string]ListMfaTotpDevicesSortOrderEnum{
	"ASC":  ListMfaTotpDevicesSortOrderAsc,
	"DESC": ListMfaTotpDevicesSortOrderDesc,
}

var mappingListMfaTotpDevicesSortOrderEnumLowerCase = map[string]ListMfaTotpDevicesSortOrderEnum{
	"asc":  ListMfaTotpDevicesSortOrderAsc,
	"desc": ListMfaTotpDevicesSortOrderDesc,
}

// GetListMfaTotpDevicesSortOrderEnumValues Enumerates the set of values for ListMfaTotpDevicesSortOrderEnum
func GetListMfaTotpDevicesSortOrderEnumValues() []ListMfaTotpDevicesSortOrderEnum {
	values := make([]ListMfaTotpDevicesSortOrderEnum, 0)
	for _, v := range mappingListMfaTotpDevicesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMfaTotpDevicesSortOrderEnumStringValues Enumerates the set of values in String for ListMfaTotpDevicesSortOrderEnum
func GetListMfaTotpDevicesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMfaTotpDevicesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMfaTotpDevicesSortOrderEnum(val string) (ListMfaTotpDevicesSortOrderEnum, bool) {
	enum, ok := mappingListMfaTotpDevicesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
