// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAddmDbParameterCategoriesRequest wrapper for the ListAddmDbParameterCategories operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAddmDbParameterCategories.go.html to see an example of how to use ListAddmDbParameterCategoriesRequest.
type ListAddmDbParameterCategoriesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Optional list of database OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAddmDbParameterCategoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Field name for sorting the database parameter categories
	SortBy ListAddmDbParameterCategoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A list of tag filters to apply.  Only resources with a defined tag matching the value will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagEquals []string `contributesTo:"query" name:"definedTagEquals" collectionFormat:"multi"`

	// A list of tag filters to apply.  Only resources with a freeform tag matching the value will be returned.
	// The key for each tag is "{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same tag name are interpreted as "OR".  Values for different tag names are interpreted as "AND".
	FreeformTagEquals []string `contributesTo:"query" name:"freeformTagEquals" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified defined tags exist will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.true" (for checking existence of a defined tag)
	// or "{namespace}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagExists []string `contributesTo:"query" name:"definedTagExists" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified freeform tags exist the value will be returned.
	// The key for each tag is "{tagName}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for different tag names are interpreted as "AND".
	FreeformTagExists []string `contributesTo:"query" name:"freeformTagExists" collectionFormat:"multi"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAddmDbParameterCategoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAddmDbParameterCategoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAddmDbParameterCategoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAddmDbParameterCategoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAddmDbParameterCategoriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAddmDbParameterCategoriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAddmDbParameterCategoriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAddmDbParameterCategoriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAddmDbParameterCategoriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAddmDbParameterCategoriesResponse wrapper for the ListAddmDbParameterCategories operation
type ListAddmDbParameterCategoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AddmDbParameterCategoryCollection instances
	AddmDbParameterCategoryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAddmDbParameterCategoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAddmDbParameterCategoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAddmDbParameterCategoriesSortOrderEnum Enum with underlying type: string
type ListAddmDbParameterCategoriesSortOrderEnum string

// Set of constants representing the allowable values for ListAddmDbParameterCategoriesSortOrderEnum
const (
	ListAddmDbParameterCategoriesSortOrderAsc  ListAddmDbParameterCategoriesSortOrderEnum = "ASC"
	ListAddmDbParameterCategoriesSortOrderDesc ListAddmDbParameterCategoriesSortOrderEnum = "DESC"
)

var mappingListAddmDbParameterCategoriesSortOrderEnum = map[string]ListAddmDbParameterCategoriesSortOrderEnum{
	"ASC":  ListAddmDbParameterCategoriesSortOrderAsc,
	"DESC": ListAddmDbParameterCategoriesSortOrderDesc,
}

var mappingListAddmDbParameterCategoriesSortOrderEnumLowerCase = map[string]ListAddmDbParameterCategoriesSortOrderEnum{
	"asc":  ListAddmDbParameterCategoriesSortOrderAsc,
	"desc": ListAddmDbParameterCategoriesSortOrderDesc,
}

// GetListAddmDbParameterCategoriesSortOrderEnumValues Enumerates the set of values for ListAddmDbParameterCategoriesSortOrderEnum
func GetListAddmDbParameterCategoriesSortOrderEnumValues() []ListAddmDbParameterCategoriesSortOrderEnum {
	values := make([]ListAddmDbParameterCategoriesSortOrderEnum, 0)
	for _, v := range mappingListAddmDbParameterCategoriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAddmDbParameterCategoriesSortOrderEnumStringValues Enumerates the set of values in String for ListAddmDbParameterCategoriesSortOrderEnum
func GetListAddmDbParameterCategoriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAddmDbParameterCategoriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAddmDbParameterCategoriesSortOrderEnum(val string) (ListAddmDbParameterCategoriesSortOrderEnum, bool) {
	enum, ok := mappingListAddmDbParameterCategoriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAddmDbParameterCategoriesSortByEnum Enum with underlying type: string
type ListAddmDbParameterCategoriesSortByEnum string

// Set of constants representing the allowable values for ListAddmDbParameterCategoriesSortByEnum
const (
	ListAddmDbParameterCategoriesSortByName ListAddmDbParameterCategoriesSortByEnum = "name"
)

var mappingListAddmDbParameterCategoriesSortByEnum = map[string]ListAddmDbParameterCategoriesSortByEnum{
	"name": ListAddmDbParameterCategoriesSortByName,
}

var mappingListAddmDbParameterCategoriesSortByEnumLowerCase = map[string]ListAddmDbParameterCategoriesSortByEnum{
	"name": ListAddmDbParameterCategoriesSortByName,
}

// GetListAddmDbParameterCategoriesSortByEnumValues Enumerates the set of values for ListAddmDbParameterCategoriesSortByEnum
func GetListAddmDbParameterCategoriesSortByEnumValues() []ListAddmDbParameterCategoriesSortByEnum {
	values := make([]ListAddmDbParameterCategoriesSortByEnum, 0)
	for _, v := range mappingListAddmDbParameterCategoriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAddmDbParameterCategoriesSortByEnumStringValues Enumerates the set of values in String for ListAddmDbParameterCategoriesSortByEnum
func GetListAddmDbParameterCategoriesSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListAddmDbParameterCategoriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAddmDbParameterCategoriesSortByEnum(val string) (ListAddmDbParameterCategoriesSortByEnum, bool) {
	enum, ok := mappingListAddmDbParameterCategoriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
