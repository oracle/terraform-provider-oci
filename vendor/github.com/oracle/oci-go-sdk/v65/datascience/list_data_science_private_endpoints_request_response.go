// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDataSciencePrivateEndpointsRequest wrapper for the ListDataSciencePrivateEndpoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListDataSciencePrivateEndpoints.go.html to see an example of how to use ListDataSciencePrivateEndpointsRequest.
type ListDataSciencePrivateEndpointsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

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

	// The lifecycle state of the private endpoint.
	LifecycleState ListDataSciencePrivateEndpointsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field used to sort the results. Multiple fields aren't supported.
	SortBy ListDataSciencePrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListDataSciencePrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

	// Resource types in the Data Science service such as notebooks.
	DataScienceResourceType ListDataSciencePrivateEndpointsDataScienceResourceTypeEnum `mandatory:"false" contributesTo:"query" name:"dataScienceResourceType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataSciencePrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataSciencePrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataSciencePrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataSciencePrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataSciencePrivateEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataSciencePrivateEndpointsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDataSciencePrivateEndpointsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSciencePrivateEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataSciencePrivateEndpointsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSciencePrivateEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataSciencePrivateEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSciencePrivateEndpointsDataScienceResourceTypeEnum(string(request.DataScienceResourceType)); !ok && request.DataScienceResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataScienceResourceType: %s. Supported values are: %s.", request.DataScienceResourceType, strings.Join(GetListDataSciencePrivateEndpointsDataScienceResourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataSciencePrivateEndpointsResponse wrapper for the ListDataSciencePrivateEndpoints operation
type ListDataSciencePrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DataSciencePrivateEndpointSummary instances
	Items []DataSciencePrivateEndpointSummary `presentIn:"body"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDataSciencePrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataSciencePrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataSciencePrivateEndpointsLifecycleStateEnum Enum with underlying type: string
type ListDataSciencePrivateEndpointsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDataSciencePrivateEndpointsLifecycleStateEnum
const (
	ListDataSciencePrivateEndpointsLifecycleStateCreating       ListDataSciencePrivateEndpointsLifecycleStateEnum = "CREATING"
	ListDataSciencePrivateEndpointsLifecycleStateActive         ListDataSciencePrivateEndpointsLifecycleStateEnum = "ACTIVE"
	ListDataSciencePrivateEndpointsLifecycleStateUpdating       ListDataSciencePrivateEndpointsLifecycleStateEnum = "UPDATING"
	ListDataSciencePrivateEndpointsLifecycleStateDeleting       ListDataSciencePrivateEndpointsLifecycleStateEnum = "DELETING"
	ListDataSciencePrivateEndpointsLifecycleStateDeleted        ListDataSciencePrivateEndpointsLifecycleStateEnum = "DELETED"
	ListDataSciencePrivateEndpointsLifecycleStateFailed         ListDataSciencePrivateEndpointsLifecycleStateEnum = "FAILED"
	ListDataSciencePrivateEndpointsLifecycleStateNeedsAttention ListDataSciencePrivateEndpointsLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListDataSciencePrivateEndpointsLifecycleStateEnum = map[string]ListDataSciencePrivateEndpointsLifecycleStateEnum{
	"CREATING":        ListDataSciencePrivateEndpointsLifecycleStateCreating,
	"ACTIVE":          ListDataSciencePrivateEndpointsLifecycleStateActive,
	"UPDATING":        ListDataSciencePrivateEndpointsLifecycleStateUpdating,
	"DELETING":        ListDataSciencePrivateEndpointsLifecycleStateDeleting,
	"DELETED":         ListDataSciencePrivateEndpointsLifecycleStateDeleted,
	"FAILED":          ListDataSciencePrivateEndpointsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListDataSciencePrivateEndpointsLifecycleStateNeedsAttention,
}

var mappingListDataSciencePrivateEndpointsLifecycleStateEnumLowerCase = map[string]ListDataSciencePrivateEndpointsLifecycleStateEnum{
	"creating":        ListDataSciencePrivateEndpointsLifecycleStateCreating,
	"active":          ListDataSciencePrivateEndpointsLifecycleStateActive,
	"updating":        ListDataSciencePrivateEndpointsLifecycleStateUpdating,
	"deleting":        ListDataSciencePrivateEndpointsLifecycleStateDeleting,
	"deleted":         ListDataSciencePrivateEndpointsLifecycleStateDeleted,
	"failed":          ListDataSciencePrivateEndpointsLifecycleStateFailed,
	"needs_attention": ListDataSciencePrivateEndpointsLifecycleStateNeedsAttention,
}

// GetListDataSciencePrivateEndpointsLifecycleStateEnumValues Enumerates the set of values for ListDataSciencePrivateEndpointsLifecycleStateEnum
func GetListDataSciencePrivateEndpointsLifecycleStateEnumValues() []ListDataSciencePrivateEndpointsLifecycleStateEnum {
	values := make([]ListDataSciencePrivateEndpointsLifecycleStateEnum, 0)
	for _, v := range mappingListDataSciencePrivateEndpointsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSciencePrivateEndpointsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDataSciencePrivateEndpointsLifecycleStateEnum
func GetListDataSciencePrivateEndpointsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListDataSciencePrivateEndpointsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSciencePrivateEndpointsLifecycleStateEnum(val string) (ListDataSciencePrivateEndpointsLifecycleStateEnum, bool) {
	enum, ok := mappingListDataSciencePrivateEndpointsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSciencePrivateEndpointsSortByEnum Enum with underlying type: string
type ListDataSciencePrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListDataSciencePrivateEndpointsSortByEnum
const (
	ListDataSciencePrivateEndpointsSortByTimecreated ListDataSciencePrivateEndpointsSortByEnum = "timeCreated"
)

var mappingListDataSciencePrivateEndpointsSortByEnum = map[string]ListDataSciencePrivateEndpointsSortByEnum{
	"timeCreated": ListDataSciencePrivateEndpointsSortByTimecreated,
}

var mappingListDataSciencePrivateEndpointsSortByEnumLowerCase = map[string]ListDataSciencePrivateEndpointsSortByEnum{
	"timecreated": ListDataSciencePrivateEndpointsSortByTimecreated,
}

// GetListDataSciencePrivateEndpointsSortByEnumValues Enumerates the set of values for ListDataSciencePrivateEndpointsSortByEnum
func GetListDataSciencePrivateEndpointsSortByEnumValues() []ListDataSciencePrivateEndpointsSortByEnum {
	values := make([]ListDataSciencePrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListDataSciencePrivateEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSciencePrivateEndpointsSortByEnumStringValues Enumerates the set of values in String for ListDataSciencePrivateEndpointsSortByEnum
func GetListDataSciencePrivateEndpointsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListDataSciencePrivateEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSciencePrivateEndpointsSortByEnum(val string) (ListDataSciencePrivateEndpointsSortByEnum, bool) {
	enum, ok := mappingListDataSciencePrivateEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSciencePrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListDataSciencePrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListDataSciencePrivateEndpointsSortOrderEnum
const (
	ListDataSciencePrivateEndpointsSortOrderAsc  ListDataSciencePrivateEndpointsSortOrderEnum = "ASC"
	ListDataSciencePrivateEndpointsSortOrderDesc ListDataSciencePrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListDataSciencePrivateEndpointsSortOrderEnum = map[string]ListDataSciencePrivateEndpointsSortOrderEnum{
	"ASC":  ListDataSciencePrivateEndpointsSortOrderAsc,
	"DESC": ListDataSciencePrivateEndpointsSortOrderDesc,
}

var mappingListDataSciencePrivateEndpointsSortOrderEnumLowerCase = map[string]ListDataSciencePrivateEndpointsSortOrderEnum{
	"asc":  ListDataSciencePrivateEndpointsSortOrderAsc,
	"desc": ListDataSciencePrivateEndpointsSortOrderDesc,
}

// GetListDataSciencePrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListDataSciencePrivateEndpointsSortOrderEnum
func GetListDataSciencePrivateEndpointsSortOrderEnumValues() []ListDataSciencePrivateEndpointsSortOrderEnum {
	values := make([]ListDataSciencePrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListDataSciencePrivateEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSciencePrivateEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListDataSciencePrivateEndpointsSortOrderEnum
func GetListDataSciencePrivateEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataSciencePrivateEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSciencePrivateEndpointsSortOrderEnum(val string) (ListDataSciencePrivateEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListDataSciencePrivateEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSciencePrivateEndpointsDataScienceResourceTypeEnum Enum with underlying type: string
type ListDataSciencePrivateEndpointsDataScienceResourceTypeEnum string

// Set of constants representing the allowable values for ListDataSciencePrivateEndpointsDataScienceResourceTypeEnum
const (
	ListDataSciencePrivateEndpointsDataScienceResourceTypeNotebookSession ListDataSciencePrivateEndpointsDataScienceResourceTypeEnum = "NOTEBOOK_SESSION"
)

var mappingListDataSciencePrivateEndpointsDataScienceResourceTypeEnum = map[string]ListDataSciencePrivateEndpointsDataScienceResourceTypeEnum{
	"NOTEBOOK_SESSION": ListDataSciencePrivateEndpointsDataScienceResourceTypeNotebookSession,
}

var mappingListDataSciencePrivateEndpointsDataScienceResourceTypeEnumLowerCase = map[string]ListDataSciencePrivateEndpointsDataScienceResourceTypeEnum{
	"notebook_session": ListDataSciencePrivateEndpointsDataScienceResourceTypeNotebookSession,
}

// GetListDataSciencePrivateEndpointsDataScienceResourceTypeEnumValues Enumerates the set of values for ListDataSciencePrivateEndpointsDataScienceResourceTypeEnum
func GetListDataSciencePrivateEndpointsDataScienceResourceTypeEnumValues() []ListDataSciencePrivateEndpointsDataScienceResourceTypeEnum {
	values := make([]ListDataSciencePrivateEndpointsDataScienceResourceTypeEnum, 0)
	for _, v := range mappingListDataSciencePrivateEndpointsDataScienceResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSciencePrivateEndpointsDataScienceResourceTypeEnumStringValues Enumerates the set of values in String for ListDataSciencePrivateEndpointsDataScienceResourceTypeEnum
func GetListDataSciencePrivateEndpointsDataScienceResourceTypeEnumStringValues() []string {
	return []string{
		"NOTEBOOK_SESSION",
	}
}

// GetMappingListDataSciencePrivateEndpointsDataScienceResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSciencePrivateEndpointsDataScienceResourceTypeEnum(val string) (ListDataSciencePrivateEndpointsDataScienceResourceTypeEnum, bool) {
	enum, ok := mappingListDataSciencePrivateEndpointsDataScienceResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
