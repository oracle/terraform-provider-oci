// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListNetworkLocationsRequest wrapper for the ListNetworkLocations operation
type ListNetworkLocationsRequest struct {

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for NAME is ascending. The NAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListNetworkLocationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListNetworkLocationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState NetworkLocationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNetworkLocationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNetworkLocationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNetworkLocationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNetworkLocationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNetworkLocationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNetworkLocationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNetworkLocationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNetworkLocationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNetworkLocationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNetworkLocationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetNetworkLocationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNetworkLocationsResponse wrapper for the ListNetworkLocations operation
type ListNetworkLocationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []NetworkLocationSummary instances
	Items []NetworkLocationSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListNetworkLocationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNetworkLocationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNetworkLocationsSortByEnum Enum with underlying type: string
type ListNetworkLocationsSortByEnum string

// Set of constants representing the allowable values for ListNetworkLocationsSortByEnum
const (
	ListNetworkLocationsSortByTimecreated ListNetworkLocationsSortByEnum = "TIMECREATED"
	ListNetworkLocationsSortByName        ListNetworkLocationsSortByEnum = "NAME"
)

var mappingListNetworkLocationsSortByEnum = map[string]ListNetworkLocationsSortByEnum{
	"TIMECREATED": ListNetworkLocationsSortByTimecreated,
	"NAME":        ListNetworkLocationsSortByName,
}

var mappingListNetworkLocationsSortByEnumLowerCase = map[string]ListNetworkLocationsSortByEnum{
	"timecreated": ListNetworkLocationsSortByTimecreated,
	"name":        ListNetworkLocationsSortByName,
}

// GetListNetworkLocationsSortByEnumValues Enumerates the set of values for ListNetworkLocationsSortByEnum
func GetListNetworkLocationsSortByEnumValues() []ListNetworkLocationsSortByEnum {
	values := make([]ListNetworkLocationsSortByEnum, 0)
	for _, v := range mappingListNetworkLocationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkLocationsSortByEnumStringValues Enumerates the set of values in String for ListNetworkLocationsSortByEnum
func GetListNetworkLocationsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListNetworkLocationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkLocationsSortByEnum(val string) (ListNetworkLocationsSortByEnum, bool) {
	enum, ok := mappingListNetworkLocationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNetworkLocationsSortOrderEnum Enum with underlying type: string
type ListNetworkLocationsSortOrderEnum string

// Set of constants representing the allowable values for ListNetworkLocationsSortOrderEnum
const (
	ListNetworkLocationsSortOrderAsc  ListNetworkLocationsSortOrderEnum = "ASC"
	ListNetworkLocationsSortOrderDesc ListNetworkLocationsSortOrderEnum = "DESC"
)

var mappingListNetworkLocationsSortOrderEnum = map[string]ListNetworkLocationsSortOrderEnum{
	"ASC":  ListNetworkLocationsSortOrderAsc,
	"DESC": ListNetworkLocationsSortOrderDesc,
}

var mappingListNetworkLocationsSortOrderEnumLowerCase = map[string]ListNetworkLocationsSortOrderEnum{
	"asc":  ListNetworkLocationsSortOrderAsc,
	"desc": ListNetworkLocationsSortOrderDesc,
}

// GetListNetworkLocationsSortOrderEnumValues Enumerates the set of values for ListNetworkLocationsSortOrderEnum
func GetListNetworkLocationsSortOrderEnumValues() []ListNetworkLocationsSortOrderEnum {
	values := make([]ListNetworkLocationsSortOrderEnum, 0)
	for _, v := range mappingListNetworkLocationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkLocationsSortOrderEnumStringValues Enumerates the set of values in String for ListNetworkLocationsSortOrderEnum
func GetListNetworkLocationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNetworkLocationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkLocationsSortOrderEnum(val string) (ListNetworkLocationsSortOrderEnum, bool) {
	enum, ok := mappingListNetworkLocationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
