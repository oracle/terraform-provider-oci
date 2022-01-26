// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListMonitorsRequest wrapper for the ListMonitors operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmsynthetics/ListMonitors.go.html to see an example of how to use ListMonitorsRequest.
type ListMonitorsRequest struct {

	// The APM domain ID the request is intended for.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only monitors using scriptId.
	ScriptId *string `mandatory:"false" contributesTo:"query" name:"scriptId"`

	// A filter to return only monitors that match the given monitor type.
	// Supported values are SCRIPTED_BROWSER, BROWSER, SCRIPTED_REST and REST.
	MonitorType *string `mandatory:"false" contributesTo:"query" name:"monitorType"`

	// A filter to return only monitors that match the status given.
	Status ListMonitorsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). Default sort order is ascending.
	SortOrder ListMonitorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order of displayName is ascending.
	// Default order of timeCreated and timeUpdated is descending.
	// The displayName sort by is case insensitive.
	SortBy ListMonitorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMonitorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMonitorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMonitorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMonitorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListMonitorsResponse wrapper for the ListMonitors operation
type ListMonitorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MonitorCollection instances
	MonitorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMonitorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMonitorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMonitorsStatusEnum Enum with underlying type: string
type ListMonitorsStatusEnum string

// Set of constants representing the allowable values for ListMonitorsStatusEnum
const (
	ListMonitorsStatusEnabled  ListMonitorsStatusEnum = "ENABLED"
	ListMonitorsStatusDisabled ListMonitorsStatusEnum = "DISABLED"
	ListMonitorsStatusInvalid  ListMonitorsStatusEnum = "INVALID"
)

var mappingListMonitorsStatus = map[string]ListMonitorsStatusEnum{
	"ENABLED":  ListMonitorsStatusEnabled,
	"DISABLED": ListMonitorsStatusDisabled,
	"INVALID":  ListMonitorsStatusInvalid,
}

// GetListMonitorsStatusEnumValues Enumerates the set of values for ListMonitorsStatusEnum
func GetListMonitorsStatusEnumValues() []ListMonitorsStatusEnum {
	values := make([]ListMonitorsStatusEnum, 0)
	for _, v := range mappingListMonitorsStatus {
		values = append(values, v)
	}
	return values
}

// ListMonitorsSortOrderEnum Enum with underlying type: string
type ListMonitorsSortOrderEnum string

// Set of constants representing the allowable values for ListMonitorsSortOrderEnum
const (
	ListMonitorsSortOrderAsc  ListMonitorsSortOrderEnum = "ASC"
	ListMonitorsSortOrderDesc ListMonitorsSortOrderEnum = "DESC"
)

var mappingListMonitorsSortOrder = map[string]ListMonitorsSortOrderEnum{
	"ASC":  ListMonitorsSortOrderAsc,
	"DESC": ListMonitorsSortOrderDesc,
}

// GetListMonitorsSortOrderEnumValues Enumerates the set of values for ListMonitorsSortOrderEnum
func GetListMonitorsSortOrderEnumValues() []ListMonitorsSortOrderEnum {
	values := make([]ListMonitorsSortOrderEnum, 0)
	for _, v := range mappingListMonitorsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListMonitorsSortByEnum Enum with underlying type: string
type ListMonitorsSortByEnum string

// Set of constants representing the allowable values for ListMonitorsSortByEnum
const (
	ListMonitorsSortByDisplayname ListMonitorsSortByEnum = "displayName"
	ListMonitorsSortByTimecreated ListMonitorsSortByEnum = "timeCreated"
	ListMonitorsSortByTimeupdated ListMonitorsSortByEnum = "timeUpdated"
	ListMonitorsSortByStatus      ListMonitorsSortByEnum = "status"
	ListMonitorsSortByMonitortype ListMonitorsSortByEnum = "monitorType"
)

var mappingListMonitorsSortBy = map[string]ListMonitorsSortByEnum{
	"displayName": ListMonitorsSortByDisplayname,
	"timeCreated": ListMonitorsSortByTimecreated,
	"timeUpdated": ListMonitorsSortByTimeupdated,
	"status":      ListMonitorsSortByStatus,
	"monitorType": ListMonitorsSortByMonitortype,
}

// GetListMonitorsSortByEnumValues Enumerates the set of values for ListMonitorsSortByEnum
func GetListMonitorsSortByEnumValues() []ListMonitorsSortByEnum {
	values := make([]ListMonitorsSortByEnum, 0)
	for _, v := range mappingListMonitorsSortBy {
		values = append(values, v)
	}
	return values
}
