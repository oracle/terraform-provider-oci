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

// ListWlpAdhocQueriesRequest wrapper for the ListWlpAdhocQueries operation
type ListWlpAdhocQueriesRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The status of the adhoc query created. Default value for state is provisioning. If no value is specified state is provisioning.
	AdhocQueryStatus ListWlpAdhocQueriesAdhocQueryStatusEnum `mandatory:"false" contributesTo:"query" name:"adhocQueryStatus" omitEmpty:"true"`

	// Start time for a filter. If start time is not specified, start time will be set to current time - 30 days.
	TimeStartedFilterQueryParam *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStartedFilterQueryParam"`

	// End time for a filter. If end time is not specified, end time will be set to current time.
	TimeEndedFilterQueryParam *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEndedFilterQueryParam"`

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
	AccessLevel ListWlpAdhocQueriesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The sort order to use
	SortOrder ListWlpAdhocQueriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListWlpAdhocQueriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWlpAdhocQueriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWlpAdhocQueriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWlpAdhocQueriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWlpAdhocQueriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWlpAdhocQueriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWlpAdhocQueriesAdhocQueryStatusEnum(string(request.AdhocQueryStatus)); !ok && request.AdhocQueryStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdhocQueryStatus: %s. Supported values are: %s.", request.AdhocQueryStatus, strings.Join(GetListWlpAdhocQueriesAdhocQueryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlpAdhocQueriesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListWlpAdhocQueriesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlpAdhocQueriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWlpAdhocQueriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlpAdhocQueriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWlpAdhocQueriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWlpAdhocQueriesResponse wrapper for the ListWlpAdhocQueries operation
type ListWlpAdhocQueriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WlpAdhocQueryCollection instances
	WlpAdhocQueryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWlpAdhocQueriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWlpAdhocQueriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWlpAdhocQueriesAdhocQueryStatusEnum Enum with underlying type: string
type ListWlpAdhocQueriesAdhocQueryStatusEnum string

// Set of constants representing the allowable values for ListWlpAdhocQueriesAdhocQueryStatusEnum
const (
	ListWlpAdhocQueriesAdhocQueryStatusCreating           ListWlpAdhocQueriesAdhocQueryStatusEnum = "CREATING"
	ListWlpAdhocQueriesAdhocQueryStatusCreated            ListWlpAdhocQueriesAdhocQueryStatusEnum = "CREATED"
	ListWlpAdhocQueriesAdhocQueryStatusInProgress         ListWlpAdhocQueriesAdhocQueryStatusEnum = "IN_PROGRESS"
	ListWlpAdhocQueriesAdhocQueryStatusPartiallyCompleted ListWlpAdhocQueriesAdhocQueryStatusEnum = "PARTIALLY_COMPLETED"
	ListWlpAdhocQueriesAdhocQueryStatusExpired            ListWlpAdhocQueriesAdhocQueryStatusEnum = "EXPIRED"
	ListWlpAdhocQueriesAdhocQueryStatusCompleted          ListWlpAdhocQueriesAdhocQueryStatusEnum = "COMPLETED"
	ListWlpAdhocQueriesAdhocQueryStatusFailed             ListWlpAdhocQueriesAdhocQueryStatusEnum = "FAILED"
)

var mappingListWlpAdhocQueriesAdhocQueryStatusEnum = map[string]ListWlpAdhocQueriesAdhocQueryStatusEnum{
	"CREATING":            ListWlpAdhocQueriesAdhocQueryStatusCreating,
	"CREATED":             ListWlpAdhocQueriesAdhocQueryStatusCreated,
	"IN_PROGRESS":         ListWlpAdhocQueriesAdhocQueryStatusInProgress,
	"PARTIALLY_COMPLETED": ListWlpAdhocQueriesAdhocQueryStatusPartiallyCompleted,
	"EXPIRED":             ListWlpAdhocQueriesAdhocQueryStatusExpired,
	"COMPLETED":           ListWlpAdhocQueriesAdhocQueryStatusCompleted,
	"FAILED":              ListWlpAdhocQueriesAdhocQueryStatusFailed,
}

var mappingListWlpAdhocQueriesAdhocQueryStatusEnumLowerCase = map[string]ListWlpAdhocQueriesAdhocQueryStatusEnum{
	"creating":            ListWlpAdhocQueriesAdhocQueryStatusCreating,
	"created":             ListWlpAdhocQueriesAdhocQueryStatusCreated,
	"in_progress":         ListWlpAdhocQueriesAdhocQueryStatusInProgress,
	"partially_completed": ListWlpAdhocQueriesAdhocQueryStatusPartiallyCompleted,
	"expired":             ListWlpAdhocQueriesAdhocQueryStatusExpired,
	"completed":           ListWlpAdhocQueriesAdhocQueryStatusCompleted,
	"failed":              ListWlpAdhocQueriesAdhocQueryStatusFailed,
}

// GetListWlpAdhocQueriesAdhocQueryStatusEnumValues Enumerates the set of values for ListWlpAdhocQueriesAdhocQueryStatusEnum
func GetListWlpAdhocQueriesAdhocQueryStatusEnumValues() []ListWlpAdhocQueriesAdhocQueryStatusEnum {
	values := make([]ListWlpAdhocQueriesAdhocQueryStatusEnum, 0)
	for _, v := range mappingListWlpAdhocQueriesAdhocQueryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlpAdhocQueriesAdhocQueryStatusEnumStringValues Enumerates the set of values in String for ListWlpAdhocQueriesAdhocQueryStatusEnum
func GetListWlpAdhocQueriesAdhocQueryStatusEnumStringValues() []string {
	return []string{
		"CREATING",
		"CREATED",
		"IN_PROGRESS",
		"PARTIALLY_COMPLETED",
		"EXPIRED",
		"COMPLETED",
		"FAILED",
	}
}

// GetMappingListWlpAdhocQueriesAdhocQueryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlpAdhocQueriesAdhocQueryStatusEnum(val string) (ListWlpAdhocQueriesAdhocQueryStatusEnum, bool) {
	enum, ok := mappingListWlpAdhocQueriesAdhocQueryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlpAdhocQueriesAccessLevelEnum Enum with underlying type: string
type ListWlpAdhocQueriesAccessLevelEnum string

// Set of constants representing the allowable values for ListWlpAdhocQueriesAccessLevelEnum
const (
	ListWlpAdhocQueriesAccessLevelRestricted ListWlpAdhocQueriesAccessLevelEnum = "RESTRICTED"
	ListWlpAdhocQueriesAccessLevelAccessible ListWlpAdhocQueriesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListWlpAdhocQueriesAccessLevelEnum = map[string]ListWlpAdhocQueriesAccessLevelEnum{
	"RESTRICTED": ListWlpAdhocQueriesAccessLevelRestricted,
	"ACCESSIBLE": ListWlpAdhocQueriesAccessLevelAccessible,
}

var mappingListWlpAdhocQueriesAccessLevelEnumLowerCase = map[string]ListWlpAdhocQueriesAccessLevelEnum{
	"restricted": ListWlpAdhocQueriesAccessLevelRestricted,
	"accessible": ListWlpAdhocQueriesAccessLevelAccessible,
}

// GetListWlpAdhocQueriesAccessLevelEnumValues Enumerates the set of values for ListWlpAdhocQueriesAccessLevelEnum
func GetListWlpAdhocQueriesAccessLevelEnumValues() []ListWlpAdhocQueriesAccessLevelEnum {
	values := make([]ListWlpAdhocQueriesAccessLevelEnum, 0)
	for _, v := range mappingListWlpAdhocQueriesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlpAdhocQueriesAccessLevelEnumStringValues Enumerates the set of values in String for ListWlpAdhocQueriesAccessLevelEnum
func GetListWlpAdhocQueriesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListWlpAdhocQueriesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlpAdhocQueriesAccessLevelEnum(val string) (ListWlpAdhocQueriesAccessLevelEnum, bool) {
	enum, ok := mappingListWlpAdhocQueriesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlpAdhocQueriesSortOrderEnum Enum with underlying type: string
type ListWlpAdhocQueriesSortOrderEnum string

// Set of constants representing the allowable values for ListWlpAdhocQueriesSortOrderEnum
const (
	ListWlpAdhocQueriesSortOrderAsc  ListWlpAdhocQueriesSortOrderEnum = "ASC"
	ListWlpAdhocQueriesSortOrderDesc ListWlpAdhocQueriesSortOrderEnum = "DESC"
)

var mappingListWlpAdhocQueriesSortOrderEnum = map[string]ListWlpAdhocQueriesSortOrderEnum{
	"ASC":  ListWlpAdhocQueriesSortOrderAsc,
	"DESC": ListWlpAdhocQueriesSortOrderDesc,
}

var mappingListWlpAdhocQueriesSortOrderEnumLowerCase = map[string]ListWlpAdhocQueriesSortOrderEnum{
	"asc":  ListWlpAdhocQueriesSortOrderAsc,
	"desc": ListWlpAdhocQueriesSortOrderDesc,
}

// GetListWlpAdhocQueriesSortOrderEnumValues Enumerates the set of values for ListWlpAdhocQueriesSortOrderEnum
func GetListWlpAdhocQueriesSortOrderEnumValues() []ListWlpAdhocQueriesSortOrderEnum {
	values := make([]ListWlpAdhocQueriesSortOrderEnum, 0)
	for _, v := range mappingListWlpAdhocQueriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlpAdhocQueriesSortOrderEnumStringValues Enumerates the set of values in String for ListWlpAdhocQueriesSortOrderEnum
func GetListWlpAdhocQueriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWlpAdhocQueriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlpAdhocQueriesSortOrderEnum(val string) (ListWlpAdhocQueriesSortOrderEnum, bool) {
	enum, ok := mappingListWlpAdhocQueriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlpAdhocQueriesSortByEnum Enum with underlying type: string
type ListWlpAdhocQueriesSortByEnum string

// Set of constants representing the allowable values for ListWlpAdhocQueriesSortByEnum
const (
	ListWlpAdhocQueriesSortByTimecreated ListWlpAdhocQueriesSortByEnum = "timeCreated"
	ListWlpAdhocQueriesSortByDisplayname ListWlpAdhocQueriesSortByEnum = "displayName"
)

var mappingListWlpAdhocQueriesSortByEnum = map[string]ListWlpAdhocQueriesSortByEnum{
	"timeCreated": ListWlpAdhocQueriesSortByTimecreated,
	"displayName": ListWlpAdhocQueriesSortByDisplayname,
}

var mappingListWlpAdhocQueriesSortByEnumLowerCase = map[string]ListWlpAdhocQueriesSortByEnum{
	"timecreated": ListWlpAdhocQueriesSortByTimecreated,
	"displayname": ListWlpAdhocQueriesSortByDisplayname,
}

// GetListWlpAdhocQueriesSortByEnumValues Enumerates the set of values for ListWlpAdhocQueriesSortByEnum
func GetListWlpAdhocQueriesSortByEnumValues() []ListWlpAdhocQueriesSortByEnum {
	values := make([]ListWlpAdhocQueriesSortByEnum, 0)
	for _, v := range mappingListWlpAdhocQueriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlpAdhocQueriesSortByEnumStringValues Enumerates the set of values in String for ListWlpAdhocQueriesSortByEnum
func GetListWlpAdhocQueriesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListWlpAdhocQueriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlpAdhocQueriesSortByEnum(val string) (ListWlpAdhocQueriesSortByEnum, bool) {
	enum, ok := mappingListWlpAdhocQueriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
