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

// ListFsuHomesRequest wrapper for the ListFsuHomes operation
type ListFsuHomesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListFsuHomesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFsuHomesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListFsuHomesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFsuHomesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFsuHomesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFsuHomesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFsuHomesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFsuHomesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFsuHomesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListFsuHomesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuHomesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFsuHomesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuHomesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFsuHomesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFsuHomesResponse wrapper for the ListFsuHomes operation
type ListFsuHomesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FsuHomeCollection instances
	FsuHomeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFsuHomesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFsuHomesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFsuHomesLifecycleStateEnum Enum with underlying type: string
type ListFsuHomesLifecycleStateEnum string

// Set of constants representing the allowable values for ListFsuHomesLifecycleStateEnum
const (
	ListFsuHomesLifecycleStateCreating       ListFsuHomesLifecycleStateEnum = "CREATING"
	ListFsuHomesLifecycleStateUpdating       ListFsuHomesLifecycleStateEnum = "UPDATING"
	ListFsuHomesLifecycleStateActive         ListFsuHomesLifecycleStateEnum = "ACTIVE"
	ListFsuHomesLifecycleStateNeedsAttention ListFsuHomesLifecycleStateEnum = "NEEDS_ATTENTION"
	ListFsuHomesLifecycleStateDeleting       ListFsuHomesLifecycleStateEnum = "DELETING"
	ListFsuHomesLifecycleStateDeleted        ListFsuHomesLifecycleStateEnum = "DELETED"
	ListFsuHomesLifecycleStateFailed         ListFsuHomesLifecycleStateEnum = "FAILED"
)

var mappingListFsuHomesLifecycleStateEnum = map[string]ListFsuHomesLifecycleStateEnum{
	"CREATING":        ListFsuHomesLifecycleStateCreating,
	"UPDATING":        ListFsuHomesLifecycleStateUpdating,
	"ACTIVE":          ListFsuHomesLifecycleStateActive,
	"NEEDS_ATTENTION": ListFsuHomesLifecycleStateNeedsAttention,
	"DELETING":        ListFsuHomesLifecycleStateDeleting,
	"DELETED":         ListFsuHomesLifecycleStateDeleted,
	"FAILED":          ListFsuHomesLifecycleStateFailed,
}

var mappingListFsuHomesLifecycleStateEnumLowerCase = map[string]ListFsuHomesLifecycleStateEnum{
	"creating":        ListFsuHomesLifecycleStateCreating,
	"updating":        ListFsuHomesLifecycleStateUpdating,
	"active":          ListFsuHomesLifecycleStateActive,
	"needs_attention": ListFsuHomesLifecycleStateNeedsAttention,
	"deleting":        ListFsuHomesLifecycleStateDeleting,
	"deleted":         ListFsuHomesLifecycleStateDeleted,
	"failed":          ListFsuHomesLifecycleStateFailed,
}

// GetListFsuHomesLifecycleStateEnumValues Enumerates the set of values for ListFsuHomesLifecycleStateEnum
func GetListFsuHomesLifecycleStateEnumValues() []ListFsuHomesLifecycleStateEnum {
	values := make([]ListFsuHomesLifecycleStateEnum, 0)
	for _, v := range mappingListFsuHomesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuHomesLifecycleStateEnumStringValues Enumerates the set of values in String for ListFsuHomesLifecycleStateEnum
func GetListFsuHomesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListFsuHomesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuHomesLifecycleStateEnum(val string) (ListFsuHomesLifecycleStateEnum, bool) {
	enum, ok := mappingListFsuHomesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuHomesSortOrderEnum Enum with underlying type: string
type ListFsuHomesSortOrderEnum string

// Set of constants representing the allowable values for ListFsuHomesSortOrderEnum
const (
	ListFsuHomesSortOrderAsc  ListFsuHomesSortOrderEnum = "ASC"
	ListFsuHomesSortOrderDesc ListFsuHomesSortOrderEnum = "DESC"
)

var mappingListFsuHomesSortOrderEnum = map[string]ListFsuHomesSortOrderEnum{
	"ASC":  ListFsuHomesSortOrderAsc,
	"DESC": ListFsuHomesSortOrderDesc,
}

var mappingListFsuHomesSortOrderEnumLowerCase = map[string]ListFsuHomesSortOrderEnum{
	"asc":  ListFsuHomesSortOrderAsc,
	"desc": ListFsuHomesSortOrderDesc,
}

// GetListFsuHomesSortOrderEnumValues Enumerates the set of values for ListFsuHomesSortOrderEnum
func GetListFsuHomesSortOrderEnumValues() []ListFsuHomesSortOrderEnum {
	values := make([]ListFsuHomesSortOrderEnum, 0)
	for _, v := range mappingListFsuHomesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuHomesSortOrderEnumStringValues Enumerates the set of values in String for ListFsuHomesSortOrderEnum
func GetListFsuHomesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFsuHomesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuHomesSortOrderEnum(val string) (ListFsuHomesSortOrderEnum, bool) {
	enum, ok := mappingListFsuHomesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuHomesSortByEnum Enum with underlying type: string
type ListFsuHomesSortByEnum string

// Set of constants representing the allowable values for ListFsuHomesSortByEnum
const (
	ListFsuHomesSortByTimecreated ListFsuHomesSortByEnum = "timeCreated"
	ListFsuHomesSortByDisplayname ListFsuHomesSortByEnum = "displayName"
)

var mappingListFsuHomesSortByEnum = map[string]ListFsuHomesSortByEnum{
	"timeCreated": ListFsuHomesSortByTimecreated,
	"displayName": ListFsuHomesSortByDisplayname,
}

var mappingListFsuHomesSortByEnumLowerCase = map[string]ListFsuHomesSortByEnum{
	"timecreated": ListFsuHomesSortByTimecreated,
	"displayname": ListFsuHomesSortByDisplayname,
}

// GetListFsuHomesSortByEnumValues Enumerates the set of values for ListFsuHomesSortByEnum
func GetListFsuHomesSortByEnumValues() []ListFsuHomesSortByEnum {
	values := make([]ListFsuHomesSortByEnum, 0)
	for _, v := range mappingListFsuHomesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuHomesSortByEnumStringValues Enumerates the set of values in String for ListFsuHomesSortByEnum
func GetListFsuHomesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListFsuHomesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuHomesSortByEnum(val string) (ListFsuHomesSortByEnum, bool) {
	enum, ok := mappingListFsuHomesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
