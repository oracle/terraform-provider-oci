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

// ListEnrollmentStatusesRequest wrapper for the ListEnrollmentStatuses operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListEnrollmentStatuses.go.html to see an example of how to use ListEnrollmentStatusesRequest.
type ListEnrollmentStatusesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListEnrollmentStatusesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for NAME is ascending. The NAME sort order is case sensitive.
	SortBy ListEnrollmentStatusesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter that returns results that match the lifecycle state specified.
	LifecycleState ListEnrollmentStatusesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter that returns results that match the Cloud Advisor enrollment status specified.
	Status ListEnrollmentStatusesStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEnrollmentStatusesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEnrollmentStatusesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEnrollmentStatusesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEnrollmentStatusesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEnrollmentStatusesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEnrollmentStatusesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEnrollmentStatusesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEnrollmentStatusesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEnrollmentStatusesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEnrollmentStatusesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListEnrollmentStatusesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEnrollmentStatusesStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListEnrollmentStatusesStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEnrollmentStatusesResponse wrapper for the ListEnrollmentStatuses operation
type ListEnrollmentStatusesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EnrollmentStatusCollection instances
	EnrollmentStatusCollection `presentIn:"body"`

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

func (response ListEnrollmentStatusesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEnrollmentStatusesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEnrollmentStatusesSortOrderEnum Enum with underlying type: string
type ListEnrollmentStatusesSortOrderEnum string

// Set of constants representing the allowable values for ListEnrollmentStatusesSortOrderEnum
const (
	ListEnrollmentStatusesSortOrderAsc  ListEnrollmentStatusesSortOrderEnum = "ASC"
	ListEnrollmentStatusesSortOrderDesc ListEnrollmentStatusesSortOrderEnum = "DESC"
)

var mappingListEnrollmentStatusesSortOrderEnum = map[string]ListEnrollmentStatusesSortOrderEnum{
	"ASC":  ListEnrollmentStatusesSortOrderAsc,
	"DESC": ListEnrollmentStatusesSortOrderDesc,
}

var mappingListEnrollmentStatusesSortOrderEnumLowerCase = map[string]ListEnrollmentStatusesSortOrderEnum{
	"asc":  ListEnrollmentStatusesSortOrderAsc,
	"desc": ListEnrollmentStatusesSortOrderDesc,
}

// GetListEnrollmentStatusesSortOrderEnumValues Enumerates the set of values for ListEnrollmentStatusesSortOrderEnum
func GetListEnrollmentStatusesSortOrderEnumValues() []ListEnrollmentStatusesSortOrderEnum {
	values := make([]ListEnrollmentStatusesSortOrderEnum, 0)
	for _, v := range mappingListEnrollmentStatusesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEnrollmentStatusesSortOrderEnumStringValues Enumerates the set of values in String for ListEnrollmentStatusesSortOrderEnum
func GetListEnrollmentStatusesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEnrollmentStatusesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEnrollmentStatusesSortOrderEnum(val string) (ListEnrollmentStatusesSortOrderEnum, bool) {
	enum, ok := mappingListEnrollmentStatusesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEnrollmentStatusesSortByEnum Enum with underlying type: string
type ListEnrollmentStatusesSortByEnum string

// Set of constants representing the allowable values for ListEnrollmentStatusesSortByEnum
const (
	ListEnrollmentStatusesSortByName        ListEnrollmentStatusesSortByEnum = "NAME"
	ListEnrollmentStatusesSortByTimecreated ListEnrollmentStatusesSortByEnum = "TIMECREATED"
)

var mappingListEnrollmentStatusesSortByEnum = map[string]ListEnrollmentStatusesSortByEnum{
	"NAME":        ListEnrollmentStatusesSortByName,
	"TIMECREATED": ListEnrollmentStatusesSortByTimecreated,
}

var mappingListEnrollmentStatusesSortByEnumLowerCase = map[string]ListEnrollmentStatusesSortByEnum{
	"name":        ListEnrollmentStatusesSortByName,
	"timecreated": ListEnrollmentStatusesSortByTimecreated,
}

// GetListEnrollmentStatusesSortByEnumValues Enumerates the set of values for ListEnrollmentStatusesSortByEnum
func GetListEnrollmentStatusesSortByEnumValues() []ListEnrollmentStatusesSortByEnum {
	values := make([]ListEnrollmentStatusesSortByEnum, 0)
	for _, v := range mappingListEnrollmentStatusesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEnrollmentStatusesSortByEnumStringValues Enumerates the set of values in String for ListEnrollmentStatusesSortByEnum
func GetListEnrollmentStatusesSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIMECREATED",
	}
}

// GetMappingListEnrollmentStatusesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEnrollmentStatusesSortByEnum(val string) (ListEnrollmentStatusesSortByEnum, bool) {
	enum, ok := mappingListEnrollmentStatusesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEnrollmentStatusesLifecycleStateEnum Enum with underlying type: string
type ListEnrollmentStatusesLifecycleStateEnum string

// Set of constants representing the allowable values for ListEnrollmentStatusesLifecycleStateEnum
const (
	ListEnrollmentStatusesLifecycleStateActive    ListEnrollmentStatusesLifecycleStateEnum = "ACTIVE"
	ListEnrollmentStatusesLifecycleStateFailed    ListEnrollmentStatusesLifecycleStateEnum = "FAILED"
	ListEnrollmentStatusesLifecycleStateInactive  ListEnrollmentStatusesLifecycleStateEnum = "INACTIVE"
	ListEnrollmentStatusesLifecycleStateAttaching ListEnrollmentStatusesLifecycleStateEnum = "ATTACHING"
	ListEnrollmentStatusesLifecycleStateDetaching ListEnrollmentStatusesLifecycleStateEnum = "DETACHING"
	ListEnrollmentStatusesLifecycleStateDeleting  ListEnrollmentStatusesLifecycleStateEnum = "DELETING"
	ListEnrollmentStatusesLifecycleStateDeleted   ListEnrollmentStatusesLifecycleStateEnum = "DELETED"
	ListEnrollmentStatusesLifecycleStateUpdating  ListEnrollmentStatusesLifecycleStateEnum = "UPDATING"
	ListEnrollmentStatusesLifecycleStateCreating  ListEnrollmentStatusesLifecycleStateEnum = "CREATING"
)

var mappingListEnrollmentStatusesLifecycleStateEnum = map[string]ListEnrollmentStatusesLifecycleStateEnum{
	"ACTIVE":    ListEnrollmentStatusesLifecycleStateActive,
	"FAILED":    ListEnrollmentStatusesLifecycleStateFailed,
	"INACTIVE":  ListEnrollmentStatusesLifecycleStateInactive,
	"ATTACHING": ListEnrollmentStatusesLifecycleStateAttaching,
	"DETACHING": ListEnrollmentStatusesLifecycleStateDetaching,
	"DELETING":  ListEnrollmentStatusesLifecycleStateDeleting,
	"DELETED":   ListEnrollmentStatusesLifecycleStateDeleted,
	"UPDATING":  ListEnrollmentStatusesLifecycleStateUpdating,
	"CREATING":  ListEnrollmentStatusesLifecycleStateCreating,
}

var mappingListEnrollmentStatusesLifecycleStateEnumLowerCase = map[string]ListEnrollmentStatusesLifecycleStateEnum{
	"active":    ListEnrollmentStatusesLifecycleStateActive,
	"failed":    ListEnrollmentStatusesLifecycleStateFailed,
	"inactive":  ListEnrollmentStatusesLifecycleStateInactive,
	"attaching": ListEnrollmentStatusesLifecycleStateAttaching,
	"detaching": ListEnrollmentStatusesLifecycleStateDetaching,
	"deleting":  ListEnrollmentStatusesLifecycleStateDeleting,
	"deleted":   ListEnrollmentStatusesLifecycleStateDeleted,
	"updating":  ListEnrollmentStatusesLifecycleStateUpdating,
	"creating":  ListEnrollmentStatusesLifecycleStateCreating,
}

// GetListEnrollmentStatusesLifecycleStateEnumValues Enumerates the set of values for ListEnrollmentStatusesLifecycleStateEnum
func GetListEnrollmentStatusesLifecycleStateEnumValues() []ListEnrollmentStatusesLifecycleStateEnum {
	values := make([]ListEnrollmentStatusesLifecycleStateEnum, 0)
	for _, v := range mappingListEnrollmentStatusesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListEnrollmentStatusesLifecycleStateEnumStringValues Enumerates the set of values in String for ListEnrollmentStatusesLifecycleStateEnum
func GetListEnrollmentStatusesLifecycleStateEnumStringValues() []string {
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

// GetMappingListEnrollmentStatusesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEnrollmentStatusesLifecycleStateEnum(val string) (ListEnrollmentStatusesLifecycleStateEnum, bool) {
	enum, ok := mappingListEnrollmentStatusesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEnrollmentStatusesStatusEnum Enum with underlying type: string
type ListEnrollmentStatusesStatusEnum string

// Set of constants representing the allowable values for ListEnrollmentStatusesStatusEnum
const (
	ListEnrollmentStatusesStatusActive   ListEnrollmentStatusesStatusEnum = "ACTIVE"
	ListEnrollmentStatusesStatusInactive ListEnrollmentStatusesStatusEnum = "INACTIVE"
)

var mappingListEnrollmentStatusesStatusEnum = map[string]ListEnrollmentStatusesStatusEnum{
	"ACTIVE":   ListEnrollmentStatusesStatusActive,
	"INACTIVE": ListEnrollmentStatusesStatusInactive,
}

var mappingListEnrollmentStatusesStatusEnumLowerCase = map[string]ListEnrollmentStatusesStatusEnum{
	"active":   ListEnrollmentStatusesStatusActive,
	"inactive": ListEnrollmentStatusesStatusInactive,
}

// GetListEnrollmentStatusesStatusEnumValues Enumerates the set of values for ListEnrollmentStatusesStatusEnum
func GetListEnrollmentStatusesStatusEnumValues() []ListEnrollmentStatusesStatusEnum {
	values := make([]ListEnrollmentStatusesStatusEnum, 0)
	for _, v := range mappingListEnrollmentStatusesStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListEnrollmentStatusesStatusEnumStringValues Enumerates the set of values in String for ListEnrollmentStatusesStatusEnum
func GetListEnrollmentStatusesStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingListEnrollmentStatusesStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEnrollmentStatusesStatusEnum(val string) (ListEnrollmentStatusesStatusEnum, bool) {
	enum, ok := mappingListEnrollmentStatusesStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
