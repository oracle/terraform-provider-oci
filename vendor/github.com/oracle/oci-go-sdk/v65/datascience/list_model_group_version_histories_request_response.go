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

// ListModelGroupVersionHistoriesRequest wrapper for the ListModelGroupVersionHistories operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListModelGroupVersionHistories.go.html to see an example of how to use ListModelGroupVersionHistoriesRequest.
type ListModelGroupVersionHistoriesRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// <b>Filter</b> results by OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources matching the given lifecycleState.
	LifecycleState ListModelGroupVersionHistoriesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

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
	SortOrder ListModelGroupVersionHistoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, the results are shown
	// in descending order. All other fields default to ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListModelGroupVersionHistoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModelGroupVersionHistoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModelGroupVersionHistoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModelGroupVersionHistoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModelGroupVersionHistoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListModelGroupVersionHistoriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListModelGroupVersionHistoriesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListModelGroupVersionHistoriesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModelGroupVersionHistoriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListModelGroupVersionHistoriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModelGroupVersionHistoriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListModelGroupVersionHistoriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListModelGroupVersionHistoriesResponse wrapper for the ListModelGroupVersionHistories operation
type ListModelGroupVersionHistoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ModelGroupVersionHistorySummary instances
	Items []ModelGroupVersionHistorySummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListModelGroupVersionHistoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModelGroupVersionHistoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModelGroupVersionHistoriesLifecycleStateEnum Enum with underlying type: string
type ListModelGroupVersionHistoriesLifecycleStateEnum string

// Set of constants representing the allowable values for ListModelGroupVersionHistoriesLifecycleStateEnum
const (
	ListModelGroupVersionHistoriesLifecycleStateActive   ListModelGroupVersionHistoriesLifecycleStateEnum = "ACTIVE"
	ListModelGroupVersionHistoriesLifecycleStateDeleted  ListModelGroupVersionHistoriesLifecycleStateEnum = "DELETED"
	ListModelGroupVersionHistoriesLifecycleStateFailed   ListModelGroupVersionHistoriesLifecycleStateEnum = "FAILED"
	ListModelGroupVersionHistoriesLifecycleStateDeleting ListModelGroupVersionHistoriesLifecycleStateEnum = "DELETING"
)

var mappingListModelGroupVersionHistoriesLifecycleStateEnum = map[string]ListModelGroupVersionHistoriesLifecycleStateEnum{
	"ACTIVE":   ListModelGroupVersionHistoriesLifecycleStateActive,
	"DELETED":  ListModelGroupVersionHistoriesLifecycleStateDeleted,
	"FAILED":   ListModelGroupVersionHistoriesLifecycleStateFailed,
	"DELETING": ListModelGroupVersionHistoriesLifecycleStateDeleting,
}

var mappingListModelGroupVersionHistoriesLifecycleStateEnumLowerCase = map[string]ListModelGroupVersionHistoriesLifecycleStateEnum{
	"active":   ListModelGroupVersionHistoriesLifecycleStateActive,
	"deleted":  ListModelGroupVersionHistoriesLifecycleStateDeleted,
	"failed":   ListModelGroupVersionHistoriesLifecycleStateFailed,
	"deleting": ListModelGroupVersionHistoriesLifecycleStateDeleting,
}

// GetListModelGroupVersionHistoriesLifecycleStateEnumValues Enumerates the set of values for ListModelGroupVersionHistoriesLifecycleStateEnum
func GetListModelGroupVersionHistoriesLifecycleStateEnumValues() []ListModelGroupVersionHistoriesLifecycleStateEnum {
	values := make([]ListModelGroupVersionHistoriesLifecycleStateEnum, 0)
	for _, v := range mappingListModelGroupVersionHistoriesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelGroupVersionHistoriesLifecycleStateEnumStringValues Enumerates the set of values in String for ListModelGroupVersionHistoriesLifecycleStateEnum
func GetListModelGroupVersionHistoriesLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"FAILED",
		"DELETING",
	}
}

// GetMappingListModelGroupVersionHistoriesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelGroupVersionHistoriesLifecycleStateEnum(val string) (ListModelGroupVersionHistoriesLifecycleStateEnum, bool) {
	enum, ok := mappingListModelGroupVersionHistoriesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModelGroupVersionHistoriesSortOrderEnum Enum with underlying type: string
type ListModelGroupVersionHistoriesSortOrderEnum string

// Set of constants representing the allowable values for ListModelGroupVersionHistoriesSortOrderEnum
const (
	ListModelGroupVersionHistoriesSortOrderAsc  ListModelGroupVersionHistoriesSortOrderEnum = "ASC"
	ListModelGroupVersionHistoriesSortOrderDesc ListModelGroupVersionHistoriesSortOrderEnum = "DESC"
)

var mappingListModelGroupVersionHistoriesSortOrderEnum = map[string]ListModelGroupVersionHistoriesSortOrderEnum{
	"ASC":  ListModelGroupVersionHistoriesSortOrderAsc,
	"DESC": ListModelGroupVersionHistoriesSortOrderDesc,
}

var mappingListModelGroupVersionHistoriesSortOrderEnumLowerCase = map[string]ListModelGroupVersionHistoriesSortOrderEnum{
	"asc":  ListModelGroupVersionHistoriesSortOrderAsc,
	"desc": ListModelGroupVersionHistoriesSortOrderDesc,
}

// GetListModelGroupVersionHistoriesSortOrderEnumValues Enumerates the set of values for ListModelGroupVersionHistoriesSortOrderEnum
func GetListModelGroupVersionHistoriesSortOrderEnumValues() []ListModelGroupVersionHistoriesSortOrderEnum {
	values := make([]ListModelGroupVersionHistoriesSortOrderEnum, 0)
	for _, v := range mappingListModelGroupVersionHistoriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelGroupVersionHistoriesSortOrderEnumStringValues Enumerates the set of values in String for ListModelGroupVersionHistoriesSortOrderEnum
func GetListModelGroupVersionHistoriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListModelGroupVersionHistoriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelGroupVersionHistoriesSortOrderEnum(val string) (ListModelGroupVersionHistoriesSortOrderEnum, bool) {
	enum, ok := mappingListModelGroupVersionHistoriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModelGroupVersionHistoriesSortByEnum Enum with underlying type: string
type ListModelGroupVersionHistoriesSortByEnum string

// Set of constants representing the allowable values for ListModelGroupVersionHistoriesSortByEnum
const (
	ListModelGroupVersionHistoriesSortByTimecreated    ListModelGroupVersionHistoriesSortByEnum = "timeCreated"
	ListModelGroupVersionHistoriesSortByDisplayname    ListModelGroupVersionHistoriesSortByEnum = "displayName"
	ListModelGroupVersionHistoriesSortByLifecyclestate ListModelGroupVersionHistoriesSortByEnum = "lifecycleState"
)

var mappingListModelGroupVersionHistoriesSortByEnum = map[string]ListModelGroupVersionHistoriesSortByEnum{
	"timeCreated":    ListModelGroupVersionHistoriesSortByTimecreated,
	"displayName":    ListModelGroupVersionHistoriesSortByDisplayname,
	"lifecycleState": ListModelGroupVersionHistoriesSortByLifecyclestate,
}

var mappingListModelGroupVersionHistoriesSortByEnumLowerCase = map[string]ListModelGroupVersionHistoriesSortByEnum{
	"timecreated":    ListModelGroupVersionHistoriesSortByTimecreated,
	"displayname":    ListModelGroupVersionHistoriesSortByDisplayname,
	"lifecyclestate": ListModelGroupVersionHistoriesSortByLifecyclestate,
}

// GetListModelGroupVersionHistoriesSortByEnumValues Enumerates the set of values for ListModelGroupVersionHistoriesSortByEnum
func GetListModelGroupVersionHistoriesSortByEnumValues() []ListModelGroupVersionHistoriesSortByEnum {
	values := make([]ListModelGroupVersionHistoriesSortByEnum, 0)
	for _, v := range mappingListModelGroupVersionHistoriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelGroupVersionHistoriesSortByEnumStringValues Enumerates the set of values in String for ListModelGroupVersionHistoriesSortByEnum
func GetListModelGroupVersionHistoriesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"lifecycleState",
	}
}

// GetMappingListModelGroupVersionHistoriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelGroupVersionHistoriesSortByEnum(val string) (ListModelGroupVersionHistoriesSortByEnum, bool) {
	enum, ok := mappingListModelGroupVersionHistoriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
