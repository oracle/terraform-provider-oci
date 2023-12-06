// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDrPlansRequest wrapper for the ListDrPlans operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/ListDrPlans.go.html to see an example of how to use ListDrPlansRequest.
type ListDrPlansRequest struct {

	// The OCID of the DR protection group. Mandatory query param.
	// Example: `ocid1.drprotectiongroup.oc1..uniqueID`
	DrProtectionGroupId *string `mandatory:"true" contributesTo:"query" name:"drProtectionGroupId"`

	// A filter to return only DR plans that match the given lifecycle state.
	LifecycleState ListDrPlansLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The OCID of the DR plan.
	// Example: `ocid1.drplan.oc1..uniqueID`
	DrPlanId *string `mandatory:"false" contributesTo:"query" name:"drPlanId"`

	// The DR plan type.
	DrPlanType ListDrPlansDrPlanTypeEnum `mandatory:"false" contributesTo:"query" name:"drPlanType" omitEmpty:"true"`

	// A filter to return only resources that match the given display name.
	// Example: `MyResourceDisplayName`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `100`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDrPlansSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	// Default order for displayName is ascending. If no value is specified timeCreated is default.
	// Example: `MyResourceDisplayName`
	SortBy ListDrPlansSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDrPlansRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDrPlansRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDrPlansRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDrPlansRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDrPlansRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDrPlansLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDrPlansLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDrPlansDrPlanTypeEnum(string(request.DrPlanType)); !ok && request.DrPlanType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrPlanType: %s. Supported values are: %s.", request.DrPlanType, strings.Join(GetListDrPlansDrPlanTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDrPlansSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDrPlansSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDrPlansSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDrPlansSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDrPlansResponse wrapper for the ListDrPlans operation
type ListDrPlansResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DrPlanCollection instances
	DrPlanCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDrPlansResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDrPlansResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDrPlansLifecycleStateEnum Enum with underlying type: string
type ListDrPlansLifecycleStateEnum string

// Set of constants representing the allowable values for ListDrPlansLifecycleStateEnum
const (
	ListDrPlansLifecycleStateCreating       ListDrPlansLifecycleStateEnum = "CREATING"
	ListDrPlansLifecycleStateUpdating       ListDrPlansLifecycleStateEnum = "UPDATING"
	ListDrPlansLifecycleStateActive         ListDrPlansLifecycleStateEnum = "ACTIVE"
	ListDrPlansLifecycleStateInactive       ListDrPlansLifecycleStateEnum = "INACTIVE"
	ListDrPlansLifecycleStateDeleting       ListDrPlansLifecycleStateEnum = "DELETING"
	ListDrPlansLifecycleStateDeleted        ListDrPlansLifecycleStateEnum = "DELETED"
	ListDrPlansLifecycleStateFailed         ListDrPlansLifecycleStateEnum = "FAILED"
	ListDrPlansLifecycleStateNeedsAttention ListDrPlansLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListDrPlansLifecycleStateEnum = map[string]ListDrPlansLifecycleStateEnum{
	"CREATING":        ListDrPlansLifecycleStateCreating,
	"UPDATING":        ListDrPlansLifecycleStateUpdating,
	"ACTIVE":          ListDrPlansLifecycleStateActive,
	"INACTIVE":        ListDrPlansLifecycleStateInactive,
	"DELETING":        ListDrPlansLifecycleStateDeleting,
	"DELETED":         ListDrPlansLifecycleStateDeleted,
	"FAILED":          ListDrPlansLifecycleStateFailed,
	"NEEDS_ATTENTION": ListDrPlansLifecycleStateNeedsAttention,
}

var mappingListDrPlansLifecycleStateEnumLowerCase = map[string]ListDrPlansLifecycleStateEnum{
	"creating":        ListDrPlansLifecycleStateCreating,
	"updating":        ListDrPlansLifecycleStateUpdating,
	"active":          ListDrPlansLifecycleStateActive,
	"inactive":        ListDrPlansLifecycleStateInactive,
	"deleting":        ListDrPlansLifecycleStateDeleting,
	"deleted":         ListDrPlansLifecycleStateDeleted,
	"failed":          ListDrPlansLifecycleStateFailed,
	"needs_attention": ListDrPlansLifecycleStateNeedsAttention,
}

// GetListDrPlansLifecycleStateEnumValues Enumerates the set of values for ListDrPlansLifecycleStateEnum
func GetListDrPlansLifecycleStateEnumValues() []ListDrPlansLifecycleStateEnum {
	values := make([]ListDrPlansLifecycleStateEnum, 0)
	for _, v := range mappingListDrPlansLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrPlansLifecycleStateEnumStringValues Enumerates the set of values in String for ListDrPlansLifecycleStateEnum
func GetListDrPlansLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListDrPlansLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrPlansLifecycleStateEnum(val string) (ListDrPlansLifecycleStateEnum, bool) {
	enum, ok := mappingListDrPlansLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDrPlansDrPlanTypeEnum Enum with underlying type: string
type ListDrPlansDrPlanTypeEnum string

// Set of constants representing the allowable values for ListDrPlansDrPlanTypeEnum
const (
	ListDrPlansDrPlanTypeSwitchover ListDrPlansDrPlanTypeEnum = "SWITCHOVER"
	ListDrPlansDrPlanTypeFailover   ListDrPlansDrPlanTypeEnum = "FAILOVER"
	ListDrPlansDrPlanTypeStartDrill ListDrPlansDrPlanTypeEnum = "START_DRILL"
	ListDrPlansDrPlanTypeStopDrill  ListDrPlansDrPlanTypeEnum = "STOP_DRILL"
)

var mappingListDrPlansDrPlanTypeEnum = map[string]ListDrPlansDrPlanTypeEnum{
	"SWITCHOVER":  ListDrPlansDrPlanTypeSwitchover,
	"FAILOVER":    ListDrPlansDrPlanTypeFailover,
	"START_DRILL": ListDrPlansDrPlanTypeStartDrill,
	"STOP_DRILL":  ListDrPlansDrPlanTypeStopDrill,
}

var mappingListDrPlansDrPlanTypeEnumLowerCase = map[string]ListDrPlansDrPlanTypeEnum{
	"switchover":  ListDrPlansDrPlanTypeSwitchover,
	"failover":    ListDrPlansDrPlanTypeFailover,
	"start_drill": ListDrPlansDrPlanTypeStartDrill,
	"stop_drill":  ListDrPlansDrPlanTypeStopDrill,
}

// GetListDrPlansDrPlanTypeEnumValues Enumerates the set of values for ListDrPlansDrPlanTypeEnum
func GetListDrPlansDrPlanTypeEnumValues() []ListDrPlansDrPlanTypeEnum {
	values := make([]ListDrPlansDrPlanTypeEnum, 0)
	for _, v := range mappingListDrPlansDrPlanTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrPlansDrPlanTypeEnumStringValues Enumerates the set of values in String for ListDrPlansDrPlanTypeEnum
func GetListDrPlansDrPlanTypeEnumStringValues() []string {
	return []string{
		"SWITCHOVER",
		"FAILOVER",
		"START_DRILL",
		"STOP_DRILL",
	}
}

// GetMappingListDrPlansDrPlanTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrPlansDrPlanTypeEnum(val string) (ListDrPlansDrPlanTypeEnum, bool) {
	enum, ok := mappingListDrPlansDrPlanTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDrPlansSortOrderEnum Enum with underlying type: string
type ListDrPlansSortOrderEnum string

// Set of constants representing the allowable values for ListDrPlansSortOrderEnum
const (
	ListDrPlansSortOrderAsc  ListDrPlansSortOrderEnum = "ASC"
	ListDrPlansSortOrderDesc ListDrPlansSortOrderEnum = "DESC"
)

var mappingListDrPlansSortOrderEnum = map[string]ListDrPlansSortOrderEnum{
	"ASC":  ListDrPlansSortOrderAsc,
	"DESC": ListDrPlansSortOrderDesc,
}

var mappingListDrPlansSortOrderEnumLowerCase = map[string]ListDrPlansSortOrderEnum{
	"asc":  ListDrPlansSortOrderAsc,
	"desc": ListDrPlansSortOrderDesc,
}

// GetListDrPlansSortOrderEnumValues Enumerates the set of values for ListDrPlansSortOrderEnum
func GetListDrPlansSortOrderEnumValues() []ListDrPlansSortOrderEnum {
	values := make([]ListDrPlansSortOrderEnum, 0)
	for _, v := range mappingListDrPlansSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrPlansSortOrderEnumStringValues Enumerates the set of values in String for ListDrPlansSortOrderEnum
func GetListDrPlansSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDrPlansSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrPlansSortOrderEnum(val string) (ListDrPlansSortOrderEnum, bool) {
	enum, ok := mappingListDrPlansSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDrPlansSortByEnum Enum with underlying type: string
type ListDrPlansSortByEnum string

// Set of constants representing the allowable values for ListDrPlansSortByEnum
const (
	ListDrPlansSortByTimecreated ListDrPlansSortByEnum = "timeCreated"
	ListDrPlansSortByDisplayname ListDrPlansSortByEnum = "displayName"
)

var mappingListDrPlansSortByEnum = map[string]ListDrPlansSortByEnum{
	"timeCreated": ListDrPlansSortByTimecreated,
	"displayName": ListDrPlansSortByDisplayname,
}

var mappingListDrPlansSortByEnumLowerCase = map[string]ListDrPlansSortByEnum{
	"timecreated": ListDrPlansSortByTimecreated,
	"displayname": ListDrPlansSortByDisplayname,
}

// GetListDrPlansSortByEnumValues Enumerates the set of values for ListDrPlansSortByEnum
func GetListDrPlansSortByEnumValues() []ListDrPlansSortByEnum {
	values := make([]ListDrPlansSortByEnum, 0)
	for _, v := range mappingListDrPlansSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrPlansSortByEnumStringValues Enumerates the set of values in String for ListDrPlansSortByEnum
func GetListDrPlansSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDrPlansSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrPlansSortByEnum(val string) (ListDrPlansSortByEnum, bool) {
	enum, ok := mappingListDrPlansSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
