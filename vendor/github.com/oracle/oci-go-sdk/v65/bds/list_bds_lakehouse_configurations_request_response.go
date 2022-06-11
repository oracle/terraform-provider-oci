// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBdsLakehouseConfigurationsRequest wrapper for the ListBdsLakehouseConfigurations operation
type ListBdsLakehouseConfigurationsRequest struct {

	// The OCID of the cluster.
	BdsInstanceId *string `mandatory:"true" contributesTo:"path" name:"bdsInstanceId"`

	// The lakehouse configuration ID.
	LakehouseConfigId *string `mandatory:"false" contributesTo:"query" name:"lakehouseConfigId"`

	// The current state of the lakehouse in the lakehouse configuration lifecycle.
	LifecycleState BdsLakehouseConfigurationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The ID of the API key that is associated with the lakehouse.
	BdsApiKeyId *string `mandatory:"false" contributesTo:"query" name:"bdsApiKeyId"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListBdsLakehouseConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListBdsLakehouseConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBdsLakehouseConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBdsLakehouseConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBdsLakehouseConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBdsLakehouseConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBdsLakehouseConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsLakehouseConfigurationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBdsLakehouseConfigurationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBdsLakehouseConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBdsLakehouseConfigurationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBdsLakehouseConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBdsLakehouseConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBdsLakehouseConfigurationsResponse wrapper for the ListBdsLakehouseConfigurations operation
type ListBdsLakehouseConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []BdsLakehouseConfigurationSummary instances
	Items []BdsLakehouseConfigurationSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBdsLakehouseConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBdsLakehouseConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBdsLakehouseConfigurationsSortByEnum Enum with underlying type: string
type ListBdsLakehouseConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListBdsLakehouseConfigurationsSortByEnum
const (
	ListBdsLakehouseConfigurationsSortByTimecreated ListBdsLakehouseConfigurationsSortByEnum = "timeCreated"
	ListBdsLakehouseConfigurationsSortByDisplayname ListBdsLakehouseConfigurationsSortByEnum = "displayName"
)

var mappingListBdsLakehouseConfigurationsSortByEnum = map[string]ListBdsLakehouseConfigurationsSortByEnum{
	"timeCreated": ListBdsLakehouseConfigurationsSortByTimecreated,
	"displayName": ListBdsLakehouseConfigurationsSortByDisplayname,
}

var mappingListBdsLakehouseConfigurationsSortByEnumLowerCase = map[string]ListBdsLakehouseConfigurationsSortByEnum{
	"timecreated": ListBdsLakehouseConfigurationsSortByTimecreated,
	"displayname": ListBdsLakehouseConfigurationsSortByDisplayname,
}

// GetListBdsLakehouseConfigurationsSortByEnumValues Enumerates the set of values for ListBdsLakehouseConfigurationsSortByEnum
func GetListBdsLakehouseConfigurationsSortByEnumValues() []ListBdsLakehouseConfigurationsSortByEnum {
	values := make([]ListBdsLakehouseConfigurationsSortByEnum, 0)
	for _, v := range mappingListBdsLakehouseConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBdsLakehouseConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListBdsLakehouseConfigurationsSortByEnum
func GetListBdsLakehouseConfigurationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListBdsLakehouseConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBdsLakehouseConfigurationsSortByEnum(val string) (ListBdsLakehouseConfigurationsSortByEnum, bool) {
	enum, ok := mappingListBdsLakehouseConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBdsLakehouseConfigurationsSortOrderEnum Enum with underlying type: string
type ListBdsLakehouseConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListBdsLakehouseConfigurationsSortOrderEnum
const (
	ListBdsLakehouseConfigurationsSortOrderAsc  ListBdsLakehouseConfigurationsSortOrderEnum = "ASC"
	ListBdsLakehouseConfigurationsSortOrderDesc ListBdsLakehouseConfigurationsSortOrderEnum = "DESC"
)

var mappingListBdsLakehouseConfigurationsSortOrderEnum = map[string]ListBdsLakehouseConfigurationsSortOrderEnum{
	"ASC":  ListBdsLakehouseConfigurationsSortOrderAsc,
	"DESC": ListBdsLakehouseConfigurationsSortOrderDesc,
}

var mappingListBdsLakehouseConfigurationsSortOrderEnumLowerCase = map[string]ListBdsLakehouseConfigurationsSortOrderEnum{
	"asc":  ListBdsLakehouseConfigurationsSortOrderAsc,
	"desc": ListBdsLakehouseConfigurationsSortOrderDesc,
}

// GetListBdsLakehouseConfigurationsSortOrderEnumValues Enumerates the set of values for ListBdsLakehouseConfigurationsSortOrderEnum
func GetListBdsLakehouseConfigurationsSortOrderEnumValues() []ListBdsLakehouseConfigurationsSortOrderEnum {
	values := make([]ListBdsLakehouseConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListBdsLakehouseConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBdsLakehouseConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListBdsLakehouseConfigurationsSortOrderEnum
func GetListBdsLakehouseConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBdsLakehouseConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBdsLakehouseConfigurationsSortOrderEnum(val string) (ListBdsLakehouseConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListBdsLakehouseConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
