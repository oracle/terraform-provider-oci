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

// ListDatabaseToolsDatabaseApiGatewayConfigsRequest wrapper for the ListDatabaseToolsDatabaseApiGatewayConfigs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsDatabaseApiGatewayConfigs.go.html to see an example of how to use ListDatabaseToolsDatabaseApiGatewayConfigsRequest.
type ListDatabaseToolsDatabaseApiGatewayConfigsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources only when their `lifecycleState` matches the specified `lifecycleState`.
	LifecycleState ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources with one of the specified type values.
	Type []DatabaseApiGatewayConfigTypeEnum `contributesTo:"query" name:"type" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseToolsDatabaseApiGatewayConfigsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseToolsDatabaseApiGatewayConfigsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseToolsDatabaseApiGatewayConfigsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseToolsDatabaseApiGatewayConfigsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseToolsDatabaseApiGatewayConfigsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseToolsDatabaseApiGatewayConfigsSortByEnumStringValues(), ",")))
	}
	for _, val := range request.Type {
		if _, ok := GetMappingDatabaseApiGatewayConfigTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", val, strings.Join(GetDatabaseApiGatewayConfigTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseToolsDatabaseApiGatewayConfigsResponse wrapper for the ListDatabaseToolsDatabaseApiGatewayConfigs operation
type ListDatabaseToolsDatabaseApiGatewayConfigsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseToolsDatabaseApiGatewayConfigCollection instances
	DatabaseToolsDatabaseApiGatewayConfigCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseToolsDatabaseApiGatewayConfigsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseToolsDatabaseApiGatewayConfigsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum Enum with underlying type: string
type ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum
const (
	ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateActive  ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum = "ACTIVE"
	ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateDeleted ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum = "DELETED"
)

var mappingListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum = map[string]ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum{
	"ACTIVE":  ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateActive,
	"DELETED": ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateDeleted,
}

var mappingListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnumLowerCase = map[string]ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum{
	"active":  ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateActive,
	"deleted": ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateDeleted,
}

// GetListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnumValues Enumerates the set of values for ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnumValues() []ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum {
	values := make([]ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum, 0)
	for _, v := range mappingListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum(val string) (ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum, bool) {
	enum, ok := mappingListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum Enum with underlying type: string
type ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum
const (
	ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderAsc  ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum = "ASC"
	ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderDesc ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum = "DESC"
)

var mappingListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum = map[string]ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum{
	"ASC":  ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderAsc,
	"DESC": ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderDesc,
}

var mappingListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnumLowerCase = map[string]ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum{
	"asc":  ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderAsc,
	"desc": ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderDesc,
}

// GetListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnumValues Enumerates the set of values for ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnumValues() []ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum {
	values := make([]ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum(val string) (ListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseToolsDatabaseApiGatewayConfigsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum Enum with underlying type: string
type ListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum
const (
	ListDatabaseToolsDatabaseApiGatewayConfigsSortByTimecreated ListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum = "timeCreated"
	ListDatabaseToolsDatabaseApiGatewayConfigsSortByDisplayname ListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum = "displayName"
)

var mappingListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum = map[string]ListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum{
	"timeCreated": ListDatabaseToolsDatabaseApiGatewayConfigsSortByTimecreated,
	"displayName": ListDatabaseToolsDatabaseApiGatewayConfigsSortByDisplayname,
}

var mappingListDatabaseToolsDatabaseApiGatewayConfigsSortByEnumLowerCase = map[string]ListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum{
	"timecreated": ListDatabaseToolsDatabaseApiGatewayConfigsSortByTimecreated,
	"displayname": ListDatabaseToolsDatabaseApiGatewayConfigsSortByDisplayname,
}

// GetListDatabaseToolsDatabaseApiGatewayConfigsSortByEnumValues Enumerates the set of values for ListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigsSortByEnumValues() []ListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum {
	values := make([]ListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum, 0)
	for _, v := range mappingListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseToolsDatabaseApiGatewayConfigsSortByEnumStringValues Enumerates the set of values in String for ListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum
func GetListDatabaseToolsDatabaseApiGatewayConfigsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum(val string) (ListDatabaseToolsDatabaseApiGatewayConfigsSortByEnum, bool) {
	enum, ok := mappingListDatabaseToolsDatabaseApiGatewayConfigsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
