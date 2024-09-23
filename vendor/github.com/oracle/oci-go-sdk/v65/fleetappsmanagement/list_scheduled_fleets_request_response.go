// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListScheduledFleetsRequest wrapper for the ListScheduledFleets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListScheduledFleets.go.html to see an example of how to use ListScheduledFleetsRequest.
type ListScheduledFleetsRequest struct {

	// unique SchedulerDefinition identifier
	SchedulerDefinitionId *string `mandatory:"true" contributesTo:"path" name:"schedulerDefinitionId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListScheduledFleetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for displayName is ascending.
	SortBy ListScheduledFleetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListScheduledFleetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListScheduledFleetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListScheduledFleetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListScheduledFleetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListScheduledFleetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListScheduledFleetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListScheduledFleetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListScheduledFleetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListScheduledFleetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListScheduledFleetsResponse wrapper for the ListScheduledFleets operation
type ListScheduledFleetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ScheduledFleetCollection instances
	ScheduledFleetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListScheduledFleetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListScheduledFleetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListScheduledFleetsSortOrderEnum Enum with underlying type: string
type ListScheduledFleetsSortOrderEnum string

// Set of constants representing the allowable values for ListScheduledFleetsSortOrderEnum
const (
	ListScheduledFleetsSortOrderAsc  ListScheduledFleetsSortOrderEnum = "ASC"
	ListScheduledFleetsSortOrderDesc ListScheduledFleetsSortOrderEnum = "DESC"
)

var mappingListScheduledFleetsSortOrderEnum = map[string]ListScheduledFleetsSortOrderEnum{
	"ASC":  ListScheduledFleetsSortOrderAsc,
	"DESC": ListScheduledFleetsSortOrderDesc,
}

var mappingListScheduledFleetsSortOrderEnumLowerCase = map[string]ListScheduledFleetsSortOrderEnum{
	"asc":  ListScheduledFleetsSortOrderAsc,
	"desc": ListScheduledFleetsSortOrderDesc,
}

// GetListScheduledFleetsSortOrderEnumValues Enumerates the set of values for ListScheduledFleetsSortOrderEnum
func GetListScheduledFleetsSortOrderEnumValues() []ListScheduledFleetsSortOrderEnum {
	values := make([]ListScheduledFleetsSortOrderEnum, 0)
	for _, v := range mappingListScheduledFleetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledFleetsSortOrderEnumStringValues Enumerates the set of values in String for ListScheduledFleetsSortOrderEnum
func GetListScheduledFleetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListScheduledFleetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledFleetsSortOrderEnum(val string) (ListScheduledFleetsSortOrderEnum, bool) {
	enum, ok := mappingListScheduledFleetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListScheduledFleetsSortByEnum Enum with underlying type: string
type ListScheduledFleetsSortByEnum string

// Set of constants representing the allowable values for ListScheduledFleetsSortByEnum
const (
	ListScheduledFleetsSortByDisplayname ListScheduledFleetsSortByEnum = "displayName"
)

var mappingListScheduledFleetsSortByEnum = map[string]ListScheduledFleetsSortByEnum{
	"displayName": ListScheduledFleetsSortByDisplayname,
}

var mappingListScheduledFleetsSortByEnumLowerCase = map[string]ListScheduledFleetsSortByEnum{
	"displayname": ListScheduledFleetsSortByDisplayname,
}

// GetListScheduledFleetsSortByEnumValues Enumerates the set of values for ListScheduledFleetsSortByEnum
func GetListScheduledFleetsSortByEnumValues() []ListScheduledFleetsSortByEnum {
	values := make([]ListScheduledFleetsSortByEnum, 0)
	for _, v := range mappingListScheduledFleetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledFleetsSortByEnumStringValues Enumerates the set of values in String for ListScheduledFleetsSortByEnum
func GetListScheduledFleetsSortByEnumStringValues() []string {
	return []string{
		"displayName",
	}
}

// GetMappingListScheduledFleetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledFleetsSortByEnum(val string) (ListScheduledFleetsSortByEnum, bool) {
	enum, ok := mappingListScheduledFleetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
