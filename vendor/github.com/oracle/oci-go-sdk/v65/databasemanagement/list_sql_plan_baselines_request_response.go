// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSqlPlanBaselinesRequest wrapper for the ListSqlPlanBaselines operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListSqlPlanBaselines.go.html to see an example of how to use ListSqlPlanBaselinesRequest.
type ListSqlPlanBaselinesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// A filter to return only SQL plan baselines that match the plan name.
	PlanName *string `mandatory:"false" contributesTo:"query" name:"planName"`

	// A filter to return all the SQL plan baselines for the specified SQL handle.
	SqlHandle *string `mandatory:"false" contributesTo:"query" name:"sqlHandle"`

	// A filter to return all the SQL plan baselines that match the SQL text. By default, the search
	// is case insensitive. To run an exact or case-sensitive search, double-quote the search string.
	// You may also use the '%' symbol as a wildcard.
	SqlText *string `mandatory:"false" contributesTo:"query" name:"sqlText"`

	// A filter to return only SQL plan baselines that are either enabled or not enabled.
	// By default, all SQL plan baselines are returned.
	IsEnabled *bool `mandatory:"false" contributesTo:"query" name:"isEnabled"`

	// A filter to return only SQL plan baselines that are either accepted or not accepted.
	// By default, all SQL plan baselines are returned.
	IsAccepted *bool `mandatory:"false" contributesTo:"query" name:"isAccepted"`

	// A filter to return only SQL plan baselines that were either reproduced or
	// not reproduced by the optimizer. By default, all SQL plan baselines are returned.
	IsReproduced *bool `mandatory:"false" contributesTo:"query" name:"isReproduced"`

	// A filter to return only SQL plan baselines that are either fixed or not fixed.
	// By default, all SQL plan baselines are returned.
	IsFixed *bool `mandatory:"false" contributesTo:"query" name:"isFixed"`

	// A filter to return only SQL plan baselines that are either adaptive or not adaptive.
	// By default, all SQL plan baselines are returned.
	IsAdaptive *bool `mandatory:"false" contributesTo:"query" name:"isAdaptive"`

	// A filter to return all the SQL plan baselines that match the origin.
	Origin ListSqlPlanBaselinesOriginEnum `mandatory:"false" contributesTo:"query" name:"origin" omitEmpty:"true"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort the SQL plan baseline summary data.
	SortBy ListSqlPlanBaselinesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Descending order is the default order.
	SortOrder ListSqlPlanBaselinesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlPlanBaselinesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlPlanBaselinesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlPlanBaselinesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlPlanBaselinesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlPlanBaselinesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlPlanBaselinesOriginEnum(string(request.Origin)); !ok && request.Origin != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Origin: %s. Supported values are: %s.", request.Origin, strings.Join(GetListSqlPlanBaselinesOriginEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlPlanBaselinesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSqlPlanBaselinesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlPlanBaselinesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSqlPlanBaselinesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlPlanBaselinesResponse wrapper for the ListSqlPlanBaselines operation
type ListSqlPlanBaselinesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlPlanBaselineCollection instances
	SqlPlanBaselineCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSqlPlanBaselinesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlPlanBaselinesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlPlanBaselinesOriginEnum Enum with underlying type: string
type ListSqlPlanBaselinesOriginEnum string

// Set of constants representing the allowable values for ListSqlPlanBaselinesOriginEnum
const (
	ListSqlPlanBaselinesOriginAddmSqltune               ListSqlPlanBaselinesOriginEnum = "ADDM_SQLTUNE"
	ListSqlPlanBaselinesOriginAutoCapture               ListSqlPlanBaselinesOriginEnum = "AUTO_CAPTURE"
	ListSqlPlanBaselinesOriginAutoSqltune               ListSqlPlanBaselinesOriginEnum = "AUTO_SQLTUNE"
	ListSqlPlanBaselinesOriginEvolveAutoIndexLoad       ListSqlPlanBaselinesOriginEnum = "EVOLVE_AUTO_INDEX_LOAD"
	ListSqlPlanBaselinesOriginEvolveCreateFromAdaptive  ListSqlPlanBaselinesOriginEnum = "EVOLVE_CREATE_FROM_ADAPTIVE"
	ListSqlPlanBaselinesOriginEvolveLoadFromSts         ListSqlPlanBaselinesOriginEnum = "EVOLVE_LOAD_FROM_STS"
	ListSqlPlanBaselinesOriginEvolveLoadFromAwr         ListSqlPlanBaselinesOriginEnum = "EVOLVE_LOAD_FROM_AWR"
	ListSqlPlanBaselinesOriginEvolveLoadFromCursorCache ListSqlPlanBaselinesOriginEnum = "EVOLVE_LOAD_FROM_CURSOR_CACHE"
	ListSqlPlanBaselinesOriginManualLoad                ListSqlPlanBaselinesOriginEnum = "MANUAL_LOAD"
	ListSqlPlanBaselinesOriginManualLoadFromAwr         ListSqlPlanBaselinesOriginEnum = "MANUAL_LOAD_FROM_AWR"
	ListSqlPlanBaselinesOriginManualLoadFromCursorCache ListSqlPlanBaselinesOriginEnum = "MANUAL_LOAD_FROM_CURSOR_CACHE"
	ListSqlPlanBaselinesOriginManualLoadFromSts         ListSqlPlanBaselinesOriginEnum = "MANUAL_LOAD_FROM_STS"
	ListSqlPlanBaselinesOriginManualSqltune             ListSqlPlanBaselinesOriginEnum = "MANUAL_SQLTUNE"
	ListSqlPlanBaselinesOriginStoredOutline             ListSqlPlanBaselinesOriginEnum = "STORED_OUTLINE"
	ListSqlPlanBaselinesOriginUnknown                   ListSqlPlanBaselinesOriginEnum = "UNKNOWN"
)

var mappingListSqlPlanBaselinesOriginEnum = map[string]ListSqlPlanBaselinesOriginEnum{
	"ADDM_SQLTUNE":                  ListSqlPlanBaselinesOriginAddmSqltune,
	"AUTO_CAPTURE":                  ListSqlPlanBaselinesOriginAutoCapture,
	"AUTO_SQLTUNE":                  ListSqlPlanBaselinesOriginAutoSqltune,
	"EVOLVE_AUTO_INDEX_LOAD":        ListSqlPlanBaselinesOriginEvolveAutoIndexLoad,
	"EVOLVE_CREATE_FROM_ADAPTIVE":   ListSqlPlanBaselinesOriginEvolveCreateFromAdaptive,
	"EVOLVE_LOAD_FROM_STS":          ListSqlPlanBaselinesOriginEvolveLoadFromSts,
	"EVOLVE_LOAD_FROM_AWR":          ListSqlPlanBaselinesOriginEvolveLoadFromAwr,
	"EVOLVE_LOAD_FROM_CURSOR_CACHE": ListSqlPlanBaselinesOriginEvolveLoadFromCursorCache,
	"MANUAL_LOAD":                   ListSqlPlanBaselinesOriginManualLoad,
	"MANUAL_LOAD_FROM_AWR":          ListSqlPlanBaselinesOriginManualLoadFromAwr,
	"MANUAL_LOAD_FROM_CURSOR_CACHE": ListSqlPlanBaselinesOriginManualLoadFromCursorCache,
	"MANUAL_LOAD_FROM_STS":          ListSqlPlanBaselinesOriginManualLoadFromSts,
	"MANUAL_SQLTUNE":                ListSqlPlanBaselinesOriginManualSqltune,
	"STORED_OUTLINE":                ListSqlPlanBaselinesOriginStoredOutline,
	"UNKNOWN":                       ListSqlPlanBaselinesOriginUnknown,
}

var mappingListSqlPlanBaselinesOriginEnumLowerCase = map[string]ListSqlPlanBaselinesOriginEnum{
	"addm_sqltune":                  ListSqlPlanBaselinesOriginAddmSqltune,
	"auto_capture":                  ListSqlPlanBaselinesOriginAutoCapture,
	"auto_sqltune":                  ListSqlPlanBaselinesOriginAutoSqltune,
	"evolve_auto_index_load":        ListSqlPlanBaselinesOriginEvolveAutoIndexLoad,
	"evolve_create_from_adaptive":   ListSqlPlanBaselinesOriginEvolveCreateFromAdaptive,
	"evolve_load_from_sts":          ListSqlPlanBaselinesOriginEvolveLoadFromSts,
	"evolve_load_from_awr":          ListSqlPlanBaselinesOriginEvolveLoadFromAwr,
	"evolve_load_from_cursor_cache": ListSqlPlanBaselinesOriginEvolveLoadFromCursorCache,
	"manual_load":                   ListSqlPlanBaselinesOriginManualLoad,
	"manual_load_from_awr":          ListSqlPlanBaselinesOriginManualLoadFromAwr,
	"manual_load_from_cursor_cache": ListSqlPlanBaselinesOriginManualLoadFromCursorCache,
	"manual_load_from_sts":          ListSqlPlanBaselinesOriginManualLoadFromSts,
	"manual_sqltune":                ListSqlPlanBaselinesOriginManualSqltune,
	"stored_outline":                ListSqlPlanBaselinesOriginStoredOutline,
	"unknown":                       ListSqlPlanBaselinesOriginUnknown,
}

// GetListSqlPlanBaselinesOriginEnumValues Enumerates the set of values for ListSqlPlanBaselinesOriginEnum
func GetListSqlPlanBaselinesOriginEnumValues() []ListSqlPlanBaselinesOriginEnum {
	values := make([]ListSqlPlanBaselinesOriginEnum, 0)
	for _, v := range mappingListSqlPlanBaselinesOriginEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlPlanBaselinesOriginEnumStringValues Enumerates the set of values in String for ListSqlPlanBaselinesOriginEnum
func GetListSqlPlanBaselinesOriginEnumStringValues() []string {
	return []string{
		"ADDM_SQLTUNE",
		"AUTO_CAPTURE",
		"AUTO_SQLTUNE",
		"EVOLVE_AUTO_INDEX_LOAD",
		"EVOLVE_CREATE_FROM_ADAPTIVE",
		"EVOLVE_LOAD_FROM_STS",
		"EVOLVE_LOAD_FROM_AWR",
		"EVOLVE_LOAD_FROM_CURSOR_CACHE",
		"MANUAL_LOAD",
		"MANUAL_LOAD_FROM_AWR",
		"MANUAL_LOAD_FROM_CURSOR_CACHE",
		"MANUAL_LOAD_FROM_STS",
		"MANUAL_SQLTUNE",
		"STORED_OUTLINE",
		"UNKNOWN",
	}
}

// GetMappingListSqlPlanBaselinesOriginEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlPlanBaselinesOriginEnum(val string) (ListSqlPlanBaselinesOriginEnum, bool) {
	enum, ok := mappingListSqlPlanBaselinesOriginEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlPlanBaselinesSortByEnum Enum with underlying type: string
type ListSqlPlanBaselinesSortByEnum string

// Set of constants representing the allowable values for ListSqlPlanBaselinesSortByEnum
const (
	ListSqlPlanBaselinesSortByTimecreated      ListSqlPlanBaselinesSortByEnum = "timeCreated"
	ListSqlPlanBaselinesSortByTimelastmodified ListSqlPlanBaselinesSortByEnum = "timeLastModified"
)

var mappingListSqlPlanBaselinesSortByEnum = map[string]ListSqlPlanBaselinesSortByEnum{
	"timeCreated":      ListSqlPlanBaselinesSortByTimecreated,
	"timeLastModified": ListSqlPlanBaselinesSortByTimelastmodified,
}

var mappingListSqlPlanBaselinesSortByEnumLowerCase = map[string]ListSqlPlanBaselinesSortByEnum{
	"timecreated":      ListSqlPlanBaselinesSortByTimecreated,
	"timelastmodified": ListSqlPlanBaselinesSortByTimelastmodified,
}

// GetListSqlPlanBaselinesSortByEnumValues Enumerates the set of values for ListSqlPlanBaselinesSortByEnum
func GetListSqlPlanBaselinesSortByEnumValues() []ListSqlPlanBaselinesSortByEnum {
	values := make([]ListSqlPlanBaselinesSortByEnum, 0)
	for _, v := range mappingListSqlPlanBaselinesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlPlanBaselinesSortByEnumStringValues Enumerates the set of values in String for ListSqlPlanBaselinesSortByEnum
func GetListSqlPlanBaselinesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeLastModified",
	}
}

// GetMappingListSqlPlanBaselinesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlPlanBaselinesSortByEnum(val string) (ListSqlPlanBaselinesSortByEnum, bool) {
	enum, ok := mappingListSqlPlanBaselinesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlPlanBaselinesSortOrderEnum Enum with underlying type: string
type ListSqlPlanBaselinesSortOrderEnum string

// Set of constants representing the allowable values for ListSqlPlanBaselinesSortOrderEnum
const (
	ListSqlPlanBaselinesSortOrderAsc  ListSqlPlanBaselinesSortOrderEnum = "ASC"
	ListSqlPlanBaselinesSortOrderDesc ListSqlPlanBaselinesSortOrderEnum = "DESC"
)

var mappingListSqlPlanBaselinesSortOrderEnum = map[string]ListSqlPlanBaselinesSortOrderEnum{
	"ASC":  ListSqlPlanBaselinesSortOrderAsc,
	"DESC": ListSqlPlanBaselinesSortOrderDesc,
}

var mappingListSqlPlanBaselinesSortOrderEnumLowerCase = map[string]ListSqlPlanBaselinesSortOrderEnum{
	"asc":  ListSqlPlanBaselinesSortOrderAsc,
	"desc": ListSqlPlanBaselinesSortOrderDesc,
}

// GetListSqlPlanBaselinesSortOrderEnumValues Enumerates the set of values for ListSqlPlanBaselinesSortOrderEnum
func GetListSqlPlanBaselinesSortOrderEnumValues() []ListSqlPlanBaselinesSortOrderEnum {
	values := make([]ListSqlPlanBaselinesSortOrderEnum, 0)
	for _, v := range mappingListSqlPlanBaselinesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlPlanBaselinesSortOrderEnumStringValues Enumerates the set of values in String for ListSqlPlanBaselinesSortOrderEnum
func GetListSqlPlanBaselinesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSqlPlanBaselinesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlPlanBaselinesSortOrderEnum(val string) (ListSqlPlanBaselinesSortOrderEnum, bool) {
	enum, ok := mappingListSqlPlanBaselinesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
