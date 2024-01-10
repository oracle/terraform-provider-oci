// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package queue

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListQueuesRequest wrapper for the ListQueues operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/queue/ListQueues.go.html to see an example of how to use ListQueuesRequest.
type ListQueuesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState QueueLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The unique queue identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListQueuesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListQueuesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListQueuesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListQueuesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListQueuesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListQueuesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListQueuesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQueueLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetQueueLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListQueuesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListQueuesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListQueuesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListQueuesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListQueuesResponse wrapper for the ListQueues operation
type ListQueuesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of QueueCollection instances
	QueueCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListQueuesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListQueuesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListQueuesSortOrderEnum Enum with underlying type: string
type ListQueuesSortOrderEnum string

// Set of constants representing the allowable values for ListQueuesSortOrderEnum
const (
	ListQueuesSortOrderAsc  ListQueuesSortOrderEnum = "ASC"
	ListQueuesSortOrderDesc ListQueuesSortOrderEnum = "DESC"
)

var mappingListQueuesSortOrderEnum = map[string]ListQueuesSortOrderEnum{
	"ASC":  ListQueuesSortOrderAsc,
	"DESC": ListQueuesSortOrderDesc,
}

var mappingListQueuesSortOrderEnumLowerCase = map[string]ListQueuesSortOrderEnum{
	"asc":  ListQueuesSortOrderAsc,
	"desc": ListQueuesSortOrderDesc,
}

// GetListQueuesSortOrderEnumValues Enumerates the set of values for ListQueuesSortOrderEnum
func GetListQueuesSortOrderEnumValues() []ListQueuesSortOrderEnum {
	values := make([]ListQueuesSortOrderEnum, 0)
	for _, v := range mappingListQueuesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListQueuesSortOrderEnumStringValues Enumerates the set of values in String for ListQueuesSortOrderEnum
func GetListQueuesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListQueuesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListQueuesSortOrderEnum(val string) (ListQueuesSortOrderEnum, bool) {
	enum, ok := mappingListQueuesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListQueuesSortByEnum Enum with underlying type: string
type ListQueuesSortByEnum string

// Set of constants representing the allowable values for ListQueuesSortByEnum
const (
	ListQueuesSortByTimecreated ListQueuesSortByEnum = "timeCreated"
	ListQueuesSortByDisplayname ListQueuesSortByEnum = "displayName"
)

var mappingListQueuesSortByEnum = map[string]ListQueuesSortByEnum{
	"timeCreated": ListQueuesSortByTimecreated,
	"displayName": ListQueuesSortByDisplayname,
}

var mappingListQueuesSortByEnumLowerCase = map[string]ListQueuesSortByEnum{
	"timecreated": ListQueuesSortByTimecreated,
	"displayname": ListQueuesSortByDisplayname,
}

// GetListQueuesSortByEnumValues Enumerates the set of values for ListQueuesSortByEnum
func GetListQueuesSortByEnumValues() []ListQueuesSortByEnum {
	values := make([]ListQueuesSortByEnum, 0)
	for _, v := range mappingListQueuesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListQueuesSortByEnumStringValues Enumerates the set of values in String for ListQueuesSortByEnum
func GetListQueuesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListQueuesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListQueuesSortByEnum(val string) (ListQueuesSortByEnum, bool) {
	enum, ok := mappingListQueuesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
