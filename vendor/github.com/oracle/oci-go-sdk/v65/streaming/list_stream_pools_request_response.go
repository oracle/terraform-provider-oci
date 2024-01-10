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

// ListStreamPoolsRequest wrapper for the ListStreamPools operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/streaming/ListStreamPools.go.html to see an example of how to use ListStreamPoolsRequest.
type ListStreamPoolsRequest struct {

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
	SortBy ListStreamPoolsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListStreamPoolsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState StreamPoolSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStreamPoolsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStreamPoolsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListStreamPoolsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStreamPoolsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListStreamPoolsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListStreamPoolsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListStreamPoolsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStreamPoolsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListStreamPoolsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingStreamPoolSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetStreamPoolSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListStreamPoolsResponse wrapper for the ListStreamPools operation
type ListStreamPoolsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []StreamPoolSummary instances
	Items []StreamPoolSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results exist. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListStreamPoolsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStreamPoolsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStreamPoolsSortByEnum Enum with underlying type: string
type ListStreamPoolsSortByEnum string

// Set of constants representing the allowable values for ListStreamPoolsSortByEnum
const (
	ListStreamPoolsSortByName        ListStreamPoolsSortByEnum = "NAME"
	ListStreamPoolsSortByTimecreated ListStreamPoolsSortByEnum = "TIMECREATED"
)

var mappingListStreamPoolsSortByEnum = map[string]ListStreamPoolsSortByEnum{
	"NAME":        ListStreamPoolsSortByName,
	"TIMECREATED": ListStreamPoolsSortByTimecreated,
}

var mappingListStreamPoolsSortByEnumLowerCase = map[string]ListStreamPoolsSortByEnum{
	"name":        ListStreamPoolsSortByName,
	"timecreated": ListStreamPoolsSortByTimecreated,
}

// GetListStreamPoolsSortByEnumValues Enumerates the set of values for ListStreamPoolsSortByEnum
func GetListStreamPoolsSortByEnumValues() []ListStreamPoolsSortByEnum {
	values := make([]ListStreamPoolsSortByEnum, 0)
	for _, v := range mappingListStreamPoolsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListStreamPoolsSortByEnumStringValues Enumerates the set of values in String for ListStreamPoolsSortByEnum
func GetListStreamPoolsSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIMECREATED",
	}
}

// GetMappingListStreamPoolsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStreamPoolsSortByEnum(val string) (ListStreamPoolsSortByEnum, bool) {
	enum, ok := mappingListStreamPoolsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStreamPoolsSortOrderEnum Enum with underlying type: string
type ListStreamPoolsSortOrderEnum string

// Set of constants representing the allowable values for ListStreamPoolsSortOrderEnum
const (
	ListStreamPoolsSortOrderAsc  ListStreamPoolsSortOrderEnum = "ASC"
	ListStreamPoolsSortOrderDesc ListStreamPoolsSortOrderEnum = "DESC"
)

var mappingListStreamPoolsSortOrderEnum = map[string]ListStreamPoolsSortOrderEnum{
	"ASC":  ListStreamPoolsSortOrderAsc,
	"DESC": ListStreamPoolsSortOrderDesc,
}

var mappingListStreamPoolsSortOrderEnumLowerCase = map[string]ListStreamPoolsSortOrderEnum{
	"asc":  ListStreamPoolsSortOrderAsc,
	"desc": ListStreamPoolsSortOrderDesc,
}

// GetListStreamPoolsSortOrderEnumValues Enumerates the set of values for ListStreamPoolsSortOrderEnum
func GetListStreamPoolsSortOrderEnumValues() []ListStreamPoolsSortOrderEnum {
	values := make([]ListStreamPoolsSortOrderEnum, 0)
	for _, v := range mappingListStreamPoolsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListStreamPoolsSortOrderEnumStringValues Enumerates the set of values in String for ListStreamPoolsSortOrderEnum
func GetListStreamPoolsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListStreamPoolsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStreamPoolsSortOrderEnum(val string) (ListStreamPoolsSortOrderEnum, bool) {
	enum, ok := mappingListStreamPoolsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
