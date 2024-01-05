// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAvailabilityHistoriesRequest wrapper for the ListAvailabilityHistories operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementagent/ListAvailabilityHistories.go.html to see an example of how to use ListAvailabilityHistoriesRequest.
type ListAvailabilityHistoriesRequest struct {

	// Unique Management Agent identifier
	ManagementAgentId *string `mandatory:"true" contributesTo:"path" name:"managementAgentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Filter to limit the availability history results to that of time after the input time including the boundary record.
	// Defaulted to current date minus one year.
	// The date and time to be given as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 5.6.
	TimeAvailabilityStatusEndedGreaterThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeAvailabilityStatusEndedGreaterThan"`

	// Filter to limit the availability history results to that of time before the input time including the boundary record
	// Defaulted to current date.
	// The date and time to be given as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 5.6.
	TimeAvailabilityStatusStartedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeAvailabilityStatusStartedLessThan"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAvailabilityHistoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Default order for timeAvailabilityStatusStarted is descending.
	SortBy ListAvailabilityHistoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailabilityHistoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailabilityHistoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAvailabilityHistoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailabilityHistoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAvailabilityHistoriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAvailabilityHistoriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAvailabilityHistoriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAvailabilityHistoriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAvailabilityHistoriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAvailabilityHistoriesResponse wrapper for the ListAvailabilityHistories operation
type ListAvailabilityHistoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AvailabilityHistorySummary instances
	Items []AvailabilityHistorySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAvailabilityHistoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailabilityHistoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailabilityHistoriesSortOrderEnum Enum with underlying type: string
type ListAvailabilityHistoriesSortOrderEnum string

// Set of constants representing the allowable values for ListAvailabilityHistoriesSortOrderEnum
const (
	ListAvailabilityHistoriesSortOrderAsc  ListAvailabilityHistoriesSortOrderEnum = "ASC"
	ListAvailabilityHistoriesSortOrderDesc ListAvailabilityHistoriesSortOrderEnum = "DESC"
)

var mappingListAvailabilityHistoriesSortOrderEnum = map[string]ListAvailabilityHistoriesSortOrderEnum{
	"ASC":  ListAvailabilityHistoriesSortOrderAsc,
	"DESC": ListAvailabilityHistoriesSortOrderDesc,
}

var mappingListAvailabilityHistoriesSortOrderEnumLowerCase = map[string]ListAvailabilityHistoriesSortOrderEnum{
	"asc":  ListAvailabilityHistoriesSortOrderAsc,
	"desc": ListAvailabilityHistoriesSortOrderDesc,
}

// GetListAvailabilityHistoriesSortOrderEnumValues Enumerates the set of values for ListAvailabilityHistoriesSortOrderEnum
func GetListAvailabilityHistoriesSortOrderEnumValues() []ListAvailabilityHistoriesSortOrderEnum {
	values := make([]ListAvailabilityHistoriesSortOrderEnum, 0)
	for _, v := range mappingListAvailabilityHistoriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailabilityHistoriesSortOrderEnumStringValues Enumerates the set of values in String for ListAvailabilityHistoriesSortOrderEnum
func GetListAvailabilityHistoriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAvailabilityHistoriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailabilityHistoriesSortOrderEnum(val string) (ListAvailabilityHistoriesSortOrderEnum, bool) {
	enum, ok := mappingListAvailabilityHistoriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAvailabilityHistoriesSortByEnum Enum with underlying type: string
type ListAvailabilityHistoriesSortByEnum string

// Set of constants representing the allowable values for ListAvailabilityHistoriesSortByEnum
const (
	ListAvailabilityHistoriesSortByTimeavailabilitystatusstarted ListAvailabilityHistoriesSortByEnum = "timeAvailabilityStatusStarted"
)

var mappingListAvailabilityHistoriesSortByEnum = map[string]ListAvailabilityHistoriesSortByEnum{
	"timeAvailabilityStatusStarted": ListAvailabilityHistoriesSortByTimeavailabilitystatusstarted,
}

var mappingListAvailabilityHistoriesSortByEnumLowerCase = map[string]ListAvailabilityHistoriesSortByEnum{
	"timeavailabilitystatusstarted": ListAvailabilityHistoriesSortByTimeavailabilitystatusstarted,
}

// GetListAvailabilityHistoriesSortByEnumValues Enumerates the set of values for ListAvailabilityHistoriesSortByEnum
func GetListAvailabilityHistoriesSortByEnumValues() []ListAvailabilityHistoriesSortByEnum {
	values := make([]ListAvailabilityHistoriesSortByEnum, 0)
	for _, v := range mappingListAvailabilityHistoriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailabilityHistoriesSortByEnumStringValues Enumerates the set of values in String for ListAvailabilityHistoriesSortByEnum
func GetListAvailabilityHistoriesSortByEnumStringValues() []string {
	return []string{
		"timeAvailabilityStatusStarted",
	}
}

// GetMappingListAvailabilityHistoriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailabilityHistoriesSortByEnum(val string) (ListAvailabilityHistoriesSortByEnum, bool) {
	enum, ok := mappingListAvailabilityHistoriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
