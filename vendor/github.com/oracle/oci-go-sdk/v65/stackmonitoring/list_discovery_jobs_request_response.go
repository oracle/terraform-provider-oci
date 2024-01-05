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

// ListDiscoveryJobsRequest wrapper for the ListDiscoveryJobs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListDiscoveryJobs.go.html to see an example of how to use ListDiscoveryJobsRequest.
type ListDiscoveryJobsRequest struct {

	// The ID of the compartment in which data is listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only discovery jobs that match the entire resource name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDiscoveryJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeUpdated is descending. Default order for resourceName is ascending.
	SortBy ListDiscoveryJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDiscoveryJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDiscoveryJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDiscoveryJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDiscoveryJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDiscoveryJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDiscoveryJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDiscoveryJobsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiscoveryJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDiscoveryJobsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDiscoveryJobsResponse wrapper for the ListDiscoveryJobs operation
type ListDiscoveryJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DiscoveryJobCollection instances
	DiscoveryJobCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDiscoveryJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDiscoveryJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDiscoveryJobsSortOrderEnum Enum with underlying type: string
type ListDiscoveryJobsSortOrderEnum string

// Set of constants representing the allowable values for ListDiscoveryJobsSortOrderEnum
const (
	ListDiscoveryJobsSortOrderAsc  ListDiscoveryJobsSortOrderEnum = "ASC"
	ListDiscoveryJobsSortOrderDesc ListDiscoveryJobsSortOrderEnum = "DESC"
)

var mappingListDiscoveryJobsSortOrderEnum = map[string]ListDiscoveryJobsSortOrderEnum{
	"ASC":  ListDiscoveryJobsSortOrderAsc,
	"DESC": ListDiscoveryJobsSortOrderDesc,
}

var mappingListDiscoveryJobsSortOrderEnumLowerCase = map[string]ListDiscoveryJobsSortOrderEnum{
	"asc":  ListDiscoveryJobsSortOrderAsc,
	"desc": ListDiscoveryJobsSortOrderDesc,
}

// GetListDiscoveryJobsSortOrderEnumValues Enumerates the set of values for ListDiscoveryJobsSortOrderEnum
func GetListDiscoveryJobsSortOrderEnumValues() []ListDiscoveryJobsSortOrderEnum {
	values := make([]ListDiscoveryJobsSortOrderEnum, 0)
	for _, v := range mappingListDiscoveryJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryJobsSortOrderEnumStringValues Enumerates the set of values in String for ListDiscoveryJobsSortOrderEnum
func GetListDiscoveryJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDiscoveryJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryJobsSortOrderEnum(val string) (ListDiscoveryJobsSortOrderEnum, bool) {
	enum, ok := mappingListDiscoveryJobsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDiscoveryJobsSortByEnum Enum with underlying type: string
type ListDiscoveryJobsSortByEnum string

// Set of constants representing the allowable values for ListDiscoveryJobsSortByEnum
const (
	ListDiscoveryJobsSortByTimeupdated  ListDiscoveryJobsSortByEnum = "timeUpdated"
	ListDiscoveryJobsSortByResourcename ListDiscoveryJobsSortByEnum = "resourceName"
)

var mappingListDiscoveryJobsSortByEnum = map[string]ListDiscoveryJobsSortByEnum{
	"timeUpdated":  ListDiscoveryJobsSortByTimeupdated,
	"resourceName": ListDiscoveryJobsSortByResourcename,
}

var mappingListDiscoveryJobsSortByEnumLowerCase = map[string]ListDiscoveryJobsSortByEnum{
	"timeupdated":  ListDiscoveryJobsSortByTimeupdated,
	"resourcename": ListDiscoveryJobsSortByResourcename,
}

// GetListDiscoveryJobsSortByEnumValues Enumerates the set of values for ListDiscoveryJobsSortByEnum
func GetListDiscoveryJobsSortByEnumValues() []ListDiscoveryJobsSortByEnum {
	values := make([]ListDiscoveryJobsSortByEnum, 0)
	for _, v := range mappingListDiscoveryJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryJobsSortByEnumStringValues Enumerates the set of values in String for ListDiscoveryJobsSortByEnum
func GetListDiscoveryJobsSortByEnumStringValues() []string {
	return []string{
		"timeUpdated",
		"resourceName",
	}
}

// GetMappingListDiscoveryJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryJobsSortByEnum(val string) (ListDiscoveryJobsSortByEnum, bool) {
	enum, ok := mappingListDiscoveryJobsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
