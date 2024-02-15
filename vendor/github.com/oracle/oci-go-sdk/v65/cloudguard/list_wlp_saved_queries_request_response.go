// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWlpSavedQueriesRequest wrapper for the ListWlpSavedQueries operation
type ListWlpSavedQueriesRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// returned depending on the setting of `accessLevel`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel ListWlpSavedQueriesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The sort order to use
	SortOrder ListWlpSavedQueriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListWlpSavedQueriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWlpSavedQueriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWlpSavedQueriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWlpSavedQueriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWlpSavedQueriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWlpSavedQueriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWlpSavedQueriesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListWlpSavedQueriesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlpSavedQueriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWlpSavedQueriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlpSavedQueriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWlpSavedQueriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWlpSavedQueriesResponse wrapper for the ListWlpSavedQueries operation
type ListWlpSavedQueriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WlpSavedQueryCollection instances
	WlpSavedQueryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWlpSavedQueriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWlpSavedQueriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWlpSavedQueriesAccessLevelEnum Enum with underlying type: string
type ListWlpSavedQueriesAccessLevelEnum string

// Set of constants representing the allowable values for ListWlpSavedQueriesAccessLevelEnum
const (
	ListWlpSavedQueriesAccessLevelRestricted ListWlpSavedQueriesAccessLevelEnum = "RESTRICTED"
	ListWlpSavedQueriesAccessLevelAccessible ListWlpSavedQueriesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListWlpSavedQueriesAccessLevelEnum = map[string]ListWlpSavedQueriesAccessLevelEnum{
	"RESTRICTED": ListWlpSavedQueriesAccessLevelRestricted,
	"ACCESSIBLE": ListWlpSavedQueriesAccessLevelAccessible,
}

var mappingListWlpSavedQueriesAccessLevelEnumLowerCase = map[string]ListWlpSavedQueriesAccessLevelEnum{
	"restricted": ListWlpSavedQueriesAccessLevelRestricted,
	"accessible": ListWlpSavedQueriesAccessLevelAccessible,
}

// GetListWlpSavedQueriesAccessLevelEnumValues Enumerates the set of values for ListWlpSavedQueriesAccessLevelEnum
func GetListWlpSavedQueriesAccessLevelEnumValues() []ListWlpSavedQueriesAccessLevelEnum {
	values := make([]ListWlpSavedQueriesAccessLevelEnum, 0)
	for _, v := range mappingListWlpSavedQueriesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlpSavedQueriesAccessLevelEnumStringValues Enumerates the set of values in String for ListWlpSavedQueriesAccessLevelEnum
func GetListWlpSavedQueriesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListWlpSavedQueriesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlpSavedQueriesAccessLevelEnum(val string) (ListWlpSavedQueriesAccessLevelEnum, bool) {
	enum, ok := mappingListWlpSavedQueriesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlpSavedQueriesSortOrderEnum Enum with underlying type: string
type ListWlpSavedQueriesSortOrderEnum string

// Set of constants representing the allowable values for ListWlpSavedQueriesSortOrderEnum
const (
	ListWlpSavedQueriesSortOrderAsc  ListWlpSavedQueriesSortOrderEnum = "ASC"
	ListWlpSavedQueriesSortOrderDesc ListWlpSavedQueriesSortOrderEnum = "DESC"
)

var mappingListWlpSavedQueriesSortOrderEnum = map[string]ListWlpSavedQueriesSortOrderEnum{
	"ASC":  ListWlpSavedQueriesSortOrderAsc,
	"DESC": ListWlpSavedQueriesSortOrderDesc,
}

var mappingListWlpSavedQueriesSortOrderEnumLowerCase = map[string]ListWlpSavedQueriesSortOrderEnum{
	"asc":  ListWlpSavedQueriesSortOrderAsc,
	"desc": ListWlpSavedQueriesSortOrderDesc,
}

// GetListWlpSavedQueriesSortOrderEnumValues Enumerates the set of values for ListWlpSavedQueriesSortOrderEnum
func GetListWlpSavedQueriesSortOrderEnumValues() []ListWlpSavedQueriesSortOrderEnum {
	values := make([]ListWlpSavedQueriesSortOrderEnum, 0)
	for _, v := range mappingListWlpSavedQueriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlpSavedQueriesSortOrderEnumStringValues Enumerates the set of values in String for ListWlpSavedQueriesSortOrderEnum
func GetListWlpSavedQueriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWlpSavedQueriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlpSavedQueriesSortOrderEnum(val string) (ListWlpSavedQueriesSortOrderEnum, bool) {
	enum, ok := mappingListWlpSavedQueriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlpSavedQueriesSortByEnum Enum with underlying type: string
type ListWlpSavedQueriesSortByEnum string

// Set of constants representing the allowable values for ListWlpSavedQueriesSortByEnum
const (
	ListWlpSavedQueriesSortByTimecreated ListWlpSavedQueriesSortByEnum = "timeCreated"
	ListWlpSavedQueriesSortByDisplayname ListWlpSavedQueriesSortByEnum = "displayName"
)

var mappingListWlpSavedQueriesSortByEnum = map[string]ListWlpSavedQueriesSortByEnum{
	"timeCreated": ListWlpSavedQueriesSortByTimecreated,
	"displayName": ListWlpSavedQueriesSortByDisplayname,
}

var mappingListWlpSavedQueriesSortByEnumLowerCase = map[string]ListWlpSavedQueriesSortByEnum{
	"timecreated": ListWlpSavedQueriesSortByTimecreated,
	"displayname": ListWlpSavedQueriesSortByDisplayname,
}

// GetListWlpSavedQueriesSortByEnumValues Enumerates the set of values for ListWlpSavedQueriesSortByEnum
func GetListWlpSavedQueriesSortByEnumValues() []ListWlpSavedQueriesSortByEnum {
	values := make([]ListWlpSavedQueriesSortByEnum, 0)
	for _, v := range mappingListWlpSavedQueriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlpSavedQueriesSortByEnumStringValues Enumerates the set of values in String for ListWlpSavedQueriesSortByEnum
func GetListWlpSavedQueriesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListWlpSavedQueriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlpSavedQueriesSortByEnum(val string) (ListWlpSavedQueriesSortByEnum, bool) {
	enum, ok := mappingListWlpSavedQueriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
