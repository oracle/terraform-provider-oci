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

// ListAssessmentObjectTypesRequest wrapper for the ListAssessmentObjectTypes operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAssessmentObjectTypes.go.html to see an example of how to use ListAssessmentObjectTypesRequest.
type ListAssessmentObjectTypesRequest struct {

	// The connection type for assessment objects.
	ConnectionType ListAssessmentObjectTypesConnectionTypeEnum `mandatory:"true" contributesTo:"query" name:"connectionType" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for name is custom based on it's usage frequency. If no value is specified name is default.
	SortBy ListAssessmentObjectTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAssessmentObjectTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssessmentObjectTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssessmentObjectTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssessmentObjectTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssessmentObjectTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAssessmentObjectTypesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAssessmentObjectTypesConnectionTypeEnum(string(request.ConnectionType)); !ok && request.ConnectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionType: %s. Supported values are: %s.", request.ConnectionType, strings.Join(GetListAssessmentObjectTypesConnectionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssessmentObjectTypesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAssessmentObjectTypesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssessmentObjectTypesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAssessmentObjectTypesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAssessmentObjectTypesResponse wrapper for the ListAssessmentObjectTypes operation
type ListAssessmentObjectTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssessmentObjectTypeSummaryCollection instances
	AssessmentObjectTypeSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAssessmentObjectTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssessmentObjectTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssessmentObjectTypesConnectionTypeEnum Enum with underlying type: string
type ListAssessmentObjectTypesConnectionTypeEnum string

// Set of constants representing the allowable values for ListAssessmentObjectTypesConnectionTypeEnum
const (
	ListAssessmentObjectTypesConnectionTypeMysql  ListAssessmentObjectTypesConnectionTypeEnum = "MYSQL"
	ListAssessmentObjectTypesConnectionTypeOracle ListAssessmentObjectTypesConnectionTypeEnum = "ORACLE"
)

var mappingListAssessmentObjectTypesConnectionTypeEnum = map[string]ListAssessmentObjectTypesConnectionTypeEnum{
	"MYSQL":  ListAssessmentObjectTypesConnectionTypeMysql,
	"ORACLE": ListAssessmentObjectTypesConnectionTypeOracle,
}

var mappingListAssessmentObjectTypesConnectionTypeEnumLowerCase = map[string]ListAssessmentObjectTypesConnectionTypeEnum{
	"mysql":  ListAssessmentObjectTypesConnectionTypeMysql,
	"oracle": ListAssessmentObjectTypesConnectionTypeOracle,
}

// GetListAssessmentObjectTypesConnectionTypeEnumValues Enumerates the set of values for ListAssessmentObjectTypesConnectionTypeEnum
func GetListAssessmentObjectTypesConnectionTypeEnumValues() []ListAssessmentObjectTypesConnectionTypeEnum {
	values := make([]ListAssessmentObjectTypesConnectionTypeEnum, 0)
	for _, v := range mappingListAssessmentObjectTypesConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssessmentObjectTypesConnectionTypeEnumStringValues Enumerates the set of values in String for ListAssessmentObjectTypesConnectionTypeEnum
func GetListAssessmentObjectTypesConnectionTypeEnumStringValues() []string {
	return []string{
		"MYSQL",
		"ORACLE",
	}
}

// GetMappingListAssessmentObjectTypesConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssessmentObjectTypesConnectionTypeEnum(val string) (ListAssessmentObjectTypesConnectionTypeEnum, bool) {
	enum, ok := mappingListAssessmentObjectTypesConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssessmentObjectTypesSortByEnum Enum with underlying type: string
type ListAssessmentObjectTypesSortByEnum string

// Set of constants representing the allowable values for ListAssessmentObjectTypesSortByEnum
const (
	ListAssessmentObjectTypesSortByName ListAssessmentObjectTypesSortByEnum = "name"
)

var mappingListAssessmentObjectTypesSortByEnum = map[string]ListAssessmentObjectTypesSortByEnum{
	"name": ListAssessmentObjectTypesSortByName,
}

var mappingListAssessmentObjectTypesSortByEnumLowerCase = map[string]ListAssessmentObjectTypesSortByEnum{
	"name": ListAssessmentObjectTypesSortByName,
}

// GetListAssessmentObjectTypesSortByEnumValues Enumerates the set of values for ListAssessmentObjectTypesSortByEnum
func GetListAssessmentObjectTypesSortByEnumValues() []ListAssessmentObjectTypesSortByEnum {
	values := make([]ListAssessmentObjectTypesSortByEnum, 0)
	for _, v := range mappingListAssessmentObjectTypesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssessmentObjectTypesSortByEnumStringValues Enumerates the set of values in String for ListAssessmentObjectTypesSortByEnum
func GetListAssessmentObjectTypesSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListAssessmentObjectTypesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssessmentObjectTypesSortByEnum(val string) (ListAssessmentObjectTypesSortByEnum, bool) {
	enum, ok := mappingListAssessmentObjectTypesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssessmentObjectTypesSortOrderEnum Enum with underlying type: string
type ListAssessmentObjectTypesSortOrderEnum string

// Set of constants representing the allowable values for ListAssessmentObjectTypesSortOrderEnum
const (
	ListAssessmentObjectTypesSortOrderAsc  ListAssessmentObjectTypesSortOrderEnum = "ASC"
	ListAssessmentObjectTypesSortOrderDesc ListAssessmentObjectTypesSortOrderEnum = "DESC"
)

var mappingListAssessmentObjectTypesSortOrderEnum = map[string]ListAssessmentObjectTypesSortOrderEnum{
	"ASC":  ListAssessmentObjectTypesSortOrderAsc,
	"DESC": ListAssessmentObjectTypesSortOrderDesc,
}

var mappingListAssessmentObjectTypesSortOrderEnumLowerCase = map[string]ListAssessmentObjectTypesSortOrderEnum{
	"asc":  ListAssessmentObjectTypesSortOrderAsc,
	"desc": ListAssessmentObjectTypesSortOrderDesc,
}

// GetListAssessmentObjectTypesSortOrderEnumValues Enumerates the set of values for ListAssessmentObjectTypesSortOrderEnum
func GetListAssessmentObjectTypesSortOrderEnumValues() []ListAssessmentObjectTypesSortOrderEnum {
	values := make([]ListAssessmentObjectTypesSortOrderEnum, 0)
	for _, v := range mappingListAssessmentObjectTypesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssessmentObjectTypesSortOrderEnumStringValues Enumerates the set of values in String for ListAssessmentObjectTypesSortOrderEnum
func GetListAssessmentObjectTypesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAssessmentObjectTypesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssessmentObjectTypesSortOrderEnum(val string) (ListAssessmentObjectTypesSortOrderEnum, bool) {
	enum, ok := mappingListAssessmentObjectTypesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
