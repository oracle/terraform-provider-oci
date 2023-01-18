// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListIngressGatewayRouteTablesRequest wrapper for the ListIngressGatewayRouteTables operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListIngressGatewayRouteTables.go.html to see an example of how to use ListIngressGatewayRouteTablesRequest.
type ListIngressGatewayRouteTablesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListIngressGatewayRouteTablesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for 'timeCreated' is descending. Default order for 'name' is ascending.
	SortBy ListIngressGatewayRouteTablesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Unique IngressGateway identifier.
	IngressGatewayId *string `mandatory:"false" contributesTo:"query" name:"ingressGatewayId"`

	// Unique IngressGatewayRouteTable identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the life cycle state given.
	LifecycleState IngressGatewayRouteTableLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIngressGatewayRouteTablesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIngressGatewayRouteTablesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIngressGatewayRouteTablesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIngressGatewayRouteTablesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListIngressGatewayRouteTablesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListIngressGatewayRouteTablesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListIngressGatewayRouteTablesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIngressGatewayRouteTablesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListIngressGatewayRouteTablesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIngressGatewayRouteTableLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetIngressGatewayRouteTableLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIngressGatewayRouteTablesResponse wrapper for the ListIngressGatewayRouteTables operation
type ListIngressGatewayRouteTablesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of IngressGatewayRouteTableCollection instances
	IngressGatewayRouteTableCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListIngressGatewayRouteTablesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIngressGatewayRouteTablesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIngressGatewayRouteTablesSortOrderEnum Enum with underlying type: string
type ListIngressGatewayRouteTablesSortOrderEnum string

// Set of constants representing the allowable values for ListIngressGatewayRouteTablesSortOrderEnum
const (
	ListIngressGatewayRouteTablesSortOrderAsc  ListIngressGatewayRouteTablesSortOrderEnum = "ASC"
	ListIngressGatewayRouteTablesSortOrderDesc ListIngressGatewayRouteTablesSortOrderEnum = "DESC"
)

var mappingListIngressGatewayRouteTablesSortOrderEnum = map[string]ListIngressGatewayRouteTablesSortOrderEnum{
	"ASC":  ListIngressGatewayRouteTablesSortOrderAsc,
	"DESC": ListIngressGatewayRouteTablesSortOrderDesc,
}

var mappingListIngressGatewayRouteTablesSortOrderEnumLowerCase = map[string]ListIngressGatewayRouteTablesSortOrderEnum{
	"asc":  ListIngressGatewayRouteTablesSortOrderAsc,
	"desc": ListIngressGatewayRouteTablesSortOrderDesc,
}

// GetListIngressGatewayRouteTablesSortOrderEnumValues Enumerates the set of values for ListIngressGatewayRouteTablesSortOrderEnum
func GetListIngressGatewayRouteTablesSortOrderEnumValues() []ListIngressGatewayRouteTablesSortOrderEnum {
	values := make([]ListIngressGatewayRouteTablesSortOrderEnum, 0)
	for _, v := range mappingListIngressGatewayRouteTablesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIngressGatewayRouteTablesSortOrderEnumStringValues Enumerates the set of values in String for ListIngressGatewayRouteTablesSortOrderEnum
func GetListIngressGatewayRouteTablesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIngressGatewayRouteTablesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIngressGatewayRouteTablesSortOrderEnum(val string) (ListIngressGatewayRouteTablesSortOrderEnum, bool) {
	enum, ok := mappingListIngressGatewayRouteTablesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIngressGatewayRouteTablesSortByEnum Enum with underlying type: string
type ListIngressGatewayRouteTablesSortByEnum string

// Set of constants representing the allowable values for ListIngressGatewayRouteTablesSortByEnum
const (
	ListIngressGatewayRouteTablesSortById          ListIngressGatewayRouteTablesSortByEnum = "id"
	ListIngressGatewayRouteTablesSortByTimecreated ListIngressGatewayRouteTablesSortByEnum = "timeCreated"
	ListIngressGatewayRouteTablesSortByName        ListIngressGatewayRouteTablesSortByEnum = "name"
)

var mappingListIngressGatewayRouteTablesSortByEnum = map[string]ListIngressGatewayRouteTablesSortByEnum{
	"id":          ListIngressGatewayRouteTablesSortById,
	"timeCreated": ListIngressGatewayRouteTablesSortByTimecreated,
	"name":        ListIngressGatewayRouteTablesSortByName,
}

var mappingListIngressGatewayRouteTablesSortByEnumLowerCase = map[string]ListIngressGatewayRouteTablesSortByEnum{
	"id":          ListIngressGatewayRouteTablesSortById,
	"timecreated": ListIngressGatewayRouteTablesSortByTimecreated,
	"name":        ListIngressGatewayRouteTablesSortByName,
}

// GetListIngressGatewayRouteTablesSortByEnumValues Enumerates the set of values for ListIngressGatewayRouteTablesSortByEnum
func GetListIngressGatewayRouteTablesSortByEnumValues() []ListIngressGatewayRouteTablesSortByEnum {
	values := make([]ListIngressGatewayRouteTablesSortByEnum, 0)
	for _, v := range mappingListIngressGatewayRouteTablesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListIngressGatewayRouteTablesSortByEnumStringValues Enumerates the set of values in String for ListIngressGatewayRouteTablesSortByEnum
func GetListIngressGatewayRouteTablesSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeCreated",
		"name",
	}
}

// GetMappingListIngressGatewayRouteTablesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIngressGatewayRouteTablesSortByEnum(val string) (ListIngressGatewayRouteTablesSortByEnum, bool) {
	enum, ok := mappingListIngressGatewayRouteTablesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
