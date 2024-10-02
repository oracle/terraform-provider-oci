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

// ListZprPolicyWorkRequestsRequest wrapper for the ListZprPolicyWorkRequests operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/ListZprPolicyWorkRequests.go.html to see an example of how to use ListZprPolicyWorkRequestsRequest.
type ListZprPolicyWorkRequestsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the asynchronous work request.
	WorkRequestId *string `mandatory:"false" contributesTo:"query" name:"workRequestId"`

	// A filter to return only the resources that match the given lifecycle state.
	Status ListZprPolicyWorkRequestsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

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
	SortOrder ListZprPolicyWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for `timeAccepted` is descending.
	SortBy ListZprPolicyWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListZprPolicyWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListZprPolicyWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListZprPolicyWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request ListZprPolicyWorkRequestsRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListZprPolicyWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListZprPolicyWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListZprPolicyWorkRequestsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListZprPolicyWorkRequestsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListZprPolicyWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListZprPolicyWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListZprPolicyWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListZprPolicyWorkRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListZprPolicyWorkRequestsResponse wrapper for the ListZprPolicyWorkRequests operation
type ListZprPolicyWorkRequestsResponse struct {

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

func (response ListZprPolicyWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListZprPolicyWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListZprPolicyWorkRequestsStatusEnum Enum with underlying type: string
type ListZprPolicyWorkRequestsStatusEnum string

// Set of constants representing the allowable values for ListZprPolicyWorkRequestsStatusEnum
const (
	ListZprPolicyWorkRequestsStatusAccepted       ListZprPolicyWorkRequestsStatusEnum = "ACCEPTED"
	ListZprPolicyWorkRequestsStatusInProgress     ListZprPolicyWorkRequestsStatusEnum = "IN_PROGRESS"
	ListZprPolicyWorkRequestsStatusWaiting        ListZprPolicyWorkRequestsStatusEnum = "WAITING"
	ListZprPolicyWorkRequestsStatusNeedsAttention ListZprPolicyWorkRequestsStatusEnum = "NEEDS_ATTENTION"
	ListZprPolicyWorkRequestsStatusFailed         ListZprPolicyWorkRequestsStatusEnum = "FAILED"
	ListZprPolicyWorkRequestsStatusSucceeded      ListZprPolicyWorkRequestsStatusEnum = "SUCCEEDED"
	ListZprPolicyWorkRequestsStatusCanceling      ListZprPolicyWorkRequestsStatusEnum = "CANCELING"
	ListZprPolicyWorkRequestsStatusCanceled       ListZprPolicyWorkRequestsStatusEnum = "CANCELED"
)

var mappingListZprPolicyWorkRequestsStatusEnum = map[string]ListZprPolicyWorkRequestsStatusEnum{
	"ACCEPTED":        ListZprPolicyWorkRequestsStatusAccepted,
	"IN_PROGRESS":     ListZprPolicyWorkRequestsStatusInProgress,
	"WAITING":         ListZprPolicyWorkRequestsStatusWaiting,
	"NEEDS_ATTENTION": ListZprPolicyWorkRequestsStatusNeedsAttention,
	"FAILED":          ListZprPolicyWorkRequestsStatusFailed,
	"SUCCEEDED":       ListZprPolicyWorkRequestsStatusSucceeded,
	"CANCELING":       ListZprPolicyWorkRequestsStatusCanceling,
	"CANCELED":        ListZprPolicyWorkRequestsStatusCanceled,
}

var mappingListZprPolicyWorkRequestsStatusEnumLowerCase = map[string]ListZprPolicyWorkRequestsStatusEnum{
	"accepted":        ListZprPolicyWorkRequestsStatusAccepted,
	"in_progress":     ListZprPolicyWorkRequestsStatusInProgress,
	"waiting":         ListZprPolicyWorkRequestsStatusWaiting,
	"needs_attention": ListZprPolicyWorkRequestsStatusNeedsAttention,
	"failed":          ListZprPolicyWorkRequestsStatusFailed,
	"succeeded":       ListZprPolicyWorkRequestsStatusSucceeded,
	"canceling":       ListZprPolicyWorkRequestsStatusCanceling,
	"canceled":        ListZprPolicyWorkRequestsStatusCanceled,
}

// GetListZprPolicyWorkRequestsStatusEnumValues Enumerates the set of values for ListZprPolicyWorkRequestsStatusEnum
func GetListZprPolicyWorkRequestsStatusEnumValues() []ListZprPolicyWorkRequestsStatusEnum {
	values := make([]ListZprPolicyWorkRequestsStatusEnum, 0)
	for _, v := range mappingListZprPolicyWorkRequestsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprPolicyWorkRequestsStatusEnumStringValues Enumerates the set of values in String for ListZprPolicyWorkRequestsStatusEnum
func GetListZprPolicyWorkRequestsStatusEnumStringValues() []string {
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

// GetMappingListZprPolicyWorkRequestsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprPolicyWorkRequestsStatusEnum(val string) (ListZprPolicyWorkRequestsStatusEnum, bool) {
	enum, ok := mappingListZprPolicyWorkRequestsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListZprPolicyWorkRequestsSortOrderEnum Enum with underlying type: string
type ListZprPolicyWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListZprPolicyWorkRequestsSortOrderEnum
const (
	ListZprPolicyWorkRequestsSortOrderAsc  ListZprPolicyWorkRequestsSortOrderEnum = "ASC"
	ListZprPolicyWorkRequestsSortOrderDesc ListZprPolicyWorkRequestsSortOrderEnum = "DESC"
)

var mappingListZprPolicyWorkRequestsSortOrderEnum = map[string]ListZprPolicyWorkRequestsSortOrderEnum{
	"ASC":  ListZprPolicyWorkRequestsSortOrderAsc,
	"DESC": ListZprPolicyWorkRequestsSortOrderDesc,
}

var mappingListZprPolicyWorkRequestsSortOrderEnumLowerCase = map[string]ListZprPolicyWorkRequestsSortOrderEnum{
	"asc":  ListZprPolicyWorkRequestsSortOrderAsc,
	"desc": ListZprPolicyWorkRequestsSortOrderDesc,
}

// GetListZprPolicyWorkRequestsSortOrderEnumValues Enumerates the set of values for ListZprPolicyWorkRequestsSortOrderEnum
func GetListZprPolicyWorkRequestsSortOrderEnumValues() []ListZprPolicyWorkRequestsSortOrderEnum {
	values := make([]ListZprPolicyWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListZprPolicyWorkRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprPolicyWorkRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListZprPolicyWorkRequestsSortOrderEnum
func GetListZprPolicyWorkRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListZprPolicyWorkRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprPolicyWorkRequestsSortOrderEnum(val string) (ListZprPolicyWorkRequestsSortOrderEnum, bool) {
	enum, ok := mappingListZprPolicyWorkRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListZprPolicyWorkRequestsSortByEnum Enum with underlying type: string
type ListZprPolicyWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListZprPolicyWorkRequestsSortByEnum
const (
	ListZprPolicyWorkRequestsSortByTimeaccepted ListZprPolicyWorkRequestsSortByEnum = "timeAccepted"
)

var mappingListZprPolicyWorkRequestsSortByEnum = map[string]ListZprPolicyWorkRequestsSortByEnum{
	"timeAccepted": ListZprPolicyWorkRequestsSortByTimeaccepted,
}

var mappingListZprPolicyWorkRequestsSortByEnumLowerCase = map[string]ListZprPolicyWorkRequestsSortByEnum{
	"timeaccepted": ListZprPolicyWorkRequestsSortByTimeaccepted,
}

// GetListZprPolicyWorkRequestsSortByEnumValues Enumerates the set of values for ListZprPolicyWorkRequestsSortByEnum
func GetListZprPolicyWorkRequestsSortByEnumValues() []ListZprPolicyWorkRequestsSortByEnum {
	values := make([]ListZprPolicyWorkRequestsSortByEnum, 0)
	for _, v := range mappingListZprPolicyWorkRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprPolicyWorkRequestsSortByEnumStringValues Enumerates the set of values in String for ListZprPolicyWorkRequestsSortByEnum
func GetListZprPolicyWorkRequestsSortByEnumStringValues() []string {
	return []string{
		"timeAccepted",
	}
}

// GetMappingListZprPolicyWorkRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprPolicyWorkRequestsSortByEnum(val string) (ListZprPolicyWorkRequestsSortByEnum, bool) {
	enum, ok := mappingListZprPolicyWorkRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
