// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package batch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBatchJobTasksRequest wrapper for the ListBatchJobTasks operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/batch/ListBatchJobTasks.go.html to see an example of how to use ListBatchJobTasksRequest.
type ListBatchJobTasksRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch job.
	BatchJobId *string `mandatory:"true" contributesTo:"path" name:"batchJobId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState BatchTaskLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The name of the task.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListBatchJobTasksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `name` is ascending.
	SortBy ListBatchJobTasksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBatchJobTasksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBatchJobTasksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBatchJobTasksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBatchJobTasksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBatchJobTasksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBatchTaskLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBatchTaskLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBatchJobTasksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBatchJobTasksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBatchJobTasksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBatchJobTasksSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBatchJobTasksResponse wrapper for the ListBatchJobTasks operation
type ListBatchJobTasksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BatchTaskCollection instances
	BatchTaskCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBatchJobTasksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBatchJobTasksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBatchJobTasksSortOrderEnum Enum with underlying type: string
type ListBatchJobTasksSortOrderEnum string

// Set of constants representing the allowable values for ListBatchJobTasksSortOrderEnum
const (
	ListBatchJobTasksSortOrderAsc  ListBatchJobTasksSortOrderEnum = "ASC"
	ListBatchJobTasksSortOrderDesc ListBatchJobTasksSortOrderEnum = "DESC"
)

var mappingListBatchJobTasksSortOrderEnum = map[string]ListBatchJobTasksSortOrderEnum{
	"ASC":  ListBatchJobTasksSortOrderAsc,
	"DESC": ListBatchJobTasksSortOrderDesc,
}

var mappingListBatchJobTasksSortOrderEnumLowerCase = map[string]ListBatchJobTasksSortOrderEnum{
	"asc":  ListBatchJobTasksSortOrderAsc,
	"desc": ListBatchJobTasksSortOrderDesc,
}

// GetListBatchJobTasksSortOrderEnumValues Enumerates the set of values for ListBatchJobTasksSortOrderEnum
func GetListBatchJobTasksSortOrderEnumValues() []ListBatchJobTasksSortOrderEnum {
	values := make([]ListBatchJobTasksSortOrderEnum, 0)
	for _, v := range mappingListBatchJobTasksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBatchJobTasksSortOrderEnumStringValues Enumerates the set of values in String for ListBatchJobTasksSortOrderEnum
func GetListBatchJobTasksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBatchJobTasksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBatchJobTasksSortOrderEnum(val string) (ListBatchJobTasksSortOrderEnum, bool) {
	enum, ok := mappingListBatchJobTasksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBatchJobTasksSortByEnum Enum with underlying type: string
type ListBatchJobTasksSortByEnum string

// Set of constants representing the allowable values for ListBatchJobTasksSortByEnum
const (
	ListBatchJobTasksSortByName ListBatchJobTasksSortByEnum = "name"
)

var mappingListBatchJobTasksSortByEnum = map[string]ListBatchJobTasksSortByEnum{
	"name": ListBatchJobTasksSortByName,
}

var mappingListBatchJobTasksSortByEnumLowerCase = map[string]ListBatchJobTasksSortByEnum{
	"name": ListBatchJobTasksSortByName,
}

// GetListBatchJobTasksSortByEnumValues Enumerates the set of values for ListBatchJobTasksSortByEnum
func GetListBatchJobTasksSortByEnumValues() []ListBatchJobTasksSortByEnum {
	values := make([]ListBatchJobTasksSortByEnum, 0)
	for _, v := range mappingListBatchJobTasksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBatchJobTasksSortByEnumStringValues Enumerates the set of values in String for ListBatchJobTasksSortByEnum
func GetListBatchJobTasksSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListBatchJobTasksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBatchJobTasksSortByEnum(val string) (ListBatchJobTasksSortByEnum, bool) {
	enum, ok := mappingListBatchJobTasksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
