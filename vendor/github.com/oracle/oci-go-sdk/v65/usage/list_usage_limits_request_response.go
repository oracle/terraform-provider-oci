// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package usage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListUsageLimitsRequest wrapper for the ListUsageLimits operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usage/ListUsageLimits.go.html to see an example of how to use ListUsageLimitsRequest.
type ListUsageLimitsRequest struct {

	// The OCID of the root compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The subscription ID for which rewards information is requested for.
	SubscriptionId *string `mandatory:"true" contributesTo:"query" name:"subscriptionId"`

	// Hard or soft limit. Hard limits lead to breaches, soft to alerts.
	LimitType *string `mandatory:"false" contributesTo:"query" name:"limitType"`

	// Resource Name.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// Service Name.
	ServiceType *string `mandatory:"false" contributesTo:"query" name:"serviceType"`

	// Unique, Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The value of the 'opc-next-page' response header from the previous call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, which can be ascending (ASC) or descending (DESC).
	SortOrder ListUsageLimitsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Supports one sort order.
	SortBy ListUsageLimitsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUsageLimitsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUsageLimitsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUsageLimitsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUsageLimitsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUsageLimitsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUsageLimitsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUsageLimitsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUsageLimitsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUsageLimitsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUsageLimitsResponse wrapper for the ListUsageLimits operation
type ListUsageLimitsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UsageLimitCollection instances
	UsageLimitCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListUsageLimitsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUsageLimitsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUsageLimitsSortOrderEnum Enum with underlying type: string
type ListUsageLimitsSortOrderEnum string

// Set of constants representing the allowable values for ListUsageLimitsSortOrderEnum
const (
	ListUsageLimitsSortOrderAsc  ListUsageLimitsSortOrderEnum = "ASC"
	ListUsageLimitsSortOrderDesc ListUsageLimitsSortOrderEnum = "DESC"
)

var mappingListUsageLimitsSortOrderEnum = map[string]ListUsageLimitsSortOrderEnum{
	"ASC":  ListUsageLimitsSortOrderAsc,
	"DESC": ListUsageLimitsSortOrderDesc,
}

var mappingListUsageLimitsSortOrderEnumLowerCase = map[string]ListUsageLimitsSortOrderEnum{
	"asc":  ListUsageLimitsSortOrderAsc,
	"desc": ListUsageLimitsSortOrderDesc,
}

// GetListUsageLimitsSortOrderEnumValues Enumerates the set of values for ListUsageLimitsSortOrderEnum
func GetListUsageLimitsSortOrderEnumValues() []ListUsageLimitsSortOrderEnum {
	values := make([]ListUsageLimitsSortOrderEnum, 0)
	for _, v := range mappingListUsageLimitsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUsageLimitsSortOrderEnumStringValues Enumerates the set of values in String for ListUsageLimitsSortOrderEnum
func GetListUsageLimitsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUsageLimitsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUsageLimitsSortOrderEnum(val string) (ListUsageLimitsSortOrderEnum, bool) {
	enum, ok := mappingListUsageLimitsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUsageLimitsSortByEnum Enum with underlying type: string
type ListUsageLimitsSortByEnum string

// Set of constants representing the allowable values for ListUsageLimitsSortByEnum
const (
	ListUsageLimitsSortByTimecreated ListUsageLimitsSortByEnum = "TIMECREATED"
	ListUsageLimitsSortByTimestart   ListUsageLimitsSortByEnum = "TIMESTART"
)

var mappingListUsageLimitsSortByEnum = map[string]ListUsageLimitsSortByEnum{
	"TIMECREATED": ListUsageLimitsSortByTimecreated,
	"TIMESTART":   ListUsageLimitsSortByTimestart,
}

var mappingListUsageLimitsSortByEnumLowerCase = map[string]ListUsageLimitsSortByEnum{
	"timecreated": ListUsageLimitsSortByTimecreated,
	"timestart":   ListUsageLimitsSortByTimestart,
}

// GetListUsageLimitsSortByEnumValues Enumerates the set of values for ListUsageLimitsSortByEnum
func GetListUsageLimitsSortByEnumValues() []ListUsageLimitsSortByEnum {
	values := make([]ListUsageLimitsSortByEnum, 0)
	for _, v := range mappingListUsageLimitsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUsageLimitsSortByEnumStringValues Enumerates the set of values in String for ListUsageLimitsSortByEnum
func GetListUsageLimitsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"TIMESTART",
	}
}

// GetMappingListUsageLimitsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUsageLimitsSortByEnum(val string) (ListUsageLimitsSortByEnum, bool) {
	enum, ok := mappingListUsageLimitsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
