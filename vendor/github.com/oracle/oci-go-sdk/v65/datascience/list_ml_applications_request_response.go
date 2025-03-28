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

// ListMlApplicationsRequest wrapper for the ListMlApplications operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListMlApplications.go.html to see an example of how to use ListMlApplicationsRequest.
type ListMlApplicationsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// unique MlApplication identifier
	MlApplicationId *string `mandatory:"false" contributesTo:"query" name:"mlApplicationId"`

	// If it is true search must include all results from descendant compartments. Value true is allowed only if compartmentId refers to root compartment.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only resources with lifecycleState matching the given lifecycleState.
	LifecycleState MlApplicationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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
	SortOrder ListMlApplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending.
	SortBy ListMlApplicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMlApplicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMlApplicationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMlApplicationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMlApplicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMlApplicationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMlApplicationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMlApplicationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMlApplicationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMlApplicationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMlApplicationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMlApplicationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMlApplicationsResponse wrapper for the ListMlApplications operation
type ListMlApplicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MlApplicationCollection instances
	MlApplicationCollection `presentIn:"body"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMlApplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMlApplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMlApplicationsSortOrderEnum Enum with underlying type: string
type ListMlApplicationsSortOrderEnum string

// Set of constants representing the allowable values for ListMlApplicationsSortOrderEnum
const (
	ListMlApplicationsSortOrderAsc  ListMlApplicationsSortOrderEnum = "ASC"
	ListMlApplicationsSortOrderDesc ListMlApplicationsSortOrderEnum = "DESC"
)

var mappingListMlApplicationsSortOrderEnum = map[string]ListMlApplicationsSortOrderEnum{
	"ASC":  ListMlApplicationsSortOrderAsc,
	"DESC": ListMlApplicationsSortOrderDesc,
}

var mappingListMlApplicationsSortOrderEnumLowerCase = map[string]ListMlApplicationsSortOrderEnum{
	"asc":  ListMlApplicationsSortOrderAsc,
	"desc": ListMlApplicationsSortOrderDesc,
}

// GetListMlApplicationsSortOrderEnumValues Enumerates the set of values for ListMlApplicationsSortOrderEnum
func GetListMlApplicationsSortOrderEnumValues() []ListMlApplicationsSortOrderEnum {
	values := make([]ListMlApplicationsSortOrderEnum, 0)
	for _, v := range mappingListMlApplicationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMlApplicationsSortOrderEnumStringValues Enumerates the set of values in String for ListMlApplicationsSortOrderEnum
func GetListMlApplicationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMlApplicationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMlApplicationsSortOrderEnum(val string) (ListMlApplicationsSortOrderEnum, bool) {
	enum, ok := mappingListMlApplicationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMlApplicationsSortByEnum Enum with underlying type: string
type ListMlApplicationsSortByEnum string

// Set of constants representing the allowable values for ListMlApplicationsSortByEnum
const (
	ListMlApplicationsSortByTimecreated ListMlApplicationsSortByEnum = "timeCreated"
	ListMlApplicationsSortByName        ListMlApplicationsSortByEnum = "name"
)

var mappingListMlApplicationsSortByEnum = map[string]ListMlApplicationsSortByEnum{
	"timeCreated": ListMlApplicationsSortByTimecreated,
	"name":        ListMlApplicationsSortByName,
}

var mappingListMlApplicationsSortByEnumLowerCase = map[string]ListMlApplicationsSortByEnum{
	"timecreated": ListMlApplicationsSortByTimecreated,
	"name":        ListMlApplicationsSortByName,
}

// GetListMlApplicationsSortByEnumValues Enumerates the set of values for ListMlApplicationsSortByEnum
func GetListMlApplicationsSortByEnumValues() []ListMlApplicationsSortByEnum {
	values := make([]ListMlApplicationsSortByEnum, 0)
	for _, v := range mappingListMlApplicationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMlApplicationsSortByEnumStringValues Enumerates the set of values in String for ListMlApplicationsSortByEnum
func GetListMlApplicationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListMlApplicationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMlApplicationsSortByEnum(val string) (ListMlApplicationsSortByEnum, bool) {
	enum, ok := mappingListMlApplicationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
