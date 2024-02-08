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

// ListAuditPolicyAnalyticsRequest wrapper for the ListAuditPolicyAnalytics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditPolicyAnalytics.go.html to see an example of how to use ListAuditPolicyAnalyticsRequest.
type ListAuditPolicyAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListAuditPolicyAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The group by parameter to summarize audit policy aggregation.
	GroupBy []ListAuditPolicyAnalyticsGroupByEnum `contributesTo:"query" name:"groupBy" omitEmpty:"true" collectionFormat:"multi"`

	// The category to which the audit policy belongs to.
	AuditPolicyCategory ListAuditPolicyAnalyticsAuditPolicyCategoryEnum `mandatory:"false" contributesTo:"query" name:"auditPolicyCategory" omitEmpty:"true"`

	// In case of seeded policies, it is the policy name defined by Data Safe.
	// In case of custom Policies, it is the policy name that is used to create the policies on the target database.
	// In case of Oracle Pre-seeded policies, it is the default policy name of the same.
	AuditPolicyName *string `mandatory:"false" contributesTo:"query" name:"auditPolicyName"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// The current state of the audit policy.
	LifecycleState ListAuditPolicyAnalyticsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAuditPolicyAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAuditPolicyAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAuditPolicyAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAuditPolicyAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAuditPolicyAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAuditPolicyAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListAuditPolicyAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	for _, val := range request.GroupBy {
		if _, ok := GetMappingListAuditPolicyAnalyticsGroupByEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", val, strings.Join(GetListAuditPolicyAnalyticsGroupByEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListAuditPolicyAnalyticsAuditPolicyCategoryEnum(string(request.AuditPolicyCategory)); !ok && request.AuditPolicyCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuditPolicyCategory: %s. Supported values are: %s.", request.AuditPolicyCategory, strings.Join(GetListAuditPolicyAnalyticsAuditPolicyCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditPolicyAnalyticsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAuditPolicyAnalyticsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAuditPolicyAnalyticsResponse wrapper for the ListAuditPolicyAnalytics operation
type ListAuditPolicyAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AuditPolicyAnalyticCollection instances
	AuditPolicyAnalyticCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListAuditPolicyAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAuditPolicyAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAuditPolicyAnalyticsAccessLevelEnum Enum with underlying type: string
type ListAuditPolicyAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListAuditPolicyAnalyticsAccessLevelEnum
const (
	ListAuditPolicyAnalyticsAccessLevelRestricted ListAuditPolicyAnalyticsAccessLevelEnum = "RESTRICTED"
	ListAuditPolicyAnalyticsAccessLevelAccessible ListAuditPolicyAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListAuditPolicyAnalyticsAccessLevelEnum = map[string]ListAuditPolicyAnalyticsAccessLevelEnum{
	"RESTRICTED": ListAuditPolicyAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListAuditPolicyAnalyticsAccessLevelAccessible,
}

var mappingListAuditPolicyAnalyticsAccessLevelEnumLowerCase = map[string]ListAuditPolicyAnalyticsAccessLevelEnum{
	"restricted": ListAuditPolicyAnalyticsAccessLevelRestricted,
	"accessible": ListAuditPolicyAnalyticsAccessLevelAccessible,
}

// GetListAuditPolicyAnalyticsAccessLevelEnumValues Enumerates the set of values for ListAuditPolicyAnalyticsAccessLevelEnum
func GetListAuditPolicyAnalyticsAccessLevelEnumValues() []ListAuditPolicyAnalyticsAccessLevelEnum {
	values := make([]ListAuditPolicyAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListAuditPolicyAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditPolicyAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListAuditPolicyAnalyticsAccessLevelEnum
func GetListAuditPolicyAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListAuditPolicyAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditPolicyAnalyticsAccessLevelEnum(val string) (ListAuditPolicyAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListAuditPolicyAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditPolicyAnalyticsGroupByEnum Enum with underlying type: string
type ListAuditPolicyAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListAuditPolicyAnalyticsGroupByEnum
const (
	ListAuditPolicyAnalyticsGroupByAuditpolicycategory ListAuditPolicyAnalyticsGroupByEnum = "auditPolicyCategory"
	ListAuditPolicyAnalyticsGroupByAuditpolicyname     ListAuditPolicyAnalyticsGroupByEnum = "auditPolicyName"
	ListAuditPolicyAnalyticsGroupByTargetid            ListAuditPolicyAnalyticsGroupByEnum = "targetId"
)

var mappingListAuditPolicyAnalyticsGroupByEnum = map[string]ListAuditPolicyAnalyticsGroupByEnum{
	"auditPolicyCategory": ListAuditPolicyAnalyticsGroupByAuditpolicycategory,
	"auditPolicyName":     ListAuditPolicyAnalyticsGroupByAuditpolicyname,
	"targetId":            ListAuditPolicyAnalyticsGroupByTargetid,
}

var mappingListAuditPolicyAnalyticsGroupByEnumLowerCase = map[string]ListAuditPolicyAnalyticsGroupByEnum{
	"auditpolicycategory": ListAuditPolicyAnalyticsGroupByAuditpolicycategory,
	"auditpolicyname":     ListAuditPolicyAnalyticsGroupByAuditpolicyname,
	"targetid":            ListAuditPolicyAnalyticsGroupByTargetid,
}

// GetListAuditPolicyAnalyticsGroupByEnumValues Enumerates the set of values for ListAuditPolicyAnalyticsGroupByEnum
func GetListAuditPolicyAnalyticsGroupByEnumValues() []ListAuditPolicyAnalyticsGroupByEnum {
	values := make([]ListAuditPolicyAnalyticsGroupByEnum, 0)
	for _, v := range mappingListAuditPolicyAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditPolicyAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListAuditPolicyAnalyticsGroupByEnum
func GetListAuditPolicyAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"auditPolicyCategory",
		"auditPolicyName",
		"targetId",
	}
}

// GetMappingListAuditPolicyAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditPolicyAnalyticsGroupByEnum(val string) (ListAuditPolicyAnalyticsGroupByEnum, bool) {
	enum, ok := mappingListAuditPolicyAnalyticsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditPolicyAnalyticsAuditPolicyCategoryEnum Enum with underlying type: string
type ListAuditPolicyAnalyticsAuditPolicyCategoryEnum string

// Set of constants representing the allowable values for ListAuditPolicyAnalyticsAuditPolicyCategoryEnum
const (
	ListAuditPolicyAnalyticsAuditPolicyCategoryBasicActivity       ListAuditPolicyAnalyticsAuditPolicyCategoryEnum = "BASIC_ACTIVITY"
	ListAuditPolicyAnalyticsAuditPolicyCategoryAdminUserActivity   ListAuditPolicyAnalyticsAuditPolicyCategoryEnum = "ADMIN_USER_ACTIVITY"
	ListAuditPolicyAnalyticsAuditPolicyCategoryUserActivity        ListAuditPolicyAnalyticsAuditPolicyCategoryEnum = "USER_ACTIVITY"
	ListAuditPolicyAnalyticsAuditPolicyCategoryOraclePredefined    ListAuditPolicyAnalyticsAuditPolicyCategoryEnum = "ORACLE_PREDEFINED"
	ListAuditPolicyAnalyticsAuditPolicyCategoryComplianceStandard  ListAuditPolicyAnalyticsAuditPolicyCategoryEnum = "COMPLIANCE_STANDARD"
	ListAuditPolicyAnalyticsAuditPolicyCategoryCustom              ListAuditPolicyAnalyticsAuditPolicyCategoryEnum = "CUSTOM"
	ListAuditPolicyAnalyticsAuditPolicyCategorySqlFirewallAuditing ListAuditPolicyAnalyticsAuditPolicyCategoryEnum = "SQL_FIREWALL_AUDITING"
)

var mappingListAuditPolicyAnalyticsAuditPolicyCategoryEnum = map[string]ListAuditPolicyAnalyticsAuditPolicyCategoryEnum{
	"BASIC_ACTIVITY":        ListAuditPolicyAnalyticsAuditPolicyCategoryBasicActivity,
	"ADMIN_USER_ACTIVITY":   ListAuditPolicyAnalyticsAuditPolicyCategoryAdminUserActivity,
	"USER_ACTIVITY":         ListAuditPolicyAnalyticsAuditPolicyCategoryUserActivity,
	"ORACLE_PREDEFINED":     ListAuditPolicyAnalyticsAuditPolicyCategoryOraclePredefined,
	"COMPLIANCE_STANDARD":   ListAuditPolicyAnalyticsAuditPolicyCategoryComplianceStandard,
	"CUSTOM":                ListAuditPolicyAnalyticsAuditPolicyCategoryCustom,
	"SQL_FIREWALL_AUDITING": ListAuditPolicyAnalyticsAuditPolicyCategorySqlFirewallAuditing,
}

var mappingListAuditPolicyAnalyticsAuditPolicyCategoryEnumLowerCase = map[string]ListAuditPolicyAnalyticsAuditPolicyCategoryEnum{
	"basic_activity":        ListAuditPolicyAnalyticsAuditPolicyCategoryBasicActivity,
	"admin_user_activity":   ListAuditPolicyAnalyticsAuditPolicyCategoryAdminUserActivity,
	"user_activity":         ListAuditPolicyAnalyticsAuditPolicyCategoryUserActivity,
	"oracle_predefined":     ListAuditPolicyAnalyticsAuditPolicyCategoryOraclePredefined,
	"compliance_standard":   ListAuditPolicyAnalyticsAuditPolicyCategoryComplianceStandard,
	"custom":                ListAuditPolicyAnalyticsAuditPolicyCategoryCustom,
	"sql_firewall_auditing": ListAuditPolicyAnalyticsAuditPolicyCategorySqlFirewallAuditing,
}

// GetListAuditPolicyAnalyticsAuditPolicyCategoryEnumValues Enumerates the set of values for ListAuditPolicyAnalyticsAuditPolicyCategoryEnum
func GetListAuditPolicyAnalyticsAuditPolicyCategoryEnumValues() []ListAuditPolicyAnalyticsAuditPolicyCategoryEnum {
	values := make([]ListAuditPolicyAnalyticsAuditPolicyCategoryEnum, 0)
	for _, v := range mappingListAuditPolicyAnalyticsAuditPolicyCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditPolicyAnalyticsAuditPolicyCategoryEnumStringValues Enumerates the set of values in String for ListAuditPolicyAnalyticsAuditPolicyCategoryEnum
func GetListAuditPolicyAnalyticsAuditPolicyCategoryEnumStringValues() []string {
	return []string{
		"BASIC_ACTIVITY",
		"ADMIN_USER_ACTIVITY",
		"USER_ACTIVITY",
		"ORACLE_PREDEFINED",
		"COMPLIANCE_STANDARD",
		"CUSTOM",
		"SQL_FIREWALL_AUDITING",
	}
}

// GetMappingListAuditPolicyAnalyticsAuditPolicyCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditPolicyAnalyticsAuditPolicyCategoryEnum(val string) (ListAuditPolicyAnalyticsAuditPolicyCategoryEnum, bool) {
	enum, ok := mappingListAuditPolicyAnalyticsAuditPolicyCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditPolicyAnalyticsLifecycleStateEnum Enum with underlying type: string
type ListAuditPolicyAnalyticsLifecycleStateEnum string

// Set of constants representing the allowable values for ListAuditPolicyAnalyticsLifecycleStateEnum
const (
	ListAuditPolicyAnalyticsLifecycleStateCreating       ListAuditPolicyAnalyticsLifecycleStateEnum = "CREATING"
	ListAuditPolicyAnalyticsLifecycleStateUpdating       ListAuditPolicyAnalyticsLifecycleStateEnum = "UPDATING"
	ListAuditPolicyAnalyticsLifecycleStateActive         ListAuditPolicyAnalyticsLifecycleStateEnum = "ACTIVE"
	ListAuditPolicyAnalyticsLifecycleStateFailed         ListAuditPolicyAnalyticsLifecycleStateEnum = "FAILED"
	ListAuditPolicyAnalyticsLifecycleStateNeedsAttention ListAuditPolicyAnalyticsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListAuditPolicyAnalyticsLifecycleStateDeleting       ListAuditPolicyAnalyticsLifecycleStateEnum = "DELETING"
	ListAuditPolicyAnalyticsLifecycleStateDeleted        ListAuditPolicyAnalyticsLifecycleStateEnum = "DELETED"
)

var mappingListAuditPolicyAnalyticsLifecycleStateEnum = map[string]ListAuditPolicyAnalyticsLifecycleStateEnum{
	"CREATING":        ListAuditPolicyAnalyticsLifecycleStateCreating,
	"UPDATING":        ListAuditPolicyAnalyticsLifecycleStateUpdating,
	"ACTIVE":          ListAuditPolicyAnalyticsLifecycleStateActive,
	"FAILED":          ListAuditPolicyAnalyticsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListAuditPolicyAnalyticsLifecycleStateNeedsAttention,
	"DELETING":        ListAuditPolicyAnalyticsLifecycleStateDeleting,
	"DELETED":         ListAuditPolicyAnalyticsLifecycleStateDeleted,
}

var mappingListAuditPolicyAnalyticsLifecycleStateEnumLowerCase = map[string]ListAuditPolicyAnalyticsLifecycleStateEnum{
	"creating":        ListAuditPolicyAnalyticsLifecycleStateCreating,
	"updating":        ListAuditPolicyAnalyticsLifecycleStateUpdating,
	"active":          ListAuditPolicyAnalyticsLifecycleStateActive,
	"failed":          ListAuditPolicyAnalyticsLifecycleStateFailed,
	"needs_attention": ListAuditPolicyAnalyticsLifecycleStateNeedsAttention,
	"deleting":        ListAuditPolicyAnalyticsLifecycleStateDeleting,
	"deleted":         ListAuditPolicyAnalyticsLifecycleStateDeleted,
}

// GetListAuditPolicyAnalyticsLifecycleStateEnumValues Enumerates the set of values for ListAuditPolicyAnalyticsLifecycleStateEnum
func GetListAuditPolicyAnalyticsLifecycleStateEnumValues() []ListAuditPolicyAnalyticsLifecycleStateEnum {
	values := make([]ListAuditPolicyAnalyticsLifecycleStateEnum, 0)
	for _, v := range mappingListAuditPolicyAnalyticsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditPolicyAnalyticsLifecycleStateEnumStringValues Enumerates the set of values in String for ListAuditPolicyAnalyticsLifecycleStateEnum
func GetListAuditPolicyAnalyticsLifecycleStateEnumStringValues() []string {
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

// GetMappingListAuditPolicyAnalyticsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditPolicyAnalyticsLifecycleStateEnum(val string) (ListAuditPolicyAnalyticsLifecycleStateEnum, bool) {
	enum, ok := mappingListAuditPolicyAnalyticsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
