// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRoutingProfilesRequest wrapper for the ListRoutingProfiles operation
type ListRoutingProfilesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycle state matches the given value.
	LifecycleState RoutingProfileLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the routing profile.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListRoutingProfilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListRoutingProfilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRoutingProfilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRoutingProfilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRoutingProfilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRoutingProfilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRoutingProfilesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRoutingProfileLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRoutingProfileLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRoutingProfilesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRoutingProfilesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRoutingProfilesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRoutingProfilesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRoutingProfilesResponse wrapper for the ListRoutingProfiles operation
type ListRoutingProfilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RoutingProfileCollection instances
	RoutingProfileCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRoutingProfilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRoutingProfilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRoutingProfilesSortOrderEnum Enum with underlying type: string
type ListRoutingProfilesSortOrderEnum string

// Set of constants representing the allowable values for ListRoutingProfilesSortOrderEnum
const (
	ListRoutingProfilesSortOrderAsc  ListRoutingProfilesSortOrderEnum = "ASC"
	ListRoutingProfilesSortOrderDesc ListRoutingProfilesSortOrderEnum = "DESC"
)

var mappingListRoutingProfilesSortOrderEnum = map[string]ListRoutingProfilesSortOrderEnum{
	"ASC":  ListRoutingProfilesSortOrderAsc,
	"DESC": ListRoutingProfilesSortOrderDesc,
}

var mappingListRoutingProfilesSortOrderEnumLowerCase = map[string]ListRoutingProfilesSortOrderEnum{
	"asc":  ListRoutingProfilesSortOrderAsc,
	"desc": ListRoutingProfilesSortOrderDesc,
}

// GetListRoutingProfilesSortOrderEnumValues Enumerates the set of values for ListRoutingProfilesSortOrderEnum
func GetListRoutingProfilesSortOrderEnumValues() []ListRoutingProfilesSortOrderEnum {
	values := make([]ListRoutingProfilesSortOrderEnum, 0)
	for _, v := range mappingListRoutingProfilesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRoutingProfilesSortOrderEnumStringValues Enumerates the set of values in String for ListRoutingProfilesSortOrderEnum
func GetListRoutingProfilesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRoutingProfilesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRoutingProfilesSortOrderEnum(val string) (ListRoutingProfilesSortOrderEnum, bool) {
	enum, ok := mappingListRoutingProfilesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRoutingProfilesSortByEnum Enum with underlying type: string
type ListRoutingProfilesSortByEnum string

// Set of constants representing the allowable values for ListRoutingProfilesSortByEnum
const (
	ListRoutingProfilesSortByDisplayname ListRoutingProfilesSortByEnum = "displayName"
	ListRoutingProfilesSortByTimecreated ListRoutingProfilesSortByEnum = "timeCreated"
)

var mappingListRoutingProfilesSortByEnum = map[string]ListRoutingProfilesSortByEnum{
	"displayName": ListRoutingProfilesSortByDisplayname,
	"timeCreated": ListRoutingProfilesSortByTimecreated,
}

var mappingListRoutingProfilesSortByEnumLowerCase = map[string]ListRoutingProfilesSortByEnum{
	"displayname": ListRoutingProfilesSortByDisplayname,
	"timecreated": ListRoutingProfilesSortByTimecreated,
}

// GetListRoutingProfilesSortByEnumValues Enumerates the set of values for ListRoutingProfilesSortByEnum
func GetListRoutingProfilesSortByEnumValues() []ListRoutingProfilesSortByEnum {
	values := make([]ListRoutingProfilesSortByEnum, 0)
	for _, v := range mappingListRoutingProfilesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRoutingProfilesSortByEnumStringValues Enumerates the set of values in String for ListRoutingProfilesSortByEnum
func GetListRoutingProfilesSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListRoutingProfilesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRoutingProfilesSortByEnum(val string) (ListRoutingProfilesSortByEnum, bool) {
	enum, ok := mappingListRoutingProfilesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
