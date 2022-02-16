// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListResponderActivitiesRequest wrapper for the ListResponderActivities operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResponderActivities.go.html to see an example of how to use ListResponderActivitiesRequest.
type ListResponderActivitiesRequest struct {

	// OCId of the problem.
	ProblemId *string `mandatory:"true" contributesTo:"path" name:"problemId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListResponderActivitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for responderRuleName is ascending. If no value is specified timeCreated is default.
	SortBy ListResponderActivitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResponderActivitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResponderActivitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResponderActivitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResponderActivitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResponderActivitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResponderActivitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResponderActivitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResponderActivitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResponderActivitiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResponderActivitiesResponse wrapper for the ListResponderActivities operation
type ListResponderActivitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResponderActivityCollection instances
	ResponderActivityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResponderActivitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResponderActivitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResponderActivitiesSortOrderEnum Enum with underlying type: string
type ListResponderActivitiesSortOrderEnum string

// Set of constants representing the allowable values for ListResponderActivitiesSortOrderEnum
const (
	ListResponderActivitiesSortOrderAsc  ListResponderActivitiesSortOrderEnum = "ASC"
	ListResponderActivitiesSortOrderDesc ListResponderActivitiesSortOrderEnum = "DESC"
)

var mappingListResponderActivitiesSortOrderEnum = map[string]ListResponderActivitiesSortOrderEnum{
	"ASC":  ListResponderActivitiesSortOrderAsc,
	"DESC": ListResponderActivitiesSortOrderDesc,
}

// GetListResponderActivitiesSortOrderEnumValues Enumerates the set of values for ListResponderActivitiesSortOrderEnum
func GetListResponderActivitiesSortOrderEnumValues() []ListResponderActivitiesSortOrderEnum {
	values := make([]ListResponderActivitiesSortOrderEnum, 0)
	for _, v := range mappingListResponderActivitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderActivitiesSortOrderEnumStringValues Enumerates the set of values in String for ListResponderActivitiesSortOrderEnum
func GetListResponderActivitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResponderActivitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderActivitiesSortOrderEnum(val string) (ListResponderActivitiesSortOrderEnum, bool) {
	mappingListResponderActivitiesSortOrderEnumIgnoreCase := make(map[string]ListResponderActivitiesSortOrderEnum)
	for k, v := range mappingListResponderActivitiesSortOrderEnum {
		mappingListResponderActivitiesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListResponderActivitiesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListResponderActivitiesSortByEnum Enum with underlying type: string
type ListResponderActivitiesSortByEnum string

// Set of constants representing the allowable values for ListResponderActivitiesSortByEnum
const (
	ListResponderActivitiesSortByTimecreated       ListResponderActivitiesSortByEnum = "timeCreated"
	ListResponderActivitiesSortByResponderrulename ListResponderActivitiesSortByEnum = "responderRuleName"
)

var mappingListResponderActivitiesSortByEnum = map[string]ListResponderActivitiesSortByEnum{
	"timeCreated":       ListResponderActivitiesSortByTimecreated,
	"responderRuleName": ListResponderActivitiesSortByResponderrulename,
}

// GetListResponderActivitiesSortByEnumValues Enumerates the set of values for ListResponderActivitiesSortByEnum
func GetListResponderActivitiesSortByEnumValues() []ListResponderActivitiesSortByEnum {
	values := make([]ListResponderActivitiesSortByEnum, 0)
	for _, v := range mappingListResponderActivitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderActivitiesSortByEnumStringValues Enumerates the set of values in String for ListResponderActivitiesSortByEnum
func GetListResponderActivitiesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"responderRuleName",
	}
}

// GetMappingListResponderActivitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderActivitiesSortByEnum(val string) (ListResponderActivitiesSortByEnum, bool) {
	mappingListResponderActivitiesSortByEnumIgnoreCase := make(map[string]ListResponderActivitiesSortByEnum)
	for k, v := range mappingListResponderActivitiesSortByEnum {
		mappingListResponderActivitiesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListResponderActivitiesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
