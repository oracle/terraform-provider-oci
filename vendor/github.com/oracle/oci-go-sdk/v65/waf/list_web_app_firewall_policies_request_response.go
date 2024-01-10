// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waf

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWebAppFirewallPoliciesRequest wrapper for the ListWebAppFirewallPolicies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ListWebAppFirewallPolicies.go.html to see an example of how to use ListWebAppFirewallPoliciesRequest.
type ListWebAppFirewallPoliciesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycleState.
	LifecycleState []WebAppFirewallPolicyLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only the WebAppFirewallPolicy with the given OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListWebAppFirewallPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending.
	// Default order for displayName is ascending.
	// If no value is specified timeCreated is default.
	SortBy ListWebAppFirewallPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWebAppFirewallPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWebAppFirewallPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWebAppFirewallPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWebAppFirewallPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWebAppFirewallPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingWebAppFirewallPolicyLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetWebAppFirewallPolicyLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListWebAppFirewallPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWebAppFirewallPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWebAppFirewallPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWebAppFirewallPoliciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWebAppFirewallPoliciesResponse wrapper for the ListWebAppFirewallPolicies operation
type ListWebAppFirewallPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WebAppFirewallPolicyCollection instances
	WebAppFirewallPolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWebAppFirewallPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWebAppFirewallPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWebAppFirewallPoliciesSortOrderEnum Enum with underlying type: string
type ListWebAppFirewallPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListWebAppFirewallPoliciesSortOrderEnum
const (
	ListWebAppFirewallPoliciesSortOrderAsc  ListWebAppFirewallPoliciesSortOrderEnum = "ASC"
	ListWebAppFirewallPoliciesSortOrderDesc ListWebAppFirewallPoliciesSortOrderEnum = "DESC"
)

var mappingListWebAppFirewallPoliciesSortOrderEnum = map[string]ListWebAppFirewallPoliciesSortOrderEnum{
	"ASC":  ListWebAppFirewallPoliciesSortOrderAsc,
	"DESC": ListWebAppFirewallPoliciesSortOrderDesc,
}

var mappingListWebAppFirewallPoliciesSortOrderEnumLowerCase = map[string]ListWebAppFirewallPoliciesSortOrderEnum{
	"asc":  ListWebAppFirewallPoliciesSortOrderAsc,
	"desc": ListWebAppFirewallPoliciesSortOrderDesc,
}

// GetListWebAppFirewallPoliciesSortOrderEnumValues Enumerates the set of values for ListWebAppFirewallPoliciesSortOrderEnum
func GetListWebAppFirewallPoliciesSortOrderEnumValues() []ListWebAppFirewallPoliciesSortOrderEnum {
	values := make([]ListWebAppFirewallPoliciesSortOrderEnum, 0)
	for _, v := range mappingListWebAppFirewallPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWebAppFirewallPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListWebAppFirewallPoliciesSortOrderEnum
func GetListWebAppFirewallPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWebAppFirewallPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWebAppFirewallPoliciesSortOrderEnum(val string) (ListWebAppFirewallPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListWebAppFirewallPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWebAppFirewallPoliciesSortByEnum Enum with underlying type: string
type ListWebAppFirewallPoliciesSortByEnum string

// Set of constants representing the allowable values for ListWebAppFirewallPoliciesSortByEnum
const (
	ListWebAppFirewallPoliciesSortByTimecreated ListWebAppFirewallPoliciesSortByEnum = "timeCreated"
	ListWebAppFirewallPoliciesSortByDisplayname ListWebAppFirewallPoliciesSortByEnum = "displayName"
)

var mappingListWebAppFirewallPoliciesSortByEnum = map[string]ListWebAppFirewallPoliciesSortByEnum{
	"timeCreated": ListWebAppFirewallPoliciesSortByTimecreated,
	"displayName": ListWebAppFirewallPoliciesSortByDisplayname,
}

var mappingListWebAppFirewallPoliciesSortByEnumLowerCase = map[string]ListWebAppFirewallPoliciesSortByEnum{
	"timecreated": ListWebAppFirewallPoliciesSortByTimecreated,
	"displayname": ListWebAppFirewallPoliciesSortByDisplayname,
}

// GetListWebAppFirewallPoliciesSortByEnumValues Enumerates the set of values for ListWebAppFirewallPoliciesSortByEnum
func GetListWebAppFirewallPoliciesSortByEnumValues() []ListWebAppFirewallPoliciesSortByEnum {
	values := make([]ListWebAppFirewallPoliciesSortByEnum, 0)
	for _, v := range mappingListWebAppFirewallPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWebAppFirewallPoliciesSortByEnumStringValues Enumerates the set of values in String for ListWebAppFirewallPoliciesSortByEnum
func GetListWebAppFirewallPoliciesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListWebAppFirewallPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWebAppFirewallPoliciesSortByEnum(val string) (ListWebAppFirewallPoliciesSortByEnum, bool) {
	enum, ok := mappingListWebAppFirewallPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
