// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package clusterplacementgroups

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListClusterPlacementGroupsRequest wrapper for the ListClusterPlacementGroups operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/clusterplacementgroups/ListClusterPlacementGroups.go.html to see an example of how to use ListClusterPlacementGroupsRequest.
type ListClusterPlacementGroupsRequest struct {

	// A filter to return only the resources that match the specified compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources that match the specified lifecycle state.
	LifecycleState ClusterPlacementGroupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire display name specified.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only the resources that match the specified availability domain.
	Ad *string `mandatory:"false" contributesTo:"query" name:"ad"`

	// A filter to return only the resources that match the specified unique cluster placement group identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListClusterPlacementGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order. The default order for `timeCreated` is descending. The default order for `name` is ascending.
	SortBy ListClusterPlacementGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// When set to `true`, cluster placement groups in all compartments under the specified compartment are returned. The default is set to `false`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListClusterPlacementGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListClusterPlacementGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListClusterPlacementGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListClusterPlacementGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListClusterPlacementGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingClusterPlacementGroupLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetClusterPlacementGroupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListClusterPlacementGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListClusterPlacementGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListClusterPlacementGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListClusterPlacementGroupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListClusterPlacementGroupsResponse wrapper for the ListClusterPlacementGroups operation
type ListClusterPlacementGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ClusterPlacementGroupCollection instances
	ClusterPlacementGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items to get. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items. For information about pagination, see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListClusterPlacementGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListClusterPlacementGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListClusterPlacementGroupsSortOrderEnum Enum with underlying type: string
type ListClusterPlacementGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListClusterPlacementGroupsSortOrderEnum
const (
	ListClusterPlacementGroupsSortOrderAsc  ListClusterPlacementGroupsSortOrderEnum = "ASC"
	ListClusterPlacementGroupsSortOrderDesc ListClusterPlacementGroupsSortOrderEnum = "DESC"
)

var mappingListClusterPlacementGroupsSortOrderEnum = map[string]ListClusterPlacementGroupsSortOrderEnum{
	"ASC":  ListClusterPlacementGroupsSortOrderAsc,
	"DESC": ListClusterPlacementGroupsSortOrderDesc,
}

var mappingListClusterPlacementGroupsSortOrderEnumLowerCase = map[string]ListClusterPlacementGroupsSortOrderEnum{
	"asc":  ListClusterPlacementGroupsSortOrderAsc,
	"desc": ListClusterPlacementGroupsSortOrderDesc,
}

// GetListClusterPlacementGroupsSortOrderEnumValues Enumerates the set of values for ListClusterPlacementGroupsSortOrderEnum
func GetListClusterPlacementGroupsSortOrderEnumValues() []ListClusterPlacementGroupsSortOrderEnum {
	values := make([]ListClusterPlacementGroupsSortOrderEnum, 0)
	for _, v := range mappingListClusterPlacementGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListClusterPlacementGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListClusterPlacementGroupsSortOrderEnum
func GetListClusterPlacementGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListClusterPlacementGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListClusterPlacementGroupsSortOrderEnum(val string) (ListClusterPlacementGroupsSortOrderEnum, bool) {
	enum, ok := mappingListClusterPlacementGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListClusterPlacementGroupsSortByEnum Enum with underlying type: string
type ListClusterPlacementGroupsSortByEnum string

// Set of constants representing the allowable values for ListClusterPlacementGroupsSortByEnum
const (
	ListClusterPlacementGroupsSortByTimecreated ListClusterPlacementGroupsSortByEnum = "timeCreated"
	ListClusterPlacementGroupsSortByName        ListClusterPlacementGroupsSortByEnum = "name"
	ListClusterPlacementGroupsSortById          ListClusterPlacementGroupsSortByEnum = "id"
)

var mappingListClusterPlacementGroupsSortByEnum = map[string]ListClusterPlacementGroupsSortByEnum{
	"timeCreated": ListClusterPlacementGroupsSortByTimecreated,
	"name":        ListClusterPlacementGroupsSortByName,
	"id":          ListClusterPlacementGroupsSortById,
}

var mappingListClusterPlacementGroupsSortByEnumLowerCase = map[string]ListClusterPlacementGroupsSortByEnum{
	"timecreated": ListClusterPlacementGroupsSortByTimecreated,
	"name":        ListClusterPlacementGroupsSortByName,
	"id":          ListClusterPlacementGroupsSortById,
}

// GetListClusterPlacementGroupsSortByEnumValues Enumerates the set of values for ListClusterPlacementGroupsSortByEnum
func GetListClusterPlacementGroupsSortByEnumValues() []ListClusterPlacementGroupsSortByEnum {
	values := make([]ListClusterPlacementGroupsSortByEnum, 0)
	for _, v := range mappingListClusterPlacementGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListClusterPlacementGroupsSortByEnumStringValues Enumerates the set of values in String for ListClusterPlacementGroupsSortByEnum
func GetListClusterPlacementGroupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
		"id",
	}
}

// GetMappingListClusterPlacementGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListClusterPlacementGroupsSortByEnum(val string) (ListClusterPlacementGroupsSortByEnum, bool) {
	enum, ok := mappingListClusterPlacementGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
