// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEmailPrivateEndpointsRequest wrapper for the ListEmailPrivateEndpoints operation
type ListEmailPrivateEndpointsRequest struct {

	// The OCID for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The request ID for tracing from the system
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to only return resources that match the given id exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. `1` is the minimum, `1000` is the maximum. For important details about
	// how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending or descending order.
	SortOrder ListEmailPrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used to sort the results. Multiple fields are not supported.
	SortBy ListEmailPrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter returned list by specified lifecycle state. This parameter is case-insensitive.
	LifecycleState EmailPrivateEndpointLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEmailPrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEmailPrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEmailPrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEmailPrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEmailPrivateEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEmailPrivateEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEmailPrivateEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEmailPrivateEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEmailPrivateEndpointsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmailPrivateEndpointLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetEmailPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEmailPrivateEndpointsResponse wrapper for the ListEmailPrivateEndpoints operation
type ListEmailPrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EmailPrivateEndpointCollection instances
	EmailPrivateEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEmailPrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEmailPrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEmailPrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListEmailPrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListEmailPrivateEndpointsSortOrderEnum
const (
	ListEmailPrivateEndpointsSortOrderAsc  ListEmailPrivateEndpointsSortOrderEnum = "ASC"
	ListEmailPrivateEndpointsSortOrderDesc ListEmailPrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListEmailPrivateEndpointsSortOrderEnum = map[string]ListEmailPrivateEndpointsSortOrderEnum{
	"ASC":  ListEmailPrivateEndpointsSortOrderAsc,
	"DESC": ListEmailPrivateEndpointsSortOrderDesc,
}

var mappingListEmailPrivateEndpointsSortOrderEnumLowerCase = map[string]ListEmailPrivateEndpointsSortOrderEnum{
	"asc":  ListEmailPrivateEndpointsSortOrderAsc,
	"desc": ListEmailPrivateEndpointsSortOrderDesc,
}

// GetListEmailPrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListEmailPrivateEndpointsSortOrderEnum
func GetListEmailPrivateEndpointsSortOrderEnumValues() []ListEmailPrivateEndpointsSortOrderEnum {
	values := make([]ListEmailPrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListEmailPrivateEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailPrivateEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListEmailPrivateEndpointsSortOrderEnum
func GetListEmailPrivateEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEmailPrivateEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailPrivateEndpointsSortOrderEnum(val string) (ListEmailPrivateEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListEmailPrivateEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEmailPrivateEndpointsSortByEnum Enum with underlying type: string
type ListEmailPrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListEmailPrivateEndpointsSortByEnum
const (
	ListEmailPrivateEndpointsSortByTimecreated ListEmailPrivateEndpointsSortByEnum = "TIMECREATED"
	ListEmailPrivateEndpointsSortByDisplayname ListEmailPrivateEndpointsSortByEnum = "DISPLAYNAME"
)

var mappingListEmailPrivateEndpointsSortByEnum = map[string]ListEmailPrivateEndpointsSortByEnum{
	"TIMECREATED": ListEmailPrivateEndpointsSortByTimecreated,
	"DISPLAYNAME": ListEmailPrivateEndpointsSortByDisplayname,
}

var mappingListEmailPrivateEndpointsSortByEnumLowerCase = map[string]ListEmailPrivateEndpointsSortByEnum{
	"timecreated": ListEmailPrivateEndpointsSortByTimecreated,
	"displayname": ListEmailPrivateEndpointsSortByDisplayname,
}

// GetListEmailPrivateEndpointsSortByEnumValues Enumerates the set of values for ListEmailPrivateEndpointsSortByEnum
func GetListEmailPrivateEndpointsSortByEnumValues() []ListEmailPrivateEndpointsSortByEnum {
	values := make([]ListEmailPrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListEmailPrivateEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailPrivateEndpointsSortByEnumStringValues Enumerates the set of values in String for ListEmailPrivateEndpointsSortByEnum
func GetListEmailPrivateEndpointsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListEmailPrivateEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailPrivateEndpointsSortByEnum(val string) (ListEmailPrivateEndpointsSortByEnum, bool) {
	enum, ok := mappingListEmailPrivateEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
