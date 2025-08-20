// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListKnowledgeGroupsRequest wrapper for the ListKnowledgeGroups operation
type ListKnowledgeGroupsRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// List only Knowledge Group resources with this name. Names are unique and may not change.
	// Example: `MyKnowledgeGroup`
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListKnowledgeGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `timeCreated`.
	// The default sort order for `timeCreated` and `timeUpdated` is descending.
	// For all other sort fields the default sort order is ascending.
	SortBy ListKnowledgeGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListKnowledgeGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListKnowledgeGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListKnowledgeGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListKnowledgeGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListKnowledgeGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListKnowledgeGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListKnowledgeGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListKnowledgeGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListKnowledgeGroupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListKnowledgeGroupsResponse wrapper for the ListKnowledgeGroups operation
type ListKnowledgeGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of KnowledgeGroupCollection instances
	KnowledgeGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// When you are paging through a list, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the
	// `page` query parameter for the subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of results that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListKnowledgeGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListKnowledgeGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListKnowledgeGroupsSortOrderEnum Enum with underlying type: string
type ListKnowledgeGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListKnowledgeGroupsSortOrderEnum
const (
	ListKnowledgeGroupsSortOrderAsc  ListKnowledgeGroupsSortOrderEnum = "ASC"
	ListKnowledgeGroupsSortOrderDesc ListKnowledgeGroupsSortOrderEnum = "DESC"
)

var mappingListKnowledgeGroupsSortOrderEnum = map[string]ListKnowledgeGroupsSortOrderEnum{
	"ASC":  ListKnowledgeGroupsSortOrderAsc,
	"DESC": ListKnowledgeGroupsSortOrderDesc,
}

var mappingListKnowledgeGroupsSortOrderEnumLowerCase = map[string]ListKnowledgeGroupsSortOrderEnum{
	"asc":  ListKnowledgeGroupsSortOrderAsc,
	"desc": ListKnowledgeGroupsSortOrderDesc,
}

// GetListKnowledgeGroupsSortOrderEnumValues Enumerates the set of values for ListKnowledgeGroupsSortOrderEnum
func GetListKnowledgeGroupsSortOrderEnumValues() []ListKnowledgeGroupsSortOrderEnum {
	values := make([]ListKnowledgeGroupsSortOrderEnum, 0)
	for _, v := range mappingListKnowledgeGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListKnowledgeGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListKnowledgeGroupsSortOrderEnum
func GetListKnowledgeGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListKnowledgeGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListKnowledgeGroupsSortOrderEnum(val string) (ListKnowledgeGroupsSortOrderEnum, bool) {
	enum, ok := mappingListKnowledgeGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListKnowledgeGroupsSortByEnum Enum with underlying type: string
type ListKnowledgeGroupsSortByEnum string

// Set of constants representing the allowable values for ListKnowledgeGroupsSortByEnum
const (
	ListKnowledgeGroupsSortByTimecreated ListKnowledgeGroupsSortByEnum = "timeCreated"
	ListKnowledgeGroupsSortByTimeupdated ListKnowledgeGroupsSortByEnum = "timeUpdated"
	ListKnowledgeGroupsSortByName        ListKnowledgeGroupsSortByEnum = "name"
)

var mappingListKnowledgeGroupsSortByEnum = map[string]ListKnowledgeGroupsSortByEnum{
	"timeCreated": ListKnowledgeGroupsSortByTimecreated,
	"timeUpdated": ListKnowledgeGroupsSortByTimeupdated,
	"name":        ListKnowledgeGroupsSortByName,
}

var mappingListKnowledgeGroupsSortByEnumLowerCase = map[string]ListKnowledgeGroupsSortByEnum{
	"timecreated": ListKnowledgeGroupsSortByTimecreated,
	"timeupdated": ListKnowledgeGroupsSortByTimeupdated,
	"name":        ListKnowledgeGroupsSortByName,
}

// GetListKnowledgeGroupsSortByEnumValues Enumerates the set of values for ListKnowledgeGroupsSortByEnum
func GetListKnowledgeGroupsSortByEnumValues() []ListKnowledgeGroupsSortByEnum {
	values := make([]ListKnowledgeGroupsSortByEnum, 0)
	for _, v := range mappingListKnowledgeGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListKnowledgeGroupsSortByEnumStringValues Enumerates the set of values in String for ListKnowledgeGroupsSortByEnum
func GetListKnowledgeGroupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"name",
	}
}

// GetMappingListKnowledgeGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListKnowledgeGroupsSortByEnum(val string) (ListKnowledgeGroupsSortByEnum, bool) {
	enum, ok := mappingListKnowledgeGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
