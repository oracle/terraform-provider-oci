// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEdgeSubnetsRequest wrapper for the ListEdgeSubnets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waas/ListEdgeSubnets.go.html to see an example of how to use ListEdgeSubnetsRequest.
type ListEdgeSubnetsRequest struct {

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated call. If unspecified, defaults to `10`.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous paginated call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The value by which edge node subnets are sorted in a paginated 'List' call. If unspecified, defaults to `timeModified`.
	SortBy ListEdgeSubnetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The value of the sorting direction of resources in a paginated 'List' call. If unspecified, defaults to `DESC`.
	SortOrder ListEdgeSubnetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEdgeSubnetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEdgeSubnetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEdgeSubnetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEdgeSubnetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEdgeSubnetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEdgeSubnetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEdgeSubnetsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEdgeSubnetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEdgeSubnetsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEdgeSubnetsResponse wrapper for the ListEdgeSubnets operation
type ListEdgeSubnetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []EdgeSubnet instances
	Items []EdgeSubnet `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent `GET` request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListEdgeSubnetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEdgeSubnetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEdgeSubnetsSortByEnum Enum with underlying type: string
type ListEdgeSubnetsSortByEnum string

// Set of constants representing the allowable values for ListEdgeSubnetsSortByEnum
const (
	ListEdgeSubnetsSortByCidr         ListEdgeSubnetsSortByEnum = "cidr"
	ListEdgeSubnetsSortByRegion       ListEdgeSubnetsSortByEnum = "region"
	ListEdgeSubnetsSortByTimemodified ListEdgeSubnetsSortByEnum = "timeModified"
)

var mappingListEdgeSubnetsSortByEnum = map[string]ListEdgeSubnetsSortByEnum{
	"cidr":         ListEdgeSubnetsSortByCidr,
	"region":       ListEdgeSubnetsSortByRegion,
	"timeModified": ListEdgeSubnetsSortByTimemodified,
}

var mappingListEdgeSubnetsSortByEnumLowerCase = map[string]ListEdgeSubnetsSortByEnum{
	"cidr":         ListEdgeSubnetsSortByCidr,
	"region":       ListEdgeSubnetsSortByRegion,
	"timemodified": ListEdgeSubnetsSortByTimemodified,
}

// GetListEdgeSubnetsSortByEnumValues Enumerates the set of values for ListEdgeSubnetsSortByEnum
func GetListEdgeSubnetsSortByEnumValues() []ListEdgeSubnetsSortByEnum {
	values := make([]ListEdgeSubnetsSortByEnum, 0)
	for _, v := range mappingListEdgeSubnetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEdgeSubnetsSortByEnumStringValues Enumerates the set of values in String for ListEdgeSubnetsSortByEnum
func GetListEdgeSubnetsSortByEnumStringValues() []string {
	return []string{
		"cidr",
		"region",
		"timeModified",
	}
}

// GetMappingListEdgeSubnetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEdgeSubnetsSortByEnum(val string) (ListEdgeSubnetsSortByEnum, bool) {
	enum, ok := mappingListEdgeSubnetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEdgeSubnetsSortOrderEnum Enum with underlying type: string
type ListEdgeSubnetsSortOrderEnum string

// Set of constants representing the allowable values for ListEdgeSubnetsSortOrderEnum
const (
	ListEdgeSubnetsSortOrderAsc  ListEdgeSubnetsSortOrderEnum = "ASC"
	ListEdgeSubnetsSortOrderDesc ListEdgeSubnetsSortOrderEnum = "DESC"
)

var mappingListEdgeSubnetsSortOrderEnum = map[string]ListEdgeSubnetsSortOrderEnum{
	"ASC":  ListEdgeSubnetsSortOrderAsc,
	"DESC": ListEdgeSubnetsSortOrderDesc,
}

var mappingListEdgeSubnetsSortOrderEnumLowerCase = map[string]ListEdgeSubnetsSortOrderEnum{
	"asc":  ListEdgeSubnetsSortOrderAsc,
	"desc": ListEdgeSubnetsSortOrderDesc,
}

// GetListEdgeSubnetsSortOrderEnumValues Enumerates the set of values for ListEdgeSubnetsSortOrderEnum
func GetListEdgeSubnetsSortOrderEnumValues() []ListEdgeSubnetsSortOrderEnum {
	values := make([]ListEdgeSubnetsSortOrderEnum, 0)
	for _, v := range mappingListEdgeSubnetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEdgeSubnetsSortOrderEnumStringValues Enumerates the set of values in String for ListEdgeSubnetsSortOrderEnum
func GetListEdgeSubnetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEdgeSubnetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEdgeSubnetsSortOrderEnum(val string) (ListEdgeSubnetsSortOrderEnum, bool) {
	enum, ok := mappingListEdgeSubnetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
