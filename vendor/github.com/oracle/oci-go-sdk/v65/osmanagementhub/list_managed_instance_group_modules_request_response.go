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

// ListManagedInstanceGroupModulesRequest wrapper for the ListManagedInstanceGroupModules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceGroupModules.go.html to see an example of how to use ListManagedInstanceGroupModulesRequest.
type ListManagedInstanceGroupModulesRequest struct {

	// The managed instance group OCID.
	ManagedInstanceGroupId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceGroupId"`

	// The OCID of the compartment that contains the resources to list.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return resources that may partially match the name given.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// The name of the stream of the containing module.  This parameter
	// is required if a profileName is specified.
	StreamName *string `mandatory:"false" contributesTo:"query" name:"streamName"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListManagedInstanceGroupModulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for name is ascending.
	SortBy ListManagedInstanceGroupModulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceGroupModulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceGroupModulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceGroupModulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceGroupModulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceGroupModulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceGroupModulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceGroupModulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceGroupModulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceGroupModulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceGroupModulesResponse wrapper for the ListManagedInstanceGroupModules operation
type ListManagedInstanceGroupModulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceGroupModuleCollection instances
	ManagedInstanceGroupModuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the asynchronous work. You can use this to query its status.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceGroupModulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceGroupModulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceGroupModulesSortOrderEnum Enum with underlying type: string
type ListManagedInstanceGroupModulesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupModulesSortOrderEnum
const (
	ListManagedInstanceGroupModulesSortOrderAsc  ListManagedInstanceGroupModulesSortOrderEnum = "ASC"
	ListManagedInstanceGroupModulesSortOrderDesc ListManagedInstanceGroupModulesSortOrderEnum = "DESC"
)

var mappingListManagedInstanceGroupModulesSortOrderEnum = map[string]ListManagedInstanceGroupModulesSortOrderEnum{
	"ASC":  ListManagedInstanceGroupModulesSortOrderAsc,
	"DESC": ListManagedInstanceGroupModulesSortOrderDesc,
}

var mappingListManagedInstanceGroupModulesSortOrderEnumLowerCase = map[string]ListManagedInstanceGroupModulesSortOrderEnum{
	"asc":  ListManagedInstanceGroupModulesSortOrderAsc,
	"desc": ListManagedInstanceGroupModulesSortOrderDesc,
}

// GetListManagedInstanceGroupModulesSortOrderEnumValues Enumerates the set of values for ListManagedInstanceGroupModulesSortOrderEnum
func GetListManagedInstanceGroupModulesSortOrderEnumValues() []ListManagedInstanceGroupModulesSortOrderEnum {
	values := make([]ListManagedInstanceGroupModulesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceGroupModulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupModulesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupModulesSortOrderEnum
func GetListManagedInstanceGroupModulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceGroupModulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupModulesSortOrderEnum(val string) (ListManagedInstanceGroupModulesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupModulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceGroupModulesSortByEnum Enum with underlying type: string
type ListManagedInstanceGroupModulesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupModulesSortByEnum
const (
	ListManagedInstanceGroupModulesSortByName ListManagedInstanceGroupModulesSortByEnum = "name"
)

var mappingListManagedInstanceGroupModulesSortByEnum = map[string]ListManagedInstanceGroupModulesSortByEnum{
	"name": ListManagedInstanceGroupModulesSortByName,
}

var mappingListManagedInstanceGroupModulesSortByEnumLowerCase = map[string]ListManagedInstanceGroupModulesSortByEnum{
	"name": ListManagedInstanceGroupModulesSortByName,
}

// GetListManagedInstanceGroupModulesSortByEnumValues Enumerates the set of values for ListManagedInstanceGroupModulesSortByEnum
func GetListManagedInstanceGroupModulesSortByEnumValues() []ListManagedInstanceGroupModulesSortByEnum {
	values := make([]ListManagedInstanceGroupModulesSortByEnum, 0)
	for _, v := range mappingListManagedInstanceGroupModulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupModulesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupModulesSortByEnum
func GetListManagedInstanceGroupModulesSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListManagedInstanceGroupModulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupModulesSortByEnum(val string) (ListManagedInstanceGroupModulesSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupModulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
