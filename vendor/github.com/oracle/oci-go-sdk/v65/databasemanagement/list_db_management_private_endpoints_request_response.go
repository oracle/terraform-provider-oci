// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDbManagementPrivateEndpointsRequest wrapper for the ListDbManagementPrivateEndpoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListDbManagementPrivateEndpoints.go.html to see an example of how to use ListDbManagementPrivateEndpointsRequest.
type ListDbManagementPrivateEndpointsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VCN.
	VcnId *string `mandatory:"false" contributesTo:"query" name:"vcnId"`

	// The option to filter Database Management private endpoints that can used for Oracle Databases in a cluster. This should be used along with the vcnId query parameter.
	IsCluster *bool `mandatory:"false" contributesTo:"query" name:"isCluster"`

	// The lifecycle state of a resource.
	LifecycleState ListDbManagementPrivateEndpointsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListDbManagementPrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListDbManagementPrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbManagementPrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbManagementPrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbManagementPrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbManagementPrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDbManagementPrivateEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDbManagementPrivateEndpointsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDbManagementPrivateEndpointsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbManagementPrivateEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDbManagementPrivateEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbManagementPrivateEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDbManagementPrivateEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDbManagementPrivateEndpointsResponse wrapper for the ListDbManagementPrivateEndpoints operation
type ListDbManagementPrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DbManagementPrivateEndpointCollection instances
	DbManagementPrivateEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbManagementPrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbManagementPrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbManagementPrivateEndpointsLifecycleStateEnum Enum with underlying type: string
type ListDbManagementPrivateEndpointsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDbManagementPrivateEndpointsLifecycleStateEnum
const (
	ListDbManagementPrivateEndpointsLifecycleStateCreating ListDbManagementPrivateEndpointsLifecycleStateEnum = "CREATING"
	ListDbManagementPrivateEndpointsLifecycleStateUpdating ListDbManagementPrivateEndpointsLifecycleStateEnum = "UPDATING"
	ListDbManagementPrivateEndpointsLifecycleStateActive   ListDbManagementPrivateEndpointsLifecycleStateEnum = "ACTIVE"
	ListDbManagementPrivateEndpointsLifecycleStateDeleting ListDbManagementPrivateEndpointsLifecycleStateEnum = "DELETING"
	ListDbManagementPrivateEndpointsLifecycleStateDeleted  ListDbManagementPrivateEndpointsLifecycleStateEnum = "DELETED"
	ListDbManagementPrivateEndpointsLifecycleStateFailed   ListDbManagementPrivateEndpointsLifecycleStateEnum = "FAILED"
)

var mappingListDbManagementPrivateEndpointsLifecycleStateEnum = map[string]ListDbManagementPrivateEndpointsLifecycleStateEnum{
	"CREATING": ListDbManagementPrivateEndpointsLifecycleStateCreating,
	"UPDATING": ListDbManagementPrivateEndpointsLifecycleStateUpdating,
	"ACTIVE":   ListDbManagementPrivateEndpointsLifecycleStateActive,
	"DELETING": ListDbManagementPrivateEndpointsLifecycleStateDeleting,
	"DELETED":  ListDbManagementPrivateEndpointsLifecycleStateDeleted,
	"FAILED":   ListDbManagementPrivateEndpointsLifecycleStateFailed,
}

var mappingListDbManagementPrivateEndpointsLifecycleStateEnumLowerCase = map[string]ListDbManagementPrivateEndpointsLifecycleStateEnum{
	"creating": ListDbManagementPrivateEndpointsLifecycleStateCreating,
	"updating": ListDbManagementPrivateEndpointsLifecycleStateUpdating,
	"active":   ListDbManagementPrivateEndpointsLifecycleStateActive,
	"deleting": ListDbManagementPrivateEndpointsLifecycleStateDeleting,
	"deleted":  ListDbManagementPrivateEndpointsLifecycleStateDeleted,
	"failed":   ListDbManagementPrivateEndpointsLifecycleStateFailed,
}

// GetListDbManagementPrivateEndpointsLifecycleStateEnumValues Enumerates the set of values for ListDbManagementPrivateEndpointsLifecycleStateEnum
func GetListDbManagementPrivateEndpointsLifecycleStateEnumValues() []ListDbManagementPrivateEndpointsLifecycleStateEnum {
	values := make([]ListDbManagementPrivateEndpointsLifecycleStateEnum, 0)
	for _, v := range mappingListDbManagementPrivateEndpointsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbManagementPrivateEndpointsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDbManagementPrivateEndpointsLifecycleStateEnum
func GetListDbManagementPrivateEndpointsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListDbManagementPrivateEndpointsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbManagementPrivateEndpointsLifecycleStateEnum(val string) (ListDbManagementPrivateEndpointsLifecycleStateEnum, bool) {
	enum, ok := mappingListDbManagementPrivateEndpointsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbManagementPrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListDbManagementPrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListDbManagementPrivateEndpointsSortOrderEnum
const (
	ListDbManagementPrivateEndpointsSortOrderAsc  ListDbManagementPrivateEndpointsSortOrderEnum = "ASC"
	ListDbManagementPrivateEndpointsSortOrderDesc ListDbManagementPrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListDbManagementPrivateEndpointsSortOrderEnum = map[string]ListDbManagementPrivateEndpointsSortOrderEnum{
	"ASC":  ListDbManagementPrivateEndpointsSortOrderAsc,
	"DESC": ListDbManagementPrivateEndpointsSortOrderDesc,
}

var mappingListDbManagementPrivateEndpointsSortOrderEnumLowerCase = map[string]ListDbManagementPrivateEndpointsSortOrderEnum{
	"asc":  ListDbManagementPrivateEndpointsSortOrderAsc,
	"desc": ListDbManagementPrivateEndpointsSortOrderDesc,
}

// GetListDbManagementPrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListDbManagementPrivateEndpointsSortOrderEnum
func GetListDbManagementPrivateEndpointsSortOrderEnumValues() []ListDbManagementPrivateEndpointsSortOrderEnum {
	values := make([]ListDbManagementPrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListDbManagementPrivateEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbManagementPrivateEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListDbManagementPrivateEndpointsSortOrderEnum
func GetListDbManagementPrivateEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDbManagementPrivateEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbManagementPrivateEndpointsSortOrderEnum(val string) (ListDbManagementPrivateEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListDbManagementPrivateEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbManagementPrivateEndpointsSortByEnum Enum with underlying type: string
type ListDbManagementPrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListDbManagementPrivateEndpointsSortByEnum
const (
	ListDbManagementPrivateEndpointsSortByTimecreated ListDbManagementPrivateEndpointsSortByEnum = "TIMECREATED"
	ListDbManagementPrivateEndpointsSortByName        ListDbManagementPrivateEndpointsSortByEnum = "NAME"
)

var mappingListDbManagementPrivateEndpointsSortByEnum = map[string]ListDbManagementPrivateEndpointsSortByEnum{
	"TIMECREATED": ListDbManagementPrivateEndpointsSortByTimecreated,
	"NAME":        ListDbManagementPrivateEndpointsSortByName,
}

var mappingListDbManagementPrivateEndpointsSortByEnumLowerCase = map[string]ListDbManagementPrivateEndpointsSortByEnum{
	"timecreated": ListDbManagementPrivateEndpointsSortByTimecreated,
	"name":        ListDbManagementPrivateEndpointsSortByName,
}

// GetListDbManagementPrivateEndpointsSortByEnumValues Enumerates the set of values for ListDbManagementPrivateEndpointsSortByEnum
func GetListDbManagementPrivateEndpointsSortByEnumValues() []ListDbManagementPrivateEndpointsSortByEnum {
	values := make([]ListDbManagementPrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListDbManagementPrivateEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbManagementPrivateEndpointsSortByEnumStringValues Enumerates the set of values in String for ListDbManagementPrivateEndpointsSortByEnum
func GetListDbManagementPrivateEndpointsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListDbManagementPrivateEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbManagementPrivateEndpointsSortByEnum(val string) (ListDbManagementPrivateEndpointsSortByEnum, bool) {
	enum, ok := mappingListDbManagementPrivateEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
