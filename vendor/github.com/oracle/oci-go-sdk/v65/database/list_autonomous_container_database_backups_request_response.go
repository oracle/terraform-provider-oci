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

// ListAutonomousContainerDatabaseBackupsRequest wrapper for the ListAutonomousContainerDatabaseBackups operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousContainerDatabaseBackups.go.html to see an example of how to use ListAutonomousContainerDatabaseBackupsRequest.
type ListAutonomousContainerDatabaseBackupsRequest struct {

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The Autonomous Container Database OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	AutonomousContainerDatabaseId *string `mandatory:"false" contributesTo:"query" name:"autonomousContainerDatabaseId"`

	// call for all remote backups
	IsRemote *bool `mandatory:"false" contributesTo:"query" name:"isRemote"`

	// A filter to return only resources that match the given Infrastructure Type.
	InfrastructureType AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum `mandatory:"false" contributesTo:"query" name:"infrastructureType" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	// **Note:** If you do not include the availability domain filter, the resources are grouped by availability domain, then sorted.
	SortBy ListAutonomousContainerDatabaseBackupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAutonomousContainerDatabaseBackupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutonomousContainerDatabaseBackupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutonomousContainerDatabaseBackupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutonomousContainerDatabaseBackupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutonomousContainerDatabaseBackupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutonomousContainerDatabaseBackupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum(string(request.InfrastructureType)); !ok && request.InfrastructureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InfrastructureType: %s. Supported values are: %s.", request.InfrastructureType, strings.Join(GetAutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseBackupSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAutonomousContainerDatabaseBackupSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutonomousContainerDatabaseBackupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAutonomousContainerDatabaseBackupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutonomousContainerDatabaseBackupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAutonomousContainerDatabaseBackupsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutonomousContainerDatabaseBackupsResponse wrapper for the ListAutonomousContainerDatabaseBackups operation
type ListAutonomousContainerDatabaseBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AutonomousContainerDatabaseBackupCollection instances
	AutonomousContainerDatabaseBackupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutonomousContainerDatabaseBackupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutonomousContainerDatabaseBackupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutonomousContainerDatabaseBackupsSortByEnum Enum with underlying type: string
type ListAutonomousContainerDatabaseBackupsSortByEnum string

// Set of constants representing the allowable values for ListAutonomousContainerDatabaseBackupsSortByEnum
const (
	ListAutonomousContainerDatabaseBackupsSortByTimecreated ListAutonomousContainerDatabaseBackupsSortByEnum = "TIMECREATED"
	ListAutonomousContainerDatabaseBackupsSortByDisplayname ListAutonomousContainerDatabaseBackupsSortByEnum = "DISPLAYNAME"
)

var mappingListAutonomousContainerDatabaseBackupsSortByEnum = map[string]ListAutonomousContainerDatabaseBackupsSortByEnum{
	"TIMECREATED": ListAutonomousContainerDatabaseBackupsSortByTimecreated,
	"DISPLAYNAME": ListAutonomousContainerDatabaseBackupsSortByDisplayname,
}

var mappingListAutonomousContainerDatabaseBackupsSortByEnumLowerCase = map[string]ListAutonomousContainerDatabaseBackupsSortByEnum{
	"timecreated": ListAutonomousContainerDatabaseBackupsSortByTimecreated,
	"displayname": ListAutonomousContainerDatabaseBackupsSortByDisplayname,
}

// GetListAutonomousContainerDatabaseBackupsSortByEnumValues Enumerates the set of values for ListAutonomousContainerDatabaseBackupsSortByEnum
func GetListAutonomousContainerDatabaseBackupsSortByEnumValues() []ListAutonomousContainerDatabaseBackupsSortByEnum {
	values := make([]ListAutonomousContainerDatabaseBackupsSortByEnum, 0)
	for _, v := range mappingListAutonomousContainerDatabaseBackupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousContainerDatabaseBackupsSortByEnumStringValues Enumerates the set of values in String for ListAutonomousContainerDatabaseBackupsSortByEnum
func GetListAutonomousContainerDatabaseBackupsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAutonomousContainerDatabaseBackupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousContainerDatabaseBackupsSortByEnum(val string) (ListAutonomousContainerDatabaseBackupsSortByEnum, bool) {
	enum, ok := mappingListAutonomousContainerDatabaseBackupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutonomousContainerDatabaseBackupsSortOrderEnum Enum with underlying type: string
type ListAutonomousContainerDatabaseBackupsSortOrderEnum string

// Set of constants representing the allowable values for ListAutonomousContainerDatabaseBackupsSortOrderEnum
const (
	ListAutonomousContainerDatabaseBackupsSortOrderAsc  ListAutonomousContainerDatabaseBackupsSortOrderEnum = "ASC"
	ListAutonomousContainerDatabaseBackupsSortOrderDesc ListAutonomousContainerDatabaseBackupsSortOrderEnum = "DESC"
)

var mappingListAutonomousContainerDatabaseBackupsSortOrderEnum = map[string]ListAutonomousContainerDatabaseBackupsSortOrderEnum{
	"ASC":  ListAutonomousContainerDatabaseBackupsSortOrderAsc,
	"DESC": ListAutonomousContainerDatabaseBackupsSortOrderDesc,
}

var mappingListAutonomousContainerDatabaseBackupsSortOrderEnumLowerCase = map[string]ListAutonomousContainerDatabaseBackupsSortOrderEnum{
	"asc":  ListAutonomousContainerDatabaseBackupsSortOrderAsc,
	"desc": ListAutonomousContainerDatabaseBackupsSortOrderDesc,
}

// GetListAutonomousContainerDatabaseBackupsSortOrderEnumValues Enumerates the set of values for ListAutonomousContainerDatabaseBackupsSortOrderEnum
func GetListAutonomousContainerDatabaseBackupsSortOrderEnumValues() []ListAutonomousContainerDatabaseBackupsSortOrderEnum {
	values := make([]ListAutonomousContainerDatabaseBackupsSortOrderEnum, 0)
	for _, v := range mappingListAutonomousContainerDatabaseBackupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousContainerDatabaseBackupsSortOrderEnumStringValues Enumerates the set of values in String for ListAutonomousContainerDatabaseBackupsSortOrderEnum
func GetListAutonomousContainerDatabaseBackupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAutonomousContainerDatabaseBackupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousContainerDatabaseBackupsSortOrderEnum(val string) (ListAutonomousContainerDatabaseBackupsSortOrderEnum, bool) {
	enum, ok := mappingListAutonomousContainerDatabaseBackupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
