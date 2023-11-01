// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMonitoredResourcesRequest wrapper for the ListMonitoredResources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListMonitoredResources.go.html to see an example of how to use ListMonitoredResourcesRequest.
type ListMonitoredResourcesRequest struct {

	// The ID of the compartment in which data is listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources that match exact resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return resources which were impacted as part of this work request identifier.
	WorkRequestId *string `mandatory:"false" contributesTo:"query" name:"workRequestId"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for resources is ascending.
	SortBy ListMonitoredResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMonitoredResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMonitoredResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMonitoredResourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMonitoredResourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMonitoredResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMonitoredResourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMonitoredResourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMonitoredResourcesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMonitoredResourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMonitoredResourcesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMonitoredResourcesResponse wrapper for the ListMonitoredResources operation
type ListMonitoredResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MonitoredResourceCollection instances
	MonitoredResourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. The total number of items in the result.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListMonitoredResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMonitoredResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMonitoredResourcesSortByEnum Enum with underlying type: string
type ListMonitoredResourcesSortByEnum string

// Set of constants representing the allowable values for ListMonitoredResourcesSortByEnum
const (
	ListMonitoredResourcesSortByName        ListMonitoredResourcesSortByEnum = "NAME"
	ListMonitoredResourcesSortByTimeCreated ListMonitoredResourcesSortByEnum = "TIME_CREATED"
)

var mappingListMonitoredResourcesSortByEnum = map[string]ListMonitoredResourcesSortByEnum{
	"NAME":         ListMonitoredResourcesSortByName,
	"TIME_CREATED": ListMonitoredResourcesSortByTimeCreated,
}

var mappingListMonitoredResourcesSortByEnumLowerCase = map[string]ListMonitoredResourcesSortByEnum{
	"name":         ListMonitoredResourcesSortByName,
	"time_created": ListMonitoredResourcesSortByTimeCreated,
}

// GetListMonitoredResourcesSortByEnumValues Enumerates the set of values for ListMonitoredResourcesSortByEnum
func GetListMonitoredResourcesSortByEnumValues() []ListMonitoredResourcesSortByEnum {
	values := make([]ListMonitoredResourcesSortByEnum, 0)
	for _, v := range mappingListMonitoredResourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoredResourcesSortByEnumStringValues Enumerates the set of values in String for ListMonitoredResourcesSortByEnum
func GetListMonitoredResourcesSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIME_CREATED",
	}
}

// GetMappingListMonitoredResourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoredResourcesSortByEnum(val string) (ListMonitoredResourcesSortByEnum, bool) {
	enum, ok := mappingListMonitoredResourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMonitoredResourcesSortOrderEnum Enum with underlying type: string
type ListMonitoredResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListMonitoredResourcesSortOrderEnum
const (
	ListMonitoredResourcesSortOrderAsc  ListMonitoredResourcesSortOrderEnum = "ASC"
	ListMonitoredResourcesSortOrderDesc ListMonitoredResourcesSortOrderEnum = "DESC"
)

var mappingListMonitoredResourcesSortOrderEnum = map[string]ListMonitoredResourcesSortOrderEnum{
	"ASC":  ListMonitoredResourcesSortOrderAsc,
	"DESC": ListMonitoredResourcesSortOrderDesc,
}

var mappingListMonitoredResourcesSortOrderEnumLowerCase = map[string]ListMonitoredResourcesSortOrderEnum{
	"asc":  ListMonitoredResourcesSortOrderAsc,
	"desc": ListMonitoredResourcesSortOrderDesc,
}

// GetListMonitoredResourcesSortOrderEnumValues Enumerates the set of values for ListMonitoredResourcesSortOrderEnum
func GetListMonitoredResourcesSortOrderEnumValues() []ListMonitoredResourcesSortOrderEnum {
	values := make([]ListMonitoredResourcesSortOrderEnum, 0)
	for _, v := range mappingListMonitoredResourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoredResourcesSortOrderEnumStringValues Enumerates the set of values in String for ListMonitoredResourcesSortOrderEnum
func GetListMonitoredResourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMonitoredResourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoredResourcesSortOrderEnum(val string) (ListMonitoredResourcesSortOrderEnum, bool) {
	enum, ok := mappingListMonitoredResourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
