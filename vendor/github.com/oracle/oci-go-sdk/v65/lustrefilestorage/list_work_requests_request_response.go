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

// ListWorkRequestsRequest wrapper for the ListWorkRequests operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/ListWorkRequests.go.html to see an example of how to use ListWorkRequestsRequest.
type ListWorkRequestsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the asynchronous work request.
	WorkRequestId *string `mandatory:"false" contributesTo:"query" name:"workRequestId"`

	// A filter to return only the resources that match the given lifecycle state.
	Status ListWorkRequestsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource affected by the work request.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for `timeAccepted` is descending.
	SortBy ListWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWorkRequestsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListWorkRequestsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWorkRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWorkRequestsResponse wrapper for the ListWorkRequests operation
type ListWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestSummaryCollection instances
	WorkRequestSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestsStatusEnum Enum with underlying type: string
type ListWorkRequestsStatusEnum string

// Set of constants representing the allowable values for ListWorkRequestsStatusEnum
const (
	ListWorkRequestsStatusAccepted       ListWorkRequestsStatusEnum = "ACCEPTED"
	ListWorkRequestsStatusInProgress     ListWorkRequestsStatusEnum = "IN_PROGRESS"
	ListWorkRequestsStatusWaiting        ListWorkRequestsStatusEnum = "WAITING"
	ListWorkRequestsStatusNeedsAttention ListWorkRequestsStatusEnum = "NEEDS_ATTENTION"
	ListWorkRequestsStatusFailed         ListWorkRequestsStatusEnum = "FAILED"
	ListWorkRequestsStatusSucceeded      ListWorkRequestsStatusEnum = "SUCCEEDED"
	ListWorkRequestsStatusCanceling      ListWorkRequestsStatusEnum = "CANCELING"
	ListWorkRequestsStatusCanceled       ListWorkRequestsStatusEnum = "CANCELED"
)

var mappingListWorkRequestsStatusEnum = map[string]ListWorkRequestsStatusEnum{
	"ACCEPTED":        ListWorkRequestsStatusAccepted,
	"IN_PROGRESS":     ListWorkRequestsStatusInProgress,
	"WAITING":         ListWorkRequestsStatusWaiting,
	"NEEDS_ATTENTION": ListWorkRequestsStatusNeedsAttention,
	"FAILED":          ListWorkRequestsStatusFailed,
	"SUCCEEDED":       ListWorkRequestsStatusSucceeded,
	"CANCELING":       ListWorkRequestsStatusCanceling,
	"CANCELED":        ListWorkRequestsStatusCanceled,
}

var mappingListWorkRequestsStatusEnumLowerCase = map[string]ListWorkRequestsStatusEnum{
	"accepted":        ListWorkRequestsStatusAccepted,
	"in_progress":     ListWorkRequestsStatusInProgress,
	"waiting":         ListWorkRequestsStatusWaiting,
	"needs_attention": ListWorkRequestsStatusNeedsAttention,
	"failed":          ListWorkRequestsStatusFailed,
	"succeeded":       ListWorkRequestsStatusSucceeded,
	"canceling":       ListWorkRequestsStatusCanceling,
	"canceled":        ListWorkRequestsStatusCanceled,
}

// GetListWorkRequestsStatusEnumValues Enumerates the set of values for ListWorkRequestsStatusEnum
func GetListWorkRequestsStatusEnumValues() []ListWorkRequestsStatusEnum {
	values := make([]ListWorkRequestsStatusEnum, 0)
	for _, v := range mappingListWorkRequestsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsStatusEnumStringValues Enumerates the set of values in String for ListWorkRequestsStatusEnum
func GetListWorkRequestsStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"NEEDS_ATTENTION",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingListWorkRequestsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsStatusEnum(val string) (ListWorkRequestsStatusEnum, bool) {
	enum, ok := mappingListWorkRequestsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestsSortOrderEnum Enum with underlying type: string
type ListWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortOrderEnum
const (
	ListWorkRequestsSortOrderAsc  ListWorkRequestsSortOrderEnum = "ASC"
	ListWorkRequestsSortOrderDesc ListWorkRequestsSortOrderEnum = "DESC"
)

var mappingListWorkRequestsSortOrderEnum = map[string]ListWorkRequestsSortOrderEnum{
	"ASC":  ListWorkRequestsSortOrderAsc,
	"DESC": ListWorkRequestsSortOrderDesc,
}

var mappingListWorkRequestsSortOrderEnumLowerCase = map[string]ListWorkRequestsSortOrderEnum{
	"asc":  ListWorkRequestsSortOrderAsc,
	"desc": ListWorkRequestsSortOrderDesc,
}

// GetListWorkRequestsSortOrderEnumValues Enumerates the set of values for ListWorkRequestsSortOrderEnum
func GetListWorkRequestsSortOrderEnumValues() []ListWorkRequestsSortOrderEnum {
	values := make([]ListWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListWorkRequestsSortOrderEnum
func GetListWorkRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWorkRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsSortOrderEnum(val string) (ListWorkRequestsSortOrderEnum, bool) {
	enum, ok := mappingListWorkRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestsSortByEnum Enum with underlying type: string
type ListWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortByEnum
const (
	ListWorkRequestsSortByTimeaccepted ListWorkRequestsSortByEnum = "timeAccepted"
)

var mappingListWorkRequestsSortByEnum = map[string]ListWorkRequestsSortByEnum{
	"timeAccepted": ListWorkRequestsSortByTimeaccepted,
}

var mappingListWorkRequestsSortByEnumLowerCase = map[string]ListWorkRequestsSortByEnum{
	"timeaccepted": ListWorkRequestsSortByTimeaccepted,
}

// GetListWorkRequestsSortByEnumValues Enumerates the set of values for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumValues() []ListWorkRequestsSortByEnum {
	values := make([]ListWorkRequestsSortByEnum, 0)
	for _, v := range mappingListWorkRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsSortByEnumStringValues Enumerates the set of values in String for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumStringValues() []string {
	return []string{
		"timeAccepted",
	}
}

// GetMappingListWorkRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsSortByEnum(val string) (ListWorkRequestsSortByEnum, bool) {
	enum, ok := mappingListWorkRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
