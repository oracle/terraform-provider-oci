// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListStreamPackagingConfigsRequest wrapper for the ListStreamPackagingConfigs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListStreamPackagingConfigs.go.html to see an example of how to use ListStreamPackagingConfigsRequest.
type ListStreamPackagingConfigsRequest struct {

	// Unique Stream Distribution Channel identifier.
	DistributionChannelId *string `mandatory:"true" contributesTo:"query" name:"distributionChannelId"`

	// Unique Stream Packaging Configuration identifier.
	StreamPackagingConfigId *string `mandatory:"false" contributesTo:"query" name:"streamPackagingConfigId"`

	// A filter to return only the resources with lifecycleState matching the given lifecycleState.
	LifecycleState StreamPackagingConfigLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListStreamPackagingConfigsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default
	// order for displayName is ascending.
	SortBy ListStreamPackagingConfigsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A token representing the position at which to start retrieving results. This must come from the
	// `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStreamPackagingConfigsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStreamPackagingConfigsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListStreamPackagingConfigsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStreamPackagingConfigsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListStreamPackagingConfigsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStreamPackagingConfigLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetStreamPackagingConfigLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStreamPackagingConfigsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListStreamPackagingConfigsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStreamPackagingConfigsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListStreamPackagingConfigsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListStreamPackagingConfigsResponse wrapper for the ListStreamPackagingConfigs operation
type ListStreamPackagingConfigsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of StreamPackagingConfigCollection instances
	StreamPackagingConfigCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListStreamPackagingConfigsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStreamPackagingConfigsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStreamPackagingConfigsSortOrderEnum Enum with underlying type: string
type ListStreamPackagingConfigsSortOrderEnum string

// Set of constants representing the allowable values for ListStreamPackagingConfigsSortOrderEnum
const (
	ListStreamPackagingConfigsSortOrderAsc  ListStreamPackagingConfigsSortOrderEnum = "ASC"
	ListStreamPackagingConfigsSortOrderDesc ListStreamPackagingConfigsSortOrderEnum = "DESC"
)

var mappingListStreamPackagingConfigsSortOrderEnum = map[string]ListStreamPackagingConfigsSortOrderEnum{
	"ASC":  ListStreamPackagingConfigsSortOrderAsc,
	"DESC": ListStreamPackagingConfigsSortOrderDesc,
}

var mappingListStreamPackagingConfigsSortOrderEnumLowerCase = map[string]ListStreamPackagingConfigsSortOrderEnum{
	"asc":  ListStreamPackagingConfigsSortOrderAsc,
	"desc": ListStreamPackagingConfigsSortOrderDesc,
}

// GetListStreamPackagingConfigsSortOrderEnumValues Enumerates the set of values for ListStreamPackagingConfigsSortOrderEnum
func GetListStreamPackagingConfigsSortOrderEnumValues() []ListStreamPackagingConfigsSortOrderEnum {
	values := make([]ListStreamPackagingConfigsSortOrderEnum, 0)
	for _, v := range mappingListStreamPackagingConfigsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListStreamPackagingConfigsSortOrderEnumStringValues Enumerates the set of values in String for ListStreamPackagingConfigsSortOrderEnum
func GetListStreamPackagingConfigsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListStreamPackagingConfigsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStreamPackagingConfigsSortOrderEnum(val string) (ListStreamPackagingConfigsSortOrderEnum, bool) {
	enum, ok := mappingListStreamPackagingConfigsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStreamPackagingConfigsSortByEnum Enum with underlying type: string
type ListStreamPackagingConfigsSortByEnum string

// Set of constants representing the allowable values for ListStreamPackagingConfigsSortByEnum
const (
	ListStreamPackagingConfigsSortByTimecreated ListStreamPackagingConfigsSortByEnum = "timeCreated"
	ListStreamPackagingConfigsSortByDisplayname ListStreamPackagingConfigsSortByEnum = "displayName"
)

var mappingListStreamPackagingConfigsSortByEnum = map[string]ListStreamPackagingConfigsSortByEnum{
	"timeCreated": ListStreamPackagingConfigsSortByTimecreated,
	"displayName": ListStreamPackagingConfigsSortByDisplayname,
}

var mappingListStreamPackagingConfigsSortByEnumLowerCase = map[string]ListStreamPackagingConfigsSortByEnum{
	"timecreated": ListStreamPackagingConfigsSortByTimecreated,
	"displayname": ListStreamPackagingConfigsSortByDisplayname,
}

// GetListStreamPackagingConfigsSortByEnumValues Enumerates the set of values for ListStreamPackagingConfigsSortByEnum
func GetListStreamPackagingConfigsSortByEnumValues() []ListStreamPackagingConfigsSortByEnum {
	values := make([]ListStreamPackagingConfigsSortByEnum, 0)
	for _, v := range mappingListStreamPackagingConfigsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListStreamPackagingConfigsSortByEnumStringValues Enumerates the set of values in String for ListStreamPackagingConfigsSortByEnum
func GetListStreamPackagingConfigsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListStreamPackagingConfigsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStreamPackagingConfigsSortByEnum(val string) (ListStreamPackagingConfigsSortByEnum, bool) {
	enum, ok := mappingListStreamPackagingConfigsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
