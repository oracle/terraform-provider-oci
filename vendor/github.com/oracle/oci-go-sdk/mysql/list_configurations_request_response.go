// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListConfigurationsRequest wrapper for the ListConfigurations operation
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
func (request ListConfigurationsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListConfigurationsResponse wrapper for the ListConfigurations operation
type ListConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ConfigurationSummary instances
	Items []ConfigurationSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a specific request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Opaque token representing the next page of results.
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

var mappingListConfigurationsType = map[string]ListConfigurationsTypeEnum{
	"DEFAULT": ListConfigurationsTypeDefault,
	"CUSTOM":  ListConfigurationsTypeCustom,
}

// GetListConfigurationsTypeEnumValues Enumerates the set of values for ListConfigurationsTypeEnum
func GetListConfigurationsTypeEnumValues() []ListConfigurationsTypeEnum {
	values := make([]ListConfigurationsTypeEnum, 0)
	for _, v := range mappingListConfigurationsType {
		values = append(values, v)
	}
	return values
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

var mappingListConfigurationsSortBy = map[string]ListConfigurationsSortByEnum{
	"displayName": ListConfigurationsSortByDisplayname,
	"shapeName":   ListConfigurationsSortByShapename,
	"timeCreated": ListConfigurationsSortByTimecreated,
	"timeUpdated": ListConfigurationsSortByTimeupdated,
}

// GetListConfigurationsSortByEnumValues Enumerates the set of values for ListConfigurationsSortByEnum
func GetListConfigurationsSortByEnumValues() []ListConfigurationsSortByEnum {
	values := make([]ListConfigurationsSortByEnum, 0)
	for _, v := range mappingListConfigurationsSortBy {
		values = append(values, v)
	}
	return values
}

// ListConfigurationsSortOrderEnum Enum with underlying type: string
type ListConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListConfigurationsSortOrderEnum
const (
	ListConfigurationsSortOrderAsc  ListConfigurationsSortOrderEnum = "ASC"
	ListConfigurationsSortOrderDesc ListConfigurationsSortOrderEnum = "DESC"
)

var mappingListConfigurationsSortOrder = map[string]ListConfigurationsSortOrderEnum{
	"ASC":  ListConfigurationsSortOrderAsc,
	"DESC": ListConfigurationsSortOrderDesc,
}

// GetListConfigurationsSortOrderEnumValues Enumerates the set of values for ListConfigurationsSortOrderEnum
func GetListConfigurationsSortOrderEnumValues() []ListConfigurationsSortOrderEnum {
	values := make([]ListConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListConfigurationsSortOrder {
		values = append(values, v)
	}
	return values
}
