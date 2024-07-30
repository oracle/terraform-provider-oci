// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package networkfirewall

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTunnelInspectionRulesRequest wrapper for the ListTunnelInspectionRules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListTunnelInspectionRules.go.html to see an example of how to use ListTunnelInspectionRulesRequest.
type ListTunnelInspectionRulesRequest struct {

	// Unique Network Firewall Policy identifier
	NetworkFirewallPolicyId *string `mandatory:"true" contributesTo:"path" name:"networkFirewallPolicyId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` or `opc-prev-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListTunnelInspectionRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListTunnelInspectionRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique priority order for Tunnel Inspection rules in the network firewall policy.
	TunnelInspectionRulePriorityOrder *int `mandatory:"false" contributesTo:"query" name:"tunnelInspectionRulePriorityOrder"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTunnelInspectionRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTunnelInspectionRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTunnelInspectionRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTunnelInspectionRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTunnelInspectionRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTunnelInspectionRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTunnelInspectionRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTunnelInspectionRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTunnelInspectionRulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTunnelInspectionRulesResponse wrapper for the ListTunnelInspectionRules operation
type ListTunnelInspectionRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TunnelInspectionRuleSummaryCollection instances
	TunnelInspectionRuleSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results exist. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. This is to get the page counts overall.
	OpcPageCount *string `presentIn:"header" name:"opc-page-count"`

	// For pagination of a list of items. This provides the count of total items across pages.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListTunnelInspectionRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTunnelInspectionRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTunnelInspectionRulesSortOrderEnum Enum with underlying type: string
type ListTunnelInspectionRulesSortOrderEnum string

// Set of constants representing the allowable values for ListTunnelInspectionRulesSortOrderEnum
const (
	ListTunnelInspectionRulesSortOrderAsc  ListTunnelInspectionRulesSortOrderEnum = "ASC"
	ListTunnelInspectionRulesSortOrderDesc ListTunnelInspectionRulesSortOrderEnum = "DESC"
)

var mappingListTunnelInspectionRulesSortOrderEnum = map[string]ListTunnelInspectionRulesSortOrderEnum{
	"ASC":  ListTunnelInspectionRulesSortOrderAsc,
	"DESC": ListTunnelInspectionRulesSortOrderDesc,
}

var mappingListTunnelInspectionRulesSortOrderEnumLowerCase = map[string]ListTunnelInspectionRulesSortOrderEnum{
	"asc":  ListTunnelInspectionRulesSortOrderAsc,
	"desc": ListTunnelInspectionRulesSortOrderDesc,
}

// GetListTunnelInspectionRulesSortOrderEnumValues Enumerates the set of values for ListTunnelInspectionRulesSortOrderEnum
func GetListTunnelInspectionRulesSortOrderEnumValues() []ListTunnelInspectionRulesSortOrderEnum {
	values := make([]ListTunnelInspectionRulesSortOrderEnum, 0)
	for _, v := range mappingListTunnelInspectionRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTunnelInspectionRulesSortOrderEnumStringValues Enumerates the set of values in String for ListTunnelInspectionRulesSortOrderEnum
func GetListTunnelInspectionRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTunnelInspectionRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTunnelInspectionRulesSortOrderEnum(val string) (ListTunnelInspectionRulesSortOrderEnum, bool) {
	enum, ok := mappingListTunnelInspectionRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTunnelInspectionRulesSortByEnum Enum with underlying type: string
type ListTunnelInspectionRulesSortByEnum string

// Set of constants representing the allowable values for ListTunnelInspectionRulesSortByEnum
const (
	ListTunnelInspectionRulesSortByTimecreated ListTunnelInspectionRulesSortByEnum = "timeCreated"
	ListTunnelInspectionRulesSortByDisplayname ListTunnelInspectionRulesSortByEnum = "displayName"
)

var mappingListTunnelInspectionRulesSortByEnum = map[string]ListTunnelInspectionRulesSortByEnum{
	"timeCreated": ListTunnelInspectionRulesSortByTimecreated,
	"displayName": ListTunnelInspectionRulesSortByDisplayname,
}

var mappingListTunnelInspectionRulesSortByEnumLowerCase = map[string]ListTunnelInspectionRulesSortByEnum{
	"timecreated": ListTunnelInspectionRulesSortByTimecreated,
	"displayname": ListTunnelInspectionRulesSortByDisplayname,
}

// GetListTunnelInspectionRulesSortByEnumValues Enumerates the set of values for ListTunnelInspectionRulesSortByEnum
func GetListTunnelInspectionRulesSortByEnumValues() []ListTunnelInspectionRulesSortByEnum {
	values := make([]ListTunnelInspectionRulesSortByEnum, 0)
	for _, v := range mappingListTunnelInspectionRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTunnelInspectionRulesSortByEnumStringValues Enumerates the set of values in String for ListTunnelInspectionRulesSortByEnum
func GetListTunnelInspectionRulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListTunnelInspectionRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTunnelInspectionRulesSortByEnum(val string) (ListTunnelInspectionRulesSortByEnum, bool) {
	enum, ok := mappingListTunnelInspectionRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
