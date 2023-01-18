// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBackupsRequest wrapper for the ListBackups operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/ListBackups.go.html to see an example of how to use ListBackupsRequest.
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

	// Backup creationType
	CreationType BackupCreationTypeEnum `mandatory:"false" contributesTo:"query" name:"creationType" omitEmpty:"true"`

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
func (request ListBackupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBackupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBackupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBackupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBackupLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBackupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBackupCreationTypeEnum(string(request.CreationType)); !ok && request.CreationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreationType: %s. Supported values are: %s.", request.CreationType, strings.Join(GetBackupCreationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBackupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBackupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBackupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBackupsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBackupsResponse wrapper for the ListBackups operation
type ListBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []BackupSummary instances
	Items []BackupSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
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

var mappingListBackupsSortByEnum = map[string]ListBackupsSortByEnum{
	"timeCreated": ListBackupsSortByTimecreated,
	"timeUpdated": ListBackupsSortByTimeupdated,
	"displayName": ListBackupsSortByDisplayname,
}

var mappingListBackupsSortByEnumLowerCase = map[string]ListBackupsSortByEnum{
	"timecreated": ListBackupsSortByTimecreated,
	"timeupdated": ListBackupsSortByTimeupdated,
	"displayname": ListBackupsSortByDisplayname,
}

// GetListBackupsSortByEnumValues Enumerates the set of values for ListBackupsSortByEnum
func GetListBackupsSortByEnumValues() []ListBackupsSortByEnum {
	values := make([]ListBackupsSortByEnum, 0)
	for _, v := range mappingListBackupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBackupsSortByEnumStringValues Enumerates the set of values in String for ListBackupsSortByEnum
func GetListBackupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListBackupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBackupsSortByEnum(val string) (ListBackupsSortByEnum, bool) {
	enum, ok := mappingListBackupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBackupsSortOrderEnum Enum with underlying type: string
type ListBackupsSortOrderEnum string

// Set of constants representing the allowable values for ListBackupsSortOrderEnum
const (
	ListBackupsSortOrderAsc  ListBackupsSortOrderEnum = "ASC"
	ListBackupsSortOrderDesc ListBackupsSortOrderEnum = "DESC"
)

var mappingListBackupsSortOrderEnum = map[string]ListBackupsSortOrderEnum{
	"ASC":  ListBackupsSortOrderAsc,
	"DESC": ListBackupsSortOrderDesc,
}

var mappingListBackupsSortOrderEnumLowerCase = map[string]ListBackupsSortOrderEnum{
	"asc":  ListBackupsSortOrderAsc,
	"desc": ListBackupsSortOrderDesc,
}

// GetListBackupsSortOrderEnumValues Enumerates the set of values for ListBackupsSortOrderEnum
func GetListBackupsSortOrderEnumValues() []ListBackupsSortOrderEnum {
	values := make([]ListBackupsSortOrderEnum, 0)
	for _, v := range mappingListBackupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBackupsSortOrderEnumStringValues Enumerates the set of values in String for ListBackupsSortOrderEnum
func GetListBackupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBackupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBackupsSortOrderEnum(val string) (ListBackupsSortOrderEnum, bool) {
	enum, ok := mappingListBackupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
