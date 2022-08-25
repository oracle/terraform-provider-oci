// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ons

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOptOutNumbersRequest wrapper for the ListOptOutNumbers operation
type ListOptOutNumbersRequest struct {

	// unique PhoneApplication identifier
	PhoneApplicationId *string `mandatory:"true" contributesTo:"path" name:"phoneApplicationId"`

	// A filter to return only resources that match the entire toNumber given
	ToNumber *string `mandatory:"false" contributesTo:"query" name:"toNumber"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use (ascending or descending).
	SortOrder ListOptOutNumbersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one field can be selected for sorting.
	SortBy ListOptOutNumbersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOptOutNumbersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOptOutNumbersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOptOutNumbersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOptOutNumbersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOptOutNumbersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOptOutNumbersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOptOutNumbersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOptOutNumbersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOptOutNumbersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOptOutNumbersResponse wrapper for the ListOptOutNumbers operation
type ListOptOutNumbersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OptOutNumberCollection instances
	OptOutNumberCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOptOutNumbersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOptOutNumbersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOptOutNumbersSortOrderEnum Enum with underlying type: string
type ListOptOutNumbersSortOrderEnum string

// Set of constants representing the allowable values for ListOptOutNumbersSortOrderEnum
const (
	ListOptOutNumbersSortOrderAsc  ListOptOutNumbersSortOrderEnum = "ASC"
	ListOptOutNumbersSortOrderDesc ListOptOutNumbersSortOrderEnum = "DESC"
)

var mappingListOptOutNumbersSortOrderEnum = map[string]ListOptOutNumbersSortOrderEnum{
	"ASC":  ListOptOutNumbersSortOrderAsc,
	"DESC": ListOptOutNumbersSortOrderDesc,
}

var mappingListOptOutNumbersSortOrderEnumLowerCase = map[string]ListOptOutNumbersSortOrderEnum{
	"asc":  ListOptOutNumbersSortOrderAsc,
	"desc": ListOptOutNumbersSortOrderDesc,
}

// GetListOptOutNumbersSortOrderEnumValues Enumerates the set of values for ListOptOutNumbersSortOrderEnum
func GetListOptOutNumbersSortOrderEnumValues() []ListOptOutNumbersSortOrderEnum {
	values := make([]ListOptOutNumbersSortOrderEnum, 0)
	for _, v := range mappingListOptOutNumbersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOptOutNumbersSortOrderEnumStringValues Enumerates the set of values in String for ListOptOutNumbersSortOrderEnum
func GetListOptOutNumbersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOptOutNumbersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOptOutNumbersSortOrderEnum(val string) (ListOptOutNumbersSortOrderEnum, bool) {
	enum, ok := mappingListOptOutNumbersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOptOutNumbersSortByEnum Enum with underlying type: string
type ListOptOutNumbersSortByEnum string

// Set of constants representing the allowable values for ListOptOutNumbersSortByEnum
const (
	ListOptOutNumbersSortByTimecreated    ListOptOutNumbersSortByEnum = "TIMECREATED"
	ListOptOutNumbersSortByLifecyclestate ListOptOutNumbersSortByEnum = "LIFECYCLESTATE"
)

var mappingListOptOutNumbersSortByEnum = map[string]ListOptOutNumbersSortByEnum{
	"TIMECREATED":    ListOptOutNumbersSortByTimecreated,
	"LIFECYCLESTATE": ListOptOutNumbersSortByLifecyclestate,
}

var mappingListOptOutNumbersSortByEnumLowerCase = map[string]ListOptOutNumbersSortByEnum{
	"timecreated":    ListOptOutNumbersSortByTimecreated,
	"lifecyclestate": ListOptOutNumbersSortByLifecyclestate,
}

// GetListOptOutNumbersSortByEnumValues Enumerates the set of values for ListOptOutNumbersSortByEnum
func GetListOptOutNumbersSortByEnumValues() []ListOptOutNumbersSortByEnum {
	values := make([]ListOptOutNumbersSortByEnum, 0)
	for _, v := range mappingListOptOutNumbersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOptOutNumbersSortByEnumStringValues Enumerates the set of values in String for ListOptOutNumbersSortByEnum
func GetListOptOutNumbersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"LIFECYCLESTATE",
	}
}

// GetMappingListOptOutNumbersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOptOutNumbersSortByEnum(val string) (ListOptOutNumbersSortByEnum, bool) {
	enum, ok := mappingListOptOutNumbersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
