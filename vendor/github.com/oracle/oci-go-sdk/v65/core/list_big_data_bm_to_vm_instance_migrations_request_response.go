// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBigDataBmToVmInstanceMigrationsRequest wrapper for the ListBigDataBmToVmInstanceMigrations operation
type ListBigDataBmToVmInstanceMigrationsRequest struct {

	// A filter to return only BigDataBmToVmInstanceMigration resources that match the given compartment OCID exactly.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID of the source instance.
	SourceInstanceId *string `mandatory:"false" contributesTo:"query" name:"sourceInstanceId"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListBigDataBmToVmInstanceMigrationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListBigDataBmToVmInstanceMigrationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to only return migration resources that match the given migration state or sourceInstanceIdQueryParam. The state
	// value is case-insensitive.
	LifecycleState BigDataBmToVmInstanceMigrationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBigDataBmToVmInstanceMigrationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBigDataBmToVmInstanceMigrationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBigDataBmToVmInstanceMigrationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBigDataBmToVmInstanceMigrationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBigDataBmToVmInstanceMigrationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListBigDataBmToVmInstanceMigrationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBigDataBmToVmInstanceMigrationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBigDataBmToVmInstanceMigrationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBigDataBmToVmInstanceMigrationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBigDataBmToVmInstanceMigrationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBigDataBmToVmInstanceMigrationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBigDataBmToVmInstanceMigrationsResponse wrapper for the ListBigDataBmToVmInstanceMigrations operation
type ListBigDataBmToVmInstanceMigrationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BigDataBmToVmInstanceMigrationCollection instances
	BigDataBmToVmInstanceMigrationCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListBigDataBmToVmInstanceMigrationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBigDataBmToVmInstanceMigrationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBigDataBmToVmInstanceMigrationsSortByEnum Enum with underlying type: string
type ListBigDataBmToVmInstanceMigrationsSortByEnum string

// Set of constants representing the allowable values for ListBigDataBmToVmInstanceMigrationsSortByEnum
const (
	ListBigDataBmToVmInstanceMigrationsSortByTimecreated ListBigDataBmToVmInstanceMigrationsSortByEnum = "TIMECREATED"
	ListBigDataBmToVmInstanceMigrationsSortByDisplayname ListBigDataBmToVmInstanceMigrationsSortByEnum = "DISPLAYNAME"
)

var mappingListBigDataBmToVmInstanceMigrationsSortByEnum = map[string]ListBigDataBmToVmInstanceMigrationsSortByEnum{
	"TIMECREATED": ListBigDataBmToVmInstanceMigrationsSortByTimecreated,
	"DISPLAYNAME": ListBigDataBmToVmInstanceMigrationsSortByDisplayname,
}

var mappingListBigDataBmToVmInstanceMigrationsSortByEnumLowerCase = map[string]ListBigDataBmToVmInstanceMigrationsSortByEnum{
	"timecreated": ListBigDataBmToVmInstanceMigrationsSortByTimecreated,
	"displayname": ListBigDataBmToVmInstanceMigrationsSortByDisplayname,
}

// GetListBigDataBmToVmInstanceMigrationsSortByEnumValues Enumerates the set of values for ListBigDataBmToVmInstanceMigrationsSortByEnum
func GetListBigDataBmToVmInstanceMigrationsSortByEnumValues() []ListBigDataBmToVmInstanceMigrationsSortByEnum {
	values := make([]ListBigDataBmToVmInstanceMigrationsSortByEnum, 0)
	for _, v := range mappingListBigDataBmToVmInstanceMigrationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBigDataBmToVmInstanceMigrationsSortByEnumStringValues Enumerates the set of values in String for ListBigDataBmToVmInstanceMigrationsSortByEnum
func GetListBigDataBmToVmInstanceMigrationsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListBigDataBmToVmInstanceMigrationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBigDataBmToVmInstanceMigrationsSortByEnum(val string) (ListBigDataBmToVmInstanceMigrationsSortByEnum, bool) {
	enum, ok := mappingListBigDataBmToVmInstanceMigrationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBigDataBmToVmInstanceMigrationsSortOrderEnum Enum with underlying type: string
type ListBigDataBmToVmInstanceMigrationsSortOrderEnum string

// Set of constants representing the allowable values for ListBigDataBmToVmInstanceMigrationsSortOrderEnum
const (
	ListBigDataBmToVmInstanceMigrationsSortOrderAsc  ListBigDataBmToVmInstanceMigrationsSortOrderEnum = "ASC"
	ListBigDataBmToVmInstanceMigrationsSortOrderDesc ListBigDataBmToVmInstanceMigrationsSortOrderEnum = "DESC"
)

var mappingListBigDataBmToVmInstanceMigrationsSortOrderEnum = map[string]ListBigDataBmToVmInstanceMigrationsSortOrderEnum{
	"ASC":  ListBigDataBmToVmInstanceMigrationsSortOrderAsc,
	"DESC": ListBigDataBmToVmInstanceMigrationsSortOrderDesc,
}

var mappingListBigDataBmToVmInstanceMigrationsSortOrderEnumLowerCase = map[string]ListBigDataBmToVmInstanceMigrationsSortOrderEnum{
	"asc":  ListBigDataBmToVmInstanceMigrationsSortOrderAsc,
	"desc": ListBigDataBmToVmInstanceMigrationsSortOrderDesc,
}

// GetListBigDataBmToVmInstanceMigrationsSortOrderEnumValues Enumerates the set of values for ListBigDataBmToVmInstanceMigrationsSortOrderEnum
func GetListBigDataBmToVmInstanceMigrationsSortOrderEnumValues() []ListBigDataBmToVmInstanceMigrationsSortOrderEnum {
	values := make([]ListBigDataBmToVmInstanceMigrationsSortOrderEnum, 0)
	for _, v := range mappingListBigDataBmToVmInstanceMigrationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBigDataBmToVmInstanceMigrationsSortOrderEnumStringValues Enumerates the set of values in String for ListBigDataBmToVmInstanceMigrationsSortOrderEnum
func GetListBigDataBmToVmInstanceMigrationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBigDataBmToVmInstanceMigrationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBigDataBmToVmInstanceMigrationsSortOrderEnum(val string) (ListBigDataBmToVmInstanceMigrationsSortOrderEnum, bool) {
	enum, ok := mappingListBigDataBmToVmInstanceMigrationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
