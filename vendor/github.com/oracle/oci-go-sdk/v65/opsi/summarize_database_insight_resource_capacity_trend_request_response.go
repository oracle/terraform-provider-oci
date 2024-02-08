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

// SummarizeDatabaseInsightResourceCapacityTrendRequest wrapper for the SummarizeDatabaseInsightResourceCapacityTrend operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceCapacityTrend.go.html to see an example of how to use SummarizeDatabaseInsightResourceCapacityTrendRequest.
type SummarizeDatabaseInsightResourceCapacityTrendRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter by resource metric.
	// Supported values are CPU , STORAGE, MEMORY and IO.
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

	// Filter by one or more database type.
	// Possible values are ADW-S, ATP-S, ADW-D, ATP-D, EXTERNAL-PDB, EXTERNAL-NONCDB.
	DatabaseType []SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum `contributesTo:"query" name:"databaseType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of database OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Filter by one or more cdb name.
	CdbName []string `contributesTo:"query" name:"cdbName" collectionFormat:"multi"`

	// Filter by utilization level by the following buckets:
	//   - HIGH_UTILIZATION: DBs with utilization greater or equal than 75.
	//   - LOW_UTILIZATION: DBs with utilization lower than 25.
	//   - MEDIUM_HIGH_UTILIZATION: DBs with utilization greater or equal than 50 but lower than 75.
	//   - MEDIUM_LOW_UTILIZATION: DBs with utilization greater or equal than 25 but lower than 50.
	UtilizationLevel SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum `mandatory:"false" contributesTo:"query" name:"utilizationLevel" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sorts using end timestamp , capacity or baseCapacity
	SortBy SummarizeDatabaseInsightResourceCapacityTrendSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Tablespace name for a database
	TablespaceName *string `mandatory:"false" contributesTo:"query" name:"tablespaceName"`

	// Filter by one or more hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

	// Flag to indicate if database instance level metrics should be returned. The flag is ignored when a host name filter is not applied.
	// When a hostname filter is applied this flag will determine whether to return metrics for the instances located on the specified host or for the
	// whole database which contains an instance on this host.
	IsDatabaseInstanceLevelMetrics *bool `mandatory:"false" contributesTo:"query" name:"isDatabaseInstanceLevelMetrics"`

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

	// Optional list of Exadata Insight VM cluster name.
	VmclusterName []string `contributesTo:"query" name:"vmclusterName" collectionFormat:"multi"`

	// Percent value in which a resource metric is considered highly utilized.
	HighUtilizationThreshold *int `mandatory:"false" contributesTo:"query" name:"highUtilizationThreshold"`

	// Percent value in which a resource metric is considered low utilized.
	LowUtilizationThreshold *int `mandatory:"false" contributesTo:"query" name:"lowUtilizationThreshold"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeDatabaseInsightResourceCapacityTrendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeDatabaseInsightResourceCapacityTrendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeDatabaseInsightResourceCapacityTrendRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeDatabaseInsightResourceCapacityTrendRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeDatabaseInsightResourceCapacityTrendRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.DatabaseType {
		if _, ok := GetMappingSummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", val, strings.Join(GetSummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum(string(request.UtilizationLevel)); !ok && request.UtilizationLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UtilizationLevel: %s. Supported values are: %s.", request.UtilizationLevel, strings.Join(GetSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeDatabaseInsightResourceCapacityTrendSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceCapacityTrendSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeDatabaseInsightResourceCapacityTrendSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeDatabaseInsightResourceCapacityTrendResponse wrapper for the SummarizeDatabaseInsightResourceCapacityTrend operation
type SummarizeDatabaseInsightResourceCapacityTrendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeDatabaseInsightResourceCapacityTrendAggregationCollection instances
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeDatabaseInsightResourceCapacityTrendResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeDatabaseInsightResourceCapacityTrendResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum
const (
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAdwS                 SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "ADW-S"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAtpS                 SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "ATP-S"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAdwD                 SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "ADW-D"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAtpD                 SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "ATP-D"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeExternalPdb          SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "EXTERNAL-PDB"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeExternalNoncdb       SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "EXTERNAL-NONCDB"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedVmCdb       SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "COMANAGED-VM-CDB"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedVmPdb       SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "COMANAGED-VM-PDB"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedVmNoncdb    SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "COMANAGED-VM-NONCDB"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedBmCdb       SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "COMANAGED-BM-CDB"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedBmPdb       SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "COMANAGED-BM-PDB"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedBmNoncdb    SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "COMANAGED-BM-NONCDB"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedExacsCdb    SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "COMANAGED-EXACS-CDB"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedExacsPdb    SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "COMANAGED-EXACS-PDB"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedExacsNoncdb SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "COMANAGED-EXACS-NONCDB"
)

var mappingSummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = map[string]SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum{
	"ADW-S":                  SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAdwS,
	"ATP-S":                  SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAtpS,
	"ADW-D":                  SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAdwD,
	"ATP-D":                  SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAtpD,
	"EXTERNAL-PDB":           SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeExternalPdb,
	"EXTERNAL-NONCDB":        SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeExternalNoncdb,
	"COMANAGED-VM-CDB":       SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedVmCdb,
	"COMANAGED-VM-PDB":       SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedVmPdb,
	"COMANAGED-VM-NONCDB":    SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedVmNoncdb,
	"COMANAGED-BM-CDB":       SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedBmCdb,
	"COMANAGED-BM-PDB":       SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedBmPdb,
	"COMANAGED-BM-NONCDB":    SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedBmNoncdb,
	"COMANAGED-EXACS-CDB":    SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedExacsCdb,
	"COMANAGED-EXACS-PDB":    SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedExacsPdb,
	"COMANAGED-EXACS-NONCDB": SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedExacsNoncdb,
}

var mappingSummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnumLowerCase = map[string]SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum{
	"adw-s":                  SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAdwS,
	"atp-s":                  SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAtpS,
	"adw-d":                  SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAdwD,
	"atp-d":                  SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAtpD,
	"external-pdb":           SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeExternalPdb,
	"external-noncdb":        SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeExternalNoncdb,
	"comanaged-vm-cdb":       SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedVmCdb,
	"comanaged-vm-pdb":       SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedVmPdb,
	"comanaged-vm-noncdb":    SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedVmNoncdb,
	"comanaged-bm-cdb":       SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedBmCdb,
	"comanaged-bm-pdb":       SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedBmPdb,
	"comanaged-bm-noncdb":    SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedBmNoncdb,
	"comanaged-exacs-cdb":    SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedExacsCdb,
	"comanaged-exacs-pdb":    SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedExacsPdb,
	"comanaged-exacs-noncdb": SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeComanagedExacsNoncdb,
}

// GetSummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnumValues() []SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum {
	values := make([]SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnumStringValues() []string {
	return []string{
		"ADW-S",
		"ATP-S",
		"ADW-D",
		"ATP-D",
		"EXTERNAL-PDB",
		"EXTERNAL-NONCDB",
		"COMANAGED-VM-CDB",
		"COMANAGED-VM-PDB",
		"COMANAGED-VM-NONCDB",
		"COMANAGED-BM-CDB",
		"COMANAGED-BM-PDB",
		"COMANAGED-BM-NONCDB",
		"COMANAGED-EXACS-CDB",
		"COMANAGED-EXACS-PDB",
		"COMANAGED-EXACS-NONCDB",
	}
}

// GetMappingSummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum(val string) (SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum
const (
	SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelHighUtilization       SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum = "HIGH_UTILIZATION"
	SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelLowUtilization        SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum = "LOW_UTILIZATION"
	SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelMediumHighUtilization SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum = "MEDIUM_HIGH_UTILIZATION"
	SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelMediumLowUtilization  SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum = "MEDIUM_LOW_UTILIZATION"
)

var mappingSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum = map[string]SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum{
	"HIGH_UTILIZATION":        SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelHighUtilization,
	"LOW_UTILIZATION":         SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelLowUtilization,
	"MEDIUM_HIGH_UTILIZATION": SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelMediumHighUtilization,
	"MEDIUM_LOW_UTILIZATION":  SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelMediumLowUtilization,
}

var mappingSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnumLowerCase = map[string]SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum{
	"high_utilization":        SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelHighUtilization,
	"low_utilization":         SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelLowUtilization,
	"medium_high_utilization": SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelMediumHighUtilization,
	"medium_low_utilization":  SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelMediumLowUtilization,
}

// GetSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnumValues() []SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum {
	values := make([]SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnumStringValues() []string {
	return []string{
		"HIGH_UTILIZATION",
		"LOW_UTILIZATION",
		"MEDIUM_HIGH_UTILIZATION",
		"MEDIUM_LOW_UTILIZATION",
	}
}

// GetMappingSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum(val string) (SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum
const (
	SummarizeDatabaseInsightResourceCapacityTrendSortOrderAsc  SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum = "ASC"
	SummarizeDatabaseInsightResourceCapacityTrendSortOrderDesc SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum = "DESC"
)

var mappingSummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum = map[string]SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum{
	"ASC":  SummarizeDatabaseInsightResourceCapacityTrendSortOrderAsc,
	"DESC": SummarizeDatabaseInsightResourceCapacityTrendSortOrderDesc,
}

var mappingSummarizeDatabaseInsightResourceCapacityTrendSortOrderEnumLowerCase = map[string]SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum{
	"asc":  SummarizeDatabaseInsightResourceCapacityTrendSortOrderAsc,
	"desc": SummarizeDatabaseInsightResourceCapacityTrendSortOrderDesc,
}

// GetSummarizeDatabaseInsightResourceCapacityTrendSortOrderEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendSortOrderEnumValues() []SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum {
	values := make([]SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceCapacityTrendSortOrderEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum(val string) (SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceCapacityTrendSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeDatabaseInsightResourceCapacityTrendSortByEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceCapacityTrendSortByEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceCapacityTrendSortByEnum
const (
	SummarizeDatabaseInsightResourceCapacityTrendSortByEndtimestamp SummarizeDatabaseInsightResourceCapacityTrendSortByEnum = "endTimestamp"
	SummarizeDatabaseInsightResourceCapacityTrendSortByCapacity     SummarizeDatabaseInsightResourceCapacityTrendSortByEnum = "capacity"
	SummarizeDatabaseInsightResourceCapacityTrendSortByBasecapacity SummarizeDatabaseInsightResourceCapacityTrendSortByEnum = "baseCapacity"
)

var mappingSummarizeDatabaseInsightResourceCapacityTrendSortByEnum = map[string]SummarizeDatabaseInsightResourceCapacityTrendSortByEnum{
	"endTimestamp": SummarizeDatabaseInsightResourceCapacityTrendSortByEndtimestamp,
	"capacity":     SummarizeDatabaseInsightResourceCapacityTrendSortByCapacity,
	"baseCapacity": SummarizeDatabaseInsightResourceCapacityTrendSortByBasecapacity,
}

var mappingSummarizeDatabaseInsightResourceCapacityTrendSortByEnumLowerCase = map[string]SummarizeDatabaseInsightResourceCapacityTrendSortByEnum{
	"endtimestamp": SummarizeDatabaseInsightResourceCapacityTrendSortByEndtimestamp,
	"capacity":     SummarizeDatabaseInsightResourceCapacityTrendSortByCapacity,
	"basecapacity": SummarizeDatabaseInsightResourceCapacityTrendSortByBasecapacity,
}

// GetSummarizeDatabaseInsightResourceCapacityTrendSortByEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceCapacityTrendSortByEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendSortByEnumValues() []SummarizeDatabaseInsightResourceCapacityTrendSortByEnum {
	values := make([]SummarizeDatabaseInsightResourceCapacityTrendSortByEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceCapacityTrendSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceCapacityTrendSortByEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceCapacityTrendSortByEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendSortByEnumStringValues() []string {
	return []string{
		"endTimestamp",
		"capacity",
		"baseCapacity",
	}
}

// GetMappingSummarizeDatabaseInsightResourceCapacityTrendSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceCapacityTrendSortByEnum(val string) (SummarizeDatabaseInsightResourceCapacityTrendSortByEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceCapacityTrendSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
