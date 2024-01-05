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

// ListSkillsRequest wrapper for the ListSkills operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListSkills.go.html to see an example of how to use ListSkillsRequest.
type ListSkillsRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// Unique Skill identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// List only Bot resources with this category.
	Category *string `mandatory:"false" contributesTo:"query" name:"category"`

	// List only Bot resources with this name. Names are unique and may not change.
	// Example: `MySkill`
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// List only Bot resources with this version. Versions are unique and may not change.
	// Example: `1.0`
	Version *string `mandatory:"false" contributesTo:"query" name:"version"`

	// List only Bot resources with this namespace. Namespaces may not change.
	// Example: `MyNamespace`
	Namespace *string `mandatory:"false" contributesTo:"query" name:"namespace"`

	// List only Bot resources with this platform version.
	PlatformVersion *string `mandatory:"false" contributesTo:"query" name:"platformVersion"`

	// List only the resources that are in this lifecycle state.
	LifecycleState ListSkillsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// List only Bot resources with this lifecycle details.
	LifecycleDetails *string `mandatory:"false" contributesTo:"query" name:"lifecycleDetails"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListSkillsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `timeCreated`.
	// The default sort order for `timeCreated` and `timeUpdated` is descending.
	// For all other sort fields the default sort order is ascending.
	SortBy ListSkillsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSkillsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSkillsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSkillsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSkillsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSkillsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSkillsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSkillsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSkillsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSkillsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSkillsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSkillsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSkillsResponse wrapper for the ListSkills operation
type ListSkillsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SkillCollection instances
	SkillCollection `presentIn:"body"`

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

func (response ListSkillsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSkillsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSkillsLifecycleStateEnum Enum with underlying type: string
type ListSkillsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSkillsLifecycleStateEnum
const (
	ListSkillsLifecycleStateCreating ListSkillsLifecycleStateEnum = "CREATING"
	ListSkillsLifecycleStateUpdating ListSkillsLifecycleStateEnum = "UPDATING"
	ListSkillsLifecycleStateActive   ListSkillsLifecycleStateEnum = "ACTIVE"
	ListSkillsLifecycleStateInactive ListSkillsLifecycleStateEnum = "INACTIVE"
	ListSkillsLifecycleStateDeleting ListSkillsLifecycleStateEnum = "DELETING"
	ListSkillsLifecycleStateDeleted  ListSkillsLifecycleStateEnum = "DELETED"
	ListSkillsLifecycleStateFailed   ListSkillsLifecycleStateEnum = "FAILED"
)

var mappingListSkillsLifecycleStateEnum = map[string]ListSkillsLifecycleStateEnum{
	"CREATING": ListSkillsLifecycleStateCreating,
	"UPDATING": ListSkillsLifecycleStateUpdating,
	"ACTIVE":   ListSkillsLifecycleStateActive,
	"INACTIVE": ListSkillsLifecycleStateInactive,
	"DELETING": ListSkillsLifecycleStateDeleting,
	"DELETED":  ListSkillsLifecycleStateDeleted,
	"FAILED":   ListSkillsLifecycleStateFailed,
}

var mappingListSkillsLifecycleStateEnumLowerCase = map[string]ListSkillsLifecycleStateEnum{
	"creating": ListSkillsLifecycleStateCreating,
	"updating": ListSkillsLifecycleStateUpdating,
	"active":   ListSkillsLifecycleStateActive,
	"inactive": ListSkillsLifecycleStateInactive,
	"deleting": ListSkillsLifecycleStateDeleting,
	"deleted":  ListSkillsLifecycleStateDeleted,
	"failed":   ListSkillsLifecycleStateFailed,
}

// GetListSkillsLifecycleStateEnumValues Enumerates the set of values for ListSkillsLifecycleStateEnum
func GetListSkillsLifecycleStateEnumValues() []ListSkillsLifecycleStateEnum {
	values := make([]ListSkillsLifecycleStateEnum, 0)
	for _, v := range mappingListSkillsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSkillsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSkillsLifecycleStateEnum
func GetListSkillsLifecycleStateEnumStringValues() []string {
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

// GetMappingListSkillsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSkillsLifecycleStateEnum(val string) (ListSkillsLifecycleStateEnum, bool) {
	enum, ok := mappingListSkillsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSkillsSortOrderEnum Enum with underlying type: string
type ListSkillsSortOrderEnum string

// Set of constants representing the allowable values for ListSkillsSortOrderEnum
const (
	ListSkillsSortOrderAsc  ListSkillsSortOrderEnum = "ASC"
	ListSkillsSortOrderDesc ListSkillsSortOrderEnum = "DESC"
)

var mappingListSkillsSortOrderEnum = map[string]ListSkillsSortOrderEnum{
	"ASC":  ListSkillsSortOrderAsc,
	"DESC": ListSkillsSortOrderDesc,
}

var mappingListSkillsSortOrderEnumLowerCase = map[string]ListSkillsSortOrderEnum{
	"asc":  ListSkillsSortOrderAsc,
	"desc": ListSkillsSortOrderDesc,
}

// GetListSkillsSortOrderEnumValues Enumerates the set of values for ListSkillsSortOrderEnum
func GetListSkillsSortOrderEnumValues() []ListSkillsSortOrderEnum {
	values := make([]ListSkillsSortOrderEnum, 0)
	for _, v := range mappingListSkillsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSkillsSortOrderEnumStringValues Enumerates the set of values in String for ListSkillsSortOrderEnum
func GetListSkillsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSkillsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSkillsSortOrderEnum(val string) (ListSkillsSortOrderEnum, bool) {
	enum, ok := mappingListSkillsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSkillsSortByEnum Enum with underlying type: string
type ListSkillsSortByEnum string

// Set of constants representing the allowable values for ListSkillsSortByEnum
const (
	ListSkillsSortByTimecreated ListSkillsSortByEnum = "timeCreated"
	ListSkillsSortByTimeupdated ListSkillsSortByEnum = "timeUpdated"
	ListSkillsSortByName        ListSkillsSortByEnum = "name"
)

var mappingListSkillsSortByEnum = map[string]ListSkillsSortByEnum{
	"timeCreated": ListSkillsSortByTimecreated,
	"timeUpdated": ListSkillsSortByTimeupdated,
	"name":        ListSkillsSortByName,
}

var mappingListSkillsSortByEnumLowerCase = map[string]ListSkillsSortByEnum{
	"timecreated": ListSkillsSortByTimecreated,
	"timeupdated": ListSkillsSortByTimeupdated,
	"name":        ListSkillsSortByName,
}

// GetListSkillsSortByEnumValues Enumerates the set of values for ListSkillsSortByEnum
func GetListSkillsSortByEnumValues() []ListSkillsSortByEnum {
	values := make([]ListSkillsSortByEnum, 0)
	for _, v := range mappingListSkillsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSkillsSortByEnumStringValues Enumerates the set of values in String for ListSkillsSortByEnum
func GetListSkillsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"name",
	}
}

// GetMappingListSkillsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSkillsSortByEnum(val string) (ListSkillsSortByEnum, bool) {
	enum, ok := mappingListSkillsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
