// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListNodeBackupsRequest wrapper for the ListNodeBackups operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListNodeBackups.go.html to see an example of how to use ListNodeBackupsRequest.
type ListNodeBackupsRequest struct {

	// The OCID of the cluster.
	BdsInstanceId *string `mandatory:"true" contributesTo:"path" name:"bdsInstanceId"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListNodeBackupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListNodeBackupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The node host name belonged to a node that has a node backup.
	NodeHostName *string `mandatory:"false" contributesTo:"query" name:"nodeHostName"`

	// The state of the Node's Backup.
	LifecycleState NodeBackupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The display name belonged to the node backup.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNodeBackupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNodeBackupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNodeBackupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNodeBackupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNodeBackupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNodeBackupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNodeBackupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNodeBackupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNodeBackupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNodeBackupLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetNodeBackupLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNodeBackupsResponse wrapper for the ListNodeBackups operation
type ListNodeBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []NodeBackupSummary instances
	Items []NodeBackupSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListNodeBackupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNodeBackupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNodeBackupsSortByEnum Enum with underlying type: string
type ListNodeBackupsSortByEnum string

// Set of constants representing the allowable values for ListNodeBackupsSortByEnum
const (
	ListNodeBackupsSortByTimecreated ListNodeBackupsSortByEnum = "timeCreated"
	ListNodeBackupsSortByDisplayname ListNodeBackupsSortByEnum = "displayName"
)

var mappingListNodeBackupsSortByEnum = map[string]ListNodeBackupsSortByEnum{
	"timeCreated": ListNodeBackupsSortByTimecreated,
	"displayName": ListNodeBackupsSortByDisplayname,
}

var mappingListNodeBackupsSortByEnumLowerCase = map[string]ListNodeBackupsSortByEnum{
	"timecreated": ListNodeBackupsSortByTimecreated,
	"displayname": ListNodeBackupsSortByDisplayname,
}

// GetListNodeBackupsSortByEnumValues Enumerates the set of values for ListNodeBackupsSortByEnum
func GetListNodeBackupsSortByEnumValues() []ListNodeBackupsSortByEnum {
	values := make([]ListNodeBackupsSortByEnum, 0)
	for _, v := range mappingListNodeBackupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNodeBackupsSortByEnumStringValues Enumerates the set of values in String for ListNodeBackupsSortByEnum
func GetListNodeBackupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListNodeBackupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNodeBackupsSortByEnum(val string) (ListNodeBackupsSortByEnum, bool) {
	enum, ok := mappingListNodeBackupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNodeBackupsSortOrderEnum Enum with underlying type: string
type ListNodeBackupsSortOrderEnum string

// Set of constants representing the allowable values for ListNodeBackupsSortOrderEnum
const (
	ListNodeBackupsSortOrderAsc  ListNodeBackupsSortOrderEnum = "ASC"
	ListNodeBackupsSortOrderDesc ListNodeBackupsSortOrderEnum = "DESC"
)

var mappingListNodeBackupsSortOrderEnum = map[string]ListNodeBackupsSortOrderEnum{
	"ASC":  ListNodeBackupsSortOrderAsc,
	"DESC": ListNodeBackupsSortOrderDesc,
}

var mappingListNodeBackupsSortOrderEnumLowerCase = map[string]ListNodeBackupsSortOrderEnum{
	"asc":  ListNodeBackupsSortOrderAsc,
	"desc": ListNodeBackupsSortOrderDesc,
}

// GetListNodeBackupsSortOrderEnumValues Enumerates the set of values for ListNodeBackupsSortOrderEnum
func GetListNodeBackupsSortOrderEnumValues() []ListNodeBackupsSortOrderEnum {
	values := make([]ListNodeBackupsSortOrderEnum, 0)
	for _, v := range mappingListNodeBackupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNodeBackupsSortOrderEnumStringValues Enumerates the set of values in String for ListNodeBackupsSortOrderEnum
func GetListNodeBackupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNodeBackupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNodeBackupsSortOrderEnum(val string) (ListNodeBackupsSortOrderEnum, bool) {
	enum, ok := mappingListNodeBackupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
