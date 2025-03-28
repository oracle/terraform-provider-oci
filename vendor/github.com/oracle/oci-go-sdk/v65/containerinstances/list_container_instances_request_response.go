// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package containerinstances

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListContainerInstancesRequest wrapper for the ListContainerInstances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/containerinstances/ListContainerInstances.go.html to see an example of how to use ListContainerInstancesRequest.
type ListContainerInstancesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to only return resources that match the given lifecycle state.
	LifecycleState ContainerInstanceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use (ASC) or (DESC).
	SortOrder ListContainerInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order. Default order for timeCreated is descending. Default order for displayName is ascending. If you don't specify a value, timeCreated is the default.
	SortBy ListContainerInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListContainerInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListContainerInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListContainerInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListContainerInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListContainerInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingContainerInstanceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetContainerInstanceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListContainerInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListContainerInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListContainerInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListContainerInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListContainerInstancesResponse wrapper for the ListContainerInstances operation
type ListContainerInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ContainerInstanceCollection instances
	ContainerInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListContainerInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListContainerInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListContainerInstancesSortOrderEnum Enum with underlying type: string
type ListContainerInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListContainerInstancesSortOrderEnum
const (
	ListContainerInstancesSortOrderAsc  ListContainerInstancesSortOrderEnum = "ASC"
	ListContainerInstancesSortOrderDesc ListContainerInstancesSortOrderEnum = "DESC"
)

var mappingListContainerInstancesSortOrderEnum = map[string]ListContainerInstancesSortOrderEnum{
	"ASC":  ListContainerInstancesSortOrderAsc,
	"DESC": ListContainerInstancesSortOrderDesc,
}

var mappingListContainerInstancesSortOrderEnumLowerCase = map[string]ListContainerInstancesSortOrderEnum{
	"asc":  ListContainerInstancesSortOrderAsc,
	"desc": ListContainerInstancesSortOrderDesc,
}

// GetListContainerInstancesSortOrderEnumValues Enumerates the set of values for ListContainerInstancesSortOrderEnum
func GetListContainerInstancesSortOrderEnumValues() []ListContainerInstancesSortOrderEnum {
	values := make([]ListContainerInstancesSortOrderEnum, 0)
	for _, v := range mappingListContainerInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainerInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListContainerInstancesSortOrderEnum
func GetListContainerInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListContainerInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainerInstancesSortOrderEnum(val string) (ListContainerInstancesSortOrderEnum, bool) {
	enum, ok := mappingListContainerInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListContainerInstancesSortByEnum Enum with underlying type: string
type ListContainerInstancesSortByEnum string

// Set of constants representing the allowable values for ListContainerInstancesSortByEnum
const (
	ListContainerInstancesSortByTimecreated ListContainerInstancesSortByEnum = "timeCreated"
	ListContainerInstancesSortByDisplayname ListContainerInstancesSortByEnum = "displayName"
)

var mappingListContainerInstancesSortByEnum = map[string]ListContainerInstancesSortByEnum{
	"timeCreated": ListContainerInstancesSortByTimecreated,
	"displayName": ListContainerInstancesSortByDisplayname,
}

var mappingListContainerInstancesSortByEnumLowerCase = map[string]ListContainerInstancesSortByEnum{
	"timecreated": ListContainerInstancesSortByTimecreated,
	"displayname": ListContainerInstancesSortByDisplayname,
}

// GetListContainerInstancesSortByEnumValues Enumerates the set of values for ListContainerInstancesSortByEnum
func GetListContainerInstancesSortByEnumValues() []ListContainerInstancesSortByEnum {
	values := make([]ListContainerInstancesSortByEnum, 0)
	for _, v := range mappingListContainerInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainerInstancesSortByEnumStringValues Enumerates the set of values in String for ListContainerInstancesSortByEnum
func GetListContainerInstancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListContainerInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainerInstancesSortByEnum(val string) (ListContainerInstancesSortByEnum, bool) {
	enum, ok := mappingListContainerInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
