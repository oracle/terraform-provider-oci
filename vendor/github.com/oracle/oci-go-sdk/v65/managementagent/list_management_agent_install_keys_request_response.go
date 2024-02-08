// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagementAgentInstallKeysRequest wrapper for the ListManagementAgentInstallKeys operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementagent/ListManagementAgentInstallKeys.go.html to see an example of how to use ListManagementAgentInstallKeysRequest.
type ListManagementAgentInstallKeysRequest struct {

	// The OCID of the compartment to which a request will be scoped.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// if set to true then it fetches resources for all compartments where user has access to else only on the compartment specified.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Value of this is always "ACCESSIBLE" and any other value is not supported.
	AccessLevel *string `mandatory:"false" contributesTo:"query" name:"accessLevel"`

	// Filter to return only Management Agents in the particular lifecycle state.
	LifecycleState ListManagementAgentInstallKeysLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The display name for which the Key needs to be listed.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagementAgentInstallKeysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListManagementAgentInstallKeysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementAgentInstallKeysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementAgentInstallKeysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagementAgentInstallKeysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagementAgentInstallKeysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagementAgentInstallKeysRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagementAgentInstallKeysLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListManagementAgentInstallKeysLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementAgentInstallKeysSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagementAgentInstallKeysSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementAgentInstallKeysSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagementAgentInstallKeysSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagementAgentInstallKeysResponse wrapper for the ListManagementAgentInstallKeys operation
type ListManagementAgentInstallKeysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ManagementAgentInstallKeySummary instances
	Items []ManagementAgentInstallKeySummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListManagementAgentInstallKeysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagementAgentInstallKeysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagementAgentInstallKeysLifecycleStateEnum Enum with underlying type: string
type ListManagementAgentInstallKeysLifecycleStateEnum string

// Set of constants representing the allowable values for ListManagementAgentInstallKeysLifecycleStateEnum
const (
	ListManagementAgentInstallKeysLifecycleStateCreating   ListManagementAgentInstallKeysLifecycleStateEnum = "CREATING"
	ListManagementAgentInstallKeysLifecycleStateUpdating   ListManagementAgentInstallKeysLifecycleStateEnum = "UPDATING"
	ListManagementAgentInstallKeysLifecycleStateActive     ListManagementAgentInstallKeysLifecycleStateEnum = "ACTIVE"
	ListManagementAgentInstallKeysLifecycleStateInactive   ListManagementAgentInstallKeysLifecycleStateEnum = "INACTIVE"
	ListManagementAgentInstallKeysLifecycleStateTerminated ListManagementAgentInstallKeysLifecycleStateEnum = "TERMINATED"
	ListManagementAgentInstallKeysLifecycleStateDeleting   ListManagementAgentInstallKeysLifecycleStateEnum = "DELETING"
	ListManagementAgentInstallKeysLifecycleStateDeleted    ListManagementAgentInstallKeysLifecycleStateEnum = "DELETED"
	ListManagementAgentInstallKeysLifecycleStateFailed     ListManagementAgentInstallKeysLifecycleStateEnum = "FAILED"
)

var mappingListManagementAgentInstallKeysLifecycleStateEnum = map[string]ListManagementAgentInstallKeysLifecycleStateEnum{
	"CREATING":   ListManagementAgentInstallKeysLifecycleStateCreating,
	"UPDATING":   ListManagementAgentInstallKeysLifecycleStateUpdating,
	"ACTIVE":     ListManagementAgentInstallKeysLifecycleStateActive,
	"INACTIVE":   ListManagementAgentInstallKeysLifecycleStateInactive,
	"TERMINATED": ListManagementAgentInstallKeysLifecycleStateTerminated,
	"DELETING":   ListManagementAgentInstallKeysLifecycleStateDeleting,
	"DELETED":    ListManagementAgentInstallKeysLifecycleStateDeleted,
	"FAILED":     ListManagementAgentInstallKeysLifecycleStateFailed,
}

var mappingListManagementAgentInstallKeysLifecycleStateEnumLowerCase = map[string]ListManagementAgentInstallKeysLifecycleStateEnum{
	"creating":   ListManagementAgentInstallKeysLifecycleStateCreating,
	"updating":   ListManagementAgentInstallKeysLifecycleStateUpdating,
	"active":     ListManagementAgentInstallKeysLifecycleStateActive,
	"inactive":   ListManagementAgentInstallKeysLifecycleStateInactive,
	"terminated": ListManagementAgentInstallKeysLifecycleStateTerminated,
	"deleting":   ListManagementAgentInstallKeysLifecycleStateDeleting,
	"deleted":    ListManagementAgentInstallKeysLifecycleStateDeleted,
	"failed":     ListManagementAgentInstallKeysLifecycleStateFailed,
}

// GetListManagementAgentInstallKeysLifecycleStateEnumValues Enumerates the set of values for ListManagementAgentInstallKeysLifecycleStateEnum
func GetListManagementAgentInstallKeysLifecycleStateEnumValues() []ListManagementAgentInstallKeysLifecycleStateEnum {
	values := make([]ListManagementAgentInstallKeysLifecycleStateEnum, 0)
	for _, v := range mappingListManagementAgentInstallKeysLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentInstallKeysLifecycleStateEnumStringValues Enumerates the set of values in String for ListManagementAgentInstallKeysLifecycleStateEnum
func GetListManagementAgentInstallKeysLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"TERMINATED",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListManagementAgentInstallKeysLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentInstallKeysLifecycleStateEnum(val string) (ListManagementAgentInstallKeysLifecycleStateEnum, bool) {
	enum, ok := mappingListManagementAgentInstallKeysLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementAgentInstallKeysSortOrderEnum Enum with underlying type: string
type ListManagementAgentInstallKeysSortOrderEnum string

// Set of constants representing the allowable values for ListManagementAgentInstallKeysSortOrderEnum
const (
	ListManagementAgentInstallKeysSortOrderAsc  ListManagementAgentInstallKeysSortOrderEnum = "ASC"
	ListManagementAgentInstallKeysSortOrderDesc ListManagementAgentInstallKeysSortOrderEnum = "DESC"
)

var mappingListManagementAgentInstallKeysSortOrderEnum = map[string]ListManagementAgentInstallKeysSortOrderEnum{
	"ASC":  ListManagementAgentInstallKeysSortOrderAsc,
	"DESC": ListManagementAgentInstallKeysSortOrderDesc,
}

var mappingListManagementAgentInstallKeysSortOrderEnumLowerCase = map[string]ListManagementAgentInstallKeysSortOrderEnum{
	"asc":  ListManagementAgentInstallKeysSortOrderAsc,
	"desc": ListManagementAgentInstallKeysSortOrderDesc,
}

// GetListManagementAgentInstallKeysSortOrderEnumValues Enumerates the set of values for ListManagementAgentInstallKeysSortOrderEnum
func GetListManagementAgentInstallKeysSortOrderEnumValues() []ListManagementAgentInstallKeysSortOrderEnum {
	values := make([]ListManagementAgentInstallKeysSortOrderEnum, 0)
	for _, v := range mappingListManagementAgentInstallKeysSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentInstallKeysSortOrderEnumStringValues Enumerates the set of values in String for ListManagementAgentInstallKeysSortOrderEnum
func GetListManagementAgentInstallKeysSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagementAgentInstallKeysSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentInstallKeysSortOrderEnum(val string) (ListManagementAgentInstallKeysSortOrderEnum, bool) {
	enum, ok := mappingListManagementAgentInstallKeysSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementAgentInstallKeysSortByEnum Enum with underlying type: string
type ListManagementAgentInstallKeysSortByEnum string

// Set of constants representing the allowable values for ListManagementAgentInstallKeysSortByEnum
const (
	ListManagementAgentInstallKeysSortByTimecreated ListManagementAgentInstallKeysSortByEnum = "timeCreated"
	ListManagementAgentInstallKeysSortByDisplayname ListManagementAgentInstallKeysSortByEnum = "displayName"
)

var mappingListManagementAgentInstallKeysSortByEnum = map[string]ListManagementAgentInstallKeysSortByEnum{
	"timeCreated": ListManagementAgentInstallKeysSortByTimecreated,
	"displayName": ListManagementAgentInstallKeysSortByDisplayname,
}

var mappingListManagementAgentInstallKeysSortByEnumLowerCase = map[string]ListManagementAgentInstallKeysSortByEnum{
	"timecreated": ListManagementAgentInstallKeysSortByTimecreated,
	"displayname": ListManagementAgentInstallKeysSortByDisplayname,
}

// GetListManagementAgentInstallKeysSortByEnumValues Enumerates the set of values for ListManagementAgentInstallKeysSortByEnum
func GetListManagementAgentInstallKeysSortByEnumValues() []ListManagementAgentInstallKeysSortByEnum {
	values := make([]ListManagementAgentInstallKeysSortByEnum, 0)
	for _, v := range mappingListManagementAgentInstallKeysSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementAgentInstallKeysSortByEnumStringValues Enumerates the set of values in String for ListManagementAgentInstallKeysSortByEnum
func GetListManagementAgentInstallKeysSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagementAgentInstallKeysSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementAgentInstallKeysSortByEnum(val string) (ListManagementAgentInstallKeysSortByEnum, bool) {
	enum, ok := mappingListManagementAgentInstallKeysSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
