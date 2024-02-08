// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPathAnalyzerTestsRequest wrapper for the ListPathAnalyzerTests operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vnmonitoring/ListPathAnalyzerTests.go.html to see an example of how to use ListPathAnalyzerTestsRequest.
type ListPathAnalyzerTestsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter that returns only resources whose `lifecycleState` matches the given `lifecycleState`.
	LifecycleState PathAnalyzerTestLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter that returns only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListPathAnalyzerTestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListPathAnalyzerTestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPathAnalyzerTestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPathAnalyzerTestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPathAnalyzerTestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPathAnalyzerTestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPathAnalyzerTestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPathAnalyzerTestLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPathAnalyzerTestLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPathAnalyzerTestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPathAnalyzerTestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPathAnalyzerTestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPathAnalyzerTestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPathAnalyzerTestsResponse wrapper for the ListPathAnalyzerTests operation
type ListPathAnalyzerTestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PathAnalyzerTestCollection instances
	PathAnalyzerTestCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPathAnalyzerTestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPathAnalyzerTestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPathAnalyzerTestsSortOrderEnum Enum with underlying type: string
type ListPathAnalyzerTestsSortOrderEnum string

// Set of constants representing the allowable values for ListPathAnalyzerTestsSortOrderEnum
const (
	ListPathAnalyzerTestsSortOrderAsc  ListPathAnalyzerTestsSortOrderEnum = "ASC"
	ListPathAnalyzerTestsSortOrderDesc ListPathAnalyzerTestsSortOrderEnum = "DESC"
)

var mappingListPathAnalyzerTestsSortOrderEnum = map[string]ListPathAnalyzerTestsSortOrderEnum{
	"ASC":  ListPathAnalyzerTestsSortOrderAsc,
	"DESC": ListPathAnalyzerTestsSortOrderDesc,
}

var mappingListPathAnalyzerTestsSortOrderEnumLowerCase = map[string]ListPathAnalyzerTestsSortOrderEnum{
	"asc":  ListPathAnalyzerTestsSortOrderAsc,
	"desc": ListPathAnalyzerTestsSortOrderDesc,
}

// GetListPathAnalyzerTestsSortOrderEnumValues Enumerates the set of values for ListPathAnalyzerTestsSortOrderEnum
func GetListPathAnalyzerTestsSortOrderEnumValues() []ListPathAnalyzerTestsSortOrderEnum {
	values := make([]ListPathAnalyzerTestsSortOrderEnum, 0)
	for _, v := range mappingListPathAnalyzerTestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPathAnalyzerTestsSortOrderEnumStringValues Enumerates the set of values in String for ListPathAnalyzerTestsSortOrderEnum
func GetListPathAnalyzerTestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPathAnalyzerTestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPathAnalyzerTestsSortOrderEnum(val string) (ListPathAnalyzerTestsSortOrderEnum, bool) {
	enum, ok := mappingListPathAnalyzerTestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPathAnalyzerTestsSortByEnum Enum with underlying type: string
type ListPathAnalyzerTestsSortByEnum string

// Set of constants representing the allowable values for ListPathAnalyzerTestsSortByEnum
const (
	ListPathAnalyzerTestsSortByTimecreated ListPathAnalyzerTestsSortByEnum = "TIMECREATED"
	ListPathAnalyzerTestsSortByDisplayname ListPathAnalyzerTestsSortByEnum = "DISPLAYNAME"
)

var mappingListPathAnalyzerTestsSortByEnum = map[string]ListPathAnalyzerTestsSortByEnum{
	"TIMECREATED": ListPathAnalyzerTestsSortByTimecreated,
	"DISPLAYNAME": ListPathAnalyzerTestsSortByDisplayname,
}

var mappingListPathAnalyzerTestsSortByEnumLowerCase = map[string]ListPathAnalyzerTestsSortByEnum{
	"timecreated": ListPathAnalyzerTestsSortByTimecreated,
	"displayname": ListPathAnalyzerTestsSortByDisplayname,
}

// GetListPathAnalyzerTestsSortByEnumValues Enumerates the set of values for ListPathAnalyzerTestsSortByEnum
func GetListPathAnalyzerTestsSortByEnumValues() []ListPathAnalyzerTestsSortByEnum {
	values := make([]ListPathAnalyzerTestsSortByEnum, 0)
	for _, v := range mappingListPathAnalyzerTestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPathAnalyzerTestsSortByEnumStringValues Enumerates the set of values in String for ListPathAnalyzerTestsSortByEnum
func GetListPathAnalyzerTestsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListPathAnalyzerTestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPathAnalyzerTestsSortByEnum(val string) (ListPathAnalyzerTestsSortByEnum, bool) {
	enum, ok := mappingListPathAnalyzerTestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
