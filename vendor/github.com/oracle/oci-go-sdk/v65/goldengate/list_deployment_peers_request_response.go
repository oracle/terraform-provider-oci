// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDeploymentPeersRequest wrapper for the ListDeploymentPeers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeploymentPeers.go.html to see an example of how to use ListDeploymentPeersRequest.
type ListDeploymentPeersRequest struct {

	// A unique Deployment identifier.
	DeploymentId *string `mandatory:"true" contributesTo:"path" name:"deploymentId"`

	// A filter to return only the resources that match the 'lifecycleState' given.
	LifecycleState ListDeploymentPeersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDeploymentPeersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
	SortBy ListDeploymentPeersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeploymentPeersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeploymentPeersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDeploymentPeersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeploymentPeersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDeploymentPeersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDeploymentPeersLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDeploymentPeersLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentPeersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDeploymentPeersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentPeersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDeploymentPeersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDeploymentPeersResponse wrapper for the ListDeploymentPeers operation
type ListDeploymentPeersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeploymentPeerCollection instances
	DeploymentPeerCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeploymentPeersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeploymentPeersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeploymentPeersLifecycleStateEnum Enum with underlying type: string
type ListDeploymentPeersLifecycleStateEnum string

// Set of constants representing the allowable values for ListDeploymentPeersLifecycleStateEnum
const (
	ListDeploymentPeersLifecycleStateCreating       ListDeploymentPeersLifecycleStateEnum = "CREATING"
	ListDeploymentPeersLifecycleStateUpdating       ListDeploymentPeersLifecycleStateEnum = "UPDATING"
	ListDeploymentPeersLifecycleStateActive         ListDeploymentPeersLifecycleStateEnum = "ACTIVE"
	ListDeploymentPeersLifecycleStateInactive       ListDeploymentPeersLifecycleStateEnum = "INACTIVE"
	ListDeploymentPeersLifecycleStateDeleting       ListDeploymentPeersLifecycleStateEnum = "DELETING"
	ListDeploymentPeersLifecycleStateDeleted        ListDeploymentPeersLifecycleStateEnum = "DELETED"
	ListDeploymentPeersLifecycleStateFailed         ListDeploymentPeersLifecycleStateEnum = "FAILED"
	ListDeploymentPeersLifecycleStateNeedsAttention ListDeploymentPeersLifecycleStateEnum = "NEEDS_ATTENTION"
	ListDeploymentPeersLifecycleStateInProgress     ListDeploymentPeersLifecycleStateEnum = "IN_PROGRESS"
	ListDeploymentPeersLifecycleStateCanceling      ListDeploymentPeersLifecycleStateEnum = "CANCELING"
	ListDeploymentPeersLifecycleStateCanceled       ListDeploymentPeersLifecycleStateEnum = "CANCELED"
	ListDeploymentPeersLifecycleStateSucceeded      ListDeploymentPeersLifecycleStateEnum = "SUCCEEDED"
	ListDeploymentPeersLifecycleStateWaiting        ListDeploymentPeersLifecycleStateEnum = "WAITING"
)

var mappingListDeploymentPeersLifecycleStateEnum = map[string]ListDeploymentPeersLifecycleStateEnum{
	"CREATING":        ListDeploymentPeersLifecycleStateCreating,
	"UPDATING":        ListDeploymentPeersLifecycleStateUpdating,
	"ACTIVE":          ListDeploymentPeersLifecycleStateActive,
	"INACTIVE":        ListDeploymentPeersLifecycleStateInactive,
	"DELETING":        ListDeploymentPeersLifecycleStateDeleting,
	"DELETED":         ListDeploymentPeersLifecycleStateDeleted,
	"FAILED":          ListDeploymentPeersLifecycleStateFailed,
	"NEEDS_ATTENTION": ListDeploymentPeersLifecycleStateNeedsAttention,
	"IN_PROGRESS":     ListDeploymentPeersLifecycleStateInProgress,
	"CANCELING":       ListDeploymentPeersLifecycleStateCanceling,
	"CANCELED":        ListDeploymentPeersLifecycleStateCanceled,
	"SUCCEEDED":       ListDeploymentPeersLifecycleStateSucceeded,
	"WAITING":         ListDeploymentPeersLifecycleStateWaiting,
}

var mappingListDeploymentPeersLifecycleStateEnumLowerCase = map[string]ListDeploymentPeersLifecycleStateEnum{
	"creating":        ListDeploymentPeersLifecycleStateCreating,
	"updating":        ListDeploymentPeersLifecycleStateUpdating,
	"active":          ListDeploymentPeersLifecycleStateActive,
	"inactive":        ListDeploymentPeersLifecycleStateInactive,
	"deleting":        ListDeploymentPeersLifecycleStateDeleting,
	"deleted":         ListDeploymentPeersLifecycleStateDeleted,
	"failed":          ListDeploymentPeersLifecycleStateFailed,
	"needs_attention": ListDeploymentPeersLifecycleStateNeedsAttention,
	"in_progress":     ListDeploymentPeersLifecycleStateInProgress,
	"canceling":       ListDeploymentPeersLifecycleStateCanceling,
	"canceled":        ListDeploymentPeersLifecycleStateCanceled,
	"succeeded":       ListDeploymentPeersLifecycleStateSucceeded,
	"waiting":         ListDeploymentPeersLifecycleStateWaiting,
}

// GetListDeploymentPeersLifecycleStateEnumValues Enumerates the set of values for ListDeploymentPeersLifecycleStateEnum
func GetListDeploymentPeersLifecycleStateEnumValues() []ListDeploymentPeersLifecycleStateEnum {
	values := make([]ListDeploymentPeersLifecycleStateEnum, 0)
	for _, v := range mappingListDeploymentPeersLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentPeersLifecycleStateEnumStringValues Enumerates the set of values in String for ListDeploymentPeersLifecycleStateEnum
func GetListDeploymentPeersLifecycleStateEnumStringValues() []string {
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

// GetMappingListDeploymentPeersLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentPeersLifecycleStateEnum(val string) (ListDeploymentPeersLifecycleStateEnum, bool) {
	enum, ok := mappingListDeploymentPeersLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentPeersSortOrderEnum Enum with underlying type: string
type ListDeploymentPeersSortOrderEnum string

// Set of constants representing the allowable values for ListDeploymentPeersSortOrderEnum
const (
	ListDeploymentPeersSortOrderAsc  ListDeploymentPeersSortOrderEnum = "ASC"
	ListDeploymentPeersSortOrderDesc ListDeploymentPeersSortOrderEnum = "DESC"
)

var mappingListDeploymentPeersSortOrderEnum = map[string]ListDeploymentPeersSortOrderEnum{
	"ASC":  ListDeploymentPeersSortOrderAsc,
	"DESC": ListDeploymentPeersSortOrderDesc,
}

var mappingListDeploymentPeersSortOrderEnumLowerCase = map[string]ListDeploymentPeersSortOrderEnum{
	"asc":  ListDeploymentPeersSortOrderAsc,
	"desc": ListDeploymentPeersSortOrderDesc,
}

// GetListDeploymentPeersSortOrderEnumValues Enumerates the set of values for ListDeploymentPeersSortOrderEnum
func GetListDeploymentPeersSortOrderEnumValues() []ListDeploymentPeersSortOrderEnum {
	values := make([]ListDeploymentPeersSortOrderEnum, 0)
	for _, v := range mappingListDeploymentPeersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentPeersSortOrderEnumStringValues Enumerates the set of values in String for ListDeploymentPeersSortOrderEnum
func GetListDeploymentPeersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDeploymentPeersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentPeersSortOrderEnum(val string) (ListDeploymentPeersSortOrderEnum, bool) {
	enum, ok := mappingListDeploymentPeersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentPeersSortByEnum Enum with underlying type: string
type ListDeploymentPeersSortByEnum string

// Set of constants representing the allowable values for ListDeploymentPeersSortByEnum
const (
	ListDeploymentPeersSortByTimecreated ListDeploymentPeersSortByEnum = "timeCreated"
	ListDeploymentPeersSortByDisplayname ListDeploymentPeersSortByEnum = "displayName"
)

var mappingListDeploymentPeersSortByEnum = map[string]ListDeploymentPeersSortByEnum{
	"timeCreated": ListDeploymentPeersSortByTimecreated,
	"displayName": ListDeploymentPeersSortByDisplayname,
}

var mappingListDeploymentPeersSortByEnumLowerCase = map[string]ListDeploymentPeersSortByEnum{
	"timecreated": ListDeploymentPeersSortByTimecreated,
	"displayname": ListDeploymentPeersSortByDisplayname,
}

// GetListDeploymentPeersSortByEnumValues Enumerates the set of values for ListDeploymentPeersSortByEnum
func GetListDeploymentPeersSortByEnumValues() []ListDeploymentPeersSortByEnum {
	values := make([]ListDeploymentPeersSortByEnum, 0)
	for _, v := range mappingListDeploymentPeersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentPeersSortByEnumStringValues Enumerates the set of values in String for ListDeploymentPeersSortByEnum
func GetListDeploymentPeersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDeploymentPeersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentPeersSortByEnum(val string) (ListDeploymentPeersSortByEnum, bool) {
	enum, ok := mappingListDeploymentPeersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
