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

// ListEmailSubdomainsRequest wrapper for the ListEmailSubdomains operation
type ListEmailSubdomainsRequest struct {

	// unique FusionEnvironment identifier
	FusionEnvironmentId *string `mandatory:"true" contributesTo:"path" name:"fusionEnvironmentId"`

	// unique brand identifier
	MarketingBrandId *string `mandatory:"true" contributesTo:"path" name:"marketingBrandId"`

	// unique emailSubdomain identifier
	EmailSubdomainId *string `mandatory:"false" contributesTo:"query" name:"emailSubdomainId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter that returns all resources that match the specified status
	LifecycleState EmailSubdomainLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListEmailSubdomainsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending. If no value is specified timeCreated is default.
	SortBy ListEmailSubdomainsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

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

func (request ListEmailSubdomainsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEmailSubdomainsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEmailSubdomainsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEmailSubdomainsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEmailSubdomainsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmailSubdomainLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetEmailSubdomainLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEmailSubdomainsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEmailSubdomainsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEmailSubdomainsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEmailSubdomainsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEmailSubdomainsResponse wrapper for the ListEmailSubdomains operation
type ListEmailSubdomainsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EmailSubdomainCollection instances
	EmailSubdomainCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEmailSubdomainsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEmailSubdomainsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEmailSubdomainsSortOrderEnum Enum with underlying type: string
type ListEmailSubdomainsSortOrderEnum string

// Set of constants representing the allowable values for ListEmailSubdomainsSortOrderEnum
const (
	ListEmailSubdomainsSortOrderAsc  ListEmailSubdomainsSortOrderEnum = "ASC"
	ListEmailSubdomainsSortOrderDesc ListEmailSubdomainsSortOrderEnum = "DESC"
)

var mappingListEmailSubdomainsSortOrderEnum = map[string]ListEmailSubdomainsSortOrderEnum{
	"ASC":  ListEmailSubdomainsSortOrderAsc,
	"DESC": ListEmailSubdomainsSortOrderDesc,
}

var mappingListEmailSubdomainsSortOrderEnumLowerCase = map[string]ListEmailSubdomainsSortOrderEnum{
	"asc":  ListEmailSubdomainsSortOrderAsc,
	"desc": ListEmailSubdomainsSortOrderDesc,
}

// GetListEmailSubdomainsSortOrderEnumValues Enumerates the set of values for ListEmailSubdomainsSortOrderEnum
func GetListEmailSubdomainsSortOrderEnumValues() []ListEmailSubdomainsSortOrderEnum {
	values := make([]ListEmailSubdomainsSortOrderEnum, 0)
	for _, v := range mappingListEmailSubdomainsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailSubdomainsSortOrderEnumStringValues Enumerates the set of values in String for ListEmailSubdomainsSortOrderEnum
func GetListEmailSubdomainsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEmailSubdomainsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailSubdomainsSortOrderEnum(val string) (ListEmailSubdomainsSortOrderEnum, bool) {
	enum, ok := mappingListEmailSubdomainsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEmailSubdomainsSortByEnum Enum with underlying type: string
type ListEmailSubdomainsSortByEnum string

// Set of constants representing the allowable values for ListEmailSubdomainsSortByEnum
const (
	ListEmailSubdomainsSortByTimeCreated ListEmailSubdomainsSortByEnum = "TIME_CREATED"
	ListEmailSubdomainsSortByName        ListEmailSubdomainsSortByEnum = "NAME"
	ListEmailSubdomainsSortById          ListEmailSubdomainsSortByEnum = "ID"
)

var mappingListEmailSubdomainsSortByEnum = map[string]ListEmailSubdomainsSortByEnum{
	"TIME_CREATED": ListEmailSubdomainsSortByTimeCreated,
	"NAME":         ListEmailSubdomainsSortByName,
	"ID":           ListEmailSubdomainsSortById,
}

var mappingListEmailSubdomainsSortByEnumLowerCase = map[string]ListEmailSubdomainsSortByEnum{
	"time_created": ListEmailSubdomainsSortByTimeCreated,
	"name":         ListEmailSubdomainsSortByName,
	"id":           ListEmailSubdomainsSortById,
}

// GetListEmailSubdomainsSortByEnumValues Enumerates the set of values for ListEmailSubdomainsSortByEnum
func GetListEmailSubdomainsSortByEnumValues() []ListEmailSubdomainsSortByEnum {
	values := make([]ListEmailSubdomainsSortByEnum, 0)
	for _, v := range mappingListEmailSubdomainsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailSubdomainsSortByEnumStringValues Enumerates the set of values in String for ListEmailSubdomainsSortByEnum
func GetListEmailSubdomainsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"NAME",
		"ID",
	}
}

// GetMappingListEmailSubdomainsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailSubdomainsSortByEnum(val string) (ListEmailSubdomainsSortByEnum, bool) {
	enum, ok := mappingListEmailSubdomainsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
