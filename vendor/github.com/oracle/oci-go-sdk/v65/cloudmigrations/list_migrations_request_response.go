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

// ListMigrationsRequest wrapper for the ListMigrations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudmigrations/ListMigrations.go.html to see an example of how to use ListMigrationsRequest.
type ListMigrationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources where the resource's lifecycle state matches the given lifecycle state.
	LifecycleState MigrationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire given display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique migration identifier
	MigrationId *string `mandatory:"false" contributesTo:"query" name:"migrationId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of the previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMigrationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. The default order for 'timeCreated' is descending. The default order for 'displayName' is ascending.
	SortBy ListMigrationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMigrationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMigrationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMigrationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMigrationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMigrationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMigrationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMigrationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMigrationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMigrationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMigrationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMigrationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMigrationsResponse wrapper for the ListMigrations operation
type ListMigrationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MigrationCollection instances
	MigrationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMigrationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMigrationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMigrationsSortOrderEnum Enum with underlying type: string
type ListMigrationsSortOrderEnum string

// Set of constants representing the allowable values for ListMigrationsSortOrderEnum
const (
	ListMigrationsSortOrderAsc  ListMigrationsSortOrderEnum = "ASC"
	ListMigrationsSortOrderDesc ListMigrationsSortOrderEnum = "DESC"
)

var mappingListMigrationsSortOrderEnum = map[string]ListMigrationsSortOrderEnum{
	"ASC":  ListMigrationsSortOrderAsc,
	"DESC": ListMigrationsSortOrderDesc,
}

var mappingListMigrationsSortOrderEnumLowerCase = map[string]ListMigrationsSortOrderEnum{
	"asc":  ListMigrationsSortOrderAsc,
	"desc": ListMigrationsSortOrderDesc,
}

// GetListMigrationsSortOrderEnumValues Enumerates the set of values for ListMigrationsSortOrderEnum
func GetListMigrationsSortOrderEnumValues() []ListMigrationsSortOrderEnum {
	values := make([]ListMigrationsSortOrderEnum, 0)
	for _, v := range mappingListMigrationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationsSortOrderEnumStringValues Enumerates the set of values in String for ListMigrationsSortOrderEnum
func GetListMigrationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMigrationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationsSortOrderEnum(val string) (ListMigrationsSortOrderEnum, bool) {
	enum, ok := mappingListMigrationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMigrationsSortByEnum Enum with underlying type: string
type ListMigrationsSortByEnum string

// Set of constants representing the allowable values for ListMigrationsSortByEnum
const (
	ListMigrationsSortByTimecreated ListMigrationsSortByEnum = "timeCreated"
	ListMigrationsSortByDisplayname ListMigrationsSortByEnum = "displayName"
)

var mappingListMigrationsSortByEnum = map[string]ListMigrationsSortByEnum{
	"timeCreated": ListMigrationsSortByTimecreated,
	"displayName": ListMigrationsSortByDisplayname,
}

var mappingListMigrationsSortByEnumLowerCase = map[string]ListMigrationsSortByEnum{
	"timecreated": ListMigrationsSortByTimecreated,
	"displayname": ListMigrationsSortByDisplayname,
}

// GetListMigrationsSortByEnumValues Enumerates the set of values for ListMigrationsSortByEnum
func GetListMigrationsSortByEnumValues() []ListMigrationsSortByEnum {
	values := make([]ListMigrationsSortByEnum, 0)
	for _, v := range mappingListMigrationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationsSortByEnumStringValues Enumerates the set of values in String for ListMigrationsSortByEnum
func GetListMigrationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMigrationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationsSortByEnum(val string) (ListMigrationsSortByEnum, bool) {
	enum, ok := mappingListMigrationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
