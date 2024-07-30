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

// ListAlertAnalyticsRequest wrapper for the ListAlertAnalytics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAlertAnalytics.go.html to see an example of how to use ListAlertAnalyticsRequest.
type ListAlertAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the if-match parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// An optional filter to return audit events whose creation time in the database is greater than and equal to the date-time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStarted"`

	// An optional filter to return audit events whose creation time in the database is less than and equal to the date-time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeEnded *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnded"`

	// Default time zone is UTC if no time zone provided. The date-time considerations of the resource will be in accordance with the specified time zone.
	QueryTimeZone *string `mandatory:"false" contributesTo:"query" name:"queryTimeZone"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListAlertAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListAlertAnalyticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListAlertAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2
	// of the System for Cross-Domain Identity Management (SCIM) specification, which is available
	// at RFC3339 (https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions,
	// text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format.
	// (Numeric and boolean values should not be quoted.)
	// **Example:** |
	// query=(timeCreated ge '2021-06-04T01-00-26') and (targetNames eq 'target_1')
	// query=(featureDetails.userName eq "user") and (targetNames eq "target_1")
	// Supported fields:
	// severity
	// status
	// alertType
	// targetIds
	// targetNames
	// operationTime
	// lifecycleState
	// displayName
	// timeCreated
	// timeUpdated
	// featureDetails.* (* can be any field in nestedStrMap in Feature Attributes in Alert Summary. For example -
	// userName,object,clientHostname,osUserName,clientIPs,clientId,commandText,commandParam,clientProgram,objectType,targetOwner)
	ScimQuery *string `mandatory:"false" contributesTo:"query" name:"scimQuery"`

	// Specifies a subset of summarized fields to be returned in the response.
	SummaryField []ListAlertAnalyticsSummaryFieldEnum `contributesTo:"query" name:"summaryField" omitEmpty:"true" collectionFormat:"multi"`

	// A groupBy can only be used in combination with summaryField parameter.
	// A groupBy value has to be a subset of the values mentioned in summaryField parameter.
	GroupBy []ListAlertAnalyticsGroupByEnum `contributesTo:"query" name:"groupBy" omitEmpty:"true" collectionFormat:"multi"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAlertAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAlertAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAlertAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAlertAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAlertAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAlertAnalyticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAlertAnalyticsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlertAnalyticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAlertAnalyticsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlertAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListAlertAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	for _, val := range request.SummaryField {
		if _, ok := GetMappingListAlertAnalyticsSummaryFieldEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SummaryField: %s. Supported values are: %s.", val, strings.Join(GetListAlertAnalyticsSummaryFieldEnumStringValues(), ",")))
		}
	}

	for _, val := range request.GroupBy {
		if _, ok := GetMappingListAlertAnalyticsGroupByEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", val, strings.Join(GetListAlertAnalyticsGroupByEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAlertAnalyticsResponse wrapper for the ListAlertAnalytics operation
type ListAlertAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AlertAnalyticsCollection instances
	AlertAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAlertAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAlertAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAlertAnalyticsSortOrderEnum Enum with underlying type: string
type ListAlertAnalyticsSortOrderEnum string

// Set of constants representing the allowable values for ListAlertAnalyticsSortOrderEnum
const (
	ListAlertAnalyticsSortOrderAsc  ListAlertAnalyticsSortOrderEnum = "ASC"
	ListAlertAnalyticsSortOrderDesc ListAlertAnalyticsSortOrderEnum = "DESC"
)

var mappingListAlertAnalyticsSortOrderEnum = map[string]ListAlertAnalyticsSortOrderEnum{
	"ASC":  ListAlertAnalyticsSortOrderAsc,
	"DESC": ListAlertAnalyticsSortOrderDesc,
}

var mappingListAlertAnalyticsSortOrderEnumLowerCase = map[string]ListAlertAnalyticsSortOrderEnum{
	"asc":  ListAlertAnalyticsSortOrderAsc,
	"desc": ListAlertAnalyticsSortOrderDesc,
}

// GetListAlertAnalyticsSortOrderEnumValues Enumerates the set of values for ListAlertAnalyticsSortOrderEnum
func GetListAlertAnalyticsSortOrderEnumValues() []ListAlertAnalyticsSortOrderEnum {
	values := make([]ListAlertAnalyticsSortOrderEnum, 0)
	for _, v := range mappingListAlertAnalyticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertAnalyticsSortOrderEnumStringValues Enumerates the set of values in String for ListAlertAnalyticsSortOrderEnum
func GetListAlertAnalyticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAlertAnalyticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertAnalyticsSortOrderEnum(val string) (ListAlertAnalyticsSortOrderEnum, bool) {
	enum, ok := mappingListAlertAnalyticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertAnalyticsSortByEnum Enum with underlying type: string
type ListAlertAnalyticsSortByEnum string

// Set of constants representing the allowable values for ListAlertAnalyticsSortByEnum
const (
	ListAlertAnalyticsSortByDisplayname ListAlertAnalyticsSortByEnum = "displayName"
	ListAlertAnalyticsSortByTimecreated ListAlertAnalyticsSortByEnum = "timeCreated"
)

var mappingListAlertAnalyticsSortByEnum = map[string]ListAlertAnalyticsSortByEnum{
	"displayName": ListAlertAnalyticsSortByDisplayname,
	"timeCreated": ListAlertAnalyticsSortByTimecreated,
}

var mappingListAlertAnalyticsSortByEnumLowerCase = map[string]ListAlertAnalyticsSortByEnum{
	"displayname": ListAlertAnalyticsSortByDisplayname,
	"timecreated": ListAlertAnalyticsSortByTimecreated,
}

// GetListAlertAnalyticsSortByEnumValues Enumerates the set of values for ListAlertAnalyticsSortByEnum
func GetListAlertAnalyticsSortByEnumValues() []ListAlertAnalyticsSortByEnum {
	values := make([]ListAlertAnalyticsSortByEnum, 0)
	for _, v := range mappingListAlertAnalyticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertAnalyticsSortByEnumStringValues Enumerates the set of values in String for ListAlertAnalyticsSortByEnum
func GetListAlertAnalyticsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListAlertAnalyticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertAnalyticsSortByEnum(val string) (ListAlertAnalyticsSortByEnum, bool) {
	enum, ok := mappingListAlertAnalyticsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertAnalyticsAccessLevelEnum Enum with underlying type: string
type ListAlertAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListAlertAnalyticsAccessLevelEnum
const (
	ListAlertAnalyticsAccessLevelRestricted ListAlertAnalyticsAccessLevelEnum = "RESTRICTED"
	ListAlertAnalyticsAccessLevelAccessible ListAlertAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListAlertAnalyticsAccessLevelEnum = map[string]ListAlertAnalyticsAccessLevelEnum{
	"RESTRICTED": ListAlertAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListAlertAnalyticsAccessLevelAccessible,
}

var mappingListAlertAnalyticsAccessLevelEnumLowerCase = map[string]ListAlertAnalyticsAccessLevelEnum{
	"restricted": ListAlertAnalyticsAccessLevelRestricted,
	"accessible": ListAlertAnalyticsAccessLevelAccessible,
}

// GetListAlertAnalyticsAccessLevelEnumValues Enumerates the set of values for ListAlertAnalyticsAccessLevelEnum
func GetListAlertAnalyticsAccessLevelEnumValues() []ListAlertAnalyticsAccessLevelEnum {
	values := make([]ListAlertAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListAlertAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListAlertAnalyticsAccessLevelEnum
func GetListAlertAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListAlertAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertAnalyticsAccessLevelEnum(val string) (ListAlertAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListAlertAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertAnalyticsSummaryFieldEnum Enum with underlying type: string
type ListAlertAnalyticsSummaryFieldEnum string

// Set of constants representing the allowable values for ListAlertAnalyticsSummaryFieldEnum
const (
	ListAlertAnalyticsSummaryFieldAlerttype           ListAlertAnalyticsSummaryFieldEnum = "alertType"
	ListAlertAnalyticsSummaryFieldTargetids           ListAlertAnalyticsSummaryFieldEnum = "targetIds"
	ListAlertAnalyticsSummaryFieldTargetnames         ListAlertAnalyticsSummaryFieldEnum = "targetNames"
	ListAlertAnalyticsSummaryFieldAlertseverity       ListAlertAnalyticsSummaryFieldEnum = "alertSeverity"
	ListAlertAnalyticsSummaryFieldAlertstatus         ListAlertAnalyticsSummaryFieldEnum = "alertStatus"
	ListAlertAnalyticsSummaryFieldTimecreated         ListAlertAnalyticsSummaryFieldEnum = "timeCreated"
	ListAlertAnalyticsSummaryFieldPolicyid            ListAlertAnalyticsSummaryFieldEnum = "policyId"
	ListAlertAnalyticsSummaryFieldOpen                ListAlertAnalyticsSummaryFieldEnum = "open"
	ListAlertAnalyticsSummaryFieldClosed              ListAlertAnalyticsSummaryFieldEnum = "closed"
	ListAlertAnalyticsSummaryFieldCritical            ListAlertAnalyticsSummaryFieldEnum = "critical"
	ListAlertAnalyticsSummaryFieldHigh                ListAlertAnalyticsSummaryFieldEnum = "high"
	ListAlertAnalyticsSummaryFieldMedium              ListAlertAnalyticsSummaryFieldEnum = "medium"
	ListAlertAnalyticsSummaryFieldLow                 ListAlertAnalyticsSummaryFieldEnum = "low"
	ListAlertAnalyticsSummaryFieldAlertcount          ListAlertAnalyticsSummaryFieldEnum = "alertcount"
	ListAlertAnalyticsSummaryFieldAlertpolicyrulekey  ListAlertAnalyticsSummaryFieldEnum = "alertPolicyRuleKey"
	ListAlertAnalyticsSummaryFieldAlertpolicyrulename ListAlertAnalyticsSummaryFieldEnum = "alertPolicyRuleName"
	ListAlertAnalyticsSummaryFieldThrottled           ListAlertAnalyticsSummaryFieldEnum = "throttled"
)

var mappingListAlertAnalyticsSummaryFieldEnum = map[string]ListAlertAnalyticsSummaryFieldEnum{
	"alertType":           ListAlertAnalyticsSummaryFieldAlerttype,
	"targetIds":           ListAlertAnalyticsSummaryFieldTargetids,
	"targetNames":         ListAlertAnalyticsSummaryFieldTargetnames,
	"alertSeverity":       ListAlertAnalyticsSummaryFieldAlertseverity,
	"alertStatus":         ListAlertAnalyticsSummaryFieldAlertstatus,
	"timeCreated":         ListAlertAnalyticsSummaryFieldTimecreated,
	"policyId":            ListAlertAnalyticsSummaryFieldPolicyid,
	"open":                ListAlertAnalyticsSummaryFieldOpen,
	"closed":              ListAlertAnalyticsSummaryFieldClosed,
	"critical":            ListAlertAnalyticsSummaryFieldCritical,
	"high":                ListAlertAnalyticsSummaryFieldHigh,
	"medium":              ListAlertAnalyticsSummaryFieldMedium,
	"low":                 ListAlertAnalyticsSummaryFieldLow,
	"alertcount":          ListAlertAnalyticsSummaryFieldAlertcount,
	"alertPolicyRuleKey":  ListAlertAnalyticsSummaryFieldAlertpolicyrulekey,
	"alertPolicyRuleName": ListAlertAnalyticsSummaryFieldAlertpolicyrulename,
	"throttled":           ListAlertAnalyticsSummaryFieldThrottled,
}

var mappingListAlertAnalyticsSummaryFieldEnumLowerCase = map[string]ListAlertAnalyticsSummaryFieldEnum{
	"alerttype":           ListAlertAnalyticsSummaryFieldAlerttype,
	"targetids":           ListAlertAnalyticsSummaryFieldTargetids,
	"targetnames":         ListAlertAnalyticsSummaryFieldTargetnames,
	"alertseverity":       ListAlertAnalyticsSummaryFieldAlertseverity,
	"alertstatus":         ListAlertAnalyticsSummaryFieldAlertstatus,
	"timecreated":         ListAlertAnalyticsSummaryFieldTimecreated,
	"policyid":            ListAlertAnalyticsSummaryFieldPolicyid,
	"open":                ListAlertAnalyticsSummaryFieldOpen,
	"closed":              ListAlertAnalyticsSummaryFieldClosed,
	"critical":            ListAlertAnalyticsSummaryFieldCritical,
	"high":                ListAlertAnalyticsSummaryFieldHigh,
	"medium":              ListAlertAnalyticsSummaryFieldMedium,
	"low":                 ListAlertAnalyticsSummaryFieldLow,
	"alertcount":          ListAlertAnalyticsSummaryFieldAlertcount,
	"alertpolicyrulekey":  ListAlertAnalyticsSummaryFieldAlertpolicyrulekey,
	"alertpolicyrulename": ListAlertAnalyticsSummaryFieldAlertpolicyrulename,
	"throttled":           ListAlertAnalyticsSummaryFieldThrottled,
}

// GetListAlertAnalyticsSummaryFieldEnumValues Enumerates the set of values for ListAlertAnalyticsSummaryFieldEnum
func GetListAlertAnalyticsSummaryFieldEnumValues() []ListAlertAnalyticsSummaryFieldEnum {
	values := make([]ListAlertAnalyticsSummaryFieldEnum, 0)
	for _, v := range mappingListAlertAnalyticsSummaryFieldEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertAnalyticsSummaryFieldEnumStringValues Enumerates the set of values in String for ListAlertAnalyticsSummaryFieldEnum
func GetListAlertAnalyticsSummaryFieldEnumStringValues() []string {
	return []string{
		"alertType",
		"targetIds",
		"targetNames",
		"alertSeverity",
		"alertStatus",
		"timeCreated",
		"policyId",
		"open",
		"closed",
		"critical",
		"high",
		"medium",
		"low",
		"alertcount",
		"alertPolicyRuleKey",
		"alertPolicyRuleName",
		"throttled",
	}
}

// GetMappingListAlertAnalyticsSummaryFieldEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertAnalyticsSummaryFieldEnum(val string) (ListAlertAnalyticsSummaryFieldEnum, bool) {
	enum, ok := mappingListAlertAnalyticsSummaryFieldEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertAnalyticsGroupByEnum Enum with underlying type: string
type ListAlertAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListAlertAnalyticsGroupByEnum
const (
	ListAlertAnalyticsGroupByAlerttype           ListAlertAnalyticsGroupByEnum = "alertType"
	ListAlertAnalyticsGroupByTargetids           ListAlertAnalyticsGroupByEnum = "targetIds"
	ListAlertAnalyticsGroupByTargetnames         ListAlertAnalyticsGroupByEnum = "targetNames"
	ListAlertAnalyticsGroupByAlertseverity       ListAlertAnalyticsGroupByEnum = "alertSeverity"
	ListAlertAnalyticsGroupByAlertstatus         ListAlertAnalyticsGroupByEnum = "alertStatus"
	ListAlertAnalyticsGroupByTimecreated         ListAlertAnalyticsGroupByEnum = "timeCreated"
	ListAlertAnalyticsGroupByPolicyid            ListAlertAnalyticsGroupByEnum = "policyId"
	ListAlertAnalyticsGroupByAlertpolicyrulekey  ListAlertAnalyticsGroupByEnum = "alertPolicyRuleKey"
	ListAlertAnalyticsGroupByAlertpolicyrulename ListAlertAnalyticsGroupByEnum = "alertPolicyRuleName"
)

var mappingListAlertAnalyticsGroupByEnum = map[string]ListAlertAnalyticsGroupByEnum{
	"alertType":           ListAlertAnalyticsGroupByAlerttype,
	"targetIds":           ListAlertAnalyticsGroupByTargetids,
	"targetNames":         ListAlertAnalyticsGroupByTargetnames,
	"alertSeverity":       ListAlertAnalyticsGroupByAlertseverity,
	"alertStatus":         ListAlertAnalyticsGroupByAlertstatus,
	"timeCreated":         ListAlertAnalyticsGroupByTimecreated,
	"policyId":            ListAlertAnalyticsGroupByPolicyid,
	"alertPolicyRuleKey":  ListAlertAnalyticsGroupByAlertpolicyrulekey,
	"alertPolicyRuleName": ListAlertAnalyticsGroupByAlertpolicyrulename,
}

var mappingListAlertAnalyticsGroupByEnumLowerCase = map[string]ListAlertAnalyticsGroupByEnum{
	"alerttype":           ListAlertAnalyticsGroupByAlerttype,
	"targetids":           ListAlertAnalyticsGroupByTargetids,
	"targetnames":         ListAlertAnalyticsGroupByTargetnames,
	"alertseverity":       ListAlertAnalyticsGroupByAlertseverity,
	"alertstatus":         ListAlertAnalyticsGroupByAlertstatus,
	"timecreated":         ListAlertAnalyticsGroupByTimecreated,
	"policyid":            ListAlertAnalyticsGroupByPolicyid,
	"alertpolicyrulekey":  ListAlertAnalyticsGroupByAlertpolicyrulekey,
	"alertpolicyrulename": ListAlertAnalyticsGroupByAlertpolicyrulename,
}

// GetListAlertAnalyticsGroupByEnumValues Enumerates the set of values for ListAlertAnalyticsGroupByEnum
func GetListAlertAnalyticsGroupByEnumValues() []ListAlertAnalyticsGroupByEnum {
	values := make([]ListAlertAnalyticsGroupByEnum, 0)
	for _, v := range mappingListAlertAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListAlertAnalyticsGroupByEnum
func GetListAlertAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"alertType",
		"targetIds",
		"targetNames",
		"alertSeverity",
		"alertStatus",
		"timeCreated",
		"policyId",
		"alertPolicyRuleKey",
		"alertPolicyRuleName",
	}
}

// GetMappingListAlertAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertAnalyticsGroupByEnum(val string) (ListAlertAnalyticsGroupByEnum, bool) {
	enum, ok := mappingListAlertAnalyticsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
