// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package delegateaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDelegatedResourceAccessRequestHistoriesRequest wrapper for the ListDelegatedResourceAccessRequestHistories operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ListDelegatedResourceAccessRequestHistories.go.html to see an example of how to use ListDelegatedResourceAccessRequestHistoriesRequest.
type ListDelegatedResourceAccessRequestHistoriesRequest struct {

	// Unique Delegated Resource Access Request identifier
	DelegatedResourceAccessRequestId *string `mandatory:"true" contributesTo:"path" name:"delegatedResourceAccessRequestId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDelegatedResourceAccessRequestHistoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timestamp is descending. If no value is specified, timestamp is default.
	SortBy ListDelegatedResourceAccessRequestHistoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDelegatedResourceAccessRequestHistoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDelegatedResourceAccessRequestHistoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDelegatedResourceAccessRequestHistoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDelegatedResourceAccessRequestHistoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDelegatedResourceAccessRequestHistoriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDelegatedResourceAccessRequestHistoriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDelegatedResourceAccessRequestHistoriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDelegatedResourceAccessRequestHistoriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDelegatedResourceAccessRequestHistoriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDelegatedResourceAccessRequestHistoriesResponse wrapper for the ListDelegatedResourceAccessRequestHistories operation
type ListDelegatedResourceAccessRequestHistoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DelegatedResourceAccessRequestHistoryCollection instances
	DelegatedResourceAccessRequestHistoryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDelegatedResourceAccessRequestHistoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDelegatedResourceAccessRequestHistoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDelegatedResourceAccessRequestHistoriesSortOrderEnum Enum with underlying type: string
type ListDelegatedResourceAccessRequestHistoriesSortOrderEnum string

// Set of constants representing the allowable values for ListDelegatedResourceAccessRequestHistoriesSortOrderEnum
const (
	ListDelegatedResourceAccessRequestHistoriesSortOrderAsc  ListDelegatedResourceAccessRequestHistoriesSortOrderEnum = "ASC"
	ListDelegatedResourceAccessRequestHistoriesSortOrderDesc ListDelegatedResourceAccessRequestHistoriesSortOrderEnum = "DESC"
)

var mappingListDelegatedResourceAccessRequestHistoriesSortOrderEnum = map[string]ListDelegatedResourceAccessRequestHistoriesSortOrderEnum{
	"ASC":  ListDelegatedResourceAccessRequestHistoriesSortOrderAsc,
	"DESC": ListDelegatedResourceAccessRequestHistoriesSortOrderDesc,
}

var mappingListDelegatedResourceAccessRequestHistoriesSortOrderEnumLowerCase = map[string]ListDelegatedResourceAccessRequestHistoriesSortOrderEnum{
	"asc":  ListDelegatedResourceAccessRequestHistoriesSortOrderAsc,
	"desc": ListDelegatedResourceAccessRequestHistoriesSortOrderDesc,
}

// GetListDelegatedResourceAccessRequestHistoriesSortOrderEnumValues Enumerates the set of values for ListDelegatedResourceAccessRequestHistoriesSortOrderEnum
func GetListDelegatedResourceAccessRequestHistoriesSortOrderEnumValues() []ListDelegatedResourceAccessRequestHistoriesSortOrderEnum {
	values := make([]ListDelegatedResourceAccessRequestHistoriesSortOrderEnum, 0)
	for _, v := range mappingListDelegatedResourceAccessRequestHistoriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDelegatedResourceAccessRequestHistoriesSortOrderEnumStringValues Enumerates the set of values in String for ListDelegatedResourceAccessRequestHistoriesSortOrderEnum
func GetListDelegatedResourceAccessRequestHistoriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDelegatedResourceAccessRequestHistoriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDelegatedResourceAccessRequestHistoriesSortOrderEnum(val string) (ListDelegatedResourceAccessRequestHistoriesSortOrderEnum, bool) {
	enum, ok := mappingListDelegatedResourceAccessRequestHistoriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDelegatedResourceAccessRequestHistoriesSortByEnum Enum with underlying type: string
type ListDelegatedResourceAccessRequestHistoriesSortByEnum string

// Set of constants representing the allowable values for ListDelegatedResourceAccessRequestHistoriesSortByEnum
const (
	ListDelegatedResourceAccessRequestHistoriesSortByTimestamp ListDelegatedResourceAccessRequestHistoriesSortByEnum = "timestamp"
)

var mappingListDelegatedResourceAccessRequestHistoriesSortByEnum = map[string]ListDelegatedResourceAccessRequestHistoriesSortByEnum{
	"timestamp": ListDelegatedResourceAccessRequestHistoriesSortByTimestamp,
}

var mappingListDelegatedResourceAccessRequestHistoriesSortByEnumLowerCase = map[string]ListDelegatedResourceAccessRequestHistoriesSortByEnum{
	"timestamp": ListDelegatedResourceAccessRequestHistoriesSortByTimestamp,
}

// GetListDelegatedResourceAccessRequestHistoriesSortByEnumValues Enumerates the set of values for ListDelegatedResourceAccessRequestHistoriesSortByEnum
func GetListDelegatedResourceAccessRequestHistoriesSortByEnumValues() []ListDelegatedResourceAccessRequestHistoriesSortByEnum {
	values := make([]ListDelegatedResourceAccessRequestHistoriesSortByEnum, 0)
	for _, v := range mappingListDelegatedResourceAccessRequestHistoriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDelegatedResourceAccessRequestHistoriesSortByEnumStringValues Enumerates the set of values in String for ListDelegatedResourceAccessRequestHistoriesSortByEnum
func GetListDelegatedResourceAccessRequestHistoriesSortByEnumStringValues() []string {
	return []string{
		"timestamp",
	}
}

// GetMappingListDelegatedResourceAccessRequestHistoriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDelegatedResourceAccessRequestHistoriesSortByEnum(val string) (ListDelegatedResourceAccessRequestHistoriesSortByEnum, bool) {
	enum, ok := mappingListDelegatedResourceAccessRequestHistoriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
