// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jmsjavadownloads

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListJavaDownloadTokensRequest wrapper for the ListJavaDownloadTokens operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/ListJavaDownloadTokens.go.html to see an example of how to use ListJavaDownloadTokensRequest.
type ListJavaDownloadTokensRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the tenancy.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListJavaDownloadTokensLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique JavaDownloadToken identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Unique JavaDownloadToken value.
	Value *string `mandatory:"false" contributesTo:"query" name:"value"`

	// Unique Java family version identifier.
	FamilyVersion *string `mandatory:"false" contributesTo:"query" name:"familyVersion"`

	// A filter to return only resources that match the user principal detail.
	// The search string can be any of the property values from the Principal object.
	// This object is used as response datatype for the `createdBy` and `lastUpdatedBy` fields in applicable resource.
	SearchByUser *string `mandatory:"false" contributesTo:"query" name:"searchByUser"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListJavaDownloadTokensSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. If no value is specified, _timeCreated_ is the default.
	SortBy ListJavaDownloadTokensSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJavaDownloadTokensRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJavaDownloadTokensRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJavaDownloadTokensRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJavaDownloadTokensRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJavaDownloadTokensRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJavaDownloadTokensLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListJavaDownloadTokensLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaDownloadTokensSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJavaDownloadTokensSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaDownloadTokensSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJavaDownloadTokensSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJavaDownloadTokensResponse wrapper for the ListJavaDownloadTokens operation
type ListJavaDownloadTokensResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JavaDownloadTokenCollection instances
	JavaDownloadTokenCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJavaDownloadTokensResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJavaDownloadTokensResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJavaDownloadTokensLifecycleStateEnum Enum with underlying type: string
type ListJavaDownloadTokensLifecycleStateEnum string

// Set of constants representing the allowable values for ListJavaDownloadTokensLifecycleStateEnum
const (
	ListJavaDownloadTokensLifecycleStateActive         ListJavaDownloadTokensLifecycleStateEnum = "ACTIVE"
	ListJavaDownloadTokensLifecycleStateCreating       ListJavaDownloadTokensLifecycleStateEnum = "CREATING"
	ListJavaDownloadTokensLifecycleStateDeleted        ListJavaDownloadTokensLifecycleStateEnum = "DELETED"
	ListJavaDownloadTokensLifecycleStateDeleting       ListJavaDownloadTokensLifecycleStateEnum = "DELETING"
	ListJavaDownloadTokensLifecycleStateFailed         ListJavaDownloadTokensLifecycleStateEnum = "FAILED"
	ListJavaDownloadTokensLifecycleStateNeedsAttention ListJavaDownloadTokensLifecycleStateEnum = "NEEDS_ATTENTION"
	ListJavaDownloadTokensLifecycleStateUpdating       ListJavaDownloadTokensLifecycleStateEnum = "UPDATING"
)

var mappingListJavaDownloadTokensLifecycleStateEnum = map[string]ListJavaDownloadTokensLifecycleStateEnum{
	"ACTIVE":          ListJavaDownloadTokensLifecycleStateActive,
	"CREATING":        ListJavaDownloadTokensLifecycleStateCreating,
	"DELETED":         ListJavaDownloadTokensLifecycleStateDeleted,
	"DELETING":        ListJavaDownloadTokensLifecycleStateDeleting,
	"FAILED":          ListJavaDownloadTokensLifecycleStateFailed,
	"NEEDS_ATTENTION": ListJavaDownloadTokensLifecycleStateNeedsAttention,
	"UPDATING":        ListJavaDownloadTokensLifecycleStateUpdating,
}

var mappingListJavaDownloadTokensLifecycleStateEnumLowerCase = map[string]ListJavaDownloadTokensLifecycleStateEnum{
	"active":          ListJavaDownloadTokensLifecycleStateActive,
	"creating":        ListJavaDownloadTokensLifecycleStateCreating,
	"deleted":         ListJavaDownloadTokensLifecycleStateDeleted,
	"deleting":        ListJavaDownloadTokensLifecycleStateDeleting,
	"failed":          ListJavaDownloadTokensLifecycleStateFailed,
	"needs_attention": ListJavaDownloadTokensLifecycleStateNeedsAttention,
	"updating":        ListJavaDownloadTokensLifecycleStateUpdating,
}

// GetListJavaDownloadTokensLifecycleStateEnumValues Enumerates the set of values for ListJavaDownloadTokensLifecycleStateEnum
func GetListJavaDownloadTokensLifecycleStateEnumValues() []ListJavaDownloadTokensLifecycleStateEnum {
	values := make([]ListJavaDownloadTokensLifecycleStateEnum, 0)
	for _, v := range mappingListJavaDownloadTokensLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaDownloadTokensLifecycleStateEnumStringValues Enumerates the set of values in String for ListJavaDownloadTokensLifecycleStateEnum
func GetListJavaDownloadTokensLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"NEEDS_ATTENTION",
		"UPDATING",
	}
}

// GetMappingListJavaDownloadTokensLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaDownloadTokensLifecycleStateEnum(val string) (ListJavaDownloadTokensLifecycleStateEnum, bool) {
	enum, ok := mappingListJavaDownloadTokensLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaDownloadTokensSortOrderEnum Enum with underlying type: string
type ListJavaDownloadTokensSortOrderEnum string

// Set of constants representing the allowable values for ListJavaDownloadTokensSortOrderEnum
const (
	ListJavaDownloadTokensSortOrderAsc  ListJavaDownloadTokensSortOrderEnum = "ASC"
	ListJavaDownloadTokensSortOrderDesc ListJavaDownloadTokensSortOrderEnum = "DESC"
)

var mappingListJavaDownloadTokensSortOrderEnum = map[string]ListJavaDownloadTokensSortOrderEnum{
	"ASC":  ListJavaDownloadTokensSortOrderAsc,
	"DESC": ListJavaDownloadTokensSortOrderDesc,
}

var mappingListJavaDownloadTokensSortOrderEnumLowerCase = map[string]ListJavaDownloadTokensSortOrderEnum{
	"asc":  ListJavaDownloadTokensSortOrderAsc,
	"desc": ListJavaDownloadTokensSortOrderDesc,
}

// GetListJavaDownloadTokensSortOrderEnumValues Enumerates the set of values for ListJavaDownloadTokensSortOrderEnum
func GetListJavaDownloadTokensSortOrderEnumValues() []ListJavaDownloadTokensSortOrderEnum {
	values := make([]ListJavaDownloadTokensSortOrderEnum, 0)
	for _, v := range mappingListJavaDownloadTokensSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaDownloadTokensSortOrderEnumStringValues Enumerates the set of values in String for ListJavaDownloadTokensSortOrderEnum
func GetListJavaDownloadTokensSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJavaDownloadTokensSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaDownloadTokensSortOrderEnum(val string) (ListJavaDownloadTokensSortOrderEnum, bool) {
	enum, ok := mappingListJavaDownloadTokensSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaDownloadTokensSortByEnum Enum with underlying type: string
type ListJavaDownloadTokensSortByEnum string

// Set of constants representing the allowable values for ListJavaDownloadTokensSortByEnum
const (
	ListJavaDownloadTokensSortByTimecreated ListJavaDownloadTokensSortByEnum = "timeCreated"
	ListJavaDownloadTokensSortByTimeexpires ListJavaDownloadTokensSortByEnum = "timeExpires"
	ListJavaDownloadTokensSortByState       ListJavaDownloadTokensSortByEnum = "state"
	ListJavaDownloadTokensSortByDisplayname ListJavaDownloadTokensSortByEnum = "displayName"
	ListJavaDownloadTokensSortByJavaversion ListJavaDownloadTokensSortByEnum = "javaVersion"
)

var mappingListJavaDownloadTokensSortByEnum = map[string]ListJavaDownloadTokensSortByEnum{
	"timeCreated": ListJavaDownloadTokensSortByTimecreated,
	"timeExpires": ListJavaDownloadTokensSortByTimeexpires,
	"state":       ListJavaDownloadTokensSortByState,
	"displayName": ListJavaDownloadTokensSortByDisplayname,
	"javaVersion": ListJavaDownloadTokensSortByJavaversion,
}

var mappingListJavaDownloadTokensSortByEnumLowerCase = map[string]ListJavaDownloadTokensSortByEnum{
	"timecreated": ListJavaDownloadTokensSortByTimecreated,
	"timeexpires": ListJavaDownloadTokensSortByTimeexpires,
	"state":       ListJavaDownloadTokensSortByState,
	"displayname": ListJavaDownloadTokensSortByDisplayname,
	"javaversion": ListJavaDownloadTokensSortByJavaversion,
}

// GetListJavaDownloadTokensSortByEnumValues Enumerates the set of values for ListJavaDownloadTokensSortByEnum
func GetListJavaDownloadTokensSortByEnumValues() []ListJavaDownloadTokensSortByEnum {
	values := make([]ListJavaDownloadTokensSortByEnum, 0)
	for _, v := range mappingListJavaDownloadTokensSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaDownloadTokensSortByEnumStringValues Enumerates the set of values in String for ListJavaDownloadTokensSortByEnum
func GetListJavaDownloadTokensSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeExpires",
		"state",
		"displayName",
		"javaVersion",
	}
}

// GetMappingListJavaDownloadTokensSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaDownloadTokensSortByEnum(val string) (ListJavaDownloadTokensSortByEnum, bool) {
	enum, ok := mappingListJavaDownloadTokensSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
