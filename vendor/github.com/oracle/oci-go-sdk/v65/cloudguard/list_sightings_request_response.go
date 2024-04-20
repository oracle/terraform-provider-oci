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

// ListSightingsRequest wrapper for the ListSightings operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListSightings.go.html to see an example of how to use ListSightingsRequest.
type ListSightingsRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// OCID of the problem.
	ProblemId *string `mandatory:"false" contributesTo:"query" name:"problemId"`

	// OCID of the resource profile.
	ResourceProfileId *string `mandatory:"false" contributesTo:"query" name:"resourceProfileId"`

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
	AccessLevel ListSightingsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListSightingsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListSightingsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Start time for a filter. If start time is not specified, start time will be set to current time - 30 days.
	TimeLastDetectedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastDetectedGreaterThanOrEqualTo"`

	// End time for a filter. If end time is not specified, end time will be set to current time.
	TimeLastDetectedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastDetectedLessThanOrEqualTo"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSightingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSightingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSightingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSightingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSightingsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSightingsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSightingsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSightingsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSightingsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSightingsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSightingsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSightingsResponse wrapper for the ListSightings operation
type ListSightingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SightingCollection instances
	SightingCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSightingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSightingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSightingsAccessLevelEnum Enum with underlying type: string
type ListSightingsAccessLevelEnum string

// Set of constants representing the allowable values for ListSightingsAccessLevelEnum
const (
	ListSightingsAccessLevelRestricted ListSightingsAccessLevelEnum = "RESTRICTED"
	ListSightingsAccessLevelAccessible ListSightingsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSightingsAccessLevelEnum = map[string]ListSightingsAccessLevelEnum{
	"RESTRICTED": ListSightingsAccessLevelRestricted,
	"ACCESSIBLE": ListSightingsAccessLevelAccessible,
}

var mappingListSightingsAccessLevelEnumLowerCase = map[string]ListSightingsAccessLevelEnum{
	"restricted": ListSightingsAccessLevelRestricted,
	"accessible": ListSightingsAccessLevelAccessible,
}

// GetListSightingsAccessLevelEnumValues Enumerates the set of values for ListSightingsAccessLevelEnum
func GetListSightingsAccessLevelEnumValues() []ListSightingsAccessLevelEnum {
	values := make([]ListSightingsAccessLevelEnum, 0)
	for _, v := range mappingListSightingsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSightingsAccessLevelEnumStringValues Enumerates the set of values in String for ListSightingsAccessLevelEnum
func GetListSightingsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSightingsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSightingsAccessLevelEnum(val string) (ListSightingsAccessLevelEnum, bool) {
	enum, ok := mappingListSightingsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSightingsSortOrderEnum Enum with underlying type: string
type ListSightingsSortOrderEnum string

// Set of constants representing the allowable values for ListSightingsSortOrderEnum
const (
	ListSightingsSortOrderAsc  ListSightingsSortOrderEnum = "ASC"
	ListSightingsSortOrderDesc ListSightingsSortOrderEnum = "DESC"
)

var mappingListSightingsSortOrderEnum = map[string]ListSightingsSortOrderEnum{
	"ASC":  ListSightingsSortOrderAsc,
	"DESC": ListSightingsSortOrderDesc,
}

var mappingListSightingsSortOrderEnumLowerCase = map[string]ListSightingsSortOrderEnum{
	"asc":  ListSightingsSortOrderAsc,
	"desc": ListSightingsSortOrderDesc,
}

// GetListSightingsSortOrderEnumValues Enumerates the set of values for ListSightingsSortOrderEnum
func GetListSightingsSortOrderEnumValues() []ListSightingsSortOrderEnum {
	values := make([]ListSightingsSortOrderEnum, 0)
	for _, v := range mappingListSightingsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSightingsSortOrderEnumStringValues Enumerates the set of values in String for ListSightingsSortOrderEnum
func GetListSightingsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSightingsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSightingsSortOrderEnum(val string) (ListSightingsSortOrderEnum, bool) {
	enum, ok := mappingListSightingsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSightingsSortByEnum Enum with underlying type: string
type ListSightingsSortByEnum string

// Set of constants representing the allowable values for ListSightingsSortByEnum
const (
	ListSightingsSortByTimecreated ListSightingsSortByEnum = "timeCreated"
)

var mappingListSightingsSortByEnum = map[string]ListSightingsSortByEnum{
	"timeCreated": ListSightingsSortByTimecreated,
}

var mappingListSightingsSortByEnumLowerCase = map[string]ListSightingsSortByEnum{
	"timecreated": ListSightingsSortByTimecreated,
}

// GetListSightingsSortByEnumValues Enumerates the set of values for ListSightingsSortByEnum
func GetListSightingsSortByEnumValues() []ListSightingsSortByEnum {
	values := make([]ListSightingsSortByEnum, 0)
	for _, v := range mappingListSightingsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSightingsSortByEnumStringValues Enumerates the set of values in String for ListSightingsSortByEnum
func GetListSightingsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListSightingsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSightingsSortByEnum(val string) (ListSightingsSortByEnum, bool) {
	enum, ok := mappingListSightingsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
