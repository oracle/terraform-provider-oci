// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListManagedInstanceErrataRequest wrapper for the ListManagedInstanceErrata operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListManagedInstanceErrata.go.html to see an example of how to use ListManagedInstanceErrataRequest.
type ListManagedInstanceErrataRequest struct {

	// OCID for the managed instance
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagedInstanceErrataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListManagedInstanceErrataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceErrataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceErrataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceErrataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceErrataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceErrataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceErrataSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceErrataSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceErrataSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceErrataSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceErrataResponse wrapper for the ListManagedInstanceErrata operation
type ListManagedInstanceErrataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ErratumSummary instances
	Items []ErratumSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceErrataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceErrataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceErrataSortOrderEnum Enum with underlying type: string
type ListManagedInstanceErrataSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceErrataSortOrderEnum
const (
	ListManagedInstanceErrataSortOrderAsc  ListManagedInstanceErrataSortOrderEnum = "ASC"
	ListManagedInstanceErrataSortOrderDesc ListManagedInstanceErrataSortOrderEnum = "DESC"
)

var mappingListManagedInstanceErrataSortOrderEnum = map[string]ListManagedInstanceErrataSortOrderEnum{
	"ASC":  ListManagedInstanceErrataSortOrderAsc,
	"DESC": ListManagedInstanceErrataSortOrderDesc,
}

// GetListManagedInstanceErrataSortOrderEnumValues Enumerates the set of values for ListManagedInstanceErrataSortOrderEnum
func GetListManagedInstanceErrataSortOrderEnumValues() []ListManagedInstanceErrataSortOrderEnum {
	values := make([]ListManagedInstanceErrataSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceErrataSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceErrataSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceErrataSortOrderEnum
func GetListManagedInstanceErrataSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceErrataSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceErrataSortOrderEnum(val string) (ListManagedInstanceErrataSortOrderEnum, bool) {
	mappingListManagedInstanceErrataSortOrderEnumIgnoreCase := make(map[string]ListManagedInstanceErrataSortOrderEnum)
	for k, v := range mappingListManagedInstanceErrataSortOrderEnum {
		mappingListManagedInstanceErrataSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListManagedInstanceErrataSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceErrataSortByEnum Enum with underlying type: string
type ListManagedInstanceErrataSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceErrataSortByEnum
const (
	ListManagedInstanceErrataSortByTimecreated ListManagedInstanceErrataSortByEnum = "TIMECREATED"
	ListManagedInstanceErrataSortByDisplayname ListManagedInstanceErrataSortByEnum = "DISPLAYNAME"
)

var mappingListManagedInstanceErrataSortByEnum = map[string]ListManagedInstanceErrataSortByEnum{
	"TIMECREATED": ListManagedInstanceErrataSortByTimecreated,
	"DISPLAYNAME": ListManagedInstanceErrataSortByDisplayname,
}

// GetListManagedInstanceErrataSortByEnumValues Enumerates the set of values for ListManagedInstanceErrataSortByEnum
func GetListManagedInstanceErrataSortByEnumValues() []ListManagedInstanceErrataSortByEnum {
	values := make([]ListManagedInstanceErrataSortByEnum, 0)
	for _, v := range mappingListManagedInstanceErrataSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceErrataSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceErrataSortByEnum
func GetListManagedInstanceErrataSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListManagedInstanceErrataSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceErrataSortByEnum(val string) (ListManagedInstanceErrataSortByEnum, bool) {
	mappingListManagedInstanceErrataSortByEnumIgnoreCase := make(map[string]ListManagedInstanceErrataSortByEnum)
	for k, v := range mappingListManagedInstanceErrataSortByEnum {
		mappingListManagedInstanceErrataSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListManagedInstanceErrataSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
