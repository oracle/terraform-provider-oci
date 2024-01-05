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

// ListAuditEventAnalyticsRequest wrapper for the ListAuditEventAnalytics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditEventAnalytics.go.html to see an example of how to use ListAuditEventAnalyticsRequest.
type ListAuditEventAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the if-match parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// It is usually retrieved from a previous "List" call. For details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListAuditEventAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2
	// of the System for Cross-Domain Identity Management (SCIM) specification, which is available
	// at RFC3339 (https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions,
	// text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format.
	// (Numeric and boolean values should not be quoted.)
	// **Example:** query=(operationTime ge '2021-06-04T01-00-26') and (eventName eq 'LOGON')
	ScimQuery *string `mandatory:"false" contributesTo:"query" name:"scimQuery"`

	// Specifies a subset of summarized fields to be returned in the response.
	SummaryField []ListAuditEventAnalyticsSummaryFieldEnum `contributesTo:"query" name:"summaryField" omitEmpty:"true" collectionFormat:"multi"`

	// An optional filter to return audit events whose creation time in the database is greater than and equal to the date-time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStarted"`

	// An optional filter to return audit events whose creation time in the database is less than and equal to the date-time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeEnded *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnded"`

	// Default time zone is UTC if no time zone provided. The date-time considerations of the resource will be in accordance with the specified time zone.
	QueryTimeZone *string `mandatory:"false" contributesTo:"query" name:"queryTimeZone"`

	// A groupBy can only be used in combination with summaryField parameter.
	// A groupBy value has to be a subset of the values mentioned in summaryField parameter.
	GroupBy []ListAuditEventAnalyticsGroupByEnum `contributesTo:"query" name:"groupBy" omitEmpty:"true" collectionFormat:"multi"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListAuditEventAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// If this query parameter is specified, the result is ordered based on this query parameter value.
	SortBy ListAuditEventAnalyticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAuditEventAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAuditEventAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAuditEventAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAuditEventAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAuditEventAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAuditEventAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListAuditEventAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	for _, val := range request.SummaryField {
		if _, ok := GetMappingListAuditEventAnalyticsSummaryFieldEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SummaryField: %s. Supported values are: %s.", val, strings.Join(GetListAuditEventAnalyticsSummaryFieldEnumStringValues(), ",")))
		}
	}

	for _, val := range request.GroupBy {
		if _, ok := GetMappingListAuditEventAnalyticsGroupByEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", val, strings.Join(GetListAuditEventAnalyticsGroupByEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListAuditEventAnalyticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAuditEventAnalyticsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditEventAnalyticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAuditEventAnalyticsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAuditEventAnalyticsResponse wrapper for the ListAuditEventAnalytics operation
type ListAuditEventAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AuditEventAnalyticsCollection instances
	AuditEventAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListAuditEventAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAuditEventAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAuditEventAnalyticsAccessLevelEnum Enum with underlying type: string
type ListAuditEventAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListAuditEventAnalyticsAccessLevelEnum
const (
	ListAuditEventAnalyticsAccessLevelRestricted ListAuditEventAnalyticsAccessLevelEnum = "RESTRICTED"
	ListAuditEventAnalyticsAccessLevelAccessible ListAuditEventAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListAuditEventAnalyticsAccessLevelEnum = map[string]ListAuditEventAnalyticsAccessLevelEnum{
	"RESTRICTED": ListAuditEventAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListAuditEventAnalyticsAccessLevelAccessible,
}

var mappingListAuditEventAnalyticsAccessLevelEnumLowerCase = map[string]ListAuditEventAnalyticsAccessLevelEnum{
	"restricted": ListAuditEventAnalyticsAccessLevelRestricted,
	"accessible": ListAuditEventAnalyticsAccessLevelAccessible,
}

// GetListAuditEventAnalyticsAccessLevelEnumValues Enumerates the set of values for ListAuditEventAnalyticsAccessLevelEnum
func GetListAuditEventAnalyticsAccessLevelEnumValues() []ListAuditEventAnalyticsAccessLevelEnum {
	values := make([]ListAuditEventAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListAuditEventAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditEventAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListAuditEventAnalyticsAccessLevelEnum
func GetListAuditEventAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListAuditEventAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditEventAnalyticsAccessLevelEnum(val string) (ListAuditEventAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListAuditEventAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditEventAnalyticsSummaryFieldEnum Enum with underlying type: string
type ListAuditEventAnalyticsSummaryFieldEnum string

// Set of constants representing the allowable values for ListAuditEventAnalyticsSummaryFieldEnum
const (
	ListAuditEventAnalyticsSummaryFieldAuditeventtime        ListAuditEventAnalyticsSummaryFieldEnum = "auditEventTime"
	ListAuditEventAnalyticsSummaryFieldDbusername            ListAuditEventAnalyticsSummaryFieldEnum = "dbUserName"
	ListAuditEventAnalyticsSummaryFieldTargetid              ListAuditEventAnalyticsSummaryFieldEnum = "targetId"
	ListAuditEventAnalyticsSummaryFieldTargetname            ListAuditEventAnalyticsSummaryFieldEnum = "targetName"
	ListAuditEventAnalyticsSummaryFieldTargetclass           ListAuditEventAnalyticsSummaryFieldEnum = "targetClass"
	ListAuditEventAnalyticsSummaryFieldObjecttype            ListAuditEventAnalyticsSummaryFieldEnum = "objectType"
	ListAuditEventAnalyticsSummaryFieldClienthostname        ListAuditEventAnalyticsSummaryFieldEnum = "clientHostname"
	ListAuditEventAnalyticsSummaryFieldClientprogram         ListAuditEventAnalyticsSummaryFieldEnum = "clientProgram"
	ListAuditEventAnalyticsSummaryFieldClientid              ListAuditEventAnalyticsSummaryFieldEnum = "clientId"
	ListAuditEventAnalyticsSummaryFieldAudittype             ListAuditEventAnalyticsSummaryFieldEnum = "auditType"
	ListAuditEventAnalyticsSummaryFieldEventname             ListAuditEventAnalyticsSummaryFieldEnum = "eventName"
	ListAuditEventAnalyticsSummaryFieldAllrecord             ListAuditEventAnalyticsSummaryFieldEnum = "allRecord"
	ListAuditEventAnalyticsSummaryFieldAuditsettingschange   ListAuditEventAnalyticsSummaryFieldEnum = "auditSettingsChange"
	ListAuditEventAnalyticsSummaryFieldDbschemachange        ListAuditEventAnalyticsSummaryFieldEnum = "dbSchemaChange"
	ListAuditEventAnalyticsSummaryFieldEntitlementchange     ListAuditEventAnalyticsSummaryFieldEnum = "entitlementChange"
	ListAuditEventAnalyticsSummaryFieldLoginfailure          ListAuditEventAnalyticsSummaryFieldEnum = "loginFailure"
	ListAuditEventAnalyticsSummaryFieldLoginsuccess          ListAuditEventAnalyticsSummaryFieldEnum = "loginSuccess"
	ListAuditEventAnalyticsSummaryFieldAllviolations         ListAuditEventAnalyticsSummaryFieldEnum = "allViolations"
	ListAuditEventAnalyticsSummaryFieldRealmviolations       ListAuditEventAnalyticsSummaryFieldEnum = "realmViolations"
	ListAuditEventAnalyticsSummaryFieldRuleviolations        ListAuditEventAnalyticsSummaryFieldEnum = "ruleViolations"
	ListAuditEventAnalyticsSummaryFieldDvconfigactivities    ListAuditEventAnalyticsSummaryFieldEnum = "dvconfigActivities"
	ListAuditEventAnalyticsSummaryFieldDdls                  ListAuditEventAnalyticsSummaryFieldEnum = "ddls"
	ListAuditEventAnalyticsSummaryFieldDmls                  ListAuditEventAnalyticsSummaryFieldEnum = "dmls"
	ListAuditEventAnalyticsSummaryFieldPrivilegechanges      ListAuditEventAnalyticsSummaryFieldEnum = "privilegeChanges"
	ListAuditEventAnalyticsSummaryFieldAuditsettingsenables  ListAuditEventAnalyticsSummaryFieldEnum = "auditSettingsEnables"
	ListAuditEventAnalyticsSummaryFieldAuditsettingsdisables ListAuditEventAnalyticsSummaryFieldEnum = "auditSettingsDisables"
	ListAuditEventAnalyticsSummaryFieldSelects               ListAuditEventAnalyticsSummaryFieldEnum = "selects"
	ListAuditEventAnalyticsSummaryFieldCreates               ListAuditEventAnalyticsSummaryFieldEnum = "creates"
	ListAuditEventAnalyticsSummaryFieldAlters                ListAuditEventAnalyticsSummaryFieldEnum = "alters"
	ListAuditEventAnalyticsSummaryFieldDrops                 ListAuditEventAnalyticsSummaryFieldEnum = "drops"
	ListAuditEventAnalyticsSummaryFieldGrants                ListAuditEventAnalyticsSummaryFieldEnum = "grants"
	ListAuditEventAnalyticsSummaryFieldRevokes               ListAuditEventAnalyticsSummaryFieldEnum = "revokes"
)

var mappingListAuditEventAnalyticsSummaryFieldEnum = map[string]ListAuditEventAnalyticsSummaryFieldEnum{
	"auditEventTime":        ListAuditEventAnalyticsSummaryFieldAuditeventtime,
	"dbUserName":            ListAuditEventAnalyticsSummaryFieldDbusername,
	"targetId":              ListAuditEventAnalyticsSummaryFieldTargetid,
	"targetName":            ListAuditEventAnalyticsSummaryFieldTargetname,
	"targetClass":           ListAuditEventAnalyticsSummaryFieldTargetclass,
	"objectType":            ListAuditEventAnalyticsSummaryFieldObjecttype,
	"clientHostname":        ListAuditEventAnalyticsSummaryFieldClienthostname,
	"clientProgram":         ListAuditEventAnalyticsSummaryFieldClientprogram,
	"clientId":              ListAuditEventAnalyticsSummaryFieldClientid,
	"auditType":             ListAuditEventAnalyticsSummaryFieldAudittype,
	"eventName":             ListAuditEventAnalyticsSummaryFieldEventname,
	"allRecord":             ListAuditEventAnalyticsSummaryFieldAllrecord,
	"auditSettingsChange":   ListAuditEventAnalyticsSummaryFieldAuditsettingschange,
	"dbSchemaChange":        ListAuditEventAnalyticsSummaryFieldDbschemachange,
	"entitlementChange":     ListAuditEventAnalyticsSummaryFieldEntitlementchange,
	"loginFailure":          ListAuditEventAnalyticsSummaryFieldLoginfailure,
	"loginSuccess":          ListAuditEventAnalyticsSummaryFieldLoginsuccess,
	"allViolations":         ListAuditEventAnalyticsSummaryFieldAllviolations,
	"realmViolations":       ListAuditEventAnalyticsSummaryFieldRealmviolations,
	"ruleViolations":        ListAuditEventAnalyticsSummaryFieldRuleviolations,
	"dvconfigActivities":    ListAuditEventAnalyticsSummaryFieldDvconfigactivities,
	"ddls":                  ListAuditEventAnalyticsSummaryFieldDdls,
	"dmls":                  ListAuditEventAnalyticsSummaryFieldDmls,
	"privilegeChanges":      ListAuditEventAnalyticsSummaryFieldPrivilegechanges,
	"auditSettingsEnables":  ListAuditEventAnalyticsSummaryFieldAuditsettingsenables,
	"auditSettingsDisables": ListAuditEventAnalyticsSummaryFieldAuditsettingsdisables,
	"selects":               ListAuditEventAnalyticsSummaryFieldSelects,
	"creates":               ListAuditEventAnalyticsSummaryFieldCreates,
	"alters":                ListAuditEventAnalyticsSummaryFieldAlters,
	"drops":                 ListAuditEventAnalyticsSummaryFieldDrops,
	"grants":                ListAuditEventAnalyticsSummaryFieldGrants,
	"revokes":               ListAuditEventAnalyticsSummaryFieldRevokes,
}

var mappingListAuditEventAnalyticsSummaryFieldEnumLowerCase = map[string]ListAuditEventAnalyticsSummaryFieldEnum{
	"auditeventtime":        ListAuditEventAnalyticsSummaryFieldAuditeventtime,
	"dbusername":            ListAuditEventAnalyticsSummaryFieldDbusername,
	"targetid":              ListAuditEventAnalyticsSummaryFieldTargetid,
	"targetname":            ListAuditEventAnalyticsSummaryFieldTargetname,
	"targetclass":           ListAuditEventAnalyticsSummaryFieldTargetclass,
	"objecttype":            ListAuditEventAnalyticsSummaryFieldObjecttype,
	"clienthostname":        ListAuditEventAnalyticsSummaryFieldClienthostname,
	"clientprogram":         ListAuditEventAnalyticsSummaryFieldClientprogram,
	"clientid":              ListAuditEventAnalyticsSummaryFieldClientid,
	"audittype":             ListAuditEventAnalyticsSummaryFieldAudittype,
	"eventname":             ListAuditEventAnalyticsSummaryFieldEventname,
	"allrecord":             ListAuditEventAnalyticsSummaryFieldAllrecord,
	"auditsettingschange":   ListAuditEventAnalyticsSummaryFieldAuditsettingschange,
	"dbschemachange":        ListAuditEventAnalyticsSummaryFieldDbschemachange,
	"entitlementchange":     ListAuditEventAnalyticsSummaryFieldEntitlementchange,
	"loginfailure":          ListAuditEventAnalyticsSummaryFieldLoginfailure,
	"loginsuccess":          ListAuditEventAnalyticsSummaryFieldLoginsuccess,
	"allviolations":         ListAuditEventAnalyticsSummaryFieldAllviolations,
	"realmviolations":       ListAuditEventAnalyticsSummaryFieldRealmviolations,
	"ruleviolations":        ListAuditEventAnalyticsSummaryFieldRuleviolations,
	"dvconfigactivities":    ListAuditEventAnalyticsSummaryFieldDvconfigactivities,
	"ddls":                  ListAuditEventAnalyticsSummaryFieldDdls,
	"dmls":                  ListAuditEventAnalyticsSummaryFieldDmls,
	"privilegechanges":      ListAuditEventAnalyticsSummaryFieldPrivilegechanges,
	"auditsettingsenables":  ListAuditEventAnalyticsSummaryFieldAuditsettingsenables,
	"auditsettingsdisables": ListAuditEventAnalyticsSummaryFieldAuditsettingsdisables,
	"selects":               ListAuditEventAnalyticsSummaryFieldSelects,
	"creates":               ListAuditEventAnalyticsSummaryFieldCreates,
	"alters":                ListAuditEventAnalyticsSummaryFieldAlters,
	"drops":                 ListAuditEventAnalyticsSummaryFieldDrops,
	"grants":                ListAuditEventAnalyticsSummaryFieldGrants,
	"revokes":               ListAuditEventAnalyticsSummaryFieldRevokes,
}

// GetListAuditEventAnalyticsSummaryFieldEnumValues Enumerates the set of values for ListAuditEventAnalyticsSummaryFieldEnum
func GetListAuditEventAnalyticsSummaryFieldEnumValues() []ListAuditEventAnalyticsSummaryFieldEnum {
	values := make([]ListAuditEventAnalyticsSummaryFieldEnum, 0)
	for _, v := range mappingListAuditEventAnalyticsSummaryFieldEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditEventAnalyticsSummaryFieldEnumStringValues Enumerates the set of values in String for ListAuditEventAnalyticsSummaryFieldEnum
func GetListAuditEventAnalyticsSummaryFieldEnumStringValues() []string {
	return []string{
		"auditEventTime",
		"dbUserName",
		"targetId",
		"targetName",
		"targetClass",
		"objectType",
		"clientHostname",
		"clientProgram",
		"clientId",
		"auditType",
		"eventName",
		"allRecord",
		"auditSettingsChange",
		"dbSchemaChange",
		"entitlementChange",
		"loginFailure",
		"loginSuccess",
		"allViolations",
		"realmViolations",
		"ruleViolations",
		"dvconfigActivities",
		"ddls",
		"dmls",
		"privilegeChanges",
		"auditSettingsEnables",
		"auditSettingsDisables",
		"selects",
		"creates",
		"alters",
		"drops",
		"grants",
		"revokes",
	}
}

// GetMappingListAuditEventAnalyticsSummaryFieldEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditEventAnalyticsSummaryFieldEnum(val string) (ListAuditEventAnalyticsSummaryFieldEnum, bool) {
	enum, ok := mappingListAuditEventAnalyticsSummaryFieldEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditEventAnalyticsGroupByEnum Enum with underlying type: string
type ListAuditEventAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListAuditEventAnalyticsGroupByEnum
const (
	ListAuditEventAnalyticsGroupByAuditeventtime ListAuditEventAnalyticsGroupByEnum = "auditEventTime"
	ListAuditEventAnalyticsGroupByDbusername     ListAuditEventAnalyticsGroupByEnum = "dbUserName"
	ListAuditEventAnalyticsGroupByTargetid       ListAuditEventAnalyticsGroupByEnum = "targetId"
	ListAuditEventAnalyticsGroupByTargetname     ListAuditEventAnalyticsGroupByEnum = "targetName"
	ListAuditEventAnalyticsGroupByTargetclass    ListAuditEventAnalyticsGroupByEnum = "targetClass"
	ListAuditEventAnalyticsGroupByObjecttype     ListAuditEventAnalyticsGroupByEnum = "objectType"
	ListAuditEventAnalyticsGroupByClienthostname ListAuditEventAnalyticsGroupByEnum = "clientHostname"
	ListAuditEventAnalyticsGroupByClientprogram  ListAuditEventAnalyticsGroupByEnum = "clientProgram"
	ListAuditEventAnalyticsGroupByClientid       ListAuditEventAnalyticsGroupByEnum = "clientId"
	ListAuditEventAnalyticsGroupByAudittype      ListAuditEventAnalyticsGroupByEnum = "auditType"
	ListAuditEventAnalyticsGroupByEventname      ListAuditEventAnalyticsGroupByEnum = "eventName"
)

var mappingListAuditEventAnalyticsGroupByEnum = map[string]ListAuditEventAnalyticsGroupByEnum{
	"auditEventTime": ListAuditEventAnalyticsGroupByAuditeventtime,
	"dbUserName":     ListAuditEventAnalyticsGroupByDbusername,
	"targetId":       ListAuditEventAnalyticsGroupByTargetid,
	"targetName":     ListAuditEventAnalyticsGroupByTargetname,
	"targetClass":    ListAuditEventAnalyticsGroupByTargetclass,
	"objectType":     ListAuditEventAnalyticsGroupByObjecttype,
	"clientHostname": ListAuditEventAnalyticsGroupByClienthostname,
	"clientProgram":  ListAuditEventAnalyticsGroupByClientprogram,
	"clientId":       ListAuditEventAnalyticsGroupByClientid,
	"auditType":      ListAuditEventAnalyticsGroupByAudittype,
	"eventName":      ListAuditEventAnalyticsGroupByEventname,
}

var mappingListAuditEventAnalyticsGroupByEnumLowerCase = map[string]ListAuditEventAnalyticsGroupByEnum{
	"auditeventtime": ListAuditEventAnalyticsGroupByAuditeventtime,
	"dbusername":     ListAuditEventAnalyticsGroupByDbusername,
	"targetid":       ListAuditEventAnalyticsGroupByTargetid,
	"targetname":     ListAuditEventAnalyticsGroupByTargetname,
	"targetclass":    ListAuditEventAnalyticsGroupByTargetclass,
	"objecttype":     ListAuditEventAnalyticsGroupByObjecttype,
	"clienthostname": ListAuditEventAnalyticsGroupByClienthostname,
	"clientprogram":  ListAuditEventAnalyticsGroupByClientprogram,
	"clientid":       ListAuditEventAnalyticsGroupByClientid,
	"audittype":      ListAuditEventAnalyticsGroupByAudittype,
	"eventname":      ListAuditEventAnalyticsGroupByEventname,
}

// GetListAuditEventAnalyticsGroupByEnumValues Enumerates the set of values for ListAuditEventAnalyticsGroupByEnum
func GetListAuditEventAnalyticsGroupByEnumValues() []ListAuditEventAnalyticsGroupByEnum {
	values := make([]ListAuditEventAnalyticsGroupByEnum, 0)
	for _, v := range mappingListAuditEventAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditEventAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListAuditEventAnalyticsGroupByEnum
func GetListAuditEventAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"auditEventTime",
		"dbUserName",
		"targetId",
		"targetName",
		"targetClass",
		"objectType",
		"clientHostname",
		"clientProgram",
		"clientId",
		"auditType",
		"eventName",
	}
}

// GetMappingListAuditEventAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditEventAnalyticsGroupByEnum(val string) (ListAuditEventAnalyticsGroupByEnum, bool) {
	enum, ok := mappingListAuditEventAnalyticsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditEventAnalyticsSortOrderEnum Enum with underlying type: string
type ListAuditEventAnalyticsSortOrderEnum string

// Set of constants representing the allowable values for ListAuditEventAnalyticsSortOrderEnum
const (
	ListAuditEventAnalyticsSortOrderAsc  ListAuditEventAnalyticsSortOrderEnum = "ASC"
	ListAuditEventAnalyticsSortOrderDesc ListAuditEventAnalyticsSortOrderEnum = "DESC"
)

var mappingListAuditEventAnalyticsSortOrderEnum = map[string]ListAuditEventAnalyticsSortOrderEnum{
	"ASC":  ListAuditEventAnalyticsSortOrderAsc,
	"DESC": ListAuditEventAnalyticsSortOrderDesc,
}

var mappingListAuditEventAnalyticsSortOrderEnumLowerCase = map[string]ListAuditEventAnalyticsSortOrderEnum{
	"asc":  ListAuditEventAnalyticsSortOrderAsc,
	"desc": ListAuditEventAnalyticsSortOrderDesc,
}

// GetListAuditEventAnalyticsSortOrderEnumValues Enumerates the set of values for ListAuditEventAnalyticsSortOrderEnum
func GetListAuditEventAnalyticsSortOrderEnumValues() []ListAuditEventAnalyticsSortOrderEnum {
	values := make([]ListAuditEventAnalyticsSortOrderEnum, 0)
	for _, v := range mappingListAuditEventAnalyticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditEventAnalyticsSortOrderEnumStringValues Enumerates the set of values in String for ListAuditEventAnalyticsSortOrderEnum
func GetListAuditEventAnalyticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAuditEventAnalyticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditEventAnalyticsSortOrderEnum(val string) (ListAuditEventAnalyticsSortOrderEnum, bool) {
	enum, ok := mappingListAuditEventAnalyticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditEventAnalyticsSortByEnum Enum with underlying type: string
type ListAuditEventAnalyticsSortByEnum string

// Set of constants representing the allowable values for ListAuditEventAnalyticsSortByEnum
const (
	ListAuditEventAnalyticsSortByTargetid       ListAuditEventAnalyticsSortByEnum = "targetId"
	ListAuditEventAnalyticsSortByTargetclass    ListAuditEventAnalyticsSortByEnum = "targetClass"
	ListAuditEventAnalyticsSortByTargetname     ListAuditEventAnalyticsSortByEnum = "targetName"
	ListAuditEventAnalyticsSortByObjecttype     ListAuditEventAnalyticsSortByEnum = "objectType"
	ListAuditEventAnalyticsSortByDbusername     ListAuditEventAnalyticsSortByEnum = "dbUserName"
	ListAuditEventAnalyticsSortByEventname      ListAuditEventAnalyticsSortByEnum = "eventName"
	ListAuditEventAnalyticsSortByAuditeventtime ListAuditEventAnalyticsSortByEnum = "auditEventTime"
	ListAuditEventAnalyticsSortByClienthostname ListAuditEventAnalyticsSortByEnum = "clientHostname"
	ListAuditEventAnalyticsSortByClientprogram  ListAuditEventAnalyticsSortByEnum = "clientProgram"
	ListAuditEventAnalyticsSortByClientid       ListAuditEventAnalyticsSortByEnum = "clientId"
	ListAuditEventAnalyticsSortByAudittype      ListAuditEventAnalyticsSortByEnum = "auditType"
)

var mappingListAuditEventAnalyticsSortByEnum = map[string]ListAuditEventAnalyticsSortByEnum{
	"targetId":       ListAuditEventAnalyticsSortByTargetid,
	"targetClass":    ListAuditEventAnalyticsSortByTargetclass,
	"targetName":     ListAuditEventAnalyticsSortByTargetname,
	"objectType":     ListAuditEventAnalyticsSortByObjecttype,
	"dbUserName":     ListAuditEventAnalyticsSortByDbusername,
	"eventName":      ListAuditEventAnalyticsSortByEventname,
	"auditEventTime": ListAuditEventAnalyticsSortByAuditeventtime,
	"clientHostname": ListAuditEventAnalyticsSortByClienthostname,
	"clientProgram":  ListAuditEventAnalyticsSortByClientprogram,
	"clientId":       ListAuditEventAnalyticsSortByClientid,
	"auditType":      ListAuditEventAnalyticsSortByAudittype,
}

var mappingListAuditEventAnalyticsSortByEnumLowerCase = map[string]ListAuditEventAnalyticsSortByEnum{
	"targetid":       ListAuditEventAnalyticsSortByTargetid,
	"targetclass":    ListAuditEventAnalyticsSortByTargetclass,
	"targetname":     ListAuditEventAnalyticsSortByTargetname,
	"objecttype":     ListAuditEventAnalyticsSortByObjecttype,
	"dbusername":     ListAuditEventAnalyticsSortByDbusername,
	"eventname":      ListAuditEventAnalyticsSortByEventname,
	"auditeventtime": ListAuditEventAnalyticsSortByAuditeventtime,
	"clienthostname": ListAuditEventAnalyticsSortByClienthostname,
	"clientprogram":  ListAuditEventAnalyticsSortByClientprogram,
	"clientid":       ListAuditEventAnalyticsSortByClientid,
	"audittype":      ListAuditEventAnalyticsSortByAudittype,
}

// GetListAuditEventAnalyticsSortByEnumValues Enumerates the set of values for ListAuditEventAnalyticsSortByEnum
func GetListAuditEventAnalyticsSortByEnumValues() []ListAuditEventAnalyticsSortByEnum {
	values := make([]ListAuditEventAnalyticsSortByEnum, 0)
	for _, v := range mappingListAuditEventAnalyticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditEventAnalyticsSortByEnumStringValues Enumerates the set of values in String for ListAuditEventAnalyticsSortByEnum
func GetListAuditEventAnalyticsSortByEnumStringValues() []string {
	return []string{
		"targetId",
		"targetClass",
		"targetName",
		"objectType",
		"dbUserName",
		"eventName",
		"auditEventTime",
		"clientHostname",
		"clientProgram",
		"clientId",
		"auditType",
	}
}

// GetMappingListAuditEventAnalyticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditEventAnalyticsSortByEnum(val string) (ListAuditEventAnalyticsSortByEnum, bool) {
	enum, ok := mappingListAuditEventAnalyticsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
