// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSecurityPolicyDeploymentsRequest wrapper for the ListSecurityPolicyDeployments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityPolicyDeployments.go.html to see an example of how to use ListSecurityPolicyDeploymentsRequest.
type ListSecurityPolicyDeploymentsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSecurityPolicyDeploymentsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The current state of the security policy deployment.
	LifecycleState ListSecurityPolicyDeploymentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified OCID of the security policy deployment resource.
	SecurityPolicyDeploymentId *string `mandatory:"false" contributesTo:"query" name:"securityPolicyDeploymentId"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// An optional filter to return only resources that match the specified OCID of the security policy resource.
	SecurityPolicyId *string `mandatory:"false" contributesTo:"query" name:"securityPolicyId"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSecurityPolicyDeploymentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListSecurityPolicyDeploymentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecurityPolicyDeploymentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecurityPolicyDeploymentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSecurityPolicyDeploymentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecurityPolicyDeploymentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSecurityPolicyDeploymentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSecurityPolicyDeploymentsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSecurityPolicyDeploymentsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityPolicyDeploymentsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSecurityPolicyDeploymentsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityPolicyDeploymentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSecurityPolicyDeploymentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityPolicyDeploymentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSecurityPolicyDeploymentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSecurityPolicyDeploymentsResponse wrapper for the ListSecurityPolicyDeployments operation
type ListSecurityPolicyDeploymentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SecurityPolicyDeploymentCollection instances
	SecurityPolicyDeploymentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSecurityPolicyDeploymentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecurityPolicyDeploymentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecurityPolicyDeploymentsAccessLevelEnum Enum with underlying type: string
type ListSecurityPolicyDeploymentsAccessLevelEnum string

// Set of constants representing the allowable values for ListSecurityPolicyDeploymentsAccessLevelEnum
const (
	ListSecurityPolicyDeploymentsAccessLevelRestricted ListSecurityPolicyDeploymentsAccessLevelEnum = "RESTRICTED"
	ListSecurityPolicyDeploymentsAccessLevelAccessible ListSecurityPolicyDeploymentsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSecurityPolicyDeploymentsAccessLevelEnum = map[string]ListSecurityPolicyDeploymentsAccessLevelEnum{
	"RESTRICTED": ListSecurityPolicyDeploymentsAccessLevelRestricted,
	"ACCESSIBLE": ListSecurityPolicyDeploymentsAccessLevelAccessible,
}

var mappingListSecurityPolicyDeploymentsAccessLevelEnumLowerCase = map[string]ListSecurityPolicyDeploymentsAccessLevelEnum{
	"restricted": ListSecurityPolicyDeploymentsAccessLevelRestricted,
	"accessible": ListSecurityPolicyDeploymentsAccessLevelAccessible,
}

// GetListSecurityPolicyDeploymentsAccessLevelEnumValues Enumerates the set of values for ListSecurityPolicyDeploymentsAccessLevelEnum
func GetListSecurityPolicyDeploymentsAccessLevelEnumValues() []ListSecurityPolicyDeploymentsAccessLevelEnum {
	values := make([]ListSecurityPolicyDeploymentsAccessLevelEnum, 0)
	for _, v := range mappingListSecurityPolicyDeploymentsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPolicyDeploymentsAccessLevelEnumStringValues Enumerates the set of values in String for ListSecurityPolicyDeploymentsAccessLevelEnum
func GetListSecurityPolicyDeploymentsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSecurityPolicyDeploymentsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPolicyDeploymentsAccessLevelEnum(val string) (ListSecurityPolicyDeploymentsAccessLevelEnum, bool) {
	enum, ok := mappingListSecurityPolicyDeploymentsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityPolicyDeploymentsLifecycleStateEnum Enum with underlying type: string
type ListSecurityPolicyDeploymentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSecurityPolicyDeploymentsLifecycleStateEnum
const (
	ListSecurityPolicyDeploymentsLifecycleStateCreating       ListSecurityPolicyDeploymentsLifecycleStateEnum = "CREATING"
	ListSecurityPolicyDeploymentsLifecycleStateUpdating       ListSecurityPolicyDeploymentsLifecycleStateEnum = "UPDATING"
	ListSecurityPolicyDeploymentsLifecycleStateDeployed       ListSecurityPolicyDeploymentsLifecycleStateEnum = "DEPLOYED"
	ListSecurityPolicyDeploymentsLifecycleStateNeedsAttention ListSecurityPolicyDeploymentsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListSecurityPolicyDeploymentsLifecycleStateFailed         ListSecurityPolicyDeploymentsLifecycleStateEnum = "FAILED"
	ListSecurityPolicyDeploymentsLifecycleStateDeleting       ListSecurityPolicyDeploymentsLifecycleStateEnum = "DELETING"
	ListSecurityPolicyDeploymentsLifecycleStateDeleted        ListSecurityPolicyDeploymentsLifecycleStateEnum = "DELETED"
)

var mappingListSecurityPolicyDeploymentsLifecycleStateEnum = map[string]ListSecurityPolicyDeploymentsLifecycleStateEnum{
	"CREATING":        ListSecurityPolicyDeploymentsLifecycleStateCreating,
	"UPDATING":        ListSecurityPolicyDeploymentsLifecycleStateUpdating,
	"DEPLOYED":        ListSecurityPolicyDeploymentsLifecycleStateDeployed,
	"NEEDS_ATTENTION": ListSecurityPolicyDeploymentsLifecycleStateNeedsAttention,
	"FAILED":          ListSecurityPolicyDeploymentsLifecycleStateFailed,
	"DELETING":        ListSecurityPolicyDeploymentsLifecycleStateDeleting,
	"DELETED":         ListSecurityPolicyDeploymentsLifecycleStateDeleted,
}

var mappingListSecurityPolicyDeploymentsLifecycleStateEnumLowerCase = map[string]ListSecurityPolicyDeploymentsLifecycleStateEnum{
	"creating":        ListSecurityPolicyDeploymentsLifecycleStateCreating,
	"updating":        ListSecurityPolicyDeploymentsLifecycleStateUpdating,
	"deployed":        ListSecurityPolicyDeploymentsLifecycleStateDeployed,
	"needs_attention": ListSecurityPolicyDeploymentsLifecycleStateNeedsAttention,
	"failed":          ListSecurityPolicyDeploymentsLifecycleStateFailed,
	"deleting":        ListSecurityPolicyDeploymentsLifecycleStateDeleting,
	"deleted":         ListSecurityPolicyDeploymentsLifecycleStateDeleted,
}

// GetListSecurityPolicyDeploymentsLifecycleStateEnumValues Enumerates the set of values for ListSecurityPolicyDeploymentsLifecycleStateEnum
func GetListSecurityPolicyDeploymentsLifecycleStateEnumValues() []ListSecurityPolicyDeploymentsLifecycleStateEnum {
	values := make([]ListSecurityPolicyDeploymentsLifecycleStateEnum, 0)
	for _, v := range mappingListSecurityPolicyDeploymentsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPolicyDeploymentsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSecurityPolicyDeploymentsLifecycleStateEnum
func GetListSecurityPolicyDeploymentsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"DEPLOYED",
		"NEEDS_ATTENTION",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingListSecurityPolicyDeploymentsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPolicyDeploymentsLifecycleStateEnum(val string) (ListSecurityPolicyDeploymentsLifecycleStateEnum, bool) {
	enum, ok := mappingListSecurityPolicyDeploymentsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityPolicyDeploymentsSortOrderEnum Enum with underlying type: string
type ListSecurityPolicyDeploymentsSortOrderEnum string

// Set of constants representing the allowable values for ListSecurityPolicyDeploymentsSortOrderEnum
const (
	ListSecurityPolicyDeploymentsSortOrderAsc  ListSecurityPolicyDeploymentsSortOrderEnum = "ASC"
	ListSecurityPolicyDeploymentsSortOrderDesc ListSecurityPolicyDeploymentsSortOrderEnum = "DESC"
)

var mappingListSecurityPolicyDeploymentsSortOrderEnum = map[string]ListSecurityPolicyDeploymentsSortOrderEnum{
	"ASC":  ListSecurityPolicyDeploymentsSortOrderAsc,
	"DESC": ListSecurityPolicyDeploymentsSortOrderDesc,
}

var mappingListSecurityPolicyDeploymentsSortOrderEnumLowerCase = map[string]ListSecurityPolicyDeploymentsSortOrderEnum{
	"asc":  ListSecurityPolicyDeploymentsSortOrderAsc,
	"desc": ListSecurityPolicyDeploymentsSortOrderDesc,
}

// GetListSecurityPolicyDeploymentsSortOrderEnumValues Enumerates the set of values for ListSecurityPolicyDeploymentsSortOrderEnum
func GetListSecurityPolicyDeploymentsSortOrderEnumValues() []ListSecurityPolicyDeploymentsSortOrderEnum {
	values := make([]ListSecurityPolicyDeploymentsSortOrderEnum, 0)
	for _, v := range mappingListSecurityPolicyDeploymentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPolicyDeploymentsSortOrderEnumStringValues Enumerates the set of values in String for ListSecurityPolicyDeploymentsSortOrderEnum
func GetListSecurityPolicyDeploymentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSecurityPolicyDeploymentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPolicyDeploymentsSortOrderEnum(val string) (ListSecurityPolicyDeploymentsSortOrderEnum, bool) {
	enum, ok := mappingListSecurityPolicyDeploymentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityPolicyDeploymentsSortByEnum Enum with underlying type: string
type ListSecurityPolicyDeploymentsSortByEnum string

// Set of constants representing the allowable values for ListSecurityPolicyDeploymentsSortByEnum
const (
	ListSecurityPolicyDeploymentsSortByTimecreated ListSecurityPolicyDeploymentsSortByEnum = "TIMECREATED"
	ListSecurityPolicyDeploymentsSortByDisplayname ListSecurityPolicyDeploymentsSortByEnum = "DISPLAYNAME"
)

var mappingListSecurityPolicyDeploymentsSortByEnum = map[string]ListSecurityPolicyDeploymentsSortByEnum{
	"TIMECREATED": ListSecurityPolicyDeploymentsSortByTimecreated,
	"DISPLAYNAME": ListSecurityPolicyDeploymentsSortByDisplayname,
}

var mappingListSecurityPolicyDeploymentsSortByEnumLowerCase = map[string]ListSecurityPolicyDeploymentsSortByEnum{
	"timecreated": ListSecurityPolicyDeploymentsSortByTimecreated,
	"displayname": ListSecurityPolicyDeploymentsSortByDisplayname,
}

// GetListSecurityPolicyDeploymentsSortByEnumValues Enumerates the set of values for ListSecurityPolicyDeploymentsSortByEnum
func GetListSecurityPolicyDeploymentsSortByEnumValues() []ListSecurityPolicyDeploymentsSortByEnum {
	values := make([]ListSecurityPolicyDeploymentsSortByEnum, 0)
	for _, v := range mappingListSecurityPolicyDeploymentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPolicyDeploymentsSortByEnumStringValues Enumerates the set of values in String for ListSecurityPolicyDeploymentsSortByEnum
func GetListSecurityPolicyDeploymentsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListSecurityPolicyDeploymentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPolicyDeploymentsSortByEnum(val string) (ListSecurityPolicyDeploymentsSortByEnum, bool) {
	enum, ok := mappingListSecurityPolicyDeploymentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
