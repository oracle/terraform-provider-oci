// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAdvancedLinksRequest wrapper for the ListAdvancedLinks operation
type ListAdvancedLinksRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The subject tenant name prefix
	SubjectTenantNamePrefix *string `mandatory:"false" contributesTo:"query" name:"subjectTenantNamePrefix"`

	// The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListAdvancedLinksLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListAdvancedLinksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListAdvancedLinksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAdvancedLinksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAdvancedLinksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAdvancedLinksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAdvancedLinksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAdvancedLinksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAdvancedLinksLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAdvancedLinksLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAdvancedLinksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAdvancedLinksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAdvancedLinksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAdvancedLinksSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAdvancedLinksResponse wrapper for the ListAdvancedLinks operation
type ListAdvancedLinksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AdvancedLinkCollection instances
	AdvancedLinkCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAdvancedLinksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAdvancedLinksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAdvancedLinksLifecycleStateEnum Enum with underlying type: string
type ListAdvancedLinksLifecycleStateEnum string

// Set of constants representing the allowable values for ListAdvancedLinksLifecycleStateEnum
const (
	ListAdvancedLinksLifecycleStateCreating ListAdvancedLinksLifecycleStateEnum = "CREATING"
	ListAdvancedLinksLifecycleStateUpdating ListAdvancedLinksLifecycleStateEnum = "UPDATING"
	ListAdvancedLinksLifecycleStateActive   ListAdvancedLinksLifecycleStateEnum = "ACTIVE"
	ListAdvancedLinksLifecycleStateInactive ListAdvancedLinksLifecycleStateEnum = "INACTIVE"
	ListAdvancedLinksLifecycleStateDeleting ListAdvancedLinksLifecycleStateEnum = "DELETING"
	ListAdvancedLinksLifecycleStateDeleted  ListAdvancedLinksLifecycleStateEnum = "DELETED"
	ListAdvancedLinksLifecycleStateFailed   ListAdvancedLinksLifecycleStateEnum = "FAILED"
)

var mappingListAdvancedLinksLifecycleStateEnum = map[string]ListAdvancedLinksLifecycleStateEnum{
	"CREATING": ListAdvancedLinksLifecycleStateCreating,
	"UPDATING": ListAdvancedLinksLifecycleStateUpdating,
	"ACTIVE":   ListAdvancedLinksLifecycleStateActive,
	"INACTIVE": ListAdvancedLinksLifecycleStateInactive,
	"DELETING": ListAdvancedLinksLifecycleStateDeleting,
	"DELETED":  ListAdvancedLinksLifecycleStateDeleted,
	"FAILED":   ListAdvancedLinksLifecycleStateFailed,
}

var mappingListAdvancedLinksLifecycleStateEnumLowerCase = map[string]ListAdvancedLinksLifecycleStateEnum{
	"creating": ListAdvancedLinksLifecycleStateCreating,
	"updating": ListAdvancedLinksLifecycleStateUpdating,
	"active":   ListAdvancedLinksLifecycleStateActive,
	"inactive": ListAdvancedLinksLifecycleStateInactive,
	"deleting": ListAdvancedLinksLifecycleStateDeleting,
	"deleted":  ListAdvancedLinksLifecycleStateDeleted,
	"failed":   ListAdvancedLinksLifecycleStateFailed,
}

// GetListAdvancedLinksLifecycleStateEnumValues Enumerates the set of values for ListAdvancedLinksLifecycleStateEnum
func GetListAdvancedLinksLifecycleStateEnumValues() []ListAdvancedLinksLifecycleStateEnum {
	values := make([]ListAdvancedLinksLifecycleStateEnum, 0)
	for _, v := range mappingListAdvancedLinksLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdvancedLinksLifecycleStateEnumStringValues Enumerates the set of values in String for ListAdvancedLinksLifecycleStateEnum
func GetListAdvancedLinksLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListAdvancedLinksLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdvancedLinksLifecycleStateEnum(val string) (ListAdvancedLinksLifecycleStateEnum, bool) {
	enum, ok := mappingListAdvancedLinksLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAdvancedLinksSortOrderEnum Enum with underlying type: string
type ListAdvancedLinksSortOrderEnum string

// Set of constants representing the allowable values for ListAdvancedLinksSortOrderEnum
const (
	ListAdvancedLinksSortOrderAsc  ListAdvancedLinksSortOrderEnum = "ASC"
	ListAdvancedLinksSortOrderDesc ListAdvancedLinksSortOrderEnum = "DESC"
)

var mappingListAdvancedLinksSortOrderEnum = map[string]ListAdvancedLinksSortOrderEnum{
	"ASC":  ListAdvancedLinksSortOrderAsc,
	"DESC": ListAdvancedLinksSortOrderDesc,
}

var mappingListAdvancedLinksSortOrderEnumLowerCase = map[string]ListAdvancedLinksSortOrderEnum{
	"asc":  ListAdvancedLinksSortOrderAsc,
	"desc": ListAdvancedLinksSortOrderDesc,
}

// GetListAdvancedLinksSortOrderEnumValues Enumerates the set of values for ListAdvancedLinksSortOrderEnum
func GetListAdvancedLinksSortOrderEnumValues() []ListAdvancedLinksSortOrderEnum {
	values := make([]ListAdvancedLinksSortOrderEnum, 0)
	for _, v := range mappingListAdvancedLinksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdvancedLinksSortOrderEnumStringValues Enumerates the set of values in String for ListAdvancedLinksSortOrderEnum
func GetListAdvancedLinksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAdvancedLinksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdvancedLinksSortOrderEnum(val string) (ListAdvancedLinksSortOrderEnum, bool) {
	enum, ok := mappingListAdvancedLinksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAdvancedLinksSortByEnum Enum with underlying type: string
type ListAdvancedLinksSortByEnum string

// Set of constants representing the allowable values for ListAdvancedLinksSortByEnum
const (
	ListAdvancedLinksSortByTimecreated ListAdvancedLinksSortByEnum = "timeCreated"
	ListAdvancedLinksSortByDisplayname ListAdvancedLinksSortByEnum = "displayName"
)

var mappingListAdvancedLinksSortByEnum = map[string]ListAdvancedLinksSortByEnum{
	"timeCreated": ListAdvancedLinksSortByTimecreated,
	"displayName": ListAdvancedLinksSortByDisplayname,
}

var mappingListAdvancedLinksSortByEnumLowerCase = map[string]ListAdvancedLinksSortByEnum{
	"timecreated": ListAdvancedLinksSortByTimecreated,
	"displayname": ListAdvancedLinksSortByDisplayname,
}

// GetListAdvancedLinksSortByEnumValues Enumerates the set of values for ListAdvancedLinksSortByEnum
func GetListAdvancedLinksSortByEnumValues() []ListAdvancedLinksSortByEnum {
	values := make([]ListAdvancedLinksSortByEnum, 0)
	for _, v := range mappingListAdvancedLinksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdvancedLinksSortByEnumStringValues Enumerates the set of values in String for ListAdvancedLinksSortByEnum
func GetListAdvancedLinksSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAdvancedLinksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdvancedLinksSortByEnum(val string) (ListAdvancedLinksSortByEnum, bool) {
	enum, ok := mappingListAdvancedLinksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
