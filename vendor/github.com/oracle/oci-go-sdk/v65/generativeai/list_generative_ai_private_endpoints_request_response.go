// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListGenerativeAiPrivateEndpointsRequest wrapper for the ListGenerativeAiPrivateEndpoints operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListGenerativeAiPrivateEndpoints.go.html to see an example of how to use ListGenerativeAiPrivateEndpointsRequest.
type ListGenerativeAiPrivateEndpointsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The lifecycle state of Generative AI private endpoints.
	LifecycleState GenerativeAiPrivateEndpointLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field used to sort the results. Multiple fields aren't supported.
	SortBy ListGenerativeAiPrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListGenerativeAiPrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGenerativeAiPrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGenerativeAiPrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGenerativeAiPrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGenerativeAiPrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGenerativeAiPrivateEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenerativeAiPrivateEndpointLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetGenerativeAiPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGenerativeAiPrivateEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGenerativeAiPrivateEndpointsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGenerativeAiPrivateEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGenerativeAiPrivateEndpointsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGenerativeAiPrivateEndpointsResponse wrapper for the ListGenerativeAiPrivateEndpoints operation
type ListGenerativeAiPrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GenerativeAiPrivateEndpointCollection instances
	GenerativeAiPrivateEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGenerativeAiPrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGenerativeAiPrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGenerativeAiPrivateEndpointsSortByEnum Enum with underlying type: string
type ListGenerativeAiPrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListGenerativeAiPrivateEndpointsSortByEnum
const (
	ListGenerativeAiPrivateEndpointsSortByTimecreated ListGenerativeAiPrivateEndpointsSortByEnum = "timeCreated"
)

var mappingListGenerativeAiPrivateEndpointsSortByEnum = map[string]ListGenerativeAiPrivateEndpointsSortByEnum{
	"timeCreated": ListGenerativeAiPrivateEndpointsSortByTimecreated,
}

var mappingListGenerativeAiPrivateEndpointsSortByEnumLowerCase = map[string]ListGenerativeAiPrivateEndpointsSortByEnum{
	"timecreated": ListGenerativeAiPrivateEndpointsSortByTimecreated,
}

// GetListGenerativeAiPrivateEndpointsSortByEnumValues Enumerates the set of values for ListGenerativeAiPrivateEndpointsSortByEnum
func GetListGenerativeAiPrivateEndpointsSortByEnumValues() []ListGenerativeAiPrivateEndpointsSortByEnum {
	values := make([]ListGenerativeAiPrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListGenerativeAiPrivateEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGenerativeAiPrivateEndpointsSortByEnumStringValues Enumerates the set of values in String for ListGenerativeAiPrivateEndpointsSortByEnum
func GetListGenerativeAiPrivateEndpointsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListGenerativeAiPrivateEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGenerativeAiPrivateEndpointsSortByEnum(val string) (ListGenerativeAiPrivateEndpointsSortByEnum, bool) {
	enum, ok := mappingListGenerativeAiPrivateEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGenerativeAiPrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListGenerativeAiPrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListGenerativeAiPrivateEndpointsSortOrderEnum
const (
	ListGenerativeAiPrivateEndpointsSortOrderAsc  ListGenerativeAiPrivateEndpointsSortOrderEnum = "ASC"
	ListGenerativeAiPrivateEndpointsSortOrderDesc ListGenerativeAiPrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListGenerativeAiPrivateEndpointsSortOrderEnum = map[string]ListGenerativeAiPrivateEndpointsSortOrderEnum{
	"ASC":  ListGenerativeAiPrivateEndpointsSortOrderAsc,
	"DESC": ListGenerativeAiPrivateEndpointsSortOrderDesc,
}

var mappingListGenerativeAiPrivateEndpointsSortOrderEnumLowerCase = map[string]ListGenerativeAiPrivateEndpointsSortOrderEnum{
	"asc":  ListGenerativeAiPrivateEndpointsSortOrderAsc,
	"desc": ListGenerativeAiPrivateEndpointsSortOrderDesc,
}

// GetListGenerativeAiPrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListGenerativeAiPrivateEndpointsSortOrderEnum
func GetListGenerativeAiPrivateEndpointsSortOrderEnumValues() []ListGenerativeAiPrivateEndpointsSortOrderEnum {
	values := make([]ListGenerativeAiPrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListGenerativeAiPrivateEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGenerativeAiPrivateEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListGenerativeAiPrivateEndpointsSortOrderEnum
func GetListGenerativeAiPrivateEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGenerativeAiPrivateEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGenerativeAiPrivateEndpointsSortOrderEnum(val string) (ListGenerativeAiPrivateEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListGenerativeAiPrivateEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
