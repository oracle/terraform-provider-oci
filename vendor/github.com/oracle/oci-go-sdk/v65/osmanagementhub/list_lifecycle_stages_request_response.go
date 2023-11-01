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

// ListLifecycleStagesRequest wrapper for the ListLifecycleStages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListLifecycleStages.go.html to see an example of how to use ListLifecycleStagesRequest.
type ListLifecycleStagesRequest struct {

	// The OCID of the compartment that contains the resources to list.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The OCID of the lifecycle stage.
	LifecycleStageId *string `mandatory:"false" contributesTo:"query" name:"lifecycleStageId"`

	// The OCID for the software source.
	SoftwareSourceId *string `mandatory:"false" contributesTo:"query" name:"softwareSourceId"`

	// A filter to return only profiles that match the given archType.
	ArchType ListLifecycleStagesArchTypeEnum `mandatory:"false" contributesTo:"query" name:"archType" omitEmpty:"true"`

	// A filter to return only profiles that match the given osFamily.
	OsFamily ListLifecycleStagesOsFamilyEnum `mandatory:"false" contributesTo:"query" name:"osFamily" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only lifecycle stage whose lifecycle state matches the given lifecycle state.
	LifecycleState LifecycleStageLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListLifecycleStagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListLifecycleStagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLifecycleStagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLifecycleStagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLifecycleStagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLifecycleStagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLifecycleStagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLifecycleStagesArchTypeEnum(string(request.ArchType)); !ok && request.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", request.ArchType, strings.Join(GetListLifecycleStagesArchTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLifecycleStagesOsFamilyEnum(string(request.OsFamily)); !ok && request.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", request.OsFamily, strings.Join(GetListLifecycleStagesOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStageLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetLifecycleStageLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLifecycleStagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLifecycleStagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLifecycleStagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLifecycleStagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLifecycleStagesResponse wrapper for the ListLifecycleStages operation
type ListLifecycleStagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LifecycleStageCollection instances
	LifecycleStageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLifecycleStagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLifecycleStagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLifecycleStagesArchTypeEnum Enum with underlying type: string
type ListLifecycleStagesArchTypeEnum string

// Set of constants representing the allowable values for ListLifecycleStagesArchTypeEnum
const (
	ListLifecycleStagesArchTypeX8664   ListLifecycleStagesArchTypeEnum = "X86_64"
	ListLifecycleStagesArchTypeAarch64 ListLifecycleStagesArchTypeEnum = "AARCH64"
	ListLifecycleStagesArchTypeI686    ListLifecycleStagesArchTypeEnum = "I686"
	ListLifecycleStagesArchTypeNoarch  ListLifecycleStagesArchTypeEnum = "NOARCH"
	ListLifecycleStagesArchTypeSrc     ListLifecycleStagesArchTypeEnum = "SRC"
)

var mappingListLifecycleStagesArchTypeEnum = map[string]ListLifecycleStagesArchTypeEnum{
	"X86_64":  ListLifecycleStagesArchTypeX8664,
	"AARCH64": ListLifecycleStagesArchTypeAarch64,
	"I686":    ListLifecycleStagesArchTypeI686,
	"NOARCH":  ListLifecycleStagesArchTypeNoarch,
	"SRC":     ListLifecycleStagesArchTypeSrc,
}

var mappingListLifecycleStagesArchTypeEnumLowerCase = map[string]ListLifecycleStagesArchTypeEnum{
	"x86_64":  ListLifecycleStagesArchTypeX8664,
	"aarch64": ListLifecycleStagesArchTypeAarch64,
	"i686":    ListLifecycleStagesArchTypeI686,
	"noarch":  ListLifecycleStagesArchTypeNoarch,
	"src":     ListLifecycleStagesArchTypeSrc,
}

// GetListLifecycleStagesArchTypeEnumValues Enumerates the set of values for ListLifecycleStagesArchTypeEnum
func GetListLifecycleStagesArchTypeEnumValues() []ListLifecycleStagesArchTypeEnum {
	values := make([]ListLifecycleStagesArchTypeEnum, 0)
	for _, v := range mappingListLifecycleStagesArchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListLifecycleStagesArchTypeEnumStringValues Enumerates the set of values in String for ListLifecycleStagesArchTypeEnum
func GetListLifecycleStagesArchTypeEnumStringValues() []string {
	return []string{
		"X86_64",
		"AARCH64",
		"I686",
		"NOARCH",
		"SRC",
	}
}

// GetMappingListLifecycleStagesArchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLifecycleStagesArchTypeEnum(val string) (ListLifecycleStagesArchTypeEnum, bool) {
	enum, ok := mappingListLifecycleStagesArchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLifecycleStagesOsFamilyEnum Enum with underlying type: string
type ListLifecycleStagesOsFamilyEnum string

// Set of constants representing the allowable values for ListLifecycleStagesOsFamilyEnum
const (
	ListLifecycleStagesOsFamily9 ListLifecycleStagesOsFamilyEnum = "ORACLE_LINUX_9"
	ListLifecycleStagesOsFamily8 ListLifecycleStagesOsFamilyEnum = "ORACLE_LINUX_8"
	ListLifecycleStagesOsFamily7 ListLifecycleStagesOsFamilyEnum = "ORACLE_LINUX_7"
)

var mappingListLifecycleStagesOsFamilyEnum = map[string]ListLifecycleStagesOsFamilyEnum{
	"ORACLE_LINUX_9": ListLifecycleStagesOsFamily9,
	"ORACLE_LINUX_8": ListLifecycleStagesOsFamily8,
	"ORACLE_LINUX_7": ListLifecycleStagesOsFamily7,
}

var mappingListLifecycleStagesOsFamilyEnumLowerCase = map[string]ListLifecycleStagesOsFamilyEnum{
	"oracle_linux_9": ListLifecycleStagesOsFamily9,
	"oracle_linux_8": ListLifecycleStagesOsFamily8,
	"oracle_linux_7": ListLifecycleStagesOsFamily7,
}

// GetListLifecycleStagesOsFamilyEnumValues Enumerates the set of values for ListLifecycleStagesOsFamilyEnum
func GetListLifecycleStagesOsFamilyEnumValues() []ListLifecycleStagesOsFamilyEnum {
	values := make([]ListLifecycleStagesOsFamilyEnum, 0)
	for _, v := range mappingListLifecycleStagesOsFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetListLifecycleStagesOsFamilyEnumStringValues Enumerates the set of values in String for ListLifecycleStagesOsFamilyEnum
func GetListLifecycleStagesOsFamilyEnumStringValues() []string {
	return []string{
		"ORACLE_LINUX_9",
		"ORACLE_LINUX_8",
		"ORACLE_LINUX_7",
	}
}

// GetMappingListLifecycleStagesOsFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLifecycleStagesOsFamilyEnum(val string) (ListLifecycleStagesOsFamilyEnum, bool) {
	enum, ok := mappingListLifecycleStagesOsFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLifecycleStagesSortOrderEnum Enum with underlying type: string
type ListLifecycleStagesSortOrderEnum string

// Set of constants representing the allowable values for ListLifecycleStagesSortOrderEnum
const (
	ListLifecycleStagesSortOrderAsc  ListLifecycleStagesSortOrderEnum = "ASC"
	ListLifecycleStagesSortOrderDesc ListLifecycleStagesSortOrderEnum = "DESC"
)

var mappingListLifecycleStagesSortOrderEnum = map[string]ListLifecycleStagesSortOrderEnum{
	"ASC":  ListLifecycleStagesSortOrderAsc,
	"DESC": ListLifecycleStagesSortOrderDesc,
}

var mappingListLifecycleStagesSortOrderEnumLowerCase = map[string]ListLifecycleStagesSortOrderEnum{
	"asc":  ListLifecycleStagesSortOrderAsc,
	"desc": ListLifecycleStagesSortOrderDesc,
}

// GetListLifecycleStagesSortOrderEnumValues Enumerates the set of values for ListLifecycleStagesSortOrderEnum
func GetListLifecycleStagesSortOrderEnumValues() []ListLifecycleStagesSortOrderEnum {
	values := make([]ListLifecycleStagesSortOrderEnum, 0)
	for _, v := range mappingListLifecycleStagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLifecycleStagesSortOrderEnumStringValues Enumerates the set of values in String for ListLifecycleStagesSortOrderEnum
func GetListLifecycleStagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLifecycleStagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLifecycleStagesSortOrderEnum(val string) (ListLifecycleStagesSortOrderEnum, bool) {
	enum, ok := mappingListLifecycleStagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLifecycleStagesSortByEnum Enum with underlying type: string
type ListLifecycleStagesSortByEnum string

// Set of constants representing the allowable values for ListLifecycleStagesSortByEnum
const (
	ListLifecycleStagesSortByTimecreated ListLifecycleStagesSortByEnum = "timeCreated"
	ListLifecycleStagesSortByDisplayname ListLifecycleStagesSortByEnum = "displayName"
)

var mappingListLifecycleStagesSortByEnum = map[string]ListLifecycleStagesSortByEnum{
	"timeCreated": ListLifecycleStagesSortByTimecreated,
	"displayName": ListLifecycleStagesSortByDisplayname,
}

var mappingListLifecycleStagesSortByEnumLowerCase = map[string]ListLifecycleStagesSortByEnum{
	"timecreated": ListLifecycleStagesSortByTimecreated,
	"displayname": ListLifecycleStagesSortByDisplayname,
}

// GetListLifecycleStagesSortByEnumValues Enumerates the set of values for ListLifecycleStagesSortByEnum
func GetListLifecycleStagesSortByEnumValues() []ListLifecycleStagesSortByEnum {
	values := make([]ListLifecycleStagesSortByEnum, 0)
	for _, v := range mappingListLifecycleStagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLifecycleStagesSortByEnumStringValues Enumerates the set of values in String for ListLifecycleStagesSortByEnum
func GetListLifecycleStagesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListLifecycleStagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLifecycleStagesSortByEnum(val string) (ListLifecycleStagesSortByEnum, bool) {
	enum, ok := mappingListLifecycleStagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
