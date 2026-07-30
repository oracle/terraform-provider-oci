// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDataSafeTargetsRequest wrapper for the ListDataSafeTargets operation
type ListDataSafeTargetsRequest struct {

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListDataSafeTargetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListDataSafeTargetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataSafeTargetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataSafeTargetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataSafeTargetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataSafeTargetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataSafeTargetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataSafeTargetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataSafeTargetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSafeTargetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataSafeTargetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataSafeTargetsResponse wrapper for the ListDataSafeTargets operation
type ListDataSafeTargetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DataSafeTargetSummary instances
	Items []DataSafeTargetSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataSafeTargetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataSafeTargetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataSafeTargetsSortOrderEnum Enum with underlying type: string
type ListDataSafeTargetsSortOrderEnum string

// Set of constants representing the allowable values for ListDataSafeTargetsSortOrderEnum
const (
	ListDataSafeTargetsSortOrderAsc  ListDataSafeTargetsSortOrderEnum = "ASC"
	ListDataSafeTargetsSortOrderDesc ListDataSafeTargetsSortOrderEnum = "DESC"
)

var mappingListDataSafeTargetsSortOrderEnum = map[string]ListDataSafeTargetsSortOrderEnum{
	"ASC":  ListDataSafeTargetsSortOrderAsc,
	"DESC": ListDataSafeTargetsSortOrderDesc,
}

var mappingListDataSafeTargetsSortOrderEnumLowerCase = map[string]ListDataSafeTargetsSortOrderEnum{
	"asc":  ListDataSafeTargetsSortOrderAsc,
	"desc": ListDataSafeTargetsSortOrderDesc,
}

// GetListDataSafeTargetsSortOrderEnumValues Enumerates the set of values for ListDataSafeTargetsSortOrderEnum
func GetListDataSafeTargetsSortOrderEnumValues() []ListDataSafeTargetsSortOrderEnum {
	values := make([]ListDataSafeTargetsSortOrderEnum, 0)
	for _, v := range mappingListDataSafeTargetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSafeTargetsSortOrderEnumStringValues Enumerates the set of values in String for ListDataSafeTargetsSortOrderEnum
func GetListDataSafeTargetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataSafeTargetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSafeTargetsSortOrderEnum(val string) (ListDataSafeTargetsSortOrderEnum, bool) {
	enum, ok := mappingListDataSafeTargetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSafeTargetsSortByEnum Enum with underlying type: string
type ListDataSafeTargetsSortByEnum string

// Set of constants representing the allowable values for ListDataSafeTargetsSortByEnum
const (
	ListDataSafeTargetsSortByTimecreated ListDataSafeTargetsSortByEnum = "TIMECREATED"
	ListDataSafeTargetsSortByDisplayname ListDataSafeTargetsSortByEnum = "DISPLAYNAME"
)

var mappingListDataSafeTargetsSortByEnum = map[string]ListDataSafeTargetsSortByEnum{
	"TIMECREATED": ListDataSafeTargetsSortByTimecreated,
	"DISPLAYNAME": ListDataSafeTargetsSortByDisplayname,
}

var mappingListDataSafeTargetsSortByEnumLowerCase = map[string]ListDataSafeTargetsSortByEnum{
	"timecreated": ListDataSafeTargetsSortByTimecreated,
	"displayname": ListDataSafeTargetsSortByDisplayname,
}

// GetListDataSafeTargetsSortByEnumValues Enumerates the set of values for ListDataSafeTargetsSortByEnum
func GetListDataSafeTargetsSortByEnumValues() []ListDataSafeTargetsSortByEnum {
	values := make([]ListDataSafeTargetsSortByEnum, 0)
	for _, v := range mappingListDataSafeTargetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSafeTargetsSortByEnumStringValues Enumerates the set of values in String for ListDataSafeTargetsSortByEnum
func GetListDataSafeTargetsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListDataSafeTargetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSafeTargetsSortByEnum(val string) (ListDataSafeTargetsSortByEnum, bool) {
	enum, ok := mappingListDataSafeTargetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
