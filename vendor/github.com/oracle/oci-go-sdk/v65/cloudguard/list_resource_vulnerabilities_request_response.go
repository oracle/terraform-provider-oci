// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListResourceVulnerabilitiesRequest wrapper for the ListResourceVulnerabilities operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResourceVulnerabilities.go.html to see an example of how to use ListResourceVulnerabilitiesRequest.
type ListResourceVulnerabilitiesRequest struct {

	// CloudGuard resource OCID
	ResourceId *string `mandatory:"true" contributesTo:"path" name:"resourceId"`

	// CVE ID associated with the resource.
	CveId *string `mandatory:"false" contributesTo:"query" name:"cveId"`

	// Risk level of the problem.
	RiskLevel *string `mandatory:"false" contributesTo:"query" name:"riskLevel"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListResourceVulnerabilitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListResourceVulnerabilitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourceVulnerabilitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourceVulnerabilitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourceVulnerabilitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourceVulnerabilitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResourceVulnerabilitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResourceVulnerabilitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResourceVulnerabilitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceVulnerabilitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResourceVulnerabilitiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResourceVulnerabilitiesResponse wrapper for the ListResourceVulnerabilities operation
type ListResourceVulnerabilitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourceVulnerabilityCollection instances
	ResourceVulnerabilityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResourceVulnerabilitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourceVulnerabilitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourceVulnerabilitiesSortOrderEnum Enum with underlying type: string
type ListResourceVulnerabilitiesSortOrderEnum string

// Set of constants representing the allowable values for ListResourceVulnerabilitiesSortOrderEnum
const (
	ListResourceVulnerabilitiesSortOrderAsc  ListResourceVulnerabilitiesSortOrderEnum = "ASC"
	ListResourceVulnerabilitiesSortOrderDesc ListResourceVulnerabilitiesSortOrderEnum = "DESC"
)

var mappingListResourceVulnerabilitiesSortOrderEnum = map[string]ListResourceVulnerabilitiesSortOrderEnum{
	"ASC":  ListResourceVulnerabilitiesSortOrderAsc,
	"DESC": ListResourceVulnerabilitiesSortOrderDesc,
}

var mappingListResourceVulnerabilitiesSortOrderEnumLowerCase = map[string]ListResourceVulnerabilitiesSortOrderEnum{
	"asc":  ListResourceVulnerabilitiesSortOrderAsc,
	"desc": ListResourceVulnerabilitiesSortOrderDesc,
}

// GetListResourceVulnerabilitiesSortOrderEnumValues Enumerates the set of values for ListResourceVulnerabilitiesSortOrderEnum
func GetListResourceVulnerabilitiesSortOrderEnumValues() []ListResourceVulnerabilitiesSortOrderEnum {
	values := make([]ListResourceVulnerabilitiesSortOrderEnum, 0)
	for _, v := range mappingListResourceVulnerabilitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceVulnerabilitiesSortOrderEnumStringValues Enumerates the set of values in String for ListResourceVulnerabilitiesSortOrderEnum
func GetListResourceVulnerabilitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResourceVulnerabilitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceVulnerabilitiesSortOrderEnum(val string) (ListResourceVulnerabilitiesSortOrderEnum, bool) {
	enum, ok := mappingListResourceVulnerabilitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceVulnerabilitiesSortByEnum Enum with underlying type: string
type ListResourceVulnerabilitiesSortByEnum string

// Set of constants representing the allowable values for ListResourceVulnerabilitiesSortByEnum
const (
	ListResourceVulnerabilitiesSortByTimecreated ListResourceVulnerabilitiesSortByEnum = "timeCreated"
	ListResourceVulnerabilitiesSortByDisplayname ListResourceVulnerabilitiesSortByEnum = "displayName"
)

var mappingListResourceVulnerabilitiesSortByEnum = map[string]ListResourceVulnerabilitiesSortByEnum{
	"timeCreated": ListResourceVulnerabilitiesSortByTimecreated,
	"displayName": ListResourceVulnerabilitiesSortByDisplayname,
}

var mappingListResourceVulnerabilitiesSortByEnumLowerCase = map[string]ListResourceVulnerabilitiesSortByEnum{
	"timecreated": ListResourceVulnerabilitiesSortByTimecreated,
	"displayname": ListResourceVulnerabilitiesSortByDisplayname,
}

// GetListResourceVulnerabilitiesSortByEnumValues Enumerates the set of values for ListResourceVulnerabilitiesSortByEnum
func GetListResourceVulnerabilitiesSortByEnumValues() []ListResourceVulnerabilitiesSortByEnum {
	values := make([]ListResourceVulnerabilitiesSortByEnum, 0)
	for _, v := range mappingListResourceVulnerabilitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceVulnerabilitiesSortByEnumStringValues Enumerates the set of values in String for ListResourceVulnerabilitiesSortByEnum
func GetListResourceVulnerabilitiesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListResourceVulnerabilitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceVulnerabilitiesSortByEnum(val string) (ListResourceVulnerabilitiesSortByEnum, bool) {
	enum, ok := mappingListResourceVulnerabilitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
