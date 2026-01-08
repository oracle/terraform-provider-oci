// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMarketingBrandsRequest wrapper for the ListMarketingBrands operation
type ListMarketingBrandsRequest struct {

	// unique FusionEnvironment identifier
	FusionEnvironmentId *string `mandatory:"true" contributesTo:"path" name:"fusionEnvironmentId"`

	// unique brand identifier
	MarketingBrandId *string `mandatory:"false" contributesTo:"query" name:"marketingBrandId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter that returns all resources that match the specified status
	LifecycleState MarketingBrandLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListMarketingBrandsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending. If no value is specified timeCreated is default.
	SortBy ListMarketingBrandsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMarketingBrandsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMarketingBrandsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMarketingBrandsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMarketingBrandsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMarketingBrandsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMarketingBrandLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMarketingBrandLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMarketingBrandsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMarketingBrandsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMarketingBrandsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMarketingBrandsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMarketingBrandsResponse wrapper for the ListMarketingBrands operation
type ListMarketingBrandsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MarketingBrandCollection instances
	MarketingBrandCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMarketingBrandsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMarketingBrandsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMarketingBrandsSortOrderEnum Enum with underlying type: string
type ListMarketingBrandsSortOrderEnum string

// Set of constants representing the allowable values for ListMarketingBrandsSortOrderEnum
const (
	ListMarketingBrandsSortOrderAsc  ListMarketingBrandsSortOrderEnum = "ASC"
	ListMarketingBrandsSortOrderDesc ListMarketingBrandsSortOrderEnum = "DESC"
)

var mappingListMarketingBrandsSortOrderEnum = map[string]ListMarketingBrandsSortOrderEnum{
	"ASC":  ListMarketingBrandsSortOrderAsc,
	"DESC": ListMarketingBrandsSortOrderDesc,
}

var mappingListMarketingBrandsSortOrderEnumLowerCase = map[string]ListMarketingBrandsSortOrderEnum{
	"asc":  ListMarketingBrandsSortOrderAsc,
	"desc": ListMarketingBrandsSortOrderDesc,
}

// GetListMarketingBrandsSortOrderEnumValues Enumerates the set of values for ListMarketingBrandsSortOrderEnum
func GetListMarketingBrandsSortOrderEnumValues() []ListMarketingBrandsSortOrderEnum {
	values := make([]ListMarketingBrandsSortOrderEnum, 0)
	for _, v := range mappingListMarketingBrandsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMarketingBrandsSortOrderEnumStringValues Enumerates the set of values in String for ListMarketingBrandsSortOrderEnum
func GetListMarketingBrandsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMarketingBrandsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMarketingBrandsSortOrderEnum(val string) (ListMarketingBrandsSortOrderEnum, bool) {
	enum, ok := mappingListMarketingBrandsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMarketingBrandsSortByEnum Enum with underlying type: string
type ListMarketingBrandsSortByEnum string

// Set of constants representing the allowable values for ListMarketingBrandsSortByEnum
const (
	ListMarketingBrandsSortByTimeCreated ListMarketingBrandsSortByEnum = "TIME_CREATED"
	ListMarketingBrandsSortByName        ListMarketingBrandsSortByEnum = "NAME"
	ListMarketingBrandsSortById          ListMarketingBrandsSortByEnum = "ID"
)

var mappingListMarketingBrandsSortByEnum = map[string]ListMarketingBrandsSortByEnum{
	"TIME_CREATED": ListMarketingBrandsSortByTimeCreated,
	"NAME":         ListMarketingBrandsSortByName,
	"ID":           ListMarketingBrandsSortById,
}

var mappingListMarketingBrandsSortByEnumLowerCase = map[string]ListMarketingBrandsSortByEnum{
	"time_created": ListMarketingBrandsSortByTimeCreated,
	"name":         ListMarketingBrandsSortByName,
	"id":           ListMarketingBrandsSortById,
}

// GetListMarketingBrandsSortByEnumValues Enumerates the set of values for ListMarketingBrandsSortByEnum
func GetListMarketingBrandsSortByEnumValues() []ListMarketingBrandsSortByEnum {
	values := make([]ListMarketingBrandsSortByEnum, 0)
	for _, v := range mappingListMarketingBrandsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMarketingBrandsSortByEnumStringValues Enumerates the set of values in String for ListMarketingBrandsSortByEnum
func GetListMarketingBrandsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"NAME",
		"ID",
	}
}

// GetMappingListMarketingBrandsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMarketingBrandsSortByEnum(val string) (ListMarketingBrandsSortByEnum, bool) {
	enum, ok := mappingListMarketingBrandsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
