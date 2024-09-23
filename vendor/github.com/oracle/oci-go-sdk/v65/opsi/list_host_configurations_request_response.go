// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListHostConfigurationsRequest wrapper for the ListHostConfigurations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListHostConfigurations.go.html to see an example of how to use ListHostConfigurationsRequest.
type ListHostConfigurationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique Enterprise Manager bridge identifier
	EnterpriseManagerBridgeId *string `mandatory:"false" contributesTo:"query" name:"enterpriseManagerBridgeId"`

	// Optional list of host insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Filter by one or more platform types.
	// Supported platformType(s) for MACS-managed external host insight: [LINUX, SOLARIS, WINDOWS].
	// Supported platformType(s) for MACS-managed cloud host insight: [LINUX].
	// Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS, ZLINUX, WINDOWS, AIX, HP-UX].
	PlatformType []ListHostConfigurationsPlatformTypeEnum `contributesTo:"query" name:"platformType" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListHostConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Host configuration list sort options.
	SortBy ListHostConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A list of tag filters to apply.  Only resources with a defined tag matching the value will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagEquals []string `contributesTo:"query" name:"definedTagEquals" collectionFormat:"multi"`

	// A list of tag filters to apply.  Only resources with a freeform tag matching the value will be returned.
	// The key for each tag is "{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same tag name are interpreted as "OR".  Values for different tag names are interpreted as "AND".
	FreeformTagEquals []string `contributesTo:"query" name:"freeformTagEquals" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified defined tags exist will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.true" (for checking existence of a defined tag)
	// or "{namespace}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagExists []string `contributesTo:"query" name:"definedTagExists" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified freeform tags exist the value will be returned.
	// The key for each tag is "{tagName}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for different tag names are interpreted as "AND".
	FreeformTagExists []string `contributesTo:"query" name:"freeformTagExists" collectionFormat:"multi"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Filter by one or more host types.
	// Possible values are CLOUD-HOST, EXTERNAL-HOST, COMANAGED-VM-HOST, COMANAGED-BM-HOST, COMANAGED-EXACS-HOST, COMANAGED-EXACC-HOST
	HostType []string `contributesTo:"query" name:"hostType" collectionFormat:"multi"`

	// Optional OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host (Compute Id)
	HostId *string `mandatory:"false" contributesTo:"query" name:"hostId"`

	// Optional list of Exadata Insight VM cluster name.
	VmclusterName []string `contributesTo:"query" name:"vmclusterName" collectionFormat:"multi"`

	// Resource Status
	Status []ResourceStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListHostConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListHostConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListHostConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListHostConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListHostConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.PlatformType {
		if _, ok := GetMappingListHostConfigurationsPlatformTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", val, strings.Join(GetListHostConfigurationsPlatformTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListHostConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListHostConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHostConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListHostConfigurationsSortByEnumStringValues(), ",")))
	}
	for _, val := range request.Status {
		if _, ok := GetMappingResourceStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetResourceStatusEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListHostConfigurationsResponse wrapper for the ListHostConfigurations operation
type ListHostConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of HostConfigurationCollection instances
	HostConfigurationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. The total number of items in the result.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListHostConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListHostConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListHostConfigurationsPlatformTypeEnum Enum with underlying type: string
type ListHostConfigurationsPlatformTypeEnum string

// Set of constants representing the allowable values for ListHostConfigurationsPlatformTypeEnum
const (
	ListHostConfigurationsPlatformTypeLinux   ListHostConfigurationsPlatformTypeEnum = "LINUX"
	ListHostConfigurationsPlatformTypeSolaris ListHostConfigurationsPlatformTypeEnum = "SOLARIS"
	ListHostConfigurationsPlatformTypeSunos   ListHostConfigurationsPlatformTypeEnum = "SUNOS"
	ListHostConfigurationsPlatformTypeZlinux  ListHostConfigurationsPlatformTypeEnum = "ZLINUX"
	ListHostConfigurationsPlatformTypeWindows ListHostConfigurationsPlatformTypeEnum = "WINDOWS"
	ListHostConfigurationsPlatformTypeAix     ListHostConfigurationsPlatformTypeEnum = "AIX"
	ListHostConfigurationsPlatformTypeHpUx    ListHostConfigurationsPlatformTypeEnum = "HP_UX"
)

var mappingListHostConfigurationsPlatformTypeEnum = map[string]ListHostConfigurationsPlatformTypeEnum{
	"LINUX":   ListHostConfigurationsPlatformTypeLinux,
	"SOLARIS": ListHostConfigurationsPlatformTypeSolaris,
	"SUNOS":   ListHostConfigurationsPlatformTypeSunos,
	"ZLINUX":  ListHostConfigurationsPlatformTypeZlinux,
	"WINDOWS": ListHostConfigurationsPlatformTypeWindows,
	"AIX":     ListHostConfigurationsPlatformTypeAix,
	"HP_UX":   ListHostConfigurationsPlatformTypeHpUx,
}

var mappingListHostConfigurationsPlatformTypeEnumLowerCase = map[string]ListHostConfigurationsPlatformTypeEnum{
	"linux":   ListHostConfigurationsPlatformTypeLinux,
	"solaris": ListHostConfigurationsPlatformTypeSolaris,
	"sunos":   ListHostConfigurationsPlatformTypeSunos,
	"zlinux":  ListHostConfigurationsPlatformTypeZlinux,
	"windows": ListHostConfigurationsPlatformTypeWindows,
	"aix":     ListHostConfigurationsPlatformTypeAix,
	"hp_ux":   ListHostConfigurationsPlatformTypeHpUx,
}

// GetListHostConfigurationsPlatformTypeEnumValues Enumerates the set of values for ListHostConfigurationsPlatformTypeEnum
func GetListHostConfigurationsPlatformTypeEnumValues() []ListHostConfigurationsPlatformTypeEnum {
	values := make([]ListHostConfigurationsPlatformTypeEnum, 0)
	for _, v := range mappingListHostConfigurationsPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListHostConfigurationsPlatformTypeEnumStringValues Enumerates the set of values in String for ListHostConfigurationsPlatformTypeEnum
func GetListHostConfigurationsPlatformTypeEnumStringValues() []string {
	return []string{
		"LINUX",
		"SOLARIS",
		"SUNOS",
		"ZLINUX",
		"WINDOWS",
		"AIX",
		"HP_UX",
	}
}

// GetMappingListHostConfigurationsPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHostConfigurationsPlatformTypeEnum(val string) (ListHostConfigurationsPlatformTypeEnum, bool) {
	enum, ok := mappingListHostConfigurationsPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListHostConfigurationsSortOrderEnum Enum with underlying type: string
type ListHostConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListHostConfigurationsSortOrderEnum
const (
	ListHostConfigurationsSortOrderAsc  ListHostConfigurationsSortOrderEnum = "ASC"
	ListHostConfigurationsSortOrderDesc ListHostConfigurationsSortOrderEnum = "DESC"
)

var mappingListHostConfigurationsSortOrderEnum = map[string]ListHostConfigurationsSortOrderEnum{
	"ASC":  ListHostConfigurationsSortOrderAsc,
	"DESC": ListHostConfigurationsSortOrderDesc,
}

var mappingListHostConfigurationsSortOrderEnumLowerCase = map[string]ListHostConfigurationsSortOrderEnum{
	"asc":  ListHostConfigurationsSortOrderAsc,
	"desc": ListHostConfigurationsSortOrderDesc,
}

// GetListHostConfigurationsSortOrderEnumValues Enumerates the set of values for ListHostConfigurationsSortOrderEnum
func GetListHostConfigurationsSortOrderEnumValues() []ListHostConfigurationsSortOrderEnum {
	values := make([]ListHostConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListHostConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListHostConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListHostConfigurationsSortOrderEnum
func GetListHostConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListHostConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHostConfigurationsSortOrderEnum(val string) (ListHostConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListHostConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListHostConfigurationsSortByEnum Enum with underlying type: string
type ListHostConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListHostConfigurationsSortByEnum
const (
	ListHostConfigurationsSortByHostname     ListHostConfigurationsSortByEnum = "hostName"
	ListHostConfigurationsSortByPlatformtype ListHostConfigurationsSortByEnum = "platformType"
)

var mappingListHostConfigurationsSortByEnum = map[string]ListHostConfigurationsSortByEnum{
	"hostName":     ListHostConfigurationsSortByHostname,
	"platformType": ListHostConfigurationsSortByPlatformtype,
}

var mappingListHostConfigurationsSortByEnumLowerCase = map[string]ListHostConfigurationsSortByEnum{
	"hostname":     ListHostConfigurationsSortByHostname,
	"platformtype": ListHostConfigurationsSortByPlatformtype,
}

// GetListHostConfigurationsSortByEnumValues Enumerates the set of values for ListHostConfigurationsSortByEnum
func GetListHostConfigurationsSortByEnumValues() []ListHostConfigurationsSortByEnum {
	values := make([]ListHostConfigurationsSortByEnum, 0)
	for _, v := range mappingListHostConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListHostConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListHostConfigurationsSortByEnum
func GetListHostConfigurationsSortByEnumStringValues() []string {
	return []string{
		"hostName",
		"platformType",
	}
}

// GetMappingListHostConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHostConfigurationsSortByEnum(val string) (ListHostConfigurationsSortByEnum, bool) {
	enum, ok := mappingListHostConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
