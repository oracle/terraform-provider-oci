// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFleetsRequest wrapper for the ListFleets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListFleets.go.html to see an example of how to use ListFleetsRequest.
type ListFleetsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The ID of the Fleet.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The state of the lifecycle.
	LifecycleState ListFleetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListFleetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort Fleets. Only one sort order may be provided.
	// Default order for _timeCreated_, _approximateJreCount_, _approximateInstallationCount_,
	// _approximateApplicationCount_ and _approximateManagedInstanceCount_  is **descending**.
	// Default order for _displayName_ is **ascending**.
	// If no value is specified _timeCreated_ is default.
	SortBy ListFleetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Filter the list with displayName contains the given value.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFleetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFleetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFleetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFleetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFleetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFleetsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListFleetsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFleetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFleetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFleetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFleetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFleetsResponse wrapper for the ListFleets operation
type ListFleetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FleetCollection instances
	FleetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFleetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFleetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFleetsLifecycleStateEnum Enum with underlying type: string
type ListFleetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListFleetsLifecycleStateEnum
const (
	ListFleetsLifecycleStateActive         ListFleetsLifecycleStateEnum = "ACTIVE"
	ListFleetsLifecycleStateCreating       ListFleetsLifecycleStateEnum = "CREATING"
	ListFleetsLifecycleStateDeleted        ListFleetsLifecycleStateEnum = "DELETED"
	ListFleetsLifecycleStateDeleting       ListFleetsLifecycleStateEnum = "DELETING"
	ListFleetsLifecycleStateFailed         ListFleetsLifecycleStateEnum = "FAILED"
	ListFleetsLifecycleStateNeedsAttention ListFleetsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListFleetsLifecycleStateUpdating       ListFleetsLifecycleStateEnum = "UPDATING"
)

var mappingListFleetsLifecycleStateEnum = map[string]ListFleetsLifecycleStateEnum{
	"ACTIVE":          ListFleetsLifecycleStateActive,
	"CREATING":        ListFleetsLifecycleStateCreating,
	"DELETED":         ListFleetsLifecycleStateDeleted,
	"DELETING":        ListFleetsLifecycleStateDeleting,
	"FAILED":          ListFleetsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListFleetsLifecycleStateNeedsAttention,
	"UPDATING":        ListFleetsLifecycleStateUpdating,
}

var mappingListFleetsLifecycleStateEnumLowerCase = map[string]ListFleetsLifecycleStateEnum{
	"active":          ListFleetsLifecycleStateActive,
	"creating":        ListFleetsLifecycleStateCreating,
	"deleted":         ListFleetsLifecycleStateDeleted,
	"deleting":        ListFleetsLifecycleStateDeleting,
	"failed":          ListFleetsLifecycleStateFailed,
	"needs_attention": ListFleetsLifecycleStateNeedsAttention,
	"updating":        ListFleetsLifecycleStateUpdating,
}

// GetListFleetsLifecycleStateEnumValues Enumerates the set of values for ListFleetsLifecycleStateEnum
func GetListFleetsLifecycleStateEnumValues() []ListFleetsLifecycleStateEnum {
	values := make([]ListFleetsLifecycleStateEnum, 0)
	for _, v := range mappingListFleetsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetsLifecycleStateEnumStringValues Enumerates the set of values in String for ListFleetsLifecycleStateEnum
func GetListFleetsLifecycleStateEnumStringValues() []string {
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

// GetMappingListFleetsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetsLifecycleStateEnum(val string) (ListFleetsLifecycleStateEnum, bool) {
	enum, ok := mappingListFleetsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFleetsSortOrderEnum Enum with underlying type: string
type ListFleetsSortOrderEnum string

// Set of constants representing the allowable values for ListFleetsSortOrderEnum
const (
	ListFleetsSortOrderAsc  ListFleetsSortOrderEnum = "ASC"
	ListFleetsSortOrderDesc ListFleetsSortOrderEnum = "DESC"
)

var mappingListFleetsSortOrderEnum = map[string]ListFleetsSortOrderEnum{
	"ASC":  ListFleetsSortOrderAsc,
	"DESC": ListFleetsSortOrderDesc,
}

var mappingListFleetsSortOrderEnumLowerCase = map[string]ListFleetsSortOrderEnum{
	"asc":  ListFleetsSortOrderAsc,
	"desc": ListFleetsSortOrderDesc,
}

// GetListFleetsSortOrderEnumValues Enumerates the set of values for ListFleetsSortOrderEnum
func GetListFleetsSortOrderEnumValues() []ListFleetsSortOrderEnum {
	values := make([]ListFleetsSortOrderEnum, 0)
	for _, v := range mappingListFleetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetsSortOrderEnumStringValues Enumerates the set of values in String for ListFleetsSortOrderEnum
func GetListFleetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFleetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetsSortOrderEnum(val string) (ListFleetsSortOrderEnum, bool) {
	enum, ok := mappingListFleetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFleetsSortByEnum Enum with underlying type: string
type ListFleetsSortByEnum string

// Set of constants representing the allowable values for ListFleetsSortByEnum
const (
	ListFleetsSortByDisplayname ListFleetsSortByEnum = "displayName"
	ListFleetsSortByTimecreated ListFleetsSortByEnum = "timeCreated"
)

var mappingListFleetsSortByEnum = map[string]ListFleetsSortByEnum{
	"displayName": ListFleetsSortByDisplayname,
	"timeCreated": ListFleetsSortByTimecreated,
}

var mappingListFleetsSortByEnumLowerCase = map[string]ListFleetsSortByEnum{
	"displayname": ListFleetsSortByDisplayname,
	"timecreated": ListFleetsSortByTimecreated,
}

// GetListFleetsSortByEnumValues Enumerates the set of values for ListFleetsSortByEnum
func GetListFleetsSortByEnumValues() []ListFleetsSortByEnum {
	values := make([]ListFleetsSortByEnum, 0)
	for _, v := range mappingListFleetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetsSortByEnumStringValues Enumerates the set of values in String for ListFleetsSortByEnum
func GetListFleetsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListFleetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetsSortByEnum(val string) (ListFleetsSortByEnum, bool) {
	enum, ok := mappingListFleetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
