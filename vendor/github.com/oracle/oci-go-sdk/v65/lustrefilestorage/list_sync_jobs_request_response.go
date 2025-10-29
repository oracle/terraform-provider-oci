// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSyncJobsRequest wrapper for the ListSyncJobs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/ListSyncJobs.go.html to see an example of how to use ListSyncJobsRequest.
type ListSyncJobsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Object Storage link.
	ObjectStorageLinkId *string `mandatory:"true" contributesTo:"path" name:"objectStorageLinkId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSyncJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the
	// `if-match` parameter to the value of the etag from a previous GET or POST response for
	// that resource. The resource will be updated or deleted only if the etag you provide
	// matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState SyncJobLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListSyncJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSyncJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSyncJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSyncJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSyncJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSyncJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSyncJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSyncJobsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSyncJobLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetSyncJobLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSyncJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSyncJobsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSyncJobsResponse wrapper for the ListSyncJobs operation
type ListSyncJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SyncJobCollection instances
	SyncJobCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSyncJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSyncJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSyncJobsSortOrderEnum Enum with underlying type: string
type ListSyncJobsSortOrderEnum string

// Set of constants representing the allowable values for ListSyncJobsSortOrderEnum
const (
	ListSyncJobsSortOrderAsc  ListSyncJobsSortOrderEnum = "ASC"
	ListSyncJobsSortOrderDesc ListSyncJobsSortOrderEnum = "DESC"
)

var mappingListSyncJobsSortOrderEnum = map[string]ListSyncJobsSortOrderEnum{
	"ASC":  ListSyncJobsSortOrderAsc,
	"DESC": ListSyncJobsSortOrderDesc,
}

var mappingListSyncJobsSortOrderEnumLowerCase = map[string]ListSyncJobsSortOrderEnum{
	"asc":  ListSyncJobsSortOrderAsc,
	"desc": ListSyncJobsSortOrderDesc,
}

// GetListSyncJobsSortOrderEnumValues Enumerates the set of values for ListSyncJobsSortOrderEnum
func GetListSyncJobsSortOrderEnumValues() []ListSyncJobsSortOrderEnum {
	values := make([]ListSyncJobsSortOrderEnum, 0)
	for _, v := range mappingListSyncJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSyncJobsSortOrderEnumStringValues Enumerates the set of values in String for ListSyncJobsSortOrderEnum
func GetListSyncJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSyncJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSyncJobsSortOrderEnum(val string) (ListSyncJobsSortOrderEnum, bool) {
	enum, ok := mappingListSyncJobsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSyncJobsSortByEnum Enum with underlying type: string
type ListSyncJobsSortByEnum string

// Set of constants representing the allowable values for ListSyncJobsSortByEnum
const (
	ListSyncJobsSortByTimecreated ListSyncJobsSortByEnum = "timeCreated"
)

var mappingListSyncJobsSortByEnum = map[string]ListSyncJobsSortByEnum{
	"timeCreated": ListSyncJobsSortByTimecreated,
}

var mappingListSyncJobsSortByEnumLowerCase = map[string]ListSyncJobsSortByEnum{
	"timecreated": ListSyncJobsSortByTimecreated,
}

// GetListSyncJobsSortByEnumValues Enumerates the set of values for ListSyncJobsSortByEnum
func GetListSyncJobsSortByEnumValues() []ListSyncJobsSortByEnum {
	values := make([]ListSyncJobsSortByEnum, 0)
	for _, v := range mappingListSyncJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSyncJobsSortByEnumStringValues Enumerates the set of values in String for ListSyncJobsSortByEnum
func GetListSyncJobsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListSyncJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSyncJobsSortByEnum(val string) (ListSyncJobsSortByEnum, bool) {
	enum, ok := mappingListSyncJobsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
