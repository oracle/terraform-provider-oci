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

// ListOptimizerStatisticsCollectionOperationsRequest wrapper for the ListOptimizerStatisticsCollectionOperations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListOptimizerStatisticsCollectionOperations.go.html to see an example of how to use ListOptimizerStatisticsCollectionOperationsRequest.
type ListOptimizerStatisticsCollectionOperationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The start time of the time range to retrieve the optimizer statistics of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	StartTimeGreaterThanOrEqualTo *string `mandatory:"false" contributesTo:"query" name:"startTimeGreaterThanOrEqualTo"`

	// The end time of the time range to retrieve the optimizer statistics of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	EndTimeLessThanOrEqualTo *string `mandatory:"false" contributesTo:"query" name:"endTimeLessThanOrEqualTo"`

	// The filter types of the optimizer statistics tasks.
	TaskType ListOptimizerStatisticsCollectionOperationsTaskTypeEnum `mandatory:"false" contributesTo:"query" name:"taskType" omitEmpty:"true"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The parameter used to filter the optimizer statistics operations.
	// Any property of the OptimizerStatisticsCollectionOperationSummary can be used to define the filter condition.
	// The allowed conditional operators are AND or OR, and the allowed binary operators are are >, < and =. Any other operator is regarded invalid.
	// Example: jobName=<replace with job name> AND status=<replace with status>
	FilterBy *string `mandatory:"false" contributesTo:"query" name:"filterBy"`

	// Sorts the list of optimizer statistics operations based on a specific attribute.
	SortBy ListOptimizerStatisticsCollectionOperationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListOptimizerStatisticsCollectionOperationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOptimizerStatisticsCollectionOperationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOptimizerStatisticsCollectionOperationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOptimizerStatisticsCollectionOperationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOptimizerStatisticsCollectionOperationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOptimizerStatisticsCollectionOperationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOptimizerStatisticsCollectionOperationsTaskTypeEnum(string(request.TaskType)); !ok && request.TaskType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskType: %s. Supported values are: %s.", request.TaskType, strings.Join(GetListOptimizerStatisticsCollectionOperationsTaskTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOptimizerStatisticsCollectionOperationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOptimizerStatisticsCollectionOperationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOptimizerStatisticsCollectionOperationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOptimizerStatisticsCollectionOperationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOptimizerStatisticsCollectionOperationsResponse wrapper for the ListOptimizerStatisticsCollectionOperations operation
type ListOptimizerStatisticsCollectionOperationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OptimizerStatisticsCollectionOperationsCollection instances
	OptimizerStatisticsCollectionOperationsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOptimizerStatisticsCollectionOperationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOptimizerStatisticsCollectionOperationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOptimizerStatisticsCollectionOperationsTaskTypeEnum Enum with underlying type: string
type ListOptimizerStatisticsCollectionOperationsTaskTypeEnum string

// Set of constants representing the allowable values for ListOptimizerStatisticsCollectionOperationsTaskTypeEnum
const (
	ListOptimizerStatisticsCollectionOperationsTaskTypeAll    ListOptimizerStatisticsCollectionOperationsTaskTypeEnum = "ALL"
	ListOptimizerStatisticsCollectionOperationsTaskTypeManual ListOptimizerStatisticsCollectionOperationsTaskTypeEnum = "MANUAL"
	ListOptimizerStatisticsCollectionOperationsTaskTypeAuto   ListOptimizerStatisticsCollectionOperationsTaskTypeEnum = "AUTO"
)

var mappingListOptimizerStatisticsCollectionOperationsTaskTypeEnum = map[string]ListOptimizerStatisticsCollectionOperationsTaskTypeEnum{
	"ALL":    ListOptimizerStatisticsCollectionOperationsTaskTypeAll,
	"MANUAL": ListOptimizerStatisticsCollectionOperationsTaskTypeManual,
	"AUTO":   ListOptimizerStatisticsCollectionOperationsTaskTypeAuto,
}

var mappingListOptimizerStatisticsCollectionOperationsTaskTypeEnumLowerCase = map[string]ListOptimizerStatisticsCollectionOperationsTaskTypeEnum{
	"all":    ListOptimizerStatisticsCollectionOperationsTaskTypeAll,
	"manual": ListOptimizerStatisticsCollectionOperationsTaskTypeManual,
	"auto":   ListOptimizerStatisticsCollectionOperationsTaskTypeAuto,
}

// GetListOptimizerStatisticsCollectionOperationsTaskTypeEnumValues Enumerates the set of values for ListOptimizerStatisticsCollectionOperationsTaskTypeEnum
func GetListOptimizerStatisticsCollectionOperationsTaskTypeEnumValues() []ListOptimizerStatisticsCollectionOperationsTaskTypeEnum {
	values := make([]ListOptimizerStatisticsCollectionOperationsTaskTypeEnum, 0)
	for _, v := range mappingListOptimizerStatisticsCollectionOperationsTaskTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListOptimizerStatisticsCollectionOperationsTaskTypeEnumStringValues Enumerates the set of values in String for ListOptimizerStatisticsCollectionOperationsTaskTypeEnum
func GetListOptimizerStatisticsCollectionOperationsTaskTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"MANUAL",
		"AUTO",
	}
}

// GetMappingListOptimizerStatisticsCollectionOperationsTaskTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOptimizerStatisticsCollectionOperationsTaskTypeEnum(val string) (ListOptimizerStatisticsCollectionOperationsTaskTypeEnum, bool) {
	enum, ok := mappingListOptimizerStatisticsCollectionOperationsTaskTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOptimizerStatisticsCollectionOperationsSortByEnum Enum with underlying type: string
type ListOptimizerStatisticsCollectionOperationsSortByEnum string

// Set of constants representing the allowable values for ListOptimizerStatisticsCollectionOperationsSortByEnum
const (
	ListOptimizerStatisticsCollectionOperationsSortByStartTime ListOptimizerStatisticsCollectionOperationsSortByEnum = "START_TIME"
	ListOptimizerStatisticsCollectionOperationsSortByEndTime   ListOptimizerStatisticsCollectionOperationsSortByEnum = "END_TIME"
	ListOptimizerStatisticsCollectionOperationsSortByStatus    ListOptimizerStatisticsCollectionOperationsSortByEnum = "STATUS"
)

var mappingListOptimizerStatisticsCollectionOperationsSortByEnum = map[string]ListOptimizerStatisticsCollectionOperationsSortByEnum{
	"START_TIME": ListOptimizerStatisticsCollectionOperationsSortByStartTime,
	"END_TIME":   ListOptimizerStatisticsCollectionOperationsSortByEndTime,
	"STATUS":     ListOptimizerStatisticsCollectionOperationsSortByStatus,
}

var mappingListOptimizerStatisticsCollectionOperationsSortByEnumLowerCase = map[string]ListOptimizerStatisticsCollectionOperationsSortByEnum{
	"start_time": ListOptimizerStatisticsCollectionOperationsSortByStartTime,
	"end_time":   ListOptimizerStatisticsCollectionOperationsSortByEndTime,
	"status":     ListOptimizerStatisticsCollectionOperationsSortByStatus,
}

// GetListOptimizerStatisticsCollectionOperationsSortByEnumValues Enumerates the set of values for ListOptimizerStatisticsCollectionOperationsSortByEnum
func GetListOptimizerStatisticsCollectionOperationsSortByEnumValues() []ListOptimizerStatisticsCollectionOperationsSortByEnum {
	values := make([]ListOptimizerStatisticsCollectionOperationsSortByEnum, 0)
	for _, v := range mappingListOptimizerStatisticsCollectionOperationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOptimizerStatisticsCollectionOperationsSortByEnumStringValues Enumerates the set of values in String for ListOptimizerStatisticsCollectionOperationsSortByEnum
func GetListOptimizerStatisticsCollectionOperationsSortByEnumStringValues() []string {
	return []string{
		"START_TIME",
		"END_TIME",
		"STATUS",
	}
}

// GetMappingListOptimizerStatisticsCollectionOperationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOptimizerStatisticsCollectionOperationsSortByEnum(val string) (ListOptimizerStatisticsCollectionOperationsSortByEnum, bool) {
	enum, ok := mappingListOptimizerStatisticsCollectionOperationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOptimizerStatisticsCollectionOperationsSortOrderEnum Enum with underlying type: string
type ListOptimizerStatisticsCollectionOperationsSortOrderEnum string

// Set of constants representing the allowable values for ListOptimizerStatisticsCollectionOperationsSortOrderEnum
const (
	ListOptimizerStatisticsCollectionOperationsSortOrderAsc  ListOptimizerStatisticsCollectionOperationsSortOrderEnum = "ASC"
	ListOptimizerStatisticsCollectionOperationsSortOrderDesc ListOptimizerStatisticsCollectionOperationsSortOrderEnum = "DESC"
)

var mappingListOptimizerStatisticsCollectionOperationsSortOrderEnum = map[string]ListOptimizerStatisticsCollectionOperationsSortOrderEnum{
	"ASC":  ListOptimizerStatisticsCollectionOperationsSortOrderAsc,
	"DESC": ListOptimizerStatisticsCollectionOperationsSortOrderDesc,
}

var mappingListOptimizerStatisticsCollectionOperationsSortOrderEnumLowerCase = map[string]ListOptimizerStatisticsCollectionOperationsSortOrderEnum{
	"asc":  ListOptimizerStatisticsCollectionOperationsSortOrderAsc,
	"desc": ListOptimizerStatisticsCollectionOperationsSortOrderDesc,
}

// GetListOptimizerStatisticsCollectionOperationsSortOrderEnumValues Enumerates the set of values for ListOptimizerStatisticsCollectionOperationsSortOrderEnum
func GetListOptimizerStatisticsCollectionOperationsSortOrderEnumValues() []ListOptimizerStatisticsCollectionOperationsSortOrderEnum {
	values := make([]ListOptimizerStatisticsCollectionOperationsSortOrderEnum, 0)
	for _, v := range mappingListOptimizerStatisticsCollectionOperationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOptimizerStatisticsCollectionOperationsSortOrderEnumStringValues Enumerates the set of values in String for ListOptimizerStatisticsCollectionOperationsSortOrderEnum
func GetListOptimizerStatisticsCollectionOperationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOptimizerStatisticsCollectionOperationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOptimizerStatisticsCollectionOperationsSortOrderEnum(val string) (ListOptimizerStatisticsCollectionOperationsSortOrderEnum, bool) {
	enum, ok := mappingListOptimizerStatisticsCollectionOperationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
