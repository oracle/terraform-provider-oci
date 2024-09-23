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

// ListOnboardingPoliciesRequest wrapper for the ListOnboardingPolicies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListOnboardingPolicies.go.html to see an example of how to use ListOnboardingPoliciesRequest.
type ListOnboardingPoliciesRequest struct {

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOnboardingPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	SortBy ListOnboardingPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOnboardingPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOnboardingPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOnboardingPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOnboardingPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOnboardingPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOnboardingPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOnboardingPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOnboardingPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOnboardingPoliciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOnboardingPoliciesResponse wrapper for the ListOnboardingPolicies operation
type ListOnboardingPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OnboardingPolicyCollection instances
	OnboardingPolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOnboardingPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOnboardingPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOnboardingPoliciesSortOrderEnum Enum with underlying type: string
type ListOnboardingPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListOnboardingPoliciesSortOrderEnum
const (
	ListOnboardingPoliciesSortOrderAsc  ListOnboardingPoliciesSortOrderEnum = "ASC"
	ListOnboardingPoliciesSortOrderDesc ListOnboardingPoliciesSortOrderEnum = "DESC"
)

var mappingListOnboardingPoliciesSortOrderEnum = map[string]ListOnboardingPoliciesSortOrderEnum{
	"ASC":  ListOnboardingPoliciesSortOrderAsc,
	"DESC": ListOnboardingPoliciesSortOrderDesc,
}

var mappingListOnboardingPoliciesSortOrderEnumLowerCase = map[string]ListOnboardingPoliciesSortOrderEnum{
	"asc":  ListOnboardingPoliciesSortOrderAsc,
	"desc": ListOnboardingPoliciesSortOrderDesc,
}

// GetListOnboardingPoliciesSortOrderEnumValues Enumerates the set of values for ListOnboardingPoliciesSortOrderEnum
func GetListOnboardingPoliciesSortOrderEnumValues() []ListOnboardingPoliciesSortOrderEnum {
	values := make([]ListOnboardingPoliciesSortOrderEnum, 0)
	for _, v := range mappingListOnboardingPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOnboardingPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListOnboardingPoliciesSortOrderEnum
func GetListOnboardingPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOnboardingPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOnboardingPoliciesSortOrderEnum(val string) (ListOnboardingPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListOnboardingPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOnboardingPoliciesSortByEnum Enum with underlying type: string
type ListOnboardingPoliciesSortByEnum string

// Set of constants representing the allowable values for ListOnboardingPoliciesSortByEnum
const (
	ListOnboardingPoliciesSortByTimecreated ListOnboardingPoliciesSortByEnum = "timeCreated"
)

var mappingListOnboardingPoliciesSortByEnum = map[string]ListOnboardingPoliciesSortByEnum{
	"timeCreated": ListOnboardingPoliciesSortByTimecreated,
}

var mappingListOnboardingPoliciesSortByEnumLowerCase = map[string]ListOnboardingPoliciesSortByEnum{
	"timecreated": ListOnboardingPoliciesSortByTimecreated,
}

// GetListOnboardingPoliciesSortByEnumValues Enumerates the set of values for ListOnboardingPoliciesSortByEnum
func GetListOnboardingPoliciesSortByEnumValues() []ListOnboardingPoliciesSortByEnum {
	values := make([]ListOnboardingPoliciesSortByEnum, 0)
	for _, v := range mappingListOnboardingPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOnboardingPoliciesSortByEnumStringValues Enumerates the set of values in String for ListOnboardingPoliciesSortByEnum
func GetListOnboardingPoliciesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListOnboardingPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOnboardingPoliciesSortByEnum(val string) (ListOnboardingPoliciesSortByEnum, bool) {
	enum, ok := mappingListOnboardingPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
