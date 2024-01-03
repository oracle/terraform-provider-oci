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

// ListC3DrgRouteTablesRequest wrapper for the ListC3DrgRouteTables operation
type ListC3DrgRouteTablesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
	DrgId *string `mandatory:"true" contributesTo:"query" name:"drgId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListC3DrgRouteTablesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListC3DrgRouteTablesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the import route distribution.
	ImportDrgRouteDistributionId *string `mandatory:"false" contributesTo:"query" name:"importDrgRouteDistributionId"`

	// A filter that only returns matches for the specified lifecycle
	// state. The value is case insensitive.
	LifecycleState DrgRouteTableLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListC3DrgRouteTablesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListC3DrgRouteTablesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListC3DrgRouteTablesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListC3DrgRouteTablesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListC3DrgRouteTablesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListC3DrgRouteTablesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListC3DrgRouteTablesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListC3DrgRouteTablesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListC3DrgRouteTablesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDrgRouteTableLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDrgRouteTableLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListC3DrgRouteTablesResponse wrapper for the ListC3DrgRouteTables operation
type ListC3DrgRouteTablesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DrgRouteTable instances
	Items []DrgRouteTable `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListC3DrgRouteTablesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListC3DrgRouteTablesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListC3DrgRouteTablesSortByEnum Enum with underlying type: string
type ListC3DrgRouteTablesSortByEnum string

// Set of constants representing the allowable values for ListC3DrgRouteTablesSortByEnum
const (
	ListC3DrgRouteTablesSortByTimecreated ListC3DrgRouteTablesSortByEnum = "TIMECREATED"
	ListC3DrgRouteTablesSortByDisplayname ListC3DrgRouteTablesSortByEnum = "DISPLAYNAME"
)

var mappingListC3DrgRouteTablesSortByEnum = map[string]ListC3DrgRouteTablesSortByEnum{
	"TIMECREATED": ListC3DrgRouteTablesSortByTimecreated,
	"DISPLAYNAME": ListC3DrgRouteTablesSortByDisplayname,
}

var mappingListC3DrgRouteTablesSortByEnumLowerCase = map[string]ListC3DrgRouteTablesSortByEnum{
	"timecreated": ListC3DrgRouteTablesSortByTimecreated,
	"displayname": ListC3DrgRouteTablesSortByDisplayname,
}

// GetListC3DrgRouteTablesSortByEnumValues Enumerates the set of values for ListC3DrgRouteTablesSortByEnum
func GetListC3DrgRouteTablesSortByEnumValues() []ListC3DrgRouteTablesSortByEnum {
	values := make([]ListC3DrgRouteTablesSortByEnum, 0)
	for _, v := range mappingListC3DrgRouteTablesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListC3DrgRouteTablesSortByEnumStringValues Enumerates the set of values in String for ListC3DrgRouteTablesSortByEnum
func GetListC3DrgRouteTablesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListC3DrgRouteTablesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListC3DrgRouteTablesSortByEnum(val string) (ListC3DrgRouteTablesSortByEnum, bool) {
	enum, ok := mappingListC3DrgRouteTablesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListC3DrgRouteTablesSortOrderEnum Enum with underlying type: string
type ListC3DrgRouteTablesSortOrderEnum string

// Set of constants representing the allowable values for ListC3DrgRouteTablesSortOrderEnum
const (
	ListC3DrgRouteTablesSortOrderAsc  ListC3DrgRouteTablesSortOrderEnum = "ASC"
	ListC3DrgRouteTablesSortOrderDesc ListC3DrgRouteTablesSortOrderEnum = "DESC"
)

var mappingListC3DrgRouteTablesSortOrderEnum = map[string]ListC3DrgRouteTablesSortOrderEnum{
	"ASC":  ListC3DrgRouteTablesSortOrderAsc,
	"DESC": ListC3DrgRouteTablesSortOrderDesc,
}

var mappingListC3DrgRouteTablesSortOrderEnumLowerCase = map[string]ListC3DrgRouteTablesSortOrderEnum{
	"asc":  ListC3DrgRouteTablesSortOrderAsc,
	"desc": ListC3DrgRouteTablesSortOrderDesc,
}

// GetListC3DrgRouteTablesSortOrderEnumValues Enumerates the set of values for ListC3DrgRouteTablesSortOrderEnum
func GetListC3DrgRouteTablesSortOrderEnumValues() []ListC3DrgRouteTablesSortOrderEnum {
	values := make([]ListC3DrgRouteTablesSortOrderEnum, 0)
	for _, v := range mappingListC3DrgRouteTablesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListC3DrgRouteTablesSortOrderEnumStringValues Enumerates the set of values in String for ListC3DrgRouteTablesSortOrderEnum
func GetListC3DrgRouteTablesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListC3DrgRouteTablesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListC3DrgRouteTablesSortOrderEnum(val string) (ListC3DrgRouteTablesSortOrderEnum, bool) {
	enum, ok := mappingListC3DrgRouteTablesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
