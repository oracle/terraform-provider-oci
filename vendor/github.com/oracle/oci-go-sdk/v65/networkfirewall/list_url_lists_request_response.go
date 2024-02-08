// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package networkfirewall

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListUrlListsRequest wrapper for the ListUrlLists operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListUrlLists.go.html to see an example of how to use ListUrlListsRequest.
type ListUrlListsRequest struct {

	// Unique Network Firewall Policy identifier
	NetworkFirewallPolicyId *string `mandatory:"true" contributesTo:"path" name:"networkFirewallPolicyId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` or `opc-prev-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListUrlListsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListUrlListsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUrlListsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUrlListsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUrlListsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUrlListsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUrlListsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUrlListsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUrlListsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUrlListsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUrlListsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUrlListsResponse wrapper for the ListUrlLists operation
type ListUrlListsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UrlListSummaryCollection instances
	UrlListSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results exist. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. This is to get the page counts overall.
	OpcPageCount *string `presentIn:"header" name:"opc-page-count"`

	// For pagination of a list of items. This provides the count of total items across pages.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListUrlListsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUrlListsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUrlListsSortOrderEnum Enum with underlying type: string
type ListUrlListsSortOrderEnum string

// Set of constants representing the allowable values for ListUrlListsSortOrderEnum
const (
	ListUrlListsSortOrderAsc  ListUrlListsSortOrderEnum = "ASC"
	ListUrlListsSortOrderDesc ListUrlListsSortOrderEnum = "DESC"
)

var mappingListUrlListsSortOrderEnum = map[string]ListUrlListsSortOrderEnum{
	"ASC":  ListUrlListsSortOrderAsc,
	"DESC": ListUrlListsSortOrderDesc,
}

var mappingListUrlListsSortOrderEnumLowerCase = map[string]ListUrlListsSortOrderEnum{
	"asc":  ListUrlListsSortOrderAsc,
	"desc": ListUrlListsSortOrderDesc,
}

// GetListUrlListsSortOrderEnumValues Enumerates the set of values for ListUrlListsSortOrderEnum
func GetListUrlListsSortOrderEnumValues() []ListUrlListsSortOrderEnum {
	values := make([]ListUrlListsSortOrderEnum, 0)
	for _, v := range mappingListUrlListsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUrlListsSortOrderEnumStringValues Enumerates the set of values in String for ListUrlListsSortOrderEnum
func GetListUrlListsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUrlListsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUrlListsSortOrderEnum(val string) (ListUrlListsSortOrderEnum, bool) {
	enum, ok := mappingListUrlListsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUrlListsSortByEnum Enum with underlying type: string
type ListUrlListsSortByEnum string

// Set of constants representing the allowable values for ListUrlListsSortByEnum
const (
	ListUrlListsSortByTimecreated ListUrlListsSortByEnum = "timeCreated"
	ListUrlListsSortByDisplayname ListUrlListsSortByEnum = "displayName"
)

var mappingListUrlListsSortByEnum = map[string]ListUrlListsSortByEnum{
	"timeCreated": ListUrlListsSortByTimecreated,
	"displayName": ListUrlListsSortByDisplayname,
}

var mappingListUrlListsSortByEnumLowerCase = map[string]ListUrlListsSortByEnum{
	"timecreated": ListUrlListsSortByTimecreated,
	"displayname": ListUrlListsSortByDisplayname,
}

// GetListUrlListsSortByEnumValues Enumerates the set of values for ListUrlListsSortByEnum
func GetListUrlListsSortByEnumValues() []ListUrlListsSortByEnum {
	values := make([]ListUrlListsSortByEnum, 0)
	for _, v := range mappingListUrlListsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUrlListsSortByEnumStringValues Enumerates the set of values in String for ListUrlListsSortByEnum
func GetListUrlListsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListUrlListsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUrlListsSortByEnum(val string) (ListUrlListsSortByEnum, bool) {
	enum, ok := mappingListUrlListsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
