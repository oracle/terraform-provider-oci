// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListParameterFileVersionsRequest wrapper for the ListParameterFileVersions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListParameterFileVersions.go.html to see an example of how to use ListParameterFileVersionsRequest.
type ListParameterFileVersionsRequest struct {

	// The OCID of the job
	JobId *string `mandatory:"true" contributesTo:"path" name:"jobId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	// Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListParameterFileVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListParameterFileVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListParameterFileVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListParameterFileVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListParameterFileVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListParameterFileVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListParameterFileVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListParameterFileVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListParameterFileVersionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListParameterFileVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListParameterFileVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListParameterFileVersionsResponse wrapper for the ListParameterFileVersions operation
type ListParameterFileVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ParameterFileVersionCollection instances
	ParameterFileVersionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListParameterFileVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListParameterFileVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListParameterFileVersionsSortByEnum Enum with underlying type: string
type ListParameterFileVersionsSortByEnum string

// Set of constants representing the allowable values for ListParameterFileVersionsSortByEnum
const (
	ListParameterFileVersionsSortByTimecreated ListParameterFileVersionsSortByEnum = "timeCreated"
	ListParameterFileVersionsSortByDisplayname ListParameterFileVersionsSortByEnum = "displayName"
)

var mappingListParameterFileVersionsSortByEnum = map[string]ListParameterFileVersionsSortByEnum{
	"timeCreated": ListParameterFileVersionsSortByTimecreated,
	"displayName": ListParameterFileVersionsSortByDisplayname,
}

var mappingListParameterFileVersionsSortByEnumLowerCase = map[string]ListParameterFileVersionsSortByEnum{
	"timecreated": ListParameterFileVersionsSortByTimecreated,
	"displayname": ListParameterFileVersionsSortByDisplayname,
}

// GetListParameterFileVersionsSortByEnumValues Enumerates the set of values for ListParameterFileVersionsSortByEnum
func GetListParameterFileVersionsSortByEnumValues() []ListParameterFileVersionsSortByEnum {
	values := make([]ListParameterFileVersionsSortByEnum, 0)
	for _, v := range mappingListParameterFileVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListParameterFileVersionsSortByEnumStringValues Enumerates the set of values in String for ListParameterFileVersionsSortByEnum
func GetListParameterFileVersionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListParameterFileVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListParameterFileVersionsSortByEnum(val string) (ListParameterFileVersionsSortByEnum, bool) {
	enum, ok := mappingListParameterFileVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListParameterFileVersionsSortOrderEnum Enum with underlying type: string
type ListParameterFileVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListParameterFileVersionsSortOrderEnum
const (
	ListParameterFileVersionsSortOrderAsc  ListParameterFileVersionsSortOrderEnum = "ASC"
	ListParameterFileVersionsSortOrderDesc ListParameterFileVersionsSortOrderEnum = "DESC"
)

var mappingListParameterFileVersionsSortOrderEnum = map[string]ListParameterFileVersionsSortOrderEnum{
	"ASC":  ListParameterFileVersionsSortOrderAsc,
	"DESC": ListParameterFileVersionsSortOrderDesc,
}

var mappingListParameterFileVersionsSortOrderEnumLowerCase = map[string]ListParameterFileVersionsSortOrderEnum{
	"asc":  ListParameterFileVersionsSortOrderAsc,
	"desc": ListParameterFileVersionsSortOrderDesc,
}

// GetListParameterFileVersionsSortOrderEnumValues Enumerates the set of values for ListParameterFileVersionsSortOrderEnum
func GetListParameterFileVersionsSortOrderEnumValues() []ListParameterFileVersionsSortOrderEnum {
	values := make([]ListParameterFileVersionsSortOrderEnum, 0)
	for _, v := range mappingListParameterFileVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListParameterFileVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListParameterFileVersionsSortOrderEnum
func GetListParameterFileVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListParameterFileVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListParameterFileVersionsSortOrderEnum(val string) (ListParameterFileVersionsSortOrderEnum, bool) {
	enum, ok := mappingListParameterFileVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
