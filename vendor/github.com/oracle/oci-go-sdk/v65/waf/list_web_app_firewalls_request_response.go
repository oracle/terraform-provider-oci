// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waf

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWebAppFirewallsRequest wrapper for the ListWebAppFirewalls operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ListWebAppFirewalls.go.html to see an example of how to use ListWebAppFirewallsRequest.
type ListWebAppFirewallsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the WebAppFirewall with the given OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only the WebAppFirewall with the given OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of related WebAppFirewallPolicy.
	WebAppFirewallPolicyId *string `mandatory:"false" contributesTo:"query" name:"webAppFirewallPolicyId"`

	// A filter to return only resources that match the given lifecycleState.
	LifecycleState []WebAppFirewallLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListWebAppFirewallsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending.
	// Default order for displayName is ascending.
	// If no value is specified timeCreated is default.
	SortBy ListWebAppFirewallsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWebAppFirewallsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWebAppFirewallsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWebAppFirewallsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWebAppFirewallsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWebAppFirewallsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingWebAppFirewallLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetWebAppFirewallLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListWebAppFirewallsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWebAppFirewallsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWebAppFirewallsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWebAppFirewallsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWebAppFirewallsResponse wrapper for the ListWebAppFirewalls operation
type ListWebAppFirewallsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WebAppFirewallCollection instances
	WebAppFirewallCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWebAppFirewallsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWebAppFirewallsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWebAppFirewallsSortOrderEnum Enum with underlying type: string
type ListWebAppFirewallsSortOrderEnum string

// Set of constants representing the allowable values for ListWebAppFirewallsSortOrderEnum
const (
	ListWebAppFirewallsSortOrderAsc  ListWebAppFirewallsSortOrderEnum = "ASC"
	ListWebAppFirewallsSortOrderDesc ListWebAppFirewallsSortOrderEnum = "DESC"
)

var mappingListWebAppFirewallsSortOrderEnum = map[string]ListWebAppFirewallsSortOrderEnum{
	"ASC":  ListWebAppFirewallsSortOrderAsc,
	"DESC": ListWebAppFirewallsSortOrderDesc,
}

var mappingListWebAppFirewallsSortOrderEnumLowerCase = map[string]ListWebAppFirewallsSortOrderEnum{
	"asc":  ListWebAppFirewallsSortOrderAsc,
	"desc": ListWebAppFirewallsSortOrderDesc,
}

// GetListWebAppFirewallsSortOrderEnumValues Enumerates the set of values for ListWebAppFirewallsSortOrderEnum
func GetListWebAppFirewallsSortOrderEnumValues() []ListWebAppFirewallsSortOrderEnum {
	values := make([]ListWebAppFirewallsSortOrderEnum, 0)
	for _, v := range mappingListWebAppFirewallsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWebAppFirewallsSortOrderEnumStringValues Enumerates the set of values in String for ListWebAppFirewallsSortOrderEnum
func GetListWebAppFirewallsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWebAppFirewallsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWebAppFirewallsSortOrderEnum(val string) (ListWebAppFirewallsSortOrderEnum, bool) {
	enum, ok := mappingListWebAppFirewallsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWebAppFirewallsSortByEnum Enum with underlying type: string
type ListWebAppFirewallsSortByEnum string

// Set of constants representing the allowable values for ListWebAppFirewallsSortByEnum
const (
	ListWebAppFirewallsSortByTimecreated ListWebAppFirewallsSortByEnum = "timeCreated"
	ListWebAppFirewallsSortByDisplayname ListWebAppFirewallsSortByEnum = "displayName"
)

var mappingListWebAppFirewallsSortByEnum = map[string]ListWebAppFirewallsSortByEnum{
	"timeCreated": ListWebAppFirewallsSortByTimecreated,
	"displayName": ListWebAppFirewallsSortByDisplayname,
}

var mappingListWebAppFirewallsSortByEnumLowerCase = map[string]ListWebAppFirewallsSortByEnum{
	"timecreated": ListWebAppFirewallsSortByTimecreated,
	"displayname": ListWebAppFirewallsSortByDisplayname,
}

// GetListWebAppFirewallsSortByEnumValues Enumerates the set of values for ListWebAppFirewallsSortByEnum
func GetListWebAppFirewallsSortByEnumValues() []ListWebAppFirewallsSortByEnum {
	values := make([]ListWebAppFirewallsSortByEnum, 0)
	for _, v := range mappingListWebAppFirewallsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWebAppFirewallsSortByEnumStringValues Enumerates the set of values in String for ListWebAppFirewallsSortByEnum
func GetListWebAppFirewallsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListWebAppFirewallsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWebAppFirewallsSortByEnum(val string) (ListWebAppFirewallsSortByEnum, bool) {
	enum, ok := mappingListWebAppFirewallsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
