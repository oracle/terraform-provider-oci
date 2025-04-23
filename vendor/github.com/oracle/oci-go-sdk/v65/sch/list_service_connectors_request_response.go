// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package sch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListServiceConnectorsRequest wrapper for the ListServiceConnectors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/sch/ListServiceConnectors.go.html to see an example of how to use ListServiceConnectorsRequest.
type ListServiceConnectorsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for this request.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state.
	// Example: `ACTIVE`
	LifecycleState ListServiceConnectorsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	// Example: `example_service_connector`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return
	// in a paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListServiceConnectorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for `timeCreated` is descending.
	// Default order for `displayName` is ascending. If no value is specified `timeCreated` is default.
	SortBy ListServiceConnectorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListServiceConnectorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListServiceConnectorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListServiceConnectorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListServiceConnectorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListServiceConnectorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListServiceConnectorsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListServiceConnectorsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServiceConnectorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListServiceConnectorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServiceConnectorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListServiceConnectorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListServiceConnectorsResponse wrapper for the ListServiceConnectors operation
type ListServiceConnectorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ServiceConnectorCollection instances
	ServiceConnectorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain. For important details about
	// how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination.  When this header appears in the response,
	// previous pages of results exist. For important details about
	// how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListServiceConnectorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListServiceConnectorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListServiceConnectorsLifecycleStateEnum Enum with underlying type: string
type ListServiceConnectorsLifecycleStateEnum string

// Set of constants representing the allowable values for ListServiceConnectorsLifecycleStateEnum
const (
	ListServiceConnectorsLifecycleStateCreating       ListServiceConnectorsLifecycleStateEnum = "CREATING"
	ListServiceConnectorsLifecycleStateUpdating       ListServiceConnectorsLifecycleStateEnum = "UPDATING"
	ListServiceConnectorsLifecycleStateActive         ListServiceConnectorsLifecycleStateEnum = "ACTIVE"
	ListServiceConnectorsLifecycleStateInactive       ListServiceConnectorsLifecycleStateEnum = "INACTIVE"
	ListServiceConnectorsLifecycleStateNeedsAttention ListServiceConnectorsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListServiceConnectorsLifecycleStateDeleting       ListServiceConnectorsLifecycleStateEnum = "DELETING"
	ListServiceConnectorsLifecycleStateDeleted        ListServiceConnectorsLifecycleStateEnum = "DELETED"
	ListServiceConnectorsLifecycleStateFailed         ListServiceConnectorsLifecycleStateEnum = "FAILED"
)

var mappingListServiceConnectorsLifecycleStateEnum = map[string]ListServiceConnectorsLifecycleStateEnum{
	"CREATING":        ListServiceConnectorsLifecycleStateCreating,
	"UPDATING":        ListServiceConnectorsLifecycleStateUpdating,
	"ACTIVE":          ListServiceConnectorsLifecycleStateActive,
	"INACTIVE":        ListServiceConnectorsLifecycleStateInactive,
	"NEEDS_ATTENTION": ListServiceConnectorsLifecycleStateNeedsAttention,
	"DELETING":        ListServiceConnectorsLifecycleStateDeleting,
	"DELETED":         ListServiceConnectorsLifecycleStateDeleted,
	"FAILED":          ListServiceConnectorsLifecycleStateFailed,
}

var mappingListServiceConnectorsLifecycleStateEnumLowerCase = map[string]ListServiceConnectorsLifecycleStateEnum{
	"creating":        ListServiceConnectorsLifecycleStateCreating,
	"updating":        ListServiceConnectorsLifecycleStateUpdating,
	"active":          ListServiceConnectorsLifecycleStateActive,
	"inactive":        ListServiceConnectorsLifecycleStateInactive,
	"needs_attention": ListServiceConnectorsLifecycleStateNeedsAttention,
	"deleting":        ListServiceConnectorsLifecycleStateDeleting,
	"deleted":         ListServiceConnectorsLifecycleStateDeleted,
	"failed":          ListServiceConnectorsLifecycleStateFailed,
}

// GetListServiceConnectorsLifecycleStateEnumValues Enumerates the set of values for ListServiceConnectorsLifecycleStateEnum
func GetListServiceConnectorsLifecycleStateEnumValues() []ListServiceConnectorsLifecycleStateEnum {
	values := make([]ListServiceConnectorsLifecycleStateEnum, 0)
	for _, v := range mappingListServiceConnectorsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceConnectorsLifecycleStateEnumStringValues Enumerates the set of values in String for ListServiceConnectorsLifecycleStateEnum
func GetListServiceConnectorsLifecycleStateEnumStringValues() []string {
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

// GetMappingListServiceConnectorsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceConnectorsLifecycleStateEnum(val string) (ListServiceConnectorsLifecycleStateEnum, bool) {
	enum, ok := mappingListServiceConnectorsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListServiceConnectorsSortOrderEnum Enum with underlying type: string
type ListServiceConnectorsSortOrderEnum string

// Set of constants representing the allowable values for ListServiceConnectorsSortOrderEnum
const (
	ListServiceConnectorsSortOrderAsc  ListServiceConnectorsSortOrderEnum = "ASC"
	ListServiceConnectorsSortOrderDesc ListServiceConnectorsSortOrderEnum = "DESC"
)

var mappingListServiceConnectorsSortOrderEnum = map[string]ListServiceConnectorsSortOrderEnum{
	"ASC":  ListServiceConnectorsSortOrderAsc,
	"DESC": ListServiceConnectorsSortOrderDesc,
}

var mappingListServiceConnectorsSortOrderEnumLowerCase = map[string]ListServiceConnectorsSortOrderEnum{
	"asc":  ListServiceConnectorsSortOrderAsc,
	"desc": ListServiceConnectorsSortOrderDesc,
}

// GetListServiceConnectorsSortOrderEnumValues Enumerates the set of values for ListServiceConnectorsSortOrderEnum
func GetListServiceConnectorsSortOrderEnumValues() []ListServiceConnectorsSortOrderEnum {
	values := make([]ListServiceConnectorsSortOrderEnum, 0)
	for _, v := range mappingListServiceConnectorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceConnectorsSortOrderEnumStringValues Enumerates the set of values in String for ListServiceConnectorsSortOrderEnum
func GetListServiceConnectorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListServiceConnectorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceConnectorsSortOrderEnum(val string) (ListServiceConnectorsSortOrderEnum, bool) {
	enum, ok := mappingListServiceConnectorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListServiceConnectorsSortByEnum Enum with underlying type: string
type ListServiceConnectorsSortByEnum string

// Set of constants representing the allowable values for ListServiceConnectorsSortByEnum
const (
	ListServiceConnectorsSortByTimecreated ListServiceConnectorsSortByEnum = "timeCreated"
	ListServiceConnectorsSortByDisplayname ListServiceConnectorsSortByEnum = "displayName"
)

var mappingListServiceConnectorsSortByEnum = map[string]ListServiceConnectorsSortByEnum{
	"timeCreated": ListServiceConnectorsSortByTimecreated,
	"displayName": ListServiceConnectorsSortByDisplayname,
}

var mappingListServiceConnectorsSortByEnumLowerCase = map[string]ListServiceConnectorsSortByEnum{
	"timecreated": ListServiceConnectorsSortByTimecreated,
	"displayname": ListServiceConnectorsSortByDisplayname,
}

// GetListServiceConnectorsSortByEnumValues Enumerates the set of values for ListServiceConnectorsSortByEnum
func GetListServiceConnectorsSortByEnumValues() []ListServiceConnectorsSortByEnum {
	values := make([]ListServiceConnectorsSortByEnum, 0)
	for _, v := range mappingListServiceConnectorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceConnectorsSortByEnumStringValues Enumerates the set of values in String for ListServiceConnectorsSortByEnum
func GetListServiceConnectorsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListServiceConnectorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceConnectorsSortByEnum(val string) (ListServiceConnectorsSortByEnum, bool) {
	enum, ok := mappingListServiceConnectorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
