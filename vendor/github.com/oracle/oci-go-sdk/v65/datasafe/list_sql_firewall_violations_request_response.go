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

// ListSqlFirewallViolationsRequest wrapper for the ListSqlFirewallViolations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlFirewallViolations.go.html to see an example of how to use ListSqlFirewallViolationsRequest.
type ListSqlFirewallViolationsRequest struct {

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

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSqlFirewallViolationsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSqlFirewallViolationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// If this query parameter is specified, the result is sorted by this query parameter value.
	SortBy ListSqlFirewallViolationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2
	// of the System for Cross-Domain Identity Management (SCIM) specification, which is available
	// at RFC3339 (https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions,
	// text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format.
	// (Numeric and boolean values should not be quoted.)
	// **Example:** query=(operationTime ge '2021-06-04T01-00-26') and (violationAction eq 'BLOCKED')
	ScimQuery *string `mandatory:"false" contributesTo:"query" name:"scimQuery"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlFirewallViolationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlFirewallViolationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlFirewallViolationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlFirewallViolationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlFirewallViolationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlFirewallViolationsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSqlFirewallViolationsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlFirewallViolationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSqlFirewallViolationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlFirewallViolationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSqlFirewallViolationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlFirewallViolationsResponse wrapper for the ListSqlFirewallViolations operation
type ListSqlFirewallViolationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlFirewallViolationsCollection instances
	SqlFirewallViolationsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSqlFirewallViolationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlFirewallViolationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlFirewallViolationsAccessLevelEnum Enum with underlying type: string
type ListSqlFirewallViolationsAccessLevelEnum string

// Set of constants representing the allowable values for ListSqlFirewallViolationsAccessLevelEnum
const (
	ListSqlFirewallViolationsAccessLevelRestricted ListSqlFirewallViolationsAccessLevelEnum = "RESTRICTED"
	ListSqlFirewallViolationsAccessLevelAccessible ListSqlFirewallViolationsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSqlFirewallViolationsAccessLevelEnum = map[string]ListSqlFirewallViolationsAccessLevelEnum{
	"RESTRICTED": ListSqlFirewallViolationsAccessLevelRestricted,
	"ACCESSIBLE": ListSqlFirewallViolationsAccessLevelAccessible,
}

var mappingListSqlFirewallViolationsAccessLevelEnumLowerCase = map[string]ListSqlFirewallViolationsAccessLevelEnum{
	"restricted": ListSqlFirewallViolationsAccessLevelRestricted,
	"accessible": ListSqlFirewallViolationsAccessLevelAccessible,
}

// GetListSqlFirewallViolationsAccessLevelEnumValues Enumerates the set of values for ListSqlFirewallViolationsAccessLevelEnum
func GetListSqlFirewallViolationsAccessLevelEnumValues() []ListSqlFirewallViolationsAccessLevelEnum {
	values := make([]ListSqlFirewallViolationsAccessLevelEnum, 0)
	for _, v := range mappingListSqlFirewallViolationsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallViolationsAccessLevelEnumStringValues Enumerates the set of values in String for ListSqlFirewallViolationsAccessLevelEnum
func GetListSqlFirewallViolationsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSqlFirewallViolationsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallViolationsAccessLevelEnum(val string) (ListSqlFirewallViolationsAccessLevelEnum, bool) {
	enum, ok := mappingListSqlFirewallViolationsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallViolationsSortOrderEnum Enum with underlying type: string
type ListSqlFirewallViolationsSortOrderEnum string

// Set of constants representing the allowable values for ListSqlFirewallViolationsSortOrderEnum
const (
	ListSqlFirewallViolationsSortOrderAsc  ListSqlFirewallViolationsSortOrderEnum = "ASC"
	ListSqlFirewallViolationsSortOrderDesc ListSqlFirewallViolationsSortOrderEnum = "DESC"
)

var mappingListSqlFirewallViolationsSortOrderEnum = map[string]ListSqlFirewallViolationsSortOrderEnum{
	"ASC":  ListSqlFirewallViolationsSortOrderAsc,
	"DESC": ListSqlFirewallViolationsSortOrderDesc,
}

var mappingListSqlFirewallViolationsSortOrderEnumLowerCase = map[string]ListSqlFirewallViolationsSortOrderEnum{
	"asc":  ListSqlFirewallViolationsSortOrderAsc,
	"desc": ListSqlFirewallViolationsSortOrderDesc,
}

// GetListSqlFirewallViolationsSortOrderEnumValues Enumerates the set of values for ListSqlFirewallViolationsSortOrderEnum
func GetListSqlFirewallViolationsSortOrderEnumValues() []ListSqlFirewallViolationsSortOrderEnum {
	values := make([]ListSqlFirewallViolationsSortOrderEnum, 0)
	for _, v := range mappingListSqlFirewallViolationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallViolationsSortOrderEnumStringValues Enumerates the set of values in String for ListSqlFirewallViolationsSortOrderEnum
func GetListSqlFirewallViolationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSqlFirewallViolationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallViolationsSortOrderEnum(val string) (ListSqlFirewallViolationsSortOrderEnum, bool) {
	enum, ok := mappingListSqlFirewallViolationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallViolationsSortByEnum Enum with underlying type: string
type ListSqlFirewallViolationsSortByEnum string

// Set of constants representing the allowable values for ListSqlFirewallViolationsSortByEnum
const (
	ListSqlFirewallViolationsSortByDbusername        ListSqlFirewallViolationsSortByEnum = "dbUserName"
	ListSqlFirewallViolationsSortByTargetid          ListSqlFirewallViolationsSortByEnum = "targetId"
	ListSqlFirewallViolationsSortByTargetname        ListSqlFirewallViolationsSortByEnum = "targetName"
	ListSqlFirewallViolationsSortByOperationtime     ListSqlFirewallViolationsSortByEnum = "operationTime"
	ListSqlFirewallViolationsSortByTimecollected     ListSqlFirewallViolationsSortByEnum = "timeCollected"
	ListSqlFirewallViolationsSortByClientosusername  ListSqlFirewallViolationsSortByEnum = "clientOsUserName"
	ListSqlFirewallViolationsSortByOperation         ListSqlFirewallViolationsSortByEnum = "operation"
	ListSqlFirewallViolationsSortByCurrentdbusername ListSqlFirewallViolationsSortByEnum = "currentDbUserName"
	ListSqlFirewallViolationsSortBySqllevel          ListSqlFirewallViolationsSortByEnum = "sqlLevel"
	ListSqlFirewallViolationsSortByClientip          ListSqlFirewallViolationsSortByEnum = "clientIp"
	ListSqlFirewallViolationsSortByClientprogram     ListSqlFirewallViolationsSortByEnum = "clientProgram"
	ListSqlFirewallViolationsSortByViolationcause    ListSqlFirewallViolationsSortByEnum = "violationCause"
	ListSqlFirewallViolationsSortByViolationaction   ListSqlFirewallViolationsSortByEnum = "violationAction"
)

var mappingListSqlFirewallViolationsSortByEnum = map[string]ListSqlFirewallViolationsSortByEnum{
	"dbUserName":        ListSqlFirewallViolationsSortByDbusername,
	"targetId":          ListSqlFirewallViolationsSortByTargetid,
	"targetName":        ListSqlFirewallViolationsSortByTargetname,
	"operationTime":     ListSqlFirewallViolationsSortByOperationtime,
	"timeCollected":     ListSqlFirewallViolationsSortByTimecollected,
	"clientOsUserName":  ListSqlFirewallViolationsSortByClientosusername,
	"operation":         ListSqlFirewallViolationsSortByOperation,
	"currentDbUserName": ListSqlFirewallViolationsSortByCurrentdbusername,
	"sqlLevel":          ListSqlFirewallViolationsSortBySqllevel,
	"clientIp":          ListSqlFirewallViolationsSortByClientip,
	"clientProgram":     ListSqlFirewallViolationsSortByClientprogram,
	"violationCause":    ListSqlFirewallViolationsSortByViolationcause,
	"violationAction":   ListSqlFirewallViolationsSortByViolationaction,
}

var mappingListSqlFirewallViolationsSortByEnumLowerCase = map[string]ListSqlFirewallViolationsSortByEnum{
	"dbusername":        ListSqlFirewallViolationsSortByDbusername,
	"targetid":          ListSqlFirewallViolationsSortByTargetid,
	"targetname":        ListSqlFirewallViolationsSortByTargetname,
	"operationtime":     ListSqlFirewallViolationsSortByOperationtime,
	"timecollected":     ListSqlFirewallViolationsSortByTimecollected,
	"clientosusername":  ListSqlFirewallViolationsSortByClientosusername,
	"operation":         ListSqlFirewallViolationsSortByOperation,
	"currentdbusername": ListSqlFirewallViolationsSortByCurrentdbusername,
	"sqllevel":          ListSqlFirewallViolationsSortBySqllevel,
	"clientip":          ListSqlFirewallViolationsSortByClientip,
	"clientprogram":     ListSqlFirewallViolationsSortByClientprogram,
	"violationcause":    ListSqlFirewallViolationsSortByViolationcause,
	"violationaction":   ListSqlFirewallViolationsSortByViolationaction,
}

// GetListSqlFirewallViolationsSortByEnumValues Enumerates the set of values for ListSqlFirewallViolationsSortByEnum
func GetListSqlFirewallViolationsSortByEnumValues() []ListSqlFirewallViolationsSortByEnum {
	values := make([]ListSqlFirewallViolationsSortByEnum, 0)
	for _, v := range mappingListSqlFirewallViolationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallViolationsSortByEnumStringValues Enumerates the set of values in String for ListSqlFirewallViolationsSortByEnum
func GetListSqlFirewallViolationsSortByEnumStringValues() []string {
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
	}
}

// GetMappingListSqlFirewallViolationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallViolationsSortByEnum(val string) (ListSqlFirewallViolationsSortByEnum, bool) {
	enum, ok := mappingListSqlFirewallViolationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
