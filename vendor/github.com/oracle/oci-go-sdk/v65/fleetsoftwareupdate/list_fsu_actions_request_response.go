// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetsoftwareupdate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFsuActionsRequest wrapper for the ListFsuActions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuActions.go.html to see an example of how to use ListFsuActionsRequest.
type ListFsuActionsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose fsuCycleId matches the given fleetSoftwareUpdateCycleId.
	FsuCycleId *string `mandatory:"false" contributesTo:"query" name:"fsuCycleId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState ListFsuActionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources whose type matches the given type.
	Type ListFsuActionsTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFsuActionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListFsuActionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFsuActionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFsuActionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFsuActionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFsuActionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFsuActionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFsuActionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListFsuActionsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuActionsTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListFsuActionsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuActionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFsuActionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuActionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFsuActionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFsuActionsResponse wrapper for the ListFsuActions operation
type ListFsuActionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FsuActionSummaryCollection instances
	FsuActionSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFsuActionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFsuActionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFsuActionsLifecycleStateEnum Enum with underlying type: string
type ListFsuActionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListFsuActionsLifecycleStateEnum
const (
	ListFsuActionsLifecycleStateAccepted       ListFsuActionsLifecycleStateEnum = "ACCEPTED"
	ListFsuActionsLifecycleStateInProgress     ListFsuActionsLifecycleStateEnum = "IN_PROGRESS"
	ListFsuActionsLifecycleStateWaiting        ListFsuActionsLifecycleStateEnum = "WAITING"
	ListFsuActionsLifecycleStateUpdating       ListFsuActionsLifecycleStateEnum = "UPDATING"
	ListFsuActionsLifecycleStateFailed         ListFsuActionsLifecycleStateEnum = "FAILED"
	ListFsuActionsLifecycleStateNeedsAttention ListFsuActionsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListFsuActionsLifecycleStateSucceeded      ListFsuActionsLifecycleStateEnum = "SUCCEEDED"
	ListFsuActionsLifecycleStateCanceling      ListFsuActionsLifecycleStateEnum = "CANCELING"
	ListFsuActionsLifecycleStateCanceled       ListFsuActionsLifecycleStateEnum = "CANCELED"
	ListFsuActionsLifecycleStateUnknown        ListFsuActionsLifecycleStateEnum = "UNKNOWN"
	ListFsuActionsLifecycleStateDeleting       ListFsuActionsLifecycleStateEnum = "DELETING"
	ListFsuActionsLifecycleStateDeleted        ListFsuActionsLifecycleStateEnum = "DELETED"
)

var mappingListFsuActionsLifecycleStateEnum = map[string]ListFsuActionsLifecycleStateEnum{
	"ACCEPTED":        ListFsuActionsLifecycleStateAccepted,
	"IN_PROGRESS":     ListFsuActionsLifecycleStateInProgress,
	"WAITING":         ListFsuActionsLifecycleStateWaiting,
	"UPDATING":        ListFsuActionsLifecycleStateUpdating,
	"FAILED":          ListFsuActionsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListFsuActionsLifecycleStateNeedsAttention,
	"SUCCEEDED":       ListFsuActionsLifecycleStateSucceeded,
	"CANCELING":       ListFsuActionsLifecycleStateCanceling,
	"CANCELED":        ListFsuActionsLifecycleStateCanceled,
	"UNKNOWN":         ListFsuActionsLifecycleStateUnknown,
	"DELETING":        ListFsuActionsLifecycleStateDeleting,
	"DELETED":         ListFsuActionsLifecycleStateDeleted,
}

var mappingListFsuActionsLifecycleStateEnumLowerCase = map[string]ListFsuActionsLifecycleStateEnum{
	"accepted":        ListFsuActionsLifecycleStateAccepted,
	"in_progress":     ListFsuActionsLifecycleStateInProgress,
	"waiting":         ListFsuActionsLifecycleStateWaiting,
	"updating":        ListFsuActionsLifecycleStateUpdating,
	"failed":          ListFsuActionsLifecycleStateFailed,
	"needs_attention": ListFsuActionsLifecycleStateNeedsAttention,
	"succeeded":       ListFsuActionsLifecycleStateSucceeded,
	"canceling":       ListFsuActionsLifecycleStateCanceling,
	"canceled":        ListFsuActionsLifecycleStateCanceled,
	"unknown":         ListFsuActionsLifecycleStateUnknown,
	"deleting":        ListFsuActionsLifecycleStateDeleting,
	"deleted":         ListFsuActionsLifecycleStateDeleted,
}

// GetListFsuActionsLifecycleStateEnumValues Enumerates the set of values for ListFsuActionsLifecycleStateEnum
func GetListFsuActionsLifecycleStateEnumValues() []ListFsuActionsLifecycleStateEnum {
	values := make([]ListFsuActionsLifecycleStateEnum, 0)
	for _, v := range mappingListFsuActionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuActionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListFsuActionsLifecycleStateEnum
func GetListFsuActionsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"UPDATING",
		"FAILED",
		"NEEDS_ATTENTION",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"UNKNOWN",
		"DELETING",
		"DELETED",
	}
}

// GetMappingListFsuActionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuActionsLifecycleStateEnum(val string) (ListFsuActionsLifecycleStateEnum, bool) {
	enum, ok := mappingListFsuActionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuActionsTypeEnum Enum with underlying type: string
type ListFsuActionsTypeEnum string

// Set of constants representing the allowable values for ListFsuActionsTypeEnum
const (
	ListFsuActionsTypeStage                   ListFsuActionsTypeEnum = "STAGE"
	ListFsuActionsTypePrecheck                ListFsuActionsTypeEnum = "PRECHECK"
	ListFsuActionsTypeApply                   ListFsuActionsTypeEnum = "APPLY"
	ListFsuActionsTypeRollbackAndRemoveTarget ListFsuActionsTypeEnum = "ROLLBACK_AND_REMOVE_TARGET"
	ListFsuActionsTypeCleanup                 ListFsuActionsTypeEnum = "CLEANUP"
)

var mappingListFsuActionsTypeEnum = map[string]ListFsuActionsTypeEnum{
	"STAGE":                      ListFsuActionsTypeStage,
	"PRECHECK":                   ListFsuActionsTypePrecheck,
	"APPLY":                      ListFsuActionsTypeApply,
	"ROLLBACK_AND_REMOVE_TARGET": ListFsuActionsTypeRollbackAndRemoveTarget,
	"CLEANUP":                    ListFsuActionsTypeCleanup,
}

var mappingListFsuActionsTypeEnumLowerCase = map[string]ListFsuActionsTypeEnum{
	"stage":                      ListFsuActionsTypeStage,
	"precheck":                   ListFsuActionsTypePrecheck,
	"apply":                      ListFsuActionsTypeApply,
	"rollback_and_remove_target": ListFsuActionsTypeRollbackAndRemoveTarget,
	"cleanup":                    ListFsuActionsTypeCleanup,
}

// GetListFsuActionsTypeEnumValues Enumerates the set of values for ListFsuActionsTypeEnum
func GetListFsuActionsTypeEnumValues() []ListFsuActionsTypeEnum {
	values := make([]ListFsuActionsTypeEnum, 0)
	for _, v := range mappingListFsuActionsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuActionsTypeEnumStringValues Enumerates the set of values in String for ListFsuActionsTypeEnum
func GetListFsuActionsTypeEnumStringValues() []string {
	return []string{
		"STAGE",
		"PRECHECK",
		"APPLY",
		"ROLLBACK_AND_REMOVE_TARGET",
		"CLEANUP",
	}
}

// GetMappingListFsuActionsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuActionsTypeEnum(val string) (ListFsuActionsTypeEnum, bool) {
	enum, ok := mappingListFsuActionsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuActionsSortOrderEnum Enum with underlying type: string
type ListFsuActionsSortOrderEnum string

// Set of constants representing the allowable values for ListFsuActionsSortOrderEnum
const (
	ListFsuActionsSortOrderAsc  ListFsuActionsSortOrderEnum = "ASC"
	ListFsuActionsSortOrderDesc ListFsuActionsSortOrderEnum = "DESC"
)

var mappingListFsuActionsSortOrderEnum = map[string]ListFsuActionsSortOrderEnum{
	"ASC":  ListFsuActionsSortOrderAsc,
	"DESC": ListFsuActionsSortOrderDesc,
}

var mappingListFsuActionsSortOrderEnumLowerCase = map[string]ListFsuActionsSortOrderEnum{
	"asc":  ListFsuActionsSortOrderAsc,
	"desc": ListFsuActionsSortOrderDesc,
}

// GetListFsuActionsSortOrderEnumValues Enumerates the set of values for ListFsuActionsSortOrderEnum
func GetListFsuActionsSortOrderEnumValues() []ListFsuActionsSortOrderEnum {
	values := make([]ListFsuActionsSortOrderEnum, 0)
	for _, v := range mappingListFsuActionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuActionsSortOrderEnumStringValues Enumerates the set of values in String for ListFsuActionsSortOrderEnum
func GetListFsuActionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFsuActionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuActionsSortOrderEnum(val string) (ListFsuActionsSortOrderEnum, bool) {
	enum, ok := mappingListFsuActionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuActionsSortByEnum Enum with underlying type: string
type ListFsuActionsSortByEnum string

// Set of constants representing the allowable values for ListFsuActionsSortByEnum
const (
	ListFsuActionsSortByTimecreated ListFsuActionsSortByEnum = "timeCreated"
	ListFsuActionsSortByDisplayname ListFsuActionsSortByEnum = "displayName"
)

var mappingListFsuActionsSortByEnum = map[string]ListFsuActionsSortByEnum{
	"timeCreated": ListFsuActionsSortByTimecreated,
	"displayName": ListFsuActionsSortByDisplayname,
}

var mappingListFsuActionsSortByEnumLowerCase = map[string]ListFsuActionsSortByEnum{
	"timecreated": ListFsuActionsSortByTimecreated,
	"displayname": ListFsuActionsSortByDisplayname,
}

// GetListFsuActionsSortByEnumValues Enumerates the set of values for ListFsuActionsSortByEnum
func GetListFsuActionsSortByEnumValues() []ListFsuActionsSortByEnum {
	values := make([]ListFsuActionsSortByEnum, 0)
	for _, v := range mappingListFsuActionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuActionsSortByEnumStringValues Enumerates the set of values in String for ListFsuActionsSortByEnum
func GetListFsuActionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListFsuActionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuActionsSortByEnum(val string) (ListFsuActionsSortByEnum, bool) {
	enum, ok := mappingListFsuActionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
