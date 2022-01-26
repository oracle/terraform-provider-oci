// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// SummarizeJobExecutionsStatusesRequest wrapper for the SummarizeJobExecutionsStatuses operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeJobExecutionsStatuses.go.html to see an example of how to use SummarizeJobExecutionsStatusesRequest.
type SummarizeJobExecutionsStatusesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The start time of the time range to retrieve the status summary of job executions
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	StartTime *string `mandatory:"true" contributesTo:"query" name:"startTime"`

	// The end time of the time range to retrieve the status summary of job executions
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	EndTime *string `mandatory:"true" contributesTo:"query" name:"endTime"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The identifier of the resource.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	ManagedDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"managedDatabaseGroupId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"false" contributesTo:"query" name:"managedDatabaseId"`

	// A filter to return only resources that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy SummarizeJobExecutionsStatusesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder SummarizeJobExecutionsStatusesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeJobExecutionsStatusesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeJobExecutionsStatusesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeJobExecutionsStatusesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeJobExecutionsStatusesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeJobExecutionsStatusesResponse wrapper for the SummarizeJobExecutionsStatuses operation
type SummarizeJobExecutionsStatusesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The JobExecutionsStatusSummaryCollection instance
	JobExecutionsStatusSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response SummarizeJobExecutionsStatusesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeJobExecutionsStatusesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeJobExecutionsStatusesSortByEnum Enum with underlying type: string
type SummarizeJobExecutionsStatusesSortByEnum string

// Set of constants representing the allowable values for SummarizeJobExecutionsStatusesSortByEnum
const (
	SummarizeJobExecutionsStatusesSortByTimecreated SummarizeJobExecutionsStatusesSortByEnum = "TIMECREATED"
	SummarizeJobExecutionsStatusesSortByName        SummarizeJobExecutionsStatusesSortByEnum = "NAME"
)

var mappingSummarizeJobExecutionsStatusesSortBy = map[string]SummarizeJobExecutionsStatusesSortByEnum{
	"TIMECREATED": SummarizeJobExecutionsStatusesSortByTimecreated,
	"NAME":        SummarizeJobExecutionsStatusesSortByName,
}

// GetSummarizeJobExecutionsStatusesSortByEnumValues Enumerates the set of values for SummarizeJobExecutionsStatusesSortByEnum
func GetSummarizeJobExecutionsStatusesSortByEnumValues() []SummarizeJobExecutionsStatusesSortByEnum {
	values := make([]SummarizeJobExecutionsStatusesSortByEnum, 0)
	for _, v := range mappingSummarizeJobExecutionsStatusesSortBy {
		values = append(values, v)
	}
	return values
}

// SummarizeJobExecutionsStatusesSortOrderEnum Enum with underlying type: string
type SummarizeJobExecutionsStatusesSortOrderEnum string

// Set of constants representing the allowable values for SummarizeJobExecutionsStatusesSortOrderEnum
const (
	SummarizeJobExecutionsStatusesSortOrderAsc  SummarizeJobExecutionsStatusesSortOrderEnum = "ASC"
	SummarizeJobExecutionsStatusesSortOrderDesc SummarizeJobExecutionsStatusesSortOrderEnum = "DESC"
)

var mappingSummarizeJobExecutionsStatusesSortOrder = map[string]SummarizeJobExecutionsStatusesSortOrderEnum{
	"ASC":  SummarizeJobExecutionsStatusesSortOrderAsc,
	"DESC": SummarizeJobExecutionsStatusesSortOrderDesc,
}

// GetSummarizeJobExecutionsStatusesSortOrderEnumValues Enumerates the set of values for SummarizeJobExecutionsStatusesSortOrderEnum
func GetSummarizeJobExecutionsStatusesSortOrderEnumValues() []SummarizeJobExecutionsStatusesSortOrderEnum {
	values := make([]SummarizeJobExecutionsStatusesSortOrderEnum, 0)
	for _, v := range mappingSummarizeJobExecutionsStatusesSortOrder {
		values = append(values, v)
	}
	return values
}
