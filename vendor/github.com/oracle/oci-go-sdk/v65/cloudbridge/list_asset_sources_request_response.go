// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAssetSourcesRequest wrapper for the ListAssetSources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListAssetSources.go.html to see an example of how to use ListAssetSourcesRequest.
type ListAssetSourcesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the asset source.
	AssetSourceId *string `mandatory:"false" contributesTo:"query" name:"assetSourceId"`

	// The field to sort by. Only one sort order may be provided. By default, the timeCreated is in descending order and displayName is in ascending order.
	SortBy ListAssetSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The current state of the asset source.
	LifecycleState ListAssetSourcesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAssetSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssetSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssetSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssetSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssetSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAssetSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAssetSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAssetSourcesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssetSourcesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAssetSourcesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssetSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAssetSourcesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAssetSourcesResponse wrapper for the ListAssetSources operation
type ListAssetSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssetSourceCollection instances
	AssetSourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAssetSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssetSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssetSourcesSortByEnum Enum with underlying type: string
type ListAssetSourcesSortByEnum string

// Set of constants representing the allowable values for ListAssetSourcesSortByEnum
const (
	ListAssetSourcesSortByTimecreated ListAssetSourcesSortByEnum = "timeCreated"
	ListAssetSourcesSortByDisplayname ListAssetSourcesSortByEnum = "displayName"
)

var mappingListAssetSourcesSortByEnum = map[string]ListAssetSourcesSortByEnum{
	"timeCreated": ListAssetSourcesSortByTimecreated,
	"displayName": ListAssetSourcesSortByDisplayname,
}

var mappingListAssetSourcesSortByEnumLowerCase = map[string]ListAssetSourcesSortByEnum{
	"timecreated": ListAssetSourcesSortByTimecreated,
	"displayname": ListAssetSourcesSortByDisplayname,
}

// GetListAssetSourcesSortByEnumValues Enumerates the set of values for ListAssetSourcesSortByEnum
func GetListAssetSourcesSortByEnumValues() []ListAssetSourcesSortByEnum {
	values := make([]ListAssetSourcesSortByEnum, 0)
	for _, v := range mappingListAssetSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssetSourcesSortByEnumStringValues Enumerates the set of values in String for ListAssetSourcesSortByEnum
func GetListAssetSourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAssetSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssetSourcesSortByEnum(val string) (ListAssetSourcesSortByEnum, bool) {
	enum, ok := mappingListAssetSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssetSourcesLifecycleStateEnum Enum with underlying type: string
type ListAssetSourcesLifecycleStateEnum string

// Set of constants representing the allowable values for ListAssetSourcesLifecycleStateEnum
const (
	ListAssetSourcesLifecycleStateCreating       ListAssetSourcesLifecycleStateEnum = "CREATING"
	ListAssetSourcesLifecycleStateActive         ListAssetSourcesLifecycleStateEnum = "ACTIVE"
	ListAssetSourcesLifecycleStateDeleting       ListAssetSourcesLifecycleStateEnum = "DELETING"
	ListAssetSourcesLifecycleStateDeleted        ListAssetSourcesLifecycleStateEnum = "DELETED"
	ListAssetSourcesLifecycleStateFailed         ListAssetSourcesLifecycleStateEnum = "FAILED"
	ListAssetSourcesLifecycleStateUpdating       ListAssetSourcesLifecycleStateEnum = "UPDATING"
	ListAssetSourcesLifecycleStateNeedsAttention ListAssetSourcesLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListAssetSourcesLifecycleStateEnum = map[string]ListAssetSourcesLifecycleStateEnum{
	"CREATING":        ListAssetSourcesLifecycleStateCreating,
	"ACTIVE":          ListAssetSourcesLifecycleStateActive,
	"DELETING":        ListAssetSourcesLifecycleStateDeleting,
	"DELETED":         ListAssetSourcesLifecycleStateDeleted,
	"FAILED":          ListAssetSourcesLifecycleStateFailed,
	"UPDATING":        ListAssetSourcesLifecycleStateUpdating,
	"NEEDS_ATTENTION": ListAssetSourcesLifecycleStateNeedsAttention,
}

var mappingListAssetSourcesLifecycleStateEnumLowerCase = map[string]ListAssetSourcesLifecycleStateEnum{
	"creating":        ListAssetSourcesLifecycleStateCreating,
	"active":          ListAssetSourcesLifecycleStateActive,
	"deleting":        ListAssetSourcesLifecycleStateDeleting,
	"deleted":         ListAssetSourcesLifecycleStateDeleted,
	"failed":          ListAssetSourcesLifecycleStateFailed,
	"updating":        ListAssetSourcesLifecycleStateUpdating,
	"needs_attention": ListAssetSourcesLifecycleStateNeedsAttention,
}

// GetListAssetSourcesLifecycleStateEnumValues Enumerates the set of values for ListAssetSourcesLifecycleStateEnum
func GetListAssetSourcesLifecycleStateEnumValues() []ListAssetSourcesLifecycleStateEnum {
	values := make([]ListAssetSourcesLifecycleStateEnum, 0)
	for _, v := range mappingListAssetSourcesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssetSourcesLifecycleStateEnumStringValues Enumerates the set of values in String for ListAssetSourcesLifecycleStateEnum
func GetListAssetSourcesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListAssetSourcesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssetSourcesLifecycleStateEnum(val string) (ListAssetSourcesLifecycleStateEnum, bool) {
	enum, ok := mappingListAssetSourcesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssetSourcesSortOrderEnum Enum with underlying type: string
type ListAssetSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListAssetSourcesSortOrderEnum
const (
	ListAssetSourcesSortOrderAsc  ListAssetSourcesSortOrderEnum = "ASC"
	ListAssetSourcesSortOrderDesc ListAssetSourcesSortOrderEnum = "DESC"
)

var mappingListAssetSourcesSortOrderEnum = map[string]ListAssetSourcesSortOrderEnum{
	"ASC":  ListAssetSourcesSortOrderAsc,
	"DESC": ListAssetSourcesSortOrderDesc,
}

var mappingListAssetSourcesSortOrderEnumLowerCase = map[string]ListAssetSourcesSortOrderEnum{
	"asc":  ListAssetSourcesSortOrderAsc,
	"desc": ListAssetSourcesSortOrderDesc,
}

// GetListAssetSourcesSortOrderEnumValues Enumerates the set of values for ListAssetSourcesSortOrderEnum
func GetListAssetSourcesSortOrderEnumValues() []ListAssetSourcesSortOrderEnum {
	values := make([]ListAssetSourcesSortOrderEnum, 0)
	for _, v := range mappingListAssetSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssetSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListAssetSourcesSortOrderEnum
func GetListAssetSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAssetSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssetSourcesSortOrderEnum(val string) (ListAssetSourcesSortOrderEnum, bool) {
	enum, ok := mappingListAssetSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
