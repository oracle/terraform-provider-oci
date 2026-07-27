// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDbClusterBackupsRequest wrapper for the ListDbClusterBackups operation
type ListDbClusterBackupsRequest struct {

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The shared-storage DB cluster backup OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DbClusterBackupId *string `mandatory:"false" contributesTo:"query" name:"dbClusterBackupId"`

	// The shared-storage DB cluster backup lifecycle state.
	LifecycleState DbClusterBackupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The shared-storage DB cluster OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DbClusterId *string `mandatory:"false" contributesTo:"query" name:"dbClusterId"`

	// A filter to return only the resource matching the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to use for sorting.
	SortBy ListDbClusterBackupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use (ASC or DESC).
	SortOrder ListDbClusterBackupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated list call. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaasAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from
	// the previous list call. For information about pagination, see List
	// Pagination (https://docs.oracle.com/iaasAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbClusterBackupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbClusterBackupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbClusterBackupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbClusterBackupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDbClusterBackupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbClusterBackupLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDbClusterBackupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbClusterBackupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDbClusterBackupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbClusterBackupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDbClusterBackupsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDbClusterBackupsResponse wrapper for the ListDbClusterBackups operation
type ListDbClusterBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DbClusterBackupCollection instances
	DbClusterBackupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbClusterBackupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbClusterBackupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbClusterBackupsSortByEnum Enum with underlying type: string
type ListDbClusterBackupsSortByEnum string

// Set of constants representing the allowable values for ListDbClusterBackupsSortByEnum
const (
	ListDbClusterBackupsSortByTimecreated ListDbClusterBackupsSortByEnum = "timeCreated"
	ListDbClusterBackupsSortByTimeupdated ListDbClusterBackupsSortByEnum = "timeUpdated"
	ListDbClusterBackupsSortByDisplayname ListDbClusterBackupsSortByEnum = "displayName"
)

var mappingListDbClusterBackupsSortByEnum = map[string]ListDbClusterBackupsSortByEnum{
	"timeCreated": ListDbClusterBackupsSortByTimecreated,
	"timeUpdated": ListDbClusterBackupsSortByTimeupdated,
	"displayName": ListDbClusterBackupsSortByDisplayname,
}

var mappingListDbClusterBackupsSortByEnumLowerCase = map[string]ListDbClusterBackupsSortByEnum{
	"timecreated": ListDbClusterBackupsSortByTimecreated,
	"timeupdated": ListDbClusterBackupsSortByTimeupdated,
	"displayname": ListDbClusterBackupsSortByDisplayname,
}

// GetListDbClusterBackupsSortByEnumValues Enumerates the set of values for ListDbClusterBackupsSortByEnum
func GetListDbClusterBackupsSortByEnumValues() []ListDbClusterBackupsSortByEnum {
	values := make([]ListDbClusterBackupsSortByEnum, 0)
	for _, v := range mappingListDbClusterBackupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbClusterBackupsSortByEnumStringValues Enumerates the set of values in String for ListDbClusterBackupsSortByEnum
func GetListDbClusterBackupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListDbClusterBackupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbClusterBackupsSortByEnum(val string) (ListDbClusterBackupsSortByEnum, bool) {
	enum, ok := mappingListDbClusterBackupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbClusterBackupsSortOrderEnum Enum with underlying type: string
type ListDbClusterBackupsSortOrderEnum string

// Set of constants representing the allowable values for ListDbClusterBackupsSortOrderEnum
const (
	ListDbClusterBackupsSortOrderAsc  ListDbClusterBackupsSortOrderEnum = "ASC"
	ListDbClusterBackupsSortOrderDesc ListDbClusterBackupsSortOrderEnum = "DESC"
)

var mappingListDbClusterBackupsSortOrderEnum = map[string]ListDbClusterBackupsSortOrderEnum{
	"ASC":  ListDbClusterBackupsSortOrderAsc,
	"DESC": ListDbClusterBackupsSortOrderDesc,
}

var mappingListDbClusterBackupsSortOrderEnumLowerCase = map[string]ListDbClusterBackupsSortOrderEnum{
	"asc":  ListDbClusterBackupsSortOrderAsc,
	"desc": ListDbClusterBackupsSortOrderDesc,
}

// GetListDbClusterBackupsSortOrderEnumValues Enumerates the set of values for ListDbClusterBackupsSortOrderEnum
func GetListDbClusterBackupsSortOrderEnumValues() []ListDbClusterBackupsSortOrderEnum {
	values := make([]ListDbClusterBackupsSortOrderEnum, 0)
	for _, v := range mappingListDbClusterBackupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbClusterBackupsSortOrderEnumStringValues Enumerates the set of values in String for ListDbClusterBackupsSortOrderEnum
func GetListDbClusterBackupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDbClusterBackupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbClusterBackupsSortOrderEnum(val string) (ListDbClusterBackupsSortOrderEnum, bool) {
	enum, ok := mappingListDbClusterBackupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
