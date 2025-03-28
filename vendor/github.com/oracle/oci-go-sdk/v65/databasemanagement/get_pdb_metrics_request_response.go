// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetPdbMetricsRequest wrapper for the GetPdbMetrics operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetPdbMetrics.go.html to see an example of how to use GetPdbMetricsRequest.
type GetPdbMetricsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The start time of the time range to retrieve the health metrics of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	StartTime *string `mandatory:"true" contributesTo:"query" name:"startTime"`

	// The end time of the time range to retrieve the health metrics of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	EndTime *string `mandatory:"true" contributesTo:"query" name:"endTime"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The time window used for metrics comparison.
	CompareType GetPdbMetricsCompareTypeEnum `mandatory:"false" contributesTo:"query" name:"compareType" omitEmpty:"true"`

	// The filter used to retrieve a specific set of metrics by passing the desired metric names with a comma separator. Note that, by default, the service returns all supported metrics.
	FilterByMetricNames *string `mandatory:"false" contributesTo:"query" name:"filterByMetricNames"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetPdbMetricsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetPdbMetricsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetPdbMetricsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetPdbMetricsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetPdbMetricsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetPdbMetricsCompareTypeEnum(string(request.CompareType)); !ok && request.CompareType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompareType: %s. Supported values are: %s.", request.CompareType, strings.Join(GetGetPdbMetricsCompareTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetPdbMetricsResponse wrapper for the GetPdbMetrics operation
type GetPdbMetricsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The PdbMetrics instance
	PdbMetrics `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetPdbMetricsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetPdbMetricsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetPdbMetricsCompareTypeEnum Enum with underlying type: string
type GetPdbMetricsCompareTypeEnum string

// Set of constants representing the allowable values for GetPdbMetricsCompareTypeEnum
const (
	GetPdbMetricsCompareTypeHour GetPdbMetricsCompareTypeEnum = "HOUR"
	GetPdbMetricsCompareTypeDay  GetPdbMetricsCompareTypeEnum = "DAY"
	GetPdbMetricsCompareTypeWeek GetPdbMetricsCompareTypeEnum = "WEEK"
)

var mappingGetPdbMetricsCompareTypeEnum = map[string]GetPdbMetricsCompareTypeEnum{
	"HOUR": GetPdbMetricsCompareTypeHour,
	"DAY":  GetPdbMetricsCompareTypeDay,
	"WEEK": GetPdbMetricsCompareTypeWeek,
}

var mappingGetPdbMetricsCompareTypeEnumLowerCase = map[string]GetPdbMetricsCompareTypeEnum{
	"hour": GetPdbMetricsCompareTypeHour,
	"day":  GetPdbMetricsCompareTypeDay,
	"week": GetPdbMetricsCompareTypeWeek,
}

// GetGetPdbMetricsCompareTypeEnumValues Enumerates the set of values for GetPdbMetricsCompareTypeEnum
func GetGetPdbMetricsCompareTypeEnumValues() []GetPdbMetricsCompareTypeEnum {
	values := make([]GetPdbMetricsCompareTypeEnum, 0)
	for _, v := range mappingGetPdbMetricsCompareTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetPdbMetricsCompareTypeEnumStringValues Enumerates the set of values in String for GetPdbMetricsCompareTypeEnum
func GetGetPdbMetricsCompareTypeEnumStringValues() []string {
	return []string{
		"HOUR",
		"DAY",
		"WEEK",
	}
}

// GetMappingGetPdbMetricsCompareTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetPdbMetricsCompareTypeEnum(val string) (GetPdbMetricsCompareTypeEnum, bool) {
	enum, ok := mappingGetPdbMetricsCompareTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
