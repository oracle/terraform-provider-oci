// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package iot

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListIotDomainGroupsRequest wrapper for the ListIotDomainGroups operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListIotDomainGroups.go.html to see an example of how to use ListIotDomainGroupsRequest.
type ListIotDomainGroupsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter resources by OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be a valid OCID of the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Filter resources whose display name matches the specified value.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter resources whose lifecycleState matches the specified value.
	LifecycleState IotDomainGroupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination: The value of the opc-next-page response header from the previous "List" call.
	// For important details on how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either ASC (ascending) or DESC (descending).
	SortOrder ListIotDomainGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListIotDomainGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIotDomainGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIotDomainGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIotDomainGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIotDomainGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListIotDomainGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIotDomainGroupLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetIotDomainGroupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIotDomainGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListIotDomainGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIotDomainGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListIotDomainGroupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIotDomainGroupsResponse wrapper for the ListIotDomainGroups operation
type ListIotDomainGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of IotDomainGroupCollection instances
	IotDomainGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListIotDomainGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIotDomainGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIotDomainGroupsSortOrderEnum Enum with underlying type: string
type ListIotDomainGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListIotDomainGroupsSortOrderEnum
const (
	ListIotDomainGroupsSortOrderAsc  ListIotDomainGroupsSortOrderEnum = "ASC"
	ListIotDomainGroupsSortOrderDesc ListIotDomainGroupsSortOrderEnum = "DESC"
)

var mappingListIotDomainGroupsSortOrderEnum = map[string]ListIotDomainGroupsSortOrderEnum{
	"ASC":  ListIotDomainGroupsSortOrderAsc,
	"DESC": ListIotDomainGroupsSortOrderDesc,
}

var mappingListIotDomainGroupsSortOrderEnumLowerCase = map[string]ListIotDomainGroupsSortOrderEnum{
	"asc":  ListIotDomainGroupsSortOrderAsc,
	"desc": ListIotDomainGroupsSortOrderDesc,
}

// GetListIotDomainGroupsSortOrderEnumValues Enumerates the set of values for ListIotDomainGroupsSortOrderEnum
func GetListIotDomainGroupsSortOrderEnumValues() []ListIotDomainGroupsSortOrderEnum {
	values := make([]ListIotDomainGroupsSortOrderEnum, 0)
	for _, v := range mappingListIotDomainGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIotDomainGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListIotDomainGroupsSortOrderEnum
func GetListIotDomainGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIotDomainGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIotDomainGroupsSortOrderEnum(val string) (ListIotDomainGroupsSortOrderEnum, bool) {
	enum, ok := mappingListIotDomainGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIotDomainGroupsSortByEnum Enum with underlying type: string
type ListIotDomainGroupsSortByEnum string

// Set of constants representing the allowable values for ListIotDomainGroupsSortByEnum
const (
	ListIotDomainGroupsSortByTimecreated ListIotDomainGroupsSortByEnum = "timeCreated"
	ListIotDomainGroupsSortByDisplayname ListIotDomainGroupsSortByEnum = "displayName"
)

var mappingListIotDomainGroupsSortByEnum = map[string]ListIotDomainGroupsSortByEnum{
	"timeCreated": ListIotDomainGroupsSortByTimecreated,
	"displayName": ListIotDomainGroupsSortByDisplayname,
}

var mappingListIotDomainGroupsSortByEnumLowerCase = map[string]ListIotDomainGroupsSortByEnum{
	"timecreated": ListIotDomainGroupsSortByTimecreated,
	"displayname": ListIotDomainGroupsSortByDisplayname,
}

// GetListIotDomainGroupsSortByEnumValues Enumerates the set of values for ListIotDomainGroupsSortByEnum
func GetListIotDomainGroupsSortByEnumValues() []ListIotDomainGroupsSortByEnum {
	values := make([]ListIotDomainGroupsSortByEnum, 0)
	for _, v := range mappingListIotDomainGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListIotDomainGroupsSortByEnumStringValues Enumerates the set of values in String for ListIotDomainGroupsSortByEnum
func GetListIotDomainGroupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListIotDomainGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIotDomainGroupsSortByEnum(val string) (ListIotDomainGroupsSortByEnum, bool) {
	enum, ok := mappingListIotDomainGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
