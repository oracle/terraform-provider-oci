// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExportsRequest wrapper for the ListExports operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListExports.go.html to see an example of how to use ListExportsRequest.
type ListExportsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 4096 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the export set.
	ExportSetId *string `mandatory:"false" contributesTo:"query" name:"exportSetId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system.
	FileSystemId *string `mandatory:"false" contributesTo:"query" name:"fileSystemId"`

	// Filter results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListExportsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filter results by OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for
	// the resouce type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The field to sort by. You can provide either value, but not both.
	// By default, when you sort by time created, results are shown
	// in descending order. When you sort by path, results are
	// shown in ascending alphanumeric order.
	SortBy ListExportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending. The default order is 'desc'
	// except for numeric values.
	SortOrder ListExportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExportsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListExportsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExportsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExportsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExportsResponse wrapper for the ListExports operation
type ListExportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExportSummary instances
	Items []ExportSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListExportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExportsLifecycleStateEnum Enum with underlying type: string
type ListExportsLifecycleStateEnum string

// Set of constants representing the allowable values for ListExportsLifecycleStateEnum
const (
	ListExportsLifecycleStateCreating ListExportsLifecycleStateEnum = "CREATING"
	ListExportsLifecycleStateActive   ListExportsLifecycleStateEnum = "ACTIVE"
	ListExportsLifecycleStateUpdating ListExportsLifecycleStateEnum = "UPDATING"
	ListExportsLifecycleStateDeleting ListExportsLifecycleStateEnum = "DELETING"
	ListExportsLifecycleStateDeleted  ListExportsLifecycleStateEnum = "DELETED"
	ListExportsLifecycleStateFailed   ListExportsLifecycleStateEnum = "FAILED"
)

var mappingListExportsLifecycleStateEnum = map[string]ListExportsLifecycleStateEnum{
	"CREATING": ListExportsLifecycleStateCreating,
	"ACTIVE":   ListExportsLifecycleStateActive,
	"UPDATING": ListExportsLifecycleStateUpdating,
	"DELETING": ListExportsLifecycleStateDeleting,
	"DELETED":  ListExportsLifecycleStateDeleted,
	"FAILED":   ListExportsLifecycleStateFailed,
}

var mappingListExportsLifecycleStateEnumLowerCase = map[string]ListExportsLifecycleStateEnum{
	"creating": ListExportsLifecycleStateCreating,
	"active":   ListExportsLifecycleStateActive,
	"updating": ListExportsLifecycleStateUpdating,
	"deleting": ListExportsLifecycleStateDeleting,
	"deleted":  ListExportsLifecycleStateDeleted,
	"failed":   ListExportsLifecycleStateFailed,
}

// GetListExportsLifecycleStateEnumValues Enumerates the set of values for ListExportsLifecycleStateEnum
func GetListExportsLifecycleStateEnumValues() []ListExportsLifecycleStateEnum {
	values := make([]ListExportsLifecycleStateEnum, 0)
	for _, v := range mappingListExportsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListExportsLifecycleStateEnumStringValues Enumerates the set of values in String for ListExportsLifecycleStateEnum
func GetListExportsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListExportsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExportsLifecycleStateEnum(val string) (ListExportsLifecycleStateEnum, bool) {
	enum, ok := mappingListExportsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExportsSortByEnum Enum with underlying type: string
type ListExportsSortByEnum string

// Set of constants representing the allowable values for ListExportsSortByEnum
const (
	ListExportsSortByTimecreated ListExportsSortByEnum = "TIMECREATED"
	ListExportsSortByPath        ListExportsSortByEnum = "PATH"
)

var mappingListExportsSortByEnum = map[string]ListExportsSortByEnum{
	"TIMECREATED": ListExportsSortByTimecreated,
	"PATH":        ListExportsSortByPath,
}

var mappingListExportsSortByEnumLowerCase = map[string]ListExportsSortByEnum{
	"timecreated": ListExportsSortByTimecreated,
	"path":        ListExportsSortByPath,
}

// GetListExportsSortByEnumValues Enumerates the set of values for ListExportsSortByEnum
func GetListExportsSortByEnumValues() []ListExportsSortByEnum {
	values := make([]ListExportsSortByEnum, 0)
	for _, v := range mappingListExportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExportsSortByEnumStringValues Enumerates the set of values in String for ListExportsSortByEnum
func GetListExportsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"PATH",
	}
}

// GetMappingListExportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExportsSortByEnum(val string) (ListExportsSortByEnum, bool) {
	enum, ok := mappingListExportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExportsSortOrderEnum Enum with underlying type: string
type ListExportsSortOrderEnum string

// Set of constants representing the allowable values for ListExportsSortOrderEnum
const (
	ListExportsSortOrderAsc  ListExportsSortOrderEnum = "ASC"
	ListExportsSortOrderDesc ListExportsSortOrderEnum = "DESC"
)

var mappingListExportsSortOrderEnum = map[string]ListExportsSortOrderEnum{
	"ASC":  ListExportsSortOrderAsc,
	"DESC": ListExportsSortOrderDesc,
}

var mappingListExportsSortOrderEnumLowerCase = map[string]ListExportsSortOrderEnum{
	"asc":  ListExportsSortOrderAsc,
	"desc": ListExportsSortOrderDesc,
}

// GetListExportsSortOrderEnumValues Enumerates the set of values for ListExportsSortOrderEnum
func GetListExportsSortOrderEnumValues() []ListExportsSortOrderEnum {
	values := make([]ListExportsSortOrderEnum, 0)
	for _, v := range mappingListExportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExportsSortOrderEnumStringValues Enumerates the set of values in String for ListExportsSortOrderEnum
func GetListExportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExportsSortOrderEnum(val string) (ListExportsSortOrderEnum, bool) {
	enum, ok := mappingListExportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
