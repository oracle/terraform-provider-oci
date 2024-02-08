// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDiscoveryJobsRequest wrapper for the ListDiscoveryJobs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDiscoveryJobs.go.html to see an example of how to use ListDiscoveryJobsRequest.
type ListDiscoveryJobsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListDiscoveryJobsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only the resources that match the specified discovery job OCID.
	DiscoveryJobId *string `mandatory:"false" contributesTo:"query" name:"discoveryJobId"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only the resources that match the specified lifecycle state.
	LifecycleState ListDiscoveryJobsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the specified sensitive data model OCID.
	SensitiveDataModelId *string `mandatory:"false" contributesTo:"query" name:"sensitiveDataModelId"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListDiscoveryJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for timeFinished is descending.
	// The default order for displayName is ascending.
	SortBy ListDiscoveryJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDiscoveryJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDiscoveryJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDiscoveryJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDiscoveryJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDiscoveryJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDiscoveryJobsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListDiscoveryJobsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiscoveryJobsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDiscoveryJobsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiscoveryJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDiscoveryJobsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiscoveryJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDiscoveryJobsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDiscoveryJobsResponse wrapper for the ListDiscoveryJobs operation
type ListDiscoveryJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DiscoveryJobCollection instances
	DiscoveryJobCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListDiscoveryJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDiscoveryJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDiscoveryJobsAccessLevelEnum Enum with underlying type: string
type ListDiscoveryJobsAccessLevelEnum string

// Set of constants representing the allowable values for ListDiscoveryJobsAccessLevelEnum
const (
	ListDiscoveryJobsAccessLevelRestricted ListDiscoveryJobsAccessLevelEnum = "RESTRICTED"
	ListDiscoveryJobsAccessLevelAccessible ListDiscoveryJobsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListDiscoveryJobsAccessLevelEnum = map[string]ListDiscoveryJobsAccessLevelEnum{
	"RESTRICTED": ListDiscoveryJobsAccessLevelRestricted,
	"ACCESSIBLE": ListDiscoveryJobsAccessLevelAccessible,
}

var mappingListDiscoveryJobsAccessLevelEnumLowerCase = map[string]ListDiscoveryJobsAccessLevelEnum{
	"restricted": ListDiscoveryJobsAccessLevelRestricted,
	"accessible": ListDiscoveryJobsAccessLevelAccessible,
}

// GetListDiscoveryJobsAccessLevelEnumValues Enumerates the set of values for ListDiscoveryJobsAccessLevelEnum
func GetListDiscoveryJobsAccessLevelEnumValues() []ListDiscoveryJobsAccessLevelEnum {
	values := make([]ListDiscoveryJobsAccessLevelEnum, 0)
	for _, v := range mappingListDiscoveryJobsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryJobsAccessLevelEnumStringValues Enumerates the set of values in String for ListDiscoveryJobsAccessLevelEnum
func GetListDiscoveryJobsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListDiscoveryJobsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryJobsAccessLevelEnum(val string) (ListDiscoveryJobsAccessLevelEnum, bool) {
	enum, ok := mappingListDiscoveryJobsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDiscoveryJobsLifecycleStateEnum Enum with underlying type: string
type ListDiscoveryJobsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDiscoveryJobsLifecycleStateEnum
const (
	ListDiscoveryJobsLifecycleStateCreating ListDiscoveryJobsLifecycleStateEnum = "CREATING"
	ListDiscoveryJobsLifecycleStateActive   ListDiscoveryJobsLifecycleStateEnum = "ACTIVE"
	ListDiscoveryJobsLifecycleStateUpdating ListDiscoveryJobsLifecycleStateEnum = "UPDATING"
	ListDiscoveryJobsLifecycleStateDeleting ListDiscoveryJobsLifecycleStateEnum = "DELETING"
	ListDiscoveryJobsLifecycleStateDeleted  ListDiscoveryJobsLifecycleStateEnum = "DELETED"
	ListDiscoveryJobsLifecycleStateFailed   ListDiscoveryJobsLifecycleStateEnum = "FAILED"
)

var mappingListDiscoveryJobsLifecycleStateEnum = map[string]ListDiscoveryJobsLifecycleStateEnum{
	"CREATING": ListDiscoveryJobsLifecycleStateCreating,
	"ACTIVE":   ListDiscoveryJobsLifecycleStateActive,
	"UPDATING": ListDiscoveryJobsLifecycleStateUpdating,
	"DELETING": ListDiscoveryJobsLifecycleStateDeleting,
	"DELETED":  ListDiscoveryJobsLifecycleStateDeleted,
	"FAILED":   ListDiscoveryJobsLifecycleStateFailed,
}

var mappingListDiscoveryJobsLifecycleStateEnumLowerCase = map[string]ListDiscoveryJobsLifecycleStateEnum{
	"creating": ListDiscoveryJobsLifecycleStateCreating,
	"active":   ListDiscoveryJobsLifecycleStateActive,
	"updating": ListDiscoveryJobsLifecycleStateUpdating,
	"deleting": ListDiscoveryJobsLifecycleStateDeleting,
	"deleted":  ListDiscoveryJobsLifecycleStateDeleted,
	"failed":   ListDiscoveryJobsLifecycleStateFailed,
}

// GetListDiscoveryJobsLifecycleStateEnumValues Enumerates the set of values for ListDiscoveryJobsLifecycleStateEnum
func GetListDiscoveryJobsLifecycleStateEnumValues() []ListDiscoveryJobsLifecycleStateEnum {
	values := make([]ListDiscoveryJobsLifecycleStateEnum, 0)
	for _, v := range mappingListDiscoveryJobsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryJobsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDiscoveryJobsLifecycleStateEnum
func GetListDiscoveryJobsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListDiscoveryJobsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryJobsLifecycleStateEnum(val string) (ListDiscoveryJobsLifecycleStateEnum, bool) {
	enum, ok := mappingListDiscoveryJobsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDiscoveryJobsSortOrderEnum Enum with underlying type: string
type ListDiscoveryJobsSortOrderEnum string

// Set of constants representing the allowable values for ListDiscoveryJobsSortOrderEnum
const (
	ListDiscoveryJobsSortOrderAsc  ListDiscoveryJobsSortOrderEnum = "ASC"
	ListDiscoveryJobsSortOrderDesc ListDiscoveryJobsSortOrderEnum = "DESC"
)

var mappingListDiscoveryJobsSortOrderEnum = map[string]ListDiscoveryJobsSortOrderEnum{
	"ASC":  ListDiscoveryJobsSortOrderAsc,
	"DESC": ListDiscoveryJobsSortOrderDesc,
}

var mappingListDiscoveryJobsSortOrderEnumLowerCase = map[string]ListDiscoveryJobsSortOrderEnum{
	"asc":  ListDiscoveryJobsSortOrderAsc,
	"desc": ListDiscoveryJobsSortOrderDesc,
}

// GetListDiscoveryJobsSortOrderEnumValues Enumerates the set of values for ListDiscoveryJobsSortOrderEnum
func GetListDiscoveryJobsSortOrderEnumValues() []ListDiscoveryJobsSortOrderEnum {
	values := make([]ListDiscoveryJobsSortOrderEnum, 0)
	for _, v := range mappingListDiscoveryJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryJobsSortOrderEnumStringValues Enumerates the set of values in String for ListDiscoveryJobsSortOrderEnum
func GetListDiscoveryJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDiscoveryJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryJobsSortOrderEnum(val string) (ListDiscoveryJobsSortOrderEnum, bool) {
	enum, ok := mappingListDiscoveryJobsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDiscoveryJobsSortByEnum Enum with underlying type: string
type ListDiscoveryJobsSortByEnum string

// Set of constants representing the allowable values for ListDiscoveryJobsSortByEnum
const (
	ListDiscoveryJobsSortByTimestarted ListDiscoveryJobsSortByEnum = "timeStarted"
	ListDiscoveryJobsSortByDisplayname ListDiscoveryJobsSortByEnum = "displayName"
)

var mappingListDiscoveryJobsSortByEnum = map[string]ListDiscoveryJobsSortByEnum{
	"timeStarted": ListDiscoveryJobsSortByTimestarted,
	"displayName": ListDiscoveryJobsSortByDisplayname,
}

var mappingListDiscoveryJobsSortByEnumLowerCase = map[string]ListDiscoveryJobsSortByEnum{
	"timestarted": ListDiscoveryJobsSortByTimestarted,
	"displayname": ListDiscoveryJobsSortByDisplayname,
}

// GetListDiscoveryJobsSortByEnumValues Enumerates the set of values for ListDiscoveryJobsSortByEnum
func GetListDiscoveryJobsSortByEnumValues() []ListDiscoveryJobsSortByEnum {
	values := make([]ListDiscoveryJobsSortByEnum, 0)
	for _, v := range mappingListDiscoveryJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryJobsSortByEnumStringValues Enumerates the set of values in String for ListDiscoveryJobsSortByEnum
func GetListDiscoveryJobsSortByEnumStringValues() []string {
	return []string{
		"timeStarted",
		"displayName",
	}
}

// GetMappingListDiscoveryJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryJobsSortByEnum(val string) (ListDiscoveryJobsSortByEnum, bool) {
	enum, ok := mappingListDiscoveryJobsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
