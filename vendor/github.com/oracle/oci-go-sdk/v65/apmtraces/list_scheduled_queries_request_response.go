// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListScheduledQueriesRequest wrapper for the ListScheduledQueries operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/ListScheduledQueries.go.html to see an example of how to use ListScheduledQueriesRequest.
type ListScheduledQueriesRequest struct {

	// The APM Domain ID for the intended request.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page where to start retrieving results.
	// This is usually retrieved from a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return resources that match the given display name.  This will return resources that have name starting with this filter.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The displayName sort order
	// is case-sensitive.
	SortOrder ListScheduledQueriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one "sortBy" value.
	SortBy ListScheduledQueriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListScheduledQueriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListScheduledQueriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListScheduledQueriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListScheduledQueriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListScheduledQueriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListScheduledQueriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListScheduledQueriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListScheduledQueriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListScheduledQueriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListScheduledQueriesResponse wrapper for the ListScheduledQueries operation
type ListScheduledQueriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ScheduledQueryCollection instances
	ScheduledQueryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the page parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListScheduledQueriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListScheduledQueriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListScheduledQueriesSortOrderEnum Enum with underlying type: string
type ListScheduledQueriesSortOrderEnum string

// Set of constants representing the allowable values for ListScheduledQueriesSortOrderEnum
const (
	ListScheduledQueriesSortOrderAsc  ListScheduledQueriesSortOrderEnum = "ASC"
	ListScheduledQueriesSortOrderDesc ListScheduledQueriesSortOrderEnum = "DESC"
)

var mappingListScheduledQueriesSortOrderEnum = map[string]ListScheduledQueriesSortOrderEnum{
	"ASC":  ListScheduledQueriesSortOrderAsc,
	"DESC": ListScheduledQueriesSortOrderDesc,
}

var mappingListScheduledQueriesSortOrderEnumLowerCase = map[string]ListScheduledQueriesSortOrderEnum{
	"asc":  ListScheduledQueriesSortOrderAsc,
	"desc": ListScheduledQueriesSortOrderDesc,
}

// GetListScheduledQueriesSortOrderEnumValues Enumerates the set of values for ListScheduledQueriesSortOrderEnum
func GetListScheduledQueriesSortOrderEnumValues() []ListScheduledQueriesSortOrderEnum {
	values := make([]ListScheduledQueriesSortOrderEnum, 0)
	for _, v := range mappingListScheduledQueriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledQueriesSortOrderEnumStringValues Enumerates the set of values in String for ListScheduledQueriesSortOrderEnum
func GetListScheduledQueriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListScheduledQueriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledQueriesSortOrderEnum(val string) (ListScheduledQueriesSortOrderEnum, bool) {
	enum, ok := mappingListScheduledQueriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListScheduledQueriesSortByEnum Enum with underlying type: string
type ListScheduledQueriesSortByEnum string

// Set of constants representing the allowable values for ListScheduledQueriesSortByEnum
const (
	ListScheduledQueriesSortByName     ListScheduledQueriesSortByEnum = "SCHEDULED_QUERY_NAME"
	ListScheduledQueriesSortByType     ListScheduledQueriesSortByEnum = "SCHEDULED_QUERY_TYPE"
	ListScheduledQueriesSortBySubType  ListScheduledQueriesSortByEnum = "SCHEDULED_QUERY_SUB_TYPE"
	ListScheduledQueriesSortByNextRun  ListScheduledQueriesSortByEnum = "SCHEDULED_QUERY_NEXT_RUN"
	ListScheduledQueriesSortBySchedule ListScheduledQueriesSortByEnum = "SCHEDULED_QUERY_SCHEDULE"
)

var mappingListScheduledQueriesSortByEnum = map[string]ListScheduledQueriesSortByEnum{
	"SCHEDULED_QUERY_NAME":     ListScheduledQueriesSortByName,
	"SCHEDULED_QUERY_TYPE":     ListScheduledQueriesSortByType,
	"SCHEDULED_QUERY_SUB_TYPE": ListScheduledQueriesSortBySubType,
	"SCHEDULED_QUERY_NEXT_RUN": ListScheduledQueriesSortByNextRun,
	"SCHEDULED_QUERY_SCHEDULE": ListScheduledQueriesSortBySchedule,
}

var mappingListScheduledQueriesSortByEnumLowerCase = map[string]ListScheduledQueriesSortByEnum{
	"scheduled_query_name":     ListScheduledQueriesSortByName,
	"scheduled_query_type":     ListScheduledQueriesSortByType,
	"scheduled_query_sub_type": ListScheduledQueriesSortBySubType,
	"scheduled_query_next_run": ListScheduledQueriesSortByNextRun,
	"scheduled_query_schedule": ListScheduledQueriesSortBySchedule,
}

// GetListScheduledQueriesSortByEnumValues Enumerates the set of values for ListScheduledQueriesSortByEnum
func GetListScheduledQueriesSortByEnumValues() []ListScheduledQueriesSortByEnum {
	values := make([]ListScheduledQueriesSortByEnum, 0)
	for _, v := range mappingListScheduledQueriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledQueriesSortByEnumStringValues Enumerates the set of values in String for ListScheduledQueriesSortByEnum
func GetListScheduledQueriesSortByEnumStringValues() []string {
	return []string{
		"SCHEDULED_QUERY_NAME",
		"SCHEDULED_QUERY_TYPE",
		"SCHEDULED_QUERY_SUB_TYPE",
		"SCHEDULED_QUERY_NEXT_RUN",
		"SCHEDULED_QUERY_SCHEDULE",
	}
}

// GetMappingListScheduledQueriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledQueriesSortByEnum(val string) (ListScheduledQueriesSortByEnum, bool) {
	enum, ok := mappingListScheduledQueriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
