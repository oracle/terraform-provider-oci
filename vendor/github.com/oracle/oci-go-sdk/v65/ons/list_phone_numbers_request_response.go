// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ons

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPhoneNumbersRequest wrapper for the ListPhoneNumbers operation
type ListPhoneNumbersRequest struct {

	// unique PhoneApplication identifier
	PhoneApplicationId *string `mandatory:"true" contributesTo:"query" name:"phoneApplicationId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListPhoneNumbersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique PhoneNumber identifier
	PhoneNumberId *string `mandatory:"false" contributesTo:"query" name:"phoneNumberId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use (ascending or descending).
	SortOrder ListPhoneNumbersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one field can be selected for sorting.
	SortBy ListPhoneNumbersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPhoneNumbersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPhoneNumbersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPhoneNumbersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPhoneNumbersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPhoneNumbersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPhoneNumbersLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListPhoneNumbersLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPhoneNumbersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPhoneNumbersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPhoneNumbersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPhoneNumbersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPhoneNumbersResponse wrapper for the ListPhoneNumbers operation
type ListPhoneNumbersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PhoneNumberCollection instances
	PhoneNumberCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPhoneNumbersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPhoneNumbersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPhoneNumbersLifecycleStateEnum Enum with underlying type: string
type ListPhoneNumbersLifecycleStateEnum string

// Set of constants representing the allowable values for ListPhoneNumbersLifecycleStateEnum
const (
	ListPhoneNumbersLifecycleStateCreating ListPhoneNumbersLifecycleStateEnum = "CREATING"
	ListPhoneNumbersLifecycleStateActive   ListPhoneNumbersLifecycleStateEnum = "ACTIVE"
	ListPhoneNumbersLifecycleStateUpdating ListPhoneNumbersLifecycleStateEnum = "UPDATING"
	ListPhoneNumbersLifecycleStateInactive ListPhoneNumbersLifecycleStateEnum = "INACTIVE"
	ListPhoneNumbersLifecycleStateDeleting ListPhoneNumbersLifecycleStateEnum = "DELETING"
	ListPhoneNumbersLifecycleStateDeleted  ListPhoneNumbersLifecycleStateEnum = "DELETED"
	ListPhoneNumbersLifecycleStateFailed   ListPhoneNumbersLifecycleStateEnum = "FAILED"
)

var mappingListPhoneNumbersLifecycleStateEnum = map[string]ListPhoneNumbersLifecycleStateEnum{
	"CREATING": ListPhoneNumbersLifecycleStateCreating,
	"ACTIVE":   ListPhoneNumbersLifecycleStateActive,
	"UPDATING": ListPhoneNumbersLifecycleStateUpdating,
	"INACTIVE": ListPhoneNumbersLifecycleStateInactive,
	"DELETING": ListPhoneNumbersLifecycleStateDeleting,
	"DELETED":  ListPhoneNumbersLifecycleStateDeleted,
	"FAILED":   ListPhoneNumbersLifecycleStateFailed,
}

var mappingListPhoneNumbersLifecycleStateEnumLowerCase = map[string]ListPhoneNumbersLifecycleStateEnum{
	"creating": ListPhoneNumbersLifecycleStateCreating,
	"active":   ListPhoneNumbersLifecycleStateActive,
	"updating": ListPhoneNumbersLifecycleStateUpdating,
	"inactive": ListPhoneNumbersLifecycleStateInactive,
	"deleting": ListPhoneNumbersLifecycleStateDeleting,
	"deleted":  ListPhoneNumbersLifecycleStateDeleted,
	"failed":   ListPhoneNumbersLifecycleStateFailed,
}

// GetListPhoneNumbersLifecycleStateEnumValues Enumerates the set of values for ListPhoneNumbersLifecycleStateEnum
func GetListPhoneNumbersLifecycleStateEnumValues() []ListPhoneNumbersLifecycleStateEnum {
	values := make([]ListPhoneNumbersLifecycleStateEnum, 0)
	for _, v := range mappingListPhoneNumbersLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListPhoneNumbersLifecycleStateEnumStringValues Enumerates the set of values in String for ListPhoneNumbersLifecycleStateEnum
func GetListPhoneNumbersLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListPhoneNumbersLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPhoneNumbersLifecycleStateEnum(val string) (ListPhoneNumbersLifecycleStateEnum, bool) {
	enum, ok := mappingListPhoneNumbersLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPhoneNumbersSortOrderEnum Enum with underlying type: string
type ListPhoneNumbersSortOrderEnum string

// Set of constants representing the allowable values for ListPhoneNumbersSortOrderEnum
const (
	ListPhoneNumbersSortOrderAsc  ListPhoneNumbersSortOrderEnum = "ASC"
	ListPhoneNumbersSortOrderDesc ListPhoneNumbersSortOrderEnum = "DESC"
)

var mappingListPhoneNumbersSortOrderEnum = map[string]ListPhoneNumbersSortOrderEnum{
	"ASC":  ListPhoneNumbersSortOrderAsc,
	"DESC": ListPhoneNumbersSortOrderDesc,
}

var mappingListPhoneNumbersSortOrderEnumLowerCase = map[string]ListPhoneNumbersSortOrderEnum{
	"asc":  ListPhoneNumbersSortOrderAsc,
	"desc": ListPhoneNumbersSortOrderDesc,
}

// GetListPhoneNumbersSortOrderEnumValues Enumerates the set of values for ListPhoneNumbersSortOrderEnum
func GetListPhoneNumbersSortOrderEnumValues() []ListPhoneNumbersSortOrderEnum {
	values := make([]ListPhoneNumbersSortOrderEnum, 0)
	for _, v := range mappingListPhoneNumbersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPhoneNumbersSortOrderEnumStringValues Enumerates the set of values in String for ListPhoneNumbersSortOrderEnum
func GetListPhoneNumbersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPhoneNumbersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPhoneNumbersSortOrderEnum(val string) (ListPhoneNumbersSortOrderEnum, bool) {
	enum, ok := mappingListPhoneNumbersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPhoneNumbersSortByEnum Enum with underlying type: string
type ListPhoneNumbersSortByEnum string

// Set of constants representing the allowable values for ListPhoneNumbersSortByEnum
const (
	ListPhoneNumbersSortByTimecreated    ListPhoneNumbersSortByEnum = "TIMECREATED"
	ListPhoneNumbersSortByLifecyclestate ListPhoneNumbersSortByEnum = "LIFECYCLESTATE"
)

var mappingListPhoneNumbersSortByEnum = map[string]ListPhoneNumbersSortByEnum{
	"TIMECREATED":    ListPhoneNumbersSortByTimecreated,
	"LIFECYCLESTATE": ListPhoneNumbersSortByLifecyclestate,
}

var mappingListPhoneNumbersSortByEnumLowerCase = map[string]ListPhoneNumbersSortByEnum{
	"timecreated":    ListPhoneNumbersSortByTimecreated,
	"lifecyclestate": ListPhoneNumbersSortByLifecyclestate,
}

// GetListPhoneNumbersSortByEnumValues Enumerates the set of values for ListPhoneNumbersSortByEnum
func GetListPhoneNumbersSortByEnumValues() []ListPhoneNumbersSortByEnum {
	values := make([]ListPhoneNumbersSortByEnum, 0)
	for _, v := range mappingListPhoneNumbersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPhoneNumbersSortByEnumStringValues Enumerates the set of values in String for ListPhoneNumbersSortByEnum
func GetListPhoneNumbersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"LIFECYCLESTATE",
	}
}

// GetMappingListPhoneNumbersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPhoneNumbersSortByEnum(val string) (ListPhoneNumbersSortByEnum, bool) {
	enum, ok := mappingListPhoneNumbersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
