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

// ListModuleStreamsRequest wrapper for the ListModuleStreams operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListModuleStreams.go.html to see an example of how to use ListModuleStreamsRequest.
type ListModuleStreamsRequest struct {

	// The software source OCID.
	SoftwareSourceId *string `mandatory:"true" contributesTo:"path" name:"softwareSourceId"`

	// The name of a module. This parameter is required if a
	// streamName is specified.
	ModuleName *string `mandatory:"false" contributesTo:"query" name:"moduleName"`

	// The name of the entity to be queried.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A boolean variable that is used to list only the latest versions of packages, module streams,
	// and stream profiles when set to true. All packages, module streams, and stream profiles are
	// returned when set to false.
	IsLatest *bool `mandatory:"false" contributesTo:"query" name:"isLatest"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListModuleStreamsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for moduleName is ascending.
	SortBy ListModuleStreamsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return resources that may partially match the module name given.
	ModuleNameContains *string `mandatory:"false" contributesTo:"query" name:"moduleNameContains"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModuleStreamsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModuleStreamsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModuleStreamsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModuleStreamsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListModuleStreamsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListModuleStreamsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListModuleStreamsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModuleStreamsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListModuleStreamsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListModuleStreamsResponse wrapper for the ListModuleStreams operation
type ListModuleStreamsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ModuleStreamCollection instances
	ModuleStreamCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListModuleStreamsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModuleStreamsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModuleStreamsSortOrderEnum Enum with underlying type: string
type ListModuleStreamsSortOrderEnum string

// Set of constants representing the allowable values for ListModuleStreamsSortOrderEnum
const (
	ListModuleStreamsSortOrderAsc  ListModuleStreamsSortOrderEnum = "ASC"
	ListModuleStreamsSortOrderDesc ListModuleStreamsSortOrderEnum = "DESC"
)

var mappingListModuleStreamsSortOrderEnum = map[string]ListModuleStreamsSortOrderEnum{
	"ASC":  ListModuleStreamsSortOrderAsc,
	"DESC": ListModuleStreamsSortOrderDesc,
}

var mappingListModuleStreamsSortOrderEnumLowerCase = map[string]ListModuleStreamsSortOrderEnum{
	"asc":  ListModuleStreamsSortOrderAsc,
	"desc": ListModuleStreamsSortOrderDesc,
}

// GetListModuleStreamsSortOrderEnumValues Enumerates the set of values for ListModuleStreamsSortOrderEnum
func GetListModuleStreamsSortOrderEnumValues() []ListModuleStreamsSortOrderEnum {
	values := make([]ListModuleStreamsSortOrderEnum, 0)
	for _, v := range mappingListModuleStreamsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListModuleStreamsSortOrderEnumStringValues Enumerates the set of values in String for ListModuleStreamsSortOrderEnum
func GetListModuleStreamsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListModuleStreamsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModuleStreamsSortOrderEnum(val string) (ListModuleStreamsSortOrderEnum, bool) {
	enum, ok := mappingListModuleStreamsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModuleStreamsSortByEnum Enum with underlying type: string
type ListModuleStreamsSortByEnum string

// Set of constants representing the allowable values for ListModuleStreamsSortByEnum
const (
	ListModuleStreamsSortByModulename ListModuleStreamsSortByEnum = "moduleName"
)

var mappingListModuleStreamsSortByEnum = map[string]ListModuleStreamsSortByEnum{
	"moduleName": ListModuleStreamsSortByModulename,
}

var mappingListModuleStreamsSortByEnumLowerCase = map[string]ListModuleStreamsSortByEnum{
	"modulename": ListModuleStreamsSortByModulename,
}

// GetListModuleStreamsSortByEnumValues Enumerates the set of values for ListModuleStreamsSortByEnum
func GetListModuleStreamsSortByEnumValues() []ListModuleStreamsSortByEnum {
	values := make([]ListModuleStreamsSortByEnum, 0)
	for _, v := range mappingListModuleStreamsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListModuleStreamsSortByEnumStringValues Enumerates the set of values in String for ListModuleStreamsSortByEnum
func GetListModuleStreamsSortByEnumStringValues() []string {
	return []string{
		"moduleName",
	}
}

// GetMappingListModuleStreamsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModuleStreamsSortByEnum(val string) (ListModuleStreamsSortByEnum, bool) {
	enum, ok := mappingListModuleStreamsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
