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

// ListBatchJobsRequest wrapper for the ListBatchJobs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/batch/ListBatchJobs.go.html to see an example of how to use ListBatchJobsRequest.
type ListBatchJobsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState BatchJobLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch job.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch job pool.
	BatchJobPoolId *string `mandatory:"false" contributesTo:"query" name:"batchJobPoolId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListBatchJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListBatchJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBatchJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBatchJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBatchJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBatchJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBatchJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBatchJobLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBatchJobLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBatchJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBatchJobsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBatchJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBatchJobsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBatchJobsResponse wrapper for the ListBatchJobs operation
type ListBatchJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BatchJobCollection instances
	BatchJobCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBatchJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBatchJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBatchJobsSortOrderEnum Enum with underlying type: string
type ListBatchJobsSortOrderEnum string

// Set of constants representing the allowable values for ListBatchJobsSortOrderEnum
const (
	ListBatchJobsSortOrderAsc  ListBatchJobsSortOrderEnum = "ASC"
	ListBatchJobsSortOrderDesc ListBatchJobsSortOrderEnum = "DESC"
)

var mappingListBatchJobsSortOrderEnum = map[string]ListBatchJobsSortOrderEnum{
	"ASC":  ListBatchJobsSortOrderAsc,
	"DESC": ListBatchJobsSortOrderDesc,
}

var mappingListBatchJobsSortOrderEnumLowerCase = map[string]ListBatchJobsSortOrderEnum{
	"asc":  ListBatchJobsSortOrderAsc,
	"desc": ListBatchJobsSortOrderDesc,
}

// GetListBatchJobsSortOrderEnumValues Enumerates the set of values for ListBatchJobsSortOrderEnum
func GetListBatchJobsSortOrderEnumValues() []ListBatchJobsSortOrderEnum {
	values := make([]ListBatchJobsSortOrderEnum, 0)
	for _, v := range mappingListBatchJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBatchJobsSortOrderEnumStringValues Enumerates the set of values in String for ListBatchJobsSortOrderEnum
func GetListBatchJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBatchJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBatchJobsSortOrderEnum(val string) (ListBatchJobsSortOrderEnum, bool) {
	enum, ok := mappingListBatchJobsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBatchJobsSortByEnum Enum with underlying type: string
type ListBatchJobsSortByEnum string

// Set of constants representing the allowable values for ListBatchJobsSortByEnum
const (
	ListBatchJobsSortByTimecreated ListBatchJobsSortByEnum = "timeCreated"
	ListBatchJobsSortByDisplayname ListBatchJobsSortByEnum = "displayName"
)

var mappingListBatchJobsSortByEnum = map[string]ListBatchJobsSortByEnum{
	"timeCreated": ListBatchJobsSortByTimecreated,
	"displayName": ListBatchJobsSortByDisplayname,
}

var mappingListBatchJobsSortByEnumLowerCase = map[string]ListBatchJobsSortByEnum{
	"timecreated": ListBatchJobsSortByTimecreated,
	"displayname": ListBatchJobsSortByDisplayname,
}

// GetListBatchJobsSortByEnumValues Enumerates the set of values for ListBatchJobsSortByEnum
func GetListBatchJobsSortByEnumValues() []ListBatchJobsSortByEnum {
	values := make([]ListBatchJobsSortByEnum, 0)
	for _, v := range mappingListBatchJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBatchJobsSortByEnumStringValues Enumerates the set of values in String for ListBatchJobsSortByEnum
func GetListBatchJobsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListBatchJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBatchJobsSortByEnum(val string) (ListBatchJobsSortByEnum, bool) {
	enum, ok := mappingListBatchJobsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
