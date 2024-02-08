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

// ListOdaPrivateEndpointScanProxiesRequest wrapper for the ListOdaPrivateEndpointScanProxies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListOdaPrivateEndpointScanProxies.go.html to see an example of how to use ListOdaPrivateEndpointScanProxiesRequest.
type ListOdaPrivateEndpointScanProxiesRequest struct {

	// Unique ODA Private Endpoint identifier which is the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	OdaPrivateEndpointId *string `mandatory:"true" contributesTo:"path" name:"odaPrivateEndpointId"`

	// List only the ODA Private Endpoint Scan Proxies that are in this lifecycle state.
	LifecycleState OdaPrivateEndpointScanProxyLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListOdaPrivateEndpointScanProxiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `TIMECREATED`.
	// The default sort order for `TIMECREATED` is descending, and the default sort order for `DISPLAYNAME` is ascending.
	SortBy ListOdaPrivateEndpointScanProxiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOdaPrivateEndpointScanProxiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOdaPrivateEndpointScanProxiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOdaPrivateEndpointScanProxiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOdaPrivateEndpointScanProxiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOdaPrivateEndpointScanProxiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOdaPrivateEndpointScanProxyLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOdaPrivateEndpointScanProxyLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOdaPrivateEndpointScanProxiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOdaPrivateEndpointScanProxiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOdaPrivateEndpointScanProxiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOdaPrivateEndpointScanProxiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOdaPrivateEndpointScanProxiesResponse wrapper for the ListOdaPrivateEndpointScanProxies operation
type ListOdaPrivateEndpointScanProxiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OdaPrivateEndpointScanProxyCollection instances
	OdaPrivateEndpointScanProxyCollection `presentIn:"body"`

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

func (response ListOdaPrivateEndpointScanProxiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOdaPrivateEndpointScanProxiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOdaPrivateEndpointScanProxiesSortOrderEnum Enum with underlying type: string
type ListOdaPrivateEndpointScanProxiesSortOrderEnum string

// Set of constants representing the allowable values for ListOdaPrivateEndpointScanProxiesSortOrderEnum
const (
	ListOdaPrivateEndpointScanProxiesSortOrderAsc  ListOdaPrivateEndpointScanProxiesSortOrderEnum = "ASC"
	ListOdaPrivateEndpointScanProxiesSortOrderDesc ListOdaPrivateEndpointScanProxiesSortOrderEnum = "DESC"
)

var mappingListOdaPrivateEndpointScanProxiesSortOrderEnum = map[string]ListOdaPrivateEndpointScanProxiesSortOrderEnum{
	"ASC":  ListOdaPrivateEndpointScanProxiesSortOrderAsc,
	"DESC": ListOdaPrivateEndpointScanProxiesSortOrderDesc,
}

var mappingListOdaPrivateEndpointScanProxiesSortOrderEnumLowerCase = map[string]ListOdaPrivateEndpointScanProxiesSortOrderEnum{
	"asc":  ListOdaPrivateEndpointScanProxiesSortOrderAsc,
	"desc": ListOdaPrivateEndpointScanProxiesSortOrderDesc,
}

// GetListOdaPrivateEndpointScanProxiesSortOrderEnumValues Enumerates the set of values for ListOdaPrivateEndpointScanProxiesSortOrderEnum
func GetListOdaPrivateEndpointScanProxiesSortOrderEnumValues() []ListOdaPrivateEndpointScanProxiesSortOrderEnum {
	values := make([]ListOdaPrivateEndpointScanProxiesSortOrderEnum, 0)
	for _, v := range mappingListOdaPrivateEndpointScanProxiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOdaPrivateEndpointScanProxiesSortOrderEnumStringValues Enumerates the set of values in String for ListOdaPrivateEndpointScanProxiesSortOrderEnum
func GetListOdaPrivateEndpointScanProxiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOdaPrivateEndpointScanProxiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOdaPrivateEndpointScanProxiesSortOrderEnum(val string) (ListOdaPrivateEndpointScanProxiesSortOrderEnum, bool) {
	enum, ok := mappingListOdaPrivateEndpointScanProxiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOdaPrivateEndpointScanProxiesSortByEnum Enum with underlying type: string
type ListOdaPrivateEndpointScanProxiesSortByEnum string

// Set of constants representing the allowable values for ListOdaPrivateEndpointScanProxiesSortByEnum
const (
	ListOdaPrivateEndpointScanProxiesSortByTimecreated ListOdaPrivateEndpointScanProxiesSortByEnum = "TIMECREATED"
	ListOdaPrivateEndpointScanProxiesSortByDisplayname ListOdaPrivateEndpointScanProxiesSortByEnum = "DISPLAYNAME"
)

var mappingListOdaPrivateEndpointScanProxiesSortByEnum = map[string]ListOdaPrivateEndpointScanProxiesSortByEnum{
	"TIMECREATED": ListOdaPrivateEndpointScanProxiesSortByTimecreated,
	"DISPLAYNAME": ListOdaPrivateEndpointScanProxiesSortByDisplayname,
}

var mappingListOdaPrivateEndpointScanProxiesSortByEnumLowerCase = map[string]ListOdaPrivateEndpointScanProxiesSortByEnum{
	"timecreated": ListOdaPrivateEndpointScanProxiesSortByTimecreated,
	"displayname": ListOdaPrivateEndpointScanProxiesSortByDisplayname,
}

// GetListOdaPrivateEndpointScanProxiesSortByEnumValues Enumerates the set of values for ListOdaPrivateEndpointScanProxiesSortByEnum
func GetListOdaPrivateEndpointScanProxiesSortByEnumValues() []ListOdaPrivateEndpointScanProxiesSortByEnum {
	values := make([]ListOdaPrivateEndpointScanProxiesSortByEnum, 0)
	for _, v := range mappingListOdaPrivateEndpointScanProxiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOdaPrivateEndpointScanProxiesSortByEnumStringValues Enumerates the set of values in String for ListOdaPrivateEndpointScanProxiesSortByEnum
func GetListOdaPrivateEndpointScanProxiesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListOdaPrivateEndpointScanProxiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOdaPrivateEndpointScanProxiesSortByEnum(val string) (ListOdaPrivateEndpointScanProxiesSortByEnum, bool) {
	enum, ok := mappingListOdaPrivateEndpointScanProxiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
