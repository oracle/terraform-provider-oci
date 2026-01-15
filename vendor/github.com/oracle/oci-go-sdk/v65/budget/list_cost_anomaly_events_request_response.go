// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package budget

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCostAnomalyEventsRequest wrapper for the ListCostAnomalyEvents operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/ListCostAnomalyEvents.go.html to see an example of how to use ListCostAnomalyEventsRequest.
type ListCostAnomalyEventsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListCostAnomalyEventsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. If not specified, the default is timeAnomalyEventDate.
	// The default sort order for timeAnomalyEventDate is DESC.
	// The default sort order for costAnomalyName is ASC in alphanumeric order.
	SortBy ListCostAnomalyEventsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique, non-changeable resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The cost monitor ocid.
	CostAnomalyMonitorId *string `mandatory:"false" contributesTo:"query" name:"costAnomalyMonitorId"`

	// The target tenantId ocid filter param.
	TargetTenantId []string `contributesTo:"query" name:"targetTenantId" collectionFormat:"csv"`

	// startDate for anomaly event date.
	TimeAnomalyEventStartDate *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeAnomalyEventStartDate"`

	// endDate for anomaly event date.
	TimeAnomalyEventEndDate *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeAnomalyEventEndDate"`

	// region of the anomaly event.
	Region []string `contributesTo:"query" name:"region" collectionFormat:"csv"`

	// cost impact (absolute) of the anomaly event.
	CostImpact *float64 `mandatory:"false" contributesTo:"query" name:"costImpact"`

	// cost impact (percentage) of the anomaly event.
	CostImpactPercentage *float64 `mandatory:"false" contributesTo:"query" name:"costImpactPercentage"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCostAnomalyEventsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCostAnomalyEventsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCostAnomalyEventsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCostAnomalyEventsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCostAnomalyEventsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCostAnomalyEventsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCostAnomalyEventsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCostAnomalyEventsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCostAnomalyEventsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCostAnomalyEventsResponse wrapper for the ListCostAnomalyEvents operation
type ListCostAnomalyEventsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CostAnomalyEventCollection instances
	CostAnomalyEventCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCostAnomalyEventsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCostAnomalyEventsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCostAnomalyEventsSortOrderEnum Enum with underlying type: string
type ListCostAnomalyEventsSortOrderEnum string

// Set of constants representing the allowable values for ListCostAnomalyEventsSortOrderEnum
const (
	ListCostAnomalyEventsSortOrderAsc  ListCostAnomalyEventsSortOrderEnum = "ASC"
	ListCostAnomalyEventsSortOrderDesc ListCostAnomalyEventsSortOrderEnum = "DESC"
)

var mappingListCostAnomalyEventsSortOrderEnum = map[string]ListCostAnomalyEventsSortOrderEnum{
	"ASC":  ListCostAnomalyEventsSortOrderAsc,
	"DESC": ListCostAnomalyEventsSortOrderDesc,
}

var mappingListCostAnomalyEventsSortOrderEnumLowerCase = map[string]ListCostAnomalyEventsSortOrderEnum{
	"asc":  ListCostAnomalyEventsSortOrderAsc,
	"desc": ListCostAnomalyEventsSortOrderDesc,
}

// GetListCostAnomalyEventsSortOrderEnumValues Enumerates the set of values for ListCostAnomalyEventsSortOrderEnum
func GetListCostAnomalyEventsSortOrderEnumValues() []ListCostAnomalyEventsSortOrderEnum {
	values := make([]ListCostAnomalyEventsSortOrderEnum, 0)
	for _, v := range mappingListCostAnomalyEventsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCostAnomalyEventsSortOrderEnumStringValues Enumerates the set of values in String for ListCostAnomalyEventsSortOrderEnum
func GetListCostAnomalyEventsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCostAnomalyEventsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCostAnomalyEventsSortOrderEnum(val string) (ListCostAnomalyEventsSortOrderEnum, bool) {
	enum, ok := mappingListCostAnomalyEventsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCostAnomalyEventsSortByEnum Enum with underlying type: string
type ListCostAnomalyEventsSortByEnum string

// Set of constants representing the allowable values for ListCostAnomalyEventsSortByEnum
const (
	ListCostAnomalyEventsSortByTimeanomalyeventdate ListCostAnomalyEventsSortByEnum = "timeAnomalyEventDate"
	ListCostAnomalyEventsSortByCostanomalyname      ListCostAnomalyEventsSortByEnum = "costAnomalyName"
	ListCostAnomalyEventsSortById                   ListCostAnomalyEventsSortByEnum = "id"
)

var mappingListCostAnomalyEventsSortByEnum = map[string]ListCostAnomalyEventsSortByEnum{
	"timeAnomalyEventDate": ListCostAnomalyEventsSortByTimeanomalyeventdate,
	"costAnomalyName":      ListCostAnomalyEventsSortByCostanomalyname,
	"id":                   ListCostAnomalyEventsSortById,
}

var mappingListCostAnomalyEventsSortByEnumLowerCase = map[string]ListCostAnomalyEventsSortByEnum{
	"timeanomalyeventdate": ListCostAnomalyEventsSortByTimeanomalyeventdate,
	"costanomalyname":      ListCostAnomalyEventsSortByCostanomalyname,
	"id":                   ListCostAnomalyEventsSortById,
}

// GetListCostAnomalyEventsSortByEnumValues Enumerates the set of values for ListCostAnomalyEventsSortByEnum
func GetListCostAnomalyEventsSortByEnumValues() []ListCostAnomalyEventsSortByEnum {
	values := make([]ListCostAnomalyEventsSortByEnum, 0)
	for _, v := range mappingListCostAnomalyEventsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCostAnomalyEventsSortByEnumStringValues Enumerates the set of values in String for ListCostAnomalyEventsSortByEnum
func GetListCostAnomalyEventsSortByEnumStringValues() []string {
	return []string{
		"timeAnomalyEventDate",
		"costAnomalyName",
		"id",
	}
}

// GetMappingListCostAnomalyEventsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCostAnomalyEventsSortByEnum(val string) (ListCostAnomalyEventsSortByEnum, bool) {
	enum, ok := mappingListCostAnomalyEventsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
