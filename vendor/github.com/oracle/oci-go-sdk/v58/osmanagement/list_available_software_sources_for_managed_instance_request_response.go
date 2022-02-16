// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListAvailableSoftwareSourcesForManagedInstanceRequest wrapper for the ListAvailableSoftwareSourcesForManagedInstance operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListAvailableSoftwareSourcesForManagedInstance.go.html to see an example of how to use ListAvailableSoftwareSourcesForManagedInstanceRequest.
type ListAvailableSoftwareSourcesForManagedInstanceRequest struct {

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
	SortOrder ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListAvailableSoftwareSourcesForManagedInstanceSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailableSoftwareSourcesForManagedInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailableSoftwareSourcesForManagedInstanceRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAvailableSoftwareSourcesForManagedInstanceRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailableSoftwareSourcesForManagedInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAvailableSoftwareSourcesForManagedInstanceRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAvailableSoftwareSourcesForManagedInstanceSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAvailableSoftwareSourcesForManagedInstanceSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAvailableSoftwareSourcesForManagedInstanceSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAvailableSoftwareSourcesForManagedInstanceResponse wrapper for the ListAvailableSoftwareSourcesForManagedInstance operation
type ListAvailableSoftwareSourcesForManagedInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AvailableSoftwareSourceSummary instances
	Items []AvailableSoftwareSourceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAvailableSoftwareSourcesForManagedInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailableSoftwareSourcesForManagedInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum Enum with underlying type: string
type ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum string

// Set of constants representing the allowable values for ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum
const (
	ListAvailableSoftwareSourcesForManagedInstanceSortOrderAsc  ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum = "ASC"
	ListAvailableSoftwareSourcesForManagedInstanceSortOrderDesc ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum = "DESC"
)

var mappingListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum = map[string]ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum{
	"ASC":  ListAvailableSoftwareSourcesForManagedInstanceSortOrderAsc,
	"DESC": ListAvailableSoftwareSourcesForManagedInstanceSortOrderDesc,
}

// GetListAvailableSoftwareSourcesForManagedInstanceSortOrderEnumValues Enumerates the set of values for ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum
func GetListAvailableSoftwareSourcesForManagedInstanceSortOrderEnumValues() []ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum {
	values := make([]ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum, 0)
	for _, v := range mappingListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableSoftwareSourcesForManagedInstanceSortOrderEnumStringValues Enumerates the set of values in String for ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum
func GetListAvailableSoftwareSourcesForManagedInstanceSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum(val string) (ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum, bool) {
	mappingListAvailableSoftwareSourcesForManagedInstanceSortOrderEnumIgnoreCase := make(map[string]ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum)
	for k, v := range mappingListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum {
		mappingListAvailableSoftwareSourcesForManagedInstanceSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAvailableSoftwareSourcesForManagedInstanceSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAvailableSoftwareSourcesForManagedInstanceSortByEnum Enum with underlying type: string
type ListAvailableSoftwareSourcesForManagedInstanceSortByEnum string

// Set of constants representing the allowable values for ListAvailableSoftwareSourcesForManagedInstanceSortByEnum
const (
	ListAvailableSoftwareSourcesForManagedInstanceSortByTimecreated ListAvailableSoftwareSourcesForManagedInstanceSortByEnum = "TIMECREATED"
	ListAvailableSoftwareSourcesForManagedInstanceSortByDisplayname ListAvailableSoftwareSourcesForManagedInstanceSortByEnum = "DISPLAYNAME"
)

var mappingListAvailableSoftwareSourcesForManagedInstanceSortByEnum = map[string]ListAvailableSoftwareSourcesForManagedInstanceSortByEnum{
	"TIMECREATED": ListAvailableSoftwareSourcesForManagedInstanceSortByTimecreated,
	"DISPLAYNAME": ListAvailableSoftwareSourcesForManagedInstanceSortByDisplayname,
}

// GetListAvailableSoftwareSourcesForManagedInstanceSortByEnumValues Enumerates the set of values for ListAvailableSoftwareSourcesForManagedInstanceSortByEnum
func GetListAvailableSoftwareSourcesForManagedInstanceSortByEnumValues() []ListAvailableSoftwareSourcesForManagedInstanceSortByEnum {
	values := make([]ListAvailableSoftwareSourcesForManagedInstanceSortByEnum, 0)
	for _, v := range mappingListAvailableSoftwareSourcesForManagedInstanceSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableSoftwareSourcesForManagedInstanceSortByEnumStringValues Enumerates the set of values in String for ListAvailableSoftwareSourcesForManagedInstanceSortByEnum
func GetListAvailableSoftwareSourcesForManagedInstanceSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAvailableSoftwareSourcesForManagedInstanceSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableSoftwareSourcesForManagedInstanceSortByEnum(val string) (ListAvailableSoftwareSourcesForManagedInstanceSortByEnum, bool) {
	mappingListAvailableSoftwareSourcesForManagedInstanceSortByEnumIgnoreCase := make(map[string]ListAvailableSoftwareSourcesForManagedInstanceSortByEnum)
	for k, v := range mappingListAvailableSoftwareSourcesForManagedInstanceSortByEnum {
		mappingListAvailableSoftwareSourcesForManagedInstanceSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAvailableSoftwareSourcesForManagedInstanceSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
