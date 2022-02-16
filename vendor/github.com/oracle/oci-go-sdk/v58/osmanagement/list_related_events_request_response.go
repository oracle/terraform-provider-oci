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

// ListRelatedEventsRequest wrapper for the ListRelatedEvents operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListRelatedEvents.go.html to see an example of how to use ListRelatedEventsRequest.
type ListRelatedEventsRequest struct {

	// Event fingerprint identifier
	EventFingerprint *string `mandatory:"true" contributesTo:"query" name:"eventFingerprint"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListRelatedEventsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for id is descending.
	SortBy ListRelatedEventsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// filter event occurrence. Selecting only those last occurred before given date in ISO 8601 format
	// Example: 2017-07-14T02:40:00.000Z
	LatestTimestampLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"latestTimestampLessThan"`

	// filter event occurrence. Selecting only those last occurred on or after given date in ISO 8601 format
	// Example: 2017-07-14T02:40:00.000Z
	LatestTimestampGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"latestTimestampGreaterThanOrEqualTo"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRelatedEventsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRelatedEventsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRelatedEventsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRelatedEventsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRelatedEventsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRelatedEventsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRelatedEventsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRelatedEventsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRelatedEventsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRelatedEventsResponse wrapper for the ListRelatedEvents operation
type ListRelatedEventsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RelatedEventCollection instances
	RelatedEventCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRelatedEventsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRelatedEventsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRelatedEventsSortOrderEnum Enum with underlying type: string
type ListRelatedEventsSortOrderEnum string

// Set of constants representing the allowable values for ListRelatedEventsSortOrderEnum
const (
	ListRelatedEventsSortOrderAsc  ListRelatedEventsSortOrderEnum = "ASC"
	ListRelatedEventsSortOrderDesc ListRelatedEventsSortOrderEnum = "DESC"
)

var mappingListRelatedEventsSortOrderEnum = map[string]ListRelatedEventsSortOrderEnum{
	"ASC":  ListRelatedEventsSortOrderAsc,
	"DESC": ListRelatedEventsSortOrderDesc,
}

// GetListRelatedEventsSortOrderEnumValues Enumerates the set of values for ListRelatedEventsSortOrderEnum
func GetListRelatedEventsSortOrderEnumValues() []ListRelatedEventsSortOrderEnum {
	values := make([]ListRelatedEventsSortOrderEnum, 0)
	for _, v := range mappingListRelatedEventsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRelatedEventsSortOrderEnumStringValues Enumerates the set of values in String for ListRelatedEventsSortOrderEnum
func GetListRelatedEventsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRelatedEventsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRelatedEventsSortOrderEnum(val string) (ListRelatedEventsSortOrderEnum, bool) {
	mappingListRelatedEventsSortOrderEnumIgnoreCase := make(map[string]ListRelatedEventsSortOrderEnum)
	for k, v := range mappingListRelatedEventsSortOrderEnum {
		mappingListRelatedEventsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRelatedEventsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListRelatedEventsSortByEnum Enum with underlying type: string
type ListRelatedEventsSortByEnum string

// Set of constants representing the allowable values for ListRelatedEventsSortByEnum
const (
	ListRelatedEventsSortByInstanceid       ListRelatedEventsSortByEnum = "instanceId"
	ListRelatedEventsSortById               ListRelatedEventsSortByEnum = "id"
	ListRelatedEventsSortByEventfingerprint ListRelatedEventsSortByEnum = "eventFingerprint"
)

var mappingListRelatedEventsSortByEnum = map[string]ListRelatedEventsSortByEnum{
	"instanceId":       ListRelatedEventsSortByInstanceid,
	"id":               ListRelatedEventsSortById,
	"eventFingerprint": ListRelatedEventsSortByEventfingerprint,
}

// GetListRelatedEventsSortByEnumValues Enumerates the set of values for ListRelatedEventsSortByEnum
func GetListRelatedEventsSortByEnumValues() []ListRelatedEventsSortByEnum {
	values := make([]ListRelatedEventsSortByEnum, 0)
	for _, v := range mappingListRelatedEventsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRelatedEventsSortByEnumStringValues Enumerates the set of values in String for ListRelatedEventsSortByEnum
func GetListRelatedEventsSortByEnumStringValues() []string {
	return []string{
		"instanceId",
		"id",
		"eventFingerprint",
	}
}

// GetMappingListRelatedEventsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRelatedEventsSortByEnum(val string) (ListRelatedEventsSortByEnum, bool) {
	mappingListRelatedEventsSortByEnumIgnoreCase := make(map[string]ListRelatedEventsSortByEnum)
	for k, v := range mappingListRelatedEventsSortByEnum {
		mappingListRelatedEventsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRelatedEventsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
