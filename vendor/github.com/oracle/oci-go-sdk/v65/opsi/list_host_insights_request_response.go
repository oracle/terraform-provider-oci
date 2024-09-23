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

// ListHostInsightsRequest wrapper for the ListHostInsights operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListHostInsights.go.html to see an example of how to use ListHostInsightsRequest.
type ListHostInsightsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Optional list of host insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Resource Status
	Status []ResourceStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// Lifecycle states
	LifecycleState []LifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// Filter by one or more host types.
	// Possible values are CLOUD-HOST, EXTERNAL-HOST, COMANAGED-VM-HOST, COMANAGED-BM-HOST, COMANAGED-EXACS-HOST, COMANAGED-EXACC-HOST
	HostType []string `contributesTo:"query" name:"hostType" collectionFormat:"multi"`

	// Filter by one or more platform types.
	// Supported platformType(s) for MACS-managed external host insight: [LINUX, SOLARIS, WINDOWS].
	// Supported platformType(s) for MACS-managed cloud host insight: [LINUX].
	// Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS, ZLINUX, WINDOWS, AIX, HP-UX].
	PlatformType []ListHostInsightsPlatformTypeEnum `contributesTo:"query" name:"platformType" omitEmpty:"true" collectionFormat:"multi"`

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
	SortOrder ListHostInsightsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Host insight list sort options. If `fields` parameter is selected, the `sortBy` parameter must be one of the fields specified.
	SortBy ListHostInsightsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Enterprise Manager bridge identifier
	EnterpriseManagerBridgeId *string `mandatory:"false" contributesTo:"query" name:"enterpriseManagerBridgeId"`

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of exadata insight resource.
	ExadataInsightId *string `mandatory:"false" contributesTo:"query" name:"exadataInsightId"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListHostInsightsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListHostInsightsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListHostInsightsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListHostInsightsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListHostInsightsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Status {
		if _, ok := GetMappingResourceStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetResourceStatusEnumStringValues(), ",")))
		}
	}

	for _, val := range request.LifecycleState {
		if _, ok := GetMappingLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
		}
	}

	for _, val := range request.PlatformType {
		if _, ok := GetMappingListHostInsightsPlatformTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", val, strings.Join(GetListHostInsightsPlatformTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListHostInsightsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListHostInsightsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHostInsightsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListHostInsightsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListHostInsightsResponse wrapper for the ListHostInsights operation
type ListHostInsightsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of HostInsightSummaryCollection instances
	HostInsightSummaryCollection `presentIn:"body"`

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

func (response ListHostInsightsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListHostInsightsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListHostInsightsPlatformTypeEnum Enum with underlying type: string
type ListHostInsightsPlatformTypeEnum string

// Set of constants representing the allowable values for ListHostInsightsPlatformTypeEnum
const (
	ListHostInsightsPlatformTypeLinux   ListHostInsightsPlatformTypeEnum = "LINUX"
	ListHostInsightsPlatformTypeSolaris ListHostInsightsPlatformTypeEnum = "SOLARIS"
	ListHostInsightsPlatformTypeSunos   ListHostInsightsPlatformTypeEnum = "SUNOS"
	ListHostInsightsPlatformTypeZlinux  ListHostInsightsPlatformTypeEnum = "ZLINUX"
	ListHostInsightsPlatformTypeWindows ListHostInsightsPlatformTypeEnum = "WINDOWS"
	ListHostInsightsPlatformTypeAix     ListHostInsightsPlatformTypeEnum = "AIX"
	ListHostInsightsPlatformTypeHpUx    ListHostInsightsPlatformTypeEnum = "HP_UX"
)

var mappingListHostInsightsPlatformTypeEnum = map[string]ListHostInsightsPlatformTypeEnum{
	"LINUX":   ListHostInsightsPlatformTypeLinux,
	"SOLARIS": ListHostInsightsPlatformTypeSolaris,
	"SUNOS":   ListHostInsightsPlatformTypeSunos,
	"ZLINUX":  ListHostInsightsPlatformTypeZlinux,
	"WINDOWS": ListHostInsightsPlatformTypeWindows,
	"AIX":     ListHostInsightsPlatformTypeAix,
	"HP_UX":   ListHostInsightsPlatformTypeHpUx,
}

var mappingListHostInsightsPlatformTypeEnumLowerCase = map[string]ListHostInsightsPlatformTypeEnum{
	"linux":   ListHostInsightsPlatformTypeLinux,
	"solaris": ListHostInsightsPlatformTypeSolaris,
	"sunos":   ListHostInsightsPlatformTypeSunos,
	"zlinux":  ListHostInsightsPlatformTypeZlinux,
	"windows": ListHostInsightsPlatformTypeWindows,
	"aix":     ListHostInsightsPlatformTypeAix,
	"hp_ux":   ListHostInsightsPlatformTypeHpUx,
}

// GetListHostInsightsPlatformTypeEnumValues Enumerates the set of values for ListHostInsightsPlatformTypeEnum
func GetListHostInsightsPlatformTypeEnumValues() []ListHostInsightsPlatformTypeEnum {
	values := make([]ListHostInsightsPlatformTypeEnum, 0)
	for _, v := range mappingListHostInsightsPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListHostInsightsPlatformTypeEnumStringValues Enumerates the set of values in String for ListHostInsightsPlatformTypeEnum
func GetListHostInsightsPlatformTypeEnumStringValues() []string {
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

// GetMappingListHostInsightsPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHostInsightsPlatformTypeEnum(val string) (ListHostInsightsPlatformTypeEnum, bool) {
	enum, ok := mappingListHostInsightsPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListHostInsightsSortOrderEnum Enum with underlying type: string
type ListHostInsightsSortOrderEnum string

// Set of constants representing the allowable values for ListHostInsightsSortOrderEnum
const (
	ListHostInsightsSortOrderAsc  ListHostInsightsSortOrderEnum = "ASC"
	ListHostInsightsSortOrderDesc ListHostInsightsSortOrderEnum = "DESC"
)

var mappingListHostInsightsSortOrderEnum = map[string]ListHostInsightsSortOrderEnum{
	"ASC":  ListHostInsightsSortOrderAsc,
	"DESC": ListHostInsightsSortOrderDesc,
}

var mappingListHostInsightsSortOrderEnumLowerCase = map[string]ListHostInsightsSortOrderEnum{
	"asc":  ListHostInsightsSortOrderAsc,
	"desc": ListHostInsightsSortOrderDesc,
}

// GetListHostInsightsSortOrderEnumValues Enumerates the set of values for ListHostInsightsSortOrderEnum
func GetListHostInsightsSortOrderEnumValues() []ListHostInsightsSortOrderEnum {
	values := make([]ListHostInsightsSortOrderEnum, 0)
	for _, v := range mappingListHostInsightsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListHostInsightsSortOrderEnumStringValues Enumerates the set of values in String for ListHostInsightsSortOrderEnum
func GetListHostInsightsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListHostInsightsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHostInsightsSortOrderEnum(val string) (ListHostInsightsSortOrderEnum, bool) {
	enum, ok := mappingListHostInsightsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListHostInsightsSortByEnum Enum with underlying type: string
type ListHostInsightsSortByEnum string

// Set of constants representing the allowable values for ListHostInsightsSortByEnum
const (
	ListHostInsightsSortByHostname ListHostInsightsSortByEnum = "hostName"
	ListHostInsightsSortByHosttype ListHostInsightsSortByEnum = "hostType"
)

var mappingListHostInsightsSortByEnum = map[string]ListHostInsightsSortByEnum{
	"hostName": ListHostInsightsSortByHostname,
	"hostType": ListHostInsightsSortByHosttype,
}

var mappingListHostInsightsSortByEnumLowerCase = map[string]ListHostInsightsSortByEnum{
	"hostname": ListHostInsightsSortByHostname,
	"hosttype": ListHostInsightsSortByHosttype,
}

// GetListHostInsightsSortByEnumValues Enumerates the set of values for ListHostInsightsSortByEnum
func GetListHostInsightsSortByEnumValues() []ListHostInsightsSortByEnum {
	values := make([]ListHostInsightsSortByEnum, 0)
	for _, v := range mappingListHostInsightsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListHostInsightsSortByEnumStringValues Enumerates the set of values in String for ListHostInsightsSortByEnum
func GetListHostInsightsSortByEnumStringValues() []string {
	return []string{
		"hostName",
		"hostType",
	}
}

// GetMappingListHostInsightsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHostInsightsSortByEnum(val string) (ListHostInsightsSortByEnum, bool) {
	enum, ok := mappingListHostInsightsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
