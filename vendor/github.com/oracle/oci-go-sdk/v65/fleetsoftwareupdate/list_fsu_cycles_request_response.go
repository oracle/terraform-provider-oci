// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetsoftwareupdate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFsuCyclesRequest wrapper for the ListFsuCycles operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuCycles.go.html to see an example of how to use ListFsuCyclesRequest.
type ListFsuCyclesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose fsuCollectionId matches the given fsuCollectionId.
	FsuCollectionId *string `mandatory:"false" contributesTo:"query" name:"fsuCollectionId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState ListFsuCyclesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources whose Collection type matches the given type.
	CollectionType ListFsuCyclesCollectionTypeEnum `mandatory:"false" contributesTo:"query" name:"collectionType" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only entries whose targetVersion matches the given targetVersion.
	TargetVersion *string `mandatory:"false" contributesTo:"query" name:"targetVersion"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFsuCyclesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListFsuCyclesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFsuCyclesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFsuCyclesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFsuCyclesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFsuCyclesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFsuCyclesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFsuCyclesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListFsuCyclesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuCyclesCollectionTypeEnum(string(request.CollectionType)); !ok && request.CollectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CollectionType: %s. Supported values are: %s.", request.CollectionType, strings.Join(GetListFsuCyclesCollectionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuCyclesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFsuCyclesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuCyclesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFsuCyclesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFsuCyclesResponse wrapper for the ListFsuCycles operation
type ListFsuCyclesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FsuCycleSummaryCollection instances
	FsuCycleSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFsuCyclesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFsuCyclesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFsuCyclesLifecycleStateEnum Enum with underlying type: string
type ListFsuCyclesLifecycleStateEnum string

// Set of constants representing the allowable values for ListFsuCyclesLifecycleStateEnum
const (
	ListFsuCyclesLifecycleStateCreating       ListFsuCyclesLifecycleStateEnum = "CREATING"
	ListFsuCyclesLifecycleStateActive         ListFsuCyclesLifecycleStateEnum = "ACTIVE"
	ListFsuCyclesLifecycleStateUpdating       ListFsuCyclesLifecycleStateEnum = "UPDATING"
	ListFsuCyclesLifecycleStateInProgress     ListFsuCyclesLifecycleStateEnum = "IN_PROGRESS"
	ListFsuCyclesLifecycleStateFailed         ListFsuCyclesLifecycleStateEnum = "FAILED"
	ListFsuCyclesLifecycleStateNeedsAttention ListFsuCyclesLifecycleStateEnum = "NEEDS_ATTENTION"
	ListFsuCyclesLifecycleStateSucceeded      ListFsuCyclesLifecycleStateEnum = "SUCCEEDED"
	ListFsuCyclesLifecycleStateDeleting       ListFsuCyclesLifecycleStateEnum = "DELETING"
	ListFsuCyclesLifecycleStateDeleted        ListFsuCyclesLifecycleStateEnum = "DELETED"
)

var mappingListFsuCyclesLifecycleStateEnum = map[string]ListFsuCyclesLifecycleStateEnum{
	"CREATING":        ListFsuCyclesLifecycleStateCreating,
	"ACTIVE":          ListFsuCyclesLifecycleStateActive,
	"UPDATING":        ListFsuCyclesLifecycleStateUpdating,
	"IN_PROGRESS":     ListFsuCyclesLifecycleStateInProgress,
	"FAILED":          ListFsuCyclesLifecycleStateFailed,
	"NEEDS_ATTENTION": ListFsuCyclesLifecycleStateNeedsAttention,
	"SUCCEEDED":       ListFsuCyclesLifecycleStateSucceeded,
	"DELETING":        ListFsuCyclesLifecycleStateDeleting,
	"DELETED":         ListFsuCyclesLifecycleStateDeleted,
}

var mappingListFsuCyclesLifecycleStateEnumLowerCase = map[string]ListFsuCyclesLifecycleStateEnum{
	"creating":        ListFsuCyclesLifecycleStateCreating,
	"active":          ListFsuCyclesLifecycleStateActive,
	"updating":        ListFsuCyclesLifecycleStateUpdating,
	"in_progress":     ListFsuCyclesLifecycleStateInProgress,
	"failed":          ListFsuCyclesLifecycleStateFailed,
	"needs_attention": ListFsuCyclesLifecycleStateNeedsAttention,
	"succeeded":       ListFsuCyclesLifecycleStateSucceeded,
	"deleting":        ListFsuCyclesLifecycleStateDeleting,
	"deleted":         ListFsuCyclesLifecycleStateDeleted,
}

// GetListFsuCyclesLifecycleStateEnumValues Enumerates the set of values for ListFsuCyclesLifecycleStateEnum
func GetListFsuCyclesLifecycleStateEnumValues() []ListFsuCyclesLifecycleStateEnum {
	values := make([]ListFsuCyclesLifecycleStateEnum, 0)
	for _, v := range mappingListFsuCyclesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuCyclesLifecycleStateEnumStringValues Enumerates the set of values in String for ListFsuCyclesLifecycleStateEnum
func GetListFsuCyclesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"IN_PROGRESS",
		"FAILED",
		"NEEDS_ATTENTION",
		"SUCCEEDED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingListFsuCyclesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuCyclesLifecycleStateEnum(val string) (ListFsuCyclesLifecycleStateEnum, bool) {
	enum, ok := mappingListFsuCyclesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuCyclesCollectionTypeEnum Enum with underlying type: string
type ListFsuCyclesCollectionTypeEnum string

// Set of constants representing the allowable values for ListFsuCyclesCollectionTypeEnum
const (
	ListFsuCyclesCollectionTypeDb ListFsuCyclesCollectionTypeEnum = "DB"
	ListFsuCyclesCollectionTypeGi ListFsuCyclesCollectionTypeEnum = "GI"
)

var mappingListFsuCyclesCollectionTypeEnum = map[string]ListFsuCyclesCollectionTypeEnum{
	"DB": ListFsuCyclesCollectionTypeDb,
	"GI": ListFsuCyclesCollectionTypeGi,
}

var mappingListFsuCyclesCollectionTypeEnumLowerCase = map[string]ListFsuCyclesCollectionTypeEnum{
	"db": ListFsuCyclesCollectionTypeDb,
	"gi": ListFsuCyclesCollectionTypeGi,
}

// GetListFsuCyclesCollectionTypeEnumValues Enumerates the set of values for ListFsuCyclesCollectionTypeEnum
func GetListFsuCyclesCollectionTypeEnumValues() []ListFsuCyclesCollectionTypeEnum {
	values := make([]ListFsuCyclesCollectionTypeEnum, 0)
	for _, v := range mappingListFsuCyclesCollectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuCyclesCollectionTypeEnumStringValues Enumerates the set of values in String for ListFsuCyclesCollectionTypeEnum
func GetListFsuCyclesCollectionTypeEnumStringValues() []string {
	return []string{
		"DB",
		"GI",
	}
}

// GetMappingListFsuCyclesCollectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuCyclesCollectionTypeEnum(val string) (ListFsuCyclesCollectionTypeEnum, bool) {
	enum, ok := mappingListFsuCyclesCollectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuCyclesSortOrderEnum Enum with underlying type: string
type ListFsuCyclesSortOrderEnum string

// Set of constants representing the allowable values for ListFsuCyclesSortOrderEnum
const (
	ListFsuCyclesSortOrderAsc  ListFsuCyclesSortOrderEnum = "ASC"
	ListFsuCyclesSortOrderDesc ListFsuCyclesSortOrderEnum = "DESC"
)

var mappingListFsuCyclesSortOrderEnum = map[string]ListFsuCyclesSortOrderEnum{
	"ASC":  ListFsuCyclesSortOrderAsc,
	"DESC": ListFsuCyclesSortOrderDesc,
}

var mappingListFsuCyclesSortOrderEnumLowerCase = map[string]ListFsuCyclesSortOrderEnum{
	"asc":  ListFsuCyclesSortOrderAsc,
	"desc": ListFsuCyclesSortOrderDesc,
}

// GetListFsuCyclesSortOrderEnumValues Enumerates the set of values for ListFsuCyclesSortOrderEnum
func GetListFsuCyclesSortOrderEnumValues() []ListFsuCyclesSortOrderEnum {
	values := make([]ListFsuCyclesSortOrderEnum, 0)
	for _, v := range mappingListFsuCyclesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuCyclesSortOrderEnumStringValues Enumerates the set of values in String for ListFsuCyclesSortOrderEnum
func GetListFsuCyclesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFsuCyclesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuCyclesSortOrderEnum(val string) (ListFsuCyclesSortOrderEnum, bool) {
	enum, ok := mappingListFsuCyclesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuCyclesSortByEnum Enum with underlying type: string
type ListFsuCyclesSortByEnum string

// Set of constants representing the allowable values for ListFsuCyclesSortByEnum
const (
	ListFsuCyclesSortByTimecreated ListFsuCyclesSortByEnum = "timeCreated"
	ListFsuCyclesSortByDisplayname ListFsuCyclesSortByEnum = "displayName"
)

var mappingListFsuCyclesSortByEnum = map[string]ListFsuCyclesSortByEnum{
	"timeCreated": ListFsuCyclesSortByTimecreated,
	"displayName": ListFsuCyclesSortByDisplayname,
}

var mappingListFsuCyclesSortByEnumLowerCase = map[string]ListFsuCyclesSortByEnum{
	"timecreated": ListFsuCyclesSortByTimecreated,
	"displayname": ListFsuCyclesSortByDisplayname,
}

// GetListFsuCyclesSortByEnumValues Enumerates the set of values for ListFsuCyclesSortByEnum
func GetListFsuCyclesSortByEnumValues() []ListFsuCyclesSortByEnum {
	values := make([]ListFsuCyclesSortByEnum, 0)
	for _, v := range mappingListFsuCyclesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuCyclesSortByEnumStringValues Enumerates the set of values in String for ListFsuCyclesSortByEnum
func GetListFsuCyclesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListFsuCyclesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuCyclesSortByEnum(val string) (ListFsuCyclesSortByEnum, bool) {
	enum, ok := mappingListFsuCyclesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
