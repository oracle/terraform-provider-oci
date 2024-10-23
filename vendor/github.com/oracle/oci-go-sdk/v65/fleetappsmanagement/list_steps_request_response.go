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

// ListStepsRequest wrapper for the ListSteps operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListSteps.go.html to see an example of how to use ListStepsRequest.
type ListStepsRequest struct {

	// unique SchedulerJob identifier
	SchedulerJobId *string `mandatory:"true" contributesTo:"path" name:"schedulerJobId"`

	// unique jobActivity identifier
	JobActivityId *string `mandatory:"true" contributesTo:"path" name:"jobActivityId"`

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
	SortOrder ListStepsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeStarted is descending.
	SortBy ListStepsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStepsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStepsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListStepsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStepsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListStepsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListStepsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListStepsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStepsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListStepsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListStepsResponse wrapper for the ListSteps operation
type ListStepsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of StepCollection instances
	StepCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListStepsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStepsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStepsSortOrderEnum Enum with underlying type: string
type ListStepsSortOrderEnum string

// Set of constants representing the allowable values for ListStepsSortOrderEnum
const (
	ListStepsSortOrderAsc  ListStepsSortOrderEnum = "ASC"
	ListStepsSortOrderDesc ListStepsSortOrderEnum = "DESC"
)

var mappingListStepsSortOrderEnum = map[string]ListStepsSortOrderEnum{
	"ASC":  ListStepsSortOrderAsc,
	"DESC": ListStepsSortOrderDesc,
}

var mappingListStepsSortOrderEnumLowerCase = map[string]ListStepsSortOrderEnum{
	"asc":  ListStepsSortOrderAsc,
	"desc": ListStepsSortOrderDesc,
}

// GetListStepsSortOrderEnumValues Enumerates the set of values for ListStepsSortOrderEnum
func GetListStepsSortOrderEnumValues() []ListStepsSortOrderEnum {
	values := make([]ListStepsSortOrderEnum, 0)
	for _, v := range mappingListStepsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListStepsSortOrderEnumStringValues Enumerates the set of values in String for ListStepsSortOrderEnum
func GetListStepsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListStepsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStepsSortOrderEnum(val string) (ListStepsSortOrderEnum, bool) {
	enum, ok := mappingListStepsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStepsSortByEnum Enum with underlying type: string
type ListStepsSortByEnum string

// Set of constants representing the allowable values for ListStepsSortByEnum
const (
	ListStepsSortByTimestarted ListStepsSortByEnum = "timeStarted"
)

var mappingListStepsSortByEnum = map[string]ListStepsSortByEnum{
	"timeStarted": ListStepsSortByTimestarted,
}

var mappingListStepsSortByEnumLowerCase = map[string]ListStepsSortByEnum{
	"timestarted": ListStepsSortByTimestarted,
}

// GetListStepsSortByEnumValues Enumerates the set of values for ListStepsSortByEnum
func GetListStepsSortByEnumValues() []ListStepsSortByEnum {
	values := make([]ListStepsSortByEnum, 0)
	for _, v := range mappingListStepsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListStepsSortByEnumStringValues Enumerates the set of values in String for ListStepsSortByEnum
func GetListStepsSortByEnumStringValues() []string {
	return []string{
		"timeStarted",
	}
}

// GetMappingListStepsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStepsSortByEnum(val string) (ListStepsSortByEnum, bool) {
	enum, ok := mappingListStepsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
