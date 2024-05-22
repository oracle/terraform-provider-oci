// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListNodeReplaceConfigurationsRequest wrapper for the ListNodeReplaceConfigurations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListNodeReplaceConfigurations.go.html to see an example of how to use ListNodeReplaceConfigurationsRequest.
type ListNodeReplaceConfigurationsRequest struct {

	// The OCID of the cluster.
	BdsInstanceId *string `mandatory:"true" contributesTo:"path" name:"bdsInstanceId"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListNodeReplaceConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListNodeReplaceConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The state of the NodeReplaceConfiguration.
	LifecycleState NodeReplaceConfigurationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNodeReplaceConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNodeReplaceConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNodeReplaceConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNodeReplaceConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNodeReplaceConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNodeReplaceConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNodeReplaceConfigurationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNodeReplaceConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNodeReplaceConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNodeReplaceConfigurationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetNodeReplaceConfigurationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNodeReplaceConfigurationsResponse wrapper for the ListNodeReplaceConfigurations operation
type ListNodeReplaceConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []NodeReplaceConfigurationSummary instances
	Items []NodeReplaceConfigurationSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListNodeReplaceConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNodeReplaceConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNodeReplaceConfigurationsSortByEnum Enum with underlying type: string
type ListNodeReplaceConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListNodeReplaceConfigurationsSortByEnum
const (
	ListNodeReplaceConfigurationsSortByTimecreated ListNodeReplaceConfigurationsSortByEnum = "timeCreated"
	ListNodeReplaceConfigurationsSortByDisplayname ListNodeReplaceConfigurationsSortByEnum = "displayName"
)

var mappingListNodeReplaceConfigurationsSortByEnum = map[string]ListNodeReplaceConfigurationsSortByEnum{
	"timeCreated": ListNodeReplaceConfigurationsSortByTimecreated,
	"displayName": ListNodeReplaceConfigurationsSortByDisplayname,
}

var mappingListNodeReplaceConfigurationsSortByEnumLowerCase = map[string]ListNodeReplaceConfigurationsSortByEnum{
	"timecreated": ListNodeReplaceConfigurationsSortByTimecreated,
	"displayname": ListNodeReplaceConfigurationsSortByDisplayname,
}

// GetListNodeReplaceConfigurationsSortByEnumValues Enumerates the set of values for ListNodeReplaceConfigurationsSortByEnum
func GetListNodeReplaceConfigurationsSortByEnumValues() []ListNodeReplaceConfigurationsSortByEnum {
	values := make([]ListNodeReplaceConfigurationsSortByEnum, 0)
	for _, v := range mappingListNodeReplaceConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNodeReplaceConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListNodeReplaceConfigurationsSortByEnum
func GetListNodeReplaceConfigurationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListNodeReplaceConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNodeReplaceConfigurationsSortByEnum(val string) (ListNodeReplaceConfigurationsSortByEnum, bool) {
	enum, ok := mappingListNodeReplaceConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNodeReplaceConfigurationsSortOrderEnum Enum with underlying type: string
type ListNodeReplaceConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListNodeReplaceConfigurationsSortOrderEnum
const (
	ListNodeReplaceConfigurationsSortOrderAsc  ListNodeReplaceConfigurationsSortOrderEnum = "ASC"
	ListNodeReplaceConfigurationsSortOrderDesc ListNodeReplaceConfigurationsSortOrderEnum = "DESC"
)

var mappingListNodeReplaceConfigurationsSortOrderEnum = map[string]ListNodeReplaceConfigurationsSortOrderEnum{
	"ASC":  ListNodeReplaceConfigurationsSortOrderAsc,
	"DESC": ListNodeReplaceConfigurationsSortOrderDesc,
}

var mappingListNodeReplaceConfigurationsSortOrderEnumLowerCase = map[string]ListNodeReplaceConfigurationsSortOrderEnum{
	"asc":  ListNodeReplaceConfigurationsSortOrderAsc,
	"desc": ListNodeReplaceConfigurationsSortOrderDesc,
}

// GetListNodeReplaceConfigurationsSortOrderEnumValues Enumerates the set of values for ListNodeReplaceConfigurationsSortOrderEnum
func GetListNodeReplaceConfigurationsSortOrderEnumValues() []ListNodeReplaceConfigurationsSortOrderEnum {
	values := make([]ListNodeReplaceConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListNodeReplaceConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNodeReplaceConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListNodeReplaceConfigurationsSortOrderEnum
func GetListNodeReplaceConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNodeReplaceConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNodeReplaceConfigurationsSortOrderEnum(val string) (ListNodeReplaceConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListNodeReplaceConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
