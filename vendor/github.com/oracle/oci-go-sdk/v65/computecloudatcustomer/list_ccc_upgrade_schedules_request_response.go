// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCccUpgradeSchedulesRequest wrapper for the ListCccUpgradeSchedules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computecloudatcustomer/ListCccUpgradeSchedules.go.html to see an example of how to use ListCccUpgradeSchedulesRequest.
type ListCccUpgradeSchedulesRequest struct {

	// Compute Cloud@Customer upgrade schedule
	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CccUpgradeScheduleId *string `mandatory:"false" contributesTo:"query" name:"cccUpgradeScheduleId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to
	// list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and
	// sub-compartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no
	// partial results are displayed.
	AccessLevel ListCccUpgradeSchedulesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return resources only when their lifecycleState matches the given lifecycleState.
	LifecycleState CccUpgradeScheduleLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources whose display name contains the substring.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListCccUpgradeSchedulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListCccUpgradeSchedulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCccUpgradeSchedulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCccUpgradeSchedulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCccUpgradeSchedulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCccUpgradeSchedulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCccUpgradeSchedulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCccUpgradeSchedulesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListCccUpgradeSchedulesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCccUpgradeScheduleLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCccUpgradeScheduleLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCccUpgradeSchedulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCccUpgradeSchedulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCccUpgradeSchedulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCccUpgradeSchedulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCccUpgradeSchedulesResponse wrapper for the ListCccUpgradeSchedules operation
type ListCccUpgradeSchedulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CccUpgradeScheduleCollection instances
	CccUpgradeScheduleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCccUpgradeSchedulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCccUpgradeSchedulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCccUpgradeSchedulesAccessLevelEnum Enum with underlying type: string
type ListCccUpgradeSchedulesAccessLevelEnum string

// Set of constants representing the allowable values for ListCccUpgradeSchedulesAccessLevelEnum
const (
	ListCccUpgradeSchedulesAccessLevelRestricted ListCccUpgradeSchedulesAccessLevelEnum = "RESTRICTED"
	ListCccUpgradeSchedulesAccessLevelAccessible ListCccUpgradeSchedulesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListCccUpgradeSchedulesAccessLevelEnum = map[string]ListCccUpgradeSchedulesAccessLevelEnum{
	"RESTRICTED": ListCccUpgradeSchedulesAccessLevelRestricted,
	"ACCESSIBLE": ListCccUpgradeSchedulesAccessLevelAccessible,
}

var mappingListCccUpgradeSchedulesAccessLevelEnumLowerCase = map[string]ListCccUpgradeSchedulesAccessLevelEnum{
	"restricted": ListCccUpgradeSchedulesAccessLevelRestricted,
	"accessible": ListCccUpgradeSchedulesAccessLevelAccessible,
}

// GetListCccUpgradeSchedulesAccessLevelEnumValues Enumerates the set of values for ListCccUpgradeSchedulesAccessLevelEnum
func GetListCccUpgradeSchedulesAccessLevelEnumValues() []ListCccUpgradeSchedulesAccessLevelEnum {
	values := make([]ListCccUpgradeSchedulesAccessLevelEnum, 0)
	for _, v := range mappingListCccUpgradeSchedulesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccUpgradeSchedulesAccessLevelEnumStringValues Enumerates the set of values in String for ListCccUpgradeSchedulesAccessLevelEnum
func GetListCccUpgradeSchedulesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListCccUpgradeSchedulesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccUpgradeSchedulesAccessLevelEnum(val string) (ListCccUpgradeSchedulesAccessLevelEnum, bool) {
	enum, ok := mappingListCccUpgradeSchedulesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCccUpgradeSchedulesSortOrderEnum Enum with underlying type: string
type ListCccUpgradeSchedulesSortOrderEnum string

// Set of constants representing the allowable values for ListCccUpgradeSchedulesSortOrderEnum
const (
	ListCccUpgradeSchedulesSortOrderAsc  ListCccUpgradeSchedulesSortOrderEnum = "ASC"
	ListCccUpgradeSchedulesSortOrderDesc ListCccUpgradeSchedulesSortOrderEnum = "DESC"
)

var mappingListCccUpgradeSchedulesSortOrderEnum = map[string]ListCccUpgradeSchedulesSortOrderEnum{
	"ASC":  ListCccUpgradeSchedulesSortOrderAsc,
	"DESC": ListCccUpgradeSchedulesSortOrderDesc,
}

var mappingListCccUpgradeSchedulesSortOrderEnumLowerCase = map[string]ListCccUpgradeSchedulesSortOrderEnum{
	"asc":  ListCccUpgradeSchedulesSortOrderAsc,
	"desc": ListCccUpgradeSchedulesSortOrderDesc,
}

// GetListCccUpgradeSchedulesSortOrderEnumValues Enumerates the set of values for ListCccUpgradeSchedulesSortOrderEnum
func GetListCccUpgradeSchedulesSortOrderEnumValues() []ListCccUpgradeSchedulesSortOrderEnum {
	values := make([]ListCccUpgradeSchedulesSortOrderEnum, 0)
	for _, v := range mappingListCccUpgradeSchedulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccUpgradeSchedulesSortOrderEnumStringValues Enumerates the set of values in String for ListCccUpgradeSchedulesSortOrderEnum
func GetListCccUpgradeSchedulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCccUpgradeSchedulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccUpgradeSchedulesSortOrderEnum(val string) (ListCccUpgradeSchedulesSortOrderEnum, bool) {
	enum, ok := mappingListCccUpgradeSchedulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCccUpgradeSchedulesSortByEnum Enum with underlying type: string
type ListCccUpgradeSchedulesSortByEnum string

// Set of constants representing the allowable values for ListCccUpgradeSchedulesSortByEnum
const (
	ListCccUpgradeSchedulesSortByTimecreated ListCccUpgradeSchedulesSortByEnum = "timeCreated"
	ListCccUpgradeSchedulesSortByDisplayname ListCccUpgradeSchedulesSortByEnum = "displayName"
)

var mappingListCccUpgradeSchedulesSortByEnum = map[string]ListCccUpgradeSchedulesSortByEnum{
	"timeCreated": ListCccUpgradeSchedulesSortByTimecreated,
	"displayName": ListCccUpgradeSchedulesSortByDisplayname,
}

var mappingListCccUpgradeSchedulesSortByEnumLowerCase = map[string]ListCccUpgradeSchedulesSortByEnum{
	"timecreated": ListCccUpgradeSchedulesSortByTimecreated,
	"displayname": ListCccUpgradeSchedulesSortByDisplayname,
}

// GetListCccUpgradeSchedulesSortByEnumValues Enumerates the set of values for ListCccUpgradeSchedulesSortByEnum
func GetListCccUpgradeSchedulesSortByEnumValues() []ListCccUpgradeSchedulesSortByEnum {
	values := make([]ListCccUpgradeSchedulesSortByEnum, 0)
	for _, v := range mappingListCccUpgradeSchedulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccUpgradeSchedulesSortByEnumStringValues Enumerates the set of values in String for ListCccUpgradeSchedulesSortByEnum
func GetListCccUpgradeSchedulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListCccUpgradeSchedulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccUpgradeSchedulesSortByEnum(val string) (ListCccUpgradeSchedulesSortByEnum, bool) {
	enum, ok := mappingListCccUpgradeSchedulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
