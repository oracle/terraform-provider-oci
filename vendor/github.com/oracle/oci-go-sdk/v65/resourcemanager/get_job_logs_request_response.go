// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourcemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetJobLogsRequest wrapper for the GetJobLogs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetJobLogs.go.html to see an example of how to use GetJobLogsRequest.
type GetJobLogsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
	JobId *string `mandatory:"true" contributesTo:"path" name:"jobId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter that returns only logs of a specified type.
	Type []LogEntryTypeEnum `contributesTo:"query" name:"type" omitEmpty:"true" collectionFormat:"multi"`

	// A filter that returns only log entries that match a given severity level or greater.
	LevelGreaterThanOrEqualTo LogEntryLevelEnum `mandatory:"false" contributesTo:"query" name:"levelGreaterThanOrEqualTo" omitEmpty:"true"`

	// The sort order to use when sorting returned resources. Ascending (`ASC`) or descending (`DESC`).
	SortOrder GetJobLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request GetJobLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetJobLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetJobLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetJobLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetJobLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Type {
		if _, ok := GetMappingLogEntryTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", val, strings.Join(GetLogEntryTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingLogEntryLevelEnum(string(request.LevelGreaterThanOrEqualTo)); !ok && request.LevelGreaterThanOrEqualTo != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LevelGreaterThanOrEqualTo: %s. Supported values are: %s.", request.LevelGreaterThanOrEqualTo, strings.Join(GetLogEntryLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetJobLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetJobLogsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetJobLogsResponse wrapper for the GetJobLogs operation
type GetJobLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []LogEntry instances
	Items []LogEntry `presentIn:"body"`

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

func (response GetJobLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetJobLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetJobLogsSortOrderEnum Enum with underlying type: string
type GetJobLogsSortOrderEnum string

// Set of constants representing the allowable values for GetJobLogsSortOrderEnum
const (
	GetJobLogsSortOrderAsc  GetJobLogsSortOrderEnum = "ASC"
	GetJobLogsSortOrderDesc GetJobLogsSortOrderEnum = "DESC"
)

var mappingGetJobLogsSortOrderEnum = map[string]GetJobLogsSortOrderEnum{
	"ASC":  GetJobLogsSortOrderAsc,
	"DESC": GetJobLogsSortOrderDesc,
}

var mappingGetJobLogsSortOrderEnumLowerCase = map[string]GetJobLogsSortOrderEnum{
	"asc":  GetJobLogsSortOrderAsc,
	"desc": GetJobLogsSortOrderDesc,
}

// GetGetJobLogsSortOrderEnumValues Enumerates the set of values for GetJobLogsSortOrderEnum
func GetGetJobLogsSortOrderEnumValues() []GetJobLogsSortOrderEnum {
	values := make([]GetJobLogsSortOrderEnum, 0)
	for _, v := range mappingGetJobLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetJobLogsSortOrderEnumStringValues Enumerates the set of values in String for GetJobLogsSortOrderEnum
func GetGetJobLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetJobLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetJobLogsSortOrderEnum(val string) (GetJobLogsSortOrderEnum, bool) {
	enum, ok := mappingGetJobLogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
