// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourcemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetWorkRequestLogEntriesRequest wrapper for the GetWorkRequestLogEntries operation
type GetWorkRequestLogEntriesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the work request.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter that returns only logs of a specified type.
	Type []LogEntryTypeEnum `contributesTo:"query" name:"type" omitEmpty:"true" collectionFormat:"multi"`

	// A filter that returns only log entries that match a given severity level or greater.
	LevelGreaterThanOrEqualTo LogEntryLevelEnum `mandatory:"false" contributesTo:"query" name:"levelGreaterThanOrEqualTo" omitEmpty:"true"`

	// The sort order to use when sorting returned resources. Ascending (`ASC`) or descending (`DESC`).
	SortOrder GetWorkRequestLogEntriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The number of items returned in a paginated `List` call. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the preceding `List` call.
	// For information about pagination, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Time stamp specifying the lower time limit for which logs are returned in a query.
	// Format is defined by RFC3339.
	// Example: `2020-01-01T12:00:00.000Z`
	TimestampGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timestampGreaterThanOrEqualTo"`

	// Time stamp specifying the upper time limit for which logs are returned in a query.
	// Format is defined by RFC3339.
	// Example: `2020-02-01T12:00:00.000Z`
	TimestampLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timestampLessThanOrEqualTo"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetWorkRequestLogEntriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetWorkRequestLogEntriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetWorkRequestLogEntriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetWorkRequestLogEntriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetWorkRequestLogEntriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Type {
		if _, ok := GetMappingLogEntryTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", val, strings.Join(GetLogEntryTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingLogEntryLevelEnum(string(request.LevelGreaterThanOrEqualTo)); !ok && request.LevelGreaterThanOrEqualTo != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LevelGreaterThanOrEqualTo: %s. Supported values are: %s.", request.LevelGreaterThanOrEqualTo, strings.Join(GetLogEntryLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetWorkRequestLogEntriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetWorkRequestLogEntriesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetWorkRequestLogEntriesResponse wrapper for the GetWorkRequestLogEntries operation
type GetWorkRequestLogEntriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogEntryCollection instances
	LogEntryCollection `presentIn:"body"`

	// Unique identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of paginated list items. If the `opc-next-page`
	// header appears in the response, additional pages of results remain.
	// To receive the next page, include the header value in the `page` param.
	// If the `opc-next-page` header does not appear in the response, there
	// are no more list items to get. For more information about list pagination,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response GetWorkRequestLogEntriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetWorkRequestLogEntriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetWorkRequestLogEntriesSortOrderEnum Enum with underlying type: string
type GetWorkRequestLogEntriesSortOrderEnum string

// Set of constants representing the allowable values for GetWorkRequestLogEntriesSortOrderEnum
const (
	GetWorkRequestLogEntriesSortOrderAsc  GetWorkRequestLogEntriesSortOrderEnum = "ASC"
	GetWorkRequestLogEntriesSortOrderDesc GetWorkRequestLogEntriesSortOrderEnum = "DESC"
)

var mappingGetWorkRequestLogEntriesSortOrderEnum = map[string]GetWorkRequestLogEntriesSortOrderEnum{
	"ASC":  GetWorkRequestLogEntriesSortOrderAsc,
	"DESC": GetWorkRequestLogEntriesSortOrderDesc,
}

var mappingGetWorkRequestLogEntriesSortOrderEnumLowerCase = map[string]GetWorkRequestLogEntriesSortOrderEnum{
	"asc":  GetWorkRequestLogEntriesSortOrderAsc,
	"desc": GetWorkRequestLogEntriesSortOrderDesc,
}

// GetGetWorkRequestLogEntriesSortOrderEnumValues Enumerates the set of values for GetWorkRequestLogEntriesSortOrderEnum
func GetGetWorkRequestLogEntriesSortOrderEnumValues() []GetWorkRequestLogEntriesSortOrderEnum {
	values := make([]GetWorkRequestLogEntriesSortOrderEnum, 0)
	for _, v := range mappingGetWorkRequestLogEntriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetWorkRequestLogEntriesSortOrderEnumStringValues Enumerates the set of values in String for GetWorkRequestLogEntriesSortOrderEnum
func GetGetWorkRequestLogEntriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetWorkRequestLogEntriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetWorkRequestLogEntriesSortOrderEnum(val string) (GetWorkRequestLogEntriesSortOrderEnum, bool) {
	enum, ok := mappingGetWorkRequestLogEntriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
