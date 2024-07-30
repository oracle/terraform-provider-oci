// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetsoftwareupdate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFsuDiscoveriesRequest wrapper for the ListFsuDiscoveries operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuDiscoveries.go.html to see an example of how to use ListFsuDiscoveriesRequest.
type ListFsuDiscoveriesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState ListFsuDiscoveriesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFsuDiscoveriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListFsuDiscoveriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFsuDiscoveriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFsuDiscoveriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFsuDiscoveriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFsuDiscoveriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFsuDiscoveriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFsuDiscoveriesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListFsuDiscoveriesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuDiscoveriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFsuDiscoveriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuDiscoveriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFsuDiscoveriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFsuDiscoveriesResponse wrapper for the ListFsuDiscoveries operation
type ListFsuDiscoveriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FsuDiscoverySummaryCollection instances
	FsuDiscoverySummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFsuDiscoveriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFsuDiscoveriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFsuDiscoveriesLifecycleStateEnum Enum with underlying type: string
type ListFsuDiscoveriesLifecycleStateEnum string

// Set of constants representing the allowable values for ListFsuDiscoveriesLifecycleStateEnum
const (
	ListFsuDiscoveriesLifecycleStateAccepted   ListFsuDiscoveriesLifecycleStateEnum = "ACCEPTED"
	ListFsuDiscoveriesLifecycleStateInProgress ListFsuDiscoveriesLifecycleStateEnum = "IN_PROGRESS"
	ListFsuDiscoveriesLifecycleStateFailed     ListFsuDiscoveriesLifecycleStateEnum = "FAILED"
	ListFsuDiscoveriesLifecycleStateSucceeded  ListFsuDiscoveriesLifecycleStateEnum = "SUCCEEDED"
	ListFsuDiscoveriesLifecycleStateCanceling  ListFsuDiscoveriesLifecycleStateEnum = "CANCELING"
	ListFsuDiscoveriesLifecycleStateCanceled   ListFsuDiscoveriesLifecycleStateEnum = "CANCELED"
	ListFsuDiscoveriesLifecycleStateDeleting   ListFsuDiscoveriesLifecycleStateEnum = "DELETING"
	ListFsuDiscoveriesLifecycleStateDeleted    ListFsuDiscoveriesLifecycleStateEnum = "DELETED"
)

var mappingListFsuDiscoveriesLifecycleStateEnum = map[string]ListFsuDiscoveriesLifecycleStateEnum{
	"ACCEPTED":    ListFsuDiscoveriesLifecycleStateAccepted,
	"IN_PROGRESS": ListFsuDiscoveriesLifecycleStateInProgress,
	"FAILED":      ListFsuDiscoveriesLifecycleStateFailed,
	"SUCCEEDED":   ListFsuDiscoveriesLifecycleStateSucceeded,
	"CANCELING":   ListFsuDiscoveriesLifecycleStateCanceling,
	"CANCELED":    ListFsuDiscoveriesLifecycleStateCanceled,
	"DELETING":    ListFsuDiscoveriesLifecycleStateDeleting,
	"DELETED":     ListFsuDiscoveriesLifecycleStateDeleted,
}

var mappingListFsuDiscoveriesLifecycleStateEnumLowerCase = map[string]ListFsuDiscoveriesLifecycleStateEnum{
	"accepted":    ListFsuDiscoveriesLifecycleStateAccepted,
	"in_progress": ListFsuDiscoveriesLifecycleStateInProgress,
	"failed":      ListFsuDiscoveriesLifecycleStateFailed,
	"succeeded":   ListFsuDiscoveriesLifecycleStateSucceeded,
	"canceling":   ListFsuDiscoveriesLifecycleStateCanceling,
	"canceled":    ListFsuDiscoveriesLifecycleStateCanceled,
	"deleting":    ListFsuDiscoveriesLifecycleStateDeleting,
	"deleted":     ListFsuDiscoveriesLifecycleStateDeleted,
}

// GetListFsuDiscoveriesLifecycleStateEnumValues Enumerates the set of values for ListFsuDiscoveriesLifecycleStateEnum
func GetListFsuDiscoveriesLifecycleStateEnumValues() []ListFsuDiscoveriesLifecycleStateEnum {
	values := make([]ListFsuDiscoveriesLifecycleStateEnum, 0)
	for _, v := range mappingListFsuDiscoveriesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuDiscoveriesLifecycleStateEnumStringValues Enumerates the set of values in String for ListFsuDiscoveriesLifecycleStateEnum
func GetListFsuDiscoveriesLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingListFsuDiscoveriesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuDiscoveriesLifecycleStateEnum(val string) (ListFsuDiscoveriesLifecycleStateEnum, bool) {
	enum, ok := mappingListFsuDiscoveriesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuDiscoveriesSortOrderEnum Enum with underlying type: string
type ListFsuDiscoveriesSortOrderEnum string

// Set of constants representing the allowable values for ListFsuDiscoveriesSortOrderEnum
const (
	ListFsuDiscoveriesSortOrderAsc  ListFsuDiscoveriesSortOrderEnum = "ASC"
	ListFsuDiscoveriesSortOrderDesc ListFsuDiscoveriesSortOrderEnum = "DESC"
)

var mappingListFsuDiscoveriesSortOrderEnum = map[string]ListFsuDiscoveriesSortOrderEnum{
	"ASC":  ListFsuDiscoveriesSortOrderAsc,
	"DESC": ListFsuDiscoveriesSortOrderDesc,
}

var mappingListFsuDiscoveriesSortOrderEnumLowerCase = map[string]ListFsuDiscoveriesSortOrderEnum{
	"asc":  ListFsuDiscoveriesSortOrderAsc,
	"desc": ListFsuDiscoveriesSortOrderDesc,
}

// GetListFsuDiscoveriesSortOrderEnumValues Enumerates the set of values for ListFsuDiscoveriesSortOrderEnum
func GetListFsuDiscoveriesSortOrderEnumValues() []ListFsuDiscoveriesSortOrderEnum {
	values := make([]ListFsuDiscoveriesSortOrderEnum, 0)
	for _, v := range mappingListFsuDiscoveriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuDiscoveriesSortOrderEnumStringValues Enumerates the set of values in String for ListFsuDiscoveriesSortOrderEnum
func GetListFsuDiscoveriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFsuDiscoveriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuDiscoveriesSortOrderEnum(val string) (ListFsuDiscoveriesSortOrderEnum, bool) {
	enum, ok := mappingListFsuDiscoveriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuDiscoveriesSortByEnum Enum with underlying type: string
type ListFsuDiscoveriesSortByEnum string

// Set of constants representing the allowable values for ListFsuDiscoveriesSortByEnum
const (
	ListFsuDiscoveriesSortByTimecreated ListFsuDiscoveriesSortByEnum = "timeCreated"
	ListFsuDiscoveriesSortByDisplayname ListFsuDiscoveriesSortByEnum = "displayName"
)

var mappingListFsuDiscoveriesSortByEnum = map[string]ListFsuDiscoveriesSortByEnum{
	"timeCreated": ListFsuDiscoveriesSortByTimecreated,
	"displayName": ListFsuDiscoveriesSortByDisplayname,
}

var mappingListFsuDiscoveriesSortByEnumLowerCase = map[string]ListFsuDiscoveriesSortByEnum{
	"timecreated": ListFsuDiscoveriesSortByTimecreated,
	"displayname": ListFsuDiscoveriesSortByDisplayname,
}

// GetListFsuDiscoveriesSortByEnumValues Enumerates the set of values for ListFsuDiscoveriesSortByEnum
func GetListFsuDiscoveriesSortByEnumValues() []ListFsuDiscoveriesSortByEnum {
	values := make([]ListFsuDiscoveriesSortByEnum, 0)
	for _, v := range mappingListFsuDiscoveriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuDiscoveriesSortByEnumStringValues Enumerates the set of values in String for ListFsuDiscoveriesSortByEnum
func GetListFsuDiscoveriesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListFsuDiscoveriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuDiscoveriesSortByEnum(val string) (ListFsuDiscoveriesSortByEnum, bool) {
	enum, ok := mappingListFsuDiscoveriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
