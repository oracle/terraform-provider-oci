// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDataVerificationTableRowCountsRequest wrapper for the ListDataVerificationTableRowCounts operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListDataVerificationTableRowCounts.go.html to see an example of how to use ListDataVerificationTableRowCountsRequest.
type ListDataVerificationTableRowCountsRequest struct {

	// The OCID of the migration
	MigrationId *string `mandatory:"true" contributesTo:"path" name:"migrationId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// A filter to return only results for a specific owner.
	Owner *string `mandatory:"false" contributesTo:"query" name:"owner"`

	// A filter to return only results for a specific table.
	TableName *string `mandatory:"false" contributesTo:"query" name:"tableName"`

	// Free-text filter applied by the service to relevant fields for the report.
	Filter *string `mandatory:"false" contributesTo:"query" name:"filter"`

	// Minimum absolute deltaPercent threshold (magnitude) to return.
	// The service filters results where `abs(deltaPercent) >= minAbsDeltaPercent`.
	// Must be non-negative.
	// Note: This filter applies to the object type counts report, which uses `deltaPercent`.
	// Table row count reports use `variancePercent` instead.
	MinAbsDeltaPercent *float64 `mandatory:"false" contributesTo:"query" name:"minAbsDeltaPercent"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for variancePercent is descending.
	SortBy ListDataVerificationTableRowCountsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataVerificationTableRowCountsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataVerificationTableRowCountsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataVerificationTableRowCountsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataVerificationTableRowCountsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataVerificationTableRowCountsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataVerificationTableRowCountsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataVerificationTableRowCountsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataVerificationTableRowCountsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataVerificationTableRowCountsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataVerificationTableRowCountsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataVerificationTableRowCountsResponse wrapper for the ListDataVerificationTableRowCounts operation
type ListDataVerificationTableRowCountsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataVerificationTableRowCountCollection instances
	DataVerificationTableRowCountCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataVerificationTableRowCountsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataVerificationTableRowCountsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataVerificationTableRowCountsSortByEnum Enum with underlying type: string
type ListDataVerificationTableRowCountsSortByEnum string

// Set of constants representing the allowable values for ListDataVerificationTableRowCountsSortByEnum
const (
	ListDataVerificationTableRowCountsSortByVariancepercent                    ListDataVerificationTableRowCountsSortByEnum = "variancePercent"
	ListDataVerificationTableRowCountsSortByOwner                              ListDataVerificationTableRowCountsSortByEnum = "owner"
	ListDataVerificationTableRowCountsSortByTablename                          ListDataVerificationTableRowCountsSortByEnum = "tableName"
	ListDataVerificationTableRowCountsSortByTimelastsourcestatisticscollection ListDataVerificationTableRowCountsSortByEnum = "timeLastSourceStatisticsCollection"
	ListDataVerificationTableRowCountsSortByTimelasttargetstatisticscollection ListDataVerificationTableRowCountsSortByEnum = "timeLastTargetStatisticsCollection"
	ListDataVerificationTableRowCountsSortBySourcerowcount                     ListDataVerificationTableRowCountsSortByEnum = "sourceRowCount"
	ListDataVerificationTableRowCountsSortByTargetrowcount                     ListDataVerificationTableRowCountsSortByEnum = "targetRowCount"
)

var mappingListDataVerificationTableRowCountsSortByEnum = map[string]ListDataVerificationTableRowCountsSortByEnum{
	"variancePercent":                    ListDataVerificationTableRowCountsSortByVariancepercent,
	"owner":                              ListDataVerificationTableRowCountsSortByOwner,
	"tableName":                          ListDataVerificationTableRowCountsSortByTablename,
	"timeLastSourceStatisticsCollection": ListDataVerificationTableRowCountsSortByTimelastsourcestatisticscollection,
	"timeLastTargetStatisticsCollection": ListDataVerificationTableRowCountsSortByTimelasttargetstatisticscollection,
	"sourceRowCount":                     ListDataVerificationTableRowCountsSortBySourcerowcount,
	"targetRowCount":                     ListDataVerificationTableRowCountsSortByTargetrowcount,
}

var mappingListDataVerificationTableRowCountsSortByEnumLowerCase = map[string]ListDataVerificationTableRowCountsSortByEnum{
	"variancepercent":                    ListDataVerificationTableRowCountsSortByVariancepercent,
	"owner":                              ListDataVerificationTableRowCountsSortByOwner,
	"tablename":                          ListDataVerificationTableRowCountsSortByTablename,
	"timelastsourcestatisticscollection": ListDataVerificationTableRowCountsSortByTimelastsourcestatisticscollection,
	"timelasttargetstatisticscollection": ListDataVerificationTableRowCountsSortByTimelasttargetstatisticscollection,
	"sourcerowcount":                     ListDataVerificationTableRowCountsSortBySourcerowcount,
	"targetrowcount":                     ListDataVerificationTableRowCountsSortByTargetrowcount,
}

// GetListDataVerificationTableRowCountsSortByEnumValues Enumerates the set of values for ListDataVerificationTableRowCountsSortByEnum
func GetListDataVerificationTableRowCountsSortByEnumValues() []ListDataVerificationTableRowCountsSortByEnum {
	values := make([]ListDataVerificationTableRowCountsSortByEnum, 0)
	for _, v := range mappingListDataVerificationTableRowCountsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataVerificationTableRowCountsSortByEnumStringValues Enumerates the set of values in String for ListDataVerificationTableRowCountsSortByEnum
func GetListDataVerificationTableRowCountsSortByEnumStringValues() []string {
	return []string{
		"variancePercent",
		"owner",
		"tableName",
		"timeLastSourceStatisticsCollection",
		"timeLastTargetStatisticsCollection",
		"sourceRowCount",
		"targetRowCount",
	}
}

// GetMappingListDataVerificationTableRowCountsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataVerificationTableRowCountsSortByEnum(val string) (ListDataVerificationTableRowCountsSortByEnum, bool) {
	enum, ok := mappingListDataVerificationTableRowCountsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataVerificationTableRowCountsSortOrderEnum Enum with underlying type: string
type ListDataVerificationTableRowCountsSortOrderEnum string

// Set of constants representing the allowable values for ListDataVerificationTableRowCountsSortOrderEnum
const (
	ListDataVerificationTableRowCountsSortOrderAsc  ListDataVerificationTableRowCountsSortOrderEnum = "ASC"
	ListDataVerificationTableRowCountsSortOrderDesc ListDataVerificationTableRowCountsSortOrderEnum = "DESC"
)

var mappingListDataVerificationTableRowCountsSortOrderEnum = map[string]ListDataVerificationTableRowCountsSortOrderEnum{
	"ASC":  ListDataVerificationTableRowCountsSortOrderAsc,
	"DESC": ListDataVerificationTableRowCountsSortOrderDesc,
}

var mappingListDataVerificationTableRowCountsSortOrderEnumLowerCase = map[string]ListDataVerificationTableRowCountsSortOrderEnum{
	"asc":  ListDataVerificationTableRowCountsSortOrderAsc,
	"desc": ListDataVerificationTableRowCountsSortOrderDesc,
}

// GetListDataVerificationTableRowCountsSortOrderEnumValues Enumerates the set of values for ListDataVerificationTableRowCountsSortOrderEnum
func GetListDataVerificationTableRowCountsSortOrderEnumValues() []ListDataVerificationTableRowCountsSortOrderEnum {
	values := make([]ListDataVerificationTableRowCountsSortOrderEnum, 0)
	for _, v := range mappingListDataVerificationTableRowCountsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataVerificationTableRowCountsSortOrderEnumStringValues Enumerates the set of values in String for ListDataVerificationTableRowCountsSortOrderEnum
func GetListDataVerificationTableRowCountsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataVerificationTableRowCountsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataVerificationTableRowCountsSortOrderEnum(val string) (ListDataVerificationTableRowCountsSortOrderEnum, bool) {
	enum, ok := mappingListDataVerificationTableRowCountsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
