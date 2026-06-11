// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeProtectedDatabaseAnalyticsRequest wrapper for the SummarizeProtectedDatabaseAnalytics operation
type SummarizeProtectedDatabaseAnalyticsRequest struct {

	// The compartment OCID.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The dimension used to return the aggregated metrics for the protected database.
	// You can group the metrics by the protected database health status (`health`), the real-time data protection status (`isRedoLogsEnabled`), the associated protection policy (`protectionPolicyId`), and the life cycle state (`lifecycleState`).
	// If you skip the `groupBy` parameter value, then the API returns a single aggregated value per metric, without grouping.
	// For example, for the `COUNT` metric, it returns the count of protected databases in the specified compartment.
	GroupBy []ProtectedDatabaseAnalyticsGroupByEnum `contributesTo:"query" name:"groupBy" omitEmpty:"true" collectionFormat:"multi"`

	// The protected database metric. You can specify either one or both of these metric names.
	// Allowed values are:
	//  - `COUNT`
	//  - `BACKUP_SPACE_USED_IN_GBS`
	MetricName []ProtectedDatabaseAnalyticsMetricEnum `contributesTo:"query" name:"metricName" omitEmpty:"true" collectionFormat:"multi"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeProtectedDatabaseAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeProtectedDatabaseAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeProtectedDatabaseAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeProtectedDatabaseAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeProtectedDatabaseAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.GroupBy {
		if _, ok := GetMappingProtectedDatabaseAnalyticsGroupByEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", val, strings.Join(GetProtectedDatabaseAnalyticsGroupByEnumStringValues(), ",")))
		}
	}

	for _, val := range request.MetricName {
		if _, ok := GetMappingProtectedDatabaseAnalyticsMetricEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricName: %s. Supported values are: %s.", val, strings.Join(GetProtectedDatabaseAnalyticsMetricEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeProtectedDatabaseAnalyticsResponse wrapper for the SummarizeProtectedDatabaseAnalytics operation
type SummarizeProtectedDatabaseAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProtectedDatabaseAnalyticsCollection instances
	ProtectedDatabaseAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeProtectedDatabaseAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeProtectedDatabaseAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
