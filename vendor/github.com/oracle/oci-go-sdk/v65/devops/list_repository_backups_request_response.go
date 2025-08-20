// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRepositoryBackupsRequest wrapper for the ListRepositoryBackups operation
type ListRepositoryBackupsRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID of the repository in which to list resources.
	RepositoryId *string `mandatory:"false" contributesTo:"query" name:"repositoryId"`

	// A filter used to return the backups based on the repository backup settings id.
	RepositoryBackupSettingsId *string `mandatory:"false" contributesTo:"query" name:"repositoryBackupSettingsId"`

	// Unique identifier or OCID for listing a single resource by ID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources whose lifecycle state matches the given lifecycle state.
	LifecycleState RepositoryBackupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListRepositoryBackupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.  If no value is specified timeCreated is default.
	SortBy ListRepositoryBackupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRepositoryBackupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRepositoryBackupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRepositoryBackupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRepositoryBackupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRepositoryBackupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRepositoryBackupLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRepositoryBackupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRepositoryBackupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRepositoryBackupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRepositoryBackupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRepositoryBackupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRepositoryBackupsResponse wrapper for the ListRepositoryBackups operation
type ListRepositoryBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RepositoryBackupCollection instances
	RepositoryBackupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRepositoryBackupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRepositoryBackupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRepositoryBackupsSortOrderEnum Enum with underlying type: string
type ListRepositoryBackupsSortOrderEnum string

// Set of constants representing the allowable values for ListRepositoryBackupsSortOrderEnum
const (
	ListRepositoryBackupsSortOrderAsc  ListRepositoryBackupsSortOrderEnum = "ASC"
	ListRepositoryBackupsSortOrderDesc ListRepositoryBackupsSortOrderEnum = "DESC"
)

var mappingListRepositoryBackupsSortOrderEnum = map[string]ListRepositoryBackupsSortOrderEnum{
	"ASC":  ListRepositoryBackupsSortOrderAsc,
	"DESC": ListRepositoryBackupsSortOrderDesc,
}

var mappingListRepositoryBackupsSortOrderEnumLowerCase = map[string]ListRepositoryBackupsSortOrderEnum{
	"asc":  ListRepositoryBackupsSortOrderAsc,
	"desc": ListRepositoryBackupsSortOrderDesc,
}

// GetListRepositoryBackupsSortOrderEnumValues Enumerates the set of values for ListRepositoryBackupsSortOrderEnum
func GetListRepositoryBackupsSortOrderEnumValues() []ListRepositoryBackupsSortOrderEnum {
	values := make([]ListRepositoryBackupsSortOrderEnum, 0)
	for _, v := range mappingListRepositoryBackupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRepositoryBackupsSortOrderEnumStringValues Enumerates the set of values in String for ListRepositoryBackupsSortOrderEnum
func GetListRepositoryBackupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRepositoryBackupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRepositoryBackupsSortOrderEnum(val string) (ListRepositoryBackupsSortOrderEnum, bool) {
	enum, ok := mappingListRepositoryBackupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRepositoryBackupsSortByEnum Enum with underlying type: string
type ListRepositoryBackupsSortByEnum string

// Set of constants representing the allowable values for ListRepositoryBackupsSortByEnum
const (
	ListRepositoryBackupsSortByTimecreated ListRepositoryBackupsSortByEnum = "timeCreated"
	ListRepositoryBackupsSortByDisplayname ListRepositoryBackupsSortByEnum = "displayName"
)

var mappingListRepositoryBackupsSortByEnum = map[string]ListRepositoryBackupsSortByEnum{
	"timeCreated": ListRepositoryBackupsSortByTimecreated,
	"displayName": ListRepositoryBackupsSortByDisplayname,
}

var mappingListRepositoryBackupsSortByEnumLowerCase = map[string]ListRepositoryBackupsSortByEnum{
	"timecreated": ListRepositoryBackupsSortByTimecreated,
	"displayname": ListRepositoryBackupsSortByDisplayname,
}

// GetListRepositoryBackupsSortByEnumValues Enumerates the set of values for ListRepositoryBackupsSortByEnum
func GetListRepositoryBackupsSortByEnumValues() []ListRepositoryBackupsSortByEnum {
	values := make([]ListRepositoryBackupsSortByEnum, 0)
	for _, v := range mappingListRepositoryBackupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRepositoryBackupsSortByEnumStringValues Enumerates the set of values in String for ListRepositoryBackupsSortByEnum
func GetListRepositoryBackupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListRepositoryBackupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRepositoryBackupsSortByEnum(val string) (ListRepositoryBackupsSortByEnum, bool) {
	enum, ok := mappingListRepositoryBackupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
