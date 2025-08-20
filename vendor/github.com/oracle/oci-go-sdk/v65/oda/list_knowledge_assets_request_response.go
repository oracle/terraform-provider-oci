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

// ListKnowledgeAssetsRequest wrapper for the ListKnowledgeAssets operation
type ListKnowledgeAssetsRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// Unique Knowledge Group identifier.
	KnowledgeGroupId *string `mandatory:"true" contributesTo:"path" name:"knowledgeGroupId"`

	// List only Knowledge Asset resources with this name. Names are unique and may not change.
	// Example: `MyKnowledgeAsset`
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListKnowledgeAssetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `timeCreated`.
	// The default sort order for `timeCreated` and `timeUpdated` is descending.
	// For all other sort fields the default sort order is ascending.
	SortBy ListKnowledgeAssetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListKnowledgeAssetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListKnowledgeAssetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListKnowledgeAssetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListKnowledgeAssetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListKnowledgeAssetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListKnowledgeAssetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListKnowledgeAssetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListKnowledgeAssetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListKnowledgeAssetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListKnowledgeAssetsResponse wrapper for the ListKnowledgeAssets operation
type ListKnowledgeAssetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of KnowledgeAssetCollection instances
	KnowledgeAssetCollection `presentIn:"body"`

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

func (response ListKnowledgeAssetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListKnowledgeAssetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListKnowledgeAssetsSortOrderEnum Enum with underlying type: string
type ListKnowledgeAssetsSortOrderEnum string

// Set of constants representing the allowable values for ListKnowledgeAssetsSortOrderEnum
const (
	ListKnowledgeAssetsSortOrderAsc  ListKnowledgeAssetsSortOrderEnum = "ASC"
	ListKnowledgeAssetsSortOrderDesc ListKnowledgeAssetsSortOrderEnum = "DESC"
)

var mappingListKnowledgeAssetsSortOrderEnum = map[string]ListKnowledgeAssetsSortOrderEnum{
	"ASC":  ListKnowledgeAssetsSortOrderAsc,
	"DESC": ListKnowledgeAssetsSortOrderDesc,
}

var mappingListKnowledgeAssetsSortOrderEnumLowerCase = map[string]ListKnowledgeAssetsSortOrderEnum{
	"asc":  ListKnowledgeAssetsSortOrderAsc,
	"desc": ListKnowledgeAssetsSortOrderDesc,
}

// GetListKnowledgeAssetsSortOrderEnumValues Enumerates the set of values for ListKnowledgeAssetsSortOrderEnum
func GetListKnowledgeAssetsSortOrderEnumValues() []ListKnowledgeAssetsSortOrderEnum {
	values := make([]ListKnowledgeAssetsSortOrderEnum, 0)
	for _, v := range mappingListKnowledgeAssetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListKnowledgeAssetsSortOrderEnumStringValues Enumerates the set of values in String for ListKnowledgeAssetsSortOrderEnum
func GetListKnowledgeAssetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListKnowledgeAssetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListKnowledgeAssetsSortOrderEnum(val string) (ListKnowledgeAssetsSortOrderEnum, bool) {
	enum, ok := mappingListKnowledgeAssetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListKnowledgeAssetsSortByEnum Enum with underlying type: string
type ListKnowledgeAssetsSortByEnum string

// Set of constants representing the allowable values for ListKnowledgeAssetsSortByEnum
const (
	ListKnowledgeAssetsSortByTimecreated ListKnowledgeAssetsSortByEnum = "timeCreated"
	ListKnowledgeAssetsSortByTimeupdated ListKnowledgeAssetsSortByEnum = "timeUpdated"
	ListKnowledgeAssetsSortByName        ListKnowledgeAssetsSortByEnum = "name"
)

var mappingListKnowledgeAssetsSortByEnum = map[string]ListKnowledgeAssetsSortByEnum{
	"timeCreated": ListKnowledgeAssetsSortByTimecreated,
	"timeUpdated": ListKnowledgeAssetsSortByTimeupdated,
	"name":        ListKnowledgeAssetsSortByName,
}

var mappingListKnowledgeAssetsSortByEnumLowerCase = map[string]ListKnowledgeAssetsSortByEnum{
	"timecreated": ListKnowledgeAssetsSortByTimecreated,
	"timeupdated": ListKnowledgeAssetsSortByTimeupdated,
	"name":        ListKnowledgeAssetsSortByName,
}

// GetListKnowledgeAssetsSortByEnumValues Enumerates the set of values for ListKnowledgeAssetsSortByEnum
func GetListKnowledgeAssetsSortByEnumValues() []ListKnowledgeAssetsSortByEnum {
	values := make([]ListKnowledgeAssetsSortByEnum, 0)
	for _, v := range mappingListKnowledgeAssetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListKnowledgeAssetsSortByEnumStringValues Enumerates the set of values in String for ListKnowledgeAssetsSortByEnum
func GetListKnowledgeAssetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"name",
	}
}

// GetMappingListKnowledgeAssetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListKnowledgeAssetsSortByEnum(val string) (ListKnowledgeAssetsSortByEnum, bool) {
	enum, ok := mappingListKnowledgeAssetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
