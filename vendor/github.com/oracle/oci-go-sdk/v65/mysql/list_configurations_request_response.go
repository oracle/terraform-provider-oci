// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListConfigurationsRequest wrapper for the ListConfigurations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/ListConfigurations.go.html to see an example of how to use ListConfigurationsRequest.
type ListConfigurationsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The requested Configuration instance.
	ConfigurationId *string `mandatory:"false" contributesTo:"query" name:"configurationId"`

	// Configuration Lifecycle State
	LifecycleState ConfigurationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The requested Configuration types.
	Type []ListConfigurationsTypeEnum `contributesTo:"query" name:"type" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only the resource matching the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The requested Shape name.
	ShapeName *string `mandatory:"false" contributesTo:"query" name:"shapeName"`

	// The field to sort by. Only one sort order may be provided. Time fields are default ordered as descending. Display name is default ordered as ascending.
	SortBy ListConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use (ASC or DESC).
	SortOrder ListConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated list call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from
	// the previous list call. For information about pagination, see List
	// Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConfigurationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetConfigurationLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.Type {
		if _, ok := GetMappingListConfigurationsTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", val, strings.Join(GetListConfigurationsTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConfigurationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListConfigurationsResponse wrapper for the ListConfigurations operation
type ListConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ConfigurationSummary instances
	Items []ConfigurationSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConfigurationsTypeEnum Enum with underlying type: string
type ListConfigurationsTypeEnum string

// Set of constants representing the allowable values for ListConfigurationsTypeEnum
const (
	ListConfigurationsTypeDefault ListConfigurationsTypeEnum = "DEFAULT"
	ListConfigurationsTypeCustom  ListConfigurationsTypeEnum = "CUSTOM"
)

var mappingListConfigurationsTypeEnum = map[string]ListConfigurationsTypeEnum{
	"DEFAULT": ListConfigurationsTypeDefault,
	"CUSTOM":  ListConfigurationsTypeCustom,
}

var mappingListConfigurationsTypeEnumLowerCase = map[string]ListConfigurationsTypeEnum{
	"default": ListConfigurationsTypeDefault,
	"custom":  ListConfigurationsTypeCustom,
}

// GetListConfigurationsTypeEnumValues Enumerates the set of values for ListConfigurationsTypeEnum
func GetListConfigurationsTypeEnumValues() []ListConfigurationsTypeEnum {
	values := make([]ListConfigurationsTypeEnum, 0)
	for _, v := range mappingListConfigurationsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListConfigurationsTypeEnumStringValues Enumerates the set of values in String for ListConfigurationsTypeEnum
func GetListConfigurationsTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"CUSTOM",
	}
}

// GetMappingListConfigurationsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConfigurationsTypeEnum(val string) (ListConfigurationsTypeEnum, bool) {
	enum, ok := mappingListConfigurationsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConfigurationsSortByEnum Enum with underlying type: string
type ListConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListConfigurationsSortByEnum
const (
	ListConfigurationsSortByDisplayname ListConfigurationsSortByEnum = "displayName"
	ListConfigurationsSortByShapename   ListConfigurationsSortByEnum = "shapeName"
	ListConfigurationsSortByTimecreated ListConfigurationsSortByEnum = "timeCreated"
	ListConfigurationsSortByTimeupdated ListConfigurationsSortByEnum = "timeUpdated"
)

var mappingListConfigurationsSortByEnum = map[string]ListConfigurationsSortByEnum{
	"displayName": ListConfigurationsSortByDisplayname,
	"shapeName":   ListConfigurationsSortByShapename,
	"timeCreated": ListConfigurationsSortByTimecreated,
	"timeUpdated": ListConfigurationsSortByTimeupdated,
}

var mappingListConfigurationsSortByEnumLowerCase = map[string]ListConfigurationsSortByEnum{
	"displayname": ListConfigurationsSortByDisplayname,
	"shapename":   ListConfigurationsSortByShapename,
	"timecreated": ListConfigurationsSortByTimecreated,
	"timeupdated": ListConfigurationsSortByTimeupdated,
}

// GetListConfigurationsSortByEnumValues Enumerates the set of values for ListConfigurationsSortByEnum
func GetListConfigurationsSortByEnumValues() []ListConfigurationsSortByEnum {
	values := make([]ListConfigurationsSortByEnum, 0)
	for _, v := range mappingListConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListConfigurationsSortByEnum
func GetListConfigurationsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"shapeName",
		"timeCreated",
		"timeUpdated",
	}
}

// GetMappingListConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConfigurationsSortByEnum(val string) (ListConfigurationsSortByEnum, bool) {
	enum, ok := mappingListConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConfigurationsSortOrderEnum Enum with underlying type: string
type ListConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListConfigurationsSortOrderEnum
const (
	ListConfigurationsSortOrderAsc  ListConfigurationsSortOrderEnum = "ASC"
	ListConfigurationsSortOrderDesc ListConfigurationsSortOrderEnum = "DESC"
)

var mappingListConfigurationsSortOrderEnum = map[string]ListConfigurationsSortOrderEnum{
	"ASC":  ListConfigurationsSortOrderAsc,
	"DESC": ListConfigurationsSortOrderDesc,
}

var mappingListConfigurationsSortOrderEnumLowerCase = map[string]ListConfigurationsSortOrderEnum{
	"asc":  ListConfigurationsSortOrderAsc,
	"desc": ListConfigurationsSortOrderDesc,
}

// GetListConfigurationsSortOrderEnumValues Enumerates the set of values for ListConfigurationsSortOrderEnum
func GetListConfigurationsSortOrderEnumValues() []ListConfigurationsSortOrderEnum {
	values := make([]ListConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListConfigurationsSortOrderEnum
func GetListConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConfigurationsSortOrderEnum(val string) (ListConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
