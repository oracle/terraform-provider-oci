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

// ListIngressGatewaysRequest wrapper for the ListIngressGateways operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListIngressGateways.go.html to see an example of how to use ListIngressGatewaysRequest.
type ListIngressGatewaysRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListIngressGatewaysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for 'timeCreated' is descending. Default order for 'name' is ascending.
	SortBy ListIngressGatewaysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Unique Mesh identifier.
	MeshId *string `mandatory:"false" contributesTo:"query" name:"meshId"`

	// Unique IngressGateway identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the life cycle state given.
	LifecycleState IngressGatewayLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIngressGatewaysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIngressGatewaysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIngressGatewaysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIngressGatewaysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListIngressGatewaysRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListIngressGatewaysSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListIngressGatewaysSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIngressGatewaysSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListIngressGatewaysSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIngressGatewayLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetIngressGatewayLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIngressGatewaysResponse wrapper for the ListIngressGateways operation
type ListIngressGatewaysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of IngressGatewayCollection instances
	IngressGatewayCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListIngressGatewaysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIngressGatewaysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIngressGatewaysSortOrderEnum Enum with underlying type: string
type ListIngressGatewaysSortOrderEnum string

// Set of constants representing the allowable values for ListIngressGatewaysSortOrderEnum
const (
	ListIngressGatewaysSortOrderAsc  ListIngressGatewaysSortOrderEnum = "ASC"
	ListIngressGatewaysSortOrderDesc ListIngressGatewaysSortOrderEnum = "DESC"
)

var mappingListIngressGatewaysSortOrderEnum = map[string]ListIngressGatewaysSortOrderEnum{
	"ASC":  ListIngressGatewaysSortOrderAsc,
	"DESC": ListIngressGatewaysSortOrderDesc,
}

var mappingListIngressGatewaysSortOrderEnumLowerCase = map[string]ListIngressGatewaysSortOrderEnum{
	"asc":  ListIngressGatewaysSortOrderAsc,
	"desc": ListIngressGatewaysSortOrderDesc,
}

// GetListIngressGatewaysSortOrderEnumValues Enumerates the set of values for ListIngressGatewaysSortOrderEnum
func GetListIngressGatewaysSortOrderEnumValues() []ListIngressGatewaysSortOrderEnum {
	values := make([]ListIngressGatewaysSortOrderEnum, 0)
	for _, v := range mappingListIngressGatewaysSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIngressGatewaysSortOrderEnumStringValues Enumerates the set of values in String for ListIngressGatewaysSortOrderEnum
func GetListIngressGatewaysSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIngressGatewaysSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIngressGatewaysSortOrderEnum(val string) (ListIngressGatewaysSortOrderEnum, bool) {
	enum, ok := mappingListIngressGatewaysSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIngressGatewaysSortByEnum Enum with underlying type: string
type ListIngressGatewaysSortByEnum string

// Set of constants representing the allowable values for ListIngressGatewaysSortByEnum
const (
	ListIngressGatewaysSortById          ListIngressGatewaysSortByEnum = "id"
	ListIngressGatewaysSortByTimecreated ListIngressGatewaysSortByEnum = "timeCreated"
	ListIngressGatewaysSortByName        ListIngressGatewaysSortByEnum = "name"
)

var mappingListIngressGatewaysSortByEnum = map[string]ListIngressGatewaysSortByEnum{
	"id":          ListIngressGatewaysSortById,
	"timeCreated": ListIngressGatewaysSortByTimecreated,
	"name":        ListIngressGatewaysSortByName,
}

var mappingListIngressGatewaysSortByEnumLowerCase = map[string]ListIngressGatewaysSortByEnum{
	"id":          ListIngressGatewaysSortById,
	"timecreated": ListIngressGatewaysSortByTimecreated,
	"name":        ListIngressGatewaysSortByName,
}

// GetListIngressGatewaysSortByEnumValues Enumerates the set of values for ListIngressGatewaysSortByEnum
func GetListIngressGatewaysSortByEnumValues() []ListIngressGatewaysSortByEnum {
	values := make([]ListIngressGatewaysSortByEnum, 0)
	for _, v := range mappingListIngressGatewaysSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListIngressGatewaysSortByEnumStringValues Enumerates the set of values in String for ListIngressGatewaysSortByEnum
func GetListIngressGatewaysSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeCreated",
		"name",
	}
}

// GetMappingListIngressGatewaysSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIngressGatewaysSortByEnum(val string) (ListIngressGatewaysSortByEnum, bool) {
	enum, ok := mappingListIngressGatewaysSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
