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

// ListDatabaseToolsMcpServersRequest wrapper for the ListDatabaseToolsMcpServers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsMcpServers.go.html to see an example of how to use ListDatabaseToolsMcpServersRequest.
type ListDatabaseToolsMcpServersRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources only when their `databaseToolsMcpServerLifecycleState` matches the specified `databaseToolsMcpServerLifecycleState`.
	LifecycleState ListDatabaseToolsMcpServersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources when their `databaseToolsConnectionId` matches the specified `databaseToolsConnectionId`.
	DatabaseToolsConnectionId *string `mandatory:"false" contributesTo:"query" name:"databaseToolsConnectionId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related resource.
	RelatedResourceIdentifier *string `mandatory:"false" contributesTo:"query" name:"relatedResourceIdentifier"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatabaseToolsMcpServersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDatabaseToolsMcpServersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources with one of the specified type values.
	Type []DatabaseToolsMcpServerTypeEnum `contributesTo:"query" name:"type" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseToolsMcpServersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseToolsMcpServersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseToolsMcpServersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseToolsMcpServersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseToolsMcpServersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseToolsMcpServersLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDatabaseToolsMcpServersLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsMcpServersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseToolsMcpServersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsMcpServersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseToolsMcpServersSortByEnumStringValues(), ",")))
	}
	for _, val := range request.Type {
		if _, ok := GetMappingDatabaseToolsMcpServerTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", val, strings.Join(GetDatabaseToolsMcpServerTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseToolsMcpServersResponse wrapper for the ListDatabaseToolsMcpServers operation
type ListDatabaseToolsMcpServersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseToolsMcpServerCollection instances
	DatabaseToolsMcpServerCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseToolsMcpServersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseToolsMcpServersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseToolsMcpServersLifecycleStateEnum Enum with underlying type: string
type ListDatabaseToolsMcpServersLifecycleStateEnum string

// Set of constants representing the allowable values for ListDatabaseToolsMcpServersLifecycleStateEnum
const (
	ListDatabaseToolsMcpServersLifecycleStateCreating       ListDatabaseToolsMcpServersLifecycleStateEnum = "CREATING"
	ListDatabaseToolsMcpServersLifecycleStateUpdating       ListDatabaseToolsMcpServersLifecycleStateEnum = "UPDATING"
	ListDatabaseToolsMcpServersLifecycleStateActive         ListDatabaseToolsMcpServersLifecycleStateEnum = "ACTIVE"
	ListDatabaseToolsMcpServersLifecycleStateDeleting       ListDatabaseToolsMcpServersLifecycleStateEnum = "DELETING"
	ListDatabaseToolsMcpServersLifecycleStateDeleted        ListDatabaseToolsMcpServersLifecycleStateEnum = "DELETED"
	ListDatabaseToolsMcpServersLifecycleStateFailed         ListDatabaseToolsMcpServersLifecycleStateEnum = "FAILED"
	ListDatabaseToolsMcpServersLifecycleStateNeedsAttention ListDatabaseToolsMcpServersLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListDatabaseToolsMcpServersLifecycleStateEnum = map[string]ListDatabaseToolsMcpServersLifecycleStateEnum{
	"CREATING":        ListDatabaseToolsMcpServersLifecycleStateCreating,
	"UPDATING":        ListDatabaseToolsMcpServersLifecycleStateUpdating,
	"ACTIVE":          ListDatabaseToolsMcpServersLifecycleStateActive,
	"DELETING":        ListDatabaseToolsMcpServersLifecycleStateDeleting,
	"DELETED":         ListDatabaseToolsMcpServersLifecycleStateDeleted,
	"FAILED":          ListDatabaseToolsMcpServersLifecycleStateFailed,
	"NEEDS_ATTENTION": ListDatabaseToolsMcpServersLifecycleStateNeedsAttention,
}

var mappingListDatabaseToolsMcpServersLifecycleStateEnumLowerCase = map[string]ListDatabaseToolsMcpServersLifecycleStateEnum{
	"creating":        ListDatabaseToolsMcpServersLifecycleStateCreating,
	"updating":        ListDatabaseToolsMcpServersLifecycleStateUpdating,
	"active":          ListDatabaseToolsMcpServersLifecycleStateActive,
	"deleting":        ListDatabaseToolsMcpServersLifecycleStateDeleting,
	"deleted":         ListDatabaseToolsMcpServersLifecycleStateDeleted,
	"failed":          ListDatabaseToolsMcpServersLifecycleStateFailed,
	"needs_attention": ListDatabaseToolsMcpServersLifecycleStateNeedsAttention,
}

// GetListDatabaseToolsMcpServersLifecycleStateEnumValues Enumerates the set of values for ListDatabaseToolsMcpServersLifecycleStateEnum
func GetListDatabaseToolsMcpServersLifecycleStateEnumValues() []ListDatabaseToolsMcpServersLifecycleStateEnum {
	values := make([]ListDatabaseToolsMcpServersLifecycleStateEnum, 0)
	for _, v := range mappingListDatabaseToolsMcpServersLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsMcpServersLifecycleStateEnumStringValues Enumerates the set of values in String for ListDatabaseToolsMcpServersLifecycleStateEnum
func GetListDatabaseToolsMcpServersLifecycleStateEnumStringValues() []string {
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

// GetMappingListDatabaseToolsMcpServersLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsMcpServersLifecycleStateEnum(val string) (ListDatabaseToolsMcpServersLifecycleStateEnum, bool) {
	enum, ok := mappingListDatabaseToolsMcpServersLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsMcpServersSortOrderEnum Enum with underlying type: string
type ListDatabaseToolsMcpServersSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseToolsMcpServersSortOrderEnum
const (
	ListDatabaseToolsMcpServersSortOrderAsc  ListDatabaseToolsMcpServersSortOrderEnum = "ASC"
	ListDatabaseToolsMcpServersSortOrderDesc ListDatabaseToolsMcpServersSortOrderEnum = "DESC"
)

var mappingListDatabaseToolsMcpServersSortOrderEnum = map[string]ListDatabaseToolsMcpServersSortOrderEnum{
	"ASC":  ListDatabaseToolsMcpServersSortOrderAsc,
	"DESC": ListDatabaseToolsMcpServersSortOrderDesc,
}

var mappingListDatabaseToolsMcpServersSortOrderEnumLowerCase = map[string]ListDatabaseToolsMcpServersSortOrderEnum{
	"asc":  ListDatabaseToolsMcpServersSortOrderAsc,
	"desc": ListDatabaseToolsMcpServersSortOrderDesc,
}

// GetListDatabaseToolsMcpServersSortOrderEnumValues Enumerates the set of values for ListDatabaseToolsMcpServersSortOrderEnum
func GetListDatabaseToolsMcpServersSortOrderEnumValues() []ListDatabaseToolsMcpServersSortOrderEnum {
	values := make([]ListDatabaseToolsMcpServersSortOrderEnum, 0)
	for _, v := range mappingListDatabaseToolsMcpServersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsMcpServersSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseToolsMcpServersSortOrderEnum
func GetListDatabaseToolsMcpServersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseToolsMcpServersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsMcpServersSortOrderEnum(val string) (ListDatabaseToolsMcpServersSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseToolsMcpServersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsMcpServersSortByEnum Enum with underlying type: string
type ListDatabaseToolsMcpServersSortByEnum string

// Set of constants representing the allowable values for ListDatabaseToolsMcpServersSortByEnum
const (
	ListDatabaseToolsMcpServersSortByTimecreated ListDatabaseToolsMcpServersSortByEnum = "timeCreated"
	ListDatabaseToolsMcpServersSortByDisplayname ListDatabaseToolsMcpServersSortByEnum = "displayName"
)

var mappingListDatabaseToolsMcpServersSortByEnum = map[string]ListDatabaseToolsMcpServersSortByEnum{
	"timeCreated": ListDatabaseToolsMcpServersSortByTimecreated,
	"displayName": ListDatabaseToolsMcpServersSortByDisplayname,
}

var mappingListDatabaseToolsMcpServersSortByEnumLowerCase = map[string]ListDatabaseToolsMcpServersSortByEnum{
	"timecreated": ListDatabaseToolsMcpServersSortByTimecreated,
	"displayname": ListDatabaseToolsMcpServersSortByDisplayname,
}

// GetListDatabaseToolsMcpServersSortByEnumValues Enumerates the set of values for ListDatabaseToolsMcpServersSortByEnum
func GetListDatabaseToolsMcpServersSortByEnumValues() []ListDatabaseToolsMcpServersSortByEnum {
	values := make([]ListDatabaseToolsMcpServersSortByEnum, 0)
	for _, v := range mappingListDatabaseToolsMcpServersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsMcpServersSortByEnumStringValues Enumerates the set of values in String for ListDatabaseToolsMcpServersSortByEnum
func GetListDatabaseToolsMcpServersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatabaseToolsMcpServersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsMcpServersSortByEnum(val string) (ListDatabaseToolsMcpServersSortByEnum, bool) {
	enum, ok := mappingListDatabaseToolsMcpServersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
