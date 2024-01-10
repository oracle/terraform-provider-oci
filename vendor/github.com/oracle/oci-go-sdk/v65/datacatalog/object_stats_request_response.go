// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ObjectStatsRequest wrapper for the ObjectStats operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ObjectStats.go.html to see an example of how to use ObjectStatsRequest.
type ObjectStatsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ObjectStatsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ObjectStatsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ObjectStatsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ObjectStatsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ObjectStatsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ObjectStatsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ObjectStatsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingObjectStatsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetObjectStatsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingObjectStatsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetObjectStatsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ObjectStatsResponse wrapper for the ObjectStats operation
type ObjectStatsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of string instances
	Value *string `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ObjectStatsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ObjectStatsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ObjectStatsSortByEnum Enum with underlying type: string
type ObjectStatsSortByEnum string

// Set of constants representing the allowable values for ObjectStatsSortByEnum
const (
	ObjectStatsSortByTimecreated ObjectStatsSortByEnum = "TIMECREATED"
	ObjectStatsSortByDisplayname ObjectStatsSortByEnum = "DISPLAYNAME"
)

var mappingObjectStatsSortByEnum = map[string]ObjectStatsSortByEnum{
	"TIMECREATED": ObjectStatsSortByTimecreated,
	"DISPLAYNAME": ObjectStatsSortByDisplayname,
}

var mappingObjectStatsSortByEnumLowerCase = map[string]ObjectStatsSortByEnum{
	"timecreated": ObjectStatsSortByTimecreated,
	"displayname": ObjectStatsSortByDisplayname,
}

// GetObjectStatsSortByEnumValues Enumerates the set of values for ObjectStatsSortByEnum
func GetObjectStatsSortByEnumValues() []ObjectStatsSortByEnum {
	values := make([]ObjectStatsSortByEnum, 0)
	for _, v := range mappingObjectStatsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetObjectStatsSortByEnumStringValues Enumerates the set of values in String for ObjectStatsSortByEnum
func GetObjectStatsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingObjectStatsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingObjectStatsSortByEnum(val string) (ObjectStatsSortByEnum, bool) {
	enum, ok := mappingObjectStatsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ObjectStatsSortOrderEnum Enum with underlying type: string
type ObjectStatsSortOrderEnum string

// Set of constants representing the allowable values for ObjectStatsSortOrderEnum
const (
	ObjectStatsSortOrderAsc  ObjectStatsSortOrderEnum = "ASC"
	ObjectStatsSortOrderDesc ObjectStatsSortOrderEnum = "DESC"
)

var mappingObjectStatsSortOrderEnum = map[string]ObjectStatsSortOrderEnum{
	"ASC":  ObjectStatsSortOrderAsc,
	"DESC": ObjectStatsSortOrderDesc,
}

var mappingObjectStatsSortOrderEnumLowerCase = map[string]ObjectStatsSortOrderEnum{
	"asc":  ObjectStatsSortOrderAsc,
	"desc": ObjectStatsSortOrderDesc,
}

// GetObjectStatsSortOrderEnumValues Enumerates the set of values for ObjectStatsSortOrderEnum
func GetObjectStatsSortOrderEnumValues() []ObjectStatsSortOrderEnum {
	values := make([]ObjectStatsSortOrderEnum, 0)
	for _, v := range mappingObjectStatsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetObjectStatsSortOrderEnumStringValues Enumerates the set of values in String for ObjectStatsSortOrderEnum
func GetObjectStatsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingObjectStatsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingObjectStatsSortOrderEnum(val string) (ObjectStatsSortOrderEnum, bool) {
	enum, ok := mappingObjectStatsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
