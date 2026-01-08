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

// ListMicrositesRequest wrapper for the ListMicrosites operation
type ListMicrositesRequest struct {

	// unique FusionEnvironment identifier
	FusionEnvironmentId *string `mandatory:"true" contributesTo:"path" name:"fusionEnvironmentId"`

	// unique brand identifier
	MarketingBrandId *string `mandatory:"true" contributesTo:"path" name:"marketingBrandId"`

	// unique microsite identifier
	MicrositeId *string `mandatory:"false" contributesTo:"query" name:"micrositeId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter that returns all resources that match the specified status
	LifecycleState MicrositeLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListMicrositesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending. If no value is specified timeCreated is default.
	SortBy ListMicrositesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMicrositesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMicrositesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMicrositesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMicrositesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMicrositesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMicrositeLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMicrositeLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMicrositesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMicrositesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMicrositesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMicrositesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMicrositesResponse wrapper for the ListMicrosites operation
type ListMicrositesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MicrositeCollection instances
	MicrositeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMicrositesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMicrositesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMicrositesSortOrderEnum Enum with underlying type: string
type ListMicrositesSortOrderEnum string

// Set of constants representing the allowable values for ListMicrositesSortOrderEnum
const (
	ListMicrositesSortOrderAsc  ListMicrositesSortOrderEnum = "ASC"
	ListMicrositesSortOrderDesc ListMicrositesSortOrderEnum = "DESC"
)

var mappingListMicrositesSortOrderEnum = map[string]ListMicrositesSortOrderEnum{
	"ASC":  ListMicrositesSortOrderAsc,
	"DESC": ListMicrositesSortOrderDesc,
}

var mappingListMicrositesSortOrderEnumLowerCase = map[string]ListMicrositesSortOrderEnum{
	"asc":  ListMicrositesSortOrderAsc,
	"desc": ListMicrositesSortOrderDesc,
}

// GetListMicrositesSortOrderEnumValues Enumerates the set of values for ListMicrositesSortOrderEnum
func GetListMicrositesSortOrderEnumValues() []ListMicrositesSortOrderEnum {
	values := make([]ListMicrositesSortOrderEnum, 0)
	for _, v := range mappingListMicrositesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMicrositesSortOrderEnumStringValues Enumerates the set of values in String for ListMicrositesSortOrderEnum
func GetListMicrositesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMicrositesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMicrositesSortOrderEnum(val string) (ListMicrositesSortOrderEnum, bool) {
	enum, ok := mappingListMicrositesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMicrositesSortByEnum Enum with underlying type: string
type ListMicrositesSortByEnum string

// Set of constants representing the allowable values for ListMicrositesSortByEnum
const (
	ListMicrositesSortByTimeCreated ListMicrositesSortByEnum = "TIME_CREATED"
	ListMicrositesSortByName        ListMicrositesSortByEnum = "NAME"
	ListMicrositesSortById          ListMicrositesSortByEnum = "ID"
)

var mappingListMicrositesSortByEnum = map[string]ListMicrositesSortByEnum{
	"TIME_CREATED": ListMicrositesSortByTimeCreated,
	"NAME":         ListMicrositesSortByName,
	"ID":           ListMicrositesSortById,
}

var mappingListMicrositesSortByEnumLowerCase = map[string]ListMicrositesSortByEnum{
	"time_created": ListMicrositesSortByTimeCreated,
	"name":         ListMicrositesSortByName,
	"id":           ListMicrositesSortById,
}

// GetListMicrositesSortByEnumValues Enumerates the set of values for ListMicrositesSortByEnum
func GetListMicrositesSortByEnumValues() []ListMicrositesSortByEnum {
	values := make([]ListMicrositesSortByEnum, 0)
	for _, v := range mappingListMicrositesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMicrositesSortByEnumStringValues Enumerates the set of values in String for ListMicrositesSortByEnum
func GetListMicrositesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"NAME",
		"ID",
	}
}

// GetMappingListMicrositesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMicrositesSortByEnum(val string) (ListMicrositesSortByEnum, bool) {
	enum, ok := mappingListMicrositesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
