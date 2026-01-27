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

// ListManagementAppliancesRequest wrapper for the ListManagementAppliances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ListManagementAppliances.go.html to see an example of how to use ListManagementAppliancesRequest.
type ListManagementAppliancesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management appliance.
	ManagementApplianceId *string `mandatory:"false" contributesTo:"query" name:"managementApplianceId"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The lifecycle state of the management appliance.
	LifecycleState ListManagementAppliancesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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
	SortOrder ListManagementAppliancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListManagementAppliancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request. If you need to contact Oracle about a particular
	// request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementAppliancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementAppliancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagementAppliancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagementAppliancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagementAppliancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagementAppliancesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListManagementAppliancesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementAppliancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagementAppliancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementAppliancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagementAppliancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagementAppliancesResponse wrapper for the ListManagementAppliances operation
type ListManagementAppliancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagementApplianceCollection instances
	ManagementApplianceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagementAppliancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagementAppliancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagementAppliancesLifecycleStateEnum Enum with underlying type: string
type ListManagementAppliancesLifecycleStateEnum string

// Set of constants representing the allowable values for ListManagementAppliancesLifecycleStateEnum
const (
	ListManagementAppliancesLifecycleStateCreating       ListManagementAppliancesLifecycleStateEnum = "CREATING"
	ListManagementAppliancesLifecycleStateUpdating       ListManagementAppliancesLifecycleStateEnum = "UPDATING"
	ListManagementAppliancesLifecycleStateActive         ListManagementAppliancesLifecycleStateEnum = "ACTIVE"
	ListManagementAppliancesLifecycleStateNeedsAttention ListManagementAppliancesLifecycleStateEnum = "NEEDS_ATTENTION"
	ListManagementAppliancesLifecycleStateDeleting       ListManagementAppliancesLifecycleStateEnum = "DELETING"
	ListManagementAppliancesLifecycleStateDeleted        ListManagementAppliancesLifecycleStateEnum = "DELETED"
	ListManagementAppliancesLifecycleStateFailed         ListManagementAppliancesLifecycleStateEnum = "FAILED"
)

var mappingListManagementAppliancesLifecycleStateEnum = map[string]ListManagementAppliancesLifecycleStateEnum{
	"CREATING":        ListManagementAppliancesLifecycleStateCreating,
	"UPDATING":        ListManagementAppliancesLifecycleStateUpdating,
	"ACTIVE":          ListManagementAppliancesLifecycleStateActive,
	"NEEDS_ATTENTION": ListManagementAppliancesLifecycleStateNeedsAttention,
	"DELETING":        ListManagementAppliancesLifecycleStateDeleting,
	"DELETED":         ListManagementAppliancesLifecycleStateDeleted,
	"FAILED":          ListManagementAppliancesLifecycleStateFailed,
}

var mappingListManagementAppliancesLifecycleStateEnumLowerCase = map[string]ListManagementAppliancesLifecycleStateEnum{
	"creating":        ListManagementAppliancesLifecycleStateCreating,
	"updating":        ListManagementAppliancesLifecycleStateUpdating,
	"active":          ListManagementAppliancesLifecycleStateActive,
	"needs_attention": ListManagementAppliancesLifecycleStateNeedsAttention,
	"deleting":        ListManagementAppliancesLifecycleStateDeleting,
	"deleted":         ListManagementAppliancesLifecycleStateDeleted,
	"failed":          ListManagementAppliancesLifecycleStateFailed,
}

// GetListManagementAppliancesLifecycleStateEnumValues Enumerates the set of values for ListManagementAppliancesLifecycleStateEnum
func GetListManagementAppliancesLifecycleStateEnumValues() []ListManagementAppliancesLifecycleStateEnum {
	values := make([]ListManagementAppliancesLifecycleStateEnum, 0)
	for _, v := range mappingListManagementAppliancesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAppliancesLifecycleStateEnumStringValues Enumerates the set of values in String for ListManagementAppliancesLifecycleStateEnum
func GetListManagementAppliancesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListManagementAppliancesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAppliancesLifecycleStateEnum(val string) (ListManagementAppliancesLifecycleStateEnum, bool) {
	enum, ok := mappingListManagementAppliancesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementAppliancesSortOrderEnum Enum with underlying type: string
type ListManagementAppliancesSortOrderEnum string

// Set of constants representing the allowable values for ListManagementAppliancesSortOrderEnum
const (
	ListManagementAppliancesSortOrderAsc  ListManagementAppliancesSortOrderEnum = "ASC"
	ListManagementAppliancesSortOrderDesc ListManagementAppliancesSortOrderEnum = "DESC"
)

var mappingListManagementAppliancesSortOrderEnum = map[string]ListManagementAppliancesSortOrderEnum{
	"ASC":  ListManagementAppliancesSortOrderAsc,
	"DESC": ListManagementAppliancesSortOrderDesc,
}

var mappingListManagementAppliancesSortOrderEnumLowerCase = map[string]ListManagementAppliancesSortOrderEnum{
	"asc":  ListManagementAppliancesSortOrderAsc,
	"desc": ListManagementAppliancesSortOrderDesc,
}

// GetListManagementAppliancesSortOrderEnumValues Enumerates the set of values for ListManagementAppliancesSortOrderEnum
func GetListManagementAppliancesSortOrderEnumValues() []ListManagementAppliancesSortOrderEnum {
	values := make([]ListManagementAppliancesSortOrderEnum, 0)
	for _, v := range mappingListManagementAppliancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAppliancesSortOrderEnumStringValues Enumerates the set of values in String for ListManagementAppliancesSortOrderEnum
func GetListManagementAppliancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagementAppliancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAppliancesSortOrderEnum(val string) (ListManagementAppliancesSortOrderEnum, bool) {
	enum, ok := mappingListManagementAppliancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementAppliancesSortByEnum Enum with underlying type: string
type ListManagementAppliancesSortByEnum string

// Set of constants representing the allowable values for ListManagementAppliancesSortByEnum
const (
	ListManagementAppliancesSortByTimecreated ListManagementAppliancesSortByEnum = "timeCreated"
	ListManagementAppliancesSortByDisplayname ListManagementAppliancesSortByEnum = "displayName"
)

var mappingListManagementAppliancesSortByEnum = map[string]ListManagementAppliancesSortByEnum{
	"timeCreated": ListManagementAppliancesSortByTimecreated,
	"displayName": ListManagementAppliancesSortByDisplayname,
}

var mappingListManagementAppliancesSortByEnumLowerCase = map[string]ListManagementAppliancesSortByEnum{
	"timecreated": ListManagementAppliancesSortByTimecreated,
	"displayname": ListManagementAppliancesSortByDisplayname,
}

// GetListManagementAppliancesSortByEnumValues Enumerates the set of values for ListManagementAppliancesSortByEnum
func GetListManagementAppliancesSortByEnumValues() []ListManagementAppliancesSortByEnum {
	values := make([]ListManagementAppliancesSortByEnum, 0)
	for _, v := range mappingListManagementAppliancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAppliancesSortByEnumStringValues Enumerates the set of values in String for ListManagementAppliancesSortByEnum
func GetListManagementAppliancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagementAppliancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAppliancesSortByEnum(val string) (ListManagementAppliancesSortByEnum, bool) {
	enum, ok := mappingListManagementAppliancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
