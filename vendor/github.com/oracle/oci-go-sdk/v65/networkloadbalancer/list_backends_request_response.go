// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBackendsRequest wrapper for the ListBackends operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkloadbalancer/ListBackends.go.html to see an example of how to use ListBackendsRequest.
type ListBackendsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
	NetworkLoadBalancerId *string `mandatory:"true" contributesTo:"path" name:"networkLoadBalancerId"`

	// The name of the backend set associated with the backend servers.
	// Example: `example_backend_set`
	BackendSetName *string `mandatory:"true" contributesTo:"path" name:"backendSetName"`

	// The unique Oracle-assigned identifier for the request. If you must contact Oracle about a
	// particular request, then provide the request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The system returns the requested resource, with a 200 status, only if the resource has no etag
	// matching the one specified. If the condition fails for the GET and HEAD methods, then the system returns the
	// HTTP status code `304 (Not Modified)`.
	// Example: `example-etag`
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// For list pagination. The maximum number of results per page or items to return, in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from which to start retrieving results.
	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' (ascending) or 'desc' (descending).
	SortOrder ListBackendsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. The default order for timeCreated is descending.
	// The default order for displayName is ascending. If no value is specified, then timeCreated is the default.
	SortBy ListBackendsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBackendsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBackendsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBackendsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBackendsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBackendsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListBackendsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBackendsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBackendsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBackendsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBackendsResponse wrapper for the ListBackends operation
type ListBackendsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BackendCollection instances
	BackendCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you must contact
	// Oracle about a particular request, then provide the request identifier.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBackendsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBackendsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBackendsSortOrderEnum Enum with underlying type: string
type ListBackendsSortOrderEnum string

// Set of constants representing the allowable values for ListBackendsSortOrderEnum
const (
	ListBackendsSortOrderAsc  ListBackendsSortOrderEnum = "ASC"
	ListBackendsSortOrderDesc ListBackendsSortOrderEnum = "DESC"
)

var mappingListBackendsSortOrderEnum = map[string]ListBackendsSortOrderEnum{
	"ASC":  ListBackendsSortOrderAsc,
	"DESC": ListBackendsSortOrderDesc,
}

var mappingListBackendsSortOrderEnumLowerCase = map[string]ListBackendsSortOrderEnum{
	"asc":  ListBackendsSortOrderAsc,
	"desc": ListBackendsSortOrderDesc,
}

// GetListBackendsSortOrderEnumValues Enumerates the set of values for ListBackendsSortOrderEnum
func GetListBackendsSortOrderEnumValues() []ListBackendsSortOrderEnum {
	values := make([]ListBackendsSortOrderEnum, 0)
	for _, v := range mappingListBackendsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBackendsSortOrderEnumStringValues Enumerates the set of values in String for ListBackendsSortOrderEnum
func GetListBackendsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBackendsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBackendsSortOrderEnum(val string) (ListBackendsSortOrderEnum, bool) {
	enum, ok := mappingListBackendsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBackendsSortByEnum Enum with underlying type: string
type ListBackendsSortByEnum string

// Set of constants representing the allowable values for ListBackendsSortByEnum
const (
	ListBackendsSortByTimecreated ListBackendsSortByEnum = "timeCreated"
	ListBackendsSortByDisplayname ListBackendsSortByEnum = "displayName"
)

var mappingListBackendsSortByEnum = map[string]ListBackendsSortByEnum{
	"timeCreated": ListBackendsSortByTimecreated,
	"displayName": ListBackendsSortByDisplayname,
}

var mappingListBackendsSortByEnumLowerCase = map[string]ListBackendsSortByEnum{
	"timecreated": ListBackendsSortByTimecreated,
	"displayname": ListBackendsSortByDisplayname,
}

// GetListBackendsSortByEnumValues Enumerates the set of values for ListBackendsSortByEnum
func GetListBackendsSortByEnumValues() []ListBackendsSortByEnum {
	values := make([]ListBackendsSortByEnum, 0)
	for _, v := range mappingListBackendsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBackendsSortByEnumStringValues Enumerates the set of values in String for ListBackendsSortByEnum
func GetListBackendsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListBackendsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBackendsSortByEnum(val string) (ListBackendsSortByEnum, bool) {
	enum, ok := mappingListBackendsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
