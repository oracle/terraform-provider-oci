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

// ListPatchHistoriesRequest wrapper for the ListPatchHistories operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListPatchHistories.go.html to see an example of how to use ListPatchHistoriesRequest.
type ListPatchHistoriesRequest struct {

	// The OCID of the cluster.
	BdsInstanceId *string `mandatory:"true" contributesTo:"path" name:"bdsInstanceId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The status of the patch.
	LifecycleState PatchHistorySummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListPatchHistoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The version of the patch
	PatchVersion *string `mandatory:"false" contributesTo:"query" name:"patchVersion"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListPatchHistoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The type of a BDS patch history entity.
	PatchType PatchHistorySummaryPatchTypeEnum `mandatory:"false" contributesTo:"query" name:"patchType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPatchHistoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPatchHistoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPatchHistoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPatchHistoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPatchHistoriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchHistorySummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPatchHistorySummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatchHistoriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPatchHistoriesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatchHistoriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPatchHistoriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchHistorySummaryPatchTypeEnum(string(request.PatchType)); !ok && request.PatchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchType: %s. Supported values are: %s.", request.PatchType, strings.Join(GetPatchHistorySummaryPatchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPatchHistoriesResponse wrapper for the ListPatchHistories operation
type ListPatchHistoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []PatchHistorySummary instances
	Items []PatchHistorySummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListPatchHistoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPatchHistoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPatchHistoriesSortByEnum Enum with underlying type: string
type ListPatchHistoriesSortByEnum string

// Set of constants representing the allowable values for ListPatchHistoriesSortByEnum
const (
	ListPatchHistoriesSortByTimecreated ListPatchHistoriesSortByEnum = "timeCreated"
	ListPatchHistoriesSortByDisplayname ListPatchHistoriesSortByEnum = "displayName"
)

var mappingListPatchHistoriesSortByEnum = map[string]ListPatchHistoriesSortByEnum{
	"timeCreated": ListPatchHistoriesSortByTimecreated,
	"displayName": ListPatchHistoriesSortByDisplayname,
}

var mappingListPatchHistoriesSortByEnumLowerCase = map[string]ListPatchHistoriesSortByEnum{
	"timecreated": ListPatchHistoriesSortByTimecreated,
	"displayname": ListPatchHistoriesSortByDisplayname,
}

// GetListPatchHistoriesSortByEnumValues Enumerates the set of values for ListPatchHistoriesSortByEnum
func GetListPatchHistoriesSortByEnumValues() []ListPatchHistoriesSortByEnum {
	values := make([]ListPatchHistoriesSortByEnum, 0)
	for _, v := range mappingListPatchHistoriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchHistoriesSortByEnumStringValues Enumerates the set of values in String for ListPatchHistoriesSortByEnum
func GetListPatchHistoriesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPatchHistoriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchHistoriesSortByEnum(val string) (ListPatchHistoriesSortByEnum, bool) {
	enum, ok := mappingListPatchHistoriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPatchHistoriesSortOrderEnum Enum with underlying type: string
type ListPatchHistoriesSortOrderEnum string

// Set of constants representing the allowable values for ListPatchHistoriesSortOrderEnum
const (
	ListPatchHistoriesSortOrderAsc  ListPatchHistoriesSortOrderEnum = "ASC"
	ListPatchHistoriesSortOrderDesc ListPatchHistoriesSortOrderEnum = "DESC"
)

var mappingListPatchHistoriesSortOrderEnum = map[string]ListPatchHistoriesSortOrderEnum{
	"ASC":  ListPatchHistoriesSortOrderAsc,
	"DESC": ListPatchHistoriesSortOrderDesc,
}

var mappingListPatchHistoriesSortOrderEnumLowerCase = map[string]ListPatchHistoriesSortOrderEnum{
	"asc":  ListPatchHistoriesSortOrderAsc,
	"desc": ListPatchHistoriesSortOrderDesc,
}

// GetListPatchHistoriesSortOrderEnumValues Enumerates the set of values for ListPatchHistoriesSortOrderEnum
func GetListPatchHistoriesSortOrderEnumValues() []ListPatchHistoriesSortOrderEnum {
	values := make([]ListPatchHistoriesSortOrderEnum, 0)
	for _, v := range mappingListPatchHistoriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchHistoriesSortOrderEnumStringValues Enumerates the set of values in String for ListPatchHistoriesSortOrderEnum
func GetListPatchHistoriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPatchHistoriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchHistoriesSortOrderEnum(val string) (ListPatchHistoriesSortOrderEnum, bool) {
	enum, ok := mappingListPatchHistoriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
