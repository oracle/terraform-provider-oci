// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListStorageWorkRequestsRequest wrapper for the ListStorageWorkRequests operation
//
// See also
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

var mappingListStorageWorkRequestsSortOrder = map[string]ListStorageWorkRequestsSortOrderEnum{
	"ASC":  ListStorageWorkRequestsSortOrderAsc,
	"DESC": ListStorageWorkRequestsSortOrderDesc,
}

// GetListStorageWorkRequestsSortOrderEnumValues Enumerates the set of values for ListStorageWorkRequestsSortOrderEnum
func GetListStorageWorkRequestsSortOrderEnumValues() []ListStorageWorkRequestsSortOrderEnum {
	values := make([]ListStorageWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListStorageWorkRequestsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListStorageWorkRequestsSortByEnum Enum with underlying type: string
type ListStorageWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListStorageWorkRequestsSortByEnum
const (
	ListStorageWorkRequestsSortByTimeaccepted ListStorageWorkRequestsSortByEnum = "timeAccepted"
	ListStorageWorkRequestsSortByTimeexpires  ListStorageWorkRequestsSortByEnum = "timeExpires"
	ListStorageWorkRequestsSortByTimefinished ListStorageWorkRequestsSortByEnum = "timeFinished"
)

var mappingListStorageWorkRequestsSortBy = map[string]ListStorageWorkRequestsSortByEnum{
	"timeAccepted": ListStorageWorkRequestsSortByTimeaccepted,
	"timeExpires":  ListStorageWorkRequestsSortByTimeexpires,
	"timeFinished": ListStorageWorkRequestsSortByTimefinished,
}

// GetListStorageWorkRequestsSortByEnumValues Enumerates the set of values for ListStorageWorkRequestsSortByEnum
func GetListStorageWorkRequestsSortByEnumValues() []ListStorageWorkRequestsSortByEnum {
	values := make([]ListStorageWorkRequestsSortByEnum, 0)
	for _, v := range mappingListStorageWorkRequestsSortBy {
		values = append(values, v)
	}
	return values
}

// ListStorageWorkRequestsOperationTypeEnum Enum with underlying type: string
type ListStorageWorkRequestsOperationTypeEnum string

// Set of constants representing the allowable values for ListStorageWorkRequestsOperationTypeEnum
const (
	ListStorageWorkRequestsOperationTypeOffboardTenancy            ListStorageWorkRequestsOperationTypeEnum = "OFFBOARD_TENANCY"
	ListStorageWorkRequestsOperationTypePurgeStorageData           ListStorageWorkRequestsOperationTypeEnum = "PURGE_STORAGE_DATA"
	ListStorageWorkRequestsOperationTypeRecallArchivedStorageData  ListStorageWorkRequestsOperationTypeEnum = "RECALL_ARCHIVED_STORAGE_DATA"
	ListStorageWorkRequestsOperationTypeReleaseRecalledStorageData ListStorageWorkRequestsOperationTypeEnum = "RELEASE_RECALLED_STORAGE_DATA"
	ListStorageWorkRequestsOperationTypeArchiveStorageData         ListStorageWorkRequestsOperationTypeEnum = "ARCHIVE_STORAGE_DATA"
	ListStorageWorkRequestsOperationTypeCleanupArchivalStorageData ListStorageWorkRequestsOperationTypeEnum = "CLEANUP_ARCHIVAL_STORAGE_DATA"
)

var mappingListStorageWorkRequestsOperationType = map[string]ListStorageWorkRequestsOperationTypeEnum{
	"OFFBOARD_TENANCY":              ListStorageWorkRequestsOperationTypeOffboardTenancy,
	"PURGE_STORAGE_DATA":            ListStorageWorkRequestsOperationTypePurgeStorageData,
	"RECALL_ARCHIVED_STORAGE_DATA":  ListStorageWorkRequestsOperationTypeRecallArchivedStorageData,
	"RELEASE_RECALLED_STORAGE_DATA": ListStorageWorkRequestsOperationTypeReleaseRecalledStorageData,
	"ARCHIVE_STORAGE_DATA":          ListStorageWorkRequestsOperationTypeArchiveStorageData,
	"CLEANUP_ARCHIVAL_STORAGE_DATA": ListStorageWorkRequestsOperationTypeCleanupArchivalStorageData,
}

// GetListStorageWorkRequestsOperationTypeEnumValues Enumerates the set of values for ListStorageWorkRequestsOperationTypeEnum
func GetListStorageWorkRequestsOperationTypeEnumValues() []ListStorageWorkRequestsOperationTypeEnum {
	values := make([]ListStorageWorkRequestsOperationTypeEnum, 0)
	for _, v := range mappingListStorageWorkRequestsOperationType {
		values = append(values, v)
	}
	return values
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

var mappingListStorageWorkRequestsStatus = map[string]ListStorageWorkRequestsStatusEnum{
	"ACCEPTED":    ListStorageWorkRequestsStatusAccepted,
	"CANCELED":    ListStorageWorkRequestsStatusCanceled,
	"FAILED":      ListStorageWorkRequestsStatusFailed,
	"IN_PROGRESS": ListStorageWorkRequestsStatusInProgress,
	"SUCCEEDED":   ListStorageWorkRequestsStatusSucceeded,
}

// GetListStorageWorkRequestsStatusEnumValues Enumerates the set of values for ListStorageWorkRequestsStatusEnum
func GetListStorageWorkRequestsStatusEnumValues() []ListStorageWorkRequestsStatusEnum {
	values := make([]ListStorageWorkRequestsStatusEnum, 0)
	for _, v := range mappingListStorageWorkRequestsStatus {
		values = append(values, v)
	}
	return values
}
