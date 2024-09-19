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

// ListInternalOccHandoverResourceBlocksRequest wrapper for the ListInternalOccHandoverResourceBlocks operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalOccHandoverResourceBlocks.go.html to see an example of how to use ListInternalOccHandoverResourceBlocksRequest.
type ListInternalOccHandoverResourceBlocksRequest struct {

	// The namespace enum value that needs to be passed as a required query parameter.
	Namespace ListInternalOccHandoverResourceBlocksNamespaceEnum `mandatory:"true" contributesTo:"query" name:"namespace" omitEmpty:"true"`

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The customer group ocid by which we would filter the list.
	OccCustomerGroupId *string `mandatory:"true" contributesTo:"query" name:"occCustomerGroupId"`

	// A filter to return only the list of resources that match the name provided in this filter.
	HandoverResourceName *string `mandatory:"false" contributesTo:"query" name:"handoverResourceName"`

	// This filter helps in fetching all handed over resources for which the recordDate is greater than or equal to the startDate.
	HandoverDateGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"handoverDateGreaterThanOrEqualTo"`

	// This filter helps in fetching all handed over resources for which the recordDate is less than or equal to the endDate.
	HandoverDateLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"handoverDateLessThanOrEqualTo"`

	// This filter helps in fetching the handed over resource for which the occHandoverResourceId is equal to the one provided here.
	OccHandoverResourceBlockId *string `mandatory:"false" contributesTo:"query" name:"occHandoverResourceBlockId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListInternalOccHandoverResourceBlocksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// The default order for handoverDate is chronological order(latest date item at the end).
	SortBy ListInternalOccHandoverResourceBlocksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInternalOccHandoverResourceBlocksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInternalOccHandoverResourceBlocksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInternalOccHandoverResourceBlocksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInternalOccHandoverResourceBlocksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInternalOccHandoverResourceBlocksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInternalOccHandoverResourceBlocksNamespaceEnum(string(request.Namespace)); !ok && request.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", request.Namespace, strings.Join(GetListInternalOccHandoverResourceBlocksNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalOccHandoverResourceBlocksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInternalOccHandoverResourceBlocksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalOccHandoverResourceBlocksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInternalOccHandoverResourceBlocksSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInternalOccHandoverResourceBlocksResponse wrapper for the ListInternalOccHandoverResourceBlocks operation
type ListInternalOccHandoverResourceBlocksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OccHandoverResourceBlockCollection instances
	OccHandoverResourceBlockCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInternalOccHandoverResourceBlocksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInternalOccHandoverResourceBlocksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInternalOccHandoverResourceBlocksNamespaceEnum Enum with underlying type: string
type ListInternalOccHandoverResourceBlocksNamespaceEnum string

// Set of constants representing the allowable values for ListInternalOccHandoverResourceBlocksNamespaceEnum
const (
	ListInternalOccHandoverResourceBlocksNamespaceCompute ListInternalOccHandoverResourceBlocksNamespaceEnum = "COMPUTE"
)

var mappingListInternalOccHandoverResourceBlocksNamespaceEnum = map[string]ListInternalOccHandoverResourceBlocksNamespaceEnum{
	"COMPUTE": ListInternalOccHandoverResourceBlocksNamespaceCompute,
}

var mappingListInternalOccHandoverResourceBlocksNamespaceEnumLowerCase = map[string]ListInternalOccHandoverResourceBlocksNamespaceEnum{
	"compute": ListInternalOccHandoverResourceBlocksNamespaceCompute,
}

// GetListInternalOccHandoverResourceBlocksNamespaceEnumValues Enumerates the set of values for ListInternalOccHandoverResourceBlocksNamespaceEnum
func GetListInternalOccHandoverResourceBlocksNamespaceEnumValues() []ListInternalOccHandoverResourceBlocksNamespaceEnum {
	values := make([]ListInternalOccHandoverResourceBlocksNamespaceEnum, 0)
	for _, v := range mappingListInternalOccHandoverResourceBlocksNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccHandoverResourceBlocksNamespaceEnumStringValues Enumerates the set of values in String for ListInternalOccHandoverResourceBlocksNamespaceEnum
func GetListInternalOccHandoverResourceBlocksNamespaceEnumStringValues() []string {
	return []string{
		"COMPUTE",
	}
}

// GetMappingListInternalOccHandoverResourceBlocksNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccHandoverResourceBlocksNamespaceEnum(val string) (ListInternalOccHandoverResourceBlocksNamespaceEnum, bool) {
	enum, ok := mappingListInternalOccHandoverResourceBlocksNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInternalOccHandoverResourceBlocksSortOrderEnum Enum with underlying type: string
type ListInternalOccHandoverResourceBlocksSortOrderEnum string

// Set of constants representing the allowable values for ListInternalOccHandoverResourceBlocksSortOrderEnum
const (
	ListInternalOccHandoverResourceBlocksSortOrderAsc  ListInternalOccHandoverResourceBlocksSortOrderEnum = "ASC"
	ListInternalOccHandoverResourceBlocksSortOrderDesc ListInternalOccHandoverResourceBlocksSortOrderEnum = "DESC"
)

var mappingListInternalOccHandoverResourceBlocksSortOrderEnum = map[string]ListInternalOccHandoverResourceBlocksSortOrderEnum{
	"ASC":  ListInternalOccHandoverResourceBlocksSortOrderAsc,
	"DESC": ListInternalOccHandoverResourceBlocksSortOrderDesc,
}

var mappingListInternalOccHandoverResourceBlocksSortOrderEnumLowerCase = map[string]ListInternalOccHandoverResourceBlocksSortOrderEnum{
	"asc":  ListInternalOccHandoverResourceBlocksSortOrderAsc,
	"desc": ListInternalOccHandoverResourceBlocksSortOrderDesc,
}

// GetListInternalOccHandoverResourceBlocksSortOrderEnumValues Enumerates the set of values for ListInternalOccHandoverResourceBlocksSortOrderEnum
func GetListInternalOccHandoverResourceBlocksSortOrderEnumValues() []ListInternalOccHandoverResourceBlocksSortOrderEnum {
	values := make([]ListInternalOccHandoverResourceBlocksSortOrderEnum, 0)
	for _, v := range mappingListInternalOccHandoverResourceBlocksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccHandoverResourceBlocksSortOrderEnumStringValues Enumerates the set of values in String for ListInternalOccHandoverResourceBlocksSortOrderEnum
func GetListInternalOccHandoverResourceBlocksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInternalOccHandoverResourceBlocksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccHandoverResourceBlocksSortOrderEnum(val string) (ListInternalOccHandoverResourceBlocksSortOrderEnum, bool) {
	enum, ok := mappingListInternalOccHandoverResourceBlocksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInternalOccHandoverResourceBlocksSortByEnum Enum with underlying type: string
type ListInternalOccHandoverResourceBlocksSortByEnum string

// Set of constants representing the allowable values for ListInternalOccHandoverResourceBlocksSortByEnum
const (
	ListInternalOccHandoverResourceBlocksSortByHandoverdate ListInternalOccHandoverResourceBlocksSortByEnum = "handoverDate"
)

var mappingListInternalOccHandoverResourceBlocksSortByEnum = map[string]ListInternalOccHandoverResourceBlocksSortByEnum{
	"handoverDate": ListInternalOccHandoverResourceBlocksSortByHandoverdate,
}

var mappingListInternalOccHandoverResourceBlocksSortByEnumLowerCase = map[string]ListInternalOccHandoverResourceBlocksSortByEnum{
	"handoverdate": ListInternalOccHandoverResourceBlocksSortByHandoverdate,
}

// GetListInternalOccHandoverResourceBlocksSortByEnumValues Enumerates the set of values for ListInternalOccHandoverResourceBlocksSortByEnum
func GetListInternalOccHandoverResourceBlocksSortByEnumValues() []ListInternalOccHandoverResourceBlocksSortByEnum {
	values := make([]ListInternalOccHandoverResourceBlocksSortByEnum, 0)
	for _, v := range mappingListInternalOccHandoverResourceBlocksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccHandoverResourceBlocksSortByEnumStringValues Enumerates the set of values in String for ListInternalOccHandoverResourceBlocksSortByEnum
func GetListInternalOccHandoverResourceBlocksSortByEnumStringValues() []string {
	return []string{
		"handoverDate",
	}
}

// GetMappingListInternalOccHandoverResourceBlocksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccHandoverResourceBlocksSortByEnum(val string) (ListInternalOccHandoverResourceBlocksSortByEnum, bool) {
	enum, ok := mappingListInternalOccHandoverResourceBlocksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
