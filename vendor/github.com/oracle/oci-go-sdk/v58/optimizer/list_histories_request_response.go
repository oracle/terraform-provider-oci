// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package optimizer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListHistoriesRequest wrapper for the ListHistories operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListHistories.go.html to see an example of how to use ListHistoriesRequest.
type ListHistoriesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.
	// Can only be set to true when performing ListCompartments on the tenancy (root compartment).
	CompartmentIdInSubtree *bool `mandatory:"true" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Optional. A filter that returns results that match the name specified.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Optional. A filter that returns results that match the recommendation name specified.
	RecommendationName *string `mandatory:"false" contributesTo:"query" name:"recommendationName"`

	// The unique OCID associated with the recommendation.
	RecommendationId *string `mandatory:"false" contributesTo:"query" name:"recommendationId"`

	// Optional. A filter that returns results that match the resource type specified.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListHistoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for NAME is ascending. The NAME sort order is case sensitive.
	SortBy ListHistoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter that returns results that match the lifecycle state specified.
	LifecycleState ListHistoriesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter that returns recommendations that match the status specified.
	Status ListHistoriesStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListHistoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListHistoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListHistoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListHistoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListHistoriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListHistoriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListHistoriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHistoriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListHistoriesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHistoriesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListHistoriesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHistoriesStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListHistoriesStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListHistoriesResponse wrapper for the ListHistories operation
type ListHistoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of HistoryCollection instances
	HistoryCollection `presentIn:"body"`

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

func (response ListHistoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListHistoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListHistoriesSortOrderEnum Enum with underlying type: string
type ListHistoriesSortOrderEnum string

// Set of constants representing the allowable values for ListHistoriesSortOrderEnum
const (
	ListHistoriesSortOrderAsc  ListHistoriesSortOrderEnum = "ASC"
	ListHistoriesSortOrderDesc ListHistoriesSortOrderEnum = "DESC"
)

var mappingListHistoriesSortOrderEnum = map[string]ListHistoriesSortOrderEnum{
	"ASC":  ListHistoriesSortOrderAsc,
	"DESC": ListHistoriesSortOrderDesc,
}

// GetListHistoriesSortOrderEnumValues Enumerates the set of values for ListHistoriesSortOrderEnum
func GetListHistoriesSortOrderEnumValues() []ListHistoriesSortOrderEnum {
	values := make([]ListHistoriesSortOrderEnum, 0)
	for _, v := range mappingListHistoriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListHistoriesSortOrderEnumStringValues Enumerates the set of values in String for ListHistoriesSortOrderEnum
func GetListHistoriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListHistoriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHistoriesSortOrderEnum(val string) (ListHistoriesSortOrderEnum, bool) {
	mappingListHistoriesSortOrderEnumIgnoreCase := make(map[string]ListHistoriesSortOrderEnum)
	for k, v := range mappingListHistoriesSortOrderEnum {
		mappingListHistoriesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListHistoriesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListHistoriesSortByEnum Enum with underlying type: string
type ListHistoriesSortByEnum string

// Set of constants representing the allowable values for ListHistoriesSortByEnum
const (
	ListHistoriesSortByName        ListHistoriesSortByEnum = "NAME"
	ListHistoriesSortByTimecreated ListHistoriesSortByEnum = "TIMECREATED"
)

var mappingListHistoriesSortByEnum = map[string]ListHistoriesSortByEnum{
	"NAME":        ListHistoriesSortByName,
	"TIMECREATED": ListHistoriesSortByTimecreated,
}

// GetListHistoriesSortByEnumValues Enumerates the set of values for ListHistoriesSortByEnum
func GetListHistoriesSortByEnumValues() []ListHistoriesSortByEnum {
	values := make([]ListHistoriesSortByEnum, 0)
	for _, v := range mappingListHistoriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListHistoriesSortByEnumStringValues Enumerates the set of values in String for ListHistoriesSortByEnum
func GetListHistoriesSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIMECREATED",
	}
}

// GetMappingListHistoriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHistoriesSortByEnum(val string) (ListHistoriesSortByEnum, bool) {
	mappingListHistoriesSortByEnumIgnoreCase := make(map[string]ListHistoriesSortByEnum)
	for k, v := range mappingListHistoriesSortByEnum {
		mappingListHistoriesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListHistoriesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListHistoriesLifecycleStateEnum Enum with underlying type: string
type ListHistoriesLifecycleStateEnum string

// Set of constants representing the allowable values for ListHistoriesLifecycleStateEnum
const (
	ListHistoriesLifecycleStateActive    ListHistoriesLifecycleStateEnum = "ACTIVE"
	ListHistoriesLifecycleStateFailed    ListHistoriesLifecycleStateEnum = "FAILED"
	ListHistoriesLifecycleStateInactive  ListHistoriesLifecycleStateEnum = "INACTIVE"
	ListHistoriesLifecycleStateAttaching ListHistoriesLifecycleStateEnum = "ATTACHING"
	ListHistoriesLifecycleStateDetaching ListHistoriesLifecycleStateEnum = "DETACHING"
	ListHistoriesLifecycleStateDeleting  ListHistoriesLifecycleStateEnum = "DELETING"
	ListHistoriesLifecycleStateDeleted   ListHistoriesLifecycleStateEnum = "DELETED"
	ListHistoriesLifecycleStateUpdating  ListHistoriesLifecycleStateEnum = "UPDATING"
	ListHistoriesLifecycleStateCreating  ListHistoriesLifecycleStateEnum = "CREATING"
)

var mappingListHistoriesLifecycleStateEnum = map[string]ListHistoriesLifecycleStateEnum{
	"ACTIVE":    ListHistoriesLifecycleStateActive,
	"FAILED":    ListHistoriesLifecycleStateFailed,
	"INACTIVE":  ListHistoriesLifecycleStateInactive,
	"ATTACHING": ListHistoriesLifecycleStateAttaching,
	"DETACHING": ListHistoriesLifecycleStateDetaching,
	"DELETING":  ListHistoriesLifecycleStateDeleting,
	"DELETED":   ListHistoriesLifecycleStateDeleted,
	"UPDATING":  ListHistoriesLifecycleStateUpdating,
	"CREATING":  ListHistoriesLifecycleStateCreating,
}

// GetListHistoriesLifecycleStateEnumValues Enumerates the set of values for ListHistoriesLifecycleStateEnum
func GetListHistoriesLifecycleStateEnumValues() []ListHistoriesLifecycleStateEnum {
	values := make([]ListHistoriesLifecycleStateEnum, 0)
	for _, v := range mappingListHistoriesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListHistoriesLifecycleStateEnumStringValues Enumerates the set of values in String for ListHistoriesLifecycleStateEnum
func GetListHistoriesLifecycleStateEnumStringValues() []string {
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

// GetMappingListHistoriesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHistoriesLifecycleStateEnum(val string) (ListHistoriesLifecycleStateEnum, bool) {
	mappingListHistoriesLifecycleStateEnumIgnoreCase := make(map[string]ListHistoriesLifecycleStateEnum)
	for k, v := range mappingListHistoriesLifecycleStateEnum {
		mappingListHistoriesLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListHistoriesLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListHistoriesStatusEnum Enum with underlying type: string
type ListHistoriesStatusEnum string

// Set of constants representing the allowable values for ListHistoriesStatusEnum
const (
	ListHistoriesStatusPending     ListHistoriesStatusEnum = "PENDING"
	ListHistoriesStatusDismissed   ListHistoriesStatusEnum = "DISMISSED"
	ListHistoriesStatusPostponed   ListHistoriesStatusEnum = "POSTPONED"
	ListHistoriesStatusImplemented ListHistoriesStatusEnum = "IMPLEMENTED"
)

var mappingListHistoriesStatusEnum = map[string]ListHistoriesStatusEnum{
	"PENDING":     ListHistoriesStatusPending,
	"DISMISSED":   ListHistoriesStatusDismissed,
	"POSTPONED":   ListHistoriesStatusPostponed,
	"IMPLEMENTED": ListHistoriesStatusImplemented,
}

// GetListHistoriesStatusEnumValues Enumerates the set of values for ListHistoriesStatusEnum
func GetListHistoriesStatusEnumValues() []ListHistoriesStatusEnum {
	values := make([]ListHistoriesStatusEnum, 0)
	for _, v := range mappingListHistoriesStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListHistoriesStatusEnumStringValues Enumerates the set of values in String for ListHistoriesStatusEnum
func GetListHistoriesStatusEnumStringValues() []string {
	return []string{
		"PENDING",
		"DISMISSED",
		"POSTPONED",
		"IMPLEMENTED",
	}
}

// GetMappingListHistoriesStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHistoriesStatusEnum(val string) (ListHistoriesStatusEnum, bool) {
	mappingListHistoriesStatusEnumIgnoreCase := make(map[string]ListHistoriesStatusEnum)
	for k, v := range mappingListHistoriesStatusEnum {
		mappingListHistoriesStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListHistoriesStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
