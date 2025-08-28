// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFleetErrorsRequest wrapper for the ListFleetErrors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListFleetErrors.go.html to see an example of how to use ListFleetErrorsRequest.
type ListFleetErrorsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Flag to determine whether the info should be gathered only in the compartment or in the compartment and its subcompartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// The ID of the Fleet.
	FleetId *string `mandatory:"false" contributesTo:"query" name:"fleetId"`

	// If specified, only errors with a first seen time earlier than this parameter will be included in the search (formatted according to RFC3339).
	TimeFirstSeenLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeFirstSeenLessThanOrEqualTo"`

	// If specified, only errors with a first seen time later than this parameter will be included in the search (formatted according to RFC3339).
	TimeFirstSeenGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeFirstSeenGreaterThanOrEqualTo"`

	// If specified, only errors with a last seen time earlier than this parameter will be included in the search (formatted according to RFC3339).
	TimeLastSeenLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastSeenLessThanOrEqualTo"`

	// If specified, only errors with a last seen time later than this parameter will be included in the search (formatted according to RFC3339).
	TimeLastSeenGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastSeenGreaterThanOrEqualTo"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort FleetError. Only one sort order may be provided.
	// Default order is **descending**.
	// If no value is specified _timeLastSeen_ is default.
	SortBy ListFleetErrorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListFleetErrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFleetErrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFleetErrorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFleetErrorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFleetErrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFleetErrorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFleetErrorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFleetErrorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFleetErrorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFleetErrorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFleetErrorsResponse wrapper for the ListFleetErrors operation
type ListFleetErrorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FleetErrorCollection instances
	FleetErrorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFleetErrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFleetErrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFleetErrorsSortByEnum Enum with underlying type: string
type ListFleetErrorsSortByEnum string

// Set of constants representing the allowable values for ListFleetErrorsSortByEnum
const (
	ListFleetErrorsSortByFirstSeen ListFleetErrorsSortByEnum = "TIME_FIRST_SEEN"
	ListFleetErrorsSortByLastSeen  ListFleetErrorsSortByEnum = "TIME_LAST_SEEN"
)

var mappingListFleetErrorsSortByEnum = map[string]ListFleetErrorsSortByEnum{
	"TIME_FIRST_SEEN": ListFleetErrorsSortByFirstSeen,
	"TIME_LAST_SEEN":  ListFleetErrorsSortByLastSeen,
}

var mappingListFleetErrorsSortByEnumLowerCase = map[string]ListFleetErrorsSortByEnum{
	"time_first_seen": ListFleetErrorsSortByFirstSeen,
	"time_last_seen":  ListFleetErrorsSortByLastSeen,
}

// GetListFleetErrorsSortByEnumValues Enumerates the set of values for ListFleetErrorsSortByEnum
func GetListFleetErrorsSortByEnumValues() []ListFleetErrorsSortByEnum {
	values := make([]ListFleetErrorsSortByEnum, 0)
	for _, v := range mappingListFleetErrorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetErrorsSortByEnumStringValues Enumerates the set of values in String for ListFleetErrorsSortByEnum
func GetListFleetErrorsSortByEnumStringValues() []string {
	return []string{
		"TIME_FIRST_SEEN",
		"TIME_LAST_SEEN",
	}
}

// GetMappingListFleetErrorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetErrorsSortByEnum(val string) (ListFleetErrorsSortByEnum, bool) {
	enum, ok := mappingListFleetErrorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFleetErrorsSortOrderEnum Enum with underlying type: string
type ListFleetErrorsSortOrderEnum string

// Set of constants representing the allowable values for ListFleetErrorsSortOrderEnum
const (
	ListFleetErrorsSortOrderAsc  ListFleetErrorsSortOrderEnum = "ASC"
	ListFleetErrorsSortOrderDesc ListFleetErrorsSortOrderEnum = "DESC"
)

var mappingListFleetErrorsSortOrderEnum = map[string]ListFleetErrorsSortOrderEnum{
	"ASC":  ListFleetErrorsSortOrderAsc,
	"DESC": ListFleetErrorsSortOrderDesc,
}

var mappingListFleetErrorsSortOrderEnumLowerCase = map[string]ListFleetErrorsSortOrderEnum{
	"asc":  ListFleetErrorsSortOrderAsc,
	"desc": ListFleetErrorsSortOrderDesc,
}

// GetListFleetErrorsSortOrderEnumValues Enumerates the set of values for ListFleetErrorsSortOrderEnum
func GetListFleetErrorsSortOrderEnumValues() []ListFleetErrorsSortOrderEnum {
	values := make([]ListFleetErrorsSortOrderEnum, 0)
	for _, v := range mappingListFleetErrorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetErrorsSortOrderEnumStringValues Enumerates the set of values in String for ListFleetErrorsSortOrderEnum
func GetListFleetErrorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFleetErrorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetErrorsSortOrderEnum(val string) (ListFleetErrorsSortOrderEnum, bool) {
	enum, ok := mappingListFleetErrorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
