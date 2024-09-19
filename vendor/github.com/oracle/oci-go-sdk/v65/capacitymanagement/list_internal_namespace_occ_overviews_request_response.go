// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInternalNamespaceOccOverviewsRequest wrapper for the ListInternalNamespaceOccOverviews operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalNamespaceOccOverviews.go.html to see an example of how to use ListInternalNamespaceOccOverviewsRequest.
type ListInternalNamespaceOccOverviewsRequest struct {

	// The namespace by which we would filter the list.
	Namespace ListInternalNamespaceOccOverviewsNamespaceEnum `mandatory:"true" contributesTo:"path" name:"namespace"`

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The customer group ocid by which we would filter the list.
	OccCustomerGroupId *string `mandatory:"true" contributesTo:"query" name:"occCustomerGroupId"`

	// Workload type using the resources in an availability catalog can be filtered.
	WorkloadType *string `mandatory:"false" contributesTo:"query" name:"workloadType"`

	// The month corresponding to this date would be considered as the starting point of the time period against which we would like to perform an aggregation.
	From *common.SDKTime `mandatory:"false" contributesTo:"query" name:"from"`

	// The month corresponding to this date would be considered as the ending point of the time period against which we would like to perform an aggregation.
	To *common.SDKTime `mandatory:"false" contributesTo:"query" name:"to"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListInternalNamespaceOccOverviewsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. The default order for periodValue is chronological order(latest month item at the end).
	SortBy ListInternalNamespaceOccOverviewsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInternalNamespaceOccOverviewsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInternalNamespaceOccOverviewsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInternalNamespaceOccOverviewsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInternalNamespaceOccOverviewsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInternalNamespaceOccOverviewsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInternalNamespaceOccOverviewsNamespaceEnum(string(request.Namespace)); !ok && request.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", request.Namespace, strings.Join(GetListInternalNamespaceOccOverviewsNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalNamespaceOccOverviewsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInternalNamespaceOccOverviewsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalNamespaceOccOverviewsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInternalNamespaceOccOverviewsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInternalNamespaceOccOverviewsResponse wrapper for the ListInternalNamespaceOccOverviews operation
type ListInternalNamespaceOccOverviewsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OccOverviewCollection instances
	OccOverviewCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInternalNamespaceOccOverviewsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInternalNamespaceOccOverviewsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInternalNamespaceOccOverviewsNamespaceEnum Enum with underlying type: string
type ListInternalNamespaceOccOverviewsNamespaceEnum string

// Set of constants representing the allowable values for ListInternalNamespaceOccOverviewsNamespaceEnum
const (
	ListInternalNamespaceOccOverviewsNamespaceCompute ListInternalNamespaceOccOverviewsNamespaceEnum = "COMPUTE"
)

var mappingListInternalNamespaceOccOverviewsNamespaceEnum = map[string]ListInternalNamespaceOccOverviewsNamespaceEnum{
	"COMPUTE": ListInternalNamespaceOccOverviewsNamespaceCompute,
}

var mappingListInternalNamespaceOccOverviewsNamespaceEnumLowerCase = map[string]ListInternalNamespaceOccOverviewsNamespaceEnum{
	"compute": ListInternalNamespaceOccOverviewsNamespaceCompute,
}

// GetListInternalNamespaceOccOverviewsNamespaceEnumValues Enumerates the set of values for ListInternalNamespaceOccOverviewsNamespaceEnum
func GetListInternalNamespaceOccOverviewsNamespaceEnumValues() []ListInternalNamespaceOccOverviewsNamespaceEnum {
	values := make([]ListInternalNamespaceOccOverviewsNamespaceEnum, 0)
	for _, v := range mappingListInternalNamespaceOccOverviewsNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalNamespaceOccOverviewsNamespaceEnumStringValues Enumerates the set of values in String for ListInternalNamespaceOccOverviewsNamespaceEnum
func GetListInternalNamespaceOccOverviewsNamespaceEnumStringValues() []string {
	return []string{
		"COMPUTE",
	}
}

// GetMappingListInternalNamespaceOccOverviewsNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalNamespaceOccOverviewsNamespaceEnum(val string) (ListInternalNamespaceOccOverviewsNamespaceEnum, bool) {
	enum, ok := mappingListInternalNamespaceOccOverviewsNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInternalNamespaceOccOverviewsSortOrderEnum Enum with underlying type: string
type ListInternalNamespaceOccOverviewsSortOrderEnum string

// Set of constants representing the allowable values for ListInternalNamespaceOccOverviewsSortOrderEnum
const (
	ListInternalNamespaceOccOverviewsSortOrderAsc  ListInternalNamespaceOccOverviewsSortOrderEnum = "ASC"
	ListInternalNamespaceOccOverviewsSortOrderDesc ListInternalNamespaceOccOverviewsSortOrderEnum = "DESC"
)

var mappingListInternalNamespaceOccOverviewsSortOrderEnum = map[string]ListInternalNamespaceOccOverviewsSortOrderEnum{
	"ASC":  ListInternalNamespaceOccOverviewsSortOrderAsc,
	"DESC": ListInternalNamespaceOccOverviewsSortOrderDesc,
}

var mappingListInternalNamespaceOccOverviewsSortOrderEnumLowerCase = map[string]ListInternalNamespaceOccOverviewsSortOrderEnum{
	"asc":  ListInternalNamespaceOccOverviewsSortOrderAsc,
	"desc": ListInternalNamespaceOccOverviewsSortOrderDesc,
}

// GetListInternalNamespaceOccOverviewsSortOrderEnumValues Enumerates the set of values for ListInternalNamespaceOccOverviewsSortOrderEnum
func GetListInternalNamespaceOccOverviewsSortOrderEnumValues() []ListInternalNamespaceOccOverviewsSortOrderEnum {
	values := make([]ListInternalNamespaceOccOverviewsSortOrderEnum, 0)
	for _, v := range mappingListInternalNamespaceOccOverviewsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalNamespaceOccOverviewsSortOrderEnumStringValues Enumerates the set of values in String for ListInternalNamespaceOccOverviewsSortOrderEnum
func GetListInternalNamespaceOccOverviewsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInternalNamespaceOccOverviewsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalNamespaceOccOverviewsSortOrderEnum(val string) (ListInternalNamespaceOccOverviewsSortOrderEnum, bool) {
	enum, ok := mappingListInternalNamespaceOccOverviewsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInternalNamespaceOccOverviewsSortByEnum Enum with underlying type: string
type ListInternalNamespaceOccOverviewsSortByEnum string

// Set of constants representing the allowable values for ListInternalNamespaceOccOverviewsSortByEnum
const (
	ListInternalNamespaceOccOverviewsSortByPeriodvalue ListInternalNamespaceOccOverviewsSortByEnum = "periodValue"
)

var mappingListInternalNamespaceOccOverviewsSortByEnum = map[string]ListInternalNamespaceOccOverviewsSortByEnum{
	"periodValue": ListInternalNamespaceOccOverviewsSortByPeriodvalue,
}

var mappingListInternalNamespaceOccOverviewsSortByEnumLowerCase = map[string]ListInternalNamespaceOccOverviewsSortByEnum{
	"periodvalue": ListInternalNamespaceOccOverviewsSortByPeriodvalue,
}

// GetListInternalNamespaceOccOverviewsSortByEnumValues Enumerates the set of values for ListInternalNamespaceOccOverviewsSortByEnum
func GetListInternalNamespaceOccOverviewsSortByEnumValues() []ListInternalNamespaceOccOverviewsSortByEnum {
	values := make([]ListInternalNamespaceOccOverviewsSortByEnum, 0)
	for _, v := range mappingListInternalNamespaceOccOverviewsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalNamespaceOccOverviewsSortByEnumStringValues Enumerates the set of values in String for ListInternalNamespaceOccOverviewsSortByEnum
func GetListInternalNamespaceOccOverviewsSortByEnumStringValues() []string {
	return []string{
		"periodValue",
	}
}

// GetMappingListInternalNamespaceOccOverviewsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalNamespaceOccOverviewsSortByEnum(val string) (ListInternalNamespaceOccOverviewsSortByEnum, bool) {
	enum, ok := mappingListInternalNamespaceOccOverviewsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
