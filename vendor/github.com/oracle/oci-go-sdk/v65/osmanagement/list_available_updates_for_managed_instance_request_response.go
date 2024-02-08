// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAvailableUpdatesForManagedInstanceRequest wrapper for the ListAvailableUpdatesForManagedInstance operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListAvailableUpdatesForManagedInstance.go.html to see an example of how to use ListAvailableUpdatesForManagedInstanceRequest.
type ListAvailableUpdatesForManagedInstanceRequest struct {

	// OCID for the managed instance
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAvailableUpdatesForManagedInstanceSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListAvailableUpdatesForManagedInstanceSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailableUpdatesForManagedInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailableUpdatesForManagedInstanceRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAvailableUpdatesForManagedInstanceRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailableUpdatesForManagedInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAvailableUpdatesForManagedInstanceRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAvailableUpdatesForManagedInstanceSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAvailableUpdatesForManagedInstanceSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAvailableUpdatesForManagedInstanceSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAvailableUpdatesForManagedInstanceSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAvailableUpdatesForManagedInstanceResponse wrapper for the ListAvailableUpdatesForManagedInstance operation
type ListAvailableUpdatesForManagedInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AvailableUpdateSummary instances
	Items []AvailableUpdateSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAvailableUpdatesForManagedInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailableUpdatesForManagedInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailableUpdatesForManagedInstanceSortOrderEnum Enum with underlying type: string
type ListAvailableUpdatesForManagedInstanceSortOrderEnum string

// Set of constants representing the allowable values for ListAvailableUpdatesForManagedInstanceSortOrderEnum
const (
	ListAvailableUpdatesForManagedInstanceSortOrderAsc  ListAvailableUpdatesForManagedInstanceSortOrderEnum = "ASC"
	ListAvailableUpdatesForManagedInstanceSortOrderDesc ListAvailableUpdatesForManagedInstanceSortOrderEnum = "DESC"
)

var mappingListAvailableUpdatesForManagedInstanceSortOrderEnum = map[string]ListAvailableUpdatesForManagedInstanceSortOrderEnum{
	"ASC":  ListAvailableUpdatesForManagedInstanceSortOrderAsc,
	"DESC": ListAvailableUpdatesForManagedInstanceSortOrderDesc,
}

var mappingListAvailableUpdatesForManagedInstanceSortOrderEnumLowerCase = map[string]ListAvailableUpdatesForManagedInstanceSortOrderEnum{
	"asc":  ListAvailableUpdatesForManagedInstanceSortOrderAsc,
	"desc": ListAvailableUpdatesForManagedInstanceSortOrderDesc,
}

// GetListAvailableUpdatesForManagedInstanceSortOrderEnumValues Enumerates the set of values for ListAvailableUpdatesForManagedInstanceSortOrderEnum
func GetListAvailableUpdatesForManagedInstanceSortOrderEnumValues() []ListAvailableUpdatesForManagedInstanceSortOrderEnum {
	values := make([]ListAvailableUpdatesForManagedInstanceSortOrderEnum, 0)
	for _, v := range mappingListAvailableUpdatesForManagedInstanceSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableUpdatesForManagedInstanceSortOrderEnumStringValues Enumerates the set of values in String for ListAvailableUpdatesForManagedInstanceSortOrderEnum
func GetListAvailableUpdatesForManagedInstanceSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAvailableUpdatesForManagedInstanceSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableUpdatesForManagedInstanceSortOrderEnum(val string) (ListAvailableUpdatesForManagedInstanceSortOrderEnum, bool) {
	enum, ok := mappingListAvailableUpdatesForManagedInstanceSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAvailableUpdatesForManagedInstanceSortByEnum Enum with underlying type: string
type ListAvailableUpdatesForManagedInstanceSortByEnum string

// Set of constants representing the allowable values for ListAvailableUpdatesForManagedInstanceSortByEnum
const (
	ListAvailableUpdatesForManagedInstanceSortByTimecreated ListAvailableUpdatesForManagedInstanceSortByEnum = "TIMECREATED"
	ListAvailableUpdatesForManagedInstanceSortByDisplayname ListAvailableUpdatesForManagedInstanceSortByEnum = "DISPLAYNAME"
)

var mappingListAvailableUpdatesForManagedInstanceSortByEnum = map[string]ListAvailableUpdatesForManagedInstanceSortByEnum{
	"TIMECREATED": ListAvailableUpdatesForManagedInstanceSortByTimecreated,
	"DISPLAYNAME": ListAvailableUpdatesForManagedInstanceSortByDisplayname,
}

var mappingListAvailableUpdatesForManagedInstanceSortByEnumLowerCase = map[string]ListAvailableUpdatesForManagedInstanceSortByEnum{
	"timecreated": ListAvailableUpdatesForManagedInstanceSortByTimecreated,
	"displayname": ListAvailableUpdatesForManagedInstanceSortByDisplayname,
}

// GetListAvailableUpdatesForManagedInstanceSortByEnumValues Enumerates the set of values for ListAvailableUpdatesForManagedInstanceSortByEnum
func GetListAvailableUpdatesForManagedInstanceSortByEnumValues() []ListAvailableUpdatesForManagedInstanceSortByEnum {
	values := make([]ListAvailableUpdatesForManagedInstanceSortByEnum, 0)
	for _, v := range mappingListAvailableUpdatesForManagedInstanceSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableUpdatesForManagedInstanceSortByEnumStringValues Enumerates the set of values in String for ListAvailableUpdatesForManagedInstanceSortByEnum
func GetListAvailableUpdatesForManagedInstanceSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAvailableUpdatesForManagedInstanceSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableUpdatesForManagedInstanceSortByEnum(val string) (ListAvailableUpdatesForManagedInstanceSortByEnum, bool) {
	enum, ok := mappingListAvailableUpdatesForManagedInstanceSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
