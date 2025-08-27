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

// ListModelGroupModelsRequest wrapper for the ListModelGroupModels operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListModelGroupModels.go.html to see an example of how to use ListModelGroupModelsRequest.
type ListModelGroupModelsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the modelGroup.
	ModelGroupId *string `mandatory:"true" contributesTo:"path" name:"modelGroupId"`

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListModelGroupModelsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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
	SortOrder ListModelGroupModelsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, the results are shown
	// in descending order. All other fields default to ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListModelGroupModelsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModelGroupModelsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModelGroupModelsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModelGroupModelsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModelGroupModelsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListModelGroupModelsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListModelGroupModelsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListModelGroupModelsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModelGroupModelsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListModelGroupModelsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModelGroupModelsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListModelGroupModelsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListModelGroupModelsResponse wrapper for the ListModelGroupModels operation
type ListModelGroupModelsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ModelGroupModelSummary instances
	Items []ModelGroupModelSummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListModelGroupModelsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModelGroupModelsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModelGroupModelsLifecycleStateEnum Enum with underlying type: string
type ListModelGroupModelsLifecycleStateEnum string

// Set of constants representing the allowable values for ListModelGroupModelsLifecycleStateEnum
const (
	ListModelGroupModelsLifecycleStateActive   ListModelGroupModelsLifecycleStateEnum = "ACTIVE"
	ListModelGroupModelsLifecycleStateDeleted  ListModelGroupModelsLifecycleStateEnum = "DELETED"
	ListModelGroupModelsLifecycleStateFailed   ListModelGroupModelsLifecycleStateEnum = "FAILED"
	ListModelGroupModelsLifecycleStateInactive ListModelGroupModelsLifecycleStateEnum = "INACTIVE"
)

var mappingListModelGroupModelsLifecycleStateEnum = map[string]ListModelGroupModelsLifecycleStateEnum{
	"ACTIVE":   ListModelGroupModelsLifecycleStateActive,
	"DELETED":  ListModelGroupModelsLifecycleStateDeleted,
	"FAILED":   ListModelGroupModelsLifecycleStateFailed,
	"INACTIVE": ListModelGroupModelsLifecycleStateInactive,
}

var mappingListModelGroupModelsLifecycleStateEnumLowerCase = map[string]ListModelGroupModelsLifecycleStateEnum{
	"active":   ListModelGroupModelsLifecycleStateActive,
	"deleted":  ListModelGroupModelsLifecycleStateDeleted,
	"failed":   ListModelGroupModelsLifecycleStateFailed,
	"inactive": ListModelGroupModelsLifecycleStateInactive,
}

// GetListModelGroupModelsLifecycleStateEnumValues Enumerates the set of values for ListModelGroupModelsLifecycleStateEnum
func GetListModelGroupModelsLifecycleStateEnumValues() []ListModelGroupModelsLifecycleStateEnum {
	values := make([]ListModelGroupModelsLifecycleStateEnum, 0)
	for _, v := range mappingListModelGroupModelsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelGroupModelsLifecycleStateEnumStringValues Enumerates the set of values in String for ListModelGroupModelsLifecycleStateEnum
func GetListModelGroupModelsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"FAILED",
		"INACTIVE",
	}
}

// GetMappingListModelGroupModelsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelGroupModelsLifecycleStateEnum(val string) (ListModelGroupModelsLifecycleStateEnum, bool) {
	enum, ok := mappingListModelGroupModelsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModelGroupModelsSortOrderEnum Enum with underlying type: string
type ListModelGroupModelsSortOrderEnum string

// Set of constants representing the allowable values for ListModelGroupModelsSortOrderEnum
const (
	ListModelGroupModelsSortOrderAsc  ListModelGroupModelsSortOrderEnum = "ASC"
	ListModelGroupModelsSortOrderDesc ListModelGroupModelsSortOrderEnum = "DESC"
)

var mappingListModelGroupModelsSortOrderEnum = map[string]ListModelGroupModelsSortOrderEnum{
	"ASC":  ListModelGroupModelsSortOrderAsc,
	"DESC": ListModelGroupModelsSortOrderDesc,
}

var mappingListModelGroupModelsSortOrderEnumLowerCase = map[string]ListModelGroupModelsSortOrderEnum{
	"asc":  ListModelGroupModelsSortOrderAsc,
	"desc": ListModelGroupModelsSortOrderDesc,
}

// GetListModelGroupModelsSortOrderEnumValues Enumerates the set of values for ListModelGroupModelsSortOrderEnum
func GetListModelGroupModelsSortOrderEnumValues() []ListModelGroupModelsSortOrderEnum {
	values := make([]ListModelGroupModelsSortOrderEnum, 0)
	for _, v := range mappingListModelGroupModelsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelGroupModelsSortOrderEnumStringValues Enumerates the set of values in String for ListModelGroupModelsSortOrderEnum
func GetListModelGroupModelsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListModelGroupModelsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelGroupModelsSortOrderEnum(val string) (ListModelGroupModelsSortOrderEnum, bool) {
	enum, ok := mappingListModelGroupModelsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModelGroupModelsSortByEnum Enum with underlying type: string
type ListModelGroupModelsSortByEnum string

// Set of constants representing the allowable values for ListModelGroupModelsSortByEnum
const (
	ListModelGroupModelsSortByTimecreated    ListModelGroupModelsSortByEnum = "timeCreated"
	ListModelGroupModelsSortByDisplayname    ListModelGroupModelsSortByEnum = "displayName"
	ListModelGroupModelsSortByLifecyclestate ListModelGroupModelsSortByEnum = "lifecycleState"
)

var mappingListModelGroupModelsSortByEnum = map[string]ListModelGroupModelsSortByEnum{
	"timeCreated":    ListModelGroupModelsSortByTimecreated,
	"displayName":    ListModelGroupModelsSortByDisplayname,
	"lifecycleState": ListModelGroupModelsSortByLifecyclestate,
}

var mappingListModelGroupModelsSortByEnumLowerCase = map[string]ListModelGroupModelsSortByEnum{
	"timecreated":    ListModelGroupModelsSortByTimecreated,
	"displayname":    ListModelGroupModelsSortByDisplayname,
	"lifecyclestate": ListModelGroupModelsSortByLifecyclestate,
}

// GetListModelGroupModelsSortByEnumValues Enumerates the set of values for ListModelGroupModelsSortByEnum
func GetListModelGroupModelsSortByEnumValues() []ListModelGroupModelsSortByEnum {
	values := make([]ListModelGroupModelsSortByEnum, 0)
	for _, v := range mappingListModelGroupModelsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelGroupModelsSortByEnumStringValues Enumerates the set of values in String for ListModelGroupModelsSortByEnum
func GetListModelGroupModelsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"lifecycleState",
	}
}

// GetMappingListModelGroupModelsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelGroupModelsSortByEnum(val string) (ListModelGroupModelsSortByEnum, bool) {
	enum, ok := mappingListModelGroupModelsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
