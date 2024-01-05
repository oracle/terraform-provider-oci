// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSensitiveDataModelsRequest wrapper for the ListSensitiveDataModels operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveDataModels.go.html to see an example of how to use ListSensitiveDataModelsRequest.
type ListSensitiveDataModelsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSensitiveDataModelsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only the resources that match the specified sensitive data model OCID.
	SensitiveDataModelId *string `mandatory:"false" contributesTo:"query" name:"sensitiveDataModelId"`

	// A filter to return only the resources that were created after the specified date and time, as defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// Search for resources that were created before a specific date.
	// Specifying this parameter corresponding `timeCreatedLessThan`
	// parameter will retrieve all resources created before the
	// specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as
	// defined by RFC 3339.
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSensitiveDataModelsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for timeCreated is descending.
	// The default order for displayName is ascending.
	SortBy ListSensitiveDataModelsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only the resources that match the specified lifecycle state.
	LifecycleState ListSensitiveDataModelsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSensitiveDataModelsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSensitiveDataModelsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSensitiveDataModelsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSensitiveDataModelsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSensitiveDataModelsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSensitiveDataModelsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSensitiveDataModelsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveDataModelsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSensitiveDataModelsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveDataModelsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSensitiveDataModelsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveDataModelsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSensitiveDataModelsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSensitiveDataModelsResponse wrapper for the ListSensitiveDataModels operation
type ListSensitiveDataModelsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SensitiveDataModelCollection instances
	SensitiveDataModelCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSensitiveDataModelsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSensitiveDataModelsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSensitiveDataModelsAccessLevelEnum Enum with underlying type: string
type ListSensitiveDataModelsAccessLevelEnum string

// Set of constants representing the allowable values for ListSensitiveDataModelsAccessLevelEnum
const (
	ListSensitiveDataModelsAccessLevelRestricted ListSensitiveDataModelsAccessLevelEnum = "RESTRICTED"
	ListSensitiveDataModelsAccessLevelAccessible ListSensitiveDataModelsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSensitiveDataModelsAccessLevelEnum = map[string]ListSensitiveDataModelsAccessLevelEnum{
	"RESTRICTED": ListSensitiveDataModelsAccessLevelRestricted,
	"ACCESSIBLE": ListSensitiveDataModelsAccessLevelAccessible,
}

var mappingListSensitiveDataModelsAccessLevelEnumLowerCase = map[string]ListSensitiveDataModelsAccessLevelEnum{
	"restricted": ListSensitiveDataModelsAccessLevelRestricted,
	"accessible": ListSensitiveDataModelsAccessLevelAccessible,
}

// GetListSensitiveDataModelsAccessLevelEnumValues Enumerates the set of values for ListSensitiveDataModelsAccessLevelEnum
func GetListSensitiveDataModelsAccessLevelEnumValues() []ListSensitiveDataModelsAccessLevelEnum {
	values := make([]ListSensitiveDataModelsAccessLevelEnum, 0)
	for _, v := range mappingListSensitiveDataModelsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveDataModelsAccessLevelEnumStringValues Enumerates the set of values in String for ListSensitiveDataModelsAccessLevelEnum
func GetListSensitiveDataModelsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSensitiveDataModelsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveDataModelsAccessLevelEnum(val string) (ListSensitiveDataModelsAccessLevelEnum, bool) {
	enum, ok := mappingListSensitiveDataModelsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveDataModelsSortOrderEnum Enum with underlying type: string
type ListSensitiveDataModelsSortOrderEnum string

// Set of constants representing the allowable values for ListSensitiveDataModelsSortOrderEnum
const (
	ListSensitiveDataModelsSortOrderAsc  ListSensitiveDataModelsSortOrderEnum = "ASC"
	ListSensitiveDataModelsSortOrderDesc ListSensitiveDataModelsSortOrderEnum = "DESC"
)

var mappingListSensitiveDataModelsSortOrderEnum = map[string]ListSensitiveDataModelsSortOrderEnum{
	"ASC":  ListSensitiveDataModelsSortOrderAsc,
	"DESC": ListSensitiveDataModelsSortOrderDesc,
}

var mappingListSensitiveDataModelsSortOrderEnumLowerCase = map[string]ListSensitiveDataModelsSortOrderEnum{
	"asc":  ListSensitiveDataModelsSortOrderAsc,
	"desc": ListSensitiveDataModelsSortOrderDesc,
}

// GetListSensitiveDataModelsSortOrderEnumValues Enumerates the set of values for ListSensitiveDataModelsSortOrderEnum
func GetListSensitiveDataModelsSortOrderEnumValues() []ListSensitiveDataModelsSortOrderEnum {
	values := make([]ListSensitiveDataModelsSortOrderEnum, 0)
	for _, v := range mappingListSensitiveDataModelsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveDataModelsSortOrderEnumStringValues Enumerates the set of values in String for ListSensitiveDataModelsSortOrderEnum
func GetListSensitiveDataModelsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSensitiveDataModelsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveDataModelsSortOrderEnum(val string) (ListSensitiveDataModelsSortOrderEnum, bool) {
	enum, ok := mappingListSensitiveDataModelsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveDataModelsSortByEnum Enum with underlying type: string
type ListSensitiveDataModelsSortByEnum string

// Set of constants representing the allowable values for ListSensitiveDataModelsSortByEnum
const (
	ListSensitiveDataModelsSortByTimecreated ListSensitiveDataModelsSortByEnum = "timeCreated"
	ListSensitiveDataModelsSortByDisplayname ListSensitiveDataModelsSortByEnum = "displayName"
)

var mappingListSensitiveDataModelsSortByEnum = map[string]ListSensitiveDataModelsSortByEnum{
	"timeCreated": ListSensitiveDataModelsSortByTimecreated,
	"displayName": ListSensitiveDataModelsSortByDisplayname,
}

var mappingListSensitiveDataModelsSortByEnumLowerCase = map[string]ListSensitiveDataModelsSortByEnum{
	"timecreated": ListSensitiveDataModelsSortByTimecreated,
	"displayname": ListSensitiveDataModelsSortByDisplayname,
}

// GetListSensitiveDataModelsSortByEnumValues Enumerates the set of values for ListSensitiveDataModelsSortByEnum
func GetListSensitiveDataModelsSortByEnumValues() []ListSensitiveDataModelsSortByEnum {
	values := make([]ListSensitiveDataModelsSortByEnum, 0)
	for _, v := range mappingListSensitiveDataModelsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveDataModelsSortByEnumStringValues Enumerates the set of values in String for ListSensitiveDataModelsSortByEnum
func GetListSensitiveDataModelsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSensitiveDataModelsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveDataModelsSortByEnum(val string) (ListSensitiveDataModelsSortByEnum, bool) {
	enum, ok := mappingListSensitiveDataModelsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveDataModelsLifecycleStateEnum Enum with underlying type: string
type ListSensitiveDataModelsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSensitiveDataModelsLifecycleStateEnum
const (
	ListSensitiveDataModelsLifecycleStateCreating ListSensitiveDataModelsLifecycleStateEnum = "CREATING"
	ListSensitiveDataModelsLifecycleStateActive   ListSensitiveDataModelsLifecycleStateEnum = "ACTIVE"
	ListSensitiveDataModelsLifecycleStateUpdating ListSensitiveDataModelsLifecycleStateEnum = "UPDATING"
	ListSensitiveDataModelsLifecycleStateDeleting ListSensitiveDataModelsLifecycleStateEnum = "DELETING"
	ListSensitiveDataModelsLifecycleStateDeleted  ListSensitiveDataModelsLifecycleStateEnum = "DELETED"
	ListSensitiveDataModelsLifecycleStateFailed   ListSensitiveDataModelsLifecycleStateEnum = "FAILED"
)

var mappingListSensitiveDataModelsLifecycleStateEnum = map[string]ListSensitiveDataModelsLifecycleStateEnum{
	"CREATING": ListSensitiveDataModelsLifecycleStateCreating,
	"ACTIVE":   ListSensitiveDataModelsLifecycleStateActive,
	"UPDATING": ListSensitiveDataModelsLifecycleStateUpdating,
	"DELETING": ListSensitiveDataModelsLifecycleStateDeleting,
	"DELETED":  ListSensitiveDataModelsLifecycleStateDeleted,
	"FAILED":   ListSensitiveDataModelsLifecycleStateFailed,
}

var mappingListSensitiveDataModelsLifecycleStateEnumLowerCase = map[string]ListSensitiveDataModelsLifecycleStateEnum{
	"creating": ListSensitiveDataModelsLifecycleStateCreating,
	"active":   ListSensitiveDataModelsLifecycleStateActive,
	"updating": ListSensitiveDataModelsLifecycleStateUpdating,
	"deleting": ListSensitiveDataModelsLifecycleStateDeleting,
	"deleted":  ListSensitiveDataModelsLifecycleStateDeleted,
	"failed":   ListSensitiveDataModelsLifecycleStateFailed,
}

// GetListSensitiveDataModelsLifecycleStateEnumValues Enumerates the set of values for ListSensitiveDataModelsLifecycleStateEnum
func GetListSensitiveDataModelsLifecycleStateEnumValues() []ListSensitiveDataModelsLifecycleStateEnum {
	values := make([]ListSensitiveDataModelsLifecycleStateEnum, 0)
	for _, v := range mappingListSensitiveDataModelsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveDataModelsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSensitiveDataModelsLifecycleStateEnum
func GetListSensitiveDataModelsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListSensitiveDataModelsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveDataModelsLifecycleStateEnum(val string) (ListSensitiveDataModelsLifecycleStateEnum, bool) {
	enum, ok := mappingListSensitiveDataModelsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
