// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListDatabaseRegistrationsRequest wrapper for the ListDatabaseRegistrations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDatabaseRegistrations.go.html to see an example of how to use ListDatabaseRegistrationsRequest.
type ListDatabaseRegistrationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources that match the 'lifecycleState' given.
	LifecycleState ListDatabaseRegistrationsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatabaseRegistrationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is descending.  Default order for 'displayName' is ascending. If no value is specified timeCreated is the default.
	SortBy ListDatabaseRegistrationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseRegistrationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseRegistrationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseRegistrationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseRegistrationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseRegistrationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseRegistrationsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDatabaseRegistrationsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseRegistrationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseRegistrationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseRegistrationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseRegistrationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseRegistrationsResponse wrapper for the ListDatabaseRegistrations operation
type ListDatabaseRegistrationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseRegistrationCollection instances
	DatabaseRegistrationCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseRegistrationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseRegistrationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseRegistrationsLifecycleStateEnum Enum with underlying type: string
type ListDatabaseRegistrationsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDatabaseRegistrationsLifecycleStateEnum
const (
	ListDatabaseRegistrationsLifecycleStateCreating       ListDatabaseRegistrationsLifecycleStateEnum = "CREATING"
	ListDatabaseRegistrationsLifecycleStateUpdating       ListDatabaseRegistrationsLifecycleStateEnum = "UPDATING"
	ListDatabaseRegistrationsLifecycleStateActive         ListDatabaseRegistrationsLifecycleStateEnum = "ACTIVE"
	ListDatabaseRegistrationsLifecycleStateInactive       ListDatabaseRegistrationsLifecycleStateEnum = "INACTIVE"
	ListDatabaseRegistrationsLifecycleStateDeleting       ListDatabaseRegistrationsLifecycleStateEnum = "DELETING"
	ListDatabaseRegistrationsLifecycleStateDeleted        ListDatabaseRegistrationsLifecycleStateEnum = "DELETED"
	ListDatabaseRegistrationsLifecycleStateFailed         ListDatabaseRegistrationsLifecycleStateEnum = "FAILED"
	ListDatabaseRegistrationsLifecycleStateNeedsAttention ListDatabaseRegistrationsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListDatabaseRegistrationsLifecycleStateInProgress     ListDatabaseRegistrationsLifecycleStateEnum = "IN_PROGRESS"
	ListDatabaseRegistrationsLifecycleStateCanceling      ListDatabaseRegistrationsLifecycleStateEnum = "CANCELING"
	ListDatabaseRegistrationsLifecycleStateCanceled       ListDatabaseRegistrationsLifecycleStateEnum = "CANCELED"
	ListDatabaseRegistrationsLifecycleStateSucceeded      ListDatabaseRegistrationsLifecycleStateEnum = "SUCCEEDED"
)

var mappingListDatabaseRegistrationsLifecycleStateEnum = map[string]ListDatabaseRegistrationsLifecycleStateEnum{
	"CREATING":        ListDatabaseRegistrationsLifecycleStateCreating,
	"UPDATING":        ListDatabaseRegistrationsLifecycleStateUpdating,
	"ACTIVE":          ListDatabaseRegistrationsLifecycleStateActive,
	"INACTIVE":        ListDatabaseRegistrationsLifecycleStateInactive,
	"DELETING":        ListDatabaseRegistrationsLifecycleStateDeleting,
	"DELETED":         ListDatabaseRegistrationsLifecycleStateDeleted,
	"FAILED":          ListDatabaseRegistrationsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListDatabaseRegistrationsLifecycleStateNeedsAttention,
	"IN_PROGRESS":     ListDatabaseRegistrationsLifecycleStateInProgress,
	"CANCELING":       ListDatabaseRegistrationsLifecycleStateCanceling,
	"CANCELED":        ListDatabaseRegistrationsLifecycleStateCanceled,
	"SUCCEEDED":       ListDatabaseRegistrationsLifecycleStateSucceeded,
}

// GetListDatabaseRegistrationsLifecycleStateEnumValues Enumerates the set of values for ListDatabaseRegistrationsLifecycleStateEnum
func GetListDatabaseRegistrationsLifecycleStateEnumValues() []ListDatabaseRegistrationsLifecycleStateEnum {
	values := make([]ListDatabaseRegistrationsLifecycleStateEnum, 0)
	for _, v := range mappingListDatabaseRegistrationsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseRegistrationsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDatabaseRegistrationsLifecycleStateEnum
func GetListDatabaseRegistrationsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
		"IN_PROGRESS",
		"CANCELING",
		"CANCELED",
		"SUCCEEDED",
	}
}

// GetMappingListDatabaseRegistrationsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseRegistrationsLifecycleStateEnum(val string) (ListDatabaseRegistrationsLifecycleStateEnum, bool) {
	mappingListDatabaseRegistrationsLifecycleStateEnumIgnoreCase := make(map[string]ListDatabaseRegistrationsLifecycleStateEnum)
	for k, v := range mappingListDatabaseRegistrationsLifecycleStateEnum {
		mappingListDatabaseRegistrationsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDatabaseRegistrationsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseRegistrationsSortOrderEnum Enum with underlying type: string
type ListDatabaseRegistrationsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseRegistrationsSortOrderEnum
const (
	ListDatabaseRegistrationsSortOrderAsc  ListDatabaseRegistrationsSortOrderEnum = "ASC"
	ListDatabaseRegistrationsSortOrderDesc ListDatabaseRegistrationsSortOrderEnum = "DESC"
)

var mappingListDatabaseRegistrationsSortOrderEnum = map[string]ListDatabaseRegistrationsSortOrderEnum{
	"ASC":  ListDatabaseRegistrationsSortOrderAsc,
	"DESC": ListDatabaseRegistrationsSortOrderDesc,
}

// GetListDatabaseRegistrationsSortOrderEnumValues Enumerates the set of values for ListDatabaseRegistrationsSortOrderEnum
func GetListDatabaseRegistrationsSortOrderEnumValues() []ListDatabaseRegistrationsSortOrderEnum {
	values := make([]ListDatabaseRegistrationsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseRegistrationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseRegistrationsSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseRegistrationsSortOrderEnum
func GetListDatabaseRegistrationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseRegistrationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseRegistrationsSortOrderEnum(val string) (ListDatabaseRegistrationsSortOrderEnum, bool) {
	mappingListDatabaseRegistrationsSortOrderEnumIgnoreCase := make(map[string]ListDatabaseRegistrationsSortOrderEnum)
	for k, v := range mappingListDatabaseRegistrationsSortOrderEnum {
		mappingListDatabaseRegistrationsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDatabaseRegistrationsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseRegistrationsSortByEnum Enum with underlying type: string
type ListDatabaseRegistrationsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseRegistrationsSortByEnum
const (
	ListDatabaseRegistrationsSortByTimecreated ListDatabaseRegistrationsSortByEnum = "timeCreated"
	ListDatabaseRegistrationsSortByDisplayname ListDatabaseRegistrationsSortByEnum = "displayName"
)

var mappingListDatabaseRegistrationsSortByEnum = map[string]ListDatabaseRegistrationsSortByEnum{
	"timeCreated": ListDatabaseRegistrationsSortByTimecreated,
	"displayName": ListDatabaseRegistrationsSortByDisplayname,
}

// GetListDatabaseRegistrationsSortByEnumValues Enumerates the set of values for ListDatabaseRegistrationsSortByEnum
func GetListDatabaseRegistrationsSortByEnumValues() []ListDatabaseRegistrationsSortByEnum {
	values := make([]ListDatabaseRegistrationsSortByEnum, 0)
	for _, v := range mappingListDatabaseRegistrationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseRegistrationsSortByEnumStringValues Enumerates the set of values in String for ListDatabaseRegistrationsSortByEnum
func GetListDatabaseRegistrationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatabaseRegistrationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseRegistrationsSortByEnum(val string) (ListDatabaseRegistrationsSortByEnum, bool) {
	mappingListDatabaseRegistrationsSortByEnumIgnoreCase := make(map[string]ListDatabaseRegistrationsSortByEnum)
	for k, v := range mappingListDatabaseRegistrationsSortByEnum {
		mappingListDatabaseRegistrationsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDatabaseRegistrationsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
