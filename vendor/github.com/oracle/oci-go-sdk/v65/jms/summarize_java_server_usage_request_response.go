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

// SummarizeJavaServerUsageRequest wrapper for the SummarizeJavaServerUsage operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeJavaServerUsage.go.html to see an example of how to use SummarizeJavaServerUsageRequest.
type SummarizeJavaServerUsageRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The server key.
	ServerKey *string `mandatory:"false" contributesTo:"query" name:"serverKey"`

	// Filter the list with server name contains the given value.
	ServerNameContains *string `mandatory:"false" contributesTo:"query" name:"serverNameContains"`

	// The server name.
	ServerName *string `mandatory:"false" contributesTo:"query" name:"serverName"`

	// The server version.
	ServerVersion *string `mandatory:"false" contributesTo:"query" name:"serverVersion"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder SummarizeJavaServerUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort a Java Server. Only one sort order can be provided.
	// If no value is specified _timeLastSeen_ is default.
	SortBy SummarizeJavaServerUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeJavaServerUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeJavaServerUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeJavaServerUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeJavaServerUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeJavaServerUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeJavaServerUsageSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeJavaServerUsageSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeJavaServerUsageSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeJavaServerUsageSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeJavaServerUsageResponse wrapper for the SummarizeJavaServerUsage operation
type SummarizeJavaServerUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JavaServerUsageCollection instances
	JavaServerUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeJavaServerUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeJavaServerUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeJavaServerUsageSortOrderEnum Enum with underlying type: string
type SummarizeJavaServerUsageSortOrderEnum string

// Set of constants representing the allowable values for SummarizeJavaServerUsageSortOrderEnum
const (
	SummarizeJavaServerUsageSortOrderAsc  SummarizeJavaServerUsageSortOrderEnum = "ASC"
	SummarizeJavaServerUsageSortOrderDesc SummarizeJavaServerUsageSortOrderEnum = "DESC"
)

var mappingSummarizeJavaServerUsageSortOrderEnum = map[string]SummarizeJavaServerUsageSortOrderEnum{
	"ASC":  SummarizeJavaServerUsageSortOrderAsc,
	"DESC": SummarizeJavaServerUsageSortOrderDesc,
}

var mappingSummarizeJavaServerUsageSortOrderEnumLowerCase = map[string]SummarizeJavaServerUsageSortOrderEnum{
	"asc":  SummarizeJavaServerUsageSortOrderAsc,
	"desc": SummarizeJavaServerUsageSortOrderDesc,
}

// GetSummarizeJavaServerUsageSortOrderEnumValues Enumerates the set of values for SummarizeJavaServerUsageSortOrderEnum
func GetSummarizeJavaServerUsageSortOrderEnumValues() []SummarizeJavaServerUsageSortOrderEnum {
	values := make([]SummarizeJavaServerUsageSortOrderEnum, 0)
	for _, v := range mappingSummarizeJavaServerUsageSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeJavaServerUsageSortOrderEnumStringValues Enumerates the set of values in String for SummarizeJavaServerUsageSortOrderEnum
func GetSummarizeJavaServerUsageSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeJavaServerUsageSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeJavaServerUsageSortOrderEnum(val string) (SummarizeJavaServerUsageSortOrderEnum, bool) {
	enum, ok := mappingSummarizeJavaServerUsageSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeJavaServerUsageSortByEnum Enum with underlying type: string
type SummarizeJavaServerUsageSortByEnum string

// Set of constants representing the allowable values for SummarizeJavaServerUsageSortByEnum
const (
	SummarizeJavaServerUsageSortByServername                          SummarizeJavaServerUsageSortByEnum = "serverName"
	SummarizeJavaServerUsageSortByServerversion                       SummarizeJavaServerUsageSortByEnum = "serverVersion"
	SummarizeJavaServerUsageSortByServerinstancecount                 SummarizeJavaServerUsageSortByEnum = "serverInstanceCount"
	SummarizeJavaServerUsageSortByApproximatedeployedapplicationcount SummarizeJavaServerUsageSortByEnum = "approximateDeployedApplicationCount"
	SummarizeJavaServerUsageSortByTimefirstseen                       SummarizeJavaServerUsageSortByEnum = "timeFirstSeen"
	SummarizeJavaServerUsageSortByTimelastseen                        SummarizeJavaServerUsageSortByEnum = "timeLastSeen"
)

var mappingSummarizeJavaServerUsageSortByEnum = map[string]SummarizeJavaServerUsageSortByEnum{
	"serverName":                          SummarizeJavaServerUsageSortByServername,
	"serverVersion":                       SummarizeJavaServerUsageSortByServerversion,
	"serverInstanceCount":                 SummarizeJavaServerUsageSortByServerinstancecount,
	"approximateDeployedApplicationCount": SummarizeJavaServerUsageSortByApproximatedeployedapplicationcount,
	"timeFirstSeen":                       SummarizeJavaServerUsageSortByTimefirstseen,
	"timeLastSeen":                        SummarizeJavaServerUsageSortByTimelastseen,
}

var mappingSummarizeJavaServerUsageSortByEnumLowerCase = map[string]SummarizeJavaServerUsageSortByEnum{
	"servername":                          SummarizeJavaServerUsageSortByServername,
	"serverversion":                       SummarizeJavaServerUsageSortByServerversion,
	"serverinstancecount":                 SummarizeJavaServerUsageSortByServerinstancecount,
	"approximatedeployedapplicationcount": SummarizeJavaServerUsageSortByApproximatedeployedapplicationcount,
	"timefirstseen":                       SummarizeJavaServerUsageSortByTimefirstseen,
	"timelastseen":                        SummarizeJavaServerUsageSortByTimelastseen,
}

// GetSummarizeJavaServerUsageSortByEnumValues Enumerates the set of values for SummarizeJavaServerUsageSortByEnum
func GetSummarizeJavaServerUsageSortByEnumValues() []SummarizeJavaServerUsageSortByEnum {
	values := make([]SummarizeJavaServerUsageSortByEnum, 0)
	for _, v := range mappingSummarizeJavaServerUsageSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeJavaServerUsageSortByEnumStringValues Enumerates the set of values in String for SummarizeJavaServerUsageSortByEnum
func GetSummarizeJavaServerUsageSortByEnumStringValues() []string {
	return []string{
		"serverName",
		"serverVersion",
		"serverInstanceCount",
		"approximateDeployedApplicationCount",
		"timeFirstSeen",
		"timeLastSeen",
	}
}

// GetMappingSummarizeJavaServerUsageSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeJavaServerUsageSortByEnum(val string) (SummarizeJavaServerUsageSortByEnum, bool) {
	enum, ok := mappingSummarizeJavaServerUsageSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
