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

// ListReplicationsRequest wrapper for the ListReplications operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListReplications.go.html to see an example of how to use ListReplicationsRequest.
type ListReplicationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" contributesTo:"query" name:"availabilityDomain"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 4096 is the maximum.
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
	LifecycleState ListReplicationsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Example: `My resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter results by OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for
	// the resouce type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The field to sort by. You can choose either value, but not both.
	// By default, when you sort by time created, results are shown
	// in descending order. When you sort by display name, results are
	// shown in ascending order.
	SortBy ListReplicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending. The default order is 'desc'
	// except for numeric values.
	SortOrder ListReplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the source file system.
	FileSystemId *string `mandatory:"false" contributesTo:"query" name:"fileSystemId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReplicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReplicationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListReplicationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReplicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListReplicationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListReplicationsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListReplicationsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReplicationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListReplicationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReplicationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListReplicationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListReplicationsResponse wrapper for the ListReplications operation
type ListReplicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ReplicationSummary instances
	Items []ReplicationSummary `presentIn:"body"`

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

func (response ListReplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListReplicationsLifecycleStateEnum Enum with underlying type: string
type ListReplicationsLifecycleStateEnum string

// Set of constants representing the allowable values for ListReplicationsLifecycleStateEnum
const (
	ListReplicationsLifecycleStateCreating ListReplicationsLifecycleStateEnum = "CREATING"
	ListReplicationsLifecycleStateActive   ListReplicationsLifecycleStateEnum = "ACTIVE"
	ListReplicationsLifecycleStateUpdating ListReplicationsLifecycleStateEnum = "UPDATING"
	ListReplicationsLifecycleStateDeleting ListReplicationsLifecycleStateEnum = "DELETING"
	ListReplicationsLifecycleStateDeleted  ListReplicationsLifecycleStateEnum = "DELETED"
	ListReplicationsLifecycleStateFailed   ListReplicationsLifecycleStateEnum = "FAILED"
)

var mappingListReplicationsLifecycleStateEnum = map[string]ListReplicationsLifecycleStateEnum{
	"CREATING": ListReplicationsLifecycleStateCreating,
	"ACTIVE":   ListReplicationsLifecycleStateActive,
	"UPDATING": ListReplicationsLifecycleStateUpdating,
	"DELETING": ListReplicationsLifecycleStateDeleting,
	"DELETED":  ListReplicationsLifecycleStateDeleted,
	"FAILED":   ListReplicationsLifecycleStateFailed,
}

var mappingListReplicationsLifecycleStateEnumLowerCase = map[string]ListReplicationsLifecycleStateEnum{
	"creating": ListReplicationsLifecycleStateCreating,
	"active":   ListReplicationsLifecycleStateActive,
	"updating": ListReplicationsLifecycleStateUpdating,
	"deleting": ListReplicationsLifecycleStateDeleting,
	"deleted":  ListReplicationsLifecycleStateDeleted,
	"failed":   ListReplicationsLifecycleStateFailed,
}

// GetListReplicationsLifecycleStateEnumValues Enumerates the set of values for ListReplicationsLifecycleStateEnum
func GetListReplicationsLifecycleStateEnumValues() []ListReplicationsLifecycleStateEnum {
	values := make([]ListReplicationsLifecycleStateEnum, 0)
	for _, v := range mappingListReplicationsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListReplicationsLifecycleStateEnumStringValues Enumerates the set of values in String for ListReplicationsLifecycleStateEnum
func GetListReplicationsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListReplicationsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReplicationsLifecycleStateEnum(val string) (ListReplicationsLifecycleStateEnum, bool) {
	enum, ok := mappingListReplicationsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReplicationsSortByEnum Enum with underlying type: string
type ListReplicationsSortByEnum string

// Set of constants representing the allowable values for ListReplicationsSortByEnum
const (
	ListReplicationsSortByTimecreated ListReplicationsSortByEnum = "timeCreated"
	ListReplicationsSortByDisplayname ListReplicationsSortByEnum = "displayName"
)

var mappingListReplicationsSortByEnum = map[string]ListReplicationsSortByEnum{
	"timeCreated": ListReplicationsSortByTimecreated,
	"displayName": ListReplicationsSortByDisplayname,
}

var mappingListReplicationsSortByEnumLowerCase = map[string]ListReplicationsSortByEnum{
	"timecreated": ListReplicationsSortByTimecreated,
	"displayname": ListReplicationsSortByDisplayname,
}

// GetListReplicationsSortByEnumValues Enumerates the set of values for ListReplicationsSortByEnum
func GetListReplicationsSortByEnumValues() []ListReplicationsSortByEnum {
	values := make([]ListReplicationsSortByEnum, 0)
	for _, v := range mappingListReplicationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListReplicationsSortByEnumStringValues Enumerates the set of values in String for ListReplicationsSortByEnum
func GetListReplicationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListReplicationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReplicationsSortByEnum(val string) (ListReplicationsSortByEnum, bool) {
	enum, ok := mappingListReplicationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReplicationsSortOrderEnum Enum with underlying type: string
type ListReplicationsSortOrderEnum string

// Set of constants representing the allowable values for ListReplicationsSortOrderEnum
const (
	ListReplicationsSortOrderAsc  ListReplicationsSortOrderEnum = "ASC"
	ListReplicationsSortOrderDesc ListReplicationsSortOrderEnum = "DESC"
)

var mappingListReplicationsSortOrderEnum = map[string]ListReplicationsSortOrderEnum{
	"ASC":  ListReplicationsSortOrderAsc,
	"DESC": ListReplicationsSortOrderDesc,
}

var mappingListReplicationsSortOrderEnumLowerCase = map[string]ListReplicationsSortOrderEnum{
	"asc":  ListReplicationsSortOrderAsc,
	"desc": ListReplicationsSortOrderDesc,
}

// GetListReplicationsSortOrderEnumValues Enumerates the set of values for ListReplicationsSortOrderEnum
func GetListReplicationsSortOrderEnumValues() []ListReplicationsSortOrderEnum {
	values := make([]ListReplicationsSortOrderEnum, 0)
	for _, v := range mappingListReplicationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListReplicationsSortOrderEnumStringValues Enumerates the set of values in String for ListReplicationsSortOrderEnum
func GetListReplicationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListReplicationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReplicationsSortOrderEnum(val string) (ListReplicationsSortOrderEnum, bool) {
	enum, ok := mappingListReplicationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
