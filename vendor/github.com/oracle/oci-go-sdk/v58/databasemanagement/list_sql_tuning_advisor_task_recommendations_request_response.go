// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListSqlTuningAdvisorTaskRecommendationsRequest wrapper for the ListSqlTuningAdvisorTaskRecommendations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListSqlTuningAdvisorTaskRecommendations.go.html to see an example of how to use ListSqlTuningAdvisorTaskRecommendationsRequest.
type ListSqlTuningAdvisorTaskRecommendationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The SQL tuning task identifier. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskId *int64 `mandatory:"true" contributesTo:"path" name:"sqlTuningAdvisorTaskId"`

	// The SQL object ID for the SQL tuning task. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlObjectId *int64 `mandatory:"true" contributesTo:"query" name:"sqlObjectId"`

	// The execution ID for an execution of a SQL tuning task. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ExecutionId *int64 `mandatory:"true" contributesTo:"query" name:"executionId"`

	// The possible sortBy values of an object's recommendations.
	SortBy ListSqlTuningAdvisorTaskRecommendationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Descending order is the default order.
	SortOrder ListSqlTuningAdvisorTaskRecommendationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlTuningAdvisorTaskRecommendationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlTuningAdvisorTaskRecommendationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlTuningAdvisorTaskRecommendationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlTuningAdvisorTaskRecommendationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlTuningAdvisorTaskRecommendationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlTuningAdvisorTaskRecommendationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSqlTuningAdvisorTaskRecommendationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlTuningAdvisorTaskRecommendationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSqlTuningAdvisorTaskRecommendationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlTuningAdvisorTaskRecommendationsResponse wrapper for the ListSqlTuningAdvisorTaskRecommendations operation
type ListSqlTuningAdvisorTaskRecommendationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlTuningAdvisorTaskRecommendationCollection instances
	SqlTuningAdvisorTaskRecommendationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSqlTuningAdvisorTaskRecommendationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlTuningAdvisorTaskRecommendationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlTuningAdvisorTaskRecommendationsSortByEnum Enum with underlying type: string
type ListSqlTuningAdvisorTaskRecommendationsSortByEnum string

// Set of constants representing the allowable values for ListSqlTuningAdvisorTaskRecommendationsSortByEnum
const (
	ListSqlTuningAdvisorTaskRecommendationsSortByRecommendationType ListSqlTuningAdvisorTaskRecommendationsSortByEnum = "RECOMMENDATION_TYPE"
	ListSqlTuningAdvisorTaskRecommendationsSortByBenefit            ListSqlTuningAdvisorTaskRecommendationsSortByEnum = "BENEFIT"
)

var mappingListSqlTuningAdvisorTaskRecommendationsSortByEnum = map[string]ListSqlTuningAdvisorTaskRecommendationsSortByEnum{
	"RECOMMENDATION_TYPE": ListSqlTuningAdvisorTaskRecommendationsSortByRecommendationType,
	"BENEFIT":             ListSqlTuningAdvisorTaskRecommendationsSortByBenefit,
}

// GetListSqlTuningAdvisorTaskRecommendationsSortByEnumValues Enumerates the set of values for ListSqlTuningAdvisorTaskRecommendationsSortByEnum
func GetListSqlTuningAdvisorTaskRecommendationsSortByEnumValues() []ListSqlTuningAdvisorTaskRecommendationsSortByEnum {
	values := make([]ListSqlTuningAdvisorTaskRecommendationsSortByEnum, 0)
	for _, v := range mappingListSqlTuningAdvisorTaskRecommendationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlTuningAdvisorTaskRecommendationsSortByEnumStringValues Enumerates the set of values in String for ListSqlTuningAdvisorTaskRecommendationsSortByEnum
func GetListSqlTuningAdvisorTaskRecommendationsSortByEnumStringValues() []string {
	return []string{
		"RECOMMENDATION_TYPE",
		"BENEFIT",
	}
}

// GetMappingListSqlTuningAdvisorTaskRecommendationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlTuningAdvisorTaskRecommendationsSortByEnum(val string) (ListSqlTuningAdvisorTaskRecommendationsSortByEnum, bool) {
	mappingListSqlTuningAdvisorTaskRecommendationsSortByEnumIgnoreCase := make(map[string]ListSqlTuningAdvisorTaskRecommendationsSortByEnum)
	for k, v := range mappingListSqlTuningAdvisorTaskRecommendationsSortByEnum {
		mappingListSqlTuningAdvisorTaskRecommendationsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSqlTuningAdvisorTaskRecommendationsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlTuningAdvisorTaskRecommendationsSortOrderEnum Enum with underlying type: string
type ListSqlTuningAdvisorTaskRecommendationsSortOrderEnum string

// Set of constants representing the allowable values for ListSqlTuningAdvisorTaskRecommendationsSortOrderEnum
const (
	ListSqlTuningAdvisorTaskRecommendationsSortOrderAsc  ListSqlTuningAdvisorTaskRecommendationsSortOrderEnum = "ASC"
	ListSqlTuningAdvisorTaskRecommendationsSortOrderDesc ListSqlTuningAdvisorTaskRecommendationsSortOrderEnum = "DESC"
)

var mappingListSqlTuningAdvisorTaskRecommendationsSortOrderEnum = map[string]ListSqlTuningAdvisorTaskRecommendationsSortOrderEnum{
	"ASC":  ListSqlTuningAdvisorTaskRecommendationsSortOrderAsc,
	"DESC": ListSqlTuningAdvisorTaskRecommendationsSortOrderDesc,
}

// GetListSqlTuningAdvisorTaskRecommendationsSortOrderEnumValues Enumerates the set of values for ListSqlTuningAdvisorTaskRecommendationsSortOrderEnum
func GetListSqlTuningAdvisorTaskRecommendationsSortOrderEnumValues() []ListSqlTuningAdvisorTaskRecommendationsSortOrderEnum {
	values := make([]ListSqlTuningAdvisorTaskRecommendationsSortOrderEnum, 0)
	for _, v := range mappingListSqlTuningAdvisorTaskRecommendationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlTuningAdvisorTaskRecommendationsSortOrderEnumStringValues Enumerates the set of values in String for ListSqlTuningAdvisorTaskRecommendationsSortOrderEnum
func GetListSqlTuningAdvisorTaskRecommendationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSqlTuningAdvisorTaskRecommendationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlTuningAdvisorTaskRecommendationsSortOrderEnum(val string) (ListSqlTuningAdvisorTaskRecommendationsSortOrderEnum, bool) {
	mappingListSqlTuningAdvisorTaskRecommendationsSortOrderEnumIgnoreCase := make(map[string]ListSqlTuningAdvisorTaskRecommendationsSortOrderEnum)
	for k, v := range mappingListSqlTuningAdvisorTaskRecommendationsSortOrderEnum {
		mappingListSqlTuningAdvisorTaskRecommendationsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSqlTuningAdvisorTaskRecommendationsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
