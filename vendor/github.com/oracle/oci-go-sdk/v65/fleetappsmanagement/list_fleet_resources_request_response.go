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

// ListFleetResourcesRequest wrapper for the ListFleetResources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListFleetResources.go.html to see an example of how to use ListFleetResourcesRequest.
type ListFleetResourcesRequest struct {

	// Unique Fleet identifier.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// Resource Tenancy Id
	TenancyId *string `mandatory:"false" contributesTo:"query" name:"tenancyId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState FleetResourceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources whose identifier matches the given identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources whose resourceType matches the given resourceType.
	FleetResourceType *string `mandatory:"false" contributesTo:"query" name:"fleetResourceType"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFleetResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListFleetResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFleetResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFleetResourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFleetResourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFleetResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFleetResourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFleetResourceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetFleetResourceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFleetResourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFleetResourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFleetResourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFleetResourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFleetResourcesResponse wrapper for the ListFleetResources operation
type ListFleetResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FleetResourceCollection instances
	FleetResourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFleetResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFleetResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFleetResourcesSortOrderEnum Enum with underlying type: string
type ListFleetResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListFleetResourcesSortOrderEnum
const (
	ListFleetResourcesSortOrderAsc  ListFleetResourcesSortOrderEnum = "ASC"
	ListFleetResourcesSortOrderDesc ListFleetResourcesSortOrderEnum = "DESC"
)

var mappingListFleetResourcesSortOrderEnum = map[string]ListFleetResourcesSortOrderEnum{
	"ASC":  ListFleetResourcesSortOrderAsc,
	"DESC": ListFleetResourcesSortOrderDesc,
}

var mappingListFleetResourcesSortOrderEnumLowerCase = map[string]ListFleetResourcesSortOrderEnum{
	"asc":  ListFleetResourcesSortOrderAsc,
	"desc": ListFleetResourcesSortOrderDesc,
}

// GetListFleetResourcesSortOrderEnumValues Enumerates the set of values for ListFleetResourcesSortOrderEnum
func GetListFleetResourcesSortOrderEnumValues() []ListFleetResourcesSortOrderEnum {
	values := make([]ListFleetResourcesSortOrderEnum, 0)
	for _, v := range mappingListFleetResourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetResourcesSortOrderEnumStringValues Enumerates the set of values in String for ListFleetResourcesSortOrderEnum
func GetListFleetResourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFleetResourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetResourcesSortOrderEnum(val string) (ListFleetResourcesSortOrderEnum, bool) {
	enum, ok := mappingListFleetResourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFleetResourcesSortByEnum Enum with underlying type: string
type ListFleetResourcesSortByEnum string

// Set of constants representing the allowable values for ListFleetResourcesSortByEnum
const (
	ListFleetResourcesSortByTimecreated ListFleetResourcesSortByEnum = "timeCreated"
	ListFleetResourcesSortByDisplayname ListFleetResourcesSortByEnum = "displayName"
)

var mappingListFleetResourcesSortByEnum = map[string]ListFleetResourcesSortByEnum{
	"timeCreated": ListFleetResourcesSortByTimecreated,
	"displayName": ListFleetResourcesSortByDisplayname,
}

var mappingListFleetResourcesSortByEnumLowerCase = map[string]ListFleetResourcesSortByEnum{
	"timecreated": ListFleetResourcesSortByTimecreated,
	"displayname": ListFleetResourcesSortByDisplayname,
}

// GetListFleetResourcesSortByEnumValues Enumerates the set of values for ListFleetResourcesSortByEnum
func GetListFleetResourcesSortByEnumValues() []ListFleetResourcesSortByEnum {
	values := make([]ListFleetResourcesSortByEnum, 0)
	for _, v := range mappingListFleetResourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetResourcesSortByEnumStringValues Enumerates the set of values in String for ListFleetResourcesSortByEnum
func GetListFleetResourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListFleetResourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetResourcesSortByEnum(val string) (ListFleetResourcesSortByEnum, bool) {
	enum, ok := mappingListFleetResourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
