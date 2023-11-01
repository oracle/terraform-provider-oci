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

// ListSqlFirewallViolationAnalyticsRequest wrapper for the ListSqlFirewallViolationAnalytics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlFirewallViolationAnalytics.go.html to see an example of how to use ListSqlFirewallViolationAnalyticsRequest.
type ListSqlFirewallViolationAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSqlFirewallViolationAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// An optional filter to return audit events whose creation time in the database is greater than and equal to the date-time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStarted"`

	// An optional filter to return audit events whose creation time in the database is less than and equal to the date-time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeEnded *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnded"`

	// Default time zone is UTC if no time zone provided. The date-time considerations of the resource will be in accordance with the specified time zone.
	QueryTimeZone *string `mandatory:"false" contributesTo:"query" name:"queryTimeZone"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSqlFirewallViolationAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// If this query parameter is specified, the result is sorted by this query parameter value.
	SortBy ListSqlFirewallViolationAnalyticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2
	// of the System for Cross-Domain Identity Management (SCIM) specification, which is available
	// at RFC3339 (https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions,
	// text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format.
	// (Numeric and boolean values should not be quoted.)
	// **Example:** query=(operationTime ge '2021-06-04T01-00-26') and (violationAction eq 'BLOCKED')
	ScimQuery *string `mandatory:"false" contributesTo:"query" name:"scimQuery"`

	// Specifies a subset of summarized fields to be returned in the response.
	SummaryField []ListSqlFirewallViolationAnalyticsSummaryFieldEnum `contributesTo:"query" name:"summaryField" omitEmpty:"true" collectionFormat:"multi"`

	// A groupBy can only be used in combination with summaryField parameter.
	// A groupBy value has to be a subset of the values mentioned in summaryField parameter.
	GroupBy []ListSqlFirewallViolationAnalyticsGroupByEnum `contributesTo:"query" name:"groupBy" omitEmpty:"true" collectionFormat:"multi"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlFirewallViolationAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlFirewallViolationAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlFirewallViolationAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlFirewallViolationAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlFirewallViolationAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlFirewallViolationAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSqlFirewallViolationAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlFirewallViolationAnalyticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSqlFirewallViolationAnalyticsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlFirewallViolationAnalyticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSqlFirewallViolationAnalyticsSortByEnumStringValues(), ",")))
	}
	for _, val := range request.SummaryField {
		if _, ok := GetMappingListSqlFirewallViolationAnalyticsSummaryFieldEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SummaryField: %s. Supported values are: %s.", val, strings.Join(GetListSqlFirewallViolationAnalyticsSummaryFieldEnumStringValues(), ",")))
		}
	}

	for _, val := range request.GroupBy {
		if _, ok := GetMappingListSqlFirewallViolationAnalyticsGroupByEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", val, strings.Join(GetListSqlFirewallViolationAnalyticsGroupByEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlFirewallViolationAnalyticsResponse wrapper for the ListSqlFirewallViolationAnalytics operation
type ListSqlFirewallViolationAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlFirewallViolationAnalyticsCollection instances
	SqlFirewallViolationAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSqlFirewallViolationAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlFirewallViolationAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlFirewallViolationAnalyticsAccessLevelEnum Enum with underlying type: string
type ListSqlFirewallViolationAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListSqlFirewallViolationAnalyticsAccessLevelEnum
const (
	ListSqlFirewallViolationAnalyticsAccessLevelRestricted ListSqlFirewallViolationAnalyticsAccessLevelEnum = "RESTRICTED"
	ListSqlFirewallViolationAnalyticsAccessLevelAccessible ListSqlFirewallViolationAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSqlFirewallViolationAnalyticsAccessLevelEnum = map[string]ListSqlFirewallViolationAnalyticsAccessLevelEnum{
	"RESTRICTED": ListSqlFirewallViolationAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListSqlFirewallViolationAnalyticsAccessLevelAccessible,
}

var mappingListSqlFirewallViolationAnalyticsAccessLevelEnumLowerCase = map[string]ListSqlFirewallViolationAnalyticsAccessLevelEnum{
	"restricted": ListSqlFirewallViolationAnalyticsAccessLevelRestricted,
	"accessible": ListSqlFirewallViolationAnalyticsAccessLevelAccessible,
}

// GetListSqlFirewallViolationAnalyticsAccessLevelEnumValues Enumerates the set of values for ListSqlFirewallViolationAnalyticsAccessLevelEnum
func GetListSqlFirewallViolationAnalyticsAccessLevelEnumValues() []ListSqlFirewallViolationAnalyticsAccessLevelEnum {
	values := make([]ListSqlFirewallViolationAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListSqlFirewallViolationAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallViolationAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListSqlFirewallViolationAnalyticsAccessLevelEnum
func GetListSqlFirewallViolationAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSqlFirewallViolationAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallViolationAnalyticsAccessLevelEnum(val string) (ListSqlFirewallViolationAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListSqlFirewallViolationAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallViolationAnalyticsSortOrderEnum Enum with underlying type: string
type ListSqlFirewallViolationAnalyticsSortOrderEnum string

// Set of constants representing the allowable values for ListSqlFirewallViolationAnalyticsSortOrderEnum
const (
	ListSqlFirewallViolationAnalyticsSortOrderAsc  ListSqlFirewallViolationAnalyticsSortOrderEnum = "ASC"
	ListSqlFirewallViolationAnalyticsSortOrderDesc ListSqlFirewallViolationAnalyticsSortOrderEnum = "DESC"
)

var mappingListSqlFirewallViolationAnalyticsSortOrderEnum = map[string]ListSqlFirewallViolationAnalyticsSortOrderEnum{
	"ASC":  ListSqlFirewallViolationAnalyticsSortOrderAsc,
	"DESC": ListSqlFirewallViolationAnalyticsSortOrderDesc,
}

var mappingListSqlFirewallViolationAnalyticsSortOrderEnumLowerCase = map[string]ListSqlFirewallViolationAnalyticsSortOrderEnum{
	"asc":  ListSqlFirewallViolationAnalyticsSortOrderAsc,
	"desc": ListSqlFirewallViolationAnalyticsSortOrderDesc,
}

// GetListSqlFirewallViolationAnalyticsSortOrderEnumValues Enumerates the set of values for ListSqlFirewallViolationAnalyticsSortOrderEnum
func GetListSqlFirewallViolationAnalyticsSortOrderEnumValues() []ListSqlFirewallViolationAnalyticsSortOrderEnum {
	values := make([]ListSqlFirewallViolationAnalyticsSortOrderEnum, 0)
	for _, v := range mappingListSqlFirewallViolationAnalyticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallViolationAnalyticsSortOrderEnumStringValues Enumerates the set of values in String for ListSqlFirewallViolationAnalyticsSortOrderEnum
func GetListSqlFirewallViolationAnalyticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSqlFirewallViolationAnalyticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallViolationAnalyticsSortOrderEnum(val string) (ListSqlFirewallViolationAnalyticsSortOrderEnum, bool) {
	enum, ok := mappingListSqlFirewallViolationAnalyticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallViolationAnalyticsSortByEnum Enum with underlying type: string
type ListSqlFirewallViolationAnalyticsSortByEnum string

// Set of constants representing the allowable values for ListSqlFirewallViolationAnalyticsSortByEnum
const (
	ListSqlFirewallViolationAnalyticsSortByDbusername        ListSqlFirewallViolationAnalyticsSortByEnum = "dbUserName"
	ListSqlFirewallViolationAnalyticsSortByTargetid          ListSqlFirewallViolationAnalyticsSortByEnum = "targetId"
	ListSqlFirewallViolationAnalyticsSortByTargetname        ListSqlFirewallViolationAnalyticsSortByEnum = "targetName"
	ListSqlFirewallViolationAnalyticsSortByOperationtime     ListSqlFirewallViolationAnalyticsSortByEnum = "operationTime"
	ListSqlFirewallViolationAnalyticsSortByTimecollected     ListSqlFirewallViolationAnalyticsSortByEnum = "timeCollected"
	ListSqlFirewallViolationAnalyticsSortByClientosusername  ListSqlFirewallViolationAnalyticsSortByEnum = "clientOsUserName"
	ListSqlFirewallViolationAnalyticsSortByOperation         ListSqlFirewallViolationAnalyticsSortByEnum = "operation"
	ListSqlFirewallViolationAnalyticsSortByCurrentdbusername ListSqlFirewallViolationAnalyticsSortByEnum = "currentDbUserName"
	ListSqlFirewallViolationAnalyticsSortBySqllevel          ListSqlFirewallViolationAnalyticsSortByEnum = "sqlLevel"
	ListSqlFirewallViolationAnalyticsSortByClientip          ListSqlFirewallViolationAnalyticsSortByEnum = "clientIp"
	ListSqlFirewallViolationAnalyticsSortByClientprogram     ListSqlFirewallViolationAnalyticsSortByEnum = "clientProgram"
	ListSqlFirewallViolationAnalyticsSortByViolationcause    ListSqlFirewallViolationAnalyticsSortByEnum = "violationCause"
	ListSqlFirewallViolationAnalyticsSortByViolationaction   ListSqlFirewallViolationAnalyticsSortByEnum = "violationAction"
	ListSqlFirewallViolationAnalyticsSortByViolationcount    ListSqlFirewallViolationAnalyticsSortByEnum = "violationCount"
)

var mappingListSqlFirewallViolationAnalyticsSortByEnum = map[string]ListSqlFirewallViolationAnalyticsSortByEnum{
	"dbUserName":        ListSqlFirewallViolationAnalyticsSortByDbusername,
	"targetId":          ListSqlFirewallViolationAnalyticsSortByTargetid,
	"targetName":        ListSqlFirewallViolationAnalyticsSortByTargetname,
	"operationTime":     ListSqlFirewallViolationAnalyticsSortByOperationtime,
	"timeCollected":     ListSqlFirewallViolationAnalyticsSortByTimecollected,
	"clientOsUserName":  ListSqlFirewallViolationAnalyticsSortByClientosusername,
	"operation":         ListSqlFirewallViolationAnalyticsSortByOperation,
	"currentDbUserName": ListSqlFirewallViolationAnalyticsSortByCurrentdbusername,
	"sqlLevel":          ListSqlFirewallViolationAnalyticsSortBySqllevel,
	"clientIp":          ListSqlFirewallViolationAnalyticsSortByClientip,
	"clientProgram":     ListSqlFirewallViolationAnalyticsSortByClientprogram,
	"violationCause":    ListSqlFirewallViolationAnalyticsSortByViolationcause,
	"violationAction":   ListSqlFirewallViolationAnalyticsSortByViolationaction,
	"violationCount":    ListSqlFirewallViolationAnalyticsSortByViolationcount,
}

var mappingListSqlFirewallViolationAnalyticsSortByEnumLowerCase = map[string]ListSqlFirewallViolationAnalyticsSortByEnum{
	"dbusername":        ListSqlFirewallViolationAnalyticsSortByDbusername,
	"targetid":          ListSqlFirewallViolationAnalyticsSortByTargetid,
	"targetname":        ListSqlFirewallViolationAnalyticsSortByTargetname,
	"operationtime":     ListSqlFirewallViolationAnalyticsSortByOperationtime,
	"timecollected":     ListSqlFirewallViolationAnalyticsSortByTimecollected,
	"clientosusername":  ListSqlFirewallViolationAnalyticsSortByClientosusername,
	"operation":         ListSqlFirewallViolationAnalyticsSortByOperation,
	"currentdbusername": ListSqlFirewallViolationAnalyticsSortByCurrentdbusername,
	"sqllevel":          ListSqlFirewallViolationAnalyticsSortBySqllevel,
	"clientip":          ListSqlFirewallViolationAnalyticsSortByClientip,
	"clientprogram":     ListSqlFirewallViolationAnalyticsSortByClientprogram,
	"violationcause":    ListSqlFirewallViolationAnalyticsSortByViolationcause,
	"violationaction":   ListSqlFirewallViolationAnalyticsSortByViolationaction,
	"violationcount":    ListSqlFirewallViolationAnalyticsSortByViolationcount,
}

// GetListSqlFirewallViolationAnalyticsSortByEnumValues Enumerates the set of values for ListSqlFirewallViolationAnalyticsSortByEnum
func GetListSqlFirewallViolationAnalyticsSortByEnumValues() []ListSqlFirewallViolationAnalyticsSortByEnum {
	values := make([]ListSqlFirewallViolationAnalyticsSortByEnum, 0)
	for _, v := range mappingListSqlFirewallViolationAnalyticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallViolationAnalyticsSortByEnumStringValues Enumerates the set of values in String for ListSqlFirewallViolationAnalyticsSortByEnum
func GetListSqlFirewallViolationAnalyticsSortByEnumStringValues() []string {
	return []string{
		"dbUserName",
		"targetId",
		"targetName",
		"operationTime",
		"timeCollected",
		"clientOsUserName",
		"operation",
		"currentDbUserName",
		"sqlLevel",
		"clientIp",
		"clientProgram",
		"violationCause",
		"violationAction",
		"violationCount",
	}
}

// GetMappingListSqlFirewallViolationAnalyticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallViolationAnalyticsSortByEnum(val string) (ListSqlFirewallViolationAnalyticsSortByEnum, bool) {
	enum, ok := mappingListSqlFirewallViolationAnalyticsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallViolationAnalyticsSummaryFieldEnum Enum with underlying type: string
type ListSqlFirewallViolationAnalyticsSummaryFieldEnum string

// Set of constants representing the allowable values for ListSqlFirewallViolationAnalyticsSummaryFieldEnum
const (
	ListSqlFirewallViolationAnalyticsSummaryFieldDbusername        ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "dbUserName"
	ListSqlFirewallViolationAnalyticsSummaryFieldTargetname        ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "targetName"
	ListSqlFirewallViolationAnalyticsSummaryFieldClientosusername  ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "clientOsUserName"
	ListSqlFirewallViolationAnalyticsSummaryFieldOperation         ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "operation"
	ListSqlFirewallViolationAnalyticsSummaryFieldSqltext           ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "sqlText"
	ListSqlFirewallViolationAnalyticsSummaryFieldCurrentdbusername ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "currentDbUserName"
	ListSqlFirewallViolationAnalyticsSummaryFieldSqllevel          ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "sqlLevel"
	ListSqlFirewallViolationAnalyticsSummaryFieldClientip          ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "clientIp"
	ListSqlFirewallViolationAnalyticsSummaryFieldClientprogram     ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "clientProgram"
	ListSqlFirewallViolationAnalyticsSummaryFieldViolationcause    ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "violationCause"
	ListSqlFirewallViolationAnalyticsSummaryFieldViolationaction   ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "violationAction"
	ListSqlFirewallViolationAnalyticsSummaryFieldSelects           ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "selects"
	ListSqlFirewallViolationAnalyticsSummaryFieldCreates           ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "creates"
	ListSqlFirewallViolationAnalyticsSummaryFieldAlters            ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "alters"
	ListSqlFirewallViolationAnalyticsSummaryFieldDrops             ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "drops"
	ListSqlFirewallViolationAnalyticsSummaryFieldGrants            ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "grants"
	ListSqlFirewallViolationAnalyticsSummaryFieldRevokes           ListSqlFirewallViolationAnalyticsSummaryFieldEnum = "revokes"
)

var mappingListSqlFirewallViolationAnalyticsSummaryFieldEnum = map[string]ListSqlFirewallViolationAnalyticsSummaryFieldEnum{
	"dbUserName":        ListSqlFirewallViolationAnalyticsSummaryFieldDbusername,
	"targetName":        ListSqlFirewallViolationAnalyticsSummaryFieldTargetname,
	"clientOsUserName":  ListSqlFirewallViolationAnalyticsSummaryFieldClientosusername,
	"operation":         ListSqlFirewallViolationAnalyticsSummaryFieldOperation,
	"sqlText":           ListSqlFirewallViolationAnalyticsSummaryFieldSqltext,
	"currentDbUserName": ListSqlFirewallViolationAnalyticsSummaryFieldCurrentdbusername,
	"sqlLevel":          ListSqlFirewallViolationAnalyticsSummaryFieldSqllevel,
	"clientIp":          ListSqlFirewallViolationAnalyticsSummaryFieldClientip,
	"clientProgram":     ListSqlFirewallViolationAnalyticsSummaryFieldClientprogram,
	"violationCause":    ListSqlFirewallViolationAnalyticsSummaryFieldViolationcause,
	"violationAction":   ListSqlFirewallViolationAnalyticsSummaryFieldViolationaction,
	"selects":           ListSqlFirewallViolationAnalyticsSummaryFieldSelects,
	"creates":           ListSqlFirewallViolationAnalyticsSummaryFieldCreates,
	"alters":            ListSqlFirewallViolationAnalyticsSummaryFieldAlters,
	"drops":             ListSqlFirewallViolationAnalyticsSummaryFieldDrops,
	"grants":            ListSqlFirewallViolationAnalyticsSummaryFieldGrants,
	"revokes":           ListSqlFirewallViolationAnalyticsSummaryFieldRevokes,
}

var mappingListSqlFirewallViolationAnalyticsSummaryFieldEnumLowerCase = map[string]ListSqlFirewallViolationAnalyticsSummaryFieldEnum{
	"dbusername":        ListSqlFirewallViolationAnalyticsSummaryFieldDbusername,
	"targetname":        ListSqlFirewallViolationAnalyticsSummaryFieldTargetname,
	"clientosusername":  ListSqlFirewallViolationAnalyticsSummaryFieldClientosusername,
	"operation":         ListSqlFirewallViolationAnalyticsSummaryFieldOperation,
	"sqltext":           ListSqlFirewallViolationAnalyticsSummaryFieldSqltext,
	"currentdbusername": ListSqlFirewallViolationAnalyticsSummaryFieldCurrentdbusername,
	"sqllevel":          ListSqlFirewallViolationAnalyticsSummaryFieldSqllevel,
	"clientip":          ListSqlFirewallViolationAnalyticsSummaryFieldClientip,
	"clientprogram":     ListSqlFirewallViolationAnalyticsSummaryFieldClientprogram,
	"violationcause":    ListSqlFirewallViolationAnalyticsSummaryFieldViolationcause,
	"violationaction":   ListSqlFirewallViolationAnalyticsSummaryFieldViolationaction,
	"selects":           ListSqlFirewallViolationAnalyticsSummaryFieldSelects,
	"creates":           ListSqlFirewallViolationAnalyticsSummaryFieldCreates,
	"alters":            ListSqlFirewallViolationAnalyticsSummaryFieldAlters,
	"drops":             ListSqlFirewallViolationAnalyticsSummaryFieldDrops,
	"grants":            ListSqlFirewallViolationAnalyticsSummaryFieldGrants,
	"revokes":           ListSqlFirewallViolationAnalyticsSummaryFieldRevokes,
}

// GetListSqlFirewallViolationAnalyticsSummaryFieldEnumValues Enumerates the set of values for ListSqlFirewallViolationAnalyticsSummaryFieldEnum
func GetListSqlFirewallViolationAnalyticsSummaryFieldEnumValues() []ListSqlFirewallViolationAnalyticsSummaryFieldEnum {
	values := make([]ListSqlFirewallViolationAnalyticsSummaryFieldEnum, 0)
	for _, v := range mappingListSqlFirewallViolationAnalyticsSummaryFieldEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallViolationAnalyticsSummaryFieldEnumStringValues Enumerates the set of values in String for ListSqlFirewallViolationAnalyticsSummaryFieldEnum
func GetListSqlFirewallViolationAnalyticsSummaryFieldEnumStringValues() []string {
	return []string{
		"dbUserName",
		"targetName",
		"clientOsUserName",
		"operation",
		"sqlText",
		"currentDbUserName",
		"sqlLevel",
		"clientIp",
		"clientProgram",
		"violationCause",
		"violationAction",
		"selects",
		"creates",
		"alters",
		"drops",
		"grants",
		"revokes",
	}
}

// GetMappingListSqlFirewallViolationAnalyticsSummaryFieldEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallViolationAnalyticsSummaryFieldEnum(val string) (ListSqlFirewallViolationAnalyticsSummaryFieldEnum, bool) {
	enum, ok := mappingListSqlFirewallViolationAnalyticsSummaryFieldEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallViolationAnalyticsGroupByEnum Enum with underlying type: string
type ListSqlFirewallViolationAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListSqlFirewallViolationAnalyticsGroupByEnum
const (
	ListSqlFirewallViolationAnalyticsGroupByDbusername        ListSqlFirewallViolationAnalyticsGroupByEnum = "dbUserName"
	ListSqlFirewallViolationAnalyticsGroupByTargetname        ListSqlFirewallViolationAnalyticsGroupByEnum = "targetName"
	ListSqlFirewallViolationAnalyticsGroupByOperationtime     ListSqlFirewallViolationAnalyticsGroupByEnum = "operationTime"
	ListSqlFirewallViolationAnalyticsGroupByTimecollected     ListSqlFirewallViolationAnalyticsGroupByEnum = "timeCollected"
	ListSqlFirewallViolationAnalyticsGroupByClientosusername  ListSqlFirewallViolationAnalyticsGroupByEnum = "clientOsUserName"
	ListSqlFirewallViolationAnalyticsGroupByOperation         ListSqlFirewallViolationAnalyticsGroupByEnum = "operation"
	ListSqlFirewallViolationAnalyticsGroupBySqltext           ListSqlFirewallViolationAnalyticsGroupByEnum = "sqlText"
	ListSqlFirewallViolationAnalyticsGroupByCurrentdbusername ListSqlFirewallViolationAnalyticsGroupByEnum = "currentDbUserName"
	ListSqlFirewallViolationAnalyticsGroupBySqllevel          ListSqlFirewallViolationAnalyticsGroupByEnum = "sqlLevel"
	ListSqlFirewallViolationAnalyticsGroupByClientip          ListSqlFirewallViolationAnalyticsGroupByEnum = "clientIp"
	ListSqlFirewallViolationAnalyticsGroupByClientprogram     ListSqlFirewallViolationAnalyticsGroupByEnum = "clientProgram"
	ListSqlFirewallViolationAnalyticsGroupByViolationcause    ListSqlFirewallViolationAnalyticsGroupByEnum = "violationCause"
	ListSqlFirewallViolationAnalyticsGroupByViolationaction   ListSqlFirewallViolationAnalyticsGroupByEnum = "violationAction"
)

var mappingListSqlFirewallViolationAnalyticsGroupByEnum = map[string]ListSqlFirewallViolationAnalyticsGroupByEnum{
	"dbUserName":        ListSqlFirewallViolationAnalyticsGroupByDbusername,
	"targetName":        ListSqlFirewallViolationAnalyticsGroupByTargetname,
	"operationTime":     ListSqlFirewallViolationAnalyticsGroupByOperationtime,
	"timeCollected":     ListSqlFirewallViolationAnalyticsGroupByTimecollected,
	"clientOsUserName":  ListSqlFirewallViolationAnalyticsGroupByClientosusername,
	"operation":         ListSqlFirewallViolationAnalyticsGroupByOperation,
	"sqlText":           ListSqlFirewallViolationAnalyticsGroupBySqltext,
	"currentDbUserName": ListSqlFirewallViolationAnalyticsGroupByCurrentdbusername,
	"sqlLevel":          ListSqlFirewallViolationAnalyticsGroupBySqllevel,
	"clientIp":          ListSqlFirewallViolationAnalyticsGroupByClientip,
	"clientProgram":     ListSqlFirewallViolationAnalyticsGroupByClientprogram,
	"violationCause":    ListSqlFirewallViolationAnalyticsGroupByViolationcause,
	"violationAction":   ListSqlFirewallViolationAnalyticsGroupByViolationaction,
}

var mappingListSqlFirewallViolationAnalyticsGroupByEnumLowerCase = map[string]ListSqlFirewallViolationAnalyticsGroupByEnum{
	"dbusername":        ListSqlFirewallViolationAnalyticsGroupByDbusername,
	"targetname":        ListSqlFirewallViolationAnalyticsGroupByTargetname,
	"operationtime":     ListSqlFirewallViolationAnalyticsGroupByOperationtime,
	"timecollected":     ListSqlFirewallViolationAnalyticsGroupByTimecollected,
	"clientosusername":  ListSqlFirewallViolationAnalyticsGroupByClientosusername,
	"operation":         ListSqlFirewallViolationAnalyticsGroupByOperation,
	"sqltext":           ListSqlFirewallViolationAnalyticsGroupBySqltext,
	"currentdbusername": ListSqlFirewallViolationAnalyticsGroupByCurrentdbusername,
	"sqllevel":          ListSqlFirewallViolationAnalyticsGroupBySqllevel,
	"clientip":          ListSqlFirewallViolationAnalyticsGroupByClientip,
	"clientprogram":     ListSqlFirewallViolationAnalyticsGroupByClientprogram,
	"violationcause":    ListSqlFirewallViolationAnalyticsGroupByViolationcause,
	"violationaction":   ListSqlFirewallViolationAnalyticsGroupByViolationaction,
}

// GetListSqlFirewallViolationAnalyticsGroupByEnumValues Enumerates the set of values for ListSqlFirewallViolationAnalyticsGroupByEnum
func GetListSqlFirewallViolationAnalyticsGroupByEnumValues() []ListSqlFirewallViolationAnalyticsGroupByEnum {
	values := make([]ListSqlFirewallViolationAnalyticsGroupByEnum, 0)
	for _, v := range mappingListSqlFirewallViolationAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallViolationAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListSqlFirewallViolationAnalyticsGroupByEnum
func GetListSqlFirewallViolationAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"dbUserName",
		"targetName",
		"operationTime",
		"timeCollected",
		"clientOsUserName",
		"operation",
		"sqlText",
		"currentDbUserName",
		"sqlLevel",
		"clientIp",
		"clientProgram",
		"violationCause",
		"violationAction",
	}
}

// GetMappingListSqlFirewallViolationAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallViolationAnalyticsGroupByEnum(val string) (ListSqlFirewallViolationAnalyticsGroupByEnum, bool) {
	enum, ok := mappingListSqlFirewallViolationAnalyticsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
