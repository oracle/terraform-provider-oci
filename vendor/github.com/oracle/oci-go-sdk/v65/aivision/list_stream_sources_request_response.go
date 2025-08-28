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

// ListStreamSourcesRequest wrapper for the ListStreamSources operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aivision/ListStreamSources.go.html to see an example of how to use ListStreamSourcesRequest.
type ListStreamSourcesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The filter to match projects with the given lifecycleState.
	LifecycleState StreamSourceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The filter to find the device with the given identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListStreamSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. The default order for timeCreated is descending. The default order for displayName is ascending.
	SortBy ListStreamSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStreamSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStreamSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListStreamSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStreamSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListStreamSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStreamSourceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetStreamSourceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStreamSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListStreamSourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStreamSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListStreamSourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListStreamSourcesResponse wrapper for the ListStreamSources operation
type ListStreamSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of StreamSourceCollection instances
	StreamSourceCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListStreamSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStreamSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStreamSourcesSortOrderEnum Enum with underlying type: string
type ListStreamSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListStreamSourcesSortOrderEnum
const (
	ListStreamSourcesSortOrderAsc  ListStreamSourcesSortOrderEnum = "ASC"
	ListStreamSourcesSortOrderDesc ListStreamSourcesSortOrderEnum = "DESC"
)

var mappingListStreamSourcesSortOrderEnum = map[string]ListStreamSourcesSortOrderEnum{
	"ASC":  ListStreamSourcesSortOrderAsc,
	"DESC": ListStreamSourcesSortOrderDesc,
}

var mappingListStreamSourcesSortOrderEnumLowerCase = map[string]ListStreamSourcesSortOrderEnum{
	"asc":  ListStreamSourcesSortOrderAsc,
	"desc": ListStreamSourcesSortOrderDesc,
}

// GetListStreamSourcesSortOrderEnumValues Enumerates the set of values for ListStreamSourcesSortOrderEnum
func GetListStreamSourcesSortOrderEnumValues() []ListStreamSourcesSortOrderEnum {
	values := make([]ListStreamSourcesSortOrderEnum, 0)
	for _, v := range mappingListStreamSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListStreamSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListStreamSourcesSortOrderEnum
func GetListStreamSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListStreamSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStreamSourcesSortOrderEnum(val string) (ListStreamSourcesSortOrderEnum, bool) {
	enum, ok := mappingListStreamSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStreamSourcesSortByEnum Enum with underlying type: string
type ListStreamSourcesSortByEnum string

// Set of constants representing the allowable values for ListStreamSourcesSortByEnum
const (
	ListStreamSourcesSortByTimecreated ListStreamSourcesSortByEnum = "timeCreated"
	ListStreamSourcesSortByDisplayname ListStreamSourcesSortByEnum = "displayName"
)

var mappingListStreamSourcesSortByEnum = map[string]ListStreamSourcesSortByEnum{
	"timeCreated": ListStreamSourcesSortByTimecreated,
	"displayName": ListStreamSourcesSortByDisplayname,
}

var mappingListStreamSourcesSortByEnumLowerCase = map[string]ListStreamSourcesSortByEnum{
	"timecreated": ListStreamSourcesSortByTimecreated,
	"displayname": ListStreamSourcesSortByDisplayname,
}

// GetListStreamSourcesSortByEnumValues Enumerates the set of values for ListStreamSourcesSortByEnum
func GetListStreamSourcesSortByEnumValues() []ListStreamSourcesSortByEnum {
	values := make([]ListStreamSourcesSortByEnum, 0)
	for _, v := range mappingListStreamSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListStreamSourcesSortByEnumStringValues Enumerates the set of values in String for ListStreamSourcesSortByEnum
func GetListStreamSourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListStreamSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStreamSourcesSortByEnum(val string) (ListStreamSourcesSortByEnum, bool) {
	enum, ok := mappingListStreamSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
