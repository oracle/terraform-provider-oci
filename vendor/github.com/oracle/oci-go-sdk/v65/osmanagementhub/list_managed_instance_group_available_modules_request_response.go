// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedInstanceGroupAvailableModulesRequest wrapper for the ListManagedInstanceGroupAvailableModules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceGroupAvailableModules.go.html to see an example of how to use ListManagedInstanceGroupAvailableModulesRequest.
type ListManagedInstanceGroupAvailableModulesRequest struct {

	// The managed instance group OCID.
	ManagedInstanceGroupId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceGroupId"`

	// The OCID of the compartment that contains the resources to list.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return resources that may partially match the name given.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListManagedInstanceGroupAvailableModulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for name is ascending.
	SortBy ListManagedInstanceGroupAvailableModulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceGroupAvailableModulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceGroupAvailableModulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceGroupAvailableModulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceGroupAvailableModulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceGroupAvailableModulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceGroupAvailableModulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceGroupAvailableModulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceGroupAvailableModulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceGroupAvailableModulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceGroupAvailableModulesResponse wrapper for the ListManagedInstanceGroupAvailableModules operation
type ListManagedInstanceGroupAvailableModulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceGroupAvailableModuleCollection instances
	ManagedInstanceGroupAvailableModuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceGroupAvailableModulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceGroupAvailableModulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceGroupAvailableModulesSortOrderEnum Enum with underlying type: string
type ListManagedInstanceGroupAvailableModulesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupAvailableModulesSortOrderEnum
const (
	ListManagedInstanceGroupAvailableModulesSortOrderAsc  ListManagedInstanceGroupAvailableModulesSortOrderEnum = "ASC"
	ListManagedInstanceGroupAvailableModulesSortOrderDesc ListManagedInstanceGroupAvailableModulesSortOrderEnum = "DESC"
)

var mappingListManagedInstanceGroupAvailableModulesSortOrderEnum = map[string]ListManagedInstanceGroupAvailableModulesSortOrderEnum{
	"ASC":  ListManagedInstanceGroupAvailableModulesSortOrderAsc,
	"DESC": ListManagedInstanceGroupAvailableModulesSortOrderDesc,
}

var mappingListManagedInstanceGroupAvailableModulesSortOrderEnumLowerCase = map[string]ListManagedInstanceGroupAvailableModulesSortOrderEnum{
	"asc":  ListManagedInstanceGroupAvailableModulesSortOrderAsc,
	"desc": ListManagedInstanceGroupAvailableModulesSortOrderDesc,
}

// GetListManagedInstanceGroupAvailableModulesSortOrderEnumValues Enumerates the set of values for ListManagedInstanceGroupAvailableModulesSortOrderEnum
func GetListManagedInstanceGroupAvailableModulesSortOrderEnumValues() []ListManagedInstanceGroupAvailableModulesSortOrderEnum {
	values := make([]ListManagedInstanceGroupAvailableModulesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceGroupAvailableModulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupAvailableModulesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupAvailableModulesSortOrderEnum
func GetListManagedInstanceGroupAvailableModulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceGroupAvailableModulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupAvailableModulesSortOrderEnum(val string) (ListManagedInstanceGroupAvailableModulesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupAvailableModulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceGroupAvailableModulesSortByEnum Enum with underlying type: string
type ListManagedInstanceGroupAvailableModulesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupAvailableModulesSortByEnum
const (
	ListManagedInstanceGroupAvailableModulesSortByName ListManagedInstanceGroupAvailableModulesSortByEnum = "name"
)

var mappingListManagedInstanceGroupAvailableModulesSortByEnum = map[string]ListManagedInstanceGroupAvailableModulesSortByEnum{
	"name": ListManagedInstanceGroupAvailableModulesSortByName,
}

var mappingListManagedInstanceGroupAvailableModulesSortByEnumLowerCase = map[string]ListManagedInstanceGroupAvailableModulesSortByEnum{
	"name": ListManagedInstanceGroupAvailableModulesSortByName,
}

// GetListManagedInstanceGroupAvailableModulesSortByEnumValues Enumerates the set of values for ListManagedInstanceGroupAvailableModulesSortByEnum
func GetListManagedInstanceGroupAvailableModulesSortByEnumValues() []ListManagedInstanceGroupAvailableModulesSortByEnum {
	values := make([]ListManagedInstanceGroupAvailableModulesSortByEnum, 0)
	for _, v := range mappingListManagedInstanceGroupAvailableModulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupAvailableModulesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupAvailableModulesSortByEnum
func GetListManagedInstanceGroupAvailableModulesSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListManagedInstanceGroupAvailableModulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupAvailableModulesSortByEnum(val string) (ListManagedInstanceGroupAvailableModulesSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupAvailableModulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
