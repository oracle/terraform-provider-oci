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

// SummarizeHostInsightResourceUtilizationInsightRequest wrapper for the SummarizeHostInsightResourceUtilizationInsight operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceUtilizationInsight.go.html to see an example of how to use SummarizeHostInsightResourceUtilizationInsightRequest.
type SummarizeHostInsightResourceUtilizationInsightRequest struct {

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
	PlatformType []SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum `contributesTo:"query" name:"platformType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of host insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Number of days used for utilization forecast analysis.
	ForecastDays *int `mandatory:"false" contributesTo:"query" name:"forecastDays"`

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

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeHostInsightResourceUtilizationInsightRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeHostInsightResourceUtilizationInsightRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeHostInsightResourceUtilizationInsightRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeHostInsightResourceUtilizationInsightRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeHostInsightResourceUtilizationInsightRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.PlatformType {
		if _, ok := GetMappingSummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", val, strings.Join(GetSummarizeHostInsightResourceUtilizationInsightPlatformTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightResourceUtilizationInsightResponse wrapper for the SummarizeHostInsightResourceUtilizationInsight operation
type SummarizeHostInsightResourceUtilizationInsightResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeHostInsightResourceUtilizationInsightAggregation instances
	SummarizeHostInsightResourceUtilizationInsightAggregation `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response SummarizeHostInsightResourceUtilizationInsightResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeHostInsightResourceUtilizationInsightResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum Enum with underlying type: string
type SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum
const (
	SummarizeHostInsightResourceUtilizationInsightPlatformTypeLinux   SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum = "LINUX"
	SummarizeHostInsightResourceUtilizationInsightPlatformTypeSolaris SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum = "SOLARIS"
	SummarizeHostInsightResourceUtilizationInsightPlatformTypeSunos   SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum = "SUNOS"
	SummarizeHostInsightResourceUtilizationInsightPlatformTypeZlinux  SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum = "ZLINUX"
	SummarizeHostInsightResourceUtilizationInsightPlatformTypeWindows SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum = "WINDOWS"
	SummarizeHostInsightResourceUtilizationInsightPlatformTypeAix     SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum = "AIX"
	SummarizeHostInsightResourceUtilizationInsightPlatformTypeHpUx    SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum = "HP_UX"
)

var mappingSummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum = map[string]SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum{
	"LINUX":   SummarizeHostInsightResourceUtilizationInsightPlatformTypeLinux,
	"SOLARIS": SummarizeHostInsightResourceUtilizationInsightPlatformTypeSolaris,
	"SUNOS":   SummarizeHostInsightResourceUtilizationInsightPlatformTypeSunos,
	"ZLINUX":  SummarizeHostInsightResourceUtilizationInsightPlatformTypeZlinux,
	"WINDOWS": SummarizeHostInsightResourceUtilizationInsightPlatformTypeWindows,
	"AIX":     SummarizeHostInsightResourceUtilizationInsightPlatformTypeAix,
	"HP_UX":   SummarizeHostInsightResourceUtilizationInsightPlatformTypeHpUx,
}

var mappingSummarizeHostInsightResourceUtilizationInsightPlatformTypeEnumLowerCase = map[string]SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum{
	"linux":   SummarizeHostInsightResourceUtilizationInsightPlatformTypeLinux,
	"solaris": SummarizeHostInsightResourceUtilizationInsightPlatformTypeSolaris,
	"sunos":   SummarizeHostInsightResourceUtilizationInsightPlatformTypeSunos,
	"zlinux":  SummarizeHostInsightResourceUtilizationInsightPlatformTypeZlinux,
	"windows": SummarizeHostInsightResourceUtilizationInsightPlatformTypeWindows,
	"aix":     SummarizeHostInsightResourceUtilizationInsightPlatformTypeAix,
	"hp_ux":   SummarizeHostInsightResourceUtilizationInsightPlatformTypeHpUx,
}

// GetSummarizeHostInsightResourceUtilizationInsightPlatformTypeEnumValues Enumerates the set of values for SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum
func GetSummarizeHostInsightResourceUtilizationInsightPlatformTypeEnumValues() []SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum {
	values := make([]SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceUtilizationInsightPlatformTypeEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum
func GetSummarizeHostInsightResourceUtilizationInsightPlatformTypeEnumStringValues() []string {
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

// GetMappingSummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum(val string) (SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceUtilizationInsightPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
