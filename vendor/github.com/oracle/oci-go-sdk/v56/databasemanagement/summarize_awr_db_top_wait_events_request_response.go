// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// SummarizeAwrDbTopWaitEventsRequest wrapper for the SummarizeAwrDbTopWaitEvents operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAwrDbTopWaitEvents.go.html to see an example of how to use SummarizeAwrDbTopWaitEventsRequest.
type SummarizeAwrDbTopWaitEventsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The parameter to filter the database by internal ID.
	// Note that the internal ID of the database can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs
	AwrDbId *string `mandatory:"true" contributesTo:"path" name:"awrDbId"`

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

	// The optional query parameter to filter ASH activities by FOREGROUND or BACKGROUND.
	SessionType SummarizeAwrDbTopWaitEventsSessionTypeEnum `mandatory:"false" contributesTo:"query" name:"sessionType" omitEmpty:"true"`

	// The optional query parameter to filter the database container by an exact ID value.
	// Note that the database container ID can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbSnapshotRanges
	ContainerId *int `mandatory:"false" contributesTo:"query" name:"containerId"`

	// The optional query parameter to filter the number of top categories to be returned.
	TopN *int `mandatory:"false" contributesTo:"query" name:"topN"`

	// The option to sort the AWR top event summary data.
	SortBy SummarizeAwrDbTopWaitEventsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Descending order is the default order.
	SortOrder SummarizeAwrDbTopWaitEventsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request SummarizeAwrDbTopWaitEventsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrDbTopWaitEventsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrDbTopWaitEventsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrDbTopWaitEventsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeAwrDbTopWaitEventsResponse wrapper for the SummarizeAwrDbTopWaitEvents operation
type SummarizeAwrDbTopWaitEventsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AwrDbTopWaitEventCollection instance
	AwrDbTopWaitEventCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrDbTopWaitEventsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrDbTopWaitEventsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrDbTopWaitEventsSessionTypeEnum Enum with underlying type: string
type SummarizeAwrDbTopWaitEventsSessionTypeEnum string

// Set of constants representing the allowable values for SummarizeAwrDbTopWaitEventsSessionTypeEnum
const (
	SummarizeAwrDbTopWaitEventsSessionTypeForeground SummarizeAwrDbTopWaitEventsSessionTypeEnum = "FOREGROUND"
	SummarizeAwrDbTopWaitEventsSessionTypeBackground SummarizeAwrDbTopWaitEventsSessionTypeEnum = "BACKGROUND"
	SummarizeAwrDbTopWaitEventsSessionTypeAll        SummarizeAwrDbTopWaitEventsSessionTypeEnum = "ALL"
)

var mappingSummarizeAwrDbTopWaitEventsSessionType = map[string]SummarizeAwrDbTopWaitEventsSessionTypeEnum{
	"FOREGROUND": SummarizeAwrDbTopWaitEventsSessionTypeForeground,
	"BACKGROUND": SummarizeAwrDbTopWaitEventsSessionTypeBackground,
	"ALL":        SummarizeAwrDbTopWaitEventsSessionTypeAll,
}

// GetSummarizeAwrDbTopWaitEventsSessionTypeEnumValues Enumerates the set of values for SummarizeAwrDbTopWaitEventsSessionTypeEnum
func GetSummarizeAwrDbTopWaitEventsSessionTypeEnumValues() []SummarizeAwrDbTopWaitEventsSessionTypeEnum {
	values := make([]SummarizeAwrDbTopWaitEventsSessionTypeEnum, 0)
	for _, v := range mappingSummarizeAwrDbTopWaitEventsSessionType {
		values = append(values, v)
	}
	return values
}

// SummarizeAwrDbTopWaitEventsSortByEnum Enum with underlying type: string
type SummarizeAwrDbTopWaitEventsSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrDbTopWaitEventsSortByEnum
const (
	SummarizeAwrDbTopWaitEventsSortByWaitsPersec       SummarizeAwrDbTopWaitEventsSortByEnum = "WAITS_PERSEC"
	SummarizeAwrDbTopWaitEventsSortByAvgWaitTimePersec SummarizeAwrDbTopWaitEventsSortByEnum = "AVG_WAIT_TIME_PERSEC"
)

var mappingSummarizeAwrDbTopWaitEventsSortBy = map[string]SummarizeAwrDbTopWaitEventsSortByEnum{
	"WAITS_PERSEC":         SummarizeAwrDbTopWaitEventsSortByWaitsPersec,
	"AVG_WAIT_TIME_PERSEC": SummarizeAwrDbTopWaitEventsSortByAvgWaitTimePersec,
}

// GetSummarizeAwrDbTopWaitEventsSortByEnumValues Enumerates the set of values for SummarizeAwrDbTopWaitEventsSortByEnum
func GetSummarizeAwrDbTopWaitEventsSortByEnumValues() []SummarizeAwrDbTopWaitEventsSortByEnum {
	values := make([]SummarizeAwrDbTopWaitEventsSortByEnum, 0)
	for _, v := range mappingSummarizeAwrDbTopWaitEventsSortBy {
		values = append(values, v)
	}
	return values
}

// SummarizeAwrDbTopWaitEventsSortOrderEnum Enum with underlying type: string
type SummarizeAwrDbTopWaitEventsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrDbTopWaitEventsSortOrderEnum
const (
	SummarizeAwrDbTopWaitEventsSortOrderAsc  SummarizeAwrDbTopWaitEventsSortOrderEnum = "ASC"
	SummarizeAwrDbTopWaitEventsSortOrderDesc SummarizeAwrDbTopWaitEventsSortOrderEnum = "DESC"
)

var mappingSummarizeAwrDbTopWaitEventsSortOrder = map[string]SummarizeAwrDbTopWaitEventsSortOrderEnum{
	"ASC":  SummarizeAwrDbTopWaitEventsSortOrderAsc,
	"DESC": SummarizeAwrDbTopWaitEventsSortOrderDesc,
}

// GetSummarizeAwrDbTopWaitEventsSortOrderEnumValues Enumerates the set of values for SummarizeAwrDbTopWaitEventsSortOrderEnum
func GetSummarizeAwrDbTopWaitEventsSortOrderEnumValues() []SummarizeAwrDbTopWaitEventsSortOrderEnum {
	values := make([]SummarizeAwrDbTopWaitEventsSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrDbTopWaitEventsSortOrder {
		values = append(values, v)
	}
	return values
}
