// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTargetAssetsRequest wrapper for the ListTargetAssets operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudmigrations/ListTargetAssets.go.html to see an example of how to use ListTargetAssetsRequest.
type ListTargetAssetsRequest struct {

	// Unique migration plan identifier
	MigrationPlanId *string `mandatory:"false" contributesTo:"query" name:"migrationPlanId"`

	// A filter to return only resources that match the entire given display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique target asset identifier
	TargetAssetId *string `mandatory:"false" contributesTo:"query" name:"targetAssetId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of the previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The current state of the target asset.
	LifecycleState TargetAssetLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListTargetAssetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. The default order for 'timeCreated' is descending. The default order for 'displayName' is ascending.
	SortBy ListTargetAssetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetAssetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetAssetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetAssetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetAssetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetAssetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTargetAssetLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetTargetAssetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetAssetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetAssetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetAssetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetAssetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTargetAssetsResponse wrapper for the ListTargetAssets operation
type ListTargetAssetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetAssetCollection instances
	TargetAssetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTargetAssetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetAssetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetAssetsSortOrderEnum Enum with underlying type: string
type ListTargetAssetsSortOrderEnum string

// Set of constants representing the allowable values for ListTargetAssetsSortOrderEnum
const (
	ListTargetAssetsSortOrderAsc  ListTargetAssetsSortOrderEnum = "ASC"
	ListTargetAssetsSortOrderDesc ListTargetAssetsSortOrderEnum = "DESC"
)

var mappingListTargetAssetsSortOrderEnum = map[string]ListTargetAssetsSortOrderEnum{
	"ASC":  ListTargetAssetsSortOrderAsc,
	"DESC": ListTargetAssetsSortOrderDesc,
}

var mappingListTargetAssetsSortOrderEnumLowerCase = map[string]ListTargetAssetsSortOrderEnum{
	"asc":  ListTargetAssetsSortOrderAsc,
	"desc": ListTargetAssetsSortOrderDesc,
}

// GetListTargetAssetsSortOrderEnumValues Enumerates the set of values for ListTargetAssetsSortOrderEnum
func GetListTargetAssetsSortOrderEnumValues() []ListTargetAssetsSortOrderEnum {
	values := make([]ListTargetAssetsSortOrderEnum, 0)
	for _, v := range mappingListTargetAssetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetAssetsSortOrderEnumStringValues Enumerates the set of values in String for ListTargetAssetsSortOrderEnum
func GetListTargetAssetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetAssetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetAssetsSortOrderEnum(val string) (ListTargetAssetsSortOrderEnum, bool) {
	enum, ok := mappingListTargetAssetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetAssetsSortByEnum Enum with underlying type: string
type ListTargetAssetsSortByEnum string

// Set of constants representing the allowable values for ListTargetAssetsSortByEnum
const (
	ListTargetAssetsSortByTimecreated ListTargetAssetsSortByEnum = "timeCreated"
	ListTargetAssetsSortByDisplayname ListTargetAssetsSortByEnum = "displayName"
)

var mappingListTargetAssetsSortByEnum = map[string]ListTargetAssetsSortByEnum{
	"timeCreated": ListTargetAssetsSortByTimecreated,
	"displayName": ListTargetAssetsSortByDisplayname,
}

var mappingListTargetAssetsSortByEnumLowerCase = map[string]ListTargetAssetsSortByEnum{
	"timecreated": ListTargetAssetsSortByTimecreated,
	"displayname": ListTargetAssetsSortByDisplayname,
}

// GetListTargetAssetsSortByEnumValues Enumerates the set of values for ListTargetAssetsSortByEnum
func GetListTargetAssetsSortByEnumValues() []ListTargetAssetsSortByEnum {
	values := make([]ListTargetAssetsSortByEnum, 0)
	for _, v := range mappingListTargetAssetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetAssetsSortByEnumStringValues Enumerates the set of values in String for ListTargetAssetsSortByEnum
func GetListTargetAssetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListTargetAssetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetAssetsSortByEnum(val string) (ListTargetAssetsSortByEnum, bool) {
	enum, ok := mappingListTargetAssetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
