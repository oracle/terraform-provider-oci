// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPipelinesRequest wrapper for the ListPipelines operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListPipelines.go.html to see an example of how to use ListPipelinesRequest.
type ListPipelinesRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// <b>Filter</b> results by OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The current state of the Pipeline.
	LifecycleState ListPipelinesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 100 is the maximum.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListPipelinesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, the results are shown
	// in descending order. When you sort by `displayName`, the results are
	// shown in ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListPipelinesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPipelinesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPipelinesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPipelinesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPipelinesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPipelinesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPipelinesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListPipelinesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPipelinesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPipelinesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPipelinesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPipelinesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPipelinesResponse wrapper for the ListPipelines operation
type ListPipelinesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []PipelineSummary instances
	Items []PipelineSummary `presentIn:"body"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListPipelinesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPipelinesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPipelinesLifecycleStateEnum Enum with underlying type: string
type ListPipelinesLifecycleStateEnum string

// Set of constants representing the allowable values for ListPipelinesLifecycleStateEnum
const (
	ListPipelinesLifecycleStateCreating ListPipelinesLifecycleStateEnum = "CREATING"
	ListPipelinesLifecycleStateActive   ListPipelinesLifecycleStateEnum = "ACTIVE"
	ListPipelinesLifecycleStateDeleting ListPipelinesLifecycleStateEnum = "DELETING"
	ListPipelinesLifecycleStateFailed   ListPipelinesLifecycleStateEnum = "FAILED"
	ListPipelinesLifecycleStateDeleted  ListPipelinesLifecycleStateEnum = "DELETED"
)

var mappingListPipelinesLifecycleStateEnum = map[string]ListPipelinesLifecycleStateEnum{
	"CREATING": ListPipelinesLifecycleStateCreating,
	"ACTIVE":   ListPipelinesLifecycleStateActive,
	"DELETING": ListPipelinesLifecycleStateDeleting,
	"FAILED":   ListPipelinesLifecycleStateFailed,
	"DELETED":  ListPipelinesLifecycleStateDeleted,
}

var mappingListPipelinesLifecycleStateEnumLowerCase = map[string]ListPipelinesLifecycleStateEnum{
	"creating": ListPipelinesLifecycleStateCreating,
	"active":   ListPipelinesLifecycleStateActive,
	"deleting": ListPipelinesLifecycleStateDeleting,
	"failed":   ListPipelinesLifecycleStateFailed,
	"deleted":  ListPipelinesLifecycleStateDeleted,
}

// GetListPipelinesLifecycleStateEnumValues Enumerates the set of values for ListPipelinesLifecycleStateEnum
func GetListPipelinesLifecycleStateEnumValues() []ListPipelinesLifecycleStateEnum {
	values := make([]ListPipelinesLifecycleStateEnum, 0)
	for _, v := range mappingListPipelinesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelinesLifecycleStateEnumStringValues Enumerates the set of values in String for ListPipelinesLifecycleStateEnum
func GetListPipelinesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
		"DELETED",
	}
}

// GetMappingListPipelinesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelinesLifecycleStateEnum(val string) (ListPipelinesLifecycleStateEnum, bool) {
	enum, ok := mappingListPipelinesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPipelinesSortOrderEnum Enum with underlying type: string
type ListPipelinesSortOrderEnum string

// Set of constants representing the allowable values for ListPipelinesSortOrderEnum
const (
	ListPipelinesSortOrderAsc  ListPipelinesSortOrderEnum = "ASC"
	ListPipelinesSortOrderDesc ListPipelinesSortOrderEnum = "DESC"
)

var mappingListPipelinesSortOrderEnum = map[string]ListPipelinesSortOrderEnum{
	"ASC":  ListPipelinesSortOrderAsc,
	"DESC": ListPipelinesSortOrderDesc,
}

var mappingListPipelinesSortOrderEnumLowerCase = map[string]ListPipelinesSortOrderEnum{
	"asc":  ListPipelinesSortOrderAsc,
	"desc": ListPipelinesSortOrderDesc,
}

// GetListPipelinesSortOrderEnumValues Enumerates the set of values for ListPipelinesSortOrderEnum
func GetListPipelinesSortOrderEnumValues() []ListPipelinesSortOrderEnum {
	values := make([]ListPipelinesSortOrderEnum, 0)
	for _, v := range mappingListPipelinesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelinesSortOrderEnumStringValues Enumerates the set of values in String for ListPipelinesSortOrderEnum
func GetListPipelinesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPipelinesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelinesSortOrderEnum(val string) (ListPipelinesSortOrderEnum, bool) {
	enum, ok := mappingListPipelinesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPipelinesSortByEnum Enum with underlying type: string
type ListPipelinesSortByEnum string

// Set of constants representing the allowable values for ListPipelinesSortByEnum
const (
	ListPipelinesSortByTimecreated ListPipelinesSortByEnum = "timeCreated"
	ListPipelinesSortByDisplayname ListPipelinesSortByEnum = "displayName"
)

var mappingListPipelinesSortByEnum = map[string]ListPipelinesSortByEnum{
	"timeCreated": ListPipelinesSortByTimecreated,
	"displayName": ListPipelinesSortByDisplayname,
}

var mappingListPipelinesSortByEnumLowerCase = map[string]ListPipelinesSortByEnum{
	"timecreated": ListPipelinesSortByTimecreated,
	"displayname": ListPipelinesSortByDisplayname,
}

// GetListPipelinesSortByEnumValues Enumerates the set of values for ListPipelinesSortByEnum
func GetListPipelinesSortByEnumValues() []ListPipelinesSortByEnum {
	values := make([]ListPipelinesSortByEnum, 0)
	for _, v := range mappingListPipelinesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelinesSortByEnumStringValues Enumerates the set of values in String for ListPipelinesSortByEnum
func GetListPipelinesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPipelinesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelinesSortByEnum(val string) (ListPipelinesSortByEnum, bool) {
	enum, ok := mappingListPipelinesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
