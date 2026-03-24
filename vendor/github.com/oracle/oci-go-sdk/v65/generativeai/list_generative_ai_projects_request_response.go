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

// ListGenerativeAiProjectsRequest wrapper for the ListGenerativeAiProjects operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListGenerativeAiProjects.go.html to see an example of how to use ListGenerativeAiProjectsRequest.
type ListGenerativeAiProjectsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycle state matches the given value.
	LifecycleState GenerativeAiProjectLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the generativeAiProject.
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
	SortOrder ListGenerativeAiProjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListGenerativeAiProjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGenerativeAiProjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGenerativeAiProjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGenerativeAiProjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGenerativeAiProjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGenerativeAiProjectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenerativeAiProjectLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetGenerativeAiProjectLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGenerativeAiProjectsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGenerativeAiProjectsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGenerativeAiProjectsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGenerativeAiProjectsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGenerativeAiProjectsResponse wrapper for the ListGenerativeAiProjects operation
type ListGenerativeAiProjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GenerativeAiProjectCollection instances
	GenerativeAiProjectCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGenerativeAiProjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGenerativeAiProjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGenerativeAiProjectsSortOrderEnum Enum with underlying type: string
type ListGenerativeAiProjectsSortOrderEnum string

// Set of constants representing the allowable values for ListGenerativeAiProjectsSortOrderEnum
const (
	ListGenerativeAiProjectsSortOrderAsc  ListGenerativeAiProjectsSortOrderEnum = "ASC"
	ListGenerativeAiProjectsSortOrderDesc ListGenerativeAiProjectsSortOrderEnum = "DESC"
)

var mappingListGenerativeAiProjectsSortOrderEnum = map[string]ListGenerativeAiProjectsSortOrderEnum{
	"ASC":  ListGenerativeAiProjectsSortOrderAsc,
	"DESC": ListGenerativeAiProjectsSortOrderDesc,
}

var mappingListGenerativeAiProjectsSortOrderEnumLowerCase = map[string]ListGenerativeAiProjectsSortOrderEnum{
	"asc":  ListGenerativeAiProjectsSortOrderAsc,
	"desc": ListGenerativeAiProjectsSortOrderDesc,
}

// GetListGenerativeAiProjectsSortOrderEnumValues Enumerates the set of values for ListGenerativeAiProjectsSortOrderEnum
func GetListGenerativeAiProjectsSortOrderEnumValues() []ListGenerativeAiProjectsSortOrderEnum {
	values := make([]ListGenerativeAiProjectsSortOrderEnum, 0)
	for _, v := range mappingListGenerativeAiProjectsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGenerativeAiProjectsSortOrderEnumStringValues Enumerates the set of values in String for ListGenerativeAiProjectsSortOrderEnum
func GetListGenerativeAiProjectsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGenerativeAiProjectsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGenerativeAiProjectsSortOrderEnum(val string) (ListGenerativeAiProjectsSortOrderEnum, bool) {
	enum, ok := mappingListGenerativeAiProjectsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGenerativeAiProjectsSortByEnum Enum with underlying type: string
type ListGenerativeAiProjectsSortByEnum string

// Set of constants representing the allowable values for ListGenerativeAiProjectsSortByEnum
const (
	ListGenerativeAiProjectsSortByDisplayname ListGenerativeAiProjectsSortByEnum = "displayName"
	ListGenerativeAiProjectsSortByTimecreated ListGenerativeAiProjectsSortByEnum = "timeCreated"
)

var mappingListGenerativeAiProjectsSortByEnum = map[string]ListGenerativeAiProjectsSortByEnum{
	"displayName": ListGenerativeAiProjectsSortByDisplayname,
	"timeCreated": ListGenerativeAiProjectsSortByTimecreated,
}

var mappingListGenerativeAiProjectsSortByEnumLowerCase = map[string]ListGenerativeAiProjectsSortByEnum{
	"displayname": ListGenerativeAiProjectsSortByDisplayname,
	"timecreated": ListGenerativeAiProjectsSortByTimecreated,
}

// GetListGenerativeAiProjectsSortByEnumValues Enumerates the set of values for ListGenerativeAiProjectsSortByEnum
func GetListGenerativeAiProjectsSortByEnumValues() []ListGenerativeAiProjectsSortByEnum {
	values := make([]ListGenerativeAiProjectsSortByEnum, 0)
	for _, v := range mappingListGenerativeAiProjectsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGenerativeAiProjectsSortByEnumStringValues Enumerates the set of values in String for ListGenerativeAiProjectsSortByEnum
func GetListGenerativeAiProjectsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListGenerativeAiProjectsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGenerativeAiProjectsSortByEnum(val string) (ListGenerativeAiProjectsSortByEnum, bool) {
	enum, ok := mappingListGenerativeAiProjectsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
