// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// SummarizeAwrDbSnapshotRangesRequest wrapper for the SummarizeAwrDbSnapshotRanges operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbSnapshotRanges.go.html to see an example of how to use SummarizeAwrDbSnapshotRangesRequest.
type SummarizeAwrDbSnapshotRangesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The optional single value query parameter to filter the entity name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The optional greater than or equal to query parameter to filter the timestamp.
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp.
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort the AWR summary data.
	SortBy SummarizeAwrDbSnapshotRangesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Descending order is the default order.
	SortOrder SummarizeAwrDbSnapshotRangesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request SummarizeAwrDbSnapshotRangesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDbSnapshotRangesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDbSnapshotRangesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDbSnapshotRangesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeAwrDbSnapshotRangesResponse wrapper for the SummarizeAwrDbSnapshotRanges operation
type SummarizeAwrDbSnapshotRangesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDbSnapshotRangeCollection instances
	AwrDbSnapshotRangeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDbSnapshotRangesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDbSnapshotRangesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDbSnapshotRangesSortByEnum Enum with underlying type: string
type SummarizeAwrDbSnapshotRangesSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDbSnapshotRangesSortByEnum
const (
	SummarizeAwrDbSnapshotRangesSortByEndIntervalTime SummarizeAwrDbSnapshotRangesSortByEnum = "END_INTERVAL_TIME"
	SummarizeAwrDbSnapshotRangesSortByName            SummarizeAwrDbSnapshotRangesSortByEnum = "NAME"
)

var mappingSummarizeAwrDbSnapshotRangesSortBy = map[string]SummarizeAwrDbSnapshotRangesSortByEnum{
	"END_INTERVAL_TIME": SummarizeAwrDbSnapshotRangesSortByEndIntervalTime,
	"NAME":              SummarizeAwrDbSnapshotRangesSortByName,
}

// GetSummarizeAwrDbSnapshotRangesSortByEnumValues Enumerates the set of values for SummarizeAwrDbSnapshotRangesSortByEnum
func GetSummarizeAwrDbSnapshotRangesSortByEnumValues() []SummarizeAwrDbSnapshotRangesSortByEnum {
	values := make([]SummarizeAwrDbSnapshotRangesSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDbSnapshotRangesSortBy {
		values = append(values, v)
	}
	return values
}

// SummarizeAwrDbSnapshotRangesSortOrderEnum Enum with underlying type: string
type SummarizeAwrDbSnapshotRangesSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDbSnapshotRangesSortOrderEnum
const (
	SummarizeAwrDbSnapshotRangesSortOrderAsc  SummarizeAwrDbSnapshotRangesSortOrderEnum = "ASC"
	SummarizeAwrDbSnapshotRangesSortOrderDesc SummarizeAwrDbSnapshotRangesSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDbSnapshotRangesSortOrder = map[string]SummarizeAwrDbSnapshotRangesSortOrderEnum{
	"ASC":  SummarizeAwrDbSnapshotRangesSortOrderAsc,
	"DESC": SummarizeAwrDbSnapshotRangesSortOrderDesc,
}

// GetSummarizeAwrDbSnapshotRangesSortOrderEnumValues Enumerates the set of values for SummarizeAwrDbSnapshotRangesSortOrderEnum
func GetSummarizeAwrDbSnapshotRangesSortOrderEnumValues() []SummarizeAwrDbSnapshotRangesSortOrderEnum {
	values := make([]SummarizeAwrDbSnapshotRangesSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDbSnapshotRangesSortOrder {
		values = append(values, v)
	}
	return values
}
