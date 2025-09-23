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

// ListJavaMigrationAnalysisRequest wrapper for the ListJavaMigrationAnalysis operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsutils/ListJavaMigrationAnalysis.go.html to see an example of how to use ListJavaMigrationAnalysisRequest.
type ListJavaMigrationAnalysisRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Java Migration Analysis.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The project name of the Performance Tuning Analysis to query for.
	AnalysisProjectName *string `mandatory:"false" contributesTo:"query" name:"analysisProjectName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListJavaMigrationAnalysisSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort Java Migration Analysis. Only one sort order may be provided.
	SortBy ListJavaMigrationAnalysisSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJavaMigrationAnalysisRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJavaMigrationAnalysisRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJavaMigrationAnalysisRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJavaMigrationAnalysisRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJavaMigrationAnalysisRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJavaMigrationAnalysisSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJavaMigrationAnalysisSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaMigrationAnalysisSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJavaMigrationAnalysisSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJavaMigrationAnalysisResponse wrapper for the ListJavaMigrationAnalysis operation
type ListJavaMigrationAnalysisResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JavaMigrationAnalysisCollection instances
	JavaMigrationAnalysisCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJavaMigrationAnalysisResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJavaMigrationAnalysisResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJavaMigrationAnalysisSortOrderEnum Enum with underlying type: string
type ListJavaMigrationAnalysisSortOrderEnum string

// Set of constants representing the allowable values for ListJavaMigrationAnalysisSortOrderEnum
const (
	ListJavaMigrationAnalysisSortOrderAsc  ListJavaMigrationAnalysisSortOrderEnum = "ASC"
	ListJavaMigrationAnalysisSortOrderDesc ListJavaMigrationAnalysisSortOrderEnum = "DESC"
)

var mappingListJavaMigrationAnalysisSortOrderEnum = map[string]ListJavaMigrationAnalysisSortOrderEnum{
	"ASC":  ListJavaMigrationAnalysisSortOrderAsc,
	"DESC": ListJavaMigrationAnalysisSortOrderDesc,
}

var mappingListJavaMigrationAnalysisSortOrderEnumLowerCase = map[string]ListJavaMigrationAnalysisSortOrderEnum{
	"asc":  ListJavaMigrationAnalysisSortOrderAsc,
	"desc": ListJavaMigrationAnalysisSortOrderDesc,
}

// GetListJavaMigrationAnalysisSortOrderEnumValues Enumerates the set of values for ListJavaMigrationAnalysisSortOrderEnum
func GetListJavaMigrationAnalysisSortOrderEnumValues() []ListJavaMigrationAnalysisSortOrderEnum {
	values := make([]ListJavaMigrationAnalysisSortOrderEnum, 0)
	for _, v := range mappingListJavaMigrationAnalysisSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaMigrationAnalysisSortOrderEnumStringValues Enumerates the set of values in String for ListJavaMigrationAnalysisSortOrderEnum
func GetListJavaMigrationAnalysisSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJavaMigrationAnalysisSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaMigrationAnalysisSortOrderEnum(val string) (ListJavaMigrationAnalysisSortOrderEnum, bool) {
	enum, ok := mappingListJavaMigrationAnalysisSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaMigrationAnalysisSortByEnum Enum with underlying type: string
type ListJavaMigrationAnalysisSortByEnum string

// Set of constants representing the allowable values for ListJavaMigrationAnalysisSortByEnum
const (
	ListJavaMigrationAnalysisSortByCreated  ListJavaMigrationAnalysisSortByEnum = "TIME_CREATED"
	ListJavaMigrationAnalysisSortByStarted  ListJavaMigrationAnalysisSortByEnum = "TIME_STARTED"
	ListJavaMigrationAnalysisSortByFinished ListJavaMigrationAnalysisSortByEnum = "TIME_FINISHED"
)

var mappingListJavaMigrationAnalysisSortByEnum = map[string]ListJavaMigrationAnalysisSortByEnum{
	"TIME_CREATED":  ListJavaMigrationAnalysisSortByCreated,
	"TIME_STARTED":  ListJavaMigrationAnalysisSortByStarted,
	"TIME_FINISHED": ListJavaMigrationAnalysisSortByFinished,
}

var mappingListJavaMigrationAnalysisSortByEnumLowerCase = map[string]ListJavaMigrationAnalysisSortByEnum{
	"time_created":  ListJavaMigrationAnalysisSortByCreated,
	"time_started":  ListJavaMigrationAnalysisSortByStarted,
	"time_finished": ListJavaMigrationAnalysisSortByFinished,
}

// GetListJavaMigrationAnalysisSortByEnumValues Enumerates the set of values for ListJavaMigrationAnalysisSortByEnum
func GetListJavaMigrationAnalysisSortByEnumValues() []ListJavaMigrationAnalysisSortByEnum {
	values := make([]ListJavaMigrationAnalysisSortByEnum, 0)
	for _, v := range mappingListJavaMigrationAnalysisSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaMigrationAnalysisSortByEnumStringValues Enumerates the set of values in String for ListJavaMigrationAnalysisSortByEnum
func GetListJavaMigrationAnalysisSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"TIME_STARTED",
		"TIME_FINISHED",
	}
}

// GetMappingListJavaMigrationAnalysisSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaMigrationAnalysisSortByEnum(val string) (ListJavaMigrationAnalysisSortByEnum, bool) {
	enum, ok := mappingListJavaMigrationAnalysisSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
