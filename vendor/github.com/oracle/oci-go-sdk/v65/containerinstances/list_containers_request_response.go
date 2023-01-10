// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package containerinstances

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListContainersRequest wrapper for the ListContainers operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/containerinstances/ListContainers.go.html to see an example of how to use ListContainersRequest.
type ListContainersRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState ContainerLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique ContainerInstance identifier
	ContainerInstanceId *string `mandatory:"false" contributesTo:"query" name:"containerInstanceId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListContainersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListContainersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListContainersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListContainersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListContainersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListContainersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListContainersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingContainerLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetContainerLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListContainersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListContainersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListContainersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListContainersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListContainersResponse wrapper for the ListContainers operation
type ListContainersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ContainerCollection instances
	ContainerCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListContainersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListContainersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListContainersSortOrderEnum Enum with underlying type: string
type ListContainersSortOrderEnum string

// Set of constants representing the allowable values for ListContainersSortOrderEnum
const (
	ListContainersSortOrderAsc  ListContainersSortOrderEnum = "ASC"
	ListContainersSortOrderDesc ListContainersSortOrderEnum = "DESC"
)

var mappingListContainersSortOrderEnum = map[string]ListContainersSortOrderEnum{
	"ASC":  ListContainersSortOrderAsc,
	"DESC": ListContainersSortOrderDesc,
}

var mappingListContainersSortOrderEnumLowerCase = map[string]ListContainersSortOrderEnum{
	"asc":  ListContainersSortOrderAsc,
	"desc": ListContainersSortOrderDesc,
}

// GetListContainersSortOrderEnumValues Enumerates the set of values for ListContainersSortOrderEnum
func GetListContainersSortOrderEnumValues() []ListContainersSortOrderEnum {
	values := make([]ListContainersSortOrderEnum, 0)
	for _, v := range mappingListContainersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainersSortOrderEnumStringValues Enumerates the set of values in String for ListContainersSortOrderEnum
func GetListContainersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListContainersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainersSortOrderEnum(val string) (ListContainersSortOrderEnum, bool) {
	enum, ok := mappingListContainersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListContainersSortByEnum Enum with underlying type: string
type ListContainersSortByEnum string

// Set of constants representing the allowable values for ListContainersSortByEnum
const (
	ListContainersSortByTimecreated ListContainersSortByEnum = "timeCreated"
	ListContainersSortByDisplayname ListContainersSortByEnum = "displayName"
)

var mappingListContainersSortByEnum = map[string]ListContainersSortByEnum{
	"timeCreated": ListContainersSortByTimecreated,
	"displayName": ListContainersSortByDisplayname,
}

var mappingListContainersSortByEnumLowerCase = map[string]ListContainersSortByEnum{
	"timecreated": ListContainersSortByTimecreated,
	"displayname": ListContainersSortByDisplayname,
}

// GetListContainersSortByEnumValues Enumerates the set of values for ListContainersSortByEnum
func GetListContainersSortByEnumValues() []ListContainersSortByEnum {
	values := make([]ListContainersSortByEnum, 0)
	for _, v := range mappingListContainersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainersSortByEnumStringValues Enumerates the set of values in String for ListContainersSortByEnum
func GetListContainersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListContainersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainersSortByEnum(val string) (ListContainersSortByEnum, bool) {
	enum, ok := mappingListContainersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
