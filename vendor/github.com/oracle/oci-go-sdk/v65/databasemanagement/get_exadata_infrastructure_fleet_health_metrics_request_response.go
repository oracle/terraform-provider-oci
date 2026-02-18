// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetExadataInfrastructureFleetHealthMetricsRequest wrapper for the GetExadataInfrastructureFleetHealthMetrics operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExadataInfrastructureFleetHealthMetrics.go.html to see an example of how to use GetExadataInfrastructureFleetHealthMetricsRequest.
type GetExadataInfrastructureFleetHealthMetricsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The baseline time for metrics comparison.
	CompareBaselineTime *string `mandatory:"true" contributesTo:"query" name:"compareBaselineTime"`

	// The target time for metrics comparison.
	CompareTargetTime *string `mandatory:"true" contributesTo:"query" name:"compareTargetTime"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The time window used for metrics comparison.
	CompareType GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum `mandatory:"false" contributesTo:"query" name:"compareType" omitEmpty:"true"`

	// The filter used to filter the Exadata infrastructures in the fleet by a specific deployment type.
	FilterByExadataInfrastructureDeploymentType GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum `mandatory:"false" contributesTo:"query" name:"filterByExadataInfrastructureDeploymentType" omitEmpty:"true"`

	// The filter used to filter the Exadata infrastructure in the fleet by its lifecycle state.
	// If the parameter is not provided, Exdata infrastructures in any state are returned.
	FilterByExadataInfrastructureLifecycleState ExadataInfrastructureLifecycleStateValuesStateEnum `mandatory:"false" contributesTo:"query" name:"filterByExadataInfrastructureLifecycleState" omitEmpty:"true"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy GetExadataInfrastructureFleetHealthMetricsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder GetExadataInfrastructureFleetHealthMetricsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetExadataInfrastructureFleetHealthMetricsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetExadataInfrastructureFleetHealthMetricsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetExadataInfrastructureFleetHealthMetricsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetExadataInfrastructureFleetHealthMetricsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetExadataInfrastructureFleetHealthMetricsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetExadataInfrastructureFleetHealthMetricsCompareTypeEnum(string(request.CompareType)); !ok && request.CompareType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompareType: %s. Supported values are: %s.", request.CompareType, strings.Join(GetGetExadataInfrastructureFleetHealthMetricsCompareTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum(string(request.FilterByExadataInfrastructureDeploymentType)); !ok && request.FilterByExadataInfrastructureDeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FilterByExadataInfrastructureDeploymentType: %s. Supported values are: %s.", request.FilterByExadataInfrastructureDeploymentType, strings.Join(GetGetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadataInfrastructureLifecycleStateValuesStateEnum(string(request.FilterByExadataInfrastructureLifecycleState)); !ok && request.FilterByExadataInfrastructureLifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FilterByExadataInfrastructureLifecycleState: %s. Supported values are: %s.", request.FilterByExadataInfrastructureLifecycleState, strings.Join(GetExadataInfrastructureLifecycleStateValuesStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetExadataInfrastructureFleetHealthMetricsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetGetExadataInfrastructureFleetHealthMetricsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetExadataInfrastructureFleetHealthMetricsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetExadataInfrastructureFleetHealthMetricsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetExadataInfrastructureFleetHealthMetricsResponse wrapper for the GetExadataInfrastructureFleetHealthMetrics operation
type GetExadataInfrastructureFleetHealthMetricsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ExadataInfrastructureFleetHealthMetrics instance
	ExadataInfrastructureFleetHealthMetrics `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetExadataInfrastructureFleetHealthMetricsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetExadataInfrastructureFleetHealthMetricsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum Enum with underlying type: string
type GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum string

// Set of constants representing the allowable values for GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum
const (
	GetExadataInfrastructureFleetHealthMetricsCompareTypeHour GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum = "HOUR"
	GetExadataInfrastructureFleetHealthMetricsCompareTypeDay  GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum = "DAY"
	GetExadataInfrastructureFleetHealthMetricsCompareTypeWeek GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum = "WEEK"
)

var mappingGetExadataInfrastructureFleetHealthMetricsCompareTypeEnum = map[string]GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum{
	"HOUR": GetExadataInfrastructureFleetHealthMetricsCompareTypeHour,
	"DAY":  GetExadataInfrastructureFleetHealthMetricsCompareTypeDay,
	"WEEK": GetExadataInfrastructureFleetHealthMetricsCompareTypeWeek,
}

var mappingGetExadataInfrastructureFleetHealthMetricsCompareTypeEnumLowerCase = map[string]GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum{
	"hour": GetExadataInfrastructureFleetHealthMetricsCompareTypeHour,
	"day":  GetExadataInfrastructureFleetHealthMetricsCompareTypeDay,
	"week": GetExadataInfrastructureFleetHealthMetricsCompareTypeWeek,
}

// GetGetExadataInfrastructureFleetHealthMetricsCompareTypeEnumValues Enumerates the set of values for GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum
func GetGetExadataInfrastructureFleetHealthMetricsCompareTypeEnumValues() []GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum {
	values := make([]GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum, 0)
	for _, v := range mappingGetExadataInfrastructureFleetHealthMetricsCompareTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetExadataInfrastructureFleetHealthMetricsCompareTypeEnumStringValues Enumerates the set of values in String for GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum
func GetGetExadataInfrastructureFleetHealthMetricsCompareTypeEnumStringValues() []string {
	return []string{
		"HOUR",
		"DAY",
		"WEEK",
	}
}

// GetMappingGetExadataInfrastructureFleetHealthMetricsCompareTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetExadataInfrastructureFleetHealthMetricsCompareTypeEnum(val string) (GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum, bool) {
	enum, ok := mappingGetExadataInfrastructureFleetHealthMetricsCompareTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum Enum with underlying type: string
type GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum string

// Set of constants representing the allowable values for GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum
const (
	GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeOnpremise GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum = "ONPREMISE"
	GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeExadata   GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum = "EXADATA"
	GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeExadataCc GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum = "EXADATA_CC"
)

var mappingGetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum = map[string]GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum{
	"ONPREMISE":  GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeOnpremise,
	"EXADATA":    GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeExadata,
	"EXADATA_CC": GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeExadataCc,
}

var mappingGetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnumLowerCase = map[string]GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum{
	"onpremise":  GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeOnpremise,
	"exadata":    GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeExadata,
	"exadata_cc": GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeExadataCc,
}

// GetGetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnumValues Enumerates the set of values for GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum
func GetGetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnumValues() []GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum {
	values := make([]GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum, 0)
	for _, v := range mappingGetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnumStringValues Enumerates the set of values in String for GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum
func GetGetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnumStringValues() []string {
	return []string{
		"ONPREMISE",
		"EXADATA",
		"EXADATA_CC",
	}
}

// GetMappingGetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum(val string) (GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum, bool) {
	enum, ok := mappingGetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetExadataInfrastructureFleetHealthMetricsSortByEnum Enum with underlying type: string
type GetExadataInfrastructureFleetHealthMetricsSortByEnum string

// Set of constants representing the allowable values for GetExadataInfrastructureFleetHealthMetricsSortByEnum
const (
	GetExadataInfrastructureFleetHealthMetricsSortByTimecreated GetExadataInfrastructureFleetHealthMetricsSortByEnum = "TIMECREATED"
	GetExadataInfrastructureFleetHealthMetricsSortByName        GetExadataInfrastructureFleetHealthMetricsSortByEnum = "NAME"
)

var mappingGetExadataInfrastructureFleetHealthMetricsSortByEnum = map[string]GetExadataInfrastructureFleetHealthMetricsSortByEnum{
	"TIMECREATED": GetExadataInfrastructureFleetHealthMetricsSortByTimecreated,
	"NAME":        GetExadataInfrastructureFleetHealthMetricsSortByName,
}

var mappingGetExadataInfrastructureFleetHealthMetricsSortByEnumLowerCase = map[string]GetExadataInfrastructureFleetHealthMetricsSortByEnum{
	"timecreated": GetExadataInfrastructureFleetHealthMetricsSortByTimecreated,
	"name":        GetExadataInfrastructureFleetHealthMetricsSortByName,
}

// GetGetExadataInfrastructureFleetHealthMetricsSortByEnumValues Enumerates the set of values for GetExadataInfrastructureFleetHealthMetricsSortByEnum
func GetGetExadataInfrastructureFleetHealthMetricsSortByEnumValues() []GetExadataInfrastructureFleetHealthMetricsSortByEnum {
	values := make([]GetExadataInfrastructureFleetHealthMetricsSortByEnum, 0)
	for _, v := range mappingGetExadataInfrastructureFleetHealthMetricsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetGetExadataInfrastructureFleetHealthMetricsSortByEnumStringValues Enumerates the set of values in String for GetExadataInfrastructureFleetHealthMetricsSortByEnum
func GetGetExadataInfrastructureFleetHealthMetricsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingGetExadataInfrastructureFleetHealthMetricsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetExadataInfrastructureFleetHealthMetricsSortByEnum(val string) (GetExadataInfrastructureFleetHealthMetricsSortByEnum, bool) {
	enum, ok := mappingGetExadataInfrastructureFleetHealthMetricsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetExadataInfrastructureFleetHealthMetricsSortOrderEnum Enum with underlying type: string
type GetExadataInfrastructureFleetHealthMetricsSortOrderEnum string

// Set of constants representing the allowable values for GetExadataInfrastructureFleetHealthMetricsSortOrderEnum
const (
	GetExadataInfrastructureFleetHealthMetricsSortOrderAsc  GetExadataInfrastructureFleetHealthMetricsSortOrderEnum = "ASC"
	GetExadataInfrastructureFleetHealthMetricsSortOrderDesc GetExadataInfrastructureFleetHealthMetricsSortOrderEnum = "DESC"
)

var mappingGetExadataInfrastructureFleetHealthMetricsSortOrderEnum = map[string]GetExadataInfrastructureFleetHealthMetricsSortOrderEnum{
	"ASC":  GetExadataInfrastructureFleetHealthMetricsSortOrderAsc,
	"DESC": GetExadataInfrastructureFleetHealthMetricsSortOrderDesc,
}

var mappingGetExadataInfrastructureFleetHealthMetricsSortOrderEnumLowerCase = map[string]GetExadataInfrastructureFleetHealthMetricsSortOrderEnum{
	"asc":  GetExadataInfrastructureFleetHealthMetricsSortOrderAsc,
	"desc": GetExadataInfrastructureFleetHealthMetricsSortOrderDesc,
}

// GetGetExadataInfrastructureFleetHealthMetricsSortOrderEnumValues Enumerates the set of values for GetExadataInfrastructureFleetHealthMetricsSortOrderEnum
func GetGetExadataInfrastructureFleetHealthMetricsSortOrderEnumValues() []GetExadataInfrastructureFleetHealthMetricsSortOrderEnum {
	values := make([]GetExadataInfrastructureFleetHealthMetricsSortOrderEnum, 0)
	for _, v := range mappingGetExadataInfrastructureFleetHealthMetricsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetExadataInfrastructureFleetHealthMetricsSortOrderEnumStringValues Enumerates the set of values in String for GetExadataInfrastructureFleetHealthMetricsSortOrderEnum
func GetGetExadataInfrastructureFleetHealthMetricsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetExadataInfrastructureFleetHealthMetricsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetExadataInfrastructureFleetHealthMetricsSortOrderEnum(val string) (GetExadataInfrastructureFleetHealthMetricsSortOrderEnum, bool) {
	enum, ok := mappingGetExadataInfrastructureFleetHealthMetricsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
