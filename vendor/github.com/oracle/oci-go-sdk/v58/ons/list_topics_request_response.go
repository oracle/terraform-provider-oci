// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ons

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListTopicsRequest wrapper for the ListTopics operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ons/ListTopics.go.html to see an example of how to use ListTopicsRequest.
type ListTopicsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to only return resources that match the given id exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one field can be selected for sorting.
	SortBy ListTopicsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use (ascending or descending).
	SortOrder ListTopicsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filter returned list by specified lifecycle state. This parameter is case-insensitive.
	LifecycleState NotificationTopicSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTopicsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTopicsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTopicsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTopicsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTopicsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTopicsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTopicsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTopicsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTopicsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNotificationTopicSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetNotificationTopicSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTopicsResponse wrapper for the ListTopics operation
type ListTopicsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []NotificationTopicSummary instances
	Items []NotificationTopicSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListTopicsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTopicsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTopicsSortByEnum Enum with underlying type: string
type ListTopicsSortByEnum string

// Set of constants representing the allowable values for ListTopicsSortByEnum
const (
	ListTopicsSortByTimecreated    ListTopicsSortByEnum = "TIMECREATED"
	ListTopicsSortByLifecyclestate ListTopicsSortByEnum = "LIFECYCLESTATE"
)

var mappingListTopicsSortByEnum = map[string]ListTopicsSortByEnum{
	"TIMECREATED":    ListTopicsSortByTimecreated,
	"LIFECYCLESTATE": ListTopicsSortByLifecyclestate,
}

// GetListTopicsSortByEnumValues Enumerates the set of values for ListTopicsSortByEnum
func GetListTopicsSortByEnumValues() []ListTopicsSortByEnum {
	values := make([]ListTopicsSortByEnum, 0)
	for _, v := range mappingListTopicsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTopicsSortByEnumStringValues Enumerates the set of values in String for ListTopicsSortByEnum
func GetListTopicsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"LIFECYCLESTATE",
	}
}

// GetMappingListTopicsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTopicsSortByEnum(val string) (ListTopicsSortByEnum, bool) {
	mappingListTopicsSortByEnumIgnoreCase := make(map[string]ListTopicsSortByEnum)
	for k, v := range mappingListTopicsSortByEnum {
		mappingListTopicsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListTopicsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListTopicsSortOrderEnum Enum with underlying type: string
type ListTopicsSortOrderEnum string

// Set of constants representing the allowable values for ListTopicsSortOrderEnum
const (
	ListTopicsSortOrderAsc  ListTopicsSortOrderEnum = "ASC"
	ListTopicsSortOrderDesc ListTopicsSortOrderEnum = "DESC"
)

var mappingListTopicsSortOrderEnum = map[string]ListTopicsSortOrderEnum{
	"ASC":  ListTopicsSortOrderAsc,
	"DESC": ListTopicsSortOrderDesc,
}

// GetListTopicsSortOrderEnumValues Enumerates the set of values for ListTopicsSortOrderEnum
func GetListTopicsSortOrderEnumValues() []ListTopicsSortOrderEnum {
	values := make([]ListTopicsSortOrderEnum, 0)
	for _, v := range mappingListTopicsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTopicsSortOrderEnumStringValues Enumerates the set of values in String for ListTopicsSortOrderEnum
func GetListTopicsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTopicsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTopicsSortOrderEnum(val string) (ListTopicsSortOrderEnum, bool) {
	mappingListTopicsSortOrderEnumIgnoreCase := make(map[string]ListTopicsSortOrderEnum)
	for k, v := range mappingListTopicsSortOrderEnum {
		mappingListTopicsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListTopicsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
