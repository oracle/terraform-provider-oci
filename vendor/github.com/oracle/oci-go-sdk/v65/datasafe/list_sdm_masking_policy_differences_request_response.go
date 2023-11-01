// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSdmMaskingPolicyDifferencesRequest wrapper for the ListSdmMaskingPolicyDifferences operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSdmMaskingPolicyDifferences.go.html to see an example of how to use ListSdmMaskingPolicyDifferencesRequest.
type ListSdmMaskingPolicyDifferencesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid value is ACCESSIBLE. Default is ACCESSIBLE.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	DifferenceAccessLevel ListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"differenceAccessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only the resources that match the specified sensitive data model OCID.
	SensitiveDataModelId *string `mandatory:"false" contributesTo:"query" name:"sensitiveDataModelId"`

	// A filter to return only the resources that match the specified lifecycle states.
	LifecycleState SdmMaskingPolicyDifferenceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the specified masking policy OCID.
	MaskingPolicyId *string `mandatory:"false" contributesTo:"query" name:"maskingPolicyId"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSdmMaskingPolicyDifferencesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for timeCreationStarted is descending.
	// The default order for displayName is ascending.
	SortBy ListSdmMaskingPolicyDifferencesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSdmMaskingPolicyDifferencesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSdmMaskingPolicyDifferencesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSdmMaskingPolicyDifferencesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSdmMaskingPolicyDifferencesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSdmMaskingPolicyDifferencesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum(string(request.DifferenceAccessLevel)); !ok && request.DifferenceAccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DifferenceAccessLevel: %s. Supported values are: %s.", request.DifferenceAccessLevel, strings.Join(GetListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSdmMaskingPolicyDifferenceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetSdmMaskingPolicyDifferenceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSdmMaskingPolicyDifferencesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSdmMaskingPolicyDifferencesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSdmMaskingPolicyDifferencesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSdmMaskingPolicyDifferencesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSdmMaskingPolicyDifferencesResponse wrapper for the ListSdmMaskingPolicyDifferences operation
type ListSdmMaskingPolicyDifferencesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SdmMaskingPolicyDifferenceCollection instances
	SdmMaskingPolicyDifferenceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSdmMaskingPolicyDifferencesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSdmMaskingPolicyDifferencesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum Enum with underlying type: string
type ListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum string

// Set of constants representing the allowable values for ListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum
const (
	ListSdmMaskingPolicyDifferencesDifferenceAccessLevelAccessible ListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum = map[string]ListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum{
	"ACCESSIBLE": ListSdmMaskingPolicyDifferencesDifferenceAccessLevelAccessible,
}

var mappingListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnumLowerCase = map[string]ListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum{
	"accessible": ListSdmMaskingPolicyDifferencesDifferenceAccessLevelAccessible,
}

// GetListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnumValues Enumerates the set of values for ListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum
func GetListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnumValues() []ListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum {
	values := make([]ListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum, 0)
	for _, v := range mappingListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnumStringValues Enumerates the set of values in String for ListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum
func GetListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnumStringValues() []string {
	return []string{
		"ACCESSIBLE",
	}
}

// GetMappingListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum(val string) (ListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum, bool) {
	enum, ok := mappingListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSdmMaskingPolicyDifferencesSortOrderEnum Enum with underlying type: string
type ListSdmMaskingPolicyDifferencesSortOrderEnum string

// Set of constants representing the allowable values for ListSdmMaskingPolicyDifferencesSortOrderEnum
const (
	ListSdmMaskingPolicyDifferencesSortOrderAsc  ListSdmMaskingPolicyDifferencesSortOrderEnum = "ASC"
	ListSdmMaskingPolicyDifferencesSortOrderDesc ListSdmMaskingPolicyDifferencesSortOrderEnum = "DESC"
)

var mappingListSdmMaskingPolicyDifferencesSortOrderEnum = map[string]ListSdmMaskingPolicyDifferencesSortOrderEnum{
	"ASC":  ListSdmMaskingPolicyDifferencesSortOrderAsc,
	"DESC": ListSdmMaskingPolicyDifferencesSortOrderDesc,
}

var mappingListSdmMaskingPolicyDifferencesSortOrderEnumLowerCase = map[string]ListSdmMaskingPolicyDifferencesSortOrderEnum{
	"asc":  ListSdmMaskingPolicyDifferencesSortOrderAsc,
	"desc": ListSdmMaskingPolicyDifferencesSortOrderDesc,
}

// GetListSdmMaskingPolicyDifferencesSortOrderEnumValues Enumerates the set of values for ListSdmMaskingPolicyDifferencesSortOrderEnum
func GetListSdmMaskingPolicyDifferencesSortOrderEnumValues() []ListSdmMaskingPolicyDifferencesSortOrderEnum {
	values := make([]ListSdmMaskingPolicyDifferencesSortOrderEnum, 0)
	for _, v := range mappingListSdmMaskingPolicyDifferencesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSdmMaskingPolicyDifferencesSortOrderEnumStringValues Enumerates the set of values in String for ListSdmMaskingPolicyDifferencesSortOrderEnum
func GetListSdmMaskingPolicyDifferencesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSdmMaskingPolicyDifferencesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSdmMaskingPolicyDifferencesSortOrderEnum(val string) (ListSdmMaskingPolicyDifferencesSortOrderEnum, bool) {
	enum, ok := mappingListSdmMaskingPolicyDifferencesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSdmMaskingPolicyDifferencesSortByEnum Enum with underlying type: string
type ListSdmMaskingPolicyDifferencesSortByEnum string

// Set of constants representing the allowable values for ListSdmMaskingPolicyDifferencesSortByEnum
const (
	ListSdmMaskingPolicyDifferencesSortByTimecreationstarted ListSdmMaskingPolicyDifferencesSortByEnum = "timeCreationStarted"
	ListSdmMaskingPolicyDifferencesSortByDisplayname         ListSdmMaskingPolicyDifferencesSortByEnum = "displayName"
)

var mappingListSdmMaskingPolicyDifferencesSortByEnum = map[string]ListSdmMaskingPolicyDifferencesSortByEnum{
	"timeCreationStarted": ListSdmMaskingPolicyDifferencesSortByTimecreationstarted,
	"displayName":         ListSdmMaskingPolicyDifferencesSortByDisplayname,
}

var mappingListSdmMaskingPolicyDifferencesSortByEnumLowerCase = map[string]ListSdmMaskingPolicyDifferencesSortByEnum{
	"timecreationstarted": ListSdmMaskingPolicyDifferencesSortByTimecreationstarted,
	"displayname":         ListSdmMaskingPolicyDifferencesSortByDisplayname,
}

// GetListSdmMaskingPolicyDifferencesSortByEnumValues Enumerates the set of values for ListSdmMaskingPolicyDifferencesSortByEnum
func GetListSdmMaskingPolicyDifferencesSortByEnumValues() []ListSdmMaskingPolicyDifferencesSortByEnum {
	values := make([]ListSdmMaskingPolicyDifferencesSortByEnum, 0)
	for _, v := range mappingListSdmMaskingPolicyDifferencesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSdmMaskingPolicyDifferencesSortByEnumStringValues Enumerates the set of values in String for ListSdmMaskingPolicyDifferencesSortByEnum
func GetListSdmMaskingPolicyDifferencesSortByEnumStringValues() []string {
	return []string{
		"timeCreationStarted",
		"displayName",
	}
}

// GetMappingListSdmMaskingPolicyDifferencesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSdmMaskingPolicyDifferencesSortByEnum(val string) (ListSdmMaskingPolicyDifferencesSortByEnum, bool) {
	enum, ok := mappingListSdmMaskingPolicyDifferencesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
