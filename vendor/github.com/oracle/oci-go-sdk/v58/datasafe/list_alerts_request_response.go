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

// ListAlertsRequest wrapper for the ListAlerts operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAlerts.go.html to see an example of how to use ListAlertsRequest.
type ListAlertsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return alert by it's OCID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

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
	AccessLevel ListAlertsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListAlertsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListAlertsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2
	// of the System for Cross-Domain Identity Management (SCIM) specification, which is available
	// at RFC3339 (https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions,
	// text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format.
	// (Numeric and boolean values should not be quoted.)
	// **Example:** query=(timeCreated ge '2021-06-04T01-00-26') and (targetNames eq 'target_1')
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
	ScimQuery *string `mandatory:"false" contributesTo:"query" name:"scimQuery"`

	// Specifies a subset of fields to be returned in the response.
	Field []ListAlertsFieldEnum `contributesTo:"query" name:"field" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAlertsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAlertsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAlertsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAlertsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAlertsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAlertsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListAlertsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlertsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAlertsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlertsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAlertsSortByEnumStringValues(), ",")))
	}
	for _, val := range request.Field {
		if _, ok := GetMappingListAlertsFieldEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Field: %s. Supported values are: %s.", val, strings.Join(GetListAlertsFieldEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAlertsResponse wrapper for the ListAlerts operation
type ListAlertsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AlertCollection instances
	AlertCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListAlertsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAlertsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAlertsAccessLevelEnum Enum with underlying type: string
type ListAlertsAccessLevelEnum string

// Set of constants representing the allowable values for ListAlertsAccessLevelEnum
const (
	ListAlertsAccessLevelRestricted ListAlertsAccessLevelEnum = "RESTRICTED"
	ListAlertsAccessLevelAccessible ListAlertsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListAlertsAccessLevelEnum = map[string]ListAlertsAccessLevelEnum{
	"RESTRICTED": ListAlertsAccessLevelRestricted,
	"ACCESSIBLE": ListAlertsAccessLevelAccessible,
}

// GetListAlertsAccessLevelEnumValues Enumerates the set of values for ListAlertsAccessLevelEnum
func GetListAlertsAccessLevelEnumValues() []ListAlertsAccessLevelEnum {
	values := make([]ListAlertsAccessLevelEnum, 0)
	for _, v := range mappingListAlertsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertsAccessLevelEnumStringValues Enumerates the set of values in String for ListAlertsAccessLevelEnum
func GetListAlertsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListAlertsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertsAccessLevelEnum(val string) (ListAlertsAccessLevelEnum, bool) {
	mappingListAlertsAccessLevelEnumIgnoreCase := make(map[string]ListAlertsAccessLevelEnum)
	for k, v := range mappingListAlertsAccessLevelEnum {
		mappingListAlertsAccessLevelEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAlertsAccessLevelEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertsSortOrderEnum Enum with underlying type: string
type ListAlertsSortOrderEnum string

// Set of constants representing the allowable values for ListAlertsSortOrderEnum
const (
	ListAlertsSortOrderAsc  ListAlertsSortOrderEnum = "ASC"
	ListAlertsSortOrderDesc ListAlertsSortOrderEnum = "DESC"
)

var mappingListAlertsSortOrderEnum = map[string]ListAlertsSortOrderEnum{
	"ASC":  ListAlertsSortOrderAsc,
	"DESC": ListAlertsSortOrderDesc,
}

// GetListAlertsSortOrderEnumValues Enumerates the set of values for ListAlertsSortOrderEnum
func GetListAlertsSortOrderEnumValues() []ListAlertsSortOrderEnum {
	values := make([]ListAlertsSortOrderEnum, 0)
	for _, v := range mappingListAlertsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertsSortOrderEnumStringValues Enumerates the set of values in String for ListAlertsSortOrderEnum
func GetListAlertsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAlertsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertsSortOrderEnum(val string) (ListAlertsSortOrderEnum, bool) {
	mappingListAlertsSortOrderEnumIgnoreCase := make(map[string]ListAlertsSortOrderEnum)
	for k, v := range mappingListAlertsSortOrderEnum {
		mappingListAlertsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAlertsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertsSortByEnum Enum with underlying type: string
type ListAlertsSortByEnum string

// Set of constants representing the allowable values for ListAlertsSortByEnum
const (
	ListAlertsSortByDisplayname ListAlertsSortByEnum = "displayName"
	ListAlertsSortByTimecreated ListAlertsSortByEnum = "timeCreated"
)

var mappingListAlertsSortByEnum = map[string]ListAlertsSortByEnum{
	"displayName": ListAlertsSortByDisplayname,
	"timeCreated": ListAlertsSortByTimecreated,
}

// GetListAlertsSortByEnumValues Enumerates the set of values for ListAlertsSortByEnum
func GetListAlertsSortByEnumValues() []ListAlertsSortByEnum {
	values := make([]ListAlertsSortByEnum, 0)
	for _, v := range mappingListAlertsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertsSortByEnumStringValues Enumerates the set of values in String for ListAlertsSortByEnum
func GetListAlertsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListAlertsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertsSortByEnum(val string) (ListAlertsSortByEnum, bool) {
	mappingListAlertsSortByEnumIgnoreCase := make(map[string]ListAlertsSortByEnum)
	for k, v := range mappingListAlertsSortByEnum {
		mappingListAlertsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAlertsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertsFieldEnum Enum with underlying type: string
type ListAlertsFieldEnum string

// Set of constants representing the allowable values for ListAlertsFieldEnum
const (
	ListAlertsFieldId              ListAlertsFieldEnum = "id"
	ListAlertsFieldDisplayname     ListAlertsFieldEnum = "displayName"
	ListAlertsFieldAlerttype       ListAlertsFieldEnum = "alertType"
	ListAlertsFieldTargetids       ListAlertsFieldEnum = "targetIds"
	ListAlertsFieldTargetnames     ListAlertsFieldEnum = "targetNames"
	ListAlertsFieldSeverity        ListAlertsFieldEnum = "severity"
	ListAlertsFieldStatus          ListAlertsFieldEnum = "status"
	ListAlertsFieldOperationtime   ListAlertsFieldEnum = "operationTime"
	ListAlertsFieldOperation       ListAlertsFieldEnum = "operation"
	ListAlertsFieldOperationstatus ListAlertsFieldEnum = "operationStatus"
	ListAlertsFieldTimecreated     ListAlertsFieldEnum = "timeCreated"
	ListAlertsFieldTimeupdated     ListAlertsFieldEnum = "timeUpdated"
	ListAlertsFieldPolicyid        ListAlertsFieldEnum = "policyId"
	ListAlertsFieldLifecyclestate  ListAlertsFieldEnum = "lifecycleState"
)

var mappingListAlertsFieldEnum = map[string]ListAlertsFieldEnum{
	"id":              ListAlertsFieldId,
	"displayName":     ListAlertsFieldDisplayname,
	"alertType":       ListAlertsFieldAlerttype,
	"targetIds":       ListAlertsFieldTargetids,
	"targetNames":     ListAlertsFieldTargetnames,
	"severity":        ListAlertsFieldSeverity,
	"status":          ListAlertsFieldStatus,
	"operationTime":   ListAlertsFieldOperationtime,
	"operation":       ListAlertsFieldOperation,
	"operationStatus": ListAlertsFieldOperationstatus,
	"timeCreated":     ListAlertsFieldTimecreated,
	"timeUpdated":     ListAlertsFieldTimeupdated,
	"policyId":        ListAlertsFieldPolicyid,
	"lifecycleState":  ListAlertsFieldLifecyclestate,
}

// GetListAlertsFieldEnumValues Enumerates the set of values for ListAlertsFieldEnum
func GetListAlertsFieldEnumValues() []ListAlertsFieldEnum {
	values := make([]ListAlertsFieldEnum, 0)
	for _, v := range mappingListAlertsFieldEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertsFieldEnumStringValues Enumerates the set of values in String for ListAlertsFieldEnum
func GetListAlertsFieldEnumStringValues() []string {
	return []string{
		"id",
		"displayName",
		"alertType",
		"targetIds",
		"targetNames",
		"severity",
		"status",
		"operationTime",
		"operation",
		"operationStatus",
		"timeCreated",
		"timeUpdated",
		"policyId",
		"lifecycleState",
	}
}

// GetMappingListAlertsFieldEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertsFieldEnum(val string) (ListAlertsFieldEnum, bool) {
	mappingListAlertsFieldEnumIgnoreCase := make(map[string]ListAlertsFieldEnum)
	for k, v := range mappingListAlertsFieldEnum {
		mappingListAlertsFieldEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAlertsFieldEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
