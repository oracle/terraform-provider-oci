// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDigitalAssistantParametersRequest wrapper for the ListDigitalAssistantParameters operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListDigitalAssistantParameters.go.html to see an example of how to use ListDigitalAssistantParametersRequest.
type ListDigitalAssistantParametersRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// Unique Digital Assistant identifier.
	DigitalAssistantId *string `mandatory:"true" contributesTo:"path" name:"digitalAssistantId"`

	// List only Parameters with this name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// List only the resources that are in this lifecycle state.
	LifecycleState ListDigitalAssistantParametersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListDigitalAssistantParametersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `name`.
	// The default sort order is ascending.
	SortBy ListDigitalAssistantParametersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDigitalAssistantParametersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDigitalAssistantParametersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDigitalAssistantParametersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDigitalAssistantParametersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDigitalAssistantParametersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDigitalAssistantParametersLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDigitalAssistantParametersLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDigitalAssistantParametersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDigitalAssistantParametersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDigitalAssistantParametersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDigitalAssistantParametersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDigitalAssistantParametersResponse wrapper for the ListDigitalAssistantParameters operation
type ListDigitalAssistantParametersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DigitalAssistantParameterCollection instances
	DigitalAssistantParameterCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// When you are paging through a list, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the
	// `page` query parameter for the subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of results that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListDigitalAssistantParametersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDigitalAssistantParametersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDigitalAssistantParametersLifecycleStateEnum Enum with underlying type: string
type ListDigitalAssistantParametersLifecycleStateEnum string

// Set of constants representing the allowable values for ListDigitalAssistantParametersLifecycleStateEnum
const (
	ListDigitalAssistantParametersLifecycleStateCreating ListDigitalAssistantParametersLifecycleStateEnum = "CREATING"
	ListDigitalAssistantParametersLifecycleStateUpdating ListDigitalAssistantParametersLifecycleStateEnum = "UPDATING"
	ListDigitalAssistantParametersLifecycleStateActive   ListDigitalAssistantParametersLifecycleStateEnum = "ACTIVE"
	ListDigitalAssistantParametersLifecycleStateInactive ListDigitalAssistantParametersLifecycleStateEnum = "INACTIVE"
	ListDigitalAssistantParametersLifecycleStateDeleting ListDigitalAssistantParametersLifecycleStateEnum = "DELETING"
	ListDigitalAssistantParametersLifecycleStateDeleted  ListDigitalAssistantParametersLifecycleStateEnum = "DELETED"
	ListDigitalAssistantParametersLifecycleStateFailed   ListDigitalAssistantParametersLifecycleStateEnum = "FAILED"
)

var mappingListDigitalAssistantParametersLifecycleStateEnum = map[string]ListDigitalAssistantParametersLifecycleStateEnum{
	"CREATING": ListDigitalAssistantParametersLifecycleStateCreating,
	"UPDATING": ListDigitalAssistantParametersLifecycleStateUpdating,
	"ACTIVE":   ListDigitalAssistantParametersLifecycleStateActive,
	"INACTIVE": ListDigitalAssistantParametersLifecycleStateInactive,
	"DELETING": ListDigitalAssistantParametersLifecycleStateDeleting,
	"DELETED":  ListDigitalAssistantParametersLifecycleStateDeleted,
	"FAILED":   ListDigitalAssistantParametersLifecycleStateFailed,
}

var mappingListDigitalAssistantParametersLifecycleStateEnumLowerCase = map[string]ListDigitalAssistantParametersLifecycleStateEnum{
	"creating": ListDigitalAssistantParametersLifecycleStateCreating,
	"updating": ListDigitalAssistantParametersLifecycleStateUpdating,
	"active":   ListDigitalAssistantParametersLifecycleStateActive,
	"inactive": ListDigitalAssistantParametersLifecycleStateInactive,
	"deleting": ListDigitalAssistantParametersLifecycleStateDeleting,
	"deleted":  ListDigitalAssistantParametersLifecycleStateDeleted,
	"failed":   ListDigitalAssistantParametersLifecycleStateFailed,
}

// GetListDigitalAssistantParametersLifecycleStateEnumValues Enumerates the set of values for ListDigitalAssistantParametersLifecycleStateEnum
func GetListDigitalAssistantParametersLifecycleStateEnumValues() []ListDigitalAssistantParametersLifecycleStateEnum {
	values := make([]ListDigitalAssistantParametersLifecycleStateEnum, 0)
	for _, v := range mappingListDigitalAssistantParametersLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalAssistantParametersLifecycleStateEnumStringValues Enumerates the set of values in String for ListDigitalAssistantParametersLifecycleStateEnum
func GetListDigitalAssistantParametersLifecycleStateEnumStringValues() []string {
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

// GetMappingListDigitalAssistantParametersLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalAssistantParametersLifecycleStateEnum(val string) (ListDigitalAssistantParametersLifecycleStateEnum, bool) {
	enum, ok := mappingListDigitalAssistantParametersLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDigitalAssistantParametersSortOrderEnum Enum with underlying type: string
type ListDigitalAssistantParametersSortOrderEnum string

// Set of constants representing the allowable values for ListDigitalAssistantParametersSortOrderEnum
const (
	ListDigitalAssistantParametersSortOrderAsc  ListDigitalAssistantParametersSortOrderEnum = "ASC"
	ListDigitalAssistantParametersSortOrderDesc ListDigitalAssistantParametersSortOrderEnum = "DESC"
)

var mappingListDigitalAssistantParametersSortOrderEnum = map[string]ListDigitalAssistantParametersSortOrderEnum{
	"ASC":  ListDigitalAssistantParametersSortOrderAsc,
	"DESC": ListDigitalAssistantParametersSortOrderDesc,
}

var mappingListDigitalAssistantParametersSortOrderEnumLowerCase = map[string]ListDigitalAssistantParametersSortOrderEnum{
	"asc":  ListDigitalAssistantParametersSortOrderAsc,
	"desc": ListDigitalAssistantParametersSortOrderDesc,
}

// GetListDigitalAssistantParametersSortOrderEnumValues Enumerates the set of values for ListDigitalAssistantParametersSortOrderEnum
func GetListDigitalAssistantParametersSortOrderEnumValues() []ListDigitalAssistantParametersSortOrderEnum {
	values := make([]ListDigitalAssistantParametersSortOrderEnum, 0)
	for _, v := range mappingListDigitalAssistantParametersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalAssistantParametersSortOrderEnumStringValues Enumerates the set of values in String for ListDigitalAssistantParametersSortOrderEnum
func GetListDigitalAssistantParametersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDigitalAssistantParametersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalAssistantParametersSortOrderEnum(val string) (ListDigitalAssistantParametersSortOrderEnum, bool) {
	enum, ok := mappingListDigitalAssistantParametersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDigitalAssistantParametersSortByEnum Enum with underlying type: string
type ListDigitalAssistantParametersSortByEnum string

// Set of constants representing the allowable values for ListDigitalAssistantParametersSortByEnum
const (
	ListDigitalAssistantParametersSortByName        ListDigitalAssistantParametersSortByEnum = "name"
	ListDigitalAssistantParametersSortByDisplayname ListDigitalAssistantParametersSortByEnum = "displayName"
	ListDigitalAssistantParametersSortByType        ListDigitalAssistantParametersSortByEnum = "type"
)

var mappingListDigitalAssistantParametersSortByEnum = map[string]ListDigitalAssistantParametersSortByEnum{
	"name":        ListDigitalAssistantParametersSortByName,
	"displayName": ListDigitalAssistantParametersSortByDisplayname,
	"type":        ListDigitalAssistantParametersSortByType,
}

var mappingListDigitalAssistantParametersSortByEnumLowerCase = map[string]ListDigitalAssistantParametersSortByEnum{
	"name":        ListDigitalAssistantParametersSortByName,
	"displayname": ListDigitalAssistantParametersSortByDisplayname,
	"type":        ListDigitalAssistantParametersSortByType,
}

// GetListDigitalAssistantParametersSortByEnumValues Enumerates the set of values for ListDigitalAssistantParametersSortByEnum
func GetListDigitalAssistantParametersSortByEnumValues() []ListDigitalAssistantParametersSortByEnum {
	values := make([]ListDigitalAssistantParametersSortByEnum, 0)
	for _, v := range mappingListDigitalAssistantParametersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalAssistantParametersSortByEnumStringValues Enumerates the set of values in String for ListDigitalAssistantParametersSortByEnum
func GetListDigitalAssistantParametersSortByEnumStringValues() []string {
	return []string{
		"name",
		"displayName",
		"type",
	}
}

// GetMappingListDigitalAssistantParametersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalAssistantParametersSortByEnum(val string) (ListDigitalAssistantParametersSortByEnum, bool) {
	enum, ok := mappingListDigitalAssistantParametersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
