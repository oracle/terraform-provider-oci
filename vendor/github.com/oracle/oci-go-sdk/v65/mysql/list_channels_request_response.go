// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListChannelsRequest wrapper for the ListChannels operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/ListChannels.go.html to see an example of how to use ListChannelsRequest.
type ListChannelsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The DB System OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `mandatory:"false" contributesTo:"query" name:"dbSystemId"`

	// The OCID of the Channel.
	ChannelId *string `mandatory:"false" contributesTo:"query" name:"channelId"`

	// A filter to return only the resource matching the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The LifecycleState of the Channel.
	LifecycleState ChannelLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// If true, returns only Channels that are enabled. If false, returns only
	// Channels that are disabled.
	IsEnabled *bool `mandatory:"false" contributesTo:"query" name:"isEnabled"`

	// The field to sort by. Only one sort order may be provided. Time fields are default ordered as descending. Display name is default ordered as ascending.
	SortBy ListChannelsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use (ASC or DESC).
	SortOrder ListChannelsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated list call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from
	// the previous list call. For information about pagination, see List
	// Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListChannelsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListChannelsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListChannelsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListChannelsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListChannelsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingChannelLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetChannelLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListChannelsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListChannelsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListChannelsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListChannelsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListChannelsResponse wrapper for the ListChannels operation
type ListChannelsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ChannelSummary instances
	Items []ChannelSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListChannelsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListChannelsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListChannelsSortByEnum Enum with underlying type: string
type ListChannelsSortByEnum string

// Set of constants representing the allowable values for ListChannelsSortByEnum
const (
	ListChannelsSortByDisplayname ListChannelsSortByEnum = "displayName"
	ListChannelsSortByTimecreated ListChannelsSortByEnum = "timeCreated"
)

var mappingListChannelsSortByEnum = map[string]ListChannelsSortByEnum{
	"displayName": ListChannelsSortByDisplayname,
	"timeCreated": ListChannelsSortByTimecreated,
}

var mappingListChannelsSortByEnumLowerCase = map[string]ListChannelsSortByEnum{
	"displayname": ListChannelsSortByDisplayname,
	"timecreated": ListChannelsSortByTimecreated,
}

// GetListChannelsSortByEnumValues Enumerates the set of values for ListChannelsSortByEnum
func GetListChannelsSortByEnumValues() []ListChannelsSortByEnum {
	values := make([]ListChannelsSortByEnum, 0)
	for _, v := range mappingListChannelsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListChannelsSortByEnumStringValues Enumerates the set of values in String for ListChannelsSortByEnum
func GetListChannelsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListChannelsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChannelsSortByEnum(val string) (ListChannelsSortByEnum, bool) {
	enum, ok := mappingListChannelsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListChannelsSortOrderEnum Enum with underlying type: string
type ListChannelsSortOrderEnum string

// Set of constants representing the allowable values for ListChannelsSortOrderEnum
const (
	ListChannelsSortOrderAsc  ListChannelsSortOrderEnum = "ASC"
	ListChannelsSortOrderDesc ListChannelsSortOrderEnum = "DESC"
)

var mappingListChannelsSortOrderEnum = map[string]ListChannelsSortOrderEnum{
	"ASC":  ListChannelsSortOrderAsc,
	"DESC": ListChannelsSortOrderDesc,
}

var mappingListChannelsSortOrderEnumLowerCase = map[string]ListChannelsSortOrderEnum{
	"asc":  ListChannelsSortOrderAsc,
	"desc": ListChannelsSortOrderDesc,
}

// GetListChannelsSortOrderEnumValues Enumerates the set of values for ListChannelsSortOrderEnum
func GetListChannelsSortOrderEnumValues() []ListChannelsSortOrderEnum {
	values := make([]ListChannelsSortOrderEnum, 0)
	for _, v := range mappingListChannelsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListChannelsSortOrderEnumStringValues Enumerates the set of values in String for ListChannelsSortOrderEnum
func GetListChannelsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListChannelsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChannelsSortOrderEnum(val string) (ListChannelsSortOrderEnum, bool) {
	enum, ok := mappingListChannelsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
