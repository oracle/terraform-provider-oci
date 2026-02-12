// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDependentPropertyValuesRequest wrapper for the ListDependentPropertyValues operation
type ListDependentPropertyValuesRequest struct {

	// unique Property identifier
	PropertyId *string `mandatory:"true" contributesTo:"path" name:"propertyId"`

	// Unique Dependent property value identifier
	DependentPropertyValueId *string `mandatory:"false" contributesTo:"query" name:"dependentPropertyValueId"`

	// The ID of the compartment in which to list resources.
	// Empty only if the resource OCID query param is not specified.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState DependentPropertyValueLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListDependentPropertyValuesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListDependentPropertyValuesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDependentPropertyValuesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDependentPropertyValuesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDependentPropertyValuesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDependentPropertyValuesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDependentPropertyValuesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDependentPropertyValueLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDependentPropertyValueLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDependentPropertyValuesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDependentPropertyValuesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDependentPropertyValuesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDependentPropertyValuesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDependentPropertyValuesResponse wrapper for the ListDependentPropertyValues operation
type ListDependentPropertyValuesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DependentPropertyValueCollection instances
	DependentPropertyValueCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDependentPropertyValuesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDependentPropertyValuesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDependentPropertyValuesSortOrderEnum Enum with underlying type: string
type ListDependentPropertyValuesSortOrderEnum string

// Set of constants representing the allowable values for ListDependentPropertyValuesSortOrderEnum
const (
	ListDependentPropertyValuesSortOrderAsc  ListDependentPropertyValuesSortOrderEnum = "ASC"
	ListDependentPropertyValuesSortOrderDesc ListDependentPropertyValuesSortOrderEnum = "DESC"
)

var mappingListDependentPropertyValuesSortOrderEnum = map[string]ListDependentPropertyValuesSortOrderEnum{
	"ASC":  ListDependentPropertyValuesSortOrderAsc,
	"DESC": ListDependentPropertyValuesSortOrderDesc,
}

var mappingListDependentPropertyValuesSortOrderEnumLowerCase = map[string]ListDependentPropertyValuesSortOrderEnum{
	"asc":  ListDependentPropertyValuesSortOrderAsc,
	"desc": ListDependentPropertyValuesSortOrderDesc,
}

// GetListDependentPropertyValuesSortOrderEnumValues Enumerates the set of values for ListDependentPropertyValuesSortOrderEnum
func GetListDependentPropertyValuesSortOrderEnumValues() []ListDependentPropertyValuesSortOrderEnum {
	values := make([]ListDependentPropertyValuesSortOrderEnum, 0)
	for _, v := range mappingListDependentPropertyValuesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDependentPropertyValuesSortOrderEnumStringValues Enumerates the set of values in String for ListDependentPropertyValuesSortOrderEnum
func GetListDependentPropertyValuesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDependentPropertyValuesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDependentPropertyValuesSortOrderEnum(val string) (ListDependentPropertyValuesSortOrderEnum, bool) {
	enum, ok := mappingListDependentPropertyValuesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDependentPropertyValuesSortByEnum Enum with underlying type: string
type ListDependentPropertyValuesSortByEnum string

// Set of constants representing the allowable values for ListDependentPropertyValuesSortByEnum
const (
	ListDependentPropertyValuesSortByTimecreated ListDependentPropertyValuesSortByEnum = "timeCreated"
	ListDependentPropertyValuesSortByDisplayname ListDependentPropertyValuesSortByEnum = "displayName"
)

var mappingListDependentPropertyValuesSortByEnum = map[string]ListDependentPropertyValuesSortByEnum{
	"timeCreated": ListDependentPropertyValuesSortByTimecreated,
	"displayName": ListDependentPropertyValuesSortByDisplayname,
}

var mappingListDependentPropertyValuesSortByEnumLowerCase = map[string]ListDependentPropertyValuesSortByEnum{
	"timecreated": ListDependentPropertyValuesSortByTimecreated,
	"displayname": ListDependentPropertyValuesSortByDisplayname,
}

// GetListDependentPropertyValuesSortByEnumValues Enumerates the set of values for ListDependentPropertyValuesSortByEnum
func GetListDependentPropertyValuesSortByEnumValues() []ListDependentPropertyValuesSortByEnum {
	values := make([]ListDependentPropertyValuesSortByEnum, 0)
	for _, v := range mappingListDependentPropertyValuesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDependentPropertyValuesSortByEnumStringValues Enumerates the set of values in String for ListDependentPropertyValuesSortByEnum
func GetListDependentPropertyValuesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDependentPropertyValuesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDependentPropertyValuesSortByEnum(val string) (ListDependentPropertyValuesSortByEnum, bool) {
	enum, ok := mappingListDependentPropertyValuesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
