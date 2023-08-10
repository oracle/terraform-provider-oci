// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListComputeTargetsRequest wrapper for the ListComputeTargets operation
type ListComputeTargetsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// <b>Filter</b> results by OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	//   state for the resource type.
	LifecycleState ListComputeTargetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 100 is the maximum.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListComputeTargetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, the results are shown
	// in descending order. When you sort by `displayName`, the results are
	// shown in ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListComputeTargetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListComputeTargetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListComputeTargetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListComputeTargetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListComputeTargetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListComputeTargetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListComputeTargetsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListComputeTargetsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListComputeTargetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListComputeTargetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListComputeTargetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListComputeTargetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListComputeTargetsResponse wrapper for the ListComputeTargets operation
type ListComputeTargetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ComputeTargetSummary instances
	Items []ComputeTargetSummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListComputeTargetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListComputeTargetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListComputeTargetsLifecycleStateEnum Enum with underlying type: string
type ListComputeTargetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListComputeTargetsLifecycleStateEnum
const (
	ListComputeTargetsLifecycleStateCreating ListComputeTargetsLifecycleStateEnum = "CREATING"
	ListComputeTargetsLifecycleStateActive   ListComputeTargetsLifecycleStateEnum = "ACTIVE"
	ListComputeTargetsLifecycleStateDeleting ListComputeTargetsLifecycleStateEnum = "DELETING"
	ListComputeTargetsLifecycleStateDeleted  ListComputeTargetsLifecycleStateEnum = "DELETED"
	ListComputeTargetsLifecycleStateFailed   ListComputeTargetsLifecycleStateEnum = "FAILED"
	ListComputeTargetsLifecycleStateUpdating ListComputeTargetsLifecycleStateEnum = "UPDATING"
)

var mappingListComputeTargetsLifecycleStateEnum = map[string]ListComputeTargetsLifecycleStateEnum{
	"CREATING": ListComputeTargetsLifecycleStateCreating,
	"ACTIVE":   ListComputeTargetsLifecycleStateActive,
	"DELETING": ListComputeTargetsLifecycleStateDeleting,
	"DELETED":  ListComputeTargetsLifecycleStateDeleted,
	"FAILED":   ListComputeTargetsLifecycleStateFailed,
	"UPDATING": ListComputeTargetsLifecycleStateUpdating,
}

var mappingListComputeTargetsLifecycleStateEnumLowerCase = map[string]ListComputeTargetsLifecycleStateEnum{
	"creating": ListComputeTargetsLifecycleStateCreating,
	"active":   ListComputeTargetsLifecycleStateActive,
	"deleting": ListComputeTargetsLifecycleStateDeleting,
	"deleted":  ListComputeTargetsLifecycleStateDeleted,
	"failed":   ListComputeTargetsLifecycleStateFailed,
	"updating": ListComputeTargetsLifecycleStateUpdating,
}

// GetListComputeTargetsLifecycleStateEnumValues Enumerates the set of values for ListComputeTargetsLifecycleStateEnum
func GetListComputeTargetsLifecycleStateEnumValues() []ListComputeTargetsLifecycleStateEnum {
	values := make([]ListComputeTargetsLifecycleStateEnum, 0)
	for _, v := range mappingListComputeTargetsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListComputeTargetsLifecycleStateEnumStringValues Enumerates the set of values in String for ListComputeTargetsLifecycleStateEnum
func GetListComputeTargetsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingListComputeTargetsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListComputeTargetsLifecycleStateEnum(val string) (ListComputeTargetsLifecycleStateEnum, bool) {
	enum, ok := mappingListComputeTargetsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListComputeTargetsSortOrderEnum Enum with underlying type: string
type ListComputeTargetsSortOrderEnum string

// Set of constants representing the allowable values for ListComputeTargetsSortOrderEnum
const (
	ListComputeTargetsSortOrderAsc  ListComputeTargetsSortOrderEnum = "ASC"
	ListComputeTargetsSortOrderDesc ListComputeTargetsSortOrderEnum = "DESC"
)

var mappingListComputeTargetsSortOrderEnum = map[string]ListComputeTargetsSortOrderEnum{
	"ASC":  ListComputeTargetsSortOrderAsc,
	"DESC": ListComputeTargetsSortOrderDesc,
}

var mappingListComputeTargetsSortOrderEnumLowerCase = map[string]ListComputeTargetsSortOrderEnum{
	"asc":  ListComputeTargetsSortOrderAsc,
	"desc": ListComputeTargetsSortOrderDesc,
}

// GetListComputeTargetsSortOrderEnumValues Enumerates the set of values for ListComputeTargetsSortOrderEnum
func GetListComputeTargetsSortOrderEnumValues() []ListComputeTargetsSortOrderEnum {
	values := make([]ListComputeTargetsSortOrderEnum, 0)
	for _, v := range mappingListComputeTargetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListComputeTargetsSortOrderEnumStringValues Enumerates the set of values in String for ListComputeTargetsSortOrderEnum
func GetListComputeTargetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListComputeTargetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListComputeTargetsSortOrderEnum(val string) (ListComputeTargetsSortOrderEnum, bool) {
	enum, ok := mappingListComputeTargetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListComputeTargetsSortByEnum Enum with underlying type: string
type ListComputeTargetsSortByEnum string

// Set of constants representing the allowable values for ListComputeTargetsSortByEnum
const (
	ListComputeTargetsSortByTimecreated ListComputeTargetsSortByEnum = "timeCreated"
	ListComputeTargetsSortByDisplayname ListComputeTargetsSortByEnum = "displayName"
)

var mappingListComputeTargetsSortByEnum = map[string]ListComputeTargetsSortByEnum{
	"timeCreated": ListComputeTargetsSortByTimecreated,
	"displayName": ListComputeTargetsSortByDisplayname,
}

var mappingListComputeTargetsSortByEnumLowerCase = map[string]ListComputeTargetsSortByEnum{
	"timecreated": ListComputeTargetsSortByTimecreated,
	"displayname": ListComputeTargetsSortByDisplayname,
}

// GetListComputeTargetsSortByEnumValues Enumerates the set of values for ListComputeTargetsSortByEnum
func GetListComputeTargetsSortByEnumValues() []ListComputeTargetsSortByEnum {
	values := make([]ListComputeTargetsSortByEnum, 0)
	for _, v := range mappingListComputeTargetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListComputeTargetsSortByEnumStringValues Enumerates the set of values in String for ListComputeTargetsSortByEnum
func GetListComputeTargetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListComputeTargetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListComputeTargetsSortByEnum(val string) (ListComputeTargetsSortByEnum, bool) {
	enum, ok := mappingListComputeTargetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
