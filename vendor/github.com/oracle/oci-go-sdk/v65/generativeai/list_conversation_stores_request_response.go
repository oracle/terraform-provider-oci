// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListConversationStoresRequest wrapper for the ListConversationStores operation
type ListConversationStoresRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycle state matches the given value.
	LifecycleState ConversationStoreLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the conversationStore.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListConversationStoresSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListConversationStoresSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConversationStoresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConversationStoresRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConversationStoresRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConversationStoresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListConversationStoresRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConversationStoreLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetConversationStoreLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConversationStoresSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConversationStoresSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConversationStoresSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConversationStoresSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListConversationStoresResponse wrapper for the ListConversationStores operation
type ListConversationStoresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConversationStoreCollection instances
	ConversationStoreCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListConversationStoresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConversationStoresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConversationStoresSortOrderEnum Enum with underlying type: string
type ListConversationStoresSortOrderEnum string

// Set of constants representing the allowable values for ListConversationStoresSortOrderEnum
const (
	ListConversationStoresSortOrderAsc  ListConversationStoresSortOrderEnum = "ASC"
	ListConversationStoresSortOrderDesc ListConversationStoresSortOrderEnum = "DESC"
)

var mappingListConversationStoresSortOrderEnum = map[string]ListConversationStoresSortOrderEnum{
	"ASC":  ListConversationStoresSortOrderAsc,
	"DESC": ListConversationStoresSortOrderDesc,
}

var mappingListConversationStoresSortOrderEnumLowerCase = map[string]ListConversationStoresSortOrderEnum{
	"asc":  ListConversationStoresSortOrderAsc,
	"desc": ListConversationStoresSortOrderDesc,
}

// GetListConversationStoresSortOrderEnumValues Enumerates the set of values for ListConversationStoresSortOrderEnum
func GetListConversationStoresSortOrderEnumValues() []ListConversationStoresSortOrderEnum {
	values := make([]ListConversationStoresSortOrderEnum, 0)
	for _, v := range mappingListConversationStoresSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListConversationStoresSortOrderEnumStringValues Enumerates the set of values in String for ListConversationStoresSortOrderEnum
func GetListConversationStoresSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListConversationStoresSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConversationStoresSortOrderEnum(val string) (ListConversationStoresSortOrderEnum, bool) {
	enum, ok := mappingListConversationStoresSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConversationStoresSortByEnum Enum with underlying type: string
type ListConversationStoresSortByEnum string

// Set of constants representing the allowable values for ListConversationStoresSortByEnum
const (
	ListConversationStoresSortByDisplayname ListConversationStoresSortByEnum = "displayName"
	ListConversationStoresSortByTimecreated ListConversationStoresSortByEnum = "timeCreated"
)

var mappingListConversationStoresSortByEnum = map[string]ListConversationStoresSortByEnum{
	"displayName": ListConversationStoresSortByDisplayname,
	"timeCreated": ListConversationStoresSortByTimecreated,
}

var mappingListConversationStoresSortByEnumLowerCase = map[string]ListConversationStoresSortByEnum{
	"displayname": ListConversationStoresSortByDisplayname,
	"timecreated": ListConversationStoresSortByTimecreated,
}

// GetListConversationStoresSortByEnumValues Enumerates the set of values for ListConversationStoresSortByEnum
func GetListConversationStoresSortByEnumValues() []ListConversationStoresSortByEnum {
	values := make([]ListConversationStoresSortByEnum, 0)
	for _, v := range mappingListConversationStoresSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListConversationStoresSortByEnumStringValues Enumerates the set of values in String for ListConversationStoresSortByEnum
func GetListConversationStoresSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListConversationStoresSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConversationStoresSortByEnum(val string) (ListConversationStoresSortByEnum, bool) {
	enum, ok := mappingListConversationStoresSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
