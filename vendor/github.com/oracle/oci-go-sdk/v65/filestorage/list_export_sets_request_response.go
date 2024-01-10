// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExportSetsRequest wrapper for the ListExportSets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListExportSets.go.html to see an example of how to use ListExportSetsRequest.
type ListExportSetsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" contributesTo:"query" name:"availabilityDomain"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Example: `My resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListExportSetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filter results by OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for
	// the resouce type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The field to sort by. You can provide either value, but not both.
	// By default, when you sort by time created, results are shown
	// in descending order. When you sort by display name, results are
	// shown in ascending order.
	SortBy ListExportSetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending. The default order is 'desc'
	// except for numeric values.
	SortOrder ListExportSetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExportSetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExportSetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExportSetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExportSetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExportSetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExportSetsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListExportSetsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExportSetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExportSetsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExportSetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExportSetsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExportSetsResponse wrapper for the ListExportSets operation
type ListExportSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExportSetSummary instances
	Items []ExportSetSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListExportSetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExportSetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExportSetsLifecycleStateEnum Enum with underlying type: string
type ListExportSetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListExportSetsLifecycleStateEnum
const (
	ListExportSetsLifecycleStateCreating ListExportSetsLifecycleStateEnum = "CREATING"
	ListExportSetsLifecycleStateActive   ListExportSetsLifecycleStateEnum = "ACTIVE"
	ListExportSetsLifecycleStateDeleting ListExportSetsLifecycleStateEnum = "DELETING"
	ListExportSetsLifecycleStateDeleted  ListExportSetsLifecycleStateEnum = "DELETED"
	ListExportSetsLifecycleStateFailed   ListExportSetsLifecycleStateEnum = "FAILED"
)

var mappingListExportSetsLifecycleStateEnum = map[string]ListExportSetsLifecycleStateEnum{
	"CREATING": ListExportSetsLifecycleStateCreating,
	"ACTIVE":   ListExportSetsLifecycleStateActive,
	"DELETING": ListExportSetsLifecycleStateDeleting,
	"DELETED":  ListExportSetsLifecycleStateDeleted,
	"FAILED":   ListExportSetsLifecycleStateFailed,
}

var mappingListExportSetsLifecycleStateEnumLowerCase = map[string]ListExportSetsLifecycleStateEnum{
	"creating": ListExportSetsLifecycleStateCreating,
	"active":   ListExportSetsLifecycleStateActive,
	"deleting": ListExportSetsLifecycleStateDeleting,
	"deleted":  ListExportSetsLifecycleStateDeleted,
	"failed":   ListExportSetsLifecycleStateFailed,
}

// GetListExportSetsLifecycleStateEnumValues Enumerates the set of values for ListExportSetsLifecycleStateEnum
func GetListExportSetsLifecycleStateEnumValues() []ListExportSetsLifecycleStateEnum {
	values := make([]ListExportSetsLifecycleStateEnum, 0)
	for _, v := range mappingListExportSetsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListExportSetsLifecycleStateEnumStringValues Enumerates the set of values in String for ListExportSetsLifecycleStateEnum
func GetListExportSetsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListExportSetsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExportSetsLifecycleStateEnum(val string) (ListExportSetsLifecycleStateEnum, bool) {
	enum, ok := mappingListExportSetsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExportSetsSortByEnum Enum with underlying type: string
type ListExportSetsSortByEnum string

// Set of constants representing the allowable values for ListExportSetsSortByEnum
const (
	ListExportSetsSortByTimecreated ListExportSetsSortByEnum = "TIMECREATED"
	ListExportSetsSortByDisplayname ListExportSetsSortByEnum = "DISPLAYNAME"
)

var mappingListExportSetsSortByEnum = map[string]ListExportSetsSortByEnum{
	"TIMECREATED": ListExportSetsSortByTimecreated,
	"DISPLAYNAME": ListExportSetsSortByDisplayname,
}

var mappingListExportSetsSortByEnumLowerCase = map[string]ListExportSetsSortByEnum{
	"timecreated": ListExportSetsSortByTimecreated,
	"displayname": ListExportSetsSortByDisplayname,
}

// GetListExportSetsSortByEnumValues Enumerates the set of values for ListExportSetsSortByEnum
func GetListExportSetsSortByEnumValues() []ListExportSetsSortByEnum {
	values := make([]ListExportSetsSortByEnum, 0)
	for _, v := range mappingListExportSetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExportSetsSortByEnumStringValues Enumerates the set of values in String for ListExportSetsSortByEnum
func GetListExportSetsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListExportSetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExportSetsSortByEnum(val string) (ListExportSetsSortByEnum, bool) {
	enum, ok := mappingListExportSetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExportSetsSortOrderEnum Enum with underlying type: string
type ListExportSetsSortOrderEnum string

// Set of constants representing the allowable values for ListExportSetsSortOrderEnum
const (
	ListExportSetsSortOrderAsc  ListExportSetsSortOrderEnum = "ASC"
	ListExportSetsSortOrderDesc ListExportSetsSortOrderEnum = "DESC"
)

var mappingListExportSetsSortOrderEnum = map[string]ListExportSetsSortOrderEnum{
	"ASC":  ListExportSetsSortOrderAsc,
	"DESC": ListExportSetsSortOrderDesc,
}

var mappingListExportSetsSortOrderEnumLowerCase = map[string]ListExportSetsSortOrderEnum{
	"asc":  ListExportSetsSortOrderAsc,
	"desc": ListExportSetsSortOrderDesc,
}

// GetListExportSetsSortOrderEnumValues Enumerates the set of values for ListExportSetsSortOrderEnum
func GetListExportSetsSortOrderEnumValues() []ListExportSetsSortOrderEnum {
	values := make([]ListExportSetsSortOrderEnum, 0)
	for _, v := range mappingListExportSetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExportSetsSortOrderEnumStringValues Enumerates the set of values in String for ListExportSetsSortOrderEnum
func GetListExportSetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExportSetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExportSetsSortOrderEnum(val string) (ListExportSetsSortOrderEnum, bool) {
	enum, ok := mappingListExportSetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
