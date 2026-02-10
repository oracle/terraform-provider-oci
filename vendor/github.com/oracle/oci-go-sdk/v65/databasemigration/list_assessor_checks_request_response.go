// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAssessorChecksRequest wrapper for the ListAssessorChecks operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAssessorChecks.go.html to see an example of how to use ListAssessorChecksRequest.
type ListAssessorChecksRequest struct {

	// The OCID of the Assessment
	AssessmentId *string `mandatory:"true" contributesTo:"path" name:"assessmentId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the Assessor
	AssessorName *string `mandatory:"true" contributesTo:"path" name:"assessorName"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	// Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListAssessorChecksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAssessorChecksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssessorChecksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssessorChecksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssessorChecksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssessorChecksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAssessorChecksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAssessorChecksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAssessorChecksSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssessorChecksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAssessorChecksSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAssessorChecksResponse wrapper for the ListAssessorChecks operation
type ListAssessorChecksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssessorCheckSummaryCollection instances
	AssessorCheckSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAssessorChecksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssessorChecksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssessorChecksSortByEnum Enum with underlying type: string
type ListAssessorChecksSortByEnum string

// Set of constants representing the allowable values for ListAssessorChecksSortByEnum
const (
	ListAssessorChecksSortByTimecreated ListAssessorChecksSortByEnum = "timeCreated"
	ListAssessorChecksSortByDisplayname ListAssessorChecksSortByEnum = "displayName"
)

var mappingListAssessorChecksSortByEnum = map[string]ListAssessorChecksSortByEnum{
	"timeCreated": ListAssessorChecksSortByTimecreated,
	"displayName": ListAssessorChecksSortByDisplayname,
}

var mappingListAssessorChecksSortByEnumLowerCase = map[string]ListAssessorChecksSortByEnum{
	"timecreated": ListAssessorChecksSortByTimecreated,
	"displayname": ListAssessorChecksSortByDisplayname,
}

// GetListAssessorChecksSortByEnumValues Enumerates the set of values for ListAssessorChecksSortByEnum
func GetListAssessorChecksSortByEnumValues() []ListAssessorChecksSortByEnum {
	values := make([]ListAssessorChecksSortByEnum, 0)
	for _, v := range mappingListAssessorChecksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssessorChecksSortByEnumStringValues Enumerates the set of values in String for ListAssessorChecksSortByEnum
func GetListAssessorChecksSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAssessorChecksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssessorChecksSortByEnum(val string) (ListAssessorChecksSortByEnum, bool) {
	enum, ok := mappingListAssessorChecksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssessorChecksSortOrderEnum Enum with underlying type: string
type ListAssessorChecksSortOrderEnum string

// Set of constants representing the allowable values for ListAssessorChecksSortOrderEnum
const (
	ListAssessorChecksSortOrderAsc  ListAssessorChecksSortOrderEnum = "ASC"
	ListAssessorChecksSortOrderDesc ListAssessorChecksSortOrderEnum = "DESC"
)

var mappingListAssessorChecksSortOrderEnum = map[string]ListAssessorChecksSortOrderEnum{
	"ASC":  ListAssessorChecksSortOrderAsc,
	"DESC": ListAssessorChecksSortOrderDesc,
}

var mappingListAssessorChecksSortOrderEnumLowerCase = map[string]ListAssessorChecksSortOrderEnum{
	"asc":  ListAssessorChecksSortOrderAsc,
	"desc": ListAssessorChecksSortOrderDesc,
}

// GetListAssessorChecksSortOrderEnumValues Enumerates the set of values for ListAssessorChecksSortOrderEnum
func GetListAssessorChecksSortOrderEnumValues() []ListAssessorChecksSortOrderEnum {
	values := make([]ListAssessorChecksSortOrderEnum, 0)
	for _, v := range mappingListAssessorChecksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssessorChecksSortOrderEnumStringValues Enumerates the set of values in String for ListAssessorChecksSortOrderEnum
func GetListAssessorChecksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAssessorChecksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssessorChecksSortOrderEnum(val string) (ListAssessorChecksSortOrderEnum, bool) {
	enum, ok := mappingListAssessorChecksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
