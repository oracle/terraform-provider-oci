// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package announcementsservice

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAnnouncementSubscriptionsRequest wrapper for the ListAnnouncementSubscriptions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/ListAnnouncementSubscriptions.go.html to see an example of how to use ListAnnouncementSubscriptionsRequest.
type ListAnnouncementSubscriptionsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only announcement subscriptions that match the given lifecycle state.
	LifecycleState AnnouncementSubscriptionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID of the announcement subscription.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, whether ascending ('ASC') or descending ('DESC').
	SortOrder ListAnnouncementSubscriptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The criteria to sort by. You can specify only one sort order. The default sort order for the creation date of resources is descending. The default sort order for display names is ascending.
	SortBy ListAnnouncementSubscriptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the complete request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAnnouncementSubscriptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAnnouncementSubscriptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAnnouncementSubscriptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAnnouncementSubscriptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAnnouncementSubscriptionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAnnouncementSubscriptionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAnnouncementSubscriptionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAnnouncementSubscriptionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAnnouncementSubscriptionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAnnouncementSubscriptionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAnnouncementSubscriptionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAnnouncementSubscriptionsResponse wrapper for the ListAnnouncementSubscriptions operation
type ListAnnouncementSubscriptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AnnouncementSubscriptionCollection instances
	AnnouncementSubscriptionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAnnouncementSubscriptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAnnouncementSubscriptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAnnouncementSubscriptionsSortOrderEnum Enum with underlying type: string
type ListAnnouncementSubscriptionsSortOrderEnum string

// Set of constants representing the allowable values for ListAnnouncementSubscriptionsSortOrderEnum
const (
	ListAnnouncementSubscriptionsSortOrderAsc  ListAnnouncementSubscriptionsSortOrderEnum = "ASC"
	ListAnnouncementSubscriptionsSortOrderDesc ListAnnouncementSubscriptionsSortOrderEnum = "DESC"
)

var mappingListAnnouncementSubscriptionsSortOrderEnum = map[string]ListAnnouncementSubscriptionsSortOrderEnum{
	"ASC":  ListAnnouncementSubscriptionsSortOrderAsc,
	"DESC": ListAnnouncementSubscriptionsSortOrderDesc,
}

var mappingListAnnouncementSubscriptionsSortOrderEnumLowerCase = map[string]ListAnnouncementSubscriptionsSortOrderEnum{
	"asc":  ListAnnouncementSubscriptionsSortOrderAsc,
	"desc": ListAnnouncementSubscriptionsSortOrderDesc,
}

// GetListAnnouncementSubscriptionsSortOrderEnumValues Enumerates the set of values for ListAnnouncementSubscriptionsSortOrderEnum
func GetListAnnouncementSubscriptionsSortOrderEnumValues() []ListAnnouncementSubscriptionsSortOrderEnum {
	values := make([]ListAnnouncementSubscriptionsSortOrderEnum, 0)
	for _, v := range mappingListAnnouncementSubscriptionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnnouncementSubscriptionsSortOrderEnumStringValues Enumerates the set of values in String for ListAnnouncementSubscriptionsSortOrderEnum
func GetListAnnouncementSubscriptionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAnnouncementSubscriptionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnnouncementSubscriptionsSortOrderEnum(val string) (ListAnnouncementSubscriptionsSortOrderEnum, bool) {
	enum, ok := mappingListAnnouncementSubscriptionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAnnouncementSubscriptionsSortByEnum Enum with underlying type: string
type ListAnnouncementSubscriptionsSortByEnum string

// Set of constants representing the allowable values for ListAnnouncementSubscriptionsSortByEnum
const (
	ListAnnouncementSubscriptionsSortByTimecreated ListAnnouncementSubscriptionsSortByEnum = "timeCreated"
	ListAnnouncementSubscriptionsSortByDisplayname ListAnnouncementSubscriptionsSortByEnum = "displayName"
)

var mappingListAnnouncementSubscriptionsSortByEnum = map[string]ListAnnouncementSubscriptionsSortByEnum{
	"timeCreated": ListAnnouncementSubscriptionsSortByTimecreated,
	"displayName": ListAnnouncementSubscriptionsSortByDisplayname,
}

var mappingListAnnouncementSubscriptionsSortByEnumLowerCase = map[string]ListAnnouncementSubscriptionsSortByEnum{
	"timecreated": ListAnnouncementSubscriptionsSortByTimecreated,
	"displayname": ListAnnouncementSubscriptionsSortByDisplayname,
}

// GetListAnnouncementSubscriptionsSortByEnumValues Enumerates the set of values for ListAnnouncementSubscriptionsSortByEnum
func GetListAnnouncementSubscriptionsSortByEnumValues() []ListAnnouncementSubscriptionsSortByEnum {
	values := make([]ListAnnouncementSubscriptionsSortByEnum, 0)
	for _, v := range mappingListAnnouncementSubscriptionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnnouncementSubscriptionsSortByEnumStringValues Enumerates the set of values in String for ListAnnouncementSubscriptionsSortByEnum
func GetListAnnouncementSubscriptionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAnnouncementSubscriptionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnnouncementSubscriptionsSortByEnum(val string) (ListAnnouncementSubscriptionsSortByEnum, bool) {
	enum, ok := mappingListAnnouncementSubscriptionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
