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

// ListSqlFirewallPolicyAnalyticsRequest wrapper for the ListSqlFirewallPolicyAnalytics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlFirewallPolicyAnalytics.go.html to see an example of how to use ListSqlFirewallPolicyAnalyticsRequest.
type ListSqlFirewallPolicyAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSqlFirewallPolicyAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The group by parameter to summarize SQL Firewall policy aggregation.
	GroupBy []ListSqlFirewallPolicyAnalyticsGroupByEnum `contributesTo:"query" name:"groupBy" omitEmpty:"true" collectionFormat:"multi"`

	// The current state of the SQL Firewall policy.
	LifecycleState ListSqlFirewallPolicyAnalyticsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified OCID of the security policy resource.
	SecurityPolicyId *string `mandatory:"false" contributesTo:"query" name:"securityPolicyId"`

	// An optional filter to return the summary of the SQL Firewall policies created after the date-time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStarted"`

	// An optional filter to return the summary of the SQL Firewall policies created before the date-time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeEnded *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnded"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlFirewallPolicyAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlFirewallPolicyAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlFirewallPolicyAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlFirewallPolicyAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlFirewallPolicyAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlFirewallPolicyAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSqlFirewallPolicyAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	for _, val := range request.GroupBy {
		if _, ok := GetMappingListSqlFirewallPolicyAnalyticsGroupByEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", val, strings.Join(GetListSqlFirewallPolicyAnalyticsGroupByEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListSqlFirewallPolicyAnalyticsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSqlFirewallPolicyAnalyticsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlFirewallPolicyAnalyticsResponse wrapper for the ListSqlFirewallPolicyAnalytics operation
type ListSqlFirewallPolicyAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlFirewallPolicyAnalyticsCollection instances
	SqlFirewallPolicyAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSqlFirewallPolicyAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlFirewallPolicyAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlFirewallPolicyAnalyticsAccessLevelEnum Enum with underlying type: string
type ListSqlFirewallPolicyAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListSqlFirewallPolicyAnalyticsAccessLevelEnum
const (
	ListSqlFirewallPolicyAnalyticsAccessLevelRestricted ListSqlFirewallPolicyAnalyticsAccessLevelEnum = "RESTRICTED"
	ListSqlFirewallPolicyAnalyticsAccessLevelAccessible ListSqlFirewallPolicyAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSqlFirewallPolicyAnalyticsAccessLevelEnum = map[string]ListSqlFirewallPolicyAnalyticsAccessLevelEnum{
	"RESTRICTED": ListSqlFirewallPolicyAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListSqlFirewallPolicyAnalyticsAccessLevelAccessible,
}

var mappingListSqlFirewallPolicyAnalyticsAccessLevelEnumLowerCase = map[string]ListSqlFirewallPolicyAnalyticsAccessLevelEnum{
	"restricted": ListSqlFirewallPolicyAnalyticsAccessLevelRestricted,
	"accessible": ListSqlFirewallPolicyAnalyticsAccessLevelAccessible,
}

// GetListSqlFirewallPolicyAnalyticsAccessLevelEnumValues Enumerates the set of values for ListSqlFirewallPolicyAnalyticsAccessLevelEnum
func GetListSqlFirewallPolicyAnalyticsAccessLevelEnumValues() []ListSqlFirewallPolicyAnalyticsAccessLevelEnum {
	values := make([]ListSqlFirewallPolicyAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListSqlFirewallPolicyAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallPolicyAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListSqlFirewallPolicyAnalyticsAccessLevelEnum
func GetListSqlFirewallPolicyAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSqlFirewallPolicyAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallPolicyAnalyticsAccessLevelEnum(val string) (ListSqlFirewallPolicyAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListSqlFirewallPolicyAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallPolicyAnalyticsGroupByEnum Enum with underlying type: string
type ListSqlFirewallPolicyAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListSqlFirewallPolicyAnalyticsGroupByEnum
const (
	ListSqlFirewallPolicyAnalyticsGroupByViolationaction  ListSqlFirewallPolicyAnalyticsGroupByEnum = "violationAction"
	ListSqlFirewallPolicyAnalyticsGroupByEnforcementscope ListSqlFirewallPolicyAnalyticsGroupByEnum = "enforcementScope"
	ListSqlFirewallPolicyAnalyticsGroupBySecuritypolicyid ListSqlFirewallPolicyAnalyticsGroupByEnum = "securityPolicyId"
	ListSqlFirewallPolicyAnalyticsGroupByLifecyclestate   ListSqlFirewallPolicyAnalyticsGroupByEnum = "lifecycleState"
)

var mappingListSqlFirewallPolicyAnalyticsGroupByEnum = map[string]ListSqlFirewallPolicyAnalyticsGroupByEnum{
	"violationAction":  ListSqlFirewallPolicyAnalyticsGroupByViolationaction,
	"enforcementScope": ListSqlFirewallPolicyAnalyticsGroupByEnforcementscope,
	"securityPolicyId": ListSqlFirewallPolicyAnalyticsGroupBySecuritypolicyid,
	"lifecycleState":   ListSqlFirewallPolicyAnalyticsGroupByLifecyclestate,
}

var mappingListSqlFirewallPolicyAnalyticsGroupByEnumLowerCase = map[string]ListSqlFirewallPolicyAnalyticsGroupByEnum{
	"violationaction":  ListSqlFirewallPolicyAnalyticsGroupByViolationaction,
	"enforcementscope": ListSqlFirewallPolicyAnalyticsGroupByEnforcementscope,
	"securitypolicyid": ListSqlFirewallPolicyAnalyticsGroupBySecuritypolicyid,
	"lifecyclestate":   ListSqlFirewallPolicyAnalyticsGroupByLifecyclestate,
}

// GetListSqlFirewallPolicyAnalyticsGroupByEnumValues Enumerates the set of values for ListSqlFirewallPolicyAnalyticsGroupByEnum
func GetListSqlFirewallPolicyAnalyticsGroupByEnumValues() []ListSqlFirewallPolicyAnalyticsGroupByEnum {
	values := make([]ListSqlFirewallPolicyAnalyticsGroupByEnum, 0)
	for _, v := range mappingListSqlFirewallPolicyAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallPolicyAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListSqlFirewallPolicyAnalyticsGroupByEnum
func GetListSqlFirewallPolicyAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"violationAction",
		"enforcementScope",
		"securityPolicyId",
		"lifecycleState",
	}
}

// GetMappingListSqlFirewallPolicyAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallPolicyAnalyticsGroupByEnum(val string) (ListSqlFirewallPolicyAnalyticsGroupByEnum, bool) {
	enum, ok := mappingListSqlFirewallPolicyAnalyticsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallPolicyAnalyticsLifecycleStateEnum Enum with underlying type: string
type ListSqlFirewallPolicyAnalyticsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSqlFirewallPolicyAnalyticsLifecycleStateEnum
const (
	ListSqlFirewallPolicyAnalyticsLifecycleStateCreating       ListSqlFirewallPolicyAnalyticsLifecycleStateEnum = "CREATING"
	ListSqlFirewallPolicyAnalyticsLifecycleStateUpdating       ListSqlFirewallPolicyAnalyticsLifecycleStateEnum = "UPDATING"
	ListSqlFirewallPolicyAnalyticsLifecycleStateActive         ListSqlFirewallPolicyAnalyticsLifecycleStateEnum = "ACTIVE"
	ListSqlFirewallPolicyAnalyticsLifecycleStateInactive       ListSqlFirewallPolicyAnalyticsLifecycleStateEnum = "INACTIVE"
	ListSqlFirewallPolicyAnalyticsLifecycleStateFailed         ListSqlFirewallPolicyAnalyticsLifecycleStateEnum = "FAILED"
	ListSqlFirewallPolicyAnalyticsLifecycleStateDeleting       ListSqlFirewallPolicyAnalyticsLifecycleStateEnum = "DELETING"
	ListSqlFirewallPolicyAnalyticsLifecycleStateDeleted        ListSqlFirewallPolicyAnalyticsLifecycleStateEnum = "DELETED"
	ListSqlFirewallPolicyAnalyticsLifecycleStateNeedsAttention ListSqlFirewallPolicyAnalyticsLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListSqlFirewallPolicyAnalyticsLifecycleStateEnum = map[string]ListSqlFirewallPolicyAnalyticsLifecycleStateEnum{
	"CREATING":        ListSqlFirewallPolicyAnalyticsLifecycleStateCreating,
	"UPDATING":        ListSqlFirewallPolicyAnalyticsLifecycleStateUpdating,
	"ACTIVE":          ListSqlFirewallPolicyAnalyticsLifecycleStateActive,
	"INACTIVE":        ListSqlFirewallPolicyAnalyticsLifecycleStateInactive,
	"FAILED":          ListSqlFirewallPolicyAnalyticsLifecycleStateFailed,
	"DELETING":        ListSqlFirewallPolicyAnalyticsLifecycleStateDeleting,
	"DELETED":         ListSqlFirewallPolicyAnalyticsLifecycleStateDeleted,
	"NEEDS_ATTENTION": ListSqlFirewallPolicyAnalyticsLifecycleStateNeedsAttention,
}

var mappingListSqlFirewallPolicyAnalyticsLifecycleStateEnumLowerCase = map[string]ListSqlFirewallPolicyAnalyticsLifecycleStateEnum{
	"creating":        ListSqlFirewallPolicyAnalyticsLifecycleStateCreating,
	"updating":        ListSqlFirewallPolicyAnalyticsLifecycleStateUpdating,
	"active":          ListSqlFirewallPolicyAnalyticsLifecycleStateActive,
	"inactive":        ListSqlFirewallPolicyAnalyticsLifecycleStateInactive,
	"failed":          ListSqlFirewallPolicyAnalyticsLifecycleStateFailed,
	"deleting":        ListSqlFirewallPolicyAnalyticsLifecycleStateDeleting,
	"deleted":         ListSqlFirewallPolicyAnalyticsLifecycleStateDeleted,
	"needs_attention": ListSqlFirewallPolicyAnalyticsLifecycleStateNeedsAttention,
}

// GetListSqlFirewallPolicyAnalyticsLifecycleStateEnumValues Enumerates the set of values for ListSqlFirewallPolicyAnalyticsLifecycleStateEnum
func GetListSqlFirewallPolicyAnalyticsLifecycleStateEnumValues() []ListSqlFirewallPolicyAnalyticsLifecycleStateEnum {
	values := make([]ListSqlFirewallPolicyAnalyticsLifecycleStateEnum, 0)
	for _, v := range mappingListSqlFirewallPolicyAnalyticsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallPolicyAnalyticsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSqlFirewallPolicyAnalyticsLifecycleStateEnum
func GetListSqlFirewallPolicyAnalyticsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"FAILED",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListSqlFirewallPolicyAnalyticsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallPolicyAnalyticsLifecycleStateEnum(val string) (ListSqlFirewallPolicyAnalyticsLifecycleStateEnum, bool) {
	enum, ok := mappingListSqlFirewallPolicyAnalyticsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
