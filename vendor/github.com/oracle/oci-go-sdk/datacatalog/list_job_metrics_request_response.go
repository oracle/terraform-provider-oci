// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListJobMetricsRequest wrapper for the ListJobMetrics operation
type ListJobMetricsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique job key.
	JobKey *string `mandatory:"true" contributesTo:"path" name:"jobKey"`

	// The key of the job execution.
	JobExecutionKey *string `mandatory:"true" contributesTo:"path" name:"jobExecutionKey"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu".
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// Category of this metric.
	Category *string `mandatory:"false" contributesTo:"query" name:"category"`

	// Sub category of this metric under the category. Used for aggregating values. May be null.
	SubCategory *string `mandatory:"false" contributesTo:"query" name:"subCategory"`

	// Unit of this metric.
	Unit *string `mandatory:"false" contributesTo:"query" name:"unit"`

	// Value of this metric.
	Value *string `mandatory:"false" contributesTo:"query" name:"value"`

	// Batch key for grouping, may be null.
	BatchKey *string `mandatory:"false" contributesTo:"query" name:"batchKey"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// The time the metric was logged or captured in the system where the job executed.
	// An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeInserted *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeInserted"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Specifies the fields to return in a job metric summary response.
	Fields []ListJobMetricsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListJobMetricsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListJobMetricsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListJobMetricsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJobMetricsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobMetricsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListJobMetricsResponse wrapper for the ListJobMetrics operation
type ListJobMetricsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JobMetricCollection instances
	JobMetricCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJobMetricsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJobMetricsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJobMetricsFieldsEnum Enum with underlying type: string
type ListJobMetricsFieldsEnum string

// Set of constants representing the allowable values for ListJobMetricsFieldsEnum
const (
	ListJobMetricsFieldsKey             ListJobMetricsFieldsEnum = "key"
	ListJobMetricsFieldsDescription     ListJobMetricsFieldsEnum = "description"
	ListJobMetricsFieldsDisplayname     ListJobMetricsFieldsEnum = "displayName"
	ListJobMetricsFieldsTimeinserted    ListJobMetricsFieldsEnum = "timeInserted"
	ListJobMetricsFieldsCategory        ListJobMetricsFieldsEnum = "category"
	ListJobMetricsFieldsSubcategory     ListJobMetricsFieldsEnum = "subCategory"
	ListJobMetricsFieldsUnit            ListJobMetricsFieldsEnum = "unit"
	ListJobMetricsFieldsValue           ListJobMetricsFieldsEnum = "value"
	ListJobMetricsFieldsBatchkey        ListJobMetricsFieldsEnum = "batchKey"
	ListJobMetricsFieldsJobexecutionkey ListJobMetricsFieldsEnum = "jobExecutionKey"
	ListJobMetricsFieldsTimecreated     ListJobMetricsFieldsEnum = "timeCreated"
	ListJobMetricsFieldsUri             ListJobMetricsFieldsEnum = "uri"
)

var mappingListJobMetricsFields = map[string]ListJobMetricsFieldsEnum{
	"key":             ListJobMetricsFieldsKey,
	"description":     ListJobMetricsFieldsDescription,
	"displayName":     ListJobMetricsFieldsDisplayname,
	"timeInserted":    ListJobMetricsFieldsTimeinserted,
	"category":        ListJobMetricsFieldsCategory,
	"subCategory":     ListJobMetricsFieldsSubcategory,
	"unit":            ListJobMetricsFieldsUnit,
	"value":           ListJobMetricsFieldsValue,
	"batchKey":        ListJobMetricsFieldsBatchkey,
	"jobExecutionKey": ListJobMetricsFieldsJobexecutionkey,
	"timeCreated":     ListJobMetricsFieldsTimecreated,
	"uri":             ListJobMetricsFieldsUri,
}

// GetListJobMetricsFieldsEnumValues Enumerates the set of values for ListJobMetricsFieldsEnum
func GetListJobMetricsFieldsEnumValues() []ListJobMetricsFieldsEnum {
	values := make([]ListJobMetricsFieldsEnum, 0)
	for _, v := range mappingListJobMetricsFields {
		values = append(values, v)
	}
	return values
}

// ListJobMetricsSortByEnum Enum with underlying type: string
type ListJobMetricsSortByEnum string

// Set of constants representing the allowable values for ListJobMetricsSortByEnum
const (
	ListJobMetricsSortByTimecreated ListJobMetricsSortByEnum = "TIMECREATED"
	ListJobMetricsSortByDisplayname ListJobMetricsSortByEnum = "DISPLAYNAME"
)

var mappingListJobMetricsSortBy = map[string]ListJobMetricsSortByEnum{
	"TIMECREATED": ListJobMetricsSortByTimecreated,
	"DISPLAYNAME": ListJobMetricsSortByDisplayname,
}

// GetListJobMetricsSortByEnumValues Enumerates the set of values for ListJobMetricsSortByEnum
func GetListJobMetricsSortByEnumValues() []ListJobMetricsSortByEnum {
	values := make([]ListJobMetricsSortByEnum, 0)
	for _, v := range mappingListJobMetricsSortBy {
		values = append(values, v)
	}
	return values
}

// ListJobMetricsSortOrderEnum Enum with underlying type: string
type ListJobMetricsSortOrderEnum string

// Set of constants representing the allowable values for ListJobMetricsSortOrderEnum
const (
	ListJobMetricsSortOrderAsc  ListJobMetricsSortOrderEnum = "ASC"
	ListJobMetricsSortOrderDesc ListJobMetricsSortOrderEnum = "DESC"
)

var mappingListJobMetricsSortOrder = map[string]ListJobMetricsSortOrderEnum{
	"ASC":  ListJobMetricsSortOrderAsc,
	"DESC": ListJobMetricsSortOrderDesc,
}

// GetListJobMetricsSortOrderEnumValues Enumerates the set of values for ListJobMetricsSortOrderEnum
func GetListJobMetricsSortOrderEnumValues() []ListJobMetricsSortOrderEnum {
	values := make([]ListJobMetricsSortOrderEnum, 0)
	for _, v := range mappingListJobMetricsSortOrder {
		values = append(values, v)
	}
	return values
}
