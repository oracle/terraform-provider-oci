// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListModelDeploymentsRequest wrapper for the ListModelDeployments operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListModelDeployments.go.html to see an example of how to use ListModelDeploymentsRequest.
type ListModelDeploymentsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListModelDeploymentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListModelDeploymentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, results are shown
	// in descending order. When you sort by `displayName`, results are
	// shown in ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListModelDeploymentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModelDeploymentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModelDeploymentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModelDeploymentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModelDeploymentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListModelDeploymentsResponse wrapper for the ListModelDeployments operation
type ListModelDeploymentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ModelDeploymentSummary instances
	Items []ModelDeploymentSummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListModelDeploymentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModelDeploymentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModelDeploymentsLifecycleStateEnum Enum with underlying type: string
type ListModelDeploymentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListModelDeploymentsLifecycleStateEnum
const (
	ListModelDeploymentsLifecycleStateCreating       ListModelDeploymentsLifecycleStateEnum = "CREATING"
	ListModelDeploymentsLifecycleStateActive         ListModelDeploymentsLifecycleStateEnum = "ACTIVE"
	ListModelDeploymentsLifecycleStateDeleting       ListModelDeploymentsLifecycleStateEnum = "DELETING"
	ListModelDeploymentsLifecycleStateFailed         ListModelDeploymentsLifecycleStateEnum = "FAILED"
	ListModelDeploymentsLifecycleStateInactive       ListModelDeploymentsLifecycleStateEnum = "INACTIVE"
	ListModelDeploymentsLifecycleStateUpdating       ListModelDeploymentsLifecycleStateEnum = "UPDATING"
	ListModelDeploymentsLifecycleStateDeleted        ListModelDeploymentsLifecycleStateEnum = "DELETED"
	ListModelDeploymentsLifecycleStateNeedsAttention ListModelDeploymentsLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListModelDeploymentsLifecycleState = map[string]ListModelDeploymentsLifecycleStateEnum{
	"CREATING":        ListModelDeploymentsLifecycleStateCreating,
	"ACTIVE":          ListModelDeploymentsLifecycleStateActive,
	"DELETING":        ListModelDeploymentsLifecycleStateDeleting,
	"FAILED":          ListModelDeploymentsLifecycleStateFailed,
	"INACTIVE":        ListModelDeploymentsLifecycleStateInactive,
	"UPDATING":        ListModelDeploymentsLifecycleStateUpdating,
	"DELETED":         ListModelDeploymentsLifecycleStateDeleted,
	"NEEDS_ATTENTION": ListModelDeploymentsLifecycleStateNeedsAttention,
}

// GetListModelDeploymentsLifecycleStateEnumValues Enumerates the set of values for ListModelDeploymentsLifecycleStateEnum
func GetListModelDeploymentsLifecycleStateEnumValues() []ListModelDeploymentsLifecycleStateEnum {
	values := make([]ListModelDeploymentsLifecycleStateEnum, 0)
	for _, v := range mappingListModelDeploymentsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListModelDeploymentsSortOrderEnum Enum with underlying type: string
type ListModelDeploymentsSortOrderEnum string

// Set of constants representing the allowable values for ListModelDeploymentsSortOrderEnum
const (
	ListModelDeploymentsSortOrderAsc  ListModelDeploymentsSortOrderEnum = "ASC"
	ListModelDeploymentsSortOrderDesc ListModelDeploymentsSortOrderEnum = "DESC"
)

var mappingListModelDeploymentsSortOrder = map[string]ListModelDeploymentsSortOrderEnum{
	"ASC":  ListModelDeploymentsSortOrderAsc,
	"DESC": ListModelDeploymentsSortOrderDesc,
}

// GetListModelDeploymentsSortOrderEnumValues Enumerates the set of values for ListModelDeploymentsSortOrderEnum
func GetListModelDeploymentsSortOrderEnumValues() []ListModelDeploymentsSortOrderEnum {
	values := make([]ListModelDeploymentsSortOrderEnum, 0)
	for _, v := range mappingListModelDeploymentsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListModelDeploymentsSortByEnum Enum with underlying type: string
type ListModelDeploymentsSortByEnum string

// Set of constants representing the allowable values for ListModelDeploymentsSortByEnum
const (
	ListModelDeploymentsSortByTimecreated ListModelDeploymentsSortByEnum = "timeCreated"
	ListModelDeploymentsSortByDisplayname ListModelDeploymentsSortByEnum = "displayName"
)

var mappingListModelDeploymentsSortBy = map[string]ListModelDeploymentsSortByEnum{
	"timeCreated": ListModelDeploymentsSortByTimecreated,
	"displayName": ListModelDeploymentsSortByDisplayname,
}

// GetListModelDeploymentsSortByEnumValues Enumerates the set of values for ListModelDeploymentsSortByEnum
func GetListModelDeploymentsSortByEnumValues() []ListModelDeploymentsSortByEnum {
	values := make([]ListModelDeploymentsSortByEnum, 0)
	for _, v := range mappingListModelDeploymentsSortBy {
		values = append(values, v)
	}
	return values
}
