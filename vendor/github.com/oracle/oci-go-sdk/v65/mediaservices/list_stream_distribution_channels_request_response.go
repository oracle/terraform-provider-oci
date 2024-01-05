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

// ListStreamDistributionChannelsRequest wrapper for the ListStreamDistributionChannels operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListStreamDistributionChannels.go.html to see an example of how to use ListStreamDistributionChannelsRequest.
type ListStreamDistributionChannelsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique Stream Distribution Channel identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only the resources with lifecycleState matching the given lifecycleState.
	LifecycleState StreamDistributionChannelLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListStreamDistributionChannelsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default
	// order for displayName is ascending.
	SortBy ListStreamDistributionChannelsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

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

func (request ListStreamDistributionChannelsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStreamDistributionChannelsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListStreamDistributionChannelsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStreamDistributionChannelsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListStreamDistributionChannelsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStreamDistributionChannelLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetStreamDistributionChannelLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStreamDistributionChannelsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListStreamDistributionChannelsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStreamDistributionChannelsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListStreamDistributionChannelsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListStreamDistributionChannelsResponse wrapper for the ListStreamDistributionChannels operation
type ListStreamDistributionChannelsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of StreamDistributionChannelCollection instances
	StreamDistributionChannelCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListStreamDistributionChannelsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStreamDistributionChannelsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStreamDistributionChannelsSortOrderEnum Enum with underlying type: string
type ListStreamDistributionChannelsSortOrderEnum string

// Set of constants representing the allowable values for ListStreamDistributionChannelsSortOrderEnum
const (
	ListStreamDistributionChannelsSortOrderAsc  ListStreamDistributionChannelsSortOrderEnum = "ASC"
	ListStreamDistributionChannelsSortOrderDesc ListStreamDistributionChannelsSortOrderEnum = "DESC"
)

var mappingListStreamDistributionChannelsSortOrderEnum = map[string]ListStreamDistributionChannelsSortOrderEnum{
	"ASC":  ListStreamDistributionChannelsSortOrderAsc,
	"DESC": ListStreamDistributionChannelsSortOrderDesc,
}

var mappingListStreamDistributionChannelsSortOrderEnumLowerCase = map[string]ListStreamDistributionChannelsSortOrderEnum{
	"asc":  ListStreamDistributionChannelsSortOrderAsc,
	"desc": ListStreamDistributionChannelsSortOrderDesc,
}

// GetListStreamDistributionChannelsSortOrderEnumValues Enumerates the set of values for ListStreamDistributionChannelsSortOrderEnum
func GetListStreamDistributionChannelsSortOrderEnumValues() []ListStreamDistributionChannelsSortOrderEnum {
	values := make([]ListStreamDistributionChannelsSortOrderEnum, 0)
	for _, v := range mappingListStreamDistributionChannelsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListStreamDistributionChannelsSortOrderEnumStringValues Enumerates the set of values in String for ListStreamDistributionChannelsSortOrderEnum
func GetListStreamDistributionChannelsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListStreamDistributionChannelsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStreamDistributionChannelsSortOrderEnum(val string) (ListStreamDistributionChannelsSortOrderEnum, bool) {
	enum, ok := mappingListStreamDistributionChannelsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStreamDistributionChannelsSortByEnum Enum with underlying type: string
type ListStreamDistributionChannelsSortByEnum string

// Set of constants representing the allowable values for ListStreamDistributionChannelsSortByEnum
const (
	ListStreamDistributionChannelsSortByTimecreated ListStreamDistributionChannelsSortByEnum = "timeCreated"
	ListStreamDistributionChannelsSortByDisplayname ListStreamDistributionChannelsSortByEnum = "displayName"
)

var mappingListStreamDistributionChannelsSortByEnum = map[string]ListStreamDistributionChannelsSortByEnum{
	"timeCreated": ListStreamDistributionChannelsSortByTimecreated,
	"displayName": ListStreamDistributionChannelsSortByDisplayname,
}

var mappingListStreamDistributionChannelsSortByEnumLowerCase = map[string]ListStreamDistributionChannelsSortByEnum{
	"timecreated": ListStreamDistributionChannelsSortByTimecreated,
	"displayname": ListStreamDistributionChannelsSortByDisplayname,
}

// GetListStreamDistributionChannelsSortByEnumValues Enumerates the set of values for ListStreamDistributionChannelsSortByEnum
func GetListStreamDistributionChannelsSortByEnumValues() []ListStreamDistributionChannelsSortByEnum {
	values := make([]ListStreamDistributionChannelsSortByEnum, 0)
	for _, v := range mappingListStreamDistributionChannelsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListStreamDistributionChannelsSortByEnumStringValues Enumerates the set of values in String for ListStreamDistributionChannelsSortByEnum
func GetListStreamDistributionChannelsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListStreamDistributionChannelsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStreamDistributionChannelsSortByEnum(val string) (ListStreamDistributionChannelsSortByEnum, bool) {
	enum, ok := mappingListStreamDistributionChannelsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
