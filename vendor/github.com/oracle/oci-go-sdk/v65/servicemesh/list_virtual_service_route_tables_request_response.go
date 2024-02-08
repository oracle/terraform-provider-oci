// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListVirtualServiceRouteTablesRequest wrapper for the ListVirtualServiceRouteTables operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListVirtualServiceRouteTables.go.html to see an example of how to use ListVirtualServiceRouteTablesRequest.
type ListVirtualServiceRouteTablesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListVirtualServiceRouteTablesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for 'timeCreated' is descending. Default order for 'name' is ascending.
	SortBy ListVirtualServiceRouteTablesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Unique VirtualService identifier.
	VirtualServiceId *string `mandatory:"false" contributesTo:"query" name:"virtualServiceId"`

	// Unique VirtualServiceRouteTable identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the life cycle state given.
	LifecycleState VirtualServiceRouteTableLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVirtualServiceRouteTablesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVirtualServiceRouteTablesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVirtualServiceRouteTablesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVirtualServiceRouteTablesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVirtualServiceRouteTablesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListVirtualServiceRouteTablesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVirtualServiceRouteTablesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVirtualServiceRouteTablesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVirtualServiceRouteTablesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVirtualServiceRouteTableLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetVirtualServiceRouteTableLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVirtualServiceRouteTablesResponse wrapper for the ListVirtualServiceRouteTables operation
type ListVirtualServiceRouteTablesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of VirtualServiceRouteTableCollection instances
	VirtualServiceRouteTableCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVirtualServiceRouteTablesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVirtualServiceRouteTablesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVirtualServiceRouteTablesSortOrderEnum Enum with underlying type: string
type ListVirtualServiceRouteTablesSortOrderEnum string

// Set of constants representing the allowable values for ListVirtualServiceRouteTablesSortOrderEnum
const (
	ListVirtualServiceRouteTablesSortOrderAsc  ListVirtualServiceRouteTablesSortOrderEnum = "ASC"
	ListVirtualServiceRouteTablesSortOrderDesc ListVirtualServiceRouteTablesSortOrderEnum = "DESC"
)

var mappingListVirtualServiceRouteTablesSortOrderEnum = map[string]ListVirtualServiceRouteTablesSortOrderEnum{
	"ASC":  ListVirtualServiceRouteTablesSortOrderAsc,
	"DESC": ListVirtualServiceRouteTablesSortOrderDesc,
}

var mappingListVirtualServiceRouteTablesSortOrderEnumLowerCase = map[string]ListVirtualServiceRouteTablesSortOrderEnum{
	"asc":  ListVirtualServiceRouteTablesSortOrderAsc,
	"desc": ListVirtualServiceRouteTablesSortOrderDesc,
}

// GetListVirtualServiceRouteTablesSortOrderEnumValues Enumerates the set of values for ListVirtualServiceRouteTablesSortOrderEnum
func GetListVirtualServiceRouteTablesSortOrderEnumValues() []ListVirtualServiceRouteTablesSortOrderEnum {
	values := make([]ListVirtualServiceRouteTablesSortOrderEnum, 0)
	for _, v := range mappingListVirtualServiceRouteTablesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVirtualServiceRouteTablesSortOrderEnumStringValues Enumerates the set of values in String for ListVirtualServiceRouteTablesSortOrderEnum
func GetListVirtualServiceRouteTablesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVirtualServiceRouteTablesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVirtualServiceRouteTablesSortOrderEnum(val string) (ListVirtualServiceRouteTablesSortOrderEnum, bool) {
	enum, ok := mappingListVirtualServiceRouteTablesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVirtualServiceRouteTablesSortByEnum Enum with underlying type: string
type ListVirtualServiceRouteTablesSortByEnum string

// Set of constants representing the allowable values for ListVirtualServiceRouteTablesSortByEnum
const (
	ListVirtualServiceRouteTablesSortById          ListVirtualServiceRouteTablesSortByEnum = "id"
	ListVirtualServiceRouteTablesSortByTimecreated ListVirtualServiceRouteTablesSortByEnum = "timeCreated"
	ListVirtualServiceRouteTablesSortByName        ListVirtualServiceRouteTablesSortByEnum = "name"
)

var mappingListVirtualServiceRouteTablesSortByEnum = map[string]ListVirtualServiceRouteTablesSortByEnum{
	"id":          ListVirtualServiceRouteTablesSortById,
	"timeCreated": ListVirtualServiceRouteTablesSortByTimecreated,
	"name":        ListVirtualServiceRouteTablesSortByName,
}

var mappingListVirtualServiceRouteTablesSortByEnumLowerCase = map[string]ListVirtualServiceRouteTablesSortByEnum{
	"id":          ListVirtualServiceRouteTablesSortById,
	"timecreated": ListVirtualServiceRouteTablesSortByTimecreated,
	"name":        ListVirtualServiceRouteTablesSortByName,
}

// GetListVirtualServiceRouteTablesSortByEnumValues Enumerates the set of values for ListVirtualServiceRouteTablesSortByEnum
func GetListVirtualServiceRouteTablesSortByEnumValues() []ListVirtualServiceRouteTablesSortByEnum {
	values := make([]ListVirtualServiceRouteTablesSortByEnum, 0)
	for _, v := range mappingListVirtualServiceRouteTablesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVirtualServiceRouteTablesSortByEnumStringValues Enumerates the set of values in String for ListVirtualServiceRouteTablesSortByEnum
func GetListVirtualServiceRouteTablesSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeCreated",
		"name",
	}
}

// GetMappingListVirtualServiceRouteTablesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVirtualServiceRouteTablesSortByEnum(val string) (ListVirtualServiceRouteTablesSortByEnum, bool) {
	enum, ok := mappingListVirtualServiceRouteTablesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
