// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOciCacheBackupsRequest wrapper for the ListOciCacheBackups operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListOciCacheBackups.go.html to see an example of how to use ListOciCacheBackupsRequest.
type ListOciCacheBackupsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return the OCI Cache Backup resources, whose lifecycle state matches with the given lifecycle state.
	LifecycleState OciCacheBackupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique OCI Cache Backup identifier.
	OciCacheBackupId *string `mandatory:"false" contributesTo:"query" name:"ociCacheBackupId"`

	// A filter to return the OCI Cache Backup resources, whose source cluster ID matches with the given source cluster ID.
	SourceClusterId *string `mandatory:"false" contributesTo:"query" name:"sourceClusterId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOciCacheBackupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListOciCacheBackupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOciCacheBackupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOciCacheBackupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOciCacheBackupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOciCacheBackupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOciCacheBackupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciCacheBackupLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOciCacheBackupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOciCacheBackupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOciCacheBackupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOciCacheBackupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOciCacheBackupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOciCacheBackupsResponse wrapper for the ListOciCacheBackups operation
type ListOciCacheBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OciCacheBackupCollection instances
	OciCacheBackupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOciCacheBackupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOciCacheBackupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOciCacheBackupsSortOrderEnum Enum with underlying type: string
type ListOciCacheBackupsSortOrderEnum string

// Set of constants representing the allowable values for ListOciCacheBackupsSortOrderEnum
const (
	ListOciCacheBackupsSortOrderAsc  ListOciCacheBackupsSortOrderEnum = "ASC"
	ListOciCacheBackupsSortOrderDesc ListOciCacheBackupsSortOrderEnum = "DESC"
)

var mappingListOciCacheBackupsSortOrderEnum = map[string]ListOciCacheBackupsSortOrderEnum{
	"ASC":  ListOciCacheBackupsSortOrderAsc,
	"DESC": ListOciCacheBackupsSortOrderDesc,
}

var mappingListOciCacheBackupsSortOrderEnumLowerCase = map[string]ListOciCacheBackupsSortOrderEnum{
	"asc":  ListOciCacheBackupsSortOrderAsc,
	"desc": ListOciCacheBackupsSortOrderDesc,
}

// GetListOciCacheBackupsSortOrderEnumValues Enumerates the set of values for ListOciCacheBackupsSortOrderEnum
func GetListOciCacheBackupsSortOrderEnumValues() []ListOciCacheBackupsSortOrderEnum {
	values := make([]ListOciCacheBackupsSortOrderEnum, 0)
	for _, v := range mappingListOciCacheBackupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOciCacheBackupsSortOrderEnumStringValues Enumerates the set of values in String for ListOciCacheBackupsSortOrderEnum
func GetListOciCacheBackupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOciCacheBackupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOciCacheBackupsSortOrderEnum(val string) (ListOciCacheBackupsSortOrderEnum, bool) {
	enum, ok := mappingListOciCacheBackupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOciCacheBackupsSortByEnum Enum with underlying type: string
type ListOciCacheBackupsSortByEnum string

// Set of constants representing the allowable values for ListOciCacheBackupsSortByEnum
const (
	ListOciCacheBackupsSortByTimecreated ListOciCacheBackupsSortByEnum = "timeCreated"
	ListOciCacheBackupsSortByDisplayname ListOciCacheBackupsSortByEnum = "displayName"
)

var mappingListOciCacheBackupsSortByEnum = map[string]ListOciCacheBackupsSortByEnum{
	"timeCreated": ListOciCacheBackupsSortByTimecreated,
	"displayName": ListOciCacheBackupsSortByDisplayname,
}

var mappingListOciCacheBackupsSortByEnumLowerCase = map[string]ListOciCacheBackupsSortByEnum{
	"timecreated": ListOciCacheBackupsSortByTimecreated,
	"displayname": ListOciCacheBackupsSortByDisplayname,
}

// GetListOciCacheBackupsSortByEnumValues Enumerates the set of values for ListOciCacheBackupsSortByEnum
func GetListOciCacheBackupsSortByEnumValues() []ListOciCacheBackupsSortByEnum {
	values := make([]ListOciCacheBackupsSortByEnum, 0)
	for _, v := range mappingListOciCacheBackupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOciCacheBackupsSortByEnumStringValues Enumerates the set of values in String for ListOciCacheBackupsSortByEnum
func GetListOciCacheBackupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOciCacheBackupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOciCacheBackupsSortByEnum(val string) (ListOciCacheBackupsSortByEnum, bool) {
	enum, ok := mappingListOciCacheBackupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
