// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package healthchecks

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListPingProbeResultsRequest wrapper for the ListPingProbeResults operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/ListPingProbeResults.go.html to see an example of how to use ListPingProbeResultsRequest.
type ListPingProbeResultsRequest struct {

	// The OCID of a monitor or on-demand probe.
	ProbeConfigurationId *string `mandatory:"true" contributesTo:"path" name:"probeConfigurationId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Returns results with a `startTime` equal to or greater than the specified value.
	StartTimeGreaterThanOrEqualTo *float64 `mandatory:"false" contributesTo:"query" name:"startTimeGreaterThanOrEqualTo"`

	// Returns results with a `startTime` equal to or less than the specified value.
	StartTimeLessThanOrEqualTo *float64 `mandatory:"false" contributesTo:"query" name:"startTimeLessThanOrEqualTo"`

	// Controls the sort order of results.
	SortOrder ListPingProbeResultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filters results that match the `target`.
	Target *string `mandatory:"false" contributesTo:"query" name:"target"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPingProbeResultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPingProbeResultsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPingProbeResultsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPingProbeResultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPingProbeResultsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPingProbeResultsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPingProbeResultsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPingProbeResultsResponse wrapper for the ListPingProbeResults operation
type ListPingProbeResultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []PingProbeResultSummary instances
	Items []PingProbeResultSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list,
	// if this header appears in the response, then there may be
	// additional items still to get. Include this value as the `page`
	// parameter for the subsequent GET request. For information about
	// pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPingProbeResultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPingProbeResultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPingProbeResultsSortOrderEnum Enum with underlying type: string
type ListPingProbeResultsSortOrderEnum string

// Set of constants representing the allowable values for ListPingProbeResultsSortOrderEnum
const (
	ListPingProbeResultsSortOrderAsc  ListPingProbeResultsSortOrderEnum = "ASC"
	ListPingProbeResultsSortOrderDesc ListPingProbeResultsSortOrderEnum = "DESC"
)

var mappingListPingProbeResultsSortOrderEnum = map[string]ListPingProbeResultsSortOrderEnum{
	"ASC":  ListPingProbeResultsSortOrderAsc,
	"DESC": ListPingProbeResultsSortOrderDesc,
}

// GetListPingProbeResultsSortOrderEnumValues Enumerates the set of values for ListPingProbeResultsSortOrderEnum
func GetListPingProbeResultsSortOrderEnumValues() []ListPingProbeResultsSortOrderEnum {
	values := make([]ListPingProbeResultsSortOrderEnum, 0)
	for _, v := range mappingListPingProbeResultsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPingProbeResultsSortOrderEnumStringValues Enumerates the set of values in String for ListPingProbeResultsSortOrderEnum
func GetListPingProbeResultsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPingProbeResultsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPingProbeResultsSortOrderEnum(val string) (ListPingProbeResultsSortOrderEnum, bool) {
	mappingListPingProbeResultsSortOrderEnumIgnoreCase := make(map[string]ListPingProbeResultsSortOrderEnum)
	for k, v := range mappingListPingProbeResultsSortOrderEnum {
		mappingListPingProbeResultsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListPingProbeResultsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
