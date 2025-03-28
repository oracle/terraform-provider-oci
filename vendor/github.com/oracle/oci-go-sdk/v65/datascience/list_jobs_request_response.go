// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListJobsRequest wrapper for the ListJobs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListJobs.go.html to see an example of how to use ListJobsRequest.
type ListJobsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// <b>Filter</b> results by OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	//   state for the resource type.
	LifecycleState ListJobsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 100 is the maximum.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, the results are shown
	// in descending order. When you sort by `displayName`, the results are
	// shown in ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJobsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListJobsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJobsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJobsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJobsResponse wrapper for the ListJobs operation
type ListJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []JobSummary instances
	Items []JobSummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJobsLifecycleStateEnum Enum with underlying type: string
type ListJobsLifecycleStateEnum string

// Set of constants representing the allowable values for ListJobsLifecycleStateEnum
const (
	ListJobsLifecycleStateCreating ListJobsLifecycleStateEnum = "CREATING"
	ListJobsLifecycleStateActive   ListJobsLifecycleStateEnum = "ACTIVE"
	ListJobsLifecycleStateDeleting ListJobsLifecycleStateEnum = "DELETING"
	ListJobsLifecycleStateFailed   ListJobsLifecycleStateEnum = "FAILED"
	ListJobsLifecycleStateDeleted  ListJobsLifecycleStateEnum = "DELETED"
)

var mappingListJobsLifecycleStateEnum = map[string]ListJobsLifecycleStateEnum{
	"CREATING": ListJobsLifecycleStateCreating,
	"ACTIVE":   ListJobsLifecycleStateActive,
	"DELETING": ListJobsLifecycleStateDeleting,
	"FAILED":   ListJobsLifecycleStateFailed,
	"DELETED":  ListJobsLifecycleStateDeleted,
}

var mappingListJobsLifecycleStateEnumLowerCase = map[string]ListJobsLifecycleStateEnum{
	"creating": ListJobsLifecycleStateCreating,
	"active":   ListJobsLifecycleStateActive,
	"deleting": ListJobsLifecycleStateDeleting,
	"failed":   ListJobsLifecycleStateFailed,
	"deleted":  ListJobsLifecycleStateDeleted,
}

// GetListJobsLifecycleStateEnumValues Enumerates the set of values for ListJobsLifecycleStateEnum
func GetListJobsLifecycleStateEnumValues() []ListJobsLifecycleStateEnum {
	values := make([]ListJobsLifecycleStateEnum, 0)
	for _, v := range mappingListJobsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobsLifecycleStateEnumStringValues Enumerates the set of values in String for ListJobsLifecycleStateEnum
func GetListJobsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
		"DELETED",
	}
}

// GetMappingListJobsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobsLifecycleStateEnum(val string) (ListJobsLifecycleStateEnum, bool) {
	enum, ok := mappingListJobsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobsSortOrderEnum Enum with underlying type: string
type ListJobsSortOrderEnum string

// Set of constants representing the allowable values for ListJobsSortOrderEnum
const (
	ListJobsSortOrderAsc  ListJobsSortOrderEnum = "ASC"
	ListJobsSortOrderDesc ListJobsSortOrderEnum = "DESC"
)

var mappingListJobsSortOrderEnum = map[string]ListJobsSortOrderEnum{
	"ASC":  ListJobsSortOrderAsc,
	"DESC": ListJobsSortOrderDesc,
}

var mappingListJobsSortOrderEnumLowerCase = map[string]ListJobsSortOrderEnum{
	"asc":  ListJobsSortOrderAsc,
	"desc": ListJobsSortOrderDesc,
}

// GetListJobsSortOrderEnumValues Enumerates the set of values for ListJobsSortOrderEnum
func GetListJobsSortOrderEnumValues() []ListJobsSortOrderEnum {
	values := make([]ListJobsSortOrderEnum, 0)
	for _, v := range mappingListJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobsSortOrderEnumStringValues Enumerates the set of values in String for ListJobsSortOrderEnum
func GetListJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobsSortOrderEnum(val string) (ListJobsSortOrderEnum, bool) {
	enum, ok := mappingListJobsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJobsSortByEnum Enum with underlying type: string
type ListJobsSortByEnum string

// Set of constants representing the allowable values for ListJobsSortByEnum
const (
	ListJobsSortByTimecreated ListJobsSortByEnum = "timeCreated"
	ListJobsSortByDisplayname ListJobsSortByEnum = "displayName"
)

var mappingListJobsSortByEnum = map[string]ListJobsSortByEnum{
	"timeCreated": ListJobsSortByTimecreated,
	"displayName": ListJobsSortByDisplayname,
}

var mappingListJobsSortByEnumLowerCase = map[string]ListJobsSortByEnum{
	"timecreated": ListJobsSortByTimecreated,
	"displayname": ListJobsSortByDisplayname,
}

// GetListJobsSortByEnumValues Enumerates the set of values for ListJobsSortByEnum
func GetListJobsSortByEnumValues() []ListJobsSortByEnum {
	values := make([]ListJobsSortByEnum, 0)
	for _, v := range mappingListJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJobsSortByEnumStringValues Enumerates the set of values in String for ListJobsSortByEnum
func GetListJobsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJobsSortByEnum(val string) (ListJobsSortByEnum, bool) {
	enum, ok := mappingListJobsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
