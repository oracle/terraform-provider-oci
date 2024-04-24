// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListProfilesRequest wrapper for the ListProfiles operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListProfiles.go.html to see an example of how to use ListProfilesRequest.
type ListProfilesRequest struct {

	// The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return registration profiles that match the given profile type.
	ProfileType []ProfileTypeEnum `contributesTo:"query" name:"profileType" omitEmpty:"true" collectionFormat:"multi"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the registration profile. A filter used to return the specified profile.
	ProfileId *string `mandatory:"false" contributesTo:"query" name:"profileId"`

	// A filter to return only resources that match the given operating system family.
	OsFamily ListProfilesOsFamilyEnum `mandatory:"false" contributesTo:"query" name:"osFamily" omitEmpty:"true"`

	// A filter to return only profiles that match the given archType.
	ArchType ListProfilesArchTypeEnum `mandatory:"false" contributesTo:"query" name:"archType" omitEmpty:"true"`

	// A filter to return profiles that match the given instance type.
	RegistrationType []ProfileRegistrationTypeEnum `contributesTo:"query" name:"registrationType" omitEmpty:"true" collectionFormat:"multi"`

	// A boolean variable that is used to list only the default profile resources.
	IsDefaultProfile *bool `mandatory:"false" contributesTo:"query" name:"isDefaultProfile"`

	// A filter to return only service-provided profiles.
	IsServiceProvidedProfile *bool `mandatory:"false" contributesTo:"query" name:"isServiceProvidedProfile"`

	// A filter to return only resources that match the given vendor name.
	VendorName ListProfilesVendorNameEnum `mandatory:"false" contributesTo:"query" name:"vendorName" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only registration profiles in the given state.
	LifecycleState ProfileLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListProfilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending.
	// Default order for displayName is ascending.
	SortBy ListProfilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProfilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProfilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProfilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProfilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProfilesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.ProfileType {
		if _, ok := GetMappingProfileTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProfileType: %s. Supported values are: %s.", val, strings.Join(GetProfileTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListProfilesOsFamilyEnum(string(request.OsFamily)); !ok && request.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", request.OsFamily, strings.Join(GetListProfilesOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProfilesArchTypeEnum(string(request.ArchType)); !ok && request.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", request.ArchType, strings.Join(GetListProfilesArchTypeEnumStringValues(), ",")))
	}
	for _, val := range request.RegistrationType {
		if _, ok := GetMappingProfileRegistrationTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RegistrationType: %s. Supported values are: %s.", val, strings.Join(GetProfileRegistrationTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListProfilesVendorNameEnum(string(request.VendorName)); !ok && request.VendorName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VendorName: %s. Supported values are: %s.", request.VendorName, strings.Join(GetListProfilesVendorNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingProfileLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetProfileLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProfilesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProfilesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProfilesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProfilesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProfilesResponse wrapper for the ListProfiles operation
type ListProfilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProfileCollection instances
	ProfileCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProfilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProfilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProfilesOsFamilyEnum Enum with underlying type: string
type ListProfilesOsFamilyEnum string

// Set of constants representing the allowable values for ListProfilesOsFamilyEnum
const (
	ListProfilesOsFamilyOracleLinux9      ListProfilesOsFamilyEnum = "ORACLE_LINUX_9"
	ListProfilesOsFamilyOracleLinux8      ListProfilesOsFamilyEnum = "ORACLE_LINUX_8"
	ListProfilesOsFamilyOracleLinux7      ListProfilesOsFamilyEnum = "ORACLE_LINUX_7"
	ListProfilesOsFamilyOracleLinux6      ListProfilesOsFamilyEnum = "ORACLE_LINUX_6"
	ListProfilesOsFamilyWindowsServer2016 ListProfilesOsFamilyEnum = "WINDOWS_SERVER_2016"
	ListProfilesOsFamilyWindowsServer2019 ListProfilesOsFamilyEnum = "WINDOWS_SERVER_2019"
	ListProfilesOsFamilyWindowsServer2022 ListProfilesOsFamilyEnum = "WINDOWS_SERVER_2022"
	ListProfilesOsFamilyAll               ListProfilesOsFamilyEnum = "ALL"
)

var mappingListProfilesOsFamilyEnum = map[string]ListProfilesOsFamilyEnum{
	"ORACLE_LINUX_9":      ListProfilesOsFamilyOracleLinux9,
	"ORACLE_LINUX_8":      ListProfilesOsFamilyOracleLinux8,
	"ORACLE_LINUX_7":      ListProfilesOsFamilyOracleLinux7,
	"ORACLE_LINUX_6":      ListProfilesOsFamilyOracleLinux6,
	"WINDOWS_SERVER_2016": ListProfilesOsFamilyWindowsServer2016,
	"WINDOWS_SERVER_2019": ListProfilesOsFamilyWindowsServer2019,
	"WINDOWS_SERVER_2022": ListProfilesOsFamilyWindowsServer2022,
	"ALL":                 ListProfilesOsFamilyAll,
}

var mappingListProfilesOsFamilyEnumLowerCase = map[string]ListProfilesOsFamilyEnum{
	"oracle_linux_9":      ListProfilesOsFamilyOracleLinux9,
	"oracle_linux_8":      ListProfilesOsFamilyOracleLinux8,
	"oracle_linux_7":      ListProfilesOsFamilyOracleLinux7,
	"oracle_linux_6":      ListProfilesOsFamilyOracleLinux6,
	"windows_server_2016": ListProfilesOsFamilyWindowsServer2016,
	"windows_server_2019": ListProfilesOsFamilyWindowsServer2019,
	"windows_server_2022": ListProfilesOsFamilyWindowsServer2022,
	"all":                 ListProfilesOsFamilyAll,
}

// GetListProfilesOsFamilyEnumValues Enumerates the set of values for ListProfilesOsFamilyEnum
func GetListProfilesOsFamilyEnumValues() []ListProfilesOsFamilyEnum {
	values := make([]ListProfilesOsFamilyEnum, 0)
	for _, v := range mappingListProfilesOsFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetListProfilesOsFamilyEnumStringValues Enumerates the set of values in String for ListProfilesOsFamilyEnum
func GetListProfilesOsFamilyEnumStringValues() []string {
	return []string{
		"ORACLE_LINUX_9",
		"ORACLE_LINUX_8",
		"ORACLE_LINUX_7",
		"ORACLE_LINUX_6",
		"WINDOWS_SERVER_2016",
		"WINDOWS_SERVER_2019",
		"WINDOWS_SERVER_2022",
		"ALL",
	}
}

// GetMappingListProfilesOsFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProfilesOsFamilyEnum(val string) (ListProfilesOsFamilyEnum, bool) {
	enum, ok := mappingListProfilesOsFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProfilesArchTypeEnum Enum with underlying type: string
type ListProfilesArchTypeEnum string

// Set of constants representing the allowable values for ListProfilesArchTypeEnum
const (
	ListProfilesArchTypeX8664   ListProfilesArchTypeEnum = "X86_64"
	ListProfilesArchTypeAarch64 ListProfilesArchTypeEnum = "AARCH64"
	ListProfilesArchTypeI686    ListProfilesArchTypeEnum = "I686"
	ListProfilesArchTypeNoarch  ListProfilesArchTypeEnum = "NOARCH"
	ListProfilesArchTypeSrc     ListProfilesArchTypeEnum = "SRC"
)

var mappingListProfilesArchTypeEnum = map[string]ListProfilesArchTypeEnum{
	"X86_64":  ListProfilesArchTypeX8664,
	"AARCH64": ListProfilesArchTypeAarch64,
	"I686":    ListProfilesArchTypeI686,
	"NOARCH":  ListProfilesArchTypeNoarch,
	"SRC":     ListProfilesArchTypeSrc,
}

var mappingListProfilesArchTypeEnumLowerCase = map[string]ListProfilesArchTypeEnum{
	"x86_64":  ListProfilesArchTypeX8664,
	"aarch64": ListProfilesArchTypeAarch64,
	"i686":    ListProfilesArchTypeI686,
	"noarch":  ListProfilesArchTypeNoarch,
	"src":     ListProfilesArchTypeSrc,
}

// GetListProfilesArchTypeEnumValues Enumerates the set of values for ListProfilesArchTypeEnum
func GetListProfilesArchTypeEnumValues() []ListProfilesArchTypeEnum {
	values := make([]ListProfilesArchTypeEnum, 0)
	for _, v := range mappingListProfilesArchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListProfilesArchTypeEnumStringValues Enumerates the set of values in String for ListProfilesArchTypeEnum
func GetListProfilesArchTypeEnumStringValues() []string {
	return []string{
		"X86_64",
		"AARCH64",
		"I686",
		"NOARCH",
		"SRC",
	}
}

// GetMappingListProfilesArchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProfilesArchTypeEnum(val string) (ListProfilesArchTypeEnum, bool) {
	enum, ok := mappingListProfilesArchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProfilesVendorNameEnum Enum with underlying type: string
type ListProfilesVendorNameEnum string

// Set of constants representing the allowable values for ListProfilesVendorNameEnum
const (
	ListProfilesVendorNameOracle    ListProfilesVendorNameEnum = "ORACLE"
	ListProfilesVendorNameMicrosoft ListProfilesVendorNameEnum = "MICROSOFT"
)

var mappingListProfilesVendorNameEnum = map[string]ListProfilesVendorNameEnum{
	"ORACLE":    ListProfilesVendorNameOracle,
	"MICROSOFT": ListProfilesVendorNameMicrosoft,
}

var mappingListProfilesVendorNameEnumLowerCase = map[string]ListProfilesVendorNameEnum{
	"oracle":    ListProfilesVendorNameOracle,
	"microsoft": ListProfilesVendorNameMicrosoft,
}

// GetListProfilesVendorNameEnumValues Enumerates the set of values for ListProfilesVendorNameEnum
func GetListProfilesVendorNameEnumValues() []ListProfilesVendorNameEnum {
	values := make([]ListProfilesVendorNameEnum, 0)
	for _, v := range mappingListProfilesVendorNameEnum {
		values = append(values, v)
	}
	return values
}

// GetListProfilesVendorNameEnumStringValues Enumerates the set of values in String for ListProfilesVendorNameEnum
func GetListProfilesVendorNameEnumStringValues() []string {
	return []string{
		"ORACLE",
		"MICROSOFT",
	}
}

// GetMappingListProfilesVendorNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProfilesVendorNameEnum(val string) (ListProfilesVendorNameEnum, bool) {
	enum, ok := mappingListProfilesVendorNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProfilesSortOrderEnum Enum with underlying type: string
type ListProfilesSortOrderEnum string

// Set of constants representing the allowable values for ListProfilesSortOrderEnum
const (
	ListProfilesSortOrderAsc  ListProfilesSortOrderEnum = "ASC"
	ListProfilesSortOrderDesc ListProfilesSortOrderEnum = "DESC"
)

var mappingListProfilesSortOrderEnum = map[string]ListProfilesSortOrderEnum{
	"ASC":  ListProfilesSortOrderAsc,
	"DESC": ListProfilesSortOrderDesc,
}

var mappingListProfilesSortOrderEnumLowerCase = map[string]ListProfilesSortOrderEnum{
	"asc":  ListProfilesSortOrderAsc,
	"desc": ListProfilesSortOrderDesc,
}

// GetListProfilesSortOrderEnumValues Enumerates the set of values for ListProfilesSortOrderEnum
func GetListProfilesSortOrderEnumValues() []ListProfilesSortOrderEnum {
	values := make([]ListProfilesSortOrderEnum, 0)
	for _, v := range mappingListProfilesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProfilesSortOrderEnumStringValues Enumerates the set of values in String for ListProfilesSortOrderEnum
func GetListProfilesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProfilesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProfilesSortOrderEnum(val string) (ListProfilesSortOrderEnum, bool) {
	enum, ok := mappingListProfilesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProfilesSortByEnum Enum with underlying type: string
type ListProfilesSortByEnum string

// Set of constants representing the allowable values for ListProfilesSortByEnum
const (
	ListProfilesSortByTimecreated ListProfilesSortByEnum = "timeCreated"
	ListProfilesSortByDisplayname ListProfilesSortByEnum = "displayName"
)

var mappingListProfilesSortByEnum = map[string]ListProfilesSortByEnum{
	"timeCreated": ListProfilesSortByTimecreated,
	"displayName": ListProfilesSortByDisplayname,
}

var mappingListProfilesSortByEnumLowerCase = map[string]ListProfilesSortByEnum{
	"timecreated": ListProfilesSortByTimecreated,
	"displayname": ListProfilesSortByDisplayname,
}

// GetListProfilesSortByEnumValues Enumerates the set of values for ListProfilesSortByEnum
func GetListProfilesSortByEnumValues() []ListProfilesSortByEnum {
	values := make([]ListProfilesSortByEnum, 0)
	for _, v := range mappingListProfilesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProfilesSortByEnumStringValues Enumerates the set of values in String for ListProfilesSortByEnum
func GetListProfilesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListProfilesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProfilesSortByEnum(val string) (ListProfilesSortByEnum, bool) {
	enum, ok := mappingListProfilesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
