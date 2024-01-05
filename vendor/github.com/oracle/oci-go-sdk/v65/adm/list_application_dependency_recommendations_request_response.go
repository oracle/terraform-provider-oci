// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package adm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListApplicationDependencyRecommendationsRequest wrapper for the ListApplicationDependencyRecommendations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListApplicationDependencyRecommendations.go.html to see an example of how to use ListApplicationDependencyRecommendationsRequest.
type ListApplicationDependencyRecommendationsRequest struct {

	// Unique Remediation Run identifier path parameter.
	RemediationRunId *string `mandatory:"true" contributesTo:"path" name:"remediationRunId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListApplicationDependencyRecommendationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire GAV (Group Artifact Version) identifier given.
	Gav *string `mandatory:"false" contributesTo:"query" name:"gav"`

	// The field to sort by. Only one sort order may be provided.
	// If sort order is dfs, the nodes are returned by going through the application dependency tree in a depth-first manner. Children are sorted based on their GAV property alphabetically (either ascending or descending, depending on the order parameter). Default order is ascending.
	// If sort order is bfs, the nodes are returned by going through the application dependency tree in a breadth-first manner. Children are sorted based on their GAV property alphabetically (either ascending or descending, depending on the order parameter). Default order is ascending.
	// Default order for gav is ascending where ascending corresponds to alphanumerical order.
	// Default order for nodeId is ascending where ascending corresponds to alphanumerical order.
	// Sorting by DFS or BFS cannot be used in conjunction with the following query parameters: "gav", "cvssV2GreaterThanOrEqual", "cvssV3GreaterThanOrEqual" and "vulnerabilityId".
	SortBy ListApplicationDependencyRecommendationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListApplicationDependencyRecommendationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApplicationDependencyRecommendationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListApplicationDependencyRecommendationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApplicationDependencyRecommendationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListApplicationDependencyRecommendationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListApplicationDependencyRecommendationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListApplicationDependencyRecommendationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApplicationDependencyRecommendationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListApplicationDependencyRecommendationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListApplicationDependencyRecommendationsResponse wrapper for the ListApplicationDependencyRecommendations operation
type ListApplicationDependencyRecommendationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ApplicationDependencyRecommendationCollection instances
	ApplicationDependencyRecommendationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListApplicationDependencyRecommendationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApplicationDependencyRecommendationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApplicationDependencyRecommendationsSortOrderEnum Enum with underlying type: string
type ListApplicationDependencyRecommendationsSortOrderEnum string

// Set of constants representing the allowable values for ListApplicationDependencyRecommendationsSortOrderEnum
const (
	ListApplicationDependencyRecommendationsSortOrderAsc  ListApplicationDependencyRecommendationsSortOrderEnum = "ASC"
	ListApplicationDependencyRecommendationsSortOrderDesc ListApplicationDependencyRecommendationsSortOrderEnum = "DESC"
)

var mappingListApplicationDependencyRecommendationsSortOrderEnum = map[string]ListApplicationDependencyRecommendationsSortOrderEnum{
	"ASC":  ListApplicationDependencyRecommendationsSortOrderAsc,
	"DESC": ListApplicationDependencyRecommendationsSortOrderDesc,
}

var mappingListApplicationDependencyRecommendationsSortOrderEnumLowerCase = map[string]ListApplicationDependencyRecommendationsSortOrderEnum{
	"asc":  ListApplicationDependencyRecommendationsSortOrderAsc,
	"desc": ListApplicationDependencyRecommendationsSortOrderDesc,
}

// GetListApplicationDependencyRecommendationsSortOrderEnumValues Enumerates the set of values for ListApplicationDependencyRecommendationsSortOrderEnum
func GetListApplicationDependencyRecommendationsSortOrderEnumValues() []ListApplicationDependencyRecommendationsSortOrderEnum {
	values := make([]ListApplicationDependencyRecommendationsSortOrderEnum, 0)
	for _, v := range mappingListApplicationDependencyRecommendationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplicationDependencyRecommendationsSortOrderEnumStringValues Enumerates the set of values in String for ListApplicationDependencyRecommendationsSortOrderEnum
func GetListApplicationDependencyRecommendationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListApplicationDependencyRecommendationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplicationDependencyRecommendationsSortOrderEnum(val string) (ListApplicationDependencyRecommendationsSortOrderEnum, bool) {
	enum, ok := mappingListApplicationDependencyRecommendationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListApplicationDependencyRecommendationsSortByEnum Enum with underlying type: string
type ListApplicationDependencyRecommendationsSortByEnum string

// Set of constants representing the allowable values for ListApplicationDependencyRecommendationsSortByEnum
const (
	ListApplicationDependencyRecommendationsSortByGav    ListApplicationDependencyRecommendationsSortByEnum = "gav"
	ListApplicationDependencyRecommendationsSortByNodeid ListApplicationDependencyRecommendationsSortByEnum = "nodeId"
	ListApplicationDependencyRecommendationsSortByDfs    ListApplicationDependencyRecommendationsSortByEnum = "dfs"
	ListApplicationDependencyRecommendationsSortByBfs    ListApplicationDependencyRecommendationsSortByEnum = "bfs"
)

var mappingListApplicationDependencyRecommendationsSortByEnum = map[string]ListApplicationDependencyRecommendationsSortByEnum{
	"gav":    ListApplicationDependencyRecommendationsSortByGav,
	"nodeId": ListApplicationDependencyRecommendationsSortByNodeid,
	"dfs":    ListApplicationDependencyRecommendationsSortByDfs,
	"bfs":    ListApplicationDependencyRecommendationsSortByBfs,
}

var mappingListApplicationDependencyRecommendationsSortByEnumLowerCase = map[string]ListApplicationDependencyRecommendationsSortByEnum{
	"gav":    ListApplicationDependencyRecommendationsSortByGav,
	"nodeid": ListApplicationDependencyRecommendationsSortByNodeid,
	"dfs":    ListApplicationDependencyRecommendationsSortByDfs,
	"bfs":    ListApplicationDependencyRecommendationsSortByBfs,
}

// GetListApplicationDependencyRecommendationsSortByEnumValues Enumerates the set of values for ListApplicationDependencyRecommendationsSortByEnum
func GetListApplicationDependencyRecommendationsSortByEnumValues() []ListApplicationDependencyRecommendationsSortByEnum {
	values := make([]ListApplicationDependencyRecommendationsSortByEnum, 0)
	for _, v := range mappingListApplicationDependencyRecommendationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplicationDependencyRecommendationsSortByEnumStringValues Enumerates the set of values in String for ListApplicationDependencyRecommendationsSortByEnum
func GetListApplicationDependencyRecommendationsSortByEnumStringValues() []string {
	return []string{
		"gav",
		"nodeId",
		"dfs",
		"bfs",
	}
}

// GetMappingListApplicationDependencyRecommendationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplicationDependencyRecommendationsSortByEnum(val string) (ListApplicationDependencyRecommendationsSortByEnum, bool) {
	enum, ok := mappingListApplicationDependencyRecommendationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
