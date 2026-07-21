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

// ListBdsCapacityReservationAssociatedConfigurationsRequest wrapper for the ListBdsCapacityReservationAssociatedConfigurations operation
type ListBdsCapacityReservationAssociatedConfigurationsRequest struct {

	// The OCID of the BDS capacity reservation.
	BdsCapacityReservationId *string `mandatory:"true" contributesTo:"path" name:"bdsCapacityReservationId"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The lifecycle state of the BDS capacity reservation configuration.
	LifecycleState BdsCapacityReservationConfigurationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListBdsCapacityReservationAssociatedConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBdsCapacityReservationAssociatedConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBdsCapacityReservationAssociatedConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBdsCapacityReservationAssociatedConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBdsCapacityReservationAssociatedConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBdsCapacityReservationAssociatedConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsCapacityReservationConfigurationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBdsCapacityReservationConfigurationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBdsCapacityReservationAssociatedConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBdsCapacityReservationAssociatedConfigurationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBdsCapacityReservationAssociatedConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBdsCapacityReservationAssociatedConfigurationsResponse wrapper for the ListBdsCapacityReservationAssociatedConfigurations operation
type ListBdsCapacityReservationAssociatedConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BdsCapacityReservationAssociatedConfigurationCollection instances
	BdsCapacityReservationAssociatedConfigurationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBdsCapacityReservationAssociatedConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBdsCapacityReservationAssociatedConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBdsCapacityReservationAssociatedConfigurationsSortByEnum Enum with underlying type: string
type ListBdsCapacityReservationAssociatedConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListBdsCapacityReservationAssociatedConfigurationsSortByEnum
const (
	ListBdsCapacityReservationAssociatedConfigurationsSortByTimecreated ListBdsCapacityReservationAssociatedConfigurationsSortByEnum = "timeCreated"
	ListBdsCapacityReservationAssociatedConfigurationsSortByDisplayname ListBdsCapacityReservationAssociatedConfigurationsSortByEnum = "displayName"
)

var mappingListBdsCapacityReservationAssociatedConfigurationsSortByEnum = map[string]ListBdsCapacityReservationAssociatedConfigurationsSortByEnum{
	"timeCreated": ListBdsCapacityReservationAssociatedConfigurationsSortByTimecreated,
	"displayName": ListBdsCapacityReservationAssociatedConfigurationsSortByDisplayname,
}

var mappingListBdsCapacityReservationAssociatedConfigurationsSortByEnumLowerCase = map[string]ListBdsCapacityReservationAssociatedConfigurationsSortByEnum{
	"timecreated": ListBdsCapacityReservationAssociatedConfigurationsSortByTimecreated,
	"displayname": ListBdsCapacityReservationAssociatedConfigurationsSortByDisplayname,
}

// GetListBdsCapacityReservationAssociatedConfigurationsSortByEnumValues Enumerates the set of values for ListBdsCapacityReservationAssociatedConfigurationsSortByEnum
func GetListBdsCapacityReservationAssociatedConfigurationsSortByEnumValues() []ListBdsCapacityReservationAssociatedConfigurationsSortByEnum {
	values := make([]ListBdsCapacityReservationAssociatedConfigurationsSortByEnum, 0)
	for _, v := range mappingListBdsCapacityReservationAssociatedConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBdsCapacityReservationAssociatedConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListBdsCapacityReservationAssociatedConfigurationsSortByEnum
func GetListBdsCapacityReservationAssociatedConfigurationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListBdsCapacityReservationAssociatedConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBdsCapacityReservationAssociatedConfigurationsSortByEnum(val string) (ListBdsCapacityReservationAssociatedConfigurationsSortByEnum, bool) {
	enum, ok := mappingListBdsCapacityReservationAssociatedConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum Enum with underlying type: string
type ListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum
const (
	ListBdsCapacityReservationAssociatedConfigurationsSortOrderAsc  ListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum = "ASC"
	ListBdsCapacityReservationAssociatedConfigurationsSortOrderDesc ListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum = "DESC"
)

var mappingListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum = map[string]ListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum{
	"ASC":  ListBdsCapacityReservationAssociatedConfigurationsSortOrderAsc,
	"DESC": ListBdsCapacityReservationAssociatedConfigurationsSortOrderDesc,
}

var mappingListBdsCapacityReservationAssociatedConfigurationsSortOrderEnumLowerCase = map[string]ListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum{
	"asc":  ListBdsCapacityReservationAssociatedConfigurationsSortOrderAsc,
	"desc": ListBdsCapacityReservationAssociatedConfigurationsSortOrderDesc,
}

// GetListBdsCapacityReservationAssociatedConfigurationsSortOrderEnumValues Enumerates the set of values for ListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum
func GetListBdsCapacityReservationAssociatedConfigurationsSortOrderEnumValues() []ListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum {
	values := make([]ListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBdsCapacityReservationAssociatedConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum
func GetListBdsCapacityReservationAssociatedConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum(val string) (ListBdsCapacityReservationAssociatedConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListBdsCapacityReservationAssociatedConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
