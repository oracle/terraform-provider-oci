// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jmsutils

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPerformanceTuningAnalysisRequest wrapper for the ListPerformanceTuningAnalysis operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/ListPerformanceTuningAnalysis.go.html to see an example of how to use ListPerformanceTuningAnalysisRequest.
type ListPerformanceTuningAnalysisRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Performance Tuning Analysis.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The project name of the Performance Tuning Analysis to query for.
	AnalysisProjectName *string `mandatory:"false" contributesTo:"query" name:"analysisProjectName"`

	// The result of the Performance Tuning Analysis to query for.
	PerformanceTuningAnalysisResult PerformanceTuningAnalysisResultEnum `mandatory:"false" contributesTo:"query" name:"performanceTuningAnalysisResult" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListPerformanceTuningAnalysisSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort Performance Tuning Analysis. Only one sort order may be provided.
	SortBy ListPerformanceTuningAnalysisSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPerformanceTuningAnalysisRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPerformanceTuningAnalysisRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPerformanceTuningAnalysisRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPerformanceTuningAnalysisRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPerformanceTuningAnalysisRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPerformanceTuningAnalysisResultEnum(string(request.PerformanceTuningAnalysisResult)); !ok && request.PerformanceTuningAnalysisResult != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PerformanceTuningAnalysisResult: %s. Supported values are: %s.", request.PerformanceTuningAnalysisResult, strings.Join(GetPerformanceTuningAnalysisResultEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPerformanceTuningAnalysisSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPerformanceTuningAnalysisSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPerformanceTuningAnalysisSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPerformanceTuningAnalysisSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPerformanceTuningAnalysisResponse wrapper for the ListPerformanceTuningAnalysis operation
type ListPerformanceTuningAnalysisResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PerformanceTuningAnalysisCollection instances
	PerformanceTuningAnalysisCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPerformanceTuningAnalysisResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPerformanceTuningAnalysisResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPerformanceTuningAnalysisSortOrderEnum Enum with underlying type: string
type ListPerformanceTuningAnalysisSortOrderEnum string

// Set of constants representing the allowable values for ListPerformanceTuningAnalysisSortOrderEnum
const (
	ListPerformanceTuningAnalysisSortOrderAsc  ListPerformanceTuningAnalysisSortOrderEnum = "ASC"
	ListPerformanceTuningAnalysisSortOrderDesc ListPerformanceTuningAnalysisSortOrderEnum = "DESC"
)

var mappingListPerformanceTuningAnalysisSortOrderEnum = map[string]ListPerformanceTuningAnalysisSortOrderEnum{
	"ASC":  ListPerformanceTuningAnalysisSortOrderAsc,
	"DESC": ListPerformanceTuningAnalysisSortOrderDesc,
}

var mappingListPerformanceTuningAnalysisSortOrderEnumLowerCase = map[string]ListPerformanceTuningAnalysisSortOrderEnum{
	"asc":  ListPerformanceTuningAnalysisSortOrderAsc,
	"desc": ListPerformanceTuningAnalysisSortOrderDesc,
}

// GetListPerformanceTuningAnalysisSortOrderEnumValues Enumerates the set of values for ListPerformanceTuningAnalysisSortOrderEnum
func GetListPerformanceTuningAnalysisSortOrderEnumValues() []ListPerformanceTuningAnalysisSortOrderEnum {
	values := make([]ListPerformanceTuningAnalysisSortOrderEnum, 0)
	for _, v := range mappingListPerformanceTuningAnalysisSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPerformanceTuningAnalysisSortOrderEnumStringValues Enumerates the set of values in String for ListPerformanceTuningAnalysisSortOrderEnum
func GetListPerformanceTuningAnalysisSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPerformanceTuningAnalysisSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPerformanceTuningAnalysisSortOrderEnum(val string) (ListPerformanceTuningAnalysisSortOrderEnum, bool) {
	enum, ok := mappingListPerformanceTuningAnalysisSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPerformanceTuningAnalysisSortByEnum Enum with underlying type: string
type ListPerformanceTuningAnalysisSortByEnum string

// Set of constants representing the allowable values for ListPerformanceTuningAnalysisSortByEnum
const (
	ListPerformanceTuningAnalysisSortByCreated  ListPerformanceTuningAnalysisSortByEnum = "TIME_CREATED"
	ListPerformanceTuningAnalysisSortByStarted  ListPerformanceTuningAnalysisSortByEnum = "TIME_STARTED"
	ListPerformanceTuningAnalysisSortByFinished ListPerformanceTuningAnalysisSortByEnum = "TIME_FINISHED"
)

var mappingListPerformanceTuningAnalysisSortByEnum = map[string]ListPerformanceTuningAnalysisSortByEnum{
	"TIME_CREATED":  ListPerformanceTuningAnalysisSortByCreated,
	"TIME_STARTED":  ListPerformanceTuningAnalysisSortByStarted,
	"TIME_FINISHED": ListPerformanceTuningAnalysisSortByFinished,
}

var mappingListPerformanceTuningAnalysisSortByEnumLowerCase = map[string]ListPerformanceTuningAnalysisSortByEnum{
	"time_created":  ListPerformanceTuningAnalysisSortByCreated,
	"time_started":  ListPerformanceTuningAnalysisSortByStarted,
	"time_finished": ListPerformanceTuningAnalysisSortByFinished,
}

// GetListPerformanceTuningAnalysisSortByEnumValues Enumerates the set of values for ListPerformanceTuningAnalysisSortByEnum
func GetListPerformanceTuningAnalysisSortByEnumValues() []ListPerformanceTuningAnalysisSortByEnum {
	values := make([]ListPerformanceTuningAnalysisSortByEnum, 0)
	for _, v := range mappingListPerformanceTuningAnalysisSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPerformanceTuningAnalysisSortByEnumStringValues Enumerates the set of values in String for ListPerformanceTuningAnalysisSortByEnum
func GetListPerformanceTuningAnalysisSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"TIME_STARTED",
		"TIME_FINISHED",
	}
}

// GetMappingListPerformanceTuningAnalysisSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPerformanceTuningAnalysisSortByEnum(val string) (ListPerformanceTuningAnalysisSortByEnum, bool) {
	enum, ok := mappingListPerformanceTuningAnalysisSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
