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

// ListOptimizerStatisticsCollectionAggregationsRequest wrapper for the ListOptimizerStatisticsCollectionAggregations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListOptimizerStatisticsCollectionAggregations.go.html to see an example of how to use ListOptimizerStatisticsCollectionAggregationsRequest.
type ListOptimizerStatisticsCollectionAggregationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The optimizer statistics tasks grouped by type.
	GroupType ListOptimizerStatisticsCollectionAggregationsGroupTypeEnum `mandatory:"true" contributesTo:"query" name:"groupType" omitEmpty:"true"`

	// The start time of the time range to retrieve the optimizer statistics of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	StartTimeGreaterThanOrEqualTo *string `mandatory:"false" contributesTo:"query" name:"startTimeGreaterThanOrEqualTo"`

	// The end time of the time range to retrieve the optimizer statistics of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	EndTimeLessThanOrEqualTo *string `mandatory:"false" contributesTo:"query" name:"endTimeLessThanOrEqualTo"`

	// The filter types of the optimizer statistics tasks.
	TaskType ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum `mandatory:"false" contributesTo:"query" name:"taskType" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOptimizerStatisticsCollectionAggregationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOptimizerStatisticsCollectionAggregationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOptimizerStatisticsCollectionAggregationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOptimizerStatisticsCollectionAggregationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOptimizerStatisticsCollectionAggregationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOptimizerStatisticsCollectionAggregationsGroupTypeEnum(string(request.GroupType)); !ok && request.GroupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupType: %s. Supported values are: %s.", request.GroupType, strings.Join(GetListOptimizerStatisticsCollectionAggregationsGroupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOptimizerStatisticsCollectionAggregationsTaskTypeEnum(string(request.TaskType)); !ok && request.TaskType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskType: %s. Supported values are: %s.", request.TaskType, strings.Join(GetListOptimizerStatisticsCollectionAggregationsTaskTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOptimizerStatisticsCollectionAggregationsResponse wrapper for the ListOptimizerStatisticsCollectionAggregations operation
type ListOptimizerStatisticsCollectionAggregationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OptimizerStatisticsCollectionAggregationsCollection instances
	OptimizerStatisticsCollectionAggregationsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOptimizerStatisticsCollectionAggregationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOptimizerStatisticsCollectionAggregationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOptimizerStatisticsCollectionAggregationsGroupTypeEnum Enum with underlying type: string
type ListOptimizerStatisticsCollectionAggregationsGroupTypeEnum string

// Set of constants representing the allowable values for ListOptimizerStatisticsCollectionAggregationsGroupTypeEnum
const (
	ListOptimizerStatisticsCollectionAggregationsGroupTypeStatus        ListOptimizerStatisticsCollectionAggregationsGroupTypeEnum = "TASK_STATUS"
	ListOptimizerStatisticsCollectionAggregationsGroupTypeObjectsStatus ListOptimizerStatisticsCollectionAggregationsGroupTypeEnum = "TASK_OBJECTS_STATUS"
)

var mappingListOptimizerStatisticsCollectionAggregationsGroupTypeEnum = map[string]ListOptimizerStatisticsCollectionAggregationsGroupTypeEnum{
	"TASK_STATUS":         ListOptimizerStatisticsCollectionAggregationsGroupTypeStatus,
	"TASK_OBJECTS_STATUS": ListOptimizerStatisticsCollectionAggregationsGroupTypeObjectsStatus,
}

var mappingListOptimizerStatisticsCollectionAggregationsGroupTypeEnumLowerCase = map[string]ListOptimizerStatisticsCollectionAggregationsGroupTypeEnum{
	"task_status":         ListOptimizerStatisticsCollectionAggregationsGroupTypeStatus,
	"task_objects_status": ListOptimizerStatisticsCollectionAggregationsGroupTypeObjectsStatus,
}

// GetListOptimizerStatisticsCollectionAggregationsGroupTypeEnumValues Enumerates the set of values for ListOptimizerStatisticsCollectionAggregationsGroupTypeEnum
func GetListOptimizerStatisticsCollectionAggregationsGroupTypeEnumValues() []ListOptimizerStatisticsCollectionAggregationsGroupTypeEnum {
	values := make([]ListOptimizerStatisticsCollectionAggregationsGroupTypeEnum, 0)
	for _, v := range mappingListOptimizerStatisticsCollectionAggregationsGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListOptimizerStatisticsCollectionAggregationsGroupTypeEnumStringValues Enumerates the set of values in String for ListOptimizerStatisticsCollectionAggregationsGroupTypeEnum
func GetListOptimizerStatisticsCollectionAggregationsGroupTypeEnumStringValues() []string {
	return []string{
		"TASK_STATUS",
		"TASK_OBJECTS_STATUS",
	}
}

// GetMappingListOptimizerStatisticsCollectionAggregationsGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOptimizerStatisticsCollectionAggregationsGroupTypeEnum(val string) (ListOptimizerStatisticsCollectionAggregationsGroupTypeEnum, bool) {
	enum, ok := mappingListOptimizerStatisticsCollectionAggregationsGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum Enum with underlying type: string
type ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum string

// Set of constants representing the allowable values for ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum
const (
	ListOptimizerStatisticsCollectionAggregationsTaskTypeAll    ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum = "ALL"
	ListOptimizerStatisticsCollectionAggregationsTaskTypeManual ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum = "MANUAL"
	ListOptimizerStatisticsCollectionAggregationsTaskTypeAuto   ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum = "AUTO"
)

var mappingListOptimizerStatisticsCollectionAggregationsTaskTypeEnum = map[string]ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum{
	"ALL":    ListOptimizerStatisticsCollectionAggregationsTaskTypeAll,
	"MANUAL": ListOptimizerStatisticsCollectionAggregationsTaskTypeManual,
	"AUTO":   ListOptimizerStatisticsCollectionAggregationsTaskTypeAuto,
}

var mappingListOptimizerStatisticsCollectionAggregationsTaskTypeEnumLowerCase = map[string]ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum{
	"all":    ListOptimizerStatisticsCollectionAggregationsTaskTypeAll,
	"manual": ListOptimizerStatisticsCollectionAggregationsTaskTypeManual,
	"auto":   ListOptimizerStatisticsCollectionAggregationsTaskTypeAuto,
}

// GetListOptimizerStatisticsCollectionAggregationsTaskTypeEnumValues Enumerates the set of values for ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum
func GetListOptimizerStatisticsCollectionAggregationsTaskTypeEnumValues() []ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum {
	values := make([]ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum, 0)
	for _, v := range mappingListOptimizerStatisticsCollectionAggregationsTaskTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListOptimizerStatisticsCollectionAggregationsTaskTypeEnumStringValues Enumerates the set of values in String for ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum
func GetListOptimizerStatisticsCollectionAggregationsTaskTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"MANUAL",
		"AUTO",
	}
}

// GetMappingListOptimizerStatisticsCollectionAggregationsTaskTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOptimizerStatisticsCollectionAggregationsTaskTypeEnum(val string) (ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum, bool) {
	enum, ok := mappingListOptimizerStatisticsCollectionAggregationsTaskTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
