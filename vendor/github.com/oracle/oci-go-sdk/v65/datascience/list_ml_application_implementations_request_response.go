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

// ListMlApplicationImplementationsRequest wrapper for the ListMlApplicationImplementations operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListMlApplicationImplementations.go.html to see an example of how to use ListMlApplicationImplementationsRequest.
type ListMlApplicationImplementationsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// unique MlApplicationImplementation identifier
	MlApplicationImplementationId *string `mandatory:"false" contributesTo:"query" name:"mlApplicationImplementationId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// unique MlApplication identifier
	MlApplicationId *string `mandatory:"false" contributesTo:"query" name:"mlApplicationId"`

	// If it is true search must include all results from descendant compartments. Value true is allowed only if compartmentId refers to root compartment.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// A filter to return only resources with lifecycleState matching the given lifecycleState.
	LifecycleState MlApplicationImplementationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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
	SortOrder ListMlApplicationImplementationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending.
	SortBy ListMlApplicationImplementationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMlApplicationImplementationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMlApplicationImplementationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMlApplicationImplementationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMlApplicationImplementationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMlApplicationImplementationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMlApplicationImplementationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMlApplicationImplementationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMlApplicationImplementationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMlApplicationImplementationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMlApplicationImplementationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMlApplicationImplementationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMlApplicationImplementationsResponse wrapper for the ListMlApplicationImplementations operation
type ListMlApplicationImplementationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MlApplicationImplementationCollection instances
	MlApplicationImplementationCollection `presentIn:"body"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMlApplicationImplementationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMlApplicationImplementationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMlApplicationImplementationsSortOrderEnum Enum with underlying type: string
type ListMlApplicationImplementationsSortOrderEnum string

// Set of constants representing the allowable values for ListMlApplicationImplementationsSortOrderEnum
const (
	ListMlApplicationImplementationsSortOrderAsc  ListMlApplicationImplementationsSortOrderEnum = "ASC"
	ListMlApplicationImplementationsSortOrderDesc ListMlApplicationImplementationsSortOrderEnum = "DESC"
)

var mappingListMlApplicationImplementationsSortOrderEnum = map[string]ListMlApplicationImplementationsSortOrderEnum{
	"ASC":  ListMlApplicationImplementationsSortOrderAsc,
	"DESC": ListMlApplicationImplementationsSortOrderDesc,
}

var mappingListMlApplicationImplementationsSortOrderEnumLowerCase = map[string]ListMlApplicationImplementationsSortOrderEnum{
	"asc":  ListMlApplicationImplementationsSortOrderAsc,
	"desc": ListMlApplicationImplementationsSortOrderDesc,
}

// GetListMlApplicationImplementationsSortOrderEnumValues Enumerates the set of values for ListMlApplicationImplementationsSortOrderEnum
func GetListMlApplicationImplementationsSortOrderEnumValues() []ListMlApplicationImplementationsSortOrderEnum {
	values := make([]ListMlApplicationImplementationsSortOrderEnum, 0)
	for _, v := range mappingListMlApplicationImplementationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMlApplicationImplementationsSortOrderEnumStringValues Enumerates the set of values in String for ListMlApplicationImplementationsSortOrderEnum
func GetListMlApplicationImplementationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMlApplicationImplementationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMlApplicationImplementationsSortOrderEnum(val string) (ListMlApplicationImplementationsSortOrderEnum, bool) {
	enum, ok := mappingListMlApplicationImplementationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMlApplicationImplementationsSortByEnum Enum with underlying type: string
type ListMlApplicationImplementationsSortByEnum string

// Set of constants representing the allowable values for ListMlApplicationImplementationsSortByEnum
const (
	ListMlApplicationImplementationsSortByTimecreated ListMlApplicationImplementationsSortByEnum = "timeCreated"
	ListMlApplicationImplementationsSortByName        ListMlApplicationImplementationsSortByEnum = "name"
)

var mappingListMlApplicationImplementationsSortByEnum = map[string]ListMlApplicationImplementationsSortByEnum{
	"timeCreated": ListMlApplicationImplementationsSortByTimecreated,
	"name":        ListMlApplicationImplementationsSortByName,
}

var mappingListMlApplicationImplementationsSortByEnumLowerCase = map[string]ListMlApplicationImplementationsSortByEnum{
	"timecreated": ListMlApplicationImplementationsSortByTimecreated,
	"name":        ListMlApplicationImplementationsSortByName,
}

// GetListMlApplicationImplementationsSortByEnumValues Enumerates the set of values for ListMlApplicationImplementationsSortByEnum
func GetListMlApplicationImplementationsSortByEnumValues() []ListMlApplicationImplementationsSortByEnum {
	values := make([]ListMlApplicationImplementationsSortByEnum, 0)
	for _, v := range mappingListMlApplicationImplementationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMlApplicationImplementationsSortByEnumStringValues Enumerates the set of values in String for ListMlApplicationImplementationsSortByEnum
func GetListMlApplicationImplementationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListMlApplicationImplementationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMlApplicationImplementationsSortByEnum(val string) (ListMlApplicationImplementationsSortByEnum, bool) {
	enum, ok := mappingListMlApplicationImplementationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
