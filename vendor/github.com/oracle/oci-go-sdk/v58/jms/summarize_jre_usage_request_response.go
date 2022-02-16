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

// SummarizeJreUsageRequest wrapper for the SummarizeJreUsage operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeJreUsage.go.html to see an example of how to use SummarizeJreUsageRequest.
type SummarizeJreUsageRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The vendor of the Java Runtime.
	JreVendor *string `mandatory:"false" contributesTo:"query" name:"jreVendor"`

	// The distribution of the Java Runtime.
	JreDistribution *string `mandatory:"false" contributesTo:"query" name:"jreDistribution"`

	// The version of the Java Runtime.
	JreVersion *string `mandatory:"false" contributesTo:"query" name:"jreVersion"`

	// The Fleet-unique identifier of the related application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The Fleet-unique identifier of the related managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// Additional fields to include into the returned model on top of the required ones.
	// This parameter can also include 'approximateApplicationCount', 'approximateInstallationCount' and 'approximateManagedInstanceCount'.
	// For example 'approximateApplicationCount,approximateManagedInstanceCount'.
	Fields []SummarizeJreUsageFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder SummarizeJreUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort JRE usages. Only one sort order may be provided.
	// Default order for _timeFirstSeen_, _timeLastSeen_, and _version_ is **descending**.
	// Default order for _timeFirstSeen_, _timeLastSeen_, _version_, _approximateInstallationCount_,
	// _approximateApplicationCount_ and _approximateManagedInstanceCount_  is **descending**.
	// Default order for _distribution_, _vendor_, and _osName_ is **ascending**.
	// If no value is specified _timeLastSeen_ is default.
	SortBy SummarizeJreUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The operating system type.
	OsFamily []OsFamilyEnum `contributesTo:"query" name:"osFamily" omitEmpty:"true" collectionFormat:"multi"`

	// The security status of the Java Runtime.
	JreSecurityStatus SummarizeJreUsageJreSecurityStatusEnum `mandatory:"false" contributesTo:"query" name:"jreSecurityStatus" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeJreUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeJreUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeJreUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeJreUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeJreUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingSummarizeJreUsageFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetSummarizeJreUsageFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSummarizeJreUsageSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeJreUsageSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeJreUsageSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeJreUsageSortByEnumStringValues(), ",")))
	}
	for _, val := range request.OsFamily {
		if _, ok := GetMappingOsFamilyEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", val, strings.Join(GetOsFamilyEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSummarizeJreUsageJreSecurityStatusEnum(string(request.JreSecurityStatus)); !ok && request.JreSecurityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JreSecurityStatus: %s. Supported values are: %s.", request.JreSecurityStatus, strings.Join(GetSummarizeJreUsageJreSecurityStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeJreUsageResponse wrapper for the SummarizeJreUsage operation
type SummarizeJreUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JreUsageCollection instances
	JreUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeJreUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeJreUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeJreUsageSortOrderEnum Enum with underlying type: string
type SummarizeJreUsageSortOrderEnum string

// Set of constants representing the allowable values for SummarizeJreUsageSortOrderEnum
const (
	SummarizeJreUsageSortOrderAsc  SummarizeJreUsageSortOrderEnum = "ASC"
	SummarizeJreUsageSortOrderDesc SummarizeJreUsageSortOrderEnum = "DESC"
)

var mappingSummarizeJreUsageSortOrderEnum = map[string]SummarizeJreUsageSortOrderEnum{
	"ASC":  SummarizeJreUsageSortOrderAsc,
	"DESC": SummarizeJreUsageSortOrderDesc,
}

// GetSummarizeJreUsageSortOrderEnumValues Enumerates the set of values for SummarizeJreUsageSortOrderEnum
func GetSummarizeJreUsageSortOrderEnumValues() []SummarizeJreUsageSortOrderEnum {
	values := make([]SummarizeJreUsageSortOrderEnum, 0)
	for _, v := range mappingSummarizeJreUsageSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeJreUsageSortOrderEnumStringValues Enumerates the set of values in String for SummarizeJreUsageSortOrderEnum
func GetSummarizeJreUsageSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeJreUsageSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeJreUsageSortOrderEnum(val string) (SummarizeJreUsageSortOrderEnum, bool) {
	mappingSummarizeJreUsageSortOrderEnumIgnoreCase := make(map[string]SummarizeJreUsageSortOrderEnum)
	for k, v := range mappingSummarizeJreUsageSortOrderEnum {
		mappingSummarizeJreUsageSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeJreUsageSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeJreUsageSortByEnum Enum with underlying type: string
type SummarizeJreUsageSortByEnum string

// Set of constants representing the allowable values for SummarizeJreUsageSortByEnum
const (
	SummarizeJreUsageSortByDistribution                    SummarizeJreUsageSortByEnum = "distribution"
	SummarizeJreUsageSortByTimefirstseen                   SummarizeJreUsageSortByEnum = "timeFirstSeen"
	SummarizeJreUsageSortByTimelastseen                    SummarizeJreUsageSortByEnum = "timeLastSeen"
	SummarizeJreUsageSortByVendor                          SummarizeJreUsageSortByEnum = "vendor"
	SummarizeJreUsageSortByVersion                         SummarizeJreUsageSortByEnum = "version"
	SummarizeJreUsageSortByApproximateinstallationcount    SummarizeJreUsageSortByEnum = "approximateInstallationCount"
	SummarizeJreUsageSortByApproximateapplicationcount     SummarizeJreUsageSortByEnum = "approximateApplicationCount"
	SummarizeJreUsageSortByApproximatemanagedinstancecount SummarizeJreUsageSortByEnum = "approximateManagedInstanceCount"
	SummarizeJreUsageSortByOsname                          SummarizeJreUsageSortByEnum = "osName"
	SummarizeJreUsageSortBySecuritystatus                  SummarizeJreUsageSortByEnum = "securityStatus"
)

var mappingSummarizeJreUsageSortByEnum = map[string]SummarizeJreUsageSortByEnum{
	"distribution":                    SummarizeJreUsageSortByDistribution,
	"timeFirstSeen":                   SummarizeJreUsageSortByTimefirstseen,
	"timeLastSeen":                    SummarizeJreUsageSortByTimelastseen,
	"vendor":                          SummarizeJreUsageSortByVendor,
	"version":                         SummarizeJreUsageSortByVersion,
	"approximateInstallationCount":    SummarizeJreUsageSortByApproximateinstallationcount,
	"approximateApplicationCount":     SummarizeJreUsageSortByApproximateapplicationcount,
	"approximateManagedInstanceCount": SummarizeJreUsageSortByApproximatemanagedinstancecount,
	"osName":                          SummarizeJreUsageSortByOsname,
	"securityStatus":                  SummarizeJreUsageSortBySecuritystatus,
}

// GetSummarizeJreUsageSortByEnumValues Enumerates the set of values for SummarizeJreUsageSortByEnum
func GetSummarizeJreUsageSortByEnumValues() []SummarizeJreUsageSortByEnum {
	values := make([]SummarizeJreUsageSortByEnum, 0)
	for _, v := range mappingSummarizeJreUsageSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeJreUsageSortByEnumStringValues Enumerates the set of values in String for SummarizeJreUsageSortByEnum
func GetSummarizeJreUsageSortByEnumStringValues() []string {
	return []string{
		"distribution",
		"timeFirstSeen",
		"timeLastSeen",
		"vendor",
		"version",
		"approximateInstallationCount",
		"approximateApplicationCount",
		"approximateManagedInstanceCount",
		"osName",
		"securityStatus",
	}
}

// GetMappingSummarizeJreUsageSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeJreUsageSortByEnum(val string) (SummarizeJreUsageSortByEnum, bool) {
	mappingSummarizeJreUsageSortByEnumIgnoreCase := make(map[string]SummarizeJreUsageSortByEnum)
	for k, v := range mappingSummarizeJreUsageSortByEnum {
		mappingSummarizeJreUsageSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeJreUsageSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeJreUsageJreSecurityStatusEnum Enum with underlying type: string
type SummarizeJreUsageJreSecurityStatusEnum string

// Set of constants representing the allowable values for SummarizeJreUsageJreSecurityStatusEnum
const (
	SummarizeJreUsageJreSecurityStatusUnknown         SummarizeJreUsageJreSecurityStatusEnum = "UNKNOWN"
	SummarizeJreUsageJreSecurityStatusUpToDate        SummarizeJreUsageJreSecurityStatusEnum = "UP_TO_DATE"
	SummarizeJreUsageJreSecurityStatusUpdateRequired  SummarizeJreUsageJreSecurityStatusEnum = "UPDATE_REQUIRED"
	SummarizeJreUsageJreSecurityStatusUpgradeRequired SummarizeJreUsageJreSecurityStatusEnum = "UPGRADE_REQUIRED"
)

var mappingSummarizeJreUsageJreSecurityStatusEnum = map[string]SummarizeJreUsageJreSecurityStatusEnum{
	"UNKNOWN":          SummarizeJreUsageJreSecurityStatusUnknown,
	"UP_TO_DATE":       SummarizeJreUsageJreSecurityStatusUpToDate,
	"UPDATE_REQUIRED":  SummarizeJreUsageJreSecurityStatusUpdateRequired,
	"UPGRADE_REQUIRED": SummarizeJreUsageJreSecurityStatusUpgradeRequired,
}

// GetSummarizeJreUsageJreSecurityStatusEnumValues Enumerates the set of values for SummarizeJreUsageJreSecurityStatusEnum
func GetSummarizeJreUsageJreSecurityStatusEnumValues() []SummarizeJreUsageJreSecurityStatusEnum {
	values := make([]SummarizeJreUsageJreSecurityStatusEnum, 0)
	for _, v := range mappingSummarizeJreUsageJreSecurityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeJreUsageJreSecurityStatusEnumStringValues Enumerates the set of values in String for SummarizeJreUsageJreSecurityStatusEnum
func GetSummarizeJreUsageJreSecurityStatusEnumStringValues() []string {
	return []string{
		"UNKNOWN",
		"UP_TO_DATE",
		"UPDATE_REQUIRED",
		"UPGRADE_REQUIRED",
	}
}

// GetMappingSummarizeJreUsageJreSecurityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeJreUsageJreSecurityStatusEnum(val string) (SummarizeJreUsageJreSecurityStatusEnum, bool) {
	mappingSummarizeJreUsageJreSecurityStatusEnumIgnoreCase := make(map[string]SummarizeJreUsageJreSecurityStatusEnum)
	for k, v := range mappingSummarizeJreUsageJreSecurityStatusEnum {
		mappingSummarizeJreUsageJreSecurityStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeJreUsageJreSecurityStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
