// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeHostInsightResourceCapacityTrendRequest wrapper for the SummarizeHostInsightResourceCapacityTrend operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceCapacityTrend.go.html to see an example of how to use SummarizeHostInsightResourceCapacityTrendRequest.
type SummarizeHostInsightResourceCapacityTrendRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter by host resource metric.
	// Supported values are CPU, MEMORY, LOGICAL_MEMORY, STORAGE and NETWORK.
	ResourceMetric *string `mandatory:"true" contributesTo:"query" name:"resourceMetric"`

	// Specify time period in ISO 8601 format with respect to current time.
	// Default is last 30 days represented by P30D.
	// If timeInterval is specified, then timeIntervalStart and timeIntervalEnd will be ignored.
	// Examples  P90D (last 90 days), P4W (last 4 weeks), P2M (last 2 months), P1Y (last 12 months), . Maximum value allowed is 25 months prior to current time (P25M).
	AnalysisTimeInterval *string `mandatory:"false" contributesTo:"query" name:"analysisTimeInterval"`

	// Analysis start time in UTC in ISO 8601 format(inclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// The minimum allowed value is 2 years prior to the current day.
	// timeIntervalStart and timeIntervalEnd parameters are used together.
	// If analysisTimeInterval is specified, this parameter is ignored.
	TimeIntervalStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalStart"`

	// Analysis end time in UTC in ISO 8601 format(exclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// timeIntervalStart and timeIntervalEnd are used together.
	// If timeIntervalEnd is not specified, current time is used as timeIntervalEnd.
	TimeIntervalEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalEnd"`

	// Filter by one or more platform types.
	// Supported platformType(s) for MACS-managed external host insight: [LINUX, SOLARIS, WINDOWS].
	// Supported platformType(s) for MACS-managed cloud host insight: [LINUX].
	// Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS, ZLINUX, WINDOWS, AIX, HP-UX].
	PlatformType []SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum `contributesTo:"query" name:"platformType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of host insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Filter by utilization level by the following buckets:
	//   - HIGH_UTILIZATION: DBs with utilization greater or equal than 75.
	//   - LOW_UTILIZATION: DBs with utilization lower than 25.
	//   - MEDIUM_HIGH_UTILIZATION: DBs with utilization greater or equal than 50 but lower than 75.
	//   - MEDIUM_LOW_UTILIZATION: DBs with utilization greater or equal than 25 but lower than 50.
	UtilizationLevel SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum `mandatory:"false" contributesTo:"query" name:"utilizationLevel" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeHostInsightResourceCapacityTrendSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sorts using end timestamp or capacity
	SortBy SummarizeHostInsightResourceCapacityTrendSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

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

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Filter by one or more host types.
	// Possible values are CLOUD-HOST, EXTERNAL-HOST, COMANAGED-VM-HOST, COMANAGED-BM-HOST, COMANAGED-EXACS-HOST
	HostType []string `contributesTo:"query" name:"hostType" collectionFormat:"multi"`

	// Optional OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host (Compute Id)
	HostId *string `mandatory:"false" contributesTo:"query" name:"hostId"`

	// Optional list of Exadata Insight VM cluster name.
	VmclusterName []string `contributesTo:"query" name:"vmclusterName" collectionFormat:"multi"`

	// Percent value in which a resource metric is considered highly utilized.
	HighUtilizationThreshold *int `mandatory:"false" contributesTo:"query" name:"highUtilizationThreshold"`

	// Percent value in which a resource metric is considered low utilized.
	LowUtilizationThreshold *int `mandatory:"false" contributesTo:"query" name:"lowUtilizationThreshold"`

	// Resource Status
	Status []ResourceStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeHostInsightResourceCapacityTrendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeHostInsightResourceCapacityTrendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeHostInsightResourceCapacityTrendRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeHostInsightResourceCapacityTrendRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeHostInsightResourceCapacityTrendRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.PlatformType {
		if _, ok := GetMappingSummarizeHostInsightResourceCapacityTrendPlatformTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", val, strings.Join(GetSummarizeHostInsightResourceCapacityTrendPlatformTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum(string(request.UtilizationLevel)); !ok && request.UtilizationLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UtilizationLevel: %s. Supported values are: %s.", request.UtilizationLevel, strings.Join(GetSummarizeHostInsightResourceCapacityTrendUtilizationLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeHostInsightResourceCapacityTrendSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeHostInsightResourceCapacityTrendSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeHostInsightResourceCapacityTrendSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeHostInsightResourceCapacityTrendSortByEnumStringValues(), ",")))
	}
	for _, val := range request.Status {
		if _, ok := GetMappingResourceStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetResourceStatusEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightResourceCapacityTrendResponse wrapper for the SummarizeHostInsightResourceCapacityTrend operation
type SummarizeHostInsightResourceCapacityTrendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeHostInsightResourceCapacityTrendAggregationCollection instances
	SummarizeHostInsightResourceCapacityTrendAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeHostInsightResourceCapacityTrendResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeHostInsightResourceCapacityTrendResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum Enum with underlying type: string
type SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum
const (
	SummarizeHostInsightResourceCapacityTrendPlatformTypeLinux   SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum = "LINUX"
	SummarizeHostInsightResourceCapacityTrendPlatformTypeSolaris SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum = "SOLARIS"
	SummarizeHostInsightResourceCapacityTrendPlatformTypeSunos   SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum = "SUNOS"
	SummarizeHostInsightResourceCapacityTrendPlatformTypeZlinux  SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum = "ZLINUX"
	SummarizeHostInsightResourceCapacityTrendPlatformTypeWindows SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum = "WINDOWS"
	SummarizeHostInsightResourceCapacityTrendPlatformTypeAix     SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum = "AIX"
	SummarizeHostInsightResourceCapacityTrendPlatformTypeHpUx    SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum = "HP_UX"
)

var mappingSummarizeHostInsightResourceCapacityTrendPlatformTypeEnum = map[string]SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum{
	"LINUX":   SummarizeHostInsightResourceCapacityTrendPlatformTypeLinux,
	"SOLARIS": SummarizeHostInsightResourceCapacityTrendPlatformTypeSolaris,
	"SUNOS":   SummarizeHostInsightResourceCapacityTrendPlatformTypeSunos,
	"ZLINUX":  SummarizeHostInsightResourceCapacityTrendPlatformTypeZlinux,
	"WINDOWS": SummarizeHostInsightResourceCapacityTrendPlatformTypeWindows,
	"AIX":     SummarizeHostInsightResourceCapacityTrendPlatformTypeAix,
	"HP_UX":   SummarizeHostInsightResourceCapacityTrendPlatformTypeHpUx,
}

var mappingSummarizeHostInsightResourceCapacityTrendPlatformTypeEnumLowerCase = map[string]SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum{
	"linux":   SummarizeHostInsightResourceCapacityTrendPlatformTypeLinux,
	"solaris": SummarizeHostInsightResourceCapacityTrendPlatformTypeSolaris,
	"sunos":   SummarizeHostInsightResourceCapacityTrendPlatformTypeSunos,
	"zlinux":  SummarizeHostInsightResourceCapacityTrendPlatformTypeZlinux,
	"windows": SummarizeHostInsightResourceCapacityTrendPlatformTypeWindows,
	"aix":     SummarizeHostInsightResourceCapacityTrendPlatformTypeAix,
	"hp_ux":   SummarizeHostInsightResourceCapacityTrendPlatformTypeHpUx,
}

// GetSummarizeHostInsightResourceCapacityTrendPlatformTypeEnumValues Enumerates the set of values for SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum
func GetSummarizeHostInsightResourceCapacityTrendPlatformTypeEnumValues() []SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum {
	values := make([]SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceCapacityTrendPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceCapacityTrendPlatformTypeEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum
func GetSummarizeHostInsightResourceCapacityTrendPlatformTypeEnumStringValues() []string {
	return []string{
		"LINUX",
		"SOLARIS",
		"SUNOS",
		"ZLINUX",
		"WINDOWS",
		"AIX",
		"HP_UX",
	}
}

// GetMappingSummarizeHostInsightResourceCapacityTrendPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceCapacityTrendPlatformTypeEnum(val string) (SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceCapacityTrendPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum Enum with underlying type: string
type SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum
const (
	SummarizeHostInsightResourceCapacityTrendUtilizationLevelHighUtilization       SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum = "HIGH_UTILIZATION"
	SummarizeHostInsightResourceCapacityTrendUtilizationLevelLowUtilization        SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum = "LOW_UTILIZATION"
	SummarizeHostInsightResourceCapacityTrendUtilizationLevelMediumHighUtilization SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum = "MEDIUM_HIGH_UTILIZATION"
	SummarizeHostInsightResourceCapacityTrendUtilizationLevelMediumLowUtilization  SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum = "MEDIUM_LOW_UTILIZATION"
)

var mappingSummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum = map[string]SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum{
	"HIGH_UTILIZATION":        SummarizeHostInsightResourceCapacityTrendUtilizationLevelHighUtilization,
	"LOW_UTILIZATION":         SummarizeHostInsightResourceCapacityTrendUtilizationLevelLowUtilization,
	"MEDIUM_HIGH_UTILIZATION": SummarizeHostInsightResourceCapacityTrendUtilizationLevelMediumHighUtilization,
	"MEDIUM_LOW_UTILIZATION":  SummarizeHostInsightResourceCapacityTrendUtilizationLevelMediumLowUtilization,
}

var mappingSummarizeHostInsightResourceCapacityTrendUtilizationLevelEnumLowerCase = map[string]SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum{
	"high_utilization":        SummarizeHostInsightResourceCapacityTrendUtilizationLevelHighUtilization,
	"low_utilization":         SummarizeHostInsightResourceCapacityTrendUtilizationLevelLowUtilization,
	"medium_high_utilization": SummarizeHostInsightResourceCapacityTrendUtilizationLevelMediumHighUtilization,
	"medium_low_utilization":  SummarizeHostInsightResourceCapacityTrendUtilizationLevelMediumLowUtilization,
}

// GetSummarizeHostInsightResourceCapacityTrendUtilizationLevelEnumValues Enumerates the set of values for SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum
func GetSummarizeHostInsightResourceCapacityTrendUtilizationLevelEnumValues() []SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum {
	values := make([]SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceCapacityTrendUtilizationLevelEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum
func GetSummarizeHostInsightResourceCapacityTrendUtilizationLevelEnumStringValues() []string {
	return []string{
		"HIGH_UTILIZATION",
		"LOW_UTILIZATION",
		"MEDIUM_HIGH_UTILIZATION",
		"MEDIUM_LOW_UTILIZATION",
	}
}

// GetMappingSummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum(val string) (SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceCapacityTrendUtilizationLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeHostInsightResourceCapacityTrendSortOrderEnum Enum with underlying type: string
type SummarizeHostInsightResourceCapacityTrendSortOrderEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceCapacityTrendSortOrderEnum
const (
	SummarizeHostInsightResourceCapacityTrendSortOrderAsc  SummarizeHostInsightResourceCapacityTrendSortOrderEnum = "ASC"
	SummarizeHostInsightResourceCapacityTrendSortOrderDesc SummarizeHostInsightResourceCapacityTrendSortOrderEnum = "DESC"
)

var mappingSummarizeHostInsightResourceCapacityTrendSortOrderEnum = map[string]SummarizeHostInsightResourceCapacityTrendSortOrderEnum{
	"ASC":  SummarizeHostInsightResourceCapacityTrendSortOrderAsc,
	"DESC": SummarizeHostInsightResourceCapacityTrendSortOrderDesc,
}

var mappingSummarizeHostInsightResourceCapacityTrendSortOrderEnumLowerCase = map[string]SummarizeHostInsightResourceCapacityTrendSortOrderEnum{
	"asc":  SummarizeHostInsightResourceCapacityTrendSortOrderAsc,
	"desc": SummarizeHostInsightResourceCapacityTrendSortOrderDesc,
}

// GetSummarizeHostInsightResourceCapacityTrendSortOrderEnumValues Enumerates the set of values for SummarizeHostInsightResourceCapacityTrendSortOrderEnum
func GetSummarizeHostInsightResourceCapacityTrendSortOrderEnumValues() []SummarizeHostInsightResourceCapacityTrendSortOrderEnum {
	values := make([]SummarizeHostInsightResourceCapacityTrendSortOrderEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceCapacityTrendSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceCapacityTrendSortOrderEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceCapacityTrendSortOrderEnum
func GetSummarizeHostInsightResourceCapacityTrendSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeHostInsightResourceCapacityTrendSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceCapacityTrendSortOrderEnum(val string) (SummarizeHostInsightResourceCapacityTrendSortOrderEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceCapacityTrendSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeHostInsightResourceCapacityTrendSortByEnum Enum with underlying type: string
type SummarizeHostInsightResourceCapacityTrendSortByEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceCapacityTrendSortByEnum
const (
	SummarizeHostInsightResourceCapacityTrendSortByEndtimestamp SummarizeHostInsightResourceCapacityTrendSortByEnum = "endTimestamp"
	SummarizeHostInsightResourceCapacityTrendSortByCapacity     SummarizeHostInsightResourceCapacityTrendSortByEnum = "capacity"
)

var mappingSummarizeHostInsightResourceCapacityTrendSortByEnum = map[string]SummarizeHostInsightResourceCapacityTrendSortByEnum{
	"endTimestamp": SummarizeHostInsightResourceCapacityTrendSortByEndtimestamp,
	"capacity":     SummarizeHostInsightResourceCapacityTrendSortByCapacity,
}

var mappingSummarizeHostInsightResourceCapacityTrendSortByEnumLowerCase = map[string]SummarizeHostInsightResourceCapacityTrendSortByEnum{
	"endtimestamp": SummarizeHostInsightResourceCapacityTrendSortByEndtimestamp,
	"capacity":     SummarizeHostInsightResourceCapacityTrendSortByCapacity,
}

// GetSummarizeHostInsightResourceCapacityTrendSortByEnumValues Enumerates the set of values for SummarizeHostInsightResourceCapacityTrendSortByEnum
func GetSummarizeHostInsightResourceCapacityTrendSortByEnumValues() []SummarizeHostInsightResourceCapacityTrendSortByEnum {
	values := make([]SummarizeHostInsightResourceCapacityTrendSortByEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceCapacityTrendSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceCapacityTrendSortByEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceCapacityTrendSortByEnum
func GetSummarizeHostInsightResourceCapacityTrendSortByEnumStringValues() []string {
	return []string{
		"endTimestamp",
		"capacity",
	}
}

// GetMappingSummarizeHostInsightResourceCapacityTrendSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceCapacityTrendSortByEnum(val string) (SummarizeHostInsightResourceCapacityTrendSortByEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceCapacityTrendSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
