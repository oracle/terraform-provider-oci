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

// ListFsuCollectionsRequest wrapper for the ListFsuCollections operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuCollections.go.html to see an example of how to use ListFsuCollectionsRequest.
type ListFsuCollectionsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState ListFsuCollectionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources whose type matches the given type.
	Type ListFsuCollectionsTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFsuCollectionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListFsuCollectionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFsuCollectionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFsuCollectionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFsuCollectionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFsuCollectionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFsuCollectionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFsuCollectionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListFsuCollectionsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuCollectionsTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListFsuCollectionsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuCollectionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFsuCollectionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuCollectionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFsuCollectionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFsuCollectionsResponse wrapper for the ListFsuCollections operation
type ListFsuCollectionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FsuCollectionSummaryCollection instances
	FsuCollectionSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFsuCollectionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFsuCollectionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFsuCollectionsLifecycleStateEnum Enum with underlying type: string
type ListFsuCollectionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListFsuCollectionsLifecycleStateEnum
const (
	ListFsuCollectionsLifecycleStateCreating       ListFsuCollectionsLifecycleStateEnum = "CREATING"
	ListFsuCollectionsLifecycleStateUpdating       ListFsuCollectionsLifecycleStateEnum = "UPDATING"
	ListFsuCollectionsLifecycleStateActive         ListFsuCollectionsLifecycleStateEnum = "ACTIVE"
	ListFsuCollectionsLifecycleStateNeedsAttention ListFsuCollectionsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListFsuCollectionsLifecycleStateDeleting       ListFsuCollectionsLifecycleStateEnum = "DELETING"
	ListFsuCollectionsLifecycleStateDeleted        ListFsuCollectionsLifecycleStateEnum = "DELETED"
	ListFsuCollectionsLifecycleStateFailed         ListFsuCollectionsLifecycleStateEnum = "FAILED"
)

var mappingListFsuCollectionsLifecycleStateEnum = map[string]ListFsuCollectionsLifecycleStateEnum{
	"CREATING":        ListFsuCollectionsLifecycleStateCreating,
	"UPDATING":        ListFsuCollectionsLifecycleStateUpdating,
	"ACTIVE":          ListFsuCollectionsLifecycleStateActive,
	"NEEDS_ATTENTION": ListFsuCollectionsLifecycleStateNeedsAttention,
	"DELETING":        ListFsuCollectionsLifecycleStateDeleting,
	"DELETED":         ListFsuCollectionsLifecycleStateDeleted,
	"FAILED":          ListFsuCollectionsLifecycleStateFailed,
}

var mappingListFsuCollectionsLifecycleStateEnumLowerCase = map[string]ListFsuCollectionsLifecycleStateEnum{
	"creating":        ListFsuCollectionsLifecycleStateCreating,
	"updating":        ListFsuCollectionsLifecycleStateUpdating,
	"active":          ListFsuCollectionsLifecycleStateActive,
	"needs_attention": ListFsuCollectionsLifecycleStateNeedsAttention,
	"deleting":        ListFsuCollectionsLifecycleStateDeleting,
	"deleted":         ListFsuCollectionsLifecycleStateDeleted,
	"failed":          ListFsuCollectionsLifecycleStateFailed,
}

// GetListFsuCollectionsLifecycleStateEnumValues Enumerates the set of values for ListFsuCollectionsLifecycleStateEnum
func GetListFsuCollectionsLifecycleStateEnumValues() []ListFsuCollectionsLifecycleStateEnum {
	values := make([]ListFsuCollectionsLifecycleStateEnum, 0)
	for _, v := range mappingListFsuCollectionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuCollectionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListFsuCollectionsLifecycleStateEnum
func GetListFsuCollectionsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListFsuCollectionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuCollectionsLifecycleStateEnum(val string) (ListFsuCollectionsLifecycleStateEnum, bool) {
	enum, ok := mappingListFsuCollectionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuCollectionsTypeEnum Enum with underlying type: string
type ListFsuCollectionsTypeEnum string

// Set of constants representing the allowable values for ListFsuCollectionsTypeEnum
const (
	ListFsuCollectionsTypeDb ListFsuCollectionsTypeEnum = "DB"
	ListFsuCollectionsTypeGi ListFsuCollectionsTypeEnum = "GI"
)

var mappingListFsuCollectionsTypeEnum = map[string]ListFsuCollectionsTypeEnum{
	"DB": ListFsuCollectionsTypeDb,
	"GI": ListFsuCollectionsTypeGi,
}

var mappingListFsuCollectionsTypeEnumLowerCase = map[string]ListFsuCollectionsTypeEnum{
	"db": ListFsuCollectionsTypeDb,
	"gi": ListFsuCollectionsTypeGi,
}

// GetListFsuCollectionsTypeEnumValues Enumerates the set of values for ListFsuCollectionsTypeEnum
func GetListFsuCollectionsTypeEnumValues() []ListFsuCollectionsTypeEnum {
	values := make([]ListFsuCollectionsTypeEnum, 0)
	for _, v := range mappingListFsuCollectionsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuCollectionsTypeEnumStringValues Enumerates the set of values in String for ListFsuCollectionsTypeEnum
func GetListFsuCollectionsTypeEnumStringValues() []string {
	return []string{
		"DB",
		"GI",
	}
}

// GetMappingListFsuCollectionsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuCollectionsTypeEnum(val string) (ListFsuCollectionsTypeEnum, bool) {
	enum, ok := mappingListFsuCollectionsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuCollectionsSortOrderEnum Enum with underlying type: string
type ListFsuCollectionsSortOrderEnum string

// Set of constants representing the allowable values for ListFsuCollectionsSortOrderEnum
const (
	ListFsuCollectionsSortOrderAsc  ListFsuCollectionsSortOrderEnum = "ASC"
	ListFsuCollectionsSortOrderDesc ListFsuCollectionsSortOrderEnum = "DESC"
)

var mappingListFsuCollectionsSortOrderEnum = map[string]ListFsuCollectionsSortOrderEnum{
	"ASC":  ListFsuCollectionsSortOrderAsc,
	"DESC": ListFsuCollectionsSortOrderDesc,
}

var mappingListFsuCollectionsSortOrderEnumLowerCase = map[string]ListFsuCollectionsSortOrderEnum{
	"asc":  ListFsuCollectionsSortOrderAsc,
	"desc": ListFsuCollectionsSortOrderDesc,
}

// GetListFsuCollectionsSortOrderEnumValues Enumerates the set of values for ListFsuCollectionsSortOrderEnum
func GetListFsuCollectionsSortOrderEnumValues() []ListFsuCollectionsSortOrderEnum {
	values := make([]ListFsuCollectionsSortOrderEnum, 0)
	for _, v := range mappingListFsuCollectionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuCollectionsSortOrderEnumStringValues Enumerates the set of values in String for ListFsuCollectionsSortOrderEnum
func GetListFsuCollectionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFsuCollectionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuCollectionsSortOrderEnum(val string) (ListFsuCollectionsSortOrderEnum, bool) {
	enum, ok := mappingListFsuCollectionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuCollectionsSortByEnum Enum with underlying type: string
type ListFsuCollectionsSortByEnum string

// Set of constants representing the allowable values for ListFsuCollectionsSortByEnum
const (
	ListFsuCollectionsSortByTimecreated ListFsuCollectionsSortByEnum = "timeCreated"
	ListFsuCollectionsSortByDisplayname ListFsuCollectionsSortByEnum = "displayName"
)

var mappingListFsuCollectionsSortByEnum = map[string]ListFsuCollectionsSortByEnum{
	"timeCreated": ListFsuCollectionsSortByTimecreated,
	"displayName": ListFsuCollectionsSortByDisplayname,
}

var mappingListFsuCollectionsSortByEnumLowerCase = map[string]ListFsuCollectionsSortByEnum{
	"timecreated": ListFsuCollectionsSortByTimecreated,
	"displayname": ListFsuCollectionsSortByDisplayname,
}

// GetListFsuCollectionsSortByEnumValues Enumerates the set of values for ListFsuCollectionsSortByEnum
func GetListFsuCollectionsSortByEnumValues() []ListFsuCollectionsSortByEnum {
	values := make([]ListFsuCollectionsSortByEnum, 0)
	for _, v := range mappingListFsuCollectionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuCollectionsSortByEnumStringValues Enumerates the set of values in String for ListFsuCollectionsSortByEnum
func GetListFsuCollectionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListFsuCollectionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuCollectionsSortByEnum(val string) (ListFsuCollectionsSortByEnum, bool) {
	enum, ok := mappingListFsuCollectionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
