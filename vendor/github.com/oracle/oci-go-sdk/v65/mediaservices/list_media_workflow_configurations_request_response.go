// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMediaWorkflowConfigurationsRequest wrapper for the ListMediaWorkflowConfigurations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListMediaWorkflowConfigurations.go.html to see an example of how to use ListMediaWorkflowConfigurationsRequest.
type ListMediaWorkflowConfigurationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources with lifecycleState matching the given lifecycleState.
	LifecycleState MediaWorkflowLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique MediaWorkflowConfiguration identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the
	// `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMediaWorkflowConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default
	// order for displayName is ascending.
	SortBy ListMediaWorkflowConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMediaWorkflowConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMediaWorkflowConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMediaWorkflowConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMediaWorkflowConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMediaWorkflowConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMediaWorkflowLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMediaWorkflowLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMediaWorkflowConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMediaWorkflowConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMediaWorkflowConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMediaWorkflowConfigurationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMediaWorkflowConfigurationsResponse wrapper for the ListMediaWorkflowConfigurations operation
type ListMediaWorkflowConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MediaWorkflowConfigurationCollection instances
	MediaWorkflowConfigurationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMediaWorkflowConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMediaWorkflowConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMediaWorkflowConfigurationsSortOrderEnum Enum with underlying type: string
type ListMediaWorkflowConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListMediaWorkflowConfigurationsSortOrderEnum
const (
	ListMediaWorkflowConfigurationsSortOrderAsc  ListMediaWorkflowConfigurationsSortOrderEnum = "ASC"
	ListMediaWorkflowConfigurationsSortOrderDesc ListMediaWorkflowConfigurationsSortOrderEnum = "DESC"
)

var mappingListMediaWorkflowConfigurationsSortOrderEnum = map[string]ListMediaWorkflowConfigurationsSortOrderEnum{
	"ASC":  ListMediaWorkflowConfigurationsSortOrderAsc,
	"DESC": ListMediaWorkflowConfigurationsSortOrderDesc,
}

var mappingListMediaWorkflowConfigurationsSortOrderEnumLowerCase = map[string]ListMediaWorkflowConfigurationsSortOrderEnum{
	"asc":  ListMediaWorkflowConfigurationsSortOrderAsc,
	"desc": ListMediaWorkflowConfigurationsSortOrderDesc,
}

// GetListMediaWorkflowConfigurationsSortOrderEnumValues Enumerates the set of values for ListMediaWorkflowConfigurationsSortOrderEnum
func GetListMediaWorkflowConfigurationsSortOrderEnumValues() []ListMediaWorkflowConfigurationsSortOrderEnum {
	values := make([]ListMediaWorkflowConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListMediaWorkflowConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaWorkflowConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListMediaWorkflowConfigurationsSortOrderEnum
func GetListMediaWorkflowConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMediaWorkflowConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaWorkflowConfigurationsSortOrderEnum(val string) (ListMediaWorkflowConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListMediaWorkflowConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMediaWorkflowConfigurationsSortByEnum Enum with underlying type: string
type ListMediaWorkflowConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListMediaWorkflowConfigurationsSortByEnum
const (
	ListMediaWorkflowConfigurationsSortByTimecreated ListMediaWorkflowConfigurationsSortByEnum = "timeCreated"
	ListMediaWorkflowConfigurationsSortByDisplayname ListMediaWorkflowConfigurationsSortByEnum = "displayName"
)

var mappingListMediaWorkflowConfigurationsSortByEnum = map[string]ListMediaWorkflowConfigurationsSortByEnum{
	"timeCreated": ListMediaWorkflowConfigurationsSortByTimecreated,
	"displayName": ListMediaWorkflowConfigurationsSortByDisplayname,
}

var mappingListMediaWorkflowConfigurationsSortByEnumLowerCase = map[string]ListMediaWorkflowConfigurationsSortByEnum{
	"timecreated": ListMediaWorkflowConfigurationsSortByTimecreated,
	"displayname": ListMediaWorkflowConfigurationsSortByDisplayname,
}

// GetListMediaWorkflowConfigurationsSortByEnumValues Enumerates the set of values for ListMediaWorkflowConfigurationsSortByEnum
func GetListMediaWorkflowConfigurationsSortByEnumValues() []ListMediaWorkflowConfigurationsSortByEnum {
	values := make([]ListMediaWorkflowConfigurationsSortByEnum, 0)
	for _, v := range mappingListMediaWorkflowConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaWorkflowConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListMediaWorkflowConfigurationsSortByEnum
func GetListMediaWorkflowConfigurationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMediaWorkflowConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaWorkflowConfigurationsSortByEnum(val string) (ListMediaWorkflowConfigurationsSortByEnum, bool) {
	enum, ok := mappingListMediaWorkflowConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
