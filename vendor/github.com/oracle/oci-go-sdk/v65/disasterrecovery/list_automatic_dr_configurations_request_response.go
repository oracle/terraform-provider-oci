// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAutomaticDrConfigurationsRequest wrapper for the ListAutomaticDrConfigurations operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/ListAutomaticDrConfigurations.go.html to see an example of how to use ListAutomaticDrConfigurationsRequest.
type ListAutomaticDrConfigurationsRequest struct {

	// The OCID of the DR protection group. Mandatory query param.
	// Example: `ocid1.drprotectiongroup.oc1..uniqueID`
	DrProtectionGroupId *string `mandatory:"true" contributesTo:"query" name:"drProtectionGroupId"`

	// The OCID of the automatic DR configuration.
	// Example: `ocid1.automaticDrConfiguration.oc1..uniqueID`
	AutomaticDrConfigurationId *string `mandatory:"false" contributesTo:"query" name:"automaticDrConfigurationId"`

	// A filter to return only Automatic DR configurations that match the given lifecycle state.
	LifecycleState ListAutomaticDrConfigurationsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only Automatic DR configurations that do not match the given lifecycle state.
	LifecycleStateNotEqualTo ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum `mandatory:"false" contributesTo:"query" name:"lifecycleStateNotEqualTo" omitEmpty:"true"`

	// A filter to return only resources that match the given display name.
	// Example: `MyResourceDisplayName`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `100`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAutomaticDrConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	// Default order for displayName is ascending. If no value is specified timeCreated is default.
	// Example: `MyResourceDisplayName`
	SortBy ListAutomaticDrConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutomaticDrConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutomaticDrConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutomaticDrConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutomaticDrConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutomaticDrConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAutomaticDrConfigurationsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAutomaticDrConfigurationsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum(string(request.LifecycleStateNotEqualTo)); !ok && request.LifecycleStateNotEqualTo != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleStateNotEqualTo: %s. Supported values are: %s.", request.LifecycleStateNotEqualTo, strings.Join(GetListAutomaticDrConfigurationsLifecycleStateNotEqualToEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutomaticDrConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAutomaticDrConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutomaticDrConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAutomaticDrConfigurationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutomaticDrConfigurationsResponse wrapper for the ListAutomaticDrConfigurations operation
type ListAutomaticDrConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AutomaticDrConfigurationCollection instances
	AutomaticDrConfigurationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutomaticDrConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutomaticDrConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutomaticDrConfigurationsLifecycleStateEnum Enum with underlying type: string
type ListAutomaticDrConfigurationsLifecycleStateEnum string

// Set of constants representing the allowable values for ListAutomaticDrConfigurationsLifecycleStateEnum
const (
	ListAutomaticDrConfigurationsLifecycleStateCreating       ListAutomaticDrConfigurationsLifecycleStateEnum = "CREATING"
	ListAutomaticDrConfigurationsLifecycleStateUpdating       ListAutomaticDrConfigurationsLifecycleStateEnum = "UPDATING"
	ListAutomaticDrConfigurationsLifecycleStateActive         ListAutomaticDrConfigurationsLifecycleStateEnum = "ACTIVE"
	ListAutomaticDrConfigurationsLifecycleStateInactive       ListAutomaticDrConfigurationsLifecycleStateEnum = "INACTIVE"
	ListAutomaticDrConfigurationsLifecycleStateNeedsAttention ListAutomaticDrConfigurationsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListAutomaticDrConfigurationsLifecycleStateDeleting       ListAutomaticDrConfigurationsLifecycleStateEnum = "DELETING"
	ListAutomaticDrConfigurationsLifecycleStateDeleted        ListAutomaticDrConfigurationsLifecycleStateEnum = "DELETED"
	ListAutomaticDrConfigurationsLifecycleStateFailed         ListAutomaticDrConfigurationsLifecycleStateEnum = "FAILED"
)

var mappingListAutomaticDrConfigurationsLifecycleStateEnum = map[string]ListAutomaticDrConfigurationsLifecycleStateEnum{
	"CREATING":        ListAutomaticDrConfigurationsLifecycleStateCreating,
	"UPDATING":        ListAutomaticDrConfigurationsLifecycleStateUpdating,
	"ACTIVE":          ListAutomaticDrConfigurationsLifecycleStateActive,
	"INACTIVE":        ListAutomaticDrConfigurationsLifecycleStateInactive,
	"NEEDS_ATTENTION": ListAutomaticDrConfigurationsLifecycleStateNeedsAttention,
	"DELETING":        ListAutomaticDrConfigurationsLifecycleStateDeleting,
	"DELETED":         ListAutomaticDrConfigurationsLifecycleStateDeleted,
	"FAILED":          ListAutomaticDrConfigurationsLifecycleStateFailed,
}

var mappingListAutomaticDrConfigurationsLifecycleStateEnumLowerCase = map[string]ListAutomaticDrConfigurationsLifecycleStateEnum{
	"creating":        ListAutomaticDrConfigurationsLifecycleStateCreating,
	"updating":        ListAutomaticDrConfigurationsLifecycleStateUpdating,
	"active":          ListAutomaticDrConfigurationsLifecycleStateActive,
	"inactive":        ListAutomaticDrConfigurationsLifecycleStateInactive,
	"needs_attention": ListAutomaticDrConfigurationsLifecycleStateNeedsAttention,
	"deleting":        ListAutomaticDrConfigurationsLifecycleStateDeleting,
	"deleted":         ListAutomaticDrConfigurationsLifecycleStateDeleted,
	"failed":          ListAutomaticDrConfigurationsLifecycleStateFailed,
}

// GetListAutomaticDrConfigurationsLifecycleStateEnumValues Enumerates the set of values for ListAutomaticDrConfigurationsLifecycleStateEnum
func GetListAutomaticDrConfigurationsLifecycleStateEnumValues() []ListAutomaticDrConfigurationsLifecycleStateEnum {
	values := make([]ListAutomaticDrConfigurationsLifecycleStateEnum, 0)
	for _, v := range mappingListAutomaticDrConfigurationsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutomaticDrConfigurationsLifecycleStateEnumStringValues Enumerates the set of values in String for ListAutomaticDrConfigurationsLifecycleStateEnum
func GetListAutomaticDrConfigurationsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListAutomaticDrConfigurationsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutomaticDrConfigurationsLifecycleStateEnum(val string) (ListAutomaticDrConfigurationsLifecycleStateEnum, bool) {
	enum, ok := mappingListAutomaticDrConfigurationsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum Enum with underlying type: string
type ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum string

// Set of constants representing the allowable values for ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum
const (
	ListAutomaticDrConfigurationsLifecycleStateNotEqualToCreating       ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum = "CREATING"
	ListAutomaticDrConfigurationsLifecycleStateNotEqualToUpdating       ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum = "UPDATING"
	ListAutomaticDrConfigurationsLifecycleStateNotEqualToActive         ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum = "ACTIVE"
	ListAutomaticDrConfigurationsLifecycleStateNotEqualToInactive       ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum = "INACTIVE"
	ListAutomaticDrConfigurationsLifecycleStateNotEqualToNeedsAttention ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum = "NEEDS_ATTENTION"
	ListAutomaticDrConfigurationsLifecycleStateNotEqualToDeleting       ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum = "DELETING"
	ListAutomaticDrConfigurationsLifecycleStateNotEqualToDeleted        ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum = "DELETED"
	ListAutomaticDrConfigurationsLifecycleStateNotEqualToFailed         ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum = "FAILED"
)

var mappingListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum = map[string]ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum{
	"CREATING":        ListAutomaticDrConfigurationsLifecycleStateNotEqualToCreating,
	"UPDATING":        ListAutomaticDrConfigurationsLifecycleStateNotEqualToUpdating,
	"ACTIVE":          ListAutomaticDrConfigurationsLifecycleStateNotEqualToActive,
	"INACTIVE":        ListAutomaticDrConfigurationsLifecycleStateNotEqualToInactive,
	"NEEDS_ATTENTION": ListAutomaticDrConfigurationsLifecycleStateNotEqualToNeedsAttention,
	"DELETING":        ListAutomaticDrConfigurationsLifecycleStateNotEqualToDeleting,
	"DELETED":         ListAutomaticDrConfigurationsLifecycleStateNotEqualToDeleted,
	"FAILED":          ListAutomaticDrConfigurationsLifecycleStateNotEqualToFailed,
}

var mappingListAutomaticDrConfigurationsLifecycleStateNotEqualToEnumLowerCase = map[string]ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum{
	"creating":        ListAutomaticDrConfigurationsLifecycleStateNotEqualToCreating,
	"updating":        ListAutomaticDrConfigurationsLifecycleStateNotEqualToUpdating,
	"active":          ListAutomaticDrConfigurationsLifecycleStateNotEqualToActive,
	"inactive":        ListAutomaticDrConfigurationsLifecycleStateNotEqualToInactive,
	"needs_attention": ListAutomaticDrConfigurationsLifecycleStateNotEqualToNeedsAttention,
	"deleting":        ListAutomaticDrConfigurationsLifecycleStateNotEqualToDeleting,
	"deleted":         ListAutomaticDrConfigurationsLifecycleStateNotEqualToDeleted,
	"failed":          ListAutomaticDrConfigurationsLifecycleStateNotEqualToFailed,
}

// GetListAutomaticDrConfigurationsLifecycleStateNotEqualToEnumValues Enumerates the set of values for ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum
func GetListAutomaticDrConfigurationsLifecycleStateNotEqualToEnumValues() []ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum {
	values := make([]ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum, 0)
	for _, v := range mappingListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutomaticDrConfigurationsLifecycleStateNotEqualToEnumStringValues Enumerates the set of values in String for ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum
func GetListAutomaticDrConfigurationsLifecycleStateNotEqualToEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum(val string) (ListAutomaticDrConfigurationsLifecycleStateNotEqualToEnum, bool) {
	enum, ok := mappingListAutomaticDrConfigurationsLifecycleStateNotEqualToEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutomaticDrConfigurationsSortOrderEnum Enum with underlying type: string
type ListAutomaticDrConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListAutomaticDrConfigurationsSortOrderEnum
const (
	ListAutomaticDrConfigurationsSortOrderAsc  ListAutomaticDrConfigurationsSortOrderEnum = "ASC"
	ListAutomaticDrConfigurationsSortOrderDesc ListAutomaticDrConfigurationsSortOrderEnum = "DESC"
)

var mappingListAutomaticDrConfigurationsSortOrderEnum = map[string]ListAutomaticDrConfigurationsSortOrderEnum{
	"ASC":  ListAutomaticDrConfigurationsSortOrderAsc,
	"DESC": ListAutomaticDrConfigurationsSortOrderDesc,
}

var mappingListAutomaticDrConfigurationsSortOrderEnumLowerCase = map[string]ListAutomaticDrConfigurationsSortOrderEnum{
	"asc":  ListAutomaticDrConfigurationsSortOrderAsc,
	"desc": ListAutomaticDrConfigurationsSortOrderDesc,
}

// GetListAutomaticDrConfigurationsSortOrderEnumValues Enumerates the set of values for ListAutomaticDrConfigurationsSortOrderEnum
func GetListAutomaticDrConfigurationsSortOrderEnumValues() []ListAutomaticDrConfigurationsSortOrderEnum {
	values := make([]ListAutomaticDrConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListAutomaticDrConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutomaticDrConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListAutomaticDrConfigurationsSortOrderEnum
func GetListAutomaticDrConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAutomaticDrConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutomaticDrConfigurationsSortOrderEnum(val string) (ListAutomaticDrConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListAutomaticDrConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutomaticDrConfigurationsSortByEnum Enum with underlying type: string
type ListAutomaticDrConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListAutomaticDrConfigurationsSortByEnum
const (
	ListAutomaticDrConfigurationsSortByTimecreated ListAutomaticDrConfigurationsSortByEnum = "timeCreated"
	ListAutomaticDrConfigurationsSortByDisplayname ListAutomaticDrConfigurationsSortByEnum = "displayName"
)

var mappingListAutomaticDrConfigurationsSortByEnum = map[string]ListAutomaticDrConfigurationsSortByEnum{
	"timeCreated": ListAutomaticDrConfigurationsSortByTimecreated,
	"displayName": ListAutomaticDrConfigurationsSortByDisplayname,
}

var mappingListAutomaticDrConfigurationsSortByEnumLowerCase = map[string]ListAutomaticDrConfigurationsSortByEnum{
	"timecreated": ListAutomaticDrConfigurationsSortByTimecreated,
	"displayname": ListAutomaticDrConfigurationsSortByDisplayname,
}

// GetListAutomaticDrConfigurationsSortByEnumValues Enumerates the set of values for ListAutomaticDrConfigurationsSortByEnum
func GetListAutomaticDrConfigurationsSortByEnumValues() []ListAutomaticDrConfigurationsSortByEnum {
	values := make([]ListAutomaticDrConfigurationsSortByEnum, 0)
	for _, v := range mappingListAutomaticDrConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutomaticDrConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListAutomaticDrConfigurationsSortByEnum
func GetListAutomaticDrConfigurationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAutomaticDrConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutomaticDrConfigurationsSortByEnum(val string) (ListAutomaticDrConfigurationsSortByEnum, bool) {
	enum, ok := mappingListAutomaticDrConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
