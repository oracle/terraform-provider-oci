// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMigrationPlansRequest wrapper for the ListMigrationPlans operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudmigrations/ListMigrationPlans.go.html to see an example of how to use ListMigrationPlansRequest.
type ListMigrationPlansRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique migration identifier
	MigrationId *string `mandatory:"false" contributesTo:"query" name:"migrationId"`

	// A filter to return only resources that match the entire given display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique migration plan identifier
	MigrationPlanId *string `mandatory:"false" contributesTo:"query" name:"migrationPlanId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of the previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The current state of the migration plan.
	LifecycleState MigrationPlanLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMigrationPlansSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. The default order for 'timeCreated' is descending. The default order for 'displayName' is ascending.
	SortBy ListMigrationPlansSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMigrationPlansRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMigrationPlansRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMigrationPlansRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMigrationPlansRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMigrationPlansRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMigrationPlanLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMigrationPlanLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMigrationPlansSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMigrationPlansSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMigrationPlansSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMigrationPlansSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMigrationPlansResponse wrapper for the ListMigrationPlans operation
type ListMigrationPlansResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MigrationPlanCollection instances
	MigrationPlanCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMigrationPlansResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMigrationPlansResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMigrationPlansSortOrderEnum Enum with underlying type: string
type ListMigrationPlansSortOrderEnum string

// Set of constants representing the allowable values for ListMigrationPlansSortOrderEnum
const (
	ListMigrationPlansSortOrderAsc  ListMigrationPlansSortOrderEnum = "ASC"
	ListMigrationPlansSortOrderDesc ListMigrationPlansSortOrderEnum = "DESC"
)

var mappingListMigrationPlansSortOrderEnum = map[string]ListMigrationPlansSortOrderEnum{
	"ASC":  ListMigrationPlansSortOrderAsc,
	"DESC": ListMigrationPlansSortOrderDesc,
}

var mappingListMigrationPlansSortOrderEnumLowerCase = map[string]ListMigrationPlansSortOrderEnum{
	"asc":  ListMigrationPlansSortOrderAsc,
	"desc": ListMigrationPlansSortOrderDesc,
}

// GetListMigrationPlansSortOrderEnumValues Enumerates the set of values for ListMigrationPlansSortOrderEnum
func GetListMigrationPlansSortOrderEnumValues() []ListMigrationPlansSortOrderEnum {
	values := make([]ListMigrationPlansSortOrderEnum, 0)
	for _, v := range mappingListMigrationPlansSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationPlansSortOrderEnumStringValues Enumerates the set of values in String for ListMigrationPlansSortOrderEnum
func GetListMigrationPlansSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMigrationPlansSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationPlansSortOrderEnum(val string) (ListMigrationPlansSortOrderEnum, bool) {
	enum, ok := mappingListMigrationPlansSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMigrationPlansSortByEnum Enum with underlying type: string
type ListMigrationPlansSortByEnum string

// Set of constants representing the allowable values for ListMigrationPlansSortByEnum
const (
	ListMigrationPlansSortByTimecreated ListMigrationPlansSortByEnum = "timeCreated"
	ListMigrationPlansSortByDisplayname ListMigrationPlansSortByEnum = "displayName"
)

var mappingListMigrationPlansSortByEnum = map[string]ListMigrationPlansSortByEnum{
	"timeCreated": ListMigrationPlansSortByTimecreated,
	"displayName": ListMigrationPlansSortByDisplayname,
}

var mappingListMigrationPlansSortByEnumLowerCase = map[string]ListMigrationPlansSortByEnum{
	"timecreated": ListMigrationPlansSortByTimecreated,
	"displayname": ListMigrationPlansSortByDisplayname,
}

// GetListMigrationPlansSortByEnumValues Enumerates the set of values for ListMigrationPlansSortByEnum
func GetListMigrationPlansSortByEnumValues() []ListMigrationPlansSortByEnum {
	values := make([]ListMigrationPlansSortByEnum, 0)
	for _, v := range mappingListMigrationPlansSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationPlansSortByEnumStringValues Enumerates the set of values in String for ListMigrationPlansSortByEnum
func GetListMigrationPlansSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMigrationPlansSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationPlansSortByEnum(val string) (ListMigrationPlansSortByEnum, bool) {
	enum, ok := mappingListMigrationPlansSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
