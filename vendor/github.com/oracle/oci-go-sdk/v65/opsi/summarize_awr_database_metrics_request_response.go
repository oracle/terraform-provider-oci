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

// SummarizeAwrDatabaseMetricsRequest wrapper for the SummarizeAwrDatabaseMetrics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseMetrics.go.html to see an example of how to use SummarizeAwrDatabaseMetricsRequest.
type SummarizeAwrDatabaseMetricsRequest struct {

	// Unique Awr Hub identifier
	AwrHubId *string `mandatory:"true" contributesTo:"path" name:"awrHubId"`

	// The internal ID of the database. The internal ID of the database is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /awrHubs/{awrHubId}/awrDatabases
	AwrSourceDatabaseIdentifier *string `mandatory:"true" contributesTo:"query" name:"awrSourceDatabaseIdentifier"`

	// The required multiple value query parameter to filter the entity name.
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// The optional single value query parameter to filter by database instance number.
	InstanceNumber *string `mandatory:"false" contributesTo:"query" name:"instanceNumber"`

	// The optional greater than or equal to filter on the snapshot ID.
	BeginSnapshotIdentifierGreaterThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"beginSnapshotIdentifierGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the snapshot Identifier.
	EndSnapshotIdentifierLessThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"endSnapshotIdentifierLessThanOrEqualTo"`

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

	// The option to sort the AWR time series summary data.
	SortBy SummarizeAwrDatabaseMetricsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeAwrDatabaseMetricsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeAwrDatabaseMetricsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDatabaseMetricsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDatabaseMetricsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDatabaseMetricsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAwrDatabaseMetricsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAwrDatabaseMetricsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeAwrDatabaseMetricsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAwrDatabaseMetricsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeAwrDatabaseMetricsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAwrDatabaseMetricsResponse wrapper for the SummarizeAwrDatabaseMetrics operation
type SummarizeAwrDatabaseMetricsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDatabaseMetricCollection instances
	AwrDatabaseMetricCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDatabaseMetricsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDatabaseMetricsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDatabaseMetricsSortByEnum Enum with underlying type: string
type SummarizeAwrDatabaseMetricsSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseMetricsSortByEnum
const (
	SummarizeAwrDatabaseMetricsSortByTimestamp SummarizeAwrDatabaseMetricsSortByEnum = "TIMESTAMP"
	SummarizeAwrDatabaseMetricsSortByName      SummarizeAwrDatabaseMetricsSortByEnum = "NAME"
)

var mappingSummarizeAwrDatabaseMetricsSortByEnum = map[string]SummarizeAwrDatabaseMetricsSortByEnum{
	"TIMESTAMP": SummarizeAwrDatabaseMetricsSortByTimestamp,
	"NAME":      SummarizeAwrDatabaseMetricsSortByName,
}

var mappingSummarizeAwrDatabaseMetricsSortByEnumLowerCase = map[string]SummarizeAwrDatabaseMetricsSortByEnum{
	"timestamp": SummarizeAwrDatabaseMetricsSortByTimestamp,
	"name":      SummarizeAwrDatabaseMetricsSortByName,
}

// GetSummarizeAwrDatabaseMetricsSortByEnumValues Enumerates the set of values for SummarizeAwrDatabaseMetricsSortByEnum
func GetSummarizeAwrDatabaseMetricsSortByEnumValues() []SummarizeAwrDatabaseMetricsSortByEnum {
	values := make([]SummarizeAwrDatabaseMetricsSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseMetricsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseMetricsSortByEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseMetricsSortByEnum
func GetSummarizeAwrDatabaseMetricsSortByEnumStringValues() []string {
	return []string{
		"TIMESTAMP",
		"NAME",
	}
}

// GetMappingSummarizeAwrDatabaseMetricsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseMetricsSortByEnum(val string) (SummarizeAwrDatabaseMetricsSortByEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseMetricsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAwrDatabaseMetricsSortOrderEnum Enum with underlying type: string
type SummarizeAwrDatabaseMetricsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDatabaseMetricsSortOrderEnum
const (
	SummarizeAwrDatabaseMetricsSortOrderAsc  SummarizeAwrDatabaseMetricsSortOrderEnum = "ASC"
	SummarizeAwrDatabaseMetricsSortOrderDesc SummarizeAwrDatabaseMetricsSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDatabaseMetricsSortOrderEnum = map[string]SummarizeAwrDatabaseMetricsSortOrderEnum{
	"ASC":  SummarizeAwrDatabaseMetricsSortOrderAsc,
	"DESC": SummarizeAwrDatabaseMetricsSortOrderDesc,
}

var mappingSummarizeAwrDatabaseMetricsSortOrderEnumLowerCase = map[string]SummarizeAwrDatabaseMetricsSortOrderEnum{
	"asc":  SummarizeAwrDatabaseMetricsSortOrderAsc,
	"desc": SummarizeAwrDatabaseMetricsSortOrderDesc,
}

// GetSummarizeAwrDatabaseMetricsSortOrderEnumValues Enumerates the set of values for SummarizeAwrDatabaseMetricsSortOrderEnum
func GetSummarizeAwrDatabaseMetricsSortOrderEnumValues() []SummarizeAwrDatabaseMetricsSortOrderEnum {
	values := make([]SummarizeAwrDatabaseMetricsSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDatabaseMetricsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAwrDatabaseMetricsSortOrderEnumStringValues Enumerates the set of values in String for SummarizeAwrDatabaseMetricsSortOrderEnum
func GetSummarizeAwrDatabaseMetricsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeAwrDatabaseMetricsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAwrDatabaseMetricsSortOrderEnum(val string) (SummarizeAwrDatabaseMetricsSortOrderEnum, bool) {
	enum, ok := mappingSummarizeAwrDatabaseMetricsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
