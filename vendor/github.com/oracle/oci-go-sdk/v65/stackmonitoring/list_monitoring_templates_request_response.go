// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMonitoringTemplatesRequest wrapper for the ListMonitoringTemplates operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListMonitoringTemplates.go.html to see an example of how to use ListMonitoringTemplatesRequest.
type ListMonitoringTemplatesRequest struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The ID of the compartment in which data is listed.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return monitoring template based on input monitoringTemplateId
	MonitoringTemplateId *string `mandatory:"false" contributesTo:"query" name:"monitoringTemplateId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMonitoringTemplatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeUpdated is descending.
	SortBy ListMonitoringTemplatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return monitoring template based on name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return monitoring template based on input status
	Status ListMonitoringTemplatesStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// A filter to return monitoring template based on Lifecycle State
	LifecycleState ListMonitoringTemplatesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Multiple resource types filter.
	ResourceTypes []string `contributesTo:"query" name:"resourceTypes" collectionFormat:"multi"`

	// metricName filter.
	MetricName []string `contributesTo:"query" name:"metricName" collectionFormat:"multi"`

	// namespace filter.
	Namespace []string `contributesTo:"query" name:"namespace" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMonitoringTemplatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMonitoringTemplatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMonitoringTemplatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMonitoringTemplatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMonitoringTemplatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMonitoringTemplatesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMonitoringTemplatesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMonitoringTemplatesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMonitoringTemplatesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMonitoringTemplatesStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListMonitoringTemplatesStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMonitoringTemplatesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListMonitoringTemplatesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMonitoringTemplatesResponse wrapper for the ListMonitoringTemplates operation
type ListMonitoringTemplatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MonitoringTemplateCollection instances
	MonitoringTemplateCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMonitoringTemplatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMonitoringTemplatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMonitoringTemplatesSortOrderEnum Enum with underlying type: string
type ListMonitoringTemplatesSortOrderEnum string

// Set of constants representing the allowable values for ListMonitoringTemplatesSortOrderEnum
const (
	ListMonitoringTemplatesSortOrderAsc  ListMonitoringTemplatesSortOrderEnum = "ASC"
	ListMonitoringTemplatesSortOrderDesc ListMonitoringTemplatesSortOrderEnum = "DESC"
)

var mappingListMonitoringTemplatesSortOrderEnum = map[string]ListMonitoringTemplatesSortOrderEnum{
	"ASC":  ListMonitoringTemplatesSortOrderAsc,
	"DESC": ListMonitoringTemplatesSortOrderDesc,
}

var mappingListMonitoringTemplatesSortOrderEnumLowerCase = map[string]ListMonitoringTemplatesSortOrderEnum{
	"asc":  ListMonitoringTemplatesSortOrderAsc,
	"desc": ListMonitoringTemplatesSortOrderDesc,
}

// GetListMonitoringTemplatesSortOrderEnumValues Enumerates the set of values for ListMonitoringTemplatesSortOrderEnum
func GetListMonitoringTemplatesSortOrderEnumValues() []ListMonitoringTemplatesSortOrderEnum {
	values := make([]ListMonitoringTemplatesSortOrderEnum, 0)
	for _, v := range mappingListMonitoringTemplatesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoringTemplatesSortOrderEnumStringValues Enumerates the set of values in String for ListMonitoringTemplatesSortOrderEnum
func GetListMonitoringTemplatesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMonitoringTemplatesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoringTemplatesSortOrderEnum(val string) (ListMonitoringTemplatesSortOrderEnum, bool) {
	enum, ok := mappingListMonitoringTemplatesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMonitoringTemplatesSortByEnum Enum with underlying type: string
type ListMonitoringTemplatesSortByEnum string

// Set of constants representing the allowable values for ListMonitoringTemplatesSortByEnum
const (
	ListMonitoringTemplatesSortByDisplayname    ListMonitoringTemplatesSortByEnum = "displayName"
	ListMonitoringTemplatesSortByLifecyclestate ListMonitoringTemplatesSortByEnum = "lifeCycleState"
	ListMonitoringTemplatesSortByStatus         ListMonitoringTemplatesSortByEnum = "status"
	ListMonitoringTemplatesSortByTimeupdated    ListMonitoringTemplatesSortByEnum = "timeUpdated"
	ListMonitoringTemplatesSortByTimecreated    ListMonitoringTemplatesSortByEnum = "timeCreated"
)

var mappingListMonitoringTemplatesSortByEnum = map[string]ListMonitoringTemplatesSortByEnum{
	"displayName":    ListMonitoringTemplatesSortByDisplayname,
	"lifeCycleState": ListMonitoringTemplatesSortByLifecyclestate,
	"status":         ListMonitoringTemplatesSortByStatus,
	"timeUpdated":    ListMonitoringTemplatesSortByTimeupdated,
	"timeCreated":    ListMonitoringTemplatesSortByTimecreated,
}

var mappingListMonitoringTemplatesSortByEnumLowerCase = map[string]ListMonitoringTemplatesSortByEnum{
	"displayname":    ListMonitoringTemplatesSortByDisplayname,
	"lifecyclestate": ListMonitoringTemplatesSortByLifecyclestate,
	"status":         ListMonitoringTemplatesSortByStatus,
	"timeupdated":    ListMonitoringTemplatesSortByTimeupdated,
	"timecreated":    ListMonitoringTemplatesSortByTimecreated,
}

// GetListMonitoringTemplatesSortByEnumValues Enumerates the set of values for ListMonitoringTemplatesSortByEnum
func GetListMonitoringTemplatesSortByEnumValues() []ListMonitoringTemplatesSortByEnum {
	values := make([]ListMonitoringTemplatesSortByEnum, 0)
	for _, v := range mappingListMonitoringTemplatesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoringTemplatesSortByEnumStringValues Enumerates the set of values in String for ListMonitoringTemplatesSortByEnum
func GetListMonitoringTemplatesSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"lifeCycleState",
		"status",
		"timeUpdated",
		"timeCreated",
	}
}

// GetMappingListMonitoringTemplatesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoringTemplatesSortByEnum(val string) (ListMonitoringTemplatesSortByEnum, bool) {
	enum, ok := mappingListMonitoringTemplatesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMonitoringTemplatesStatusEnum Enum with underlying type: string
type ListMonitoringTemplatesStatusEnum string

// Set of constants representing the allowable values for ListMonitoringTemplatesStatusEnum
const (
	ListMonitoringTemplatesStatusNotApplied     ListMonitoringTemplatesStatusEnum = "NOT_APPLIED"
	ListMonitoringTemplatesStatusApplied        ListMonitoringTemplatesStatusEnum = "APPLIED"
	ListMonitoringTemplatesStatusPartialApplied ListMonitoringTemplatesStatusEnum = "PARTIAL_APPLIED"
)

var mappingListMonitoringTemplatesStatusEnum = map[string]ListMonitoringTemplatesStatusEnum{
	"NOT_APPLIED":     ListMonitoringTemplatesStatusNotApplied,
	"APPLIED":         ListMonitoringTemplatesStatusApplied,
	"PARTIAL_APPLIED": ListMonitoringTemplatesStatusPartialApplied,
}

var mappingListMonitoringTemplatesStatusEnumLowerCase = map[string]ListMonitoringTemplatesStatusEnum{
	"not_applied":     ListMonitoringTemplatesStatusNotApplied,
	"applied":         ListMonitoringTemplatesStatusApplied,
	"partial_applied": ListMonitoringTemplatesStatusPartialApplied,
}

// GetListMonitoringTemplatesStatusEnumValues Enumerates the set of values for ListMonitoringTemplatesStatusEnum
func GetListMonitoringTemplatesStatusEnumValues() []ListMonitoringTemplatesStatusEnum {
	values := make([]ListMonitoringTemplatesStatusEnum, 0)
	for _, v := range mappingListMonitoringTemplatesStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoringTemplatesStatusEnumStringValues Enumerates the set of values in String for ListMonitoringTemplatesStatusEnum
func GetListMonitoringTemplatesStatusEnumStringValues() []string {
	return []string{
		"NOT_APPLIED",
		"APPLIED",
		"PARTIAL_APPLIED",
	}
}

// GetMappingListMonitoringTemplatesStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoringTemplatesStatusEnum(val string) (ListMonitoringTemplatesStatusEnum, bool) {
	enum, ok := mappingListMonitoringTemplatesStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMonitoringTemplatesLifecycleStateEnum Enum with underlying type: string
type ListMonitoringTemplatesLifecycleStateEnum string

// Set of constants representing the allowable values for ListMonitoringTemplatesLifecycleStateEnum
const (
	ListMonitoringTemplatesLifecycleStateCreating ListMonitoringTemplatesLifecycleStateEnum = "CREATING"
	ListMonitoringTemplatesLifecycleStateActive   ListMonitoringTemplatesLifecycleStateEnum = "ACTIVE"
	ListMonitoringTemplatesLifecycleStateInactive ListMonitoringTemplatesLifecycleStateEnum = "INACTIVE"
	ListMonitoringTemplatesLifecycleStateUpdating ListMonitoringTemplatesLifecycleStateEnum = "UPDATING"
	ListMonitoringTemplatesLifecycleStateDeleted  ListMonitoringTemplatesLifecycleStateEnum = "DELETED"
)

var mappingListMonitoringTemplatesLifecycleStateEnum = map[string]ListMonitoringTemplatesLifecycleStateEnum{
	"CREATING": ListMonitoringTemplatesLifecycleStateCreating,
	"ACTIVE":   ListMonitoringTemplatesLifecycleStateActive,
	"INACTIVE": ListMonitoringTemplatesLifecycleStateInactive,
	"UPDATING": ListMonitoringTemplatesLifecycleStateUpdating,
	"DELETED":  ListMonitoringTemplatesLifecycleStateDeleted,
}

var mappingListMonitoringTemplatesLifecycleStateEnumLowerCase = map[string]ListMonitoringTemplatesLifecycleStateEnum{
	"creating": ListMonitoringTemplatesLifecycleStateCreating,
	"active":   ListMonitoringTemplatesLifecycleStateActive,
	"inactive": ListMonitoringTemplatesLifecycleStateInactive,
	"updating": ListMonitoringTemplatesLifecycleStateUpdating,
	"deleted":  ListMonitoringTemplatesLifecycleStateDeleted,
}

// GetListMonitoringTemplatesLifecycleStateEnumValues Enumerates the set of values for ListMonitoringTemplatesLifecycleStateEnum
func GetListMonitoringTemplatesLifecycleStateEnumValues() []ListMonitoringTemplatesLifecycleStateEnum {
	values := make([]ListMonitoringTemplatesLifecycleStateEnum, 0)
	for _, v := range mappingListMonitoringTemplatesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoringTemplatesLifecycleStateEnumStringValues Enumerates the set of values in String for ListMonitoringTemplatesLifecycleStateEnum
func GetListMonitoringTemplatesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETED",
	}
}

// GetMappingListMonitoringTemplatesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoringTemplatesLifecycleStateEnum(val string) (ListMonitoringTemplatesLifecycleStateEnum, bool) {
	enum, ok := mappingListMonitoringTemplatesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
