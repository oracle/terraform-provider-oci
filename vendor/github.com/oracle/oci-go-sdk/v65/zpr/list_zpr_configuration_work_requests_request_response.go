// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package zpr

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListZprConfigurationWorkRequestsRequest wrapper for the ListZprConfigurationWorkRequests operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/ListZprConfigurationWorkRequests.go.html to see an example of how to use ListZprConfigurationWorkRequestsRequest.
type ListZprConfigurationWorkRequestsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the asynchronous work request.
	WorkRequestId *string `mandatory:"false" contributesTo:"query" name:"workRequestId"`

	// A filter to return only the resources that match the given lifecycle state.
	Status ListZprConfigurationWorkRequestsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource affected by the work request.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListZprConfigurationWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for `timeAccepted` is descending.
	SortBy ListZprConfigurationWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListZprConfigurationWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListZprConfigurationWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListZprConfigurationWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request ListZprConfigurationWorkRequestsRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListZprConfigurationWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListZprConfigurationWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListZprConfigurationWorkRequestsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListZprConfigurationWorkRequestsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListZprConfigurationWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListZprConfigurationWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListZprConfigurationWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListZprConfigurationWorkRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListZprConfigurationWorkRequestsResponse wrapper for the ListZprConfigurationWorkRequests operation
type ListZprConfigurationWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestSummaryCollection instances
	WorkRequestSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListZprConfigurationWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListZprConfigurationWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListZprConfigurationWorkRequestsStatusEnum Enum with underlying type: string
type ListZprConfigurationWorkRequestsStatusEnum string

// Set of constants representing the allowable values for ListZprConfigurationWorkRequestsStatusEnum
const (
	ListZprConfigurationWorkRequestsStatusAccepted       ListZprConfigurationWorkRequestsStatusEnum = "ACCEPTED"
	ListZprConfigurationWorkRequestsStatusInProgress     ListZprConfigurationWorkRequestsStatusEnum = "IN_PROGRESS"
	ListZprConfigurationWorkRequestsStatusWaiting        ListZprConfigurationWorkRequestsStatusEnum = "WAITING"
	ListZprConfigurationWorkRequestsStatusNeedsAttention ListZprConfigurationWorkRequestsStatusEnum = "NEEDS_ATTENTION"
	ListZprConfigurationWorkRequestsStatusFailed         ListZprConfigurationWorkRequestsStatusEnum = "FAILED"
	ListZprConfigurationWorkRequestsStatusSucceeded      ListZprConfigurationWorkRequestsStatusEnum = "SUCCEEDED"
	ListZprConfigurationWorkRequestsStatusCanceling      ListZprConfigurationWorkRequestsStatusEnum = "CANCELING"
	ListZprConfigurationWorkRequestsStatusCanceled       ListZprConfigurationWorkRequestsStatusEnum = "CANCELED"
)

var mappingListZprConfigurationWorkRequestsStatusEnum = map[string]ListZprConfigurationWorkRequestsStatusEnum{
	"ACCEPTED":        ListZprConfigurationWorkRequestsStatusAccepted,
	"IN_PROGRESS":     ListZprConfigurationWorkRequestsStatusInProgress,
	"WAITING":         ListZprConfigurationWorkRequestsStatusWaiting,
	"NEEDS_ATTENTION": ListZprConfigurationWorkRequestsStatusNeedsAttention,
	"FAILED":          ListZprConfigurationWorkRequestsStatusFailed,
	"SUCCEEDED":       ListZprConfigurationWorkRequestsStatusSucceeded,
	"CANCELING":       ListZprConfigurationWorkRequestsStatusCanceling,
	"CANCELED":        ListZprConfigurationWorkRequestsStatusCanceled,
}

var mappingListZprConfigurationWorkRequestsStatusEnumLowerCase = map[string]ListZprConfigurationWorkRequestsStatusEnum{
	"accepted":        ListZprConfigurationWorkRequestsStatusAccepted,
	"in_progress":     ListZprConfigurationWorkRequestsStatusInProgress,
	"waiting":         ListZprConfigurationWorkRequestsStatusWaiting,
	"needs_attention": ListZprConfigurationWorkRequestsStatusNeedsAttention,
	"failed":          ListZprConfigurationWorkRequestsStatusFailed,
	"succeeded":       ListZprConfigurationWorkRequestsStatusSucceeded,
	"canceling":       ListZprConfigurationWorkRequestsStatusCanceling,
	"canceled":        ListZprConfigurationWorkRequestsStatusCanceled,
}

// GetListZprConfigurationWorkRequestsStatusEnumValues Enumerates the set of values for ListZprConfigurationWorkRequestsStatusEnum
func GetListZprConfigurationWorkRequestsStatusEnumValues() []ListZprConfigurationWorkRequestsStatusEnum {
	values := make([]ListZprConfigurationWorkRequestsStatusEnum, 0)
	for _, v := range mappingListZprConfigurationWorkRequestsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprConfigurationWorkRequestsStatusEnumStringValues Enumerates the set of values in String for ListZprConfigurationWorkRequestsStatusEnum
func GetListZprConfigurationWorkRequestsStatusEnumStringValues() []string {
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

// GetMappingListZprConfigurationWorkRequestsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprConfigurationWorkRequestsStatusEnum(val string) (ListZprConfigurationWorkRequestsStatusEnum, bool) {
	enum, ok := mappingListZprConfigurationWorkRequestsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListZprConfigurationWorkRequestsSortOrderEnum Enum with underlying type: string
type ListZprConfigurationWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListZprConfigurationWorkRequestsSortOrderEnum
const (
	ListZprConfigurationWorkRequestsSortOrderAsc  ListZprConfigurationWorkRequestsSortOrderEnum = "ASC"
	ListZprConfigurationWorkRequestsSortOrderDesc ListZprConfigurationWorkRequestsSortOrderEnum = "DESC"
)

var mappingListZprConfigurationWorkRequestsSortOrderEnum = map[string]ListZprConfigurationWorkRequestsSortOrderEnum{
	"ASC":  ListZprConfigurationWorkRequestsSortOrderAsc,
	"DESC": ListZprConfigurationWorkRequestsSortOrderDesc,
}

var mappingListZprConfigurationWorkRequestsSortOrderEnumLowerCase = map[string]ListZprConfigurationWorkRequestsSortOrderEnum{
	"asc":  ListZprConfigurationWorkRequestsSortOrderAsc,
	"desc": ListZprConfigurationWorkRequestsSortOrderDesc,
}

// GetListZprConfigurationWorkRequestsSortOrderEnumValues Enumerates the set of values for ListZprConfigurationWorkRequestsSortOrderEnum
func GetListZprConfigurationWorkRequestsSortOrderEnumValues() []ListZprConfigurationWorkRequestsSortOrderEnum {
	values := make([]ListZprConfigurationWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListZprConfigurationWorkRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprConfigurationWorkRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListZprConfigurationWorkRequestsSortOrderEnum
func GetListZprConfigurationWorkRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListZprConfigurationWorkRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprConfigurationWorkRequestsSortOrderEnum(val string) (ListZprConfigurationWorkRequestsSortOrderEnum, bool) {
	enum, ok := mappingListZprConfigurationWorkRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListZprConfigurationWorkRequestsSortByEnum Enum with underlying type: string
type ListZprConfigurationWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListZprConfigurationWorkRequestsSortByEnum
const (
	ListZprConfigurationWorkRequestsSortByTimeaccepted ListZprConfigurationWorkRequestsSortByEnum = "timeAccepted"
)

var mappingListZprConfigurationWorkRequestsSortByEnum = map[string]ListZprConfigurationWorkRequestsSortByEnum{
	"timeAccepted": ListZprConfigurationWorkRequestsSortByTimeaccepted,
}

var mappingListZprConfigurationWorkRequestsSortByEnumLowerCase = map[string]ListZprConfigurationWorkRequestsSortByEnum{
	"timeaccepted": ListZprConfigurationWorkRequestsSortByTimeaccepted,
}

// GetListZprConfigurationWorkRequestsSortByEnumValues Enumerates the set of values for ListZprConfigurationWorkRequestsSortByEnum
func GetListZprConfigurationWorkRequestsSortByEnumValues() []ListZprConfigurationWorkRequestsSortByEnum {
	values := make([]ListZprConfigurationWorkRequestsSortByEnum, 0)
	for _, v := range mappingListZprConfigurationWorkRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprConfigurationWorkRequestsSortByEnumStringValues Enumerates the set of values in String for ListZprConfigurationWorkRequestsSortByEnum
func GetListZprConfigurationWorkRequestsSortByEnumStringValues() []string {
	return []string{
		"timeAccepted",
	}
}

// GetMappingListZprConfigurationWorkRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprConfigurationWorkRequestsSortByEnum(val string) (ListZprConfigurationWorkRequestsSortByEnum, bool) {
	enum, ok := mappingListZprConfigurationWorkRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
