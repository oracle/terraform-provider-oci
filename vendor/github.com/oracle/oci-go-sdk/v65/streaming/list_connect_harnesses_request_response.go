// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package streaming

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListConnectHarnessesRequest wrapper for the ListConnectHarnesses operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/ListConnectHarnesses.go.html to see an example of how to use ListConnectHarnessesRequest.
type ListConnectHarnessesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given ID exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return. The value must be between 1 and 50. The default is 10.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide no more than one sort order. By default, `TIMECREATED` sorts results in descending order and `NAME` sorts results in ascending order.
	SortBy ListConnectHarnessesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListConnectHarnessesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState ConnectHarnessSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConnectHarnessesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConnectHarnessesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConnectHarnessesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConnectHarnessesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListConnectHarnessesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListConnectHarnessesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConnectHarnessesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConnectHarnessesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConnectHarnessesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConnectHarnessSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetConnectHarnessSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListConnectHarnessesResponse wrapper for the ListConnectHarnesses operation
type ListConnectHarnessesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ConnectHarnessSummary instances
	Items []ConnectHarnessSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results exist. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListConnectHarnessesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConnectHarnessesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConnectHarnessesSortByEnum Enum with underlying type: string
type ListConnectHarnessesSortByEnum string

// Set of constants representing the allowable values for ListConnectHarnessesSortByEnum
const (
	ListConnectHarnessesSortByName        ListConnectHarnessesSortByEnum = "NAME"
	ListConnectHarnessesSortByTimecreated ListConnectHarnessesSortByEnum = "TIMECREATED"
)

var mappingListConnectHarnessesSortByEnum = map[string]ListConnectHarnessesSortByEnum{
	"NAME":        ListConnectHarnessesSortByName,
	"TIMECREATED": ListConnectHarnessesSortByTimecreated,
}

var mappingListConnectHarnessesSortByEnumLowerCase = map[string]ListConnectHarnessesSortByEnum{
	"name":        ListConnectHarnessesSortByName,
	"timecreated": ListConnectHarnessesSortByTimecreated,
}

// GetListConnectHarnessesSortByEnumValues Enumerates the set of values for ListConnectHarnessesSortByEnum
func GetListConnectHarnessesSortByEnumValues() []ListConnectHarnessesSortByEnum {
	values := make([]ListConnectHarnessesSortByEnum, 0)
	for _, v := range mappingListConnectHarnessesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectHarnessesSortByEnumStringValues Enumerates the set of values in String for ListConnectHarnessesSortByEnum
func GetListConnectHarnessesSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIMECREATED",
	}
}

// GetMappingListConnectHarnessesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectHarnessesSortByEnum(val string) (ListConnectHarnessesSortByEnum, bool) {
	enum, ok := mappingListConnectHarnessesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConnectHarnessesSortOrderEnum Enum with underlying type: string
type ListConnectHarnessesSortOrderEnum string

// Set of constants representing the allowable values for ListConnectHarnessesSortOrderEnum
const (
	ListConnectHarnessesSortOrderAsc  ListConnectHarnessesSortOrderEnum = "ASC"
	ListConnectHarnessesSortOrderDesc ListConnectHarnessesSortOrderEnum = "DESC"
)

var mappingListConnectHarnessesSortOrderEnum = map[string]ListConnectHarnessesSortOrderEnum{
	"ASC":  ListConnectHarnessesSortOrderAsc,
	"DESC": ListConnectHarnessesSortOrderDesc,
}

var mappingListConnectHarnessesSortOrderEnumLowerCase = map[string]ListConnectHarnessesSortOrderEnum{
	"asc":  ListConnectHarnessesSortOrderAsc,
	"desc": ListConnectHarnessesSortOrderDesc,
}

// GetListConnectHarnessesSortOrderEnumValues Enumerates the set of values for ListConnectHarnessesSortOrderEnum
func GetListConnectHarnessesSortOrderEnumValues() []ListConnectHarnessesSortOrderEnum {
	values := make([]ListConnectHarnessesSortOrderEnum, 0)
	for _, v := range mappingListConnectHarnessesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectHarnessesSortOrderEnumStringValues Enumerates the set of values in String for ListConnectHarnessesSortOrderEnum
func GetListConnectHarnessesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListConnectHarnessesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectHarnessesSortOrderEnum(val string) (ListConnectHarnessesSortOrderEnum, bool) {
	enum, ok := mappingListConnectHarnessesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
