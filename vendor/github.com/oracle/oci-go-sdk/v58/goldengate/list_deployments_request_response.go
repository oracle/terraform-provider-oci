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

// ListDeploymentsRequest wrapper for the ListDeployments operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeployments.go.html to see an example of how to use ListDeploymentsRequest.
type ListDeploymentsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources that match the 'lifecycleState' given.
	LifecycleState ListDeploymentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the 'lifecycleSubState' given.
	LifecycleSubState ListDeploymentsLifecycleSubStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleSubState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only the resources that match the 'fqdn' given.
	Fqdn *string `mandatory:"false" contributesTo:"query" name:"fqdn"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDeploymentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is descending.  Default order for 'displayName' is ascending. If no value is specified timeCreated is the default.
	SortBy ListDeploymentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeploymentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeploymentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDeploymentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeploymentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDeploymentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDeploymentsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDeploymentsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentsLifecycleSubStateEnum(string(request.LifecycleSubState)); !ok && request.LifecycleSubState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubState: %s. Supported values are: %s.", request.LifecycleSubState, strings.Join(GetListDeploymentsLifecycleSubStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDeploymentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDeploymentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDeploymentsResponse wrapper for the ListDeployments operation
type ListDeploymentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeploymentCollection instances
	DeploymentCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeploymentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeploymentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeploymentsLifecycleStateEnum Enum with underlying type: string
type ListDeploymentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDeploymentsLifecycleStateEnum
const (
	ListDeploymentsLifecycleStateCreating       ListDeploymentsLifecycleStateEnum = "CREATING"
	ListDeploymentsLifecycleStateUpdating       ListDeploymentsLifecycleStateEnum = "UPDATING"
	ListDeploymentsLifecycleStateActive         ListDeploymentsLifecycleStateEnum = "ACTIVE"
	ListDeploymentsLifecycleStateInactive       ListDeploymentsLifecycleStateEnum = "INACTIVE"
	ListDeploymentsLifecycleStateDeleting       ListDeploymentsLifecycleStateEnum = "DELETING"
	ListDeploymentsLifecycleStateDeleted        ListDeploymentsLifecycleStateEnum = "DELETED"
	ListDeploymentsLifecycleStateFailed         ListDeploymentsLifecycleStateEnum = "FAILED"
	ListDeploymentsLifecycleStateNeedsAttention ListDeploymentsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListDeploymentsLifecycleStateInProgress     ListDeploymentsLifecycleStateEnum = "IN_PROGRESS"
	ListDeploymentsLifecycleStateCanceling      ListDeploymentsLifecycleStateEnum = "CANCELING"
	ListDeploymentsLifecycleStateCanceled       ListDeploymentsLifecycleStateEnum = "CANCELED"
	ListDeploymentsLifecycleStateSucceeded      ListDeploymentsLifecycleStateEnum = "SUCCEEDED"
)

var mappingListDeploymentsLifecycleStateEnum = map[string]ListDeploymentsLifecycleStateEnum{
	"CREATING":        ListDeploymentsLifecycleStateCreating,
	"UPDATING":        ListDeploymentsLifecycleStateUpdating,
	"ACTIVE":          ListDeploymentsLifecycleStateActive,
	"INACTIVE":        ListDeploymentsLifecycleStateInactive,
	"DELETING":        ListDeploymentsLifecycleStateDeleting,
	"DELETED":         ListDeploymentsLifecycleStateDeleted,
	"FAILED":          ListDeploymentsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListDeploymentsLifecycleStateNeedsAttention,
	"IN_PROGRESS":     ListDeploymentsLifecycleStateInProgress,
	"CANCELING":       ListDeploymentsLifecycleStateCanceling,
	"CANCELED":        ListDeploymentsLifecycleStateCanceled,
	"SUCCEEDED":       ListDeploymentsLifecycleStateSucceeded,
}

// GetListDeploymentsLifecycleStateEnumValues Enumerates the set of values for ListDeploymentsLifecycleStateEnum
func GetListDeploymentsLifecycleStateEnumValues() []ListDeploymentsLifecycleStateEnum {
	values := make([]ListDeploymentsLifecycleStateEnum, 0)
	for _, v := range mappingListDeploymentsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDeploymentsLifecycleStateEnum
func GetListDeploymentsLifecycleStateEnumStringValues() []string {
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

// GetMappingListDeploymentsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentsLifecycleStateEnum(val string) (ListDeploymentsLifecycleStateEnum, bool) {
	mappingListDeploymentsLifecycleStateEnumIgnoreCase := make(map[string]ListDeploymentsLifecycleStateEnum)
	for k, v := range mappingListDeploymentsLifecycleStateEnum {
		mappingListDeploymentsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDeploymentsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentsLifecycleSubStateEnum Enum with underlying type: string
type ListDeploymentsLifecycleSubStateEnum string

// Set of constants representing the allowable values for ListDeploymentsLifecycleSubStateEnum
const (
	ListDeploymentsLifecycleSubStateRecovering       ListDeploymentsLifecycleSubStateEnum = "RECOVERING"
	ListDeploymentsLifecycleSubStateStarting         ListDeploymentsLifecycleSubStateEnum = "STARTING"
	ListDeploymentsLifecycleSubStateStopping         ListDeploymentsLifecycleSubStateEnum = "STOPPING"
	ListDeploymentsLifecycleSubStateMoving           ListDeploymentsLifecycleSubStateEnum = "MOVING"
	ListDeploymentsLifecycleSubStateUpgrading        ListDeploymentsLifecycleSubStateEnum = "UPGRADING"
	ListDeploymentsLifecycleSubStateRestoring        ListDeploymentsLifecycleSubStateEnum = "RESTORING"
	ListDeploymentsLifecycleSubStateBackupInProgress ListDeploymentsLifecycleSubStateEnum = "BACKUP_IN_PROGRESS"
)

var mappingListDeploymentsLifecycleSubStateEnum = map[string]ListDeploymentsLifecycleSubStateEnum{
	"RECOVERING":         ListDeploymentsLifecycleSubStateRecovering,
	"STARTING":           ListDeploymentsLifecycleSubStateStarting,
	"STOPPING":           ListDeploymentsLifecycleSubStateStopping,
	"MOVING":             ListDeploymentsLifecycleSubStateMoving,
	"UPGRADING":          ListDeploymentsLifecycleSubStateUpgrading,
	"RESTORING":          ListDeploymentsLifecycleSubStateRestoring,
	"BACKUP_IN_PROGRESS": ListDeploymentsLifecycleSubStateBackupInProgress,
}

// GetListDeploymentsLifecycleSubStateEnumValues Enumerates the set of values for ListDeploymentsLifecycleSubStateEnum
func GetListDeploymentsLifecycleSubStateEnumValues() []ListDeploymentsLifecycleSubStateEnum {
	values := make([]ListDeploymentsLifecycleSubStateEnum, 0)
	for _, v := range mappingListDeploymentsLifecycleSubStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentsLifecycleSubStateEnumStringValues Enumerates the set of values in String for ListDeploymentsLifecycleSubStateEnum
func GetListDeploymentsLifecycleSubStateEnumStringValues() []string {
	return []string{
		"RECOVERING",
		"STARTING",
		"STOPPING",
		"MOVING",
		"UPGRADING",
		"RESTORING",
		"BACKUP_IN_PROGRESS",
	}
}

// GetMappingListDeploymentsLifecycleSubStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentsLifecycleSubStateEnum(val string) (ListDeploymentsLifecycleSubStateEnum, bool) {
	mappingListDeploymentsLifecycleSubStateEnumIgnoreCase := make(map[string]ListDeploymentsLifecycleSubStateEnum)
	for k, v := range mappingListDeploymentsLifecycleSubStateEnum {
		mappingListDeploymentsLifecycleSubStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDeploymentsLifecycleSubStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentsSortOrderEnum Enum with underlying type: string
type ListDeploymentsSortOrderEnum string

// Set of constants representing the allowable values for ListDeploymentsSortOrderEnum
const (
	ListDeploymentsSortOrderAsc  ListDeploymentsSortOrderEnum = "ASC"
	ListDeploymentsSortOrderDesc ListDeploymentsSortOrderEnum = "DESC"
)

var mappingListDeploymentsSortOrderEnum = map[string]ListDeploymentsSortOrderEnum{
	"ASC":  ListDeploymentsSortOrderAsc,
	"DESC": ListDeploymentsSortOrderDesc,
}

// GetListDeploymentsSortOrderEnumValues Enumerates the set of values for ListDeploymentsSortOrderEnum
func GetListDeploymentsSortOrderEnumValues() []ListDeploymentsSortOrderEnum {
	values := make([]ListDeploymentsSortOrderEnum, 0)
	for _, v := range mappingListDeploymentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentsSortOrderEnumStringValues Enumerates the set of values in String for ListDeploymentsSortOrderEnum
func GetListDeploymentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDeploymentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentsSortOrderEnum(val string) (ListDeploymentsSortOrderEnum, bool) {
	mappingListDeploymentsSortOrderEnumIgnoreCase := make(map[string]ListDeploymentsSortOrderEnum)
	for k, v := range mappingListDeploymentsSortOrderEnum {
		mappingListDeploymentsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDeploymentsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentsSortByEnum Enum with underlying type: string
type ListDeploymentsSortByEnum string

// Set of constants representing the allowable values for ListDeploymentsSortByEnum
const (
	ListDeploymentsSortByTimecreated ListDeploymentsSortByEnum = "timeCreated"
	ListDeploymentsSortByDisplayname ListDeploymentsSortByEnum = "displayName"
)

var mappingListDeploymentsSortByEnum = map[string]ListDeploymentsSortByEnum{
	"timeCreated": ListDeploymentsSortByTimecreated,
	"displayName": ListDeploymentsSortByDisplayname,
}

// GetListDeploymentsSortByEnumValues Enumerates the set of values for ListDeploymentsSortByEnum
func GetListDeploymentsSortByEnumValues() []ListDeploymentsSortByEnum {
	values := make([]ListDeploymentsSortByEnum, 0)
	for _, v := range mappingListDeploymentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentsSortByEnumStringValues Enumerates the set of values in String for ListDeploymentsSortByEnum
func GetListDeploymentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDeploymentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentsSortByEnum(val string) (ListDeploymentsSortByEnum, bool) {
	mappingListDeploymentsSortByEnumIgnoreCase := make(map[string]ListDeploymentsSortByEnum)
	for k, v := range mappingListDeploymentsSortByEnum {
		mappingListDeploymentsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDeploymentsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
