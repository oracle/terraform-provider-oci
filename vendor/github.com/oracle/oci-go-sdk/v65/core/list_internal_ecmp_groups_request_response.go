// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInternalEcmpGroupsRequest wrapper for the ListInternalEcmpGroups operation
type ListInternalEcmpGroupsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to only return resources that match the given gatewayType `gatewayType`. The `gatewayType` value is the string
	// representation of enum - `SERVICEGATEWAY`, `NATGATEWAY`, `PRIVATEACCESSGATEWAY`.
	GatewayType InternalEcmpGroupGatewayTypeEnum `mandatory:"false" contributesTo:"query" name:"gatewayType" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListInternalEcmpGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListInternalEcmpGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given Internal Ecmp Group lifecycle state.
	// The state value is not case-sensitive.
	LifecycleState InternalEcmpGroupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given Internal Ecmp Group name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInternalEcmpGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInternalEcmpGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInternalEcmpGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInternalEcmpGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInternalEcmpGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternalEcmpGroupGatewayTypeEnum(string(request.GatewayType)); !ok && request.GatewayType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GatewayType: %s. Supported values are: %s.", request.GatewayType, strings.Join(GetInternalEcmpGroupGatewayTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalEcmpGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInternalEcmpGroupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalEcmpGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInternalEcmpGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInternalEcmpGroupLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetInternalEcmpGroupLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInternalEcmpGroupsResponse wrapper for the ListInternalEcmpGroups operation
type ListInternalEcmpGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []InternalEcmpGroup instances
	Items []InternalEcmpGroup `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListInternalEcmpGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInternalEcmpGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInternalEcmpGroupsSortByEnum Enum with underlying type: string
type ListInternalEcmpGroupsSortByEnum string

// Set of constants representing the allowable values for ListInternalEcmpGroupsSortByEnum
const (
	ListInternalEcmpGroupsSortByTimecreated ListInternalEcmpGroupsSortByEnum = "TIMECREATED"
	ListInternalEcmpGroupsSortByDisplayname ListInternalEcmpGroupsSortByEnum = "DISPLAYNAME"
)

var mappingListInternalEcmpGroupsSortByEnum = map[string]ListInternalEcmpGroupsSortByEnum{
	"TIMECREATED": ListInternalEcmpGroupsSortByTimecreated,
	"DISPLAYNAME": ListInternalEcmpGroupsSortByDisplayname,
}

var mappingListInternalEcmpGroupsSortByEnumLowerCase = map[string]ListInternalEcmpGroupsSortByEnum{
	"timecreated": ListInternalEcmpGroupsSortByTimecreated,
	"displayname": ListInternalEcmpGroupsSortByDisplayname,
}

// GetListInternalEcmpGroupsSortByEnumValues Enumerates the set of values for ListInternalEcmpGroupsSortByEnum
func GetListInternalEcmpGroupsSortByEnumValues() []ListInternalEcmpGroupsSortByEnum {
	values := make([]ListInternalEcmpGroupsSortByEnum, 0)
	for _, v := range mappingListInternalEcmpGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalEcmpGroupsSortByEnumStringValues Enumerates the set of values in String for ListInternalEcmpGroupsSortByEnum
func GetListInternalEcmpGroupsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListInternalEcmpGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalEcmpGroupsSortByEnum(val string) (ListInternalEcmpGroupsSortByEnum, bool) {
	enum, ok := mappingListInternalEcmpGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInternalEcmpGroupsSortOrderEnum Enum with underlying type: string
type ListInternalEcmpGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListInternalEcmpGroupsSortOrderEnum
const (
	ListInternalEcmpGroupsSortOrderAsc  ListInternalEcmpGroupsSortOrderEnum = "ASC"
	ListInternalEcmpGroupsSortOrderDesc ListInternalEcmpGroupsSortOrderEnum = "DESC"
)

var mappingListInternalEcmpGroupsSortOrderEnum = map[string]ListInternalEcmpGroupsSortOrderEnum{
	"ASC":  ListInternalEcmpGroupsSortOrderAsc,
	"DESC": ListInternalEcmpGroupsSortOrderDesc,
}

var mappingListInternalEcmpGroupsSortOrderEnumLowerCase = map[string]ListInternalEcmpGroupsSortOrderEnum{
	"asc":  ListInternalEcmpGroupsSortOrderAsc,
	"desc": ListInternalEcmpGroupsSortOrderDesc,
}

// GetListInternalEcmpGroupsSortOrderEnumValues Enumerates the set of values for ListInternalEcmpGroupsSortOrderEnum
func GetListInternalEcmpGroupsSortOrderEnumValues() []ListInternalEcmpGroupsSortOrderEnum {
	values := make([]ListInternalEcmpGroupsSortOrderEnum, 0)
	for _, v := range mappingListInternalEcmpGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalEcmpGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListInternalEcmpGroupsSortOrderEnum
func GetListInternalEcmpGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInternalEcmpGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalEcmpGroupsSortOrderEnum(val string) (ListInternalEcmpGroupsSortOrderEnum, bool) {
	enum, ok := mappingListInternalEcmpGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
