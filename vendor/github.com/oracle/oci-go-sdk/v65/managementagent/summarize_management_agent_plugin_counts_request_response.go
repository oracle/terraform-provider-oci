// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeManagementAgentPluginCountsRequest wrapper for the SummarizeManagementAgentPluginCounts operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementagent/SummarizeManagementAgentPluginCounts.go.html to see an example of how to use SummarizeManagementAgentPluginCountsRequest.
type SummarizeManagementAgentPluginCountsRequest struct {

	// The OCID of the compartment to which a request will be scoped.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The field by which to group Management Agent Plugins
	GroupBy SummarizeManagementAgentPluginCountsGroupByEnum `mandatory:"true" contributesTo:"query" name:"groupBy" omitEmpty:"true"`

	// if set to true then it fetches resources for all compartments where user has access to else only on the compartment specified.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeManagementAgentPluginCountsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeManagementAgentPluginCountsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeManagementAgentPluginCountsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeManagementAgentPluginCountsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeManagementAgentPluginCountsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeManagementAgentPluginCountsGroupByEnum(string(request.GroupBy)); !ok && request.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", request.GroupBy, strings.Join(GetSummarizeManagementAgentPluginCountsGroupByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeManagementAgentPluginCountsResponse wrapper for the SummarizeManagementAgentPluginCounts operation
type SummarizeManagementAgentPluginCountsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagementAgentPluginAggregationCollection instances
	ManagementAgentPluginAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeManagementAgentPluginCountsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeManagementAgentPluginCountsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeManagementAgentPluginCountsGroupByEnum Enum with underlying type: string
type SummarizeManagementAgentPluginCountsGroupByEnum string

// Set of constants representing the allowable values for SummarizeManagementAgentPluginCountsGroupByEnum
const (
	SummarizeManagementAgentPluginCountsGroupByPluginname SummarizeManagementAgentPluginCountsGroupByEnum = "pluginName"
)

var mappingSummarizeManagementAgentPluginCountsGroupByEnum = map[string]SummarizeManagementAgentPluginCountsGroupByEnum{
	"pluginName": SummarizeManagementAgentPluginCountsGroupByPluginname,
}

var mappingSummarizeManagementAgentPluginCountsGroupByEnumLowerCase = map[string]SummarizeManagementAgentPluginCountsGroupByEnum{
	"pluginname": SummarizeManagementAgentPluginCountsGroupByPluginname,
}

// GetSummarizeManagementAgentPluginCountsGroupByEnumValues Enumerates the set of values for SummarizeManagementAgentPluginCountsGroupByEnum
func GetSummarizeManagementAgentPluginCountsGroupByEnumValues() []SummarizeManagementAgentPluginCountsGroupByEnum {
	values := make([]SummarizeManagementAgentPluginCountsGroupByEnum, 0)
	for _, v := range mappingSummarizeManagementAgentPluginCountsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeManagementAgentPluginCountsGroupByEnumStringValues Enumerates the set of values in String for SummarizeManagementAgentPluginCountsGroupByEnum
func GetSummarizeManagementAgentPluginCountsGroupByEnumStringValues() []string {
	return []string{
		"pluginName",
	}
}

// GetMappingSummarizeManagementAgentPluginCountsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeManagementAgentPluginCountsGroupByEnum(val string) (SummarizeManagementAgentPluginCountsGroupByEnum, bool) {
	enum, ok := mappingSummarizeManagementAgentPluginCountsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
