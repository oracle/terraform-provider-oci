// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMountTargetsRequest wrapper for the ListMountTargets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListMountTargets.go.html to see an example of how to use ListMountTargetsRequest.
type ListMountTargetsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" contributesTo:"query" name:"availabilityDomain"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Example: `My resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the export set.
	ExportSetId *string `mandatory:"false" contributesTo:"query" name:"exportSetId"`

	// Filter results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListMountTargetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filter results by OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for
	// the resouce type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The field to sort by. You can choose either value, but not both.
	// By default, when you sort by time created, results are shown
	// in descending order. When you sort by display name, results are
	// shown in ascending order.
	SortBy ListMountTargetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending. The default order is 'desc'
	// except for numeric values.
	SortOrder ListMountTargetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMountTargetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMountTargetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMountTargetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMountTargetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMountTargetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMountTargetsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListMountTargetsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMountTargetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMountTargetsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMountTargetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMountTargetsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMountTargetsResponse wrapper for the ListMountTargets operation
type ListMountTargetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []MountTargetSummary instances
	Items []MountTargetSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListMountTargetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMountTargetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMountTargetsLifecycleStateEnum Enum with underlying type: string
type ListMountTargetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListMountTargetsLifecycleStateEnum
const (
	ListMountTargetsLifecycleStateCreating ListMountTargetsLifecycleStateEnum = "CREATING"
	ListMountTargetsLifecycleStateActive   ListMountTargetsLifecycleStateEnum = "ACTIVE"
	ListMountTargetsLifecycleStateUpdating ListMountTargetsLifecycleStateEnum = "UPDATING"
	ListMountTargetsLifecycleStateDeleting ListMountTargetsLifecycleStateEnum = "DELETING"
	ListMountTargetsLifecycleStateDeleted  ListMountTargetsLifecycleStateEnum = "DELETED"
	ListMountTargetsLifecycleStateFailed   ListMountTargetsLifecycleStateEnum = "FAILED"
)

var mappingListMountTargetsLifecycleStateEnum = map[string]ListMountTargetsLifecycleStateEnum{
	"CREATING": ListMountTargetsLifecycleStateCreating,
	"ACTIVE":   ListMountTargetsLifecycleStateActive,
	"UPDATING": ListMountTargetsLifecycleStateUpdating,
	"DELETING": ListMountTargetsLifecycleStateDeleting,
	"DELETED":  ListMountTargetsLifecycleStateDeleted,
	"FAILED":   ListMountTargetsLifecycleStateFailed,
}

var mappingListMountTargetsLifecycleStateEnumLowerCase = map[string]ListMountTargetsLifecycleStateEnum{
	"creating": ListMountTargetsLifecycleStateCreating,
	"active":   ListMountTargetsLifecycleStateActive,
	"updating": ListMountTargetsLifecycleStateUpdating,
	"deleting": ListMountTargetsLifecycleStateDeleting,
	"deleted":  ListMountTargetsLifecycleStateDeleted,
	"failed":   ListMountTargetsLifecycleStateFailed,
}

// GetListMountTargetsLifecycleStateEnumValues Enumerates the set of values for ListMountTargetsLifecycleStateEnum
func GetListMountTargetsLifecycleStateEnumValues() []ListMountTargetsLifecycleStateEnum {
	values := make([]ListMountTargetsLifecycleStateEnum, 0)
	for _, v := range mappingListMountTargetsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListMountTargetsLifecycleStateEnumStringValues Enumerates the set of values in String for ListMountTargetsLifecycleStateEnum
func GetListMountTargetsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListMountTargetsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMountTargetsLifecycleStateEnum(val string) (ListMountTargetsLifecycleStateEnum, bool) {
	enum, ok := mappingListMountTargetsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMountTargetsSortByEnum Enum with underlying type: string
type ListMountTargetsSortByEnum string

// Set of constants representing the allowable values for ListMountTargetsSortByEnum
const (
	ListMountTargetsSortByTimecreated ListMountTargetsSortByEnum = "TIMECREATED"
	ListMountTargetsSortByDisplayname ListMountTargetsSortByEnum = "DISPLAYNAME"
)

var mappingListMountTargetsSortByEnum = map[string]ListMountTargetsSortByEnum{
	"TIMECREATED": ListMountTargetsSortByTimecreated,
	"DISPLAYNAME": ListMountTargetsSortByDisplayname,
}

var mappingListMountTargetsSortByEnumLowerCase = map[string]ListMountTargetsSortByEnum{
	"timecreated": ListMountTargetsSortByTimecreated,
	"displayname": ListMountTargetsSortByDisplayname,
}

// GetListMountTargetsSortByEnumValues Enumerates the set of values for ListMountTargetsSortByEnum
func GetListMountTargetsSortByEnumValues() []ListMountTargetsSortByEnum {
	values := make([]ListMountTargetsSortByEnum, 0)
	for _, v := range mappingListMountTargetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMountTargetsSortByEnumStringValues Enumerates the set of values in String for ListMountTargetsSortByEnum
func GetListMountTargetsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListMountTargetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMountTargetsSortByEnum(val string) (ListMountTargetsSortByEnum, bool) {
	enum, ok := mappingListMountTargetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMountTargetsSortOrderEnum Enum with underlying type: string
type ListMountTargetsSortOrderEnum string

// Set of constants representing the allowable values for ListMountTargetsSortOrderEnum
const (
	ListMountTargetsSortOrderAsc  ListMountTargetsSortOrderEnum = "ASC"
	ListMountTargetsSortOrderDesc ListMountTargetsSortOrderEnum = "DESC"
)

var mappingListMountTargetsSortOrderEnum = map[string]ListMountTargetsSortOrderEnum{
	"ASC":  ListMountTargetsSortOrderAsc,
	"DESC": ListMountTargetsSortOrderDesc,
}

var mappingListMountTargetsSortOrderEnumLowerCase = map[string]ListMountTargetsSortOrderEnum{
	"asc":  ListMountTargetsSortOrderAsc,
	"desc": ListMountTargetsSortOrderDesc,
}

// GetListMountTargetsSortOrderEnumValues Enumerates the set of values for ListMountTargetsSortOrderEnum
func GetListMountTargetsSortOrderEnumValues() []ListMountTargetsSortOrderEnum {
	values := make([]ListMountTargetsSortOrderEnum, 0)
	for _, v := range mappingListMountTargetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMountTargetsSortOrderEnumStringValues Enumerates the set of values in String for ListMountTargetsSortOrderEnum
func GetListMountTargetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMountTargetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMountTargetsSortOrderEnum(val string) (ListMountTargetsSortOrderEnum, bool) {
	enum, ok := mappingListMountTargetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
