// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFleetAutoExecutionsRequest wrapper for the ListFleetAutoExecutions operation
type ListFleetAutoExecutionsRequest struct {

	// Unique Fleet identifier.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Returns results where the execution time is greater than or equal to the specified value
	TimeExecutionGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeExecutionGreaterThanOrEqualTo"`

	// Returns results where the execution time is less than the specified value
	TimeExecutionLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeExecutionLessThan"`

	// The field to sort by. Only one sort order may be provided.  Default order for timeCreated and timeScheduled is descending.
	SortBy ListFleetAutoExecutionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFleetAutoExecutionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFleetAutoExecutionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFleetAutoExecutionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFleetAutoExecutionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFleetAutoExecutionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFleetAutoExecutionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFleetAutoExecutionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFleetAutoExecutionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFleetAutoExecutionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFleetAutoExecutionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFleetAutoExecutionsResponse wrapper for the ListFleetAutoExecutions operation
type ListFleetAutoExecutionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AutoExecutionCollection instances
	AutoExecutionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFleetAutoExecutionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFleetAutoExecutionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFleetAutoExecutionsSortByEnum Enum with underlying type: string
type ListFleetAutoExecutionsSortByEnum string

// Set of constants representing the allowable values for ListFleetAutoExecutionsSortByEnum
const (
	ListFleetAutoExecutionsSortByTimecreated   ListFleetAutoExecutionsSortByEnum = "timeCreated"
	ListFleetAutoExecutionsSortByTimescheduled ListFleetAutoExecutionsSortByEnum = "timeScheduled"
)

var mappingListFleetAutoExecutionsSortByEnum = map[string]ListFleetAutoExecutionsSortByEnum{
	"timeCreated":   ListFleetAutoExecutionsSortByTimecreated,
	"timeScheduled": ListFleetAutoExecutionsSortByTimescheduled,
}

var mappingListFleetAutoExecutionsSortByEnumLowerCase = map[string]ListFleetAutoExecutionsSortByEnum{
	"timecreated":   ListFleetAutoExecutionsSortByTimecreated,
	"timescheduled": ListFleetAutoExecutionsSortByTimescheduled,
}

// GetListFleetAutoExecutionsSortByEnumValues Enumerates the set of values for ListFleetAutoExecutionsSortByEnum
func GetListFleetAutoExecutionsSortByEnumValues() []ListFleetAutoExecutionsSortByEnum {
	values := make([]ListFleetAutoExecutionsSortByEnum, 0)
	for _, v := range mappingListFleetAutoExecutionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetAutoExecutionsSortByEnumStringValues Enumerates the set of values in String for ListFleetAutoExecutionsSortByEnum
func GetListFleetAutoExecutionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeScheduled",
	}
}

// GetMappingListFleetAutoExecutionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetAutoExecutionsSortByEnum(val string) (ListFleetAutoExecutionsSortByEnum, bool) {
	enum, ok := mappingListFleetAutoExecutionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFleetAutoExecutionsSortOrderEnum Enum with underlying type: string
type ListFleetAutoExecutionsSortOrderEnum string

// Set of constants representing the allowable values for ListFleetAutoExecutionsSortOrderEnum
const (
	ListFleetAutoExecutionsSortOrderAsc  ListFleetAutoExecutionsSortOrderEnum = "ASC"
	ListFleetAutoExecutionsSortOrderDesc ListFleetAutoExecutionsSortOrderEnum = "DESC"
)

var mappingListFleetAutoExecutionsSortOrderEnum = map[string]ListFleetAutoExecutionsSortOrderEnum{
	"ASC":  ListFleetAutoExecutionsSortOrderAsc,
	"DESC": ListFleetAutoExecutionsSortOrderDesc,
}

var mappingListFleetAutoExecutionsSortOrderEnumLowerCase = map[string]ListFleetAutoExecutionsSortOrderEnum{
	"asc":  ListFleetAutoExecutionsSortOrderAsc,
	"desc": ListFleetAutoExecutionsSortOrderDesc,
}

// GetListFleetAutoExecutionsSortOrderEnumValues Enumerates the set of values for ListFleetAutoExecutionsSortOrderEnum
func GetListFleetAutoExecutionsSortOrderEnumValues() []ListFleetAutoExecutionsSortOrderEnum {
	values := make([]ListFleetAutoExecutionsSortOrderEnum, 0)
	for _, v := range mappingListFleetAutoExecutionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetAutoExecutionsSortOrderEnumStringValues Enumerates the set of values in String for ListFleetAutoExecutionsSortOrderEnum
func GetListFleetAutoExecutionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFleetAutoExecutionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetAutoExecutionsSortOrderEnum(val string) (ListFleetAutoExecutionsSortOrderEnum, bool) {
	enum, ok := mappingListFleetAutoExecutionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
