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

// ListSqlCollectionsRequest wrapper for the ListSqlCollections operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlCollections.go.html to see an example of how to use ListSqlCollectionsRequest.
type ListSqlCollectionsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSqlCollectionsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The current state of the SQL collection.
	LifecycleState ListSqlCollectionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified OCID of the SQL collection resource.
	SqlCollectionId *string `mandatory:"false" contributesTo:"query" name:"sqlCollectionId"`

	// A filter to return only the resources that were created after the specified date and time, as defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// Search for resources that were created before a specific date.
	// Specifying this parameter corresponding `timeCreatedLessThan`
	// parameter will retrieve all resources created before the
	// specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as
	// defined by RFC 3339.
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only items that match the specified user name.
	DbUserName *string `mandatory:"false" contributesTo:"query" name:"dbUserName"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSqlCollectionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting parameter order (sortOrder) can be specified.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListSqlCollectionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlCollectionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlCollectionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlCollectionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlCollectionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlCollectionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlCollectionsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSqlCollectionsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlCollectionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSqlCollectionsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlCollectionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSqlCollectionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlCollectionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSqlCollectionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlCollectionsResponse wrapper for the ListSqlCollections operation
type ListSqlCollectionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlCollectionCollection instances
	SqlCollectionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSqlCollectionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlCollectionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlCollectionsAccessLevelEnum Enum with underlying type: string
type ListSqlCollectionsAccessLevelEnum string

// Set of constants representing the allowable values for ListSqlCollectionsAccessLevelEnum
const (
	ListSqlCollectionsAccessLevelRestricted ListSqlCollectionsAccessLevelEnum = "RESTRICTED"
	ListSqlCollectionsAccessLevelAccessible ListSqlCollectionsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSqlCollectionsAccessLevelEnum = map[string]ListSqlCollectionsAccessLevelEnum{
	"RESTRICTED": ListSqlCollectionsAccessLevelRestricted,
	"ACCESSIBLE": ListSqlCollectionsAccessLevelAccessible,
}

var mappingListSqlCollectionsAccessLevelEnumLowerCase = map[string]ListSqlCollectionsAccessLevelEnum{
	"restricted": ListSqlCollectionsAccessLevelRestricted,
	"accessible": ListSqlCollectionsAccessLevelAccessible,
}

// GetListSqlCollectionsAccessLevelEnumValues Enumerates the set of values for ListSqlCollectionsAccessLevelEnum
func GetListSqlCollectionsAccessLevelEnumValues() []ListSqlCollectionsAccessLevelEnum {
	values := make([]ListSqlCollectionsAccessLevelEnum, 0)
	for _, v := range mappingListSqlCollectionsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlCollectionsAccessLevelEnumStringValues Enumerates the set of values in String for ListSqlCollectionsAccessLevelEnum
func GetListSqlCollectionsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSqlCollectionsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlCollectionsAccessLevelEnum(val string) (ListSqlCollectionsAccessLevelEnum, bool) {
	enum, ok := mappingListSqlCollectionsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlCollectionsLifecycleStateEnum Enum with underlying type: string
type ListSqlCollectionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSqlCollectionsLifecycleStateEnum
const (
	ListSqlCollectionsLifecycleStateCreating       ListSqlCollectionsLifecycleStateEnum = "CREATING"
	ListSqlCollectionsLifecycleStateUpdating       ListSqlCollectionsLifecycleStateEnum = "UPDATING"
	ListSqlCollectionsLifecycleStateCollecting     ListSqlCollectionsLifecycleStateEnum = "COLLECTING"
	ListSqlCollectionsLifecycleStateCompleted      ListSqlCollectionsLifecycleStateEnum = "COMPLETED"
	ListSqlCollectionsLifecycleStateInactive       ListSqlCollectionsLifecycleStateEnum = "INACTIVE"
	ListSqlCollectionsLifecycleStateFailed         ListSqlCollectionsLifecycleStateEnum = "FAILED"
	ListSqlCollectionsLifecycleStateDeleting       ListSqlCollectionsLifecycleStateEnum = "DELETING"
	ListSqlCollectionsLifecycleStateDeleted        ListSqlCollectionsLifecycleStateEnum = "DELETED"
	ListSqlCollectionsLifecycleStateNeedsAttention ListSqlCollectionsLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListSqlCollectionsLifecycleStateEnum = map[string]ListSqlCollectionsLifecycleStateEnum{
	"CREATING":        ListSqlCollectionsLifecycleStateCreating,
	"UPDATING":        ListSqlCollectionsLifecycleStateUpdating,
	"COLLECTING":      ListSqlCollectionsLifecycleStateCollecting,
	"COMPLETED":       ListSqlCollectionsLifecycleStateCompleted,
	"INACTIVE":        ListSqlCollectionsLifecycleStateInactive,
	"FAILED":          ListSqlCollectionsLifecycleStateFailed,
	"DELETING":        ListSqlCollectionsLifecycleStateDeleting,
	"DELETED":         ListSqlCollectionsLifecycleStateDeleted,
	"NEEDS_ATTENTION": ListSqlCollectionsLifecycleStateNeedsAttention,
}

var mappingListSqlCollectionsLifecycleStateEnumLowerCase = map[string]ListSqlCollectionsLifecycleStateEnum{
	"creating":        ListSqlCollectionsLifecycleStateCreating,
	"updating":        ListSqlCollectionsLifecycleStateUpdating,
	"collecting":      ListSqlCollectionsLifecycleStateCollecting,
	"completed":       ListSqlCollectionsLifecycleStateCompleted,
	"inactive":        ListSqlCollectionsLifecycleStateInactive,
	"failed":          ListSqlCollectionsLifecycleStateFailed,
	"deleting":        ListSqlCollectionsLifecycleStateDeleting,
	"deleted":         ListSqlCollectionsLifecycleStateDeleted,
	"needs_attention": ListSqlCollectionsLifecycleStateNeedsAttention,
}

// GetListSqlCollectionsLifecycleStateEnumValues Enumerates the set of values for ListSqlCollectionsLifecycleStateEnum
func GetListSqlCollectionsLifecycleStateEnumValues() []ListSqlCollectionsLifecycleStateEnum {
	values := make([]ListSqlCollectionsLifecycleStateEnum, 0)
	for _, v := range mappingListSqlCollectionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlCollectionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSqlCollectionsLifecycleStateEnum
func GetListSqlCollectionsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"COLLECTING",
		"COMPLETED",
		"INACTIVE",
		"FAILED",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListSqlCollectionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlCollectionsLifecycleStateEnum(val string) (ListSqlCollectionsLifecycleStateEnum, bool) {
	enum, ok := mappingListSqlCollectionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlCollectionsSortOrderEnum Enum with underlying type: string
type ListSqlCollectionsSortOrderEnum string

// Set of constants representing the allowable values for ListSqlCollectionsSortOrderEnum
const (
	ListSqlCollectionsSortOrderAsc  ListSqlCollectionsSortOrderEnum = "ASC"
	ListSqlCollectionsSortOrderDesc ListSqlCollectionsSortOrderEnum = "DESC"
)

var mappingListSqlCollectionsSortOrderEnum = map[string]ListSqlCollectionsSortOrderEnum{
	"ASC":  ListSqlCollectionsSortOrderAsc,
	"DESC": ListSqlCollectionsSortOrderDesc,
}

var mappingListSqlCollectionsSortOrderEnumLowerCase = map[string]ListSqlCollectionsSortOrderEnum{
	"asc":  ListSqlCollectionsSortOrderAsc,
	"desc": ListSqlCollectionsSortOrderDesc,
}

// GetListSqlCollectionsSortOrderEnumValues Enumerates the set of values for ListSqlCollectionsSortOrderEnum
func GetListSqlCollectionsSortOrderEnumValues() []ListSqlCollectionsSortOrderEnum {
	values := make([]ListSqlCollectionsSortOrderEnum, 0)
	for _, v := range mappingListSqlCollectionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlCollectionsSortOrderEnumStringValues Enumerates the set of values in String for ListSqlCollectionsSortOrderEnum
func GetListSqlCollectionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSqlCollectionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlCollectionsSortOrderEnum(val string) (ListSqlCollectionsSortOrderEnum, bool) {
	enum, ok := mappingListSqlCollectionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlCollectionsSortByEnum Enum with underlying type: string
type ListSqlCollectionsSortByEnum string

// Set of constants representing the allowable values for ListSqlCollectionsSortByEnum
const (
	ListSqlCollectionsSortByTimecreated     ListSqlCollectionsSortByEnum = "TIMECREATED"
	ListSqlCollectionsSortByDisplayname     ListSqlCollectionsSortByEnum = "DISPLAYNAME"
	ListSqlCollectionsSortByTimelaststarted ListSqlCollectionsSortByEnum = "TIMELASTSTARTED"
)

var mappingListSqlCollectionsSortByEnum = map[string]ListSqlCollectionsSortByEnum{
	"TIMECREATED":     ListSqlCollectionsSortByTimecreated,
	"DISPLAYNAME":     ListSqlCollectionsSortByDisplayname,
	"TIMELASTSTARTED": ListSqlCollectionsSortByTimelaststarted,
}

var mappingListSqlCollectionsSortByEnumLowerCase = map[string]ListSqlCollectionsSortByEnum{
	"timecreated":     ListSqlCollectionsSortByTimecreated,
	"displayname":     ListSqlCollectionsSortByDisplayname,
	"timelaststarted": ListSqlCollectionsSortByTimelaststarted,
}

// GetListSqlCollectionsSortByEnumValues Enumerates the set of values for ListSqlCollectionsSortByEnum
func GetListSqlCollectionsSortByEnumValues() []ListSqlCollectionsSortByEnum {
	values := make([]ListSqlCollectionsSortByEnum, 0)
	for _, v := range mappingListSqlCollectionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlCollectionsSortByEnumStringValues Enumerates the set of values in String for ListSqlCollectionsSortByEnum
func GetListSqlCollectionsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
		"TIMELASTSTARTED",
	}
}

// GetMappingListSqlCollectionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlCollectionsSortByEnum(val string) (ListSqlCollectionsSortByEnum, bool) {
	enum, ok := mappingListSqlCollectionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
