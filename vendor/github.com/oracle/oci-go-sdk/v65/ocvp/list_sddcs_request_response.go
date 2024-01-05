// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSddcsRequest wrapper for the ListSddcs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ListSddcs.go.html to see an example of how to use ListSddcsRequest.
type ListSddcsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain that the Compute instances are running in.
	// Example: `Uocm:PHX-AD-1`
	ComputeAvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"computeAvailabilityDomain"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

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
	SortOrder ListSddcsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListSddcsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request. If you need to contact Oracle about a particular
	// request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The lifecycle state of the resource.
	LifecycleState ListSddcsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSddcsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSddcsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSddcsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSddcsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSddcsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSddcsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSddcsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSddcsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSddcsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSddcsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSddcsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSddcsResponse wrapper for the ListSddcs operation
type ListSddcsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SddcCollection instances
	SddcCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSddcsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSddcsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSddcsSortOrderEnum Enum with underlying type: string
type ListSddcsSortOrderEnum string

// Set of constants representing the allowable values for ListSddcsSortOrderEnum
const (
	ListSddcsSortOrderAsc  ListSddcsSortOrderEnum = "ASC"
	ListSddcsSortOrderDesc ListSddcsSortOrderEnum = "DESC"
)

var mappingListSddcsSortOrderEnum = map[string]ListSddcsSortOrderEnum{
	"ASC":  ListSddcsSortOrderAsc,
	"DESC": ListSddcsSortOrderDesc,
}

var mappingListSddcsSortOrderEnumLowerCase = map[string]ListSddcsSortOrderEnum{
	"asc":  ListSddcsSortOrderAsc,
	"desc": ListSddcsSortOrderDesc,
}

// GetListSddcsSortOrderEnumValues Enumerates the set of values for ListSddcsSortOrderEnum
func GetListSddcsSortOrderEnumValues() []ListSddcsSortOrderEnum {
	values := make([]ListSddcsSortOrderEnum, 0)
	for _, v := range mappingListSddcsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSddcsSortOrderEnumStringValues Enumerates the set of values in String for ListSddcsSortOrderEnum
func GetListSddcsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSddcsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSddcsSortOrderEnum(val string) (ListSddcsSortOrderEnum, bool) {
	enum, ok := mappingListSddcsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSddcsSortByEnum Enum with underlying type: string
type ListSddcsSortByEnum string

// Set of constants representing the allowable values for ListSddcsSortByEnum
const (
	ListSddcsSortByTimecreated ListSddcsSortByEnum = "timeCreated"
	ListSddcsSortByDisplayname ListSddcsSortByEnum = "displayName"
)

var mappingListSddcsSortByEnum = map[string]ListSddcsSortByEnum{
	"timeCreated": ListSddcsSortByTimecreated,
	"displayName": ListSddcsSortByDisplayname,
}

var mappingListSddcsSortByEnumLowerCase = map[string]ListSddcsSortByEnum{
	"timecreated": ListSddcsSortByTimecreated,
	"displayname": ListSddcsSortByDisplayname,
}

// GetListSddcsSortByEnumValues Enumerates the set of values for ListSddcsSortByEnum
func GetListSddcsSortByEnumValues() []ListSddcsSortByEnum {
	values := make([]ListSddcsSortByEnum, 0)
	for _, v := range mappingListSddcsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSddcsSortByEnumStringValues Enumerates the set of values in String for ListSddcsSortByEnum
func GetListSddcsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSddcsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSddcsSortByEnum(val string) (ListSddcsSortByEnum, bool) {
	enum, ok := mappingListSddcsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSddcsLifecycleStateEnum Enum with underlying type: string
type ListSddcsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSddcsLifecycleStateEnum
const (
	ListSddcsLifecycleStateCreating ListSddcsLifecycleStateEnum = "CREATING"
	ListSddcsLifecycleStateUpdating ListSddcsLifecycleStateEnum = "UPDATING"
	ListSddcsLifecycleStateActive   ListSddcsLifecycleStateEnum = "ACTIVE"
	ListSddcsLifecycleStateDeleting ListSddcsLifecycleStateEnum = "DELETING"
	ListSddcsLifecycleStateDeleted  ListSddcsLifecycleStateEnum = "DELETED"
	ListSddcsLifecycleStateFailed   ListSddcsLifecycleStateEnum = "FAILED"
)

var mappingListSddcsLifecycleStateEnum = map[string]ListSddcsLifecycleStateEnum{
	"CREATING": ListSddcsLifecycleStateCreating,
	"UPDATING": ListSddcsLifecycleStateUpdating,
	"ACTIVE":   ListSddcsLifecycleStateActive,
	"DELETING": ListSddcsLifecycleStateDeleting,
	"DELETED":  ListSddcsLifecycleStateDeleted,
	"FAILED":   ListSddcsLifecycleStateFailed,
}

var mappingListSddcsLifecycleStateEnumLowerCase = map[string]ListSddcsLifecycleStateEnum{
	"creating": ListSddcsLifecycleStateCreating,
	"updating": ListSddcsLifecycleStateUpdating,
	"active":   ListSddcsLifecycleStateActive,
	"deleting": ListSddcsLifecycleStateDeleting,
	"deleted":  ListSddcsLifecycleStateDeleted,
	"failed":   ListSddcsLifecycleStateFailed,
}

// GetListSddcsLifecycleStateEnumValues Enumerates the set of values for ListSddcsLifecycleStateEnum
func GetListSddcsLifecycleStateEnumValues() []ListSddcsLifecycleStateEnum {
	values := make([]ListSddcsLifecycleStateEnum, 0)
	for _, v := range mappingListSddcsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSddcsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSddcsLifecycleStateEnum
func GetListSddcsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListSddcsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSddcsLifecycleStateEnum(val string) (ListSddcsLifecycleStateEnum, bool) {
	enum, ok := mappingListSddcsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
