// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourceanalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMonitoredRegionsRequest wrapper for the ListMonitoredRegions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/ListMonitoredRegions.go.html to see an example of how to use ListMonitoredRegionsRequest.
type ListMonitoredRegionsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a ResourceAnalyticsInstance.
	ResourceAnalyticsInstanceId *string `mandatory:"false" contributesTo:"query" name:"resourceAnalyticsInstanceId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState MonitoredRegionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the MonitoredRegion.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMonitoredRegionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `TIME_CREATED`
	// is descending. Default order for `REGION_ID` is ascending.
	SortBy ListMonitoredRegionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMonitoredRegionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMonitoredRegionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMonitoredRegionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMonitoredRegionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMonitoredRegionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMonitoredRegionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMonitoredRegionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMonitoredRegionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMonitoredRegionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMonitoredRegionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMonitoredRegionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMonitoredRegionsResponse wrapper for the ListMonitoredRegions operation
type ListMonitoredRegionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MonitoredRegionCollection instances
	MonitoredRegionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMonitoredRegionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMonitoredRegionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMonitoredRegionsSortOrderEnum Enum with underlying type: string
type ListMonitoredRegionsSortOrderEnum string

// Set of constants representing the allowable values for ListMonitoredRegionsSortOrderEnum
const (
	ListMonitoredRegionsSortOrderAsc  ListMonitoredRegionsSortOrderEnum = "ASC"
	ListMonitoredRegionsSortOrderDesc ListMonitoredRegionsSortOrderEnum = "DESC"
)

var mappingListMonitoredRegionsSortOrderEnum = map[string]ListMonitoredRegionsSortOrderEnum{
	"ASC":  ListMonitoredRegionsSortOrderAsc,
	"DESC": ListMonitoredRegionsSortOrderDesc,
}

var mappingListMonitoredRegionsSortOrderEnumLowerCase = map[string]ListMonitoredRegionsSortOrderEnum{
	"asc":  ListMonitoredRegionsSortOrderAsc,
	"desc": ListMonitoredRegionsSortOrderDesc,
}

// GetListMonitoredRegionsSortOrderEnumValues Enumerates the set of values for ListMonitoredRegionsSortOrderEnum
func GetListMonitoredRegionsSortOrderEnumValues() []ListMonitoredRegionsSortOrderEnum {
	values := make([]ListMonitoredRegionsSortOrderEnum, 0)
	for _, v := range mappingListMonitoredRegionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoredRegionsSortOrderEnumStringValues Enumerates the set of values in String for ListMonitoredRegionsSortOrderEnum
func GetListMonitoredRegionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMonitoredRegionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoredRegionsSortOrderEnum(val string) (ListMonitoredRegionsSortOrderEnum, bool) {
	enum, ok := mappingListMonitoredRegionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMonitoredRegionsSortByEnum Enum with underlying type: string
type ListMonitoredRegionsSortByEnum string

// Set of constants representing the allowable values for ListMonitoredRegionsSortByEnum
const (
	ListMonitoredRegionsSortByTimeCreated ListMonitoredRegionsSortByEnum = "TIME_CREATED"
	ListMonitoredRegionsSortByRegionId    ListMonitoredRegionsSortByEnum = "REGION_ID"
)

var mappingListMonitoredRegionsSortByEnum = map[string]ListMonitoredRegionsSortByEnum{
	"TIME_CREATED": ListMonitoredRegionsSortByTimeCreated,
	"REGION_ID":    ListMonitoredRegionsSortByRegionId,
}

var mappingListMonitoredRegionsSortByEnumLowerCase = map[string]ListMonitoredRegionsSortByEnum{
	"time_created": ListMonitoredRegionsSortByTimeCreated,
	"region_id":    ListMonitoredRegionsSortByRegionId,
}

// GetListMonitoredRegionsSortByEnumValues Enumerates the set of values for ListMonitoredRegionsSortByEnum
func GetListMonitoredRegionsSortByEnumValues() []ListMonitoredRegionsSortByEnum {
	values := make([]ListMonitoredRegionsSortByEnum, 0)
	for _, v := range mappingListMonitoredRegionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoredRegionsSortByEnumStringValues Enumerates the set of values in String for ListMonitoredRegionsSortByEnum
func GetListMonitoredRegionsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"REGION_ID",
	}
}

// GetMappingListMonitoredRegionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoredRegionsSortByEnum(val string) (ListMonitoredRegionsSortByEnum, bool) {
	enum, ok := mappingListMonitoredRegionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
