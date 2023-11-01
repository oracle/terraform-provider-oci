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

// ListLifecycleEnvironmentsRequest wrapper for the ListLifecycleEnvironments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListLifecycleEnvironments.go.html to see an example of how to use ListLifecycleEnvironmentsRequest.
type ListLifecycleEnvironmentsRequest struct {

	// The OCID of the compartment that contains the resources to list.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The OCID of the lifecycle environment.
	LifecycleEnvironmentId *string `mandatory:"false" contributesTo:"query" name:"lifecycleEnvironmentId"`

	// A filter to return only profiles that match the given archType.
	ArchType ListLifecycleEnvironmentsArchTypeEnum `mandatory:"false" contributesTo:"query" name:"archType" omitEmpty:"true"`

	// A filter to return only profiles that match the given osFamily.
	OsFamily ListLifecycleEnvironmentsOsFamilyEnum `mandatory:"false" contributesTo:"query" name:"osFamily" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only the lifecycle environments that match the display name given.
	LifecycleState LifecycleEnvironmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListLifecycleEnvironmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListLifecycleEnvironmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLifecycleEnvironmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLifecycleEnvironmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLifecycleEnvironmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLifecycleEnvironmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLifecycleEnvironmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLifecycleEnvironmentsArchTypeEnum(string(request.ArchType)); !ok && request.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", request.ArchType, strings.Join(GetListLifecycleEnvironmentsArchTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLifecycleEnvironmentsOsFamilyEnum(string(request.OsFamily)); !ok && request.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", request.OsFamily, strings.Join(GetListLifecycleEnvironmentsOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleEnvironmentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetLifecycleEnvironmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLifecycleEnvironmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLifecycleEnvironmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLifecycleEnvironmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLifecycleEnvironmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLifecycleEnvironmentsResponse wrapper for the ListLifecycleEnvironments operation
type ListLifecycleEnvironmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LifecycleEnvironmentCollection instances
	LifecycleEnvironmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLifecycleEnvironmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLifecycleEnvironmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLifecycleEnvironmentsArchTypeEnum Enum with underlying type: string
type ListLifecycleEnvironmentsArchTypeEnum string

// Set of constants representing the allowable values for ListLifecycleEnvironmentsArchTypeEnum
const (
	ListLifecycleEnvironmentsArchTypeX8664   ListLifecycleEnvironmentsArchTypeEnum = "X86_64"
	ListLifecycleEnvironmentsArchTypeAarch64 ListLifecycleEnvironmentsArchTypeEnum = "AARCH64"
	ListLifecycleEnvironmentsArchTypeI686    ListLifecycleEnvironmentsArchTypeEnum = "I686"
	ListLifecycleEnvironmentsArchTypeNoarch  ListLifecycleEnvironmentsArchTypeEnum = "NOARCH"
	ListLifecycleEnvironmentsArchTypeSrc     ListLifecycleEnvironmentsArchTypeEnum = "SRC"
)

var mappingListLifecycleEnvironmentsArchTypeEnum = map[string]ListLifecycleEnvironmentsArchTypeEnum{
	"X86_64":  ListLifecycleEnvironmentsArchTypeX8664,
	"AARCH64": ListLifecycleEnvironmentsArchTypeAarch64,
	"I686":    ListLifecycleEnvironmentsArchTypeI686,
	"NOARCH":  ListLifecycleEnvironmentsArchTypeNoarch,
	"SRC":     ListLifecycleEnvironmentsArchTypeSrc,
}

var mappingListLifecycleEnvironmentsArchTypeEnumLowerCase = map[string]ListLifecycleEnvironmentsArchTypeEnum{
	"x86_64":  ListLifecycleEnvironmentsArchTypeX8664,
	"aarch64": ListLifecycleEnvironmentsArchTypeAarch64,
	"i686":    ListLifecycleEnvironmentsArchTypeI686,
	"noarch":  ListLifecycleEnvironmentsArchTypeNoarch,
	"src":     ListLifecycleEnvironmentsArchTypeSrc,
}

// GetListLifecycleEnvironmentsArchTypeEnumValues Enumerates the set of values for ListLifecycleEnvironmentsArchTypeEnum
func GetListLifecycleEnvironmentsArchTypeEnumValues() []ListLifecycleEnvironmentsArchTypeEnum {
	values := make([]ListLifecycleEnvironmentsArchTypeEnum, 0)
	for _, v := range mappingListLifecycleEnvironmentsArchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListLifecycleEnvironmentsArchTypeEnumStringValues Enumerates the set of values in String for ListLifecycleEnvironmentsArchTypeEnum
func GetListLifecycleEnvironmentsArchTypeEnumStringValues() []string {
	return []string{
		"X86_64",
		"AARCH64",
		"I686",
		"NOARCH",
		"SRC",
	}
}

// GetMappingListLifecycleEnvironmentsArchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLifecycleEnvironmentsArchTypeEnum(val string) (ListLifecycleEnvironmentsArchTypeEnum, bool) {
	enum, ok := mappingListLifecycleEnvironmentsArchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLifecycleEnvironmentsOsFamilyEnum Enum with underlying type: string
type ListLifecycleEnvironmentsOsFamilyEnum string

// Set of constants representing the allowable values for ListLifecycleEnvironmentsOsFamilyEnum
const (
	ListLifecycleEnvironmentsOsFamily9 ListLifecycleEnvironmentsOsFamilyEnum = "ORACLE_LINUX_9"
	ListLifecycleEnvironmentsOsFamily8 ListLifecycleEnvironmentsOsFamilyEnum = "ORACLE_LINUX_8"
	ListLifecycleEnvironmentsOsFamily7 ListLifecycleEnvironmentsOsFamilyEnum = "ORACLE_LINUX_7"
)

var mappingListLifecycleEnvironmentsOsFamilyEnum = map[string]ListLifecycleEnvironmentsOsFamilyEnum{
	"ORACLE_LINUX_9": ListLifecycleEnvironmentsOsFamily9,
	"ORACLE_LINUX_8": ListLifecycleEnvironmentsOsFamily8,
	"ORACLE_LINUX_7": ListLifecycleEnvironmentsOsFamily7,
}

var mappingListLifecycleEnvironmentsOsFamilyEnumLowerCase = map[string]ListLifecycleEnvironmentsOsFamilyEnum{
	"oracle_linux_9": ListLifecycleEnvironmentsOsFamily9,
	"oracle_linux_8": ListLifecycleEnvironmentsOsFamily8,
	"oracle_linux_7": ListLifecycleEnvironmentsOsFamily7,
}

// GetListLifecycleEnvironmentsOsFamilyEnumValues Enumerates the set of values for ListLifecycleEnvironmentsOsFamilyEnum
func GetListLifecycleEnvironmentsOsFamilyEnumValues() []ListLifecycleEnvironmentsOsFamilyEnum {
	values := make([]ListLifecycleEnvironmentsOsFamilyEnum, 0)
	for _, v := range mappingListLifecycleEnvironmentsOsFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetListLifecycleEnvironmentsOsFamilyEnumStringValues Enumerates the set of values in String for ListLifecycleEnvironmentsOsFamilyEnum
func GetListLifecycleEnvironmentsOsFamilyEnumStringValues() []string {
	return []string{
		"ORACLE_LINUX_9",
		"ORACLE_LINUX_8",
		"ORACLE_LINUX_7",
	}
}

// GetMappingListLifecycleEnvironmentsOsFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLifecycleEnvironmentsOsFamilyEnum(val string) (ListLifecycleEnvironmentsOsFamilyEnum, bool) {
	enum, ok := mappingListLifecycleEnvironmentsOsFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLifecycleEnvironmentsSortOrderEnum Enum with underlying type: string
type ListLifecycleEnvironmentsSortOrderEnum string

// Set of constants representing the allowable values for ListLifecycleEnvironmentsSortOrderEnum
const (
	ListLifecycleEnvironmentsSortOrderAsc  ListLifecycleEnvironmentsSortOrderEnum = "ASC"
	ListLifecycleEnvironmentsSortOrderDesc ListLifecycleEnvironmentsSortOrderEnum = "DESC"
)

var mappingListLifecycleEnvironmentsSortOrderEnum = map[string]ListLifecycleEnvironmentsSortOrderEnum{
	"ASC":  ListLifecycleEnvironmentsSortOrderAsc,
	"DESC": ListLifecycleEnvironmentsSortOrderDesc,
}

var mappingListLifecycleEnvironmentsSortOrderEnumLowerCase = map[string]ListLifecycleEnvironmentsSortOrderEnum{
	"asc":  ListLifecycleEnvironmentsSortOrderAsc,
	"desc": ListLifecycleEnvironmentsSortOrderDesc,
}

// GetListLifecycleEnvironmentsSortOrderEnumValues Enumerates the set of values for ListLifecycleEnvironmentsSortOrderEnum
func GetListLifecycleEnvironmentsSortOrderEnumValues() []ListLifecycleEnvironmentsSortOrderEnum {
	values := make([]ListLifecycleEnvironmentsSortOrderEnum, 0)
	for _, v := range mappingListLifecycleEnvironmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLifecycleEnvironmentsSortOrderEnumStringValues Enumerates the set of values in String for ListLifecycleEnvironmentsSortOrderEnum
func GetListLifecycleEnvironmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLifecycleEnvironmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLifecycleEnvironmentsSortOrderEnum(val string) (ListLifecycleEnvironmentsSortOrderEnum, bool) {
	enum, ok := mappingListLifecycleEnvironmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLifecycleEnvironmentsSortByEnum Enum with underlying type: string
type ListLifecycleEnvironmentsSortByEnum string

// Set of constants representing the allowable values for ListLifecycleEnvironmentsSortByEnum
const (
	ListLifecycleEnvironmentsSortByTimecreated ListLifecycleEnvironmentsSortByEnum = "timeCreated"
	ListLifecycleEnvironmentsSortByDisplayname ListLifecycleEnvironmentsSortByEnum = "displayName"
)

var mappingListLifecycleEnvironmentsSortByEnum = map[string]ListLifecycleEnvironmentsSortByEnum{
	"timeCreated": ListLifecycleEnvironmentsSortByTimecreated,
	"displayName": ListLifecycleEnvironmentsSortByDisplayname,
}

var mappingListLifecycleEnvironmentsSortByEnumLowerCase = map[string]ListLifecycleEnvironmentsSortByEnum{
	"timecreated": ListLifecycleEnvironmentsSortByTimecreated,
	"displayname": ListLifecycleEnvironmentsSortByDisplayname,
}

// GetListLifecycleEnvironmentsSortByEnumValues Enumerates the set of values for ListLifecycleEnvironmentsSortByEnum
func GetListLifecycleEnvironmentsSortByEnumValues() []ListLifecycleEnvironmentsSortByEnum {
	values := make([]ListLifecycleEnvironmentsSortByEnum, 0)
	for _, v := range mappingListLifecycleEnvironmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLifecycleEnvironmentsSortByEnumStringValues Enumerates the set of values in String for ListLifecycleEnvironmentsSortByEnum
func GetListLifecycleEnvironmentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListLifecycleEnvironmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLifecycleEnvironmentsSortByEnum(val string) (ListLifecycleEnvironmentsSortByEnum, bool) {
	enum, ok := mappingListLifecycleEnvironmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
