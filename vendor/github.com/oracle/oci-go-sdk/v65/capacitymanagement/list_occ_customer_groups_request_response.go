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

// ListOccCustomerGroupsRequest wrapper for the ListOccCustomerGroups operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccCustomerGroups.go.html to see an example of how to use ListOccCustomerGroupsRequest.
type ListOccCustomerGroupsRequest struct {

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A query filter to return the list result based on status.
	Status OccCustomerGroupStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// A filter to return only the resources that match the entire display name. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A query filter to return the list result based on the customer group OCID. This is done for users who have INSPECT permission but do not have READ permission.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOccCustomerGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for name and compartment ID is ascending. Default order for time created is descending.
	SortBy ListOccCustomerGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccCustomerGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccCustomerGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccCustomerGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccCustomerGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccCustomerGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOccCustomerGroupStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetOccCustomerGroupStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccCustomerGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccCustomerGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccCustomerGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccCustomerGroupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccCustomerGroupsResponse wrapper for the ListOccCustomerGroups operation
type ListOccCustomerGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OccCustomerGroupCollection instances
	OccCustomerGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOccCustomerGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccCustomerGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccCustomerGroupsSortOrderEnum Enum with underlying type: string
type ListOccCustomerGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListOccCustomerGroupsSortOrderEnum
const (
	ListOccCustomerGroupsSortOrderAsc  ListOccCustomerGroupsSortOrderEnum = "ASC"
	ListOccCustomerGroupsSortOrderDesc ListOccCustomerGroupsSortOrderEnum = "DESC"
)

var mappingListOccCustomerGroupsSortOrderEnum = map[string]ListOccCustomerGroupsSortOrderEnum{
	"ASC":  ListOccCustomerGroupsSortOrderAsc,
	"DESC": ListOccCustomerGroupsSortOrderDesc,
}

var mappingListOccCustomerGroupsSortOrderEnumLowerCase = map[string]ListOccCustomerGroupsSortOrderEnum{
	"asc":  ListOccCustomerGroupsSortOrderAsc,
	"desc": ListOccCustomerGroupsSortOrderDesc,
}

// GetListOccCustomerGroupsSortOrderEnumValues Enumerates the set of values for ListOccCustomerGroupsSortOrderEnum
func GetListOccCustomerGroupsSortOrderEnumValues() []ListOccCustomerGroupsSortOrderEnum {
	values := make([]ListOccCustomerGroupsSortOrderEnum, 0)
	for _, v := range mappingListOccCustomerGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccCustomerGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListOccCustomerGroupsSortOrderEnum
func GetListOccCustomerGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccCustomerGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccCustomerGroupsSortOrderEnum(val string) (ListOccCustomerGroupsSortOrderEnum, bool) {
	enum, ok := mappingListOccCustomerGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccCustomerGroupsSortByEnum Enum with underlying type: string
type ListOccCustomerGroupsSortByEnum string

// Set of constants representing the allowable values for ListOccCustomerGroupsSortByEnum
const (
	ListOccCustomerGroupsSortByCompartmentid ListOccCustomerGroupsSortByEnum = "compartmentId"
	ListOccCustomerGroupsSortByName          ListOccCustomerGroupsSortByEnum = "name"
	ListOccCustomerGroupsSortByTimecreated   ListOccCustomerGroupsSortByEnum = "timeCreated"
)

var mappingListOccCustomerGroupsSortByEnum = map[string]ListOccCustomerGroupsSortByEnum{
	"compartmentId": ListOccCustomerGroupsSortByCompartmentid,
	"name":          ListOccCustomerGroupsSortByName,
	"timeCreated":   ListOccCustomerGroupsSortByTimecreated,
}

var mappingListOccCustomerGroupsSortByEnumLowerCase = map[string]ListOccCustomerGroupsSortByEnum{
	"compartmentid": ListOccCustomerGroupsSortByCompartmentid,
	"name":          ListOccCustomerGroupsSortByName,
	"timecreated":   ListOccCustomerGroupsSortByTimecreated,
}

// GetListOccCustomerGroupsSortByEnumValues Enumerates the set of values for ListOccCustomerGroupsSortByEnum
func GetListOccCustomerGroupsSortByEnumValues() []ListOccCustomerGroupsSortByEnum {
	values := make([]ListOccCustomerGroupsSortByEnum, 0)
	for _, v := range mappingListOccCustomerGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccCustomerGroupsSortByEnumStringValues Enumerates the set of values in String for ListOccCustomerGroupsSortByEnum
func GetListOccCustomerGroupsSortByEnumStringValues() []string {
	return []string{
		"compartmentId",
		"name",
		"timeCreated",
	}
}

// GetMappingListOccCustomerGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccCustomerGroupsSortByEnum(val string) (ListOccCustomerGroupsSortByEnum, bool) {
	enum, ok := mappingListOccCustomerGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
