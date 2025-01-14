// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRecipesRequest wrapper for the ListRecipes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListRecipes.go.html to see an example of how to use ListRecipesRequest.
type ListRecipesRequest struct {

	// The OCID of the compartment that contains the work request. Work requests should be scoped
	// to the same compartment as the resource the work request affects. If the work request concerns
	// multiple resources, and those resources are not in the same compartment, it is up to the service team
	// to pick the primary resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The pipeline's recipe type. The default value is ZERO_ETL.
	RecipeType ListRecipesRecipeTypeEnum `mandatory:"false" contributesTo:"query" name:"recipeType" omitEmpty:"true"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListRecipesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
	SortBy ListRecipesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRecipesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRecipesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRecipesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRecipesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRecipesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRecipesRecipeTypeEnum(string(request.RecipeType)); !ok && request.RecipeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecipeType: %s. Supported values are: %s.", request.RecipeType, strings.Join(GetListRecipesRecipeTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecipesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRecipesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecipesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRecipesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRecipesResponse wrapper for the ListRecipes operation
type ListRecipesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecipeSummaryCollection instances
	RecipeSummaryCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRecipesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRecipesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRecipesRecipeTypeEnum Enum with underlying type: string
type ListRecipesRecipeTypeEnum string

// Set of constants representing the allowable values for ListRecipesRecipeTypeEnum
const (
	ListRecipesRecipeTypeZeroEtl ListRecipesRecipeTypeEnum = "ZERO_ETL"
)

var mappingListRecipesRecipeTypeEnum = map[string]ListRecipesRecipeTypeEnum{
	"ZERO_ETL": ListRecipesRecipeTypeZeroEtl,
}

var mappingListRecipesRecipeTypeEnumLowerCase = map[string]ListRecipesRecipeTypeEnum{
	"zero_etl": ListRecipesRecipeTypeZeroEtl,
}

// GetListRecipesRecipeTypeEnumValues Enumerates the set of values for ListRecipesRecipeTypeEnum
func GetListRecipesRecipeTypeEnumValues() []ListRecipesRecipeTypeEnum {
	values := make([]ListRecipesRecipeTypeEnum, 0)
	for _, v := range mappingListRecipesRecipeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecipesRecipeTypeEnumStringValues Enumerates the set of values in String for ListRecipesRecipeTypeEnum
func GetListRecipesRecipeTypeEnumStringValues() []string {
	return []string{
		"ZERO_ETL",
	}
}

// GetMappingListRecipesRecipeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecipesRecipeTypeEnum(val string) (ListRecipesRecipeTypeEnum, bool) {
	enum, ok := mappingListRecipesRecipeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecipesSortOrderEnum Enum with underlying type: string
type ListRecipesSortOrderEnum string

// Set of constants representing the allowable values for ListRecipesSortOrderEnum
const (
	ListRecipesSortOrderAsc  ListRecipesSortOrderEnum = "ASC"
	ListRecipesSortOrderDesc ListRecipesSortOrderEnum = "DESC"
)

var mappingListRecipesSortOrderEnum = map[string]ListRecipesSortOrderEnum{
	"ASC":  ListRecipesSortOrderAsc,
	"DESC": ListRecipesSortOrderDesc,
}

var mappingListRecipesSortOrderEnumLowerCase = map[string]ListRecipesSortOrderEnum{
	"asc":  ListRecipesSortOrderAsc,
	"desc": ListRecipesSortOrderDesc,
}

// GetListRecipesSortOrderEnumValues Enumerates the set of values for ListRecipesSortOrderEnum
func GetListRecipesSortOrderEnumValues() []ListRecipesSortOrderEnum {
	values := make([]ListRecipesSortOrderEnum, 0)
	for _, v := range mappingListRecipesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecipesSortOrderEnumStringValues Enumerates the set of values in String for ListRecipesSortOrderEnum
func GetListRecipesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRecipesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecipesSortOrderEnum(val string) (ListRecipesSortOrderEnum, bool) {
	enum, ok := mappingListRecipesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecipesSortByEnum Enum with underlying type: string
type ListRecipesSortByEnum string

// Set of constants representing the allowable values for ListRecipesSortByEnum
const (
	ListRecipesSortByTimecreated ListRecipesSortByEnum = "timeCreated"
	ListRecipesSortByDisplayname ListRecipesSortByEnum = "displayName"
)

var mappingListRecipesSortByEnum = map[string]ListRecipesSortByEnum{
	"timeCreated": ListRecipesSortByTimecreated,
	"displayName": ListRecipesSortByDisplayname,
}

var mappingListRecipesSortByEnumLowerCase = map[string]ListRecipesSortByEnum{
	"timecreated": ListRecipesSortByTimecreated,
	"displayname": ListRecipesSortByDisplayname,
}

// GetListRecipesSortByEnumValues Enumerates the set of values for ListRecipesSortByEnum
func GetListRecipesSortByEnumValues() []ListRecipesSortByEnum {
	values := make([]ListRecipesSortByEnum, 0)
	for _, v := range mappingListRecipesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecipesSortByEnumStringValues Enumerates the set of values in String for ListRecipesSortByEnum
func GetListRecipesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListRecipesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecipesSortByEnum(val string) (ListRecipesSortByEnum, bool) {
	enum, ok := mappingListRecipesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
