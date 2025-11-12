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

// ListReportMetadataRequest wrapper for the ListReportMetadata operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListReportMetadata.go.html to see an example of how to use ListReportMetadataRequest.
type ListReportMetadataRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return data for given report name.
	ReportName *string `mandatory:"false" contributesTo:"query" name:"reportName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListReportMetadataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListReportMetadataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReportMetadataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReportMetadataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListReportMetadataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReportMetadataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListReportMetadataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListReportMetadataSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListReportMetadataSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReportMetadataSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListReportMetadataSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListReportMetadataResponse wrapper for the ListReportMetadata operation
type ListReportMetadataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ReportMetadataCollection instances
	ReportMetadataCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListReportMetadataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReportMetadataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListReportMetadataSortOrderEnum Enum with underlying type: string
type ListReportMetadataSortOrderEnum string

// Set of constants representing the allowable values for ListReportMetadataSortOrderEnum
const (
	ListReportMetadataSortOrderAsc  ListReportMetadataSortOrderEnum = "ASC"
	ListReportMetadataSortOrderDesc ListReportMetadataSortOrderEnum = "DESC"
)

var mappingListReportMetadataSortOrderEnum = map[string]ListReportMetadataSortOrderEnum{
	"ASC":  ListReportMetadataSortOrderAsc,
	"DESC": ListReportMetadataSortOrderDesc,
}

var mappingListReportMetadataSortOrderEnumLowerCase = map[string]ListReportMetadataSortOrderEnum{
	"asc":  ListReportMetadataSortOrderAsc,
	"desc": ListReportMetadataSortOrderDesc,
}

// GetListReportMetadataSortOrderEnumValues Enumerates the set of values for ListReportMetadataSortOrderEnum
func GetListReportMetadataSortOrderEnumValues() []ListReportMetadataSortOrderEnum {
	values := make([]ListReportMetadataSortOrderEnum, 0)
	for _, v := range mappingListReportMetadataSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportMetadataSortOrderEnumStringValues Enumerates the set of values in String for ListReportMetadataSortOrderEnum
func GetListReportMetadataSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListReportMetadataSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportMetadataSortOrderEnum(val string) (ListReportMetadataSortOrderEnum, bool) {
	enum, ok := mappingListReportMetadataSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReportMetadataSortByEnum Enum with underlying type: string
type ListReportMetadataSortByEnum string

// Set of constants representing the allowable values for ListReportMetadataSortByEnum
const (
	ListReportMetadataSortByTimecreated ListReportMetadataSortByEnum = "timeCreated"
	ListReportMetadataSortByDisplayname ListReportMetadataSortByEnum = "displayName"
)

var mappingListReportMetadataSortByEnum = map[string]ListReportMetadataSortByEnum{
	"timeCreated": ListReportMetadataSortByTimecreated,
	"displayName": ListReportMetadataSortByDisplayname,
}

var mappingListReportMetadataSortByEnumLowerCase = map[string]ListReportMetadataSortByEnum{
	"timecreated": ListReportMetadataSortByTimecreated,
	"displayname": ListReportMetadataSortByDisplayname,
}

// GetListReportMetadataSortByEnumValues Enumerates the set of values for ListReportMetadataSortByEnum
func GetListReportMetadataSortByEnumValues() []ListReportMetadataSortByEnum {
	values := make([]ListReportMetadataSortByEnum, 0)
	for _, v := range mappingListReportMetadataSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportMetadataSortByEnumStringValues Enumerates the set of values in String for ListReportMetadataSortByEnum
func GetListReportMetadataSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListReportMetadataSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportMetadataSortByEnum(val string) (ListReportMetadataSortByEnum, bool) {
	enum, ok := mappingListReportMetadataSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
