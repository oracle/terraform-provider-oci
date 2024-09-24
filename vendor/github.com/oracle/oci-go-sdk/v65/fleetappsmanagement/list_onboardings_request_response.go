// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOnboardingsRequest wrapper for the ListOnboardings operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListOnboardings.go.html to see an example of how to use ListOnboardingsRequest.
type ListOnboardingsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState OnboardingLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// unique onboarding identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOnboardingsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListOnboardingsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOnboardingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOnboardingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOnboardingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOnboardingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOnboardingsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOnboardingLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOnboardingLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOnboardingsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOnboardingsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOnboardingsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOnboardingsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOnboardingsResponse wrapper for the ListOnboardings operation
type ListOnboardingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OnboardingCollection instances
	OnboardingCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOnboardingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOnboardingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOnboardingsSortOrderEnum Enum with underlying type: string
type ListOnboardingsSortOrderEnum string

// Set of constants representing the allowable values for ListOnboardingsSortOrderEnum
const (
	ListOnboardingsSortOrderAsc  ListOnboardingsSortOrderEnum = "ASC"
	ListOnboardingsSortOrderDesc ListOnboardingsSortOrderEnum = "DESC"
)

var mappingListOnboardingsSortOrderEnum = map[string]ListOnboardingsSortOrderEnum{
	"ASC":  ListOnboardingsSortOrderAsc,
	"DESC": ListOnboardingsSortOrderDesc,
}

var mappingListOnboardingsSortOrderEnumLowerCase = map[string]ListOnboardingsSortOrderEnum{
	"asc":  ListOnboardingsSortOrderAsc,
	"desc": ListOnboardingsSortOrderDesc,
}

// GetListOnboardingsSortOrderEnumValues Enumerates the set of values for ListOnboardingsSortOrderEnum
func GetListOnboardingsSortOrderEnumValues() []ListOnboardingsSortOrderEnum {
	values := make([]ListOnboardingsSortOrderEnum, 0)
	for _, v := range mappingListOnboardingsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOnboardingsSortOrderEnumStringValues Enumerates the set of values in String for ListOnboardingsSortOrderEnum
func GetListOnboardingsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOnboardingsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOnboardingsSortOrderEnum(val string) (ListOnboardingsSortOrderEnum, bool) {
	enum, ok := mappingListOnboardingsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOnboardingsSortByEnum Enum with underlying type: string
type ListOnboardingsSortByEnum string

// Set of constants representing the allowable values for ListOnboardingsSortByEnum
const (
	ListOnboardingsSortByTimecreated ListOnboardingsSortByEnum = "timeCreated"
	ListOnboardingsSortByDisplayname ListOnboardingsSortByEnum = "displayName"
)

var mappingListOnboardingsSortByEnum = map[string]ListOnboardingsSortByEnum{
	"timeCreated": ListOnboardingsSortByTimecreated,
	"displayName": ListOnboardingsSortByDisplayname,
}

var mappingListOnboardingsSortByEnumLowerCase = map[string]ListOnboardingsSortByEnum{
	"timecreated": ListOnboardingsSortByTimecreated,
	"displayname": ListOnboardingsSortByDisplayname,
}

// GetListOnboardingsSortByEnumValues Enumerates the set of values for ListOnboardingsSortByEnum
func GetListOnboardingsSortByEnumValues() []ListOnboardingsSortByEnum {
	values := make([]ListOnboardingsSortByEnum, 0)
	for _, v := range mappingListOnboardingsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOnboardingsSortByEnumStringValues Enumerates the set of values in String for ListOnboardingsSortByEnum
func GetListOnboardingsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOnboardingsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOnboardingsSortByEnum(val string) (ListOnboardingsSortByEnum, bool) {
	enum, ok := mappingListOnboardingsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
