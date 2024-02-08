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

// ListSqlFirewallAllowedSqlsRequest wrapper for the ListSqlFirewallAllowedSqls operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlFirewallAllowedSqls.go.html to see an example of how to use ListSqlFirewallAllowedSqlsRequest.
type ListSqlFirewallAllowedSqlsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSqlFirewallAllowedSqlsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2
	// of the System for Cross-Domain Identity Management (SCIM) specification, which is available
	// at RFC3339 (https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions,
	// text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format.
	// (Numeric and boolean values should not be quoted.)
	// **Example:** query=(currentUser eq 'SCOTT') and (topLevel eq 'YES')
	ScimQuery *string `mandatory:"false" contributesTo:"query" name:"scimQuery"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSqlFirewallAllowedSqlsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort parameter should be provided.
	SortBy ListSqlFirewallAllowedSqlsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlFirewallAllowedSqlsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlFirewallAllowedSqlsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlFirewallAllowedSqlsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlFirewallAllowedSqlsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlFirewallAllowedSqlsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlFirewallAllowedSqlsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSqlFirewallAllowedSqlsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlFirewallAllowedSqlsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSqlFirewallAllowedSqlsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlFirewallAllowedSqlsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSqlFirewallAllowedSqlsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlFirewallAllowedSqlsResponse wrapper for the ListSqlFirewallAllowedSqls operation
type ListSqlFirewallAllowedSqlsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlFirewallAllowedSqlCollection instances
	SqlFirewallAllowedSqlCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSqlFirewallAllowedSqlsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlFirewallAllowedSqlsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlFirewallAllowedSqlsAccessLevelEnum Enum with underlying type: string
type ListSqlFirewallAllowedSqlsAccessLevelEnum string

// Set of constants representing the allowable values for ListSqlFirewallAllowedSqlsAccessLevelEnum
const (
	ListSqlFirewallAllowedSqlsAccessLevelRestricted ListSqlFirewallAllowedSqlsAccessLevelEnum = "RESTRICTED"
	ListSqlFirewallAllowedSqlsAccessLevelAccessible ListSqlFirewallAllowedSqlsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSqlFirewallAllowedSqlsAccessLevelEnum = map[string]ListSqlFirewallAllowedSqlsAccessLevelEnum{
	"RESTRICTED": ListSqlFirewallAllowedSqlsAccessLevelRestricted,
	"ACCESSIBLE": ListSqlFirewallAllowedSqlsAccessLevelAccessible,
}

var mappingListSqlFirewallAllowedSqlsAccessLevelEnumLowerCase = map[string]ListSqlFirewallAllowedSqlsAccessLevelEnum{
	"restricted": ListSqlFirewallAllowedSqlsAccessLevelRestricted,
	"accessible": ListSqlFirewallAllowedSqlsAccessLevelAccessible,
}

// GetListSqlFirewallAllowedSqlsAccessLevelEnumValues Enumerates the set of values for ListSqlFirewallAllowedSqlsAccessLevelEnum
func GetListSqlFirewallAllowedSqlsAccessLevelEnumValues() []ListSqlFirewallAllowedSqlsAccessLevelEnum {
	values := make([]ListSqlFirewallAllowedSqlsAccessLevelEnum, 0)
	for _, v := range mappingListSqlFirewallAllowedSqlsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallAllowedSqlsAccessLevelEnumStringValues Enumerates the set of values in String for ListSqlFirewallAllowedSqlsAccessLevelEnum
func GetListSqlFirewallAllowedSqlsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSqlFirewallAllowedSqlsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallAllowedSqlsAccessLevelEnum(val string) (ListSqlFirewallAllowedSqlsAccessLevelEnum, bool) {
	enum, ok := mappingListSqlFirewallAllowedSqlsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallAllowedSqlsSortOrderEnum Enum with underlying type: string
type ListSqlFirewallAllowedSqlsSortOrderEnum string

// Set of constants representing the allowable values for ListSqlFirewallAllowedSqlsSortOrderEnum
const (
	ListSqlFirewallAllowedSqlsSortOrderAsc  ListSqlFirewallAllowedSqlsSortOrderEnum = "ASC"
	ListSqlFirewallAllowedSqlsSortOrderDesc ListSqlFirewallAllowedSqlsSortOrderEnum = "DESC"
)

var mappingListSqlFirewallAllowedSqlsSortOrderEnum = map[string]ListSqlFirewallAllowedSqlsSortOrderEnum{
	"ASC":  ListSqlFirewallAllowedSqlsSortOrderAsc,
	"DESC": ListSqlFirewallAllowedSqlsSortOrderDesc,
}

var mappingListSqlFirewallAllowedSqlsSortOrderEnumLowerCase = map[string]ListSqlFirewallAllowedSqlsSortOrderEnum{
	"asc":  ListSqlFirewallAllowedSqlsSortOrderAsc,
	"desc": ListSqlFirewallAllowedSqlsSortOrderDesc,
}

// GetListSqlFirewallAllowedSqlsSortOrderEnumValues Enumerates the set of values for ListSqlFirewallAllowedSqlsSortOrderEnum
func GetListSqlFirewallAllowedSqlsSortOrderEnumValues() []ListSqlFirewallAllowedSqlsSortOrderEnum {
	values := make([]ListSqlFirewallAllowedSqlsSortOrderEnum, 0)
	for _, v := range mappingListSqlFirewallAllowedSqlsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallAllowedSqlsSortOrderEnumStringValues Enumerates the set of values in String for ListSqlFirewallAllowedSqlsSortOrderEnum
func GetListSqlFirewallAllowedSqlsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSqlFirewallAllowedSqlsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallAllowedSqlsSortOrderEnum(val string) (ListSqlFirewallAllowedSqlsSortOrderEnum, bool) {
	enum, ok := mappingListSqlFirewallAllowedSqlsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallAllowedSqlsSortByEnum Enum with underlying type: string
type ListSqlFirewallAllowedSqlsSortByEnum string

// Set of constants representing the allowable values for ListSqlFirewallAllowedSqlsSortByEnum
const (
	ListSqlFirewallAllowedSqlsSortByDisplayname   ListSqlFirewallAllowedSqlsSortByEnum = "displayName"
	ListSqlFirewallAllowedSqlsSortByTimecollected ListSqlFirewallAllowedSqlsSortByEnum = "timeCollected"
)

var mappingListSqlFirewallAllowedSqlsSortByEnum = map[string]ListSqlFirewallAllowedSqlsSortByEnum{
	"displayName":   ListSqlFirewallAllowedSqlsSortByDisplayname,
	"timeCollected": ListSqlFirewallAllowedSqlsSortByTimecollected,
}

var mappingListSqlFirewallAllowedSqlsSortByEnumLowerCase = map[string]ListSqlFirewallAllowedSqlsSortByEnum{
	"displayname":   ListSqlFirewallAllowedSqlsSortByDisplayname,
	"timecollected": ListSqlFirewallAllowedSqlsSortByTimecollected,
}

// GetListSqlFirewallAllowedSqlsSortByEnumValues Enumerates the set of values for ListSqlFirewallAllowedSqlsSortByEnum
func GetListSqlFirewallAllowedSqlsSortByEnumValues() []ListSqlFirewallAllowedSqlsSortByEnum {
	values := make([]ListSqlFirewallAllowedSqlsSortByEnum, 0)
	for _, v := range mappingListSqlFirewallAllowedSqlsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallAllowedSqlsSortByEnumStringValues Enumerates the set of values in String for ListSqlFirewallAllowedSqlsSortByEnum
func GetListSqlFirewallAllowedSqlsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCollected",
	}
}

// GetMappingListSqlFirewallAllowedSqlsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallAllowedSqlsSortByEnum(val string) (ListSqlFirewallAllowedSqlsSortByEnum, bool) {
	enum, ok := mappingListSqlFirewallAllowedSqlsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
