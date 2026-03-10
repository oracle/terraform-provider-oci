// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package email

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
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/email/ListWorkRequests.go.html to see an example of how to use ListWorkRequestsRequest.
type ListWorkRequestsRequest struct {

	// The OCID for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ID of the asynchronous work request.
	WorkRequestId *string `mandatory:"false" contributesTo:"query" name:"workRequestId"`

	// The request ID for tracing from the system
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. `1` is the minimum, `1000` is the maximum. For important details about
	// how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to return only resources their lifecycleState matches the given OperationStatus.
	Status ListWorkRequestsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeAccepted is descending.
	SortBy ListWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending or descending order.
	SortOrder ListWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources matching the given operation type.
	OperationType ListWorkRequestsOperationTypeEnum `mandatory:"false" contributesTo:"query" name:"operationType" omitEmpty:"true"`

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
	if _, ok := GetMappingListWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWorkRequestsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsOperationTypeEnum(string(request.OperationType)); !ok && request.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", request.OperationType, strings.Join(GetListWorkRequestsOperationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
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

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
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

// ListWorkRequestsOperationTypeEnum Enum with underlying type: string
type ListWorkRequestsOperationTypeEnum string

// Set of constants representing the allowable values for ListWorkRequestsOperationTypeEnum
const (
	ListWorkRequestsOperationTypeCreateDkim        ListWorkRequestsOperationTypeEnum = "CREATE_DKIM"
	ListWorkRequestsOperationTypeDeleteDkim        ListWorkRequestsOperationTypeEnum = "DELETE_DKIM"
	ListWorkRequestsOperationTypeMoveDkim          ListWorkRequestsOperationTypeEnum = "MOVE_DKIM"
	ListWorkRequestsOperationTypeUpdateDkim        ListWorkRequestsOperationTypeEnum = "UPDATE_DKIM"
	ListWorkRequestsOperationTypeCreateEmailDomain ListWorkRequestsOperationTypeEnum = "CREATE_EMAIL_DOMAIN"
	ListWorkRequestsOperationTypeDeleteEmailDomain ListWorkRequestsOperationTypeEnum = "DELETE_EMAIL_DOMAIN"
	ListWorkRequestsOperationTypeMoveEmailDomain   ListWorkRequestsOperationTypeEnum = "MOVE_EMAIL_DOMAIN"
	ListWorkRequestsOperationTypeUpdateEmailDomain ListWorkRequestsOperationTypeEnum = "UPDATE_EMAIL_DOMAIN"
	ListWorkRequestsOperationTypeCreateReturnPath  ListWorkRequestsOperationTypeEnum = "CREATE_RETURN_PATH"
	ListWorkRequestsOperationTypeDeleteReturnPath  ListWorkRequestsOperationTypeEnum = "DELETE_RETURN_PATH"
	ListWorkRequestsOperationTypeUpdateReturnPath  ListWorkRequestsOperationTypeEnum = "UPDATE_RETURN_PATH"
	ListWorkRequestsOperationTypeCreateIpPool      ListWorkRequestsOperationTypeEnum = "CREATE_IP_POOL"
	ListWorkRequestsOperationTypeUpdateIpPool      ListWorkRequestsOperationTypeEnum = "UPDATE_IP_POOL"
	ListWorkRequestsOperationTypeDeleteIpPool      ListWorkRequestsOperationTypeEnum = "DELETE_IP_POOL"
	ListWorkRequestsOperationTypeMoveIpPool        ListWorkRequestsOperationTypeEnum = "MOVE_IP_POOL"
)

var mappingListWorkRequestsOperationTypeEnum = map[string]ListWorkRequestsOperationTypeEnum{
	"CREATE_DKIM":         ListWorkRequestsOperationTypeCreateDkim,
	"DELETE_DKIM":         ListWorkRequestsOperationTypeDeleteDkim,
	"MOVE_DKIM":           ListWorkRequestsOperationTypeMoveDkim,
	"UPDATE_DKIM":         ListWorkRequestsOperationTypeUpdateDkim,
	"CREATE_EMAIL_DOMAIN": ListWorkRequestsOperationTypeCreateEmailDomain,
	"DELETE_EMAIL_DOMAIN": ListWorkRequestsOperationTypeDeleteEmailDomain,
	"MOVE_EMAIL_DOMAIN":   ListWorkRequestsOperationTypeMoveEmailDomain,
	"UPDATE_EMAIL_DOMAIN": ListWorkRequestsOperationTypeUpdateEmailDomain,
	"CREATE_RETURN_PATH":  ListWorkRequestsOperationTypeCreateReturnPath,
	"DELETE_RETURN_PATH":  ListWorkRequestsOperationTypeDeleteReturnPath,
	"UPDATE_RETURN_PATH":  ListWorkRequestsOperationTypeUpdateReturnPath,
	"CREATE_IP_POOL":      ListWorkRequestsOperationTypeCreateIpPool,
	"UPDATE_IP_POOL":      ListWorkRequestsOperationTypeUpdateIpPool,
	"DELETE_IP_POOL":      ListWorkRequestsOperationTypeDeleteIpPool,
	"MOVE_IP_POOL":        ListWorkRequestsOperationTypeMoveIpPool,
}

var mappingListWorkRequestsOperationTypeEnumLowerCase = map[string]ListWorkRequestsOperationTypeEnum{
	"create_dkim":         ListWorkRequestsOperationTypeCreateDkim,
	"delete_dkim":         ListWorkRequestsOperationTypeDeleteDkim,
	"move_dkim":           ListWorkRequestsOperationTypeMoveDkim,
	"update_dkim":         ListWorkRequestsOperationTypeUpdateDkim,
	"create_email_domain": ListWorkRequestsOperationTypeCreateEmailDomain,
	"delete_email_domain": ListWorkRequestsOperationTypeDeleteEmailDomain,
	"move_email_domain":   ListWorkRequestsOperationTypeMoveEmailDomain,
	"update_email_domain": ListWorkRequestsOperationTypeUpdateEmailDomain,
	"create_return_path":  ListWorkRequestsOperationTypeCreateReturnPath,
	"delete_return_path":  ListWorkRequestsOperationTypeDeleteReturnPath,
	"update_return_path":  ListWorkRequestsOperationTypeUpdateReturnPath,
	"create_ip_pool":      ListWorkRequestsOperationTypeCreateIpPool,
	"update_ip_pool":      ListWorkRequestsOperationTypeUpdateIpPool,
	"delete_ip_pool":      ListWorkRequestsOperationTypeDeleteIpPool,
	"move_ip_pool":        ListWorkRequestsOperationTypeMoveIpPool,
}

// GetListWorkRequestsOperationTypeEnumValues Enumerates the set of values for ListWorkRequestsOperationTypeEnum
func GetListWorkRequestsOperationTypeEnumValues() []ListWorkRequestsOperationTypeEnum {
	values := make([]ListWorkRequestsOperationTypeEnum, 0)
	for _, v := range mappingListWorkRequestsOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsOperationTypeEnumStringValues Enumerates the set of values in String for ListWorkRequestsOperationTypeEnum
func GetListWorkRequestsOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_DKIM",
		"DELETE_DKIM",
		"MOVE_DKIM",
		"UPDATE_DKIM",
		"CREATE_EMAIL_DOMAIN",
		"DELETE_EMAIL_DOMAIN",
		"MOVE_EMAIL_DOMAIN",
		"UPDATE_EMAIL_DOMAIN",
		"CREATE_RETURN_PATH",
		"DELETE_RETURN_PATH",
		"UPDATE_RETURN_PATH",
		"CREATE_IP_POOL",
		"UPDATE_IP_POOL",
		"DELETE_IP_POOL",
		"MOVE_IP_POOL",
	}
}

// GetMappingListWorkRequestsOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsOperationTypeEnum(val string) (ListWorkRequestsOperationTypeEnum, bool) {
	enum, ok := mappingListWorkRequestsOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
