// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMlApplicationInstancesRequest wrapper for the ListMlApplicationInstances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListMlApplicationInstances.go.html to see an example of how to use ListMlApplicationInstancesRequest.
type ListMlApplicationInstancesRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique MlApplication identifier
	MlApplicationId *string `mandatory:"false" contributesTo:"query" name:"mlApplicationId"`

	// A filter to return only resources matching the given lifecycleState.
	LifecycleState MlApplicationInstanceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 100 is the maximum.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListMlApplicationInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending.
	SortBy ListMlApplicationInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMlApplicationInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMlApplicationInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMlApplicationInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMlApplicationInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMlApplicationInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMlApplicationInstanceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMlApplicationInstanceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMlApplicationInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMlApplicationInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMlApplicationInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMlApplicationInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMlApplicationInstancesResponse wrapper for the ListMlApplicationInstances operation
type ListMlApplicationInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MlApplicationInstanceCollection instances
	MlApplicationInstanceCollection `presentIn:"body"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMlApplicationInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMlApplicationInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMlApplicationInstancesSortOrderEnum Enum with underlying type: string
type ListMlApplicationInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListMlApplicationInstancesSortOrderEnum
const (
	ListMlApplicationInstancesSortOrderAsc  ListMlApplicationInstancesSortOrderEnum = "ASC"
	ListMlApplicationInstancesSortOrderDesc ListMlApplicationInstancesSortOrderEnum = "DESC"
)

var mappingListMlApplicationInstancesSortOrderEnum = map[string]ListMlApplicationInstancesSortOrderEnum{
	"ASC":  ListMlApplicationInstancesSortOrderAsc,
	"DESC": ListMlApplicationInstancesSortOrderDesc,
}

var mappingListMlApplicationInstancesSortOrderEnumLowerCase = map[string]ListMlApplicationInstancesSortOrderEnum{
	"asc":  ListMlApplicationInstancesSortOrderAsc,
	"desc": ListMlApplicationInstancesSortOrderDesc,
}

// GetListMlApplicationInstancesSortOrderEnumValues Enumerates the set of values for ListMlApplicationInstancesSortOrderEnum
func GetListMlApplicationInstancesSortOrderEnumValues() []ListMlApplicationInstancesSortOrderEnum {
	values := make([]ListMlApplicationInstancesSortOrderEnum, 0)
	for _, v := range mappingListMlApplicationInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMlApplicationInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListMlApplicationInstancesSortOrderEnum
func GetListMlApplicationInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMlApplicationInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMlApplicationInstancesSortOrderEnum(val string) (ListMlApplicationInstancesSortOrderEnum, bool) {
	enum, ok := mappingListMlApplicationInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMlApplicationInstancesSortByEnum Enum with underlying type: string
type ListMlApplicationInstancesSortByEnum string

// Set of constants representing the allowable values for ListMlApplicationInstancesSortByEnum
const (
	ListMlApplicationInstancesSortByTimecreated ListMlApplicationInstancesSortByEnum = "timeCreated"
	ListMlApplicationInstancesSortByName        ListMlApplicationInstancesSortByEnum = "name"
)

var mappingListMlApplicationInstancesSortByEnum = map[string]ListMlApplicationInstancesSortByEnum{
	"timeCreated": ListMlApplicationInstancesSortByTimecreated,
	"name":        ListMlApplicationInstancesSortByName,
}

var mappingListMlApplicationInstancesSortByEnumLowerCase = map[string]ListMlApplicationInstancesSortByEnum{
	"timecreated": ListMlApplicationInstancesSortByTimecreated,
	"name":        ListMlApplicationInstancesSortByName,
}

// GetListMlApplicationInstancesSortByEnumValues Enumerates the set of values for ListMlApplicationInstancesSortByEnum
func GetListMlApplicationInstancesSortByEnumValues() []ListMlApplicationInstancesSortByEnum {
	values := make([]ListMlApplicationInstancesSortByEnum, 0)
	for _, v := range mappingListMlApplicationInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMlApplicationInstancesSortByEnumStringValues Enumerates the set of values in String for ListMlApplicationInstancesSortByEnum
func GetListMlApplicationInstancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListMlApplicationInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMlApplicationInstancesSortByEnum(val string) (ListMlApplicationInstancesSortByEnum, bool) {
	enum, ok := mappingListMlApplicationInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
