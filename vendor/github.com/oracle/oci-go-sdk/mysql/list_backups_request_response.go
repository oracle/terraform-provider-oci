// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListBackupsRequest wrapper for the ListBackups operation
type ListBackupsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Backup OCID
	BackupId *string `mandatory:"false" contributesTo:"query" name:"backupId"`

	// Backup Lifecycle State
	LifecycleState BackupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The DB System OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `mandatory:"false" contributesTo:"query" name:"dbSystemId"`

	// A filter to return only the resource matching the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. Only one sort order may be provided. Time fields are default ordered as descending.
	SortBy ListBackupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use (ASC or DESC).
	SortOrder ListBackupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated list call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from
	// the previous list call. For information about pagination, see List
	// Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBackupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBackupsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBackupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListBackupsResponse wrapper for the ListBackups operation
type ListBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []BackupSummary instances
	Items []BackupSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a specific request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Opaque token representing the next page of results.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBackupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBackupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBackupsSortByEnum Enum with underlying type: string
type ListBackupsSortByEnum string

// Set of constants representing the allowable values for ListBackupsSortByEnum
const (
	ListBackupsSortByTimecreated ListBackupsSortByEnum = "timeCreated"
	ListBackupsSortByTimeupdated ListBackupsSortByEnum = "timeUpdated"
	ListBackupsSortByDisplayname ListBackupsSortByEnum = "displayName"
)

var mappingListBackupsSortBy = map[string]ListBackupsSortByEnum{
	"timeCreated": ListBackupsSortByTimecreated,
	"timeUpdated": ListBackupsSortByTimeupdated,
	"displayName": ListBackupsSortByDisplayname,
}

// GetListBackupsSortByEnumValues Enumerates the set of values for ListBackupsSortByEnum
func GetListBackupsSortByEnumValues() []ListBackupsSortByEnum {
	values := make([]ListBackupsSortByEnum, 0)
	for _, v := range mappingListBackupsSortBy {
		values = append(values, v)
	}
	return values
}

// ListBackupsSortOrderEnum Enum with underlying type: string
type ListBackupsSortOrderEnum string

// Set of constants representing the allowable values for ListBackupsSortOrderEnum
const (
	ListBackupsSortOrderAsc  ListBackupsSortOrderEnum = "ASC"
	ListBackupsSortOrderDesc ListBackupsSortOrderEnum = "DESC"
)

var mappingListBackupsSortOrder = map[string]ListBackupsSortOrderEnum{
	"ASC":  ListBackupsSortOrderAsc,
	"DESC": ListBackupsSortOrderDesc,
}

// GetListBackupsSortOrderEnumValues Enumerates the set of values for ListBackupsSortOrderEnum
func GetListBackupsSortOrderEnumValues() []ListBackupsSortOrderEnum {
	values := make([]ListBackupsSortOrderEnum, 0)
	for _, v := range mappingListBackupsSortOrder {
		values = append(values, v)
	}
	return values
}
