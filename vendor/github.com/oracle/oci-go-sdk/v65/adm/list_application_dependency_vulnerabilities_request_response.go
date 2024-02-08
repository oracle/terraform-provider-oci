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

// ListApplicationDependencyVulnerabilitiesRequest wrapper for the ListApplicationDependencyVulnerabilities operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListApplicationDependencyVulnerabilities.go.html to see an example of how to use ListApplicationDependencyVulnerabilitiesRequest.
type ListApplicationDependencyVulnerabilitiesRequest struct {

	// Unique Vulnerability Audit identifier path parameter.
	VulnerabilityAuditId *string `mandatory:"true" contributesTo:"path" name:"vulnerabilityAuditId"`

	// A filter to return only Vulnerability Audits that match the specified id.
	VulnerabilityId *string `mandatory:"false" contributesTo:"query" name:"vulnerabilityId"`

	// A filter that returns only Vulnerabilities that have a Common Vulnerability Scoring System Version 3 (CVSS V3) greater than or equal to the specified value.
	CvssV3GreaterThanOrEqual *float32 `mandatory:"false" contributesTo:"query" name:"cvssV3GreaterThanOrEqual"`

	// A filter that returns only Vulnerabilities that have a Common Vulnerability Scoring System Version 2 (CVSS V2) greater than or equal to the specified value.
	CvssV2GreaterThanOrEqual *float32 `mandatory:"false" contributesTo:"query" name:"cvssV2GreaterThanOrEqual"`

	// A filter that returns only Vulnerabilities that have a severity greater than or equal to the specified value.
	SeverityGreaterThanOrEqual ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum `mandatory:"false" contributesTo:"query" name:"severityGreaterThanOrEqual" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListApplicationDependencyVulnerabilitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// If sort order is dfs, the nodes are returned by going through the application dependency tree in a depth-first manner. Children are sorted based on their GAV property alphabetically (either ascending or descending, depending on the order parameter). Default order is ascending.
	// If sort order is bfs, the nodes are returned by going through the application dependency tree in a breadth-first manner. Children are sorted based on their GAV property alphabetically (either ascending or descending, depending on the order parameter). Default order is ascending.
	// Default order for gav is ascending where ascending corresponds to alphanumerical order.
	// Default order for purl is ascending where ascending corresponds to alphabetical order
	// Default order for nodeId is ascending where ascending corresponds to alphanumerical order.
	// Sorting by DFS or BFS cannot be used in conjunction with the following query parameters: "gav", "cvssV2GreaterThanOrEqual", "cvssV3GreaterThanOrEqual" and "vulnerabilityId".
	SortBy ListApplicationDependencyVulnerabilitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to override the top level root identifier with the new given value. The application dependency tree will only be traversed from the given node.
	// Query parameters "cvssV2GreaterThanOrEqual", "cvssV3GreaterThanOrEqual", "gav" and "vulnerabilityId" cannot be used in conjunction with this parameter.
	RootNodeId *string `mandatory:"false" contributesTo:"query" name:"rootNodeId"`

	// A filter to limit depth of the application dependencies tree traversal.
	// Additionally query parameters such as "cvssV2GreaterThanOrEqual", "cvssV3GreaterThanOrEqual", "gav" and "vulnerabilityId" can't be used in conjunction with this latter.
	Depth *int `mandatory:"false" contributesTo:"query" name:"depth"`

	// A filter to return only resources that match the entire GAV (Group Artifact Version) identifier given.
	Gav *string `mandatory:"false" contributesTo:"query" name:"gav"`

	// A filter to return only resources that match the entire PURL given (https://github.com/package-url/purl-spec/).
	Purl *string `mandatory:"false" contributesTo:"query" name:"purl"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListApplicationDependencyVulnerabilitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApplicationDependencyVulnerabilitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListApplicationDependencyVulnerabilitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApplicationDependencyVulnerabilitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListApplicationDependencyVulnerabilitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum(string(request.SeverityGreaterThanOrEqual)); !ok && request.SeverityGreaterThanOrEqual != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SeverityGreaterThanOrEqual: %s. Supported values are: %s.", request.SeverityGreaterThanOrEqual, strings.Join(GetListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApplicationDependencyVulnerabilitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListApplicationDependencyVulnerabilitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApplicationDependencyVulnerabilitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListApplicationDependencyVulnerabilitiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListApplicationDependencyVulnerabilitiesResponse wrapper for the ListApplicationDependencyVulnerabilities operation
type ListApplicationDependencyVulnerabilitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ApplicationDependencyVulnerabilityCollection instances
	ApplicationDependencyVulnerabilityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListApplicationDependencyVulnerabilitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApplicationDependencyVulnerabilitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum Enum with underlying type: string
type ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum string

// Set of constants representing the allowable values for ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum
const (
	ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualNone     ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum = "NONE"
	ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualLow      ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum = "LOW"
	ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualMedium   ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum = "MEDIUM"
	ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualHigh     ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum = "HIGH"
	ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualCritical ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum = "CRITICAL"
)

var mappingListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum = map[string]ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum{
	"NONE":     ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualNone,
	"LOW":      ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualLow,
	"MEDIUM":   ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualMedium,
	"HIGH":     ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualHigh,
	"CRITICAL": ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualCritical,
}

var mappingListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnumLowerCase = map[string]ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum{
	"none":     ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualNone,
	"low":      ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualLow,
	"medium":   ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualMedium,
	"high":     ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualHigh,
	"critical": ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualCritical,
}

// GetListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnumValues Enumerates the set of values for ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum
func GetListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnumValues() []ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum {
	values := make([]ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum, 0)
	for _, v := range mappingListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnumStringValues Enumerates the set of values in String for ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum
func GetListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnumStringValues() []string {
	return []string{
		"NONE",
		"LOW",
		"MEDIUM",
		"HIGH",
		"CRITICAL",
	}
}

// GetMappingListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum(val string) (ListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnum, bool) {
	enum, ok := mappingListApplicationDependencyVulnerabilitiesSeverityGreaterThanOrEqualEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListApplicationDependencyVulnerabilitiesSortOrderEnum Enum with underlying type: string
type ListApplicationDependencyVulnerabilitiesSortOrderEnum string

// Set of constants representing the allowable values for ListApplicationDependencyVulnerabilitiesSortOrderEnum
const (
	ListApplicationDependencyVulnerabilitiesSortOrderAsc  ListApplicationDependencyVulnerabilitiesSortOrderEnum = "ASC"
	ListApplicationDependencyVulnerabilitiesSortOrderDesc ListApplicationDependencyVulnerabilitiesSortOrderEnum = "DESC"
)

var mappingListApplicationDependencyVulnerabilitiesSortOrderEnum = map[string]ListApplicationDependencyVulnerabilitiesSortOrderEnum{
	"ASC":  ListApplicationDependencyVulnerabilitiesSortOrderAsc,
	"DESC": ListApplicationDependencyVulnerabilitiesSortOrderDesc,
}

var mappingListApplicationDependencyVulnerabilitiesSortOrderEnumLowerCase = map[string]ListApplicationDependencyVulnerabilitiesSortOrderEnum{
	"asc":  ListApplicationDependencyVulnerabilitiesSortOrderAsc,
	"desc": ListApplicationDependencyVulnerabilitiesSortOrderDesc,
}

// GetListApplicationDependencyVulnerabilitiesSortOrderEnumValues Enumerates the set of values for ListApplicationDependencyVulnerabilitiesSortOrderEnum
func GetListApplicationDependencyVulnerabilitiesSortOrderEnumValues() []ListApplicationDependencyVulnerabilitiesSortOrderEnum {
	values := make([]ListApplicationDependencyVulnerabilitiesSortOrderEnum, 0)
	for _, v := range mappingListApplicationDependencyVulnerabilitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplicationDependencyVulnerabilitiesSortOrderEnumStringValues Enumerates the set of values in String for ListApplicationDependencyVulnerabilitiesSortOrderEnum
func GetListApplicationDependencyVulnerabilitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListApplicationDependencyVulnerabilitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplicationDependencyVulnerabilitiesSortOrderEnum(val string) (ListApplicationDependencyVulnerabilitiesSortOrderEnum, bool) {
	enum, ok := mappingListApplicationDependencyVulnerabilitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListApplicationDependencyVulnerabilitiesSortByEnum Enum with underlying type: string
type ListApplicationDependencyVulnerabilitiesSortByEnum string

// Set of constants representing the allowable values for ListApplicationDependencyVulnerabilitiesSortByEnum
const (
	ListApplicationDependencyVulnerabilitiesSortByGav    ListApplicationDependencyVulnerabilitiesSortByEnum = "gav"
	ListApplicationDependencyVulnerabilitiesSortByPurl   ListApplicationDependencyVulnerabilitiesSortByEnum = "purl"
	ListApplicationDependencyVulnerabilitiesSortByNodeid ListApplicationDependencyVulnerabilitiesSortByEnum = "nodeId"
	ListApplicationDependencyVulnerabilitiesSortByDfs    ListApplicationDependencyVulnerabilitiesSortByEnum = "dfs"
	ListApplicationDependencyVulnerabilitiesSortByBfs    ListApplicationDependencyVulnerabilitiesSortByEnum = "bfs"
)

var mappingListApplicationDependencyVulnerabilitiesSortByEnum = map[string]ListApplicationDependencyVulnerabilitiesSortByEnum{
	"gav":    ListApplicationDependencyVulnerabilitiesSortByGav,
	"purl":   ListApplicationDependencyVulnerabilitiesSortByPurl,
	"nodeId": ListApplicationDependencyVulnerabilitiesSortByNodeid,
	"dfs":    ListApplicationDependencyVulnerabilitiesSortByDfs,
	"bfs":    ListApplicationDependencyVulnerabilitiesSortByBfs,
}

var mappingListApplicationDependencyVulnerabilitiesSortByEnumLowerCase = map[string]ListApplicationDependencyVulnerabilitiesSortByEnum{
	"gav":    ListApplicationDependencyVulnerabilitiesSortByGav,
	"purl":   ListApplicationDependencyVulnerabilitiesSortByPurl,
	"nodeid": ListApplicationDependencyVulnerabilitiesSortByNodeid,
	"dfs":    ListApplicationDependencyVulnerabilitiesSortByDfs,
	"bfs":    ListApplicationDependencyVulnerabilitiesSortByBfs,
}

// GetListApplicationDependencyVulnerabilitiesSortByEnumValues Enumerates the set of values for ListApplicationDependencyVulnerabilitiesSortByEnum
func GetListApplicationDependencyVulnerabilitiesSortByEnumValues() []ListApplicationDependencyVulnerabilitiesSortByEnum {
	values := make([]ListApplicationDependencyVulnerabilitiesSortByEnum, 0)
	for _, v := range mappingListApplicationDependencyVulnerabilitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplicationDependencyVulnerabilitiesSortByEnumStringValues Enumerates the set of values in String for ListApplicationDependencyVulnerabilitiesSortByEnum
func GetListApplicationDependencyVulnerabilitiesSortByEnumStringValues() []string {
	return []string{
		"gav",
		"purl",
		"nodeId",
		"dfs",
		"bfs",
	}
}

// GetMappingListApplicationDependencyVulnerabilitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplicationDependencyVulnerabilitiesSortByEnum(val string) (ListApplicationDependencyVulnerabilitiesSortByEnum, bool) {
	enum, ok := mappingListApplicationDependencyVulnerabilitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
