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

// SummarizeJavaServerInstanceUsageRequest wrapper for the SummarizeJavaServerInstanceUsage operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeJavaServerInstanceUsage.go.html to see an example of how to use SummarizeJavaServerInstanceUsageRequest.
type SummarizeJavaServerInstanceUsageRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The server key.
	ServerKey *string `mandatory:"false" contributesTo:"query" name:"serverKey"`

	// The Java Server instance key.
	ServerInstanceKey *string `mandatory:"false" contributesTo:"query" name:"serverInstanceKey"`

	// The Fleet-unique identifier of the managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The deployed application key.
	ApplicationKey *string `mandatory:"false" contributesTo:"query" name:"applicationKey"`

	// The library key.
	LibraryKey *string `mandatory:"false" contributesTo:"query" name:"libraryKey"`

	// Filter the list with the Java Server instance name contains the given value.
	ServerInstanceNameContains *string `mandatory:"false" contributesTo:"query" name:"serverInstanceNameContains"`

	// The Java Server instance name.
	ServerInstanceName *string `mandatory:"false" contributesTo:"query" name:"serverInstanceName"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder SummarizeJavaServerInstanceUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the Java Server instances. Only one sort order can be provided.
	// If no value is specified _timeLastSeen_ is default.
	SortBy SummarizeJavaServerInstanceUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeJavaServerInstanceUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeJavaServerInstanceUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeJavaServerInstanceUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeJavaServerInstanceUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeJavaServerInstanceUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeJavaServerInstanceUsageSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeJavaServerInstanceUsageSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeJavaServerInstanceUsageSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeJavaServerInstanceUsageSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeJavaServerInstanceUsageResponse wrapper for the SummarizeJavaServerInstanceUsage operation
type SummarizeJavaServerInstanceUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JavaServerInstanceUsageCollection instances
	JavaServerInstanceUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeJavaServerInstanceUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeJavaServerInstanceUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeJavaServerInstanceUsageSortOrderEnum Enum with underlying type: string
type SummarizeJavaServerInstanceUsageSortOrderEnum string

// Set of constants representing the allowable values for SummarizeJavaServerInstanceUsageSortOrderEnum
const (
	SummarizeJavaServerInstanceUsageSortOrderAsc  SummarizeJavaServerInstanceUsageSortOrderEnum = "ASC"
	SummarizeJavaServerInstanceUsageSortOrderDesc SummarizeJavaServerInstanceUsageSortOrderEnum = "DESC"
)

var mappingSummarizeJavaServerInstanceUsageSortOrderEnum = map[string]SummarizeJavaServerInstanceUsageSortOrderEnum{
	"ASC":  SummarizeJavaServerInstanceUsageSortOrderAsc,
	"DESC": SummarizeJavaServerInstanceUsageSortOrderDesc,
}

var mappingSummarizeJavaServerInstanceUsageSortOrderEnumLowerCase = map[string]SummarizeJavaServerInstanceUsageSortOrderEnum{
	"asc":  SummarizeJavaServerInstanceUsageSortOrderAsc,
	"desc": SummarizeJavaServerInstanceUsageSortOrderDesc,
}

// GetSummarizeJavaServerInstanceUsageSortOrderEnumValues Enumerates the set of values for SummarizeJavaServerInstanceUsageSortOrderEnum
func GetSummarizeJavaServerInstanceUsageSortOrderEnumValues() []SummarizeJavaServerInstanceUsageSortOrderEnum {
	values := make([]SummarizeJavaServerInstanceUsageSortOrderEnum, 0)
	for _, v := range mappingSummarizeJavaServerInstanceUsageSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeJavaServerInstanceUsageSortOrderEnumStringValues Enumerates the set of values in String for SummarizeJavaServerInstanceUsageSortOrderEnum
func GetSummarizeJavaServerInstanceUsageSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeJavaServerInstanceUsageSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeJavaServerInstanceUsageSortOrderEnum(val string) (SummarizeJavaServerInstanceUsageSortOrderEnum, bool) {
	enum, ok := mappingSummarizeJavaServerInstanceUsageSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeJavaServerInstanceUsageSortByEnum Enum with underlying type: string
type SummarizeJavaServerInstanceUsageSortByEnum string

// Set of constants representing the allowable values for SummarizeJavaServerInstanceUsageSortByEnum
const (
	SummarizeJavaServerInstanceUsageSortByServerinstancename                  SummarizeJavaServerInstanceUsageSortByEnum = "serverInstanceName"
	SummarizeJavaServerInstanceUsageSortByManagedinstancename                 SummarizeJavaServerInstanceUsageSortByEnum = "managedInstanceName"
	SummarizeJavaServerInstanceUsageSortByApproximatedeployedapplicationcount SummarizeJavaServerInstanceUsageSortByEnum = "approximateDeployedApplicationCount"
	SummarizeJavaServerInstanceUsageSortByTimefirstseen                       SummarizeJavaServerInstanceUsageSortByEnum = "timeFirstSeen"
	SummarizeJavaServerInstanceUsageSortByTimelastseen                        SummarizeJavaServerInstanceUsageSortByEnum = "timeLastSeen"
)

var mappingSummarizeJavaServerInstanceUsageSortByEnum = map[string]SummarizeJavaServerInstanceUsageSortByEnum{
	"serverInstanceName":                  SummarizeJavaServerInstanceUsageSortByServerinstancename,
	"managedInstanceName":                 SummarizeJavaServerInstanceUsageSortByManagedinstancename,
	"approximateDeployedApplicationCount": SummarizeJavaServerInstanceUsageSortByApproximatedeployedapplicationcount,
	"timeFirstSeen":                       SummarizeJavaServerInstanceUsageSortByTimefirstseen,
	"timeLastSeen":                        SummarizeJavaServerInstanceUsageSortByTimelastseen,
}

var mappingSummarizeJavaServerInstanceUsageSortByEnumLowerCase = map[string]SummarizeJavaServerInstanceUsageSortByEnum{
	"serverinstancename":                  SummarizeJavaServerInstanceUsageSortByServerinstancename,
	"managedinstancename":                 SummarizeJavaServerInstanceUsageSortByManagedinstancename,
	"approximatedeployedapplicationcount": SummarizeJavaServerInstanceUsageSortByApproximatedeployedapplicationcount,
	"timefirstseen":                       SummarizeJavaServerInstanceUsageSortByTimefirstseen,
	"timelastseen":                        SummarizeJavaServerInstanceUsageSortByTimelastseen,
}

// GetSummarizeJavaServerInstanceUsageSortByEnumValues Enumerates the set of values for SummarizeJavaServerInstanceUsageSortByEnum
func GetSummarizeJavaServerInstanceUsageSortByEnumValues() []SummarizeJavaServerInstanceUsageSortByEnum {
	values := make([]SummarizeJavaServerInstanceUsageSortByEnum, 0)
	for _, v := range mappingSummarizeJavaServerInstanceUsageSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeJavaServerInstanceUsageSortByEnumStringValues Enumerates the set of values in String for SummarizeJavaServerInstanceUsageSortByEnum
func GetSummarizeJavaServerInstanceUsageSortByEnumStringValues() []string {
	return []string{
		"serverInstanceName",
		"managedInstanceName",
		"approximateDeployedApplicationCount",
		"timeFirstSeen",
		"timeLastSeen",
	}
}

// GetMappingSummarizeJavaServerInstanceUsageSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeJavaServerInstanceUsageSortByEnum(val string) (SummarizeJavaServerInstanceUsageSortByEnum, bool) {
	enum, ok := mappingSummarizeJavaServerInstanceUsageSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
