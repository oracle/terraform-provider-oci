// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListCustomProtectionRulesRequest wrapper for the ListCustomProtectionRules operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waas/ListCustomProtectionRules.go.html to see an example of how to use ListCustomProtectionRulesRequest.
type ListCustomProtectionRulesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This number is generated when the compartment is created.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated call. If unspecified, defaults to `10`.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous paginated call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The value by which custom protection rules are sorted in a paginated 'List' call. If unspecified, defaults to `timeCreated`.
	SortBy ListCustomProtectionRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The value of the sorting direction of resources in a paginated 'List' call. If unspecified, defaults to `DESC`.
	SortOrder ListCustomProtectionRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filter custom protection rules using a list of custom protection rule OCIDs.
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Filter custom protection rules using a list of display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// Filter Custom Protection rules using a list of lifecycle states.
	LifecycleState []LifecycleStatesEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter that matches Custom Protection rules created on or after the specified date-time.
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// A filter that matches custom protection rules created before the specified date-time.
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCustomProtectionRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCustomProtectionRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCustomProtectionRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCustomProtectionRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCustomProtectionRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCustomProtectionRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCustomProtectionRulesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCustomProtectionRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCustomProtectionRulesSortOrderEnumStringValues(), ",")))
	}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingLifecycleStatesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCustomProtectionRulesResponse wrapper for the ListCustomProtectionRules operation
type ListCustomProtectionRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CustomProtectionRuleSummary instances
	Items []CustomProtectionRuleSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results may remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCustomProtectionRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCustomProtectionRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCustomProtectionRulesSortByEnum Enum with underlying type: string
type ListCustomProtectionRulesSortByEnum string

// Set of constants representing the allowable values for ListCustomProtectionRulesSortByEnum
const (
	ListCustomProtectionRulesSortById                ListCustomProtectionRulesSortByEnum = "id"
	ListCustomProtectionRulesSortByCompartmentid     ListCustomProtectionRulesSortByEnum = "compartmentId"
	ListCustomProtectionRulesSortByDisplayname       ListCustomProtectionRulesSortByEnum = "displayName"
	ListCustomProtectionRulesSortByModsecurityruleid ListCustomProtectionRulesSortByEnum = "modSecurityRuleId"
	ListCustomProtectionRulesSortByTimecreated       ListCustomProtectionRulesSortByEnum = "timeCreated"
)

var mappingListCustomProtectionRulesSortByEnum = map[string]ListCustomProtectionRulesSortByEnum{
	"id":                ListCustomProtectionRulesSortById,
	"compartmentId":     ListCustomProtectionRulesSortByCompartmentid,
	"displayName":       ListCustomProtectionRulesSortByDisplayname,
	"modSecurityRuleId": ListCustomProtectionRulesSortByModsecurityruleid,
	"timeCreated":       ListCustomProtectionRulesSortByTimecreated,
}

// GetListCustomProtectionRulesSortByEnumValues Enumerates the set of values for ListCustomProtectionRulesSortByEnum
func GetListCustomProtectionRulesSortByEnumValues() []ListCustomProtectionRulesSortByEnum {
	values := make([]ListCustomProtectionRulesSortByEnum, 0)
	for _, v := range mappingListCustomProtectionRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomProtectionRulesSortByEnumStringValues Enumerates the set of values in String for ListCustomProtectionRulesSortByEnum
func GetListCustomProtectionRulesSortByEnumStringValues() []string {
	return []string{
		"id",
		"compartmentId",
		"displayName",
		"modSecurityRuleId",
		"timeCreated",
	}
}

// GetMappingListCustomProtectionRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomProtectionRulesSortByEnum(val string) (ListCustomProtectionRulesSortByEnum, bool) {
	mappingListCustomProtectionRulesSortByEnumIgnoreCase := make(map[string]ListCustomProtectionRulesSortByEnum)
	for k, v := range mappingListCustomProtectionRulesSortByEnum {
		mappingListCustomProtectionRulesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListCustomProtectionRulesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListCustomProtectionRulesSortOrderEnum Enum with underlying type: string
type ListCustomProtectionRulesSortOrderEnum string

// Set of constants representing the allowable values for ListCustomProtectionRulesSortOrderEnum
const (
	ListCustomProtectionRulesSortOrderAsc  ListCustomProtectionRulesSortOrderEnum = "ASC"
	ListCustomProtectionRulesSortOrderDesc ListCustomProtectionRulesSortOrderEnum = "DESC"
)

var mappingListCustomProtectionRulesSortOrderEnum = map[string]ListCustomProtectionRulesSortOrderEnum{
	"ASC":  ListCustomProtectionRulesSortOrderAsc,
	"DESC": ListCustomProtectionRulesSortOrderDesc,
}

// GetListCustomProtectionRulesSortOrderEnumValues Enumerates the set of values for ListCustomProtectionRulesSortOrderEnum
func GetListCustomProtectionRulesSortOrderEnumValues() []ListCustomProtectionRulesSortOrderEnum {
	values := make([]ListCustomProtectionRulesSortOrderEnum, 0)
	for _, v := range mappingListCustomProtectionRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomProtectionRulesSortOrderEnumStringValues Enumerates the set of values in String for ListCustomProtectionRulesSortOrderEnum
func GetListCustomProtectionRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCustomProtectionRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomProtectionRulesSortOrderEnum(val string) (ListCustomProtectionRulesSortOrderEnum, bool) {
	mappingListCustomProtectionRulesSortOrderEnumIgnoreCase := make(map[string]ListCustomProtectionRulesSortOrderEnum)
	for k, v := range mappingListCustomProtectionRulesSortOrderEnum {
		mappingListCustomProtectionRulesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListCustomProtectionRulesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
