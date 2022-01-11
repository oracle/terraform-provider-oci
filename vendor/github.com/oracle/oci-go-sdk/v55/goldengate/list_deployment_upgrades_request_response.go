// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v55/common"
	"net/http"
	"strings"
)

// ListDeploymentUpgradesRequest wrapper for the ListDeploymentUpgrades operation
type ListDeploymentUpgradesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ID of the deployment in which to list resources.
	DeploymentId *string `mandatory:"false" contributesTo:"query" name:"deploymentId"`

	// A filter to return only the resources that match the 'lifecycleState' given.
	LifecycleState ListDeploymentUpgradesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDeploymentUpgradesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is descending.  Default order for 'displayName' is ascending. If no value is specified timeCreated is the default.
	SortBy ListDeploymentUpgradesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeploymentUpgradesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeploymentUpgradesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDeploymentUpgradesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeploymentUpgradesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDeploymentUpgradesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingListDeploymentUpgradesLifecycleStateEnum[string(request.LifecycleState)]; !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDeploymentUpgradesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := mappingListDeploymentUpgradesSortOrderEnum[string(request.SortOrder)]; !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDeploymentUpgradesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := mappingListDeploymentUpgradesSortByEnum[string(request.SortBy)]; !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDeploymentUpgradesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDeploymentUpgradesResponse wrapper for the ListDeploymentUpgrades operation
type ListDeploymentUpgradesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeploymentUpgradeCollection instances
	DeploymentUpgradeCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeploymentUpgradesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeploymentUpgradesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeploymentUpgradesLifecycleStateEnum Enum with underlying type: string
type ListDeploymentUpgradesLifecycleStateEnum string

// Set of constants representing the allowable values for ListDeploymentUpgradesLifecycleStateEnum
const (
	ListDeploymentUpgradesLifecycleStateCreating       ListDeploymentUpgradesLifecycleStateEnum = "CREATING"
	ListDeploymentUpgradesLifecycleStateUpdating       ListDeploymentUpgradesLifecycleStateEnum = "UPDATING"
	ListDeploymentUpgradesLifecycleStateActive         ListDeploymentUpgradesLifecycleStateEnum = "ACTIVE"
	ListDeploymentUpgradesLifecycleStateInactive       ListDeploymentUpgradesLifecycleStateEnum = "INACTIVE"
	ListDeploymentUpgradesLifecycleStateDeleting       ListDeploymentUpgradesLifecycleStateEnum = "DELETING"
	ListDeploymentUpgradesLifecycleStateDeleted        ListDeploymentUpgradesLifecycleStateEnum = "DELETED"
	ListDeploymentUpgradesLifecycleStateFailed         ListDeploymentUpgradesLifecycleStateEnum = "FAILED"
	ListDeploymentUpgradesLifecycleStateNeedsAttention ListDeploymentUpgradesLifecycleStateEnum = "NEEDS_ATTENTION"
	ListDeploymentUpgradesLifecycleStateInProgress     ListDeploymentUpgradesLifecycleStateEnum = "IN_PROGRESS"
	ListDeploymentUpgradesLifecycleStateCanceling      ListDeploymentUpgradesLifecycleStateEnum = "CANCELING"
	ListDeploymentUpgradesLifecycleStateCanceled       ListDeploymentUpgradesLifecycleStateEnum = "CANCELED"
	ListDeploymentUpgradesLifecycleStateSucceeded      ListDeploymentUpgradesLifecycleStateEnum = "SUCCEEDED"
)

var mappingListDeploymentUpgradesLifecycleStateEnum = map[string]ListDeploymentUpgradesLifecycleStateEnum{
	"CREATING":        ListDeploymentUpgradesLifecycleStateCreating,
	"UPDATING":        ListDeploymentUpgradesLifecycleStateUpdating,
	"ACTIVE":          ListDeploymentUpgradesLifecycleStateActive,
	"INACTIVE":        ListDeploymentUpgradesLifecycleStateInactive,
	"DELETING":        ListDeploymentUpgradesLifecycleStateDeleting,
	"DELETED":         ListDeploymentUpgradesLifecycleStateDeleted,
	"FAILED":          ListDeploymentUpgradesLifecycleStateFailed,
	"NEEDS_ATTENTION": ListDeploymentUpgradesLifecycleStateNeedsAttention,
	"IN_PROGRESS":     ListDeploymentUpgradesLifecycleStateInProgress,
	"CANCELING":       ListDeploymentUpgradesLifecycleStateCanceling,
	"CANCELED":        ListDeploymentUpgradesLifecycleStateCanceled,
	"SUCCEEDED":       ListDeploymentUpgradesLifecycleStateSucceeded,
}

// GetListDeploymentUpgradesLifecycleStateEnumValues Enumerates the set of values for ListDeploymentUpgradesLifecycleStateEnum
func GetListDeploymentUpgradesLifecycleStateEnumValues() []ListDeploymentUpgradesLifecycleStateEnum {
	values := make([]ListDeploymentUpgradesLifecycleStateEnum, 0)
	for _, v := range mappingListDeploymentUpgradesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentUpgradesLifecycleStateEnumStringValues Enumerates the set of values in String for ListDeploymentUpgradesLifecycleStateEnum
func GetListDeploymentUpgradesLifecycleStateEnumStringValues() []string {
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

// ListDeploymentUpgradesSortOrderEnum Enum with underlying type: string
type ListDeploymentUpgradesSortOrderEnum string

// Set of constants representing the allowable values for ListDeploymentUpgradesSortOrderEnum
const (
	ListDeploymentUpgradesSortOrderAsc  ListDeploymentUpgradesSortOrderEnum = "ASC"
	ListDeploymentUpgradesSortOrderDesc ListDeploymentUpgradesSortOrderEnum = "DESC"
)

var mappingListDeploymentUpgradesSortOrderEnum = map[string]ListDeploymentUpgradesSortOrderEnum{
	"ASC":  ListDeploymentUpgradesSortOrderAsc,
	"DESC": ListDeploymentUpgradesSortOrderDesc,
}

// GetListDeploymentUpgradesSortOrderEnumValues Enumerates the set of values for ListDeploymentUpgradesSortOrderEnum
func GetListDeploymentUpgradesSortOrderEnumValues() []ListDeploymentUpgradesSortOrderEnum {
	values := make([]ListDeploymentUpgradesSortOrderEnum, 0)
	for _, v := range mappingListDeploymentUpgradesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentUpgradesSortOrderEnumStringValues Enumerates the set of values in String for ListDeploymentUpgradesSortOrderEnum
func GetListDeploymentUpgradesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// ListDeploymentUpgradesSortByEnum Enum with underlying type: string
type ListDeploymentUpgradesSortByEnum string

// Set of constants representing the allowable values for ListDeploymentUpgradesSortByEnum
const (
	ListDeploymentUpgradesSortByTimecreated ListDeploymentUpgradesSortByEnum = "timeCreated"
	ListDeploymentUpgradesSortByDisplayname ListDeploymentUpgradesSortByEnum = "displayName"
)

var mappingListDeploymentUpgradesSortByEnum = map[string]ListDeploymentUpgradesSortByEnum{
	"timeCreated": ListDeploymentUpgradesSortByTimecreated,
	"displayName": ListDeploymentUpgradesSortByDisplayname,
}

// GetListDeploymentUpgradesSortByEnumValues Enumerates the set of values for ListDeploymentUpgradesSortByEnum
func GetListDeploymentUpgradesSortByEnumValues() []ListDeploymentUpgradesSortByEnum {
	values := make([]ListDeploymentUpgradesSortByEnum, 0)
	for _, v := range mappingListDeploymentUpgradesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentUpgradesSortByEnumStringValues Enumerates the set of values in String for ListDeploymentUpgradesSortByEnum
func GetListDeploymentUpgradesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}
