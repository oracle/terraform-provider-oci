// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMigrationAssetsRequest wrapper for the ListMigrationAssets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudmigrations/ListMigrationAssets.go.html to see an example of how to use ListMigrationAssetsRequest.
type ListMigrationAssetsRequest struct {

	// Unique migration identifier
	MigrationId *string `mandatory:"false" contributesTo:"query" name:"migrationId"`

	// A filter to return only resources that match the entire given display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique migration asset identifier
	MigrationAssetId *string `mandatory:"false" contributesTo:"query" name:"migrationAssetId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of the previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The current state of the migration asset.
	LifecycleState MigrationAssetLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMigrationAssetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. The default order for 'timeCreated' is descending. The default order for 'displayName' is ascending.
	SortBy ListMigrationAssetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMigrationAssetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMigrationAssetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMigrationAssetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMigrationAssetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMigrationAssetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMigrationAssetLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMigrationAssetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMigrationAssetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMigrationAssetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMigrationAssetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMigrationAssetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMigrationAssetsResponse wrapper for the ListMigrationAssets operation
type ListMigrationAssetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MigrationAssetCollection instances
	MigrationAssetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMigrationAssetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMigrationAssetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMigrationAssetsSortOrderEnum Enum with underlying type: string
type ListMigrationAssetsSortOrderEnum string

// Set of constants representing the allowable values for ListMigrationAssetsSortOrderEnum
const (
	ListMigrationAssetsSortOrderAsc  ListMigrationAssetsSortOrderEnum = "ASC"
	ListMigrationAssetsSortOrderDesc ListMigrationAssetsSortOrderEnum = "DESC"
)

var mappingListMigrationAssetsSortOrderEnum = map[string]ListMigrationAssetsSortOrderEnum{
	"ASC":  ListMigrationAssetsSortOrderAsc,
	"DESC": ListMigrationAssetsSortOrderDesc,
}

var mappingListMigrationAssetsSortOrderEnumLowerCase = map[string]ListMigrationAssetsSortOrderEnum{
	"asc":  ListMigrationAssetsSortOrderAsc,
	"desc": ListMigrationAssetsSortOrderDesc,
}

// GetListMigrationAssetsSortOrderEnumValues Enumerates the set of values for ListMigrationAssetsSortOrderEnum
func GetListMigrationAssetsSortOrderEnumValues() []ListMigrationAssetsSortOrderEnum {
	values := make([]ListMigrationAssetsSortOrderEnum, 0)
	for _, v := range mappingListMigrationAssetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationAssetsSortOrderEnumStringValues Enumerates the set of values in String for ListMigrationAssetsSortOrderEnum
func GetListMigrationAssetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMigrationAssetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationAssetsSortOrderEnum(val string) (ListMigrationAssetsSortOrderEnum, bool) {
	enum, ok := mappingListMigrationAssetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMigrationAssetsSortByEnum Enum with underlying type: string
type ListMigrationAssetsSortByEnum string

// Set of constants representing the allowable values for ListMigrationAssetsSortByEnum
const (
	ListMigrationAssetsSortByTimecreated ListMigrationAssetsSortByEnum = "timeCreated"
	ListMigrationAssetsSortByDisplayname ListMigrationAssetsSortByEnum = "displayName"
)

var mappingListMigrationAssetsSortByEnum = map[string]ListMigrationAssetsSortByEnum{
	"timeCreated": ListMigrationAssetsSortByTimecreated,
	"displayName": ListMigrationAssetsSortByDisplayname,
}

var mappingListMigrationAssetsSortByEnumLowerCase = map[string]ListMigrationAssetsSortByEnum{
	"timecreated": ListMigrationAssetsSortByTimecreated,
	"displayname": ListMigrationAssetsSortByDisplayname,
}

// GetListMigrationAssetsSortByEnumValues Enumerates the set of values for ListMigrationAssetsSortByEnum
func GetListMigrationAssetsSortByEnumValues() []ListMigrationAssetsSortByEnum {
	values := make([]ListMigrationAssetsSortByEnum, 0)
	for _, v := range mappingListMigrationAssetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationAssetsSortByEnumStringValues Enumerates the set of values in String for ListMigrationAssetsSortByEnum
func GetListMigrationAssetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMigrationAssetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationAssetsSortByEnum(val string) (ListMigrationAssetsSortByEnum, bool) {
	enum, ok := mappingListMigrationAssetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
