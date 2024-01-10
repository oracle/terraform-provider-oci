// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListConfigsRequest wrapper for the ListConfigs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListConfigs.go.html to see an example of how to use ListConfigsRequest.
type ListConfigsRequest struct {

	// The ID of the compartment in which data is listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only configuration items for a given config type.
	Type ConfigConfigTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The current state of the Config.
	LifecycleState ConfigLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListConfigsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for 'timeCreated' is descending.
	// Default order for 'displayName' and 'configType' is ascending.
	SortBy ListConfigsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConfigsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConfigsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConfigsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConfigsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListConfigsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConfigConfigTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetConfigConfigTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConfigLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetConfigLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConfigsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConfigsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConfigsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConfigsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListConfigsResponse wrapper for the ListConfigs operation
type ListConfigsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConfigCollection instances
	ConfigCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListConfigsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConfigsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConfigsSortOrderEnum Enum with underlying type: string
type ListConfigsSortOrderEnum string

// Set of constants representing the allowable values for ListConfigsSortOrderEnum
const (
	ListConfigsSortOrderAsc  ListConfigsSortOrderEnum = "ASC"
	ListConfigsSortOrderDesc ListConfigsSortOrderEnum = "DESC"
)

var mappingListConfigsSortOrderEnum = map[string]ListConfigsSortOrderEnum{
	"ASC":  ListConfigsSortOrderAsc,
	"DESC": ListConfigsSortOrderDesc,
}

var mappingListConfigsSortOrderEnumLowerCase = map[string]ListConfigsSortOrderEnum{
	"asc":  ListConfigsSortOrderAsc,
	"desc": ListConfigsSortOrderDesc,
}

// GetListConfigsSortOrderEnumValues Enumerates the set of values for ListConfigsSortOrderEnum
func GetListConfigsSortOrderEnumValues() []ListConfigsSortOrderEnum {
	values := make([]ListConfigsSortOrderEnum, 0)
	for _, v := range mappingListConfigsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListConfigsSortOrderEnumStringValues Enumerates the set of values in String for ListConfigsSortOrderEnum
func GetListConfigsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListConfigsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConfigsSortOrderEnum(val string) (ListConfigsSortOrderEnum, bool) {
	enum, ok := mappingListConfigsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConfigsSortByEnum Enum with underlying type: string
type ListConfigsSortByEnum string

// Set of constants representing the allowable values for ListConfigsSortByEnum
const (
	ListConfigsSortByTimecreated ListConfigsSortByEnum = "timeCreated"
	ListConfigsSortByConfigtype  ListConfigsSortByEnum = "configType"
	ListConfigsSortByDisplayname ListConfigsSortByEnum = "displayName"
)

var mappingListConfigsSortByEnum = map[string]ListConfigsSortByEnum{
	"timeCreated": ListConfigsSortByTimecreated,
	"configType":  ListConfigsSortByConfigtype,
	"displayName": ListConfigsSortByDisplayname,
}

var mappingListConfigsSortByEnumLowerCase = map[string]ListConfigsSortByEnum{
	"timecreated": ListConfigsSortByTimecreated,
	"configtype":  ListConfigsSortByConfigtype,
	"displayname": ListConfigsSortByDisplayname,
}

// GetListConfigsSortByEnumValues Enumerates the set of values for ListConfigsSortByEnum
func GetListConfigsSortByEnumValues() []ListConfigsSortByEnum {
	values := make([]ListConfigsSortByEnum, 0)
	for _, v := range mappingListConfigsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListConfigsSortByEnumStringValues Enumerates the set of values in String for ListConfigsSortByEnum
func GetListConfigsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"configType",
		"displayName",
	}
}

// GetMappingListConfigsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConfigsSortByEnum(val string) (ListConfigsSortByEnum, bool) {
	enum, ok := mappingListConfigsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
