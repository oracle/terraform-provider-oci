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

// ListTargetAlertPolicyAssociationsRequest wrapper for the ListTargetAlertPolicyAssociations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTargetAlertPolicyAssociations.go.html to see an example of how to use ListTargetAlertPolicyAssociationsRequest.
type ListTargetAlertPolicyAssociationsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only items related to a specific target-alert policy association ID.
	TargetAlertPolicyAssociationId *string `mandatory:"false" contributesTo:"query" name:"targetAlertPolicyAssociationId"`

	// A filter to return policy by it's OCID.
	AlertPolicyId *string `mandatory:"false" contributesTo:"query" name:"alertPolicyId"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// An optional filter to return only alert policies that have the given life-cycle state.
	LifecycleState ListTargetAlertPolicyAssociationsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListTargetAlertPolicyAssociationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort parameter may be provided.
	SortBy ListTargetAlertPolicyAssociationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

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

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListTargetAlertPolicyAssociationsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetAlertPolicyAssociationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetAlertPolicyAssociationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetAlertPolicyAssociationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetAlertPolicyAssociationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetAlertPolicyAssociationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetAlertPolicyAssociationsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTargetAlertPolicyAssociationsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetAlertPolicyAssociationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetAlertPolicyAssociationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetAlertPolicyAssociationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetAlertPolicyAssociationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetAlertPolicyAssociationsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListTargetAlertPolicyAssociationsAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTargetAlertPolicyAssociationsResponse wrapper for the ListTargetAlertPolicyAssociations operation
type ListTargetAlertPolicyAssociationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetAlertPolicyAssociationCollection instances
	TargetAlertPolicyAssociationCollection `presentIn:"body"`

	// For optimistic concurrency control. For more information, see ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven)
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListTargetAlertPolicyAssociationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetAlertPolicyAssociationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetAlertPolicyAssociationsLifecycleStateEnum Enum with underlying type: string
type ListTargetAlertPolicyAssociationsLifecycleStateEnum string

// Set of constants representing the allowable values for ListTargetAlertPolicyAssociationsLifecycleStateEnum
const (
	ListTargetAlertPolicyAssociationsLifecycleStateCreating ListTargetAlertPolicyAssociationsLifecycleStateEnum = "CREATING"
	ListTargetAlertPolicyAssociationsLifecycleStateUpdating ListTargetAlertPolicyAssociationsLifecycleStateEnum = "UPDATING"
	ListTargetAlertPolicyAssociationsLifecycleStateActive   ListTargetAlertPolicyAssociationsLifecycleStateEnum = "ACTIVE"
	ListTargetAlertPolicyAssociationsLifecycleStateDeleting ListTargetAlertPolicyAssociationsLifecycleStateEnum = "DELETING"
	ListTargetAlertPolicyAssociationsLifecycleStateDeleted  ListTargetAlertPolicyAssociationsLifecycleStateEnum = "DELETED"
	ListTargetAlertPolicyAssociationsLifecycleStateFailed   ListTargetAlertPolicyAssociationsLifecycleStateEnum = "FAILED"
)

var mappingListTargetAlertPolicyAssociationsLifecycleStateEnum = map[string]ListTargetAlertPolicyAssociationsLifecycleStateEnum{
	"CREATING": ListTargetAlertPolicyAssociationsLifecycleStateCreating,
	"UPDATING": ListTargetAlertPolicyAssociationsLifecycleStateUpdating,
	"ACTIVE":   ListTargetAlertPolicyAssociationsLifecycleStateActive,
	"DELETING": ListTargetAlertPolicyAssociationsLifecycleStateDeleting,
	"DELETED":  ListTargetAlertPolicyAssociationsLifecycleStateDeleted,
	"FAILED":   ListTargetAlertPolicyAssociationsLifecycleStateFailed,
}

var mappingListTargetAlertPolicyAssociationsLifecycleStateEnumLowerCase = map[string]ListTargetAlertPolicyAssociationsLifecycleStateEnum{
	"creating": ListTargetAlertPolicyAssociationsLifecycleStateCreating,
	"updating": ListTargetAlertPolicyAssociationsLifecycleStateUpdating,
	"active":   ListTargetAlertPolicyAssociationsLifecycleStateActive,
	"deleting": ListTargetAlertPolicyAssociationsLifecycleStateDeleting,
	"deleted":  ListTargetAlertPolicyAssociationsLifecycleStateDeleted,
	"failed":   ListTargetAlertPolicyAssociationsLifecycleStateFailed,
}

// GetListTargetAlertPolicyAssociationsLifecycleStateEnumValues Enumerates the set of values for ListTargetAlertPolicyAssociationsLifecycleStateEnum
func GetListTargetAlertPolicyAssociationsLifecycleStateEnumValues() []ListTargetAlertPolicyAssociationsLifecycleStateEnum {
	values := make([]ListTargetAlertPolicyAssociationsLifecycleStateEnum, 0)
	for _, v := range mappingListTargetAlertPolicyAssociationsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetAlertPolicyAssociationsLifecycleStateEnumStringValues Enumerates the set of values in String for ListTargetAlertPolicyAssociationsLifecycleStateEnum
func GetListTargetAlertPolicyAssociationsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListTargetAlertPolicyAssociationsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetAlertPolicyAssociationsLifecycleStateEnum(val string) (ListTargetAlertPolicyAssociationsLifecycleStateEnum, bool) {
	enum, ok := mappingListTargetAlertPolicyAssociationsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetAlertPolicyAssociationsSortOrderEnum Enum with underlying type: string
type ListTargetAlertPolicyAssociationsSortOrderEnum string

// Set of constants representing the allowable values for ListTargetAlertPolicyAssociationsSortOrderEnum
const (
	ListTargetAlertPolicyAssociationsSortOrderAsc  ListTargetAlertPolicyAssociationsSortOrderEnum = "ASC"
	ListTargetAlertPolicyAssociationsSortOrderDesc ListTargetAlertPolicyAssociationsSortOrderEnum = "DESC"
)

var mappingListTargetAlertPolicyAssociationsSortOrderEnum = map[string]ListTargetAlertPolicyAssociationsSortOrderEnum{
	"ASC":  ListTargetAlertPolicyAssociationsSortOrderAsc,
	"DESC": ListTargetAlertPolicyAssociationsSortOrderDesc,
}

var mappingListTargetAlertPolicyAssociationsSortOrderEnumLowerCase = map[string]ListTargetAlertPolicyAssociationsSortOrderEnum{
	"asc":  ListTargetAlertPolicyAssociationsSortOrderAsc,
	"desc": ListTargetAlertPolicyAssociationsSortOrderDesc,
}

// GetListTargetAlertPolicyAssociationsSortOrderEnumValues Enumerates the set of values for ListTargetAlertPolicyAssociationsSortOrderEnum
func GetListTargetAlertPolicyAssociationsSortOrderEnumValues() []ListTargetAlertPolicyAssociationsSortOrderEnum {
	values := make([]ListTargetAlertPolicyAssociationsSortOrderEnum, 0)
	for _, v := range mappingListTargetAlertPolicyAssociationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetAlertPolicyAssociationsSortOrderEnumStringValues Enumerates the set of values in String for ListTargetAlertPolicyAssociationsSortOrderEnum
func GetListTargetAlertPolicyAssociationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetAlertPolicyAssociationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetAlertPolicyAssociationsSortOrderEnum(val string) (ListTargetAlertPolicyAssociationsSortOrderEnum, bool) {
	enum, ok := mappingListTargetAlertPolicyAssociationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetAlertPolicyAssociationsSortByEnum Enum with underlying type: string
type ListTargetAlertPolicyAssociationsSortByEnum string

// Set of constants representing the allowable values for ListTargetAlertPolicyAssociationsSortByEnum
const (
	ListTargetAlertPolicyAssociationsSortByDisplayname ListTargetAlertPolicyAssociationsSortByEnum = "DISPLAYNAME"
	ListTargetAlertPolicyAssociationsSortByTimecreated ListTargetAlertPolicyAssociationsSortByEnum = "TIMECREATED"
	ListTargetAlertPolicyAssociationsSortByTimeupdated ListTargetAlertPolicyAssociationsSortByEnum = "TIMEUPDATED"
)

var mappingListTargetAlertPolicyAssociationsSortByEnum = map[string]ListTargetAlertPolicyAssociationsSortByEnum{
	"DISPLAYNAME": ListTargetAlertPolicyAssociationsSortByDisplayname,
	"TIMECREATED": ListTargetAlertPolicyAssociationsSortByTimecreated,
	"TIMEUPDATED": ListTargetAlertPolicyAssociationsSortByTimeupdated,
}

var mappingListTargetAlertPolicyAssociationsSortByEnumLowerCase = map[string]ListTargetAlertPolicyAssociationsSortByEnum{
	"displayname": ListTargetAlertPolicyAssociationsSortByDisplayname,
	"timecreated": ListTargetAlertPolicyAssociationsSortByTimecreated,
	"timeupdated": ListTargetAlertPolicyAssociationsSortByTimeupdated,
}

// GetListTargetAlertPolicyAssociationsSortByEnumValues Enumerates the set of values for ListTargetAlertPolicyAssociationsSortByEnum
func GetListTargetAlertPolicyAssociationsSortByEnumValues() []ListTargetAlertPolicyAssociationsSortByEnum {
	values := make([]ListTargetAlertPolicyAssociationsSortByEnum, 0)
	for _, v := range mappingListTargetAlertPolicyAssociationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetAlertPolicyAssociationsSortByEnumStringValues Enumerates the set of values in String for ListTargetAlertPolicyAssociationsSortByEnum
func GetListTargetAlertPolicyAssociationsSortByEnumStringValues() []string {
	return []string{
		"DISPLAYNAME",
		"TIMECREATED",
		"TIMEUPDATED",
	}
}

// GetMappingListTargetAlertPolicyAssociationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetAlertPolicyAssociationsSortByEnum(val string) (ListTargetAlertPolicyAssociationsSortByEnum, bool) {
	enum, ok := mappingListTargetAlertPolicyAssociationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetAlertPolicyAssociationsAccessLevelEnum Enum with underlying type: string
type ListTargetAlertPolicyAssociationsAccessLevelEnum string

// Set of constants representing the allowable values for ListTargetAlertPolicyAssociationsAccessLevelEnum
const (
	ListTargetAlertPolicyAssociationsAccessLevelRestricted ListTargetAlertPolicyAssociationsAccessLevelEnum = "RESTRICTED"
	ListTargetAlertPolicyAssociationsAccessLevelAccessible ListTargetAlertPolicyAssociationsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListTargetAlertPolicyAssociationsAccessLevelEnum = map[string]ListTargetAlertPolicyAssociationsAccessLevelEnum{
	"RESTRICTED": ListTargetAlertPolicyAssociationsAccessLevelRestricted,
	"ACCESSIBLE": ListTargetAlertPolicyAssociationsAccessLevelAccessible,
}

var mappingListTargetAlertPolicyAssociationsAccessLevelEnumLowerCase = map[string]ListTargetAlertPolicyAssociationsAccessLevelEnum{
	"restricted": ListTargetAlertPolicyAssociationsAccessLevelRestricted,
	"accessible": ListTargetAlertPolicyAssociationsAccessLevelAccessible,
}

// GetListTargetAlertPolicyAssociationsAccessLevelEnumValues Enumerates the set of values for ListTargetAlertPolicyAssociationsAccessLevelEnum
func GetListTargetAlertPolicyAssociationsAccessLevelEnumValues() []ListTargetAlertPolicyAssociationsAccessLevelEnum {
	values := make([]ListTargetAlertPolicyAssociationsAccessLevelEnum, 0)
	for _, v := range mappingListTargetAlertPolicyAssociationsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetAlertPolicyAssociationsAccessLevelEnumStringValues Enumerates the set of values in String for ListTargetAlertPolicyAssociationsAccessLevelEnum
func GetListTargetAlertPolicyAssociationsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListTargetAlertPolicyAssociationsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetAlertPolicyAssociationsAccessLevelEnum(val string) (ListTargetAlertPolicyAssociationsAccessLevelEnum, bool) {
	enum, ok := mappingListTargetAlertPolicyAssociationsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
