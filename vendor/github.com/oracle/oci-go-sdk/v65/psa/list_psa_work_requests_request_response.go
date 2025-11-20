// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package psa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPsaWorkRequestsRequest wrapper for the ListPsaWorkRequests operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/ListPsaWorkRequests.go.html to see an example of how to use ListPsaWorkRequestsRequest.
type ListPsaWorkRequestsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the asynchronous work request.
	WorkRequestId *string `mandatory:"false" contributesTo:"query" name:"workRequestId"`

	// A filter to return only the resources that match the given lifecycle state.
	Status ListPsaWorkRequestsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

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
	SortOrder ListPsaWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for `timeAccepted` is descending.
	SortBy ListPsaWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPsaWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPsaWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPsaWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPsaWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPsaWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPsaWorkRequestsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListPsaWorkRequestsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPsaWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPsaWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPsaWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPsaWorkRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPsaWorkRequestsResponse wrapper for the ListPsaWorkRequests operation
type ListPsaWorkRequestsResponse struct {

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

func (response ListPsaWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPsaWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPsaWorkRequestsStatusEnum Enum with underlying type: string
type ListPsaWorkRequestsStatusEnum string

// Set of constants representing the allowable values for ListPsaWorkRequestsStatusEnum
const (
	ListPsaWorkRequestsStatusAccepted       ListPsaWorkRequestsStatusEnum = "ACCEPTED"
	ListPsaWorkRequestsStatusInProgress     ListPsaWorkRequestsStatusEnum = "IN_PROGRESS"
	ListPsaWorkRequestsStatusWaiting        ListPsaWorkRequestsStatusEnum = "WAITING"
	ListPsaWorkRequestsStatusNeedsAttention ListPsaWorkRequestsStatusEnum = "NEEDS_ATTENTION"
	ListPsaWorkRequestsStatusFailed         ListPsaWorkRequestsStatusEnum = "FAILED"
	ListPsaWorkRequestsStatusSucceeded      ListPsaWorkRequestsStatusEnum = "SUCCEEDED"
	ListPsaWorkRequestsStatusCancelling     ListPsaWorkRequestsStatusEnum = "CANCELLING"
	ListPsaWorkRequestsStatusCancelled      ListPsaWorkRequestsStatusEnum = "CANCELLED"
)

var mappingListPsaWorkRequestsStatusEnum = map[string]ListPsaWorkRequestsStatusEnum{
	"ACCEPTED":        ListPsaWorkRequestsStatusAccepted,
	"IN_PROGRESS":     ListPsaWorkRequestsStatusInProgress,
	"WAITING":         ListPsaWorkRequestsStatusWaiting,
	"NEEDS_ATTENTION": ListPsaWorkRequestsStatusNeedsAttention,
	"FAILED":          ListPsaWorkRequestsStatusFailed,
	"SUCCEEDED":       ListPsaWorkRequestsStatusSucceeded,
	"CANCELLING":      ListPsaWorkRequestsStatusCancelling,
	"CANCELLED":       ListPsaWorkRequestsStatusCancelled,
}

var mappingListPsaWorkRequestsStatusEnumLowerCase = map[string]ListPsaWorkRequestsStatusEnum{
	"accepted":        ListPsaWorkRequestsStatusAccepted,
	"in_progress":     ListPsaWorkRequestsStatusInProgress,
	"waiting":         ListPsaWorkRequestsStatusWaiting,
	"needs_attention": ListPsaWorkRequestsStatusNeedsAttention,
	"failed":          ListPsaWorkRequestsStatusFailed,
	"succeeded":       ListPsaWorkRequestsStatusSucceeded,
	"cancelling":      ListPsaWorkRequestsStatusCancelling,
	"cancelled":       ListPsaWorkRequestsStatusCancelled,
}

// GetListPsaWorkRequestsStatusEnumValues Enumerates the set of values for ListPsaWorkRequestsStatusEnum
func GetListPsaWorkRequestsStatusEnumValues() []ListPsaWorkRequestsStatusEnum {
	values := make([]ListPsaWorkRequestsStatusEnum, 0)
	for _, v := range mappingListPsaWorkRequestsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListPsaWorkRequestsStatusEnumStringValues Enumerates the set of values in String for ListPsaWorkRequestsStatusEnum
func GetListPsaWorkRequestsStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"NEEDS_ATTENTION",
		"FAILED",
		"SUCCEEDED",
		"CANCELLING",
		"CANCELLED",
	}
}

// GetMappingListPsaWorkRequestsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPsaWorkRequestsStatusEnum(val string) (ListPsaWorkRequestsStatusEnum, bool) {
	enum, ok := mappingListPsaWorkRequestsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPsaWorkRequestsSortOrderEnum Enum with underlying type: string
type ListPsaWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListPsaWorkRequestsSortOrderEnum
const (
	ListPsaWorkRequestsSortOrderAsc  ListPsaWorkRequestsSortOrderEnum = "ASC"
	ListPsaWorkRequestsSortOrderDesc ListPsaWorkRequestsSortOrderEnum = "DESC"
)

var mappingListPsaWorkRequestsSortOrderEnum = map[string]ListPsaWorkRequestsSortOrderEnum{
	"ASC":  ListPsaWorkRequestsSortOrderAsc,
	"DESC": ListPsaWorkRequestsSortOrderDesc,
}

var mappingListPsaWorkRequestsSortOrderEnumLowerCase = map[string]ListPsaWorkRequestsSortOrderEnum{
	"asc":  ListPsaWorkRequestsSortOrderAsc,
	"desc": ListPsaWorkRequestsSortOrderDesc,
}

// GetListPsaWorkRequestsSortOrderEnumValues Enumerates the set of values for ListPsaWorkRequestsSortOrderEnum
func GetListPsaWorkRequestsSortOrderEnumValues() []ListPsaWorkRequestsSortOrderEnum {
	values := make([]ListPsaWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListPsaWorkRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPsaWorkRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListPsaWorkRequestsSortOrderEnum
func GetListPsaWorkRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPsaWorkRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPsaWorkRequestsSortOrderEnum(val string) (ListPsaWorkRequestsSortOrderEnum, bool) {
	enum, ok := mappingListPsaWorkRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPsaWorkRequestsSortByEnum Enum with underlying type: string
type ListPsaWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListPsaWorkRequestsSortByEnum
const (
	ListPsaWorkRequestsSortByTimeaccepted ListPsaWorkRequestsSortByEnum = "timeAccepted"
)

var mappingListPsaWorkRequestsSortByEnum = map[string]ListPsaWorkRequestsSortByEnum{
	"timeAccepted": ListPsaWorkRequestsSortByTimeaccepted,
}

var mappingListPsaWorkRequestsSortByEnumLowerCase = map[string]ListPsaWorkRequestsSortByEnum{
	"timeaccepted": ListPsaWorkRequestsSortByTimeaccepted,
}

// GetListPsaWorkRequestsSortByEnumValues Enumerates the set of values for ListPsaWorkRequestsSortByEnum
func GetListPsaWorkRequestsSortByEnumValues() []ListPsaWorkRequestsSortByEnum {
	values := make([]ListPsaWorkRequestsSortByEnum, 0)
	for _, v := range mappingListPsaWorkRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPsaWorkRequestsSortByEnumStringValues Enumerates the set of values in String for ListPsaWorkRequestsSortByEnum
func GetListPsaWorkRequestsSortByEnumStringValues() []string {
	return []string{
		"timeAccepted",
	}
}

// GetMappingListPsaWorkRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPsaWorkRequestsSortByEnum(val string) (ListPsaWorkRequestsSortByEnum, bool) {
	enum, ok := mappingListPsaWorkRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
