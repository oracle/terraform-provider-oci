// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetDatabaseFleetHealthMetricsRequest wrapper for the GetDatabaseFleetHealthMetrics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetDatabaseFleetHealthMetrics.go.html to see an example of how to use GetDatabaseFleetHealthMetricsRequest.
type GetDatabaseFleetHealthMetricsRequest struct {

	// The baseline time for metrics comparison.
	CompareBaselineTime *string `mandatory:"true" contributesTo:"query" name:"compareBaselineTime"`

	// The target time for metrics comparison.
	CompareTargetTime *string `mandatory:"true" contributesTo:"query" name:"compareTargetTime"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	ManagedDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"managedDatabaseGroupId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The time window used for metrics comparison.
	CompareType GetDatabaseFleetHealthMetricsCompareTypeEnum `mandatory:"false" contributesTo:"query" name:"compareType" omitEmpty:"true"`

	// The filter used to retrieve a specific set of metrics by passing the desired metric names with a comma separator. Note that, by default, the service returns all supported metrics.
	FilterByMetricNames *string `mandatory:"false" contributesTo:"query" name:"filterByMetricNames"`

	// The filter used to filter the databases in the fleet by a specific Oracle Database type.
	FilterByDatabaseType *string `mandatory:"false" contributesTo:"query" name:"filterByDatabaseType"`

	// The filter used to filter the databases in the fleet by a specific Oracle Database subtype.
	FilterByDatabaseSubType *string `mandatory:"false" contributesTo:"query" name:"filterByDatabaseSubType"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy GetDatabaseFleetHealthMetricsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder GetDatabaseFleetHealthMetricsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The filter used to filter the databases in the fleet by a specific Oracle Database deployment type.
	FilterByDatabaseDeploymentType *string `mandatory:"false" contributesTo:"query" name:"filterByDatabaseDeploymentType"`

	// The filter used to filter the databases in the fleet by a specific Oracle Database version.
	FilterByDatabaseVersion *string `mandatory:"false" contributesTo:"query" name:"filterByDatabaseVersion"`

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

func (request GetDatabaseFleetHealthMetricsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDatabaseFleetHealthMetricsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetDatabaseFleetHealthMetricsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDatabaseFleetHealthMetricsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetDatabaseFleetHealthMetricsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetDatabaseFleetHealthMetricsCompareTypeEnum(string(request.CompareType)); !ok && request.CompareType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompareType: %s. Supported values are: %s.", request.CompareType, strings.Join(GetGetDatabaseFleetHealthMetricsCompareTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetDatabaseFleetHealthMetricsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetGetDatabaseFleetHealthMetricsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetDatabaseFleetHealthMetricsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetDatabaseFleetHealthMetricsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetDatabaseFleetHealthMetricsResponse wrapper for the GetDatabaseFleetHealthMetrics operation
type GetDatabaseFleetHealthMetricsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseFleetHealthMetrics instances
	DatabaseFleetHealthMetrics `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response GetDatabaseFleetHealthMetricsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDatabaseFleetHealthMetricsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetDatabaseFleetHealthMetricsCompareTypeEnum Enum with underlying type: string
type GetDatabaseFleetHealthMetricsCompareTypeEnum string

// Set of constants representing the allowable values for GetDatabaseFleetHealthMetricsCompareTypeEnum
const (
	GetDatabaseFleetHealthMetricsCompareTypeHour GetDatabaseFleetHealthMetricsCompareTypeEnum = "HOUR"
	GetDatabaseFleetHealthMetricsCompareTypeDay  GetDatabaseFleetHealthMetricsCompareTypeEnum = "DAY"
	GetDatabaseFleetHealthMetricsCompareTypeWeek GetDatabaseFleetHealthMetricsCompareTypeEnum = "WEEK"
)

var mappingGetDatabaseFleetHealthMetricsCompareTypeEnum = map[string]GetDatabaseFleetHealthMetricsCompareTypeEnum{
	"HOUR": GetDatabaseFleetHealthMetricsCompareTypeHour,
	"DAY":  GetDatabaseFleetHealthMetricsCompareTypeDay,
	"WEEK": GetDatabaseFleetHealthMetricsCompareTypeWeek,
}

var mappingGetDatabaseFleetHealthMetricsCompareTypeEnumLowerCase = map[string]GetDatabaseFleetHealthMetricsCompareTypeEnum{
	"hour": GetDatabaseFleetHealthMetricsCompareTypeHour,
	"day":  GetDatabaseFleetHealthMetricsCompareTypeDay,
	"week": GetDatabaseFleetHealthMetricsCompareTypeWeek,
}

// GetGetDatabaseFleetHealthMetricsCompareTypeEnumValues Enumerates the set of values for GetDatabaseFleetHealthMetricsCompareTypeEnum
func GetGetDatabaseFleetHealthMetricsCompareTypeEnumValues() []GetDatabaseFleetHealthMetricsCompareTypeEnum {
	values := make([]GetDatabaseFleetHealthMetricsCompareTypeEnum, 0)
	for _, v := range mappingGetDatabaseFleetHealthMetricsCompareTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDatabaseFleetHealthMetricsCompareTypeEnumStringValues Enumerates the set of values in String for GetDatabaseFleetHealthMetricsCompareTypeEnum
func GetGetDatabaseFleetHealthMetricsCompareTypeEnumStringValues() []string {
	return []string{
		"HOUR",
		"DAY",
		"WEEK",
	}
}

// GetMappingGetDatabaseFleetHealthMetricsCompareTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDatabaseFleetHealthMetricsCompareTypeEnum(val string) (GetDatabaseFleetHealthMetricsCompareTypeEnum, bool) {
	enum, ok := mappingGetDatabaseFleetHealthMetricsCompareTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetDatabaseFleetHealthMetricsSortByEnum Enum with underlying type: string
type GetDatabaseFleetHealthMetricsSortByEnum string

// Set of constants representing the allowable values for GetDatabaseFleetHealthMetricsSortByEnum
const (
	GetDatabaseFleetHealthMetricsSortByTimecreated GetDatabaseFleetHealthMetricsSortByEnum = "TIMECREATED"
	GetDatabaseFleetHealthMetricsSortByName        GetDatabaseFleetHealthMetricsSortByEnum = "NAME"
)

var mappingGetDatabaseFleetHealthMetricsSortByEnum = map[string]GetDatabaseFleetHealthMetricsSortByEnum{
	"TIMECREATED": GetDatabaseFleetHealthMetricsSortByTimecreated,
	"NAME":        GetDatabaseFleetHealthMetricsSortByName,
}

var mappingGetDatabaseFleetHealthMetricsSortByEnumLowerCase = map[string]GetDatabaseFleetHealthMetricsSortByEnum{
	"timecreated": GetDatabaseFleetHealthMetricsSortByTimecreated,
	"name":        GetDatabaseFleetHealthMetricsSortByName,
}

// GetGetDatabaseFleetHealthMetricsSortByEnumValues Enumerates the set of values for GetDatabaseFleetHealthMetricsSortByEnum
func GetGetDatabaseFleetHealthMetricsSortByEnumValues() []GetDatabaseFleetHealthMetricsSortByEnum {
	values := make([]GetDatabaseFleetHealthMetricsSortByEnum, 0)
	for _, v := range mappingGetDatabaseFleetHealthMetricsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDatabaseFleetHealthMetricsSortByEnumStringValues Enumerates the set of values in String for GetDatabaseFleetHealthMetricsSortByEnum
func GetGetDatabaseFleetHealthMetricsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingGetDatabaseFleetHealthMetricsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDatabaseFleetHealthMetricsSortByEnum(val string) (GetDatabaseFleetHealthMetricsSortByEnum, bool) {
	enum, ok := mappingGetDatabaseFleetHealthMetricsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetDatabaseFleetHealthMetricsSortOrderEnum Enum with underlying type: string
type GetDatabaseFleetHealthMetricsSortOrderEnum string

// Set of constants representing the allowable values for GetDatabaseFleetHealthMetricsSortOrderEnum
const (
	GetDatabaseFleetHealthMetricsSortOrderAsc  GetDatabaseFleetHealthMetricsSortOrderEnum = "ASC"
	GetDatabaseFleetHealthMetricsSortOrderDesc GetDatabaseFleetHealthMetricsSortOrderEnum = "DESC"
)

var mappingGetDatabaseFleetHealthMetricsSortOrderEnum = map[string]GetDatabaseFleetHealthMetricsSortOrderEnum{
	"ASC":  GetDatabaseFleetHealthMetricsSortOrderAsc,
	"DESC": GetDatabaseFleetHealthMetricsSortOrderDesc,
}

var mappingGetDatabaseFleetHealthMetricsSortOrderEnumLowerCase = map[string]GetDatabaseFleetHealthMetricsSortOrderEnum{
	"asc":  GetDatabaseFleetHealthMetricsSortOrderAsc,
	"desc": GetDatabaseFleetHealthMetricsSortOrderDesc,
}

// GetGetDatabaseFleetHealthMetricsSortOrderEnumValues Enumerates the set of values for GetDatabaseFleetHealthMetricsSortOrderEnum
func GetGetDatabaseFleetHealthMetricsSortOrderEnumValues() []GetDatabaseFleetHealthMetricsSortOrderEnum {
	values := make([]GetDatabaseFleetHealthMetricsSortOrderEnum, 0)
	for _, v := range mappingGetDatabaseFleetHealthMetricsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDatabaseFleetHealthMetricsSortOrderEnumStringValues Enumerates the set of values in String for GetDatabaseFleetHealthMetricsSortOrderEnum
func GetGetDatabaseFleetHealthMetricsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetDatabaseFleetHealthMetricsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDatabaseFleetHealthMetricsSortOrderEnum(val string) (GetDatabaseFleetHealthMetricsSortOrderEnum, bool) {
	enum, ok := mappingGetDatabaseFleetHealthMetricsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
