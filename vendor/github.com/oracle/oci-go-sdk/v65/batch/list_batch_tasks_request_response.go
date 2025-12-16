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

// ListBatchTasksRequest wrapper for the ListBatchTasks operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/batch/ListBatchTasks.go.html to see an example of how to use ListBatchTasksRequest.
type ListBatchTasksRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch job.
	BatchJobId *string `mandatory:"false" contributesTo:"query" name:"batchJobId"`

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
	SortOrder ListBatchTasksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `name` is ascending.
	SortBy ListBatchTasksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBatchTasksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBatchTasksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBatchTasksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBatchTasksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBatchTasksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBatchTaskLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBatchTaskLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBatchTasksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBatchTasksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBatchTasksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBatchTasksSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBatchTasksResponse wrapper for the ListBatchTasks operation
type ListBatchTasksResponse struct {

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

func (response ListBatchTasksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBatchTasksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBatchTasksSortOrderEnum Enum with underlying type: string
type ListBatchTasksSortOrderEnum string

// Set of constants representing the allowable values for ListBatchTasksSortOrderEnum
const (
	ListBatchTasksSortOrderAsc  ListBatchTasksSortOrderEnum = "ASC"
	ListBatchTasksSortOrderDesc ListBatchTasksSortOrderEnum = "DESC"
)

var mappingListBatchTasksSortOrderEnum = map[string]ListBatchTasksSortOrderEnum{
	"ASC":  ListBatchTasksSortOrderAsc,
	"DESC": ListBatchTasksSortOrderDesc,
}

var mappingListBatchTasksSortOrderEnumLowerCase = map[string]ListBatchTasksSortOrderEnum{
	"asc":  ListBatchTasksSortOrderAsc,
	"desc": ListBatchTasksSortOrderDesc,
}

// GetListBatchTasksSortOrderEnumValues Enumerates the set of values for ListBatchTasksSortOrderEnum
func GetListBatchTasksSortOrderEnumValues() []ListBatchTasksSortOrderEnum {
	values := make([]ListBatchTasksSortOrderEnum, 0)
	for _, v := range mappingListBatchTasksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBatchTasksSortOrderEnumStringValues Enumerates the set of values in String for ListBatchTasksSortOrderEnum
func GetListBatchTasksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBatchTasksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBatchTasksSortOrderEnum(val string) (ListBatchTasksSortOrderEnum, bool) {
	enum, ok := mappingListBatchTasksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBatchTasksSortByEnum Enum with underlying type: string
type ListBatchTasksSortByEnum string

// Set of constants representing the allowable values for ListBatchTasksSortByEnum
const (
	ListBatchTasksSortByName ListBatchTasksSortByEnum = "name"
)

var mappingListBatchTasksSortByEnum = map[string]ListBatchTasksSortByEnum{
	"name": ListBatchTasksSortByName,
}

var mappingListBatchTasksSortByEnumLowerCase = map[string]ListBatchTasksSortByEnum{
	"name": ListBatchTasksSortByName,
}

// GetListBatchTasksSortByEnumValues Enumerates the set of values for ListBatchTasksSortByEnum
func GetListBatchTasksSortByEnumValues() []ListBatchTasksSortByEnum {
	values := make([]ListBatchTasksSortByEnum, 0)
	for _, v := range mappingListBatchTasksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBatchTasksSortByEnumStringValues Enumerates the set of values in String for ListBatchTasksSortByEnum
func GetListBatchTasksSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListBatchTasksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBatchTasksSortByEnum(val string) (ListBatchTasksSortByEnum, bool) {
	enum, ok := mappingListBatchTasksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
