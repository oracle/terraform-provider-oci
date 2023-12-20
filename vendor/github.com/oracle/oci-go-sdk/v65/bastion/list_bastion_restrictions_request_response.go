// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bastion

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBastionRestrictionsRequest wrapper for the ListBastionRestrictions operation
type ListBastionRestrictionsRequest struct {

	// The unique identifier (OCID) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListBastionRestrictionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique bastion restriction identifier (OCID).
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListBastionRestrictionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending. If no value is specified timeCreated is default.
	SortBy ListBastionRestrictionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBastionRestrictionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBastionRestrictionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBastionRestrictionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBastionRestrictionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBastionRestrictionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListBastionRestrictionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListBastionRestrictionsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBastionRestrictionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBastionRestrictionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBastionRestrictionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBastionRestrictionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBastionRestrictionsResponse wrapper for the ListBastionRestrictions operation
type ListBastionRestrictionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BastionRestrictionCollection instances
	BastionRestrictionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBastionRestrictionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBastionRestrictionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBastionRestrictionsLifecycleStateEnum Enum with underlying type: string
type ListBastionRestrictionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListBastionRestrictionsLifecycleStateEnum
const (
	ListBastionRestrictionsLifecycleStateCreating ListBastionRestrictionsLifecycleStateEnum = "CREATING"
	ListBastionRestrictionsLifecycleStateActive   ListBastionRestrictionsLifecycleStateEnum = "ACTIVE"
	ListBastionRestrictionsLifecycleStateDeleting ListBastionRestrictionsLifecycleStateEnum = "DELETING"
	ListBastionRestrictionsLifecycleStateDeleted  ListBastionRestrictionsLifecycleStateEnum = "DELETED"
	ListBastionRestrictionsLifecycleStateFailed   ListBastionRestrictionsLifecycleStateEnum = "FAILED"
)

var mappingListBastionRestrictionsLifecycleStateEnum = map[string]ListBastionRestrictionsLifecycleStateEnum{
	"CREATING": ListBastionRestrictionsLifecycleStateCreating,
	"ACTIVE":   ListBastionRestrictionsLifecycleStateActive,
	"DELETING": ListBastionRestrictionsLifecycleStateDeleting,
	"DELETED":  ListBastionRestrictionsLifecycleStateDeleted,
	"FAILED":   ListBastionRestrictionsLifecycleStateFailed,
}

var mappingListBastionRestrictionsLifecycleStateEnumLowerCase = map[string]ListBastionRestrictionsLifecycleStateEnum{
	"creating": ListBastionRestrictionsLifecycleStateCreating,
	"active":   ListBastionRestrictionsLifecycleStateActive,
	"deleting": ListBastionRestrictionsLifecycleStateDeleting,
	"deleted":  ListBastionRestrictionsLifecycleStateDeleted,
	"failed":   ListBastionRestrictionsLifecycleStateFailed,
}

// GetListBastionRestrictionsLifecycleStateEnumValues Enumerates the set of values for ListBastionRestrictionsLifecycleStateEnum
func GetListBastionRestrictionsLifecycleStateEnumValues() []ListBastionRestrictionsLifecycleStateEnum {
	values := make([]ListBastionRestrictionsLifecycleStateEnum, 0)
	for _, v := range mappingListBastionRestrictionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListBastionRestrictionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListBastionRestrictionsLifecycleStateEnum
func GetListBastionRestrictionsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListBastionRestrictionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBastionRestrictionsLifecycleStateEnum(val string) (ListBastionRestrictionsLifecycleStateEnum, bool) {
	enum, ok := mappingListBastionRestrictionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBastionRestrictionsSortOrderEnum Enum with underlying type: string
type ListBastionRestrictionsSortOrderEnum string

// Set of constants representing the allowable values for ListBastionRestrictionsSortOrderEnum
const (
	ListBastionRestrictionsSortOrderAsc  ListBastionRestrictionsSortOrderEnum = "ASC"
	ListBastionRestrictionsSortOrderDesc ListBastionRestrictionsSortOrderEnum = "DESC"
)

var mappingListBastionRestrictionsSortOrderEnum = map[string]ListBastionRestrictionsSortOrderEnum{
	"ASC":  ListBastionRestrictionsSortOrderAsc,
	"DESC": ListBastionRestrictionsSortOrderDesc,
}

var mappingListBastionRestrictionsSortOrderEnumLowerCase = map[string]ListBastionRestrictionsSortOrderEnum{
	"asc":  ListBastionRestrictionsSortOrderAsc,
	"desc": ListBastionRestrictionsSortOrderDesc,
}

// GetListBastionRestrictionsSortOrderEnumValues Enumerates the set of values for ListBastionRestrictionsSortOrderEnum
func GetListBastionRestrictionsSortOrderEnumValues() []ListBastionRestrictionsSortOrderEnum {
	values := make([]ListBastionRestrictionsSortOrderEnum, 0)
	for _, v := range mappingListBastionRestrictionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBastionRestrictionsSortOrderEnumStringValues Enumerates the set of values in String for ListBastionRestrictionsSortOrderEnum
func GetListBastionRestrictionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBastionRestrictionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBastionRestrictionsSortOrderEnum(val string) (ListBastionRestrictionsSortOrderEnum, bool) {
	enum, ok := mappingListBastionRestrictionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBastionRestrictionsSortByEnum Enum with underlying type: string
type ListBastionRestrictionsSortByEnum string

// Set of constants representing the allowable values for ListBastionRestrictionsSortByEnum
const (
	ListBastionRestrictionsSortByTimecreated ListBastionRestrictionsSortByEnum = "timeCreated"
	ListBastionRestrictionsSortByName        ListBastionRestrictionsSortByEnum = "name"
)

var mappingListBastionRestrictionsSortByEnum = map[string]ListBastionRestrictionsSortByEnum{
	"timeCreated": ListBastionRestrictionsSortByTimecreated,
	"name":        ListBastionRestrictionsSortByName,
}

var mappingListBastionRestrictionsSortByEnumLowerCase = map[string]ListBastionRestrictionsSortByEnum{
	"timecreated": ListBastionRestrictionsSortByTimecreated,
	"name":        ListBastionRestrictionsSortByName,
}

// GetListBastionRestrictionsSortByEnumValues Enumerates the set of values for ListBastionRestrictionsSortByEnum
func GetListBastionRestrictionsSortByEnumValues() []ListBastionRestrictionsSortByEnum {
	values := make([]ListBastionRestrictionsSortByEnum, 0)
	for _, v := range mappingListBastionRestrictionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBastionRestrictionsSortByEnumStringValues Enumerates the set of values in String for ListBastionRestrictionsSortByEnum
func GetListBastionRestrictionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListBastionRestrictionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBastionRestrictionsSortByEnum(val string) (ListBastionRestrictionsSortByEnum, bool) {
	enum, ok := mappingListBastionRestrictionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
