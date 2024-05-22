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

// ListNodeBackupConfigurationsRequest wrapper for the ListNodeBackupConfigurations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListNodeBackupConfigurations.go.html to see an example of how to use ListNodeBackupConfigurationsRequest.
type ListNodeBackupConfigurationsRequest struct {

	// The OCID of the cluster.
	BdsInstanceId *string `mandatory:"true" contributesTo:"path" name:"bdsInstanceId"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListNodeBackupConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListNodeBackupConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The state of the NodeBackupConfiguration configuration.
	LifecycleState NodeBackupConfigurationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNodeBackupConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNodeBackupConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNodeBackupConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNodeBackupConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNodeBackupConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNodeBackupConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNodeBackupConfigurationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNodeBackupConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNodeBackupConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNodeBackupConfigurationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetNodeBackupConfigurationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNodeBackupConfigurationsResponse wrapper for the ListNodeBackupConfigurations operation
type ListNodeBackupConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []NodeBackupConfigurationSummary instances
	Items []NodeBackupConfigurationSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListNodeBackupConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNodeBackupConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNodeBackupConfigurationsSortByEnum Enum with underlying type: string
type ListNodeBackupConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListNodeBackupConfigurationsSortByEnum
const (
	ListNodeBackupConfigurationsSortByTimecreated ListNodeBackupConfigurationsSortByEnum = "timeCreated"
	ListNodeBackupConfigurationsSortByDisplayname ListNodeBackupConfigurationsSortByEnum = "displayName"
)

var mappingListNodeBackupConfigurationsSortByEnum = map[string]ListNodeBackupConfigurationsSortByEnum{
	"timeCreated": ListNodeBackupConfigurationsSortByTimecreated,
	"displayName": ListNodeBackupConfigurationsSortByDisplayname,
}

var mappingListNodeBackupConfigurationsSortByEnumLowerCase = map[string]ListNodeBackupConfigurationsSortByEnum{
	"timecreated": ListNodeBackupConfigurationsSortByTimecreated,
	"displayname": ListNodeBackupConfigurationsSortByDisplayname,
}

// GetListNodeBackupConfigurationsSortByEnumValues Enumerates the set of values for ListNodeBackupConfigurationsSortByEnum
func GetListNodeBackupConfigurationsSortByEnumValues() []ListNodeBackupConfigurationsSortByEnum {
	values := make([]ListNodeBackupConfigurationsSortByEnum, 0)
	for _, v := range mappingListNodeBackupConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNodeBackupConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListNodeBackupConfigurationsSortByEnum
func GetListNodeBackupConfigurationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListNodeBackupConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNodeBackupConfigurationsSortByEnum(val string) (ListNodeBackupConfigurationsSortByEnum, bool) {
	enum, ok := mappingListNodeBackupConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNodeBackupConfigurationsSortOrderEnum Enum with underlying type: string
type ListNodeBackupConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListNodeBackupConfigurationsSortOrderEnum
const (
	ListNodeBackupConfigurationsSortOrderAsc  ListNodeBackupConfigurationsSortOrderEnum = "ASC"
	ListNodeBackupConfigurationsSortOrderDesc ListNodeBackupConfigurationsSortOrderEnum = "DESC"
)

var mappingListNodeBackupConfigurationsSortOrderEnum = map[string]ListNodeBackupConfigurationsSortOrderEnum{
	"ASC":  ListNodeBackupConfigurationsSortOrderAsc,
	"DESC": ListNodeBackupConfigurationsSortOrderDesc,
}

var mappingListNodeBackupConfigurationsSortOrderEnumLowerCase = map[string]ListNodeBackupConfigurationsSortOrderEnum{
	"asc":  ListNodeBackupConfigurationsSortOrderAsc,
	"desc": ListNodeBackupConfigurationsSortOrderDesc,
}

// GetListNodeBackupConfigurationsSortOrderEnumValues Enumerates the set of values for ListNodeBackupConfigurationsSortOrderEnum
func GetListNodeBackupConfigurationsSortOrderEnumValues() []ListNodeBackupConfigurationsSortOrderEnum {
	values := make([]ListNodeBackupConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListNodeBackupConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNodeBackupConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListNodeBackupConfigurationsSortOrderEnum
func GetListNodeBackupConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNodeBackupConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNodeBackupConfigurationsSortOrderEnum(val string) (ListNodeBackupConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListNodeBackupConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
