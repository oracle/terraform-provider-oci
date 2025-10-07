// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package aidataplatform

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAiDataPlatformsRequest wrapper for the ListAiDataPlatforms operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aidataplatform/ListAiDataPlatforms.go.html to see an example of how to use ListAiDataPlatformsRequest.
type ListAiDataPlatformsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState AiDataPlatformLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to exclude resources that match the given lifecycle state. The
	// state value is case-insensitive.
	ExcludeLifecycleState AiDataPlatformLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"excludeLifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the AiDataPlatform.
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
	SortOrder ListAiDataPlatformsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListAiDataPlatformsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// This flag will determine if legacy instances will be returned.
	IncludeLegacy *string `mandatory:"false" contributesTo:"query" name:"includeLegacy"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAiDataPlatformsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAiDataPlatformsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAiDataPlatformsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAiDataPlatformsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAiDataPlatformsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAiDataPlatformLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAiDataPlatformLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAiDataPlatformLifecycleStateEnum(string(request.ExcludeLifecycleState)); !ok && request.ExcludeLifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExcludeLifecycleState: %s. Supported values are: %s.", request.ExcludeLifecycleState, strings.Join(GetAiDataPlatformLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAiDataPlatformsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAiDataPlatformsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAiDataPlatformsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAiDataPlatformsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAiDataPlatformsResponse wrapper for the ListAiDataPlatforms operation
type ListAiDataPlatformsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AiDataPlatformCollection instances
	AiDataPlatformCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAiDataPlatformsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAiDataPlatformsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAiDataPlatformsSortOrderEnum Enum with underlying type: string
type ListAiDataPlatformsSortOrderEnum string

// Set of constants representing the allowable values for ListAiDataPlatformsSortOrderEnum
const (
	ListAiDataPlatformsSortOrderAsc  ListAiDataPlatformsSortOrderEnum = "ASC"
	ListAiDataPlatformsSortOrderDesc ListAiDataPlatformsSortOrderEnum = "DESC"
)

var mappingListAiDataPlatformsSortOrderEnum = map[string]ListAiDataPlatformsSortOrderEnum{
	"ASC":  ListAiDataPlatformsSortOrderAsc,
	"DESC": ListAiDataPlatformsSortOrderDesc,
}

var mappingListAiDataPlatformsSortOrderEnumLowerCase = map[string]ListAiDataPlatformsSortOrderEnum{
	"asc":  ListAiDataPlatformsSortOrderAsc,
	"desc": ListAiDataPlatformsSortOrderDesc,
}

// GetListAiDataPlatformsSortOrderEnumValues Enumerates the set of values for ListAiDataPlatformsSortOrderEnum
func GetListAiDataPlatformsSortOrderEnumValues() []ListAiDataPlatformsSortOrderEnum {
	values := make([]ListAiDataPlatformsSortOrderEnum, 0)
	for _, v := range mappingListAiDataPlatformsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAiDataPlatformsSortOrderEnumStringValues Enumerates the set of values in String for ListAiDataPlatformsSortOrderEnum
func GetListAiDataPlatformsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAiDataPlatformsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAiDataPlatformsSortOrderEnum(val string) (ListAiDataPlatformsSortOrderEnum, bool) {
	enum, ok := mappingListAiDataPlatformsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAiDataPlatformsSortByEnum Enum with underlying type: string
type ListAiDataPlatformsSortByEnum string

// Set of constants representing the allowable values for ListAiDataPlatformsSortByEnum
const (
	ListAiDataPlatformsSortByTimecreated ListAiDataPlatformsSortByEnum = "timeCreated"
	ListAiDataPlatformsSortByDisplayname ListAiDataPlatformsSortByEnum = "displayName"
)

var mappingListAiDataPlatformsSortByEnum = map[string]ListAiDataPlatformsSortByEnum{
	"timeCreated": ListAiDataPlatformsSortByTimecreated,
	"displayName": ListAiDataPlatformsSortByDisplayname,
}

var mappingListAiDataPlatformsSortByEnumLowerCase = map[string]ListAiDataPlatformsSortByEnum{
	"timecreated": ListAiDataPlatformsSortByTimecreated,
	"displayname": ListAiDataPlatformsSortByDisplayname,
}

// GetListAiDataPlatformsSortByEnumValues Enumerates the set of values for ListAiDataPlatformsSortByEnum
func GetListAiDataPlatformsSortByEnumValues() []ListAiDataPlatformsSortByEnum {
	values := make([]ListAiDataPlatformsSortByEnum, 0)
	for _, v := range mappingListAiDataPlatformsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAiDataPlatformsSortByEnumStringValues Enumerates the set of values in String for ListAiDataPlatformsSortByEnum
func GetListAiDataPlatformsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAiDataPlatformsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAiDataPlatformsSortByEnum(val string) (ListAiDataPlatformsSortByEnum, bool) {
	enum, ok := mappingListAiDataPlatformsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
