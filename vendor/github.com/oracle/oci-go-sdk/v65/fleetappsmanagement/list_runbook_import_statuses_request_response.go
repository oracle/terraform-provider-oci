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

// ListRunbookImportStatusesRequest wrapper for the ListRunbookImportStatuses operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListRunbookImportStatuses.go.html to see an example of how to use ListRunbookImportStatusesRequest.
type ListRunbookImportStatusesRequest struct {

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
	SortOrder ListRunbookImportStatusesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListRunbookImportStatusesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRunbookImportStatusesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRunbookImportStatusesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRunbookImportStatusesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRunbookImportStatusesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRunbookImportStatusesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRunbookImportStatusesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRunbookImportStatusesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRunbookImportStatusesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRunbookImportStatusesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRunbookImportStatusesResponse wrapper for the ListRunbookImportStatuses operation
type ListRunbookImportStatusesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RunbookImportStatusCollection instances
	RunbookImportStatusCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRunbookImportStatusesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRunbookImportStatusesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRunbookImportStatusesSortOrderEnum Enum with underlying type: string
type ListRunbookImportStatusesSortOrderEnum string

// Set of constants representing the allowable values for ListRunbookImportStatusesSortOrderEnum
const (
	ListRunbookImportStatusesSortOrderAsc  ListRunbookImportStatusesSortOrderEnum = "ASC"
	ListRunbookImportStatusesSortOrderDesc ListRunbookImportStatusesSortOrderEnum = "DESC"
)

var mappingListRunbookImportStatusesSortOrderEnum = map[string]ListRunbookImportStatusesSortOrderEnum{
	"ASC":  ListRunbookImportStatusesSortOrderAsc,
	"DESC": ListRunbookImportStatusesSortOrderDesc,
}

var mappingListRunbookImportStatusesSortOrderEnumLowerCase = map[string]ListRunbookImportStatusesSortOrderEnum{
	"asc":  ListRunbookImportStatusesSortOrderAsc,
	"desc": ListRunbookImportStatusesSortOrderDesc,
}

// GetListRunbookImportStatusesSortOrderEnumValues Enumerates the set of values for ListRunbookImportStatusesSortOrderEnum
func GetListRunbookImportStatusesSortOrderEnumValues() []ListRunbookImportStatusesSortOrderEnum {
	values := make([]ListRunbookImportStatusesSortOrderEnum, 0)
	for _, v := range mappingListRunbookImportStatusesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRunbookImportStatusesSortOrderEnumStringValues Enumerates the set of values in String for ListRunbookImportStatusesSortOrderEnum
func GetListRunbookImportStatusesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRunbookImportStatusesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRunbookImportStatusesSortOrderEnum(val string) (ListRunbookImportStatusesSortOrderEnum, bool) {
	enum, ok := mappingListRunbookImportStatusesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRunbookImportStatusesSortByEnum Enum with underlying type: string
type ListRunbookImportStatusesSortByEnum string

// Set of constants representing the allowable values for ListRunbookImportStatusesSortByEnum
const (
	ListRunbookImportStatusesSortByTimecreated ListRunbookImportStatusesSortByEnum = "timeCreated"
	ListRunbookImportStatusesSortByDisplayname ListRunbookImportStatusesSortByEnum = "displayName"
)

var mappingListRunbookImportStatusesSortByEnum = map[string]ListRunbookImportStatusesSortByEnum{
	"timeCreated": ListRunbookImportStatusesSortByTimecreated,
	"displayName": ListRunbookImportStatusesSortByDisplayname,
}

var mappingListRunbookImportStatusesSortByEnumLowerCase = map[string]ListRunbookImportStatusesSortByEnum{
	"timecreated": ListRunbookImportStatusesSortByTimecreated,
	"displayname": ListRunbookImportStatusesSortByDisplayname,
}

// GetListRunbookImportStatusesSortByEnumValues Enumerates the set of values for ListRunbookImportStatusesSortByEnum
func GetListRunbookImportStatusesSortByEnumValues() []ListRunbookImportStatusesSortByEnum {
	values := make([]ListRunbookImportStatusesSortByEnum, 0)
	for _, v := range mappingListRunbookImportStatusesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRunbookImportStatusesSortByEnumStringValues Enumerates the set of values in String for ListRunbookImportStatusesSortByEnum
func GetListRunbookImportStatusesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListRunbookImportStatusesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRunbookImportStatusesSortByEnum(val string) (ListRunbookImportStatusesSortByEnum, bool) {
	enum, ok := mappingListRunbookImportStatusesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
