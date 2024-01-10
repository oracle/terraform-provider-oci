// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package announcementsservice

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAnnouncementsRequest wrapper for the ListAnnouncements operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/ListAnnouncements.go.html to see an example of how to use ListAnnouncementsRequest.
type ListAnnouncementsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The type of announcement.
	AnnouncementType *string `mandatory:"false" contributesTo:"query" name:"announcementType"`

	// The announcement's current lifecycle state.
	LifecycleState ListAnnouncementsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Whether the announcement is displayed as a console banner.
	IsBanner *bool `mandatory:"false" contributesTo:"query" name:"isBanner"`

	// The criteria to sort by. You can specify only one sort order.
	SortBy ListAnnouncementsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use. (Sorting by `announcementType` orders the announcements list according to importance.)
	SortOrder ListAnnouncementsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The boundary for the earliest `timeOneValue` date on announcements that you want to see.
	TimeOneEarliestTime *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeOneEarliestTime"`

	// The boundary for the latest `timeOneValue` date on announcements that you want to see.
	TimeOneLatestTime *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeOneLatestTime"`

	// A filter to return only announcements that match a specific environment name.
	EnvironmentName *string `mandatory:"false" contributesTo:"query" name:"environmentName"`

	// A filter to return only announcements affecting a specific service.
	Service *string `mandatory:"false" contributesTo:"query" name:"service"`

	// A filter to return only announcements affecting a specific platform.
	PlatformType ListAnnouncementsPlatformTypeEnum `mandatory:"false" contributesTo:"query" name:"platformType" omitEmpty:"true"`

	// Exclude The type of announcement.
	ExcludeAnnouncementTypes []string `contributesTo:"query" name:"excludeAnnouncementTypes" collectionFormat:"multi"`

	// A filter to display only the latest announcement in a chain.
	ShouldShowOnlyLatestInChain *bool `mandatory:"false" contributesTo:"query" name:"shouldShowOnlyLatestInChain"`

	// A filter to return only announcements belonging to the specified announcement chain ID.
	ChainId *string `mandatory:"false" contributesTo:"query" name:"chainId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the complete request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAnnouncementsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAnnouncementsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAnnouncementsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAnnouncementsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAnnouncementsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAnnouncementsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAnnouncementsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAnnouncementsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAnnouncementsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAnnouncementsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAnnouncementsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAnnouncementsPlatformTypeEnum(string(request.PlatformType)); !ok && request.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", request.PlatformType, strings.Join(GetListAnnouncementsPlatformTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAnnouncementsResponse wrapper for the ListAnnouncements operation
type ListAnnouncementsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AnnouncementsCollection instances
	AnnouncementsCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAnnouncementsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAnnouncementsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAnnouncementsLifecycleStateEnum Enum with underlying type: string
type ListAnnouncementsLifecycleStateEnum string

// Set of constants representing the allowable values for ListAnnouncementsLifecycleStateEnum
const (
	ListAnnouncementsLifecycleStateActive   ListAnnouncementsLifecycleStateEnum = "ACTIVE"
	ListAnnouncementsLifecycleStateInactive ListAnnouncementsLifecycleStateEnum = "INACTIVE"
)

var mappingListAnnouncementsLifecycleStateEnum = map[string]ListAnnouncementsLifecycleStateEnum{
	"ACTIVE":   ListAnnouncementsLifecycleStateActive,
	"INACTIVE": ListAnnouncementsLifecycleStateInactive,
}

var mappingListAnnouncementsLifecycleStateEnumLowerCase = map[string]ListAnnouncementsLifecycleStateEnum{
	"active":   ListAnnouncementsLifecycleStateActive,
	"inactive": ListAnnouncementsLifecycleStateInactive,
}

// GetListAnnouncementsLifecycleStateEnumValues Enumerates the set of values for ListAnnouncementsLifecycleStateEnum
func GetListAnnouncementsLifecycleStateEnumValues() []ListAnnouncementsLifecycleStateEnum {
	values := make([]ListAnnouncementsLifecycleStateEnum, 0)
	for _, v := range mappingListAnnouncementsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnnouncementsLifecycleStateEnumStringValues Enumerates the set of values in String for ListAnnouncementsLifecycleStateEnum
func GetListAnnouncementsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingListAnnouncementsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnnouncementsLifecycleStateEnum(val string) (ListAnnouncementsLifecycleStateEnum, bool) {
	enum, ok := mappingListAnnouncementsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAnnouncementsSortByEnum Enum with underlying type: string
type ListAnnouncementsSortByEnum string

// Set of constants representing the allowable values for ListAnnouncementsSortByEnum
const (
	ListAnnouncementsSortByTimeonevalue          ListAnnouncementsSortByEnum = "timeOneValue"
	ListAnnouncementsSortByTimetwovalue          ListAnnouncementsSortByEnum = "timeTwoValue"
	ListAnnouncementsSortByTimecreated           ListAnnouncementsSortByEnum = "timeCreated"
	ListAnnouncementsSortByReferenceticketnumber ListAnnouncementsSortByEnum = "referenceTicketNumber"
	ListAnnouncementsSortBySummary               ListAnnouncementsSortByEnum = "summary"
	ListAnnouncementsSortByAnnouncementtype      ListAnnouncementsSortByEnum = "announcementType"
)

var mappingListAnnouncementsSortByEnum = map[string]ListAnnouncementsSortByEnum{
	"timeOneValue":          ListAnnouncementsSortByTimeonevalue,
	"timeTwoValue":          ListAnnouncementsSortByTimetwovalue,
	"timeCreated":           ListAnnouncementsSortByTimecreated,
	"referenceTicketNumber": ListAnnouncementsSortByReferenceticketnumber,
	"summary":               ListAnnouncementsSortBySummary,
	"announcementType":      ListAnnouncementsSortByAnnouncementtype,
}

var mappingListAnnouncementsSortByEnumLowerCase = map[string]ListAnnouncementsSortByEnum{
	"timeonevalue":          ListAnnouncementsSortByTimeonevalue,
	"timetwovalue":          ListAnnouncementsSortByTimetwovalue,
	"timecreated":           ListAnnouncementsSortByTimecreated,
	"referenceticketnumber": ListAnnouncementsSortByReferenceticketnumber,
	"summary":               ListAnnouncementsSortBySummary,
	"announcementtype":      ListAnnouncementsSortByAnnouncementtype,
}

// GetListAnnouncementsSortByEnumValues Enumerates the set of values for ListAnnouncementsSortByEnum
func GetListAnnouncementsSortByEnumValues() []ListAnnouncementsSortByEnum {
	values := make([]ListAnnouncementsSortByEnum, 0)
	for _, v := range mappingListAnnouncementsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnnouncementsSortByEnumStringValues Enumerates the set of values in String for ListAnnouncementsSortByEnum
func GetListAnnouncementsSortByEnumStringValues() []string {
	return []string{
		"timeOneValue",
		"timeTwoValue",
		"timeCreated",
		"referenceTicketNumber",
		"summary",
		"announcementType",
	}
}

// GetMappingListAnnouncementsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnnouncementsSortByEnum(val string) (ListAnnouncementsSortByEnum, bool) {
	enum, ok := mappingListAnnouncementsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAnnouncementsSortOrderEnum Enum with underlying type: string
type ListAnnouncementsSortOrderEnum string

// Set of constants representing the allowable values for ListAnnouncementsSortOrderEnum
const (
	ListAnnouncementsSortOrderAsc  ListAnnouncementsSortOrderEnum = "ASC"
	ListAnnouncementsSortOrderDesc ListAnnouncementsSortOrderEnum = "DESC"
)

var mappingListAnnouncementsSortOrderEnum = map[string]ListAnnouncementsSortOrderEnum{
	"ASC":  ListAnnouncementsSortOrderAsc,
	"DESC": ListAnnouncementsSortOrderDesc,
}

var mappingListAnnouncementsSortOrderEnumLowerCase = map[string]ListAnnouncementsSortOrderEnum{
	"asc":  ListAnnouncementsSortOrderAsc,
	"desc": ListAnnouncementsSortOrderDesc,
}

// GetListAnnouncementsSortOrderEnumValues Enumerates the set of values for ListAnnouncementsSortOrderEnum
func GetListAnnouncementsSortOrderEnumValues() []ListAnnouncementsSortOrderEnum {
	values := make([]ListAnnouncementsSortOrderEnum, 0)
	for _, v := range mappingListAnnouncementsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnnouncementsSortOrderEnumStringValues Enumerates the set of values in String for ListAnnouncementsSortOrderEnum
func GetListAnnouncementsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAnnouncementsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnnouncementsSortOrderEnum(val string) (ListAnnouncementsSortOrderEnum, bool) {
	enum, ok := mappingListAnnouncementsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAnnouncementsPlatformTypeEnum Enum with underlying type: string
type ListAnnouncementsPlatformTypeEnum string

// Set of constants representing the allowable values for ListAnnouncementsPlatformTypeEnum
const (
	ListAnnouncementsPlatformTypeIaas ListAnnouncementsPlatformTypeEnum = "IAAS"
	ListAnnouncementsPlatformTypeSaas ListAnnouncementsPlatformTypeEnum = "SAAS"
)

var mappingListAnnouncementsPlatformTypeEnum = map[string]ListAnnouncementsPlatformTypeEnum{
	"IAAS": ListAnnouncementsPlatformTypeIaas,
	"SAAS": ListAnnouncementsPlatformTypeSaas,
}

var mappingListAnnouncementsPlatformTypeEnumLowerCase = map[string]ListAnnouncementsPlatformTypeEnum{
	"iaas": ListAnnouncementsPlatformTypeIaas,
	"saas": ListAnnouncementsPlatformTypeSaas,
}

// GetListAnnouncementsPlatformTypeEnumValues Enumerates the set of values for ListAnnouncementsPlatformTypeEnum
func GetListAnnouncementsPlatformTypeEnumValues() []ListAnnouncementsPlatformTypeEnum {
	values := make([]ListAnnouncementsPlatformTypeEnum, 0)
	for _, v := range mappingListAnnouncementsPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnnouncementsPlatformTypeEnumStringValues Enumerates the set of values in String for ListAnnouncementsPlatformTypeEnum
func GetListAnnouncementsPlatformTypeEnumStringValues() []string {
	return []string{
		"IAAS",
		"SAAS",
	}
}

// GetMappingListAnnouncementsPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnnouncementsPlatformTypeEnum(val string) (ListAnnouncementsPlatformTypeEnum, bool) {
	enum, ok := mappingListAnnouncementsPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
