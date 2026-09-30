// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSubsettingErrorsRequest wrapper for the ListSubsettingErrors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSubsettingErrors.go.html to see an example of how to use ListSubsettingErrorsRequest.
type ListSubsettingErrorsRequest struct {

	// The OCID of the subsetting report.
	SubsettingReportId *string `mandatory:"true" contributesTo:"path" name:"subsettingReportId"`

	// A filter to return only subsetting errors that match the specified step name.
	StepName ListSubsettingErrorsStepNameEnum `mandatory:"false" contributesTo:"query" name:"stepName" omitEmpty:"true"`

	// The field to sort by. The default order will be ascending.
	SortBy ListSubsettingErrorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSubsettingErrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubsettingErrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubsettingErrorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubsettingErrorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubsettingErrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubsettingErrorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubsettingErrorsStepNameEnum(string(request.StepName)); !ok && request.StepName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StepName: %s. Supported values are: %s.", request.StepName, strings.Join(GetListSubsettingErrorsStepNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingErrorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubsettingErrorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingErrorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubsettingErrorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubsettingErrorsResponse wrapper for the ListSubsettingErrors operation
type ListSubsettingErrorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubsettingErrorCollection instances
	SubsettingErrorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSubsettingErrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubsettingErrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubsettingErrorsStepNameEnum Enum with underlying type: string
type ListSubsettingErrorsStepNameEnum string

// Set of constants representing the allowable values for ListSubsettingErrorsStepNameEnum
const (
	ListSubsettingErrorsStepNameValidate              ListSubsettingErrorsStepNameEnum = "VALIDATE"
	ListSubsettingErrorsStepNameIdentifyRowsForSubset ListSubsettingErrorsStepNameEnum = "IDENTIFY_ROWS_FOR_SUBSET"
	ListSubsettingErrorsStepNameGenerateScript        ListSubsettingErrorsStepNameEnum = "GENERATE_SCRIPT"
	ListSubsettingErrorsStepNameExecuteSubsetting     ListSubsettingErrorsStepNameEnum = "EXECUTE_SUBSETTING"
	ListSubsettingErrorsStepNamePreSubsetting         ListSubsettingErrorsStepNameEnum = "PRE_SUBSETTING"
	ListSubsettingErrorsStepNamePostSubsetting        ListSubsettingErrorsStepNameEnum = "POST_SUBSETTING"
)

var mappingListSubsettingErrorsStepNameEnum = map[string]ListSubsettingErrorsStepNameEnum{
	"VALIDATE":                 ListSubsettingErrorsStepNameValidate,
	"IDENTIFY_ROWS_FOR_SUBSET": ListSubsettingErrorsStepNameIdentifyRowsForSubset,
	"GENERATE_SCRIPT":          ListSubsettingErrorsStepNameGenerateScript,
	"EXECUTE_SUBSETTING":       ListSubsettingErrorsStepNameExecuteSubsetting,
	"PRE_SUBSETTING":           ListSubsettingErrorsStepNamePreSubsetting,
	"POST_SUBSETTING":          ListSubsettingErrorsStepNamePostSubsetting,
}

var mappingListSubsettingErrorsStepNameEnumLowerCase = map[string]ListSubsettingErrorsStepNameEnum{
	"validate":                 ListSubsettingErrorsStepNameValidate,
	"identify_rows_for_subset": ListSubsettingErrorsStepNameIdentifyRowsForSubset,
	"generate_script":          ListSubsettingErrorsStepNameGenerateScript,
	"execute_subsetting":       ListSubsettingErrorsStepNameExecuteSubsetting,
	"pre_subsetting":           ListSubsettingErrorsStepNamePreSubsetting,
	"post_subsetting":          ListSubsettingErrorsStepNamePostSubsetting,
}

// GetListSubsettingErrorsStepNameEnumValues Enumerates the set of values for ListSubsettingErrorsStepNameEnum
func GetListSubsettingErrorsStepNameEnumValues() []ListSubsettingErrorsStepNameEnum {
	values := make([]ListSubsettingErrorsStepNameEnum, 0)
	for _, v := range mappingListSubsettingErrorsStepNameEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingErrorsStepNameEnumStringValues Enumerates the set of values in String for ListSubsettingErrorsStepNameEnum
func GetListSubsettingErrorsStepNameEnumStringValues() []string {
	return []string{
		"VALIDATE",
		"IDENTIFY_ROWS_FOR_SUBSET",
		"GENERATE_SCRIPT",
		"EXECUTE_SUBSETTING",
		"PRE_SUBSETTING",
		"POST_SUBSETTING",
	}
}

// GetMappingListSubsettingErrorsStepNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingErrorsStepNameEnum(val string) (ListSubsettingErrorsStepNameEnum, bool) {
	enum, ok := mappingListSubsettingErrorsStepNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingErrorsSortByEnum Enum with underlying type: string
type ListSubsettingErrorsSortByEnum string

// Set of constants representing the allowable values for ListSubsettingErrorsSortByEnum
const (
	ListSubsettingErrorsSortByStepname    ListSubsettingErrorsSortByEnum = "stepName"
	ListSubsettingErrorsSortByTimecreated ListSubsettingErrorsSortByEnum = "timeCreated"
)

var mappingListSubsettingErrorsSortByEnum = map[string]ListSubsettingErrorsSortByEnum{
	"stepName":    ListSubsettingErrorsSortByStepname,
	"timeCreated": ListSubsettingErrorsSortByTimecreated,
}

var mappingListSubsettingErrorsSortByEnumLowerCase = map[string]ListSubsettingErrorsSortByEnum{
	"stepname":    ListSubsettingErrorsSortByStepname,
	"timecreated": ListSubsettingErrorsSortByTimecreated,
}

// GetListSubsettingErrorsSortByEnumValues Enumerates the set of values for ListSubsettingErrorsSortByEnum
func GetListSubsettingErrorsSortByEnumValues() []ListSubsettingErrorsSortByEnum {
	values := make([]ListSubsettingErrorsSortByEnum, 0)
	for _, v := range mappingListSubsettingErrorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingErrorsSortByEnumStringValues Enumerates the set of values in String for ListSubsettingErrorsSortByEnum
func GetListSubsettingErrorsSortByEnumStringValues() []string {
	return []string{
		"stepName",
		"timeCreated",
	}
}

// GetMappingListSubsettingErrorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingErrorsSortByEnum(val string) (ListSubsettingErrorsSortByEnum, bool) {
	enum, ok := mappingListSubsettingErrorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingErrorsSortOrderEnum Enum with underlying type: string
type ListSubsettingErrorsSortOrderEnum string

// Set of constants representing the allowable values for ListSubsettingErrorsSortOrderEnum
const (
	ListSubsettingErrorsSortOrderAsc  ListSubsettingErrorsSortOrderEnum = "ASC"
	ListSubsettingErrorsSortOrderDesc ListSubsettingErrorsSortOrderEnum = "DESC"
)

var mappingListSubsettingErrorsSortOrderEnum = map[string]ListSubsettingErrorsSortOrderEnum{
	"ASC":  ListSubsettingErrorsSortOrderAsc,
	"DESC": ListSubsettingErrorsSortOrderDesc,
}

var mappingListSubsettingErrorsSortOrderEnumLowerCase = map[string]ListSubsettingErrorsSortOrderEnum{
	"asc":  ListSubsettingErrorsSortOrderAsc,
	"desc": ListSubsettingErrorsSortOrderDesc,
}

// GetListSubsettingErrorsSortOrderEnumValues Enumerates the set of values for ListSubsettingErrorsSortOrderEnum
func GetListSubsettingErrorsSortOrderEnumValues() []ListSubsettingErrorsSortOrderEnum {
	values := make([]ListSubsettingErrorsSortOrderEnum, 0)
	for _, v := range mappingListSubsettingErrorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingErrorsSortOrderEnumStringValues Enumerates the set of values in String for ListSubsettingErrorsSortOrderEnum
func GetListSubsettingErrorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubsettingErrorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingErrorsSortOrderEnum(val string) (ListSubsettingErrorsSortOrderEnum, bool) {
	enum, ok := mappingListSubsettingErrorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
