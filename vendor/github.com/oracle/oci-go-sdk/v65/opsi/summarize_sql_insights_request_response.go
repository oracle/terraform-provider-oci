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

// SummarizeSqlInsightsRequest wrapper for the SummarizeSqlInsights operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlInsights.go.html to see an example of how to use SummarizeSqlInsightsRequest.
type SummarizeSqlInsightsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter by one or more database type.
	// Possible values are ADW-S, ATP-S, ADW-D, ATP-D, EXTERNAL-PDB, EXTERNAL-NONCDB.
	DatabaseType []SummarizeSqlInsightsDatabaseTypeEnum `contributesTo:"query" name:"databaseType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of database OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Filter by one or more cdb name.
	CdbName []string `contributesTo:"query" name:"cdbName" collectionFormat:"multi"`

	// Filter by one or more hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

	// Filter sqls by percentage of db time.
	DatabaseTimePctGreaterThan *float64 `mandatory:"false" contributesTo:"query" name:"databaseTimePctGreaterThan"`

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

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

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

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeSqlInsightsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeSqlInsightsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeSqlInsightsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeSqlInsightsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeSqlInsightsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.DatabaseType {
		if _, ok := GetMappingSummarizeSqlInsightsDatabaseTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", val, strings.Join(GetSummarizeSqlInsightsDatabaseTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeSqlInsightsResponse wrapper for the SummarizeSqlInsights operation
type SummarizeSqlInsightsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlInsightAggregationCollection instances
	SqlInsightAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeSqlInsightsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeSqlInsightsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeSqlInsightsDatabaseTypeEnum Enum with underlying type: string
type SummarizeSqlInsightsDatabaseTypeEnum string

// Set of constants representing the allowable values for SummarizeSqlInsightsDatabaseTypeEnum
const (
	SummarizeSqlInsightsDatabaseTypeAdwS                 SummarizeSqlInsightsDatabaseTypeEnum = "ADW-S"
	SummarizeSqlInsightsDatabaseTypeAtpS                 SummarizeSqlInsightsDatabaseTypeEnum = "ATP-S"
	SummarizeSqlInsightsDatabaseTypeAdwD                 SummarizeSqlInsightsDatabaseTypeEnum = "ADW-D"
	SummarizeSqlInsightsDatabaseTypeAtpD                 SummarizeSqlInsightsDatabaseTypeEnum = "ATP-D"
	SummarizeSqlInsightsDatabaseTypeExternalPdb          SummarizeSqlInsightsDatabaseTypeEnum = "EXTERNAL-PDB"
	SummarizeSqlInsightsDatabaseTypeExternalNoncdb       SummarizeSqlInsightsDatabaseTypeEnum = "EXTERNAL-NONCDB"
	SummarizeSqlInsightsDatabaseTypeComanagedVmCdb       SummarizeSqlInsightsDatabaseTypeEnum = "COMANAGED-VM-CDB"
	SummarizeSqlInsightsDatabaseTypeComanagedVmPdb       SummarizeSqlInsightsDatabaseTypeEnum = "COMANAGED-VM-PDB"
	SummarizeSqlInsightsDatabaseTypeComanagedVmNoncdb    SummarizeSqlInsightsDatabaseTypeEnum = "COMANAGED-VM-NONCDB"
	SummarizeSqlInsightsDatabaseTypeComanagedBmCdb       SummarizeSqlInsightsDatabaseTypeEnum = "COMANAGED-BM-CDB"
	SummarizeSqlInsightsDatabaseTypeComanagedBmPdb       SummarizeSqlInsightsDatabaseTypeEnum = "COMANAGED-BM-PDB"
	SummarizeSqlInsightsDatabaseTypeComanagedBmNoncdb    SummarizeSqlInsightsDatabaseTypeEnum = "COMANAGED-BM-NONCDB"
	SummarizeSqlInsightsDatabaseTypeComanagedExacsCdb    SummarizeSqlInsightsDatabaseTypeEnum = "COMANAGED-EXACS-CDB"
	SummarizeSqlInsightsDatabaseTypeComanagedExacsPdb    SummarizeSqlInsightsDatabaseTypeEnum = "COMANAGED-EXACS-PDB"
	SummarizeSqlInsightsDatabaseTypeComanagedExacsNoncdb SummarizeSqlInsightsDatabaseTypeEnum = "COMANAGED-EXACS-NONCDB"
)

var mappingSummarizeSqlInsightsDatabaseTypeEnum = map[string]SummarizeSqlInsightsDatabaseTypeEnum{
	"ADW-S":                  SummarizeSqlInsightsDatabaseTypeAdwS,
	"ATP-S":                  SummarizeSqlInsightsDatabaseTypeAtpS,
	"ADW-D":                  SummarizeSqlInsightsDatabaseTypeAdwD,
	"ATP-D":                  SummarizeSqlInsightsDatabaseTypeAtpD,
	"EXTERNAL-PDB":           SummarizeSqlInsightsDatabaseTypeExternalPdb,
	"EXTERNAL-NONCDB":        SummarizeSqlInsightsDatabaseTypeExternalNoncdb,
	"COMANAGED-VM-CDB":       SummarizeSqlInsightsDatabaseTypeComanagedVmCdb,
	"COMANAGED-VM-PDB":       SummarizeSqlInsightsDatabaseTypeComanagedVmPdb,
	"COMANAGED-VM-NONCDB":    SummarizeSqlInsightsDatabaseTypeComanagedVmNoncdb,
	"COMANAGED-BM-CDB":       SummarizeSqlInsightsDatabaseTypeComanagedBmCdb,
	"COMANAGED-BM-PDB":       SummarizeSqlInsightsDatabaseTypeComanagedBmPdb,
	"COMANAGED-BM-NONCDB":    SummarizeSqlInsightsDatabaseTypeComanagedBmNoncdb,
	"COMANAGED-EXACS-CDB":    SummarizeSqlInsightsDatabaseTypeComanagedExacsCdb,
	"COMANAGED-EXACS-PDB":    SummarizeSqlInsightsDatabaseTypeComanagedExacsPdb,
	"COMANAGED-EXACS-NONCDB": SummarizeSqlInsightsDatabaseTypeComanagedExacsNoncdb,
}

var mappingSummarizeSqlInsightsDatabaseTypeEnumLowerCase = map[string]SummarizeSqlInsightsDatabaseTypeEnum{
	"adw-s":                  SummarizeSqlInsightsDatabaseTypeAdwS,
	"atp-s":                  SummarizeSqlInsightsDatabaseTypeAtpS,
	"adw-d":                  SummarizeSqlInsightsDatabaseTypeAdwD,
	"atp-d":                  SummarizeSqlInsightsDatabaseTypeAtpD,
	"external-pdb":           SummarizeSqlInsightsDatabaseTypeExternalPdb,
	"external-noncdb":        SummarizeSqlInsightsDatabaseTypeExternalNoncdb,
	"comanaged-vm-cdb":       SummarizeSqlInsightsDatabaseTypeComanagedVmCdb,
	"comanaged-vm-pdb":       SummarizeSqlInsightsDatabaseTypeComanagedVmPdb,
	"comanaged-vm-noncdb":    SummarizeSqlInsightsDatabaseTypeComanagedVmNoncdb,
	"comanaged-bm-cdb":       SummarizeSqlInsightsDatabaseTypeComanagedBmCdb,
	"comanaged-bm-pdb":       SummarizeSqlInsightsDatabaseTypeComanagedBmPdb,
	"comanaged-bm-noncdb":    SummarizeSqlInsightsDatabaseTypeComanagedBmNoncdb,
	"comanaged-exacs-cdb":    SummarizeSqlInsightsDatabaseTypeComanagedExacsCdb,
	"comanaged-exacs-pdb":    SummarizeSqlInsightsDatabaseTypeComanagedExacsPdb,
	"comanaged-exacs-noncdb": SummarizeSqlInsightsDatabaseTypeComanagedExacsNoncdb,
}

// GetSummarizeSqlInsightsDatabaseTypeEnumValues Enumerates the set of values for SummarizeSqlInsightsDatabaseTypeEnum
func GetSummarizeSqlInsightsDatabaseTypeEnumValues() []SummarizeSqlInsightsDatabaseTypeEnum {
	values := make([]SummarizeSqlInsightsDatabaseTypeEnum, 0)
	for _, v := range mappingSummarizeSqlInsightsDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeSqlInsightsDatabaseTypeEnumStringValues Enumerates the set of values in String for SummarizeSqlInsightsDatabaseTypeEnum
func GetSummarizeSqlInsightsDatabaseTypeEnumStringValues() []string {
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

// GetMappingSummarizeSqlInsightsDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeSqlInsightsDatabaseTypeEnum(val string) (SummarizeSqlInsightsDatabaseTypeEnum, bool) {
	enum, ok := mappingSummarizeSqlInsightsDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
