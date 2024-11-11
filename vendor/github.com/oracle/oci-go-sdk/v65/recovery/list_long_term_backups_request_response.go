// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLongTermBackupsRequest wrapper for the ListLongTermBackups operation
type ListLongTermBackupsRequest struct {

	// The compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The long-term backup OCID. Use longTermBackupId to filter a long-term backup based on its unique identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the entire 'displayname' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The protected database OCID. Use protectedDatabaseId to list the long-term backups of a specific protected database.
	ProtectedDatabaseId *string `mandatory:"false" contributesTo:"query" name:"protectedDatabaseId"`

	// A filter to return only the resources that match the specified lifecycle state.
	LifecycleState ListLongTermBackupsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	// Allowed values are:
	//   - ASC
	//   - DESC
	SortOrder ListLongTermBackupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (sortOrder). Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If you do not specify a value, then TIMECREATED is used as the default sort order.
	// Allowed values are:
	//   - TIMECREATED
	//   - DISPLAYNAME
	SortBy ListLongTermBackupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLongTermBackupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLongTermBackupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLongTermBackupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLongTermBackupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLongTermBackupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLongTermBackupsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListLongTermBackupsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLongTermBackupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLongTermBackupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLongTermBackupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLongTermBackupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLongTermBackupsResponse wrapper for the ListLongTermBackups operation
type ListLongTermBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LongTermBackupCollection instances
	LongTermBackupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLongTermBackupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLongTermBackupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLongTermBackupsLifecycleStateEnum Enum with underlying type: string
type ListLongTermBackupsLifecycleStateEnum string

// Set of constants representing the allowable values for ListLongTermBackupsLifecycleStateEnum
const (
	ListLongTermBackupsLifecycleStateCreating        ListLongTermBackupsLifecycleStateEnum = "CREATING"
	ListLongTermBackupsLifecycleStateUpdating        ListLongTermBackupsLifecycleStateEnum = "UPDATING"
	ListLongTermBackupsLifecycleStateActive          ListLongTermBackupsLifecycleStateEnum = "ACTIVE"
	ListLongTermBackupsLifecycleStateDeleteScheduled ListLongTermBackupsLifecycleStateEnum = "DELETE_SCHEDULED"
	ListLongTermBackupsLifecycleStateDeleting        ListLongTermBackupsLifecycleStateEnum = "DELETING"
	ListLongTermBackupsLifecycleStateDeleted         ListLongTermBackupsLifecycleStateEnum = "DELETED"
	ListLongTermBackupsLifecycleStateFailed          ListLongTermBackupsLifecycleStateEnum = "FAILED"
)

var mappingListLongTermBackupsLifecycleStateEnum = map[string]ListLongTermBackupsLifecycleStateEnum{
	"CREATING":         ListLongTermBackupsLifecycleStateCreating,
	"UPDATING":         ListLongTermBackupsLifecycleStateUpdating,
	"ACTIVE":           ListLongTermBackupsLifecycleStateActive,
	"DELETE_SCHEDULED": ListLongTermBackupsLifecycleStateDeleteScheduled,
	"DELETING":         ListLongTermBackupsLifecycleStateDeleting,
	"DELETED":          ListLongTermBackupsLifecycleStateDeleted,
	"FAILED":           ListLongTermBackupsLifecycleStateFailed,
}

var mappingListLongTermBackupsLifecycleStateEnumLowerCase = map[string]ListLongTermBackupsLifecycleStateEnum{
	"creating":         ListLongTermBackupsLifecycleStateCreating,
	"updating":         ListLongTermBackupsLifecycleStateUpdating,
	"active":           ListLongTermBackupsLifecycleStateActive,
	"delete_scheduled": ListLongTermBackupsLifecycleStateDeleteScheduled,
	"deleting":         ListLongTermBackupsLifecycleStateDeleting,
	"deleted":          ListLongTermBackupsLifecycleStateDeleted,
	"failed":           ListLongTermBackupsLifecycleStateFailed,
}

// GetListLongTermBackupsLifecycleStateEnumValues Enumerates the set of values for ListLongTermBackupsLifecycleStateEnum
func GetListLongTermBackupsLifecycleStateEnumValues() []ListLongTermBackupsLifecycleStateEnum {
	values := make([]ListLongTermBackupsLifecycleStateEnum, 0)
	for _, v := range mappingListLongTermBackupsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListLongTermBackupsLifecycleStateEnumStringValues Enumerates the set of values in String for ListLongTermBackupsLifecycleStateEnum
func GetListLongTermBackupsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETE_SCHEDULED",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListLongTermBackupsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLongTermBackupsLifecycleStateEnum(val string) (ListLongTermBackupsLifecycleStateEnum, bool) {
	enum, ok := mappingListLongTermBackupsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLongTermBackupsSortOrderEnum Enum with underlying type: string
type ListLongTermBackupsSortOrderEnum string

// Set of constants representing the allowable values for ListLongTermBackupsSortOrderEnum
const (
	ListLongTermBackupsSortOrderAsc  ListLongTermBackupsSortOrderEnum = "ASC"
	ListLongTermBackupsSortOrderDesc ListLongTermBackupsSortOrderEnum = "DESC"
)

var mappingListLongTermBackupsSortOrderEnum = map[string]ListLongTermBackupsSortOrderEnum{
	"ASC":  ListLongTermBackupsSortOrderAsc,
	"DESC": ListLongTermBackupsSortOrderDesc,
}

var mappingListLongTermBackupsSortOrderEnumLowerCase = map[string]ListLongTermBackupsSortOrderEnum{
	"asc":  ListLongTermBackupsSortOrderAsc,
	"desc": ListLongTermBackupsSortOrderDesc,
}

// GetListLongTermBackupsSortOrderEnumValues Enumerates the set of values for ListLongTermBackupsSortOrderEnum
func GetListLongTermBackupsSortOrderEnumValues() []ListLongTermBackupsSortOrderEnum {
	values := make([]ListLongTermBackupsSortOrderEnum, 0)
	for _, v := range mappingListLongTermBackupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLongTermBackupsSortOrderEnumStringValues Enumerates the set of values in String for ListLongTermBackupsSortOrderEnum
func GetListLongTermBackupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLongTermBackupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLongTermBackupsSortOrderEnum(val string) (ListLongTermBackupsSortOrderEnum, bool) {
	enum, ok := mappingListLongTermBackupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLongTermBackupsSortByEnum Enum with underlying type: string
type ListLongTermBackupsSortByEnum string

// Set of constants representing the allowable values for ListLongTermBackupsSortByEnum
const (
	ListLongTermBackupsSortByTimecreated ListLongTermBackupsSortByEnum = "timeCreated"
	ListLongTermBackupsSortByDisplayname ListLongTermBackupsSortByEnum = "displayName"
)

var mappingListLongTermBackupsSortByEnum = map[string]ListLongTermBackupsSortByEnum{
	"timeCreated": ListLongTermBackupsSortByTimecreated,
	"displayName": ListLongTermBackupsSortByDisplayname,
}

var mappingListLongTermBackupsSortByEnumLowerCase = map[string]ListLongTermBackupsSortByEnum{
	"timecreated": ListLongTermBackupsSortByTimecreated,
	"displayname": ListLongTermBackupsSortByDisplayname,
}

// GetListLongTermBackupsSortByEnumValues Enumerates the set of values for ListLongTermBackupsSortByEnum
func GetListLongTermBackupsSortByEnumValues() []ListLongTermBackupsSortByEnum {
	values := make([]ListLongTermBackupsSortByEnum, 0)
	for _, v := range mappingListLongTermBackupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLongTermBackupsSortByEnumStringValues Enumerates the set of values in String for ListLongTermBackupsSortByEnum
func GetListLongTermBackupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListLongTermBackupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLongTermBackupsSortByEnum(val string) (ListLongTermBackupsSortByEnum, bool) {
	enum, ok := mappingListLongTermBackupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
