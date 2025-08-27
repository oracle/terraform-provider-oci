// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetDatabaseFleetHaOverviewMetricsRequest wrapper for the GetDatabaseFleetHaOverviewMetrics operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetDatabaseFleetHaOverviewMetrics.go.html to see an example of how to use GetDatabaseFleetHaOverviewMetricsRequest.
type GetDatabaseFleetHaOverviewMetricsRequest struct {

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	ManagedDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"managedDatabaseGroupId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The filter used to retrieve a specific set of metrics by passing the desired metric names with a comma separator. Note that, by default, the service returns all supported metrics.
	FilterByMetricNames *string `mandatory:"false" contributesTo:"query" name:"filterByMetricNames"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The
	// default sort order for `DATABASENAME` is ascending and it is case-sensitive.
	SortBy GetDatabaseFleetHaOverviewMetricsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder GetDatabaseFleetHaOverviewMetricsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A list of tag filters to apply.  Only resources with a defined tag matching the value will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagEquals []string `contributesTo:"query" name:"definedTagEquals" collectionFormat:"multi"`

	// A list of tag filters to apply.  Only resources with a freeform tag matching the value will be returned.
	// The key for each tag is "{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same tag name are interpreted as "OR".  Values for different tag names are interpreted as "AND".
	FreeformTagEquals []string `contributesTo:"query" name:"freeformTagEquals" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified defined tags exist will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.true" (for checking existence of a defined tag)
	// or "{namespace}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagExists []string `contributesTo:"query" name:"definedTagExists" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified freeform tags exist the value will be returned.
	// The key for each tag is "{tagName}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for different tag names are interpreted as "AND".
	FreeformTagExists []string `contributesTo:"query" name:"freeformTagExists" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDatabaseFleetHaOverviewMetricsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDatabaseFleetHaOverviewMetricsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetDatabaseFleetHaOverviewMetricsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDatabaseFleetHaOverviewMetricsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetDatabaseFleetHaOverviewMetricsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetDatabaseFleetHaOverviewMetricsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetGetDatabaseFleetHaOverviewMetricsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetDatabaseFleetHaOverviewMetricsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetDatabaseFleetHaOverviewMetricsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetDatabaseFleetHaOverviewMetricsResponse wrapper for the GetDatabaseFleetHaOverviewMetrics operation
type GetDatabaseFleetHaOverviewMetricsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseFleetHaOverviewMetrics instances
	DatabaseFleetHaOverviewMetrics `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response GetDatabaseFleetHaOverviewMetricsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDatabaseFleetHaOverviewMetricsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetDatabaseFleetHaOverviewMetricsSortByEnum Enum with underlying type: string
type GetDatabaseFleetHaOverviewMetricsSortByEnum string

// Set of constants representing the allowable values for GetDatabaseFleetHaOverviewMetricsSortByEnum
const (
	GetDatabaseFleetHaOverviewMetricsSortByDatabasename GetDatabaseFleetHaOverviewMetricsSortByEnum = "DATABASENAME"
)

var mappingGetDatabaseFleetHaOverviewMetricsSortByEnum = map[string]GetDatabaseFleetHaOverviewMetricsSortByEnum{
	"DATABASENAME": GetDatabaseFleetHaOverviewMetricsSortByDatabasename,
}

var mappingGetDatabaseFleetHaOverviewMetricsSortByEnumLowerCase = map[string]GetDatabaseFleetHaOverviewMetricsSortByEnum{
	"databasename": GetDatabaseFleetHaOverviewMetricsSortByDatabasename,
}

// GetGetDatabaseFleetHaOverviewMetricsSortByEnumValues Enumerates the set of values for GetDatabaseFleetHaOverviewMetricsSortByEnum
func GetGetDatabaseFleetHaOverviewMetricsSortByEnumValues() []GetDatabaseFleetHaOverviewMetricsSortByEnum {
	values := make([]GetDatabaseFleetHaOverviewMetricsSortByEnum, 0)
	for _, v := range mappingGetDatabaseFleetHaOverviewMetricsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDatabaseFleetHaOverviewMetricsSortByEnumStringValues Enumerates the set of values in String for GetDatabaseFleetHaOverviewMetricsSortByEnum
func GetGetDatabaseFleetHaOverviewMetricsSortByEnumStringValues() []string {
	return []string{
		"DATABASENAME",
	}
}

// GetMappingGetDatabaseFleetHaOverviewMetricsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDatabaseFleetHaOverviewMetricsSortByEnum(val string) (GetDatabaseFleetHaOverviewMetricsSortByEnum, bool) {
	enum, ok := mappingGetDatabaseFleetHaOverviewMetricsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetDatabaseFleetHaOverviewMetricsSortOrderEnum Enum with underlying type: string
type GetDatabaseFleetHaOverviewMetricsSortOrderEnum string

// Set of constants representing the allowable values for GetDatabaseFleetHaOverviewMetricsSortOrderEnum
const (
	GetDatabaseFleetHaOverviewMetricsSortOrderAsc  GetDatabaseFleetHaOverviewMetricsSortOrderEnum = "ASC"
	GetDatabaseFleetHaOverviewMetricsSortOrderDesc GetDatabaseFleetHaOverviewMetricsSortOrderEnum = "DESC"
)

var mappingGetDatabaseFleetHaOverviewMetricsSortOrderEnum = map[string]GetDatabaseFleetHaOverviewMetricsSortOrderEnum{
	"ASC":  GetDatabaseFleetHaOverviewMetricsSortOrderAsc,
	"DESC": GetDatabaseFleetHaOverviewMetricsSortOrderDesc,
}

var mappingGetDatabaseFleetHaOverviewMetricsSortOrderEnumLowerCase = map[string]GetDatabaseFleetHaOverviewMetricsSortOrderEnum{
	"asc":  GetDatabaseFleetHaOverviewMetricsSortOrderAsc,
	"desc": GetDatabaseFleetHaOverviewMetricsSortOrderDesc,
}

// GetGetDatabaseFleetHaOverviewMetricsSortOrderEnumValues Enumerates the set of values for GetDatabaseFleetHaOverviewMetricsSortOrderEnum
func GetGetDatabaseFleetHaOverviewMetricsSortOrderEnumValues() []GetDatabaseFleetHaOverviewMetricsSortOrderEnum {
	values := make([]GetDatabaseFleetHaOverviewMetricsSortOrderEnum, 0)
	for _, v := range mappingGetDatabaseFleetHaOverviewMetricsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDatabaseFleetHaOverviewMetricsSortOrderEnumStringValues Enumerates the set of values in String for GetDatabaseFleetHaOverviewMetricsSortOrderEnum
func GetGetDatabaseFleetHaOverviewMetricsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetDatabaseFleetHaOverviewMetricsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDatabaseFleetHaOverviewMetricsSortOrderEnum(val string) (GetDatabaseFleetHaOverviewMetricsSortOrderEnum, bool) {
	enum, ok := mappingGetDatabaseFleetHaOverviewMetricsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
