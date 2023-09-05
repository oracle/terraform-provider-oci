// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package queue

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListConsumerGroupsRequest wrapper for the ListConsumerGroups operation
type ListConsumerGroupsRequest struct {

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ConsumerGroupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The unique consumer group identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The unique queue identifier.
	QueueId *string `mandatory:"false" contributesTo:"query" name:"queueId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListConsumerGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListConsumerGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConsumerGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConsumerGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConsumerGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConsumerGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListConsumerGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConsumerGroupLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetConsumerGroupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConsumerGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConsumerGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConsumerGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConsumerGroupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListConsumerGroupsResponse wrapper for the ListConsumerGroups operation
type ListConsumerGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConsumerGroupCollection instances
	ConsumerGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListConsumerGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConsumerGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConsumerGroupsSortOrderEnum Enum with underlying type: string
type ListConsumerGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListConsumerGroupsSortOrderEnum
const (
	ListConsumerGroupsSortOrderAsc  ListConsumerGroupsSortOrderEnum = "ASC"
	ListConsumerGroupsSortOrderDesc ListConsumerGroupsSortOrderEnum = "DESC"
)

var mappingListConsumerGroupsSortOrderEnum = map[string]ListConsumerGroupsSortOrderEnum{
	"ASC":  ListConsumerGroupsSortOrderAsc,
	"DESC": ListConsumerGroupsSortOrderDesc,
}

var mappingListConsumerGroupsSortOrderEnumLowerCase = map[string]ListConsumerGroupsSortOrderEnum{
	"asc":  ListConsumerGroupsSortOrderAsc,
	"desc": ListConsumerGroupsSortOrderDesc,
}

// GetListConsumerGroupsSortOrderEnumValues Enumerates the set of values for ListConsumerGroupsSortOrderEnum
func GetListConsumerGroupsSortOrderEnumValues() []ListConsumerGroupsSortOrderEnum {
	values := make([]ListConsumerGroupsSortOrderEnum, 0)
	for _, v := range mappingListConsumerGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListConsumerGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListConsumerGroupsSortOrderEnum
func GetListConsumerGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListConsumerGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConsumerGroupsSortOrderEnum(val string) (ListConsumerGroupsSortOrderEnum, bool) {
	enum, ok := mappingListConsumerGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConsumerGroupsSortByEnum Enum with underlying type: string
type ListConsumerGroupsSortByEnum string

// Set of constants representing the allowable values for ListConsumerGroupsSortByEnum
const (
	ListConsumerGroupsSortByTimecreated ListConsumerGroupsSortByEnum = "timeCreated"
	ListConsumerGroupsSortByDisplayname ListConsumerGroupsSortByEnum = "displayName"
)

var mappingListConsumerGroupsSortByEnum = map[string]ListConsumerGroupsSortByEnum{
	"timeCreated": ListConsumerGroupsSortByTimecreated,
	"displayName": ListConsumerGroupsSortByDisplayname,
}

var mappingListConsumerGroupsSortByEnumLowerCase = map[string]ListConsumerGroupsSortByEnum{
	"timecreated": ListConsumerGroupsSortByTimecreated,
	"displayname": ListConsumerGroupsSortByDisplayname,
}

// GetListConsumerGroupsSortByEnumValues Enumerates the set of values for ListConsumerGroupsSortByEnum
func GetListConsumerGroupsSortByEnumValues() []ListConsumerGroupsSortByEnum {
	values := make([]ListConsumerGroupsSortByEnum, 0)
	for _, v := range mappingListConsumerGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListConsumerGroupsSortByEnumStringValues Enumerates the set of values in String for ListConsumerGroupsSortByEnum
func GetListConsumerGroupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListConsumerGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConsumerGroupsSortByEnum(val string) (ListConsumerGroupsSortByEnum, bool) {
	enum, ok := mappingListConsumerGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
