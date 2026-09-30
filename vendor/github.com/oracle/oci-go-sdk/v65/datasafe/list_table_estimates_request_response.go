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

// ListTableEstimatesRequest wrapper for the ListTableEstimates operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTableEstimates.go.html to see an example of how to use ListTableEstimatesRequest.
type ListTableEstimatesRequest struct {

	// The OCID of the subsetting policy.
	SubsettingPolicyId *string `mandatory:"true" contributesTo:"path" name:"subsettingPolicyId"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object name.
	ObjectName []string `contributesTo:"query" name:"objectName" collectionFormat:"multi"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListTableEstimatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for schemaName is ascending.
	SortBy ListTableEstimatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTableEstimatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTableEstimatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTableEstimatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTableEstimatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTableEstimatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTableEstimatesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTableEstimatesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTableEstimatesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTableEstimatesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTableEstimatesResponse wrapper for the ListTableEstimates operation
type ListTableEstimatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TableEstimateCollection instances
	TableEstimateCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListTableEstimatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTableEstimatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTableEstimatesSortOrderEnum Enum with underlying type: string
type ListTableEstimatesSortOrderEnum string

// Set of constants representing the allowable values for ListTableEstimatesSortOrderEnum
const (
	ListTableEstimatesSortOrderAsc  ListTableEstimatesSortOrderEnum = "ASC"
	ListTableEstimatesSortOrderDesc ListTableEstimatesSortOrderEnum = "DESC"
)

var mappingListTableEstimatesSortOrderEnum = map[string]ListTableEstimatesSortOrderEnum{
	"ASC":  ListTableEstimatesSortOrderAsc,
	"DESC": ListTableEstimatesSortOrderDesc,
}

var mappingListTableEstimatesSortOrderEnumLowerCase = map[string]ListTableEstimatesSortOrderEnum{
	"asc":  ListTableEstimatesSortOrderAsc,
	"desc": ListTableEstimatesSortOrderDesc,
}

// GetListTableEstimatesSortOrderEnumValues Enumerates the set of values for ListTableEstimatesSortOrderEnum
func GetListTableEstimatesSortOrderEnumValues() []ListTableEstimatesSortOrderEnum {
	values := make([]ListTableEstimatesSortOrderEnum, 0)
	for _, v := range mappingListTableEstimatesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTableEstimatesSortOrderEnumStringValues Enumerates the set of values in String for ListTableEstimatesSortOrderEnum
func GetListTableEstimatesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTableEstimatesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTableEstimatesSortOrderEnum(val string) (ListTableEstimatesSortOrderEnum, bool) {
	enum, ok := mappingListTableEstimatesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTableEstimatesSortByEnum Enum with underlying type: string
type ListTableEstimatesSortByEnum string

// Set of constants representing the allowable values for ListTableEstimatesSortByEnum
const (
	ListTableEstimatesSortBySchemaname ListTableEstimatesSortByEnum = "schemaName"
	ListTableEstimatesSortByTablename  ListTableEstimatesSortByEnum = "tableName"
)

var mappingListTableEstimatesSortByEnum = map[string]ListTableEstimatesSortByEnum{
	"schemaName": ListTableEstimatesSortBySchemaname,
	"tableName":  ListTableEstimatesSortByTablename,
}

var mappingListTableEstimatesSortByEnumLowerCase = map[string]ListTableEstimatesSortByEnum{
	"schemaname": ListTableEstimatesSortBySchemaname,
	"tablename":  ListTableEstimatesSortByTablename,
}

// GetListTableEstimatesSortByEnumValues Enumerates the set of values for ListTableEstimatesSortByEnum
func GetListTableEstimatesSortByEnumValues() []ListTableEstimatesSortByEnum {
	values := make([]ListTableEstimatesSortByEnum, 0)
	for _, v := range mappingListTableEstimatesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTableEstimatesSortByEnumStringValues Enumerates the set of values in String for ListTableEstimatesSortByEnum
func GetListTableEstimatesSortByEnumStringValues() []string {
	return []string{
		"schemaName",
		"tableName",
	}
}

// GetMappingListTableEstimatesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTableEstimatesSortByEnum(val string) (ListTableEstimatesSortByEnum, bool) {
	enum, ok := mappingListTableEstimatesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
