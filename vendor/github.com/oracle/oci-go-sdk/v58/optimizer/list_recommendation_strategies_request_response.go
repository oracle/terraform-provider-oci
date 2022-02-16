// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package optimizer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListRecommendationStrategiesRequest wrapper for the ListRecommendationStrategies operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListRecommendationStrategies.go.html to see an example of how to use ListRecommendationStrategiesRequest.
type ListRecommendationStrategiesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.
	// Can only be set to true when performing ListCompartments on the tenancy (root compartment).
	CompartmentIdInSubtree *bool `mandatory:"true" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Optional. A filter that returns results that match the name specified.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Optional. A filter that returns results that match the recommendation name specified.
	RecommendationName *string `mandatory:"false" contributesTo:"query" name:"recommendationName"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListRecommendationStrategiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for NAME is ascending. The NAME sort order is case sensitive.
	SortBy ListRecommendationStrategiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRecommendationStrategiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRecommendationStrategiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRecommendationStrategiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRecommendationStrategiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRecommendationStrategiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRecommendationStrategiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRecommendationStrategiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecommendationStrategiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRecommendationStrategiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRecommendationStrategiesResponse wrapper for the ListRecommendationStrategies operation
type ListRecommendationStrategiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecommendationStrategyCollection instances
	RecommendationStrategyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results exist.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListRecommendationStrategiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRecommendationStrategiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRecommendationStrategiesSortOrderEnum Enum with underlying type: string
type ListRecommendationStrategiesSortOrderEnum string

// Set of constants representing the allowable values for ListRecommendationStrategiesSortOrderEnum
const (
	ListRecommendationStrategiesSortOrderAsc  ListRecommendationStrategiesSortOrderEnum = "ASC"
	ListRecommendationStrategiesSortOrderDesc ListRecommendationStrategiesSortOrderEnum = "DESC"
)

var mappingListRecommendationStrategiesSortOrderEnum = map[string]ListRecommendationStrategiesSortOrderEnum{
	"ASC":  ListRecommendationStrategiesSortOrderAsc,
	"DESC": ListRecommendationStrategiesSortOrderDesc,
}

// GetListRecommendationStrategiesSortOrderEnumValues Enumerates the set of values for ListRecommendationStrategiesSortOrderEnum
func GetListRecommendationStrategiesSortOrderEnumValues() []ListRecommendationStrategiesSortOrderEnum {
	values := make([]ListRecommendationStrategiesSortOrderEnum, 0)
	for _, v := range mappingListRecommendationStrategiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendationStrategiesSortOrderEnumStringValues Enumerates the set of values in String for ListRecommendationStrategiesSortOrderEnum
func GetListRecommendationStrategiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRecommendationStrategiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendationStrategiesSortOrderEnum(val string) (ListRecommendationStrategiesSortOrderEnum, bool) {
	mappingListRecommendationStrategiesSortOrderEnumIgnoreCase := make(map[string]ListRecommendationStrategiesSortOrderEnum)
	for k, v := range mappingListRecommendationStrategiesSortOrderEnum {
		mappingListRecommendationStrategiesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRecommendationStrategiesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecommendationStrategiesSortByEnum Enum with underlying type: string
type ListRecommendationStrategiesSortByEnum string

// Set of constants representing the allowable values for ListRecommendationStrategiesSortByEnum
const (
	ListRecommendationStrategiesSortByName        ListRecommendationStrategiesSortByEnum = "NAME"
	ListRecommendationStrategiesSortByTimecreated ListRecommendationStrategiesSortByEnum = "TIMECREATED"
)

var mappingListRecommendationStrategiesSortByEnum = map[string]ListRecommendationStrategiesSortByEnum{
	"NAME":        ListRecommendationStrategiesSortByName,
	"TIMECREATED": ListRecommendationStrategiesSortByTimecreated,
}

// GetListRecommendationStrategiesSortByEnumValues Enumerates the set of values for ListRecommendationStrategiesSortByEnum
func GetListRecommendationStrategiesSortByEnumValues() []ListRecommendationStrategiesSortByEnum {
	values := make([]ListRecommendationStrategiesSortByEnum, 0)
	for _, v := range mappingListRecommendationStrategiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendationStrategiesSortByEnumStringValues Enumerates the set of values in String for ListRecommendationStrategiesSortByEnum
func GetListRecommendationStrategiesSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIMECREATED",
	}
}

// GetMappingListRecommendationStrategiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendationStrategiesSortByEnum(val string) (ListRecommendationStrategiesSortByEnum, bool) {
	mappingListRecommendationStrategiesSortByEnumIgnoreCase := make(map[string]ListRecommendationStrategiesSortByEnum)
	for k, v := range mappingListRecommendationStrategiesSortByEnum {
		mappingListRecommendationStrategiesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRecommendationStrategiesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
