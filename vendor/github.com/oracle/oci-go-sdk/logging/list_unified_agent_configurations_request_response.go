// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package logging

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListUnifiedAgentConfigurationsRequest wrapper for the ListUnifiedAgentConfigurations operation
type ListUnifiedAgentConfigurationsRequest struct {

	// Compartment OCID to list resources in. Please see compartmentIdInSubtree
	//      for nested compartments traversal.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Custom log OCID to list resources with the log as destination.
	LogId *string `mandatory:"false" contributesTo:"query" name:"logId"`

	// Specifies whether or not nested compartments should be traversed. Defaults to false.
	IsCompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"isCompartmentIdInSubtree"`

	// The OCID of a group or a dynamic group.
	GroupId *string `mandatory:"false" contributesTo:"query" name:"groupId"`

	// Resource name
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Lifecycle state of the log object
	LifecycleState ListUnifiedAgentConfigurationsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` or `opc-previous-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by (one column only). Default sort order is
	// ascending exception of `timeCreated` and `timeLastModified` columns (descending).
	SortBy ListUnifiedAgentConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'
	SortOrder ListUnifiedAgentConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUnifiedAgentConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUnifiedAgentConfigurationsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUnifiedAgentConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListUnifiedAgentConfigurationsResponse wrapper for the ListUnifiedAgentConfigurations operation
type ListUnifiedAgentConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UnifiedAgentConfigurationCollection instances
	UnifiedAgentConfigurationCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages
	// of results exist. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListUnifiedAgentConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUnifiedAgentConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUnifiedAgentConfigurationsLifecycleStateEnum Enum with underlying type: string
type ListUnifiedAgentConfigurationsLifecycleStateEnum string

// Set of constants representing the allowable values for ListUnifiedAgentConfigurationsLifecycleStateEnum
const (
	ListUnifiedAgentConfigurationsLifecycleStateCreating ListUnifiedAgentConfigurationsLifecycleStateEnum = "CREATING"
	ListUnifiedAgentConfigurationsLifecycleStateActive   ListUnifiedAgentConfigurationsLifecycleStateEnum = "ACTIVE"
	ListUnifiedAgentConfigurationsLifecycleStateUpdating ListUnifiedAgentConfigurationsLifecycleStateEnum = "UPDATING"
	ListUnifiedAgentConfigurationsLifecycleStateInactive ListUnifiedAgentConfigurationsLifecycleStateEnum = "INACTIVE"
	ListUnifiedAgentConfigurationsLifecycleStateDeleting ListUnifiedAgentConfigurationsLifecycleStateEnum = "DELETING"
	ListUnifiedAgentConfigurationsLifecycleStateFailed   ListUnifiedAgentConfigurationsLifecycleStateEnum = "FAILED"
)

var mappingListUnifiedAgentConfigurationsLifecycleState = map[string]ListUnifiedAgentConfigurationsLifecycleStateEnum{
	"CREATING": ListUnifiedAgentConfigurationsLifecycleStateCreating,
	"ACTIVE":   ListUnifiedAgentConfigurationsLifecycleStateActive,
	"UPDATING": ListUnifiedAgentConfigurationsLifecycleStateUpdating,
	"INACTIVE": ListUnifiedAgentConfigurationsLifecycleStateInactive,
	"DELETING": ListUnifiedAgentConfigurationsLifecycleStateDeleting,
	"FAILED":   ListUnifiedAgentConfigurationsLifecycleStateFailed,
}

// GetListUnifiedAgentConfigurationsLifecycleStateEnumValues Enumerates the set of values for ListUnifiedAgentConfigurationsLifecycleStateEnum
func GetListUnifiedAgentConfigurationsLifecycleStateEnumValues() []ListUnifiedAgentConfigurationsLifecycleStateEnum {
	values := make([]ListUnifiedAgentConfigurationsLifecycleStateEnum, 0)
	for _, v := range mappingListUnifiedAgentConfigurationsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListUnifiedAgentConfigurationsSortByEnum Enum with underlying type: string
type ListUnifiedAgentConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListUnifiedAgentConfigurationsSortByEnum
const (
	ListUnifiedAgentConfigurationsSortByTimecreated ListUnifiedAgentConfigurationsSortByEnum = "timeCreated"
	ListUnifiedAgentConfigurationsSortByDisplayname ListUnifiedAgentConfigurationsSortByEnum = "displayName"
)

var mappingListUnifiedAgentConfigurationsSortBy = map[string]ListUnifiedAgentConfigurationsSortByEnum{
	"timeCreated": ListUnifiedAgentConfigurationsSortByTimecreated,
	"displayName": ListUnifiedAgentConfigurationsSortByDisplayname,
}

// GetListUnifiedAgentConfigurationsSortByEnumValues Enumerates the set of values for ListUnifiedAgentConfigurationsSortByEnum
func GetListUnifiedAgentConfigurationsSortByEnumValues() []ListUnifiedAgentConfigurationsSortByEnum {
	values := make([]ListUnifiedAgentConfigurationsSortByEnum, 0)
	for _, v := range mappingListUnifiedAgentConfigurationsSortBy {
		values = append(values, v)
	}
	return values
}

// ListUnifiedAgentConfigurationsSortOrderEnum Enum with underlying type: string
type ListUnifiedAgentConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListUnifiedAgentConfigurationsSortOrderEnum
const (
	ListUnifiedAgentConfigurationsSortOrderAsc  ListUnifiedAgentConfigurationsSortOrderEnum = "ASC"
	ListUnifiedAgentConfigurationsSortOrderDesc ListUnifiedAgentConfigurationsSortOrderEnum = "DESC"
)

var mappingListUnifiedAgentConfigurationsSortOrder = map[string]ListUnifiedAgentConfigurationsSortOrderEnum{
	"ASC":  ListUnifiedAgentConfigurationsSortOrderAsc,
	"DESC": ListUnifiedAgentConfigurationsSortOrderDesc,
}

// GetListUnifiedAgentConfigurationsSortOrderEnumValues Enumerates the set of values for ListUnifiedAgentConfigurationsSortOrderEnum
func GetListUnifiedAgentConfigurationsSortOrderEnumValues() []ListUnifiedAgentConfigurationsSortOrderEnum {
	values := make([]ListUnifiedAgentConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListUnifiedAgentConfigurationsSortOrder {
		values = append(values, v)
	}
	return values
}
