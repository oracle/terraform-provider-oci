// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEnvironmentsRequest wrapper for the ListEnvironments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListEnvironments.go.html to see an example of how to use ListEnvironmentsRequest.
type ListEnvironmentsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources where their lifecycleState matches the given lifecycleState.
	LifecycleState EnvironmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given environment ID.
	EnvironmentId *string `mandatory:"false" contributesTo:"query" name:"environmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListEnvironmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListEnvironmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEnvironmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEnvironmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEnvironmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEnvironmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEnvironmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEnvironmentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetEnvironmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEnvironmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEnvironmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEnvironmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEnvironmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEnvironmentsResponse wrapper for the ListEnvironments operation
type ListEnvironmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EnvironmentCollection instances
	EnvironmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEnvironmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEnvironmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEnvironmentsSortOrderEnum Enum with underlying type: string
type ListEnvironmentsSortOrderEnum string

// Set of constants representing the allowable values for ListEnvironmentsSortOrderEnum
const (
	ListEnvironmentsSortOrderAsc  ListEnvironmentsSortOrderEnum = "ASC"
	ListEnvironmentsSortOrderDesc ListEnvironmentsSortOrderEnum = "DESC"
)

var mappingListEnvironmentsSortOrderEnum = map[string]ListEnvironmentsSortOrderEnum{
	"ASC":  ListEnvironmentsSortOrderAsc,
	"DESC": ListEnvironmentsSortOrderDesc,
}

var mappingListEnvironmentsSortOrderEnumLowerCase = map[string]ListEnvironmentsSortOrderEnum{
	"asc":  ListEnvironmentsSortOrderAsc,
	"desc": ListEnvironmentsSortOrderDesc,
}

// GetListEnvironmentsSortOrderEnumValues Enumerates the set of values for ListEnvironmentsSortOrderEnum
func GetListEnvironmentsSortOrderEnumValues() []ListEnvironmentsSortOrderEnum {
	values := make([]ListEnvironmentsSortOrderEnum, 0)
	for _, v := range mappingListEnvironmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEnvironmentsSortOrderEnumStringValues Enumerates the set of values in String for ListEnvironmentsSortOrderEnum
func GetListEnvironmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEnvironmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEnvironmentsSortOrderEnum(val string) (ListEnvironmentsSortOrderEnum, bool) {
	enum, ok := mappingListEnvironmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEnvironmentsSortByEnum Enum with underlying type: string
type ListEnvironmentsSortByEnum string

// Set of constants representing the allowable values for ListEnvironmentsSortByEnum
const (
	ListEnvironmentsSortByTimecreated ListEnvironmentsSortByEnum = "timeCreated"
	ListEnvironmentsSortByTimeupdated ListEnvironmentsSortByEnum = "timeUpdated"
	ListEnvironmentsSortByDisplayname ListEnvironmentsSortByEnum = "displayName"
)

var mappingListEnvironmentsSortByEnum = map[string]ListEnvironmentsSortByEnum{
	"timeCreated": ListEnvironmentsSortByTimecreated,
	"timeUpdated": ListEnvironmentsSortByTimeupdated,
	"displayName": ListEnvironmentsSortByDisplayname,
}

var mappingListEnvironmentsSortByEnumLowerCase = map[string]ListEnvironmentsSortByEnum{
	"timecreated": ListEnvironmentsSortByTimecreated,
	"timeupdated": ListEnvironmentsSortByTimeupdated,
	"displayname": ListEnvironmentsSortByDisplayname,
}

// GetListEnvironmentsSortByEnumValues Enumerates the set of values for ListEnvironmentsSortByEnum
func GetListEnvironmentsSortByEnumValues() []ListEnvironmentsSortByEnum {
	values := make([]ListEnvironmentsSortByEnum, 0)
	for _, v := range mappingListEnvironmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEnvironmentsSortByEnumStringValues Enumerates the set of values in String for ListEnvironmentsSortByEnum
func GetListEnvironmentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListEnvironmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEnvironmentsSortByEnum(val string) (ListEnvironmentsSortByEnum, bool) {
	enum, ok := mappingListEnvironmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
