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

// ListTaskRecordsRequest wrapper for the ListTaskRecords operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListTaskRecords.go.html to see an example of how to use ListTaskRecordsRequest.
type ListTaskRecordsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The platform for the Task.
	Platform *string `mandatory:"false" contributesTo:"query" name:"platform"`

	// The type of the Task.
	Type TaskRecordTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique TaskDetail identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The current state of the Task.
	LifecycleState TaskRecordLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListTaskRecordsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListTaskRecordsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTaskRecordsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTaskRecordsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTaskRecordsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTaskRecordsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTaskRecordsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTaskRecordTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetTaskRecordTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaskRecordLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetTaskRecordLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTaskRecordsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTaskRecordsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTaskRecordsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTaskRecordsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTaskRecordsResponse wrapper for the ListTaskRecords operation
type ListTaskRecordsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TaskRecordCollection instances
	TaskRecordCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTaskRecordsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTaskRecordsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTaskRecordsSortByEnum Enum with underlying type: string
type ListTaskRecordsSortByEnum string

// Set of constants representing the allowable values for ListTaskRecordsSortByEnum
const (
	ListTaskRecordsSortByTimecreated ListTaskRecordsSortByEnum = "timeCreated"
	ListTaskRecordsSortByDisplayname ListTaskRecordsSortByEnum = "displayName"
)

var mappingListTaskRecordsSortByEnum = map[string]ListTaskRecordsSortByEnum{
	"timeCreated": ListTaskRecordsSortByTimecreated,
	"displayName": ListTaskRecordsSortByDisplayname,
}

var mappingListTaskRecordsSortByEnumLowerCase = map[string]ListTaskRecordsSortByEnum{
	"timecreated": ListTaskRecordsSortByTimecreated,
	"displayname": ListTaskRecordsSortByDisplayname,
}

// GetListTaskRecordsSortByEnumValues Enumerates the set of values for ListTaskRecordsSortByEnum
func GetListTaskRecordsSortByEnumValues() []ListTaskRecordsSortByEnum {
	values := make([]ListTaskRecordsSortByEnum, 0)
	for _, v := range mappingListTaskRecordsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTaskRecordsSortByEnumStringValues Enumerates the set of values in String for ListTaskRecordsSortByEnum
func GetListTaskRecordsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListTaskRecordsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTaskRecordsSortByEnum(val string) (ListTaskRecordsSortByEnum, bool) {
	enum, ok := mappingListTaskRecordsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTaskRecordsSortOrderEnum Enum with underlying type: string
type ListTaskRecordsSortOrderEnum string

// Set of constants representing the allowable values for ListTaskRecordsSortOrderEnum
const (
	ListTaskRecordsSortOrderAsc  ListTaskRecordsSortOrderEnum = "ASC"
	ListTaskRecordsSortOrderDesc ListTaskRecordsSortOrderEnum = "DESC"
)

var mappingListTaskRecordsSortOrderEnum = map[string]ListTaskRecordsSortOrderEnum{
	"ASC":  ListTaskRecordsSortOrderAsc,
	"DESC": ListTaskRecordsSortOrderDesc,
}

var mappingListTaskRecordsSortOrderEnumLowerCase = map[string]ListTaskRecordsSortOrderEnum{
	"asc":  ListTaskRecordsSortOrderAsc,
	"desc": ListTaskRecordsSortOrderDesc,
}

// GetListTaskRecordsSortOrderEnumValues Enumerates the set of values for ListTaskRecordsSortOrderEnum
func GetListTaskRecordsSortOrderEnumValues() []ListTaskRecordsSortOrderEnum {
	values := make([]ListTaskRecordsSortOrderEnum, 0)
	for _, v := range mappingListTaskRecordsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTaskRecordsSortOrderEnumStringValues Enumerates the set of values in String for ListTaskRecordsSortOrderEnum
func GetListTaskRecordsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTaskRecordsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTaskRecordsSortOrderEnum(val string) (ListTaskRecordsSortOrderEnum, bool) {
	enum, ok := mappingListTaskRecordsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
