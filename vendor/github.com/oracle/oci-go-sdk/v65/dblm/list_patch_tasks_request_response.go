// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPatchTasksRequest wrapper for the ListPatchTasks operation
type ListPatchTasksRequest struct {

	// Unique PatchOperation identifier
	PatchOperationId *string `mandatory:"true" contributesTo:"path" name:"patchOperationId"`

	// The required ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique PatchTask key
	PatchTaskKey *int64 `mandatory:"false" contributesTo:"query" name:"patchTaskKey"`

	// A filter to return only those patch tasks whose status matches the given status.
	Status PatchTaskStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// A filter to return only those patch tasks whose type matches the given type.
	Type PatchTaskTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// A filter to return only resources that match the given resource id.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// A filter to return only resources that match the given resource name.
	ResourceName *string `mandatory:"false" contributesTo:"query" name:"resourceName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListPatchTasksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by.
	SortBy ListPatchTasksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to retrieve details when calling GET APIs
	IsDetailed *bool `mandatory:"false" contributesTo:"query" name:"isDetailed"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPatchTasksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPatchTasksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPatchTasksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPatchTasksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPatchTasksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchTaskStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetPatchTaskStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchTaskTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetPatchTaskTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatchTasksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPatchTasksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatchTasksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPatchTasksSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPatchTasksResponse wrapper for the ListPatchTasks operation
type ListPatchTasksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PatchTaskCollection instances
	PatchTaskCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPatchTasksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPatchTasksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPatchTasksSortOrderEnum Enum with underlying type: string
type ListPatchTasksSortOrderEnum string

// Set of constants representing the allowable values for ListPatchTasksSortOrderEnum
const (
	ListPatchTasksSortOrderAsc  ListPatchTasksSortOrderEnum = "ASC"
	ListPatchTasksSortOrderDesc ListPatchTasksSortOrderEnum = "DESC"
)

var mappingListPatchTasksSortOrderEnum = map[string]ListPatchTasksSortOrderEnum{
	"ASC":  ListPatchTasksSortOrderAsc,
	"DESC": ListPatchTasksSortOrderDesc,
}

var mappingListPatchTasksSortOrderEnumLowerCase = map[string]ListPatchTasksSortOrderEnum{
	"asc":  ListPatchTasksSortOrderAsc,
	"desc": ListPatchTasksSortOrderDesc,
}

// GetListPatchTasksSortOrderEnumValues Enumerates the set of values for ListPatchTasksSortOrderEnum
func GetListPatchTasksSortOrderEnumValues() []ListPatchTasksSortOrderEnum {
	values := make([]ListPatchTasksSortOrderEnum, 0)
	for _, v := range mappingListPatchTasksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchTasksSortOrderEnumStringValues Enumerates the set of values in String for ListPatchTasksSortOrderEnum
func GetListPatchTasksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPatchTasksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchTasksSortOrderEnum(val string) (ListPatchTasksSortOrderEnum, bool) {
	enum, ok := mappingListPatchTasksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPatchTasksSortByEnum Enum with underlying type: string
type ListPatchTasksSortByEnum string

// Set of constants representing the allowable values for ListPatchTasksSortByEnum
const (
	ListPatchTasksSortByTimecreated          ListPatchTasksSortByEnum = "timeCreated"
	ListPatchTasksSortByTimestarted          ListPatchTasksSortByEnum = "timeStarted"
	ListPatchTasksSortByTimecompleted        ListPatchTasksSortByEnum = "timeCompleted"
	ListPatchTasksSortByTimeelapsedinseconds ListPatchTasksSortByEnum = "timeElapsedInSeconds"
	ListPatchTasksSortByStatus               ListPatchTasksSortByEnum = "status"
)

var mappingListPatchTasksSortByEnum = map[string]ListPatchTasksSortByEnum{
	"timeCreated":          ListPatchTasksSortByTimecreated,
	"timeStarted":          ListPatchTasksSortByTimestarted,
	"timeCompleted":        ListPatchTasksSortByTimecompleted,
	"timeElapsedInSeconds": ListPatchTasksSortByTimeelapsedinseconds,
	"status":               ListPatchTasksSortByStatus,
}

var mappingListPatchTasksSortByEnumLowerCase = map[string]ListPatchTasksSortByEnum{
	"timecreated":          ListPatchTasksSortByTimecreated,
	"timestarted":          ListPatchTasksSortByTimestarted,
	"timecompleted":        ListPatchTasksSortByTimecompleted,
	"timeelapsedinseconds": ListPatchTasksSortByTimeelapsedinseconds,
	"status":               ListPatchTasksSortByStatus,
}

// GetListPatchTasksSortByEnumValues Enumerates the set of values for ListPatchTasksSortByEnum
func GetListPatchTasksSortByEnumValues() []ListPatchTasksSortByEnum {
	values := make([]ListPatchTasksSortByEnum, 0)
	for _, v := range mappingListPatchTasksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchTasksSortByEnumStringValues Enumerates the set of values in String for ListPatchTasksSortByEnum
func GetListPatchTasksSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeStarted",
		"timeCompleted",
		"timeElapsedInSeconds",
		"status",
	}
}

// GetMappingListPatchTasksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchTasksSortByEnum(val string) (ListPatchTasksSortByEnum, bool) {
	enum, ok := mappingListPatchTasksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
