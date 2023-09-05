// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListVmwareBillingLinksRequest wrapper for the ListVmwareBillingLinks operation
type ListVmwareBillingLinksRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListVmwareBillingLinksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListVmwareBillingLinksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request. If you need to contact Oracle about a particular
	// request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The lifecycle state of the resource.
	LifecycleState ListVmwareBillingLinksLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match or support the given customer account id.
	VmwareAccountId *string `mandatory:"false" contributesTo:"query" name:"vmwareAccountId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVmwareBillingLinksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVmwareBillingLinksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVmwareBillingLinksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVmwareBillingLinksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVmwareBillingLinksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListVmwareBillingLinksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVmwareBillingLinksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVmwareBillingLinksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVmwareBillingLinksSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVmwareBillingLinksLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListVmwareBillingLinksLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVmwareBillingLinksResponse wrapper for the ListVmwareBillingLinks operation
type ListVmwareBillingLinksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of VmwareBillingLinkCollection instances
	VmwareBillingLinkCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVmwareBillingLinksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVmwareBillingLinksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVmwareBillingLinksSortOrderEnum Enum with underlying type: string
type ListVmwareBillingLinksSortOrderEnum string

// Set of constants representing the allowable values for ListVmwareBillingLinksSortOrderEnum
const (
	ListVmwareBillingLinksSortOrderAsc  ListVmwareBillingLinksSortOrderEnum = "ASC"
	ListVmwareBillingLinksSortOrderDesc ListVmwareBillingLinksSortOrderEnum = "DESC"
)

var mappingListVmwareBillingLinksSortOrderEnum = map[string]ListVmwareBillingLinksSortOrderEnum{
	"ASC":  ListVmwareBillingLinksSortOrderAsc,
	"DESC": ListVmwareBillingLinksSortOrderDesc,
}

var mappingListVmwareBillingLinksSortOrderEnumLowerCase = map[string]ListVmwareBillingLinksSortOrderEnum{
	"asc":  ListVmwareBillingLinksSortOrderAsc,
	"desc": ListVmwareBillingLinksSortOrderDesc,
}

// GetListVmwareBillingLinksSortOrderEnumValues Enumerates the set of values for ListVmwareBillingLinksSortOrderEnum
func GetListVmwareBillingLinksSortOrderEnumValues() []ListVmwareBillingLinksSortOrderEnum {
	values := make([]ListVmwareBillingLinksSortOrderEnum, 0)
	for _, v := range mappingListVmwareBillingLinksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVmwareBillingLinksSortOrderEnumStringValues Enumerates the set of values in String for ListVmwareBillingLinksSortOrderEnum
func GetListVmwareBillingLinksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVmwareBillingLinksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVmwareBillingLinksSortOrderEnum(val string) (ListVmwareBillingLinksSortOrderEnum, bool) {
	enum, ok := mappingListVmwareBillingLinksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVmwareBillingLinksSortByEnum Enum with underlying type: string
type ListVmwareBillingLinksSortByEnum string

// Set of constants representing the allowable values for ListVmwareBillingLinksSortByEnum
const (
	ListVmwareBillingLinksSortByTimecreated ListVmwareBillingLinksSortByEnum = "timeCreated"
	ListVmwareBillingLinksSortByDisplayname ListVmwareBillingLinksSortByEnum = "displayName"
)

var mappingListVmwareBillingLinksSortByEnum = map[string]ListVmwareBillingLinksSortByEnum{
	"timeCreated": ListVmwareBillingLinksSortByTimecreated,
	"displayName": ListVmwareBillingLinksSortByDisplayname,
}

var mappingListVmwareBillingLinksSortByEnumLowerCase = map[string]ListVmwareBillingLinksSortByEnum{
	"timecreated": ListVmwareBillingLinksSortByTimecreated,
	"displayname": ListVmwareBillingLinksSortByDisplayname,
}

// GetListVmwareBillingLinksSortByEnumValues Enumerates the set of values for ListVmwareBillingLinksSortByEnum
func GetListVmwareBillingLinksSortByEnumValues() []ListVmwareBillingLinksSortByEnum {
	values := make([]ListVmwareBillingLinksSortByEnum, 0)
	for _, v := range mappingListVmwareBillingLinksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVmwareBillingLinksSortByEnumStringValues Enumerates the set of values in String for ListVmwareBillingLinksSortByEnum
func GetListVmwareBillingLinksSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListVmwareBillingLinksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVmwareBillingLinksSortByEnum(val string) (ListVmwareBillingLinksSortByEnum, bool) {
	enum, ok := mappingListVmwareBillingLinksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVmwareBillingLinksLifecycleStateEnum Enum with underlying type: string
type ListVmwareBillingLinksLifecycleStateEnum string

// Set of constants representing the allowable values for ListVmwareBillingLinksLifecycleStateEnum
const (
	ListVmwareBillingLinksLifecycleStateCreating ListVmwareBillingLinksLifecycleStateEnum = "CREATING"
	ListVmwareBillingLinksLifecycleStateUpdating ListVmwareBillingLinksLifecycleStateEnum = "UPDATING"
	ListVmwareBillingLinksLifecycleStateActive   ListVmwareBillingLinksLifecycleStateEnum = "ACTIVE"
	ListVmwareBillingLinksLifecycleStateDeleting ListVmwareBillingLinksLifecycleStateEnum = "DELETING"
	ListVmwareBillingLinksLifecycleStateDeleted  ListVmwareBillingLinksLifecycleStateEnum = "DELETED"
	ListVmwareBillingLinksLifecycleStateFailed   ListVmwareBillingLinksLifecycleStateEnum = "FAILED"
)

var mappingListVmwareBillingLinksLifecycleStateEnum = map[string]ListVmwareBillingLinksLifecycleStateEnum{
	"CREATING": ListVmwareBillingLinksLifecycleStateCreating,
	"UPDATING": ListVmwareBillingLinksLifecycleStateUpdating,
	"ACTIVE":   ListVmwareBillingLinksLifecycleStateActive,
	"DELETING": ListVmwareBillingLinksLifecycleStateDeleting,
	"DELETED":  ListVmwareBillingLinksLifecycleStateDeleted,
	"FAILED":   ListVmwareBillingLinksLifecycleStateFailed,
}

var mappingListVmwareBillingLinksLifecycleStateEnumLowerCase = map[string]ListVmwareBillingLinksLifecycleStateEnum{
	"creating": ListVmwareBillingLinksLifecycleStateCreating,
	"updating": ListVmwareBillingLinksLifecycleStateUpdating,
	"active":   ListVmwareBillingLinksLifecycleStateActive,
	"deleting": ListVmwareBillingLinksLifecycleStateDeleting,
	"deleted":  ListVmwareBillingLinksLifecycleStateDeleted,
	"failed":   ListVmwareBillingLinksLifecycleStateFailed,
}

// GetListVmwareBillingLinksLifecycleStateEnumValues Enumerates the set of values for ListVmwareBillingLinksLifecycleStateEnum
func GetListVmwareBillingLinksLifecycleStateEnumValues() []ListVmwareBillingLinksLifecycleStateEnum {
	values := make([]ListVmwareBillingLinksLifecycleStateEnum, 0)
	for _, v := range mappingListVmwareBillingLinksLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListVmwareBillingLinksLifecycleStateEnumStringValues Enumerates the set of values in String for ListVmwareBillingLinksLifecycleStateEnum
func GetListVmwareBillingLinksLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListVmwareBillingLinksLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVmwareBillingLinksLifecycleStateEnum(val string) (ListVmwareBillingLinksLifecycleStateEnum, bool) {
	enum, ok := mappingListVmwareBillingLinksLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
