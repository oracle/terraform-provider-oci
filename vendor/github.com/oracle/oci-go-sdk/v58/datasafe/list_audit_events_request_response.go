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

// ListAuditEventsRequest wrapper for the ListAuditEvents operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditEvents.go.html to see an example of how to use ListAuditEventsRequest.
type ListAuditEventsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListAuditEventsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// It is usually retrieved from a previous "List" call. For details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2
	// of the System for Cross-Domain Identity Management (SCIM) specification, which is available
	// at RFC3339 (https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions,
	// text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format.
	// (Numeric and boolean values should not be quoted.)
	// **Example:** query=(operationTime ge '2021-06-04T01-00-26') and (eventName eq 'LOGON')
	ScimQuery *string `mandatory:"false" contributesTo:"query" name:"scimQuery"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListAuditEventsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// If this query parameter is specified, the result is sorted by this query parameter value.
	SortBy ListAuditEventsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAuditEventsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAuditEventsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAuditEventsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAuditEventsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAuditEventsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAuditEventsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListAuditEventsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditEventsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAuditEventsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditEventsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAuditEventsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAuditEventsResponse wrapper for the ListAuditEvents operation
type ListAuditEventsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AuditEventCollection instances
	AuditEventCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListAuditEventsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAuditEventsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAuditEventsAccessLevelEnum Enum with underlying type: string
type ListAuditEventsAccessLevelEnum string

// Set of constants representing the allowable values for ListAuditEventsAccessLevelEnum
const (
	ListAuditEventsAccessLevelRestricted ListAuditEventsAccessLevelEnum = "RESTRICTED"
	ListAuditEventsAccessLevelAccessible ListAuditEventsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListAuditEventsAccessLevelEnum = map[string]ListAuditEventsAccessLevelEnum{
	"RESTRICTED": ListAuditEventsAccessLevelRestricted,
	"ACCESSIBLE": ListAuditEventsAccessLevelAccessible,
}

// GetListAuditEventsAccessLevelEnumValues Enumerates the set of values for ListAuditEventsAccessLevelEnum
func GetListAuditEventsAccessLevelEnumValues() []ListAuditEventsAccessLevelEnum {
	values := make([]ListAuditEventsAccessLevelEnum, 0)
	for _, v := range mappingListAuditEventsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditEventsAccessLevelEnumStringValues Enumerates the set of values in String for ListAuditEventsAccessLevelEnum
func GetListAuditEventsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListAuditEventsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditEventsAccessLevelEnum(val string) (ListAuditEventsAccessLevelEnum, bool) {
	mappingListAuditEventsAccessLevelEnumIgnoreCase := make(map[string]ListAuditEventsAccessLevelEnum)
	for k, v := range mappingListAuditEventsAccessLevelEnum {
		mappingListAuditEventsAccessLevelEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAuditEventsAccessLevelEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditEventsSortOrderEnum Enum with underlying type: string
type ListAuditEventsSortOrderEnum string

// Set of constants representing the allowable values for ListAuditEventsSortOrderEnum
const (
	ListAuditEventsSortOrderAsc  ListAuditEventsSortOrderEnum = "ASC"
	ListAuditEventsSortOrderDesc ListAuditEventsSortOrderEnum = "DESC"
)

var mappingListAuditEventsSortOrderEnum = map[string]ListAuditEventsSortOrderEnum{
	"ASC":  ListAuditEventsSortOrderAsc,
	"DESC": ListAuditEventsSortOrderDesc,
}

// GetListAuditEventsSortOrderEnumValues Enumerates the set of values for ListAuditEventsSortOrderEnum
func GetListAuditEventsSortOrderEnumValues() []ListAuditEventsSortOrderEnum {
	values := make([]ListAuditEventsSortOrderEnum, 0)
	for _, v := range mappingListAuditEventsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditEventsSortOrderEnumStringValues Enumerates the set of values in String for ListAuditEventsSortOrderEnum
func GetListAuditEventsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAuditEventsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditEventsSortOrderEnum(val string) (ListAuditEventsSortOrderEnum, bool) {
	mappingListAuditEventsSortOrderEnumIgnoreCase := make(map[string]ListAuditEventsSortOrderEnum)
	for k, v := range mappingListAuditEventsSortOrderEnum {
		mappingListAuditEventsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAuditEventsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditEventsSortByEnum Enum with underlying type: string
type ListAuditEventsSortByEnum string

// Set of constants representing the allowable values for ListAuditEventsSortByEnum
const (
	ListAuditEventsSortByDbusername              ListAuditEventsSortByEnum = "dbUserName"
	ListAuditEventsSortByTargetname              ListAuditEventsSortByEnum = "targetName"
	ListAuditEventsSortByDatabasetype            ListAuditEventsSortByEnum = "databaseType"
	ListAuditEventsSortByTargetclass             ListAuditEventsSortByEnum = "targetClass"
	ListAuditEventsSortByAuditeventtime          ListAuditEventsSortByEnum = "auditEventTime"
	ListAuditEventsSortByTimecollected           ListAuditEventsSortByEnum = "timeCollected"
	ListAuditEventsSortByOsusername              ListAuditEventsSortByEnum = "osUserName"
	ListAuditEventsSortByOperation               ListAuditEventsSortByEnum = "operation"
	ListAuditEventsSortByOperationstatus         ListAuditEventsSortByEnum = "operationStatus"
	ListAuditEventsSortByEventname               ListAuditEventsSortByEnum = "eventName"
	ListAuditEventsSortByErrorcode               ListAuditEventsSortByEnum = "errorCode"
	ListAuditEventsSortByErrormessage            ListAuditEventsSortByEnum = "errorMessage"
	ListAuditEventsSortByObjecttype              ListAuditEventsSortByEnum = "objectType"
	ListAuditEventsSortByObjectname              ListAuditEventsSortByEnum = "objectName"
	ListAuditEventsSortByObjectowner             ListAuditEventsSortByEnum = "objectOwner"
	ListAuditEventsSortByClienthostname          ListAuditEventsSortByEnum = "clientHostname"
	ListAuditEventsSortByClientip                ListAuditEventsSortByEnum = "clientIp"
	ListAuditEventsSortByIsalerted               ListAuditEventsSortByEnum = "isAlerted"
	ListAuditEventsSortByActiontaken             ListAuditEventsSortByEnum = "actionTaken"
	ListAuditEventsSortByClientprogram           ListAuditEventsSortByEnum = "clientProgram"
	ListAuditEventsSortByCommandtext             ListAuditEventsSortByEnum = "commandText"
	ListAuditEventsSortByCommandparam            ListAuditEventsSortByEnum = "commandParam"
	ListAuditEventsSortByExtendedeventattributes ListAuditEventsSortByEnum = "extendedEventAttributes"
	ListAuditEventsSortByAuditlocation           ListAuditEventsSortByEnum = "auditLocation"
	ListAuditEventsSortByOsterminal              ListAuditEventsSortByEnum = "osTerminal"
	ListAuditEventsSortByClientid                ListAuditEventsSortByEnum = "clientId"
	ListAuditEventsSortByAuditpolicies           ListAuditEventsSortByEnum = "auditPolicies"
	ListAuditEventsSortByAudittype               ListAuditEventsSortByEnum = "auditType"
)

var mappingListAuditEventsSortByEnum = map[string]ListAuditEventsSortByEnum{
	"dbUserName":              ListAuditEventsSortByDbusername,
	"targetName":              ListAuditEventsSortByTargetname,
	"databaseType":            ListAuditEventsSortByDatabasetype,
	"targetClass":             ListAuditEventsSortByTargetclass,
	"auditEventTime":          ListAuditEventsSortByAuditeventtime,
	"timeCollected":           ListAuditEventsSortByTimecollected,
	"osUserName":              ListAuditEventsSortByOsusername,
	"operation":               ListAuditEventsSortByOperation,
	"operationStatus":         ListAuditEventsSortByOperationstatus,
	"eventName":               ListAuditEventsSortByEventname,
	"errorCode":               ListAuditEventsSortByErrorcode,
	"errorMessage":            ListAuditEventsSortByErrormessage,
	"objectType":              ListAuditEventsSortByObjecttype,
	"objectName":              ListAuditEventsSortByObjectname,
	"objectOwner":             ListAuditEventsSortByObjectowner,
	"clientHostname":          ListAuditEventsSortByClienthostname,
	"clientIp":                ListAuditEventsSortByClientip,
	"isAlerted":               ListAuditEventsSortByIsalerted,
	"actionTaken":             ListAuditEventsSortByActiontaken,
	"clientProgram":           ListAuditEventsSortByClientprogram,
	"commandText":             ListAuditEventsSortByCommandtext,
	"commandParam":            ListAuditEventsSortByCommandparam,
	"extendedEventAttributes": ListAuditEventsSortByExtendedeventattributes,
	"auditLocation":           ListAuditEventsSortByAuditlocation,
	"osTerminal":              ListAuditEventsSortByOsterminal,
	"clientId":                ListAuditEventsSortByClientid,
	"auditPolicies":           ListAuditEventsSortByAuditpolicies,
	"auditType":               ListAuditEventsSortByAudittype,
}

// GetListAuditEventsSortByEnumValues Enumerates the set of values for ListAuditEventsSortByEnum
func GetListAuditEventsSortByEnumValues() []ListAuditEventsSortByEnum {
	values := make([]ListAuditEventsSortByEnum, 0)
	for _, v := range mappingListAuditEventsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditEventsSortByEnumStringValues Enumerates the set of values in String for ListAuditEventsSortByEnum
func GetListAuditEventsSortByEnumStringValues() []string {
	return []string{
		"dbUserName",
		"targetName",
		"databaseType",
		"targetClass",
		"auditEventTime",
		"timeCollected",
		"osUserName",
		"operation",
		"operationStatus",
		"eventName",
		"errorCode",
		"errorMessage",
		"objectType",
		"objectName",
		"objectOwner",
		"clientHostname",
		"clientIp",
		"isAlerted",
		"actionTaken",
		"clientProgram",
		"commandText",
		"commandParam",
		"extendedEventAttributes",
		"auditLocation",
		"osTerminal",
		"clientId",
		"auditPolicies",
		"auditType",
	}
}

// GetMappingListAuditEventsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditEventsSortByEnum(val string) (ListAuditEventsSortByEnum, bool) {
	mappingListAuditEventsSortByEnumIgnoreCase := make(map[string]ListAuditEventsSortByEnum)
	for k, v := range mappingListAuditEventsSortByEnum {
		mappingListAuditEventsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAuditEventsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
