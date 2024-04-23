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

// ListAdhocQueriesRequest wrapper for the ListAdhocQueries operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListAdhocQueries.go.html to see an example of how to use ListAdhocQueriesRequest.
type ListAdhocQueriesRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The status of the adhoc query created. Default value for state is provisioning. If no value is specified state is provisioning.
	AdhocQueryStatus ListAdhocQueriesAdhocQueryStatusEnum `mandatory:"false" contributesTo:"query" name:"adhocQueryStatus" omitEmpty:"true"`

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
	AccessLevel ListAdhocQueriesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The sort order to use
	SortOrder ListAdhocQueriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListAdhocQueriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAdhocQueriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAdhocQueriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAdhocQueriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAdhocQueriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAdhocQueriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAdhocQueriesAdhocQueryStatusEnum(string(request.AdhocQueryStatus)); !ok && request.AdhocQueryStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdhocQueryStatus: %s. Supported values are: %s.", request.AdhocQueryStatus, strings.Join(GetListAdhocQueriesAdhocQueryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAdhocQueriesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListAdhocQueriesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAdhocQueriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAdhocQueriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAdhocQueriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAdhocQueriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAdhocQueriesResponse wrapper for the ListAdhocQueries operation
type ListAdhocQueriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AdhocQueryCollection instances
	AdhocQueryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAdhocQueriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAdhocQueriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAdhocQueriesAdhocQueryStatusEnum Enum with underlying type: string
type ListAdhocQueriesAdhocQueryStatusEnum string

// Set of constants representing the allowable values for ListAdhocQueriesAdhocQueryStatusEnum
const (
	ListAdhocQueriesAdhocQueryStatusCreating           ListAdhocQueriesAdhocQueryStatusEnum = "CREATING"
	ListAdhocQueriesAdhocQueryStatusCreated            ListAdhocQueriesAdhocQueryStatusEnum = "CREATED"
	ListAdhocQueriesAdhocQueryStatusInProgress         ListAdhocQueriesAdhocQueryStatusEnum = "IN_PROGRESS"
	ListAdhocQueriesAdhocQueryStatusPartiallyCompleted ListAdhocQueriesAdhocQueryStatusEnum = "PARTIALLY_COMPLETED"
	ListAdhocQueriesAdhocQueryStatusExpired            ListAdhocQueriesAdhocQueryStatusEnum = "EXPIRED"
	ListAdhocQueriesAdhocQueryStatusCompleted          ListAdhocQueriesAdhocQueryStatusEnum = "COMPLETED"
	ListAdhocQueriesAdhocQueryStatusFailed             ListAdhocQueriesAdhocQueryStatusEnum = "FAILED"
)

var mappingListAdhocQueriesAdhocQueryStatusEnum = map[string]ListAdhocQueriesAdhocQueryStatusEnum{
	"CREATING":            ListAdhocQueriesAdhocQueryStatusCreating,
	"CREATED":             ListAdhocQueriesAdhocQueryStatusCreated,
	"IN_PROGRESS":         ListAdhocQueriesAdhocQueryStatusInProgress,
	"PARTIALLY_COMPLETED": ListAdhocQueriesAdhocQueryStatusPartiallyCompleted,
	"EXPIRED":             ListAdhocQueriesAdhocQueryStatusExpired,
	"COMPLETED":           ListAdhocQueriesAdhocQueryStatusCompleted,
	"FAILED":              ListAdhocQueriesAdhocQueryStatusFailed,
}

var mappingListAdhocQueriesAdhocQueryStatusEnumLowerCase = map[string]ListAdhocQueriesAdhocQueryStatusEnum{
	"creating":            ListAdhocQueriesAdhocQueryStatusCreating,
	"created":             ListAdhocQueriesAdhocQueryStatusCreated,
	"in_progress":         ListAdhocQueriesAdhocQueryStatusInProgress,
	"partially_completed": ListAdhocQueriesAdhocQueryStatusPartiallyCompleted,
	"expired":             ListAdhocQueriesAdhocQueryStatusExpired,
	"completed":           ListAdhocQueriesAdhocQueryStatusCompleted,
	"failed":              ListAdhocQueriesAdhocQueryStatusFailed,
}

// GetListAdhocQueriesAdhocQueryStatusEnumValues Enumerates the set of values for ListAdhocQueriesAdhocQueryStatusEnum
func GetListAdhocQueriesAdhocQueryStatusEnumValues() []ListAdhocQueriesAdhocQueryStatusEnum {
	values := make([]ListAdhocQueriesAdhocQueryStatusEnum, 0)
	for _, v := range mappingListAdhocQueriesAdhocQueryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdhocQueriesAdhocQueryStatusEnumStringValues Enumerates the set of values in String for ListAdhocQueriesAdhocQueryStatusEnum
func GetListAdhocQueriesAdhocQueryStatusEnumStringValues() []string {
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

// GetMappingListAdhocQueriesAdhocQueryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdhocQueriesAdhocQueryStatusEnum(val string) (ListAdhocQueriesAdhocQueryStatusEnum, bool) {
	enum, ok := mappingListAdhocQueriesAdhocQueryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAdhocQueriesAccessLevelEnum Enum with underlying type: string
type ListAdhocQueriesAccessLevelEnum string

// Set of constants representing the allowable values for ListAdhocQueriesAccessLevelEnum
const (
	ListAdhocQueriesAccessLevelRestricted ListAdhocQueriesAccessLevelEnum = "RESTRICTED"
	ListAdhocQueriesAccessLevelAccessible ListAdhocQueriesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListAdhocQueriesAccessLevelEnum = map[string]ListAdhocQueriesAccessLevelEnum{
	"RESTRICTED": ListAdhocQueriesAccessLevelRestricted,
	"ACCESSIBLE": ListAdhocQueriesAccessLevelAccessible,
}

var mappingListAdhocQueriesAccessLevelEnumLowerCase = map[string]ListAdhocQueriesAccessLevelEnum{
	"restricted": ListAdhocQueriesAccessLevelRestricted,
	"accessible": ListAdhocQueriesAccessLevelAccessible,
}

// GetListAdhocQueriesAccessLevelEnumValues Enumerates the set of values for ListAdhocQueriesAccessLevelEnum
func GetListAdhocQueriesAccessLevelEnumValues() []ListAdhocQueriesAccessLevelEnum {
	values := make([]ListAdhocQueriesAccessLevelEnum, 0)
	for _, v := range mappingListAdhocQueriesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdhocQueriesAccessLevelEnumStringValues Enumerates the set of values in String for ListAdhocQueriesAccessLevelEnum
func GetListAdhocQueriesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListAdhocQueriesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdhocQueriesAccessLevelEnum(val string) (ListAdhocQueriesAccessLevelEnum, bool) {
	enum, ok := mappingListAdhocQueriesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAdhocQueriesSortOrderEnum Enum with underlying type: string
type ListAdhocQueriesSortOrderEnum string

// Set of constants representing the allowable values for ListAdhocQueriesSortOrderEnum
const (
	ListAdhocQueriesSortOrderAsc  ListAdhocQueriesSortOrderEnum = "ASC"
	ListAdhocQueriesSortOrderDesc ListAdhocQueriesSortOrderEnum = "DESC"
)

var mappingListAdhocQueriesSortOrderEnum = map[string]ListAdhocQueriesSortOrderEnum{
	"ASC":  ListAdhocQueriesSortOrderAsc,
	"DESC": ListAdhocQueriesSortOrderDesc,
}

var mappingListAdhocQueriesSortOrderEnumLowerCase = map[string]ListAdhocQueriesSortOrderEnum{
	"asc":  ListAdhocQueriesSortOrderAsc,
	"desc": ListAdhocQueriesSortOrderDesc,
}

// GetListAdhocQueriesSortOrderEnumValues Enumerates the set of values for ListAdhocQueriesSortOrderEnum
func GetListAdhocQueriesSortOrderEnumValues() []ListAdhocQueriesSortOrderEnum {
	values := make([]ListAdhocQueriesSortOrderEnum, 0)
	for _, v := range mappingListAdhocQueriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdhocQueriesSortOrderEnumStringValues Enumerates the set of values in String for ListAdhocQueriesSortOrderEnum
func GetListAdhocQueriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAdhocQueriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdhocQueriesSortOrderEnum(val string) (ListAdhocQueriesSortOrderEnum, bool) {
	enum, ok := mappingListAdhocQueriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAdhocQueriesSortByEnum Enum with underlying type: string
type ListAdhocQueriesSortByEnum string

// Set of constants representing the allowable values for ListAdhocQueriesSortByEnum
const (
	ListAdhocQueriesSortByTimecreated ListAdhocQueriesSortByEnum = "timeCreated"
	ListAdhocQueriesSortByDisplayname ListAdhocQueriesSortByEnum = "displayName"
)

var mappingListAdhocQueriesSortByEnum = map[string]ListAdhocQueriesSortByEnum{
	"timeCreated": ListAdhocQueriesSortByTimecreated,
	"displayName": ListAdhocQueriesSortByDisplayname,
}

var mappingListAdhocQueriesSortByEnumLowerCase = map[string]ListAdhocQueriesSortByEnum{
	"timecreated": ListAdhocQueriesSortByTimecreated,
	"displayname": ListAdhocQueriesSortByDisplayname,
}

// GetListAdhocQueriesSortByEnumValues Enumerates the set of values for ListAdhocQueriesSortByEnum
func GetListAdhocQueriesSortByEnumValues() []ListAdhocQueriesSortByEnum {
	values := make([]ListAdhocQueriesSortByEnum, 0)
	for _, v := range mappingListAdhocQueriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAdhocQueriesSortByEnumStringValues Enumerates the set of values in String for ListAdhocQueriesSortByEnum
func GetListAdhocQueriesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAdhocQueriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAdhocQueriesSortByEnum(val string) (ListAdhocQueriesSortByEnum, bool) {
	enum, ok := mappingListAdhocQueriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
