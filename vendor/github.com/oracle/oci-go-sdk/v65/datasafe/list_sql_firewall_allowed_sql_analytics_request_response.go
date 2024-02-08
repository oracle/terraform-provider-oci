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

// ListSqlFirewallAllowedSqlAnalyticsRequest wrapper for the ListSqlFirewallAllowedSqlAnalytics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlFirewallAllowedSqlAnalytics.go.html to see an example of how to use ListSqlFirewallAllowedSqlAnalyticsRequest.
type ListSqlFirewallAllowedSqlAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

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

	// The group by parameter to summarize the allowed SQL aggregation.
	GroupBy []ListSqlFirewallAllowedSqlAnalyticsGroupByEnum `contributesTo:"query" name:"groupBy" omitEmpty:"true" collectionFormat:"multi"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlFirewallAllowedSqlAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlFirewallAllowedSqlAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlFirewallAllowedSqlAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlFirewallAllowedSqlAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlFirewallAllowedSqlAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSqlFirewallAllowedSqlAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	for _, val := range request.GroupBy {
		if _, ok := GetMappingListSqlFirewallAllowedSqlAnalyticsGroupByEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", val, strings.Join(GetListSqlFirewallAllowedSqlAnalyticsGroupByEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlFirewallAllowedSqlAnalyticsResponse wrapper for the ListSqlFirewallAllowedSqlAnalytics operation
type ListSqlFirewallAllowedSqlAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlFirewallAllowedSqlAnalyticsCollection instances
	SqlFirewallAllowedSqlAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSqlFirewallAllowedSqlAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlFirewallAllowedSqlAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum Enum with underlying type: string
type ListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum
const (
	ListSqlFirewallAllowedSqlAnalyticsAccessLevelRestricted ListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum = "RESTRICTED"
	ListSqlFirewallAllowedSqlAnalyticsAccessLevelAccessible ListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum = map[string]ListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum{
	"RESTRICTED": ListSqlFirewallAllowedSqlAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListSqlFirewallAllowedSqlAnalyticsAccessLevelAccessible,
}

var mappingListSqlFirewallAllowedSqlAnalyticsAccessLevelEnumLowerCase = map[string]ListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum{
	"restricted": ListSqlFirewallAllowedSqlAnalyticsAccessLevelRestricted,
	"accessible": ListSqlFirewallAllowedSqlAnalyticsAccessLevelAccessible,
}

// GetListSqlFirewallAllowedSqlAnalyticsAccessLevelEnumValues Enumerates the set of values for ListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum
func GetListSqlFirewallAllowedSqlAnalyticsAccessLevelEnumValues() []ListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum {
	values := make([]ListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallAllowedSqlAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum
func GetListSqlFirewallAllowedSqlAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum(val string) (ListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListSqlFirewallAllowedSqlAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallAllowedSqlAnalyticsGroupByEnum Enum with underlying type: string
type ListSqlFirewallAllowedSqlAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListSqlFirewallAllowedSqlAnalyticsGroupByEnum
const (
	ListSqlFirewallAllowedSqlAnalyticsGroupByDbusername          ListSqlFirewallAllowedSqlAnalyticsGroupByEnum = "dbUserName"
	ListSqlFirewallAllowedSqlAnalyticsGroupBySqllevel            ListSqlFirewallAllowedSqlAnalyticsGroupByEnum = "sqlLevel"
	ListSqlFirewallAllowedSqlAnalyticsGroupBySqlfirewallpolicyid ListSqlFirewallAllowedSqlAnalyticsGroupByEnum = "sqlFirewallPolicyId"
	ListSqlFirewallAllowedSqlAnalyticsGroupByLifecyclestate      ListSqlFirewallAllowedSqlAnalyticsGroupByEnum = "lifecycleState"
)

var mappingListSqlFirewallAllowedSqlAnalyticsGroupByEnum = map[string]ListSqlFirewallAllowedSqlAnalyticsGroupByEnum{
	"dbUserName":          ListSqlFirewallAllowedSqlAnalyticsGroupByDbusername,
	"sqlLevel":            ListSqlFirewallAllowedSqlAnalyticsGroupBySqllevel,
	"sqlFirewallPolicyId": ListSqlFirewallAllowedSqlAnalyticsGroupBySqlfirewallpolicyid,
	"lifecycleState":      ListSqlFirewallAllowedSqlAnalyticsGroupByLifecyclestate,
}

var mappingListSqlFirewallAllowedSqlAnalyticsGroupByEnumLowerCase = map[string]ListSqlFirewallAllowedSqlAnalyticsGroupByEnum{
	"dbusername":          ListSqlFirewallAllowedSqlAnalyticsGroupByDbusername,
	"sqllevel":            ListSqlFirewallAllowedSqlAnalyticsGroupBySqllevel,
	"sqlfirewallpolicyid": ListSqlFirewallAllowedSqlAnalyticsGroupBySqlfirewallpolicyid,
	"lifecyclestate":      ListSqlFirewallAllowedSqlAnalyticsGroupByLifecyclestate,
}

// GetListSqlFirewallAllowedSqlAnalyticsGroupByEnumValues Enumerates the set of values for ListSqlFirewallAllowedSqlAnalyticsGroupByEnum
func GetListSqlFirewallAllowedSqlAnalyticsGroupByEnumValues() []ListSqlFirewallAllowedSqlAnalyticsGroupByEnum {
	values := make([]ListSqlFirewallAllowedSqlAnalyticsGroupByEnum, 0)
	for _, v := range mappingListSqlFirewallAllowedSqlAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallAllowedSqlAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListSqlFirewallAllowedSqlAnalyticsGroupByEnum
func GetListSqlFirewallAllowedSqlAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"dbUserName",
		"sqlLevel",
		"sqlFirewallPolicyId",
		"lifecycleState",
	}
}

// GetMappingListSqlFirewallAllowedSqlAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallAllowedSqlAnalyticsGroupByEnum(val string) (ListSqlFirewallAllowedSqlAnalyticsGroupByEnum, bool) {
	enum, ok := mappingListSqlFirewallAllowedSqlAnalyticsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
