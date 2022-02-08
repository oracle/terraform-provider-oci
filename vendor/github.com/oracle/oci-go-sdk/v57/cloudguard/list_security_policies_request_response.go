// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"net/http"
	"strings"
)

// ListSecurityPoliciesRequest wrapper for the ListSecurityPolicies operation
type ListSecurityPoliciesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListSecurityPoliciesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique SecurityPolicy identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListSecurityPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListSecurityPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecurityPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecurityPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSecurityPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecurityPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSecurityPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingListSecurityPoliciesLifecycleStateEnum[string(request.LifecycleState)]; !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSecurityPoliciesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := mappingListSecurityPoliciesSortOrderEnum[string(request.SortOrder)]; !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSecurityPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := mappingListSecurityPoliciesSortByEnum[string(request.SortBy)]; !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSecurityPoliciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSecurityPoliciesResponse wrapper for the ListSecurityPolicies operation
type ListSecurityPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SecurityPolicyCollection instances
	SecurityPolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSecurityPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecurityPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecurityPoliciesLifecycleStateEnum Enum with underlying type: string
type ListSecurityPoliciesLifecycleStateEnum string

// Set of constants representing the allowable values for ListSecurityPoliciesLifecycleStateEnum
const (
	ListSecurityPoliciesLifecycleStateCreating ListSecurityPoliciesLifecycleStateEnum = "CREATING"
	ListSecurityPoliciesLifecycleStateUpdating ListSecurityPoliciesLifecycleStateEnum = "UPDATING"
	ListSecurityPoliciesLifecycleStateActive   ListSecurityPoliciesLifecycleStateEnum = "ACTIVE"
	ListSecurityPoliciesLifecycleStateInactive ListSecurityPoliciesLifecycleStateEnum = "INACTIVE"
	ListSecurityPoliciesLifecycleStateDeleting ListSecurityPoliciesLifecycleStateEnum = "DELETING"
	ListSecurityPoliciesLifecycleStateDeleted  ListSecurityPoliciesLifecycleStateEnum = "DELETED"
	ListSecurityPoliciesLifecycleStateFailed   ListSecurityPoliciesLifecycleStateEnum = "FAILED"
)

var mappingListSecurityPoliciesLifecycleStateEnum = map[string]ListSecurityPoliciesLifecycleStateEnum{
	"CREATING": ListSecurityPoliciesLifecycleStateCreating,
	"UPDATING": ListSecurityPoliciesLifecycleStateUpdating,
	"ACTIVE":   ListSecurityPoliciesLifecycleStateActive,
	"INACTIVE": ListSecurityPoliciesLifecycleStateInactive,
	"DELETING": ListSecurityPoliciesLifecycleStateDeleting,
	"DELETED":  ListSecurityPoliciesLifecycleStateDeleted,
	"FAILED":   ListSecurityPoliciesLifecycleStateFailed,
}

// GetListSecurityPoliciesLifecycleStateEnumValues Enumerates the set of values for ListSecurityPoliciesLifecycleStateEnum
func GetListSecurityPoliciesLifecycleStateEnumValues() []ListSecurityPoliciesLifecycleStateEnum {
	values := make([]ListSecurityPoliciesLifecycleStateEnum, 0)
	for _, v := range mappingListSecurityPoliciesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPoliciesLifecycleStateEnumStringValues Enumerates the set of values in String for ListSecurityPoliciesLifecycleStateEnum
func GetListSecurityPoliciesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// ListSecurityPoliciesSortOrderEnum Enum with underlying type: string
type ListSecurityPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListSecurityPoliciesSortOrderEnum
const (
	ListSecurityPoliciesSortOrderAsc  ListSecurityPoliciesSortOrderEnum = "ASC"
	ListSecurityPoliciesSortOrderDesc ListSecurityPoliciesSortOrderEnum = "DESC"
)

var mappingListSecurityPoliciesSortOrderEnum = map[string]ListSecurityPoliciesSortOrderEnum{
	"ASC":  ListSecurityPoliciesSortOrderAsc,
	"DESC": ListSecurityPoliciesSortOrderDesc,
}

// GetListSecurityPoliciesSortOrderEnumValues Enumerates the set of values for ListSecurityPoliciesSortOrderEnum
func GetListSecurityPoliciesSortOrderEnumValues() []ListSecurityPoliciesSortOrderEnum {
	values := make([]ListSecurityPoliciesSortOrderEnum, 0)
	for _, v := range mappingListSecurityPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListSecurityPoliciesSortOrderEnum
func GetListSecurityPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// ListSecurityPoliciesSortByEnum Enum with underlying type: string
type ListSecurityPoliciesSortByEnum string

// Set of constants representing the allowable values for ListSecurityPoliciesSortByEnum
const (
	ListSecurityPoliciesSortByTimecreated ListSecurityPoliciesSortByEnum = "timeCreated"
	ListSecurityPoliciesSortByDisplayname ListSecurityPoliciesSortByEnum = "displayName"
)

var mappingListSecurityPoliciesSortByEnum = map[string]ListSecurityPoliciesSortByEnum{
	"timeCreated": ListSecurityPoliciesSortByTimecreated,
	"displayName": ListSecurityPoliciesSortByDisplayname,
}

// GetListSecurityPoliciesSortByEnumValues Enumerates the set of values for ListSecurityPoliciesSortByEnum
func GetListSecurityPoliciesSortByEnumValues() []ListSecurityPoliciesSortByEnum {
	values := make([]ListSecurityPoliciesSortByEnum, 0)
	for _, v := range mappingListSecurityPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPoliciesSortByEnumStringValues Enumerates the set of values in String for ListSecurityPoliciesSortByEnum
func GetListSecurityPoliciesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}
