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

// ListOccOverviewsRequest wrapper for the ListOccOverviews operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccOverviews.go.html to see an example of how to use ListOccOverviewsRequest.
type ListOccOverviewsRequest struct {

	// The namespace by which we would filter the list.
	Namespace ListOccOverviewsNamespaceEnum `mandatory:"true" contributesTo:"path" name:"namespace"`

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The month corresponding to this date would be considered as the starting point of the time period against which we would like to perform an aggregation.
	From *common.SDKTime `mandatory:"false" contributesTo:"query" name:"from"`

	// The month corresponding to this date would be considered as the ending point of the time period against which we would like to perform an aggregation.
	To *common.SDKTime `mandatory:"false" contributesTo:"query" name:"to"`

	// Workload type using the resources in an availability catalog can be filtered.
	WorkloadType *string `mandatory:"false" contributesTo:"query" name:"workloadType"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOccOverviewsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. The default order for periodValue is chronological order(latest month item at the end).
	SortBy ListOccOverviewsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccOverviewsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccOverviewsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccOverviewsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccOverviewsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccOverviewsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOccOverviewsNamespaceEnum(string(request.Namespace)); !ok && request.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", request.Namespace, strings.Join(GetListOccOverviewsNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccOverviewsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccOverviewsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccOverviewsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccOverviewsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccOverviewsResponse wrapper for the ListOccOverviews operation
type ListOccOverviewsResponse struct {

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

func (response ListOccOverviewsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccOverviewsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccOverviewsNamespaceEnum Enum with underlying type: string
type ListOccOverviewsNamespaceEnum string

// Set of constants representing the allowable values for ListOccOverviewsNamespaceEnum
const (
	ListOccOverviewsNamespaceCompute ListOccOverviewsNamespaceEnum = "COMPUTE"
)

var mappingListOccOverviewsNamespaceEnum = map[string]ListOccOverviewsNamespaceEnum{
	"COMPUTE": ListOccOverviewsNamespaceCompute,
}

var mappingListOccOverviewsNamespaceEnumLowerCase = map[string]ListOccOverviewsNamespaceEnum{
	"compute": ListOccOverviewsNamespaceCompute,
}

// GetListOccOverviewsNamespaceEnumValues Enumerates the set of values for ListOccOverviewsNamespaceEnum
func GetListOccOverviewsNamespaceEnumValues() []ListOccOverviewsNamespaceEnum {
	values := make([]ListOccOverviewsNamespaceEnum, 0)
	for _, v := range mappingListOccOverviewsNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccOverviewsNamespaceEnumStringValues Enumerates the set of values in String for ListOccOverviewsNamespaceEnum
func GetListOccOverviewsNamespaceEnumStringValues() []string {
	return []string{
		"COMPUTE",
	}
}

// GetMappingListOccOverviewsNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccOverviewsNamespaceEnum(val string) (ListOccOverviewsNamespaceEnum, bool) {
	enum, ok := mappingListOccOverviewsNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccOverviewsSortOrderEnum Enum with underlying type: string
type ListOccOverviewsSortOrderEnum string

// Set of constants representing the allowable values for ListOccOverviewsSortOrderEnum
const (
	ListOccOverviewsSortOrderAsc  ListOccOverviewsSortOrderEnum = "ASC"
	ListOccOverviewsSortOrderDesc ListOccOverviewsSortOrderEnum = "DESC"
)

var mappingListOccOverviewsSortOrderEnum = map[string]ListOccOverviewsSortOrderEnum{
	"ASC":  ListOccOverviewsSortOrderAsc,
	"DESC": ListOccOverviewsSortOrderDesc,
}

var mappingListOccOverviewsSortOrderEnumLowerCase = map[string]ListOccOverviewsSortOrderEnum{
	"asc":  ListOccOverviewsSortOrderAsc,
	"desc": ListOccOverviewsSortOrderDesc,
}

// GetListOccOverviewsSortOrderEnumValues Enumerates the set of values for ListOccOverviewsSortOrderEnum
func GetListOccOverviewsSortOrderEnumValues() []ListOccOverviewsSortOrderEnum {
	values := make([]ListOccOverviewsSortOrderEnum, 0)
	for _, v := range mappingListOccOverviewsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccOverviewsSortOrderEnumStringValues Enumerates the set of values in String for ListOccOverviewsSortOrderEnum
func GetListOccOverviewsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccOverviewsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccOverviewsSortOrderEnum(val string) (ListOccOverviewsSortOrderEnum, bool) {
	enum, ok := mappingListOccOverviewsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccOverviewsSortByEnum Enum with underlying type: string
type ListOccOverviewsSortByEnum string

// Set of constants representing the allowable values for ListOccOverviewsSortByEnum
const (
	ListOccOverviewsSortByPeriodvalue ListOccOverviewsSortByEnum = "periodValue"
)

var mappingListOccOverviewsSortByEnum = map[string]ListOccOverviewsSortByEnum{
	"periodValue": ListOccOverviewsSortByPeriodvalue,
}

var mappingListOccOverviewsSortByEnumLowerCase = map[string]ListOccOverviewsSortByEnum{
	"periodvalue": ListOccOverviewsSortByPeriodvalue,
}

// GetListOccOverviewsSortByEnumValues Enumerates the set of values for ListOccOverviewsSortByEnum
func GetListOccOverviewsSortByEnumValues() []ListOccOverviewsSortByEnum {
	values := make([]ListOccOverviewsSortByEnum, 0)
	for _, v := range mappingListOccOverviewsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccOverviewsSortByEnumStringValues Enumerates the set of values in String for ListOccOverviewsSortByEnum
func GetListOccOverviewsSortByEnumStringValues() []string {
	return []string{
		"periodValue",
	}
}

// GetMappingListOccOverviewsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccOverviewsSortByEnum(val string) (ListOccOverviewsSortByEnum, bool) {
	enum, ok := mappingListOccOverviewsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
