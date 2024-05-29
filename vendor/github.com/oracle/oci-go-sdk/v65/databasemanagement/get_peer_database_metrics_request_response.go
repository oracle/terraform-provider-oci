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

// GetPeerDatabaseMetricsRequest wrapper for the GetPeerDatabaseMetrics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetPeerDatabaseMetrics.go.html to see an example of how to use GetPeerDatabaseMetricsRequest.
type GetPeerDatabaseMetricsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The start time of the time range to retrieve the health metrics of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	StartTime *string `mandatory:"true" contributesTo:"query" name:"startTime"`

	// The end time of the time range to retrieve the health metrics of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	EndTime *string `mandatory:"true" contributesTo:"query" name:"endTime"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment for which peer database metrics are required.
	// This is not a mandatory parameter and in its absence, all the peer database metrics are returned.
	PeerDatabaseCompartmentId *string `mandatory:"false" contributesTo:"query" name:"peerDatabaseCompartmentId"`

	// The time window used for metrics comparison.
	CompareType GetPeerDatabaseMetricsCompareTypeEnum `mandatory:"false" contributesTo:"query" name:"compareType" omitEmpty:"true"`

	// The filter used to retrieve a specific set of metrics by passing the desired metric names with a comma separator. Note that, by default, the service returns all supported metrics.
	FilterByMetricNames *string `mandatory:"false" contributesTo:"query" name:"filterByMetricNames"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetPeerDatabaseMetricsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetPeerDatabaseMetricsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetPeerDatabaseMetricsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetPeerDatabaseMetricsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetPeerDatabaseMetricsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetPeerDatabaseMetricsCompareTypeEnum(string(request.CompareType)); !ok && request.CompareType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompareType: %s. Supported values are: %s.", request.CompareType, strings.Join(GetGetPeerDatabaseMetricsCompareTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetPeerDatabaseMetricsResponse wrapper for the GetPeerDatabaseMetrics operation
type GetPeerDatabaseMetricsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The PeerDatabaseMetrics instance
	PeerDatabaseMetrics `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetPeerDatabaseMetricsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetPeerDatabaseMetricsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetPeerDatabaseMetricsCompareTypeEnum Enum with underlying type: string
type GetPeerDatabaseMetricsCompareTypeEnum string

// Set of constants representing the allowable values for GetPeerDatabaseMetricsCompareTypeEnum
const (
	GetPeerDatabaseMetricsCompareTypeHour GetPeerDatabaseMetricsCompareTypeEnum = "HOUR"
	GetPeerDatabaseMetricsCompareTypeDay  GetPeerDatabaseMetricsCompareTypeEnum = "DAY"
	GetPeerDatabaseMetricsCompareTypeWeek GetPeerDatabaseMetricsCompareTypeEnum = "WEEK"
)

var mappingGetPeerDatabaseMetricsCompareTypeEnum = map[string]GetPeerDatabaseMetricsCompareTypeEnum{
	"HOUR": GetPeerDatabaseMetricsCompareTypeHour,
	"DAY":  GetPeerDatabaseMetricsCompareTypeDay,
	"WEEK": GetPeerDatabaseMetricsCompareTypeWeek,
}

var mappingGetPeerDatabaseMetricsCompareTypeEnumLowerCase = map[string]GetPeerDatabaseMetricsCompareTypeEnum{
	"hour": GetPeerDatabaseMetricsCompareTypeHour,
	"day":  GetPeerDatabaseMetricsCompareTypeDay,
	"week": GetPeerDatabaseMetricsCompareTypeWeek,
}

// GetGetPeerDatabaseMetricsCompareTypeEnumValues Enumerates the set of values for GetPeerDatabaseMetricsCompareTypeEnum
func GetGetPeerDatabaseMetricsCompareTypeEnumValues() []GetPeerDatabaseMetricsCompareTypeEnum {
	values := make([]GetPeerDatabaseMetricsCompareTypeEnum, 0)
	for _, v := range mappingGetPeerDatabaseMetricsCompareTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetPeerDatabaseMetricsCompareTypeEnumStringValues Enumerates the set of values in String for GetPeerDatabaseMetricsCompareTypeEnum
func GetGetPeerDatabaseMetricsCompareTypeEnumStringValues() []string {
	return []string{
		"HOUR",
		"DAY",
		"WEEK",
	}
}

// GetMappingGetPeerDatabaseMetricsCompareTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetPeerDatabaseMetricsCompareTypeEnum(val string) (GetPeerDatabaseMetricsCompareTypeEnum, bool) {
	enum, ok := mappingGetPeerDatabaseMetricsCompareTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
