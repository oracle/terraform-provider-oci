// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListChannelsRequest wrapper for the ListChannels operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListChannels.go.html to see an example of how to use ListChannelsRequest.
type ListChannelsRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// Unique Channel identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// List only the information for Channels with this name. Channels names are unique and may not change.
	// Example: `MyChannel`
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// List only Channels with this category.
	Category ListChannelsCategoryEnum `mandatory:"false" contributesTo:"query" name:"category" omitEmpty:"true"`

	// List only Channels of this type.
	Type ListChannelsTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// List only the resources that are in this lifecycle state.
	LifecycleState ListChannelsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListChannelsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `timeCreated`.
	// The default sort order for `timeCreated` and `timeUpdated` is descending, and the default sort order for `name` is ascending.
	SortBy ListChannelsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListChannelsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListChannelsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListChannelsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListChannelsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListChannelsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListChannelsCategoryEnum(string(request.Category)); !ok && request.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", request.Category, strings.Join(GetListChannelsCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListChannelsTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListChannelsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListChannelsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListChannelsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListChannelsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListChannelsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListChannelsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListChannelsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListChannelsResponse wrapper for the ListChannels operation
type ListChannelsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ChannelCollection instances
	ChannelCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// When you are paging through a list, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the
	// `page` query parameter for the subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of results that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListChannelsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListChannelsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListChannelsCategoryEnum Enum with underlying type: string
type ListChannelsCategoryEnum string

// Set of constants representing the allowable values for ListChannelsCategoryEnum
const (
	ListChannelsCategoryAgent       ListChannelsCategoryEnum = "AGENT"
	ListChannelsCategoryApplication ListChannelsCategoryEnum = "APPLICATION"
	ListChannelsCategoryBot         ListChannelsCategoryEnum = "BOT"
	ListChannelsCategoryBotAsAgent  ListChannelsCategoryEnum = "BOT_AS_AGENT"
	ListChannelsCategorySystem      ListChannelsCategoryEnum = "SYSTEM"
	ListChannelsCategoryEvent       ListChannelsCategoryEnum = "EVENT"
)

var mappingListChannelsCategoryEnum = map[string]ListChannelsCategoryEnum{
	"AGENT":        ListChannelsCategoryAgent,
	"APPLICATION":  ListChannelsCategoryApplication,
	"BOT":          ListChannelsCategoryBot,
	"BOT_AS_AGENT": ListChannelsCategoryBotAsAgent,
	"SYSTEM":       ListChannelsCategorySystem,
	"EVENT":        ListChannelsCategoryEvent,
}

var mappingListChannelsCategoryEnumLowerCase = map[string]ListChannelsCategoryEnum{
	"agent":        ListChannelsCategoryAgent,
	"application":  ListChannelsCategoryApplication,
	"bot":          ListChannelsCategoryBot,
	"bot_as_agent": ListChannelsCategoryBotAsAgent,
	"system":       ListChannelsCategorySystem,
	"event":        ListChannelsCategoryEvent,
}

// GetListChannelsCategoryEnumValues Enumerates the set of values for ListChannelsCategoryEnum
func GetListChannelsCategoryEnumValues() []ListChannelsCategoryEnum {
	values := make([]ListChannelsCategoryEnum, 0)
	for _, v := range mappingListChannelsCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetListChannelsCategoryEnumStringValues Enumerates the set of values in String for ListChannelsCategoryEnum
func GetListChannelsCategoryEnumStringValues() []string {
	return []string{
		"AGENT",
		"APPLICATION",
		"BOT",
		"BOT_AS_AGENT",
		"SYSTEM",
		"EVENT",
	}
}

// GetMappingListChannelsCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChannelsCategoryEnum(val string) (ListChannelsCategoryEnum, bool) {
	enum, ok := mappingListChannelsCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListChannelsTypeEnum Enum with underlying type: string
type ListChannelsTypeEnum string

// Set of constants representing the allowable values for ListChannelsTypeEnum
const (
	ListChannelsTypeAndroid      ListChannelsTypeEnum = "ANDROID"
	ListChannelsTypeAppevent     ListChannelsTypeEnum = "APPEVENT"
	ListChannelsTypeApplication  ListChannelsTypeEnum = "APPLICATION"
	ListChannelsTypeCortana      ListChannelsTypeEnum = "CORTANA"
	ListChannelsTypeFacebook     ListChannelsTypeEnum = "FACEBOOK"
	ListChannelsTypeIos          ListChannelsTypeEnum = "IOS"
	ListChannelsTypeMsteams      ListChannelsTypeEnum = "MSTEAMS"
	ListChannelsTypeOss          ListChannelsTypeEnum = "OSS"
	ListChannelsTypeOsvc         ListChannelsTypeEnum = "OSVC"
	ListChannelsTypeServicecloud ListChannelsTypeEnum = "SERVICECLOUD"
	ListChannelsTypeSlack        ListChannelsTypeEnum = "SLACK"
	ListChannelsTypeTest         ListChannelsTypeEnum = "TEST"
	ListChannelsTypeTwilio       ListChannelsTypeEnum = "TWILIO"
	ListChannelsTypeWeb          ListChannelsTypeEnum = "WEB"
	ListChannelsTypeWebhook      ListChannelsTypeEnum = "WEBHOOK"
)

var mappingListChannelsTypeEnum = map[string]ListChannelsTypeEnum{
	"ANDROID":      ListChannelsTypeAndroid,
	"APPEVENT":     ListChannelsTypeAppevent,
	"APPLICATION":  ListChannelsTypeApplication,
	"CORTANA":      ListChannelsTypeCortana,
	"FACEBOOK":     ListChannelsTypeFacebook,
	"IOS":          ListChannelsTypeIos,
	"MSTEAMS":      ListChannelsTypeMsteams,
	"OSS":          ListChannelsTypeOss,
	"OSVC":         ListChannelsTypeOsvc,
	"SERVICECLOUD": ListChannelsTypeServicecloud,
	"SLACK":        ListChannelsTypeSlack,
	"TEST":         ListChannelsTypeTest,
	"TWILIO":       ListChannelsTypeTwilio,
	"WEB":          ListChannelsTypeWeb,
	"WEBHOOK":      ListChannelsTypeWebhook,
}

var mappingListChannelsTypeEnumLowerCase = map[string]ListChannelsTypeEnum{
	"android":      ListChannelsTypeAndroid,
	"appevent":     ListChannelsTypeAppevent,
	"application":  ListChannelsTypeApplication,
	"cortana":      ListChannelsTypeCortana,
	"facebook":     ListChannelsTypeFacebook,
	"ios":          ListChannelsTypeIos,
	"msteams":      ListChannelsTypeMsteams,
	"oss":          ListChannelsTypeOss,
	"osvc":         ListChannelsTypeOsvc,
	"servicecloud": ListChannelsTypeServicecloud,
	"slack":        ListChannelsTypeSlack,
	"test":         ListChannelsTypeTest,
	"twilio":       ListChannelsTypeTwilio,
	"web":          ListChannelsTypeWeb,
	"webhook":      ListChannelsTypeWebhook,
}

// GetListChannelsTypeEnumValues Enumerates the set of values for ListChannelsTypeEnum
func GetListChannelsTypeEnumValues() []ListChannelsTypeEnum {
	values := make([]ListChannelsTypeEnum, 0)
	for _, v := range mappingListChannelsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListChannelsTypeEnumStringValues Enumerates the set of values in String for ListChannelsTypeEnum
func GetListChannelsTypeEnumStringValues() []string {
	return []string{
		"ANDROID",
		"APPEVENT",
		"APPLICATION",
		"CORTANA",
		"FACEBOOK",
		"IOS",
		"MSTEAMS",
		"OSS",
		"OSVC",
		"SERVICECLOUD",
		"SLACK",
		"TEST",
		"TWILIO",
		"WEB",
		"WEBHOOK",
	}
}

// GetMappingListChannelsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChannelsTypeEnum(val string) (ListChannelsTypeEnum, bool) {
	enum, ok := mappingListChannelsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListChannelsLifecycleStateEnum Enum with underlying type: string
type ListChannelsLifecycleStateEnum string

// Set of constants representing the allowable values for ListChannelsLifecycleStateEnum
const (
	ListChannelsLifecycleStateCreating ListChannelsLifecycleStateEnum = "CREATING"
	ListChannelsLifecycleStateUpdating ListChannelsLifecycleStateEnum = "UPDATING"
	ListChannelsLifecycleStateActive   ListChannelsLifecycleStateEnum = "ACTIVE"
	ListChannelsLifecycleStateInactive ListChannelsLifecycleStateEnum = "INACTIVE"
	ListChannelsLifecycleStateDeleting ListChannelsLifecycleStateEnum = "DELETING"
	ListChannelsLifecycleStateDeleted  ListChannelsLifecycleStateEnum = "DELETED"
	ListChannelsLifecycleStateFailed   ListChannelsLifecycleStateEnum = "FAILED"
)

var mappingListChannelsLifecycleStateEnum = map[string]ListChannelsLifecycleStateEnum{
	"CREATING": ListChannelsLifecycleStateCreating,
	"UPDATING": ListChannelsLifecycleStateUpdating,
	"ACTIVE":   ListChannelsLifecycleStateActive,
	"INACTIVE": ListChannelsLifecycleStateInactive,
	"DELETING": ListChannelsLifecycleStateDeleting,
	"DELETED":  ListChannelsLifecycleStateDeleted,
	"FAILED":   ListChannelsLifecycleStateFailed,
}

var mappingListChannelsLifecycleStateEnumLowerCase = map[string]ListChannelsLifecycleStateEnum{
	"creating": ListChannelsLifecycleStateCreating,
	"updating": ListChannelsLifecycleStateUpdating,
	"active":   ListChannelsLifecycleStateActive,
	"inactive": ListChannelsLifecycleStateInactive,
	"deleting": ListChannelsLifecycleStateDeleting,
	"deleted":  ListChannelsLifecycleStateDeleted,
	"failed":   ListChannelsLifecycleStateFailed,
}

// GetListChannelsLifecycleStateEnumValues Enumerates the set of values for ListChannelsLifecycleStateEnum
func GetListChannelsLifecycleStateEnumValues() []ListChannelsLifecycleStateEnum {
	values := make([]ListChannelsLifecycleStateEnum, 0)
	for _, v := range mappingListChannelsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListChannelsLifecycleStateEnumStringValues Enumerates the set of values in String for ListChannelsLifecycleStateEnum
func GetListChannelsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListChannelsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChannelsLifecycleStateEnum(val string) (ListChannelsLifecycleStateEnum, bool) {
	enum, ok := mappingListChannelsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListChannelsSortOrderEnum Enum with underlying type: string
type ListChannelsSortOrderEnum string

// Set of constants representing the allowable values for ListChannelsSortOrderEnum
const (
	ListChannelsSortOrderAsc  ListChannelsSortOrderEnum = "ASC"
	ListChannelsSortOrderDesc ListChannelsSortOrderEnum = "DESC"
)

var mappingListChannelsSortOrderEnum = map[string]ListChannelsSortOrderEnum{
	"ASC":  ListChannelsSortOrderAsc,
	"DESC": ListChannelsSortOrderDesc,
}

var mappingListChannelsSortOrderEnumLowerCase = map[string]ListChannelsSortOrderEnum{
	"asc":  ListChannelsSortOrderAsc,
	"desc": ListChannelsSortOrderDesc,
}

// GetListChannelsSortOrderEnumValues Enumerates the set of values for ListChannelsSortOrderEnum
func GetListChannelsSortOrderEnumValues() []ListChannelsSortOrderEnum {
	values := make([]ListChannelsSortOrderEnum, 0)
	for _, v := range mappingListChannelsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListChannelsSortOrderEnumStringValues Enumerates the set of values in String for ListChannelsSortOrderEnum
func GetListChannelsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListChannelsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChannelsSortOrderEnum(val string) (ListChannelsSortOrderEnum, bool) {
	enum, ok := mappingListChannelsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListChannelsSortByEnum Enum with underlying type: string
type ListChannelsSortByEnum string

// Set of constants representing the allowable values for ListChannelsSortByEnum
const (
	ListChannelsSortByTimecreated ListChannelsSortByEnum = "timeCreated"
	ListChannelsSortByTimeupdated ListChannelsSortByEnum = "timeUpdated"
	ListChannelsSortByName        ListChannelsSortByEnum = "name"
)

var mappingListChannelsSortByEnum = map[string]ListChannelsSortByEnum{
	"timeCreated": ListChannelsSortByTimecreated,
	"timeUpdated": ListChannelsSortByTimeupdated,
	"name":        ListChannelsSortByName,
}

var mappingListChannelsSortByEnumLowerCase = map[string]ListChannelsSortByEnum{
	"timecreated": ListChannelsSortByTimecreated,
	"timeupdated": ListChannelsSortByTimeupdated,
	"name":        ListChannelsSortByName,
}

// GetListChannelsSortByEnumValues Enumerates the set of values for ListChannelsSortByEnum
func GetListChannelsSortByEnumValues() []ListChannelsSortByEnum {
	values := make([]ListChannelsSortByEnum, 0)
	for _, v := range mappingListChannelsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListChannelsSortByEnumStringValues Enumerates the set of values in String for ListChannelsSortByEnum
func GetListChannelsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"name",
	}
}

// GetMappingListChannelsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChannelsSortByEnum(val string) (ListChannelsSortByEnum, bool) {
	enum, ok := mappingListChannelsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
