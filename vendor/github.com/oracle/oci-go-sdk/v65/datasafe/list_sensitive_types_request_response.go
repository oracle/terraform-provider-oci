// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSensitiveTypesRequest wrapper for the ListSensitiveTypes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveTypes.go.html to see an example of how to use ListSensitiveTypesRequest.
type ListSensitiveTypesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSensitiveTypesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only items related to a specific sensitive type OCID.
	SensitiveTypeId *string `mandatory:"false" contributesTo:"query" name:"sensitiveTypeId"`

	// A filter to return the sensitive type resources based on the value of their source attribute.
	SensitiveTypeSource ListSensitiveTypesSensitiveTypeSourceEnum `mandatory:"false" contributesTo:"query" name:"sensitiveTypeSource" omitEmpty:"true"`

	// A filter to return the sensitive type resources based on the value of their entityType attribute.
	EntityType ListSensitiveTypesEntityTypeEnum `mandatory:"false" contributesTo:"query" name:"entityType" omitEmpty:"true"`

	// A filter to return only the sensitive types that are children of the sensitive category identified by the specified OCID.
	ParentCategoryId *string `mandatory:"false" contributesTo:"query" name:"parentCategoryId"`

	// A filter to return only the sensitive types that have the default masking format identified by the specified OCID.
	DefaultMaskingFormatId *string `mandatory:"false" contributesTo:"query" name:"defaultMaskingFormatId"`

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

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSensitiveTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for timeCreated is descending.
	// The default order for displayName is ascending.
	SortBy ListSensitiveTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only the resources that match the specified lifecycle state.
	LifecycleState ListSensitiveTypesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSensitiveTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSensitiveTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSensitiveTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSensitiveTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSensitiveTypesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSensitiveTypesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSensitiveTypesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveTypesSensitiveTypeSourceEnum(string(request.SensitiveTypeSource)); !ok && request.SensitiveTypeSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SensitiveTypeSource: %s. Supported values are: %s.", request.SensitiveTypeSource, strings.Join(GetListSensitiveTypesSensitiveTypeSourceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveTypesEntityTypeEnum(string(request.EntityType)); !ok && request.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", request.EntityType, strings.Join(GetListSensitiveTypesEntityTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveTypesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSensitiveTypesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveTypesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSensitiveTypesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveTypesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSensitiveTypesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSensitiveTypesResponse wrapper for the ListSensitiveTypes operation
type ListSensitiveTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SensitiveTypeCollection instances
	SensitiveTypeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSensitiveTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSensitiveTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSensitiveTypesAccessLevelEnum Enum with underlying type: string
type ListSensitiveTypesAccessLevelEnum string

// Set of constants representing the allowable values for ListSensitiveTypesAccessLevelEnum
const (
	ListSensitiveTypesAccessLevelRestricted ListSensitiveTypesAccessLevelEnum = "RESTRICTED"
	ListSensitiveTypesAccessLevelAccessible ListSensitiveTypesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSensitiveTypesAccessLevelEnum = map[string]ListSensitiveTypesAccessLevelEnum{
	"RESTRICTED": ListSensitiveTypesAccessLevelRestricted,
	"ACCESSIBLE": ListSensitiveTypesAccessLevelAccessible,
}

var mappingListSensitiveTypesAccessLevelEnumLowerCase = map[string]ListSensitiveTypesAccessLevelEnum{
	"restricted": ListSensitiveTypesAccessLevelRestricted,
	"accessible": ListSensitiveTypesAccessLevelAccessible,
}

// GetListSensitiveTypesAccessLevelEnumValues Enumerates the set of values for ListSensitiveTypesAccessLevelEnum
func GetListSensitiveTypesAccessLevelEnumValues() []ListSensitiveTypesAccessLevelEnum {
	values := make([]ListSensitiveTypesAccessLevelEnum, 0)
	for _, v := range mappingListSensitiveTypesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveTypesAccessLevelEnumStringValues Enumerates the set of values in String for ListSensitiveTypesAccessLevelEnum
func GetListSensitiveTypesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSensitiveTypesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveTypesAccessLevelEnum(val string) (ListSensitiveTypesAccessLevelEnum, bool) {
	enum, ok := mappingListSensitiveTypesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveTypesSensitiveTypeSourceEnum Enum with underlying type: string
type ListSensitiveTypesSensitiveTypeSourceEnum string

// Set of constants representing the allowable values for ListSensitiveTypesSensitiveTypeSourceEnum
const (
	ListSensitiveTypesSensitiveTypeSourceOracle ListSensitiveTypesSensitiveTypeSourceEnum = "ORACLE"
	ListSensitiveTypesSensitiveTypeSourceUser   ListSensitiveTypesSensitiveTypeSourceEnum = "USER"
)

var mappingListSensitiveTypesSensitiveTypeSourceEnum = map[string]ListSensitiveTypesSensitiveTypeSourceEnum{
	"ORACLE": ListSensitiveTypesSensitiveTypeSourceOracle,
	"USER":   ListSensitiveTypesSensitiveTypeSourceUser,
}

var mappingListSensitiveTypesSensitiveTypeSourceEnumLowerCase = map[string]ListSensitiveTypesSensitiveTypeSourceEnum{
	"oracle": ListSensitiveTypesSensitiveTypeSourceOracle,
	"user":   ListSensitiveTypesSensitiveTypeSourceUser,
}

// GetListSensitiveTypesSensitiveTypeSourceEnumValues Enumerates the set of values for ListSensitiveTypesSensitiveTypeSourceEnum
func GetListSensitiveTypesSensitiveTypeSourceEnumValues() []ListSensitiveTypesSensitiveTypeSourceEnum {
	values := make([]ListSensitiveTypesSensitiveTypeSourceEnum, 0)
	for _, v := range mappingListSensitiveTypesSensitiveTypeSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveTypesSensitiveTypeSourceEnumStringValues Enumerates the set of values in String for ListSensitiveTypesSensitiveTypeSourceEnum
func GetListSensitiveTypesSensitiveTypeSourceEnumStringValues() []string {
	return []string{
		"ORACLE",
		"USER",
	}
}

// GetMappingListSensitiveTypesSensitiveTypeSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveTypesSensitiveTypeSourceEnum(val string) (ListSensitiveTypesSensitiveTypeSourceEnum, bool) {
	enum, ok := mappingListSensitiveTypesSensitiveTypeSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveTypesEntityTypeEnum Enum with underlying type: string
type ListSensitiveTypesEntityTypeEnum string

// Set of constants representing the allowable values for ListSensitiveTypesEntityTypeEnum
const (
	ListSensitiveTypesEntityTypeType     ListSensitiveTypesEntityTypeEnum = "SENSITIVE_TYPE"
	ListSensitiveTypesEntityTypeCategory ListSensitiveTypesEntityTypeEnum = "SENSITIVE_CATEGORY"
)

var mappingListSensitiveTypesEntityTypeEnum = map[string]ListSensitiveTypesEntityTypeEnum{
	"SENSITIVE_TYPE":     ListSensitiveTypesEntityTypeType,
	"SENSITIVE_CATEGORY": ListSensitiveTypesEntityTypeCategory,
}

var mappingListSensitiveTypesEntityTypeEnumLowerCase = map[string]ListSensitiveTypesEntityTypeEnum{
	"sensitive_type":     ListSensitiveTypesEntityTypeType,
	"sensitive_category": ListSensitiveTypesEntityTypeCategory,
}

// GetListSensitiveTypesEntityTypeEnumValues Enumerates the set of values for ListSensitiveTypesEntityTypeEnum
func GetListSensitiveTypesEntityTypeEnumValues() []ListSensitiveTypesEntityTypeEnum {
	values := make([]ListSensitiveTypesEntityTypeEnum, 0)
	for _, v := range mappingListSensitiveTypesEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveTypesEntityTypeEnumStringValues Enumerates the set of values in String for ListSensitiveTypesEntityTypeEnum
func GetListSensitiveTypesEntityTypeEnumStringValues() []string {
	return []string{
		"SENSITIVE_TYPE",
		"SENSITIVE_CATEGORY",
	}
}

// GetMappingListSensitiveTypesEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveTypesEntityTypeEnum(val string) (ListSensitiveTypesEntityTypeEnum, bool) {
	enum, ok := mappingListSensitiveTypesEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveTypesSortOrderEnum Enum with underlying type: string
type ListSensitiveTypesSortOrderEnum string

// Set of constants representing the allowable values for ListSensitiveTypesSortOrderEnum
const (
	ListSensitiveTypesSortOrderAsc  ListSensitiveTypesSortOrderEnum = "ASC"
	ListSensitiveTypesSortOrderDesc ListSensitiveTypesSortOrderEnum = "DESC"
)

var mappingListSensitiveTypesSortOrderEnum = map[string]ListSensitiveTypesSortOrderEnum{
	"ASC":  ListSensitiveTypesSortOrderAsc,
	"DESC": ListSensitiveTypesSortOrderDesc,
}

var mappingListSensitiveTypesSortOrderEnumLowerCase = map[string]ListSensitiveTypesSortOrderEnum{
	"asc":  ListSensitiveTypesSortOrderAsc,
	"desc": ListSensitiveTypesSortOrderDesc,
}

// GetListSensitiveTypesSortOrderEnumValues Enumerates the set of values for ListSensitiveTypesSortOrderEnum
func GetListSensitiveTypesSortOrderEnumValues() []ListSensitiveTypesSortOrderEnum {
	values := make([]ListSensitiveTypesSortOrderEnum, 0)
	for _, v := range mappingListSensitiveTypesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveTypesSortOrderEnumStringValues Enumerates the set of values in String for ListSensitiveTypesSortOrderEnum
func GetListSensitiveTypesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSensitiveTypesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveTypesSortOrderEnum(val string) (ListSensitiveTypesSortOrderEnum, bool) {
	enum, ok := mappingListSensitiveTypesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveTypesSortByEnum Enum with underlying type: string
type ListSensitiveTypesSortByEnum string

// Set of constants representing the allowable values for ListSensitiveTypesSortByEnum
const (
	ListSensitiveTypesSortByTimecreated ListSensitiveTypesSortByEnum = "timeCreated"
	ListSensitiveTypesSortByDisplayname ListSensitiveTypesSortByEnum = "displayName"
)

var mappingListSensitiveTypesSortByEnum = map[string]ListSensitiveTypesSortByEnum{
	"timeCreated": ListSensitiveTypesSortByTimecreated,
	"displayName": ListSensitiveTypesSortByDisplayname,
}

var mappingListSensitiveTypesSortByEnumLowerCase = map[string]ListSensitiveTypesSortByEnum{
	"timecreated": ListSensitiveTypesSortByTimecreated,
	"displayname": ListSensitiveTypesSortByDisplayname,
}

// GetListSensitiveTypesSortByEnumValues Enumerates the set of values for ListSensitiveTypesSortByEnum
func GetListSensitiveTypesSortByEnumValues() []ListSensitiveTypesSortByEnum {
	values := make([]ListSensitiveTypesSortByEnum, 0)
	for _, v := range mappingListSensitiveTypesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveTypesSortByEnumStringValues Enumerates the set of values in String for ListSensitiveTypesSortByEnum
func GetListSensitiveTypesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSensitiveTypesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveTypesSortByEnum(val string) (ListSensitiveTypesSortByEnum, bool) {
	enum, ok := mappingListSensitiveTypesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveTypesLifecycleStateEnum Enum with underlying type: string
type ListSensitiveTypesLifecycleStateEnum string

// Set of constants representing the allowable values for ListSensitiveTypesLifecycleStateEnum
const (
	ListSensitiveTypesLifecycleStateCreating ListSensitiveTypesLifecycleStateEnum = "CREATING"
	ListSensitiveTypesLifecycleStateActive   ListSensitiveTypesLifecycleStateEnum = "ACTIVE"
	ListSensitiveTypesLifecycleStateUpdating ListSensitiveTypesLifecycleStateEnum = "UPDATING"
	ListSensitiveTypesLifecycleStateDeleting ListSensitiveTypesLifecycleStateEnum = "DELETING"
	ListSensitiveTypesLifecycleStateDeleted  ListSensitiveTypesLifecycleStateEnum = "DELETED"
	ListSensitiveTypesLifecycleStateFailed   ListSensitiveTypesLifecycleStateEnum = "FAILED"
)

var mappingListSensitiveTypesLifecycleStateEnum = map[string]ListSensitiveTypesLifecycleStateEnum{
	"CREATING": ListSensitiveTypesLifecycleStateCreating,
	"ACTIVE":   ListSensitiveTypesLifecycleStateActive,
	"UPDATING": ListSensitiveTypesLifecycleStateUpdating,
	"DELETING": ListSensitiveTypesLifecycleStateDeleting,
	"DELETED":  ListSensitiveTypesLifecycleStateDeleted,
	"FAILED":   ListSensitiveTypesLifecycleStateFailed,
}

var mappingListSensitiveTypesLifecycleStateEnumLowerCase = map[string]ListSensitiveTypesLifecycleStateEnum{
	"creating": ListSensitiveTypesLifecycleStateCreating,
	"active":   ListSensitiveTypesLifecycleStateActive,
	"updating": ListSensitiveTypesLifecycleStateUpdating,
	"deleting": ListSensitiveTypesLifecycleStateDeleting,
	"deleted":  ListSensitiveTypesLifecycleStateDeleted,
	"failed":   ListSensitiveTypesLifecycleStateFailed,
}

// GetListSensitiveTypesLifecycleStateEnumValues Enumerates the set of values for ListSensitiveTypesLifecycleStateEnum
func GetListSensitiveTypesLifecycleStateEnumValues() []ListSensitiveTypesLifecycleStateEnum {
	values := make([]ListSensitiveTypesLifecycleStateEnum, 0)
	for _, v := range mappingListSensitiveTypesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveTypesLifecycleStateEnumStringValues Enumerates the set of values in String for ListSensitiveTypesLifecycleStateEnum
func GetListSensitiveTypesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListSensitiveTypesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveTypesLifecycleStateEnum(val string) (ListSensitiveTypesLifecycleStateEnum, bool) {
	enum, ok := mappingListSensitiveTypesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
