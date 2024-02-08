// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDeploymentBackupsRequest wrapper for the ListDeploymentBackups operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeploymentBackups.go.html to see an example of how to use ListDeploymentBackupsRequest.
type ListDeploymentBackupsRequest struct {

	// The OCID of the compartment that contains the work request. Work requests should be scoped
	// to the same compartment as the resource the work request affects. If the work request concerns
	// multiple resources, and those resources are not in the same compartment, it is up to the service team
	// to pick the primary resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the deployment in which to list resources.
	DeploymentId *string `mandatory:"false" contributesTo:"query" name:"deploymentId"`

	// A filter to return only the resources that match the 'lifecycleState' given.
	LifecycleState ListDeploymentBackupsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDeploymentBackupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDeploymentBackupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDeploymentBackupsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDeploymentBackupsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentBackupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDeploymentBackupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentBackupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDeploymentBackupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDeploymentBackupsResponse wrapper for the ListDeploymentBackups operation
type ListDeploymentBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeploymentBackupCollection instances
	DeploymentBackupCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
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
	ListDeploymentBackupsLifecycleStateWaiting        ListDeploymentBackupsLifecycleStateEnum = "WAITING"
)

var mappingListDeploymentBackupsLifecycleStateEnum = map[string]ListDeploymentBackupsLifecycleStateEnum{
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
	"WAITING":         ListDeploymentBackupsLifecycleStateWaiting,
}

var mappingListDeploymentBackupsLifecycleStateEnumLowerCase = map[string]ListDeploymentBackupsLifecycleStateEnum{
	"creating":        ListDeploymentBackupsLifecycleStateCreating,
	"updating":        ListDeploymentBackupsLifecycleStateUpdating,
	"active":          ListDeploymentBackupsLifecycleStateActive,
	"inactive":        ListDeploymentBackupsLifecycleStateInactive,
	"deleting":        ListDeploymentBackupsLifecycleStateDeleting,
	"deleted":         ListDeploymentBackupsLifecycleStateDeleted,
	"failed":          ListDeploymentBackupsLifecycleStateFailed,
	"needs_attention": ListDeploymentBackupsLifecycleStateNeedsAttention,
	"in_progress":     ListDeploymentBackupsLifecycleStateInProgress,
	"canceling":       ListDeploymentBackupsLifecycleStateCanceling,
	"canceled":        ListDeploymentBackupsLifecycleStateCanceled,
	"succeeded":       ListDeploymentBackupsLifecycleStateSucceeded,
	"waiting":         ListDeploymentBackupsLifecycleStateWaiting,
}

// GetListDeploymentBackupsLifecycleStateEnumValues Enumerates the set of values for ListDeploymentBackupsLifecycleStateEnum
func GetListDeploymentBackupsLifecycleStateEnumValues() []ListDeploymentBackupsLifecycleStateEnum {
	values := make([]ListDeploymentBackupsLifecycleStateEnum, 0)
	for _, v := range mappingListDeploymentBackupsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentBackupsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDeploymentBackupsLifecycleStateEnum
func GetListDeploymentBackupsLifecycleStateEnumStringValues() []string {
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
		"WAITING",
	}
}

// GetMappingListDeploymentBackupsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentBackupsLifecycleStateEnum(val string) (ListDeploymentBackupsLifecycleStateEnum, bool) {
	enum, ok := mappingListDeploymentBackupsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentBackupsSortOrderEnum Enum with underlying type: string
type ListDeploymentBackupsSortOrderEnum string

// Set of constants representing the allowable values for ListDeploymentBackupsSortOrderEnum
const (
	ListDeploymentBackupsSortOrderAsc  ListDeploymentBackupsSortOrderEnum = "ASC"
	ListDeploymentBackupsSortOrderDesc ListDeploymentBackupsSortOrderEnum = "DESC"
)

var mappingListDeploymentBackupsSortOrderEnum = map[string]ListDeploymentBackupsSortOrderEnum{
	"ASC":  ListDeploymentBackupsSortOrderAsc,
	"DESC": ListDeploymentBackupsSortOrderDesc,
}

var mappingListDeploymentBackupsSortOrderEnumLowerCase = map[string]ListDeploymentBackupsSortOrderEnum{
	"asc":  ListDeploymentBackupsSortOrderAsc,
	"desc": ListDeploymentBackupsSortOrderDesc,
}

// GetListDeploymentBackupsSortOrderEnumValues Enumerates the set of values for ListDeploymentBackupsSortOrderEnum
func GetListDeploymentBackupsSortOrderEnumValues() []ListDeploymentBackupsSortOrderEnum {
	values := make([]ListDeploymentBackupsSortOrderEnum, 0)
	for _, v := range mappingListDeploymentBackupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentBackupsSortOrderEnumStringValues Enumerates the set of values in String for ListDeploymentBackupsSortOrderEnum
func GetListDeploymentBackupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDeploymentBackupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentBackupsSortOrderEnum(val string) (ListDeploymentBackupsSortOrderEnum, bool) {
	enum, ok := mappingListDeploymentBackupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentBackupsSortByEnum Enum with underlying type: string
type ListDeploymentBackupsSortByEnum string

// Set of constants representing the allowable values for ListDeploymentBackupsSortByEnum
const (
	ListDeploymentBackupsSortByTimecreated ListDeploymentBackupsSortByEnum = "timeCreated"
	ListDeploymentBackupsSortByDisplayname ListDeploymentBackupsSortByEnum = "displayName"
)

var mappingListDeploymentBackupsSortByEnum = map[string]ListDeploymentBackupsSortByEnum{
	"timeCreated": ListDeploymentBackupsSortByTimecreated,
	"displayName": ListDeploymentBackupsSortByDisplayname,
}

var mappingListDeploymentBackupsSortByEnumLowerCase = map[string]ListDeploymentBackupsSortByEnum{
	"timecreated": ListDeploymentBackupsSortByTimecreated,
	"displayname": ListDeploymentBackupsSortByDisplayname,
}

// GetListDeploymentBackupsSortByEnumValues Enumerates the set of values for ListDeploymentBackupsSortByEnum
func GetListDeploymentBackupsSortByEnumValues() []ListDeploymentBackupsSortByEnum {
	values := make([]ListDeploymentBackupsSortByEnum, 0)
	for _, v := range mappingListDeploymentBackupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentBackupsSortByEnumStringValues Enumerates the set of values in String for ListDeploymentBackupsSortByEnum
func GetListDeploymentBackupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDeploymentBackupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentBackupsSortByEnum(val string) (ListDeploymentBackupsSortByEnum, bool) {
	enum, ok := mappingListDeploymentBackupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
