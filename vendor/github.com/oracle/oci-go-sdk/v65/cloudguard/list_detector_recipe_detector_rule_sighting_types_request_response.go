// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDetectorRecipeDetectorRuleSightingTypesRequest wrapper for the ListDetectorRecipeDetectorRuleSightingTypes operation
type ListDetectorRecipeDetectorRuleSightingTypesRequest struct {

	// DetectorRecipe OCID
	DetectorRecipeId *string `mandatory:"true" contributesTo:"path" name:"detectorRecipeId"`

	// The key of Detector Rule.
	DetectorRuleId *string `mandatory:"true" contributesTo:"path" name:"detectorRuleId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListDetectorRecipeDetectorRuleSightingTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDetectorRecipeDetectorRuleSightingTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDetectorRecipeDetectorRuleSightingTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDetectorRecipeDetectorRuleSightingTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDetectorRecipeDetectorRuleSightingTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDetectorRecipeDetectorRuleSightingTypesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDetectorRecipeDetectorRuleSightingTypesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDetectorRecipeDetectorRuleSightingTypesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDetectorRecipeDetectorRuleSightingTypesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDetectorRecipeDetectorRuleSightingTypesResponse wrapper for the ListDetectorRecipeDetectorRuleSightingTypes operation
type ListDetectorRecipeDetectorRuleSightingTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DetectorRecipeDetectorRuleSightingTypeCollection instances
	DetectorRecipeDetectorRuleSightingTypeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDetectorRecipeDetectorRuleSightingTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDetectorRecipeDetectorRuleSightingTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum Enum with underlying type: string
type ListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum string

// Set of constants representing the allowable values for ListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum
const (
	ListDetectorRecipeDetectorRuleSightingTypesSortOrderAsc  ListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum = "ASC"
	ListDetectorRecipeDetectorRuleSightingTypesSortOrderDesc ListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum = "DESC"
)

var mappingListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum = map[string]ListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum{
	"ASC":  ListDetectorRecipeDetectorRuleSightingTypesSortOrderAsc,
	"DESC": ListDetectorRecipeDetectorRuleSightingTypesSortOrderDesc,
}

var mappingListDetectorRecipeDetectorRuleSightingTypesSortOrderEnumLowerCase = map[string]ListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum{
	"asc":  ListDetectorRecipeDetectorRuleSightingTypesSortOrderAsc,
	"desc": ListDetectorRecipeDetectorRuleSightingTypesSortOrderDesc,
}

// GetListDetectorRecipeDetectorRuleSightingTypesSortOrderEnumValues Enumerates the set of values for ListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum
func GetListDetectorRecipeDetectorRuleSightingTypesSortOrderEnumValues() []ListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum {
	values := make([]ListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum, 0)
	for _, v := range mappingListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDetectorRecipeDetectorRuleSightingTypesSortOrderEnumStringValues Enumerates the set of values in String for ListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum
func GetListDetectorRecipeDetectorRuleSightingTypesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum(val string) (ListDetectorRecipeDetectorRuleSightingTypesSortOrderEnum, bool) {
	enum, ok := mappingListDetectorRecipeDetectorRuleSightingTypesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDetectorRecipeDetectorRuleSightingTypesSortByEnum Enum with underlying type: string
type ListDetectorRecipeDetectorRuleSightingTypesSortByEnum string

// Set of constants representing the allowable values for ListDetectorRecipeDetectorRuleSightingTypesSortByEnum
const (
	ListDetectorRecipeDetectorRuleSightingTypesSortByTimecreated ListDetectorRecipeDetectorRuleSightingTypesSortByEnum = "timeCreated"
)

var mappingListDetectorRecipeDetectorRuleSightingTypesSortByEnum = map[string]ListDetectorRecipeDetectorRuleSightingTypesSortByEnum{
	"timeCreated": ListDetectorRecipeDetectorRuleSightingTypesSortByTimecreated,
}

var mappingListDetectorRecipeDetectorRuleSightingTypesSortByEnumLowerCase = map[string]ListDetectorRecipeDetectorRuleSightingTypesSortByEnum{
	"timecreated": ListDetectorRecipeDetectorRuleSightingTypesSortByTimecreated,
}

// GetListDetectorRecipeDetectorRuleSightingTypesSortByEnumValues Enumerates the set of values for ListDetectorRecipeDetectorRuleSightingTypesSortByEnum
func GetListDetectorRecipeDetectorRuleSightingTypesSortByEnumValues() []ListDetectorRecipeDetectorRuleSightingTypesSortByEnum {
	values := make([]ListDetectorRecipeDetectorRuleSightingTypesSortByEnum, 0)
	for _, v := range mappingListDetectorRecipeDetectorRuleSightingTypesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDetectorRecipeDetectorRuleSightingTypesSortByEnumStringValues Enumerates the set of values in String for ListDetectorRecipeDetectorRuleSightingTypesSortByEnum
func GetListDetectorRecipeDetectorRuleSightingTypesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListDetectorRecipeDetectorRuleSightingTypesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDetectorRecipeDetectorRuleSightingTypesSortByEnum(val string) (ListDetectorRecipeDetectorRuleSightingTypesSortByEnum, bool) {
	enum, ok := mappingListDetectorRecipeDetectorRuleSightingTypesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
