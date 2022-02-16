// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListBackendSetsRequest wrapper for the ListBackendSets operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkloadbalancer/ListBackendSets.go.html to see an example of how to use ListBackendSetsRequest.
type ListBackendSetsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
	NetworkLoadBalancerId *string `mandatory:"true" contributesTo:"path" name:"networkLoadBalancerId"`

	// The unique Oracle-assigned identifier for the request. If you must contact Oracle about a
	// particular request, then provide the request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The system returns the requested resource, with a 200 status, only if the resource has no etag
	// matching the one specified. If the condition fails for the GET and HEAD methods, then the system returns the
	// HTTP status code `304 (Not Modified)`.
	// Example: `example-etag`
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// For list pagination. The maximum number of results per page or items to return, in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from which to start retrieving results.
	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' (ascending) or 'desc' (descending).
	SortOrder ListBackendSetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. The default order for timeCreated is descending.
	// The default order for displayName is ascending. If no value is specified, then timeCreated is the default.
	SortBy ListBackendSetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBackendSetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBackendSetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBackendSetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBackendSetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBackendSetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListBackendSetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBackendSetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBackendSetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBackendSetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBackendSetsResponse wrapper for the ListBackendSets operation
type ListBackendSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BackendSetCollection instances
	BackendSetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you must contact
	// Oracle about a particular request, then provide the request identifier.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBackendSetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBackendSetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBackendSetsSortOrderEnum Enum with underlying type: string
type ListBackendSetsSortOrderEnum string

// Set of constants representing the allowable values for ListBackendSetsSortOrderEnum
const (
	ListBackendSetsSortOrderAsc  ListBackendSetsSortOrderEnum = "ASC"
	ListBackendSetsSortOrderDesc ListBackendSetsSortOrderEnum = "DESC"
)

var mappingListBackendSetsSortOrderEnum = map[string]ListBackendSetsSortOrderEnum{
	"ASC":  ListBackendSetsSortOrderAsc,
	"DESC": ListBackendSetsSortOrderDesc,
}

// GetListBackendSetsSortOrderEnumValues Enumerates the set of values for ListBackendSetsSortOrderEnum
func GetListBackendSetsSortOrderEnumValues() []ListBackendSetsSortOrderEnum {
	values := make([]ListBackendSetsSortOrderEnum, 0)
	for _, v := range mappingListBackendSetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBackendSetsSortOrderEnumStringValues Enumerates the set of values in String for ListBackendSetsSortOrderEnum
func GetListBackendSetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBackendSetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBackendSetsSortOrderEnum(val string) (ListBackendSetsSortOrderEnum, bool) {
	mappingListBackendSetsSortOrderEnumIgnoreCase := make(map[string]ListBackendSetsSortOrderEnum)
	for k, v := range mappingListBackendSetsSortOrderEnum {
		mappingListBackendSetsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListBackendSetsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListBackendSetsSortByEnum Enum with underlying type: string
type ListBackendSetsSortByEnum string

// Set of constants representing the allowable values for ListBackendSetsSortByEnum
const (
	ListBackendSetsSortByTimecreated ListBackendSetsSortByEnum = "timeCreated"
	ListBackendSetsSortByDisplayname ListBackendSetsSortByEnum = "displayName"
)

var mappingListBackendSetsSortByEnum = map[string]ListBackendSetsSortByEnum{
	"timeCreated": ListBackendSetsSortByTimecreated,
	"displayName": ListBackendSetsSortByDisplayname,
}

// GetListBackendSetsSortByEnumValues Enumerates the set of values for ListBackendSetsSortByEnum
func GetListBackendSetsSortByEnumValues() []ListBackendSetsSortByEnum {
	values := make([]ListBackendSetsSortByEnum, 0)
	for _, v := range mappingListBackendSetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBackendSetsSortByEnumStringValues Enumerates the set of values in String for ListBackendSetsSortByEnum
func GetListBackendSetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListBackendSetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBackendSetsSortByEnum(val string) (ListBackendSetsSortByEnum, bool) {
	mappingListBackendSetsSortByEnumIgnoreCase := make(map[string]ListBackendSetsSortByEnum)
	for k, v := range mappingListBackendSetsSortByEnum {
		mappingListBackendSetsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListBackendSetsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
