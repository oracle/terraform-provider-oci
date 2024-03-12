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

// ListOccCapacityRequestsRequest wrapper for the ListOccCapacityRequests operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccCapacityRequests.go.html to see an example of how to use ListOccCapacityRequestsRequest.
type ListOccCapacityRequestsRequest struct {

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return the list of capacity requests based on the OCID of the availability catalog against which they were created.
	OccAvailabilityCatalogId *string `mandatory:"false" contributesTo:"query" name:"occAvailabilityCatalogId"`

	// The namespace by which we would filter the list.
	Namespace ListOccCapacityRequestsNamespaceEnum `mandatory:"false" contributesTo:"query" name:"namespace" omitEmpty:"true"`

	// A filter to return only the resources that match the entire display name. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

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
	SortOrder ListOccCapacityRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. The default order for display name is ascending. The default order for time created is reverse chronological order(latest date at the top).
	SortBy ListOccCapacityRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccCapacityRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccCapacityRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccCapacityRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccCapacityRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccCapacityRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOccCapacityRequestsNamespaceEnum(string(request.Namespace)); !ok && request.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", request.Namespace, strings.Join(GetListOccCapacityRequestsNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccCapacityRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccCapacityRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccCapacityRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccCapacityRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccCapacityRequestsResponse wrapper for the ListOccCapacityRequests operation
type ListOccCapacityRequestsResponse struct {

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

func (response ListOccCapacityRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccCapacityRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccCapacityRequestsNamespaceEnum Enum with underlying type: string
type ListOccCapacityRequestsNamespaceEnum string

// Set of constants representing the allowable values for ListOccCapacityRequestsNamespaceEnum
const (
	ListOccCapacityRequestsNamespaceCompute ListOccCapacityRequestsNamespaceEnum = "COMPUTE"
)

var mappingListOccCapacityRequestsNamespaceEnum = map[string]ListOccCapacityRequestsNamespaceEnum{
	"COMPUTE": ListOccCapacityRequestsNamespaceCompute,
}

var mappingListOccCapacityRequestsNamespaceEnumLowerCase = map[string]ListOccCapacityRequestsNamespaceEnum{
	"compute": ListOccCapacityRequestsNamespaceCompute,
}

// GetListOccCapacityRequestsNamespaceEnumValues Enumerates the set of values for ListOccCapacityRequestsNamespaceEnum
func GetListOccCapacityRequestsNamespaceEnumValues() []ListOccCapacityRequestsNamespaceEnum {
	values := make([]ListOccCapacityRequestsNamespaceEnum, 0)
	for _, v := range mappingListOccCapacityRequestsNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccCapacityRequestsNamespaceEnumStringValues Enumerates the set of values in String for ListOccCapacityRequestsNamespaceEnum
func GetListOccCapacityRequestsNamespaceEnumStringValues() []string {
	return []string{
		"COMPUTE",
	}
}

// GetMappingListOccCapacityRequestsNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccCapacityRequestsNamespaceEnum(val string) (ListOccCapacityRequestsNamespaceEnum, bool) {
	enum, ok := mappingListOccCapacityRequestsNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccCapacityRequestsSortOrderEnum Enum with underlying type: string
type ListOccCapacityRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListOccCapacityRequestsSortOrderEnum
const (
	ListOccCapacityRequestsSortOrderAsc  ListOccCapacityRequestsSortOrderEnum = "ASC"
	ListOccCapacityRequestsSortOrderDesc ListOccCapacityRequestsSortOrderEnum = "DESC"
)

var mappingListOccCapacityRequestsSortOrderEnum = map[string]ListOccCapacityRequestsSortOrderEnum{
	"ASC":  ListOccCapacityRequestsSortOrderAsc,
	"DESC": ListOccCapacityRequestsSortOrderDesc,
}

var mappingListOccCapacityRequestsSortOrderEnumLowerCase = map[string]ListOccCapacityRequestsSortOrderEnum{
	"asc":  ListOccCapacityRequestsSortOrderAsc,
	"desc": ListOccCapacityRequestsSortOrderDesc,
}

// GetListOccCapacityRequestsSortOrderEnumValues Enumerates the set of values for ListOccCapacityRequestsSortOrderEnum
func GetListOccCapacityRequestsSortOrderEnumValues() []ListOccCapacityRequestsSortOrderEnum {
	values := make([]ListOccCapacityRequestsSortOrderEnum, 0)
	for _, v := range mappingListOccCapacityRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccCapacityRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListOccCapacityRequestsSortOrderEnum
func GetListOccCapacityRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccCapacityRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccCapacityRequestsSortOrderEnum(val string) (ListOccCapacityRequestsSortOrderEnum, bool) {
	enum, ok := mappingListOccCapacityRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccCapacityRequestsSortByEnum Enum with underlying type: string
type ListOccCapacityRequestsSortByEnum string

// Set of constants representing the allowable values for ListOccCapacityRequestsSortByEnum
const (
	ListOccCapacityRequestsSortByDisplayname ListOccCapacityRequestsSortByEnum = "displayName"
	ListOccCapacityRequestsSortByTimecreated ListOccCapacityRequestsSortByEnum = "timeCreated"
)

var mappingListOccCapacityRequestsSortByEnum = map[string]ListOccCapacityRequestsSortByEnum{
	"displayName": ListOccCapacityRequestsSortByDisplayname,
	"timeCreated": ListOccCapacityRequestsSortByTimecreated,
}

var mappingListOccCapacityRequestsSortByEnumLowerCase = map[string]ListOccCapacityRequestsSortByEnum{
	"displayname": ListOccCapacityRequestsSortByDisplayname,
	"timecreated": ListOccCapacityRequestsSortByTimecreated,
}

// GetListOccCapacityRequestsSortByEnumValues Enumerates the set of values for ListOccCapacityRequestsSortByEnum
func GetListOccCapacityRequestsSortByEnumValues() []ListOccCapacityRequestsSortByEnum {
	values := make([]ListOccCapacityRequestsSortByEnum, 0)
	for _, v := range mappingListOccCapacityRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccCapacityRequestsSortByEnumStringValues Enumerates the set of values in String for ListOccCapacityRequestsSortByEnum
func GetListOccCapacityRequestsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListOccCapacityRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccCapacityRequestsSortByEnum(val string) (ListOccCapacityRequestsSortByEnum, bool) {
	enum, ok := mappingListOccCapacityRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
