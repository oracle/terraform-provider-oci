// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeManagedEntityCountsRequest wrapper for the SummarizeManagedEntityCounts operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/SummarizeManagedEntityCounts.go.html to see an example of how to use SummarizeManagedEntityCountsRequest.
type SummarizeManagedEntityCountsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder SummarizeManagedEntityCountsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeManagedEntityCountsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeManagedEntityCountsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeManagedEntityCountsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeManagedEntityCountsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeManagedEntityCountsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeManagedEntityCountsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeManagedEntityCountsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeManagedEntityCountsResponse wrapper for the SummarizeManagedEntityCounts operation
type SummarizeManagedEntityCountsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedEntityAggregationCollection instances
	ManagedEntityAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeManagedEntityCountsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeManagedEntityCountsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeManagedEntityCountsSortOrderEnum Enum with underlying type: string
type SummarizeManagedEntityCountsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeManagedEntityCountsSortOrderEnum
const (
	SummarizeManagedEntityCountsSortOrderAsc  SummarizeManagedEntityCountsSortOrderEnum = "ASC"
	SummarizeManagedEntityCountsSortOrderDesc SummarizeManagedEntityCountsSortOrderEnum = "DESC"
)

var mappingSummarizeManagedEntityCountsSortOrderEnum = map[string]SummarizeManagedEntityCountsSortOrderEnum{
	"ASC":  SummarizeManagedEntityCountsSortOrderAsc,
	"DESC": SummarizeManagedEntityCountsSortOrderDesc,
}

var mappingSummarizeManagedEntityCountsSortOrderEnumLowerCase = map[string]SummarizeManagedEntityCountsSortOrderEnum{
	"asc":  SummarizeManagedEntityCountsSortOrderAsc,
	"desc": SummarizeManagedEntityCountsSortOrderDesc,
}

// GetSummarizeManagedEntityCountsSortOrderEnumValues Enumerates the set of values for SummarizeManagedEntityCountsSortOrderEnum
func GetSummarizeManagedEntityCountsSortOrderEnumValues() []SummarizeManagedEntityCountsSortOrderEnum {
	values := make([]SummarizeManagedEntityCountsSortOrderEnum, 0)
	for _, v := range mappingSummarizeManagedEntityCountsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeManagedEntityCountsSortOrderEnumStringValues Enumerates the set of values in String for SummarizeManagedEntityCountsSortOrderEnum
func GetSummarizeManagedEntityCountsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeManagedEntityCountsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeManagedEntityCountsSortOrderEnum(val string) (SummarizeManagedEntityCountsSortOrderEnum, bool) {
	enum, ok := mappingSummarizeManagedEntityCountsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
