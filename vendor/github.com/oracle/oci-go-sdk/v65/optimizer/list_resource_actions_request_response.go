// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package optimizer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListResourceActionsRequest wrapper for the ListResourceActions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListResourceActions.go.html to see an example of how to use ListResourceActionsRequest.
type ListResourceActionsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.
	// Can only be set to true when performing ListCompartments on the tenancy (root compartment).
	CompartmentIdInSubtree *bool `mandatory:"true" contributesTo:"query" name:"compartmentIdInSubtree"`

	// The unique OCID associated with the recommendation.
	RecommendationId *string `mandatory:"false" contributesTo:"query" name:"recommendationId"`

	// Optional. A filter that returns results that match the recommendation name specified.
	RecommendationName *string `mandatory:"false" contributesTo:"query" name:"recommendationName"`

	// A list of child tenancies for which the respective data will be returned. Please note that
	// the parent tenancy id can also be included in this list. For example, if there is a parent P with two
	// children A and B, to return results of only parent P and child A, this list should be populated with
	// tenancy id of parent P and child A.
	// If this list contains a tenancy id that isn't part of the organization of parent P, the request will
	// fail. That is, let's say there is an organization with parent P with children A and B, and also one
	// other tenant T that isn't part of the organization. If T is included in the list of
	// childTenancyIds, the request will fail.
	// It is important to note that if you are setting the includeOrganization parameter value as true and
	// also populating the childTenancyIds parameter with a list of child tenancies, the request will fail.
	// The childTenancyIds and includeOrganization should be used exclusively.
	// When using this parameter, please make sure to set the compartmentId with the parent tenancy ID.
	ChildTenancyIds []string `contributesTo:"query" name:"childTenancyIds" collectionFormat:"multi"`

	// When set to true, the data for all child tenancies including the parent is returned. That is, if
	// there is an organization with parent P and children A and B, to return the data for the parent P, child
	// A and child B, this parameter value should be set to true.
	// Please note that this parameter shouldn't be used along with childTenancyIds parameter. If you would like
	// to get results specifically for parent P and only child A, use the childTenancyIds parameter and populate
	// the list with tenancy id of P and A.
	// When using this parameter, please make sure to set the compartmentId with the parent tenancy ID.
	IncludeOrganization *bool `mandatory:"false" contributesTo:"query" name:"includeOrganization"`

	// Optional. A filter that returns results that match the name specified.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Optional. A filter that returns results that match the resource type specified.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListResourceActionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for NAME is ascending. The NAME sort order is case sensitive.
	SortBy ListResourceActionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter that returns results that match the lifecycle state specified.
	LifecycleState ListResourceActionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter that returns recommendations that match the status specified.
	Status ListResourceActionsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// Supplement additional resource information in extended metadata response.
	IncludeResourceMetadata *bool `mandatory:"false" contributesTo:"query" name:"includeResourceMetadata"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourceActionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourceActionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourceActionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourceActionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResourceActionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResourceActionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResourceActionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceActionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResourceActionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceActionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListResourceActionsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceActionsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListResourceActionsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResourceActionsResponse wrapper for the ListResourceActions operation
type ListResourceActionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourceActionCollection instances
	ResourceActionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results exist.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListResourceActionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourceActionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourceActionsSortOrderEnum Enum with underlying type: string
type ListResourceActionsSortOrderEnum string

// Set of constants representing the allowable values for ListResourceActionsSortOrderEnum
const (
	ListResourceActionsSortOrderAsc  ListResourceActionsSortOrderEnum = "ASC"
	ListResourceActionsSortOrderDesc ListResourceActionsSortOrderEnum = "DESC"
)

var mappingListResourceActionsSortOrderEnum = map[string]ListResourceActionsSortOrderEnum{
	"ASC":  ListResourceActionsSortOrderAsc,
	"DESC": ListResourceActionsSortOrderDesc,
}

var mappingListResourceActionsSortOrderEnumLowerCase = map[string]ListResourceActionsSortOrderEnum{
	"asc":  ListResourceActionsSortOrderAsc,
	"desc": ListResourceActionsSortOrderDesc,
}

// GetListResourceActionsSortOrderEnumValues Enumerates the set of values for ListResourceActionsSortOrderEnum
func GetListResourceActionsSortOrderEnumValues() []ListResourceActionsSortOrderEnum {
	values := make([]ListResourceActionsSortOrderEnum, 0)
	for _, v := range mappingListResourceActionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceActionsSortOrderEnumStringValues Enumerates the set of values in String for ListResourceActionsSortOrderEnum
func GetListResourceActionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResourceActionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceActionsSortOrderEnum(val string) (ListResourceActionsSortOrderEnum, bool) {
	enum, ok := mappingListResourceActionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceActionsSortByEnum Enum with underlying type: string
type ListResourceActionsSortByEnum string

// Set of constants representing the allowable values for ListResourceActionsSortByEnum
const (
	ListResourceActionsSortByName        ListResourceActionsSortByEnum = "NAME"
	ListResourceActionsSortByTimecreated ListResourceActionsSortByEnum = "TIMECREATED"
)

var mappingListResourceActionsSortByEnum = map[string]ListResourceActionsSortByEnum{
	"NAME":        ListResourceActionsSortByName,
	"TIMECREATED": ListResourceActionsSortByTimecreated,
}

var mappingListResourceActionsSortByEnumLowerCase = map[string]ListResourceActionsSortByEnum{
	"name":        ListResourceActionsSortByName,
	"timecreated": ListResourceActionsSortByTimecreated,
}

// GetListResourceActionsSortByEnumValues Enumerates the set of values for ListResourceActionsSortByEnum
func GetListResourceActionsSortByEnumValues() []ListResourceActionsSortByEnum {
	values := make([]ListResourceActionsSortByEnum, 0)
	for _, v := range mappingListResourceActionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceActionsSortByEnumStringValues Enumerates the set of values in String for ListResourceActionsSortByEnum
func GetListResourceActionsSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIMECREATED",
	}
}

// GetMappingListResourceActionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceActionsSortByEnum(val string) (ListResourceActionsSortByEnum, bool) {
	enum, ok := mappingListResourceActionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceActionsLifecycleStateEnum Enum with underlying type: string
type ListResourceActionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListResourceActionsLifecycleStateEnum
const (
	ListResourceActionsLifecycleStateActive    ListResourceActionsLifecycleStateEnum = "ACTIVE"
	ListResourceActionsLifecycleStateFailed    ListResourceActionsLifecycleStateEnum = "FAILED"
	ListResourceActionsLifecycleStateInactive  ListResourceActionsLifecycleStateEnum = "INACTIVE"
	ListResourceActionsLifecycleStateAttaching ListResourceActionsLifecycleStateEnum = "ATTACHING"
	ListResourceActionsLifecycleStateDetaching ListResourceActionsLifecycleStateEnum = "DETACHING"
	ListResourceActionsLifecycleStateDeleting  ListResourceActionsLifecycleStateEnum = "DELETING"
	ListResourceActionsLifecycleStateDeleted   ListResourceActionsLifecycleStateEnum = "DELETED"
	ListResourceActionsLifecycleStateUpdating  ListResourceActionsLifecycleStateEnum = "UPDATING"
	ListResourceActionsLifecycleStateCreating  ListResourceActionsLifecycleStateEnum = "CREATING"
)

var mappingListResourceActionsLifecycleStateEnum = map[string]ListResourceActionsLifecycleStateEnum{
	"ACTIVE":    ListResourceActionsLifecycleStateActive,
	"FAILED":    ListResourceActionsLifecycleStateFailed,
	"INACTIVE":  ListResourceActionsLifecycleStateInactive,
	"ATTACHING": ListResourceActionsLifecycleStateAttaching,
	"DETACHING": ListResourceActionsLifecycleStateDetaching,
	"DELETING":  ListResourceActionsLifecycleStateDeleting,
	"DELETED":   ListResourceActionsLifecycleStateDeleted,
	"UPDATING":  ListResourceActionsLifecycleStateUpdating,
	"CREATING":  ListResourceActionsLifecycleStateCreating,
}

var mappingListResourceActionsLifecycleStateEnumLowerCase = map[string]ListResourceActionsLifecycleStateEnum{
	"active":    ListResourceActionsLifecycleStateActive,
	"failed":    ListResourceActionsLifecycleStateFailed,
	"inactive":  ListResourceActionsLifecycleStateInactive,
	"attaching": ListResourceActionsLifecycleStateAttaching,
	"detaching": ListResourceActionsLifecycleStateDetaching,
	"deleting":  ListResourceActionsLifecycleStateDeleting,
	"deleted":   ListResourceActionsLifecycleStateDeleted,
	"updating":  ListResourceActionsLifecycleStateUpdating,
	"creating":  ListResourceActionsLifecycleStateCreating,
}

// GetListResourceActionsLifecycleStateEnumValues Enumerates the set of values for ListResourceActionsLifecycleStateEnum
func GetListResourceActionsLifecycleStateEnumValues() []ListResourceActionsLifecycleStateEnum {
	values := make([]ListResourceActionsLifecycleStateEnum, 0)
	for _, v := range mappingListResourceActionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceActionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListResourceActionsLifecycleStateEnum
func GetListResourceActionsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"FAILED",
		"INACTIVE",
		"ATTACHING",
		"DETACHING",
		"DELETING",
		"DELETED",
		"UPDATING",
		"CREATING",
	}
}

// GetMappingListResourceActionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceActionsLifecycleStateEnum(val string) (ListResourceActionsLifecycleStateEnum, bool) {
	enum, ok := mappingListResourceActionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceActionsStatusEnum Enum with underlying type: string
type ListResourceActionsStatusEnum string

// Set of constants representing the allowable values for ListResourceActionsStatusEnum
const (
	ListResourceActionsStatusPending     ListResourceActionsStatusEnum = "PENDING"
	ListResourceActionsStatusDismissed   ListResourceActionsStatusEnum = "DISMISSED"
	ListResourceActionsStatusPostponed   ListResourceActionsStatusEnum = "POSTPONED"
	ListResourceActionsStatusImplemented ListResourceActionsStatusEnum = "IMPLEMENTED"
)

var mappingListResourceActionsStatusEnum = map[string]ListResourceActionsStatusEnum{
	"PENDING":     ListResourceActionsStatusPending,
	"DISMISSED":   ListResourceActionsStatusDismissed,
	"POSTPONED":   ListResourceActionsStatusPostponed,
	"IMPLEMENTED": ListResourceActionsStatusImplemented,
}

var mappingListResourceActionsStatusEnumLowerCase = map[string]ListResourceActionsStatusEnum{
	"pending":     ListResourceActionsStatusPending,
	"dismissed":   ListResourceActionsStatusDismissed,
	"postponed":   ListResourceActionsStatusPostponed,
	"implemented": ListResourceActionsStatusImplemented,
}

// GetListResourceActionsStatusEnumValues Enumerates the set of values for ListResourceActionsStatusEnum
func GetListResourceActionsStatusEnumValues() []ListResourceActionsStatusEnum {
	values := make([]ListResourceActionsStatusEnum, 0)
	for _, v := range mappingListResourceActionsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceActionsStatusEnumStringValues Enumerates the set of values in String for ListResourceActionsStatusEnum
func GetListResourceActionsStatusEnumStringValues() []string {
	return []string{
		"PENDING",
		"DISMISSED",
		"POSTPONED",
		"IMPLEMENTED",
	}
}

// GetMappingListResourceActionsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceActionsStatusEnum(val string) (ListResourceActionsStatusEnum, bool) {
	enum, ok := mappingListResourceActionsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
