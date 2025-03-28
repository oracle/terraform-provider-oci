// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDefaultConfigurationsRequest wrapper for the ListDefaultConfigurations operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psql/ListDefaultConfigurations.go.html to see an example of how to use ListDefaultConfigurationsRequest.
type ListDefaultConfigurationsRequest struct {

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
	SortOrder ListDefaultConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListDefaultConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDefaultConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDefaultConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDefaultConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDefaultConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDefaultConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConfigurationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetConfigurationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDefaultConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDefaultConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDefaultConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDefaultConfigurationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDefaultConfigurationsResponse wrapper for the ListDefaultConfigurations operation
type ListDefaultConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DefaultConfigurationCollection instances
	DefaultConfigurationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDefaultConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDefaultConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDefaultConfigurationsSortOrderEnum Enum with underlying type: string
type ListDefaultConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListDefaultConfigurationsSortOrderEnum
const (
	ListDefaultConfigurationsSortOrderAsc  ListDefaultConfigurationsSortOrderEnum = "ASC"
	ListDefaultConfigurationsSortOrderDesc ListDefaultConfigurationsSortOrderEnum = "DESC"
)

var mappingListDefaultConfigurationsSortOrderEnum = map[string]ListDefaultConfigurationsSortOrderEnum{
	"ASC":  ListDefaultConfigurationsSortOrderAsc,
	"DESC": ListDefaultConfigurationsSortOrderDesc,
}

var mappingListDefaultConfigurationsSortOrderEnumLowerCase = map[string]ListDefaultConfigurationsSortOrderEnum{
	"asc":  ListDefaultConfigurationsSortOrderAsc,
	"desc": ListDefaultConfigurationsSortOrderDesc,
}

// GetListDefaultConfigurationsSortOrderEnumValues Enumerates the set of values for ListDefaultConfigurationsSortOrderEnum
func GetListDefaultConfigurationsSortOrderEnumValues() []ListDefaultConfigurationsSortOrderEnum {
	values := make([]ListDefaultConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListDefaultConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDefaultConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListDefaultConfigurationsSortOrderEnum
func GetListDefaultConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDefaultConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDefaultConfigurationsSortOrderEnum(val string) (ListDefaultConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListDefaultConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDefaultConfigurationsSortByEnum Enum with underlying type: string
type ListDefaultConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListDefaultConfigurationsSortByEnum
const (
	ListDefaultConfigurationsSortByTimecreated ListDefaultConfigurationsSortByEnum = "timeCreated"
	ListDefaultConfigurationsSortByDisplayname ListDefaultConfigurationsSortByEnum = "displayName"
)

var mappingListDefaultConfigurationsSortByEnum = map[string]ListDefaultConfigurationsSortByEnum{
	"timeCreated": ListDefaultConfigurationsSortByTimecreated,
	"displayName": ListDefaultConfigurationsSortByDisplayname,
}

var mappingListDefaultConfigurationsSortByEnumLowerCase = map[string]ListDefaultConfigurationsSortByEnum{
	"timecreated": ListDefaultConfigurationsSortByTimecreated,
	"displayname": ListDefaultConfigurationsSortByDisplayname,
}

// GetListDefaultConfigurationsSortByEnumValues Enumerates the set of values for ListDefaultConfigurationsSortByEnum
func GetListDefaultConfigurationsSortByEnumValues() []ListDefaultConfigurationsSortByEnum {
	values := make([]ListDefaultConfigurationsSortByEnum, 0)
	for _, v := range mappingListDefaultConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDefaultConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListDefaultConfigurationsSortByEnum
func GetListDefaultConfigurationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDefaultConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDefaultConfigurationsSortByEnum(val string) (ListDefaultConfigurationsSortByEnum, bool) {
	enum, ok := mappingListDefaultConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
