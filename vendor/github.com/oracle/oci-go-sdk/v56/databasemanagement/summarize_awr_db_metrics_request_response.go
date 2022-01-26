// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// SummarizeAwrDbMetricsRequest wrapper for the SummarizeAwrDbMetrics operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbMetrics.go.html to see an example of how to use SummarizeAwrDbMetricsRequest.
type SummarizeAwrDbMetricsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The parameter to filter the database by internal ID.
	// Note that the internal ID of the database can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs
	AwrDbId *string `mandatory:"true" contributesTo:"path" name:"awrDbId"`

	// The required multiple value query parameter to filter the entity name.
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// The optional single value query parameter to filter the database instance number.
	InstNum *string `mandatory:"false" contributesTo:"query" name:"instNum"`

	// The optional greater than or equal to filter on the snapshot ID.
	BeginSnIdGreaterThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"beginSnIdGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the snapshot ID.
	EndSnIdLessThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"endSnIdLessThanOrEqualTo"`

	// The optional greater than or equal to query parameter to filter the timestamp.
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp.
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The optional query parameter to filter the database container by an exact ID value.
	// Note that the database container ID can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbSnapshotRanges
	ContainerId *int `mandatory:"false" contributesTo:"query" name:"containerId"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in large paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort the AWR time series summary data.
	SortBy SummarizeAwrDbMetricsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Descending order is the default order.
	SortOrder SummarizeAwrDbMetricsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeAwrDbMetricsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDbMetricsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDbMetricsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDbMetricsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeAwrDbMetricsResponse wrapper for the SummarizeAwrDbMetrics operation
type SummarizeAwrDbMetricsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDbMetricCollection instances
	AwrDbMetricCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDbMetricsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDbMetricsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDbMetricsSortByEnum Enum with underlying type: string
type SummarizeAwrDbMetricsSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDbMetricsSortByEnum
const (
	SummarizeAwrDbMetricsSortByTimestamp SummarizeAwrDbMetricsSortByEnum = "TIMESTAMP"
	SummarizeAwrDbMetricsSortByName      SummarizeAwrDbMetricsSortByEnum = "NAME"
)

var mappingSummarizeAwrDbMetricsSortBy = map[string]SummarizeAwrDbMetricsSortByEnum{
	"TIMESTAMP": SummarizeAwrDbMetricsSortByTimestamp,
	"NAME":      SummarizeAwrDbMetricsSortByName,
}

// GetSummarizeAwrDbMetricsSortByEnumValues Enumerates the set of values for SummarizeAwrDbMetricsSortByEnum
func GetSummarizeAwrDbMetricsSortByEnumValues() []SummarizeAwrDbMetricsSortByEnum {
	values := make([]SummarizeAwrDbMetricsSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDbMetricsSortBy {
		values = append(values, v)
	}
	return values
}

// SummarizeAwrDbMetricsSortOrderEnum Enum with underlying type: string
type SummarizeAwrDbMetricsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDbMetricsSortOrderEnum
const (
	SummarizeAwrDbMetricsSortOrderAsc  SummarizeAwrDbMetricsSortOrderEnum = "ASC"
	SummarizeAwrDbMetricsSortOrderDesc SummarizeAwrDbMetricsSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDbMetricsSortOrder = map[string]SummarizeAwrDbMetricsSortOrderEnum{
	"ASC":  SummarizeAwrDbMetricsSortOrderAsc,
	"DESC": SummarizeAwrDbMetricsSortOrderDesc,
}

// GetSummarizeAwrDbMetricsSortOrderEnumValues Enumerates the set of values for SummarizeAwrDbMetricsSortOrderEnum
func GetSummarizeAwrDbMetricsSortOrderEnumValues() []SummarizeAwrDbMetricsSortOrderEnum {
	values := make([]SummarizeAwrDbMetricsSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDbMetricsSortOrder {
		values = append(values, v)
	}
	return values
}
