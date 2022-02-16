// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListCatalogPrivateEndpointsRequest wrapper for the ListCatalogPrivateEndpoints operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListCatalogPrivateEndpoints.go.html to see an example of how to use ListCatalogPrivateEndpointsRequest.
type ListCatalogPrivateEndpointsRequest struct {

	// The OCID of the compartment where you want to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListCatalogPrivateEndpointsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListCatalogPrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListCatalogPrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCatalogPrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCatalogPrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCatalogPrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCatalogPrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCatalogPrivateEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCatalogPrivateEndpointsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListCatalogPrivateEndpointsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCatalogPrivateEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCatalogPrivateEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCatalogPrivateEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCatalogPrivateEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCatalogPrivateEndpointsResponse wrapper for the ListCatalogPrivateEndpoints operation
type ListCatalogPrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CatalogPrivateEndpointSummary instances
	Items []CatalogPrivateEndpointSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCatalogPrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCatalogPrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCatalogPrivateEndpointsLifecycleStateEnum Enum with underlying type: string
type ListCatalogPrivateEndpointsLifecycleStateEnum string

// Set of constants representing the allowable values for ListCatalogPrivateEndpointsLifecycleStateEnum
const (
	ListCatalogPrivateEndpointsLifecycleStateCreating ListCatalogPrivateEndpointsLifecycleStateEnum = "CREATING"
	ListCatalogPrivateEndpointsLifecycleStateActive   ListCatalogPrivateEndpointsLifecycleStateEnum = "ACTIVE"
	ListCatalogPrivateEndpointsLifecycleStateInactive ListCatalogPrivateEndpointsLifecycleStateEnum = "INACTIVE"
	ListCatalogPrivateEndpointsLifecycleStateUpdating ListCatalogPrivateEndpointsLifecycleStateEnum = "UPDATING"
	ListCatalogPrivateEndpointsLifecycleStateDeleting ListCatalogPrivateEndpointsLifecycleStateEnum = "DELETING"
	ListCatalogPrivateEndpointsLifecycleStateDeleted  ListCatalogPrivateEndpointsLifecycleStateEnum = "DELETED"
	ListCatalogPrivateEndpointsLifecycleStateFailed   ListCatalogPrivateEndpointsLifecycleStateEnum = "FAILED"
	ListCatalogPrivateEndpointsLifecycleStateMoving   ListCatalogPrivateEndpointsLifecycleStateEnum = "MOVING"
)

var mappingListCatalogPrivateEndpointsLifecycleStateEnum = map[string]ListCatalogPrivateEndpointsLifecycleStateEnum{
	"CREATING": ListCatalogPrivateEndpointsLifecycleStateCreating,
	"ACTIVE":   ListCatalogPrivateEndpointsLifecycleStateActive,
	"INACTIVE": ListCatalogPrivateEndpointsLifecycleStateInactive,
	"UPDATING": ListCatalogPrivateEndpointsLifecycleStateUpdating,
	"DELETING": ListCatalogPrivateEndpointsLifecycleStateDeleting,
	"DELETED":  ListCatalogPrivateEndpointsLifecycleStateDeleted,
	"FAILED":   ListCatalogPrivateEndpointsLifecycleStateFailed,
	"MOVING":   ListCatalogPrivateEndpointsLifecycleStateMoving,
}

// GetListCatalogPrivateEndpointsLifecycleStateEnumValues Enumerates the set of values for ListCatalogPrivateEndpointsLifecycleStateEnum
func GetListCatalogPrivateEndpointsLifecycleStateEnumValues() []ListCatalogPrivateEndpointsLifecycleStateEnum {
	values := make([]ListCatalogPrivateEndpointsLifecycleStateEnum, 0)
	for _, v := range mappingListCatalogPrivateEndpointsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListCatalogPrivateEndpointsLifecycleStateEnumStringValues Enumerates the set of values in String for ListCatalogPrivateEndpointsLifecycleStateEnum
func GetListCatalogPrivateEndpointsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"MOVING",
	}
}

// GetMappingListCatalogPrivateEndpointsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCatalogPrivateEndpointsLifecycleStateEnum(val string) (ListCatalogPrivateEndpointsLifecycleStateEnum, bool) {
	mappingListCatalogPrivateEndpointsLifecycleStateEnumIgnoreCase := make(map[string]ListCatalogPrivateEndpointsLifecycleStateEnum)
	for k, v := range mappingListCatalogPrivateEndpointsLifecycleStateEnum {
		mappingListCatalogPrivateEndpointsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListCatalogPrivateEndpointsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListCatalogPrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListCatalogPrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListCatalogPrivateEndpointsSortOrderEnum
const (
	ListCatalogPrivateEndpointsSortOrderAsc  ListCatalogPrivateEndpointsSortOrderEnum = "ASC"
	ListCatalogPrivateEndpointsSortOrderDesc ListCatalogPrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListCatalogPrivateEndpointsSortOrderEnum = map[string]ListCatalogPrivateEndpointsSortOrderEnum{
	"ASC":  ListCatalogPrivateEndpointsSortOrderAsc,
	"DESC": ListCatalogPrivateEndpointsSortOrderDesc,
}

// GetListCatalogPrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListCatalogPrivateEndpointsSortOrderEnum
func GetListCatalogPrivateEndpointsSortOrderEnumValues() []ListCatalogPrivateEndpointsSortOrderEnum {
	values := make([]ListCatalogPrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListCatalogPrivateEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCatalogPrivateEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListCatalogPrivateEndpointsSortOrderEnum
func GetListCatalogPrivateEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCatalogPrivateEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCatalogPrivateEndpointsSortOrderEnum(val string) (ListCatalogPrivateEndpointsSortOrderEnum, bool) {
	mappingListCatalogPrivateEndpointsSortOrderEnumIgnoreCase := make(map[string]ListCatalogPrivateEndpointsSortOrderEnum)
	for k, v := range mappingListCatalogPrivateEndpointsSortOrderEnum {
		mappingListCatalogPrivateEndpointsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListCatalogPrivateEndpointsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListCatalogPrivateEndpointsSortByEnum Enum with underlying type: string
type ListCatalogPrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListCatalogPrivateEndpointsSortByEnum
const (
	ListCatalogPrivateEndpointsSortByTimecreated ListCatalogPrivateEndpointsSortByEnum = "TIMECREATED"
	ListCatalogPrivateEndpointsSortByDisplayname ListCatalogPrivateEndpointsSortByEnum = "DISPLAYNAME"
)

var mappingListCatalogPrivateEndpointsSortByEnum = map[string]ListCatalogPrivateEndpointsSortByEnum{
	"TIMECREATED": ListCatalogPrivateEndpointsSortByTimecreated,
	"DISPLAYNAME": ListCatalogPrivateEndpointsSortByDisplayname,
}

// GetListCatalogPrivateEndpointsSortByEnumValues Enumerates the set of values for ListCatalogPrivateEndpointsSortByEnum
func GetListCatalogPrivateEndpointsSortByEnumValues() []ListCatalogPrivateEndpointsSortByEnum {
	values := make([]ListCatalogPrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListCatalogPrivateEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCatalogPrivateEndpointsSortByEnumStringValues Enumerates the set of values in String for ListCatalogPrivateEndpointsSortByEnum
func GetListCatalogPrivateEndpointsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCatalogPrivateEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCatalogPrivateEndpointsSortByEnum(val string) (ListCatalogPrivateEndpointsSortByEnum, bool) {
	mappingListCatalogPrivateEndpointsSortByEnumIgnoreCase := make(map[string]ListCatalogPrivateEndpointsSortByEnum)
	for k, v := range mappingListCatalogPrivateEndpointsSortByEnum {
		mappingListCatalogPrivateEndpointsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListCatalogPrivateEndpointsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
