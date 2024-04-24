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

// ListAllSoftwarePackagesRequest wrapper for the ListAllSoftwarePackages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListAllSoftwarePackages.go.html to see an example of how to use ListAllSoftwarePackagesRequest.
type ListAllSoftwarePackagesRequest struct {

	// A filter to return resources that match the given user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return software packages that match the given version.
	Version *string `mandatory:"false" contributesTo:"query" name:"version"`

	// A filter to return software packages that match the given architecture.
	Architecture ListAllSoftwarePackagesArchitectureEnum `mandatory:"false" contributesTo:"query" name:"architecture" omitEmpty:"true"`

	// Indicates whether to list only the latest versions of packages, module streams, and stream profiles.
	IsLatest *bool `mandatory:"false" contributesTo:"query" name:"isLatest"`

	// A filter to return only resources that match the given operating system family.
	OsFamily ListAllSoftwarePackagesOsFamilyEnum `mandatory:"false" contributesTo:"query" name:"osFamily" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAllSoftwarePackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort packages by. Only one sort order may be provided. Default order for displayName is ascending. If no value is specified displayName is default.
	SortBy ListAllSoftwarePackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAllSoftwarePackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAllSoftwarePackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAllSoftwarePackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAllSoftwarePackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAllSoftwarePackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAllSoftwarePackagesArchitectureEnum(string(request.Architecture)); !ok && request.Architecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Architecture: %s. Supported values are: %s.", request.Architecture, strings.Join(GetListAllSoftwarePackagesArchitectureEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAllSoftwarePackagesOsFamilyEnum(string(request.OsFamily)); !ok && request.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", request.OsFamily, strings.Join(GetListAllSoftwarePackagesOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAllSoftwarePackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAllSoftwarePackagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAllSoftwarePackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAllSoftwarePackagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAllSoftwarePackagesResponse wrapper for the ListAllSoftwarePackages operation
type ListAllSoftwarePackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SoftwarePackageCollection instances
	SoftwarePackageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAllSoftwarePackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAllSoftwarePackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAllSoftwarePackagesArchitectureEnum Enum with underlying type: string
type ListAllSoftwarePackagesArchitectureEnum string

// Set of constants representing the allowable values for ListAllSoftwarePackagesArchitectureEnum
const (
	ListAllSoftwarePackagesArchitectureI386    ListAllSoftwarePackagesArchitectureEnum = "I386"
	ListAllSoftwarePackagesArchitectureI686    ListAllSoftwarePackagesArchitectureEnum = "I686"
	ListAllSoftwarePackagesArchitectureAarch64 ListAllSoftwarePackagesArchitectureEnum = "AARCH64"
	ListAllSoftwarePackagesArchitectureX8664   ListAllSoftwarePackagesArchitectureEnum = "X86_64"
	ListAllSoftwarePackagesArchitectureSrc     ListAllSoftwarePackagesArchitectureEnum = "SRC"
	ListAllSoftwarePackagesArchitectureNoarch  ListAllSoftwarePackagesArchitectureEnum = "NOARCH"
	ListAllSoftwarePackagesArchitectureOther   ListAllSoftwarePackagesArchitectureEnum = "OTHER"
)

var mappingListAllSoftwarePackagesArchitectureEnum = map[string]ListAllSoftwarePackagesArchitectureEnum{
	"I386":    ListAllSoftwarePackagesArchitectureI386,
	"I686":    ListAllSoftwarePackagesArchitectureI686,
	"AARCH64": ListAllSoftwarePackagesArchitectureAarch64,
	"X86_64":  ListAllSoftwarePackagesArchitectureX8664,
	"SRC":     ListAllSoftwarePackagesArchitectureSrc,
	"NOARCH":  ListAllSoftwarePackagesArchitectureNoarch,
	"OTHER":   ListAllSoftwarePackagesArchitectureOther,
}

var mappingListAllSoftwarePackagesArchitectureEnumLowerCase = map[string]ListAllSoftwarePackagesArchitectureEnum{
	"i386":    ListAllSoftwarePackagesArchitectureI386,
	"i686":    ListAllSoftwarePackagesArchitectureI686,
	"aarch64": ListAllSoftwarePackagesArchitectureAarch64,
	"x86_64":  ListAllSoftwarePackagesArchitectureX8664,
	"src":     ListAllSoftwarePackagesArchitectureSrc,
	"noarch":  ListAllSoftwarePackagesArchitectureNoarch,
	"other":   ListAllSoftwarePackagesArchitectureOther,
}

// GetListAllSoftwarePackagesArchitectureEnumValues Enumerates the set of values for ListAllSoftwarePackagesArchitectureEnum
func GetListAllSoftwarePackagesArchitectureEnumValues() []ListAllSoftwarePackagesArchitectureEnum {
	values := make([]ListAllSoftwarePackagesArchitectureEnum, 0)
	for _, v := range mappingListAllSoftwarePackagesArchitectureEnum {
		values = append(values, v)
	}
	return values
}

// GetListAllSoftwarePackagesArchitectureEnumStringValues Enumerates the set of values in String for ListAllSoftwarePackagesArchitectureEnum
func GetListAllSoftwarePackagesArchitectureEnumStringValues() []string {
	return []string{
		"I386",
		"I686",
		"AARCH64",
		"X86_64",
		"SRC",
		"NOARCH",
		"OTHER",
	}
}

// GetMappingListAllSoftwarePackagesArchitectureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAllSoftwarePackagesArchitectureEnum(val string) (ListAllSoftwarePackagesArchitectureEnum, bool) {
	enum, ok := mappingListAllSoftwarePackagesArchitectureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAllSoftwarePackagesOsFamilyEnum Enum with underlying type: string
type ListAllSoftwarePackagesOsFamilyEnum string

// Set of constants representing the allowable values for ListAllSoftwarePackagesOsFamilyEnum
const (
	ListAllSoftwarePackagesOsFamilyOracleLinux9      ListAllSoftwarePackagesOsFamilyEnum = "ORACLE_LINUX_9"
	ListAllSoftwarePackagesOsFamilyOracleLinux8      ListAllSoftwarePackagesOsFamilyEnum = "ORACLE_LINUX_8"
	ListAllSoftwarePackagesOsFamilyOracleLinux7      ListAllSoftwarePackagesOsFamilyEnum = "ORACLE_LINUX_7"
	ListAllSoftwarePackagesOsFamilyOracleLinux6      ListAllSoftwarePackagesOsFamilyEnum = "ORACLE_LINUX_6"
	ListAllSoftwarePackagesOsFamilyWindowsServer2016 ListAllSoftwarePackagesOsFamilyEnum = "WINDOWS_SERVER_2016"
	ListAllSoftwarePackagesOsFamilyWindowsServer2019 ListAllSoftwarePackagesOsFamilyEnum = "WINDOWS_SERVER_2019"
	ListAllSoftwarePackagesOsFamilyWindowsServer2022 ListAllSoftwarePackagesOsFamilyEnum = "WINDOWS_SERVER_2022"
	ListAllSoftwarePackagesOsFamilyAll               ListAllSoftwarePackagesOsFamilyEnum = "ALL"
)

var mappingListAllSoftwarePackagesOsFamilyEnum = map[string]ListAllSoftwarePackagesOsFamilyEnum{
	"ORACLE_LINUX_9":      ListAllSoftwarePackagesOsFamilyOracleLinux9,
	"ORACLE_LINUX_8":      ListAllSoftwarePackagesOsFamilyOracleLinux8,
	"ORACLE_LINUX_7":      ListAllSoftwarePackagesOsFamilyOracleLinux7,
	"ORACLE_LINUX_6":      ListAllSoftwarePackagesOsFamilyOracleLinux6,
	"WINDOWS_SERVER_2016": ListAllSoftwarePackagesOsFamilyWindowsServer2016,
	"WINDOWS_SERVER_2019": ListAllSoftwarePackagesOsFamilyWindowsServer2019,
	"WINDOWS_SERVER_2022": ListAllSoftwarePackagesOsFamilyWindowsServer2022,
	"ALL":                 ListAllSoftwarePackagesOsFamilyAll,
}

var mappingListAllSoftwarePackagesOsFamilyEnumLowerCase = map[string]ListAllSoftwarePackagesOsFamilyEnum{
	"oracle_linux_9":      ListAllSoftwarePackagesOsFamilyOracleLinux9,
	"oracle_linux_8":      ListAllSoftwarePackagesOsFamilyOracleLinux8,
	"oracle_linux_7":      ListAllSoftwarePackagesOsFamilyOracleLinux7,
	"oracle_linux_6":      ListAllSoftwarePackagesOsFamilyOracleLinux6,
	"windows_server_2016": ListAllSoftwarePackagesOsFamilyWindowsServer2016,
	"windows_server_2019": ListAllSoftwarePackagesOsFamilyWindowsServer2019,
	"windows_server_2022": ListAllSoftwarePackagesOsFamilyWindowsServer2022,
	"all":                 ListAllSoftwarePackagesOsFamilyAll,
}

// GetListAllSoftwarePackagesOsFamilyEnumValues Enumerates the set of values for ListAllSoftwarePackagesOsFamilyEnum
func GetListAllSoftwarePackagesOsFamilyEnumValues() []ListAllSoftwarePackagesOsFamilyEnum {
	values := make([]ListAllSoftwarePackagesOsFamilyEnum, 0)
	for _, v := range mappingListAllSoftwarePackagesOsFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetListAllSoftwarePackagesOsFamilyEnumStringValues Enumerates the set of values in String for ListAllSoftwarePackagesOsFamilyEnum
func GetListAllSoftwarePackagesOsFamilyEnumStringValues() []string {
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

// GetMappingListAllSoftwarePackagesOsFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAllSoftwarePackagesOsFamilyEnum(val string) (ListAllSoftwarePackagesOsFamilyEnum, bool) {
	enum, ok := mappingListAllSoftwarePackagesOsFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAllSoftwarePackagesSortOrderEnum Enum with underlying type: string
type ListAllSoftwarePackagesSortOrderEnum string

// Set of constants representing the allowable values for ListAllSoftwarePackagesSortOrderEnum
const (
	ListAllSoftwarePackagesSortOrderAsc  ListAllSoftwarePackagesSortOrderEnum = "ASC"
	ListAllSoftwarePackagesSortOrderDesc ListAllSoftwarePackagesSortOrderEnum = "DESC"
)

var mappingListAllSoftwarePackagesSortOrderEnum = map[string]ListAllSoftwarePackagesSortOrderEnum{
	"ASC":  ListAllSoftwarePackagesSortOrderAsc,
	"DESC": ListAllSoftwarePackagesSortOrderDesc,
}

var mappingListAllSoftwarePackagesSortOrderEnumLowerCase = map[string]ListAllSoftwarePackagesSortOrderEnum{
	"asc":  ListAllSoftwarePackagesSortOrderAsc,
	"desc": ListAllSoftwarePackagesSortOrderDesc,
}

// GetListAllSoftwarePackagesSortOrderEnumValues Enumerates the set of values for ListAllSoftwarePackagesSortOrderEnum
func GetListAllSoftwarePackagesSortOrderEnumValues() []ListAllSoftwarePackagesSortOrderEnum {
	values := make([]ListAllSoftwarePackagesSortOrderEnum, 0)
	for _, v := range mappingListAllSoftwarePackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAllSoftwarePackagesSortOrderEnumStringValues Enumerates the set of values in String for ListAllSoftwarePackagesSortOrderEnum
func GetListAllSoftwarePackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAllSoftwarePackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAllSoftwarePackagesSortOrderEnum(val string) (ListAllSoftwarePackagesSortOrderEnum, bool) {
	enum, ok := mappingListAllSoftwarePackagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAllSoftwarePackagesSortByEnum Enum with underlying type: string
type ListAllSoftwarePackagesSortByEnum string

// Set of constants representing the allowable values for ListAllSoftwarePackagesSortByEnum
const (
	ListAllSoftwarePackagesSortByDisplayname ListAllSoftwarePackagesSortByEnum = "displayName"
)

var mappingListAllSoftwarePackagesSortByEnum = map[string]ListAllSoftwarePackagesSortByEnum{
	"displayName": ListAllSoftwarePackagesSortByDisplayname,
}

var mappingListAllSoftwarePackagesSortByEnumLowerCase = map[string]ListAllSoftwarePackagesSortByEnum{
	"displayname": ListAllSoftwarePackagesSortByDisplayname,
}

// GetListAllSoftwarePackagesSortByEnumValues Enumerates the set of values for ListAllSoftwarePackagesSortByEnum
func GetListAllSoftwarePackagesSortByEnumValues() []ListAllSoftwarePackagesSortByEnum {
	values := make([]ListAllSoftwarePackagesSortByEnum, 0)
	for _, v := range mappingListAllSoftwarePackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAllSoftwarePackagesSortByEnumStringValues Enumerates the set of values in String for ListAllSoftwarePackagesSortByEnum
func GetListAllSoftwarePackagesSortByEnumStringValues() []string {
	return []string{
		"displayName",
	}
}

// GetMappingListAllSoftwarePackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAllSoftwarePackagesSortByEnum(val string) (ListAllSoftwarePackagesSortByEnum, bool) {
	enum, ok := mappingListAllSoftwarePackagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
