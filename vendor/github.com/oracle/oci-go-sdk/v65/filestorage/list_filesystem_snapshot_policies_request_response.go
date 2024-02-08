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

// ListFilesystemSnapshotPoliciesRequest wrapper for the ListFilesystemSnapshotPolicies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListFilesystemSnapshotPolicies.go.html to see an example of how to use ListFilesystemSnapshotPoliciesRequest.
type ListFilesystemSnapshotPoliciesRequest struct {

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

	// Filter results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListFilesystemSnapshotPoliciesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filter results by OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for
	// the resouce type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The field to sort by. You can provide either value, but not both.
	// By default, when you sort by time created, results are shown
	// in descending order. When you sort by displayName, results are
	// shown in ascending alphanumeric order.
	SortBy ListFilesystemSnapshotPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending. The default order is 'desc'
	// except for numeric values.
	SortOrder ListFilesystemSnapshotPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFilesystemSnapshotPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFilesystemSnapshotPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFilesystemSnapshotPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFilesystemSnapshotPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFilesystemSnapshotPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFilesystemSnapshotPoliciesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListFilesystemSnapshotPoliciesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFilesystemSnapshotPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFilesystemSnapshotPoliciesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFilesystemSnapshotPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFilesystemSnapshotPoliciesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFilesystemSnapshotPoliciesResponse wrapper for the ListFilesystemSnapshotPolicies operation
type ListFilesystemSnapshotPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []FilesystemSnapshotPolicySummary instances
	Items []FilesystemSnapshotPolicySummary `presentIn:"body"`

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

func (response ListFilesystemSnapshotPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFilesystemSnapshotPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFilesystemSnapshotPoliciesLifecycleStateEnum Enum with underlying type: string
type ListFilesystemSnapshotPoliciesLifecycleStateEnum string

// Set of constants representing the allowable values for ListFilesystemSnapshotPoliciesLifecycleStateEnum
const (
	ListFilesystemSnapshotPoliciesLifecycleStateCreating ListFilesystemSnapshotPoliciesLifecycleStateEnum = "CREATING"
	ListFilesystemSnapshotPoliciesLifecycleStateActive   ListFilesystemSnapshotPoliciesLifecycleStateEnum = "ACTIVE"
	ListFilesystemSnapshotPoliciesLifecycleStateDeleting ListFilesystemSnapshotPoliciesLifecycleStateEnum = "DELETING"
	ListFilesystemSnapshotPoliciesLifecycleStateDeleted  ListFilesystemSnapshotPoliciesLifecycleStateEnum = "DELETED"
	ListFilesystemSnapshotPoliciesLifecycleStateFailed   ListFilesystemSnapshotPoliciesLifecycleStateEnum = "FAILED"
	ListFilesystemSnapshotPoliciesLifecycleStateInactive ListFilesystemSnapshotPoliciesLifecycleStateEnum = "INACTIVE"
)

var mappingListFilesystemSnapshotPoliciesLifecycleStateEnum = map[string]ListFilesystemSnapshotPoliciesLifecycleStateEnum{
	"CREATING": ListFilesystemSnapshotPoliciesLifecycleStateCreating,
	"ACTIVE":   ListFilesystemSnapshotPoliciesLifecycleStateActive,
	"DELETING": ListFilesystemSnapshotPoliciesLifecycleStateDeleting,
	"DELETED":  ListFilesystemSnapshotPoliciesLifecycleStateDeleted,
	"FAILED":   ListFilesystemSnapshotPoliciesLifecycleStateFailed,
	"INACTIVE": ListFilesystemSnapshotPoliciesLifecycleStateInactive,
}

var mappingListFilesystemSnapshotPoliciesLifecycleStateEnumLowerCase = map[string]ListFilesystemSnapshotPoliciesLifecycleStateEnum{
	"creating": ListFilesystemSnapshotPoliciesLifecycleStateCreating,
	"active":   ListFilesystemSnapshotPoliciesLifecycleStateActive,
	"deleting": ListFilesystemSnapshotPoliciesLifecycleStateDeleting,
	"deleted":  ListFilesystemSnapshotPoliciesLifecycleStateDeleted,
	"failed":   ListFilesystemSnapshotPoliciesLifecycleStateFailed,
	"inactive": ListFilesystemSnapshotPoliciesLifecycleStateInactive,
}

// GetListFilesystemSnapshotPoliciesLifecycleStateEnumValues Enumerates the set of values for ListFilesystemSnapshotPoliciesLifecycleStateEnum
func GetListFilesystemSnapshotPoliciesLifecycleStateEnumValues() []ListFilesystemSnapshotPoliciesLifecycleStateEnum {
	values := make([]ListFilesystemSnapshotPoliciesLifecycleStateEnum, 0)
	for _, v := range mappingListFilesystemSnapshotPoliciesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListFilesystemSnapshotPoliciesLifecycleStateEnumStringValues Enumerates the set of values in String for ListFilesystemSnapshotPoliciesLifecycleStateEnum
func GetListFilesystemSnapshotPoliciesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"INACTIVE",
	}
}

// GetMappingListFilesystemSnapshotPoliciesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFilesystemSnapshotPoliciesLifecycleStateEnum(val string) (ListFilesystemSnapshotPoliciesLifecycleStateEnum, bool) {
	enum, ok := mappingListFilesystemSnapshotPoliciesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFilesystemSnapshotPoliciesSortByEnum Enum with underlying type: string
type ListFilesystemSnapshotPoliciesSortByEnum string

// Set of constants representing the allowable values for ListFilesystemSnapshotPoliciesSortByEnum
const (
	ListFilesystemSnapshotPoliciesSortByTimecreated ListFilesystemSnapshotPoliciesSortByEnum = "TIMECREATED"
	ListFilesystemSnapshotPoliciesSortByDisplayname ListFilesystemSnapshotPoliciesSortByEnum = "DISPLAYNAME"
)

var mappingListFilesystemSnapshotPoliciesSortByEnum = map[string]ListFilesystemSnapshotPoliciesSortByEnum{
	"TIMECREATED": ListFilesystemSnapshotPoliciesSortByTimecreated,
	"DISPLAYNAME": ListFilesystemSnapshotPoliciesSortByDisplayname,
}

var mappingListFilesystemSnapshotPoliciesSortByEnumLowerCase = map[string]ListFilesystemSnapshotPoliciesSortByEnum{
	"timecreated": ListFilesystemSnapshotPoliciesSortByTimecreated,
	"displayname": ListFilesystemSnapshotPoliciesSortByDisplayname,
}

// GetListFilesystemSnapshotPoliciesSortByEnumValues Enumerates the set of values for ListFilesystemSnapshotPoliciesSortByEnum
func GetListFilesystemSnapshotPoliciesSortByEnumValues() []ListFilesystemSnapshotPoliciesSortByEnum {
	values := make([]ListFilesystemSnapshotPoliciesSortByEnum, 0)
	for _, v := range mappingListFilesystemSnapshotPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFilesystemSnapshotPoliciesSortByEnumStringValues Enumerates the set of values in String for ListFilesystemSnapshotPoliciesSortByEnum
func GetListFilesystemSnapshotPoliciesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListFilesystemSnapshotPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFilesystemSnapshotPoliciesSortByEnum(val string) (ListFilesystemSnapshotPoliciesSortByEnum, bool) {
	enum, ok := mappingListFilesystemSnapshotPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFilesystemSnapshotPoliciesSortOrderEnum Enum with underlying type: string
type ListFilesystemSnapshotPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListFilesystemSnapshotPoliciesSortOrderEnum
const (
	ListFilesystemSnapshotPoliciesSortOrderAsc  ListFilesystemSnapshotPoliciesSortOrderEnum = "ASC"
	ListFilesystemSnapshotPoliciesSortOrderDesc ListFilesystemSnapshotPoliciesSortOrderEnum = "DESC"
)

var mappingListFilesystemSnapshotPoliciesSortOrderEnum = map[string]ListFilesystemSnapshotPoliciesSortOrderEnum{
	"ASC":  ListFilesystemSnapshotPoliciesSortOrderAsc,
	"DESC": ListFilesystemSnapshotPoliciesSortOrderDesc,
}

var mappingListFilesystemSnapshotPoliciesSortOrderEnumLowerCase = map[string]ListFilesystemSnapshotPoliciesSortOrderEnum{
	"asc":  ListFilesystemSnapshotPoliciesSortOrderAsc,
	"desc": ListFilesystemSnapshotPoliciesSortOrderDesc,
}

// GetListFilesystemSnapshotPoliciesSortOrderEnumValues Enumerates the set of values for ListFilesystemSnapshotPoliciesSortOrderEnum
func GetListFilesystemSnapshotPoliciesSortOrderEnumValues() []ListFilesystemSnapshotPoliciesSortOrderEnum {
	values := make([]ListFilesystemSnapshotPoliciesSortOrderEnum, 0)
	for _, v := range mappingListFilesystemSnapshotPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFilesystemSnapshotPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListFilesystemSnapshotPoliciesSortOrderEnum
func GetListFilesystemSnapshotPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFilesystemSnapshotPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFilesystemSnapshotPoliciesSortOrderEnum(val string) (ListFilesystemSnapshotPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListFilesystemSnapshotPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
