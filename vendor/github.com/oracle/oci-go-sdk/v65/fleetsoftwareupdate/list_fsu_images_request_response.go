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

// ListFsuImagesRequest wrapper for the ListFsuImages operation
type ListFsuImagesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListFsuImagesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFsuImagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListFsuImagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFsuImagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFsuImagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFsuImagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFsuImagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFsuImagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFsuImagesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListFsuImagesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuImagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFsuImagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuImagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFsuImagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFsuImagesResponse wrapper for the ListFsuImages operation
type ListFsuImagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FsuImageCollection instances
	FsuImageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFsuImagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFsuImagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFsuImagesLifecycleStateEnum Enum with underlying type: string
type ListFsuImagesLifecycleStateEnum string

// Set of constants representing the allowable values for ListFsuImagesLifecycleStateEnum
const (
	ListFsuImagesLifecycleStateCreating       ListFsuImagesLifecycleStateEnum = "CREATING"
	ListFsuImagesLifecycleStateUpdating       ListFsuImagesLifecycleStateEnum = "UPDATING"
	ListFsuImagesLifecycleStateActive         ListFsuImagesLifecycleStateEnum = "ACTIVE"
	ListFsuImagesLifecycleStateNeedsAttention ListFsuImagesLifecycleStateEnum = "NEEDS_ATTENTION"
	ListFsuImagesLifecycleStateDeleting       ListFsuImagesLifecycleStateEnum = "DELETING"
	ListFsuImagesLifecycleStateDeleted        ListFsuImagesLifecycleStateEnum = "DELETED"
	ListFsuImagesLifecycleStateFailed         ListFsuImagesLifecycleStateEnum = "FAILED"
)

var mappingListFsuImagesLifecycleStateEnum = map[string]ListFsuImagesLifecycleStateEnum{
	"CREATING":        ListFsuImagesLifecycleStateCreating,
	"UPDATING":        ListFsuImagesLifecycleStateUpdating,
	"ACTIVE":          ListFsuImagesLifecycleStateActive,
	"NEEDS_ATTENTION": ListFsuImagesLifecycleStateNeedsAttention,
	"DELETING":        ListFsuImagesLifecycleStateDeleting,
	"DELETED":         ListFsuImagesLifecycleStateDeleted,
	"FAILED":          ListFsuImagesLifecycleStateFailed,
}

var mappingListFsuImagesLifecycleStateEnumLowerCase = map[string]ListFsuImagesLifecycleStateEnum{
	"creating":        ListFsuImagesLifecycleStateCreating,
	"updating":        ListFsuImagesLifecycleStateUpdating,
	"active":          ListFsuImagesLifecycleStateActive,
	"needs_attention": ListFsuImagesLifecycleStateNeedsAttention,
	"deleting":        ListFsuImagesLifecycleStateDeleting,
	"deleted":         ListFsuImagesLifecycleStateDeleted,
	"failed":          ListFsuImagesLifecycleStateFailed,
}

// GetListFsuImagesLifecycleStateEnumValues Enumerates the set of values for ListFsuImagesLifecycleStateEnum
func GetListFsuImagesLifecycleStateEnumValues() []ListFsuImagesLifecycleStateEnum {
	values := make([]ListFsuImagesLifecycleStateEnum, 0)
	for _, v := range mappingListFsuImagesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuImagesLifecycleStateEnumStringValues Enumerates the set of values in String for ListFsuImagesLifecycleStateEnum
func GetListFsuImagesLifecycleStateEnumStringValues() []string {
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

// GetMappingListFsuImagesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuImagesLifecycleStateEnum(val string) (ListFsuImagesLifecycleStateEnum, bool) {
	enum, ok := mappingListFsuImagesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuImagesSortOrderEnum Enum with underlying type: string
type ListFsuImagesSortOrderEnum string

// Set of constants representing the allowable values for ListFsuImagesSortOrderEnum
const (
	ListFsuImagesSortOrderAsc  ListFsuImagesSortOrderEnum = "ASC"
	ListFsuImagesSortOrderDesc ListFsuImagesSortOrderEnum = "DESC"
)

var mappingListFsuImagesSortOrderEnum = map[string]ListFsuImagesSortOrderEnum{
	"ASC":  ListFsuImagesSortOrderAsc,
	"DESC": ListFsuImagesSortOrderDesc,
}

var mappingListFsuImagesSortOrderEnumLowerCase = map[string]ListFsuImagesSortOrderEnum{
	"asc":  ListFsuImagesSortOrderAsc,
	"desc": ListFsuImagesSortOrderDesc,
}

// GetListFsuImagesSortOrderEnumValues Enumerates the set of values for ListFsuImagesSortOrderEnum
func GetListFsuImagesSortOrderEnumValues() []ListFsuImagesSortOrderEnum {
	values := make([]ListFsuImagesSortOrderEnum, 0)
	for _, v := range mappingListFsuImagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuImagesSortOrderEnumStringValues Enumerates the set of values in String for ListFsuImagesSortOrderEnum
func GetListFsuImagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFsuImagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuImagesSortOrderEnum(val string) (ListFsuImagesSortOrderEnum, bool) {
	enum, ok := mappingListFsuImagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuImagesSortByEnum Enum with underlying type: string
type ListFsuImagesSortByEnum string

// Set of constants representing the allowable values for ListFsuImagesSortByEnum
const (
	ListFsuImagesSortByTimecreated ListFsuImagesSortByEnum = "timeCreated"
	ListFsuImagesSortByDisplayname ListFsuImagesSortByEnum = "displayName"
)

var mappingListFsuImagesSortByEnum = map[string]ListFsuImagesSortByEnum{
	"timeCreated": ListFsuImagesSortByTimecreated,
	"displayName": ListFsuImagesSortByDisplayname,
}

var mappingListFsuImagesSortByEnumLowerCase = map[string]ListFsuImagesSortByEnum{
	"timecreated": ListFsuImagesSortByTimecreated,
	"displayname": ListFsuImagesSortByDisplayname,
}

// GetListFsuImagesSortByEnumValues Enumerates the set of values for ListFsuImagesSortByEnum
func GetListFsuImagesSortByEnumValues() []ListFsuImagesSortByEnum {
	values := make([]ListFsuImagesSortByEnum, 0)
	for _, v := range mappingListFsuImagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuImagesSortByEnumStringValues Enumerates the set of values in String for ListFsuImagesSortByEnum
func GetListFsuImagesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListFsuImagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuImagesSortByEnum(val string) (ListFsuImagesSortByEnum, bool) {
	enum, ok := mappingListFsuImagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
