// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPatchTaskStepsRequest wrapper for the ListPatchTaskSteps operation
type ListPatchTaskStepsRequest struct {

	// Unique PatchOperation identifier
	PatchOperationId *string `mandatory:"true" contributesTo:"path" name:"patchOperationId"`

	// The required ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique PatchTask key
	PatchTaskKey *int64 `mandatory:"true" contributesTo:"query" name:"patchTaskKey"`

	// unique PatchTaskStep key
	PatchTaskStepKey *int64 `mandatory:"false" contributesTo:"query" name:"patchTaskStepKey"`

	// A filter to return only those patch task steps whose type matches the given type.
	Type PatchTaskStepTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// A filter to return only those patch task steps whose workState matches the given workState.
	WorkState PatchTaskStepWorkStateEnum `mandatory:"false" contributesTo:"query" name:"workState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListPatchTaskStepsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by.
	SortBy ListPatchTaskStepsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to retrieve details when calling GET APIs
	IsDetailed *bool `mandatory:"false" contributesTo:"query" name:"isDetailed"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPatchTaskStepsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPatchTaskStepsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPatchTaskStepsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPatchTaskStepsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPatchTaskStepsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchTaskStepTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetPatchTaskStepTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchTaskStepWorkStateEnum(string(request.WorkState)); !ok && request.WorkState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkState: %s. Supported values are: %s.", request.WorkState, strings.Join(GetPatchTaskStepWorkStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatchTaskStepsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPatchTaskStepsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatchTaskStepsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPatchTaskStepsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPatchTaskStepsResponse wrapper for the ListPatchTaskSteps operation
type ListPatchTaskStepsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PatchTaskStepCollection instances
	PatchTaskStepCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPatchTaskStepsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPatchTaskStepsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPatchTaskStepsSortOrderEnum Enum with underlying type: string
type ListPatchTaskStepsSortOrderEnum string

// Set of constants representing the allowable values for ListPatchTaskStepsSortOrderEnum
const (
	ListPatchTaskStepsSortOrderAsc  ListPatchTaskStepsSortOrderEnum = "ASC"
	ListPatchTaskStepsSortOrderDesc ListPatchTaskStepsSortOrderEnum = "DESC"
)

var mappingListPatchTaskStepsSortOrderEnum = map[string]ListPatchTaskStepsSortOrderEnum{
	"ASC":  ListPatchTaskStepsSortOrderAsc,
	"DESC": ListPatchTaskStepsSortOrderDesc,
}

var mappingListPatchTaskStepsSortOrderEnumLowerCase = map[string]ListPatchTaskStepsSortOrderEnum{
	"asc":  ListPatchTaskStepsSortOrderAsc,
	"desc": ListPatchTaskStepsSortOrderDesc,
}

// GetListPatchTaskStepsSortOrderEnumValues Enumerates the set of values for ListPatchTaskStepsSortOrderEnum
func GetListPatchTaskStepsSortOrderEnumValues() []ListPatchTaskStepsSortOrderEnum {
	values := make([]ListPatchTaskStepsSortOrderEnum, 0)
	for _, v := range mappingListPatchTaskStepsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchTaskStepsSortOrderEnumStringValues Enumerates the set of values in String for ListPatchTaskStepsSortOrderEnum
func GetListPatchTaskStepsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPatchTaskStepsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchTaskStepsSortOrderEnum(val string) (ListPatchTaskStepsSortOrderEnum, bool) {
	enum, ok := mappingListPatchTaskStepsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPatchTaskStepsSortByEnum Enum with underlying type: string
type ListPatchTaskStepsSortByEnum string

// Set of constants representing the allowable values for ListPatchTaskStepsSortByEnum
const (
	ListPatchTaskStepsSortByTimecreated          ListPatchTaskStepsSortByEnum = "timeCreated"
	ListPatchTaskStepsSortByTimestarted          ListPatchTaskStepsSortByEnum = "timeStarted"
	ListPatchTaskStepsSortByTimecompleted        ListPatchTaskStepsSortByEnum = "timeCompleted"
	ListPatchTaskStepsSortByTimeelapsedinseconds ListPatchTaskStepsSortByEnum = "timeElapsedInSeconds"
	ListPatchTaskStepsSortByWorkstate            ListPatchTaskStepsSortByEnum = "workState"
)

var mappingListPatchTaskStepsSortByEnum = map[string]ListPatchTaskStepsSortByEnum{
	"timeCreated":          ListPatchTaskStepsSortByTimecreated,
	"timeStarted":          ListPatchTaskStepsSortByTimestarted,
	"timeCompleted":        ListPatchTaskStepsSortByTimecompleted,
	"timeElapsedInSeconds": ListPatchTaskStepsSortByTimeelapsedinseconds,
	"workState":            ListPatchTaskStepsSortByWorkstate,
}

var mappingListPatchTaskStepsSortByEnumLowerCase = map[string]ListPatchTaskStepsSortByEnum{
	"timecreated":          ListPatchTaskStepsSortByTimecreated,
	"timestarted":          ListPatchTaskStepsSortByTimestarted,
	"timecompleted":        ListPatchTaskStepsSortByTimecompleted,
	"timeelapsedinseconds": ListPatchTaskStepsSortByTimeelapsedinseconds,
	"workstate":            ListPatchTaskStepsSortByWorkstate,
}

// GetListPatchTaskStepsSortByEnumValues Enumerates the set of values for ListPatchTaskStepsSortByEnum
func GetListPatchTaskStepsSortByEnumValues() []ListPatchTaskStepsSortByEnum {
	values := make([]ListPatchTaskStepsSortByEnum, 0)
	for _, v := range mappingListPatchTaskStepsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchTaskStepsSortByEnumStringValues Enumerates the set of values in String for ListPatchTaskStepsSortByEnum
func GetListPatchTaskStepsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeStarted",
		"timeCompleted",
		"timeElapsedInSeconds",
		"workState",
	}
}

// GetMappingListPatchTaskStepsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchTaskStepsSortByEnum(val string) (ListPatchTaskStepsSortByEnum, bool) {
	enum, ok := mappingListPatchTaskStepsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
