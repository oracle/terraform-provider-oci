// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"net/http"
	"strings"
)

// ListReferenceArtifactsRequest wrapper for the ListReferenceArtifacts operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataconnectivity/ListReferenceArtifacts.go.html to see an example of how to use ListReferenceArtifactsRequest.
type ListReferenceArtifactsRequest struct {

	// The registry Ocid.
	RegistryId *string `mandatory:"true" contributesTo:"path" name:"registryId"`

	// The ID of a dcms artifact (DataAsset or Endpoint).
	DcmsArtifactId *string `mandatory:"true" contributesTo:"path" name:"dcmsArtifactId"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Type of the object to filter the results with.
	Type *string `mandatory:"false" contributesTo:"query" name:"type"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListReferenceArtifactsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListReferenceArtifactsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Types which wont be listed while listing dataAsset/Connection
	ExcludeTypes []string `contributesTo:"query" name:"excludeTypes" collectionFormat:"multi"`

	// If value is FAVORITES_ONLY, then only objects marked as favorite by the requesting user will be included in result. If value is NON_FAVORITES_ONLY, then objects marked as favorites by the requesting user will be skipped. If value is ALL or if not specified, all objects, irrespective of favorites or not will be returned. Default is ALL.
	FavoritesQueryParam ListReferenceArtifactsFavoritesQueryParamEnum `mandatory:"false" contributesTo:"query" name:"favoritesQueryParam" omitEmpty:"true"`

	// Unique key of the service.
	ServiceArtifactId *string `mandatory:"false" contributesTo:"query" name:"serviceArtifactId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReferenceArtifactsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReferenceArtifactsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListReferenceArtifactsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReferenceArtifactsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListReferenceArtifactsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListReferenceArtifactsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListReferenceArtifactsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReferenceArtifactsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListReferenceArtifactsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReferenceArtifactsFavoritesQueryParamEnum(string(request.FavoritesQueryParam)); !ok && request.FavoritesQueryParam != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FavoritesQueryParam: %s. Supported values are: %s.", request.FavoritesQueryParam, strings.Join(GetListReferenceArtifactsFavoritesQueryParamEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListReferenceArtifactsResponse wrapper for the ListReferenceArtifacts operation
type ListReferenceArtifactsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ReferenceArtifactSummaryCollection instances
	ReferenceArtifactSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListReferenceArtifactsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReferenceArtifactsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListReferenceArtifactsSortByEnum Enum with underlying type: string
type ListReferenceArtifactsSortByEnum string

// Set of constants representing the allowable values for ListReferenceArtifactsSortByEnum
const (
	ListReferenceArtifactsSortById          ListReferenceArtifactsSortByEnum = "id"
	ListReferenceArtifactsSortByTimecreated ListReferenceArtifactsSortByEnum = "timeCreated"
	ListReferenceArtifactsSortByDisplayname ListReferenceArtifactsSortByEnum = "displayName"
)

var mappingListReferenceArtifactsSortByEnum = map[string]ListReferenceArtifactsSortByEnum{
	"id":          ListReferenceArtifactsSortById,
	"timeCreated": ListReferenceArtifactsSortByTimecreated,
	"displayName": ListReferenceArtifactsSortByDisplayname,
}

var mappingListReferenceArtifactsSortByEnumLowerCase = map[string]ListReferenceArtifactsSortByEnum{
	"id":          ListReferenceArtifactsSortById,
	"timecreated": ListReferenceArtifactsSortByTimecreated,
	"displayname": ListReferenceArtifactsSortByDisplayname,
}

// GetListReferenceArtifactsSortByEnumValues Enumerates the set of values for ListReferenceArtifactsSortByEnum
func GetListReferenceArtifactsSortByEnumValues() []ListReferenceArtifactsSortByEnum {
	values := make([]ListReferenceArtifactsSortByEnum, 0)
	for _, v := range mappingListReferenceArtifactsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListReferenceArtifactsSortByEnumStringValues Enumerates the set of values in String for ListReferenceArtifactsSortByEnum
func GetListReferenceArtifactsSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeCreated",
		"displayName",
	}
}

// GetMappingListReferenceArtifactsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReferenceArtifactsSortByEnum(val string) (ListReferenceArtifactsSortByEnum, bool) {
	enum, ok := mappingListReferenceArtifactsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReferenceArtifactsSortOrderEnum Enum with underlying type: string
type ListReferenceArtifactsSortOrderEnum string

// Set of constants representing the allowable values for ListReferenceArtifactsSortOrderEnum
const (
	ListReferenceArtifactsSortOrderAsc  ListReferenceArtifactsSortOrderEnum = "ASC"
	ListReferenceArtifactsSortOrderDesc ListReferenceArtifactsSortOrderEnum = "DESC"
)

var mappingListReferenceArtifactsSortOrderEnum = map[string]ListReferenceArtifactsSortOrderEnum{
	"ASC":  ListReferenceArtifactsSortOrderAsc,
	"DESC": ListReferenceArtifactsSortOrderDesc,
}

var mappingListReferenceArtifactsSortOrderEnumLowerCase = map[string]ListReferenceArtifactsSortOrderEnum{
	"asc":  ListReferenceArtifactsSortOrderAsc,
	"desc": ListReferenceArtifactsSortOrderDesc,
}

// GetListReferenceArtifactsSortOrderEnumValues Enumerates the set of values for ListReferenceArtifactsSortOrderEnum
func GetListReferenceArtifactsSortOrderEnumValues() []ListReferenceArtifactsSortOrderEnum {
	values := make([]ListReferenceArtifactsSortOrderEnum, 0)
	for _, v := range mappingListReferenceArtifactsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListReferenceArtifactsSortOrderEnumStringValues Enumerates the set of values in String for ListReferenceArtifactsSortOrderEnum
func GetListReferenceArtifactsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListReferenceArtifactsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReferenceArtifactsSortOrderEnum(val string) (ListReferenceArtifactsSortOrderEnum, bool) {
	enum, ok := mappingListReferenceArtifactsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReferenceArtifactsFavoritesQueryParamEnum Enum with underlying type: string
type ListReferenceArtifactsFavoritesQueryParamEnum string

// Set of constants representing the allowable values for ListReferenceArtifactsFavoritesQueryParamEnum
const (
	ListReferenceArtifactsFavoritesQueryParamFavoritesOnly    ListReferenceArtifactsFavoritesQueryParamEnum = "FAVORITES_ONLY"
	ListReferenceArtifactsFavoritesQueryParamNonFavoritesOnly ListReferenceArtifactsFavoritesQueryParamEnum = "NON_FAVORITES_ONLY"
	ListReferenceArtifactsFavoritesQueryParamAll              ListReferenceArtifactsFavoritesQueryParamEnum = "ALL"
)

var mappingListReferenceArtifactsFavoritesQueryParamEnum = map[string]ListReferenceArtifactsFavoritesQueryParamEnum{
	"FAVORITES_ONLY":     ListReferenceArtifactsFavoritesQueryParamFavoritesOnly,
	"NON_FAVORITES_ONLY": ListReferenceArtifactsFavoritesQueryParamNonFavoritesOnly,
	"ALL":                ListReferenceArtifactsFavoritesQueryParamAll,
}

var mappingListReferenceArtifactsFavoritesQueryParamEnumLowerCase = map[string]ListReferenceArtifactsFavoritesQueryParamEnum{
	"favorites_only":     ListReferenceArtifactsFavoritesQueryParamFavoritesOnly,
	"non_favorites_only": ListReferenceArtifactsFavoritesQueryParamNonFavoritesOnly,
	"all":                ListReferenceArtifactsFavoritesQueryParamAll,
}

// GetListReferenceArtifactsFavoritesQueryParamEnumValues Enumerates the set of values for ListReferenceArtifactsFavoritesQueryParamEnum
func GetListReferenceArtifactsFavoritesQueryParamEnumValues() []ListReferenceArtifactsFavoritesQueryParamEnum {
	values := make([]ListReferenceArtifactsFavoritesQueryParamEnum, 0)
	for _, v := range mappingListReferenceArtifactsFavoritesQueryParamEnum {
		values = append(values, v)
	}
	return values
}

// GetListReferenceArtifactsFavoritesQueryParamEnumStringValues Enumerates the set of values in String for ListReferenceArtifactsFavoritesQueryParamEnum
func GetListReferenceArtifactsFavoritesQueryParamEnumStringValues() []string {
	return []string{
		"FAVORITES_ONLY",
		"NON_FAVORITES_ONLY",
		"ALL",
	}
}

// GetMappingListReferenceArtifactsFavoritesQueryParamEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReferenceArtifactsFavoritesQueryParamEnum(val string) (ListReferenceArtifactsFavoritesQueryParamEnum, bool) {
	enum, ok := mappingListReferenceArtifactsFavoritesQueryParamEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
