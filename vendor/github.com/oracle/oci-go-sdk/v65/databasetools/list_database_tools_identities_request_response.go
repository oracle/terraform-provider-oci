// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasetools

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDatabaseToolsIdentitiesRequest wrapper for the ListDatabaseToolsIdentities operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsIdentities.go.html to see an example of how to use ListDatabaseToolsIdentitiesRequest.
type ListDatabaseToolsIdentitiesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources only when their `databaseToolsIdentityLifecycleState` matches the specified `databaseToolsIdentityLifecycleState`.
	LifecycleState ListDatabaseToolsIdentitiesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources when their `databaseToolsConnectionId` matches the specified `databaseToolsConnectionId`.
	DatabaseToolsConnectionId *string `mandatory:"false" contributesTo:"query" name:"databaseToolsConnectionId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatabaseToolsIdentitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDatabaseToolsIdentitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources with one of the specified type values.
	Type []IdentityTypeEnum `contributesTo:"query" name:"type" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseToolsIdentitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseToolsIdentitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseToolsIdentitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseToolsIdentitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseToolsIdentitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseToolsIdentitiesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDatabaseToolsIdentitiesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsIdentitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseToolsIdentitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsIdentitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseToolsIdentitiesSortByEnumStringValues(), ",")))
	}
	for _, val := range request.Type {
		if _, ok := GetMappingIdentityTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", val, strings.Join(GetIdentityTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseToolsIdentitiesResponse wrapper for the ListDatabaseToolsIdentities operation
type ListDatabaseToolsIdentitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseToolsIdentityCollection instances
	DatabaseToolsIdentityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseToolsIdentitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseToolsIdentitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseToolsIdentitiesLifecycleStateEnum Enum with underlying type: string
type ListDatabaseToolsIdentitiesLifecycleStateEnum string

// Set of constants representing the allowable values for ListDatabaseToolsIdentitiesLifecycleStateEnum
const (
	ListDatabaseToolsIdentitiesLifecycleStateCreating       ListDatabaseToolsIdentitiesLifecycleStateEnum = "CREATING"
	ListDatabaseToolsIdentitiesLifecycleStateUpdating       ListDatabaseToolsIdentitiesLifecycleStateEnum = "UPDATING"
	ListDatabaseToolsIdentitiesLifecycleStateActive         ListDatabaseToolsIdentitiesLifecycleStateEnum = "ACTIVE"
	ListDatabaseToolsIdentitiesLifecycleStateDeleting       ListDatabaseToolsIdentitiesLifecycleStateEnum = "DELETING"
	ListDatabaseToolsIdentitiesLifecycleStateDeleted        ListDatabaseToolsIdentitiesLifecycleStateEnum = "DELETED"
	ListDatabaseToolsIdentitiesLifecycleStateFailed         ListDatabaseToolsIdentitiesLifecycleStateEnum = "FAILED"
	ListDatabaseToolsIdentitiesLifecycleStateNeedsAttention ListDatabaseToolsIdentitiesLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListDatabaseToolsIdentitiesLifecycleStateEnum = map[string]ListDatabaseToolsIdentitiesLifecycleStateEnum{
	"CREATING":        ListDatabaseToolsIdentitiesLifecycleStateCreating,
	"UPDATING":        ListDatabaseToolsIdentitiesLifecycleStateUpdating,
	"ACTIVE":          ListDatabaseToolsIdentitiesLifecycleStateActive,
	"DELETING":        ListDatabaseToolsIdentitiesLifecycleStateDeleting,
	"DELETED":         ListDatabaseToolsIdentitiesLifecycleStateDeleted,
	"FAILED":          ListDatabaseToolsIdentitiesLifecycleStateFailed,
	"NEEDS_ATTENTION": ListDatabaseToolsIdentitiesLifecycleStateNeedsAttention,
}

var mappingListDatabaseToolsIdentitiesLifecycleStateEnumLowerCase = map[string]ListDatabaseToolsIdentitiesLifecycleStateEnum{
	"creating":        ListDatabaseToolsIdentitiesLifecycleStateCreating,
	"updating":        ListDatabaseToolsIdentitiesLifecycleStateUpdating,
	"active":          ListDatabaseToolsIdentitiesLifecycleStateActive,
	"deleting":        ListDatabaseToolsIdentitiesLifecycleStateDeleting,
	"deleted":         ListDatabaseToolsIdentitiesLifecycleStateDeleted,
	"failed":          ListDatabaseToolsIdentitiesLifecycleStateFailed,
	"needs_attention": ListDatabaseToolsIdentitiesLifecycleStateNeedsAttention,
}

// GetListDatabaseToolsIdentitiesLifecycleStateEnumValues Enumerates the set of values for ListDatabaseToolsIdentitiesLifecycleStateEnum
func GetListDatabaseToolsIdentitiesLifecycleStateEnumValues() []ListDatabaseToolsIdentitiesLifecycleStateEnum {
	values := make([]ListDatabaseToolsIdentitiesLifecycleStateEnum, 0)
	for _, v := range mappingListDatabaseToolsIdentitiesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsIdentitiesLifecycleStateEnumStringValues Enumerates the set of values in String for ListDatabaseToolsIdentitiesLifecycleStateEnum
func GetListDatabaseToolsIdentitiesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListDatabaseToolsIdentitiesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsIdentitiesLifecycleStateEnum(val string) (ListDatabaseToolsIdentitiesLifecycleStateEnum, bool) {
	enum, ok := mappingListDatabaseToolsIdentitiesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsIdentitiesSortOrderEnum Enum with underlying type: string
type ListDatabaseToolsIdentitiesSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseToolsIdentitiesSortOrderEnum
const (
	ListDatabaseToolsIdentitiesSortOrderAsc  ListDatabaseToolsIdentitiesSortOrderEnum = "ASC"
	ListDatabaseToolsIdentitiesSortOrderDesc ListDatabaseToolsIdentitiesSortOrderEnum = "DESC"
)

var mappingListDatabaseToolsIdentitiesSortOrderEnum = map[string]ListDatabaseToolsIdentitiesSortOrderEnum{
	"ASC":  ListDatabaseToolsIdentitiesSortOrderAsc,
	"DESC": ListDatabaseToolsIdentitiesSortOrderDesc,
}

var mappingListDatabaseToolsIdentitiesSortOrderEnumLowerCase = map[string]ListDatabaseToolsIdentitiesSortOrderEnum{
	"asc":  ListDatabaseToolsIdentitiesSortOrderAsc,
	"desc": ListDatabaseToolsIdentitiesSortOrderDesc,
}

// GetListDatabaseToolsIdentitiesSortOrderEnumValues Enumerates the set of values for ListDatabaseToolsIdentitiesSortOrderEnum
func GetListDatabaseToolsIdentitiesSortOrderEnumValues() []ListDatabaseToolsIdentitiesSortOrderEnum {
	values := make([]ListDatabaseToolsIdentitiesSortOrderEnum, 0)
	for _, v := range mappingListDatabaseToolsIdentitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsIdentitiesSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseToolsIdentitiesSortOrderEnum
func GetListDatabaseToolsIdentitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseToolsIdentitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsIdentitiesSortOrderEnum(val string) (ListDatabaseToolsIdentitiesSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseToolsIdentitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsIdentitiesSortByEnum Enum with underlying type: string
type ListDatabaseToolsIdentitiesSortByEnum string

// Set of constants representing the allowable values for ListDatabaseToolsIdentitiesSortByEnum
const (
	ListDatabaseToolsIdentitiesSortByTimecreated ListDatabaseToolsIdentitiesSortByEnum = "timeCreated"
	ListDatabaseToolsIdentitiesSortByDisplayname ListDatabaseToolsIdentitiesSortByEnum = "displayName"
)

var mappingListDatabaseToolsIdentitiesSortByEnum = map[string]ListDatabaseToolsIdentitiesSortByEnum{
	"timeCreated": ListDatabaseToolsIdentitiesSortByTimecreated,
	"displayName": ListDatabaseToolsIdentitiesSortByDisplayname,
}

var mappingListDatabaseToolsIdentitiesSortByEnumLowerCase = map[string]ListDatabaseToolsIdentitiesSortByEnum{
	"timecreated": ListDatabaseToolsIdentitiesSortByTimecreated,
	"displayname": ListDatabaseToolsIdentitiesSortByDisplayname,
}

// GetListDatabaseToolsIdentitiesSortByEnumValues Enumerates the set of values for ListDatabaseToolsIdentitiesSortByEnum
func GetListDatabaseToolsIdentitiesSortByEnumValues() []ListDatabaseToolsIdentitiesSortByEnum {
	values := make([]ListDatabaseToolsIdentitiesSortByEnum, 0)
	for _, v := range mappingListDatabaseToolsIdentitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsIdentitiesSortByEnumStringValues Enumerates the set of values in String for ListDatabaseToolsIdentitiesSortByEnum
func GetListDatabaseToolsIdentitiesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatabaseToolsIdentitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsIdentitiesSortByEnum(val string) (ListDatabaseToolsIdentitiesSortByEnum, bool) {
	enum, ok := mappingListDatabaseToolsIdentitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
