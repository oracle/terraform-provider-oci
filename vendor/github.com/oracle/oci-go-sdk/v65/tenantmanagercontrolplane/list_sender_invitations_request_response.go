// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSenderInvitationsRequest wrapper for the ListSenderInvitations operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListSenderInvitations.go.html to see an example of how to use ListSenderInvitationsRequest.
type ListSenderInvitationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The tenancy that the invitation is addressed to.
	RecipientTenancyId *string `mandatory:"false" contributesTo:"query" name:"recipientTenancyId"`

	// The lifecycle state of the resource.
	LifecycleState ListSenderInvitationsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The status of the sender invitation.
	Status ListSenderInvitationsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order can be provided.
	// * The default order for timeCreated is descending.
	// * The default order for displayName is ascending.
	// * If no value is specified, timeCreated is the default.
	SortBy ListSenderInvitationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListSenderInvitationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSenderInvitationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSenderInvitationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSenderInvitationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSenderInvitationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSenderInvitationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSenderInvitationsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSenderInvitationsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSenderInvitationsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListSenderInvitationsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSenderInvitationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSenderInvitationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSenderInvitationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSenderInvitationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSenderInvitationsResponse wrapper for the ListSenderInvitations operation
type ListSenderInvitationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SenderInvitationCollection instances
	SenderInvitationCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSenderInvitationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSenderInvitationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSenderInvitationsLifecycleStateEnum Enum with underlying type: string
type ListSenderInvitationsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSenderInvitationsLifecycleStateEnum
const (
	ListSenderInvitationsLifecycleStateCreating   ListSenderInvitationsLifecycleStateEnum = "CREATING"
	ListSenderInvitationsLifecycleStateActive     ListSenderInvitationsLifecycleStateEnum = "ACTIVE"
	ListSenderInvitationsLifecycleStateInactive   ListSenderInvitationsLifecycleStateEnum = "INACTIVE"
	ListSenderInvitationsLifecycleStateUpdating   ListSenderInvitationsLifecycleStateEnum = "UPDATING"
	ListSenderInvitationsLifecycleStateFailed     ListSenderInvitationsLifecycleStateEnum = "FAILED"
	ListSenderInvitationsLifecycleStateTerminated ListSenderInvitationsLifecycleStateEnum = "TERMINATED"
)

var mappingListSenderInvitationsLifecycleStateEnum = map[string]ListSenderInvitationsLifecycleStateEnum{
	"CREATING":   ListSenderInvitationsLifecycleStateCreating,
	"ACTIVE":     ListSenderInvitationsLifecycleStateActive,
	"INACTIVE":   ListSenderInvitationsLifecycleStateInactive,
	"UPDATING":   ListSenderInvitationsLifecycleStateUpdating,
	"FAILED":     ListSenderInvitationsLifecycleStateFailed,
	"TERMINATED": ListSenderInvitationsLifecycleStateTerminated,
}

var mappingListSenderInvitationsLifecycleStateEnumLowerCase = map[string]ListSenderInvitationsLifecycleStateEnum{
	"creating":   ListSenderInvitationsLifecycleStateCreating,
	"active":     ListSenderInvitationsLifecycleStateActive,
	"inactive":   ListSenderInvitationsLifecycleStateInactive,
	"updating":   ListSenderInvitationsLifecycleStateUpdating,
	"failed":     ListSenderInvitationsLifecycleStateFailed,
	"terminated": ListSenderInvitationsLifecycleStateTerminated,
}

// GetListSenderInvitationsLifecycleStateEnumValues Enumerates the set of values for ListSenderInvitationsLifecycleStateEnum
func GetListSenderInvitationsLifecycleStateEnumValues() []ListSenderInvitationsLifecycleStateEnum {
	values := make([]ListSenderInvitationsLifecycleStateEnum, 0)
	for _, v := range mappingListSenderInvitationsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSenderInvitationsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSenderInvitationsLifecycleStateEnum
func GetListSenderInvitationsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"FAILED",
		"TERMINATED",
	}
}

// GetMappingListSenderInvitationsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSenderInvitationsLifecycleStateEnum(val string) (ListSenderInvitationsLifecycleStateEnum, bool) {
	enum, ok := mappingListSenderInvitationsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSenderInvitationsStatusEnum Enum with underlying type: string
type ListSenderInvitationsStatusEnum string

// Set of constants representing the allowable values for ListSenderInvitationsStatusEnum
const (
	ListSenderInvitationsStatusPending  ListSenderInvitationsStatusEnum = "PENDING"
	ListSenderInvitationsStatusCanceled ListSenderInvitationsStatusEnum = "CANCELED"
	ListSenderInvitationsStatusAccepted ListSenderInvitationsStatusEnum = "ACCEPTED"
	ListSenderInvitationsStatusExpired  ListSenderInvitationsStatusEnum = "EXPIRED"
	ListSenderInvitationsStatusFailed   ListSenderInvitationsStatusEnum = "FAILED"
)

var mappingListSenderInvitationsStatusEnum = map[string]ListSenderInvitationsStatusEnum{
	"PENDING":  ListSenderInvitationsStatusPending,
	"CANCELED": ListSenderInvitationsStatusCanceled,
	"ACCEPTED": ListSenderInvitationsStatusAccepted,
	"EXPIRED":  ListSenderInvitationsStatusExpired,
	"FAILED":   ListSenderInvitationsStatusFailed,
}

var mappingListSenderInvitationsStatusEnumLowerCase = map[string]ListSenderInvitationsStatusEnum{
	"pending":  ListSenderInvitationsStatusPending,
	"canceled": ListSenderInvitationsStatusCanceled,
	"accepted": ListSenderInvitationsStatusAccepted,
	"expired":  ListSenderInvitationsStatusExpired,
	"failed":   ListSenderInvitationsStatusFailed,
}

// GetListSenderInvitationsStatusEnumValues Enumerates the set of values for ListSenderInvitationsStatusEnum
func GetListSenderInvitationsStatusEnumValues() []ListSenderInvitationsStatusEnum {
	values := make([]ListSenderInvitationsStatusEnum, 0)
	for _, v := range mappingListSenderInvitationsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListSenderInvitationsStatusEnumStringValues Enumerates the set of values in String for ListSenderInvitationsStatusEnum
func GetListSenderInvitationsStatusEnumStringValues() []string {
	return []string{
		"PENDING",
		"CANCELED",
		"ACCEPTED",
		"EXPIRED",
		"FAILED",
	}
}

// GetMappingListSenderInvitationsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSenderInvitationsStatusEnum(val string) (ListSenderInvitationsStatusEnum, bool) {
	enum, ok := mappingListSenderInvitationsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSenderInvitationsSortByEnum Enum with underlying type: string
type ListSenderInvitationsSortByEnum string

// Set of constants representing the allowable values for ListSenderInvitationsSortByEnum
const (
	ListSenderInvitationsSortByTimecreated ListSenderInvitationsSortByEnum = "timeCreated"
	ListSenderInvitationsSortByDisplayname ListSenderInvitationsSortByEnum = "displayName"
)

var mappingListSenderInvitationsSortByEnum = map[string]ListSenderInvitationsSortByEnum{
	"timeCreated": ListSenderInvitationsSortByTimecreated,
	"displayName": ListSenderInvitationsSortByDisplayname,
}

var mappingListSenderInvitationsSortByEnumLowerCase = map[string]ListSenderInvitationsSortByEnum{
	"timecreated": ListSenderInvitationsSortByTimecreated,
	"displayname": ListSenderInvitationsSortByDisplayname,
}

// GetListSenderInvitationsSortByEnumValues Enumerates the set of values for ListSenderInvitationsSortByEnum
func GetListSenderInvitationsSortByEnumValues() []ListSenderInvitationsSortByEnum {
	values := make([]ListSenderInvitationsSortByEnum, 0)
	for _, v := range mappingListSenderInvitationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSenderInvitationsSortByEnumStringValues Enumerates the set of values in String for ListSenderInvitationsSortByEnum
func GetListSenderInvitationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSenderInvitationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSenderInvitationsSortByEnum(val string) (ListSenderInvitationsSortByEnum, bool) {
	enum, ok := mappingListSenderInvitationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSenderInvitationsSortOrderEnum Enum with underlying type: string
type ListSenderInvitationsSortOrderEnum string

// Set of constants representing the allowable values for ListSenderInvitationsSortOrderEnum
const (
	ListSenderInvitationsSortOrderAsc  ListSenderInvitationsSortOrderEnum = "ASC"
	ListSenderInvitationsSortOrderDesc ListSenderInvitationsSortOrderEnum = "DESC"
)

var mappingListSenderInvitationsSortOrderEnum = map[string]ListSenderInvitationsSortOrderEnum{
	"ASC":  ListSenderInvitationsSortOrderAsc,
	"DESC": ListSenderInvitationsSortOrderDesc,
}

var mappingListSenderInvitationsSortOrderEnumLowerCase = map[string]ListSenderInvitationsSortOrderEnum{
	"asc":  ListSenderInvitationsSortOrderAsc,
	"desc": ListSenderInvitationsSortOrderDesc,
}

// GetListSenderInvitationsSortOrderEnumValues Enumerates the set of values for ListSenderInvitationsSortOrderEnum
func GetListSenderInvitationsSortOrderEnumValues() []ListSenderInvitationsSortOrderEnum {
	values := make([]ListSenderInvitationsSortOrderEnum, 0)
	for _, v := range mappingListSenderInvitationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSenderInvitationsSortOrderEnumStringValues Enumerates the set of values in String for ListSenderInvitationsSortOrderEnum
func GetListSenderInvitationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSenderInvitationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSenderInvitationsSortOrderEnum(val string) (ListSenderInvitationsSortOrderEnum, bool) {
	enum, ok := mappingListSenderInvitationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
