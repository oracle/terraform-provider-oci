// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListSqlTuningAdvisorTaskFindingsRequest wrapper for the ListSqlTuningAdvisorTaskFindings operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListSqlTuningAdvisorTaskFindings.go.html to see an example of how to use ListSqlTuningAdvisorTaskFindingsRequest.
type ListSqlTuningAdvisorTaskFindingsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The SQL tuning task identifier. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskId *int64 `mandatory:"true" contributesTo:"path" name:"sqlTuningAdvisorTaskId"`

	// The optional greater than or equal to filter on the execution ID related to a specific SQL Tuning Advisor task.
	BeginExecId *int64 `mandatory:"false" contributesTo:"query" name:"beginExecId"`

	// The optional less than or equal to query parameter to filter on the execution ID related to a specific SQL Tuning Advisor task.
	EndExecId *int64 `mandatory:"false" contributesTo:"query" name:"endExecId"`

	// The search period during which the API will search for begin and end exec id, if not supplied.
	// Unused if beginExecId and endExecId optional query params are both supplied.
	SearchPeriod ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum `mandatory:"false" contributesTo:"query" name:"searchPeriod" omitEmpty:"true"`

	// The filter used to display specific findings in the report.
	FindingFilter ListSqlTuningAdvisorTaskFindingsFindingFilterEnum `mandatory:"false" contributesTo:"query" name:"findingFilter" omitEmpty:"true"`

	// The hash value of the object for the statistic finding search.
	StatsHashFilter *string `mandatory:"false" contributesTo:"query" name:"statsHashFilter"`

	// The hash value of the index table name.
	IndexHashFilter *string `mandatory:"false" contributesTo:"query" name:"indexHashFilter"`

	// The possible sortBy values of an object's recommendations.
	SortBy ListSqlTuningAdvisorTaskFindingsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Descending order is the default order.
	SortOrder ListSqlTuningAdvisorTaskFindingsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlTuningAdvisorTaskFindingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlTuningAdvisorTaskFindingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlTuningAdvisorTaskFindingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlTuningAdvisorTaskFindingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlTuningAdvisorTaskFindingsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlTuningAdvisorTaskFindingsSearchPeriodEnum(string(request.SearchPeriod)); !ok && request.SearchPeriod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SearchPeriod: %s. Supported values are: %s.", request.SearchPeriod, strings.Join(GetListSqlTuningAdvisorTaskFindingsSearchPeriodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlTuningAdvisorTaskFindingsFindingFilterEnum(string(request.FindingFilter)); !ok && request.FindingFilter != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FindingFilter: %s. Supported values are: %s.", request.FindingFilter, strings.Join(GetListSqlTuningAdvisorTaskFindingsFindingFilterEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlTuningAdvisorTaskFindingsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSqlTuningAdvisorTaskFindingsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlTuningAdvisorTaskFindingsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSqlTuningAdvisorTaskFindingsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlTuningAdvisorTaskFindingsResponse wrapper for the ListSqlTuningAdvisorTaskFindings operation
type ListSqlTuningAdvisorTaskFindingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlTuningAdvisorTaskFindingCollection instances
	SqlTuningAdvisorTaskFindingCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSqlTuningAdvisorTaskFindingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlTuningAdvisorTaskFindingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum Enum with underlying type: string
type ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum string

// Set of constants representing the allowable values for ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum
const (
	ListSqlTuningAdvisorTaskFindingsSearchPeriodLast24hr  ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum = "LAST_24HR"
	ListSqlTuningAdvisorTaskFindingsSearchPeriodLast7day  ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum = "LAST_7DAY"
	ListSqlTuningAdvisorTaskFindingsSearchPeriodLast31day ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum = "LAST_31DAY"
	ListSqlTuningAdvisorTaskFindingsSearchPeriodSinceLast ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum = "SINCE_LAST"
	ListSqlTuningAdvisorTaskFindingsSearchPeriodAll       ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum = "ALL"
)

var mappingListSqlTuningAdvisorTaskFindingsSearchPeriodEnum = map[string]ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum{
	"LAST_24HR":  ListSqlTuningAdvisorTaskFindingsSearchPeriodLast24hr,
	"LAST_7DAY":  ListSqlTuningAdvisorTaskFindingsSearchPeriodLast7day,
	"LAST_31DAY": ListSqlTuningAdvisorTaskFindingsSearchPeriodLast31day,
	"SINCE_LAST": ListSqlTuningAdvisorTaskFindingsSearchPeriodSinceLast,
	"ALL":        ListSqlTuningAdvisorTaskFindingsSearchPeriodAll,
}

// GetListSqlTuningAdvisorTaskFindingsSearchPeriodEnumValues Enumerates the set of values for ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum
func GetListSqlTuningAdvisorTaskFindingsSearchPeriodEnumValues() []ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum {
	values := make([]ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum, 0)
	for _, v := range mappingListSqlTuningAdvisorTaskFindingsSearchPeriodEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlTuningAdvisorTaskFindingsSearchPeriodEnumStringValues Enumerates the set of values in String for ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum
func GetListSqlTuningAdvisorTaskFindingsSearchPeriodEnumStringValues() []string {
	return []string{
		"LAST_24HR",
		"LAST_7DAY",
		"LAST_31DAY",
		"SINCE_LAST",
		"ALL",
	}
}

// GetMappingListSqlTuningAdvisorTaskFindingsSearchPeriodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlTuningAdvisorTaskFindingsSearchPeriodEnum(val string) (ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum, bool) {
	mappingListSqlTuningAdvisorTaskFindingsSearchPeriodEnumIgnoreCase := make(map[string]ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum)
	for k, v := range mappingListSqlTuningAdvisorTaskFindingsSearchPeriodEnum {
		mappingListSqlTuningAdvisorTaskFindingsSearchPeriodEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSqlTuningAdvisorTaskFindingsSearchPeriodEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlTuningAdvisorTaskFindingsFindingFilterEnum Enum with underlying type: string
type ListSqlTuningAdvisorTaskFindingsFindingFilterEnum string

// Set of constants representing the allowable values for ListSqlTuningAdvisorTaskFindingsFindingFilterEnum
const (
	ListSqlTuningAdvisorTaskFindingsFindingFilterNone          ListSqlTuningAdvisorTaskFindingsFindingFilterEnum = "none"
	ListSqlTuningAdvisorTaskFindingsFindingFilterFindings      ListSqlTuningAdvisorTaskFindingsFindingFilterEnum = "FINDINGS"
	ListSqlTuningAdvisorTaskFindingsFindingFilterNofindings    ListSqlTuningAdvisorTaskFindingsFindingFilterEnum = "NOFINDINGS"
	ListSqlTuningAdvisorTaskFindingsFindingFilterErrors        ListSqlTuningAdvisorTaskFindingsFindingFilterEnum = "ERRORS"
	ListSqlTuningAdvisorTaskFindingsFindingFilterProfiles      ListSqlTuningAdvisorTaskFindingsFindingFilterEnum = "PROFILES"
	ListSqlTuningAdvisorTaskFindingsFindingFilterIndices       ListSqlTuningAdvisorTaskFindingsFindingFilterEnum = "INDICES"
	ListSqlTuningAdvisorTaskFindingsFindingFilterStats         ListSqlTuningAdvisorTaskFindingsFindingFilterEnum = "STATS"
	ListSqlTuningAdvisorTaskFindingsFindingFilterRestructure   ListSqlTuningAdvisorTaskFindingsFindingFilterEnum = "RESTRUCTURE"
	ListSqlTuningAdvisorTaskFindingsFindingFilterAlternative   ListSqlTuningAdvisorTaskFindingsFindingFilterEnum = "ALTERNATIVE"
	ListSqlTuningAdvisorTaskFindingsFindingFilterAutoProfiles  ListSqlTuningAdvisorTaskFindingsFindingFilterEnum = "AUTO_PROFILES"
	ListSqlTuningAdvisorTaskFindingsFindingFilterOtherProfiles ListSqlTuningAdvisorTaskFindingsFindingFilterEnum = "OTHER_PROFILES"
)

var mappingListSqlTuningAdvisorTaskFindingsFindingFilterEnum = map[string]ListSqlTuningAdvisorTaskFindingsFindingFilterEnum{
	"none":           ListSqlTuningAdvisorTaskFindingsFindingFilterNone,
	"FINDINGS":       ListSqlTuningAdvisorTaskFindingsFindingFilterFindings,
	"NOFINDINGS":     ListSqlTuningAdvisorTaskFindingsFindingFilterNofindings,
	"ERRORS":         ListSqlTuningAdvisorTaskFindingsFindingFilterErrors,
	"PROFILES":       ListSqlTuningAdvisorTaskFindingsFindingFilterProfiles,
	"INDICES":        ListSqlTuningAdvisorTaskFindingsFindingFilterIndices,
	"STATS":          ListSqlTuningAdvisorTaskFindingsFindingFilterStats,
	"RESTRUCTURE":    ListSqlTuningAdvisorTaskFindingsFindingFilterRestructure,
	"ALTERNATIVE":    ListSqlTuningAdvisorTaskFindingsFindingFilterAlternative,
	"AUTO_PROFILES":  ListSqlTuningAdvisorTaskFindingsFindingFilterAutoProfiles,
	"OTHER_PROFILES": ListSqlTuningAdvisorTaskFindingsFindingFilterOtherProfiles,
}

// GetListSqlTuningAdvisorTaskFindingsFindingFilterEnumValues Enumerates the set of values for ListSqlTuningAdvisorTaskFindingsFindingFilterEnum
func GetListSqlTuningAdvisorTaskFindingsFindingFilterEnumValues() []ListSqlTuningAdvisorTaskFindingsFindingFilterEnum {
	values := make([]ListSqlTuningAdvisorTaskFindingsFindingFilterEnum, 0)
	for _, v := range mappingListSqlTuningAdvisorTaskFindingsFindingFilterEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlTuningAdvisorTaskFindingsFindingFilterEnumStringValues Enumerates the set of values in String for ListSqlTuningAdvisorTaskFindingsFindingFilterEnum
func GetListSqlTuningAdvisorTaskFindingsFindingFilterEnumStringValues() []string {
	return []string{
		"none",
		"FINDINGS",
		"NOFINDINGS",
		"ERRORS",
		"PROFILES",
		"INDICES",
		"STATS",
		"RESTRUCTURE",
		"ALTERNATIVE",
		"AUTO_PROFILES",
		"OTHER_PROFILES",
	}
}

// GetMappingListSqlTuningAdvisorTaskFindingsFindingFilterEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlTuningAdvisorTaskFindingsFindingFilterEnum(val string) (ListSqlTuningAdvisorTaskFindingsFindingFilterEnum, bool) {
	mappingListSqlTuningAdvisorTaskFindingsFindingFilterEnumIgnoreCase := make(map[string]ListSqlTuningAdvisorTaskFindingsFindingFilterEnum)
	for k, v := range mappingListSqlTuningAdvisorTaskFindingsFindingFilterEnum {
		mappingListSqlTuningAdvisorTaskFindingsFindingFilterEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSqlTuningAdvisorTaskFindingsFindingFilterEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlTuningAdvisorTaskFindingsSortByEnum Enum with underlying type: string
type ListSqlTuningAdvisorTaskFindingsSortByEnum string

// Set of constants representing the allowable values for ListSqlTuningAdvisorTaskFindingsSortByEnum
const (
	ListSqlTuningAdvisorTaskFindingsSortByDbtimeBenefit ListSqlTuningAdvisorTaskFindingsSortByEnum = "DBTIME_BENEFIT"
	ListSqlTuningAdvisorTaskFindingsSortByParsingSchema ListSqlTuningAdvisorTaskFindingsSortByEnum = "PARSING_SCHEMA"
	ListSqlTuningAdvisorTaskFindingsSortBySqlId         ListSqlTuningAdvisorTaskFindingsSortByEnum = "SQL_ID"
	ListSqlTuningAdvisorTaskFindingsSortByStats         ListSqlTuningAdvisorTaskFindingsSortByEnum = "STATS"
	ListSqlTuningAdvisorTaskFindingsSortByProfiles      ListSqlTuningAdvisorTaskFindingsSortByEnum = "PROFILES"
	ListSqlTuningAdvisorTaskFindingsSortBySqlBenefit    ListSqlTuningAdvisorTaskFindingsSortByEnum = "SQL_BENEFIT"
	ListSqlTuningAdvisorTaskFindingsSortByDate          ListSqlTuningAdvisorTaskFindingsSortByEnum = "DATE"
	ListSqlTuningAdvisorTaskFindingsSortByIndices       ListSqlTuningAdvisorTaskFindingsSortByEnum = "INDICES"
	ListSqlTuningAdvisorTaskFindingsSortByRestructure   ListSqlTuningAdvisorTaskFindingsSortByEnum = "RESTRUCTURE"
	ListSqlTuningAdvisorTaskFindingsSortByAlternative   ListSqlTuningAdvisorTaskFindingsSortByEnum = "ALTERNATIVE"
	ListSqlTuningAdvisorTaskFindingsSortByMisc          ListSqlTuningAdvisorTaskFindingsSortByEnum = "MISC"
	ListSqlTuningAdvisorTaskFindingsSortByError         ListSqlTuningAdvisorTaskFindingsSortByEnum = "ERROR"
	ListSqlTuningAdvisorTaskFindingsSortByTimeouts      ListSqlTuningAdvisorTaskFindingsSortByEnum = "TIMEOUTS"
)

var mappingListSqlTuningAdvisorTaskFindingsSortByEnum = map[string]ListSqlTuningAdvisorTaskFindingsSortByEnum{
	"DBTIME_BENEFIT": ListSqlTuningAdvisorTaskFindingsSortByDbtimeBenefit,
	"PARSING_SCHEMA": ListSqlTuningAdvisorTaskFindingsSortByParsingSchema,
	"SQL_ID":         ListSqlTuningAdvisorTaskFindingsSortBySqlId,
	"STATS":          ListSqlTuningAdvisorTaskFindingsSortByStats,
	"PROFILES":       ListSqlTuningAdvisorTaskFindingsSortByProfiles,
	"SQL_BENEFIT":    ListSqlTuningAdvisorTaskFindingsSortBySqlBenefit,
	"DATE":           ListSqlTuningAdvisorTaskFindingsSortByDate,
	"INDICES":        ListSqlTuningAdvisorTaskFindingsSortByIndices,
	"RESTRUCTURE":    ListSqlTuningAdvisorTaskFindingsSortByRestructure,
	"ALTERNATIVE":    ListSqlTuningAdvisorTaskFindingsSortByAlternative,
	"MISC":           ListSqlTuningAdvisorTaskFindingsSortByMisc,
	"ERROR":          ListSqlTuningAdvisorTaskFindingsSortByError,
	"TIMEOUTS":       ListSqlTuningAdvisorTaskFindingsSortByTimeouts,
}

// GetListSqlTuningAdvisorTaskFindingsSortByEnumValues Enumerates the set of values for ListSqlTuningAdvisorTaskFindingsSortByEnum
func GetListSqlTuningAdvisorTaskFindingsSortByEnumValues() []ListSqlTuningAdvisorTaskFindingsSortByEnum {
	values := make([]ListSqlTuningAdvisorTaskFindingsSortByEnum, 0)
	for _, v := range mappingListSqlTuningAdvisorTaskFindingsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlTuningAdvisorTaskFindingsSortByEnumStringValues Enumerates the set of values in String for ListSqlTuningAdvisorTaskFindingsSortByEnum
func GetListSqlTuningAdvisorTaskFindingsSortByEnumStringValues() []string {
	return []string{
		"DBTIME_BENEFIT",
		"PARSING_SCHEMA",
		"SQL_ID",
		"STATS",
		"PROFILES",
		"SQL_BENEFIT",
		"DATE",
		"INDICES",
		"RESTRUCTURE",
		"ALTERNATIVE",
		"MISC",
		"ERROR",
		"TIMEOUTS",
	}
}

// GetMappingListSqlTuningAdvisorTaskFindingsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlTuningAdvisorTaskFindingsSortByEnum(val string) (ListSqlTuningAdvisorTaskFindingsSortByEnum, bool) {
	mappingListSqlTuningAdvisorTaskFindingsSortByEnumIgnoreCase := make(map[string]ListSqlTuningAdvisorTaskFindingsSortByEnum)
	for k, v := range mappingListSqlTuningAdvisorTaskFindingsSortByEnum {
		mappingListSqlTuningAdvisorTaskFindingsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSqlTuningAdvisorTaskFindingsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlTuningAdvisorTaskFindingsSortOrderEnum Enum with underlying type: string
type ListSqlTuningAdvisorTaskFindingsSortOrderEnum string

// Set of constants representing the allowable values for ListSqlTuningAdvisorTaskFindingsSortOrderEnum
const (
	ListSqlTuningAdvisorTaskFindingsSortOrderAsc  ListSqlTuningAdvisorTaskFindingsSortOrderEnum = "ASC"
	ListSqlTuningAdvisorTaskFindingsSortOrderDesc ListSqlTuningAdvisorTaskFindingsSortOrderEnum = "DESC"
)

var mappingListSqlTuningAdvisorTaskFindingsSortOrderEnum = map[string]ListSqlTuningAdvisorTaskFindingsSortOrderEnum{
	"ASC":  ListSqlTuningAdvisorTaskFindingsSortOrderAsc,
	"DESC": ListSqlTuningAdvisorTaskFindingsSortOrderDesc,
}

// GetListSqlTuningAdvisorTaskFindingsSortOrderEnumValues Enumerates the set of values for ListSqlTuningAdvisorTaskFindingsSortOrderEnum
func GetListSqlTuningAdvisorTaskFindingsSortOrderEnumValues() []ListSqlTuningAdvisorTaskFindingsSortOrderEnum {
	values := make([]ListSqlTuningAdvisorTaskFindingsSortOrderEnum, 0)
	for _, v := range mappingListSqlTuningAdvisorTaskFindingsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlTuningAdvisorTaskFindingsSortOrderEnumStringValues Enumerates the set of values in String for ListSqlTuningAdvisorTaskFindingsSortOrderEnum
func GetListSqlTuningAdvisorTaskFindingsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSqlTuningAdvisorTaskFindingsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlTuningAdvisorTaskFindingsSortOrderEnum(val string) (ListSqlTuningAdvisorTaskFindingsSortOrderEnum, bool) {
	mappingListSqlTuningAdvisorTaskFindingsSortOrderEnumIgnoreCase := make(map[string]ListSqlTuningAdvisorTaskFindingsSortOrderEnum)
	for k, v := range mappingListSqlTuningAdvisorTaskFindingsSortOrderEnum {
		mappingListSqlTuningAdvisorTaskFindingsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSqlTuningAdvisorTaskFindingsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
