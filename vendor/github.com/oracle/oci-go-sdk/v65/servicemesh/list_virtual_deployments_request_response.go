// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListVirtualDeploymentsRequest wrapper for the ListVirtualDeployments operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListVirtualDeployments.go.html to see an example of how to use ListVirtualDeploymentsRequest.
type ListVirtualDeploymentsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListVirtualDeploymentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for 'timeCreated' is descending. Default order for 'name' is ascending.
	SortBy ListVirtualDeploymentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Unique VirtualService identifier.
	VirtualServiceId *string `mandatory:"false" contributesTo:"query" name:"virtualServiceId"`

	// Unique VirtualDeployment identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the life cycle state given.
	LifecycleState VirtualDeploymentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVirtualDeploymentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVirtualDeploymentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVirtualDeploymentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVirtualDeploymentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVirtualDeploymentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListVirtualDeploymentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVirtualDeploymentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVirtualDeploymentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVirtualDeploymentsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVirtualDeploymentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetVirtualDeploymentLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVirtualDeploymentsResponse wrapper for the ListVirtualDeployments operation
type ListVirtualDeploymentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of VirtualDeploymentCollection instances
	VirtualDeploymentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVirtualDeploymentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVirtualDeploymentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVirtualDeploymentsSortOrderEnum Enum with underlying type: string
type ListVirtualDeploymentsSortOrderEnum string

// Set of constants representing the allowable values for ListVirtualDeploymentsSortOrderEnum
const (
	ListVirtualDeploymentsSortOrderAsc  ListVirtualDeploymentsSortOrderEnum = "ASC"
	ListVirtualDeploymentsSortOrderDesc ListVirtualDeploymentsSortOrderEnum = "DESC"
)

var mappingListVirtualDeploymentsSortOrderEnum = map[string]ListVirtualDeploymentsSortOrderEnum{
	"ASC":  ListVirtualDeploymentsSortOrderAsc,
	"DESC": ListVirtualDeploymentsSortOrderDesc,
}

var mappingListVirtualDeploymentsSortOrderEnumLowerCase = map[string]ListVirtualDeploymentsSortOrderEnum{
	"asc":  ListVirtualDeploymentsSortOrderAsc,
	"desc": ListVirtualDeploymentsSortOrderDesc,
}

// GetListVirtualDeploymentsSortOrderEnumValues Enumerates the set of values for ListVirtualDeploymentsSortOrderEnum
func GetListVirtualDeploymentsSortOrderEnumValues() []ListVirtualDeploymentsSortOrderEnum {
	values := make([]ListVirtualDeploymentsSortOrderEnum, 0)
	for _, v := range mappingListVirtualDeploymentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVirtualDeploymentsSortOrderEnumStringValues Enumerates the set of values in String for ListVirtualDeploymentsSortOrderEnum
func GetListVirtualDeploymentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVirtualDeploymentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVirtualDeploymentsSortOrderEnum(val string) (ListVirtualDeploymentsSortOrderEnum, bool) {
	enum, ok := mappingListVirtualDeploymentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVirtualDeploymentsSortByEnum Enum with underlying type: string
type ListVirtualDeploymentsSortByEnum string

// Set of constants representing the allowable values for ListVirtualDeploymentsSortByEnum
const (
	ListVirtualDeploymentsSortById          ListVirtualDeploymentsSortByEnum = "id"
	ListVirtualDeploymentsSortByTimecreated ListVirtualDeploymentsSortByEnum = "timeCreated"
	ListVirtualDeploymentsSortByName        ListVirtualDeploymentsSortByEnum = "name"
)

var mappingListVirtualDeploymentsSortByEnum = map[string]ListVirtualDeploymentsSortByEnum{
	"id":          ListVirtualDeploymentsSortById,
	"timeCreated": ListVirtualDeploymentsSortByTimecreated,
	"name":        ListVirtualDeploymentsSortByName,
}

var mappingListVirtualDeploymentsSortByEnumLowerCase = map[string]ListVirtualDeploymentsSortByEnum{
	"id":          ListVirtualDeploymentsSortById,
	"timecreated": ListVirtualDeploymentsSortByTimecreated,
	"name":        ListVirtualDeploymentsSortByName,
}

// GetListVirtualDeploymentsSortByEnumValues Enumerates the set of values for ListVirtualDeploymentsSortByEnum
func GetListVirtualDeploymentsSortByEnumValues() []ListVirtualDeploymentsSortByEnum {
	values := make([]ListVirtualDeploymentsSortByEnum, 0)
	for _, v := range mappingListVirtualDeploymentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVirtualDeploymentsSortByEnumStringValues Enumerates the set of values in String for ListVirtualDeploymentsSortByEnum
func GetListVirtualDeploymentsSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeCreated",
		"name",
	}
}

// GetMappingListVirtualDeploymentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVirtualDeploymentsSortByEnum(val string) (ListVirtualDeploymentsSortByEnum, bool) {
	enum, ok := mappingListVirtualDeploymentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
