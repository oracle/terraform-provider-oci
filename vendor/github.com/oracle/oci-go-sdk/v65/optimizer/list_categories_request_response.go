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

// ListCategoriesRequest wrapper for the ListCategories operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListCategories.go.html to see an example of how to use ListCategoriesRequest.
type ListCategoriesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.
	// Can only be set to true when performing ListCompartments on the tenancy (root compartment).
	CompartmentIdInSubtree *bool `mandatory:"true" contributesTo:"query" name:"compartmentIdInSubtree"`

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
	SortOrder ListCategoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for NAME is ascending. The NAME sort order is case sensitive.
	SortBy ListCategoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter that returns results that match the lifecycle state specified.
	LifecycleState ListCategoriesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCategoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCategoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCategoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCategoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCategoriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCategoriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCategoriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCategoriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCategoriesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCategoriesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListCategoriesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCategoriesResponse wrapper for the ListCategories operation
type ListCategoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CategoryCollection instances
	CategoryCollection `presentIn:"body"`

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

func (response ListCategoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCategoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCategoriesSortOrderEnum Enum with underlying type: string
type ListCategoriesSortOrderEnum string

// Set of constants representing the allowable values for ListCategoriesSortOrderEnum
const (
	ListCategoriesSortOrderAsc  ListCategoriesSortOrderEnum = "ASC"
	ListCategoriesSortOrderDesc ListCategoriesSortOrderEnum = "DESC"
)

var mappingListCategoriesSortOrderEnum = map[string]ListCategoriesSortOrderEnum{
	"ASC":  ListCategoriesSortOrderAsc,
	"DESC": ListCategoriesSortOrderDesc,
}

var mappingListCategoriesSortOrderEnumLowerCase = map[string]ListCategoriesSortOrderEnum{
	"asc":  ListCategoriesSortOrderAsc,
	"desc": ListCategoriesSortOrderDesc,
}

// GetListCategoriesSortOrderEnumValues Enumerates the set of values for ListCategoriesSortOrderEnum
func GetListCategoriesSortOrderEnumValues() []ListCategoriesSortOrderEnum {
	values := make([]ListCategoriesSortOrderEnum, 0)
	for _, v := range mappingListCategoriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCategoriesSortOrderEnumStringValues Enumerates the set of values in String for ListCategoriesSortOrderEnum
func GetListCategoriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCategoriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCategoriesSortOrderEnum(val string) (ListCategoriesSortOrderEnum, bool) {
	enum, ok := mappingListCategoriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCategoriesSortByEnum Enum with underlying type: string
type ListCategoriesSortByEnum string

// Set of constants representing the allowable values for ListCategoriesSortByEnum
const (
	ListCategoriesSortByName        ListCategoriesSortByEnum = "NAME"
	ListCategoriesSortByTimecreated ListCategoriesSortByEnum = "TIMECREATED"
)

var mappingListCategoriesSortByEnum = map[string]ListCategoriesSortByEnum{
	"NAME":        ListCategoriesSortByName,
	"TIMECREATED": ListCategoriesSortByTimecreated,
}

var mappingListCategoriesSortByEnumLowerCase = map[string]ListCategoriesSortByEnum{
	"name":        ListCategoriesSortByName,
	"timecreated": ListCategoriesSortByTimecreated,
}

// GetListCategoriesSortByEnumValues Enumerates the set of values for ListCategoriesSortByEnum
func GetListCategoriesSortByEnumValues() []ListCategoriesSortByEnum {
	values := make([]ListCategoriesSortByEnum, 0)
	for _, v := range mappingListCategoriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCategoriesSortByEnumStringValues Enumerates the set of values in String for ListCategoriesSortByEnum
func GetListCategoriesSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIMECREATED",
	}
}

// GetMappingListCategoriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCategoriesSortByEnum(val string) (ListCategoriesSortByEnum, bool) {
	enum, ok := mappingListCategoriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCategoriesLifecycleStateEnum Enum with underlying type: string
type ListCategoriesLifecycleStateEnum string

// Set of constants representing the allowable values for ListCategoriesLifecycleStateEnum
const (
	ListCategoriesLifecycleStateActive    ListCategoriesLifecycleStateEnum = "ACTIVE"
	ListCategoriesLifecycleStateFailed    ListCategoriesLifecycleStateEnum = "FAILED"
	ListCategoriesLifecycleStateInactive  ListCategoriesLifecycleStateEnum = "INACTIVE"
	ListCategoriesLifecycleStateAttaching ListCategoriesLifecycleStateEnum = "ATTACHING"
	ListCategoriesLifecycleStateDetaching ListCategoriesLifecycleStateEnum = "DETACHING"
	ListCategoriesLifecycleStateDeleting  ListCategoriesLifecycleStateEnum = "DELETING"
	ListCategoriesLifecycleStateDeleted   ListCategoriesLifecycleStateEnum = "DELETED"
	ListCategoriesLifecycleStateUpdating  ListCategoriesLifecycleStateEnum = "UPDATING"
	ListCategoriesLifecycleStateCreating  ListCategoriesLifecycleStateEnum = "CREATING"
)

var mappingListCategoriesLifecycleStateEnum = map[string]ListCategoriesLifecycleStateEnum{
	"ACTIVE":    ListCategoriesLifecycleStateActive,
	"FAILED":    ListCategoriesLifecycleStateFailed,
	"INACTIVE":  ListCategoriesLifecycleStateInactive,
	"ATTACHING": ListCategoriesLifecycleStateAttaching,
	"DETACHING": ListCategoriesLifecycleStateDetaching,
	"DELETING":  ListCategoriesLifecycleStateDeleting,
	"DELETED":   ListCategoriesLifecycleStateDeleted,
	"UPDATING":  ListCategoriesLifecycleStateUpdating,
	"CREATING":  ListCategoriesLifecycleStateCreating,
}

var mappingListCategoriesLifecycleStateEnumLowerCase = map[string]ListCategoriesLifecycleStateEnum{
	"active":    ListCategoriesLifecycleStateActive,
	"failed":    ListCategoriesLifecycleStateFailed,
	"inactive":  ListCategoriesLifecycleStateInactive,
	"attaching": ListCategoriesLifecycleStateAttaching,
	"detaching": ListCategoriesLifecycleStateDetaching,
	"deleting":  ListCategoriesLifecycleStateDeleting,
	"deleted":   ListCategoriesLifecycleStateDeleted,
	"updating":  ListCategoriesLifecycleStateUpdating,
	"creating":  ListCategoriesLifecycleStateCreating,
}

// GetListCategoriesLifecycleStateEnumValues Enumerates the set of values for ListCategoriesLifecycleStateEnum
func GetListCategoriesLifecycleStateEnumValues() []ListCategoriesLifecycleStateEnum {
	values := make([]ListCategoriesLifecycleStateEnum, 0)
	for _, v := range mappingListCategoriesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListCategoriesLifecycleStateEnumStringValues Enumerates the set of values in String for ListCategoriesLifecycleStateEnum
func GetListCategoriesLifecycleStateEnumStringValues() []string {
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

// GetMappingListCategoriesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCategoriesLifecycleStateEnum(val string) (ListCategoriesLifecycleStateEnum, bool) {
	enum, ok := mappingListCategoriesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
