// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRuntimeOperatorsRequest wrapper for the ListRuntimeOperators operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListRuntimeOperators.go.html to see an example of how to use ListRuntimeOperatorsRequest.
type ListRuntimeOperatorsRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The application key.
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// Runtime Pipeline Key
	RuntimePipelineKey *string `mandatory:"true" contributesTo:"path" name:"runtimePipelineKey"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Used to filter by the key of the object.
	Key []string `contributesTo:"query" name:"key" collectionFormat:"multi"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Used to filter by the identifier of the object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListRuntimeOperatorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListRuntimeOperatorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique type of the aggregator
	AggregatorType []string `contributesTo:"query" name:"aggregatorType" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRuntimeOperatorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRuntimeOperatorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRuntimeOperatorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRuntimeOperatorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRuntimeOperatorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRuntimeOperatorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRuntimeOperatorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRuntimeOperatorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRuntimeOperatorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRuntimeOperatorsResponse wrapper for the ListRuntimeOperators operation
type ListRuntimeOperatorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RuntimeOperatorSummaryCollection instances
	RuntimeOperatorSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `RuntimeOperator`s. If this header appears in the response, then this
	// is a partial list of RuntimeOperator. Include this value as the `page` parameter in a subsequent
	// GET request to get the next batch of RuntimeOperators.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRuntimeOperatorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRuntimeOperatorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRuntimeOperatorsSortOrderEnum Enum with underlying type: string
type ListRuntimeOperatorsSortOrderEnum string

// Set of constants representing the allowable values for ListRuntimeOperatorsSortOrderEnum
const (
	ListRuntimeOperatorsSortOrderAsc  ListRuntimeOperatorsSortOrderEnum = "ASC"
	ListRuntimeOperatorsSortOrderDesc ListRuntimeOperatorsSortOrderEnum = "DESC"
)

var mappingListRuntimeOperatorsSortOrderEnum = map[string]ListRuntimeOperatorsSortOrderEnum{
	"ASC":  ListRuntimeOperatorsSortOrderAsc,
	"DESC": ListRuntimeOperatorsSortOrderDesc,
}

var mappingListRuntimeOperatorsSortOrderEnumLowerCase = map[string]ListRuntimeOperatorsSortOrderEnum{
	"asc":  ListRuntimeOperatorsSortOrderAsc,
	"desc": ListRuntimeOperatorsSortOrderDesc,
}

// GetListRuntimeOperatorsSortOrderEnumValues Enumerates the set of values for ListRuntimeOperatorsSortOrderEnum
func GetListRuntimeOperatorsSortOrderEnumValues() []ListRuntimeOperatorsSortOrderEnum {
	values := make([]ListRuntimeOperatorsSortOrderEnum, 0)
	for _, v := range mappingListRuntimeOperatorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRuntimeOperatorsSortOrderEnumStringValues Enumerates the set of values in String for ListRuntimeOperatorsSortOrderEnum
func GetListRuntimeOperatorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRuntimeOperatorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRuntimeOperatorsSortOrderEnum(val string) (ListRuntimeOperatorsSortOrderEnum, bool) {
	enum, ok := mappingListRuntimeOperatorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRuntimeOperatorsSortByEnum Enum with underlying type: string
type ListRuntimeOperatorsSortByEnum string

// Set of constants representing the allowable values for ListRuntimeOperatorsSortByEnum
const (
	ListRuntimeOperatorsSortByTimeCreated ListRuntimeOperatorsSortByEnum = "TIME_CREATED"
	ListRuntimeOperatorsSortByDisplayName ListRuntimeOperatorsSortByEnum = "DISPLAY_NAME"
	ListRuntimeOperatorsSortByTimeUpdated ListRuntimeOperatorsSortByEnum = "TIME_UPDATED"
)

var mappingListRuntimeOperatorsSortByEnum = map[string]ListRuntimeOperatorsSortByEnum{
	"TIME_CREATED": ListRuntimeOperatorsSortByTimeCreated,
	"DISPLAY_NAME": ListRuntimeOperatorsSortByDisplayName,
	"TIME_UPDATED": ListRuntimeOperatorsSortByTimeUpdated,
}

var mappingListRuntimeOperatorsSortByEnumLowerCase = map[string]ListRuntimeOperatorsSortByEnum{
	"time_created": ListRuntimeOperatorsSortByTimeCreated,
	"display_name": ListRuntimeOperatorsSortByDisplayName,
	"time_updated": ListRuntimeOperatorsSortByTimeUpdated,
}

// GetListRuntimeOperatorsSortByEnumValues Enumerates the set of values for ListRuntimeOperatorsSortByEnum
func GetListRuntimeOperatorsSortByEnumValues() []ListRuntimeOperatorsSortByEnum {
	values := make([]ListRuntimeOperatorsSortByEnum, 0)
	for _, v := range mappingListRuntimeOperatorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRuntimeOperatorsSortByEnumStringValues Enumerates the set of values in String for ListRuntimeOperatorsSortByEnum
func GetListRuntimeOperatorsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListRuntimeOperatorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRuntimeOperatorsSortByEnum(val string) (ListRuntimeOperatorsSortByEnum, bool) {
	enum, ok := mappingListRuntimeOperatorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
