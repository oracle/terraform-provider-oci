// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDbNodeBootVolumeBackupsRequest wrapper for the ListDbNodeBootVolumeBackups operation
type ListDbNodeBootVolumeBackupsRequest struct {

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the boot volume backups that match the given lifecycle state exactly.
	LifecycleState DbNodeBootVolumeBackupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the boot volume backups that match the given database node.
	DbNodeId *string `mandatory:"false" contributesTo:"query" name:"dbNodeId"`

	// A filter to return only the boot volume backups that match the given DB system.
	DbSystemId *string `mandatory:"false" contributesTo:"query" name:"dbSystemId"`

	// A filter to return only the boot volume backup that matches the given OCID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for timeCreated is descending. Default order for displayName is ascending. The displayName sort order is case sensitive.
	SortBy ListDbNodeBootVolumeBackupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDbNodeBootVolumeBackupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbNodeBootVolumeBackupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbNodeBootVolumeBackupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbNodeBootVolumeBackupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbNodeBootVolumeBackupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDbNodeBootVolumeBackupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbNodeBootVolumeBackupLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDbNodeBootVolumeBackupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbNodeBootVolumeBackupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDbNodeBootVolumeBackupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbNodeBootVolumeBackupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDbNodeBootVolumeBackupsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDbNodeBootVolumeBackupsResponse wrapper for the ListDbNodeBootVolumeBackups operation
type ListDbNodeBootVolumeBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DbNodeBootVolumeBackupCollection instances
	DbNodeBootVolumeBackupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbNodeBootVolumeBackupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbNodeBootVolumeBackupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbNodeBootVolumeBackupsSortByEnum Enum with underlying type: string
type ListDbNodeBootVolumeBackupsSortByEnum string

// Set of constants representing the allowable values for ListDbNodeBootVolumeBackupsSortByEnum
const (
	ListDbNodeBootVolumeBackupsSortByTimecreated ListDbNodeBootVolumeBackupsSortByEnum = "timeCreated"
	ListDbNodeBootVolumeBackupsSortByDisplayname ListDbNodeBootVolumeBackupsSortByEnum = "displayName"
)

var mappingListDbNodeBootVolumeBackupsSortByEnum = map[string]ListDbNodeBootVolumeBackupsSortByEnum{
	"timeCreated": ListDbNodeBootVolumeBackupsSortByTimecreated,
	"displayName": ListDbNodeBootVolumeBackupsSortByDisplayname,
}

var mappingListDbNodeBootVolumeBackupsSortByEnumLowerCase = map[string]ListDbNodeBootVolumeBackupsSortByEnum{
	"timecreated": ListDbNodeBootVolumeBackupsSortByTimecreated,
	"displayname": ListDbNodeBootVolumeBackupsSortByDisplayname,
}

// GetListDbNodeBootVolumeBackupsSortByEnumValues Enumerates the set of values for ListDbNodeBootVolumeBackupsSortByEnum
func GetListDbNodeBootVolumeBackupsSortByEnumValues() []ListDbNodeBootVolumeBackupsSortByEnum {
	values := make([]ListDbNodeBootVolumeBackupsSortByEnum, 0)
	for _, v := range mappingListDbNodeBootVolumeBackupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbNodeBootVolumeBackupsSortByEnumStringValues Enumerates the set of values in String for ListDbNodeBootVolumeBackupsSortByEnum
func GetListDbNodeBootVolumeBackupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDbNodeBootVolumeBackupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbNodeBootVolumeBackupsSortByEnum(val string) (ListDbNodeBootVolumeBackupsSortByEnum, bool) {
	enum, ok := mappingListDbNodeBootVolumeBackupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbNodeBootVolumeBackupsSortOrderEnum Enum with underlying type: string
type ListDbNodeBootVolumeBackupsSortOrderEnum string

// Set of constants representing the allowable values for ListDbNodeBootVolumeBackupsSortOrderEnum
const (
	ListDbNodeBootVolumeBackupsSortOrderAsc  ListDbNodeBootVolumeBackupsSortOrderEnum = "ASC"
	ListDbNodeBootVolumeBackupsSortOrderDesc ListDbNodeBootVolumeBackupsSortOrderEnum = "DESC"
)

var mappingListDbNodeBootVolumeBackupsSortOrderEnum = map[string]ListDbNodeBootVolumeBackupsSortOrderEnum{
	"ASC":  ListDbNodeBootVolumeBackupsSortOrderAsc,
	"DESC": ListDbNodeBootVolumeBackupsSortOrderDesc,
}

var mappingListDbNodeBootVolumeBackupsSortOrderEnumLowerCase = map[string]ListDbNodeBootVolumeBackupsSortOrderEnum{
	"asc":  ListDbNodeBootVolumeBackupsSortOrderAsc,
	"desc": ListDbNodeBootVolumeBackupsSortOrderDesc,
}

// GetListDbNodeBootVolumeBackupsSortOrderEnumValues Enumerates the set of values for ListDbNodeBootVolumeBackupsSortOrderEnum
func GetListDbNodeBootVolumeBackupsSortOrderEnumValues() []ListDbNodeBootVolumeBackupsSortOrderEnum {
	values := make([]ListDbNodeBootVolumeBackupsSortOrderEnum, 0)
	for _, v := range mappingListDbNodeBootVolumeBackupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbNodeBootVolumeBackupsSortOrderEnumStringValues Enumerates the set of values in String for ListDbNodeBootVolumeBackupsSortOrderEnum
func GetListDbNodeBootVolumeBackupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDbNodeBootVolumeBackupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbNodeBootVolumeBackupsSortOrderEnum(val string) (ListDbNodeBootVolumeBackupsSortOrderEnum, bool) {
	enum, ok := mappingListDbNodeBootVolumeBackupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
