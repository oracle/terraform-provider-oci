// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRunbookExportStatusesRequest wrapper for the ListRunbookExportStatuses operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListRunbookExportStatuses.go.html to see an example of how to use ListRunbookExportStatusesRequest.
type ListRunbookExportStatusesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier or OCID for listing a single Runbook by id.
	// Either compartmentId or id must be provided.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListRunbookExportStatusesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListRunbookExportStatusesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRunbookExportStatusesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRunbookExportStatusesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRunbookExportStatusesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRunbookExportStatusesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRunbookExportStatusesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRunbookExportStatusesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRunbookExportStatusesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRunbookExportStatusesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRunbookExportStatusesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRunbookExportStatusesResponse wrapper for the ListRunbookExportStatuses operation
type ListRunbookExportStatusesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RunbookExportStatusCollection instances
	RunbookExportStatusCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRunbookExportStatusesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRunbookExportStatusesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRunbookExportStatusesSortOrderEnum Enum with underlying type: string
type ListRunbookExportStatusesSortOrderEnum string

// Set of constants representing the allowable values for ListRunbookExportStatusesSortOrderEnum
const (
	ListRunbookExportStatusesSortOrderAsc  ListRunbookExportStatusesSortOrderEnum = "ASC"
	ListRunbookExportStatusesSortOrderDesc ListRunbookExportStatusesSortOrderEnum = "DESC"
)

var mappingListRunbookExportStatusesSortOrderEnum = map[string]ListRunbookExportStatusesSortOrderEnum{
	"ASC":  ListRunbookExportStatusesSortOrderAsc,
	"DESC": ListRunbookExportStatusesSortOrderDesc,
}

var mappingListRunbookExportStatusesSortOrderEnumLowerCase = map[string]ListRunbookExportStatusesSortOrderEnum{
	"asc":  ListRunbookExportStatusesSortOrderAsc,
	"desc": ListRunbookExportStatusesSortOrderDesc,
}

// GetListRunbookExportStatusesSortOrderEnumValues Enumerates the set of values for ListRunbookExportStatusesSortOrderEnum
func GetListRunbookExportStatusesSortOrderEnumValues() []ListRunbookExportStatusesSortOrderEnum {
	values := make([]ListRunbookExportStatusesSortOrderEnum, 0)
	for _, v := range mappingListRunbookExportStatusesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRunbookExportStatusesSortOrderEnumStringValues Enumerates the set of values in String for ListRunbookExportStatusesSortOrderEnum
func GetListRunbookExportStatusesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRunbookExportStatusesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRunbookExportStatusesSortOrderEnum(val string) (ListRunbookExportStatusesSortOrderEnum, bool) {
	enum, ok := mappingListRunbookExportStatusesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRunbookExportStatusesSortByEnum Enum with underlying type: string
type ListRunbookExportStatusesSortByEnum string

// Set of constants representing the allowable values for ListRunbookExportStatusesSortByEnum
const (
	ListRunbookExportStatusesSortByTimecreated ListRunbookExportStatusesSortByEnum = "timeCreated"
	ListRunbookExportStatusesSortByDisplayname ListRunbookExportStatusesSortByEnum = "displayName"
)

var mappingListRunbookExportStatusesSortByEnum = map[string]ListRunbookExportStatusesSortByEnum{
	"timeCreated": ListRunbookExportStatusesSortByTimecreated,
	"displayName": ListRunbookExportStatusesSortByDisplayname,
}

var mappingListRunbookExportStatusesSortByEnumLowerCase = map[string]ListRunbookExportStatusesSortByEnum{
	"timecreated": ListRunbookExportStatusesSortByTimecreated,
	"displayname": ListRunbookExportStatusesSortByDisplayname,
}

// GetListRunbookExportStatusesSortByEnumValues Enumerates the set of values for ListRunbookExportStatusesSortByEnum
func GetListRunbookExportStatusesSortByEnumValues() []ListRunbookExportStatusesSortByEnum {
	values := make([]ListRunbookExportStatusesSortByEnum, 0)
	for _, v := range mappingListRunbookExportStatusesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRunbookExportStatusesSortByEnumStringValues Enumerates the set of values in String for ListRunbookExportStatusesSortByEnum
func GetListRunbookExportStatusesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListRunbookExportStatusesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRunbookExportStatusesSortByEnum(val string) (ListRunbookExportStatusesSortByEnum, bool) {
	enum, ok := mappingListRunbookExportStatusesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
