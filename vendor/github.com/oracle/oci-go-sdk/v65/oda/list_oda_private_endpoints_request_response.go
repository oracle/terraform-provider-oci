// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOdaPrivateEndpointsRequest wrapper for the ListOdaPrivateEndpoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListOdaPrivateEndpoints.go.html to see an example of how to use ListOdaPrivateEndpointsRequest.
type ListOdaPrivateEndpointsRequest struct {

	// List the ODA Private Endpoints that belong to this compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// List only the information for the Digital Assistant instance with this user-friendly name. These names don't have to be unique and may change.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// List only the ODA Private Endpoints that are in this lifecycle state.
	LifecycleState OdaPrivateEndpointLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListOdaPrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `TIMECREATED`.
	// The default sort order for `TIMECREATED` is descending, and the default sort order for `DISPLAYNAME` is ascending.
	SortBy ListOdaPrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOdaPrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOdaPrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOdaPrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOdaPrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOdaPrivateEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOdaPrivateEndpointLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOdaPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOdaPrivateEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOdaPrivateEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOdaPrivateEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOdaPrivateEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOdaPrivateEndpointsResponse wrapper for the ListOdaPrivateEndpoints operation
type ListOdaPrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OdaPrivateEndpointCollection instances
	OdaPrivateEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// When you are paging through a list, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the
	// `page` query parameter for the subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of results that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListOdaPrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOdaPrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOdaPrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListOdaPrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListOdaPrivateEndpointsSortOrderEnum
const (
	ListOdaPrivateEndpointsSortOrderAsc  ListOdaPrivateEndpointsSortOrderEnum = "ASC"
	ListOdaPrivateEndpointsSortOrderDesc ListOdaPrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListOdaPrivateEndpointsSortOrderEnum = map[string]ListOdaPrivateEndpointsSortOrderEnum{
	"ASC":  ListOdaPrivateEndpointsSortOrderAsc,
	"DESC": ListOdaPrivateEndpointsSortOrderDesc,
}

var mappingListOdaPrivateEndpointsSortOrderEnumLowerCase = map[string]ListOdaPrivateEndpointsSortOrderEnum{
	"asc":  ListOdaPrivateEndpointsSortOrderAsc,
	"desc": ListOdaPrivateEndpointsSortOrderDesc,
}

// GetListOdaPrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListOdaPrivateEndpointsSortOrderEnum
func GetListOdaPrivateEndpointsSortOrderEnumValues() []ListOdaPrivateEndpointsSortOrderEnum {
	values := make([]ListOdaPrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListOdaPrivateEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOdaPrivateEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListOdaPrivateEndpointsSortOrderEnum
func GetListOdaPrivateEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOdaPrivateEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOdaPrivateEndpointsSortOrderEnum(val string) (ListOdaPrivateEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListOdaPrivateEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOdaPrivateEndpointsSortByEnum Enum with underlying type: string
type ListOdaPrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListOdaPrivateEndpointsSortByEnum
const (
	ListOdaPrivateEndpointsSortByTimecreated ListOdaPrivateEndpointsSortByEnum = "TIMECREATED"
	ListOdaPrivateEndpointsSortByDisplayname ListOdaPrivateEndpointsSortByEnum = "DISPLAYNAME"
)

var mappingListOdaPrivateEndpointsSortByEnum = map[string]ListOdaPrivateEndpointsSortByEnum{
	"TIMECREATED": ListOdaPrivateEndpointsSortByTimecreated,
	"DISPLAYNAME": ListOdaPrivateEndpointsSortByDisplayname,
}

var mappingListOdaPrivateEndpointsSortByEnumLowerCase = map[string]ListOdaPrivateEndpointsSortByEnum{
	"timecreated": ListOdaPrivateEndpointsSortByTimecreated,
	"displayname": ListOdaPrivateEndpointsSortByDisplayname,
}

// GetListOdaPrivateEndpointsSortByEnumValues Enumerates the set of values for ListOdaPrivateEndpointsSortByEnum
func GetListOdaPrivateEndpointsSortByEnumValues() []ListOdaPrivateEndpointsSortByEnum {
	values := make([]ListOdaPrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListOdaPrivateEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOdaPrivateEndpointsSortByEnumStringValues Enumerates the set of values in String for ListOdaPrivateEndpointsSortByEnum
func GetListOdaPrivateEndpointsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListOdaPrivateEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOdaPrivateEndpointsSortByEnum(val string) (ListOdaPrivateEndpointsSortByEnum, bool) {
	enum, ok := mappingListOdaPrivateEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
