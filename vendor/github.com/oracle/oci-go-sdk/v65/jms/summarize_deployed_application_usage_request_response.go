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

// SummarizeDeployedApplicationUsageRequest wrapper for the SummarizeDeployedApplicationUsage operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeDeployedApplicationUsage.go.html to see an example of how to use SummarizeDeployedApplicationUsageRequest.
type SummarizeDeployedApplicationUsageRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The server key.
	ServerKey *string `mandatory:"false" contributesTo:"query" name:"serverKey"`

	// The Java Server instance key.
	ServerInstanceKey *string `mandatory:"false" contributesTo:"query" name:"serverInstanceKey"`

	// The Fleet-unique identifier of the managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The library key.
	LibraryKey *string `mandatory:"false" contributesTo:"query" name:"libraryKey"`

	// The deployed application key.
	ApplicationKey *string `mandatory:"false" contributesTo:"query" name:"applicationKey"`

	// Filter the list with deployed application name contains the given value.
	ApplicationNameContains *string `mandatory:"false" contributesTo:"query" name:"applicationNameContains"`

	// The deployed application name.
	ApplicationName *string `mandatory:"false" contributesTo:"query" name:"applicationName"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder SummarizeDeployedApplicationUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the deployed applications. Only one sort order can be provided.
	// If no value is specified _timeLastSeen_ is default.
	SortBy SummarizeDeployedApplicationUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeDeployedApplicationUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeDeployedApplicationUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeDeployedApplicationUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeDeployedApplicationUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeDeployedApplicationUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeDeployedApplicationUsageSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeDeployedApplicationUsageSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeDeployedApplicationUsageSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeDeployedApplicationUsageSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeDeployedApplicationUsageResponse wrapper for the SummarizeDeployedApplicationUsage operation
type SummarizeDeployedApplicationUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeployedApplicationUsageCollection instances
	DeployedApplicationUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeDeployedApplicationUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeDeployedApplicationUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeDeployedApplicationUsageSortOrderEnum Enum with underlying type: string
type SummarizeDeployedApplicationUsageSortOrderEnum string

// Set of constants representing the allowable values for SummarizeDeployedApplicationUsageSortOrderEnum
const (
	SummarizeDeployedApplicationUsageSortOrderAsc  SummarizeDeployedApplicationUsageSortOrderEnum = "ASC"
	SummarizeDeployedApplicationUsageSortOrderDesc SummarizeDeployedApplicationUsageSortOrderEnum = "DESC"
)

var mappingSummarizeDeployedApplicationUsageSortOrderEnum = map[string]SummarizeDeployedApplicationUsageSortOrderEnum{
	"ASC":  SummarizeDeployedApplicationUsageSortOrderAsc,
	"DESC": SummarizeDeployedApplicationUsageSortOrderDesc,
}

var mappingSummarizeDeployedApplicationUsageSortOrderEnumLowerCase = map[string]SummarizeDeployedApplicationUsageSortOrderEnum{
	"asc":  SummarizeDeployedApplicationUsageSortOrderAsc,
	"desc": SummarizeDeployedApplicationUsageSortOrderDesc,
}

// GetSummarizeDeployedApplicationUsageSortOrderEnumValues Enumerates the set of values for SummarizeDeployedApplicationUsageSortOrderEnum
func GetSummarizeDeployedApplicationUsageSortOrderEnumValues() []SummarizeDeployedApplicationUsageSortOrderEnum {
	values := make([]SummarizeDeployedApplicationUsageSortOrderEnum, 0)
	for _, v := range mappingSummarizeDeployedApplicationUsageSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDeployedApplicationUsageSortOrderEnumStringValues Enumerates the set of values in String for SummarizeDeployedApplicationUsageSortOrderEnum
func GetSummarizeDeployedApplicationUsageSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeDeployedApplicationUsageSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDeployedApplicationUsageSortOrderEnum(val string) (SummarizeDeployedApplicationUsageSortOrderEnum, bool) {
	enum, ok := mappingSummarizeDeployedApplicationUsageSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeDeployedApplicationUsageSortByEnum Enum with underlying type: string
type SummarizeDeployedApplicationUsageSortByEnum string

// Set of constants representing the allowable values for SummarizeDeployedApplicationUsageSortByEnum
const (
	SummarizeDeployedApplicationUsageSortByApplicationname                    SummarizeDeployedApplicationUsageSortByEnum = "applicationName"
	SummarizeDeployedApplicationUsageSortByApplicationtype                    SummarizeDeployedApplicationUsageSortByEnum = "applicationType"
	SummarizeDeployedApplicationUsageSortByIsclustered                        SummarizeDeployedApplicationUsageSortByEnum = "isClustered"
	SummarizeDeployedApplicationUsageSortByJavaserverinstancecount            SummarizeDeployedApplicationUsageSortByEnum = "javaServerInstanceCount"
	SummarizeDeployedApplicationUsageSortByApproximatejavaserverinstancecount SummarizeDeployedApplicationUsageSortByEnum = "approximateJavaServerInstanceCount"
	SummarizeDeployedApplicationUsageSortByApproximatelibrarycount            SummarizeDeployedApplicationUsageSortByEnum = "approximateLibraryCount"
	SummarizeDeployedApplicationUsageSortByTimefirstseen                      SummarizeDeployedApplicationUsageSortByEnum = "timeFirstSeen"
	SummarizeDeployedApplicationUsageSortByTimelastseen                       SummarizeDeployedApplicationUsageSortByEnum = "timeLastSeen"
)

var mappingSummarizeDeployedApplicationUsageSortByEnum = map[string]SummarizeDeployedApplicationUsageSortByEnum{
	"applicationName":                    SummarizeDeployedApplicationUsageSortByApplicationname,
	"applicationType":                    SummarizeDeployedApplicationUsageSortByApplicationtype,
	"isClustered":                        SummarizeDeployedApplicationUsageSortByIsclustered,
	"javaServerInstanceCount":            SummarizeDeployedApplicationUsageSortByJavaserverinstancecount,
	"approximateJavaServerInstanceCount": SummarizeDeployedApplicationUsageSortByApproximatejavaserverinstancecount,
	"approximateLibraryCount":            SummarizeDeployedApplicationUsageSortByApproximatelibrarycount,
	"timeFirstSeen":                      SummarizeDeployedApplicationUsageSortByTimefirstseen,
	"timeLastSeen":                       SummarizeDeployedApplicationUsageSortByTimelastseen,
}

var mappingSummarizeDeployedApplicationUsageSortByEnumLowerCase = map[string]SummarizeDeployedApplicationUsageSortByEnum{
	"applicationname":                    SummarizeDeployedApplicationUsageSortByApplicationname,
	"applicationtype":                    SummarizeDeployedApplicationUsageSortByApplicationtype,
	"isclustered":                        SummarizeDeployedApplicationUsageSortByIsclustered,
	"javaserverinstancecount":            SummarizeDeployedApplicationUsageSortByJavaserverinstancecount,
	"approximatejavaserverinstancecount": SummarizeDeployedApplicationUsageSortByApproximatejavaserverinstancecount,
	"approximatelibrarycount":            SummarizeDeployedApplicationUsageSortByApproximatelibrarycount,
	"timefirstseen":                      SummarizeDeployedApplicationUsageSortByTimefirstseen,
	"timelastseen":                       SummarizeDeployedApplicationUsageSortByTimelastseen,
}

// GetSummarizeDeployedApplicationUsageSortByEnumValues Enumerates the set of values for SummarizeDeployedApplicationUsageSortByEnum
func GetSummarizeDeployedApplicationUsageSortByEnumValues() []SummarizeDeployedApplicationUsageSortByEnum {
	values := make([]SummarizeDeployedApplicationUsageSortByEnum, 0)
	for _, v := range mappingSummarizeDeployedApplicationUsageSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDeployedApplicationUsageSortByEnumStringValues Enumerates the set of values in String for SummarizeDeployedApplicationUsageSortByEnum
func GetSummarizeDeployedApplicationUsageSortByEnumStringValues() []string {
	return []string{
		"applicationName",
		"applicationType",
		"isClustered",
		"javaServerInstanceCount",
		"approximateJavaServerInstanceCount",
		"approximateLibraryCount",
		"timeFirstSeen",
		"timeLastSeen",
	}
}

// GetMappingSummarizeDeployedApplicationUsageSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDeployedApplicationUsageSortByEnum(val string) (SummarizeDeployedApplicationUsageSortByEnum, bool) {
	enum, ok := mappingSummarizeDeployedApplicationUsageSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
