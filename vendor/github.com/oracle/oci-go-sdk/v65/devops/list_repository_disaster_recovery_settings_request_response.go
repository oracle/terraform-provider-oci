// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRepositoryDisasterRecoverySettingsRequest wrapper for the ListRepositoryDisasterRecoverySettings operation
type ListRepositoryDisasterRecoverySettingsRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier or OCID for listing a single resource by ID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to the specific disaster recovery setting.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources whose lifecycle state matches the given lifecycle state.
	LifecycleState RepositoryDisasterRecoverySettingsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListRepositoryDisasterRecoverySettingsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for display name is ascending. If no value is specified, then the default time created value is considered.
	SortBy ListRepositoryDisasterRecoverySettingsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRepositoryDisasterRecoverySettingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRepositoryDisasterRecoverySettingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRepositoryDisasterRecoverySettingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRepositoryDisasterRecoverySettingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRepositoryDisasterRecoverySettingsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRepositoryDisasterRecoverySettingsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRepositoryDisasterRecoverySettingsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRepositoryDisasterRecoverySettingsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRepositoryDisasterRecoverySettingsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRepositoryDisasterRecoverySettingsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRepositoryDisasterRecoverySettingsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRepositoryDisasterRecoverySettingsResponse wrapper for the ListRepositoryDisasterRecoverySettings operation
type ListRepositoryDisasterRecoverySettingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RepositoryDisasterRecoverySettingsCollection instances
	RepositoryDisasterRecoverySettingsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRepositoryDisasterRecoverySettingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRepositoryDisasterRecoverySettingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRepositoryDisasterRecoverySettingsSortOrderEnum Enum with underlying type: string
type ListRepositoryDisasterRecoverySettingsSortOrderEnum string

// Set of constants representing the allowable values for ListRepositoryDisasterRecoverySettingsSortOrderEnum
const (
	ListRepositoryDisasterRecoverySettingsSortOrderAsc  ListRepositoryDisasterRecoverySettingsSortOrderEnum = "ASC"
	ListRepositoryDisasterRecoverySettingsSortOrderDesc ListRepositoryDisasterRecoverySettingsSortOrderEnum = "DESC"
)

var mappingListRepositoryDisasterRecoverySettingsSortOrderEnum = map[string]ListRepositoryDisasterRecoverySettingsSortOrderEnum{
	"ASC":  ListRepositoryDisasterRecoverySettingsSortOrderAsc,
	"DESC": ListRepositoryDisasterRecoverySettingsSortOrderDesc,
}

var mappingListRepositoryDisasterRecoverySettingsSortOrderEnumLowerCase = map[string]ListRepositoryDisasterRecoverySettingsSortOrderEnum{
	"asc":  ListRepositoryDisasterRecoverySettingsSortOrderAsc,
	"desc": ListRepositoryDisasterRecoverySettingsSortOrderDesc,
}

// GetListRepositoryDisasterRecoverySettingsSortOrderEnumValues Enumerates the set of values for ListRepositoryDisasterRecoverySettingsSortOrderEnum
func GetListRepositoryDisasterRecoverySettingsSortOrderEnumValues() []ListRepositoryDisasterRecoverySettingsSortOrderEnum {
	values := make([]ListRepositoryDisasterRecoverySettingsSortOrderEnum, 0)
	for _, v := range mappingListRepositoryDisasterRecoverySettingsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRepositoryDisasterRecoverySettingsSortOrderEnumStringValues Enumerates the set of values in String for ListRepositoryDisasterRecoverySettingsSortOrderEnum
func GetListRepositoryDisasterRecoverySettingsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRepositoryDisasterRecoverySettingsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRepositoryDisasterRecoverySettingsSortOrderEnum(val string) (ListRepositoryDisasterRecoverySettingsSortOrderEnum, bool) {
	enum, ok := mappingListRepositoryDisasterRecoverySettingsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRepositoryDisasterRecoverySettingsSortByEnum Enum with underlying type: string
type ListRepositoryDisasterRecoverySettingsSortByEnum string

// Set of constants representing the allowable values for ListRepositoryDisasterRecoverySettingsSortByEnum
const (
	ListRepositoryDisasterRecoverySettingsSortByTimecreated ListRepositoryDisasterRecoverySettingsSortByEnum = "timeCreated"
	ListRepositoryDisasterRecoverySettingsSortByDisplayname ListRepositoryDisasterRecoverySettingsSortByEnum = "displayName"
)

var mappingListRepositoryDisasterRecoverySettingsSortByEnum = map[string]ListRepositoryDisasterRecoverySettingsSortByEnum{
	"timeCreated": ListRepositoryDisasterRecoverySettingsSortByTimecreated,
	"displayName": ListRepositoryDisasterRecoverySettingsSortByDisplayname,
}

var mappingListRepositoryDisasterRecoverySettingsSortByEnumLowerCase = map[string]ListRepositoryDisasterRecoverySettingsSortByEnum{
	"timecreated": ListRepositoryDisasterRecoverySettingsSortByTimecreated,
	"displayname": ListRepositoryDisasterRecoverySettingsSortByDisplayname,
}

// GetListRepositoryDisasterRecoverySettingsSortByEnumValues Enumerates the set of values for ListRepositoryDisasterRecoverySettingsSortByEnum
func GetListRepositoryDisasterRecoverySettingsSortByEnumValues() []ListRepositoryDisasterRecoverySettingsSortByEnum {
	values := make([]ListRepositoryDisasterRecoverySettingsSortByEnum, 0)
	for _, v := range mappingListRepositoryDisasterRecoverySettingsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRepositoryDisasterRecoverySettingsSortByEnumStringValues Enumerates the set of values in String for ListRepositoryDisasterRecoverySettingsSortByEnum
func GetListRepositoryDisasterRecoverySettingsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListRepositoryDisasterRecoverySettingsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRepositoryDisasterRecoverySettingsSortByEnum(val string) (ListRepositoryDisasterRecoverySettingsSortByEnum, bool) {
	enum, ok := mappingListRepositoryDisasterRecoverySettingsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
