// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLustreFileSystemsRequest wrapper for the ListLustreFileSystems operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/ListLustreFileSystems.go.html to see an example of how to use ListLustreFileSystemsRequest.
type ListLustreFileSystemsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState LustreFileSystemLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Lustre file system.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLustreFileSystemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListLustreFileSystemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLustreFileSystemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLustreFileSystemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLustreFileSystemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLustreFileSystemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLustreFileSystemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLustreFileSystemLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetLustreFileSystemLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLustreFileSystemsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLustreFileSystemsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLustreFileSystemsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLustreFileSystemsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLustreFileSystemsResponse wrapper for the ListLustreFileSystems operation
type ListLustreFileSystemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LustreFileSystemCollection instances
	LustreFileSystemCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLustreFileSystemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLustreFileSystemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLustreFileSystemsSortOrderEnum Enum with underlying type: string
type ListLustreFileSystemsSortOrderEnum string

// Set of constants representing the allowable values for ListLustreFileSystemsSortOrderEnum
const (
	ListLustreFileSystemsSortOrderAsc  ListLustreFileSystemsSortOrderEnum = "ASC"
	ListLustreFileSystemsSortOrderDesc ListLustreFileSystemsSortOrderEnum = "DESC"
)

var mappingListLustreFileSystemsSortOrderEnum = map[string]ListLustreFileSystemsSortOrderEnum{
	"ASC":  ListLustreFileSystemsSortOrderAsc,
	"DESC": ListLustreFileSystemsSortOrderDesc,
}

var mappingListLustreFileSystemsSortOrderEnumLowerCase = map[string]ListLustreFileSystemsSortOrderEnum{
	"asc":  ListLustreFileSystemsSortOrderAsc,
	"desc": ListLustreFileSystemsSortOrderDesc,
}

// GetListLustreFileSystemsSortOrderEnumValues Enumerates the set of values for ListLustreFileSystemsSortOrderEnum
func GetListLustreFileSystemsSortOrderEnumValues() []ListLustreFileSystemsSortOrderEnum {
	values := make([]ListLustreFileSystemsSortOrderEnum, 0)
	for _, v := range mappingListLustreFileSystemsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLustreFileSystemsSortOrderEnumStringValues Enumerates the set of values in String for ListLustreFileSystemsSortOrderEnum
func GetListLustreFileSystemsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLustreFileSystemsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLustreFileSystemsSortOrderEnum(val string) (ListLustreFileSystemsSortOrderEnum, bool) {
	enum, ok := mappingListLustreFileSystemsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLustreFileSystemsSortByEnum Enum with underlying type: string
type ListLustreFileSystemsSortByEnum string

// Set of constants representing the allowable values for ListLustreFileSystemsSortByEnum
const (
	ListLustreFileSystemsSortByTimecreated ListLustreFileSystemsSortByEnum = "timeCreated"
	ListLustreFileSystemsSortByDisplayname ListLustreFileSystemsSortByEnum = "displayName"
)

var mappingListLustreFileSystemsSortByEnum = map[string]ListLustreFileSystemsSortByEnum{
	"timeCreated": ListLustreFileSystemsSortByTimecreated,
	"displayName": ListLustreFileSystemsSortByDisplayname,
}

var mappingListLustreFileSystemsSortByEnumLowerCase = map[string]ListLustreFileSystemsSortByEnum{
	"timecreated": ListLustreFileSystemsSortByTimecreated,
	"displayname": ListLustreFileSystemsSortByDisplayname,
}

// GetListLustreFileSystemsSortByEnumValues Enumerates the set of values for ListLustreFileSystemsSortByEnum
func GetListLustreFileSystemsSortByEnumValues() []ListLustreFileSystemsSortByEnum {
	values := make([]ListLustreFileSystemsSortByEnum, 0)
	for _, v := range mappingListLustreFileSystemsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLustreFileSystemsSortByEnumStringValues Enumerates the set of values in String for ListLustreFileSystemsSortByEnum
func GetListLustreFileSystemsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListLustreFileSystemsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLustreFileSystemsSortByEnum(val string) (ListLustreFileSystemsSortByEnum, bool) {
	enum, ok := mappingListLustreFileSystemsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
