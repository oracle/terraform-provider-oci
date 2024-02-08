// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package optimizer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRecommendationsRequest wrapper for the ListRecommendations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListRecommendations.go.html to see an example of how to use ListRecommendationsRequest.
type ListRecommendationsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.
	// Can only be set to true when performing ListCompartments on the tenancy (root compartment).
	CompartmentIdInSubtree *bool `mandatory:"true" contributesTo:"query" name:"compartmentIdInSubtree"`

	// The unique OCID associated with the category.
	CategoryId *string `mandatory:"false" contributesTo:"query" name:"categoryId"`

	// Optional. A filter that returns results that match the category name specified.
	CategoryName *string `mandatory:"false" contributesTo:"query" name:"categoryName"`

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

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListRecommendationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for NAME is ascending. The NAME sort order is case sensitive.
	SortBy ListRecommendationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter that returns results that match the lifecycle state specified.
	LifecycleState ListRecommendationsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter that returns recommendations that match the status specified.
	Status ListRecommendationsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRecommendationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRecommendationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRecommendationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRecommendationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRecommendationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRecommendationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRecommendationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecommendationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRecommendationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecommendationsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListRecommendationsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecommendationsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListRecommendationsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRecommendationsResponse wrapper for the ListRecommendations operation
type ListRecommendationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecommendationCollection instances
	RecommendationCollection `presentIn:"body"`

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

func (response ListRecommendationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRecommendationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRecommendationsSortOrderEnum Enum with underlying type: string
type ListRecommendationsSortOrderEnum string

// Set of constants representing the allowable values for ListRecommendationsSortOrderEnum
const (
	ListRecommendationsSortOrderAsc  ListRecommendationsSortOrderEnum = "ASC"
	ListRecommendationsSortOrderDesc ListRecommendationsSortOrderEnum = "DESC"
)

var mappingListRecommendationsSortOrderEnum = map[string]ListRecommendationsSortOrderEnum{
	"ASC":  ListRecommendationsSortOrderAsc,
	"DESC": ListRecommendationsSortOrderDesc,
}

var mappingListRecommendationsSortOrderEnumLowerCase = map[string]ListRecommendationsSortOrderEnum{
	"asc":  ListRecommendationsSortOrderAsc,
	"desc": ListRecommendationsSortOrderDesc,
}

// GetListRecommendationsSortOrderEnumValues Enumerates the set of values for ListRecommendationsSortOrderEnum
func GetListRecommendationsSortOrderEnumValues() []ListRecommendationsSortOrderEnum {
	values := make([]ListRecommendationsSortOrderEnum, 0)
	for _, v := range mappingListRecommendationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendationsSortOrderEnumStringValues Enumerates the set of values in String for ListRecommendationsSortOrderEnum
func GetListRecommendationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRecommendationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendationsSortOrderEnum(val string) (ListRecommendationsSortOrderEnum, bool) {
	enum, ok := mappingListRecommendationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecommendationsSortByEnum Enum with underlying type: string
type ListRecommendationsSortByEnum string

// Set of constants representing the allowable values for ListRecommendationsSortByEnum
const (
	ListRecommendationsSortByName        ListRecommendationsSortByEnum = "NAME"
	ListRecommendationsSortByTimecreated ListRecommendationsSortByEnum = "TIMECREATED"
)

var mappingListRecommendationsSortByEnum = map[string]ListRecommendationsSortByEnum{
	"NAME":        ListRecommendationsSortByName,
	"TIMECREATED": ListRecommendationsSortByTimecreated,
}

var mappingListRecommendationsSortByEnumLowerCase = map[string]ListRecommendationsSortByEnum{
	"name":        ListRecommendationsSortByName,
	"timecreated": ListRecommendationsSortByTimecreated,
}

// GetListRecommendationsSortByEnumValues Enumerates the set of values for ListRecommendationsSortByEnum
func GetListRecommendationsSortByEnumValues() []ListRecommendationsSortByEnum {
	values := make([]ListRecommendationsSortByEnum, 0)
	for _, v := range mappingListRecommendationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendationsSortByEnumStringValues Enumerates the set of values in String for ListRecommendationsSortByEnum
func GetListRecommendationsSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIMECREATED",
	}
}

// GetMappingListRecommendationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendationsSortByEnum(val string) (ListRecommendationsSortByEnum, bool) {
	enum, ok := mappingListRecommendationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecommendationsLifecycleStateEnum Enum with underlying type: string
type ListRecommendationsLifecycleStateEnum string

// Set of constants representing the allowable values for ListRecommendationsLifecycleStateEnum
const (
	ListRecommendationsLifecycleStateActive    ListRecommendationsLifecycleStateEnum = "ACTIVE"
	ListRecommendationsLifecycleStateFailed    ListRecommendationsLifecycleStateEnum = "FAILED"
	ListRecommendationsLifecycleStateInactive  ListRecommendationsLifecycleStateEnum = "INACTIVE"
	ListRecommendationsLifecycleStateAttaching ListRecommendationsLifecycleStateEnum = "ATTACHING"
	ListRecommendationsLifecycleStateDetaching ListRecommendationsLifecycleStateEnum = "DETACHING"
	ListRecommendationsLifecycleStateDeleting  ListRecommendationsLifecycleStateEnum = "DELETING"
	ListRecommendationsLifecycleStateDeleted   ListRecommendationsLifecycleStateEnum = "DELETED"
	ListRecommendationsLifecycleStateUpdating  ListRecommendationsLifecycleStateEnum = "UPDATING"
	ListRecommendationsLifecycleStateCreating  ListRecommendationsLifecycleStateEnum = "CREATING"
)

var mappingListRecommendationsLifecycleStateEnum = map[string]ListRecommendationsLifecycleStateEnum{
	"ACTIVE":    ListRecommendationsLifecycleStateActive,
	"FAILED":    ListRecommendationsLifecycleStateFailed,
	"INACTIVE":  ListRecommendationsLifecycleStateInactive,
	"ATTACHING": ListRecommendationsLifecycleStateAttaching,
	"DETACHING": ListRecommendationsLifecycleStateDetaching,
	"DELETING":  ListRecommendationsLifecycleStateDeleting,
	"DELETED":   ListRecommendationsLifecycleStateDeleted,
	"UPDATING":  ListRecommendationsLifecycleStateUpdating,
	"CREATING":  ListRecommendationsLifecycleStateCreating,
}

var mappingListRecommendationsLifecycleStateEnumLowerCase = map[string]ListRecommendationsLifecycleStateEnum{
	"active":    ListRecommendationsLifecycleStateActive,
	"failed":    ListRecommendationsLifecycleStateFailed,
	"inactive":  ListRecommendationsLifecycleStateInactive,
	"attaching": ListRecommendationsLifecycleStateAttaching,
	"detaching": ListRecommendationsLifecycleStateDetaching,
	"deleting":  ListRecommendationsLifecycleStateDeleting,
	"deleted":   ListRecommendationsLifecycleStateDeleted,
	"updating":  ListRecommendationsLifecycleStateUpdating,
	"creating":  ListRecommendationsLifecycleStateCreating,
}

// GetListRecommendationsLifecycleStateEnumValues Enumerates the set of values for ListRecommendationsLifecycleStateEnum
func GetListRecommendationsLifecycleStateEnumValues() []ListRecommendationsLifecycleStateEnum {
	values := make([]ListRecommendationsLifecycleStateEnum, 0)
	for _, v := range mappingListRecommendationsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendationsLifecycleStateEnumStringValues Enumerates the set of values in String for ListRecommendationsLifecycleStateEnum
func GetListRecommendationsLifecycleStateEnumStringValues() []string {
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

// GetMappingListRecommendationsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendationsLifecycleStateEnum(val string) (ListRecommendationsLifecycleStateEnum, bool) {
	enum, ok := mappingListRecommendationsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecommendationsStatusEnum Enum with underlying type: string
type ListRecommendationsStatusEnum string

// Set of constants representing the allowable values for ListRecommendationsStatusEnum
const (
	ListRecommendationsStatusPending     ListRecommendationsStatusEnum = "PENDING"
	ListRecommendationsStatusDismissed   ListRecommendationsStatusEnum = "DISMISSED"
	ListRecommendationsStatusPostponed   ListRecommendationsStatusEnum = "POSTPONED"
	ListRecommendationsStatusImplemented ListRecommendationsStatusEnum = "IMPLEMENTED"
)

var mappingListRecommendationsStatusEnum = map[string]ListRecommendationsStatusEnum{
	"PENDING":     ListRecommendationsStatusPending,
	"DISMISSED":   ListRecommendationsStatusDismissed,
	"POSTPONED":   ListRecommendationsStatusPostponed,
	"IMPLEMENTED": ListRecommendationsStatusImplemented,
}

var mappingListRecommendationsStatusEnumLowerCase = map[string]ListRecommendationsStatusEnum{
	"pending":     ListRecommendationsStatusPending,
	"dismissed":   ListRecommendationsStatusDismissed,
	"postponed":   ListRecommendationsStatusPostponed,
	"implemented": ListRecommendationsStatusImplemented,
}

// GetListRecommendationsStatusEnumValues Enumerates the set of values for ListRecommendationsStatusEnum
func GetListRecommendationsStatusEnumValues() []ListRecommendationsStatusEnum {
	values := make([]ListRecommendationsStatusEnum, 0)
	for _, v := range mappingListRecommendationsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendationsStatusEnumStringValues Enumerates the set of values in String for ListRecommendationsStatusEnum
func GetListRecommendationsStatusEnumStringValues() []string {
	return []string{
		"PENDING",
		"DISMISSED",
		"POSTPONED",
		"IMPLEMENTED",
	}
}

// GetMappingListRecommendationsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendationsStatusEnum(val string) (ListRecommendationsStatusEnum, bool) {
	enum, ok := mappingListRecommendationsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
