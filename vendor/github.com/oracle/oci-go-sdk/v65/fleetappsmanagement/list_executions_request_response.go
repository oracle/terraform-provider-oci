// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExecutionsRequest wrapper for the ListExecutions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListExecutions.go.html to see an example of how to use ListExecutionsRequest.
type ListExecutionsRequest struct {

	// unique SchedulerJob identifier
	SchedulerJobId *string `mandatory:"true" contributesTo:"path" name:"schedulerJobId"`

	// unique jobActivity identifier
	JobActivityId *string `mandatory:"true" contributesTo:"path" name:"jobActivityId"`

	// unique Resource identifier
	ResourceId *string `mandatory:"true" contributesTo:"path" name:"resourceId"`

	// Task Id
	ResourceTaskId *string `mandatory:"false" contributesTo:"query" name:"resourceTaskId"`

	// Unique step name
	StepName *string `mandatory:"false" contributesTo:"query" name:"stepName"`

	// Unique target name
	TargetName *string `mandatory:"false" contributesTo:"query" name:"targetName"`

	// Task Order Sequence
	Sequence *string `mandatory:"false" contributesTo:"query" name:"sequence"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListExecutionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeStarted is descending.
	SortBy ListExecutionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExecutionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExecutionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExecutionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExecutionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExecutionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExecutionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExecutionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExecutionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExecutionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExecutionsResponse wrapper for the ListExecutions operation
type ListExecutionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExecutionCollection instances
	ExecutionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExecutionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExecutionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExecutionsSortOrderEnum Enum with underlying type: string
type ListExecutionsSortOrderEnum string

// Set of constants representing the allowable values for ListExecutionsSortOrderEnum
const (
	ListExecutionsSortOrderAsc  ListExecutionsSortOrderEnum = "ASC"
	ListExecutionsSortOrderDesc ListExecutionsSortOrderEnum = "DESC"
)

var mappingListExecutionsSortOrderEnum = map[string]ListExecutionsSortOrderEnum{
	"ASC":  ListExecutionsSortOrderAsc,
	"DESC": ListExecutionsSortOrderDesc,
}

var mappingListExecutionsSortOrderEnumLowerCase = map[string]ListExecutionsSortOrderEnum{
	"asc":  ListExecutionsSortOrderAsc,
	"desc": ListExecutionsSortOrderDesc,
}

// GetListExecutionsSortOrderEnumValues Enumerates the set of values for ListExecutionsSortOrderEnum
func GetListExecutionsSortOrderEnumValues() []ListExecutionsSortOrderEnum {
	values := make([]ListExecutionsSortOrderEnum, 0)
	for _, v := range mappingListExecutionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExecutionsSortOrderEnumStringValues Enumerates the set of values in String for ListExecutionsSortOrderEnum
func GetListExecutionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExecutionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExecutionsSortOrderEnum(val string) (ListExecutionsSortOrderEnum, bool) {
	enum, ok := mappingListExecutionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExecutionsSortByEnum Enum with underlying type: string
type ListExecutionsSortByEnum string

// Set of constants representing the allowable values for ListExecutionsSortByEnum
const (
	ListExecutionsSortByTimestarted ListExecutionsSortByEnum = "timeStarted"
)

var mappingListExecutionsSortByEnum = map[string]ListExecutionsSortByEnum{
	"timeStarted": ListExecutionsSortByTimestarted,
}

var mappingListExecutionsSortByEnumLowerCase = map[string]ListExecutionsSortByEnum{
	"timestarted": ListExecutionsSortByTimestarted,
}

// GetListExecutionsSortByEnumValues Enumerates the set of values for ListExecutionsSortByEnum
func GetListExecutionsSortByEnumValues() []ListExecutionsSortByEnum {
	values := make([]ListExecutionsSortByEnum, 0)
	for _, v := range mappingListExecutionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExecutionsSortByEnumStringValues Enumerates the set of values in String for ListExecutionsSortByEnum
func GetListExecutionsSortByEnumStringValues() []string {
	return []string{
		"timeStarted",
	}
}

// GetMappingListExecutionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExecutionsSortByEnum(val string) (ListExecutionsSortByEnum, bool) {
	enum, ok := mappingListExecutionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
