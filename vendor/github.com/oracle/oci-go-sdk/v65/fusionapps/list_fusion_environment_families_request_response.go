// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFusionEnvironmentFamiliesRequest wrapper for the ListFusionEnvironmentFamilies operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListFusionEnvironmentFamilies.go.html to see an example of how to use ListFusionEnvironmentFamiliesRequest.
type ListFusionEnvironmentFamiliesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ID of the fusion environment family in which to list resources.
	FusionEnvironmentFamilyId *string `mandatory:"false" contributesTo:"query" name:"fusionEnvironmentFamilyId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter that returns all resources that match the specified lifecycle state.
	LifecycleState FusionEnvironmentFamilyLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListFusionEnvironmentFamiliesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListFusionEnvironmentFamiliesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFusionEnvironmentFamiliesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFusionEnvironmentFamiliesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFusionEnvironmentFamiliesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFusionEnvironmentFamiliesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFusionEnvironmentFamiliesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFusionEnvironmentFamilyLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetFusionEnvironmentFamilyLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFusionEnvironmentFamiliesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFusionEnvironmentFamiliesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFusionEnvironmentFamiliesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFusionEnvironmentFamiliesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFusionEnvironmentFamiliesResponse wrapper for the ListFusionEnvironmentFamilies operation
type ListFusionEnvironmentFamiliesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FusionEnvironmentFamilyCollection instances
	FusionEnvironmentFamilyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFusionEnvironmentFamiliesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFusionEnvironmentFamiliesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFusionEnvironmentFamiliesSortOrderEnum Enum with underlying type: string
type ListFusionEnvironmentFamiliesSortOrderEnum string

// Set of constants representing the allowable values for ListFusionEnvironmentFamiliesSortOrderEnum
const (
	ListFusionEnvironmentFamiliesSortOrderAsc  ListFusionEnvironmentFamiliesSortOrderEnum = "ASC"
	ListFusionEnvironmentFamiliesSortOrderDesc ListFusionEnvironmentFamiliesSortOrderEnum = "DESC"
)

var mappingListFusionEnvironmentFamiliesSortOrderEnum = map[string]ListFusionEnvironmentFamiliesSortOrderEnum{
	"ASC":  ListFusionEnvironmentFamiliesSortOrderAsc,
	"DESC": ListFusionEnvironmentFamiliesSortOrderDesc,
}

var mappingListFusionEnvironmentFamiliesSortOrderEnumLowerCase = map[string]ListFusionEnvironmentFamiliesSortOrderEnum{
	"asc":  ListFusionEnvironmentFamiliesSortOrderAsc,
	"desc": ListFusionEnvironmentFamiliesSortOrderDesc,
}

// GetListFusionEnvironmentFamiliesSortOrderEnumValues Enumerates the set of values for ListFusionEnvironmentFamiliesSortOrderEnum
func GetListFusionEnvironmentFamiliesSortOrderEnumValues() []ListFusionEnvironmentFamiliesSortOrderEnum {
	values := make([]ListFusionEnvironmentFamiliesSortOrderEnum, 0)
	for _, v := range mappingListFusionEnvironmentFamiliesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFusionEnvironmentFamiliesSortOrderEnumStringValues Enumerates the set of values in String for ListFusionEnvironmentFamiliesSortOrderEnum
func GetListFusionEnvironmentFamiliesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFusionEnvironmentFamiliesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFusionEnvironmentFamiliesSortOrderEnum(val string) (ListFusionEnvironmentFamiliesSortOrderEnum, bool) {
	enum, ok := mappingListFusionEnvironmentFamiliesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFusionEnvironmentFamiliesSortByEnum Enum with underlying type: string
type ListFusionEnvironmentFamiliesSortByEnum string

// Set of constants representing the allowable values for ListFusionEnvironmentFamiliesSortByEnum
const (
	ListFusionEnvironmentFamiliesSortByTimeCreated ListFusionEnvironmentFamiliesSortByEnum = "TIME_CREATED"
	ListFusionEnvironmentFamiliesSortByDisplayName ListFusionEnvironmentFamiliesSortByEnum = "DISPLAY_NAME"
)

var mappingListFusionEnvironmentFamiliesSortByEnum = map[string]ListFusionEnvironmentFamiliesSortByEnum{
	"TIME_CREATED": ListFusionEnvironmentFamiliesSortByTimeCreated,
	"DISPLAY_NAME": ListFusionEnvironmentFamiliesSortByDisplayName,
}

var mappingListFusionEnvironmentFamiliesSortByEnumLowerCase = map[string]ListFusionEnvironmentFamiliesSortByEnum{
	"time_created": ListFusionEnvironmentFamiliesSortByTimeCreated,
	"display_name": ListFusionEnvironmentFamiliesSortByDisplayName,
}

// GetListFusionEnvironmentFamiliesSortByEnumValues Enumerates the set of values for ListFusionEnvironmentFamiliesSortByEnum
func GetListFusionEnvironmentFamiliesSortByEnumValues() []ListFusionEnvironmentFamiliesSortByEnum {
	values := make([]ListFusionEnvironmentFamiliesSortByEnum, 0)
	for _, v := range mappingListFusionEnvironmentFamiliesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFusionEnvironmentFamiliesSortByEnumStringValues Enumerates the set of values in String for ListFusionEnvironmentFamiliesSortByEnum
func GetListFusionEnvironmentFamiliesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListFusionEnvironmentFamiliesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFusionEnvironmentFamiliesSortByEnum(val string) (ListFusionEnvironmentFamiliesSortByEnum, bool) {
	enum, ok := mappingListFusionEnvironmentFamiliesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
