// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListDataSafePrivateEndpointsRequest wrapper for the ListDataSafePrivateEndpoints operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDataSafePrivateEndpoints.go.html to see an example of how to use ListDataSafePrivateEndpointsRequest.
type ListDataSafePrivateEndpointsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the specified VCN OCID.
	VcnId *string `mandatory:"false" contributesTo:"query" name:"vcnId"`

	// A filter to return only resources that match the specified lifecycle state.
	LifecycleState ListDataSafePrivateEndpointsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListDataSafePrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListDataSafePrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListDataSafePrivateEndpointsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataSafePrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataSafePrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataSafePrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataSafePrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataSafePrivateEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataSafePrivateEndpointsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDataSafePrivateEndpointsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSafePrivateEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataSafePrivateEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSafePrivateEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataSafePrivateEndpointsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSafePrivateEndpointsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListDataSafePrivateEndpointsAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataSafePrivateEndpointsResponse wrapper for the ListDataSafePrivateEndpoints operation
type ListDataSafePrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DataSafePrivateEndpointSummary instances
	Items []DataSafePrivateEndpointSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataSafePrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataSafePrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataSafePrivateEndpointsLifecycleStateEnum Enum with underlying type: string
type ListDataSafePrivateEndpointsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDataSafePrivateEndpointsLifecycleStateEnum
const (
	ListDataSafePrivateEndpointsLifecycleStateCreating ListDataSafePrivateEndpointsLifecycleStateEnum = "CREATING"
	ListDataSafePrivateEndpointsLifecycleStateUpdating ListDataSafePrivateEndpointsLifecycleStateEnum = "UPDATING"
	ListDataSafePrivateEndpointsLifecycleStateActive   ListDataSafePrivateEndpointsLifecycleStateEnum = "ACTIVE"
	ListDataSafePrivateEndpointsLifecycleStateDeleting ListDataSafePrivateEndpointsLifecycleStateEnum = "DELETING"
	ListDataSafePrivateEndpointsLifecycleStateDeleted  ListDataSafePrivateEndpointsLifecycleStateEnum = "DELETED"
	ListDataSafePrivateEndpointsLifecycleStateFailed   ListDataSafePrivateEndpointsLifecycleStateEnum = "FAILED"
	ListDataSafePrivateEndpointsLifecycleStateNa       ListDataSafePrivateEndpointsLifecycleStateEnum = "NA"
)

var mappingListDataSafePrivateEndpointsLifecycleStateEnum = map[string]ListDataSafePrivateEndpointsLifecycleStateEnum{
	"CREATING": ListDataSafePrivateEndpointsLifecycleStateCreating,
	"UPDATING": ListDataSafePrivateEndpointsLifecycleStateUpdating,
	"ACTIVE":   ListDataSafePrivateEndpointsLifecycleStateActive,
	"DELETING": ListDataSafePrivateEndpointsLifecycleStateDeleting,
	"DELETED":  ListDataSafePrivateEndpointsLifecycleStateDeleted,
	"FAILED":   ListDataSafePrivateEndpointsLifecycleStateFailed,
	"NA":       ListDataSafePrivateEndpointsLifecycleStateNa,
}

// GetListDataSafePrivateEndpointsLifecycleStateEnumValues Enumerates the set of values for ListDataSafePrivateEndpointsLifecycleStateEnum
func GetListDataSafePrivateEndpointsLifecycleStateEnumValues() []ListDataSafePrivateEndpointsLifecycleStateEnum {
	values := make([]ListDataSafePrivateEndpointsLifecycleStateEnum, 0)
	for _, v := range mappingListDataSafePrivateEndpointsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSafePrivateEndpointsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDataSafePrivateEndpointsLifecycleStateEnum
func GetListDataSafePrivateEndpointsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NA",
	}
}

// GetMappingListDataSafePrivateEndpointsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSafePrivateEndpointsLifecycleStateEnum(val string) (ListDataSafePrivateEndpointsLifecycleStateEnum, bool) {
	mappingListDataSafePrivateEndpointsLifecycleStateEnumIgnoreCase := make(map[string]ListDataSafePrivateEndpointsLifecycleStateEnum)
	for k, v := range mappingListDataSafePrivateEndpointsLifecycleStateEnum {
		mappingListDataSafePrivateEndpointsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDataSafePrivateEndpointsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSafePrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListDataSafePrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListDataSafePrivateEndpointsSortOrderEnum
const (
	ListDataSafePrivateEndpointsSortOrderAsc  ListDataSafePrivateEndpointsSortOrderEnum = "ASC"
	ListDataSafePrivateEndpointsSortOrderDesc ListDataSafePrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListDataSafePrivateEndpointsSortOrderEnum = map[string]ListDataSafePrivateEndpointsSortOrderEnum{
	"ASC":  ListDataSafePrivateEndpointsSortOrderAsc,
	"DESC": ListDataSafePrivateEndpointsSortOrderDesc,
}

// GetListDataSafePrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListDataSafePrivateEndpointsSortOrderEnum
func GetListDataSafePrivateEndpointsSortOrderEnumValues() []ListDataSafePrivateEndpointsSortOrderEnum {
	values := make([]ListDataSafePrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListDataSafePrivateEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSafePrivateEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListDataSafePrivateEndpointsSortOrderEnum
func GetListDataSafePrivateEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataSafePrivateEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSafePrivateEndpointsSortOrderEnum(val string) (ListDataSafePrivateEndpointsSortOrderEnum, bool) {
	mappingListDataSafePrivateEndpointsSortOrderEnumIgnoreCase := make(map[string]ListDataSafePrivateEndpointsSortOrderEnum)
	for k, v := range mappingListDataSafePrivateEndpointsSortOrderEnum {
		mappingListDataSafePrivateEndpointsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDataSafePrivateEndpointsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSafePrivateEndpointsSortByEnum Enum with underlying type: string
type ListDataSafePrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListDataSafePrivateEndpointsSortByEnum
const (
	ListDataSafePrivateEndpointsSortByTimecreated ListDataSafePrivateEndpointsSortByEnum = "TIMECREATED"
	ListDataSafePrivateEndpointsSortByDisplayname ListDataSafePrivateEndpointsSortByEnum = "DISPLAYNAME"
)

var mappingListDataSafePrivateEndpointsSortByEnum = map[string]ListDataSafePrivateEndpointsSortByEnum{
	"TIMECREATED": ListDataSafePrivateEndpointsSortByTimecreated,
	"DISPLAYNAME": ListDataSafePrivateEndpointsSortByDisplayname,
}

// GetListDataSafePrivateEndpointsSortByEnumValues Enumerates the set of values for ListDataSafePrivateEndpointsSortByEnum
func GetListDataSafePrivateEndpointsSortByEnumValues() []ListDataSafePrivateEndpointsSortByEnum {
	values := make([]ListDataSafePrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListDataSafePrivateEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSafePrivateEndpointsSortByEnumStringValues Enumerates the set of values in String for ListDataSafePrivateEndpointsSortByEnum
func GetListDataSafePrivateEndpointsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListDataSafePrivateEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSafePrivateEndpointsSortByEnum(val string) (ListDataSafePrivateEndpointsSortByEnum, bool) {
	mappingListDataSafePrivateEndpointsSortByEnumIgnoreCase := make(map[string]ListDataSafePrivateEndpointsSortByEnum)
	for k, v := range mappingListDataSafePrivateEndpointsSortByEnum {
		mappingListDataSafePrivateEndpointsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDataSafePrivateEndpointsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSafePrivateEndpointsAccessLevelEnum Enum with underlying type: string
type ListDataSafePrivateEndpointsAccessLevelEnum string

// Set of constants representing the allowable values for ListDataSafePrivateEndpointsAccessLevelEnum
const (
	ListDataSafePrivateEndpointsAccessLevelRestricted ListDataSafePrivateEndpointsAccessLevelEnum = "RESTRICTED"
	ListDataSafePrivateEndpointsAccessLevelAccessible ListDataSafePrivateEndpointsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListDataSafePrivateEndpointsAccessLevelEnum = map[string]ListDataSafePrivateEndpointsAccessLevelEnum{
	"RESTRICTED": ListDataSafePrivateEndpointsAccessLevelRestricted,
	"ACCESSIBLE": ListDataSafePrivateEndpointsAccessLevelAccessible,
}

// GetListDataSafePrivateEndpointsAccessLevelEnumValues Enumerates the set of values for ListDataSafePrivateEndpointsAccessLevelEnum
func GetListDataSafePrivateEndpointsAccessLevelEnumValues() []ListDataSafePrivateEndpointsAccessLevelEnum {
	values := make([]ListDataSafePrivateEndpointsAccessLevelEnum, 0)
	for _, v := range mappingListDataSafePrivateEndpointsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSafePrivateEndpointsAccessLevelEnumStringValues Enumerates the set of values in String for ListDataSafePrivateEndpointsAccessLevelEnum
func GetListDataSafePrivateEndpointsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListDataSafePrivateEndpointsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSafePrivateEndpointsAccessLevelEnum(val string) (ListDataSafePrivateEndpointsAccessLevelEnum, bool) {
	mappingListDataSafePrivateEndpointsAccessLevelEnumIgnoreCase := make(map[string]ListDataSafePrivateEndpointsAccessLevelEnum)
	for k, v := range mappingListDataSafePrivateEndpointsAccessLevelEnum {
		mappingListDataSafePrivateEndpointsAccessLevelEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDataSafePrivateEndpointsAccessLevelEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
