// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeApplicationInstallationUsageRequest wrapper for the SummarizeApplicationInstallationUsage operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeApplicationInstallationUsage.go.html to see an example of how to use SummarizeApplicationInstallationUsageRequest.
type SummarizeApplicationInstallationUsageRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The Fleet-unique identifier of the application installation.
	ApplicationInstallationKey *string `mandatory:"false" contributesTo:"query" name:"applicationInstallationKey"`

	// The Fleet-unique identifier of the application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter the list with displayName contains the given value.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The type of the application.
	ApplicationType *string `mandatory:"false" contributesTo:"query" name:"applicationType"`

	// Filter the list with the application installation path that contains the given value.
	AppInstallationPathContains *string `mandatory:"false" contributesTo:"query" name:"appInstallationPathContains"`

	// The vendor of the related Java Runtime.
	JreVendor *string `mandatory:"false" contributesTo:"query" name:"jreVendor"`

	// The distribution of the related Java Runtime.
	JreDistribution *string `mandatory:"false" contributesTo:"query" name:"jreDistribution"`

	// The version of the related Java Runtime.
	JreVersion *string `mandatory:"false" contributesTo:"query" name:"jreVersion"`

	// The file system path of the Java Runtime installation.
	InstallationPath *string `mandatory:"false" contributesTo:"query" name:"installationPath"`

	// The library key.
	LibraryKey *string `mandatory:"false" contributesTo:"query" name:"libraryKey"`

	// The Fleet-unique identifier of the related managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The operating system type.
	OsFamily []OsFamilyEnum `contributesTo:"query" name:"osFamily" omitEmpty:"true" collectionFormat:"multi"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder SummarizeApplicationInstallationUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort application installation views. Only one sort order may be provided.
	// Default order for _timeFirstSeen_, _timeLastSeen_, _approximateJreCount_, _approximateInstallationCount_
	// and _approximateManagedInstanceCount_  is **descending**.
	// Default order for _displayName_, _installationPath_ and _osName_ is **ascending**.
	// If no value is specified _timeLastSeen_ is default.
	SortBy SummarizeApplicationInstallationUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeApplicationInstallationUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeApplicationInstallationUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeApplicationInstallationUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeApplicationInstallationUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeApplicationInstallationUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.OsFamily {
		if _, ok := GetMappingOsFamilyEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", val, strings.Join(GetOsFamilyEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSummarizeApplicationInstallationUsageSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeApplicationInstallationUsageSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeApplicationInstallationUsageSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeApplicationInstallationUsageSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeApplicationInstallationUsageResponse wrapper for the SummarizeApplicationInstallationUsage operation
type SummarizeApplicationInstallationUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ApplicationInstallationUsageSummaryCollection instances
	ApplicationInstallationUsageSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeApplicationInstallationUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeApplicationInstallationUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeApplicationInstallationUsageSortOrderEnum Enum with underlying type: string
type SummarizeApplicationInstallationUsageSortOrderEnum string

// Set of constants representing the allowable values for SummarizeApplicationInstallationUsageSortOrderEnum
const (
	SummarizeApplicationInstallationUsageSortOrderAsc  SummarizeApplicationInstallationUsageSortOrderEnum = "ASC"
	SummarizeApplicationInstallationUsageSortOrderDesc SummarizeApplicationInstallationUsageSortOrderEnum = "DESC"
)

var mappingSummarizeApplicationInstallationUsageSortOrderEnum = map[string]SummarizeApplicationInstallationUsageSortOrderEnum{
	"ASC":  SummarizeApplicationInstallationUsageSortOrderAsc,
	"DESC": SummarizeApplicationInstallationUsageSortOrderDesc,
}

var mappingSummarizeApplicationInstallationUsageSortOrderEnumLowerCase = map[string]SummarizeApplicationInstallationUsageSortOrderEnum{
	"asc":  SummarizeApplicationInstallationUsageSortOrderAsc,
	"desc": SummarizeApplicationInstallationUsageSortOrderDesc,
}

// GetSummarizeApplicationInstallationUsageSortOrderEnumValues Enumerates the set of values for SummarizeApplicationInstallationUsageSortOrderEnum
func GetSummarizeApplicationInstallationUsageSortOrderEnumValues() []SummarizeApplicationInstallationUsageSortOrderEnum {
	values := make([]SummarizeApplicationInstallationUsageSortOrderEnum, 0)
	for _, v := range mappingSummarizeApplicationInstallationUsageSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeApplicationInstallationUsageSortOrderEnumStringValues Enumerates the set of values in String for SummarizeApplicationInstallationUsageSortOrderEnum
func GetSummarizeApplicationInstallationUsageSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeApplicationInstallationUsageSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeApplicationInstallationUsageSortOrderEnum(val string) (SummarizeApplicationInstallationUsageSortOrderEnum, bool) {
	enum, ok := mappingSummarizeApplicationInstallationUsageSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeApplicationInstallationUsageSortByEnum Enum with underlying type: string
type SummarizeApplicationInstallationUsageSortByEnum string

// Set of constants representing the allowable values for SummarizeApplicationInstallationUsageSortByEnum
const (
	SummarizeApplicationInstallationUsageSortByTimefirstseen                   SummarizeApplicationInstallationUsageSortByEnum = "timeFirstSeen"
	SummarizeApplicationInstallationUsageSortByTimelastseen                    SummarizeApplicationInstallationUsageSortByEnum = "timeLastSeen"
	SummarizeApplicationInstallationUsageSortByDisplayname                     SummarizeApplicationInstallationUsageSortByEnum = "displayName"
	SummarizeApplicationInstallationUsageSortByInstallationpath                SummarizeApplicationInstallationUsageSortByEnum = "installationPath"
	SummarizeApplicationInstallationUsageSortByOsname                          SummarizeApplicationInstallationUsageSortByEnum = "osName"
	SummarizeApplicationInstallationUsageSortByApproximatejrecount             SummarizeApplicationInstallationUsageSortByEnum = "approximateJreCount"
	SummarizeApplicationInstallationUsageSortByApproximateinstallationcount    SummarizeApplicationInstallationUsageSortByEnum = "approximateInstallationCount"
	SummarizeApplicationInstallationUsageSortByApproximatemanagedinstancecount SummarizeApplicationInstallationUsageSortByEnum = "approximateManagedInstanceCount"
)

var mappingSummarizeApplicationInstallationUsageSortByEnum = map[string]SummarizeApplicationInstallationUsageSortByEnum{
	"timeFirstSeen":                   SummarizeApplicationInstallationUsageSortByTimefirstseen,
	"timeLastSeen":                    SummarizeApplicationInstallationUsageSortByTimelastseen,
	"displayName":                     SummarizeApplicationInstallationUsageSortByDisplayname,
	"installationPath":                SummarizeApplicationInstallationUsageSortByInstallationpath,
	"osName":                          SummarizeApplicationInstallationUsageSortByOsname,
	"approximateJreCount":             SummarizeApplicationInstallationUsageSortByApproximatejrecount,
	"approximateInstallationCount":    SummarizeApplicationInstallationUsageSortByApproximateinstallationcount,
	"approximateManagedInstanceCount": SummarizeApplicationInstallationUsageSortByApproximatemanagedinstancecount,
}

var mappingSummarizeApplicationInstallationUsageSortByEnumLowerCase = map[string]SummarizeApplicationInstallationUsageSortByEnum{
	"timefirstseen":                   SummarizeApplicationInstallationUsageSortByTimefirstseen,
	"timelastseen":                    SummarizeApplicationInstallationUsageSortByTimelastseen,
	"displayname":                     SummarizeApplicationInstallationUsageSortByDisplayname,
	"installationpath":                SummarizeApplicationInstallationUsageSortByInstallationpath,
	"osname":                          SummarizeApplicationInstallationUsageSortByOsname,
	"approximatejrecount":             SummarizeApplicationInstallationUsageSortByApproximatejrecount,
	"approximateinstallationcount":    SummarizeApplicationInstallationUsageSortByApproximateinstallationcount,
	"approximatemanagedinstancecount": SummarizeApplicationInstallationUsageSortByApproximatemanagedinstancecount,
}

// GetSummarizeApplicationInstallationUsageSortByEnumValues Enumerates the set of values for SummarizeApplicationInstallationUsageSortByEnum
func GetSummarizeApplicationInstallationUsageSortByEnumValues() []SummarizeApplicationInstallationUsageSortByEnum {
	values := make([]SummarizeApplicationInstallationUsageSortByEnum, 0)
	for _, v := range mappingSummarizeApplicationInstallationUsageSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeApplicationInstallationUsageSortByEnumStringValues Enumerates the set of values in String for SummarizeApplicationInstallationUsageSortByEnum
func GetSummarizeApplicationInstallationUsageSortByEnumStringValues() []string {
	return []string{
		"timeFirstSeen",
		"timeLastSeen",
		"displayName",
		"installationPath",
		"osName",
		"approximateJreCount",
		"approximateInstallationCount",
		"approximateManagedInstanceCount",
	}
}

// GetMappingSummarizeApplicationInstallationUsageSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeApplicationInstallationUsageSortByEnum(val string) (SummarizeApplicationInstallationUsageSortByEnum, bool) {
	enum, ok := mappingSummarizeApplicationInstallationUsageSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
