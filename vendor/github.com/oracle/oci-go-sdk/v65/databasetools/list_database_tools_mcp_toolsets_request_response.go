// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasetools

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDatabaseToolsMcpToolsetsRequest wrapper for the ListDatabaseToolsMcpToolsets operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsMcpToolsets.go.html to see an example of how to use ListDatabaseToolsMcpToolsetsRequest.
type ListDatabaseToolsMcpToolsetsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources only when their `databaseToolsMcpToolsetLifecycleState` matches the specified `databaseToolsMcpToolsetLifecycleState`.
	LifecycleState ListDatabaseToolsMcpToolsetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatabaseToolsMcpToolsetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDatabaseToolsMcpToolsetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources with one of the specified type values.
	Type []DatabaseToolsMcpToolsetTypeEnum `contributesTo:"query" name:"type" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources matching the specified `databaseToolsMcpServerId`.
	DatabaseToolsMcpServerId *string `mandatory:"false" contributesTo:"query" name:"databaseToolsMcpServerId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseToolsMcpToolsetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseToolsMcpToolsetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseToolsMcpToolsetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseToolsMcpToolsetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseToolsMcpToolsetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseToolsMcpToolsetsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDatabaseToolsMcpToolsetsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsMcpToolsetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseToolsMcpToolsetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsMcpToolsetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseToolsMcpToolsetsSortByEnumStringValues(), ",")))
	}
	for _, val := range request.Type {
		if _, ok := GetMappingDatabaseToolsMcpToolsetTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", val, strings.Join(GetDatabaseToolsMcpToolsetTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseToolsMcpToolsetsResponse wrapper for the ListDatabaseToolsMcpToolsets operation
type ListDatabaseToolsMcpToolsetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseToolsMcpToolsetCollection instances
	DatabaseToolsMcpToolsetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseToolsMcpToolsetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseToolsMcpToolsetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseToolsMcpToolsetsLifecycleStateEnum Enum with underlying type: string
type ListDatabaseToolsMcpToolsetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDatabaseToolsMcpToolsetsLifecycleStateEnum
const (
	ListDatabaseToolsMcpToolsetsLifecycleStateCreating ListDatabaseToolsMcpToolsetsLifecycleStateEnum = "CREATING"
	ListDatabaseToolsMcpToolsetsLifecycleStateUpdating ListDatabaseToolsMcpToolsetsLifecycleStateEnum = "UPDATING"
	ListDatabaseToolsMcpToolsetsLifecycleStateActive   ListDatabaseToolsMcpToolsetsLifecycleStateEnum = "ACTIVE"
	ListDatabaseToolsMcpToolsetsLifecycleStateDeleting ListDatabaseToolsMcpToolsetsLifecycleStateEnum = "DELETING"
	ListDatabaseToolsMcpToolsetsLifecycleStateDeleted  ListDatabaseToolsMcpToolsetsLifecycleStateEnum = "DELETED"
	ListDatabaseToolsMcpToolsetsLifecycleStateFailed   ListDatabaseToolsMcpToolsetsLifecycleStateEnum = "FAILED"
)

var mappingListDatabaseToolsMcpToolsetsLifecycleStateEnum = map[string]ListDatabaseToolsMcpToolsetsLifecycleStateEnum{
	"CREATING": ListDatabaseToolsMcpToolsetsLifecycleStateCreating,
	"UPDATING": ListDatabaseToolsMcpToolsetsLifecycleStateUpdating,
	"ACTIVE":   ListDatabaseToolsMcpToolsetsLifecycleStateActive,
	"DELETING": ListDatabaseToolsMcpToolsetsLifecycleStateDeleting,
	"DELETED":  ListDatabaseToolsMcpToolsetsLifecycleStateDeleted,
	"FAILED":   ListDatabaseToolsMcpToolsetsLifecycleStateFailed,
}

var mappingListDatabaseToolsMcpToolsetsLifecycleStateEnumLowerCase = map[string]ListDatabaseToolsMcpToolsetsLifecycleStateEnum{
	"creating": ListDatabaseToolsMcpToolsetsLifecycleStateCreating,
	"updating": ListDatabaseToolsMcpToolsetsLifecycleStateUpdating,
	"active":   ListDatabaseToolsMcpToolsetsLifecycleStateActive,
	"deleting": ListDatabaseToolsMcpToolsetsLifecycleStateDeleting,
	"deleted":  ListDatabaseToolsMcpToolsetsLifecycleStateDeleted,
	"failed":   ListDatabaseToolsMcpToolsetsLifecycleStateFailed,
}

// GetListDatabaseToolsMcpToolsetsLifecycleStateEnumValues Enumerates the set of values for ListDatabaseToolsMcpToolsetsLifecycleStateEnum
func GetListDatabaseToolsMcpToolsetsLifecycleStateEnumValues() []ListDatabaseToolsMcpToolsetsLifecycleStateEnum {
	values := make([]ListDatabaseToolsMcpToolsetsLifecycleStateEnum, 0)
	for _, v := range mappingListDatabaseToolsMcpToolsetsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsMcpToolsetsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDatabaseToolsMcpToolsetsLifecycleStateEnum
func GetListDatabaseToolsMcpToolsetsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListDatabaseToolsMcpToolsetsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsMcpToolsetsLifecycleStateEnum(val string) (ListDatabaseToolsMcpToolsetsLifecycleStateEnum, bool) {
	enum, ok := mappingListDatabaseToolsMcpToolsetsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsMcpToolsetsSortOrderEnum Enum with underlying type: string
type ListDatabaseToolsMcpToolsetsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseToolsMcpToolsetsSortOrderEnum
const (
	ListDatabaseToolsMcpToolsetsSortOrderAsc  ListDatabaseToolsMcpToolsetsSortOrderEnum = "ASC"
	ListDatabaseToolsMcpToolsetsSortOrderDesc ListDatabaseToolsMcpToolsetsSortOrderEnum = "DESC"
)

var mappingListDatabaseToolsMcpToolsetsSortOrderEnum = map[string]ListDatabaseToolsMcpToolsetsSortOrderEnum{
	"ASC":  ListDatabaseToolsMcpToolsetsSortOrderAsc,
	"DESC": ListDatabaseToolsMcpToolsetsSortOrderDesc,
}

var mappingListDatabaseToolsMcpToolsetsSortOrderEnumLowerCase = map[string]ListDatabaseToolsMcpToolsetsSortOrderEnum{
	"asc":  ListDatabaseToolsMcpToolsetsSortOrderAsc,
	"desc": ListDatabaseToolsMcpToolsetsSortOrderDesc,
}

// GetListDatabaseToolsMcpToolsetsSortOrderEnumValues Enumerates the set of values for ListDatabaseToolsMcpToolsetsSortOrderEnum
func GetListDatabaseToolsMcpToolsetsSortOrderEnumValues() []ListDatabaseToolsMcpToolsetsSortOrderEnum {
	values := make([]ListDatabaseToolsMcpToolsetsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseToolsMcpToolsetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsMcpToolsetsSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseToolsMcpToolsetsSortOrderEnum
func GetListDatabaseToolsMcpToolsetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseToolsMcpToolsetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsMcpToolsetsSortOrderEnum(val string) (ListDatabaseToolsMcpToolsetsSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseToolsMcpToolsetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsMcpToolsetsSortByEnum Enum with underlying type: string
type ListDatabaseToolsMcpToolsetsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseToolsMcpToolsetsSortByEnum
const (
	ListDatabaseToolsMcpToolsetsSortByTimecreated ListDatabaseToolsMcpToolsetsSortByEnum = "timeCreated"
	ListDatabaseToolsMcpToolsetsSortByDisplayname ListDatabaseToolsMcpToolsetsSortByEnum = "displayName"
)

var mappingListDatabaseToolsMcpToolsetsSortByEnum = map[string]ListDatabaseToolsMcpToolsetsSortByEnum{
	"timeCreated": ListDatabaseToolsMcpToolsetsSortByTimecreated,
	"displayName": ListDatabaseToolsMcpToolsetsSortByDisplayname,
}

var mappingListDatabaseToolsMcpToolsetsSortByEnumLowerCase = map[string]ListDatabaseToolsMcpToolsetsSortByEnum{
	"timecreated": ListDatabaseToolsMcpToolsetsSortByTimecreated,
	"displayname": ListDatabaseToolsMcpToolsetsSortByDisplayname,
}

// GetListDatabaseToolsMcpToolsetsSortByEnumValues Enumerates the set of values for ListDatabaseToolsMcpToolsetsSortByEnum
func GetListDatabaseToolsMcpToolsetsSortByEnumValues() []ListDatabaseToolsMcpToolsetsSortByEnum {
	values := make([]ListDatabaseToolsMcpToolsetsSortByEnum, 0)
	for _, v := range mappingListDatabaseToolsMcpToolsetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsMcpToolsetsSortByEnumStringValues Enumerates the set of values in String for ListDatabaseToolsMcpToolsetsSortByEnum
func GetListDatabaseToolsMcpToolsetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatabaseToolsMcpToolsetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsMcpToolsetsSortByEnum(val string) (ListDatabaseToolsMcpToolsetsSortByEnum, bool) {
	enum, ok := mappingListDatabaseToolsMcpToolsetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
