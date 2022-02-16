// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// SummarizeManagedInstanceUsageRequest wrapper for the SummarizeManagedInstanceUsage operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeManagedInstanceUsage.go.html to see an example of how to use SummarizeManagedInstanceUsageRequest.
type SummarizeManagedInstanceUsageRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The Fleet-unique identifier of the managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The type of the managed instance.
	ManagedInstanceType SummarizeManagedInstanceUsageManagedInstanceTypeEnum `mandatory:"false" contributesTo:"query" name:"managedInstanceType" omitEmpty:"true"`

	// The vendor of the related Java Runtime.
	JreVendor *string `mandatory:"false" contributesTo:"query" name:"jreVendor"`

	// The distribution of the related Java Runtime.
	JreDistribution *string `mandatory:"false" contributesTo:"query" name:"jreDistribution"`

	// The version of the related Java Runtime.
	JreVersion *string `mandatory:"false" contributesTo:"query" name:"jreVersion"`

	// The file system path of the installation.
	InstallationPath *string `mandatory:"false" contributesTo:"query" name:"installationPath"`

	// The Fleet-unique identifier of the related application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// Additional fields to include into the returned model on top of the required ones.
	// This parameter can also include 'approximateJreCount', 'approximateInstallationCount' and 'approximateApplicationCount'.
	// For example 'approximateJreCount,approximateInstallationCount'.
	Fields []SummarizeManagedInstanceUsageFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder SummarizeManagedInstanceUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort managed instance views. Only one sort order may be provided.
	// Default order for _timeFirstSeen_, _timeLastSeen_, approximateJreCount_, _approximateInstallationCount_
	// and _approximateApplicationCount_  is **descending**.
	// Default order for _osName_ is **ascending**.
	// If no value is specified _timeLastSeen_ is default.
	SortBy SummarizeManagedInstanceUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The operating system type.
	OsFamily []OsFamilyEnum `contributesTo:"query" name:"osFamily" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeManagedInstanceUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeManagedInstanceUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeManagedInstanceUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeManagedInstanceUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeManagedInstanceUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeManagedInstanceUsageManagedInstanceTypeEnum(string(request.ManagedInstanceType)); !ok && request.ManagedInstanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedInstanceType: %s. Supported values are: %s.", request.ManagedInstanceType, strings.Join(GetSummarizeManagedInstanceUsageManagedInstanceTypeEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingSummarizeManagedInstanceUsageFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetSummarizeManagedInstanceUsageFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSummarizeManagedInstanceUsageSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeManagedInstanceUsageSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeManagedInstanceUsageSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeManagedInstanceUsageSortByEnumStringValues(), ",")))
	}
	for _, val := range request.OsFamily {
		if _, ok := GetMappingOsFamilyEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", val, strings.Join(GetOsFamilyEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeManagedInstanceUsageResponse wrapper for the SummarizeManagedInstanceUsage operation
type SummarizeManagedInstanceUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceUsageCollection instances
	ManagedInstanceUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeManagedInstanceUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeManagedInstanceUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeManagedInstanceUsageManagedInstanceTypeEnum Enum with underlying type: string
type SummarizeManagedInstanceUsageManagedInstanceTypeEnum string

// Set of constants representing the allowable values for SummarizeManagedInstanceUsageManagedInstanceTypeEnum
const (
	SummarizeManagedInstanceUsageManagedInstanceTypeOracleManagementAgent SummarizeManagedInstanceUsageManagedInstanceTypeEnum = "ORACLE_MANAGEMENT_AGENT"
)

var mappingSummarizeManagedInstanceUsageManagedInstanceTypeEnum = map[string]SummarizeManagedInstanceUsageManagedInstanceTypeEnum{
	"ORACLE_MANAGEMENT_AGENT": SummarizeManagedInstanceUsageManagedInstanceTypeOracleManagementAgent,
}

// GetSummarizeManagedInstanceUsageManagedInstanceTypeEnumValues Enumerates the set of values for SummarizeManagedInstanceUsageManagedInstanceTypeEnum
func GetSummarizeManagedInstanceUsageManagedInstanceTypeEnumValues() []SummarizeManagedInstanceUsageManagedInstanceTypeEnum {
	values := make([]SummarizeManagedInstanceUsageManagedInstanceTypeEnum, 0)
	for _, v := range mappingSummarizeManagedInstanceUsageManagedInstanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeManagedInstanceUsageManagedInstanceTypeEnumStringValues Enumerates the set of values in String for SummarizeManagedInstanceUsageManagedInstanceTypeEnum
func GetSummarizeManagedInstanceUsageManagedInstanceTypeEnumStringValues() []string {
	return []string{
		"ORACLE_MANAGEMENT_AGENT",
	}
}

// GetMappingSummarizeManagedInstanceUsageManagedInstanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeManagedInstanceUsageManagedInstanceTypeEnum(val string) (SummarizeManagedInstanceUsageManagedInstanceTypeEnum, bool) {
	mappingSummarizeManagedInstanceUsageManagedInstanceTypeEnumIgnoreCase := make(map[string]SummarizeManagedInstanceUsageManagedInstanceTypeEnum)
	for k, v := range mappingSummarizeManagedInstanceUsageManagedInstanceTypeEnum {
		mappingSummarizeManagedInstanceUsageManagedInstanceTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeManagedInstanceUsageManagedInstanceTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeManagedInstanceUsageSortOrderEnum Enum with underlying type: string
type SummarizeManagedInstanceUsageSortOrderEnum string

// Set of constants representing the allowable values for SummarizeManagedInstanceUsageSortOrderEnum
const (
	SummarizeManagedInstanceUsageSortOrderAsc  SummarizeManagedInstanceUsageSortOrderEnum = "ASC"
	SummarizeManagedInstanceUsageSortOrderDesc SummarizeManagedInstanceUsageSortOrderEnum = "DESC"
)

var mappingSummarizeManagedInstanceUsageSortOrderEnum = map[string]SummarizeManagedInstanceUsageSortOrderEnum{
	"ASC":  SummarizeManagedInstanceUsageSortOrderAsc,
	"DESC": SummarizeManagedInstanceUsageSortOrderDesc,
}

// GetSummarizeManagedInstanceUsageSortOrderEnumValues Enumerates the set of values for SummarizeManagedInstanceUsageSortOrderEnum
func GetSummarizeManagedInstanceUsageSortOrderEnumValues() []SummarizeManagedInstanceUsageSortOrderEnum {
	values := make([]SummarizeManagedInstanceUsageSortOrderEnum, 0)
	for _, v := range mappingSummarizeManagedInstanceUsageSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeManagedInstanceUsageSortOrderEnumStringValues Enumerates the set of values in String for SummarizeManagedInstanceUsageSortOrderEnum
func GetSummarizeManagedInstanceUsageSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeManagedInstanceUsageSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeManagedInstanceUsageSortOrderEnum(val string) (SummarizeManagedInstanceUsageSortOrderEnum, bool) {
	mappingSummarizeManagedInstanceUsageSortOrderEnumIgnoreCase := make(map[string]SummarizeManagedInstanceUsageSortOrderEnum)
	for k, v := range mappingSummarizeManagedInstanceUsageSortOrderEnum {
		mappingSummarizeManagedInstanceUsageSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeManagedInstanceUsageSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeManagedInstanceUsageSortByEnum Enum with underlying type: string
type SummarizeManagedInstanceUsageSortByEnum string

// Set of constants representing the allowable values for SummarizeManagedInstanceUsageSortByEnum
const (
	SummarizeManagedInstanceUsageSortByTimefirstseen                SummarizeManagedInstanceUsageSortByEnum = "timeFirstSeen"
	SummarizeManagedInstanceUsageSortByTimelastseen                 SummarizeManagedInstanceUsageSortByEnum = "timeLastSeen"
	SummarizeManagedInstanceUsageSortByApproximatejrecount          SummarizeManagedInstanceUsageSortByEnum = "approximateJreCount"
	SummarizeManagedInstanceUsageSortByApproximateinstallationcount SummarizeManagedInstanceUsageSortByEnum = "approximateInstallationCount"
	SummarizeManagedInstanceUsageSortByApproximateapplicationcount  SummarizeManagedInstanceUsageSortByEnum = "approximateApplicationCount"
	SummarizeManagedInstanceUsageSortByOsname                       SummarizeManagedInstanceUsageSortByEnum = "osName"
)

var mappingSummarizeManagedInstanceUsageSortByEnum = map[string]SummarizeManagedInstanceUsageSortByEnum{
	"timeFirstSeen":                SummarizeManagedInstanceUsageSortByTimefirstseen,
	"timeLastSeen":                 SummarizeManagedInstanceUsageSortByTimelastseen,
	"approximateJreCount":          SummarizeManagedInstanceUsageSortByApproximatejrecount,
	"approximateInstallationCount": SummarizeManagedInstanceUsageSortByApproximateinstallationcount,
	"approximateApplicationCount":  SummarizeManagedInstanceUsageSortByApproximateapplicationcount,
	"osName":                       SummarizeManagedInstanceUsageSortByOsname,
}

// GetSummarizeManagedInstanceUsageSortByEnumValues Enumerates the set of values for SummarizeManagedInstanceUsageSortByEnum
func GetSummarizeManagedInstanceUsageSortByEnumValues() []SummarizeManagedInstanceUsageSortByEnum {
	values := make([]SummarizeManagedInstanceUsageSortByEnum, 0)
	for _, v := range mappingSummarizeManagedInstanceUsageSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeManagedInstanceUsageSortByEnumStringValues Enumerates the set of values in String for SummarizeManagedInstanceUsageSortByEnum
func GetSummarizeManagedInstanceUsageSortByEnumStringValues() []string {
	return []string{
		"timeFirstSeen",
		"timeLastSeen",
		"approximateJreCount",
		"approximateInstallationCount",
		"approximateApplicationCount",
		"osName",
	}
}

// GetMappingSummarizeManagedInstanceUsageSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeManagedInstanceUsageSortByEnum(val string) (SummarizeManagedInstanceUsageSortByEnum, bool) {
	mappingSummarizeManagedInstanceUsageSortByEnumIgnoreCase := make(map[string]SummarizeManagedInstanceUsageSortByEnum)
	for k, v := range mappingSummarizeManagedInstanceUsageSortByEnum {
		mappingSummarizeManagedInstanceUsageSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeManagedInstanceUsageSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
