// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRunbookVersionsRequest wrapper for the ListRunbookVersions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListRunbookVersions.go.html to see an example of how to use ListRunbookVersionsRequest.
type ListRunbookVersionsRequest struct {

	// The ID of the compartment in which to list resources.
	// Empty only if the resource OCID query param is not specified.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState RunbookVersionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only schedule definitions whose associated runbookId matches the given runbookId.
	RunbookId *string `mandatory:"false" contributesTo:"query" name:"runbookId"`

	// A filter to return runbook versions whose identifier matches the given identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListRunbookVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListRunbookVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRunbookVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRunbookVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRunbookVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRunbookVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRunbookVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRunbookVersionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRunbookVersionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRunbookVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRunbookVersionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRunbookVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRunbookVersionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRunbookVersionsResponse wrapper for the ListRunbookVersions operation
type ListRunbookVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RunbookVersionCollection instances
	RunbookVersionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRunbookVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRunbookVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRunbookVersionsSortOrderEnum Enum with underlying type: string
type ListRunbookVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListRunbookVersionsSortOrderEnum
const (
	ListRunbookVersionsSortOrderAsc  ListRunbookVersionsSortOrderEnum = "ASC"
	ListRunbookVersionsSortOrderDesc ListRunbookVersionsSortOrderEnum = "DESC"
)

var mappingListRunbookVersionsSortOrderEnum = map[string]ListRunbookVersionsSortOrderEnum{
	"ASC":  ListRunbookVersionsSortOrderAsc,
	"DESC": ListRunbookVersionsSortOrderDesc,
}

var mappingListRunbookVersionsSortOrderEnumLowerCase = map[string]ListRunbookVersionsSortOrderEnum{
	"asc":  ListRunbookVersionsSortOrderAsc,
	"desc": ListRunbookVersionsSortOrderDesc,
}

// GetListRunbookVersionsSortOrderEnumValues Enumerates the set of values for ListRunbookVersionsSortOrderEnum
func GetListRunbookVersionsSortOrderEnumValues() []ListRunbookVersionsSortOrderEnum {
	values := make([]ListRunbookVersionsSortOrderEnum, 0)
	for _, v := range mappingListRunbookVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRunbookVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListRunbookVersionsSortOrderEnum
func GetListRunbookVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRunbookVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRunbookVersionsSortOrderEnum(val string) (ListRunbookVersionsSortOrderEnum, bool) {
	enum, ok := mappingListRunbookVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRunbookVersionsSortByEnum Enum with underlying type: string
type ListRunbookVersionsSortByEnum string

// Set of constants representing the allowable values for ListRunbookVersionsSortByEnum
const (
	ListRunbookVersionsSortByTimecreated ListRunbookVersionsSortByEnum = "timeCreated"
	ListRunbookVersionsSortByDisplayname ListRunbookVersionsSortByEnum = "displayName"
)

var mappingListRunbookVersionsSortByEnum = map[string]ListRunbookVersionsSortByEnum{
	"timeCreated": ListRunbookVersionsSortByTimecreated,
	"displayName": ListRunbookVersionsSortByDisplayname,
}

var mappingListRunbookVersionsSortByEnumLowerCase = map[string]ListRunbookVersionsSortByEnum{
	"timecreated": ListRunbookVersionsSortByTimecreated,
	"displayname": ListRunbookVersionsSortByDisplayname,
}

// GetListRunbookVersionsSortByEnumValues Enumerates the set of values for ListRunbookVersionsSortByEnum
func GetListRunbookVersionsSortByEnumValues() []ListRunbookVersionsSortByEnum {
	values := make([]ListRunbookVersionsSortByEnum, 0)
	for _, v := range mappingListRunbookVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRunbookVersionsSortByEnumStringValues Enumerates the set of values in String for ListRunbookVersionsSortByEnum
func GetListRunbookVersionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListRunbookVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRunbookVersionsSortByEnum(val string) (ListRunbookVersionsSortByEnum, bool) {
	enum, ok := mappingListRunbookVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
