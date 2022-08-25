// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLiveMigrationsRequest wrapper for the ListLiveMigrations operation
type ListLiveMigrationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
	NetworkLoadBalancerId *string `mandatory:"true" contributesTo:"path" name:"networkLoadBalancerId"`

	// The unique Oracle-assigned identifier for the request. If you must contact Oracle about a
	// particular request, then provide the request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page or items to return, in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from which to start retrieving results.
	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' (ascending) or 'desc' (descending).
	SortOrder ListLiveMigrationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. The default order for timeCreated is descending.
	// The default order for displayName is ascending. If no value is specified, then timeCreated is the default.
	SortBy ListLiveMigrationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLiveMigrationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLiveMigrationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLiveMigrationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLiveMigrationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLiveMigrationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLiveMigrationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLiveMigrationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLiveMigrationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLiveMigrationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLiveMigrationsResponse wrapper for the ListLiveMigrations operation
type ListLiveMigrationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LiveMigrationCollection instances
	LiveMigrationCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you must contact
	// Oracle about a particular request, then provide the request identifier.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLiveMigrationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLiveMigrationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLiveMigrationsSortOrderEnum Enum with underlying type: string
type ListLiveMigrationsSortOrderEnum string

// Set of constants representing the allowable values for ListLiveMigrationsSortOrderEnum
const (
	ListLiveMigrationsSortOrderAsc  ListLiveMigrationsSortOrderEnum = "ASC"
	ListLiveMigrationsSortOrderDesc ListLiveMigrationsSortOrderEnum = "DESC"
)

var mappingListLiveMigrationsSortOrderEnum = map[string]ListLiveMigrationsSortOrderEnum{
	"ASC":  ListLiveMigrationsSortOrderAsc,
	"DESC": ListLiveMigrationsSortOrderDesc,
}

var mappingListLiveMigrationsSortOrderEnumLowerCase = map[string]ListLiveMigrationsSortOrderEnum{
	"asc":  ListLiveMigrationsSortOrderAsc,
	"desc": ListLiveMigrationsSortOrderDesc,
}

// GetListLiveMigrationsSortOrderEnumValues Enumerates the set of values for ListLiveMigrationsSortOrderEnum
func GetListLiveMigrationsSortOrderEnumValues() []ListLiveMigrationsSortOrderEnum {
	values := make([]ListLiveMigrationsSortOrderEnum, 0)
	for _, v := range mappingListLiveMigrationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLiveMigrationsSortOrderEnumStringValues Enumerates the set of values in String for ListLiveMigrationsSortOrderEnum
func GetListLiveMigrationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLiveMigrationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLiveMigrationsSortOrderEnum(val string) (ListLiveMigrationsSortOrderEnum, bool) {
	enum, ok := mappingListLiveMigrationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLiveMigrationsSortByEnum Enum with underlying type: string
type ListLiveMigrationsSortByEnum string

// Set of constants representing the allowable values for ListLiveMigrationsSortByEnum
const (
	ListLiveMigrationsSortByTimecreated ListLiveMigrationsSortByEnum = "timeCreated"
	ListLiveMigrationsSortByDisplayname ListLiveMigrationsSortByEnum = "displayName"
)

var mappingListLiveMigrationsSortByEnum = map[string]ListLiveMigrationsSortByEnum{
	"timeCreated": ListLiveMigrationsSortByTimecreated,
	"displayName": ListLiveMigrationsSortByDisplayname,
}

var mappingListLiveMigrationsSortByEnumLowerCase = map[string]ListLiveMigrationsSortByEnum{
	"timecreated": ListLiveMigrationsSortByTimecreated,
	"displayname": ListLiveMigrationsSortByDisplayname,
}

// GetListLiveMigrationsSortByEnumValues Enumerates the set of values for ListLiveMigrationsSortByEnum
func GetListLiveMigrationsSortByEnumValues() []ListLiveMigrationsSortByEnum {
	values := make([]ListLiveMigrationsSortByEnum, 0)
	for _, v := range mappingListLiveMigrationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLiveMigrationsSortByEnumStringValues Enumerates the set of values in String for ListLiveMigrationsSortByEnum
func GetListLiveMigrationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListLiveMigrationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLiveMigrationsSortByEnum(val string) (ListLiveMigrationsSortByEnum, bool) {
	enum, ok := mappingListLiveMigrationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
