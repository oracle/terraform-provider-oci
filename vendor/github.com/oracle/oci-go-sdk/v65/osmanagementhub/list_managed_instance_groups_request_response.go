// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedInstanceGroupsRequest wrapper for the ListManagedInstanceGroups operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceGroups.go.html to see an example of how to use ListManagedInstanceGroupsRequest.
type ListManagedInstanceGroupsRequest struct {

	// The OCID of the compartment that contains the resources to list.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID of the managed instance group for which to list resources.
	ManagedInstanceGroupId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceGroupId"`

	// The OCID for the software source.
	SoftwareSourceId *string `mandatory:"false" contributesTo:"query" name:"softwareSourceId"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only profiles that match the given archType.
	ArchType ListManagedInstanceGroupsArchTypeEnum `mandatory:"false" contributesTo:"query" name:"archType" omitEmpty:"true"`

	// A filter to return only profiles that match the given osFamily.
	OsFamily ListManagedInstanceGroupsOsFamilyEnum `mandatory:"false" contributesTo:"query" name:"osFamily" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources their lifecycle state matches the given lifecycle state.
	LifecycleState ManagedInstanceGroupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListManagedInstanceGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListManagedInstanceGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceGroupsArchTypeEnum(string(request.ArchType)); !ok && request.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", request.ArchType, strings.Join(GetListManagedInstanceGroupsArchTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceGroupsOsFamilyEnum(string(request.OsFamily)); !ok && request.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", request.OsFamily, strings.Join(GetListManagedInstanceGroupsOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingManagedInstanceGroupLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetManagedInstanceGroupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceGroupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceGroupsResponse wrapper for the ListManagedInstanceGroups operation
type ListManagedInstanceGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceGroupCollection instances
	ManagedInstanceGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceGroupsArchTypeEnum Enum with underlying type: string
type ListManagedInstanceGroupsArchTypeEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupsArchTypeEnum
const (
	ListManagedInstanceGroupsArchTypeX8664   ListManagedInstanceGroupsArchTypeEnum = "X86_64"
	ListManagedInstanceGroupsArchTypeAarch64 ListManagedInstanceGroupsArchTypeEnum = "AARCH64"
	ListManagedInstanceGroupsArchTypeI686    ListManagedInstanceGroupsArchTypeEnum = "I686"
	ListManagedInstanceGroupsArchTypeNoarch  ListManagedInstanceGroupsArchTypeEnum = "NOARCH"
	ListManagedInstanceGroupsArchTypeSrc     ListManagedInstanceGroupsArchTypeEnum = "SRC"
)

var mappingListManagedInstanceGroupsArchTypeEnum = map[string]ListManagedInstanceGroupsArchTypeEnum{
	"X86_64":  ListManagedInstanceGroupsArchTypeX8664,
	"AARCH64": ListManagedInstanceGroupsArchTypeAarch64,
	"I686":    ListManagedInstanceGroupsArchTypeI686,
	"NOARCH":  ListManagedInstanceGroupsArchTypeNoarch,
	"SRC":     ListManagedInstanceGroupsArchTypeSrc,
}

var mappingListManagedInstanceGroupsArchTypeEnumLowerCase = map[string]ListManagedInstanceGroupsArchTypeEnum{
	"x86_64":  ListManagedInstanceGroupsArchTypeX8664,
	"aarch64": ListManagedInstanceGroupsArchTypeAarch64,
	"i686":    ListManagedInstanceGroupsArchTypeI686,
	"noarch":  ListManagedInstanceGroupsArchTypeNoarch,
	"src":     ListManagedInstanceGroupsArchTypeSrc,
}

// GetListManagedInstanceGroupsArchTypeEnumValues Enumerates the set of values for ListManagedInstanceGroupsArchTypeEnum
func GetListManagedInstanceGroupsArchTypeEnumValues() []ListManagedInstanceGroupsArchTypeEnum {
	values := make([]ListManagedInstanceGroupsArchTypeEnum, 0)
	for _, v := range mappingListManagedInstanceGroupsArchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupsArchTypeEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupsArchTypeEnum
func GetListManagedInstanceGroupsArchTypeEnumStringValues() []string {
	return []string{
		"X86_64",
		"AARCH64",
		"I686",
		"NOARCH",
		"SRC",
	}
}

// GetMappingListManagedInstanceGroupsArchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupsArchTypeEnum(val string) (ListManagedInstanceGroupsArchTypeEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupsArchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceGroupsOsFamilyEnum Enum with underlying type: string
type ListManagedInstanceGroupsOsFamilyEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupsOsFamilyEnum
const (
	ListManagedInstanceGroupsOsFamily9 ListManagedInstanceGroupsOsFamilyEnum = "ORACLE_LINUX_9"
	ListManagedInstanceGroupsOsFamily8 ListManagedInstanceGroupsOsFamilyEnum = "ORACLE_LINUX_8"
	ListManagedInstanceGroupsOsFamily7 ListManagedInstanceGroupsOsFamilyEnum = "ORACLE_LINUX_7"
)

var mappingListManagedInstanceGroupsOsFamilyEnum = map[string]ListManagedInstanceGroupsOsFamilyEnum{
	"ORACLE_LINUX_9": ListManagedInstanceGroupsOsFamily9,
	"ORACLE_LINUX_8": ListManagedInstanceGroupsOsFamily8,
	"ORACLE_LINUX_7": ListManagedInstanceGroupsOsFamily7,
}

var mappingListManagedInstanceGroupsOsFamilyEnumLowerCase = map[string]ListManagedInstanceGroupsOsFamilyEnum{
	"oracle_linux_9": ListManagedInstanceGroupsOsFamily9,
	"oracle_linux_8": ListManagedInstanceGroupsOsFamily8,
	"oracle_linux_7": ListManagedInstanceGroupsOsFamily7,
}

// GetListManagedInstanceGroupsOsFamilyEnumValues Enumerates the set of values for ListManagedInstanceGroupsOsFamilyEnum
func GetListManagedInstanceGroupsOsFamilyEnumValues() []ListManagedInstanceGroupsOsFamilyEnum {
	values := make([]ListManagedInstanceGroupsOsFamilyEnum, 0)
	for _, v := range mappingListManagedInstanceGroupsOsFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupsOsFamilyEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupsOsFamilyEnum
func GetListManagedInstanceGroupsOsFamilyEnumStringValues() []string {
	return []string{
		"ORACLE_LINUX_9",
		"ORACLE_LINUX_8",
		"ORACLE_LINUX_7",
	}
}

// GetMappingListManagedInstanceGroupsOsFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupsOsFamilyEnum(val string) (ListManagedInstanceGroupsOsFamilyEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupsOsFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceGroupsSortOrderEnum Enum with underlying type: string
type ListManagedInstanceGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupsSortOrderEnum
const (
	ListManagedInstanceGroupsSortOrderAsc  ListManagedInstanceGroupsSortOrderEnum = "ASC"
	ListManagedInstanceGroupsSortOrderDesc ListManagedInstanceGroupsSortOrderEnum = "DESC"
)

var mappingListManagedInstanceGroupsSortOrderEnum = map[string]ListManagedInstanceGroupsSortOrderEnum{
	"ASC":  ListManagedInstanceGroupsSortOrderAsc,
	"DESC": ListManagedInstanceGroupsSortOrderDesc,
}

var mappingListManagedInstanceGroupsSortOrderEnumLowerCase = map[string]ListManagedInstanceGroupsSortOrderEnum{
	"asc":  ListManagedInstanceGroupsSortOrderAsc,
	"desc": ListManagedInstanceGroupsSortOrderDesc,
}

// GetListManagedInstanceGroupsSortOrderEnumValues Enumerates the set of values for ListManagedInstanceGroupsSortOrderEnum
func GetListManagedInstanceGroupsSortOrderEnumValues() []ListManagedInstanceGroupsSortOrderEnum {
	values := make([]ListManagedInstanceGroupsSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupsSortOrderEnum
func GetListManagedInstanceGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupsSortOrderEnum(val string) (ListManagedInstanceGroupsSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceGroupsSortByEnum Enum with underlying type: string
type ListManagedInstanceGroupsSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupsSortByEnum
const (
	ListManagedInstanceGroupsSortByTimecreated ListManagedInstanceGroupsSortByEnum = "timeCreated"
	ListManagedInstanceGroupsSortByDisplayname ListManagedInstanceGroupsSortByEnum = "displayName"
)

var mappingListManagedInstanceGroupsSortByEnum = map[string]ListManagedInstanceGroupsSortByEnum{
	"timeCreated": ListManagedInstanceGroupsSortByTimecreated,
	"displayName": ListManagedInstanceGroupsSortByDisplayname,
}

var mappingListManagedInstanceGroupsSortByEnumLowerCase = map[string]ListManagedInstanceGroupsSortByEnum{
	"timecreated": ListManagedInstanceGroupsSortByTimecreated,
	"displayname": ListManagedInstanceGroupsSortByDisplayname,
}

// GetListManagedInstanceGroupsSortByEnumValues Enumerates the set of values for ListManagedInstanceGroupsSortByEnum
func GetListManagedInstanceGroupsSortByEnumValues() []ListManagedInstanceGroupsSortByEnum {
	values := make([]ListManagedInstanceGroupsSortByEnum, 0)
	for _, v := range mappingListManagedInstanceGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupsSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupsSortByEnum
func GetListManagedInstanceGroupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagedInstanceGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupsSortByEnum(val string) (ListManagedInstanceGroupsSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
