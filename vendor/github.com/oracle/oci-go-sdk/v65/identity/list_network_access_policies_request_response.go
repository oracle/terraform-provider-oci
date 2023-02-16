// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListNetworkAccessPoliciesRequest wrapper for the ListNetworkAccessPolicies operation
type ListNetworkAccessPoliciesRequest struct {

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for NAME is ascending. The NAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListNetworkAccessPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListNetworkAccessPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState NetworkAccessPolicyLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNetworkAccessPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNetworkAccessPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNetworkAccessPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNetworkAccessPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNetworkAccessPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNetworkAccessPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNetworkAccessPoliciesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNetworkAccessPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNetworkAccessPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNetworkAccessPolicyLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetNetworkAccessPolicyLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNetworkAccessPoliciesResponse wrapper for the ListNetworkAccessPolicies operation
type ListNetworkAccessPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []NetworkAccessPolicySummary instances
	Items []NetworkAccessPolicySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListNetworkAccessPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNetworkAccessPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNetworkAccessPoliciesSortByEnum Enum with underlying type: string
type ListNetworkAccessPoliciesSortByEnum string

// Set of constants representing the allowable values for ListNetworkAccessPoliciesSortByEnum
const (
	ListNetworkAccessPoliciesSortByTimecreated ListNetworkAccessPoliciesSortByEnum = "TIMECREATED"
	ListNetworkAccessPoliciesSortByName        ListNetworkAccessPoliciesSortByEnum = "NAME"
)

var mappingListNetworkAccessPoliciesSortByEnum = map[string]ListNetworkAccessPoliciesSortByEnum{
	"TIMECREATED": ListNetworkAccessPoliciesSortByTimecreated,
	"NAME":        ListNetworkAccessPoliciesSortByName,
}

var mappingListNetworkAccessPoliciesSortByEnumLowerCase = map[string]ListNetworkAccessPoliciesSortByEnum{
	"timecreated": ListNetworkAccessPoliciesSortByTimecreated,
	"name":        ListNetworkAccessPoliciesSortByName,
}

// GetListNetworkAccessPoliciesSortByEnumValues Enumerates the set of values for ListNetworkAccessPoliciesSortByEnum
func GetListNetworkAccessPoliciesSortByEnumValues() []ListNetworkAccessPoliciesSortByEnum {
	values := make([]ListNetworkAccessPoliciesSortByEnum, 0)
	for _, v := range mappingListNetworkAccessPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkAccessPoliciesSortByEnumStringValues Enumerates the set of values in String for ListNetworkAccessPoliciesSortByEnum
func GetListNetworkAccessPoliciesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListNetworkAccessPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkAccessPoliciesSortByEnum(val string) (ListNetworkAccessPoliciesSortByEnum, bool) {
	enum, ok := mappingListNetworkAccessPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNetworkAccessPoliciesSortOrderEnum Enum with underlying type: string
type ListNetworkAccessPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListNetworkAccessPoliciesSortOrderEnum
const (
	ListNetworkAccessPoliciesSortOrderAsc  ListNetworkAccessPoliciesSortOrderEnum = "ASC"
	ListNetworkAccessPoliciesSortOrderDesc ListNetworkAccessPoliciesSortOrderEnum = "DESC"
)

var mappingListNetworkAccessPoliciesSortOrderEnum = map[string]ListNetworkAccessPoliciesSortOrderEnum{
	"ASC":  ListNetworkAccessPoliciesSortOrderAsc,
	"DESC": ListNetworkAccessPoliciesSortOrderDesc,
}

var mappingListNetworkAccessPoliciesSortOrderEnumLowerCase = map[string]ListNetworkAccessPoliciesSortOrderEnum{
	"asc":  ListNetworkAccessPoliciesSortOrderAsc,
	"desc": ListNetworkAccessPoliciesSortOrderDesc,
}

// GetListNetworkAccessPoliciesSortOrderEnumValues Enumerates the set of values for ListNetworkAccessPoliciesSortOrderEnum
func GetListNetworkAccessPoliciesSortOrderEnumValues() []ListNetworkAccessPoliciesSortOrderEnum {
	values := make([]ListNetworkAccessPoliciesSortOrderEnum, 0)
	for _, v := range mappingListNetworkAccessPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkAccessPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListNetworkAccessPoliciesSortOrderEnum
func GetListNetworkAccessPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNetworkAccessPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkAccessPoliciesSortOrderEnum(val string) (ListNetworkAccessPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListNetworkAccessPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
