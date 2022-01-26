// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDeploymentBackupsRequest wrapper for the ListDeploymentBackups operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeploymentBackups.go.html to see an example of how to use ListDeploymentBackupsRequest.
type ListDeploymentBackupsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ID of the deployment in which to list resources.
	DeploymentId *string `mandatory:"false" contributesTo:"query" name:"deploymentId"`

	// A filter to return only the resources that match the 'lifecycleState' given.
	LifecycleState ListDeploymentBackupsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDeploymentBackupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is descending.  Default order for 'displayName' is ascending. If no value is specified timeCreated is the default.
	SortBy ListDeploymentBackupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeploymentBackupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeploymentBackupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDeploymentBackupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeploymentBackupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDeploymentBackupsResponse wrapper for the ListDeploymentBackups operation
type ListDeploymentBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeploymentBackupCollection instances
	DeploymentBackupCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeploymentBackupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeploymentBackupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeploymentBackupsLifecycleStateEnum Enum with underlying type: string
type ListDeploymentBackupsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDeploymentBackupsLifecycleStateEnum
const (
	ListDeploymentBackupsLifecycleStateCreating       ListDeploymentBackupsLifecycleStateEnum = "CREATING"
	ListDeploymentBackupsLifecycleStateUpdating       ListDeploymentBackupsLifecycleStateEnum = "UPDATING"
	ListDeploymentBackupsLifecycleStateActive         ListDeploymentBackupsLifecycleStateEnum = "ACTIVE"
	ListDeploymentBackupsLifecycleStateInactive       ListDeploymentBackupsLifecycleStateEnum = "INACTIVE"
	ListDeploymentBackupsLifecycleStateDeleting       ListDeploymentBackupsLifecycleStateEnum = "DELETING"
	ListDeploymentBackupsLifecycleStateDeleted        ListDeploymentBackupsLifecycleStateEnum = "DELETED"
	ListDeploymentBackupsLifecycleStateFailed         ListDeploymentBackupsLifecycleStateEnum = "FAILED"
	ListDeploymentBackupsLifecycleStateNeedsAttention ListDeploymentBackupsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListDeploymentBackupsLifecycleStateInProgress     ListDeploymentBackupsLifecycleStateEnum = "IN_PROGRESS"
	ListDeploymentBackupsLifecycleStateCanceling      ListDeploymentBackupsLifecycleStateEnum = "CANCELING"
	ListDeploymentBackupsLifecycleStateCanceled       ListDeploymentBackupsLifecycleStateEnum = "CANCELED"
	ListDeploymentBackupsLifecycleStateSucceeded      ListDeploymentBackupsLifecycleStateEnum = "SUCCEEDED"
)

var mappingListDeploymentBackupsLifecycleState = map[string]ListDeploymentBackupsLifecycleStateEnum{
	"CREATING":        ListDeploymentBackupsLifecycleStateCreating,
	"UPDATING":        ListDeploymentBackupsLifecycleStateUpdating,
	"ACTIVE":          ListDeploymentBackupsLifecycleStateActive,
	"INACTIVE":        ListDeploymentBackupsLifecycleStateInactive,
	"DELETING":        ListDeploymentBackupsLifecycleStateDeleting,
	"DELETED":         ListDeploymentBackupsLifecycleStateDeleted,
	"FAILED":          ListDeploymentBackupsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListDeploymentBackupsLifecycleStateNeedsAttention,
	"IN_PROGRESS":     ListDeploymentBackupsLifecycleStateInProgress,
	"CANCELING":       ListDeploymentBackupsLifecycleStateCanceling,
	"CANCELED":        ListDeploymentBackupsLifecycleStateCanceled,
	"SUCCEEDED":       ListDeploymentBackupsLifecycleStateSucceeded,
}

// GetListDeploymentBackupsLifecycleStateEnumValues Enumerates the set of values for ListDeploymentBackupsLifecycleStateEnum
func GetListDeploymentBackupsLifecycleStateEnumValues() []ListDeploymentBackupsLifecycleStateEnum {
	values := make([]ListDeploymentBackupsLifecycleStateEnum, 0)
	for _, v := range mappingListDeploymentBackupsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListDeploymentBackupsSortOrderEnum Enum with underlying type: string
type ListDeploymentBackupsSortOrderEnum string

// Set of constants representing the allowable values for ListDeploymentBackupsSortOrderEnum
const (
	ListDeploymentBackupsSortOrderAsc  ListDeploymentBackupsSortOrderEnum = "ASC"
	ListDeploymentBackupsSortOrderDesc ListDeploymentBackupsSortOrderEnum = "DESC"
)

var mappingListDeploymentBackupsSortOrder = map[string]ListDeploymentBackupsSortOrderEnum{
	"ASC":  ListDeploymentBackupsSortOrderAsc,
	"DESC": ListDeploymentBackupsSortOrderDesc,
}

// GetListDeploymentBackupsSortOrderEnumValues Enumerates the set of values for ListDeploymentBackupsSortOrderEnum
func GetListDeploymentBackupsSortOrderEnumValues() []ListDeploymentBackupsSortOrderEnum {
	values := make([]ListDeploymentBackupsSortOrderEnum, 0)
	for _, v := range mappingListDeploymentBackupsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDeploymentBackupsSortByEnum Enum with underlying type: string
type ListDeploymentBackupsSortByEnum string

// Set of constants representing the allowable values for ListDeploymentBackupsSortByEnum
const (
	ListDeploymentBackupsSortByTimecreated ListDeploymentBackupsSortByEnum = "timeCreated"
	ListDeploymentBackupsSortByDisplayname ListDeploymentBackupsSortByEnum = "displayName"
)

var mappingListDeploymentBackupsSortBy = map[string]ListDeploymentBackupsSortByEnum{
	"timeCreated": ListDeploymentBackupsSortByTimecreated,
	"displayName": ListDeploymentBackupsSortByDisplayname,
}

// GetListDeploymentBackupsSortByEnumValues Enumerates the set of values for ListDeploymentBackupsSortByEnum
func GetListDeploymentBackupsSortByEnumValues() []ListDeploymentBackupsSortByEnum {
	values := make([]ListDeploymentBackupsSortByEnum, 0)
	for _, v := range mappingListDeploymentBackupsSortBy {
		values = append(values, v)
	}
	return values
}
