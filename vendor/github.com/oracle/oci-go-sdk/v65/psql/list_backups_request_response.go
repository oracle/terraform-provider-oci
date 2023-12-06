// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBackupsRequest wrapper for the ListBackups operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psql/ListBackups.go.html to see an example of how to use ListBackupsRequest.
type ListBackupsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The start date for getting  backups. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStarted"`

	// The End date for getting  backups. An RFC3339 formatted datetime string.
	TimeEnded *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnded"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState BackupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique Backup identifier
	BackupId *string `mandatory:"false" contributesTo:"query" name:"backupId"`

	// unique DbSystem identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListBackupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListBackupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

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
	if _, ok := GetMappingListBackupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBackupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBackupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBackupsSortByEnumStringValues(), ",")))
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

	// A list of BackupCollection instances
	BackupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBackupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBackupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
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

// ListBackupsSortByEnum Enum with underlying type: string
type ListBackupsSortByEnum string

// Set of constants representing the allowable values for ListBackupsSortByEnum
const (
	ListBackupsSortByTimecreated ListBackupsSortByEnum = "timeCreated"
	ListBackupsSortByDisplayname ListBackupsSortByEnum = "displayName"
)

var mappingListBackupsSortByEnum = map[string]ListBackupsSortByEnum{
	"timeCreated": ListBackupsSortByTimecreated,
	"displayName": ListBackupsSortByDisplayname,
}

var mappingListBackupsSortByEnumLowerCase = map[string]ListBackupsSortByEnum{
	"timecreated": ListBackupsSortByTimecreated,
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
		"displayName",
	}
}

// GetMappingListBackupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBackupsSortByEnum(val string) (ListBackupsSortByEnum, bool) {
	enum, ok := mappingListBackupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
