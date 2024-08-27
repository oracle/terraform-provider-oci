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

// ListFileSystemsRequest wrapper for the ListFileSystems operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListFileSystems.go.html to see an example of how to use ListFileSystemsRequest.
type ListFileSystemsRequest struct {

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

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Example: `My resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListFileSystemsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filter results by OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for
	// the resouce type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the snapshot used to create a cloned file system. See Cloning a File System (https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm).
	SourceSnapshotId *string `mandatory:"false" contributesTo:"query" name:"sourceSnapshotId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file system that contains the source snapshot of a cloned file system. See Cloning a File System (https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm).
	ParentFileSystemId *string `mandatory:"false" contributesTo:"query" name:"parentFileSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file system snapshot policy
	// that is associated with the file systems.
	FilesystemSnapshotPolicyId *string `mandatory:"false" contributesTo:"query" name:"filesystemSnapshotPolicyId"`

	// The field to sort by. You can provide either value, but not both.
	// By default, when you sort by time created, results are shown
	// in descending order. When you sort by display name, results are
	// shown in ascending order.
	SortBy ListFileSystemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending. The default order is 'desc'
	// except for numeric values.
	SortOrder ListFileSystemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFileSystemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFileSystemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFileSystemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFileSystemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFileSystemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFileSystemsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListFileSystemsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFileSystemsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFileSystemsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFileSystemsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFileSystemsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFileSystemsResponse wrapper for the ListFileSystems operation
type ListFileSystemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []FileSystemSummary instances
	Items []FileSystemSummary `presentIn:"body"`

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

func (response ListFileSystemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFileSystemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFileSystemsLifecycleStateEnum Enum with underlying type: string
type ListFileSystemsLifecycleStateEnum string

// Set of constants representing the allowable values for ListFileSystemsLifecycleStateEnum
const (
	ListFileSystemsLifecycleStateCreating ListFileSystemsLifecycleStateEnum = "CREATING"
	ListFileSystemsLifecycleStateActive   ListFileSystemsLifecycleStateEnum = "ACTIVE"
	ListFileSystemsLifecycleStateUpdating ListFileSystemsLifecycleStateEnum = "UPDATING"
	ListFileSystemsLifecycleStateDeleting ListFileSystemsLifecycleStateEnum = "DELETING"
	ListFileSystemsLifecycleStateDeleted  ListFileSystemsLifecycleStateEnum = "DELETED"
	ListFileSystemsLifecycleStateFailed   ListFileSystemsLifecycleStateEnum = "FAILED"
)

var mappingListFileSystemsLifecycleStateEnum = map[string]ListFileSystemsLifecycleStateEnum{
	"CREATING": ListFileSystemsLifecycleStateCreating,
	"ACTIVE":   ListFileSystemsLifecycleStateActive,
	"UPDATING": ListFileSystemsLifecycleStateUpdating,
	"DELETING": ListFileSystemsLifecycleStateDeleting,
	"DELETED":  ListFileSystemsLifecycleStateDeleted,
	"FAILED":   ListFileSystemsLifecycleStateFailed,
}

var mappingListFileSystemsLifecycleStateEnumLowerCase = map[string]ListFileSystemsLifecycleStateEnum{
	"creating": ListFileSystemsLifecycleStateCreating,
	"active":   ListFileSystemsLifecycleStateActive,
	"updating": ListFileSystemsLifecycleStateUpdating,
	"deleting": ListFileSystemsLifecycleStateDeleting,
	"deleted":  ListFileSystemsLifecycleStateDeleted,
	"failed":   ListFileSystemsLifecycleStateFailed,
}

// GetListFileSystemsLifecycleStateEnumValues Enumerates the set of values for ListFileSystemsLifecycleStateEnum
func GetListFileSystemsLifecycleStateEnumValues() []ListFileSystemsLifecycleStateEnum {
	values := make([]ListFileSystemsLifecycleStateEnum, 0)
	for _, v := range mappingListFileSystemsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListFileSystemsLifecycleStateEnumStringValues Enumerates the set of values in String for ListFileSystemsLifecycleStateEnum
func GetListFileSystemsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListFileSystemsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFileSystemsLifecycleStateEnum(val string) (ListFileSystemsLifecycleStateEnum, bool) {
	enum, ok := mappingListFileSystemsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFileSystemsSortByEnum Enum with underlying type: string
type ListFileSystemsSortByEnum string

// Set of constants representing the allowable values for ListFileSystemsSortByEnum
const (
	ListFileSystemsSortByTimecreated ListFileSystemsSortByEnum = "TIMECREATED"
	ListFileSystemsSortByDisplayname ListFileSystemsSortByEnum = "DISPLAYNAME"
)

var mappingListFileSystemsSortByEnum = map[string]ListFileSystemsSortByEnum{
	"TIMECREATED": ListFileSystemsSortByTimecreated,
	"DISPLAYNAME": ListFileSystemsSortByDisplayname,
}

var mappingListFileSystemsSortByEnumLowerCase = map[string]ListFileSystemsSortByEnum{
	"timecreated": ListFileSystemsSortByTimecreated,
	"displayname": ListFileSystemsSortByDisplayname,
}

// GetListFileSystemsSortByEnumValues Enumerates the set of values for ListFileSystemsSortByEnum
func GetListFileSystemsSortByEnumValues() []ListFileSystemsSortByEnum {
	values := make([]ListFileSystemsSortByEnum, 0)
	for _, v := range mappingListFileSystemsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFileSystemsSortByEnumStringValues Enumerates the set of values in String for ListFileSystemsSortByEnum
func GetListFileSystemsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListFileSystemsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFileSystemsSortByEnum(val string) (ListFileSystemsSortByEnum, bool) {
	enum, ok := mappingListFileSystemsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFileSystemsSortOrderEnum Enum with underlying type: string
type ListFileSystemsSortOrderEnum string

// Set of constants representing the allowable values for ListFileSystemsSortOrderEnum
const (
	ListFileSystemsSortOrderAsc  ListFileSystemsSortOrderEnum = "ASC"
	ListFileSystemsSortOrderDesc ListFileSystemsSortOrderEnum = "DESC"
)

var mappingListFileSystemsSortOrderEnum = map[string]ListFileSystemsSortOrderEnum{
	"ASC":  ListFileSystemsSortOrderAsc,
	"DESC": ListFileSystemsSortOrderDesc,
}

var mappingListFileSystemsSortOrderEnumLowerCase = map[string]ListFileSystemsSortOrderEnum{
	"asc":  ListFileSystemsSortOrderAsc,
	"desc": ListFileSystemsSortOrderDesc,
}

// GetListFileSystemsSortOrderEnumValues Enumerates the set of values for ListFileSystemsSortOrderEnum
func GetListFileSystemsSortOrderEnumValues() []ListFileSystemsSortOrderEnum {
	values := make([]ListFileSystemsSortOrderEnum, 0)
	for _, v := range mappingListFileSystemsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFileSystemsSortOrderEnumStringValues Enumerates the set of values in String for ListFileSystemsSortOrderEnum
func GetListFileSystemsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFileSystemsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFileSystemsSortOrderEnum(val string) (ListFileSystemsSortOrderEnum, bool) {
	enum, ok := mappingListFileSystemsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
