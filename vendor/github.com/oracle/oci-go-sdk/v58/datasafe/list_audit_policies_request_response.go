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

// ListAuditPoliciesRequest wrapper for the ListAuditPolicies operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditPolicies.go.html to see an example of how to use ListAuditPoliciesRequest.
type ListAuditPoliciesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListAuditPoliciesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// The current state of the audit policy.
	LifecycleState ListAuditPoliciesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified id.
	AuditPolicyId *string `mandatory:"false" contributesTo:"query" name:"auditPolicyId"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListAuditPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListAuditPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAuditPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAuditPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAuditPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAuditPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAuditPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAuditPoliciesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListAuditPoliciesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditPoliciesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAuditPoliciesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAuditPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAuditPoliciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAuditPoliciesResponse wrapper for the ListAuditPolicies operation
type ListAuditPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AuditPolicyCollection instances
	AuditPolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListAuditPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAuditPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAuditPoliciesAccessLevelEnum Enum with underlying type: string
type ListAuditPoliciesAccessLevelEnum string

// Set of constants representing the allowable values for ListAuditPoliciesAccessLevelEnum
const (
	ListAuditPoliciesAccessLevelRestricted ListAuditPoliciesAccessLevelEnum = "RESTRICTED"
	ListAuditPoliciesAccessLevelAccessible ListAuditPoliciesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListAuditPoliciesAccessLevelEnum = map[string]ListAuditPoliciesAccessLevelEnum{
	"RESTRICTED": ListAuditPoliciesAccessLevelRestricted,
	"ACCESSIBLE": ListAuditPoliciesAccessLevelAccessible,
}

// GetListAuditPoliciesAccessLevelEnumValues Enumerates the set of values for ListAuditPoliciesAccessLevelEnum
func GetListAuditPoliciesAccessLevelEnumValues() []ListAuditPoliciesAccessLevelEnum {
	values := make([]ListAuditPoliciesAccessLevelEnum, 0)
	for _, v := range mappingListAuditPoliciesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditPoliciesAccessLevelEnumStringValues Enumerates the set of values in String for ListAuditPoliciesAccessLevelEnum
func GetListAuditPoliciesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListAuditPoliciesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditPoliciesAccessLevelEnum(val string) (ListAuditPoliciesAccessLevelEnum, bool) {
	mappingListAuditPoliciesAccessLevelEnumIgnoreCase := make(map[string]ListAuditPoliciesAccessLevelEnum)
	for k, v := range mappingListAuditPoliciesAccessLevelEnum {
		mappingListAuditPoliciesAccessLevelEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAuditPoliciesAccessLevelEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditPoliciesLifecycleStateEnum Enum with underlying type: string
type ListAuditPoliciesLifecycleStateEnum string

// Set of constants representing the allowable values for ListAuditPoliciesLifecycleStateEnum
const (
	ListAuditPoliciesLifecycleStateCreating       ListAuditPoliciesLifecycleStateEnum = "CREATING"
	ListAuditPoliciesLifecycleStateUpdating       ListAuditPoliciesLifecycleStateEnum = "UPDATING"
	ListAuditPoliciesLifecycleStateActive         ListAuditPoliciesLifecycleStateEnum = "ACTIVE"
	ListAuditPoliciesLifecycleStateFailed         ListAuditPoliciesLifecycleStateEnum = "FAILED"
	ListAuditPoliciesLifecycleStateNeedsAttention ListAuditPoliciesLifecycleStateEnum = "NEEDS_ATTENTION"
	ListAuditPoliciesLifecycleStateDeleting       ListAuditPoliciesLifecycleStateEnum = "DELETING"
	ListAuditPoliciesLifecycleStateDeleted        ListAuditPoliciesLifecycleStateEnum = "DELETED"
)

var mappingListAuditPoliciesLifecycleStateEnum = map[string]ListAuditPoliciesLifecycleStateEnum{
	"CREATING":        ListAuditPoliciesLifecycleStateCreating,
	"UPDATING":        ListAuditPoliciesLifecycleStateUpdating,
	"ACTIVE":          ListAuditPoliciesLifecycleStateActive,
	"FAILED":          ListAuditPoliciesLifecycleStateFailed,
	"NEEDS_ATTENTION": ListAuditPoliciesLifecycleStateNeedsAttention,
	"DELETING":        ListAuditPoliciesLifecycleStateDeleting,
	"DELETED":         ListAuditPoliciesLifecycleStateDeleted,
}

// GetListAuditPoliciesLifecycleStateEnumValues Enumerates the set of values for ListAuditPoliciesLifecycleStateEnum
func GetListAuditPoliciesLifecycleStateEnumValues() []ListAuditPoliciesLifecycleStateEnum {
	values := make([]ListAuditPoliciesLifecycleStateEnum, 0)
	for _, v := range mappingListAuditPoliciesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditPoliciesLifecycleStateEnumStringValues Enumerates the set of values in String for ListAuditPoliciesLifecycleStateEnum
func GetListAuditPoliciesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"FAILED",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
	}
}

// GetMappingListAuditPoliciesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditPoliciesLifecycleStateEnum(val string) (ListAuditPoliciesLifecycleStateEnum, bool) {
	mappingListAuditPoliciesLifecycleStateEnumIgnoreCase := make(map[string]ListAuditPoliciesLifecycleStateEnum)
	for k, v := range mappingListAuditPoliciesLifecycleStateEnum {
		mappingListAuditPoliciesLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAuditPoliciesLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditPoliciesSortOrderEnum Enum with underlying type: string
type ListAuditPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListAuditPoliciesSortOrderEnum
const (
	ListAuditPoliciesSortOrderAsc  ListAuditPoliciesSortOrderEnum = "ASC"
	ListAuditPoliciesSortOrderDesc ListAuditPoliciesSortOrderEnum = "DESC"
)

var mappingListAuditPoliciesSortOrderEnum = map[string]ListAuditPoliciesSortOrderEnum{
	"ASC":  ListAuditPoliciesSortOrderAsc,
	"DESC": ListAuditPoliciesSortOrderDesc,
}

// GetListAuditPoliciesSortOrderEnumValues Enumerates the set of values for ListAuditPoliciesSortOrderEnum
func GetListAuditPoliciesSortOrderEnumValues() []ListAuditPoliciesSortOrderEnum {
	values := make([]ListAuditPoliciesSortOrderEnum, 0)
	for _, v := range mappingListAuditPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListAuditPoliciesSortOrderEnum
func GetListAuditPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAuditPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditPoliciesSortOrderEnum(val string) (ListAuditPoliciesSortOrderEnum, bool) {
	mappingListAuditPoliciesSortOrderEnumIgnoreCase := make(map[string]ListAuditPoliciesSortOrderEnum)
	for k, v := range mappingListAuditPoliciesSortOrderEnum {
		mappingListAuditPoliciesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAuditPoliciesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditPoliciesSortByEnum Enum with underlying type: string
type ListAuditPoliciesSortByEnum string

// Set of constants representing the allowable values for ListAuditPoliciesSortByEnum
const (
	ListAuditPoliciesSortByTimecreated ListAuditPoliciesSortByEnum = "TIMECREATED"
	ListAuditPoliciesSortByDisplayname ListAuditPoliciesSortByEnum = "DISPLAYNAME"
)

var mappingListAuditPoliciesSortByEnum = map[string]ListAuditPoliciesSortByEnum{
	"TIMECREATED": ListAuditPoliciesSortByTimecreated,
	"DISPLAYNAME": ListAuditPoliciesSortByDisplayname,
}

// GetListAuditPoliciesSortByEnumValues Enumerates the set of values for ListAuditPoliciesSortByEnum
func GetListAuditPoliciesSortByEnumValues() []ListAuditPoliciesSortByEnum {
	values := make([]ListAuditPoliciesSortByEnum, 0)
	for _, v := range mappingListAuditPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditPoliciesSortByEnumStringValues Enumerates the set of values in String for ListAuditPoliciesSortByEnum
func GetListAuditPoliciesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAuditPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditPoliciesSortByEnum(val string) (ListAuditPoliciesSortByEnum, bool) {
	mappingListAuditPoliciesSortByEnumIgnoreCase := make(map[string]ListAuditPoliciesSortByEnum)
	for k, v := range mappingListAuditPoliciesSortByEnum {
		mappingListAuditPoliciesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAuditPoliciesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
