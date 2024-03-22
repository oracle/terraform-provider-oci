// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetHeatWaveFleetMetricRequest wrapper for the GetHeatWaveFleetMetric operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetHeatWaveFleetMetric.go.html to see an example of how to use GetHeatWaveFleetMetricRequest.
type GetHeatWaveFleetMetricRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The start time of the time range to retrieve the health metrics of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	StartTime *string `mandatory:"true" contributesTo:"query" name:"startTime"`

	// The end time of the time range to retrieve the health metrics of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	EndTime *string `mandatory:"true" contributesTo:"query" name:"endTime"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The filter used to retrieve a specific set of metrics by passing the desired metric names with a comma separator. Note that, by default, the service returns all supported metrics.
	FilterByMetricNames *string `mandatory:"false" contributesTo:"query" name:"filterByMetricNames"`

	// The parameter to filter by HeatWave cluster status.
	FilterByHeatWaveStatus GetHeatWaveFleetMetricFilterByHeatWaveStatusEnum `mandatory:"false" contributesTo:"query" name:"filterByHeatWaveStatus" omitEmpty:"true"`

	// The parameter to filter by HeatWave node shape.
	FilterByHeatWaveShape *string `mandatory:"false" contributesTo:"query" name:"filterByHeatWaveShape"`

	// The parameter to filter based on whether HeatWave Lakehouse is enabled for the cluster.
	IsHeatWaveLakehouseEnabled *bool `mandatory:"false" contributesTo:"query" name:"isHeatWaveLakehouseEnabled"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetHeatWaveFleetMetricRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetHeatWaveFleetMetricRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetHeatWaveFleetMetricRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetHeatWaveFleetMetricRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetHeatWaveFleetMetricRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetHeatWaveFleetMetricFilterByHeatWaveStatusEnum(string(request.FilterByHeatWaveStatus)); !ok && request.FilterByHeatWaveStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FilterByHeatWaveStatus: %s. Supported values are: %s.", request.FilterByHeatWaveStatus, strings.Join(GetGetHeatWaveFleetMetricFilterByHeatWaveStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetHeatWaveFleetMetricResponse wrapper for the GetHeatWaveFleetMetric operation
type GetHeatWaveFleetMetricResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The HeatWaveFleetMetrics instance
	HeatWaveFleetMetrics `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetHeatWaveFleetMetricResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetHeatWaveFleetMetricResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetHeatWaveFleetMetricFilterByHeatWaveStatusEnum Enum with underlying type: string
type GetHeatWaveFleetMetricFilterByHeatWaveStatusEnum string

// Set of constants representing the allowable values for GetHeatWaveFleetMetricFilterByHeatWaveStatusEnum
const (
	GetHeatWaveFleetMetricFilterByHeatWaveStatusUp      GetHeatWaveFleetMetricFilterByHeatWaveStatusEnum = "UP"
	GetHeatWaveFleetMetricFilterByHeatWaveStatusDown    GetHeatWaveFleetMetricFilterByHeatWaveStatusEnum = "DOWN"
	GetHeatWaveFleetMetricFilterByHeatWaveStatusUnknown GetHeatWaveFleetMetricFilterByHeatWaveStatusEnum = "UNKNOWN"
)

var mappingGetHeatWaveFleetMetricFilterByHeatWaveStatusEnum = map[string]GetHeatWaveFleetMetricFilterByHeatWaveStatusEnum{
	"UP":      GetHeatWaveFleetMetricFilterByHeatWaveStatusUp,
	"DOWN":    GetHeatWaveFleetMetricFilterByHeatWaveStatusDown,
	"UNKNOWN": GetHeatWaveFleetMetricFilterByHeatWaveStatusUnknown,
}

var mappingGetHeatWaveFleetMetricFilterByHeatWaveStatusEnumLowerCase = map[string]GetHeatWaveFleetMetricFilterByHeatWaveStatusEnum{
	"up":      GetHeatWaveFleetMetricFilterByHeatWaveStatusUp,
	"down":    GetHeatWaveFleetMetricFilterByHeatWaveStatusDown,
	"unknown": GetHeatWaveFleetMetricFilterByHeatWaveStatusUnknown,
}

// GetGetHeatWaveFleetMetricFilterByHeatWaveStatusEnumValues Enumerates the set of values for GetHeatWaveFleetMetricFilterByHeatWaveStatusEnum
func GetGetHeatWaveFleetMetricFilterByHeatWaveStatusEnumValues() []GetHeatWaveFleetMetricFilterByHeatWaveStatusEnum {
	values := make([]GetHeatWaveFleetMetricFilterByHeatWaveStatusEnum, 0)
	for _, v := range mappingGetHeatWaveFleetMetricFilterByHeatWaveStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetGetHeatWaveFleetMetricFilterByHeatWaveStatusEnumStringValues Enumerates the set of values in String for GetHeatWaveFleetMetricFilterByHeatWaveStatusEnum
func GetGetHeatWaveFleetMetricFilterByHeatWaveStatusEnumStringValues() []string {
	return []string{
		"UP",
		"DOWN",
		"UNKNOWN",
	}
}

// GetMappingGetHeatWaveFleetMetricFilterByHeatWaveStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetHeatWaveFleetMetricFilterByHeatWaveStatusEnum(val string) (GetHeatWaveFleetMetricFilterByHeatWaveStatusEnum, bool) {
	enum, ok := mappingGetHeatWaveFleetMetricFilterByHeatWaveStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
