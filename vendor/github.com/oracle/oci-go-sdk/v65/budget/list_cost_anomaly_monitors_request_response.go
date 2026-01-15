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

// ListCostAnomalyMonitorsRequest wrapper for the ListCostAnomalyMonitors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/ListCostAnomalyMonitors.go.html to see an example of how to use ListCostAnomalyMonitorsRequest.
type ListCostAnomalyMonitorsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListCostAnomalyMonitorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. If not specified, the default is timeCreated.
	// The default sort order for timeCreated is DESC.
	// The default sort order for displayName is ASC in alphanumeric order.
	SortBy ListCostAnomalyMonitorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The current state of the cost monitor.
	LifecycleState CostAnomalyMonitorLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique, non-changeable resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The target tenantId ocid filter param.
	TargetTenantId []string `contributesTo:"query" name:"targetTenantId" collectionFormat:"csv"`

	// Cost Anomaly Monitor target resource filter region.
	Region []string `contributesTo:"query" name:"region" collectionFormat:"csv"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCostAnomalyMonitorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCostAnomalyMonitorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCostAnomalyMonitorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCostAnomalyMonitorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCostAnomalyMonitorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCostAnomalyMonitorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCostAnomalyMonitorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCostAnomalyMonitorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCostAnomalyMonitorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCostAnomalyMonitorLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCostAnomalyMonitorLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCostAnomalyMonitorsResponse wrapper for the ListCostAnomalyMonitors operation
type ListCostAnomalyMonitorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CostAnomalyMonitorCollection instances
	CostAnomalyMonitorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCostAnomalyMonitorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCostAnomalyMonitorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCostAnomalyMonitorsSortOrderEnum Enum with underlying type: string
type ListCostAnomalyMonitorsSortOrderEnum string

// Set of constants representing the allowable values for ListCostAnomalyMonitorsSortOrderEnum
const (
	ListCostAnomalyMonitorsSortOrderAsc  ListCostAnomalyMonitorsSortOrderEnum = "ASC"
	ListCostAnomalyMonitorsSortOrderDesc ListCostAnomalyMonitorsSortOrderEnum = "DESC"
)

var mappingListCostAnomalyMonitorsSortOrderEnum = map[string]ListCostAnomalyMonitorsSortOrderEnum{
	"ASC":  ListCostAnomalyMonitorsSortOrderAsc,
	"DESC": ListCostAnomalyMonitorsSortOrderDesc,
}

var mappingListCostAnomalyMonitorsSortOrderEnumLowerCase = map[string]ListCostAnomalyMonitorsSortOrderEnum{
	"asc":  ListCostAnomalyMonitorsSortOrderAsc,
	"desc": ListCostAnomalyMonitorsSortOrderDesc,
}

// GetListCostAnomalyMonitorsSortOrderEnumValues Enumerates the set of values for ListCostAnomalyMonitorsSortOrderEnum
func GetListCostAnomalyMonitorsSortOrderEnumValues() []ListCostAnomalyMonitorsSortOrderEnum {
	values := make([]ListCostAnomalyMonitorsSortOrderEnum, 0)
	for _, v := range mappingListCostAnomalyMonitorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCostAnomalyMonitorsSortOrderEnumStringValues Enumerates the set of values in String for ListCostAnomalyMonitorsSortOrderEnum
func GetListCostAnomalyMonitorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCostAnomalyMonitorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCostAnomalyMonitorsSortOrderEnum(val string) (ListCostAnomalyMonitorsSortOrderEnum, bool) {
	enum, ok := mappingListCostAnomalyMonitorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCostAnomalyMonitorsSortByEnum Enum with underlying type: string
type ListCostAnomalyMonitorsSortByEnum string

// Set of constants representing the allowable values for ListCostAnomalyMonitorsSortByEnum
const (
	ListCostAnomalyMonitorsSortByTimecreated ListCostAnomalyMonitorsSortByEnum = "timeCreated"
	ListCostAnomalyMonitorsSortByName        ListCostAnomalyMonitorsSortByEnum = "name"
	ListCostAnomalyMonitorsSortById          ListCostAnomalyMonitorsSortByEnum = "id"
)

var mappingListCostAnomalyMonitorsSortByEnum = map[string]ListCostAnomalyMonitorsSortByEnum{
	"timeCreated": ListCostAnomalyMonitorsSortByTimecreated,
	"name":        ListCostAnomalyMonitorsSortByName,
	"id":          ListCostAnomalyMonitorsSortById,
}

var mappingListCostAnomalyMonitorsSortByEnumLowerCase = map[string]ListCostAnomalyMonitorsSortByEnum{
	"timecreated": ListCostAnomalyMonitorsSortByTimecreated,
	"name":        ListCostAnomalyMonitorsSortByName,
	"id":          ListCostAnomalyMonitorsSortById,
}

// GetListCostAnomalyMonitorsSortByEnumValues Enumerates the set of values for ListCostAnomalyMonitorsSortByEnum
func GetListCostAnomalyMonitorsSortByEnumValues() []ListCostAnomalyMonitorsSortByEnum {
	values := make([]ListCostAnomalyMonitorsSortByEnum, 0)
	for _, v := range mappingListCostAnomalyMonitorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCostAnomalyMonitorsSortByEnumStringValues Enumerates the set of values in String for ListCostAnomalyMonitorsSortByEnum
func GetListCostAnomalyMonitorsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
		"id",
	}
}

// GetMappingListCostAnomalyMonitorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCostAnomalyMonitorsSortByEnum(val string) (ListCostAnomalyMonitorsSortByEnum, bool) {
	enum, ok := mappingListCostAnomalyMonitorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
