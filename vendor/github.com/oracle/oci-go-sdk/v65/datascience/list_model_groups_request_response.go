// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListModelGroupsRequest wrapper for the ListModelGroups operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListModelGroups.go.html to see an example of how to use ListModelGroupsRequest.
type ListModelGroupsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// <b>Filter</b> results by OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources matching the given lifecycleState.
	LifecycleState ListModelGroupsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the modelGroupVersionHistory.
	ModelGroupVersionHistoryId *string `mandatory:"false" contributesTo:"query" name:"modelGroupVersionHistoryId"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 100 is the maximum.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListModelGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, the results are shown
	// in descending order. All other fields default to ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListModelGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModelGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModelGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModelGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModelGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListModelGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListModelGroupsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListModelGroupsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModelGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListModelGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModelGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListModelGroupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListModelGroupsResponse wrapper for the ListModelGroups operation
type ListModelGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ModelGroupSummary instances
	Items []ModelGroupSummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListModelGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModelGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModelGroupsLifecycleStateEnum Enum with underlying type: string
type ListModelGroupsLifecycleStateEnum string

// Set of constants representing the allowable values for ListModelGroupsLifecycleStateEnum
const (
	ListModelGroupsLifecycleStateCreating ListModelGroupsLifecycleStateEnum = "CREATING"
	ListModelGroupsLifecycleStateActive   ListModelGroupsLifecycleStateEnum = "ACTIVE"
	ListModelGroupsLifecycleStateFailed   ListModelGroupsLifecycleStateEnum = "FAILED"
	ListModelGroupsLifecycleStateInactive ListModelGroupsLifecycleStateEnum = "INACTIVE"
	ListModelGroupsLifecycleStateDeleting ListModelGroupsLifecycleStateEnum = "DELETING"
	ListModelGroupsLifecycleStateDeleted  ListModelGroupsLifecycleStateEnum = "DELETED"
)

var mappingListModelGroupsLifecycleStateEnum = map[string]ListModelGroupsLifecycleStateEnum{
	"CREATING": ListModelGroupsLifecycleStateCreating,
	"ACTIVE":   ListModelGroupsLifecycleStateActive,
	"FAILED":   ListModelGroupsLifecycleStateFailed,
	"INACTIVE": ListModelGroupsLifecycleStateInactive,
	"DELETING": ListModelGroupsLifecycleStateDeleting,
	"DELETED":  ListModelGroupsLifecycleStateDeleted,
}

var mappingListModelGroupsLifecycleStateEnumLowerCase = map[string]ListModelGroupsLifecycleStateEnum{
	"creating": ListModelGroupsLifecycleStateCreating,
	"active":   ListModelGroupsLifecycleStateActive,
	"failed":   ListModelGroupsLifecycleStateFailed,
	"inactive": ListModelGroupsLifecycleStateInactive,
	"deleting": ListModelGroupsLifecycleStateDeleting,
	"deleted":  ListModelGroupsLifecycleStateDeleted,
}

// GetListModelGroupsLifecycleStateEnumValues Enumerates the set of values for ListModelGroupsLifecycleStateEnum
func GetListModelGroupsLifecycleStateEnumValues() []ListModelGroupsLifecycleStateEnum {
	values := make([]ListModelGroupsLifecycleStateEnum, 0)
	for _, v := range mappingListModelGroupsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelGroupsLifecycleStateEnumStringValues Enumerates the set of values in String for ListModelGroupsLifecycleStateEnum
func GetListModelGroupsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"FAILED",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingListModelGroupsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelGroupsLifecycleStateEnum(val string) (ListModelGroupsLifecycleStateEnum, bool) {
	enum, ok := mappingListModelGroupsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModelGroupsSortOrderEnum Enum with underlying type: string
type ListModelGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListModelGroupsSortOrderEnum
const (
	ListModelGroupsSortOrderAsc  ListModelGroupsSortOrderEnum = "ASC"
	ListModelGroupsSortOrderDesc ListModelGroupsSortOrderEnum = "DESC"
)

var mappingListModelGroupsSortOrderEnum = map[string]ListModelGroupsSortOrderEnum{
	"ASC":  ListModelGroupsSortOrderAsc,
	"DESC": ListModelGroupsSortOrderDesc,
}

var mappingListModelGroupsSortOrderEnumLowerCase = map[string]ListModelGroupsSortOrderEnum{
	"asc":  ListModelGroupsSortOrderAsc,
	"desc": ListModelGroupsSortOrderDesc,
}

// GetListModelGroupsSortOrderEnumValues Enumerates the set of values for ListModelGroupsSortOrderEnum
func GetListModelGroupsSortOrderEnumValues() []ListModelGroupsSortOrderEnum {
	values := make([]ListModelGroupsSortOrderEnum, 0)
	for _, v := range mappingListModelGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListModelGroupsSortOrderEnum
func GetListModelGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListModelGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelGroupsSortOrderEnum(val string) (ListModelGroupsSortOrderEnum, bool) {
	enum, ok := mappingListModelGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModelGroupsSortByEnum Enum with underlying type: string
type ListModelGroupsSortByEnum string

// Set of constants representing the allowable values for ListModelGroupsSortByEnum
const (
	ListModelGroupsSortByTimecreated    ListModelGroupsSortByEnum = "timeCreated"
	ListModelGroupsSortByDisplayname    ListModelGroupsSortByEnum = "displayName"
	ListModelGroupsSortByLifecyclestate ListModelGroupsSortByEnum = "lifecycleState"
)

var mappingListModelGroupsSortByEnum = map[string]ListModelGroupsSortByEnum{
	"timeCreated":    ListModelGroupsSortByTimecreated,
	"displayName":    ListModelGroupsSortByDisplayname,
	"lifecycleState": ListModelGroupsSortByLifecyclestate,
}

var mappingListModelGroupsSortByEnumLowerCase = map[string]ListModelGroupsSortByEnum{
	"timecreated":    ListModelGroupsSortByTimecreated,
	"displayname":    ListModelGroupsSortByDisplayname,
	"lifecyclestate": ListModelGroupsSortByLifecyclestate,
}

// GetListModelGroupsSortByEnumValues Enumerates the set of values for ListModelGroupsSortByEnum
func GetListModelGroupsSortByEnumValues() []ListModelGroupsSortByEnum {
	values := make([]ListModelGroupsSortByEnum, 0)
	for _, v := range mappingListModelGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelGroupsSortByEnumStringValues Enumerates the set of values in String for ListModelGroupsSortByEnum
func GetListModelGroupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"lifecycleState",
	}
}

// GetMappingListModelGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelGroupsSortByEnum(val string) (ListModelGroupsSortByEnum, bool) {
	enum, ok := mappingListModelGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
