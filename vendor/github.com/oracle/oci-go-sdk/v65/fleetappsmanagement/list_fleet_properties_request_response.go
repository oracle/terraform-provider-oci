// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFleetPropertiesRequest wrapper for the ListFleetProperties operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListFleetProperties.go.html to see an example of how to use ListFleetPropertiesRequest.
type ListFleetPropertiesRequest struct {

	// Unique Fleet identifier.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState FleetPropertyLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources whose fleetProperty identifier matches the given identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFleetPropertiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListFleetPropertiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFleetPropertiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFleetPropertiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFleetPropertiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFleetPropertiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFleetPropertiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFleetPropertyLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetFleetPropertyLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFleetPropertiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFleetPropertiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFleetPropertiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFleetPropertiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFleetPropertiesResponse wrapper for the ListFleetProperties operation
type ListFleetPropertiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FleetPropertyCollection instances
	FleetPropertyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFleetPropertiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFleetPropertiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFleetPropertiesSortOrderEnum Enum with underlying type: string
type ListFleetPropertiesSortOrderEnum string

// Set of constants representing the allowable values for ListFleetPropertiesSortOrderEnum
const (
	ListFleetPropertiesSortOrderAsc  ListFleetPropertiesSortOrderEnum = "ASC"
	ListFleetPropertiesSortOrderDesc ListFleetPropertiesSortOrderEnum = "DESC"
)

var mappingListFleetPropertiesSortOrderEnum = map[string]ListFleetPropertiesSortOrderEnum{
	"ASC":  ListFleetPropertiesSortOrderAsc,
	"DESC": ListFleetPropertiesSortOrderDesc,
}

var mappingListFleetPropertiesSortOrderEnumLowerCase = map[string]ListFleetPropertiesSortOrderEnum{
	"asc":  ListFleetPropertiesSortOrderAsc,
	"desc": ListFleetPropertiesSortOrderDesc,
}

// GetListFleetPropertiesSortOrderEnumValues Enumerates the set of values for ListFleetPropertiesSortOrderEnum
func GetListFleetPropertiesSortOrderEnumValues() []ListFleetPropertiesSortOrderEnum {
	values := make([]ListFleetPropertiesSortOrderEnum, 0)
	for _, v := range mappingListFleetPropertiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetPropertiesSortOrderEnumStringValues Enumerates the set of values in String for ListFleetPropertiesSortOrderEnum
func GetListFleetPropertiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFleetPropertiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetPropertiesSortOrderEnum(val string) (ListFleetPropertiesSortOrderEnum, bool) {
	enum, ok := mappingListFleetPropertiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFleetPropertiesSortByEnum Enum with underlying type: string
type ListFleetPropertiesSortByEnum string

// Set of constants representing the allowable values for ListFleetPropertiesSortByEnum
const (
	ListFleetPropertiesSortByTimecreated ListFleetPropertiesSortByEnum = "timeCreated"
	ListFleetPropertiesSortByDisplayname ListFleetPropertiesSortByEnum = "displayName"
)

var mappingListFleetPropertiesSortByEnum = map[string]ListFleetPropertiesSortByEnum{
	"timeCreated": ListFleetPropertiesSortByTimecreated,
	"displayName": ListFleetPropertiesSortByDisplayname,
}

var mappingListFleetPropertiesSortByEnumLowerCase = map[string]ListFleetPropertiesSortByEnum{
	"timecreated": ListFleetPropertiesSortByTimecreated,
	"displayname": ListFleetPropertiesSortByDisplayname,
}

// GetListFleetPropertiesSortByEnumValues Enumerates the set of values for ListFleetPropertiesSortByEnum
func GetListFleetPropertiesSortByEnumValues() []ListFleetPropertiesSortByEnum {
	values := make([]ListFleetPropertiesSortByEnum, 0)
	for _, v := range mappingListFleetPropertiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetPropertiesSortByEnumStringValues Enumerates the set of values in String for ListFleetPropertiesSortByEnum
func GetListFleetPropertiesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListFleetPropertiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetPropertiesSortByEnum(val string) (ListFleetPropertiesSortByEnum, bool) {
	enum, ok := mappingListFleetPropertiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
