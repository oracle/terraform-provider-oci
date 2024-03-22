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

// GetMySqlFleetMetricRequest wrapper for the GetMySqlFleetMetric operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetMySqlFleetMetric.go.html to see an example of how to use GetMySqlFleetMetricRequest.
type GetMySqlFleetMetricRequest struct {

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

	// The parameter to filter by MySQL deployment type.
	FilterByMySqlDeploymentTypeParam GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum `mandatory:"false" contributesTo:"query" name:"filterByMySqlDeploymentTypeParam" omitEmpty:"true"`

	// The parameter to filter by MySQL Database System type.
	FilterByMdsDeploymentType GetMySqlFleetMetricFilterByMdsDeploymentTypeEnum `mandatory:"false" contributesTo:"query" name:"filterByMdsDeploymentType" omitEmpty:"true"`

	// The parameter to filter by MySQL Database status.
	FilterByMySqlStatus GetMySqlFleetMetricFilterByMySqlStatusEnum `mandatory:"false" contributesTo:"query" name:"filterByMySqlStatus" omitEmpty:"true"`

	// The parameter to filter by MySQL database version.
	FilterByMySqlDatabaseVersion *string `mandatory:"false" contributesTo:"query" name:"filterByMySqlDatabaseVersion"`

	// The parameter to filter based on whether HeatWave is enabled for the database.
	IsHeatWaveEnabled *bool `mandatory:"false" contributesTo:"query" name:"isHeatWaveEnabled"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetMySqlFleetMetricRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetMySqlFleetMetricRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetMySqlFleetMetricRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetMySqlFleetMetricRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetMySqlFleetMetricRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum(string(request.FilterByMySqlDeploymentTypeParam)); !ok && request.FilterByMySqlDeploymentTypeParam != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FilterByMySqlDeploymentTypeParam: %s. Supported values are: %s.", request.FilterByMySqlDeploymentTypeParam, strings.Join(GetGetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetMySqlFleetMetricFilterByMdsDeploymentTypeEnum(string(request.FilterByMdsDeploymentType)); !ok && request.FilterByMdsDeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FilterByMdsDeploymentType: %s. Supported values are: %s.", request.FilterByMdsDeploymentType, strings.Join(GetGetMySqlFleetMetricFilterByMdsDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetMySqlFleetMetricFilterByMySqlStatusEnum(string(request.FilterByMySqlStatus)); !ok && request.FilterByMySqlStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FilterByMySqlStatus: %s. Supported values are: %s.", request.FilterByMySqlStatus, strings.Join(GetGetMySqlFleetMetricFilterByMySqlStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetMySqlFleetMetricResponse wrapper for the GetMySqlFleetMetric operation
type GetMySqlFleetMetricResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The MySqlFleetMetrics instance
	MySqlFleetMetrics `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetMySqlFleetMetricResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetMySqlFleetMetricResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum Enum with underlying type: string
type GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum string

// Set of constants representing the allowable values for GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum
const (
	GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamOnpremise GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum = "ONPREMISE"
	GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamMds       GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum = "MDS"
)

var mappingGetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum = map[string]GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum{
	"ONPREMISE": GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamOnpremise,
	"MDS":       GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamMds,
}

var mappingGetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnumLowerCase = map[string]GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum{
	"onpremise": GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamOnpremise,
	"mds":       GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamMds,
}

// GetGetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnumValues Enumerates the set of values for GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum
func GetGetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnumValues() []GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum {
	values := make([]GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum, 0)
	for _, v := range mappingGetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum {
		values = append(values, v)
	}
	return values
}

// GetGetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnumStringValues Enumerates the set of values in String for GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum
func GetGetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnumStringValues() []string {
	return []string{
		"ONPREMISE",
		"MDS",
	}
}

// GetMappingGetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum(val string) (GetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnum, bool) {
	enum, ok := mappingGetMySqlFleetMetricFilterByMySqlDeploymentTypeParamEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetMySqlFleetMetricFilterByMdsDeploymentTypeEnum Enum with underlying type: string
type GetMySqlFleetMetricFilterByMdsDeploymentTypeEnum string

// Set of constants representing the allowable values for GetMySqlFleetMetricFilterByMdsDeploymentTypeEnum
const (
	GetMySqlFleetMetricFilterByMdsDeploymentTypeHa         GetMySqlFleetMetricFilterByMdsDeploymentTypeEnum = "HA"
	GetMySqlFleetMetricFilterByMdsDeploymentTypeHeatwave   GetMySqlFleetMetricFilterByMdsDeploymentTypeEnum = "HEATWAVE"
	GetMySqlFleetMetricFilterByMdsDeploymentTypeStandalone GetMySqlFleetMetricFilterByMdsDeploymentTypeEnum = "STANDALONE"
)

var mappingGetMySqlFleetMetricFilterByMdsDeploymentTypeEnum = map[string]GetMySqlFleetMetricFilterByMdsDeploymentTypeEnum{
	"HA":         GetMySqlFleetMetricFilterByMdsDeploymentTypeHa,
	"HEATWAVE":   GetMySqlFleetMetricFilterByMdsDeploymentTypeHeatwave,
	"STANDALONE": GetMySqlFleetMetricFilterByMdsDeploymentTypeStandalone,
}

var mappingGetMySqlFleetMetricFilterByMdsDeploymentTypeEnumLowerCase = map[string]GetMySqlFleetMetricFilterByMdsDeploymentTypeEnum{
	"ha":         GetMySqlFleetMetricFilterByMdsDeploymentTypeHa,
	"heatwave":   GetMySqlFleetMetricFilterByMdsDeploymentTypeHeatwave,
	"standalone": GetMySqlFleetMetricFilterByMdsDeploymentTypeStandalone,
}

// GetGetMySqlFleetMetricFilterByMdsDeploymentTypeEnumValues Enumerates the set of values for GetMySqlFleetMetricFilterByMdsDeploymentTypeEnum
func GetGetMySqlFleetMetricFilterByMdsDeploymentTypeEnumValues() []GetMySqlFleetMetricFilterByMdsDeploymentTypeEnum {
	values := make([]GetMySqlFleetMetricFilterByMdsDeploymentTypeEnum, 0)
	for _, v := range mappingGetMySqlFleetMetricFilterByMdsDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetMySqlFleetMetricFilterByMdsDeploymentTypeEnumStringValues Enumerates the set of values in String for GetMySqlFleetMetricFilterByMdsDeploymentTypeEnum
func GetGetMySqlFleetMetricFilterByMdsDeploymentTypeEnumStringValues() []string {
	return []string{
		"HA",
		"HEATWAVE",
		"STANDALONE",
	}
}

// GetMappingGetMySqlFleetMetricFilterByMdsDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetMySqlFleetMetricFilterByMdsDeploymentTypeEnum(val string) (GetMySqlFleetMetricFilterByMdsDeploymentTypeEnum, bool) {
	enum, ok := mappingGetMySqlFleetMetricFilterByMdsDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetMySqlFleetMetricFilterByMySqlStatusEnum Enum with underlying type: string
type GetMySqlFleetMetricFilterByMySqlStatusEnum string

// Set of constants representing the allowable values for GetMySqlFleetMetricFilterByMySqlStatusEnum
const (
	GetMySqlFleetMetricFilterByMySqlStatusUp      GetMySqlFleetMetricFilterByMySqlStatusEnum = "UP"
	GetMySqlFleetMetricFilterByMySqlStatusDown    GetMySqlFleetMetricFilterByMySqlStatusEnum = "DOWN"
	GetMySqlFleetMetricFilterByMySqlStatusUnknown GetMySqlFleetMetricFilterByMySqlStatusEnum = "UNKNOWN"
)

var mappingGetMySqlFleetMetricFilterByMySqlStatusEnum = map[string]GetMySqlFleetMetricFilterByMySqlStatusEnum{
	"UP":      GetMySqlFleetMetricFilterByMySqlStatusUp,
	"DOWN":    GetMySqlFleetMetricFilterByMySqlStatusDown,
	"UNKNOWN": GetMySqlFleetMetricFilterByMySqlStatusUnknown,
}

var mappingGetMySqlFleetMetricFilterByMySqlStatusEnumLowerCase = map[string]GetMySqlFleetMetricFilterByMySqlStatusEnum{
	"up":      GetMySqlFleetMetricFilterByMySqlStatusUp,
	"down":    GetMySqlFleetMetricFilterByMySqlStatusDown,
	"unknown": GetMySqlFleetMetricFilterByMySqlStatusUnknown,
}

// GetGetMySqlFleetMetricFilterByMySqlStatusEnumValues Enumerates the set of values for GetMySqlFleetMetricFilterByMySqlStatusEnum
func GetGetMySqlFleetMetricFilterByMySqlStatusEnumValues() []GetMySqlFleetMetricFilterByMySqlStatusEnum {
	values := make([]GetMySqlFleetMetricFilterByMySqlStatusEnum, 0)
	for _, v := range mappingGetMySqlFleetMetricFilterByMySqlStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetGetMySqlFleetMetricFilterByMySqlStatusEnumStringValues Enumerates the set of values in String for GetMySqlFleetMetricFilterByMySqlStatusEnum
func GetGetMySqlFleetMetricFilterByMySqlStatusEnumStringValues() []string {
	return []string{
		"UP",
		"DOWN",
		"UNKNOWN",
	}
}

// GetMappingGetMySqlFleetMetricFilterByMySqlStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetMySqlFleetMetricFilterByMySqlStatusEnum(val string) (GetMySqlFleetMetricFilterByMySqlStatusEnum, bool) {
	enum, ok := mappingGetMySqlFleetMetricFilterByMySqlStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
