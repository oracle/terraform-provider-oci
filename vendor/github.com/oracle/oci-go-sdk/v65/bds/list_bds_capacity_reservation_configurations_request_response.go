// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBdsCapacityReservationConfigurationsRequest wrapper for the ListBdsCapacityReservationConfigurations operation
type ListBdsCapacityReservationConfigurationsRequest struct {

	// The OCID of the cluster.
	BdsInstanceId *string `mandatory:"true" contributesTo:"path" name:"bdsInstanceId"`

	// The lifecycle state of the BDS capacity reservation configuration.
	LifecycleState BdsCapacityReservationConfigurationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListBdsCapacityReservationConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListBdsCapacityReservationConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBdsCapacityReservationConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBdsCapacityReservationConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBdsCapacityReservationConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBdsCapacityReservationConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBdsCapacityReservationConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsCapacityReservationConfigurationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBdsCapacityReservationConfigurationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBdsCapacityReservationConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBdsCapacityReservationConfigurationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBdsCapacityReservationConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBdsCapacityReservationConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBdsCapacityReservationConfigurationsResponse wrapper for the ListBdsCapacityReservationConfigurations operation
type ListBdsCapacityReservationConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BdsCapacityReservationConfigurationCollection instances
	BdsCapacityReservationConfigurationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBdsCapacityReservationConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBdsCapacityReservationConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBdsCapacityReservationConfigurationsSortByEnum Enum with underlying type: string
type ListBdsCapacityReservationConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListBdsCapacityReservationConfigurationsSortByEnum
const (
	ListBdsCapacityReservationConfigurationsSortByTimecreated ListBdsCapacityReservationConfigurationsSortByEnum = "timeCreated"
	ListBdsCapacityReservationConfigurationsSortByDisplayname ListBdsCapacityReservationConfigurationsSortByEnum = "displayName"
)

var mappingListBdsCapacityReservationConfigurationsSortByEnum = map[string]ListBdsCapacityReservationConfigurationsSortByEnum{
	"timeCreated": ListBdsCapacityReservationConfigurationsSortByTimecreated,
	"displayName": ListBdsCapacityReservationConfigurationsSortByDisplayname,
}

var mappingListBdsCapacityReservationConfigurationsSortByEnumLowerCase = map[string]ListBdsCapacityReservationConfigurationsSortByEnum{
	"timecreated": ListBdsCapacityReservationConfigurationsSortByTimecreated,
	"displayname": ListBdsCapacityReservationConfigurationsSortByDisplayname,
}

// GetListBdsCapacityReservationConfigurationsSortByEnumValues Enumerates the set of values for ListBdsCapacityReservationConfigurationsSortByEnum
func GetListBdsCapacityReservationConfigurationsSortByEnumValues() []ListBdsCapacityReservationConfigurationsSortByEnum {
	values := make([]ListBdsCapacityReservationConfigurationsSortByEnum, 0)
	for _, v := range mappingListBdsCapacityReservationConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBdsCapacityReservationConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListBdsCapacityReservationConfigurationsSortByEnum
func GetListBdsCapacityReservationConfigurationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListBdsCapacityReservationConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBdsCapacityReservationConfigurationsSortByEnum(val string) (ListBdsCapacityReservationConfigurationsSortByEnum, bool) {
	enum, ok := mappingListBdsCapacityReservationConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBdsCapacityReservationConfigurationsSortOrderEnum Enum with underlying type: string
type ListBdsCapacityReservationConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListBdsCapacityReservationConfigurationsSortOrderEnum
const (
	ListBdsCapacityReservationConfigurationsSortOrderAsc  ListBdsCapacityReservationConfigurationsSortOrderEnum = "ASC"
	ListBdsCapacityReservationConfigurationsSortOrderDesc ListBdsCapacityReservationConfigurationsSortOrderEnum = "DESC"
)

var mappingListBdsCapacityReservationConfigurationsSortOrderEnum = map[string]ListBdsCapacityReservationConfigurationsSortOrderEnum{
	"ASC":  ListBdsCapacityReservationConfigurationsSortOrderAsc,
	"DESC": ListBdsCapacityReservationConfigurationsSortOrderDesc,
}

var mappingListBdsCapacityReservationConfigurationsSortOrderEnumLowerCase = map[string]ListBdsCapacityReservationConfigurationsSortOrderEnum{
	"asc":  ListBdsCapacityReservationConfigurationsSortOrderAsc,
	"desc": ListBdsCapacityReservationConfigurationsSortOrderDesc,
}

// GetListBdsCapacityReservationConfigurationsSortOrderEnumValues Enumerates the set of values for ListBdsCapacityReservationConfigurationsSortOrderEnum
func GetListBdsCapacityReservationConfigurationsSortOrderEnumValues() []ListBdsCapacityReservationConfigurationsSortOrderEnum {
	values := make([]ListBdsCapacityReservationConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListBdsCapacityReservationConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBdsCapacityReservationConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListBdsCapacityReservationConfigurationsSortOrderEnum
func GetListBdsCapacityReservationConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBdsCapacityReservationConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBdsCapacityReservationConfigurationsSortOrderEnum(val string) (ListBdsCapacityReservationConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListBdsCapacityReservationConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
