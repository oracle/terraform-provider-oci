// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSnapshotsRequest wrapper for the ListSnapshots operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListSnapshots.go.html to see an example of how to use ListSnapshotsRequest.
type ListSnapshotsRequest struct {

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 100 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `100`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Filter results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListSnapshotsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filter results by OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for
	// the resouce type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system snapshot policy
	// that is used to create the snapshots.
	FilesystemSnapshotPolicyId *string `mandatory:"false" contributesTo:"query" name:"filesystemSnapshotPolicyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system.
	FileSystemId *string `mandatory:"false" contributesTo:"query" name:"fileSystemId"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending. The default order is 'desc'
	// except for numeric values.
	SortOrder ListSnapshotsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSnapshotsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSnapshotsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSnapshotsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSnapshotsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSnapshotsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSnapshotsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSnapshotsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSnapshotsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSnapshotsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSnapshotsResponse wrapper for the ListSnapshots operation
type ListSnapshotsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SnapshotSummary instances
	Items []SnapshotSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSnapshotsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSnapshotsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSnapshotsLifecycleStateEnum Enum with underlying type: string
type ListSnapshotsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSnapshotsLifecycleStateEnum
const (
	ListSnapshotsLifecycleStateCreating ListSnapshotsLifecycleStateEnum = "CREATING"
	ListSnapshotsLifecycleStateActive   ListSnapshotsLifecycleStateEnum = "ACTIVE"
	ListSnapshotsLifecycleStateUpdating ListSnapshotsLifecycleStateEnum = "UPDATING"
	ListSnapshotsLifecycleStateDeleting ListSnapshotsLifecycleStateEnum = "DELETING"
	ListSnapshotsLifecycleStateDeleted  ListSnapshotsLifecycleStateEnum = "DELETED"
	ListSnapshotsLifecycleStateFailed   ListSnapshotsLifecycleStateEnum = "FAILED"
)

var mappingListSnapshotsLifecycleStateEnum = map[string]ListSnapshotsLifecycleStateEnum{
	"CREATING": ListSnapshotsLifecycleStateCreating,
	"ACTIVE":   ListSnapshotsLifecycleStateActive,
	"UPDATING": ListSnapshotsLifecycleStateUpdating,
	"DELETING": ListSnapshotsLifecycleStateDeleting,
	"DELETED":  ListSnapshotsLifecycleStateDeleted,
	"FAILED":   ListSnapshotsLifecycleStateFailed,
}

var mappingListSnapshotsLifecycleStateEnumLowerCase = map[string]ListSnapshotsLifecycleStateEnum{
	"creating": ListSnapshotsLifecycleStateCreating,
	"active":   ListSnapshotsLifecycleStateActive,
	"updating": ListSnapshotsLifecycleStateUpdating,
	"deleting": ListSnapshotsLifecycleStateDeleting,
	"deleted":  ListSnapshotsLifecycleStateDeleted,
	"failed":   ListSnapshotsLifecycleStateFailed,
}

// GetListSnapshotsLifecycleStateEnumValues Enumerates the set of values for ListSnapshotsLifecycleStateEnum
func GetListSnapshotsLifecycleStateEnumValues() []ListSnapshotsLifecycleStateEnum {
	values := make([]ListSnapshotsLifecycleStateEnum, 0)
	for _, v := range mappingListSnapshotsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSnapshotsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSnapshotsLifecycleStateEnum
func GetListSnapshotsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListSnapshotsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSnapshotsLifecycleStateEnum(val string) (ListSnapshotsLifecycleStateEnum, bool) {
	enum, ok := mappingListSnapshotsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSnapshotsSortOrderEnum Enum with underlying type: string
type ListSnapshotsSortOrderEnum string

// Set of constants representing the allowable values for ListSnapshotsSortOrderEnum
const (
	ListSnapshotsSortOrderAsc  ListSnapshotsSortOrderEnum = "ASC"
	ListSnapshotsSortOrderDesc ListSnapshotsSortOrderEnum = "DESC"
)

var mappingListSnapshotsSortOrderEnum = map[string]ListSnapshotsSortOrderEnum{
	"ASC":  ListSnapshotsSortOrderAsc,
	"DESC": ListSnapshotsSortOrderDesc,
}

var mappingListSnapshotsSortOrderEnumLowerCase = map[string]ListSnapshotsSortOrderEnum{
	"asc":  ListSnapshotsSortOrderAsc,
	"desc": ListSnapshotsSortOrderDesc,
}

// GetListSnapshotsSortOrderEnumValues Enumerates the set of values for ListSnapshotsSortOrderEnum
func GetListSnapshotsSortOrderEnumValues() []ListSnapshotsSortOrderEnum {
	values := make([]ListSnapshotsSortOrderEnum, 0)
	for _, v := range mappingListSnapshotsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSnapshotsSortOrderEnumStringValues Enumerates the set of values in String for ListSnapshotsSortOrderEnum
func GetListSnapshotsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSnapshotsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSnapshotsSortOrderEnum(val string) (ListSnapshotsSortOrderEnum, bool) {
	enum, ok := mappingListSnapshotsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
