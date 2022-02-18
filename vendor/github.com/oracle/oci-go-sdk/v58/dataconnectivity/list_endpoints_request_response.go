// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListEndpointsRequest wrapper for the ListEndpoints operation
type ListEndpointsRequest struct {

	// The OCID of the compartment containing the resources you want to list.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// DCMS registry id
	RegistryId *string `mandatory:"false" contributesTo:"query" name:"registryId"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Lifecycle state of the resource.
	LifecycleState RegistryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// This parameter allows users to specify a sort field.  Default sort order is the descending order of `timeCreated` (most recently created objects at the top).  Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingRegistryLifecycleStateEnum[string(request.LifecycleState)]; !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRegistryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := mappingListEndpointsSortOrderEnum[string(request.SortOrder)]; !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := mappingListEndpointsSortByEnum[string(request.SortBy)]; !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEndpointsResponse wrapper for the ListEndpoints operation
type ListEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EndpointSummaryCollection instances
	EndpointSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `Endpoint`s. If this header appears in the response, then this
	// is a partial list of Registries. Include this value as the `page` parameter in a subsequent
	// GET request to get the next batch of Endpoints.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEndpointsSortOrderEnum Enum with underlying type: string
type ListEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListEndpointsSortOrderEnum
const (
	ListEndpointsSortOrderAsc  ListEndpointsSortOrderEnum = "ASC"
	ListEndpointsSortOrderDesc ListEndpointsSortOrderEnum = "DESC"
)

var mappingListEndpointsSortOrderEnum = map[string]ListEndpointsSortOrderEnum{
	"ASC":  ListEndpointsSortOrderAsc,
	"DESC": ListEndpointsSortOrderDesc,
}

// GetListEndpointsSortOrderEnumValues Enumerates the set of values for ListEndpointsSortOrderEnum
func GetListEndpointsSortOrderEnumValues() []ListEndpointsSortOrderEnum {
	values := make([]ListEndpointsSortOrderEnum, 0)
	for _, v := range mappingListEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListEndpointsSortOrderEnum
func GetListEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// ListEndpointsSortByEnum Enum with underlying type: string
type ListEndpointsSortByEnum string

// Set of constants representing the allowable values for ListEndpointsSortByEnum
const (
	ListEndpointsSortByTimecreated ListEndpointsSortByEnum = "TIMECREATED"
	ListEndpointsSortByDisplayname ListEndpointsSortByEnum = "DISPLAYNAME"
	ListEndpointsSortByTimeupdated ListEndpointsSortByEnum = "TIMEUPDATED"
)

var mappingListEndpointsSortByEnum = map[string]ListEndpointsSortByEnum{
	"TIMECREATED": ListEndpointsSortByTimecreated,
	"DISPLAYNAME": ListEndpointsSortByDisplayname,
	"TIMEUPDATED": ListEndpointsSortByTimeupdated,
}

// GetListEndpointsSortByEnumValues Enumerates the set of values for ListEndpointsSortByEnum
func GetListEndpointsSortByEnumValues() []ListEndpointsSortByEnum {
	values := make([]ListEndpointsSortByEnum, 0)
	for _, v := range mappingListEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEndpointsSortByEnumStringValues Enumerates the set of values in String for ListEndpointsSortByEnum
func GetListEndpointsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
		"TIMEUPDATED",
	}
}
