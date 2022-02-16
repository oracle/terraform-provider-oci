// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package limits

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListLimitDefinitionsRequest wrapper for the ListLimitDefinitions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limits/ListLimitDefinitions.go.html to see an example of how to use ListLimitDefinitionsRequest.
type ListLimitDefinitionsRequest struct {

	// The OCID of the parent compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The target service name.
	ServiceName *string `mandatory:"false" contributesTo:"query" name:"serviceName"`

	// Optional field, filter for a specific resource limit.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by.
	SortBy ListLimitDefinitionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'. By default, it is ascending.
	SortOrder ListLimitDefinitionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLimitDefinitionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLimitDefinitionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLimitDefinitionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLimitDefinitionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLimitDefinitionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLimitDefinitionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLimitDefinitionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLimitDefinitionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLimitDefinitionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLimitDefinitionsResponse wrapper for the ListLimitDefinitions operation
type ListLimitDefinitionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []LimitDefinitionSummary instances
	Items []LimitDefinitionSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLimitDefinitionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLimitDefinitionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLimitDefinitionsSortByEnum Enum with underlying type: string
type ListLimitDefinitionsSortByEnum string

// Set of constants representing the allowable values for ListLimitDefinitionsSortByEnum
const (
	ListLimitDefinitionsSortByName        ListLimitDefinitionsSortByEnum = "name"
	ListLimitDefinitionsSortByDescription ListLimitDefinitionsSortByEnum = "description"
)

var mappingListLimitDefinitionsSortByEnum = map[string]ListLimitDefinitionsSortByEnum{
	"name":        ListLimitDefinitionsSortByName,
	"description": ListLimitDefinitionsSortByDescription,
}

// GetListLimitDefinitionsSortByEnumValues Enumerates the set of values for ListLimitDefinitionsSortByEnum
func GetListLimitDefinitionsSortByEnumValues() []ListLimitDefinitionsSortByEnum {
	values := make([]ListLimitDefinitionsSortByEnum, 0)
	for _, v := range mappingListLimitDefinitionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLimitDefinitionsSortByEnumStringValues Enumerates the set of values in String for ListLimitDefinitionsSortByEnum
func GetListLimitDefinitionsSortByEnumStringValues() []string {
	return []string{
		"name",
		"description",
	}
}

// GetMappingListLimitDefinitionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLimitDefinitionsSortByEnum(val string) (ListLimitDefinitionsSortByEnum, bool) {
	mappingListLimitDefinitionsSortByEnumIgnoreCase := make(map[string]ListLimitDefinitionsSortByEnum)
	for k, v := range mappingListLimitDefinitionsSortByEnum {
		mappingListLimitDefinitionsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListLimitDefinitionsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListLimitDefinitionsSortOrderEnum Enum with underlying type: string
type ListLimitDefinitionsSortOrderEnum string

// Set of constants representing the allowable values for ListLimitDefinitionsSortOrderEnum
const (
	ListLimitDefinitionsSortOrderAsc  ListLimitDefinitionsSortOrderEnum = "ASC"
	ListLimitDefinitionsSortOrderDesc ListLimitDefinitionsSortOrderEnum = "DESC"
)

var mappingListLimitDefinitionsSortOrderEnum = map[string]ListLimitDefinitionsSortOrderEnum{
	"ASC":  ListLimitDefinitionsSortOrderAsc,
	"DESC": ListLimitDefinitionsSortOrderDesc,
}

// GetListLimitDefinitionsSortOrderEnumValues Enumerates the set of values for ListLimitDefinitionsSortOrderEnum
func GetListLimitDefinitionsSortOrderEnumValues() []ListLimitDefinitionsSortOrderEnum {
	values := make([]ListLimitDefinitionsSortOrderEnum, 0)
	for _, v := range mappingListLimitDefinitionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLimitDefinitionsSortOrderEnumStringValues Enumerates the set of values in String for ListLimitDefinitionsSortOrderEnum
func GetListLimitDefinitionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLimitDefinitionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLimitDefinitionsSortOrderEnum(val string) (ListLimitDefinitionsSortOrderEnum, bool) {
	mappingListLimitDefinitionsSortOrderEnumIgnoreCase := make(map[string]ListLimitDefinitionsSortOrderEnum)
	for k, v := range mappingListLimitDefinitionsSortOrderEnum {
		mappingListLimitDefinitionsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListLimitDefinitionsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
