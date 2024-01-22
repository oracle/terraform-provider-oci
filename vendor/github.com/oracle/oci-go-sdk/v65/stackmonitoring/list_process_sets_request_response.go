// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListProcessSetsRequest wrapper for the ListProcessSets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListProcessSets.go.html to see an example of how to use ListProcessSetsRequest.
type ListProcessSetsRequest struct {

	// The ID of the compartment in which data is listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListProcessSetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. Only one sort order may be provided. Default order for timeUpdated is descending. Default order for name is ascending.
	SortBy ListProcessSetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProcessSetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProcessSetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProcessSetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProcessSetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProcessSetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProcessSetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProcessSetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProcessSetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProcessSetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProcessSetsResponse wrapper for the ListProcessSets operation
type ListProcessSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProcessSetCollection instances
	ProcessSetCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListProcessSetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProcessSetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProcessSetsSortOrderEnum Enum with underlying type: string
type ListProcessSetsSortOrderEnum string

// Set of constants representing the allowable values for ListProcessSetsSortOrderEnum
const (
	ListProcessSetsSortOrderAsc  ListProcessSetsSortOrderEnum = "ASC"
	ListProcessSetsSortOrderDesc ListProcessSetsSortOrderEnum = "DESC"
)

var mappingListProcessSetsSortOrderEnum = map[string]ListProcessSetsSortOrderEnum{
	"ASC":  ListProcessSetsSortOrderAsc,
	"DESC": ListProcessSetsSortOrderDesc,
}

var mappingListProcessSetsSortOrderEnumLowerCase = map[string]ListProcessSetsSortOrderEnum{
	"asc":  ListProcessSetsSortOrderAsc,
	"desc": ListProcessSetsSortOrderDesc,
}

// GetListProcessSetsSortOrderEnumValues Enumerates the set of values for ListProcessSetsSortOrderEnum
func GetListProcessSetsSortOrderEnumValues() []ListProcessSetsSortOrderEnum {
	values := make([]ListProcessSetsSortOrderEnum, 0)
	for _, v := range mappingListProcessSetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProcessSetsSortOrderEnumStringValues Enumerates the set of values in String for ListProcessSetsSortOrderEnum
func GetListProcessSetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProcessSetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProcessSetsSortOrderEnum(val string) (ListProcessSetsSortOrderEnum, bool) {
	enum, ok := mappingListProcessSetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProcessSetsSortByEnum Enum with underlying type: string
type ListProcessSetsSortByEnum string

// Set of constants representing the allowable values for ListProcessSetsSortByEnum
const (
	ListProcessSetsSortByTimeupdated ListProcessSetsSortByEnum = "timeUpdated"
	ListProcessSetsSortByName        ListProcessSetsSortByEnum = "name"
)

var mappingListProcessSetsSortByEnum = map[string]ListProcessSetsSortByEnum{
	"timeUpdated": ListProcessSetsSortByTimeupdated,
	"name":        ListProcessSetsSortByName,
}

var mappingListProcessSetsSortByEnumLowerCase = map[string]ListProcessSetsSortByEnum{
	"timeupdated": ListProcessSetsSortByTimeupdated,
	"name":        ListProcessSetsSortByName,
}

// GetListProcessSetsSortByEnumValues Enumerates the set of values for ListProcessSetsSortByEnum
func GetListProcessSetsSortByEnumValues() []ListProcessSetsSortByEnum {
	values := make([]ListProcessSetsSortByEnum, 0)
	for _, v := range mappingListProcessSetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProcessSetsSortByEnumStringValues Enumerates the set of values in String for ListProcessSetsSortByEnum
func GetListProcessSetsSortByEnumStringValues() []string {
	return []string{
		"timeUpdated",
		"name",
	}
}

// GetMappingListProcessSetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProcessSetsSortByEnum(val string) (ListProcessSetsSortByEnum, bool) {
	enum, ok := mappingListProcessSetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
