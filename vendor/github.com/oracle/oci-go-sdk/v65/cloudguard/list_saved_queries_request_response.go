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

// ListSavedQueriesRequest wrapper for the ListSavedQueries operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListSavedQueries.go.html to see an example of how to use ListSavedQueriesRequest.
type ListSavedQueriesRequest struct {

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
	AccessLevel ListSavedQueriesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The sort order to use
	SortOrder ListSavedQueriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListSavedQueriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSavedQueriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSavedQueriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSavedQueriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSavedQueriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSavedQueriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSavedQueriesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSavedQueriesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSavedQueriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSavedQueriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSavedQueriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSavedQueriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSavedQueriesResponse wrapper for the ListSavedQueries operation
type ListSavedQueriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SavedQueryCollection instances
	SavedQueryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSavedQueriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSavedQueriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSavedQueriesAccessLevelEnum Enum with underlying type: string
type ListSavedQueriesAccessLevelEnum string

// Set of constants representing the allowable values for ListSavedQueriesAccessLevelEnum
const (
	ListSavedQueriesAccessLevelRestricted ListSavedQueriesAccessLevelEnum = "RESTRICTED"
	ListSavedQueriesAccessLevelAccessible ListSavedQueriesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSavedQueriesAccessLevelEnum = map[string]ListSavedQueriesAccessLevelEnum{
	"RESTRICTED": ListSavedQueriesAccessLevelRestricted,
	"ACCESSIBLE": ListSavedQueriesAccessLevelAccessible,
}

var mappingListSavedQueriesAccessLevelEnumLowerCase = map[string]ListSavedQueriesAccessLevelEnum{
	"restricted": ListSavedQueriesAccessLevelRestricted,
	"accessible": ListSavedQueriesAccessLevelAccessible,
}

// GetListSavedQueriesAccessLevelEnumValues Enumerates the set of values for ListSavedQueriesAccessLevelEnum
func GetListSavedQueriesAccessLevelEnumValues() []ListSavedQueriesAccessLevelEnum {
	values := make([]ListSavedQueriesAccessLevelEnum, 0)
	for _, v := range mappingListSavedQueriesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSavedQueriesAccessLevelEnumStringValues Enumerates the set of values in String for ListSavedQueriesAccessLevelEnum
func GetListSavedQueriesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSavedQueriesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSavedQueriesAccessLevelEnum(val string) (ListSavedQueriesAccessLevelEnum, bool) {
	enum, ok := mappingListSavedQueriesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSavedQueriesSortOrderEnum Enum with underlying type: string
type ListSavedQueriesSortOrderEnum string

// Set of constants representing the allowable values for ListSavedQueriesSortOrderEnum
const (
	ListSavedQueriesSortOrderAsc  ListSavedQueriesSortOrderEnum = "ASC"
	ListSavedQueriesSortOrderDesc ListSavedQueriesSortOrderEnum = "DESC"
)

var mappingListSavedQueriesSortOrderEnum = map[string]ListSavedQueriesSortOrderEnum{
	"ASC":  ListSavedQueriesSortOrderAsc,
	"DESC": ListSavedQueriesSortOrderDesc,
}

var mappingListSavedQueriesSortOrderEnumLowerCase = map[string]ListSavedQueriesSortOrderEnum{
	"asc":  ListSavedQueriesSortOrderAsc,
	"desc": ListSavedQueriesSortOrderDesc,
}

// GetListSavedQueriesSortOrderEnumValues Enumerates the set of values for ListSavedQueriesSortOrderEnum
func GetListSavedQueriesSortOrderEnumValues() []ListSavedQueriesSortOrderEnum {
	values := make([]ListSavedQueriesSortOrderEnum, 0)
	for _, v := range mappingListSavedQueriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSavedQueriesSortOrderEnumStringValues Enumerates the set of values in String for ListSavedQueriesSortOrderEnum
func GetListSavedQueriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSavedQueriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSavedQueriesSortOrderEnum(val string) (ListSavedQueriesSortOrderEnum, bool) {
	enum, ok := mappingListSavedQueriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSavedQueriesSortByEnum Enum with underlying type: string
type ListSavedQueriesSortByEnum string

// Set of constants representing the allowable values for ListSavedQueriesSortByEnum
const (
	ListSavedQueriesSortByTimecreated ListSavedQueriesSortByEnum = "timeCreated"
	ListSavedQueriesSortByDisplayname ListSavedQueriesSortByEnum = "displayName"
)

var mappingListSavedQueriesSortByEnum = map[string]ListSavedQueriesSortByEnum{
	"timeCreated": ListSavedQueriesSortByTimecreated,
	"displayName": ListSavedQueriesSortByDisplayname,
}

var mappingListSavedQueriesSortByEnumLowerCase = map[string]ListSavedQueriesSortByEnum{
	"timecreated": ListSavedQueriesSortByTimecreated,
	"displayname": ListSavedQueriesSortByDisplayname,
}

// GetListSavedQueriesSortByEnumValues Enumerates the set of values for ListSavedQueriesSortByEnum
func GetListSavedQueriesSortByEnumValues() []ListSavedQueriesSortByEnum {
	values := make([]ListSavedQueriesSortByEnum, 0)
	for _, v := range mappingListSavedQueriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSavedQueriesSortByEnumStringValues Enumerates the set of values in String for ListSavedQueriesSortByEnum
func GetListSavedQueriesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSavedQueriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSavedQueriesSortByEnum(val string) (ListSavedQueriesSortByEnum, bool) {
	enum, ok := mappingListSavedQueriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
