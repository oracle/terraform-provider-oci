// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeAwrDatabaseSnapshotRangesRequest wrapper for the SummarizeAwrDatabaseSnapshotRanges operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseSnapshotRanges.go.html to see an example of how to use SummarizeAwrDatabaseSnapshotRangesRequest.
type SummarizeAwrDatabaseSnapshotRangesRequest struct {

	// Unique Awr Hub identifier
	AwrHubId *string `mandatory:"true" contributesTo:"path" name:"awrHubId"`

	// The optional single value query parameter to filter the entity name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The optional greater than or equal to query parameter to filter the timestamp. The timestamp format to be followed is: YYYY-MM-DDTHH:MM:SSZ, example 2020-12-03T19:00:53Z
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp. The timestamp format to be followed is: YYYY-MM-DDTHH:MM:SSZ, example 2020-12-03T19:00:53Z
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort the AWR summary data.
	SortBy SummarizeAwrDatabaseSnapshotRangesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeAwrDatabaseSnapshotRangesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeAwrDatabaseSnapshotRangesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDatabaseSnapshotRangesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDatabaseSnapshotRangesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDatabaseSnapshotRangesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAwrDatabaseSnapshotRangesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAwrDatabaseSnapshotRangesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeAwrDatabaseSnapshotRangesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDatabaseSnapshotRangesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeAwrDatabaseSnapshotRangesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAwrDatabaseSnapshotRangesResponse wrapper for the SummarizeAwrDatabaseSnapshotRanges operation
type SummarizeAwrDatabaseSnapshotRangesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDatabaseSnapshotRangeCollection instances
	AwrDatabaseSnapshotRangeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDatabaseSnapshotRangesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDatabaseSnapshotRangesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDatabaseSnapshotRangesSortByEnum Enum with underlying type: string
type SummarizeAwrDatabaseSnapshotRangesSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseSnapshotRangesSortByEnum
const (
	SummarizeAwrDatabaseSnapshotRangesSortByEndIntervalTime SummarizeAwrDatabaseSnapshotRangesSortByEnum = "END_INTERVAL_TIME"
	SummarizeAwrDatabaseSnapshotRangesSortByName            SummarizeAwrDatabaseSnapshotRangesSortByEnum = "NAME"
)

var mappingSummarizeAwrDatabaseSnapshotRangesSortByEnum = map[string]SummarizeAwrDatabaseSnapshotRangesSortByEnum{
	"END_INTERVAL_TIME": SummarizeAwrDatabaseSnapshotRangesSortByEndIntervalTime,
	"NAME":              SummarizeAwrDatabaseSnapshotRangesSortByName,
}

var mappingSummarizeAwrDatabaseSnapshotRangesSortByEnumLowerCase = map[string]SummarizeAwrDatabaseSnapshotRangesSortByEnum{
	"end_interval_time": SummarizeAwrDatabaseSnapshotRangesSortByEndIntervalTime,
	"name":              SummarizeAwrDatabaseSnapshotRangesSortByName,
}

// GetSummarizeAwrDatabaseSnapshotRangesSortByEnumValues Enumerates the set of values for SummarizeAwrDatabaseSnapshotRangesSortByEnum
func GetSummarizeAwrDatabaseSnapshotRangesSortByEnumValues() []SummarizeAwrDatabaseSnapshotRangesSortByEnum {
	values := make([]SummarizeAwrDatabaseSnapshotRangesSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseSnapshotRangesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseSnapshotRangesSortByEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseSnapshotRangesSortByEnum
func GetSummarizeAwrDatabaseSnapshotRangesSortByEnumStringValues() []string {
	return []string{
		"END_INTERVAL_TIME",
		"NAME",
	}
}

// GetMappingSummarizeAwrDatabaseSnapshotRangesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseSnapshotRangesSortByEnum(val string) (SummarizeAwrDatabaseSnapshotRangesSortByEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseSnapshotRangesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDatabaseSnapshotRangesSortOrderEnum Enum with underlying type: string
type SummarizeAwrDatabaseSnapshotRangesSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseSnapshotRangesSortOrderEnum
const (
	SummarizeAwrDatabaseSnapshotRangesSortOrderAsc  SummarizeAwrDatabaseSnapshotRangesSortOrderEnum = "ASC"
	SummarizeAwrDatabaseSnapshotRangesSortOrderDesc SummarizeAwrDatabaseSnapshotRangesSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDatabaseSnapshotRangesSortOrderEnum = map[string]SummarizeAwrDatabaseSnapshotRangesSortOrderEnum{
	"ASC":  SummarizeAwrDatabaseSnapshotRangesSortOrderAsc,
	"DESC": SummarizeAwrDatabaseSnapshotRangesSortOrderDesc,
}

var mappingSummarizeAwrDatabaseSnapshotRangesSortOrderEnumLowerCase = map[string]SummarizeAwrDatabaseSnapshotRangesSortOrderEnum{
	"asc":  SummarizeAwrDatabaseSnapshotRangesSortOrderAsc,
	"desc": SummarizeAwrDatabaseSnapshotRangesSortOrderDesc,
}

// GetSummarizeAwrDatabaseSnapshotRangesSortOrderEnumValues Enumerates the set of values for SummarizeAwrDatabaseSnapshotRangesSortOrderEnum
func GetSummarizeAwrDatabaseSnapshotRangesSortOrderEnumValues() []SummarizeAwrDatabaseSnapshotRangesSortOrderEnum {
	values := make([]SummarizeAwrDatabaseSnapshotRangesSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseSnapshotRangesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseSnapshotRangesSortOrderEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseSnapshotRangesSortOrderEnum
func GetSummarizeAwrDatabaseSnapshotRangesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeAwrDatabaseSnapshotRangesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseSnapshotRangesSortOrderEnum(val string) (SummarizeAwrDatabaseSnapshotRangesSortOrderEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseSnapshotRangesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
