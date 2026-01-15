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

// SummarizeCostAnomalyEventAnalyticsRequest wrapper for the SummarizeCostAnomalyEventAnalytics operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/SummarizeCostAnomalyEventAnalytics.go.html to see an example of how to use SummarizeCostAnomalyEventAnalyticsRequest.
type SummarizeCostAnomalyEventAnalyticsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder SummarizeCostAnomalyEventAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. If not specified, the default is timeAnomalyEventDate.
	// The default sort order for timeAnomalyEventDate is DESC.
	// The default sort order for costAnomalyName is ASC in alphanumeric order.
	SortBy SummarizeCostAnomalyEventAnalyticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

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

func (request SummarizeCostAnomalyEventAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeCostAnomalyEventAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeCostAnomalyEventAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeCostAnomalyEventAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeCostAnomalyEventAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeCostAnomalyEventAnalyticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeCostAnomalyEventAnalyticsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeCostAnomalyEventAnalyticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeCostAnomalyEventAnalyticsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeCostAnomalyEventAnalyticsResponse wrapper for the SummarizeCostAnomalyEventAnalytics operation
type SummarizeCostAnomalyEventAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CostAnomalyEventAnalyticCollection instances
	CostAnomalyEventAnalyticCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeCostAnomalyEventAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeCostAnomalyEventAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeCostAnomalyEventAnalyticsSortOrderEnum Enum with underlying type: string
type SummarizeCostAnomalyEventAnalyticsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeCostAnomalyEventAnalyticsSortOrderEnum
const (
	SummarizeCostAnomalyEventAnalyticsSortOrderAsc  SummarizeCostAnomalyEventAnalyticsSortOrderEnum = "ASC"
	SummarizeCostAnomalyEventAnalyticsSortOrderDesc SummarizeCostAnomalyEventAnalyticsSortOrderEnum = "DESC"
)

var mappingSummarizeCostAnomalyEventAnalyticsSortOrderEnum = map[string]SummarizeCostAnomalyEventAnalyticsSortOrderEnum{
	"ASC":  SummarizeCostAnomalyEventAnalyticsSortOrderAsc,
	"DESC": SummarizeCostAnomalyEventAnalyticsSortOrderDesc,
}

var mappingSummarizeCostAnomalyEventAnalyticsSortOrderEnumLowerCase = map[string]SummarizeCostAnomalyEventAnalyticsSortOrderEnum{
	"asc":  SummarizeCostAnomalyEventAnalyticsSortOrderAsc,
	"desc": SummarizeCostAnomalyEventAnalyticsSortOrderDesc,
}

// GetSummarizeCostAnomalyEventAnalyticsSortOrderEnumValues Enumerates the set of values for SummarizeCostAnomalyEventAnalyticsSortOrderEnum
func GetSummarizeCostAnomalyEventAnalyticsSortOrderEnumValues() []SummarizeCostAnomalyEventAnalyticsSortOrderEnum {
	values := make([]SummarizeCostAnomalyEventAnalyticsSortOrderEnum, 0)
	for _, v := range mappingSummarizeCostAnomalyEventAnalyticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeCostAnomalyEventAnalyticsSortOrderEnumStringValues Enumerates the set of values in String for SummarizeCostAnomalyEventAnalyticsSortOrderEnum
func GetSummarizeCostAnomalyEventAnalyticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeCostAnomalyEventAnalyticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeCostAnomalyEventAnalyticsSortOrderEnum(val string) (SummarizeCostAnomalyEventAnalyticsSortOrderEnum, bool) {
	enum, ok := mappingSummarizeCostAnomalyEventAnalyticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeCostAnomalyEventAnalyticsSortByEnum Enum with underlying type: string
type SummarizeCostAnomalyEventAnalyticsSortByEnum string

// Set of constants representing the allowable values for SummarizeCostAnomalyEventAnalyticsSortByEnum
const (
	SummarizeCostAnomalyEventAnalyticsSortByTimeanomalyeventdate SummarizeCostAnomalyEventAnalyticsSortByEnum = "timeAnomalyEventDate"
	SummarizeCostAnomalyEventAnalyticsSortByCostanomalyname      SummarizeCostAnomalyEventAnalyticsSortByEnum = "costAnomalyName"
	SummarizeCostAnomalyEventAnalyticsSortById                   SummarizeCostAnomalyEventAnalyticsSortByEnum = "id"
)

var mappingSummarizeCostAnomalyEventAnalyticsSortByEnum = map[string]SummarizeCostAnomalyEventAnalyticsSortByEnum{
	"timeAnomalyEventDate": SummarizeCostAnomalyEventAnalyticsSortByTimeanomalyeventdate,
	"costAnomalyName":      SummarizeCostAnomalyEventAnalyticsSortByCostanomalyname,
	"id":                   SummarizeCostAnomalyEventAnalyticsSortById,
}

var mappingSummarizeCostAnomalyEventAnalyticsSortByEnumLowerCase = map[string]SummarizeCostAnomalyEventAnalyticsSortByEnum{
	"timeanomalyeventdate": SummarizeCostAnomalyEventAnalyticsSortByTimeanomalyeventdate,
	"costanomalyname":      SummarizeCostAnomalyEventAnalyticsSortByCostanomalyname,
	"id":                   SummarizeCostAnomalyEventAnalyticsSortById,
}

// GetSummarizeCostAnomalyEventAnalyticsSortByEnumValues Enumerates the set of values for SummarizeCostAnomalyEventAnalyticsSortByEnum
func GetSummarizeCostAnomalyEventAnalyticsSortByEnumValues() []SummarizeCostAnomalyEventAnalyticsSortByEnum {
	values := make([]SummarizeCostAnomalyEventAnalyticsSortByEnum, 0)
	for _, v := range mappingSummarizeCostAnomalyEventAnalyticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeCostAnomalyEventAnalyticsSortByEnumStringValues Enumerates the set of values in String for SummarizeCostAnomalyEventAnalyticsSortByEnum
func GetSummarizeCostAnomalyEventAnalyticsSortByEnumStringValues() []string {
	return []string{
		"timeAnomalyEventDate",
		"costAnomalyName",
		"id",
	}
}

// GetMappingSummarizeCostAnomalyEventAnalyticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeCostAnomalyEventAnalyticsSortByEnum(val string) (SummarizeCostAnomalyEventAnalyticsSortByEnum, bool) {
	enum, ok := mappingSummarizeCostAnomalyEventAnalyticsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
