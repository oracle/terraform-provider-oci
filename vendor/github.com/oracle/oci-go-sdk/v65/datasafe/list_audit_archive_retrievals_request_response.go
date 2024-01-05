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

// ListAuditArchiveRetrievalsRequest wrapper for the ListAuditArchiveRetrievals operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditArchiveRetrievals.go.html to see an example of how to use ListAuditArchiveRetrievalsRequest.
type ListAuditArchiveRetrievalsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListAuditArchiveRetrievalsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// OCID of the archive retrieval.
	AuditArchiveRetrievalId *string `mandatory:"false" contributesTo:"query" name:"auditArchiveRetrievalId"`

	// The OCID of the target associated with the archive retrieval.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that matches the specified lifecycle state.
	LifecycleState ListAuditArchiveRetrievalsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The date time when retrieved archive data will be deleted from Data Safe and unloaded back into archival.
	TimeOfExpiry *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeOfExpiry"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListAuditArchiveRetrievalsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListAuditArchiveRetrievalsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAuditArchiveRetrievalsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAuditArchiveRetrievalsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAuditArchiveRetrievalsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAuditArchiveRetrievalsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAuditArchiveRetrievalsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAuditArchiveRetrievalsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListAuditArchiveRetrievalsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditArchiveRetrievalsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAuditArchiveRetrievalsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditArchiveRetrievalsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAuditArchiveRetrievalsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditArchiveRetrievalsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAuditArchiveRetrievalsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAuditArchiveRetrievalsResponse wrapper for the ListAuditArchiveRetrievals operation
type ListAuditArchiveRetrievalsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AuditArchiveRetrievalCollection instances
	AuditArchiveRetrievalCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListAuditArchiveRetrievalsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAuditArchiveRetrievalsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAuditArchiveRetrievalsAccessLevelEnum Enum with underlying type: string
type ListAuditArchiveRetrievalsAccessLevelEnum string

// Set of constants representing the allowable values for ListAuditArchiveRetrievalsAccessLevelEnum
const (
	ListAuditArchiveRetrievalsAccessLevelRestricted ListAuditArchiveRetrievalsAccessLevelEnum = "RESTRICTED"
	ListAuditArchiveRetrievalsAccessLevelAccessible ListAuditArchiveRetrievalsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListAuditArchiveRetrievalsAccessLevelEnum = map[string]ListAuditArchiveRetrievalsAccessLevelEnum{
	"RESTRICTED": ListAuditArchiveRetrievalsAccessLevelRestricted,
	"ACCESSIBLE": ListAuditArchiveRetrievalsAccessLevelAccessible,
}

var mappingListAuditArchiveRetrievalsAccessLevelEnumLowerCase = map[string]ListAuditArchiveRetrievalsAccessLevelEnum{
	"restricted": ListAuditArchiveRetrievalsAccessLevelRestricted,
	"accessible": ListAuditArchiveRetrievalsAccessLevelAccessible,
}

// GetListAuditArchiveRetrievalsAccessLevelEnumValues Enumerates the set of values for ListAuditArchiveRetrievalsAccessLevelEnum
func GetListAuditArchiveRetrievalsAccessLevelEnumValues() []ListAuditArchiveRetrievalsAccessLevelEnum {
	values := make([]ListAuditArchiveRetrievalsAccessLevelEnum, 0)
	for _, v := range mappingListAuditArchiveRetrievalsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditArchiveRetrievalsAccessLevelEnumStringValues Enumerates the set of values in String for ListAuditArchiveRetrievalsAccessLevelEnum
func GetListAuditArchiveRetrievalsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListAuditArchiveRetrievalsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditArchiveRetrievalsAccessLevelEnum(val string) (ListAuditArchiveRetrievalsAccessLevelEnum, bool) {
	enum, ok := mappingListAuditArchiveRetrievalsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditArchiveRetrievalsLifecycleStateEnum Enum with underlying type: string
type ListAuditArchiveRetrievalsLifecycleStateEnum string

// Set of constants representing the allowable values for ListAuditArchiveRetrievalsLifecycleStateEnum
const (
	ListAuditArchiveRetrievalsLifecycleStateCreating       ListAuditArchiveRetrievalsLifecycleStateEnum = "CREATING"
	ListAuditArchiveRetrievalsLifecycleStateActive         ListAuditArchiveRetrievalsLifecycleStateEnum = "ACTIVE"
	ListAuditArchiveRetrievalsLifecycleStateNeedsAttention ListAuditArchiveRetrievalsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListAuditArchiveRetrievalsLifecycleStateFailed         ListAuditArchiveRetrievalsLifecycleStateEnum = "FAILED"
	ListAuditArchiveRetrievalsLifecycleStateDeleting       ListAuditArchiveRetrievalsLifecycleStateEnum = "DELETING"
	ListAuditArchiveRetrievalsLifecycleStateDeleted        ListAuditArchiveRetrievalsLifecycleStateEnum = "DELETED"
	ListAuditArchiveRetrievalsLifecycleStateUpdating       ListAuditArchiveRetrievalsLifecycleStateEnum = "UPDATING"
)

var mappingListAuditArchiveRetrievalsLifecycleStateEnum = map[string]ListAuditArchiveRetrievalsLifecycleStateEnum{
	"CREATING":        ListAuditArchiveRetrievalsLifecycleStateCreating,
	"ACTIVE":          ListAuditArchiveRetrievalsLifecycleStateActive,
	"NEEDS_ATTENTION": ListAuditArchiveRetrievalsLifecycleStateNeedsAttention,
	"FAILED":          ListAuditArchiveRetrievalsLifecycleStateFailed,
	"DELETING":        ListAuditArchiveRetrievalsLifecycleStateDeleting,
	"DELETED":         ListAuditArchiveRetrievalsLifecycleStateDeleted,
	"UPDATING":        ListAuditArchiveRetrievalsLifecycleStateUpdating,
}

var mappingListAuditArchiveRetrievalsLifecycleStateEnumLowerCase = map[string]ListAuditArchiveRetrievalsLifecycleStateEnum{
	"creating":        ListAuditArchiveRetrievalsLifecycleStateCreating,
	"active":          ListAuditArchiveRetrievalsLifecycleStateActive,
	"needs_attention": ListAuditArchiveRetrievalsLifecycleStateNeedsAttention,
	"failed":          ListAuditArchiveRetrievalsLifecycleStateFailed,
	"deleting":        ListAuditArchiveRetrievalsLifecycleStateDeleting,
	"deleted":         ListAuditArchiveRetrievalsLifecycleStateDeleted,
	"updating":        ListAuditArchiveRetrievalsLifecycleStateUpdating,
}

// GetListAuditArchiveRetrievalsLifecycleStateEnumValues Enumerates the set of values for ListAuditArchiveRetrievalsLifecycleStateEnum
func GetListAuditArchiveRetrievalsLifecycleStateEnumValues() []ListAuditArchiveRetrievalsLifecycleStateEnum {
	values := make([]ListAuditArchiveRetrievalsLifecycleStateEnum, 0)
	for _, v := range mappingListAuditArchiveRetrievalsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditArchiveRetrievalsLifecycleStateEnumStringValues Enumerates the set of values in String for ListAuditArchiveRetrievalsLifecycleStateEnum
func GetListAuditArchiveRetrievalsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"FAILED",
		"DELETING",
		"DELETED",
		"UPDATING",
	}
}

// GetMappingListAuditArchiveRetrievalsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditArchiveRetrievalsLifecycleStateEnum(val string) (ListAuditArchiveRetrievalsLifecycleStateEnum, bool) {
	enum, ok := mappingListAuditArchiveRetrievalsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditArchiveRetrievalsSortOrderEnum Enum with underlying type: string
type ListAuditArchiveRetrievalsSortOrderEnum string

// Set of constants representing the allowable values for ListAuditArchiveRetrievalsSortOrderEnum
const (
	ListAuditArchiveRetrievalsSortOrderAsc  ListAuditArchiveRetrievalsSortOrderEnum = "ASC"
	ListAuditArchiveRetrievalsSortOrderDesc ListAuditArchiveRetrievalsSortOrderEnum = "DESC"
)

var mappingListAuditArchiveRetrievalsSortOrderEnum = map[string]ListAuditArchiveRetrievalsSortOrderEnum{
	"ASC":  ListAuditArchiveRetrievalsSortOrderAsc,
	"DESC": ListAuditArchiveRetrievalsSortOrderDesc,
}

var mappingListAuditArchiveRetrievalsSortOrderEnumLowerCase = map[string]ListAuditArchiveRetrievalsSortOrderEnum{
	"asc":  ListAuditArchiveRetrievalsSortOrderAsc,
	"desc": ListAuditArchiveRetrievalsSortOrderDesc,
}

// GetListAuditArchiveRetrievalsSortOrderEnumValues Enumerates the set of values for ListAuditArchiveRetrievalsSortOrderEnum
func GetListAuditArchiveRetrievalsSortOrderEnumValues() []ListAuditArchiveRetrievalsSortOrderEnum {
	values := make([]ListAuditArchiveRetrievalsSortOrderEnum, 0)
	for _, v := range mappingListAuditArchiveRetrievalsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditArchiveRetrievalsSortOrderEnumStringValues Enumerates the set of values in String for ListAuditArchiveRetrievalsSortOrderEnum
func GetListAuditArchiveRetrievalsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAuditArchiveRetrievalsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditArchiveRetrievalsSortOrderEnum(val string) (ListAuditArchiveRetrievalsSortOrderEnum, bool) {
	enum, ok := mappingListAuditArchiveRetrievalsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditArchiveRetrievalsSortByEnum Enum with underlying type: string
type ListAuditArchiveRetrievalsSortByEnum string

// Set of constants representing the allowable values for ListAuditArchiveRetrievalsSortByEnum
const (
	ListAuditArchiveRetrievalsSortByTimecreated ListAuditArchiveRetrievalsSortByEnum = "TIMECREATED"
	ListAuditArchiveRetrievalsSortByDisplayname ListAuditArchiveRetrievalsSortByEnum = "DISPLAYNAME"
)

var mappingListAuditArchiveRetrievalsSortByEnum = map[string]ListAuditArchiveRetrievalsSortByEnum{
	"TIMECREATED": ListAuditArchiveRetrievalsSortByTimecreated,
	"DISPLAYNAME": ListAuditArchiveRetrievalsSortByDisplayname,
}

var mappingListAuditArchiveRetrievalsSortByEnumLowerCase = map[string]ListAuditArchiveRetrievalsSortByEnum{
	"timecreated": ListAuditArchiveRetrievalsSortByTimecreated,
	"displayname": ListAuditArchiveRetrievalsSortByDisplayname,
}

// GetListAuditArchiveRetrievalsSortByEnumValues Enumerates the set of values for ListAuditArchiveRetrievalsSortByEnum
func GetListAuditArchiveRetrievalsSortByEnumValues() []ListAuditArchiveRetrievalsSortByEnum {
	values := make([]ListAuditArchiveRetrievalsSortByEnum, 0)
	for _, v := range mappingListAuditArchiveRetrievalsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditArchiveRetrievalsSortByEnumStringValues Enumerates the set of values in String for ListAuditArchiveRetrievalsSortByEnum
func GetListAuditArchiveRetrievalsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAuditArchiveRetrievalsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditArchiveRetrievalsSortByEnum(val string) (ListAuditArchiveRetrievalsSortByEnum, bool) {
	enum, ok := mappingListAuditArchiveRetrievalsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
