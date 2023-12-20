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

// ListDatastoreClustersRequest wrapper for the ListDatastoreClusters operation
type ListDatastoreClustersRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Datastore Cluster.
	DatastoreClusterId *string `mandatory:"false" contributesTo:"query" name:"datastoreClusterId"`

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
	SortOrder ListDatastoreClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListDatastoreClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request. If you need to contact Oracle about a particular
	// request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The lifecycle state of the resource.
	LifecycleState ListDatastoreClustersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatastoreClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatastoreClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatastoreClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatastoreClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatastoreClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatastoreClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatastoreClustersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatastoreClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatastoreClustersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatastoreClustersLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDatastoreClustersLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatastoreClustersResponse wrapper for the ListDatastoreClusters operation
type ListDatastoreClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatastoreClusterCollection instances
	DatastoreClusterCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatastoreClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatastoreClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatastoreClustersSortOrderEnum Enum with underlying type: string
type ListDatastoreClustersSortOrderEnum string

// Set of constants representing the allowable values for ListDatastoreClustersSortOrderEnum
const (
	ListDatastoreClustersSortOrderAsc  ListDatastoreClustersSortOrderEnum = "ASC"
	ListDatastoreClustersSortOrderDesc ListDatastoreClustersSortOrderEnum = "DESC"
)

var mappingListDatastoreClustersSortOrderEnum = map[string]ListDatastoreClustersSortOrderEnum{
	"ASC":  ListDatastoreClustersSortOrderAsc,
	"DESC": ListDatastoreClustersSortOrderDesc,
}

var mappingListDatastoreClustersSortOrderEnumLowerCase = map[string]ListDatastoreClustersSortOrderEnum{
	"asc":  ListDatastoreClustersSortOrderAsc,
	"desc": ListDatastoreClustersSortOrderDesc,
}

// GetListDatastoreClustersSortOrderEnumValues Enumerates the set of values for ListDatastoreClustersSortOrderEnum
func GetListDatastoreClustersSortOrderEnumValues() []ListDatastoreClustersSortOrderEnum {
	values := make([]ListDatastoreClustersSortOrderEnum, 0)
	for _, v := range mappingListDatastoreClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatastoreClustersSortOrderEnumStringValues Enumerates the set of values in String for ListDatastoreClustersSortOrderEnum
func GetListDatastoreClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatastoreClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatastoreClustersSortOrderEnum(val string) (ListDatastoreClustersSortOrderEnum, bool) {
	enum, ok := mappingListDatastoreClustersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatastoreClustersSortByEnum Enum with underlying type: string
type ListDatastoreClustersSortByEnum string

// Set of constants representing the allowable values for ListDatastoreClustersSortByEnum
const (
	ListDatastoreClustersSortByTimecreated ListDatastoreClustersSortByEnum = "timeCreated"
	ListDatastoreClustersSortByDisplayname ListDatastoreClustersSortByEnum = "displayName"
)

var mappingListDatastoreClustersSortByEnum = map[string]ListDatastoreClustersSortByEnum{
	"timeCreated": ListDatastoreClustersSortByTimecreated,
	"displayName": ListDatastoreClustersSortByDisplayname,
}

var mappingListDatastoreClustersSortByEnumLowerCase = map[string]ListDatastoreClustersSortByEnum{
	"timecreated": ListDatastoreClustersSortByTimecreated,
	"displayname": ListDatastoreClustersSortByDisplayname,
}

// GetListDatastoreClustersSortByEnumValues Enumerates the set of values for ListDatastoreClustersSortByEnum
func GetListDatastoreClustersSortByEnumValues() []ListDatastoreClustersSortByEnum {
	values := make([]ListDatastoreClustersSortByEnum, 0)
	for _, v := range mappingListDatastoreClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatastoreClustersSortByEnumStringValues Enumerates the set of values in String for ListDatastoreClustersSortByEnum
func GetListDatastoreClustersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatastoreClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatastoreClustersSortByEnum(val string) (ListDatastoreClustersSortByEnum, bool) {
	enum, ok := mappingListDatastoreClustersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatastoreClustersLifecycleStateEnum Enum with underlying type: string
type ListDatastoreClustersLifecycleStateEnum string

// Set of constants representing the allowable values for ListDatastoreClustersLifecycleStateEnum
const (
	ListDatastoreClustersLifecycleStateCreating ListDatastoreClustersLifecycleStateEnum = "CREATING"
	ListDatastoreClustersLifecycleStateUpdating ListDatastoreClustersLifecycleStateEnum = "UPDATING"
	ListDatastoreClustersLifecycleStateActive   ListDatastoreClustersLifecycleStateEnum = "ACTIVE"
	ListDatastoreClustersLifecycleStateDeleting ListDatastoreClustersLifecycleStateEnum = "DELETING"
	ListDatastoreClustersLifecycleStateDeleted  ListDatastoreClustersLifecycleStateEnum = "DELETED"
	ListDatastoreClustersLifecycleStateFailed   ListDatastoreClustersLifecycleStateEnum = "FAILED"
)

var mappingListDatastoreClustersLifecycleStateEnum = map[string]ListDatastoreClustersLifecycleStateEnum{
	"CREATING": ListDatastoreClustersLifecycleStateCreating,
	"UPDATING": ListDatastoreClustersLifecycleStateUpdating,
	"ACTIVE":   ListDatastoreClustersLifecycleStateActive,
	"DELETING": ListDatastoreClustersLifecycleStateDeleting,
	"DELETED":  ListDatastoreClustersLifecycleStateDeleted,
	"FAILED":   ListDatastoreClustersLifecycleStateFailed,
}

var mappingListDatastoreClustersLifecycleStateEnumLowerCase = map[string]ListDatastoreClustersLifecycleStateEnum{
	"creating": ListDatastoreClustersLifecycleStateCreating,
	"updating": ListDatastoreClustersLifecycleStateUpdating,
	"active":   ListDatastoreClustersLifecycleStateActive,
	"deleting": ListDatastoreClustersLifecycleStateDeleting,
	"deleted":  ListDatastoreClustersLifecycleStateDeleted,
	"failed":   ListDatastoreClustersLifecycleStateFailed,
}

// GetListDatastoreClustersLifecycleStateEnumValues Enumerates the set of values for ListDatastoreClustersLifecycleStateEnum
func GetListDatastoreClustersLifecycleStateEnumValues() []ListDatastoreClustersLifecycleStateEnum {
	values := make([]ListDatastoreClustersLifecycleStateEnum, 0)
	for _, v := range mappingListDatastoreClustersLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatastoreClustersLifecycleStateEnumStringValues Enumerates the set of values in String for ListDatastoreClustersLifecycleStateEnum
func GetListDatastoreClustersLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListDatastoreClustersLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatastoreClustersLifecycleStateEnum(val string) (ListDatastoreClustersLifecycleStateEnum, bool) {
	enum, ok := mappingListDatastoreClustersLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
