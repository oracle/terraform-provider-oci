// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMetricExtensionsRequest wrapper for the ListMetricExtensions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListMetricExtensions.go.html to see an example of how to use ListMetricExtensionsRequest.
type ListMetricExtensionsRequest struct {

	// The ID of the compartment in which data is listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for resources is ascending.
	SortBy ListMetricExtensionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMetricExtensionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return resources based on resource type.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// A filter to return resources based on name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return resources based on status e.g. Draft or Published
	Status ListMetricExtensionsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// A filter to return metric extensions based on Lifecycle State
	LifecycleState ListMetricExtensionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return metric extensions based on input resource Id on which metric extension is enabled
	EnabledOnResourceId *string `mandatory:"false" contributesTo:"query" name:"enabledOnResourceId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMetricExtensionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMetricExtensionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMetricExtensionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMetricExtensionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMetricExtensionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMetricExtensionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMetricExtensionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMetricExtensionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMetricExtensionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMetricExtensionsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListMetricExtensionsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMetricExtensionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListMetricExtensionsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMetricExtensionsResponse wrapper for the ListMetricExtensions operation
type ListMetricExtensionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MetricExtensionCollection instances
	MetricExtensionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMetricExtensionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMetricExtensionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMetricExtensionsSortByEnum Enum with underlying type: string
type ListMetricExtensionsSortByEnum string

// Set of constants representing the allowable values for ListMetricExtensionsSortByEnum
const (
	ListMetricExtensionsSortByName        ListMetricExtensionsSortByEnum = "NAME"
	ListMetricExtensionsSortByTimeCreated ListMetricExtensionsSortByEnum = "TIME_CREATED"
)

var mappingListMetricExtensionsSortByEnum = map[string]ListMetricExtensionsSortByEnum{
	"NAME":         ListMetricExtensionsSortByName,
	"TIME_CREATED": ListMetricExtensionsSortByTimeCreated,
}

var mappingListMetricExtensionsSortByEnumLowerCase = map[string]ListMetricExtensionsSortByEnum{
	"name":         ListMetricExtensionsSortByName,
	"time_created": ListMetricExtensionsSortByTimeCreated,
}

// GetListMetricExtensionsSortByEnumValues Enumerates the set of values for ListMetricExtensionsSortByEnum
func GetListMetricExtensionsSortByEnumValues() []ListMetricExtensionsSortByEnum {
	values := make([]ListMetricExtensionsSortByEnum, 0)
	for _, v := range mappingListMetricExtensionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMetricExtensionsSortByEnumStringValues Enumerates the set of values in String for ListMetricExtensionsSortByEnum
func GetListMetricExtensionsSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIME_CREATED",
	}
}

// GetMappingListMetricExtensionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMetricExtensionsSortByEnum(val string) (ListMetricExtensionsSortByEnum, bool) {
	enum, ok := mappingListMetricExtensionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMetricExtensionsSortOrderEnum Enum with underlying type: string
type ListMetricExtensionsSortOrderEnum string

// Set of constants representing the allowable values for ListMetricExtensionsSortOrderEnum
const (
	ListMetricExtensionsSortOrderAsc  ListMetricExtensionsSortOrderEnum = "ASC"
	ListMetricExtensionsSortOrderDesc ListMetricExtensionsSortOrderEnum = "DESC"
)

var mappingListMetricExtensionsSortOrderEnum = map[string]ListMetricExtensionsSortOrderEnum{
	"ASC":  ListMetricExtensionsSortOrderAsc,
	"DESC": ListMetricExtensionsSortOrderDesc,
}

var mappingListMetricExtensionsSortOrderEnumLowerCase = map[string]ListMetricExtensionsSortOrderEnum{
	"asc":  ListMetricExtensionsSortOrderAsc,
	"desc": ListMetricExtensionsSortOrderDesc,
}

// GetListMetricExtensionsSortOrderEnumValues Enumerates the set of values for ListMetricExtensionsSortOrderEnum
func GetListMetricExtensionsSortOrderEnumValues() []ListMetricExtensionsSortOrderEnum {
	values := make([]ListMetricExtensionsSortOrderEnum, 0)
	for _, v := range mappingListMetricExtensionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMetricExtensionsSortOrderEnumStringValues Enumerates the set of values in String for ListMetricExtensionsSortOrderEnum
func GetListMetricExtensionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMetricExtensionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMetricExtensionsSortOrderEnum(val string) (ListMetricExtensionsSortOrderEnum, bool) {
	enum, ok := mappingListMetricExtensionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMetricExtensionsStatusEnum Enum with underlying type: string
type ListMetricExtensionsStatusEnum string

// Set of constants representing the allowable values for ListMetricExtensionsStatusEnum
const (
	ListMetricExtensionsStatusDraft     ListMetricExtensionsStatusEnum = "DRAFT"
	ListMetricExtensionsStatusPublished ListMetricExtensionsStatusEnum = "PUBLISHED"
)

var mappingListMetricExtensionsStatusEnum = map[string]ListMetricExtensionsStatusEnum{
	"DRAFT":     ListMetricExtensionsStatusDraft,
	"PUBLISHED": ListMetricExtensionsStatusPublished,
}

var mappingListMetricExtensionsStatusEnumLowerCase = map[string]ListMetricExtensionsStatusEnum{
	"draft":     ListMetricExtensionsStatusDraft,
	"published": ListMetricExtensionsStatusPublished,
}

// GetListMetricExtensionsStatusEnumValues Enumerates the set of values for ListMetricExtensionsStatusEnum
func GetListMetricExtensionsStatusEnumValues() []ListMetricExtensionsStatusEnum {
	values := make([]ListMetricExtensionsStatusEnum, 0)
	for _, v := range mappingListMetricExtensionsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListMetricExtensionsStatusEnumStringValues Enumerates the set of values in String for ListMetricExtensionsStatusEnum
func GetListMetricExtensionsStatusEnumStringValues() []string {
	return []string{
		"DRAFT",
		"PUBLISHED",
	}
}

// GetMappingListMetricExtensionsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMetricExtensionsStatusEnum(val string) (ListMetricExtensionsStatusEnum, bool) {
	enum, ok := mappingListMetricExtensionsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMetricExtensionsLifecycleStateEnum Enum with underlying type: string
type ListMetricExtensionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListMetricExtensionsLifecycleStateEnum
const (
	ListMetricExtensionsLifecycleStateActive  ListMetricExtensionsLifecycleStateEnum = "ACTIVE"
	ListMetricExtensionsLifecycleStateDeleted ListMetricExtensionsLifecycleStateEnum = "DELETED"
)

var mappingListMetricExtensionsLifecycleStateEnum = map[string]ListMetricExtensionsLifecycleStateEnum{
	"ACTIVE":  ListMetricExtensionsLifecycleStateActive,
	"DELETED": ListMetricExtensionsLifecycleStateDeleted,
}

var mappingListMetricExtensionsLifecycleStateEnumLowerCase = map[string]ListMetricExtensionsLifecycleStateEnum{
	"active":  ListMetricExtensionsLifecycleStateActive,
	"deleted": ListMetricExtensionsLifecycleStateDeleted,
}

// GetListMetricExtensionsLifecycleStateEnumValues Enumerates the set of values for ListMetricExtensionsLifecycleStateEnum
func GetListMetricExtensionsLifecycleStateEnumValues() []ListMetricExtensionsLifecycleStateEnum {
	values := make([]ListMetricExtensionsLifecycleStateEnum, 0)
	for _, v := range mappingListMetricExtensionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListMetricExtensionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListMetricExtensionsLifecycleStateEnum
func GetListMetricExtensionsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListMetricExtensionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMetricExtensionsLifecycleStateEnum(val string) (ListMetricExtensionsLifecycleStateEnum, bool) {
	enum, ok := mappingListMetricExtensionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
