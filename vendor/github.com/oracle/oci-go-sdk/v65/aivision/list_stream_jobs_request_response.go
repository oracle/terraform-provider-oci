// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package aivision

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListStreamJobsRequest wrapper for the ListStreamJobs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aivision/ListStreamJobs.go.html to see an example of how to use ListStreamJobsRequest.
type ListStreamJobsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The filter to match projects with the given lifecycleState.
	LifecycleState StreamJobLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The filter to find the streamjob with the given identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListStreamJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. The default order for timeCreated is descending. The default order for displayName is ascending.
	SortBy ListStreamJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStreamJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStreamJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListStreamJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStreamJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListStreamJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStreamJobLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetStreamJobLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStreamJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListStreamJobsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStreamJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListStreamJobsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListStreamJobsResponse wrapper for the ListStreamJobs operation
type ListStreamJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of StreamJobCollection instances
	StreamJobCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListStreamJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStreamJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStreamJobsSortOrderEnum Enum with underlying type: string
type ListStreamJobsSortOrderEnum string

// Set of constants representing the allowable values for ListStreamJobsSortOrderEnum
const (
	ListStreamJobsSortOrderAsc  ListStreamJobsSortOrderEnum = "ASC"
	ListStreamJobsSortOrderDesc ListStreamJobsSortOrderEnum = "DESC"
)

var mappingListStreamJobsSortOrderEnum = map[string]ListStreamJobsSortOrderEnum{
	"ASC":  ListStreamJobsSortOrderAsc,
	"DESC": ListStreamJobsSortOrderDesc,
}

var mappingListStreamJobsSortOrderEnumLowerCase = map[string]ListStreamJobsSortOrderEnum{
	"asc":  ListStreamJobsSortOrderAsc,
	"desc": ListStreamJobsSortOrderDesc,
}

// GetListStreamJobsSortOrderEnumValues Enumerates the set of values for ListStreamJobsSortOrderEnum
func GetListStreamJobsSortOrderEnumValues() []ListStreamJobsSortOrderEnum {
	values := make([]ListStreamJobsSortOrderEnum, 0)
	for _, v := range mappingListStreamJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListStreamJobsSortOrderEnumStringValues Enumerates the set of values in String for ListStreamJobsSortOrderEnum
func GetListStreamJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListStreamJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStreamJobsSortOrderEnum(val string) (ListStreamJobsSortOrderEnum, bool) {
	enum, ok := mappingListStreamJobsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStreamJobsSortByEnum Enum with underlying type: string
type ListStreamJobsSortByEnum string

// Set of constants representing the allowable values for ListStreamJobsSortByEnum
const (
	ListStreamJobsSortByTimecreated ListStreamJobsSortByEnum = "timeCreated"
	ListStreamJobsSortByDisplayname ListStreamJobsSortByEnum = "displayName"
)

var mappingListStreamJobsSortByEnum = map[string]ListStreamJobsSortByEnum{
	"timeCreated": ListStreamJobsSortByTimecreated,
	"displayName": ListStreamJobsSortByDisplayname,
}

var mappingListStreamJobsSortByEnumLowerCase = map[string]ListStreamJobsSortByEnum{
	"timecreated": ListStreamJobsSortByTimecreated,
	"displayname": ListStreamJobsSortByDisplayname,
}

// GetListStreamJobsSortByEnumValues Enumerates the set of values for ListStreamJobsSortByEnum
func GetListStreamJobsSortByEnumValues() []ListStreamJobsSortByEnum {
	values := make([]ListStreamJobsSortByEnum, 0)
	for _, v := range mappingListStreamJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListStreamJobsSortByEnumStringValues Enumerates the set of values in String for ListStreamJobsSortByEnum
func GetListStreamJobsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListStreamJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStreamJobsSortByEnum(val string) (ListStreamJobsSortByEnum, bool) {
	enum, ok := mappingListStreamJobsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
