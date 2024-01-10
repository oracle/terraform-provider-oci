// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bastion

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSessionsRequest wrapper for the ListSessions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bastion/ListSessions.go.html to see an example of how to use ListSessionsRequest.
type ListSessionsRequest struct {

	// The unique identifier (OCID) of the bastion in which to list sessions.
	BastionId *string `mandatory:"true" contributesTo:"query" name:"bastionId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	SessionLifecycleState ListSessionsSessionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"sessionLifecycleState" omitEmpty:"true"`

	// The unique identifier (OCID) of the session in which to list resources.
	SessionId *string `mandatory:"false" contributesTo:"query" name:"sessionId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListSessionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListSessionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSessionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSessionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSessionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSessionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSessionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSessionsSessionLifecycleStateEnum(string(request.SessionLifecycleState)); !ok && request.SessionLifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SessionLifecycleState: %s. Supported values are: %s.", request.SessionLifecycleState, strings.Join(GetListSessionsSessionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSessionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSessionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSessionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSessionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSessionsResponse wrapper for the ListSessions operation
type ListSessionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SessionSummary instances
	Items []SessionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSessionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSessionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSessionsSessionLifecycleStateEnum Enum with underlying type: string
type ListSessionsSessionLifecycleStateEnum string

// Set of constants representing the allowable values for ListSessionsSessionLifecycleStateEnum
const (
	ListSessionsSessionLifecycleStateCreating ListSessionsSessionLifecycleStateEnum = "CREATING"
	ListSessionsSessionLifecycleStateActive   ListSessionsSessionLifecycleStateEnum = "ACTIVE"
	ListSessionsSessionLifecycleStateDeleting ListSessionsSessionLifecycleStateEnum = "DELETING"
	ListSessionsSessionLifecycleStateDeleted  ListSessionsSessionLifecycleStateEnum = "DELETED"
	ListSessionsSessionLifecycleStateFailed   ListSessionsSessionLifecycleStateEnum = "FAILED"
)

var mappingListSessionsSessionLifecycleStateEnum = map[string]ListSessionsSessionLifecycleStateEnum{
	"CREATING": ListSessionsSessionLifecycleStateCreating,
	"ACTIVE":   ListSessionsSessionLifecycleStateActive,
	"DELETING": ListSessionsSessionLifecycleStateDeleting,
	"DELETED":  ListSessionsSessionLifecycleStateDeleted,
	"FAILED":   ListSessionsSessionLifecycleStateFailed,
}

var mappingListSessionsSessionLifecycleStateEnumLowerCase = map[string]ListSessionsSessionLifecycleStateEnum{
	"creating": ListSessionsSessionLifecycleStateCreating,
	"active":   ListSessionsSessionLifecycleStateActive,
	"deleting": ListSessionsSessionLifecycleStateDeleting,
	"deleted":  ListSessionsSessionLifecycleStateDeleted,
	"failed":   ListSessionsSessionLifecycleStateFailed,
}

// GetListSessionsSessionLifecycleStateEnumValues Enumerates the set of values for ListSessionsSessionLifecycleStateEnum
func GetListSessionsSessionLifecycleStateEnumValues() []ListSessionsSessionLifecycleStateEnum {
	values := make([]ListSessionsSessionLifecycleStateEnum, 0)
	for _, v := range mappingListSessionsSessionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSessionsSessionLifecycleStateEnumStringValues Enumerates the set of values in String for ListSessionsSessionLifecycleStateEnum
func GetListSessionsSessionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListSessionsSessionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSessionsSessionLifecycleStateEnum(val string) (ListSessionsSessionLifecycleStateEnum, bool) {
	enum, ok := mappingListSessionsSessionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSessionsSortOrderEnum Enum with underlying type: string
type ListSessionsSortOrderEnum string

// Set of constants representing the allowable values for ListSessionsSortOrderEnum
const (
	ListSessionsSortOrderAsc  ListSessionsSortOrderEnum = "ASC"
	ListSessionsSortOrderDesc ListSessionsSortOrderEnum = "DESC"
)

var mappingListSessionsSortOrderEnum = map[string]ListSessionsSortOrderEnum{
	"ASC":  ListSessionsSortOrderAsc,
	"DESC": ListSessionsSortOrderDesc,
}

var mappingListSessionsSortOrderEnumLowerCase = map[string]ListSessionsSortOrderEnum{
	"asc":  ListSessionsSortOrderAsc,
	"desc": ListSessionsSortOrderDesc,
}

// GetListSessionsSortOrderEnumValues Enumerates the set of values for ListSessionsSortOrderEnum
func GetListSessionsSortOrderEnumValues() []ListSessionsSortOrderEnum {
	values := make([]ListSessionsSortOrderEnum, 0)
	for _, v := range mappingListSessionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSessionsSortOrderEnumStringValues Enumerates the set of values in String for ListSessionsSortOrderEnum
func GetListSessionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSessionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSessionsSortOrderEnum(val string) (ListSessionsSortOrderEnum, bool) {
	enum, ok := mappingListSessionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSessionsSortByEnum Enum with underlying type: string
type ListSessionsSortByEnum string

// Set of constants representing the allowable values for ListSessionsSortByEnum
const (
	ListSessionsSortByTimecreated ListSessionsSortByEnum = "timeCreated"
	ListSessionsSortByDisplayname ListSessionsSortByEnum = "displayName"
)

var mappingListSessionsSortByEnum = map[string]ListSessionsSortByEnum{
	"timeCreated": ListSessionsSortByTimecreated,
	"displayName": ListSessionsSortByDisplayname,
}

var mappingListSessionsSortByEnumLowerCase = map[string]ListSessionsSortByEnum{
	"timecreated": ListSessionsSortByTimecreated,
	"displayname": ListSessionsSortByDisplayname,
}

// GetListSessionsSortByEnumValues Enumerates the set of values for ListSessionsSortByEnum
func GetListSessionsSortByEnumValues() []ListSessionsSortByEnum {
	values := make([]ListSessionsSortByEnum, 0)
	for _, v := range mappingListSessionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSessionsSortByEnumStringValues Enumerates the set of values in String for ListSessionsSortByEnum
func GetListSessionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSessionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSessionsSortByEnum(val string) (ListSessionsSortByEnum, bool) {
	enum, ok := mappingListSessionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
