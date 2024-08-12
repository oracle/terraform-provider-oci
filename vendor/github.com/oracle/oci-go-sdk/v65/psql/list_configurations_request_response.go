// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListConfigurationsRequest wrapper for the ListConfigurations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psql/ListConfigurations.go.html to see an example of how to use ListConfigurationsRequest.
type ListConfigurationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources if their `lifecycleState` matches the given `lifecycleState`.
	LifecycleState ConfigurationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Version of the PostgreSQL database, such as 14.9.
	DbVersion *string `mandatory:"false" contributesTo:"query" name:"dbVersion"`

	// The name of the shape for the configuration.
	// Example: `VM.Standard.E4.Flex`
	Shape *string `mandatory:"false" contributesTo:"query" name:"shape"`

	// A unique identifier for the configuration.
	ConfigurationId *string `mandatory:"false" contributesTo:"query" name:"configurationId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

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
	if _, ok := GetMappingListConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConfigurationsSortByEnumStringValues(), ",")))
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

	// A list of ConfigurationCollection instances
	ConfigurationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
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

// ListConfigurationsSortByEnum Enum with underlying type: string
type ListConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListConfigurationsSortByEnum
const (
	ListConfigurationsSortByTimecreated ListConfigurationsSortByEnum = "timeCreated"
	ListConfigurationsSortByDisplayname ListConfigurationsSortByEnum = "displayName"
)

var mappingListConfigurationsSortByEnum = map[string]ListConfigurationsSortByEnum{
	"timeCreated": ListConfigurationsSortByTimecreated,
	"displayName": ListConfigurationsSortByDisplayname,
}

var mappingListConfigurationsSortByEnumLowerCase = map[string]ListConfigurationsSortByEnum{
	"timecreated": ListConfigurationsSortByTimecreated,
	"displayname": ListConfigurationsSortByDisplayname,
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
		"timeCreated",
		"displayName",
	}
}

// GetMappingListConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConfigurationsSortByEnum(val string) (ListConfigurationsSortByEnum, bool) {
	enum, ok := mappingListConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
