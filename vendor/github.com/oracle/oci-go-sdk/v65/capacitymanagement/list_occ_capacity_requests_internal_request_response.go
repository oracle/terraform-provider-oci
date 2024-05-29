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

// ListOccCapacityRequestsInternalRequest wrapper for the ListOccCapacityRequestsInternal operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccCapacityRequestsInternal.go.html to see an example of how to use ListOccCapacityRequestsInternalRequest.
type ListOccCapacityRequestsInternalRequest struct {

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The customer group ocid by which we would filter the list.
	OccCustomerGroupId *string `mandatory:"false" contributesTo:"query" name:"occCustomerGroupId"`

	// A filter to return the list of capacity requests based on the OCID of the availability catalog against which they were created.
	OccAvailabilityCatalogId *string `mandatory:"false" contributesTo:"query" name:"occAvailabilityCatalogId"`

	// The namespace by which we would filter the list.
	Namespace ListOccCapacityRequestsInternalNamespaceEnum `mandatory:"false" contributesTo:"query" name:"namespace" omitEmpty:"true"`

	// A filter to return only the resources that match the entire display name. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only the resources that match the request type. The match is not case sensitive.
	RequestType OccCapacityRequestRequestTypeEnum `mandatory:"false" contributesTo:"query" name:"requestType" omitEmpty:"true"`

	// A filter to return the list of capacity requests based on the OCID of the capacity request. This is done for the users who have INSPECT permission on the resource but do not have READ permission.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOccCapacityRequestsInternalSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. The default order for display name is ascending. The default order for time created is reverse chronological order(latest date at the top).
	SortBy ListOccCapacityRequestsInternalSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccCapacityRequestsInternalRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccCapacityRequestsInternalRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccCapacityRequestsInternalRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccCapacityRequestsInternalRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccCapacityRequestsInternalRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOccCapacityRequestsInternalNamespaceEnum(string(request.Namespace)); !ok && request.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", request.Namespace, strings.Join(GetListOccCapacityRequestsInternalNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccCapacityRequestRequestTypeEnum(string(request.RequestType)); !ok && request.RequestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestType: %s. Supported values are: %s.", request.RequestType, strings.Join(GetOccCapacityRequestRequestTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccCapacityRequestsInternalSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccCapacityRequestsInternalSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccCapacityRequestsInternalSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccCapacityRequestsInternalSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccCapacityRequestsInternalResponse wrapper for the ListOccCapacityRequestsInternal operation
type ListOccCapacityRequestsInternalResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OccCapacityRequestCollection instances
	OccCapacityRequestCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOccCapacityRequestsInternalResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccCapacityRequestsInternalResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccCapacityRequestsInternalNamespaceEnum Enum with underlying type: string
type ListOccCapacityRequestsInternalNamespaceEnum string

// Set of constants representing the allowable values for ListOccCapacityRequestsInternalNamespaceEnum
const (
	ListOccCapacityRequestsInternalNamespaceCompute ListOccCapacityRequestsInternalNamespaceEnum = "COMPUTE"
)

var mappingListOccCapacityRequestsInternalNamespaceEnum = map[string]ListOccCapacityRequestsInternalNamespaceEnum{
	"COMPUTE": ListOccCapacityRequestsInternalNamespaceCompute,
}

var mappingListOccCapacityRequestsInternalNamespaceEnumLowerCase = map[string]ListOccCapacityRequestsInternalNamespaceEnum{
	"compute": ListOccCapacityRequestsInternalNamespaceCompute,
}

// GetListOccCapacityRequestsInternalNamespaceEnumValues Enumerates the set of values for ListOccCapacityRequestsInternalNamespaceEnum
func GetListOccCapacityRequestsInternalNamespaceEnumValues() []ListOccCapacityRequestsInternalNamespaceEnum {
	values := make([]ListOccCapacityRequestsInternalNamespaceEnum, 0)
	for _, v := range mappingListOccCapacityRequestsInternalNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccCapacityRequestsInternalNamespaceEnumStringValues Enumerates the set of values in String for ListOccCapacityRequestsInternalNamespaceEnum
func GetListOccCapacityRequestsInternalNamespaceEnumStringValues() []string {
	return []string{
		"COMPUTE",
	}
}

// GetMappingListOccCapacityRequestsInternalNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccCapacityRequestsInternalNamespaceEnum(val string) (ListOccCapacityRequestsInternalNamespaceEnum, bool) {
	enum, ok := mappingListOccCapacityRequestsInternalNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccCapacityRequestsInternalSortOrderEnum Enum with underlying type: string
type ListOccCapacityRequestsInternalSortOrderEnum string

// Set of constants representing the allowable values for ListOccCapacityRequestsInternalSortOrderEnum
const (
	ListOccCapacityRequestsInternalSortOrderAsc  ListOccCapacityRequestsInternalSortOrderEnum = "ASC"
	ListOccCapacityRequestsInternalSortOrderDesc ListOccCapacityRequestsInternalSortOrderEnum = "DESC"
)

var mappingListOccCapacityRequestsInternalSortOrderEnum = map[string]ListOccCapacityRequestsInternalSortOrderEnum{
	"ASC":  ListOccCapacityRequestsInternalSortOrderAsc,
	"DESC": ListOccCapacityRequestsInternalSortOrderDesc,
}

var mappingListOccCapacityRequestsInternalSortOrderEnumLowerCase = map[string]ListOccCapacityRequestsInternalSortOrderEnum{
	"asc":  ListOccCapacityRequestsInternalSortOrderAsc,
	"desc": ListOccCapacityRequestsInternalSortOrderDesc,
}

// GetListOccCapacityRequestsInternalSortOrderEnumValues Enumerates the set of values for ListOccCapacityRequestsInternalSortOrderEnum
func GetListOccCapacityRequestsInternalSortOrderEnumValues() []ListOccCapacityRequestsInternalSortOrderEnum {
	values := make([]ListOccCapacityRequestsInternalSortOrderEnum, 0)
	for _, v := range mappingListOccCapacityRequestsInternalSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccCapacityRequestsInternalSortOrderEnumStringValues Enumerates the set of values in String for ListOccCapacityRequestsInternalSortOrderEnum
func GetListOccCapacityRequestsInternalSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccCapacityRequestsInternalSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccCapacityRequestsInternalSortOrderEnum(val string) (ListOccCapacityRequestsInternalSortOrderEnum, bool) {
	enum, ok := mappingListOccCapacityRequestsInternalSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccCapacityRequestsInternalSortByEnum Enum with underlying type: string
type ListOccCapacityRequestsInternalSortByEnum string

// Set of constants representing the allowable values for ListOccCapacityRequestsInternalSortByEnum
const (
	ListOccCapacityRequestsInternalSortByDisplayname ListOccCapacityRequestsInternalSortByEnum = "displayName"
	ListOccCapacityRequestsInternalSortByTimecreated ListOccCapacityRequestsInternalSortByEnum = "timeCreated"
)

var mappingListOccCapacityRequestsInternalSortByEnum = map[string]ListOccCapacityRequestsInternalSortByEnum{
	"displayName": ListOccCapacityRequestsInternalSortByDisplayname,
	"timeCreated": ListOccCapacityRequestsInternalSortByTimecreated,
}

var mappingListOccCapacityRequestsInternalSortByEnumLowerCase = map[string]ListOccCapacityRequestsInternalSortByEnum{
	"displayname": ListOccCapacityRequestsInternalSortByDisplayname,
	"timecreated": ListOccCapacityRequestsInternalSortByTimecreated,
}

// GetListOccCapacityRequestsInternalSortByEnumValues Enumerates the set of values for ListOccCapacityRequestsInternalSortByEnum
func GetListOccCapacityRequestsInternalSortByEnumValues() []ListOccCapacityRequestsInternalSortByEnum {
	values := make([]ListOccCapacityRequestsInternalSortByEnum, 0)
	for _, v := range mappingListOccCapacityRequestsInternalSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccCapacityRequestsInternalSortByEnumStringValues Enumerates the set of values in String for ListOccCapacityRequestsInternalSortByEnum
func GetListOccCapacityRequestsInternalSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListOccCapacityRequestsInternalSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccCapacityRequestsInternalSortByEnum(val string) (ListOccCapacityRequestsInternalSortByEnum, bool) {
	enum, ok := mappingListOccCapacityRequestsInternalSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
