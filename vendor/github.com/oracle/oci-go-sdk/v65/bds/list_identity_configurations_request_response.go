// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListIdentityConfigurationsRequest wrapper for the ListIdentityConfigurations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListIdentityConfigurations.go.html to see an example of how to use ListIdentityConfigurationsRequest.
type ListIdentityConfigurationsRequest struct {

	// The OCID of the cluster.
	BdsInstanceId *string `mandatory:"true" contributesTo:"path" name:"bdsInstanceId"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListIdentityConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListIdentityConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The state of the identity config
	LifecycleState IdentityConfigurationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIdentityConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIdentityConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIdentityConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIdentityConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListIdentityConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListIdentityConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListIdentityConfigurationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIdentityConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListIdentityConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIdentityConfigurationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetIdentityConfigurationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIdentityConfigurationsResponse wrapper for the ListIdentityConfigurations operation
type ListIdentityConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []IdentityConfigurationSummary instances
	Items []IdentityConfigurationSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListIdentityConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIdentityConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIdentityConfigurationsSortByEnum Enum with underlying type: string
type ListIdentityConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListIdentityConfigurationsSortByEnum
const (
	ListIdentityConfigurationsSortByTimecreated ListIdentityConfigurationsSortByEnum = "timeCreated"
	ListIdentityConfigurationsSortByDisplayname ListIdentityConfigurationsSortByEnum = "displayName"
)

var mappingListIdentityConfigurationsSortByEnum = map[string]ListIdentityConfigurationsSortByEnum{
	"timeCreated": ListIdentityConfigurationsSortByTimecreated,
	"displayName": ListIdentityConfigurationsSortByDisplayname,
}

var mappingListIdentityConfigurationsSortByEnumLowerCase = map[string]ListIdentityConfigurationsSortByEnum{
	"timecreated": ListIdentityConfigurationsSortByTimecreated,
	"displayname": ListIdentityConfigurationsSortByDisplayname,
}

// GetListIdentityConfigurationsSortByEnumValues Enumerates the set of values for ListIdentityConfigurationsSortByEnum
func GetListIdentityConfigurationsSortByEnumValues() []ListIdentityConfigurationsSortByEnum {
	values := make([]ListIdentityConfigurationsSortByEnum, 0)
	for _, v := range mappingListIdentityConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListIdentityConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListIdentityConfigurationsSortByEnum
func GetListIdentityConfigurationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListIdentityConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIdentityConfigurationsSortByEnum(val string) (ListIdentityConfigurationsSortByEnum, bool) {
	enum, ok := mappingListIdentityConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIdentityConfigurationsSortOrderEnum Enum with underlying type: string
type ListIdentityConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListIdentityConfigurationsSortOrderEnum
const (
	ListIdentityConfigurationsSortOrderAsc  ListIdentityConfigurationsSortOrderEnum = "ASC"
	ListIdentityConfigurationsSortOrderDesc ListIdentityConfigurationsSortOrderEnum = "DESC"
)

var mappingListIdentityConfigurationsSortOrderEnum = map[string]ListIdentityConfigurationsSortOrderEnum{
	"ASC":  ListIdentityConfigurationsSortOrderAsc,
	"DESC": ListIdentityConfigurationsSortOrderDesc,
}

var mappingListIdentityConfigurationsSortOrderEnumLowerCase = map[string]ListIdentityConfigurationsSortOrderEnum{
	"asc":  ListIdentityConfigurationsSortOrderAsc,
	"desc": ListIdentityConfigurationsSortOrderDesc,
}

// GetListIdentityConfigurationsSortOrderEnumValues Enumerates the set of values for ListIdentityConfigurationsSortOrderEnum
func GetListIdentityConfigurationsSortOrderEnumValues() []ListIdentityConfigurationsSortOrderEnum {
	values := make([]ListIdentityConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListIdentityConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIdentityConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListIdentityConfigurationsSortOrderEnum
func GetListIdentityConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIdentityConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIdentityConfigurationsSortOrderEnum(val string) (ListIdentityConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListIdentityConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
