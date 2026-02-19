// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListGovernanceBenchmarkAttachmentsRequest wrapper for the ListGovernanceBenchmarkAttachments operation
type ListGovernanceBenchmarkAttachmentsRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID of the subject tenantId in which to list resources
	SubjectTenantId *string `mandatory:"false" contributesTo:"query" name:"subjectTenantId"`

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
	AccessLevel ListGovernanceBenchmarkAttachmentsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListGovernanceBenchmarkAttachmentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListGovernanceBenchmarkAttachmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListGovernanceBenchmarkAttachmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGovernanceBenchmarkAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGovernanceBenchmarkAttachmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGovernanceBenchmarkAttachmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGovernanceBenchmarkAttachmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGovernanceBenchmarkAttachmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGovernanceBenchmarkAttachmentsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListGovernanceBenchmarkAttachmentsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceBenchmarkAttachmentsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListGovernanceBenchmarkAttachmentsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceBenchmarkAttachmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGovernanceBenchmarkAttachmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceBenchmarkAttachmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGovernanceBenchmarkAttachmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGovernanceBenchmarkAttachmentsResponse wrapper for the ListGovernanceBenchmarkAttachments operation
type ListGovernanceBenchmarkAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GovernanceBenchmarkAttachmentCollection instances
	GovernanceBenchmarkAttachmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGovernanceBenchmarkAttachmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGovernanceBenchmarkAttachmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGovernanceBenchmarkAttachmentsAccessLevelEnum Enum with underlying type: string
type ListGovernanceBenchmarkAttachmentsAccessLevelEnum string

// Set of constants representing the allowable values for ListGovernanceBenchmarkAttachmentsAccessLevelEnum
const (
	ListGovernanceBenchmarkAttachmentsAccessLevelRestricted ListGovernanceBenchmarkAttachmentsAccessLevelEnum = "RESTRICTED"
	ListGovernanceBenchmarkAttachmentsAccessLevelAccessible ListGovernanceBenchmarkAttachmentsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListGovernanceBenchmarkAttachmentsAccessLevelEnum = map[string]ListGovernanceBenchmarkAttachmentsAccessLevelEnum{
	"RESTRICTED": ListGovernanceBenchmarkAttachmentsAccessLevelRestricted,
	"ACCESSIBLE": ListGovernanceBenchmarkAttachmentsAccessLevelAccessible,
}

var mappingListGovernanceBenchmarkAttachmentsAccessLevelEnumLowerCase = map[string]ListGovernanceBenchmarkAttachmentsAccessLevelEnum{
	"restricted": ListGovernanceBenchmarkAttachmentsAccessLevelRestricted,
	"accessible": ListGovernanceBenchmarkAttachmentsAccessLevelAccessible,
}

// GetListGovernanceBenchmarkAttachmentsAccessLevelEnumValues Enumerates the set of values for ListGovernanceBenchmarkAttachmentsAccessLevelEnum
func GetListGovernanceBenchmarkAttachmentsAccessLevelEnumValues() []ListGovernanceBenchmarkAttachmentsAccessLevelEnum {
	values := make([]ListGovernanceBenchmarkAttachmentsAccessLevelEnum, 0)
	for _, v := range mappingListGovernanceBenchmarkAttachmentsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceBenchmarkAttachmentsAccessLevelEnumStringValues Enumerates the set of values in String for ListGovernanceBenchmarkAttachmentsAccessLevelEnum
func GetListGovernanceBenchmarkAttachmentsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListGovernanceBenchmarkAttachmentsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceBenchmarkAttachmentsAccessLevelEnum(val string) (ListGovernanceBenchmarkAttachmentsAccessLevelEnum, bool) {
	enum, ok := mappingListGovernanceBenchmarkAttachmentsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceBenchmarkAttachmentsLifecycleStateEnum Enum with underlying type: string
type ListGovernanceBenchmarkAttachmentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListGovernanceBenchmarkAttachmentsLifecycleStateEnum
const (
	ListGovernanceBenchmarkAttachmentsLifecycleStateCreating ListGovernanceBenchmarkAttachmentsLifecycleStateEnum = "CREATING"
	ListGovernanceBenchmarkAttachmentsLifecycleStateUpdating ListGovernanceBenchmarkAttachmentsLifecycleStateEnum = "UPDATING"
	ListGovernanceBenchmarkAttachmentsLifecycleStateActive   ListGovernanceBenchmarkAttachmentsLifecycleStateEnum = "ACTIVE"
	ListGovernanceBenchmarkAttachmentsLifecycleStateInactive ListGovernanceBenchmarkAttachmentsLifecycleStateEnum = "INACTIVE"
	ListGovernanceBenchmarkAttachmentsLifecycleStateDeleting ListGovernanceBenchmarkAttachmentsLifecycleStateEnum = "DELETING"
	ListGovernanceBenchmarkAttachmentsLifecycleStateDeleted  ListGovernanceBenchmarkAttachmentsLifecycleStateEnum = "DELETED"
	ListGovernanceBenchmarkAttachmentsLifecycleStateFailed   ListGovernanceBenchmarkAttachmentsLifecycleStateEnum = "FAILED"
)

var mappingListGovernanceBenchmarkAttachmentsLifecycleStateEnum = map[string]ListGovernanceBenchmarkAttachmentsLifecycleStateEnum{
	"CREATING": ListGovernanceBenchmarkAttachmentsLifecycleStateCreating,
	"UPDATING": ListGovernanceBenchmarkAttachmentsLifecycleStateUpdating,
	"ACTIVE":   ListGovernanceBenchmarkAttachmentsLifecycleStateActive,
	"INACTIVE": ListGovernanceBenchmarkAttachmentsLifecycleStateInactive,
	"DELETING": ListGovernanceBenchmarkAttachmentsLifecycleStateDeleting,
	"DELETED":  ListGovernanceBenchmarkAttachmentsLifecycleStateDeleted,
	"FAILED":   ListGovernanceBenchmarkAttachmentsLifecycleStateFailed,
}

var mappingListGovernanceBenchmarkAttachmentsLifecycleStateEnumLowerCase = map[string]ListGovernanceBenchmarkAttachmentsLifecycleStateEnum{
	"creating": ListGovernanceBenchmarkAttachmentsLifecycleStateCreating,
	"updating": ListGovernanceBenchmarkAttachmentsLifecycleStateUpdating,
	"active":   ListGovernanceBenchmarkAttachmentsLifecycleStateActive,
	"inactive": ListGovernanceBenchmarkAttachmentsLifecycleStateInactive,
	"deleting": ListGovernanceBenchmarkAttachmentsLifecycleStateDeleting,
	"deleted":  ListGovernanceBenchmarkAttachmentsLifecycleStateDeleted,
	"failed":   ListGovernanceBenchmarkAttachmentsLifecycleStateFailed,
}

// GetListGovernanceBenchmarkAttachmentsLifecycleStateEnumValues Enumerates the set of values for ListGovernanceBenchmarkAttachmentsLifecycleStateEnum
func GetListGovernanceBenchmarkAttachmentsLifecycleStateEnumValues() []ListGovernanceBenchmarkAttachmentsLifecycleStateEnum {
	values := make([]ListGovernanceBenchmarkAttachmentsLifecycleStateEnum, 0)
	for _, v := range mappingListGovernanceBenchmarkAttachmentsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceBenchmarkAttachmentsLifecycleStateEnumStringValues Enumerates the set of values in String for ListGovernanceBenchmarkAttachmentsLifecycleStateEnum
func GetListGovernanceBenchmarkAttachmentsLifecycleStateEnumStringValues() []string {
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

// GetMappingListGovernanceBenchmarkAttachmentsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceBenchmarkAttachmentsLifecycleStateEnum(val string) (ListGovernanceBenchmarkAttachmentsLifecycleStateEnum, bool) {
	enum, ok := mappingListGovernanceBenchmarkAttachmentsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceBenchmarkAttachmentsSortOrderEnum Enum with underlying type: string
type ListGovernanceBenchmarkAttachmentsSortOrderEnum string

// Set of constants representing the allowable values for ListGovernanceBenchmarkAttachmentsSortOrderEnum
const (
	ListGovernanceBenchmarkAttachmentsSortOrderAsc  ListGovernanceBenchmarkAttachmentsSortOrderEnum = "ASC"
	ListGovernanceBenchmarkAttachmentsSortOrderDesc ListGovernanceBenchmarkAttachmentsSortOrderEnum = "DESC"
)

var mappingListGovernanceBenchmarkAttachmentsSortOrderEnum = map[string]ListGovernanceBenchmarkAttachmentsSortOrderEnum{
	"ASC":  ListGovernanceBenchmarkAttachmentsSortOrderAsc,
	"DESC": ListGovernanceBenchmarkAttachmentsSortOrderDesc,
}

var mappingListGovernanceBenchmarkAttachmentsSortOrderEnumLowerCase = map[string]ListGovernanceBenchmarkAttachmentsSortOrderEnum{
	"asc":  ListGovernanceBenchmarkAttachmentsSortOrderAsc,
	"desc": ListGovernanceBenchmarkAttachmentsSortOrderDesc,
}

// GetListGovernanceBenchmarkAttachmentsSortOrderEnumValues Enumerates the set of values for ListGovernanceBenchmarkAttachmentsSortOrderEnum
func GetListGovernanceBenchmarkAttachmentsSortOrderEnumValues() []ListGovernanceBenchmarkAttachmentsSortOrderEnum {
	values := make([]ListGovernanceBenchmarkAttachmentsSortOrderEnum, 0)
	for _, v := range mappingListGovernanceBenchmarkAttachmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceBenchmarkAttachmentsSortOrderEnumStringValues Enumerates the set of values in String for ListGovernanceBenchmarkAttachmentsSortOrderEnum
func GetListGovernanceBenchmarkAttachmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGovernanceBenchmarkAttachmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceBenchmarkAttachmentsSortOrderEnum(val string) (ListGovernanceBenchmarkAttachmentsSortOrderEnum, bool) {
	enum, ok := mappingListGovernanceBenchmarkAttachmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceBenchmarkAttachmentsSortByEnum Enum with underlying type: string
type ListGovernanceBenchmarkAttachmentsSortByEnum string

// Set of constants representing the allowable values for ListGovernanceBenchmarkAttachmentsSortByEnum
const (
	ListGovernanceBenchmarkAttachmentsSortByTimecreated ListGovernanceBenchmarkAttachmentsSortByEnum = "timeCreated"
	ListGovernanceBenchmarkAttachmentsSortByDisplayname ListGovernanceBenchmarkAttachmentsSortByEnum = "displayName"
)

var mappingListGovernanceBenchmarkAttachmentsSortByEnum = map[string]ListGovernanceBenchmarkAttachmentsSortByEnum{
	"timeCreated": ListGovernanceBenchmarkAttachmentsSortByTimecreated,
	"displayName": ListGovernanceBenchmarkAttachmentsSortByDisplayname,
}

var mappingListGovernanceBenchmarkAttachmentsSortByEnumLowerCase = map[string]ListGovernanceBenchmarkAttachmentsSortByEnum{
	"timecreated": ListGovernanceBenchmarkAttachmentsSortByTimecreated,
	"displayname": ListGovernanceBenchmarkAttachmentsSortByDisplayname,
}

// GetListGovernanceBenchmarkAttachmentsSortByEnumValues Enumerates the set of values for ListGovernanceBenchmarkAttachmentsSortByEnum
func GetListGovernanceBenchmarkAttachmentsSortByEnumValues() []ListGovernanceBenchmarkAttachmentsSortByEnum {
	values := make([]ListGovernanceBenchmarkAttachmentsSortByEnum, 0)
	for _, v := range mappingListGovernanceBenchmarkAttachmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceBenchmarkAttachmentsSortByEnumStringValues Enumerates the set of values in String for ListGovernanceBenchmarkAttachmentsSortByEnum
func GetListGovernanceBenchmarkAttachmentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListGovernanceBenchmarkAttachmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceBenchmarkAttachmentsSortByEnum(val string) (ListGovernanceBenchmarkAttachmentsSortByEnum, bool) {
	enum, ok := mappingListGovernanceBenchmarkAttachmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
