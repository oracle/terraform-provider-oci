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

// ListReplicationTargetsRequest wrapper for the ListReplicationTargets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListReplicationTargets.go.html to see an example of how to use ListReplicationTargetsRequest.
type ListReplicationTargetsRequest struct {

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

	// Filter results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListReplicationTargetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Example: `My resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter results by OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for
	// the resouce type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The field to sort by. You can choose either value, but not both.
	// By default, when you sort by `timeCreated`, results are shown
	// in descending order. When you sort by `displayName`, results are
	// shown in ascending order.
	SortBy ListReplicationTargetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending. The default order is 'desc'
	// except for numeric values.
	SortOrder ListReplicationTargetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReplicationTargetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReplicationTargetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListReplicationTargetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReplicationTargetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListReplicationTargetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListReplicationTargetsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListReplicationTargetsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReplicationTargetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListReplicationTargetsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReplicationTargetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListReplicationTargetsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListReplicationTargetsResponse wrapper for the ListReplicationTargets operation
type ListReplicationTargetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ReplicationTargetSummary instances
	Items []ReplicationTargetSummary `presentIn:"body"`

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

func (response ListReplicationTargetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReplicationTargetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListReplicationTargetsLifecycleStateEnum Enum with underlying type: string
type ListReplicationTargetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListReplicationTargetsLifecycleStateEnum
const (
	ListReplicationTargetsLifecycleStateCreating ListReplicationTargetsLifecycleStateEnum = "CREATING"
	ListReplicationTargetsLifecycleStateActive   ListReplicationTargetsLifecycleStateEnum = "ACTIVE"
	ListReplicationTargetsLifecycleStateUpdating ListReplicationTargetsLifecycleStateEnum = "UPDATING"
	ListReplicationTargetsLifecycleStateDeleting ListReplicationTargetsLifecycleStateEnum = "DELETING"
	ListReplicationTargetsLifecycleStateDeleted  ListReplicationTargetsLifecycleStateEnum = "DELETED"
	ListReplicationTargetsLifecycleStateFailed   ListReplicationTargetsLifecycleStateEnum = "FAILED"
)

var mappingListReplicationTargetsLifecycleStateEnum = map[string]ListReplicationTargetsLifecycleStateEnum{
	"CREATING": ListReplicationTargetsLifecycleStateCreating,
	"ACTIVE":   ListReplicationTargetsLifecycleStateActive,
	"UPDATING": ListReplicationTargetsLifecycleStateUpdating,
	"DELETING": ListReplicationTargetsLifecycleStateDeleting,
	"DELETED":  ListReplicationTargetsLifecycleStateDeleted,
	"FAILED":   ListReplicationTargetsLifecycleStateFailed,
}

var mappingListReplicationTargetsLifecycleStateEnumLowerCase = map[string]ListReplicationTargetsLifecycleStateEnum{
	"creating": ListReplicationTargetsLifecycleStateCreating,
	"active":   ListReplicationTargetsLifecycleStateActive,
	"updating": ListReplicationTargetsLifecycleStateUpdating,
	"deleting": ListReplicationTargetsLifecycleStateDeleting,
	"deleted":  ListReplicationTargetsLifecycleStateDeleted,
	"failed":   ListReplicationTargetsLifecycleStateFailed,
}

// GetListReplicationTargetsLifecycleStateEnumValues Enumerates the set of values for ListReplicationTargetsLifecycleStateEnum
func GetListReplicationTargetsLifecycleStateEnumValues() []ListReplicationTargetsLifecycleStateEnum {
	values := make([]ListReplicationTargetsLifecycleStateEnum, 0)
	for _, v := range mappingListReplicationTargetsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListReplicationTargetsLifecycleStateEnumStringValues Enumerates the set of values in String for ListReplicationTargetsLifecycleStateEnum
func GetListReplicationTargetsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListReplicationTargetsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReplicationTargetsLifecycleStateEnum(val string) (ListReplicationTargetsLifecycleStateEnum, bool) {
	enum, ok := mappingListReplicationTargetsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReplicationTargetsSortByEnum Enum with underlying type: string
type ListReplicationTargetsSortByEnum string

// Set of constants representing the allowable values for ListReplicationTargetsSortByEnum
const (
	ListReplicationTargetsSortByTimecreated ListReplicationTargetsSortByEnum = "timeCreated"
	ListReplicationTargetsSortByDisplayname ListReplicationTargetsSortByEnum = "displayName"
)

var mappingListReplicationTargetsSortByEnum = map[string]ListReplicationTargetsSortByEnum{
	"timeCreated": ListReplicationTargetsSortByTimecreated,
	"displayName": ListReplicationTargetsSortByDisplayname,
}

var mappingListReplicationTargetsSortByEnumLowerCase = map[string]ListReplicationTargetsSortByEnum{
	"timecreated": ListReplicationTargetsSortByTimecreated,
	"displayname": ListReplicationTargetsSortByDisplayname,
}

// GetListReplicationTargetsSortByEnumValues Enumerates the set of values for ListReplicationTargetsSortByEnum
func GetListReplicationTargetsSortByEnumValues() []ListReplicationTargetsSortByEnum {
	values := make([]ListReplicationTargetsSortByEnum, 0)
	for _, v := range mappingListReplicationTargetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListReplicationTargetsSortByEnumStringValues Enumerates the set of values in String for ListReplicationTargetsSortByEnum
func GetListReplicationTargetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListReplicationTargetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReplicationTargetsSortByEnum(val string) (ListReplicationTargetsSortByEnum, bool) {
	enum, ok := mappingListReplicationTargetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReplicationTargetsSortOrderEnum Enum with underlying type: string
type ListReplicationTargetsSortOrderEnum string

// Set of constants representing the allowable values for ListReplicationTargetsSortOrderEnum
const (
	ListReplicationTargetsSortOrderAsc  ListReplicationTargetsSortOrderEnum = "ASC"
	ListReplicationTargetsSortOrderDesc ListReplicationTargetsSortOrderEnum = "DESC"
)

var mappingListReplicationTargetsSortOrderEnum = map[string]ListReplicationTargetsSortOrderEnum{
	"ASC":  ListReplicationTargetsSortOrderAsc,
	"DESC": ListReplicationTargetsSortOrderDesc,
}

var mappingListReplicationTargetsSortOrderEnumLowerCase = map[string]ListReplicationTargetsSortOrderEnum{
	"asc":  ListReplicationTargetsSortOrderAsc,
	"desc": ListReplicationTargetsSortOrderDesc,
}

// GetListReplicationTargetsSortOrderEnumValues Enumerates the set of values for ListReplicationTargetsSortOrderEnum
func GetListReplicationTargetsSortOrderEnumValues() []ListReplicationTargetsSortOrderEnum {
	values := make([]ListReplicationTargetsSortOrderEnum, 0)
	for _, v := range mappingListReplicationTargetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListReplicationTargetsSortOrderEnumStringValues Enumerates the set of values in String for ListReplicationTargetsSortOrderEnum
func GetListReplicationTargetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListReplicationTargetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReplicationTargetsSortOrderEnum(val string) (ListReplicationTargetsSortOrderEnum, bool) {
	enum, ok := mappingListReplicationTargetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
