// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPropertiesRequest wrapper for the ListProperties operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListProperties.go.html to see an example of how to use ListPropertiesRequest.
type ListPropertiesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState PropertyLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique Property identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources their scope matches the given lifecycleState.
	Scope ListPropertiesScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListPropertiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListPropertiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPropertiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPropertiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPropertiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPropertiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPropertiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPropertyLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPropertyLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPropertiesScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetListPropertiesScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPropertiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPropertiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPropertiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPropertiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPropertiesResponse wrapper for the ListProperties operation
type ListPropertiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PropertyCollection instances
	PropertyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPropertiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPropertiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPropertiesScopeEnum Enum with underlying type: string
type ListPropertiesScopeEnum string

// Set of constants representing the allowable values for ListPropertiesScopeEnum
const (
	ListPropertiesScopeTaxonomy       ListPropertiesScopeEnum = "TAXONOMY"
	ListPropertiesScopePlatformConfig ListPropertiesScopeEnum = "PLATFORM_CONFIG"
)

var mappingListPropertiesScopeEnum = map[string]ListPropertiesScopeEnum{
	"TAXONOMY":        ListPropertiesScopeTaxonomy,
	"PLATFORM_CONFIG": ListPropertiesScopePlatformConfig,
}

var mappingListPropertiesScopeEnumLowerCase = map[string]ListPropertiesScopeEnum{
	"taxonomy":        ListPropertiesScopeTaxonomy,
	"platform_config": ListPropertiesScopePlatformConfig,
}

// GetListPropertiesScopeEnumValues Enumerates the set of values for ListPropertiesScopeEnum
func GetListPropertiesScopeEnumValues() []ListPropertiesScopeEnum {
	values := make([]ListPropertiesScopeEnum, 0)
	for _, v := range mappingListPropertiesScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetListPropertiesScopeEnumStringValues Enumerates the set of values in String for ListPropertiesScopeEnum
func GetListPropertiesScopeEnumStringValues() []string {
	return []string{
		"TAXONOMY",
		"PLATFORM_CONFIG",
	}
}

// GetMappingListPropertiesScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPropertiesScopeEnum(val string) (ListPropertiesScopeEnum, bool) {
	enum, ok := mappingListPropertiesScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPropertiesSortOrderEnum Enum with underlying type: string
type ListPropertiesSortOrderEnum string

// Set of constants representing the allowable values for ListPropertiesSortOrderEnum
const (
	ListPropertiesSortOrderAsc  ListPropertiesSortOrderEnum = "ASC"
	ListPropertiesSortOrderDesc ListPropertiesSortOrderEnum = "DESC"
)

var mappingListPropertiesSortOrderEnum = map[string]ListPropertiesSortOrderEnum{
	"ASC":  ListPropertiesSortOrderAsc,
	"DESC": ListPropertiesSortOrderDesc,
}

var mappingListPropertiesSortOrderEnumLowerCase = map[string]ListPropertiesSortOrderEnum{
	"asc":  ListPropertiesSortOrderAsc,
	"desc": ListPropertiesSortOrderDesc,
}

// GetListPropertiesSortOrderEnumValues Enumerates the set of values for ListPropertiesSortOrderEnum
func GetListPropertiesSortOrderEnumValues() []ListPropertiesSortOrderEnum {
	values := make([]ListPropertiesSortOrderEnum, 0)
	for _, v := range mappingListPropertiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPropertiesSortOrderEnumStringValues Enumerates the set of values in String for ListPropertiesSortOrderEnum
func GetListPropertiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPropertiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPropertiesSortOrderEnum(val string) (ListPropertiesSortOrderEnum, bool) {
	enum, ok := mappingListPropertiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPropertiesSortByEnum Enum with underlying type: string
type ListPropertiesSortByEnum string

// Set of constants representing the allowable values for ListPropertiesSortByEnum
const (
	ListPropertiesSortByTimecreated ListPropertiesSortByEnum = "timeCreated"
	ListPropertiesSortByDisplayname ListPropertiesSortByEnum = "displayName"
)

var mappingListPropertiesSortByEnum = map[string]ListPropertiesSortByEnum{
	"timeCreated": ListPropertiesSortByTimecreated,
	"displayName": ListPropertiesSortByDisplayname,
}

var mappingListPropertiesSortByEnumLowerCase = map[string]ListPropertiesSortByEnum{
	"timecreated": ListPropertiesSortByTimecreated,
	"displayname": ListPropertiesSortByDisplayname,
}

// GetListPropertiesSortByEnumValues Enumerates the set of values for ListPropertiesSortByEnum
func GetListPropertiesSortByEnumValues() []ListPropertiesSortByEnum {
	values := make([]ListPropertiesSortByEnum, 0)
	for _, v := range mappingListPropertiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPropertiesSortByEnumStringValues Enumerates the set of values in String for ListPropertiesSortByEnum
func GetListPropertiesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPropertiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPropertiesSortByEnum(val string) (ListPropertiesSortByEnum, bool) {
	enum, ok := mappingListPropertiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
