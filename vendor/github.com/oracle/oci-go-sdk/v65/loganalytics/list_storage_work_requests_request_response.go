// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListStorageWorkRequestsRequest wrapper for the ListStorageWorkRequests operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListStorageWorkRequests.go.html to see an example of how to use ListStorageWorkRequestsRequest.
type ListStorageWorkRequestsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListStorageWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// This is the query parameter of which field to sort by. Only one sort order may be provided. Default order for timeAccepted
	// is descending. If no value is specified timeAccepted is default.
	SortBy ListStorageWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The is the work request type query parameter
	OperationType ListStorageWorkRequestsOperationTypeEnum `mandatory:"false" contributesTo:"query" name:"operationType" omitEmpty:"true"`

	// The is the work request status query parameter
	Status ListStorageWorkRequestsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The is the query parameter of when the processing of work request was started
	TimeStarted *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStarted"`

	// The is the query parameter of when the processing of work request was finished
	TimeFinished *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeFinished"`

	// This is the query parameter of purge policy name
	PolicyName *string `mandatory:"false" contributesTo:"query" name:"policyName"`

	// This is the query parameter of purge policy ID
	PolicyId *string `mandatory:"false" contributesTo:"query" name:"policyId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStorageWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStorageWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListStorageWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStorageWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListStorageWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListStorageWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListStorageWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStorageWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListStorageWorkRequestsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStorageWorkRequestsOperationTypeEnum(string(request.OperationType)); !ok && request.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", request.OperationType, strings.Join(GetListStorageWorkRequestsOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStorageWorkRequestsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListStorageWorkRequestsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListStorageWorkRequestsResponse wrapper for the ListStorageWorkRequests operation
type ListStorageWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of StorageWorkRequestCollection instances
	StorageWorkRequestCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListStorageWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStorageWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStorageWorkRequestsSortOrderEnum Enum with underlying type: string
type ListStorageWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListStorageWorkRequestsSortOrderEnum
const (
	ListStorageWorkRequestsSortOrderAsc  ListStorageWorkRequestsSortOrderEnum = "ASC"
	ListStorageWorkRequestsSortOrderDesc ListStorageWorkRequestsSortOrderEnum = "DESC"
)

var mappingListStorageWorkRequestsSortOrderEnum = map[string]ListStorageWorkRequestsSortOrderEnum{
	"ASC":  ListStorageWorkRequestsSortOrderAsc,
	"DESC": ListStorageWorkRequestsSortOrderDesc,
}

var mappingListStorageWorkRequestsSortOrderEnumLowerCase = map[string]ListStorageWorkRequestsSortOrderEnum{
	"asc":  ListStorageWorkRequestsSortOrderAsc,
	"desc": ListStorageWorkRequestsSortOrderDesc,
}

// GetListStorageWorkRequestsSortOrderEnumValues Enumerates the set of values for ListStorageWorkRequestsSortOrderEnum
func GetListStorageWorkRequestsSortOrderEnumValues() []ListStorageWorkRequestsSortOrderEnum {
	values := make([]ListStorageWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListStorageWorkRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListStorageWorkRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListStorageWorkRequestsSortOrderEnum
func GetListStorageWorkRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListStorageWorkRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStorageWorkRequestsSortOrderEnum(val string) (ListStorageWorkRequestsSortOrderEnum, bool) {
	enum, ok := mappingListStorageWorkRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStorageWorkRequestsSortByEnum Enum with underlying type: string
type ListStorageWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListStorageWorkRequestsSortByEnum
const (
	ListStorageWorkRequestsSortByTimeaccepted ListStorageWorkRequestsSortByEnum = "timeAccepted"
	ListStorageWorkRequestsSortByTimeexpires  ListStorageWorkRequestsSortByEnum = "timeExpires"
	ListStorageWorkRequestsSortByTimefinished ListStorageWorkRequestsSortByEnum = "timeFinished"
)

var mappingListStorageWorkRequestsSortByEnum = map[string]ListStorageWorkRequestsSortByEnum{
	"timeAccepted": ListStorageWorkRequestsSortByTimeaccepted,
	"timeExpires":  ListStorageWorkRequestsSortByTimeexpires,
	"timeFinished": ListStorageWorkRequestsSortByTimefinished,
}

var mappingListStorageWorkRequestsSortByEnumLowerCase = map[string]ListStorageWorkRequestsSortByEnum{
	"timeaccepted": ListStorageWorkRequestsSortByTimeaccepted,
	"timeexpires":  ListStorageWorkRequestsSortByTimeexpires,
	"timefinished": ListStorageWorkRequestsSortByTimefinished,
}

// GetListStorageWorkRequestsSortByEnumValues Enumerates the set of values for ListStorageWorkRequestsSortByEnum
func GetListStorageWorkRequestsSortByEnumValues() []ListStorageWorkRequestsSortByEnum {
	values := make([]ListStorageWorkRequestsSortByEnum, 0)
	for _, v := range mappingListStorageWorkRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListStorageWorkRequestsSortByEnumStringValues Enumerates the set of values in String for ListStorageWorkRequestsSortByEnum
func GetListStorageWorkRequestsSortByEnumStringValues() []string {
	return []string{
		"timeAccepted",
		"timeExpires",
		"timeFinished",
	}
}

// GetMappingListStorageWorkRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStorageWorkRequestsSortByEnum(val string) (ListStorageWorkRequestsSortByEnum, bool) {
	enum, ok := mappingListStorageWorkRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStorageWorkRequestsOperationTypeEnum Enum with underlying type: string
type ListStorageWorkRequestsOperationTypeEnum string

// Set of constants representing the allowable values for ListStorageWorkRequestsOperationTypeEnum
const (
	ListStorageWorkRequestsOperationTypeOffboardTenancy            ListStorageWorkRequestsOperationTypeEnum = "OFFBOARD_TENANCY"
	ListStorageWorkRequestsOperationTypePurgeStorageData           ListStorageWorkRequestsOperationTypeEnum = "PURGE_STORAGE_DATA"
	ListStorageWorkRequestsOperationTypeRecallArchivedStorageData  ListStorageWorkRequestsOperationTypeEnum = "RECALL_ARCHIVED_STORAGE_DATA"
	ListStorageWorkRequestsOperationTypeReleaseRecalledStorageData ListStorageWorkRequestsOperationTypeEnum = "RELEASE_RECALLED_STORAGE_DATA"
	ListStorageWorkRequestsOperationTypePurgeArchivalData          ListStorageWorkRequestsOperationTypeEnum = "PURGE_ARCHIVAL_DATA"
	ListStorageWorkRequestsOperationTypeArchiveStorageData         ListStorageWorkRequestsOperationTypeEnum = "ARCHIVE_STORAGE_DATA"
	ListStorageWorkRequestsOperationTypeCleanupArchivalStorageData ListStorageWorkRequestsOperationTypeEnum = "CLEANUP_ARCHIVAL_STORAGE_DATA"
	ListStorageWorkRequestsOperationTypeEncryptActiveData          ListStorageWorkRequestsOperationTypeEnum = "ENCRYPT_ACTIVE_DATA"
	ListStorageWorkRequestsOperationTypeEncryptArchivalData        ListStorageWorkRequestsOperationTypeEnum = "ENCRYPT_ARCHIVAL_DATA"
)

var mappingListStorageWorkRequestsOperationTypeEnum = map[string]ListStorageWorkRequestsOperationTypeEnum{
	"OFFBOARD_TENANCY":              ListStorageWorkRequestsOperationTypeOffboardTenancy,
	"PURGE_STORAGE_DATA":            ListStorageWorkRequestsOperationTypePurgeStorageData,
	"RECALL_ARCHIVED_STORAGE_DATA":  ListStorageWorkRequestsOperationTypeRecallArchivedStorageData,
	"RELEASE_RECALLED_STORAGE_DATA": ListStorageWorkRequestsOperationTypeReleaseRecalledStorageData,
	"PURGE_ARCHIVAL_DATA":           ListStorageWorkRequestsOperationTypePurgeArchivalData,
	"ARCHIVE_STORAGE_DATA":          ListStorageWorkRequestsOperationTypeArchiveStorageData,
	"CLEANUP_ARCHIVAL_STORAGE_DATA": ListStorageWorkRequestsOperationTypeCleanupArchivalStorageData,
	"ENCRYPT_ACTIVE_DATA":           ListStorageWorkRequestsOperationTypeEncryptActiveData,
	"ENCRYPT_ARCHIVAL_DATA":         ListStorageWorkRequestsOperationTypeEncryptArchivalData,
}

var mappingListStorageWorkRequestsOperationTypeEnumLowerCase = map[string]ListStorageWorkRequestsOperationTypeEnum{
	"offboard_tenancy":              ListStorageWorkRequestsOperationTypeOffboardTenancy,
	"purge_storage_data":            ListStorageWorkRequestsOperationTypePurgeStorageData,
	"recall_archived_storage_data":  ListStorageWorkRequestsOperationTypeRecallArchivedStorageData,
	"release_recalled_storage_data": ListStorageWorkRequestsOperationTypeReleaseRecalledStorageData,
	"purge_archival_data":           ListStorageWorkRequestsOperationTypePurgeArchivalData,
	"archive_storage_data":          ListStorageWorkRequestsOperationTypeArchiveStorageData,
	"cleanup_archival_storage_data": ListStorageWorkRequestsOperationTypeCleanupArchivalStorageData,
	"encrypt_active_data":           ListStorageWorkRequestsOperationTypeEncryptActiveData,
	"encrypt_archival_data":         ListStorageWorkRequestsOperationTypeEncryptArchivalData,
}

// GetListStorageWorkRequestsOperationTypeEnumValues Enumerates the set of values for ListStorageWorkRequestsOperationTypeEnum
func GetListStorageWorkRequestsOperationTypeEnumValues() []ListStorageWorkRequestsOperationTypeEnum {
	values := make([]ListStorageWorkRequestsOperationTypeEnum, 0)
	for _, v := range mappingListStorageWorkRequestsOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListStorageWorkRequestsOperationTypeEnumStringValues Enumerates the set of values in String for ListStorageWorkRequestsOperationTypeEnum
func GetListStorageWorkRequestsOperationTypeEnumStringValues() []string {
	return []string{
		"OFFBOARD_TENANCY",
		"PURGE_STORAGE_DATA",
		"RECALL_ARCHIVED_STORAGE_DATA",
		"RELEASE_RECALLED_STORAGE_DATA",
		"PURGE_ARCHIVAL_DATA",
		"ARCHIVE_STORAGE_DATA",
		"CLEANUP_ARCHIVAL_STORAGE_DATA",
		"ENCRYPT_ACTIVE_DATA",
		"ENCRYPT_ARCHIVAL_DATA",
	}
}

// GetMappingListStorageWorkRequestsOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStorageWorkRequestsOperationTypeEnum(val string) (ListStorageWorkRequestsOperationTypeEnum, bool) {
	enum, ok := mappingListStorageWorkRequestsOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStorageWorkRequestsStatusEnum Enum with underlying type: string
type ListStorageWorkRequestsStatusEnum string

// Set of constants representing the allowable values for ListStorageWorkRequestsStatusEnum
const (
	ListStorageWorkRequestsStatusAccepted   ListStorageWorkRequestsStatusEnum = "ACCEPTED"
	ListStorageWorkRequestsStatusCanceled   ListStorageWorkRequestsStatusEnum = "CANCELED"
	ListStorageWorkRequestsStatusFailed     ListStorageWorkRequestsStatusEnum = "FAILED"
	ListStorageWorkRequestsStatusInProgress ListStorageWorkRequestsStatusEnum = "IN_PROGRESS"
	ListStorageWorkRequestsStatusSucceeded  ListStorageWorkRequestsStatusEnum = "SUCCEEDED"
)

var mappingListStorageWorkRequestsStatusEnum = map[string]ListStorageWorkRequestsStatusEnum{
	"ACCEPTED":    ListStorageWorkRequestsStatusAccepted,
	"CANCELED":    ListStorageWorkRequestsStatusCanceled,
	"FAILED":      ListStorageWorkRequestsStatusFailed,
	"IN_PROGRESS": ListStorageWorkRequestsStatusInProgress,
	"SUCCEEDED":   ListStorageWorkRequestsStatusSucceeded,
}

var mappingListStorageWorkRequestsStatusEnumLowerCase = map[string]ListStorageWorkRequestsStatusEnum{
	"accepted":    ListStorageWorkRequestsStatusAccepted,
	"canceled":    ListStorageWorkRequestsStatusCanceled,
	"failed":      ListStorageWorkRequestsStatusFailed,
	"in_progress": ListStorageWorkRequestsStatusInProgress,
	"succeeded":   ListStorageWorkRequestsStatusSucceeded,
}

// GetListStorageWorkRequestsStatusEnumValues Enumerates the set of values for ListStorageWorkRequestsStatusEnum
func GetListStorageWorkRequestsStatusEnumValues() []ListStorageWorkRequestsStatusEnum {
	values := make([]ListStorageWorkRequestsStatusEnum, 0)
	for _, v := range mappingListStorageWorkRequestsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListStorageWorkRequestsStatusEnumStringValues Enumerates the set of values in String for ListStorageWorkRequestsStatusEnum
func GetListStorageWorkRequestsStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"CANCELED",
		"FAILED",
		"IN_PROGRESS",
		"SUCCEEDED",
	}
}

// GetMappingListStorageWorkRequestsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStorageWorkRequestsStatusEnum(val string) (ListStorageWorkRequestsStatusEnum, bool) {
	enum, ok := mappingListStorageWorkRequestsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
