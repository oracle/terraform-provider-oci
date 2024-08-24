// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package delegateaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDelegatedResourceAccessRequestsRequest wrapper for the ListDelegatedResourceAccessRequests operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ListDelegatedResourceAccessRequests.go.html to see an example of how to use ListDelegatedResourceAccessRequestsRequest.
type ListDelegatedResourceAccessRequestsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// unique Delegation Control identifier
	DelegationControlId *string `mandatory:"false" contributesTo:"query" name:"delegationControlId"`

	// A filter to return only Delegated Resource Access Requests for the given resource identifier.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// A filter to return only Delegated Resource Access Requests whose lifecycleState matches the given Delegated Resource Access Request lifecycleState.
	LifecycleState DelegatedResourceAccessRequestLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only Delegated Resource Access Requests whose status matches the given Delegated Resource Access Request status.
	RequestStatus ListDelegatedResourceAccessRequestsRequestStatusEnum `mandatory:"false" contributesTo:"query" name:"requestStatus" omitEmpty:"true"`

	// Query start time in UTC in ISO 8601 format(inclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// timeIntervalStart and timeIntervalEnd parameters cannot be used together.
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// Query end time in UTC in ISO 8601 format(inclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// timeIntervalStart and timeIntervalEnd parameters cannot be used together.
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDelegatedResourceAccessRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListDelegatedResourceAccessRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDelegatedResourceAccessRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDelegatedResourceAccessRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDelegatedResourceAccessRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDelegatedResourceAccessRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDelegatedResourceAccessRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDelegatedResourceAccessRequestLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDelegatedResourceAccessRequestLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDelegatedResourceAccessRequestsRequestStatusEnum(string(request.RequestStatus)); !ok && request.RequestStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestStatus: %s. Supported values are: %s.", request.RequestStatus, strings.Join(GetListDelegatedResourceAccessRequestsRequestStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDelegatedResourceAccessRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDelegatedResourceAccessRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDelegatedResourceAccessRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDelegatedResourceAccessRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDelegatedResourceAccessRequestsResponse wrapper for the ListDelegatedResourceAccessRequests operation
type ListDelegatedResourceAccessRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DelegatedResourceAccessRequestSummaryCollection instances
	DelegatedResourceAccessRequestSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDelegatedResourceAccessRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDelegatedResourceAccessRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDelegatedResourceAccessRequestsRequestStatusEnum Enum with underlying type: string
type ListDelegatedResourceAccessRequestsRequestStatusEnum string

// Set of constants representing the allowable values for ListDelegatedResourceAccessRequestsRequestStatusEnum
const (
	ListDelegatedResourceAccessRequestsRequestStatusCreated                   ListDelegatedResourceAccessRequestsRequestStatusEnum = "CREATED"
	ListDelegatedResourceAccessRequestsRequestStatusApprovalWaiting           ListDelegatedResourceAccessRequestsRequestStatusEnum = "APPROVAL_WAITING"
	ListDelegatedResourceAccessRequestsRequestStatusOperatorAssignmentWaiting ListDelegatedResourceAccessRequestsRequestStatusEnum = "OPERATOR_ASSIGNMENT_WAITING"
	ListDelegatedResourceAccessRequestsRequestStatusPreapproved               ListDelegatedResourceAccessRequestsRequestStatusEnum = "PREAPPROVED"
	ListDelegatedResourceAccessRequestsRequestStatusApproved                  ListDelegatedResourceAccessRequestsRequestStatusEnum = "APPROVED"
	ListDelegatedResourceAccessRequestsRequestStatusApprovedForFuture         ListDelegatedResourceAccessRequestsRequestStatusEnum = "APPROVED_FOR_FUTURE"
	ListDelegatedResourceAccessRequestsRequestStatusRejected                  ListDelegatedResourceAccessRequestsRequestStatusEnum = "REJECTED"
	ListDelegatedResourceAccessRequestsRequestStatusDeployed                  ListDelegatedResourceAccessRequestsRequestStatusEnum = "DEPLOYED"
	ListDelegatedResourceAccessRequestsRequestStatusDeployFailed              ListDelegatedResourceAccessRequestsRequestStatusEnum = "DEPLOY_FAILED"
	ListDelegatedResourceAccessRequestsRequestStatusUndeployed                ListDelegatedResourceAccessRequestsRequestStatusEnum = "UNDEPLOYED"
	ListDelegatedResourceAccessRequestsRequestStatusUndeployFailed            ListDelegatedResourceAccessRequestsRequestStatusEnum = "UNDEPLOY_FAILED"
	ListDelegatedResourceAccessRequestsRequestStatusCloseFailed               ListDelegatedResourceAccessRequestsRequestStatusEnum = "CLOSE_FAILED"
	ListDelegatedResourceAccessRequestsRequestStatusRevokeFailed              ListDelegatedResourceAccessRequestsRequestStatusEnum = "REVOKE_FAILED"
	ListDelegatedResourceAccessRequestsRequestStatusExpiryFailed              ListDelegatedResourceAccessRequestsRequestStatusEnum = "EXPIRY_FAILED"
	ListDelegatedResourceAccessRequestsRequestStatusRevoking                  ListDelegatedResourceAccessRequestsRequestStatusEnum = "REVOKING"
	ListDelegatedResourceAccessRequestsRequestStatusRevoked                   ListDelegatedResourceAccessRequestsRequestStatusEnum = "REVOKED"
	ListDelegatedResourceAccessRequestsRequestStatusExtending                 ListDelegatedResourceAccessRequestsRequestStatusEnum = "EXTENDING"
	ListDelegatedResourceAccessRequestsRequestStatusExtended                  ListDelegatedResourceAccessRequestsRequestStatusEnum = "EXTENDED"
	ListDelegatedResourceAccessRequestsRequestStatusExtensionRejected         ListDelegatedResourceAccessRequestsRequestStatusEnum = "EXTENSION_REJECTED"
	ListDelegatedResourceAccessRequestsRequestStatusExtensionFailed           ListDelegatedResourceAccessRequestsRequestStatusEnum = "EXTENSION_FAILED"
	ListDelegatedResourceAccessRequestsRequestStatusCompleting                ListDelegatedResourceAccessRequestsRequestStatusEnum = "COMPLETING"
	ListDelegatedResourceAccessRequestsRequestStatusCompleted                 ListDelegatedResourceAccessRequestsRequestStatusEnum = "COMPLETED"
	ListDelegatedResourceAccessRequestsRequestStatusExpired                   ListDelegatedResourceAccessRequestsRequestStatusEnum = "EXPIRED"
)

var mappingListDelegatedResourceAccessRequestsRequestStatusEnum = map[string]ListDelegatedResourceAccessRequestsRequestStatusEnum{
	"CREATED":                     ListDelegatedResourceAccessRequestsRequestStatusCreated,
	"APPROVAL_WAITING":            ListDelegatedResourceAccessRequestsRequestStatusApprovalWaiting,
	"OPERATOR_ASSIGNMENT_WAITING": ListDelegatedResourceAccessRequestsRequestStatusOperatorAssignmentWaiting,
	"PREAPPROVED":                 ListDelegatedResourceAccessRequestsRequestStatusPreapproved,
	"APPROVED":                    ListDelegatedResourceAccessRequestsRequestStatusApproved,
	"APPROVED_FOR_FUTURE":         ListDelegatedResourceAccessRequestsRequestStatusApprovedForFuture,
	"REJECTED":                    ListDelegatedResourceAccessRequestsRequestStatusRejected,
	"DEPLOYED":                    ListDelegatedResourceAccessRequestsRequestStatusDeployed,
	"DEPLOY_FAILED":               ListDelegatedResourceAccessRequestsRequestStatusDeployFailed,
	"UNDEPLOYED":                  ListDelegatedResourceAccessRequestsRequestStatusUndeployed,
	"UNDEPLOY_FAILED":             ListDelegatedResourceAccessRequestsRequestStatusUndeployFailed,
	"CLOSE_FAILED":                ListDelegatedResourceAccessRequestsRequestStatusCloseFailed,
	"REVOKE_FAILED":               ListDelegatedResourceAccessRequestsRequestStatusRevokeFailed,
	"EXPIRY_FAILED":               ListDelegatedResourceAccessRequestsRequestStatusExpiryFailed,
	"REVOKING":                    ListDelegatedResourceAccessRequestsRequestStatusRevoking,
	"REVOKED":                     ListDelegatedResourceAccessRequestsRequestStatusRevoked,
	"EXTENDING":                   ListDelegatedResourceAccessRequestsRequestStatusExtending,
	"EXTENDED":                    ListDelegatedResourceAccessRequestsRequestStatusExtended,
	"EXTENSION_REJECTED":          ListDelegatedResourceAccessRequestsRequestStatusExtensionRejected,
	"EXTENSION_FAILED":            ListDelegatedResourceAccessRequestsRequestStatusExtensionFailed,
	"COMPLETING":                  ListDelegatedResourceAccessRequestsRequestStatusCompleting,
	"COMPLETED":                   ListDelegatedResourceAccessRequestsRequestStatusCompleted,
	"EXPIRED":                     ListDelegatedResourceAccessRequestsRequestStatusExpired,
}

var mappingListDelegatedResourceAccessRequestsRequestStatusEnumLowerCase = map[string]ListDelegatedResourceAccessRequestsRequestStatusEnum{
	"created":                     ListDelegatedResourceAccessRequestsRequestStatusCreated,
	"approval_waiting":            ListDelegatedResourceAccessRequestsRequestStatusApprovalWaiting,
	"operator_assignment_waiting": ListDelegatedResourceAccessRequestsRequestStatusOperatorAssignmentWaiting,
	"preapproved":                 ListDelegatedResourceAccessRequestsRequestStatusPreapproved,
	"approved":                    ListDelegatedResourceAccessRequestsRequestStatusApproved,
	"approved_for_future":         ListDelegatedResourceAccessRequestsRequestStatusApprovedForFuture,
	"rejected":                    ListDelegatedResourceAccessRequestsRequestStatusRejected,
	"deployed":                    ListDelegatedResourceAccessRequestsRequestStatusDeployed,
	"deploy_failed":               ListDelegatedResourceAccessRequestsRequestStatusDeployFailed,
	"undeployed":                  ListDelegatedResourceAccessRequestsRequestStatusUndeployed,
	"undeploy_failed":             ListDelegatedResourceAccessRequestsRequestStatusUndeployFailed,
	"close_failed":                ListDelegatedResourceAccessRequestsRequestStatusCloseFailed,
	"revoke_failed":               ListDelegatedResourceAccessRequestsRequestStatusRevokeFailed,
	"expiry_failed":               ListDelegatedResourceAccessRequestsRequestStatusExpiryFailed,
	"revoking":                    ListDelegatedResourceAccessRequestsRequestStatusRevoking,
	"revoked":                     ListDelegatedResourceAccessRequestsRequestStatusRevoked,
	"extending":                   ListDelegatedResourceAccessRequestsRequestStatusExtending,
	"extended":                    ListDelegatedResourceAccessRequestsRequestStatusExtended,
	"extension_rejected":          ListDelegatedResourceAccessRequestsRequestStatusExtensionRejected,
	"extension_failed":            ListDelegatedResourceAccessRequestsRequestStatusExtensionFailed,
	"completing":                  ListDelegatedResourceAccessRequestsRequestStatusCompleting,
	"completed":                   ListDelegatedResourceAccessRequestsRequestStatusCompleted,
	"expired":                     ListDelegatedResourceAccessRequestsRequestStatusExpired,
}

// GetListDelegatedResourceAccessRequestsRequestStatusEnumValues Enumerates the set of values for ListDelegatedResourceAccessRequestsRequestStatusEnum
func GetListDelegatedResourceAccessRequestsRequestStatusEnumValues() []ListDelegatedResourceAccessRequestsRequestStatusEnum {
	values := make([]ListDelegatedResourceAccessRequestsRequestStatusEnum, 0)
	for _, v := range mappingListDelegatedResourceAccessRequestsRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListDelegatedResourceAccessRequestsRequestStatusEnumStringValues Enumerates the set of values in String for ListDelegatedResourceAccessRequestsRequestStatusEnum
func GetListDelegatedResourceAccessRequestsRequestStatusEnumStringValues() []string {
	return []string{
		"CREATED",
		"APPROVAL_WAITING",
		"OPERATOR_ASSIGNMENT_WAITING",
		"PREAPPROVED",
		"APPROVED",
		"APPROVED_FOR_FUTURE",
		"REJECTED",
		"DEPLOYED",
		"DEPLOY_FAILED",
		"UNDEPLOYED",
		"UNDEPLOY_FAILED",
		"CLOSE_FAILED",
		"REVOKE_FAILED",
		"EXPIRY_FAILED",
		"REVOKING",
		"REVOKED",
		"EXTENDING",
		"EXTENDED",
		"EXTENSION_REJECTED",
		"EXTENSION_FAILED",
		"COMPLETING",
		"COMPLETED",
		"EXPIRED",
	}
}

// GetMappingListDelegatedResourceAccessRequestsRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDelegatedResourceAccessRequestsRequestStatusEnum(val string) (ListDelegatedResourceAccessRequestsRequestStatusEnum, bool) {
	enum, ok := mappingListDelegatedResourceAccessRequestsRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDelegatedResourceAccessRequestsSortOrderEnum Enum with underlying type: string
type ListDelegatedResourceAccessRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListDelegatedResourceAccessRequestsSortOrderEnum
const (
	ListDelegatedResourceAccessRequestsSortOrderAsc  ListDelegatedResourceAccessRequestsSortOrderEnum = "ASC"
	ListDelegatedResourceAccessRequestsSortOrderDesc ListDelegatedResourceAccessRequestsSortOrderEnum = "DESC"
)

var mappingListDelegatedResourceAccessRequestsSortOrderEnum = map[string]ListDelegatedResourceAccessRequestsSortOrderEnum{
	"ASC":  ListDelegatedResourceAccessRequestsSortOrderAsc,
	"DESC": ListDelegatedResourceAccessRequestsSortOrderDesc,
}

var mappingListDelegatedResourceAccessRequestsSortOrderEnumLowerCase = map[string]ListDelegatedResourceAccessRequestsSortOrderEnum{
	"asc":  ListDelegatedResourceAccessRequestsSortOrderAsc,
	"desc": ListDelegatedResourceAccessRequestsSortOrderDesc,
}

// GetListDelegatedResourceAccessRequestsSortOrderEnumValues Enumerates the set of values for ListDelegatedResourceAccessRequestsSortOrderEnum
func GetListDelegatedResourceAccessRequestsSortOrderEnumValues() []ListDelegatedResourceAccessRequestsSortOrderEnum {
	values := make([]ListDelegatedResourceAccessRequestsSortOrderEnum, 0)
	for _, v := range mappingListDelegatedResourceAccessRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDelegatedResourceAccessRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListDelegatedResourceAccessRequestsSortOrderEnum
func GetListDelegatedResourceAccessRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDelegatedResourceAccessRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDelegatedResourceAccessRequestsSortOrderEnum(val string) (ListDelegatedResourceAccessRequestsSortOrderEnum, bool) {
	enum, ok := mappingListDelegatedResourceAccessRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDelegatedResourceAccessRequestsSortByEnum Enum with underlying type: string
type ListDelegatedResourceAccessRequestsSortByEnum string

// Set of constants representing the allowable values for ListDelegatedResourceAccessRequestsSortByEnum
const (
	ListDelegatedResourceAccessRequestsSortByTimecreated ListDelegatedResourceAccessRequestsSortByEnum = "timeCreated"
	ListDelegatedResourceAccessRequestsSortByDisplayname ListDelegatedResourceAccessRequestsSortByEnum = "displayName"
)

var mappingListDelegatedResourceAccessRequestsSortByEnum = map[string]ListDelegatedResourceAccessRequestsSortByEnum{
	"timeCreated": ListDelegatedResourceAccessRequestsSortByTimecreated,
	"displayName": ListDelegatedResourceAccessRequestsSortByDisplayname,
}

var mappingListDelegatedResourceAccessRequestsSortByEnumLowerCase = map[string]ListDelegatedResourceAccessRequestsSortByEnum{
	"timecreated": ListDelegatedResourceAccessRequestsSortByTimecreated,
	"displayname": ListDelegatedResourceAccessRequestsSortByDisplayname,
}

// GetListDelegatedResourceAccessRequestsSortByEnumValues Enumerates the set of values for ListDelegatedResourceAccessRequestsSortByEnum
func GetListDelegatedResourceAccessRequestsSortByEnumValues() []ListDelegatedResourceAccessRequestsSortByEnum {
	values := make([]ListDelegatedResourceAccessRequestsSortByEnum, 0)
	for _, v := range mappingListDelegatedResourceAccessRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDelegatedResourceAccessRequestsSortByEnumStringValues Enumerates the set of values in String for ListDelegatedResourceAccessRequestsSortByEnum
func GetListDelegatedResourceAccessRequestsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDelegatedResourceAccessRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDelegatedResourceAccessRequestsSortByEnum(val string) (ListDelegatedResourceAccessRequestsSortByEnum, bool) {
	enum, ok := mappingListDelegatedResourceAccessRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
