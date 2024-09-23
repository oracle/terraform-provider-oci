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

// SummarizeDatabaseInsightResourceUsageTrendRequest wrapper for the SummarizeDatabaseInsightResourceUsageTrend operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceUsageTrend.go.html to see an example of how to use SummarizeDatabaseInsightResourceUsageTrendRequest.
type SummarizeDatabaseInsightResourceUsageTrendRequest struct {

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
	DatabaseType []SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum `contributesTo:"query" name:"databaseType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of database OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeDatabaseInsightResourceUsageTrendSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sorts using end timestamp, usage or capacity
	SortBy SummarizeDatabaseInsightResourceUsageTrendSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

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

	// Filter by one or more cdb name.
	CdbName []string `contributesTo:"query" name:"cdbName" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeDatabaseInsightResourceUsageTrendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeDatabaseInsightResourceUsageTrendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeDatabaseInsightResourceUsageTrendRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeDatabaseInsightResourceUsageTrendRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeDatabaseInsightResourceUsageTrendRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.DatabaseType {
		if _, ok := GetMappingSummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", val, strings.Join(GetSummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSummarizeDatabaseInsightResourceUsageTrendSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeDatabaseInsightResourceUsageTrendSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceUsageTrendSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeDatabaseInsightResourceUsageTrendSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeDatabaseInsightResourceUsageTrendResponse wrapper for the SummarizeDatabaseInsightResourceUsageTrend operation
type SummarizeDatabaseInsightResourceUsageTrendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeDatabaseInsightResourceUsageTrendAggregationCollection instances
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeDatabaseInsightResourceUsageTrendResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeDatabaseInsightResourceUsageTrendResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum
const (
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeAdwS                 SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "ADW-S"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeAtpS                 SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "ATP-S"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeAdwD                 SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "ADW-D"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeAtpD                 SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "ATP-D"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeExternalPdb          SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "EXTERNAL-PDB"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeExternalNoncdb       SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "EXTERNAL-NONCDB"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedVmCdb       SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "COMANAGED-VM-CDB"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedVmPdb       SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "COMANAGED-VM-PDB"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedVmNoncdb    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "COMANAGED-VM-NONCDB"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedBmCdb       SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "COMANAGED-BM-CDB"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedBmPdb       SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "COMANAGED-BM-PDB"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedBmNoncdb    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "COMANAGED-BM-NONCDB"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExacsCdb    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "COMANAGED-EXACS-CDB"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExacsPdb    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "COMANAGED-EXACS-PDB"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExacsNoncdb SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "COMANAGED-EXACS-NONCDB"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExaccCdb    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "COMANAGED-EXACC-CDB"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExaccPdb    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "COMANAGED-EXACC-PDB"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExaccNoncdb SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "COMANAGED-EXACC-NONCDB"
	SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeMdsMysql             SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = "MDS-MYSQL"
)

var mappingSummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum = map[string]SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum{
	"ADW-S":                  SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeAdwS,
	"ATP-S":                  SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeAtpS,
	"ADW-D":                  SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeAdwD,
	"ATP-D":                  SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeAtpD,
	"EXTERNAL-PDB":           SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeExternalPdb,
	"EXTERNAL-NONCDB":        SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeExternalNoncdb,
	"COMANAGED-VM-CDB":       SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedVmCdb,
	"COMANAGED-VM-PDB":       SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedVmPdb,
	"COMANAGED-VM-NONCDB":    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedVmNoncdb,
	"COMANAGED-BM-CDB":       SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedBmCdb,
	"COMANAGED-BM-PDB":       SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedBmPdb,
	"COMANAGED-BM-NONCDB":    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedBmNoncdb,
	"COMANAGED-EXACS-CDB":    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExacsCdb,
	"COMANAGED-EXACS-PDB":    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExacsPdb,
	"COMANAGED-EXACS-NONCDB": SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExacsNoncdb,
	"COMANAGED-EXACC-CDB":    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExaccCdb,
	"COMANAGED-EXACC-PDB":    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExaccPdb,
	"COMANAGED-EXACC-NONCDB": SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExaccNoncdb,
	"MDS-MYSQL":              SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeMdsMysql,
}

var mappingSummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnumLowerCase = map[string]SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum{
	"adw-s":                  SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeAdwS,
	"atp-s":                  SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeAtpS,
	"adw-d":                  SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeAdwD,
	"atp-d":                  SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeAtpD,
	"external-pdb":           SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeExternalPdb,
	"external-noncdb":        SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeExternalNoncdb,
	"comanaged-vm-cdb":       SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedVmCdb,
	"comanaged-vm-pdb":       SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedVmPdb,
	"comanaged-vm-noncdb":    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedVmNoncdb,
	"comanaged-bm-cdb":       SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedBmCdb,
	"comanaged-bm-pdb":       SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedBmPdb,
	"comanaged-bm-noncdb":    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedBmNoncdb,
	"comanaged-exacs-cdb":    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExacsCdb,
	"comanaged-exacs-pdb":    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExacsPdb,
	"comanaged-exacs-noncdb": SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExacsNoncdb,
	"comanaged-exacc-cdb":    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExaccCdb,
	"comanaged-exacc-pdb":    SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExaccPdb,
	"comanaged-exacc-noncdb": SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeComanagedExaccNoncdb,
	"mds-mysql":              SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeMdsMysql,
}

// GetSummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum
func GetSummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnumValues() []SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum {
	values := make([]SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum
func GetSummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnumStringValues() []string {
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
		"COMANAGED-EXACC-CDB",
		"COMANAGED-EXACC-PDB",
		"COMANAGED-EXACC-NONCDB",
		"MDS-MYSQL",
	}
}

// GetMappingSummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum(val string) (SummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceUsageTrendDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeDatabaseInsightResourceUsageTrendSortOrderEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceUsageTrendSortOrderEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceUsageTrendSortOrderEnum
const (
	SummarizeDatabaseInsightResourceUsageTrendSortOrderAsc  SummarizeDatabaseInsightResourceUsageTrendSortOrderEnum = "ASC"
	SummarizeDatabaseInsightResourceUsageTrendSortOrderDesc SummarizeDatabaseInsightResourceUsageTrendSortOrderEnum = "DESC"
)

var mappingSummarizeDatabaseInsightResourceUsageTrendSortOrderEnum = map[string]SummarizeDatabaseInsightResourceUsageTrendSortOrderEnum{
	"ASC":  SummarizeDatabaseInsightResourceUsageTrendSortOrderAsc,
	"DESC": SummarizeDatabaseInsightResourceUsageTrendSortOrderDesc,
}

var mappingSummarizeDatabaseInsightResourceUsageTrendSortOrderEnumLowerCase = map[string]SummarizeDatabaseInsightResourceUsageTrendSortOrderEnum{
	"asc":  SummarizeDatabaseInsightResourceUsageTrendSortOrderAsc,
	"desc": SummarizeDatabaseInsightResourceUsageTrendSortOrderDesc,
}

// GetSummarizeDatabaseInsightResourceUsageTrendSortOrderEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceUsageTrendSortOrderEnum
func GetSummarizeDatabaseInsightResourceUsageTrendSortOrderEnumValues() []SummarizeDatabaseInsightResourceUsageTrendSortOrderEnum {
	values := make([]SummarizeDatabaseInsightResourceUsageTrendSortOrderEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceUsageTrendSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceUsageTrendSortOrderEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceUsageTrendSortOrderEnum
func GetSummarizeDatabaseInsightResourceUsageTrendSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeDatabaseInsightResourceUsageTrendSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceUsageTrendSortOrderEnum(val string) (SummarizeDatabaseInsightResourceUsageTrendSortOrderEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceUsageTrendSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeDatabaseInsightResourceUsageTrendSortByEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceUsageTrendSortByEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceUsageTrendSortByEnum
const (
	SummarizeDatabaseInsightResourceUsageTrendSortByEndtimestamp SummarizeDatabaseInsightResourceUsageTrendSortByEnum = "endTimestamp"
	SummarizeDatabaseInsightResourceUsageTrendSortByUsage        SummarizeDatabaseInsightResourceUsageTrendSortByEnum = "usage"
	SummarizeDatabaseInsightResourceUsageTrendSortByCapacity     SummarizeDatabaseInsightResourceUsageTrendSortByEnum = "capacity"
)

var mappingSummarizeDatabaseInsightResourceUsageTrendSortByEnum = map[string]SummarizeDatabaseInsightResourceUsageTrendSortByEnum{
	"endTimestamp": SummarizeDatabaseInsightResourceUsageTrendSortByEndtimestamp,
	"usage":        SummarizeDatabaseInsightResourceUsageTrendSortByUsage,
	"capacity":     SummarizeDatabaseInsightResourceUsageTrendSortByCapacity,
}

var mappingSummarizeDatabaseInsightResourceUsageTrendSortByEnumLowerCase = map[string]SummarizeDatabaseInsightResourceUsageTrendSortByEnum{
	"endtimestamp": SummarizeDatabaseInsightResourceUsageTrendSortByEndtimestamp,
	"usage":        SummarizeDatabaseInsightResourceUsageTrendSortByUsage,
	"capacity":     SummarizeDatabaseInsightResourceUsageTrendSortByCapacity,
}

// GetSummarizeDatabaseInsightResourceUsageTrendSortByEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceUsageTrendSortByEnum
func GetSummarizeDatabaseInsightResourceUsageTrendSortByEnumValues() []SummarizeDatabaseInsightResourceUsageTrendSortByEnum {
	values := make([]SummarizeDatabaseInsightResourceUsageTrendSortByEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceUsageTrendSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceUsageTrendSortByEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceUsageTrendSortByEnum
func GetSummarizeDatabaseInsightResourceUsageTrendSortByEnumStringValues() []string {
	return []string{
		"endTimestamp",
		"usage",
		"capacity",
	}
}

// GetMappingSummarizeDatabaseInsightResourceUsageTrendSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceUsageTrendSortByEnum(val string) (SummarizeDatabaseInsightResourceUsageTrendSortByEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceUsageTrendSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
