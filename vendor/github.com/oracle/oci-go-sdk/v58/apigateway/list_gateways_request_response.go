// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apigateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListGatewaysRequest wrapper for the ListGateways operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/ListGateways.go.html to see an example of how to use ListGatewaysRequest.
type ListGatewaysRequest struct {

	// The ocid of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter gateways by the certificate ocid.
	CertificateId *string `mandatory:"false" contributesTo:"query" name:"certificateId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given lifecycle state.
	// Example: `SUCCEEDED`
	LifecycleState GatewayLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'. The default order depends on the sortBy value.
	SortOrder ListGatewaysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for `timeCreated` is descending. Default order for
	// `displayName` is ascending. The `displayName` sort order is case
	// sensitive.
	SortBy ListGatewaysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request id for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGatewaysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGatewaysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGatewaysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGatewaysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGatewaysRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGatewayLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetGatewayLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGatewaysSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGatewaysSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGatewaysSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGatewaysSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGatewaysResponse wrapper for the ListGateways operation
type ListGatewaysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GatewayCollection instances
	GatewayCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request
	// id.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain. For important details about how
	// pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response,
	// additional pages of results were seen previously. For important details
	// about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListGatewaysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGatewaysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGatewaysSortOrderEnum Enum with underlying type: string
type ListGatewaysSortOrderEnum string

// Set of constants representing the allowable values for ListGatewaysSortOrderEnum
const (
	ListGatewaysSortOrderAsc  ListGatewaysSortOrderEnum = "ASC"
	ListGatewaysSortOrderDesc ListGatewaysSortOrderEnum = "DESC"
)

var mappingListGatewaysSortOrderEnum = map[string]ListGatewaysSortOrderEnum{
	"ASC":  ListGatewaysSortOrderAsc,
	"DESC": ListGatewaysSortOrderDesc,
}

// GetListGatewaysSortOrderEnumValues Enumerates the set of values for ListGatewaysSortOrderEnum
func GetListGatewaysSortOrderEnumValues() []ListGatewaysSortOrderEnum {
	values := make([]ListGatewaysSortOrderEnum, 0)
	for _, v := range mappingListGatewaysSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGatewaysSortOrderEnumStringValues Enumerates the set of values in String for ListGatewaysSortOrderEnum
func GetListGatewaysSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGatewaysSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGatewaysSortOrderEnum(val string) (ListGatewaysSortOrderEnum, bool) {
	mappingListGatewaysSortOrderEnumIgnoreCase := make(map[string]ListGatewaysSortOrderEnum)
	for k, v := range mappingListGatewaysSortOrderEnum {
		mappingListGatewaysSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListGatewaysSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListGatewaysSortByEnum Enum with underlying type: string
type ListGatewaysSortByEnum string

// Set of constants representing the allowable values for ListGatewaysSortByEnum
const (
	ListGatewaysSortByTimecreated ListGatewaysSortByEnum = "timeCreated"
	ListGatewaysSortByDisplayname ListGatewaysSortByEnum = "displayName"
)

var mappingListGatewaysSortByEnum = map[string]ListGatewaysSortByEnum{
	"timeCreated": ListGatewaysSortByTimecreated,
	"displayName": ListGatewaysSortByDisplayname,
}

// GetListGatewaysSortByEnumValues Enumerates the set of values for ListGatewaysSortByEnum
func GetListGatewaysSortByEnumValues() []ListGatewaysSortByEnum {
	values := make([]ListGatewaysSortByEnum, 0)
	for _, v := range mappingListGatewaysSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGatewaysSortByEnumStringValues Enumerates the set of values in String for ListGatewaysSortByEnum
func GetListGatewaysSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListGatewaysSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGatewaysSortByEnum(val string) (ListGatewaysSortByEnum, bool) {
	mappingListGatewaysSortByEnumIgnoreCase := make(map[string]ListGatewaysSortByEnum)
	for k, v := range mappingListGatewaysSortByEnum {
		mappingListGatewaysSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListGatewaysSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
