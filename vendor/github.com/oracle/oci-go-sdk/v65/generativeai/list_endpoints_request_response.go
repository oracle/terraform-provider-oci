// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEndpointsRequest wrapper for the ListEndpoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListEndpoints.go.html to see an example of how to use ListEndpointsRequest.
type ListEndpointsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that their lifecycle state matches the given lifecycle state.
	LifecycleState EndpointLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the endpoint.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEndpointLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetEndpointLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEndpointsResponse wrapper for the ListEndpoints operation
type ListEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EndpointCollection instances
	EndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEndpointsSortOrderEnum Enum with underlying type: string
type ListEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListEndpointsSortOrderEnum
const (
	ListEndpointsSortOrderAsc  ListEndpointsSortOrderEnum = "ASC"
	ListEndpointsSortOrderDesc ListEndpointsSortOrderEnum = "DESC"
)

var mappingListEndpointsSortOrderEnum = map[string]ListEndpointsSortOrderEnum{
	"ASC":  ListEndpointsSortOrderAsc,
	"DESC": ListEndpointsSortOrderDesc,
}

var mappingListEndpointsSortOrderEnumLowerCase = map[string]ListEndpointsSortOrderEnum{
	"asc":  ListEndpointsSortOrderAsc,
	"desc": ListEndpointsSortOrderDesc,
}

// GetListEndpointsSortOrderEnumValues Enumerates the set of values for ListEndpointsSortOrderEnum
func GetListEndpointsSortOrderEnumValues() []ListEndpointsSortOrderEnum {
	values := make([]ListEndpointsSortOrderEnum, 0)
	for _, v := range mappingListEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListEndpointsSortOrderEnum
func GetListEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEndpointsSortOrderEnum(val string) (ListEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEndpointsSortByEnum Enum with underlying type: string
type ListEndpointsSortByEnum string

// Set of constants representing the allowable values for ListEndpointsSortByEnum
const (
	ListEndpointsSortByDisplayname ListEndpointsSortByEnum = "displayName"
	ListEndpointsSortByTimecreated ListEndpointsSortByEnum = "timeCreated"
)

var mappingListEndpointsSortByEnum = map[string]ListEndpointsSortByEnum{
	"displayName": ListEndpointsSortByDisplayname,
	"timeCreated": ListEndpointsSortByTimecreated,
}

var mappingListEndpointsSortByEnumLowerCase = map[string]ListEndpointsSortByEnum{
	"displayname": ListEndpointsSortByDisplayname,
	"timecreated": ListEndpointsSortByTimecreated,
}

// GetListEndpointsSortByEnumValues Enumerates the set of values for ListEndpointsSortByEnum
func GetListEndpointsSortByEnumValues() []ListEndpointsSortByEnum {
	values := make([]ListEndpointsSortByEnum, 0)
	for _, v := range mappingListEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEndpointsSortByEnumStringValues Enumerates the set of values in String for ListEndpointsSortByEnum
func GetListEndpointsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEndpointsSortByEnum(val string) (ListEndpointsSortByEnum, bool) {
	enum, ok := mappingListEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
