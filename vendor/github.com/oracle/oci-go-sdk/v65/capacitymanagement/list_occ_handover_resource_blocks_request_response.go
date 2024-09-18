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

// ListOccHandoverResourceBlocksRequest wrapper for the ListOccHandoverResourceBlocks operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccHandoverResourceBlocks.go.html to see an example of how to use ListOccHandoverResourceBlocksRequest.
type ListOccHandoverResourceBlocksRequest struct {

	// The namespace by which we would filter the list.
	Namespace ListOccHandoverResourceBlocksNamespaceEnum `mandatory:"false" contributesTo:"query" name:"namespace" omitEmpty:"true"`

	// The OCID of the compartment or tenancy in which resources are to be listed.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

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
	SortOrder ListOccHandoverResourceBlocksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// The default order for handoverDate is chronological order(latest date item at the end).
	SortBy ListOccHandoverResourceBlocksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccHandoverResourceBlocksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccHandoverResourceBlocksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccHandoverResourceBlocksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccHandoverResourceBlocksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccHandoverResourceBlocksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOccHandoverResourceBlocksNamespaceEnum(string(request.Namespace)); !ok && request.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", request.Namespace, strings.Join(GetListOccHandoverResourceBlocksNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccHandoverResourceBlocksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccHandoverResourceBlocksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccHandoverResourceBlocksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccHandoverResourceBlocksSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccHandoverResourceBlocksResponse wrapper for the ListOccHandoverResourceBlocks operation
type ListOccHandoverResourceBlocksResponse struct {

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

func (response ListOccHandoverResourceBlocksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccHandoverResourceBlocksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccHandoverResourceBlocksNamespaceEnum Enum with underlying type: string
type ListOccHandoverResourceBlocksNamespaceEnum string

// Set of constants representing the allowable values for ListOccHandoverResourceBlocksNamespaceEnum
const (
	ListOccHandoverResourceBlocksNamespaceCompute ListOccHandoverResourceBlocksNamespaceEnum = "COMPUTE"
)

var mappingListOccHandoverResourceBlocksNamespaceEnum = map[string]ListOccHandoverResourceBlocksNamespaceEnum{
	"COMPUTE": ListOccHandoverResourceBlocksNamespaceCompute,
}

var mappingListOccHandoverResourceBlocksNamespaceEnumLowerCase = map[string]ListOccHandoverResourceBlocksNamespaceEnum{
	"compute": ListOccHandoverResourceBlocksNamespaceCompute,
}

// GetListOccHandoverResourceBlocksNamespaceEnumValues Enumerates the set of values for ListOccHandoverResourceBlocksNamespaceEnum
func GetListOccHandoverResourceBlocksNamespaceEnumValues() []ListOccHandoverResourceBlocksNamespaceEnum {
	values := make([]ListOccHandoverResourceBlocksNamespaceEnum, 0)
	for _, v := range mappingListOccHandoverResourceBlocksNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccHandoverResourceBlocksNamespaceEnumStringValues Enumerates the set of values in String for ListOccHandoverResourceBlocksNamespaceEnum
func GetListOccHandoverResourceBlocksNamespaceEnumStringValues() []string {
	return []string{
		"COMPUTE",
	}
}

// GetMappingListOccHandoverResourceBlocksNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccHandoverResourceBlocksNamespaceEnum(val string) (ListOccHandoverResourceBlocksNamespaceEnum, bool) {
	enum, ok := mappingListOccHandoverResourceBlocksNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccHandoverResourceBlocksSortOrderEnum Enum with underlying type: string
type ListOccHandoverResourceBlocksSortOrderEnum string

// Set of constants representing the allowable values for ListOccHandoverResourceBlocksSortOrderEnum
const (
	ListOccHandoverResourceBlocksSortOrderAsc  ListOccHandoverResourceBlocksSortOrderEnum = "ASC"
	ListOccHandoverResourceBlocksSortOrderDesc ListOccHandoverResourceBlocksSortOrderEnum = "DESC"
)

var mappingListOccHandoverResourceBlocksSortOrderEnum = map[string]ListOccHandoverResourceBlocksSortOrderEnum{
	"ASC":  ListOccHandoverResourceBlocksSortOrderAsc,
	"DESC": ListOccHandoverResourceBlocksSortOrderDesc,
}

var mappingListOccHandoverResourceBlocksSortOrderEnumLowerCase = map[string]ListOccHandoverResourceBlocksSortOrderEnum{
	"asc":  ListOccHandoverResourceBlocksSortOrderAsc,
	"desc": ListOccHandoverResourceBlocksSortOrderDesc,
}

// GetListOccHandoverResourceBlocksSortOrderEnumValues Enumerates the set of values for ListOccHandoverResourceBlocksSortOrderEnum
func GetListOccHandoverResourceBlocksSortOrderEnumValues() []ListOccHandoverResourceBlocksSortOrderEnum {
	values := make([]ListOccHandoverResourceBlocksSortOrderEnum, 0)
	for _, v := range mappingListOccHandoverResourceBlocksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccHandoverResourceBlocksSortOrderEnumStringValues Enumerates the set of values in String for ListOccHandoverResourceBlocksSortOrderEnum
func GetListOccHandoverResourceBlocksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccHandoverResourceBlocksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccHandoverResourceBlocksSortOrderEnum(val string) (ListOccHandoverResourceBlocksSortOrderEnum, bool) {
	enum, ok := mappingListOccHandoverResourceBlocksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccHandoverResourceBlocksSortByEnum Enum with underlying type: string
type ListOccHandoverResourceBlocksSortByEnum string

// Set of constants representing the allowable values for ListOccHandoverResourceBlocksSortByEnum
const (
	ListOccHandoverResourceBlocksSortByHandoverdate ListOccHandoverResourceBlocksSortByEnum = "handoverDate"
)

var mappingListOccHandoverResourceBlocksSortByEnum = map[string]ListOccHandoverResourceBlocksSortByEnum{
	"handoverDate": ListOccHandoverResourceBlocksSortByHandoverdate,
}

var mappingListOccHandoverResourceBlocksSortByEnumLowerCase = map[string]ListOccHandoverResourceBlocksSortByEnum{
	"handoverdate": ListOccHandoverResourceBlocksSortByHandoverdate,
}

// GetListOccHandoverResourceBlocksSortByEnumValues Enumerates the set of values for ListOccHandoverResourceBlocksSortByEnum
func GetListOccHandoverResourceBlocksSortByEnumValues() []ListOccHandoverResourceBlocksSortByEnum {
	values := make([]ListOccHandoverResourceBlocksSortByEnum, 0)
	for _, v := range mappingListOccHandoverResourceBlocksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccHandoverResourceBlocksSortByEnumStringValues Enumerates the set of values in String for ListOccHandoverResourceBlocksSortByEnum
func GetListOccHandoverResourceBlocksSortByEnumStringValues() []string {
	return []string{
		"handoverDate",
	}
}

// GetMappingListOccHandoverResourceBlocksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccHandoverResourceBlocksSortByEnum(val string) (ListOccHandoverResourceBlocksSortByEnum, bool) {
	enum, ok := mappingListOccHandoverResourceBlocksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
