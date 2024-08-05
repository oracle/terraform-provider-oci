// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAgentInstallersRequest wrapper for the ListAgentInstallers operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListAgentInstallers.go.html to see an example of how to use ListAgentInstallersRequest.
type ListAgentInstallersRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The ID of the Fleet.
	FleetId *string `mandatory:"false" contributesTo:"query" name:"fleetId"`

	// The platform architecture for the agent installer.
	PlatformArchitecture ListAgentInstallersPlatformArchitectureEnum `mandatory:"false" contributesTo:"query" name:"platformArchitecture" omitEmpty:"true"`

	// The OS family for the agent installer.
	OsFamily ListAgentInstallersOsFamilyEnum `mandatory:"false" contributesTo:"query" name:"osFamily" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListAgentInstallersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the agent installer. Only one sort order can be provided.
	// Default order for _agentInstallerId_, _osFamily_, _platformArchitecture_ is **ascending**.
	// If no value is specified _agentInstallerId_ is default.
	SortBy ListAgentInstallersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAgentInstallersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAgentInstallersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAgentInstallersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAgentInstallersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAgentInstallersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAgentInstallersPlatformArchitectureEnum(string(request.PlatformArchitecture)); !ok && request.PlatformArchitecture != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformArchitecture: %s. Supported values are: %s.", request.PlatformArchitecture, strings.Join(GetListAgentInstallersPlatformArchitectureEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAgentInstallersOsFamilyEnum(string(request.OsFamily)); !ok && request.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", request.OsFamily, strings.Join(GetListAgentInstallersOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAgentInstallersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAgentInstallersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAgentInstallersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAgentInstallersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAgentInstallersResponse wrapper for the ListAgentInstallers operation
type ListAgentInstallersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AgentInstallerCollection instances
	AgentInstallerCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAgentInstallersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAgentInstallersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAgentInstallersPlatformArchitectureEnum Enum with underlying type: string
type ListAgentInstallersPlatformArchitectureEnum string

// Set of constants representing the allowable values for ListAgentInstallersPlatformArchitectureEnum
const (
	ListAgentInstallersPlatformArchitectureX8664   ListAgentInstallersPlatformArchitectureEnum = "X86_64"
	ListAgentInstallersPlatformArchitectureX86     ListAgentInstallersPlatformArchitectureEnum = "X86"
	ListAgentInstallersPlatformArchitectureAarch64 ListAgentInstallersPlatformArchitectureEnum = "AARCH64"
)

var mappingListAgentInstallersPlatformArchitectureEnum = map[string]ListAgentInstallersPlatformArchitectureEnum{
	"X86_64":  ListAgentInstallersPlatformArchitectureX8664,
	"X86":     ListAgentInstallersPlatformArchitectureX86,
	"AARCH64": ListAgentInstallersPlatformArchitectureAarch64,
}

var mappingListAgentInstallersPlatformArchitectureEnumLowerCase = map[string]ListAgentInstallersPlatformArchitectureEnum{
	"x86_64":  ListAgentInstallersPlatformArchitectureX8664,
	"x86":     ListAgentInstallersPlatformArchitectureX86,
	"aarch64": ListAgentInstallersPlatformArchitectureAarch64,
}

// GetListAgentInstallersPlatformArchitectureEnumValues Enumerates the set of values for ListAgentInstallersPlatformArchitectureEnum
func GetListAgentInstallersPlatformArchitectureEnumValues() []ListAgentInstallersPlatformArchitectureEnum {
	values := make([]ListAgentInstallersPlatformArchitectureEnum, 0)
	for _, v := range mappingListAgentInstallersPlatformArchitectureEnum {
		values = append(values, v)
	}
	return values
}

// GetListAgentInstallersPlatformArchitectureEnumStringValues Enumerates the set of values in String for ListAgentInstallersPlatformArchitectureEnum
func GetListAgentInstallersPlatformArchitectureEnumStringValues() []string {
	return []string{
		"X86_64",
		"X86",
		"AARCH64",
	}
}

// GetMappingListAgentInstallersPlatformArchitectureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAgentInstallersPlatformArchitectureEnum(val string) (ListAgentInstallersPlatformArchitectureEnum, bool) {
	enum, ok := mappingListAgentInstallersPlatformArchitectureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAgentInstallersOsFamilyEnum Enum with underlying type: string
type ListAgentInstallersOsFamilyEnum string

// Set of constants representing the allowable values for ListAgentInstallersOsFamilyEnum
const (
	ListAgentInstallersOsFamilyLinux   ListAgentInstallersOsFamilyEnum = "LINUX"
	ListAgentInstallersOsFamilyWindows ListAgentInstallersOsFamilyEnum = "WINDOWS"
	ListAgentInstallersOsFamilyMacos   ListAgentInstallersOsFamilyEnum = "MACOS"
	ListAgentInstallersOsFamilyUnknown ListAgentInstallersOsFamilyEnum = "UNKNOWN"
)

var mappingListAgentInstallersOsFamilyEnum = map[string]ListAgentInstallersOsFamilyEnum{
	"LINUX":   ListAgentInstallersOsFamilyLinux,
	"WINDOWS": ListAgentInstallersOsFamilyWindows,
	"MACOS":   ListAgentInstallersOsFamilyMacos,
	"UNKNOWN": ListAgentInstallersOsFamilyUnknown,
}

var mappingListAgentInstallersOsFamilyEnumLowerCase = map[string]ListAgentInstallersOsFamilyEnum{
	"linux":   ListAgentInstallersOsFamilyLinux,
	"windows": ListAgentInstallersOsFamilyWindows,
	"macos":   ListAgentInstallersOsFamilyMacos,
	"unknown": ListAgentInstallersOsFamilyUnknown,
}

// GetListAgentInstallersOsFamilyEnumValues Enumerates the set of values for ListAgentInstallersOsFamilyEnum
func GetListAgentInstallersOsFamilyEnumValues() []ListAgentInstallersOsFamilyEnum {
	values := make([]ListAgentInstallersOsFamilyEnum, 0)
	for _, v := range mappingListAgentInstallersOsFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetListAgentInstallersOsFamilyEnumStringValues Enumerates the set of values in String for ListAgentInstallersOsFamilyEnum
func GetListAgentInstallersOsFamilyEnumStringValues() []string {
	return []string{
		"LINUX",
		"WINDOWS",
		"MACOS",
		"UNKNOWN",
	}
}

// GetMappingListAgentInstallersOsFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAgentInstallersOsFamilyEnum(val string) (ListAgentInstallersOsFamilyEnum, bool) {
	enum, ok := mappingListAgentInstallersOsFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAgentInstallersSortOrderEnum Enum with underlying type: string
type ListAgentInstallersSortOrderEnum string

// Set of constants representing the allowable values for ListAgentInstallersSortOrderEnum
const (
	ListAgentInstallersSortOrderAsc  ListAgentInstallersSortOrderEnum = "ASC"
	ListAgentInstallersSortOrderDesc ListAgentInstallersSortOrderEnum = "DESC"
)

var mappingListAgentInstallersSortOrderEnum = map[string]ListAgentInstallersSortOrderEnum{
	"ASC":  ListAgentInstallersSortOrderAsc,
	"DESC": ListAgentInstallersSortOrderDesc,
}

var mappingListAgentInstallersSortOrderEnumLowerCase = map[string]ListAgentInstallersSortOrderEnum{
	"asc":  ListAgentInstallersSortOrderAsc,
	"desc": ListAgentInstallersSortOrderDesc,
}

// GetListAgentInstallersSortOrderEnumValues Enumerates the set of values for ListAgentInstallersSortOrderEnum
func GetListAgentInstallersSortOrderEnumValues() []ListAgentInstallersSortOrderEnum {
	values := make([]ListAgentInstallersSortOrderEnum, 0)
	for _, v := range mappingListAgentInstallersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAgentInstallersSortOrderEnumStringValues Enumerates the set of values in String for ListAgentInstallersSortOrderEnum
func GetListAgentInstallersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAgentInstallersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAgentInstallersSortOrderEnum(val string) (ListAgentInstallersSortOrderEnum, bool) {
	enum, ok := mappingListAgentInstallersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAgentInstallersSortByEnum Enum with underlying type: string
type ListAgentInstallersSortByEnum string

// Set of constants representing the allowable values for ListAgentInstallersSortByEnum
const (
	ListAgentInstallersSortByAgentinstallerid     ListAgentInstallersSortByEnum = "agentInstallerId"
	ListAgentInstallersSortByOsfamily             ListAgentInstallersSortByEnum = "osFamily"
	ListAgentInstallersSortByPlatformarchitecture ListAgentInstallersSortByEnum = "platformArchitecture"
)

var mappingListAgentInstallersSortByEnum = map[string]ListAgentInstallersSortByEnum{
	"agentInstallerId":     ListAgentInstallersSortByAgentinstallerid,
	"osFamily":             ListAgentInstallersSortByOsfamily,
	"platformArchitecture": ListAgentInstallersSortByPlatformarchitecture,
}

var mappingListAgentInstallersSortByEnumLowerCase = map[string]ListAgentInstallersSortByEnum{
	"agentinstallerid":     ListAgentInstallersSortByAgentinstallerid,
	"osfamily":             ListAgentInstallersSortByOsfamily,
	"platformarchitecture": ListAgentInstallersSortByPlatformarchitecture,
}

// GetListAgentInstallersSortByEnumValues Enumerates the set of values for ListAgentInstallersSortByEnum
func GetListAgentInstallersSortByEnumValues() []ListAgentInstallersSortByEnum {
	values := make([]ListAgentInstallersSortByEnum, 0)
	for _, v := range mappingListAgentInstallersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAgentInstallersSortByEnumStringValues Enumerates the set of values in String for ListAgentInstallersSortByEnum
func GetListAgentInstallersSortByEnumStringValues() []string {
	return []string{
		"agentInstallerId",
		"osFamily",
		"platformArchitecture",
	}
}

// GetMappingListAgentInstallersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAgentInstallersSortByEnum(val string) (ListAgentInstallersSortByEnum, bool) {
	enum, ok := mappingListAgentInstallersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
