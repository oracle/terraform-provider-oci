// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waf

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListWebAppFirewallPoliciesRequest wrapper for the ListWebAppFirewallPolicies operation
//
// See also
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

var mappingListWebAppFirewallPoliciesSortOrder = map[string]ListWebAppFirewallPoliciesSortOrderEnum{
	"ASC":  ListWebAppFirewallPoliciesSortOrderAsc,
	"DESC": ListWebAppFirewallPoliciesSortOrderDesc,
}

// GetListWebAppFirewallPoliciesSortOrderEnumValues Enumerates the set of values for ListWebAppFirewallPoliciesSortOrderEnum
func GetListWebAppFirewallPoliciesSortOrderEnumValues() []ListWebAppFirewallPoliciesSortOrderEnum {
	values := make([]ListWebAppFirewallPoliciesSortOrderEnum, 0)
	for _, v := range mappingListWebAppFirewallPoliciesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListWebAppFirewallPoliciesSortByEnum Enum with underlying type: string
type ListWebAppFirewallPoliciesSortByEnum string

// Set of constants representing the allowable values for ListWebAppFirewallPoliciesSortByEnum
const (
	ListWebAppFirewallPoliciesSortByTimecreated ListWebAppFirewallPoliciesSortByEnum = "timeCreated"
	ListWebAppFirewallPoliciesSortByDisplayname ListWebAppFirewallPoliciesSortByEnum = "displayName"
)

var mappingListWebAppFirewallPoliciesSortBy = map[string]ListWebAppFirewallPoliciesSortByEnum{
	"timeCreated": ListWebAppFirewallPoliciesSortByTimecreated,
	"displayName": ListWebAppFirewallPoliciesSortByDisplayname,
}

// GetListWebAppFirewallPoliciesSortByEnumValues Enumerates the set of values for ListWebAppFirewallPoliciesSortByEnum
func GetListWebAppFirewallPoliciesSortByEnumValues() []ListWebAppFirewallPoliciesSortByEnum {
	values := make([]ListWebAppFirewallPoliciesSortByEnum, 0)
	for _, v := range mappingListWebAppFirewallPoliciesSortBy {
		values = append(values, v)
	}
	return values
}
