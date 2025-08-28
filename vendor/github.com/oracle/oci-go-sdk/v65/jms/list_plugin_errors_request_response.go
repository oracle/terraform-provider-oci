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

// ListPluginErrorsRequest wrapper for the ListPluginErrors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListPluginErrors.go.html to see an example of how to use ListPluginErrorsRequest.
type ListPluginErrorsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Flag to determine whether the info should be gathered only in the compartment or in the compartment and its subcompartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// The Fleet-unique identifier of the managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

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

	// The field to sort PluginError. Only one sort order may be provided.
	// Default order is **descending**.
	// If no value is specified _timeLastSeen_ is default.
	SortBy ListPluginErrorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListPluginErrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPluginErrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPluginErrorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPluginErrorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPluginErrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPluginErrorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPluginErrorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPluginErrorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPluginErrorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPluginErrorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPluginErrorsResponse wrapper for the ListPluginErrors operation
type ListPluginErrorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PluginErrorCollection instances
	PluginErrorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPluginErrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPluginErrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPluginErrorsSortByEnum Enum with underlying type: string
type ListPluginErrorsSortByEnum string

// Set of constants representing the allowable values for ListPluginErrorsSortByEnum
const (
	ListPluginErrorsSortByFirstSeen ListPluginErrorsSortByEnum = "TIME_FIRST_SEEN"
	ListPluginErrorsSortByLastSeen  ListPluginErrorsSortByEnum = "TIME_LAST_SEEN"
)

var mappingListPluginErrorsSortByEnum = map[string]ListPluginErrorsSortByEnum{
	"TIME_FIRST_SEEN": ListPluginErrorsSortByFirstSeen,
	"TIME_LAST_SEEN":  ListPluginErrorsSortByLastSeen,
}

var mappingListPluginErrorsSortByEnumLowerCase = map[string]ListPluginErrorsSortByEnum{
	"time_first_seen": ListPluginErrorsSortByFirstSeen,
	"time_last_seen":  ListPluginErrorsSortByLastSeen,
}

// GetListPluginErrorsSortByEnumValues Enumerates the set of values for ListPluginErrorsSortByEnum
func GetListPluginErrorsSortByEnumValues() []ListPluginErrorsSortByEnum {
	values := make([]ListPluginErrorsSortByEnum, 0)
	for _, v := range mappingListPluginErrorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPluginErrorsSortByEnumStringValues Enumerates the set of values in String for ListPluginErrorsSortByEnum
func GetListPluginErrorsSortByEnumStringValues() []string {
	return []string{
		"TIME_FIRST_SEEN",
		"TIME_LAST_SEEN",
	}
}

// GetMappingListPluginErrorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPluginErrorsSortByEnum(val string) (ListPluginErrorsSortByEnum, bool) {
	enum, ok := mappingListPluginErrorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPluginErrorsSortOrderEnum Enum with underlying type: string
type ListPluginErrorsSortOrderEnum string

// Set of constants representing the allowable values for ListPluginErrorsSortOrderEnum
const (
	ListPluginErrorsSortOrderAsc  ListPluginErrorsSortOrderEnum = "ASC"
	ListPluginErrorsSortOrderDesc ListPluginErrorsSortOrderEnum = "DESC"
)

var mappingListPluginErrorsSortOrderEnum = map[string]ListPluginErrorsSortOrderEnum{
	"ASC":  ListPluginErrorsSortOrderAsc,
	"DESC": ListPluginErrorsSortOrderDesc,
}

var mappingListPluginErrorsSortOrderEnumLowerCase = map[string]ListPluginErrorsSortOrderEnum{
	"asc":  ListPluginErrorsSortOrderAsc,
	"desc": ListPluginErrorsSortOrderDesc,
}

// GetListPluginErrorsSortOrderEnumValues Enumerates the set of values for ListPluginErrorsSortOrderEnum
func GetListPluginErrorsSortOrderEnumValues() []ListPluginErrorsSortOrderEnum {
	values := make([]ListPluginErrorsSortOrderEnum, 0)
	for _, v := range mappingListPluginErrorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPluginErrorsSortOrderEnumStringValues Enumerates the set of values in String for ListPluginErrorsSortOrderEnum
func GetListPluginErrorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPluginErrorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPluginErrorsSortOrderEnum(val string) (ListPluginErrorsSortOrderEnum, bool) {
	enum, ok := mappingListPluginErrorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
