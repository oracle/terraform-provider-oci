// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeManagedInstanceAnalyticsRequest wrapper for the SummarizeManagedInstanceAnalytics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/SummarizeManagedInstanceAnalytics.go.html to see an example of how to use SummarizeManagedInstanceAnalyticsRequest.
type SummarizeManagedInstanceAnalyticsRequest struct {

	// A filter to return only metrics whose name matches the given metric names.
	MetricNames []MetricNameEnum `contributesTo:"query" name:"metricNames" omitEmpty:"true" collectionFormat:"multi"`

	// This compartmentId is used to list managed instances within a compartment.
	// Or serve as an additional filter to restrict only managed instances with in certain compartment if other filter presents.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID of the managed instance group for which to list resources.
	ManagedInstanceGroupId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceGroupId"`

	// The OCID of the lifecycle environment.
	LifecycleEnvironmentId *string `mandatory:"false" contributesTo:"query" name:"lifecycleEnvironmentId"`

	// The OCID of the lifecycle stage for which to list resources.
	LifecycleStageId *string `mandatory:"false" contributesTo:"query" name:"lifecycleStageId"`

	// A filter to return only instances whose managed instance status matches the given status.
	Status []ManagedInstanceStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// Filter instances by Location. Used when report target type is compartment or group.
	InstanceLocation SummarizeManagedInstanceAnalyticsInstanceLocationEnum `mandatory:"false" contributesTo:"query" name:"instanceLocation" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Default order for name is ascending.
	SortBy SummarizeManagedInstanceAnalyticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder SummarizeManagedInstanceAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeManagedInstanceAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeManagedInstanceAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeManagedInstanceAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeManagedInstanceAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeManagedInstanceAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.MetricNames {
		if _, ok := GetMappingMetricNameEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricNames: %s. Supported values are: %s.", val, strings.Join(GetMetricNameEnumStringValues(), ",")))
		}
	}

	for _, val := range request.Status {
		if _, ok := GetMappingManagedInstanceStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetManagedInstanceStatusEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSummarizeManagedInstanceAnalyticsInstanceLocationEnum(string(request.InstanceLocation)); !ok && request.InstanceLocation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstanceLocation: %s. Supported values are: %s.", request.InstanceLocation, strings.Join(GetSummarizeManagedInstanceAnalyticsInstanceLocationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeManagedInstanceAnalyticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeManagedInstanceAnalyticsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeManagedInstanceAnalyticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeManagedInstanceAnalyticsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeManagedInstanceAnalyticsResponse wrapper for the SummarizeManagedInstanceAnalytics operation
type SummarizeManagedInstanceAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceAnalyticCollection instances
	ManagedInstanceAnalyticCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeManagedInstanceAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeManagedInstanceAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeManagedInstanceAnalyticsInstanceLocationEnum Enum with underlying type: string
type SummarizeManagedInstanceAnalyticsInstanceLocationEnum string

// Set of constants representing the allowable values for SummarizeManagedInstanceAnalyticsInstanceLocationEnum
const (
	SummarizeManagedInstanceAnalyticsInstanceLocationOnPremise  SummarizeManagedInstanceAnalyticsInstanceLocationEnum = "ON_PREMISE"
	SummarizeManagedInstanceAnalyticsInstanceLocationOciCompute SummarizeManagedInstanceAnalyticsInstanceLocationEnum = "OCI_COMPUTE"
	SummarizeManagedInstanceAnalyticsInstanceLocationAzure      SummarizeManagedInstanceAnalyticsInstanceLocationEnum = "AZURE"
	SummarizeManagedInstanceAnalyticsInstanceLocationEc2        SummarizeManagedInstanceAnalyticsInstanceLocationEnum = "EC2"
)

var mappingSummarizeManagedInstanceAnalyticsInstanceLocationEnum = map[string]SummarizeManagedInstanceAnalyticsInstanceLocationEnum{
	"ON_PREMISE":  SummarizeManagedInstanceAnalyticsInstanceLocationOnPremise,
	"OCI_COMPUTE": SummarizeManagedInstanceAnalyticsInstanceLocationOciCompute,
	"AZURE":       SummarizeManagedInstanceAnalyticsInstanceLocationAzure,
	"EC2":         SummarizeManagedInstanceAnalyticsInstanceLocationEc2,
}

var mappingSummarizeManagedInstanceAnalyticsInstanceLocationEnumLowerCase = map[string]SummarizeManagedInstanceAnalyticsInstanceLocationEnum{
	"on_premise":  SummarizeManagedInstanceAnalyticsInstanceLocationOnPremise,
	"oci_compute": SummarizeManagedInstanceAnalyticsInstanceLocationOciCompute,
	"azure":       SummarizeManagedInstanceAnalyticsInstanceLocationAzure,
	"ec2":         SummarizeManagedInstanceAnalyticsInstanceLocationEc2,
}

// GetSummarizeManagedInstanceAnalyticsInstanceLocationEnumValues Enumerates the set of values for SummarizeManagedInstanceAnalyticsInstanceLocationEnum
func GetSummarizeManagedInstanceAnalyticsInstanceLocationEnumValues() []SummarizeManagedInstanceAnalyticsInstanceLocationEnum {
	values := make([]SummarizeManagedInstanceAnalyticsInstanceLocationEnum, 0)
	for _, v := range mappingSummarizeManagedInstanceAnalyticsInstanceLocationEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeManagedInstanceAnalyticsInstanceLocationEnumStringValues Enumerates the set of values in String for SummarizeManagedInstanceAnalyticsInstanceLocationEnum
func GetSummarizeManagedInstanceAnalyticsInstanceLocationEnumStringValues() []string {
	return []string{
		"ON_PREMISE",
		"OCI_COMPUTE",
		"AZURE",
		"EC2",
	}
}

// GetMappingSummarizeManagedInstanceAnalyticsInstanceLocationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeManagedInstanceAnalyticsInstanceLocationEnum(val string) (SummarizeManagedInstanceAnalyticsInstanceLocationEnum, bool) {
	enum, ok := mappingSummarizeManagedInstanceAnalyticsInstanceLocationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeManagedInstanceAnalyticsSortByEnum Enum with underlying type: string
type SummarizeManagedInstanceAnalyticsSortByEnum string

// Set of constants representing the allowable values for SummarizeManagedInstanceAnalyticsSortByEnum
const (
	SummarizeManagedInstanceAnalyticsSortByName SummarizeManagedInstanceAnalyticsSortByEnum = "name"
)

var mappingSummarizeManagedInstanceAnalyticsSortByEnum = map[string]SummarizeManagedInstanceAnalyticsSortByEnum{
	"name": SummarizeManagedInstanceAnalyticsSortByName,
}

var mappingSummarizeManagedInstanceAnalyticsSortByEnumLowerCase = map[string]SummarizeManagedInstanceAnalyticsSortByEnum{
	"name": SummarizeManagedInstanceAnalyticsSortByName,
}

// GetSummarizeManagedInstanceAnalyticsSortByEnumValues Enumerates the set of values for SummarizeManagedInstanceAnalyticsSortByEnum
func GetSummarizeManagedInstanceAnalyticsSortByEnumValues() []SummarizeManagedInstanceAnalyticsSortByEnum {
	values := make([]SummarizeManagedInstanceAnalyticsSortByEnum, 0)
	for _, v := range mappingSummarizeManagedInstanceAnalyticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeManagedInstanceAnalyticsSortByEnumStringValues Enumerates the set of values in String for SummarizeManagedInstanceAnalyticsSortByEnum
func GetSummarizeManagedInstanceAnalyticsSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingSummarizeManagedInstanceAnalyticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeManagedInstanceAnalyticsSortByEnum(val string) (SummarizeManagedInstanceAnalyticsSortByEnum, bool) {
	enum, ok := mappingSummarizeManagedInstanceAnalyticsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeManagedInstanceAnalyticsSortOrderEnum Enum with underlying type: string
type SummarizeManagedInstanceAnalyticsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeManagedInstanceAnalyticsSortOrderEnum
const (
	SummarizeManagedInstanceAnalyticsSortOrderAsc  SummarizeManagedInstanceAnalyticsSortOrderEnum = "ASC"
	SummarizeManagedInstanceAnalyticsSortOrderDesc SummarizeManagedInstanceAnalyticsSortOrderEnum = "DESC"
)

var mappingSummarizeManagedInstanceAnalyticsSortOrderEnum = map[string]SummarizeManagedInstanceAnalyticsSortOrderEnum{
	"ASC":  SummarizeManagedInstanceAnalyticsSortOrderAsc,
	"DESC": SummarizeManagedInstanceAnalyticsSortOrderDesc,
}

var mappingSummarizeManagedInstanceAnalyticsSortOrderEnumLowerCase = map[string]SummarizeManagedInstanceAnalyticsSortOrderEnum{
	"asc":  SummarizeManagedInstanceAnalyticsSortOrderAsc,
	"desc": SummarizeManagedInstanceAnalyticsSortOrderDesc,
}

// GetSummarizeManagedInstanceAnalyticsSortOrderEnumValues Enumerates the set of values for SummarizeManagedInstanceAnalyticsSortOrderEnum
func GetSummarizeManagedInstanceAnalyticsSortOrderEnumValues() []SummarizeManagedInstanceAnalyticsSortOrderEnum {
	values := make([]SummarizeManagedInstanceAnalyticsSortOrderEnum, 0)
	for _, v := range mappingSummarizeManagedInstanceAnalyticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeManagedInstanceAnalyticsSortOrderEnumStringValues Enumerates the set of values in String for SummarizeManagedInstanceAnalyticsSortOrderEnum
func GetSummarizeManagedInstanceAnalyticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeManagedInstanceAnalyticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeManagedInstanceAnalyticsSortOrderEnum(val string) (SummarizeManagedInstanceAnalyticsSortOrderEnum, bool) {
	enum, ok := mappingSummarizeManagedInstanceAnalyticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
