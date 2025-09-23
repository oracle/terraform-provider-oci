// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMaskingErrorsRequest wrapper for the ListMaskingErrors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingErrors.go.html to see an example of how to use ListMaskingErrorsRequest.
type ListMaskingErrorsRequest struct {

	// The OCID of the masking report.
	MaskingReportId *string `mandatory:"true" contributesTo:"path" name:"maskingReportId"`

	// A filter to return only masking errors that match the specified step name.
	StepName ListMaskingErrorsStepNameEnum `mandatory:"false" contributesTo:"query" name:"stepName" omitEmpty:"true"`

	// The field to sort by. The default order will be ascending.
	SortBy ListMaskingErrorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListMaskingErrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaskingErrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaskingErrorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaskingErrorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaskingErrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaskingErrorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMaskingErrorsStepNameEnum(string(request.StepName)); !ok && request.StepName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StepName: %s. Supported values are: %s.", request.StepName, strings.Join(GetListMaskingErrorsStepNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingErrorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaskingErrorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingErrorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaskingErrorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaskingErrorsResponse wrapper for the ListMaskingErrors operation
type ListMaskingErrorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaskingErrorCollection instances
	MaskingErrorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListMaskingErrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaskingErrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaskingErrorsStepNameEnum Enum with underlying type: string
type ListMaskingErrorsStepNameEnum string

// Set of constants representing the allowable values for ListMaskingErrorsStepNameEnum
const (
	ListMaskingErrorsStepNameValidate       ListMaskingErrorsStepNameEnum = "VALIDATE"
	ListMaskingErrorsStepNameGenerateScript ListMaskingErrorsStepNameEnum = "GENERATE_SCRIPT"
	ListMaskingErrorsStepNameExecuteMasking ListMaskingErrorsStepNameEnum = "EXECUTE_MASKING"
	ListMaskingErrorsStepNamePreMasking     ListMaskingErrorsStepNameEnum = "PRE_MASKING"
	ListMaskingErrorsStepNamePostMasking    ListMaskingErrorsStepNameEnum = "POST_MASKING"
)

var mappingListMaskingErrorsStepNameEnum = map[string]ListMaskingErrorsStepNameEnum{
	"VALIDATE":        ListMaskingErrorsStepNameValidate,
	"GENERATE_SCRIPT": ListMaskingErrorsStepNameGenerateScript,
	"EXECUTE_MASKING": ListMaskingErrorsStepNameExecuteMasking,
	"PRE_MASKING":     ListMaskingErrorsStepNamePreMasking,
	"POST_MASKING":    ListMaskingErrorsStepNamePostMasking,
}

var mappingListMaskingErrorsStepNameEnumLowerCase = map[string]ListMaskingErrorsStepNameEnum{
	"validate":        ListMaskingErrorsStepNameValidate,
	"generate_script": ListMaskingErrorsStepNameGenerateScript,
	"execute_masking": ListMaskingErrorsStepNameExecuteMasking,
	"pre_masking":     ListMaskingErrorsStepNamePreMasking,
	"post_masking":    ListMaskingErrorsStepNamePostMasking,
}

// GetListMaskingErrorsStepNameEnumValues Enumerates the set of values for ListMaskingErrorsStepNameEnum
func GetListMaskingErrorsStepNameEnumValues() []ListMaskingErrorsStepNameEnum {
	values := make([]ListMaskingErrorsStepNameEnum, 0)
	for _, v := range mappingListMaskingErrorsStepNameEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingErrorsStepNameEnumStringValues Enumerates the set of values in String for ListMaskingErrorsStepNameEnum
func GetListMaskingErrorsStepNameEnumStringValues() []string {
	return []string{
		"VALIDATE",
		"GENERATE_SCRIPT",
		"EXECUTE_MASKING",
		"PRE_MASKING",
		"POST_MASKING",
	}
}

// GetMappingListMaskingErrorsStepNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingErrorsStepNameEnum(val string) (ListMaskingErrorsStepNameEnum, bool) {
	enum, ok := mappingListMaskingErrorsStepNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingErrorsSortByEnum Enum with underlying type: string
type ListMaskingErrorsSortByEnum string

// Set of constants representing the allowable values for ListMaskingErrorsSortByEnum
const (
	ListMaskingErrorsSortByStepname    ListMaskingErrorsSortByEnum = "stepName"
	ListMaskingErrorsSortByTimecreated ListMaskingErrorsSortByEnum = "timeCreated"
)

var mappingListMaskingErrorsSortByEnum = map[string]ListMaskingErrorsSortByEnum{
	"stepName":    ListMaskingErrorsSortByStepname,
	"timeCreated": ListMaskingErrorsSortByTimecreated,
}

var mappingListMaskingErrorsSortByEnumLowerCase = map[string]ListMaskingErrorsSortByEnum{
	"stepname":    ListMaskingErrorsSortByStepname,
	"timecreated": ListMaskingErrorsSortByTimecreated,
}

// GetListMaskingErrorsSortByEnumValues Enumerates the set of values for ListMaskingErrorsSortByEnum
func GetListMaskingErrorsSortByEnumValues() []ListMaskingErrorsSortByEnum {
	values := make([]ListMaskingErrorsSortByEnum, 0)
	for _, v := range mappingListMaskingErrorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingErrorsSortByEnumStringValues Enumerates the set of values in String for ListMaskingErrorsSortByEnum
func GetListMaskingErrorsSortByEnumStringValues() []string {
	return []string{
		"stepName",
		"timeCreated",
	}
}

// GetMappingListMaskingErrorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingErrorsSortByEnum(val string) (ListMaskingErrorsSortByEnum, bool) {
	enum, ok := mappingListMaskingErrorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingErrorsSortOrderEnum Enum with underlying type: string
type ListMaskingErrorsSortOrderEnum string

// Set of constants representing the allowable values for ListMaskingErrorsSortOrderEnum
const (
	ListMaskingErrorsSortOrderAsc  ListMaskingErrorsSortOrderEnum = "ASC"
	ListMaskingErrorsSortOrderDesc ListMaskingErrorsSortOrderEnum = "DESC"
)

var mappingListMaskingErrorsSortOrderEnum = map[string]ListMaskingErrorsSortOrderEnum{
	"ASC":  ListMaskingErrorsSortOrderAsc,
	"DESC": ListMaskingErrorsSortOrderDesc,
}

var mappingListMaskingErrorsSortOrderEnumLowerCase = map[string]ListMaskingErrorsSortOrderEnum{
	"asc":  ListMaskingErrorsSortOrderAsc,
	"desc": ListMaskingErrorsSortOrderDesc,
}

// GetListMaskingErrorsSortOrderEnumValues Enumerates the set of values for ListMaskingErrorsSortOrderEnum
func GetListMaskingErrorsSortOrderEnumValues() []ListMaskingErrorsSortOrderEnum {
	values := make([]ListMaskingErrorsSortOrderEnum, 0)
	for _, v := range mappingListMaskingErrorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingErrorsSortOrderEnumStringValues Enumerates the set of values in String for ListMaskingErrorsSortOrderEnum
func GetListMaskingErrorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaskingErrorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingErrorsSortOrderEnum(val string) (ListMaskingErrorsSortOrderEnum, bool) {
	enum, ok := mappingListMaskingErrorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
