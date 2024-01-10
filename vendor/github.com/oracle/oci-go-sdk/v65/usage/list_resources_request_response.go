// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package usage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListResourcesRequest wrapper for the ListResources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usage/ListResources.go.html to see an example of how to use ListResourcesRequest.
type ListResourcesRequest struct {

	// Service Name.
	ServiceName *string `mandatory:"true" contributesTo:"query" name:"serviceName"`

	// The OCID of the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Subscription or entitlement Id.
	EntitlementId *string `mandatory:"false" contributesTo:"query" name:"entitlementId"`

	// Unique, Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The value of the 'opc-next-page' response header from the previous call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, which can be ascending (ASC) or descending (DESC).
	SortOrder ListResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Supports one sort order.
	SortBy ListResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResourcesResponse wrapper for the ListResources operation
type ListResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourcesCollection instances
	ResourcesCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourcesSortOrderEnum Enum with underlying type: string
type ListResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListResourcesSortOrderEnum
const (
	ListResourcesSortOrderAsc  ListResourcesSortOrderEnum = "ASC"
	ListResourcesSortOrderDesc ListResourcesSortOrderEnum = "DESC"
)

var mappingListResourcesSortOrderEnum = map[string]ListResourcesSortOrderEnum{
	"ASC":  ListResourcesSortOrderAsc,
	"DESC": ListResourcesSortOrderDesc,
}

var mappingListResourcesSortOrderEnumLowerCase = map[string]ListResourcesSortOrderEnum{
	"asc":  ListResourcesSortOrderAsc,
	"desc": ListResourcesSortOrderDesc,
}

// GetListResourcesSortOrderEnumValues Enumerates the set of values for ListResourcesSortOrderEnum
func GetListResourcesSortOrderEnumValues() []ListResourcesSortOrderEnum {
	values := make([]ListResourcesSortOrderEnum, 0)
	for _, v := range mappingListResourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourcesSortOrderEnumStringValues Enumerates the set of values in String for ListResourcesSortOrderEnum
func GetListResourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourcesSortOrderEnum(val string) (ListResourcesSortOrderEnum, bool) {
	enum, ok := mappingListResourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourcesSortByEnum Enum with underlying type: string
type ListResourcesSortByEnum string

// Set of constants representing the allowable values for ListResourcesSortByEnum
const (
	ListResourcesSortByTimecreated ListResourcesSortByEnum = "TIMECREATED"
	ListResourcesSortByTimestart   ListResourcesSortByEnum = "TIMESTART"
)

var mappingListResourcesSortByEnum = map[string]ListResourcesSortByEnum{
	"TIMECREATED": ListResourcesSortByTimecreated,
	"TIMESTART":   ListResourcesSortByTimestart,
}

var mappingListResourcesSortByEnumLowerCase = map[string]ListResourcesSortByEnum{
	"timecreated": ListResourcesSortByTimecreated,
	"timestart":   ListResourcesSortByTimestart,
}

// GetListResourcesSortByEnumValues Enumerates the set of values for ListResourcesSortByEnum
func GetListResourcesSortByEnumValues() []ListResourcesSortByEnum {
	values := make([]ListResourcesSortByEnum, 0)
	for _, v := range mappingListResourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourcesSortByEnumStringValues Enumerates the set of values in String for ListResourcesSortByEnum
func GetListResourcesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"TIMESTART",
	}
}

// GetMappingListResourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourcesSortByEnum(val string) (ListResourcesSortByEnum, bool) {
	enum, ok := mappingListResourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
