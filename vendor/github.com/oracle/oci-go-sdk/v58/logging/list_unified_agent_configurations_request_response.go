// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListUnifiedAgentConfigurationsRequest wrapper for the ListUnifiedAgentConfigurations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/logging/ListUnifiedAgentConfigurations.go.html to see an example of how to use ListUnifiedAgentConfigurationsRequest.
type ListUnifiedAgentConfigurationsRequest struct {

	// Compartment OCID to list resources in. See compartmentIdInSubtree
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

	// The sort order to use, whether 'asc' or 'desc'.
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
func (request ListUnifiedAgentConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUnifiedAgentConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUnifiedAgentConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUnifiedAgentConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUnifiedAgentConfigurationsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListUnifiedAgentConfigurationsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUnifiedAgentConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUnifiedAgentConfigurationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUnifiedAgentConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUnifiedAgentConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListUnifiedAgentConfigurationsLifecycleStateEnum = map[string]ListUnifiedAgentConfigurationsLifecycleStateEnum{
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
	for _, v := range mappingListUnifiedAgentConfigurationsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListUnifiedAgentConfigurationsLifecycleStateEnumStringValues Enumerates the set of values in String for ListUnifiedAgentConfigurationsLifecycleStateEnum
func GetListUnifiedAgentConfigurationsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"INACTIVE",
		"DELETING",
		"FAILED",
	}
}

// GetMappingListUnifiedAgentConfigurationsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUnifiedAgentConfigurationsLifecycleStateEnum(val string) (ListUnifiedAgentConfigurationsLifecycleStateEnum, bool) {
	mappingListUnifiedAgentConfigurationsLifecycleStateEnumIgnoreCase := make(map[string]ListUnifiedAgentConfigurationsLifecycleStateEnum)
	for k, v := range mappingListUnifiedAgentConfigurationsLifecycleStateEnum {
		mappingListUnifiedAgentConfigurationsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListUnifiedAgentConfigurationsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListUnifiedAgentConfigurationsSortByEnum Enum with underlying type: string
type ListUnifiedAgentConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListUnifiedAgentConfigurationsSortByEnum
const (
	ListUnifiedAgentConfigurationsSortByTimecreated ListUnifiedAgentConfigurationsSortByEnum = "timeCreated"
	ListUnifiedAgentConfigurationsSortByDisplayname ListUnifiedAgentConfigurationsSortByEnum = "displayName"
)

var mappingListUnifiedAgentConfigurationsSortByEnum = map[string]ListUnifiedAgentConfigurationsSortByEnum{
	"timeCreated": ListUnifiedAgentConfigurationsSortByTimecreated,
	"displayName": ListUnifiedAgentConfigurationsSortByDisplayname,
}

// GetListUnifiedAgentConfigurationsSortByEnumValues Enumerates the set of values for ListUnifiedAgentConfigurationsSortByEnum
func GetListUnifiedAgentConfigurationsSortByEnumValues() []ListUnifiedAgentConfigurationsSortByEnum {
	values := make([]ListUnifiedAgentConfigurationsSortByEnum, 0)
	for _, v := range mappingListUnifiedAgentConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUnifiedAgentConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListUnifiedAgentConfigurationsSortByEnum
func GetListUnifiedAgentConfigurationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListUnifiedAgentConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUnifiedAgentConfigurationsSortByEnum(val string) (ListUnifiedAgentConfigurationsSortByEnum, bool) {
	mappingListUnifiedAgentConfigurationsSortByEnumIgnoreCase := make(map[string]ListUnifiedAgentConfigurationsSortByEnum)
	for k, v := range mappingListUnifiedAgentConfigurationsSortByEnum {
		mappingListUnifiedAgentConfigurationsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListUnifiedAgentConfigurationsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListUnifiedAgentConfigurationsSortOrderEnum Enum with underlying type: string
type ListUnifiedAgentConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListUnifiedAgentConfigurationsSortOrderEnum
const (
	ListUnifiedAgentConfigurationsSortOrderAsc  ListUnifiedAgentConfigurationsSortOrderEnum = "ASC"
	ListUnifiedAgentConfigurationsSortOrderDesc ListUnifiedAgentConfigurationsSortOrderEnum = "DESC"
)

var mappingListUnifiedAgentConfigurationsSortOrderEnum = map[string]ListUnifiedAgentConfigurationsSortOrderEnum{
	"ASC":  ListUnifiedAgentConfigurationsSortOrderAsc,
	"DESC": ListUnifiedAgentConfigurationsSortOrderDesc,
}

// GetListUnifiedAgentConfigurationsSortOrderEnumValues Enumerates the set of values for ListUnifiedAgentConfigurationsSortOrderEnum
func GetListUnifiedAgentConfigurationsSortOrderEnumValues() []ListUnifiedAgentConfigurationsSortOrderEnum {
	values := make([]ListUnifiedAgentConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListUnifiedAgentConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUnifiedAgentConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListUnifiedAgentConfigurationsSortOrderEnum
func GetListUnifiedAgentConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUnifiedAgentConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUnifiedAgentConfigurationsSortOrderEnum(val string) (ListUnifiedAgentConfigurationsSortOrderEnum, bool) {
	mappingListUnifiedAgentConfigurationsSortOrderEnumIgnoreCase := make(map[string]ListUnifiedAgentConfigurationsSortOrderEnum)
	for k, v := range mappingListUnifiedAgentConfigurationsSortOrderEnum {
		mappingListUnifiedAgentConfigurationsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListUnifiedAgentConfigurationsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
