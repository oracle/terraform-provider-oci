// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeSchedulerJobCountsRequest wrapper for the SummarizeSchedulerJobCounts operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/SummarizeSchedulerJobCounts.go.html to see an example of how to use SummarizeSchedulerJobCountsRequest.
type SummarizeSchedulerJobCountsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder SummarizeSchedulerJobCountsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeSchedulerJobCountsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeSchedulerJobCountsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeSchedulerJobCountsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeSchedulerJobCountsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeSchedulerJobCountsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeSchedulerJobCountsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeSchedulerJobCountsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeSchedulerJobCountsResponse wrapper for the SummarizeSchedulerJobCounts operation
type SummarizeSchedulerJobCountsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SchedulerJobAggregationCollection instances
	SchedulerJobAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeSchedulerJobCountsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeSchedulerJobCountsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeSchedulerJobCountsSortOrderEnum Enum with underlying type: string
type SummarizeSchedulerJobCountsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeSchedulerJobCountsSortOrderEnum
const (
	SummarizeSchedulerJobCountsSortOrderAsc  SummarizeSchedulerJobCountsSortOrderEnum = "ASC"
	SummarizeSchedulerJobCountsSortOrderDesc SummarizeSchedulerJobCountsSortOrderEnum = "DESC"
)

var mappingSummarizeSchedulerJobCountsSortOrderEnum = map[string]SummarizeSchedulerJobCountsSortOrderEnum{
	"ASC":  SummarizeSchedulerJobCountsSortOrderAsc,
	"DESC": SummarizeSchedulerJobCountsSortOrderDesc,
}

var mappingSummarizeSchedulerJobCountsSortOrderEnumLowerCase = map[string]SummarizeSchedulerJobCountsSortOrderEnum{
	"asc":  SummarizeSchedulerJobCountsSortOrderAsc,
	"desc": SummarizeSchedulerJobCountsSortOrderDesc,
}

// GetSummarizeSchedulerJobCountsSortOrderEnumValues Enumerates the set of values for SummarizeSchedulerJobCountsSortOrderEnum
func GetSummarizeSchedulerJobCountsSortOrderEnumValues() []SummarizeSchedulerJobCountsSortOrderEnum {
	values := make([]SummarizeSchedulerJobCountsSortOrderEnum, 0)
	for _, v := range mappingSummarizeSchedulerJobCountsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeSchedulerJobCountsSortOrderEnumStringValues Enumerates the set of values in String for SummarizeSchedulerJobCountsSortOrderEnum
func GetSummarizeSchedulerJobCountsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeSchedulerJobCountsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeSchedulerJobCountsSortOrderEnum(val string) (SummarizeSchedulerJobCountsSortOrderEnum, bool) {
	enum, ok := mappingSummarizeSchedulerJobCountsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
