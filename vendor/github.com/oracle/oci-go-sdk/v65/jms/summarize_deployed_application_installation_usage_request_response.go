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

// SummarizeDeployedApplicationInstallationUsageRequest wrapper for the SummarizeDeployedApplicationInstallationUsage operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeDeployedApplicationInstallationUsage.go.html to see an example of how to use SummarizeDeployedApplicationInstallationUsageRequest.
type SummarizeDeployedApplicationInstallationUsageRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The server key.
	ServerKey *string `mandatory:"false" contributesTo:"query" name:"serverKey"`

	// The Java Server instance key.
	ServerInstanceKey *string `mandatory:"false" contributesTo:"query" name:"serverInstanceKey"`

	// The Fleet-unique identifier of the managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The deployed application installation key.
	ApplicationInstallationKey *string `mandatory:"false" contributesTo:"query" name:"applicationInstallationKey"`

	// The deployed application key.
	ApplicationKey *string `mandatory:"false" contributesTo:"query" name:"applicationKey"`

	// Filter the list with deployed application name contains the given value.
	ApplicationNameContains *string `mandatory:"false" contributesTo:"query" name:"applicationNameContains"`

	// The deployed application name.
	ApplicationName *string `mandatory:"false" contributesTo:"query" name:"applicationName"`

	// Filter the list with application source path contains the given value.
	ApplicationSourcePathContains *string `mandatory:"false" contributesTo:"query" name:"applicationSourcePathContains"`

	// The library key.
	LibraryKey *string `mandatory:"false" contributesTo:"query" name:"libraryKey"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder SummarizeDeployedApplicationInstallationUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the deployed application installations. Only one sort order can be provided.
	// If no value is specified _timeLastSeen_ is default.
	//
	SortBy SummarizeDeployedApplicationInstallationUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeDeployedApplicationInstallationUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeDeployedApplicationInstallationUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeDeployedApplicationInstallationUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeDeployedApplicationInstallationUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeDeployedApplicationInstallationUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeDeployedApplicationInstallationUsageSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeDeployedApplicationInstallationUsageSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeDeployedApplicationInstallationUsageSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeDeployedApplicationInstallationUsageSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeDeployedApplicationInstallationUsageResponse wrapper for the SummarizeDeployedApplicationInstallationUsage operation
type SummarizeDeployedApplicationInstallationUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeployedApplicationInstallationUsageSummaryCollection instances
	DeployedApplicationInstallationUsageSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeDeployedApplicationInstallationUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeDeployedApplicationInstallationUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeDeployedApplicationInstallationUsageSortOrderEnum Enum with underlying type: string
type SummarizeDeployedApplicationInstallationUsageSortOrderEnum string

// Set of constants representing the allowable values for SummarizeDeployedApplicationInstallationUsageSortOrderEnum
const (
	SummarizeDeployedApplicationInstallationUsageSortOrderAsc  SummarizeDeployedApplicationInstallationUsageSortOrderEnum = "ASC"
	SummarizeDeployedApplicationInstallationUsageSortOrderDesc SummarizeDeployedApplicationInstallationUsageSortOrderEnum = "DESC"
)

var mappingSummarizeDeployedApplicationInstallationUsageSortOrderEnum = map[string]SummarizeDeployedApplicationInstallationUsageSortOrderEnum{
	"ASC":  SummarizeDeployedApplicationInstallationUsageSortOrderAsc,
	"DESC": SummarizeDeployedApplicationInstallationUsageSortOrderDesc,
}

var mappingSummarizeDeployedApplicationInstallationUsageSortOrderEnumLowerCase = map[string]SummarizeDeployedApplicationInstallationUsageSortOrderEnum{
	"asc":  SummarizeDeployedApplicationInstallationUsageSortOrderAsc,
	"desc": SummarizeDeployedApplicationInstallationUsageSortOrderDesc,
}

// GetSummarizeDeployedApplicationInstallationUsageSortOrderEnumValues Enumerates the set of values for SummarizeDeployedApplicationInstallationUsageSortOrderEnum
func GetSummarizeDeployedApplicationInstallationUsageSortOrderEnumValues() []SummarizeDeployedApplicationInstallationUsageSortOrderEnum {
	values := make([]SummarizeDeployedApplicationInstallationUsageSortOrderEnum, 0)
	for _, v := range mappingSummarizeDeployedApplicationInstallationUsageSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDeployedApplicationInstallationUsageSortOrderEnumStringValues Enumerates the set of values in String for SummarizeDeployedApplicationInstallationUsageSortOrderEnum
func GetSummarizeDeployedApplicationInstallationUsageSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeDeployedApplicationInstallationUsageSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDeployedApplicationInstallationUsageSortOrderEnum(val string) (SummarizeDeployedApplicationInstallationUsageSortOrderEnum, bool) {
	enum, ok := mappingSummarizeDeployedApplicationInstallationUsageSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeDeployedApplicationInstallationUsageSortByEnum Enum with underlying type: string
type SummarizeDeployedApplicationInstallationUsageSortByEnum string

// Set of constants representing the allowable values for SummarizeDeployedApplicationInstallationUsageSortByEnum
const (
	SummarizeDeployedApplicationInstallationUsageSortByApplicationname         SummarizeDeployedApplicationInstallationUsageSortByEnum = "applicationName"
	SummarizeDeployedApplicationInstallationUsageSortByApplicationtype         SummarizeDeployedApplicationInstallationUsageSortByEnum = "applicationType"
	SummarizeDeployedApplicationInstallationUsageSortByApplicationsourcepath   SummarizeDeployedApplicationInstallationUsageSortByEnum = "applicationSourcePath"
	SummarizeDeployedApplicationInstallationUsageSortByIsclustered             SummarizeDeployedApplicationInstallationUsageSortByEnum = "isClustered"
	SummarizeDeployedApplicationInstallationUsageSortByJavaserverinstancecount SummarizeDeployedApplicationInstallationUsageSortByEnum = "javaServerInstanceCount"
	SummarizeDeployedApplicationInstallationUsageSortByTimefirstseen           SummarizeDeployedApplicationInstallationUsageSortByEnum = "timeFirstSeen"
	SummarizeDeployedApplicationInstallationUsageSortByTimelastseen            SummarizeDeployedApplicationInstallationUsageSortByEnum = "timeLastSeen"
)

var mappingSummarizeDeployedApplicationInstallationUsageSortByEnum = map[string]SummarizeDeployedApplicationInstallationUsageSortByEnum{
	"applicationName":         SummarizeDeployedApplicationInstallationUsageSortByApplicationname,
	"applicationType":         SummarizeDeployedApplicationInstallationUsageSortByApplicationtype,
	"applicationSourcePath":   SummarizeDeployedApplicationInstallationUsageSortByApplicationsourcepath,
	"isClustered":             SummarizeDeployedApplicationInstallationUsageSortByIsclustered,
	"javaServerInstanceCount": SummarizeDeployedApplicationInstallationUsageSortByJavaserverinstancecount,
	"timeFirstSeen":           SummarizeDeployedApplicationInstallationUsageSortByTimefirstseen,
	"timeLastSeen":            SummarizeDeployedApplicationInstallationUsageSortByTimelastseen,
}

var mappingSummarizeDeployedApplicationInstallationUsageSortByEnumLowerCase = map[string]SummarizeDeployedApplicationInstallationUsageSortByEnum{
	"applicationname":         SummarizeDeployedApplicationInstallationUsageSortByApplicationname,
	"applicationtype":         SummarizeDeployedApplicationInstallationUsageSortByApplicationtype,
	"applicationsourcepath":   SummarizeDeployedApplicationInstallationUsageSortByApplicationsourcepath,
	"isclustered":             SummarizeDeployedApplicationInstallationUsageSortByIsclustered,
	"javaserverinstancecount": SummarizeDeployedApplicationInstallationUsageSortByJavaserverinstancecount,
	"timefirstseen":           SummarizeDeployedApplicationInstallationUsageSortByTimefirstseen,
	"timelastseen":            SummarizeDeployedApplicationInstallationUsageSortByTimelastseen,
}

// GetSummarizeDeployedApplicationInstallationUsageSortByEnumValues Enumerates the set of values for SummarizeDeployedApplicationInstallationUsageSortByEnum
func GetSummarizeDeployedApplicationInstallationUsageSortByEnumValues() []SummarizeDeployedApplicationInstallationUsageSortByEnum {
	values := make([]SummarizeDeployedApplicationInstallationUsageSortByEnum, 0)
	for _, v := range mappingSummarizeDeployedApplicationInstallationUsageSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDeployedApplicationInstallationUsageSortByEnumStringValues Enumerates the set of values in String for SummarizeDeployedApplicationInstallationUsageSortByEnum
func GetSummarizeDeployedApplicationInstallationUsageSortByEnumStringValues() []string {
	return []string{
		"applicationName",
		"applicationType",
		"applicationSourcePath",
		"isClustered",
		"javaServerInstanceCount",
		"timeFirstSeen",
		"timeLastSeen",
	}
}

// GetMappingSummarizeDeployedApplicationInstallationUsageSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDeployedApplicationInstallationUsageSortByEnum(val string) (SummarizeDeployedApplicationInstallationUsageSortByEnum, bool) {
	enum, ok := mappingSummarizeDeployedApplicationInstallationUsageSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
