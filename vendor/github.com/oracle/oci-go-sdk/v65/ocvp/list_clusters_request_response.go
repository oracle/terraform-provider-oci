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

// ListClustersRequest wrapper for the ListClusters operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ListClusters.go.html to see an example of how to use ListClustersRequest.
type ListClustersRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SDDC.
	SddcId *string `mandatory:"false" contributesTo:"query" name:"sddcId"`

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
	SortOrder ListClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request. If you need to contact Oracle about a particular
	// request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The lifecycle state of the resource.
	LifecycleState ListClustersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment as optional parameter.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListClustersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListClustersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListClustersLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListClustersLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListClustersResponse wrapper for the ListClusters operation
type ListClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ClusterCollection instances
	ClusterCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListClustersSortOrderEnum Enum with underlying type: string
type ListClustersSortOrderEnum string

// Set of constants representing the allowable values for ListClustersSortOrderEnum
const (
	ListClustersSortOrderAsc  ListClustersSortOrderEnum = "ASC"
	ListClustersSortOrderDesc ListClustersSortOrderEnum = "DESC"
)

var mappingListClustersSortOrderEnum = map[string]ListClustersSortOrderEnum{
	"ASC":  ListClustersSortOrderAsc,
	"DESC": ListClustersSortOrderDesc,
}

var mappingListClustersSortOrderEnumLowerCase = map[string]ListClustersSortOrderEnum{
	"asc":  ListClustersSortOrderAsc,
	"desc": ListClustersSortOrderDesc,
}

// GetListClustersSortOrderEnumValues Enumerates the set of values for ListClustersSortOrderEnum
func GetListClustersSortOrderEnumValues() []ListClustersSortOrderEnum {
	values := make([]ListClustersSortOrderEnum, 0)
	for _, v := range mappingListClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListClustersSortOrderEnumStringValues Enumerates the set of values in String for ListClustersSortOrderEnum
func GetListClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListClustersSortOrderEnum(val string) (ListClustersSortOrderEnum, bool) {
	enum, ok := mappingListClustersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListClustersSortByEnum Enum with underlying type: string
type ListClustersSortByEnum string

// Set of constants representing the allowable values for ListClustersSortByEnum
const (
	ListClustersSortByTimecreated ListClustersSortByEnum = "timeCreated"
	ListClustersSortByDisplayname ListClustersSortByEnum = "displayName"
)

var mappingListClustersSortByEnum = map[string]ListClustersSortByEnum{
	"timeCreated": ListClustersSortByTimecreated,
	"displayName": ListClustersSortByDisplayname,
}

var mappingListClustersSortByEnumLowerCase = map[string]ListClustersSortByEnum{
	"timecreated": ListClustersSortByTimecreated,
	"displayname": ListClustersSortByDisplayname,
}

// GetListClustersSortByEnumValues Enumerates the set of values for ListClustersSortByEnum
func GetListClustersSortByEnumValues() []ListClustersSortByEnum {
	values := make([]ListClustersSortByEnum, 0)
	for _, v := range mappingListClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListClustersSortByEnumStringValues Enumerates the set of values in String for ListClustersSortByEnum
func GetListClustersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListClustersSortByEnum(val string) (ListClustersSortByEnum, bool) {
	enum, ok := mappingListClustersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListClustersLifecycleStateEnum Enum with underlying type: string
type ListClustersLifecycleStateEnum string

// Set of constants representing the allowable values for ListClustersLifecycleStateEnum
const (
	ListClustersLifecycleStateCreating ListClustersLifecycleStateEnum = "CREATING"
	ListClustersLifecycleStateUpdating ListClustersLifecycleStateEnum = "UPDATING"
	ListClustersLifecycleStateActive   ListClustersLifecycleStateEnum = "ACTIVE"
	ListClustersLifecycleStateDeleting ListClustersLifecycleStateEnum = "DELETING"
	ListClustersLifecycleStateDeleted  ListClustersLifecycleStateEnum = "DELETED"
	ListClustersLifecycleStateFailed   ListClustersLifecycleStateEnum = "FAILED"
)

var mappingListClustersLifecycleStateEnum = map[string]ListClustersLifecycleStateEnum{
	"CREATING": ListClustersLifecycleStateCreating,
	"UPDATING": ListClustersLifecycleStateUpdating,
	"ACTIVE":   ListClustersLifecycleStateActive,
	"DELETING": ListClustersLifecycleStateDeleting,
	"DELETED":  ListClustersLifecycleStateDeleted,
	"FAILED":   ListClustersLifecycleStateFailed,
}

var mappingListClustersLifecycleStateEnumLowerCase = map[string]ListClustersLifecycleStateEnum{
	"creating": ListClustersLifecycleStateCreating,
	"updating": ListClustersLifecycleStateUpdating,
	"active":   ListClustersLifecycleStateActive,
	"deleting": ListClustersLifecycleStateDeleting,
	"deleted":  ListClustersLifecycleStateDeleted,
	"failed":   ListClustersLifecycleStateFailed,
}

// GetListClustersLifecycleStateEnumValues Enumerates the set of values for ListClustersLifecycleStateEnum
func GetListClustersLifecycleStateEnumValues() []ListClustersLifecycleStateEnum {
	values := make([]ListClustersLifecycleStateEnum, 0)
	for _, v := range mappingListClustersLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListClustersLifecycleStateEnumStringValues Enumerates the set of values in String for ListClustersLifecycleStateEnum
func GetListClustersLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListClustersLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListClustersLifecycleStateEnum(val string) (ListClustersLifecycleStateEnum, bool) {
	enum, ok := mappingListClustersLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
