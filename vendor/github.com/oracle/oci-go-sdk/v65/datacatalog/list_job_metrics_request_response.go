// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListJobMetricsRequest wrapper for the ListJobMetrics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListJobMetrics.go.html to see an example of how to use ListJobMetricsRequest.
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
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
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
func (request ListJobMetricsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJobMetricsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobMetricsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJobMetricsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingListJobMetricsFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListJobMetricsFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListJobMetricsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJobMetricsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobMetricsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJobMetricsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListJobMetricsFieldsEnum = map[string]ListJobMetricsFieldsEnum{
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

var mappingListJobMetricsFieldsEnumLowerCase = map[string]ListJobMetricsFieldsEnum{
	"key":             ListJobMetricsFieldsKey,
	"description":     ListJobMetricsFieldsDescription,
	"displayname":     ListJobMetricsFieldsDisplayname,
	"timeinserted":    ListJobMetricsFieldsTimeinserted,
	"category":        ListJobMetricsFieldsCategory,
	"subcategory":     ListJobMetricsFieldsSubcategory,
	"unit":            ListJobMetricsFieldsUnit,
	"value":           ListJobMetricsFieldsValue,
	"batchkey":        ListJobMetricsFieldsBatchkey,
	"jobexecutionkey": ListJobMetricsFieldsJobexecutionkey,
	"timecreated":     ListJobMetricsFieldsTimecreated,
	"uri":             ListJobMetricsFieldsUri,
}

// GetListJobMetricsFieldsEnumValues Enumerates the set of values for ListJobMetricsFieldsEnum
func GetListJobMetricsFieldsEnumValues() []ListJobMetricsFieldsEnum {
	values := make([]ListJobMetricsFieldsEnum, 0)
	for _, v := range mappingListJobMetricsFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobMetricsFieldsEnumStringValues Enumerates the set of values in String for ListJobMetricsFieldsEnum
func GetListJobMetricsFieldsEnumStringValues() []string {
	return []string{
		"key",
		"description",
		"displayName",
		"timeInserted",
		"category",
		"subCategory",
		"unit",
		"value",
		"batchKey",
		"jobExecutionKey",
		"timeCreated",
		"uri",
	}
}

// GetMappingListJobMetricsFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobMetricsFieldsEnum(val string) (ListJobMetricsFieldsEnum, bool) {
	enum, ok := mappingListJobMetricsFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobMetricsSortByEnum Enum with underlying type: string
type ListJobMetricsSortByEnum string

// Set of constants representing the allowable values for ListJobMetricsSortByEnum
const (
	ListJobMetricsSortByTimecreated ListJobMetricsSortByEnum = "TIMECREATED"
	ListJobMetricsSortByDisplayname ListJobMetricsSortByEnum = "DISPLAYNAME"
)

var mappingListJobMetricsSortByEnum = map[string]ListJobMetricsSortByEnum{
	"TIMECREATED": ListJobMetricsSortByTimecreated,
	"DISPLAYNAME": ListJobMetricsSortByDisplayname,
}

var mappingListJobMetricsSortByEnumLowerCase = map[string]ListJobMetricsSortByEnum{
	"timecreated": ListJobMetricsSortByTimecreated,
	"displayname": ListJobMetricsSortByDisplayname,
}

// GetListJobMetricsSortByEnumValues Enumerates the set of values for ListJobMetricsSortByEnum
func GetListJobMetricsSortByEnumValues() []ListJobMetricsSortByEnum {
	values := make([]ListJobMetricsSortByEnum, 0)
	for _, v := range mappingListJobMetricsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobMetricsSortByEnumStringValues Enumerates the set of values in String for ListJobMetricsSortByEnum
func GetListJobMetricsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListJobMetricsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobMetricsSortByEnum(val string) (ListJobMetricsSortByEnum, bool) {
	enum, ok := mappingListJobMetricsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobMetricsSortOrderEnum Enum with underlying type: string
type ListJobMetricsSortOrderEnum string

// Set of constants representing the allowable values for ListJobMetricsSortOrderEnum
const (
	ListJobMetricsSortOrderAsc  ListJobMetricsSortOrderEnum = "ASC"
	ListJobMetricsSortOrderDesc ListJobMetricsSortOrderEnum = "DESC"
)

var mappingListJobMetricsSortOrderEnum = map[string]ListJobMetricsSortOrderEnum{
	"ASC":  ListJobMetricsSortOrderAsc,
	"DESC": ListJobMetricsSortOrderDesc,
}

var mappingListJobMetricsSortOrderEnumLowerCase = map[string]ListJobMetricsSortOrderEnum{
	"asc":  ListJobMetricsSortOrderAsc,
	"desc": ListJobMetricsSortOrderDesc,
}

// GetListJobMetricsSortOrderEnumValues Enumerates the set of values for ListJobMetricsSortOrderEnum
func GetListJobMetricsSortOrderEnumValues() []ListJobMetricsSortOrderEnum {
	values := make([]ListJobMetricsSortOrderEnum, 0)
	for _, v := range mappingListJobMetricsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobMetricsSortOrderEnumStringValues Enumerates the set of values in String for ListJobMetricsSortOrderEnum
func GetListJobMetricsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJobMetricsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobMetricsSortOrderEnum(val string) (ListJobMetricsSortOrderEnum, bool) {
	enum, ok := mappingListJobMetricsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
