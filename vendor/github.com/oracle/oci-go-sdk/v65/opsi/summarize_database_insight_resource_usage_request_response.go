// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeDatabaseInsightResourceUsageRequest wrapper for the SummarizeDatabaseInsightResourceUsage operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceUsage.go.html to see an example of how to use SummarizeDatabaseInsightResourceUsageRequest.
type SummarizeDatabaseInsightResourceUsageRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
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
	DatabaseType []SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum `contributesTo:"query" name:"databaseType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of database OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Optional list of database insight resource OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Optional list of exadata insight resource OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Filter by one or more hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

	// Flag to indicate if database instance level metrics should be returned. The flag is ignored when a host name filter is not applied.
	// When a hostname filter is applied this flag will determine whether to return metrics for the instances located on the specified host or for the
	// whole database which contains an instance on this host.
	IsDatabaseInstanceLevelMetrics *bool `mandatory:"false" contributesTo:"query" name:"isDatabaseInstanceLevelMetrics"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Percentile values of daily usage to be used for computing the aggregate resource usage.
	Percentile *int `mandatory:"false" contributesTo:"query" name:"percentile"`

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

func (request SummarizeDatabaseInsightResourceUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeDatabaseInsightResourceUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeDatabaseInsightResourceUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeDatabaseInsightResourceUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeDatabaseInsightResourceUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.DatabaseType {
		if _, ok := GetMappingSummarizeDatabaseInsightResourceUsageDatabaseTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", val, strings.Join(GetSummarizeDatabaseInsightResourceUsageDatabaseTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeDatabaseInsightResourceUsageResponse wrapper for the SummarizeDatabaseInsightResourceUsage operation
type SummarizeDatabaseInsightResourceUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeDatabaseInsightResourceUsageAggregation instances
	SummarizeDatabaseInsightResourceUsageAggregation `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeDatabaseInsightResourceUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeDatabaseInsightResourceUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum
const (
	SummarizeDatabaseInsightResourceUsageDatabaseTypeAdwS                 SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "ADW-S"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeAtpS                 SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "ATP-S"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeAdwD                 SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "ADW-D"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeAtpD                 SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "ATP-D"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeExternalPdb          SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "EXTERNAL-PDB"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeExternalNoncdb       SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "EXTERNAL-NONCDB"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedVmCdb       SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "COMANAGED-VM-CDB"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedVmPdb       SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "COMANAGED-VM-PDB"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedVmNoncdb    SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "COMANAGED-VM-NONCDB"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedBmCdb       SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "COMANAGED-BM-CDB"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedBmPdb       SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "COMANAGED-BM-PDB"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedBmNoncdb    SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "COMANAGED-BM-NONCDB"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExacsCdb    SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "COMANAGED-EXACS-CDB"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExacsPdb    SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "COMANAGED-EXACS-PDB"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExacsNoncdb SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "COMANAGED-EXACS-NONCDB"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExaccCdb    SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "COMANAGED-EXACC-CDB"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExaccPdb    SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "COMANAGED-EXACC-PDB"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExaccNoncdb SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "COMANAGED-EXACC-NONCDB"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeMdsMysql             SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "MDS-MYSQL"
	SummarizeDatabaseInsightResourceUsageDatabaseTypeExternalMysql        SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = "EXTERNAL-MYSQL"
)

var mappingSummarizeDatabaseInsightResourceUsageDatabaseTypeEnum = map[string]SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum{
	"ADW-S":                  SummarizeDatabaseInsightResourceUsageDatabaseTypeAdwS,
	"ATP-S":                  SummarizeDatabaseInsightResourceUsageDatabaseTypeAtpS,
	"ADW-D":                  SummarizeDatabaseInsightResourceUsageDatabaseTypeAdwD,
	"ATP-D":                  SummarizeDatabaseInsightResourceUsageDatabaseTypeAtpD,
	"EXTERNAL-PDB":           SummarizeDatabaseInsightResourceUsageDatabaseTypeExternalPdb,
	"EXTERNAL-NONCDB":        SummarizeDatabaseInsightResourceUsageDatabaseTypeExternalNoncdb,
	"COMANAGED-VM-CDB":       SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedVmCdb,
	"COMANAGED-VM-PDB":       SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedVmPdb,
	"COMANAGED-VM-NONCDB":    SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedVmNoncdb,
	"COMANAGED-BM-CDB":       SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedBmCdb,
	"COMANAGED-BM-PDB":       SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedBmPdb,
	"COMANAGED-BM-NONCDB":    SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedBmNoncdb,
	"COMANAGED-EXACS-CDB":    SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExacsCdb,
	"COMANAGED-EXACS-PDB":    SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExacsPdb,
	"COMANAGED-EXACS-NONCDB": SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExacsNoncdb,
	"COMANAGED-EXACC-CDB":    SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExaccCdb,
	"COMANAGED-EXACC-PDB":    SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExaccPdb,
	"COMANAGED-EXACC-NONCDB": SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExaccNoncdb,
	"MDS-MYSQL":              SummarizeDatabaseInsightResourceUsageDatabaseTypeMdsMysql,
	"EXTERNAL-MYSQL":         SummarizeDatabaseInsightResourceUsageDatabaseTypeExternalMysql,
}

var mappingSummarizeDatabaseInsightResourceUsageDatabaseTypeEnumLowerCase = map[string]SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum{
	"adw-s":                  SummarizeDatabaseInsightResourceUsageDatabaseTypeAdwS,
	"atp-s":                  SummarizeDatabaseInsightResourceUsageDatabaseTypeAtpS,
	"adw-d":                  SummarizeDatabaseInsightResourceUsageDatabaseTypeAdwD,
	"atp-d":                  SummarizeDatabaseInsightResourceUsageDatabaseTypeAtpD,
	"external-pdb":           SummarizeDatabaseInsightResourceUsageDatabaseTypeExternalPdb,
	"external-noncdb":        SummarizeDatabaseInsightResourceUsageDatabaseTypeExternalNoncdb,
	"comanaged-vm-cdb":       SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedVmCdb,
	"comanaged-vm-pdb":       SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedVmPdb,
	"comanaged-vm-noncdb":    SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedVmNoncdb,
	"comanaged-bm-cdb":       SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedBmCdb,
	"comanaged-bm-pdb":       SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedBmPdb,
	"comanaged-bm-noncdb":    SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedBmNoncdb,
	"comanaged-exacs-cdb":    SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExacsCdb,
	"comanaged-exacs-pdb":    SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExacsPdb,
	"comanaged-exacs-noncdb": SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExacsNoncdb,
	"comanaged-exacc-cdb":    SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExaccCdb,
	"comanaged-exacc-pdb":    SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExaccPdb,
	"comanaged-exacc-noncdb": SummarizeDatabaseInsightResourceUsageDatabaseTypeComanagedExaccNoncdb,
	"mds-mysql":              SummarizeDatabaseInsightResourceUsageDatabaseTypeMdsMysql,
	"external-mysql":         SummarizeDatabaseInsightResourceUsageDatabaseTypeExternalMysql,
}

// GetSummarizeDatabaseInsightResourceUsageDatabaseTypeEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum
func GetSummarizeDatabaseInsightResourceUsageDatabaseTypeEnumValues() []SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum {
	values := make([]SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceUsageDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceUsageDatabaseTypeEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum
func GetSummarizeDatabaseInsightResourceUsageDatabaseTypeEnumStringValues() []string {
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
		"EXTERNAL-MYSQL",
	}
}

// GetMappingSummarizeDatabaseInsightResourceUsageDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceUsageDatabaseTypeEnum(val string) (SummarizeDatabaseInsightResourceUsageDatabaseTypeEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceUsageDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
