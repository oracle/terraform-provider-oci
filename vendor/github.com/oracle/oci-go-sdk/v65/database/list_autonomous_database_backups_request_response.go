// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAutonomousDatabaseBackupsRequest wrapper for the ListAutonomousDatabaseBackups operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousDatabaseBackups.go.html to see an example of how to use ListAutonomousDatabaseBackupsRequest.
type ListAutonomousDatabaseBackupsRequest struct {

	// The database OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	AutonomousDatabaseId *string `mandatory:"false" contributesTo:"query" name:"autonomousDatabaseId"`

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	// **Note:** If you do not include the availability domain filter, the resources are grouped by availability domain, then sorted.
	SortBy ListAutonomousDatabaseBackupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAutonomousDatabaseBackupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState AutonomousDatabaseBackupSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only backups that matches with the given type of Backup.
	Type *string `mandatory:"false" contributesTo:"query" name:"type"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutonomousDatabaseBackupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutonomousDatabaseBackupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutonomousDatabaseBackupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutonomousDatabaseBackupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutonomousDatabaseBackupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAutonomousDatabaseBackupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAutonomousDatabaseBackupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutonomousDatabaseBackupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAutonomousDatabaseBackupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseBackupSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAutonomousDatabaseBackupSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutonomousDatabaseBackupsResponse wrapper for the ListAutonomousDatabaseBackups operation
type ListAutonomousDatabaseBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AutonomousDatabaseBackupSummary instances
	Items []AutonomousDatabaseBackupSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutonomousDatabaseBackupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutonomousDatabaseBackupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutonomousDatabaseBackupsSortByEnum Enum with underlying type: string
type ListAutonomousDatabaseBackupsSortByEnum string

// Set of constants representing the allowable values for ListAutonomousDatabaseBackupsSortByEnum
const (
	ListAutonomousDatabaseBackupsSortByTimecreated ListAutonomousDatabaseBackupsSortByEnum = "TIMECREATED"
	ListAutonomousDatabaseBackupsSortByDisplayname ListAutonomousDatabaseBackupsSortByEnum = "DISPLAYNAME"
)

var mappingListAutonomousDatabaseBackupsSortByEnum = map[string]ListAutonomousDatabaseBackupsSortByEnum{
	"TIMECREATED": ListAutonomousDatabaseBackupsSortByTimecreated,
	"DISPLAYNAME": ListAutonomousDatabaseBackupsSortByDisplayname,
}

var mappingListAutonomousDatabaseBackupsSortByEnumLowerCase = map[string]ListAutonomousDatabaseBackupsSortByEnum{
	"timecreated": ListAutonomousDatabaseBackupsSortByTimecreated,
	"displayname": ListAutonomousDatabaseBackupsSortByDisplayname,
}

// GetListAutonomousDatabaseBackupsSortByEnumValues Enumerates the set of values for ListAutonomousDatabaseBackupsSortByEnum
func GetListAutonomousDatabaseBackupsSortByEnumValues() []ListAutonomousDatabaseBackupsSortByEnum {
	values := make([]ListAutonomousDatabaseBackupsSortByEnum, 0)
	for _, v := range mappingListAutonomousDatabaseBackupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousDatabaseBackupsSortByEnumStringValues Enumerates the set of values in String for ListAutonomousDatabaseBackupsSortByEnum
func GetListAutonomousDatabaseBackupsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAutonomousDatabaseBackupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousDatabaseBackupsSortByEnum(val string) (ListAutonomousDatabaseBackupsSortByEnum, bool) {
	enum, ok := mappingListAutonomousDatabaseBackupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutonomousDatabaseBackupsSortOrderEnum Enum with underlying type: string
type ListAutonomousDatabaseBackupsSortOrderEnum string

// Set of constants representing the allowable values for ListAutonomousDatabaseBackupsSortOrderEnum
const (
	ListAutonomousDatabaseBackupsSortOrderAsc  ListAutonomousDatabaseBackupsSortOrderEnum = "ASC"
	ListAutonomousDatabaseBackupsSortOrderDesc ListAutonomousDatabaseBackupsSortOrderEnum = "DESC"
)

var mappingListAutonomousDatabaseBackupsSortOrderEnum = map[string]ListAutonomousDatabaseBackupsSortOrderEnum{
	"ASC":  ListAutonomousDatabaseBackupsSortOrderAsc,
	"DESC": ListAutonomousDatabaseBackupsSortOrderDesc,
}

var mappingListAutonomousDatabaseBackupsSortOrderEnumLowerCase = map[string]ListAutonomousDatabaseBackupsSortOrderEnum{
	"asc":  ListAutonomousDatabaseBackupsSortOrderAsc,
	"desc": ListAutonomousDatabaseBackupsSortOrderDesc,
}

// GetListAutonomousDatabaseBackupsSortOrderEnumValues Enumerates the set of values for ListAutonomousDatabaseBackupsSortOrderEnum
func GetListAutonomousDatabaseBackupsSortOrderEnumValues() []ListAutonomousDatabaseBackupsSortOrderEnum {
	values := make([]ListAutonomousDatabaseBackupsSortOrderEnum, 0)
	for _, v := range mappingListAutonomousDatabaseBackupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousDatabaseBackupsSortOrderEnumStringValues Enumerates the set of values in String for ListAutonomousDatabaseBackupsSortOrderEnum
func GetListAutonomousDatabaseBackupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAutonomousDatabaseBackupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousDatabaseBackupsSortOrderEnum(val string) (ListAutonomousDatabaseBackupsSortOrderEnum, bool) {
	enum, ok := mappingListAutonomousDatabaseBackupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
