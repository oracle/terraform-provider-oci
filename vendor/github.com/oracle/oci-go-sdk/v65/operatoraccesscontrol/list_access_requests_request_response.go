// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package operatoraccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAccessRequestsRequest wrapper for the ListAccessRequests operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ListAccessRequests.go.html to see an example of how to use ListAccessRequestsRequest.
type ListAccessRequestsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given ResourceName.
	ResourceName *string `mandatory:"false" contributesTo:"query" name:"resourceName"`

	// A filter to return only lists of resources that match the entire given service type.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// A filter to return only resources whose lifecycleState matches the given AccessRequest lifecycleState.
	LifecycleState ListAccessRequestsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Query start time in UTC in ISO 8601 format(inclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// timeIntervalStart and timeIntervalEnd parameters are used together.
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// Query start time in UTC in ISO 8601 format(inclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// timeIntervalStart and timeIntervalEnd parameters are used together.
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAccessRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListAccessRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAccessRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAccessRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAccessRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAccessRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAccessRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAccessRequestsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAccessRequestsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAccessRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAccessRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAccessRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAccessRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAccessRequestsResponse wrapper for the ListAccessRequests operation
type ListAccessRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AccessRequestCollection instances
	AccessRequestCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAccessRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAccessRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAccessRequestsLifecycleStateEnum Enum with underlying type: string
type ListAccessRequestsLifecycleStateEnum string

// Set of constants representing the allowable values for ListAccessRequestsLifecycleStateEnum
const (
	ListAccessRequestsLifecycleStateCreated           ListAccessRequestsLifecycleStateEnum = "CREATED"
	ListAccessRequestsLifecycleStateApprovalwaiting   ListAccessRequestsLifecycleStateEnum = "APPROVALWAITING"
	ListAccessRequestsLifecycleStatePreapproved       ListAccessRequestsLifecycleStateEnum = "PREAPPROVED"
	ListAccessRequestsLifecycleStateApproved          ListAccessRequestsLifecycleStateEnum = "APPROVED"
	ListAccessRequestsLifecycleStateMoreinfo          ListAccessRequestsLifecycleStateEnum = "MOREINFO"
	ListAccessRequestsLifecycleStateRejected          ListAccessRequestsLifecycleStateEnum = "REJECTED"
	ListAccessRequestsLifecycleStateDeployed          ListAccessRequestsLifecycleStateEnum = "DEPLOYED"
	ListAccessRequestsLifecycleStateDeployfailed      ListAccessRequestsLifecycleStateEnum = "DEPLOYFAILED"
	ListAccessRequestsLifecycleStateUndeployed        ListAccessRequestsLifecycleStateEnum = "UNDEPLOYED"
	ListAccessRequestsLifecycleStateUndeployfailed    ListAccessRequestsLifecycleStateEnum = "UNDEPLOYFAILED"
	ListAccessRequestsLifecycleStateClosefailed       ListAccessRequestsLifecycleStateEnum = "CLOSEFAILED"
	ListAccessRequestsLifecycleStateRevokefailed      ListAccessRequestsLifecycleStateEnum = "REVOKEFAILED"
	ListAccessRequestsLifecycleStateExpiryfailed      ListAccessRequestsLifecycleStateEnum = "EXPIRYFAILED"
	ListAccessRequestsLifecycleStateRevoking          ListAccessRequestsLifecycleStateEnum = "REVOKING"
	ListAccessRequestsLifecycleStateRevoked           ListAccessRequestsLifecycleStateEnum = "REVOKED"
	ListAccessRequestsLifecycleStateExtending         ListAccessRequestsLifecycleStateEnum = "EXTENDING"
	ListAccessRequestsLifecycleStateExtended          ListAccessRequestsLifecycleStateEnum = "EXTENDED"
	ListAccessRequestsLifecycleStateExtensionrejected ListAccessRequestsLifecycleStateEnum = "EXTENSIONREJECTED"
	ListAccessRequestsLifecycleStateCompleting        ListAccessRequestsLifecycleStateEnum = "COMPLETING"
	ListAccessRequestsLifecycleStateCompleted         ListAccessRequestsLifecycleStateEnum = "COMPLETED"
	ListAccessRequestsLifecycleStateExpired           ListAccessRequestsLifecycleStateEnum = "EXPIRED"
	ListAccessRequestsLifecycleStateApprovedforfuture ListAccessRequestsLifecycleStateEnum = "APPROVEDFORFUTURE"
	ListAccessRequestsLifecycleStateInreview          ListAccessRequestsLifecycleStateEnum = "INREVIEW"
)

var mappingListAccessRequestsLifecycleStateEnum = map[string]ListAccessRequestsLifecycleStateEnum{
	"CREATED":           ListAccessRequestsLifecycleStateCreated,
	"APPROVALWAITING":   ListAccessRequestsLifecycleStateApprovalwaiting,
	"PREAPPROVED":       ListAccessRequestsLifecycleStatePreapproved,
	"APPROVED":          ListAccessRequestsLifecycleStateApproved,
	"MOREINFO":          ListAccessRequestsLifecycleStateMoreinfo,
	"REJECTED":          ListAccessRequestsLifecycleStateRejected,
	"DEPLOYED":          ListAccessRequestsLifecycleStateDeployed,
	"DEPLOYFAILED":      ListAccessRequestsLifecycleStateDeployfailed,
	"UNDEPLOYED":        ListAccessRequestsLifecycleStateUndeployed,
	"UNDEPLOYFAILED":    ListAccessRequestsLifecycleStateUndeployfailed,
	"CLOSEFAILED":       ListAccessRequestsLifecycleStateClosefailed,
	"REVOKEFAILED":      ListAccessRequestsLifecycleStateRevokefailed,
	"EXPIRYFAILED":      ListAccessRequestsLifecycleStateExpiryfailed,
	"REVOKING":          ListAccessRequestsLifecycleStateRevoking,
	"REVOKED":           ListAccessRequestsLifecycleStateRevoked,
	"EXTENDING":         ListAccessRequestsLifecycleStateExtending,
	"EXTENDED":          ListAccessRequestsLifecycleStateExtended,
	"EXTENSIONREJECTED": ListAccessRequestsLifecycleStateExtensionrejected,
	"COMPLETING":        ListAccessRequestsLifecycleStateCompleting,
	"COMPLETED":         ListAccessRequestsLifecycleStateCompleted,
	"EXPIRED":           ListAccessRequestsLifecycleStateExpired,
	"APPROVEDFORFUTURE": ListAccessRequestsLifecycleStateApprovedforfuture,
	"INREVIEW":          ListAccessRequestsLifecycleStateInreview,
}

var mappingListAccessRequestsLifecycleStateEnumLowerCase = map[string]ListAccessRequestsLifecycleStateEnum{
	"created":           ListAccessRequestsLifecycleStateCreated,
	"approvalwaiting":   ListAccessRequestsLifecycleStateApprovalwaiting,
	"preapproved":       ListAccessRequestsLifecycleStatePreapproved,
	"approved":          ListAccessRequestsLifecycleStateApproved,
	"moreinfo":          ListAccessRequestsLifecycleStateMoreinfo,
	"rejected":          ListAccessRequestsLifecycleStateRejected,
	"deployed":          ListAccessRequestsLifecycleStateDeployed,
	"deployfailed":      ListAccessRequestsLifecycleStateDeployfailed,
	"undeployed":        ListAccessRequestsLifecycleStateUndeployed,
	"undeployfailed":    ListAccessRequestsLifecycleStateUndeployfailed,
	"closefailed":       ListAccessRequestsLifecycleStateClosefailed,
	"revokefailed":      ListAccessRequestsLifecycleStateRevokefailed,
	"expiryfailed":      ListAccessRequestsLifecycleStateExpiryfailed,
	"revoking":          ListAccessRequestsLifecycleStateRevoking,
	"revoked":           ListAccessRequestsLifecycleStateRevoked,
	"extending":         ListAccessRequestsLifecycleStateExtending,
	"extended":          ListAccessRequestsLifecycleStateExtended,
	"extensionrejected": ListAccessRequestsLifecycleStateExtensionrejected,
	"completing":        ListAccessRequestsLifecycleStateCompleting,
	"completed":         ListAccessRequestsLifecycleStateCompleted,
	"expired":           ListAccessRequestsLifecycleStateExpired,
	"approvedforfuture": ListAccessRequestsLifecycleStateApprovedforfuture,
	"inreview":          ListAccessRequestsLifecycleStateInreview,
}

// GetListAccessRequestsLifecycleStateEnumValues Enumerates the set of values for ListAccessRequestsLifecycleStateEnum
func GetListAccessRequestsLifecycleStateEnumValues() []ListAccessRequestsLifecycleStateEnum {
	values := make([]ListAccessRequestsLifecycleStateEnum, 0)
	for _, v := range mappingListAccessRequestsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAccessRequestsLifecycleStateEnumStringValues Enumerates the set of values in String for ListAccessRequestsLifecycleStateEnum
func GetListAccessRequestsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATED",
		"APPROVALWAITING",
		"PREAPPROVED",
		"APPROVED",
		"MOREINFO",
		"REJECTED",
		"DEPLOYED",
		"DEPLOYFAILED",
		"UNDEPLOYED",
		"UNDEPLOYFAILED",
		"CLOSEFAILED",
		"REVOKEFAILED",
		"EXPIRYFAILED",
		"REVOKING",
		"REVOKED",
		"EXTENDING",
		"EXTENDED",
		"EXTENSIONREJECTED",
		"COMPLETING",
		"COMPLETED",
		"EXPIRED",
		"APPROVEDFORFUTURE",
		"INREVIEW",
	}
}

// GetMappingListAccessRequestsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAccessRequestsLifecycleStateEnum(val string) (ListAccessRequestsLifecycleStateEnum, bool) {
	enum, ok := mappingListAccessRequestsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAccessRequestsSortOrderEnum Enum with underlying type: string
type ListAccessRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListAccessRequestsSortOrderEnum
const (
	ListAccessRequestsSortOrderAsc  ListAccessRequestsSortOrderEnum = "ASC"
	ListAccessRequestsSortOrderDesc ListAccessRequestsSortOrderEnum = "DESC"
)

var mappingListAccessRequestsSortOrderEnum = map[string]ListAccessRequestsSortOrderEnum{
	"ASC":  ListAccessRequestsSortOrderAsc,
	"DESC": ListAccessRequestsSortOrderDesc,
}

var mappingListAccessRequestsSortOrderEnumLowerCase = map[string]ListAccessRequestsSortOrderEnum{
	"asc":  ListAccessRequestsSortOrderAsc,
	"desc": ListAccessRequestsSortOrderDesc,
}

// GetListAccessRequestsSortOrderEnumValues Enumerates the set of values for ListAccessRequestsSortOrderEnum
func GetListAccessRequestsSortOrderEnumValues() []ListAccessRequestsSortOrderEnum {
	values := make([]ListAccessRequestsSortOrderEnum, 0)
	for _, v := range mappingListAccessRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAccessRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListAccessRequestsSortOrderEnum
func GetListAccessRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAccessRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAccessRequestsSortOrderEnum(val string) (ListAccessRequestsSortOrderEnum, bool) {
	enum, ok := mappingListAccessRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAccessRequestsSortByEnum Enum with underlying type: string
type ListAccessRequestsSortByEnum string

// Set of constants representing the allowable values for ListAccessRequestsSortByEnum
const (
	ListAccessRequestsSortByTimecreated ListAccessRequestsSortByEnum = "timeCreated"
	ListAccessRequestsSortByDisplayname ListAccessRequestsSortByEnum = "displayName"
)

var mappingListAccessRequestsSortByEnum = map[string]ListAccessRequestsSortByEnum{
	"timeCreated": ListAccessRequestsSortByTimecreated,
	"displayName": ListAccessRequestsSortByDisplayname,
}

var mappingListAccessRequestsSortByEnumLowerCase = map[string]ListAccessRequestsSortByEnum{
	"timecreated": ListAccessRequestsSortByTimecreated,
	"displayname": ListAccessRequestsSortByDisplayname,
}

// GetListAccessRequestsSortByEnumValues Enumerates the set of values for ListAccessRequestsSortByEnum
func GetListAccessRequestsSortByEnumValues() []ListAccessRequestsSortByEnum {
	values := make([]ListAccessRequestsSortByEnum, 0)
	for _, v := range mappingListAccessRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAccessRequestsSortByEnumStringValues Enumerates the set of values in String for ListAccessRequestsSortByEnum
func GetListAccessRequestsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAccessRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAccessRequestsSortByEnum(val string) (ListAccessRequestsSortByEnum, bool) {
	enum, ok := mappingListAccessRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
