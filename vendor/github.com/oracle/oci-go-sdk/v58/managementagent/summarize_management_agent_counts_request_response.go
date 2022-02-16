// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// SummarizeManagementAgentCountsRequest wrapper for the SummarizeManagementAgentCounts operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementagent/SummarizeManagementAgentCounts.go.html to see an example of how to use SummarizeManagementAgentCountsRequest.
type SummarizeManagementAgentCountsRequest struct {

	// The OCID of the compartment to which a request will be scoped.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The field by which to group Management Agents. Currently, only one groupBy dimension is supported at a time.
	GroupBy []ManagementAgentGroupByEnum `contributesTo:"query" name:"groupBy" omitEmpty:"true" collectionFormat:"multi"`

	// When set to true then agents that have at least one plugin deployed will be returned. When set to false only agents that have no plugins deployed will be returned.
	HasPlugins *bool `mandatory:"false" contributesTo:"query" name:"hasPlugins"`

	// A filter to return either agents or gateway types depending upon install type selected by user. By default both install type will be returned.
	InstallType SummarizeManagementAgentCountsInstallTypeEnum `mandatory:"false" contributesTo:"query" name:"installType" omitEmpty:"true"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeManagementAgentCountsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeManagementAgentCountsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeManagementAgentCountsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeManagementAgentCountsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeManagementAgentCountsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.GroupBy {
		if _, ok := GetMappingManagementAgentGroupByEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", val, strings.Join(GetManagementAgentGroupByEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSummarizeManagementAgentCountsInstallTypeEnum(string(request.InstallType)); !ok && request.InstallType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstallType: %s. Supported values are: %s.", request.InstallType, strings.Join(GetSummarizeManagementAgentCountsInstallTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeManagementAgentCountsResponse wrapper for the SummarizeManagementAgentCounts operation
type SummarizeManagementAgentCountsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagementAgentAggregationCollection instances
	ManagementAgentAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeManagementAgentCountsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeManagementAgentCountsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeManagementAgentCountsInstallTypeEnum Enum with underlying type: string
type SummarizeManagementAgentCountsInstallTypeEnum string

// Set of constants representing the allowable values for SummarizeManagementAgentCountsInstallTypeEnum
const (
	SummarizeManagementAgentCountsInstallTypeAgent   SummarizeManagementAgentCountsInstallTypeEnum = "AGENT"
	SummarizeManagementAgentCountsInstallTypeGateway SummarizeManagementAgentCountsInstallTypeEnum = "GATEWAY"
)

var mappingSummarizeManagementAgentCountsInstallTypeEnum = map[string]SummarizeManagementAgentCountsInstallTypeEnum{
	"AGENT":   SummarizeManagementAgentCountsInstallTypeAgent,
	"GATEWAY": SummarizeManagementAgentCountsInstallTypeGateway,
}

// GetSummarizeManagementAgentCountsInstallTypeEnumValues Enumerates the set of values for SummarizeManagementAgentCountsInstallTypeEnum
func GetSummarizeManagementAgentCountsInstallTypeEnumValues() []SummarizeManagementAgentCountsInstallTypeEnum {
	values := make([]SummarizeManagementAgentCountsInstallTypeEnum, 0)
	for _, v := range mappingSummarizeManagementAgentCountsInstallTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeManagementAgentCountsInstallTypeEnumStringValues Enumerates the set of values in String for SummarizeManagementAgentCountsInstallTypeEnum
func GetSummarizeManagementAgentCountsInstallTypeEnumStringValues() []string {
	return []string{
		"AGENT",
		"GATEWAY",
	}
}

// GetMappingSummarizeManagementAgentCountsInstallTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeManagementAgentCountsInstallTypeEnum(val string) (SummarizeManagementAgentCountsInstallTypeEnum, bool) {
	mappingSummarizeManagementAgentCountsInstallTypeEnumIgnoreCase := make(map[string]SummarizeManagementAgentCountsInstallTypeEnum)
	for k, v := range mappingSummarizeManagementAgentCountsInstallTypeEnum {
		mappingSummarizeManagementAgentCountsInstallTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeManagementAgentCountsInstallTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
