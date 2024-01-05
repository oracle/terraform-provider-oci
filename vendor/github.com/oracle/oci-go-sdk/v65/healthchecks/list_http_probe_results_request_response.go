// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package healthchecks

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListHttpProbeResultsRequest wrapper for the ListHttpProbeResults operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/ListHttpProbeResults.go.html to see an example of how to use ListHttpProbeResultsRequest.
type ListHttpProbeResultsRequest struct {

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
	SortOrder ListHttpProbeResultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filters results that match the `target`.
	Target *string `mandatory:"false" contributesTo:"query" name:"target"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListHttpProbeResultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListHttpProbeResultsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListHttpProbeResultsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListHttpProbeResultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListHttpProbeResultsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListHttpProbeResultsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListHttpProbeResultsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListHttpProbeResultsResponse wrapper for the ListHttpProbeResults operation
type ListHttpProbeResultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []HttpProbeResultSummary instances
	Items []HttpProbeResultSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListHttpProbeResultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListHttpProbeResultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListHttpProbeResultsSortOrderEnum Enum with underlying type: string
type ListHttpProbeResultsSortOrderEnum string

// Set of constants representing the allowable values for ListHttpProbeResultsSortOrderEnum
const (
	ListHttpProbeResultsSortOrderAsc  ListHttpProbeResultsSortOrderEnum = "ASC"
	ListHttpProbeResultsSortOrderDesc ListHttpProbeResultsSortOrderEnum = "DESC"
)

var mappingListHttpProbeResultsSortOrderEnum = map[string]ListHttpProbeResultsSortOrderEnum{
	"ASC":  ListHttpProbeResultsSortOrderAsc,
	"DESC": ListHttpProbeResultsSortOrderDesc,
}

var mappingListHttpProbeResultsSortOrderEnumLowerCase = map[string]ListHttpProbeResultsSortOrderEnum{
	"asc":  ListHttpProbeResultsSortOrderAsc,
	"desc": ListHttpProbeResultsSortOrderDesc,
}

// GetListHttpProbeResultsSortOrderEnumValues Enumerates the set of values for ListHttpProbeResultsSortOrderEnum
func GetListHttpProbeResultsSortOrderEnumValues() []ListHttpProbeResultsSortOrderEnum {
	values := make([]ListHttpProbeResultsSortOrderEnum, 0)
	for _, v := range mappingListHttpProbeResultsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListHttpProbeResultsSortOrderEnumStringValues Enumerates the set of values in String for ListHttpProbeResultsSortOrderEnum
func GetListHttpProbeResultsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListHttpProbeResultsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHttpProbeResultsSortOrderEnum(val string) (ListHttpProbeResultsSortOrderEnum, bool) {
	enum, ok := mappingListHttpProbeResultsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
