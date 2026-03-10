// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListByolAllocationsRequest wrapper for the ListByolAllocations operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ListByolAllocations.go.html to see an example of how to use ListByolAllocationsRequest.
type ListByolAllocationsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL Allocation.
	ByolAllocationId *string `mandatory:"false" contributesTo:"query" name:"byolAllocationId"`

	// A filter to return only resources whose lifecycle state matches the given value.
	LifecycleState ByolAllocationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources whose softwareType matches the given value.
	SoftwareType ByolAllocationSoftwareTypeEnum `mandatory:"false" contributesTo:"query" name:"softwareType" omitEmpty:"true"`

	// A filter to return only resources whose availableUnits greater than or equal to the given value.
	AvailableUnitsGreaterThanOrEqualTo *float32 `mandatory:"false" contributesTo:"query" name:"availableUnitsGreaterThanOrEqualTo"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL.
	ByolId *string `mandatory:"false" contributesTo:"query" name:"byolId"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListByolAllocationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// `timeTermStart` is descending. Default order for `timeCreated` is descending.
	// Default order for `displayName` is ascending. The `displayName`
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListByolAllocationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request. If you need to contact Oracle about a particular
	// request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListByolAllocationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListByolAllocationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListByolAllocationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListByolAllocationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListByolAllocationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingByolAllocationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetByolAllocationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingByolAllocationSoftwareTypeEnum(string(request.SoftwareType)); !ok && request.SoftwareType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareType: %s. Supported values are: %s.", request.SoftwareType, strings.Join(GetByolAllocationSoftwareTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListByolAllocationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListByolAllocationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListByolAllocationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListByolAllocationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListByolAllocationsResponse wrapper for the ListByolAllocations operation
type ListByolAllocationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ByolAllocationCollection instances
	ByolAllocationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListByolAllocationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListByolAllocationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListByolAllocationsSortOrderEnum Enum with underlying type: string
type ListByolAllocationsSortOrderEnum string

// Set of constants representing the allowable values for ListByolAllocationsSortOrderEnum
const (
	ListByolAllocationsSortOrderAsc  ListByolAllocationsSortOrderEnum = "ASC"
	ListByolAllocationsSortOrderDesc ListByolAllocationsSortOrderEnum = "DESC"
)

var mappingListByolAllocationsSortOrderEnum = map[string]ListByolAllocationsSortOrderEnum{
	"ASC":  ListByolAllocationsSortOrderAsc,
	"DESC": ListByolAllocationsSortOrderDesc,
}

var mappingListByolAllocationsSortOrderEnumLowerCase = map[string]ListByolAllocationsSortOrderEnum{
	"asc":  ListByolAllocationsSortOrderAsc,
	"desc": ListByolAllocationsSortOrderDesc,
}

// GetListByolAllocationsSortOrderEnumValues Enumerates the set of values for ListByolAllocationsSortOrderEnum
func GetListByolAllocationsSortOrderEnumValues() []ListByolAllocationsSortOrderEnum {
	values := make([]ListByolAllocationsSortOrderEnum, 0)
	for _, v := range mappingListByolAllocationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListByolAllocationsSortOrderEnumStringValues Enumerates the set of values in String for ListByolAllocationsSortOrderEnum
func GetListByolAllocationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListByolAllocationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListByolAllocationsSortOrderEnum(val string) (ListByolAllocationsSortOrderEnum, bool) {
	enum, ok := mappingListByolAllocationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListByolAllocationsSortByEnum Enum with underlying type: string
type ListByolAllocationsSortByEnum string

// Set of constants representing the allowable values for ListByolAllocationsSortByEnum
const (
	ListByolAllocationsSortByDisplayname   ListByolAllocationsSortByEnum = "displayName"
	ListByolAllocationsSortByTimecreated   ListByolAllocationsSortByEnum = "timeCreated"
	ListByolAllocationsSortByTimetermstart ListByolAllocationsSortByEnum = "timeTermStart"
)

var mappingListByolAllocationsSortByEnum = map[string]ListByolAllocationsSortByEnum{
	"displayName":   ListByolAllocationsSortByDisplayname,
	"timeCreated":   ListByolAllocationsSortByTimecreated,
	"timeTermStart": ListByolAllocationsSortByTimetermstart,
}

var mappingListByolAllocationsSortByEnumLowerCase = map[string]ListByolAllocationsSortByEnum{
	"displayname":   ListByolAllocationsSortByDisplayname,
	"timecreated":   ListByolAllocationsSortByTimecreated,
	"timetermstart": ListByolAllocationsSortByTimetermstart,
}

// GetListByolAllocationsSortByEnumValues Enumerates the set of values for ListByolAllocationsSortByEnum
func GetListByolAllocationsSortByEnumValues() []ListByolAllocationsSortByEnum {
	values := make([]ListByolAllocationsSortByEnum, 0)
	for _, v := range mappingListByolAllocationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListByolAllocationsSortByEnumStringValues Enumerates the set of values in String for ListByolAllocationsSortByEnum
func GetListByolAllocationsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
		"timeTermStart",
	}
}

// GetMappingListByolAllocationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListByolAllocationsSortByEnum(val string) (ListByolAllocationsSortByEnum, bool) {
	enum, ok := mappingListByolAllocationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
