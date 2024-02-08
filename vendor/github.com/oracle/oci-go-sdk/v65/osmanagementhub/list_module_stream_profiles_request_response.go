// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListModuleStreamProfilesRequest wrapper for the ListModuleStreamProfiles operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListModuleStreamProfiles.go.html to see an example of how to use ListModuleStreamProfilesRequest.
type ListModuleStreamProfilesRequest struct {

	// The software source OCID.
	SoftwareSourceId *string `mandatory:"true" contributesTo:"path" name:"softwareSourceId"`

	// The name of a module. This parameter is required if a
	// streamName is specified.
	ModuleName *string `mandatory:"false" contributesTo:"query" name:"moduleName"`

	// The name of the stream of the containing module.  This parameter
	// is required if a profileName is specified.
	StreamName *string `mandatory:"false" contributesTo:"query" name:"streamName"`

	// The name of the entity to be queried.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListModuleStreamProfilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for moduleName is ascending.
	SortBy ListModuleStreamProfilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModuleStreamProfilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModuleStreamProfilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModuleStreamProfilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModuleStreamProfilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListModuleStreamProfilesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListModuleStreamProfilesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListModuleStreamProfilesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModuleStreamProfilesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListModuleStreamProfilesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListModuleStreamProfilesResponse wrapper for the ListModuleStreamProfiles operation
type ListModuleStreamProfilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ModuleStreamProfileCollection instances
	ModuleStreamProfileCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListModuleStreamProfilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModuleStreamProfilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModuleStreamProfilesSortOrderEnum Enum with underlying type: string
type ListModuleStreamProfilesSortOrderEnum string

// Set of constants representing the allowable values for ListModuleStreamProfilesSortOrderEnum
const (
	ListModuleStreamProfilesSortOrderAsc  ListModuleStreamProfilesSortOrderEnum = "ASC"
	ListModuleStreamProfilesSortOrderDesc ListModuleStreamProfilesSortOrderEnum = "DESC"
)

var mappingListModuleStreamProfilesSortOrderEnum = map[string]ListModuleStreamProfilesSortOrderEnum{
	"ASC":  ListModuleStreamProfilesSortOrderAsc,
	"DESC": ListModuleStreamProfilesSortOrderDesc,
}

var mappingListModuleStreamProfilesSortOrderEnumLowerCase = map[string]ListModuleStreamProfilesSortOrderEnum{
	"asc":  ListModuleStreamProfilesSortOrderAsc,
	"desc": ListModuleStreamProfilesSortOrderDesc,
}

// GetListModuleStreamProfilesSortOrderEnumValues Enumerates the set of values for ListModuleStreamProfilesSortOrderEnum
func GetListModuleStreamProfilesSortOrderEnumValues() []ListModuleStreamProfilesSortOrderEnum {
	values := make([]ListModuleStreamProfilesSortOrderEnum, 0)
	for _, v := range mappingListModuleStreamProfilesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListModuleStreamProfilesSortOrderEnumStringValues Enumerates the set of values in String for ListModuleStreamProfilesSortOrderEnum
func GetListModuleStreamProfilesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListModuleStreamProfilesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModuleStreamProfilesSortOrderEnum(val string) (ListModuleStreamProfilesSortOrderEnum, bool) {
	enum, ok := mappingListModuleStreamProfilesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModuleStreamProfilesSortByEnum Enum with underlying type: string
type ListModuleStreamProfilesSortByEnum string

// Set of constants representing the allowable values for ListModuleStreamProfilesSortByEnum
const (
	ListModuleStreamProfilesSortByModulename ListModuleStreamProfilesSortByEnum = "moduleName"
)

var mappingListModuleStreamProfilesSortByEnum = map[string]ListModuleStreamProfilesSortByEnum{
	"moduleName": ListModuleStreamProfilesSortByModulename,
}

var mappingListModuleStreamProfilesSortByEnumLowerCase = map[string]ListModuleStreamProfilesSortByEnum{
	"modulename": ListModuleStreamProfilesSortByModulename,
}

// GetListModuleStreamProfilesSortByEnumValues Enumerates the set of values for ListModuleStreamProfilesSortByEnum
func GetListModuleStreamProfilesSortByEnumValues() []ListModuleStreamProfilesSortByEnum {
	values := make([]ListModuleStreamProfilesSortByEnum, 0)
	for _, v := range mappingListModuleStreamProfilesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListModuleStreamProfilesSortByEnumStringValues Enumerates the set of values in String for ListModuleStreamProfilesSortByEnum
func GetListModuleStreamProfilesSortByEnumStringValues() []string {
	return []string{
		"moduleName",
	}
}

// GetMappingListModuleStreamProfilesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModuleStreamProfilesSortByEnum(val string) (ListModuleStreamProfilesSortByEnum, bool) {
	enum, ok := mappingListModuleStreamProfilesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
