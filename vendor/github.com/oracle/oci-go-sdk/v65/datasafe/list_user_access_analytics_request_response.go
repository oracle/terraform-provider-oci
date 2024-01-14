// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListUserAccessAnalyticsRequest wrapper for the ListUserAccessAnalytics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUserAccessAnalytics.go.html to see an example of how to use ListUserAccessAnalyticsRequest.
type ListUserAccessAnalyticsRequest struct {

	// The OCID of the user assessment.
	UserAssessmentId *string `mandatory:"true" contributesTo:"path" name:"userAssessmentId"`

	// The field to sort by. Only one sort parameter may be provided.
	SortBy ListUserAccessAnalyticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListUserAccessAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUserAccessAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUserAccessAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUserAccessAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUserAccessAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUserAccessAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUserAccessAnalyticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUserAccessAnalyticsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUserAccessAnalyticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUserAccessAnalyticsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUserAccessAnalyticsResponse wrapper for the ListUserAccessAnalytics operation
type ListUserAccessAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UserAccessAnalyticsCollection instances
	UserAccessAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListUserAccessAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUserAccessAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUserAccessAnalyticsSortByEnum Enum with underlying type: string
type ListUserAccessAnalyticsSortByEnum string

// Set of constants representing the allowable values for ListUserAccessAnalyticsSortByEnum
const (
	ListUserAccessAnalyticsSortByUsername ListUserAccessAnalyticsSortByEnum = "USERNAME"
	ListUserAccessAnalyticsSortByCount    ListUserAccessAnalyticsSortByEnum = "COUNT"
)

var mappingListUserAccessAnalyticsSortByEnum = map[string]ListUserAccessAnalyticsSortByEnum{
	"USERNAME": ListUserAccessAnalyticsSortByUsername,
	"COUNT":    ListUserAccessAnalyticsSortByCount,
}

var mappingListUserAccessAnalyticsSortByEnumLowerCase = map[string]ListUserAccessAnalyticsSortByEnum{
	"username": ListUserAccessAnalyticsSortByUsername,
	"count":    ListUserAccessAnalyticsSortByCount,
}

// GetListUserAccessAnalyticsSortByEnumValues Enumerates the set of values for ListUserAccessAnalyticsSortByEnum
func GetListUserAccessAnalyticsSortByEnumValues() []ListUserAccessAnalyticsSortByEnum {
	values := make([]ListUserAccessAnalyticsSortByEnum, 0)
	for _, v := range mappingListUserAccessAnalyticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUserAccessAnalyticsSortByEnumStringValues Enumerates the set of values in String for ListUserAccessAnalyticsSortByEnum
func GetListUserAccessAnalyticsSortByEnumStringValues() []string {
	return []string{
		"USERNAME",
		"COUNT",
	}
}

// GetMappingListUserAccessAnalyticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUserAccessAnalyticsSortByEnum(val string) (ListUserAccessAnalyticsSortByEnum, bool) {
	enum, ok := mappingListUserAccessAnalyticsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUserAccessAnalyticsSortOrderEnum Enum with underlying type: string
type ListUserAccessAnalyticsSortOrderEnum string

// Set of constants representing the allowable values for ListUserAccessAnalyticsSortOrderEnum
const (
	ListUserAccessAnalyticsSortOrderAsc  ListUserAccessAnalyticsSortOrderEnum = "ASC"
	ListUserAccessAnalyticsSortOrderDesc ListUserAccessAnalyticsSortOrderEnum = "DESC"
)

var mappingListUserAccessAnalyticsSortOrderEnum = map[string]ListUserAccessAnalyticsSortOrderEnum{
	"ASC":  ListUserAccessAnalyticsSortOrderAsc,
	"DESC": ListUserAccessAnalyticsSortOrderDesc,
}

var mappingListUserAccessAnalyticsSortOrderEnumLowerCase = map[string]ListUserAccessAnalyticsSortOrderEnum{
	"asc":  ListUserAccessAnalyticsSortOrderAsc,
	"desc": ListUserAccessAnalyticsSortOrderDesc,
}

// GetListUserAccessAnalyticsSortOrderEnumValues Enumerates the set of values for ListUserAccessAnalyticsSortOrderEnum
func GetListUserAccessAnalyticsSortOrderEnumValues() []ListUserAccessAnalyticsSortOrderEnum {
	values := make([]ListUserAccessAnalyticsSortOrderEnum, 0)
	for _, v := range mappingListUserAccessAnalyticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUserAccessAnalyticsSortOrderEnumStringValues Enumerates the set of values in String for ListUserAccessAnalyticsSortOrderEnum
func GetListUserAccessAnalyticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUserAccessAnalyticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUserAccessAnalyticsSortOrderEnum(val string) (ListUserAccessAnalyticsSortOrderEnum, bool) {
	enum, ok := mappingListUserAccessAnalyticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
