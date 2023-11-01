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

// ListManagedInstanceModulesRequest wrapper for the ListManagedInstanceModules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceModules.go.html to see an example of how to use ListManagedInstanceModulesRequest.
type ListManagedInstanceModulesRequest struct {

	// The OCID of the managed instance.
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

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
	SortOrder ListManagedInstanceModulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for name is ascending.
	SortBy ListManagedInstanceModulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceModulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceModulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceModulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceModulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceModulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceModulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceModulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceModulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceModulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceModulesResponse wrapper for the ListManagedInstanceModules operation
type ListManagedInstanceModulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceModuleCollection instances
	ManagedInstanceModuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the asynchronous work. You can use this to query its status.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceModulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceModulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceModulesSortOrderEnum Enum with underlying type: string
type ListManagedInstanceModulesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceModulesSortOrderEnum
const (
	ListManagedInstanceModulesSortOrderAsc  ListManagedInstanceModulesSortOrderEnum = "ASC"
	ListManagedInstanceModulesSortOrderDesc ListManagedInstanceModulesSortOrderEnum = "DESC"
)

var mappingListManagedInstanceModulesSortOrderEnum = map[string]ListManagedInstanceModulesSortOrderEnum{
	"ASC":  ListManagedInstanceModulesSortOrderAsc,
	"DESC": ListManagedInstanceModulesSortOrderDesc,
}

var mappingListManagedInstanceModulesSortOrderEnumLowerCase = map[string]ListManagedInstanceModulesSortOrderEnum{
	"asc":  ListManagedInstanceModulesSortOrderAsc,
	"desc": ListManagedInstanceModulesSortOrderDesc,
}

// GetListManagedInstanceModulesSortOrderEnumValues Enumerates the set of values for ListManagedInstanceModulesSortOrderEnum
func GetListManagedInstanceModulesSortOrderEnumValues() []ListManagedInstanceModulesSortOrderEnum {
	values := make([]ListManagedInstanceModulesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceModulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceModulesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceModulesSortOrderEnum
func GetListManagedInstanceModulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceModulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceModulesSortOrderEnum(val string) (ListManagedInstanceModulesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceModulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceModulesSortByEnum Enum with underlying type: string
type ListManagedInstanceModulesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceModulesSortByEnum
const (
	ListManagedInstanceModulesSortByName ListManagedInstanceModulesSortByEnum = "name"
)

var mappingListManagedInstanceModulesSortByEnum = map[string]ListManagedInstanceModulesSortByEnum{
	"name": ListManagedInstanceModulesSortByName,
}

var mappingListManagedInstanceModulesSortByEnumLowerCase = map[string]ListManagedInstanceModulesSortByEnum{
	"name": ListManagedInstanceModulesSortByName,
}

// GetListManagedInstanceModulesSortByEnumValues Enumerates the set of values for ListManagedInstanceModulesSortByEnum
func GetListManagedInstanceModulesSortByEnumValues() []ListManagedInstanceModulesSortByEnum {
	values := make([]ListManagedInstanceModulesSortByEnum, 0)
	for _, v := range mappingListManagedInstanceModulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceModulesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceModulesSortByEnum
func GetListManagedInstanceModulesSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListManagedInstanceModulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceModulesSortByEnum(val string) (ListManagedInstanceModulesSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceModulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
