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

// ListSecurityPoliciesRequest wrapper for the ListSecurityPolicies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityPolicies.go.html to see an example of how to use ListSecurityPoliciesRequest.
type ListSecurityPoliciesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSecurityPoliciesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The current state of the security policy.
	LifecycleState ListSecurityPoliciesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified OCID of the security policy resource.
	SecurityPolicyId *string `mandatory:"false" contributesTo:"query" name:"securityPolicyId"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSecurityPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListSecurityPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecurityPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecurityPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSecurityPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecurityPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSecurityPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSecurityPoliciesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSecurityPoliciesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityPoliciesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSecurityPoliciesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSecurityPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSecurityPoliciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSecurityPoliciesResponse wrapper for the ListSecurityPolicies operation
type ListSecurityPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SecurityPolicyCollection instances
	SecurityPolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSecurityPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecurityPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecurityPoliciesAccessLevelEnum Enum with underlying type: string
type ListSecurityPoliciesAccessLevelEnum string

// Set of constants representing the allowable values for ListSecurityPoliciesAccessLevelEnum
const (
	ListSecurityPoliciesAccessLevelRestricted ListSecurityPoliciesAccessLevelEnum = "RESTRICTED"
	ListSecurityPoliciesAccessLevelAccessible ListSecurityPoliciesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSecurityPoliciesAccessLevelEnum = map[string]ListSecurityPoliciesAccessLevelEnum{
	"RESTRICTED": ListSecurityPoliciesAccessLevelRestricted,
	"ACCESSIBLE": ListSecurityPoliciesAccessLevelAccessible,
}

var mappingListSecurityPoliciesAccessLevelEnumLowerCase = map[string]ListSecurityPoliciesAccessLevelEnum{
	"restricted": ListSecurityPoliciesAccessLevelRestricted,
	"accessible": ListSecurityPoliciesAccessLevelAccessible,
}

// GetListSecurityPoliciesAccessLevelEnumValues Enumerates the set of values for ListSecurityPoliciesAccessLevelEnum
func GetListSecurityPoliciesAccessLevelEnumValues() []ListSecurityPoliciesAccessLevelEnum {
	values := make([]ListSecurityPoliciesAccessLevelEnum, 0)
	for _, v := range mappingListSecurityPoliciesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPoliciesAccessLevelEnumStringValues Enumerates the set of values in String for ListSecurityPoliciesAccessLevelEnum
func GetListSecurityPoliciesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSecurityPoliciesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPoliciesAccessLevelEnum(val string) (ListSecurityPoliciesAccessLevelEnum, bool) {
	enum, ok := mappingListSecurityPoliciesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityPoliciesLifecycleStateEnum Enum with underlying type: string
type ListSecurityPoliciesLifecycleStateEnum string

// Set of constants representing the allowable values for ListSecurityPoliciesLifecycleStateEnum
const (
	ListSecurityPoliciesLifecycleStateCreating ListSecurityPoliciesLifecycleStateEnum = "CREATING"
	ListSecurityPoliciesLifecycleStateUpdating ListSecurityPoliciesLifecycleStateEnum = "UPDATING"
	ListSecurityPoliciesLifecycleStateActive   ListSecurityPoliciesLifecycleStateEnum = "ACTIVE"
	ListSecurityPoliciesLifecycleStateFailed   ListSecurityPoliciesLifecycleStateEnum = "FAILED"
	ListSecurityPoliciesLifecycleStateDeleting ListSecurityPoliciesLifecycleStateEnum = "DELETING"
	ListSecurityPoliciesLifecycleStateDeleted  ListSecurityPoliciesLifecycleStateEnum = "DELETED"
)

var mappingListSecurityPoliciesLifecycleStateEnum = map[string]ListSecurityPoliciesLifecycleStateEnum{
	"CREATING": ListSecurityPoliciesLifecycleStateCreating,
	"UPDATING": ListSecurityPoliciesLifecycleStateUpdating,
	"ACTIVE":   ListSecurityPoliciesLifecycleStateActive,
	"FAILED":   ListSecurityPoliciesLifecycleStateFailed,
	"DELETING": ListSecurityPoliciesLifecycleStateDeleting,
	"DELETED":  ListSecurityPoliciesLifecycleStateDeleted,
}

var mappingListSecurityPoliciesLifecycleStateEnumLowerCase = map[string]ListSecurityPoliciesLifecycleStateEnum{
	"creating": ListSecurityPoliciesLifecycleStateCreating,
	"updating": ListSecurityPoliciesLifecycleStateUpdating,
	"active":   ListSecurityPoliciesLifecycleStateActive,
	"failed":   ListSecurityPoliciesLifecycleStateFailed,
	"deleting": ListSecurityPoliciesLifecycleStateDeleting,
	"deleted":  ListSecurityPoliciesLifecycleStateDeleted,
}

// GetListSecurityPoliciesLifecycleStateEnumValues Enumerates the set of values for ListSecurityPoliciesLifecycleStateEnum
func GetListSecurityPoliciesLifecycleStateEnumValues() []ListSecurityPoliciesLifecycleStateEnum {
	values := make([]ListSecurityPoliciesLifecycleStateEnum, 0)
	for _, v := range mappingListSecurityPoliciesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPoliciesLifecycleStateEnumStringValues Enumerates the set of values in String for ListSecurityPoliciesLifecycleStateEnum
func GetListSecurityPoliciesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingListSecurityPoliciesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPoliciesLifecycleStateEnum(val string) (ListSecurityPoliciesLifecycleStateEnum, bool) {
	enum, ok := mappingListSecurityPoliciesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityPoliciesSortOrderEnum Enum with underlying type: string
type ListSecurityPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListSecurityPoliciesSortOrderEnum
const (
	ListSecurityPoliciesSortOrderAsc  ListSecurityPoliciesSortOrderEnum = "ASC"
	ListSecurityPoliciesSortOrderDesc ListSecurityPoliciesSortOrderEnum = "DESC"
)

var mappingListSecurityPoliciesSortOrderEnum = map[string]ListSecurityPoliciesSortOrderEnum{
	"ASC":  ListSecurityPoliciesSortOrderAsc,
	"DESC": ListSecurityPoliciesSortOrderDesc,
}

var mappingListSecurityPoliciesSortOrderEnumLowerCase = map[string]ListSecurityPoliciesSortOrderEnum{
	"asc":  ListSecurityPoliciesSortOrderAsc,
	"desc": ListSecurityPoliciesSortOrderDesc,
}

// GetListSecurityPoliciesSortOrderEnumValues Enumerates the set of values for ListSecurityPoliciesSortOrderEnum
func GetListSecurityPoliciesSortOrderEnumValues() []ListSecurityPoliciesSortOrderEnum {
	values := make([]ListSecurityPoliciesSortOrderEnum, 0)
	for _, v := range mappingListSecurityPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListSecurityPoliciesSortOrderEnum
func GetListSecurityPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSecurityPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPoliciesSortOrderEnum(val string) (ListSecurityPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListSecurityPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityPoliciesSortByEnum Enum with underlying type: string
type ListSecurityPoliciesSortByEnum string

// Set of constants representing the allowable values for ListSecurityPoliciesSortByEnum
const (
	ListSecurityPoliciesSortByTimecreated ListSecurityPoliciesSortByEnum = "TIMECREATED"
	ListSecurityPoliciesSortByDisplayname ListSecurityPoliciesSortByEnum = "DISPLAYNAME"
)

var mappingListSecurityPoliciesSortByEnum = map[string]ListSecurityPoliciesSortByEnum{
	"TIMECREATED": ListSecurityPoliciesSortByTimecreated,
	"DISPLAYNAME": ListSecurityPoliciesSortByDisplayname,
}

var mappingListSecurityPoliciesSortByEnumLowerCase = map[string]ListSecurityPoliciesSortByEnum{
	"timecreated": ListSecurityPoliciesSortByTimecreated,
	"displayname": ListSecurityPoliciesSortByDisplayname,
}

// GetListSecurityPoliciesSortByEnumValues Enumerates the set of values for ListSecurityPoliciesSortByEnum
func GetListSecurityPoliciesSortByEnumValues() []ListSecurityPoliciesSortByEnum {
	values := make([]ListSecurityPoliciesSortByEnum, 0)
	for _, v := range mappingListSecurityPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPoliciesSortByEnumStringValues Enumerates the set of values in String for ListSecurityPoliciesSortByEnum
func GetListSecurityPoliciesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListSecurityPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPoliciesSortByEnum(val string) (ListSecurityPoliciesSortByEnum, bool) {
	enum, ok := mappingListSecurityPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
