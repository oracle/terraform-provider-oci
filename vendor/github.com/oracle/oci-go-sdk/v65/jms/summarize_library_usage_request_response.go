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

// SummarizeLibraryUsageRequest wrapper for the SummarizeLibraryUsage operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeLibraryUsage.go.html to see an example of how to use SummarizeLibraryUsageRequest.
type SummarizeLibraryUsageRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The Java Server instance key.
	ServerInstanceKey *string `mandatory:"false" contributesTo:"query" name:"serverInstanceKey"`

	// The Fleet-unique identifier of the managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The deployed application key.
	ApplicationKey *string `mandatory:"false" contributesTo:"query" name:"applicationKey"`

	// The library key.
	LibraryKey *string `mandatory:"false" contributesTo:"query" name:"libraryKey"`

	// Filter the list with library name contains the given value.
	LibraryNameContains *string `mandatory:"false" contributesTo:"query" name:"libraryNameContains"`

	// The library name.
	LibraryName *string `mandatory:"false" contributesTo:"query" name:"libraryName"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder SummarizeLibraryUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort libraries.  Only one sort order may be provided.
	// If no value is specified _timeLastSeen_ is default.
	SortBy SummarizeLibraryUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeLibraryUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeLibraryUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeLibraryUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeLibraryUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeLibraryUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeLibraryUsageSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeLibraryUsageSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeLibraryUsageSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeLibraryUsageSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeLibraryUsageResponse wrapper for the SummarizeLibraryUsage operation
type SummarizeLibraryUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LibraryUsageCollection instances
	LibraryUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeLibraryUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeLibraryUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeLibraryUsageSortOrderEnum Enum with underlying type: string
type SummarizeLibraryUsageSortOrderEnum string

// Set of constants representing the allowable values for SummarizeLibraryUsageSortOrderEnum
const (
	SummarizeLibraryUsageSortOrderAsc  SummarizeLibraryUsageSortOrderEnum = "ASC"
	SummarizeLibraryUsageSortOrderDesc SummarizeLibraryUsageSortOrderEnum = "DESC"
)

var mappingSummarizeLibraryUsageSortOrderEnum = map[string]SummarizeLibraryUsageSortOrderEnum{
	"ASC":  SummarizeLibraryUsageSortOrderAsc,
	"DESC": SummarizeLibraryUsageSortOrderDesc,
}

var mappingSummarizeLibraryUsageSortOrderEnumLowerCase = map[string]SummarizeLibraryUsageSortOrderEnum{
	"asc":  SummarizeLibraryUsageSortOrderAsc,
	"desc": SummarizeLibraryUsageSortOrderDesc,
}

// GetSummarizeLibraryUsageSortOrderEnumValues Enumerates the set of values for SummarizeLibraryUsageSortOrderEnum
func GetSummarizeLibraryUsageSortOrderEnumValues() []SummarizeLibraryUsageSortOrderEnum {
	values := make([]SummarizeLibraryUsageSortOrderEnum, 0)
	for _, v := range mappingSummarizeLibraryUsageSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeLibraryUsageSortOrderEnumStringValues Enumerates the set of values in String for SummarizeLibraryUsageSortOrderEnum
func GetSummarizeLibraryUsageSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeLibraryUsageSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeLibraryUsageSortOrderEnum(val string) (SummarizeLibraryUsageSortOrderEnum, bool) {
	enum, ok := mappingSummarizeLibraryUsageSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeLibraryUsageSortByEnum Enum with underlying type: string
type SummarizeLibraryUsageSortByEnum string

// Set of constants representing the allowable values for SummarizeLibraryUsageSortByEnum
const (
	SummarizeLibraryUsageSortByApplicationcount         SummarizeLibraryUsageSortByEnum = "applicationCount"
	SummarizeLibraryUsageSortByJavaserverinstancecount  SummarizeLibraryUsageSortByEnum = "javaServerInstanceCount"
	SummarizeLibraryUsageSortByCvssscore                SummarizeLibraryUsageSortByEnum = "cvssScore"
	SummarizeLibraryUsageSortByDeployedapplicationcount SummarizeLibraryUsageSortByEnum = "deployedApplicationCount"
	SummarizeLibraryUsageSortByLibraryname              SummarizeLibraryUsageSortByEnum = "libraryName"
	SummarizeLibraryUsageSortByLibraryversion           SummarizeLibraryUsageSortByEnum = "libraryVersion"
	SummarizeLibraryUsageSortByManagedinstancecount     SummarizeLibraryUsageSortByEnum = "managedInstanceCount"
	SummarizeLibraryUsageSortByTimefirstseen            SummarizeLibraryUsageSortByEnum = "timeFirstSeen"
	SummarizeLibraryUsageSortByTimelastseen             SummarizeLibraryUsageSortByEnum = "timeLastSeen"
)

var mappingSummarizeLibraryUsageSortByEnum = map[string]SummarizeLibraryUsageSortByEnum{
	"applicationCount":         SummarizeLibraryUsageSortByApplicationcount,
	"javaServerInstanceCount":  SummarizeLibraryUsageSortByJavaserverinstancecount,
	"cvssScore":                SummarizeLibraryUsageSortByCvssscore,
	"deployedApplicationCount": SummarizeLibraryUsageSortByDeployedapplicationcount,
	"libraryName":              SummarizeLibraryUsageSortByLibraryname,
	"libraryVersion":           SummarizeLibraryUsageSortByLibraryversion,
	"managedInstanceCount":     SummarizeLibraryUsageSortByManagedinstancecount,
	"timeFirstSeen":            SummarizeLibraryUsageSortByTimefirstseen,
	"timeLastSeen":             SummarizeLibraryUsageSortByTimelastseen,
}

var mappingSummarizeLibraryUsageSortByEnumLowerCase = map[string]SummarizeLibraryUsageSortByEnum{
	"applicationcount":         SummarizeLibraryUsageSortByApplicationcount,
	"javaserverinstancecount":  SummarizeLibraryUsageSortByJavaserverinstancecount,
	"cvssscore":                SummarizeLibraryUsageSortByCvssscore,
	"deployedapplicationcount": SummarizeLibraryUsageSortByDeployedapplicationcount,
	"libraryname":              SummarizeLibraryUsageSortByLibraryname,
	"libraryversion":           SummarizeLibraryUsageSortByLibraryversion,
	"managedinstancecount":     SummarizeLibraryUsageSortByManagedinstancecount,
	"timefirstseen":            SummarizeLibraryUsageSortByTimefirstseen,
	"timelastseen":             SummarizeLibraryUsageSortByTimelastseen,
}

// GetSummarizeLibraryUsageSortByEnumValues Enumerates the set of values for SummarizeLibraryUsageSortByEnum
func GetSummarizeLibraryUsageSortByEnumValues() []SummarizeLibraryUsageSortByEnum {
	values := make([]SummarizeLibraryUsageSortByEnum, 0)
	for _, v := range mappingSummarizeLibraryUsageSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeLibraryUsageSortByEnumStringValues Enumerates the set of values in String for SummarizeLibraryUsageSortByEnum
func GetSummarizeLibraryUsageSortByEnumStringValues() []string {
	return []string{
		"applicationCount",
		"javaServerInstanceCount",
		"cvssScore",
		"deployedApplicationCount",
		"libraryName",
		"libraryVersion",
		"managedInstanceCount",
		"timeFirstSeen",
		"timeLastSeen",
	}
}

// GetMappingSummarizeLibraryUsageSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeLibraryUsageSortByEnum(val string) (SummarizeLibraryUsageSortByEnum, bool) {
	enum, ok := mappingSummarizeLibraryUsageSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
