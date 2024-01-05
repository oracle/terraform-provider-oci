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

// ListDigitalAssistantsRequest wrapper for the ListDigitalAssistants operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListDigitalAssistants.go.html to see an example of how to use ListDigitalAssistantsRequest.
type ListDigitalAssistantsRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// Unique Digital Assistant identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// List only Bot resources with this category.
	Category *string `mandatory:"false" contributesTo:"query" name:"category"`

	// List only Bot resources with this name. Names are unique and may not change.
	// Example: `MySkill`
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// List only Bot resources with this version. Versions are unique and may not change.
	// Example: `1.0`
	Version *string `mandatory:"false" contributesTo:"query" name:"version"`

	// List only Bot resources with this namespace. Namespaces may not change.
	// Example: `MyNamespace`
	Namespace *string `mandatory:"false" contributesTo:"query" name:"namespace"`

	// List only Bot resources with this platform version.
	PlatformVersion *string `mandatory:"false" contributesTo:"query" name:"platformVersion"`

	// List only the resources that are in this lifecycle state.
	LifecycleState ListDigitalAssistantsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// List only Bot resources with this lifecycle details.
	LifecycleDetails *string `mandatory:"false" contributesTo:"query" name:"lifecycleDetails"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListDigitalAssistantsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `timeCreated`.
	// The default sort order for `timeCreated` and `timeUpdated` is descending.
	// For all other sort fields the default sort order is ascending.
	SortBy ListDigitalAssistantsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDigitalAssistantsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDigitalAssistantsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDigitalAssistantsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDigitalAssistantsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDigitalAssistantsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDigitalAssistantsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDigitalAssistantsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDigitalAssistantsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDigitalAssistantsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDigitalAssistantsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDigitalAssistantsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDigitalAssistantsResponse wrapper for the ListDigitalAssistants operation
type ListDigitalAssistantsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DigitalAssistantCollection instances
	DigitalAssistantCollection `presentIn:"body"`

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

func (response ListDigitalAssistantsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDigitalAssistantsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDigitalAssistantsLifecycleStateEnum Enum with underlying type: string
type ListDigitalAssistantsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDigitalAssistantsLifecycleStateEnum
const (
	ListDigitalAssistantsLifecycleStateCreating ListDigitalAssistantsLifecycleStateEnum = "CREATING"
	ListDigitalAssistantsLifecycleStateUpdating ListDigitalAssistantsLifecycleStateEnum = "UPDATING"
	ListDigitalAssistantsLifecycleStateActive   ListDigitalAssistantsLifecycleStateEnum = "ACTIVE"
	ListDigitalAssistantsLifecycleStateInactive ListDigitalAssistantsLifecycleStateEnum = "INACTIVE"
	ListDigitalAssistantsLifecycleStateDeleting ListDigitalAssistantsLifecycleStateEnum = "DELETING"
	ListDigitalAssistantsLifecycleStateDeleted  ListDigitalAssistantsLifecycleStateEnum = "DELETED"
	ListDigitalAssistantsLifecycleStateFailed   ListDigitalAssistantsLifecycleStateEnum = "FAILED"
)

var mappingListDigitalAssistantsLifecycleStateEnum = map[string]ListDigitalAssistantsLifecycleStateEnum{
	"CREATING": ListDigitalAssistantsLifecycleStateCreating,
	"UPDATING": ListDigitalAssistantsLifecycleStateUpdating,
	"ACTIVE":   ListDigitalAssistantsLifecycleStateActive,
	"INACTIVE": ListDigitalAssistantsLifecycleStateInactive,
	"DELETING": ListDigitalAssistantsLifecycleStateDeleting,
	"DELETED":  ListDigitalAssistantsLifecycleStateDeleted,
	"FAILED":   ListDigitalAssistantsLifecycleStateFailed,
}

var mappingListDigitalAssistantsLifecycleStateEnumLowerCase = map[string]ListDigitalAssistantsLifecycleStateEnum{
	"creating": ListDigitalAssistantsLifecycleStateCreating,
	"updating": ListDigitalAssistantsLifecycleStateUpdating,
	"active":   ListDigitalAssistantsLifecycleStateActive,
	"inactive": ListDigitalAssistantsLifecycleStateInactive,
	"deleting": ListDigitalAssistantsLifecycleStateDeleting,
	"deleted":  ListDigitalAssistantsLifecycleStateDeleted,
	"failed":   ListDigitalAssistantsLifecycleStateFailed,
}

// GetListDigitalAssistantsLifecycleStateEnumValues Enumerates the set of values for ListDigitalAssistantsLifecycleStateEnum
func GetListDigitalAssistantsLifecycleStateEnumValues() []ListDigitalAssistantsLifecycleStateEnum {
	values := make([]ListDigitalAssistantsLifecycleStateEnum, 0)
	for _, v := range mappingListDigitalAssistantsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalAssistantsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDigitalAssistantsLifecycleStateEnum
func GetListDigitalAssistantsLifecycleStateEnumStringValues() []string {
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

// GetMappingListDigitalAssistantsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalAssistantsLifecycleStateEnum(val string) (ListDigitalAssistantsLifecycleStateEnum, bool) {
	enum, ok := mappingListDigitalAssistantsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDigitalAssistantsSortOrderEnum Enum with underlying type: string
type ListDigitalAssistantsSortOrderEnum string

// Set of constants representing the allowable values for ListDigitalAssistantsSortOrderEnum
const (
	ListDigitalAssistantsSortOrderAsc  ListDigitalAssistantsSortOrderEnum = "ASC"
	ListDigitalAssistantsSortOrderDesc ListDigitalAssistantsSortOrderEnum = "DESC"
)

var mappingListDigitalAssistantsSortOrderEnum = map[string]ListDigitalAssistantsSortOrderEnum{
	"ASC":  ListDigitalAssistantsSortOrderAsc,
	"DESC": ListDigitalAssistantsSortOrderDesc,
}

var mappingListDigitalAssistantsSortOrderEnumLowerCase = map[string]ListDigitalAssistantsSortOrderEnum{
	"asc":  ListDigitalAssistantsSortOrderAsc,
	"desc": ListDigitalAssistantsSortOrderDesc,
}

// GetListDigitalAssistantsSortOrderEnumValues Enumerates the set of values for ListDigitalAssistantsSortOrderEnum
func GetListDigitalAssistantsSortOrderEnumValues() []ListDigitalAssistantsSortOrderEnum {
	values := make([]ListDigitalAssistantsSortOrderEnum, 0)
	for _, v := range mappingListDigitalAssistantsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalAssistantsSortOrderEnumStringValues Enumerates the set of values in String for ListDigitalAssistantsSortOrderEnum
func GetListDigitalAssistantsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDigitalAssistantsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalAssistantsSortOrderEnum(val string) (ListDigitalAssistantsSortOrderEnum, bool) {
	enum, ok := mappingListDigitalAssistantsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDigitalAssistantsSortByEnum Enum with underlying type: string
type ListDigitalAssistantsSortByEnum string

// Set of constants representing the allowable values for ListDigitalAssistantsSortByEnum
const (
	ListDigitalAssistantsSortByTimecreated ListDigitalAssistantsSortByEnum = "timeCreated"
	ListDigitalAssistantsSortByTimeupdated ListDigitalAssistantsSortByEnum = "timeUpdated"
	ListDigitalAssistantsSortByName        ListDigitalAssistantsSortByEnum = "name"
)

var mappingListDigitalAssistantsSortByEnum = map[string]ListDigitalAssistantsSortByEnum{
	"timeCreated": ListDigitalAssistantsSortByTimecreated,
	"timeUpdated": ListDigitalAssistantsSortByTimeupdated,
	"name":        ListDigitalAssistantsSortByName,
}

var mappingListDigitalAssistantsSortByEnumLowerCase = map[string]ListDigitalAssistantsSortByEnum{
	"timecreated": ListDigitalAssistantsSortByTimecreated,
	"timeupdated": ListDigitalAssistantsSortByTimeupdated,
	"name":        ListDigitalAssistantsSortByName,
}

// GetListDigitalAssistantsSortByEnumValues Enumerates the set of values for ListDigitalAssistantsSortByEnum
func GetListDigitalAssistantsSortByEnumValues() []ListDigitalAssistantsSortByEnum {
	values := make([]ListDigitalAssistantsSortByEnum, 0)
	for _, v := range mappingListDigitalAssistantsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDigitalAssistantsSortByEnumStringValues Enumerates the set of values in String for ListDigitalAssistantsSortByEnum
func GetListDigitalAssistantsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"name",
	}
}

// GetMappingListDigitalAssistantsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDigitalAssistantsSortByEnum(val string) (ListDigitalAssistantsSortByEnum, bool) {
	enum, ok := mappingListDigitalAssistantsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
