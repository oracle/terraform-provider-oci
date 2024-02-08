// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLogAnalyticsEmBridgesRequest wrapper for the ListLogAnalyticsEmBridges operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogAnalyticsEmBridges.go.html to see an example of how to use ListLogAnalyticsEmBridgesRequest.
type ListLogAnalyticsEmBridgesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only log analytics enterprise manager bridge name whose name matches the entire name given. The match
	// is case-insensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only log analytics enterprise manager bridges matching all the lifecycle states specified for this parameter.
	LifecycleState []EmBridgeLifecycleStatesEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only log analytics enterprise manager bridges whose lifecycleDetails contains the specified string.
	LifecycleDetailsContains *string `mandatory:"false" contributesTo:"query" name:"lifecycleDetailsContains"`

	// Filter by the processing status of the latest upload from enterprise manager.
	ImportStatus []EmBridgeLatestImportProcessingStatusEnum `contributesTo:"query" name:"importStatus" omitEmpty:"true" collectionFormat:"multi"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLogAnalyticsEmBridgesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort enterprise manager bridges by. Only one sort order may be provided. Default order for timeCreated and timeUpdated
	// is descending. Default order for enterprise manager name is ascending. If no value is specified timeCreated is default.
	SortBy ListLogAnalyticsEmBridgesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogAnalyticsEmBridgesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogAnalyticsEmBridgesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogAnalyticsEmBridgesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogAnalyticsEmBridgesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLogAnalyticsEmBridgesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingEmBridgeLifecycleStatesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetEmBridgeLifecycleStatesEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ImportStatus {
		if _, ok := GetMappingEmBridgeLatestImportProcessingStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImportStatus: %s. Supported values are: %s.", val, strings.Join(GetEmBridgeLatestImportProcessingStatusEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListLogAnalyticsEmBridgesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLogAnalyticsEmBridgesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogAnalyticsEmBridgesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLogAnalyticsEmBridgesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLogAnalyticsEmBridgesResponse wrapper for the ListLogAnalyticsEmBridges operation
type ListLogAnalyticsEmBridgesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsEmBridgeCollection instances
	LogAnalyticsEmBridgeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLogAnalyticsEmBridgesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogAnalyticsEmBridgesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogAnalyticsEmBridgesSortOrderEnum Enum with underlying type: string
type ListLogAnalyticsEmBridgesSortOrderEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEmBridgesSortOrderEnum
const (
	ListLogAnalyticsEmBridgesSortOrderAsc  ListLogAnalyticsEmBridgesSortOrderEnum = "ASC"
	ListLogAnalyticsEmBridgesSortOrderDesc ListLogAnalyticsEmBridgesSortOrderEnum = "DESC"
)

var mappingListLogAnalyticsEmBridgesSortOrderEnum = map[string]ListLogAnalyticsEmBridgesSortOrderEnum{
	"ASC":  ListLogAnalyticsEmBridgesSortOrderAsc,
	"DESC": ListLogAnalyticsEmBridgesSortOrderDesc,
}

var mappingListLogAnalyticsEmBridgesSortOrderEnumLowerCase = map[string]ListLogAnalyticsEmBridgesSortOrderEnum{
	"asc":  ListLogAnalyticsEmBridgesSortOrderAsc,
	"desc": ListLogAnalyticsEmBridgesSortOrderDesc,
}

// GetListLogAnalyticsEmBridgesSortOrderEnumValues Enumerates the set of values for ListLogAnalyticsEmBridgesSortOrderEnum
func GetListLogAnalyticsEmBridgesSortOrderEnumValues() []ListLogAnalyticsEmBridgesSortOrderEnum {
	values := make([]ListLogAnalyticsEmBridgesSortOrderEnum, 0)
	for _, v := range mappingListLogAnalyticsEmBridgesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogAnalyticsEmBridgesSortOrderEnumStringValues Enumerates the set of values in String for ListLogAnalyticsEmBridgesSortOrderEnum
func GetListLogAnalyticsEmBridgesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLogAnalyticsEmBridgesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogAnalyticsEmBridgesSortOrderEnum(val string) (ListLogAnalyticsEmBridgesSortOrderEnum, bool) {
	enum, ok := mappingListLogAnalyticsEmBridgesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogAnalyticsEmBridgesSortByEnum Enum with underlying type: string
type ListLogAnalyticsEmBridgesSortByEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEmBridgesSortByEnum
const (
	ListLogAnalyticsEmBridgesSortByTimecreated ListLogAnalyticsEmBridgesSortByEnum = "timeCreated"
	ListLogAnalyticsEmBridgesSortByTimeupdated ListLogAnalyticsEmBridgesSortByEnum = "timeUpdated"
	ListLogAnalyticsEmBridgesSortByDisplayname ListLogAnalyticsEmBridgesSortByEnum = "displayName"
)

var mappingListLogAnalyticsEmBridgesSortByEnum = map[string]ListLogAnalyticsEmBridgesSortByEnum{
	"timeCreated": ListLogAnalyticsEmBridgesSortByTimecreated,
	"timeUpdated": ListLogAnalyticsEmBridgesSortByTimeupdated,
	"displayName": ListLogAnalyticsEmBridgesSortByDisplayname,
}

var mappingListLogAnalyticsEmBridgesSortByEnumLowerCase = map[string]ListLogAnalyticsEmBridgesSortByEnum{
	"timecreated": ListLogAnalyticsEmBridgesSortByTimecreated,
	"timeupdated": ListLogAnalyticsEmBridgesSortByTimeupdated,
	"displayname": ListLogAnalyticsEmBridgesSortByDisplayname,
}

// GetListLogAnalyticsEmBridgesSortByEnumValues Enumerates the set of values for ListLogAnalyticsEmBridgesSortByEnum
func GetListLogAnalyticsEmBridgesSortByEnumValues() []ListLogAnalyticsEmBridgesSortByEnum {
	values := make([]ListLogAnalyticsEmBridgesSortByEnum, 0)
	for _, v := range mappingListLogAnalyticsEmBridgesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogAnalyticsEmBridgesSortByEnumStringValues Enumerates the set of values in String for ListLogAnalyticsEmBridgesSortByEnum
func GetListLogAnalyticsEmBridgesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListLogAnalyticsEmBridgesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogAnalyticsEmBridgesSortByEnum(val string) (ListLogAnalyticsEmBridgesSortByEnum, bool) {
	enum, ok := mappingListLogAnalyticsEmBridgesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
