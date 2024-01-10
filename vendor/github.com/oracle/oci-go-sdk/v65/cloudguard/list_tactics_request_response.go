// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTacticsRequest wrapper for the ListTactics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListTactics.go.html to see an example of how to use ListTacticsRequest.
type ListTacticsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListTacticsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTacticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for displayName is ascending. If no value is specified displayName is default.
	SortBy ListTacticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTacticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTacticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTacticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTacticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTacticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTacticsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTacticsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTacticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTacticsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTacticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTacticsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTacticsResponse wrapper for the ListTactics operation
type ListTacticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TacticCollection instances
	TacticCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTacticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTacticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTacticsLifecycleStateEnum Enum with underlying type: string
type ListTacticsLifecycleStateEnum string

// Set of constants representing the allowable values for ListTacticsLifecycleStateEnum
const (
	ListTacticsLifecycleStateCreating ListTacticsLifecycleStateEnum = "CREATING"
	ListTacticsLifecycleStateUpdating ListTacticsLifecycleStateEnum = "UPDATING"
	ListTacticsLifecycleStateActive   ListTacticsLifecycleStateEnum = "ACTIVE"
	ListTacticsLifecycleStateInactive ListTacticsLifecycleStateEnum = "INACTIVE"
	ListTacticsLifecycleStateDeleting ListTacticsLifecycleStateEnum = "DELETING"
	ListTacticsLifecycleStateDeleted  ListTacticsLifecycleStateEnum = "DELETED"
	ListTacticsLifecycleStateFailed   ListTacticsLifecycleStateEnum = "FAILED"
)

var mappingListTacticsLifecycleStateEnum = map[string]ListTacticsLifecycleStateEnum{
	"CREATING": ListTacticsLifecycleStateCreating,
	"UPDATING": ListTacticsLifecycleStateUpdating,
	"ACTIVE":   ListTacticsLifecycleStateActive,
	"INACTIVE": ListTacticsLifecycleStateInactive,
	"DELETING": ListTacticsLifecycleStateDeleting,
	"DELETED":  ListTacticsLifecycleStateDeleted,
	"FAILED":   ListTacticsLifecycleStateFailed,
}

var mappingListTacticsLifecycleStateEnumLowerCase = map[string]ListTacticsLifecycleStateEnum{
	"creating": ListTacticsLifecycleStateCreating,
	"updating": ListTacticsLifecycleStateUpdating,
	"active":   ListTacticsLifecycleStateActive,
	"inactive": ListTacticsLifecycleStateInactive,
	"deleting": ListTacticsLifecycleStateDeleting,
	"deleted":  ListTacticsLifecycleStateDeleted,
	"failed":   ListTacticsLifecycleStateFailed,
}

// GetListTacticsLifecycleStateEnumValues Enumerates the set of values for ListTacticsLifecycleStateEnum
func GetListTacticsLifecycleStateEnumValues() []ListTacticsLifecycleStateEnum {
	values := make([]ListTacticsLifecycleStateEnum, 0)
	for _, v := range mappingListTacticsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTacticsLifecycleStateEnumStringValues Enumerates the set of values in String for ListTacticsLifecycleStateEnum
func GetListTacticsLifecycleStateEnumStringValues() []string {
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

// GetMappingListTacticsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTacticsLifecycleStateEnum(val string) (ListTacticsLifecycleStateEnum, bool) {
	enum, ok := mappingListTacticsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTacticsSortOrderEnum Enum with underlying type: string
type ListTacticsSortOrderEnum string

// Set of constants representing the allowable values for ListTacticsSortOrderEnum
const (
	ListTacticsSortOrderAsc  ListTacticsSortOrderEnum = "ASC"
	ListTacticsSortOrderDesc ListTacticsSortOrderEnum = "DESC"
)

var mappingListTacticsSortOrderEnum = map[string]ListTacticsSortOrderEnum{
	"ASC":  ListTacticsSortOrderAsc,
	"DESC": ListTacticsSortOrderDesc,
}

var mappingListTacticsSortOrderEnumLowerCase = map[string]ListTacticsSortOrderEnum{
	"asc":  ListTacticsSortOrderAsc,
	"desc": ListTacticsSortOrderDesc,
}

// GetListTacticsSortOrderEnumValues Enumerates the set of values for ListTacticsSortOrderEnum
func GetListTacticsSortOrderEnumValues() []ListTacticsSortOrderEnum {
	values := make([]ListTacticsSortOrderEnum, 0)
	for _, v := range mappingListTacticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTacticsSortOrderEnumStringValues Enumerates the set of values in String for ListTacticsSortOrderEnum
func GetListTacticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTacticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTacticsSortOrderEnum(val string) (ListTacticsSortOrderEnum, bool) {
	enum, ok := mappingListTacticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTacticsSortByEnum Enum with underlying type: string
type ListTacticsSortByEnum string

// Set of constants representing the allowable values for ListTacticsSortByEnum
const (
	ListTacticsSortByDisplayname ListTacticsSortByEnum = "displayName"
)

var mappingListTacticsSortByEnum = map[string]ListTacticsSortByEnum{
	"displayName": ListTacticsSortByDisplayname,
}

var mappingListTacticsSortByEnumLowerCase = map[string]ListTacticsSortByEnum{
	"displayname": ListTacticsSortByDisplayname,
}

// GetListTacticsSortByEnumValues Enumerates the set of values for ListTacticsSortByEnum
func GetListTacticsSortByEnumValues() []ListTacticsSortByEnum {
	values := make([]ListTacticsSortByEnum, 0)
	for _, v := range mappingListTacticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTacticsSortByEnumStringValues Enumerates the set of values in String for ListTacticsSortByEnum
func GetListTacticsSortByEnumStringValues() []string {
	return []string{
		"displayName",
	}
}

// GetMappingListTacticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTacticsSortByEnum(val string) (ListTacticsSortByEnum, bool) {
	enum, ok := mappingListTacticsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
