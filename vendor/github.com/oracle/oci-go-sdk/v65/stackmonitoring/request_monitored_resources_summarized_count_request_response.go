// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// RequestMonitoredResourcesSummarizedCountRequest wrapper for the RequestMonitoredResourcesSummarizedCount operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/RequestMonitoredResourcesSummarizedCount.go.html to see an example of how to use RequestMonitoredResourcesSummarizedCountRequest.
type RequestMonitoredResourcesSummarizedCountRequest struct {

	// The ID of the compartment in which data is listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The field to group by. Default group by is 'resourceType'.
	GroupBy RequestMonitoredResourcesSummarizedCountGroupByEnum `mandatory:"false" contributesTo:"query" name:"groupBy" omitEmpty:"true"`

	// Filter to return resource counts that match with the given licence edition.
	License RequestMonitoredResourcesSummarizedCountLicenseEnum `mandatory:"false" contributesTo:"query" name:"license" omitEmpty:"true"`

	// A filter to return resource counts that match exact resource type.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// If this query parameter is specified, the result is sorted by this query parameter value.
	SortBy RequestMonitoredResourcesSummarizedCountSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder RequestMonitoredResourcesSummarizedCountSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request RequestMonitoredResourcesSummarizedCountRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RequestMonitoredResourcesSummarizedCountRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request RequestMonitoredResourcesSummarizedCountRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RequestMonitoredResourcesSummarizedCountRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request RequestMonitoredResourcesSummarizedCountRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRequestMonitoredResourcesSummarizedCountGroupByEnum(string(request.GroupBy)); !ok && request.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", request.GroupBy, strings.Join(GetRequestMonitoredResourcesSummarizedCountGroupByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestMonitoredResourcesSummarizedCountLicenseEnum(string(request.License)); !ok && request.License != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for License: %s. Supported values are: %s.", request.License, strings.Join(GetRequestMonitoredResourcesSummarizedCountLicenseEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestMonitoredResourcesSummarizedCountSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetRequestMonitoredResourcesSummarizedCountSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestMonitoredResourcesSummarizedCountSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetRequestMonitoredResourcesSummarizedCountSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RequestMonitoredResourcesSummarizedCountResponse wrapper for the RequestMonitoredResourcesSummarizedCount operation
type RequestMonitoredResourcesSummarizedCountResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MonitoredResourcesCountAggregationCollection instances
	MonitoredResourcesCountAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. The total number of items in the result.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response RequestMonitoredResourcesSummarizedCountResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RequestMonitoredResourcesSummarizedCountResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// RequestMonitoredResourcesSummarizedCountGroupByEnum Enum with underlying type: string
type RequestMonitoredResourcesSummarizedCountGroupByEnum string

// Set of constants representing the allowable values for RequestMonitoredResourcesSummarizedCountGroupByEnum
const (
	RequestMonitoredResourcesSummarizedCountGroupByResourcetype     RequestMonitoredResourcesSummarizedCountGroupByEnum = "resourceType"
	RequestMonitoredResourcesSummarizedCountGroupByLicense          RequestMonitoredResourcesSummarizedCountGroupByEnum = "license"
	RequestMonitoredResourcesSummarizedCountGroupByParentresourceid RequestMonitoredResourcesSummarizedCountGroupByEnum = "parentResourceId"
)

var mappingRequestMonitoredResourcesSummarizedCountGroupByEnum = map[string]RequestMonitoredResourcesSummarizedCountGroupByEnum{
	"resourceType":     RequestMonitoredResourcesSummarizedCountGroupByResourcetype,
	"license":          RequestMonitoredResourcesSummarizedCountGroupByLicense,
	"parentResourceId": RequestMonitoredResourcesSummarizedCountGroupByParentresourceid,
}

var mappingRequestMonitoredResourcesSummarizedCountGroupByEnumLowerCase = map[string]RequestMonitoredResourcesSummarizedCountGroupByEnum{
	"resourcetype":     RequestMonitoredResourcesSummarizedCountGroupByResourcetype,
	"license":          RequestMonitoredResourcesSummarizedCountGroupByLicense,
	"parentresourceid": RequestMonitoredResourcesSummarizedCountGroupByParentresourceid,
}

// GetRequestMonitoredResourcesSummarizedCountGroupByEnumValues Enumerates the set of values for RequestMonitoredResourcesSummarizedCountGroupByEnum
func GetRequestMonitoredResourcesSummarizedCountGroupByEnumValues() []RequestMonitoredResourcesSummarizedCountGroupByEnum {
	values := make([]RequestMonitoredResourcesSummarizedCountGroupByEnum, 0)
	for _, v := range mappingRequestMonitoredResourcesSummarizedCountGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestMonitoredResourcesSummarizedCountGroupByEnumStringValues Enumerates the set of values in String for RequestMonitoredResourcesSummarizedCountGroupByEnum
func GetRequestMonitoredResourcesSummarizedCountGroupByEnumStringValues() []string {
	return []string{
		"resourceType",
		"license",
		"parentResourceId",
	}
}

// GetMappingRequestMonitoredResourcesSummarizedCountGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestMonitoredResourcesSummarizedCountGroupByEnum(val string) (RequestMonitoredResourcesSummarizedCountGroupByEnum, bool) {
	enum, ok := mappingRequestMonitoredResourcesSummarizedCountGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RequestMonitoredResourcesSummarizedCountLicenseEnum Enum with underlying type: string
type RequestMonitoredResourcesSummarizedCountLicenseEnum string

// Set of constants representing the allowable values for RequestMonitoredResourcesSummarizedCountLicenseEnum
const (
	RequestMonitoredResourcesSummarizedCountLicenseStandardEdition   RequestMonitoredResourcesSummarizedCountLicenseEnum = "STANDARD_EDITION"
	RequestMonitoredResourcesSummarizedCountLicenseEnterpriseEdition RequestMonitoredResourcesSummarizedCountLicenseEnum = "ENTERPRISE_EDITION"
)

var mappingRequestMonitoredResourcesSummarizedCountLicenseEnum = map[string]RequestMonitoredResourcesSummarizedCountLicenseEnum{
	"STANDARD_EDITION":   RequestMonitoredResourcesSummarizedCountLicenseStandardEdition,
	"ENTERPRISE_EDITION": RequestMonitoredResourcesSummarizedCountLicenseEnterpriseEdition,
}

var mappingRequestMonitoredResourcesSummarizedCountLicenseEnumLowerCase = map[string]RequestMonitoredResourcesSummarizedCountLicenseEnum{
	"standard_edition":   RequestMonitoredResourcesSummarizedCountLicenseStandardEdition,
	"enterprise_edition": RequestMonitoredResourcesSummarizedCountLicenseEnterpriseEdition,
}

// GetRequestMonitoredResourcesSummarizedCountLicenseEnumValues Enumerates the set of values for RequestMonitoredResourcesSummarizedCountLicenseEnum
func GetRequestMonitoredResourcesSummarizedCountLicenseEnumValues() []RequestMonitoredResourcesSummarizedCountLicenseEnum {
	values := make([]RequestMonitoredResourcesSummarizedCountLicenseEnum, 0)
	for _, v := range mappingRequestMonitoredResourcesSummarizedCountLicenseEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestMonitoredResourcesSummarizedCountLicenseEnumStringValues Enumerates the set of values in String for RequestMonitoredResourcesSummarizedCountLicenseEnum
func GetRequestMonitoredResourcesSummarizedCountLicenseEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
	}
}

// GetMappingRequestMonitoredResourcesSummarizedCountLicenseEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestMonitoredResourcesSummarizedCountLicenseEnum(val string) (RequestMonitoredResourcesSummarizedCountLicenseEnum, bool) {
	enum, ok := mappingRequestMonitoredResourcesSummarizedCountLicenseEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RequestMonitoredResourcesSummarizedCountSortByEnum Enum with underlying type: string
type RequestMonitoredResourcesSummarizedCountSortByEnum string

// Set of constants representing the allowable values for RequestMonitoredResourcesSummarizedCountSortByEnum
const (
	RequestMonitoredResourcesSummarizedCountSortByCount RequestMonitoredResourcesSummarizedCountSortByEnum = "count"
)

var mappingRequestMonitoredResourcesSummarizedCountSortByEnum = map[string]RequestMonitoredResourcesSummarizedCountSortByEnum{
	"count": RequestMonitoredResourcesSummarizedCountSortByCount,
}

var mappingRequestMonitoredResourcesSummarizedCountSortByEnumLowerCase = map[string]RequestMonitoredResourcesSummarizedCountSortByEnum{
	"count": RequestMonitoredResourcesSummarizedCountSortByCount,
}

// GetRequestMonitoredResourcesSummarizedCountSortByEnumValues Enumerates the set of values for RequestMonitoredResourcesSummarizedCountSortByEnum
func GetRequestMonitoredResourcesSummarizedCountSortByEnumValues() []RequestMonitoredResourcesSummarizedCountSortByEnum {
	values := make([]RequestMonitoredResourcesSummarizedCountSortByEnum, 0)
	for _, v := range mappingRequestMonitoredResourcesSummarizedCountSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestMonitoredResourcesSummarizedCountSortByEnumStringValues Enumerates the set of values in String for RequestMonitoredResourcesSummarizedCountSortByEnum
func GetRequestMonitoredResourcesSummarizedCountSortByEnumStringValues() []string {
	return []string{
		"count",
	}
}

// GetMappingRequestMonitoredResourcesSummarizedCountSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestMonitoredResourcesSummarizedCountSortByEnum(val string) (RequestMonitoredResourcesSummarizedCountSortByEnum, bool) {
	enum, ok := mappingRequestMonitoredResourcesSummarizedCountSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RequestMonitoredResourcesSummarizedCountSortOrderEnum Enum with underlying type: string
type RequestMonitoredResourcesSummarizedCountSortOrderEnum string

// Set of constants representing the allowable values for RequestMonitoredResourcesSummarizedCountSortOrderEnum
const (
	RequestMonitoredResourcesSummarizedCountSortOrderAsc  RequestMonitoredResourcesSummarizedCountSortOrderEnum = "ASC"
	RequestMonitoredResourcesSummarizedCountSortOrderDesc RequestMonitoredResourcesSummarizedCountSortOrderEnum = "DESC"
)

var mappingRequestMonitoredResourcesSummarizedCountSortOrderEnum = map[string]RequestMonitoredResourcesSummarizedCountSortOrderEnum{
	"ASC":  RequestMonitoredResourcesSummarizedCountSortOrderAsc,
	"DESC": RequestMonitoredResourcesSummarizedCountSortOrderDesc,
}

var mappingRequestMonitoredResourcesSummarizedCountSortOrderEnumLowerCase = map[string]RequestMonitoredResourcesSummarizedCountSortOrderEnum{
	"asc":  RequestMonitoredResourcesSummarizedCountSortOrderAsc,
	"desc": RequestMonitoredResourcesSummarizedCountSortOrderDesc,
}

// GetRequestMonitoredResourcesSummarizedCountSortOrderEnumValues Enumerates the set of values for RequestMonitoredResourcesSummarizedCountSortOrderEnum
func GetRequestMonitoredResourcesSummarizedCountSortOrderEnumValues() []RequestMonitoredResourcesSummarizedCountSortOrderEnum {
	values := make([]RequestMonitoredResourcesSummarizedCountSortOrderEnum, 0)
	for _, v := range mappingRequestMonitoredResourcesSummarizedCountSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestMonitoredResourcesSummarizedCountSortOrderEnumStringValues Enumerates the set of values in String for RequestMonitoredResourcesSummarizedCountSortOrderEnum
func GetRequestMonitoredResourcesSummarizedCountSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingRequestMonitoredResourcesSummarizedCountSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestMonitoredResourcesSummarizedCountSortOrderEnum(val string) (RequestMonitoredResourcesSummarizedCountSortOrderEnum, bool) {
	enum, ok := mappingRequestMonitoredResourcesSummarizedCountSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
