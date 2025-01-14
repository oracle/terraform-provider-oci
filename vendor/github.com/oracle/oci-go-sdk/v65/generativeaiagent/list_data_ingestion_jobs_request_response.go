// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeaiagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDataIngestionJobsRequest wrapper for the ListDataIngestionJobs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiagent/ListDataIngestionJobs.go.html to see an example of how to use ListDataIngestionJobsRequest.
type ListDataIngestionJobsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the data source.
	DataSourceId *string `mandatory:"false" contributesTo:"query" name:"dataSourceId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState DataIngestionJobLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDataIngestionJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListDataIngestionJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataIngestionJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataIngestionJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataIngestionJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataIngestionJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataIngestionJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataIngestionJobLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDataIngestionJobLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataIngestionJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataIngestionJobsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataIngestionJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataIngestionJobsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataIngestionJobsResponse wrapper for the ListDataIngestionJobs operation
type ListDataIngestionJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataIngestionJobCollection instances
	DataIngestionJobCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataIngestionJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataIngestionJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataIngestionJobsSortOrderEnum Enum with underlying type: string
type ListDataIngestionJobsSortOrderEnum string

// Set of constants representing the allowable values for ListDataIngestionJobsSortOrderEnum
const (
	ListDataIngestionJobsSortOrderAsc  ListDataIngestionJobsSortOrderEnum = "ASC"
	ListDataIngestionJobsSortOrderDesc ListDataIngestionJobsSortOrderEnum = "DESC"
)

var mappingListDataIngestionJobsSortOrderEnum = map[string]ListDataIngestionJobsSortOrderEnum{
	"ASC":  ListDataIngestionJobsSortOrderAsc,
	"DESC": ListDataIngestionJobsSortOrderDesc,
}

var mappingListDataIngestionJobsSortOrderEnumLowerCase = map[string]ListDataIngestionJobsSortOrderEnum{
	"asc":  ListDataIngestionJobsSortOrderAsc,
	"desc": ListDataIngestionJobsSortOrderDesc,
}

// GetListDataIngestionJobsSortOrderEnumValues Enumerates the set of values for ListDataIngestionJobsSortOrderEnum
func GetListDataIngestionJobsSortOrderEnumValues() []ListDataIngestionJobsSortOrderEnum {
	values := make([]ListDataIngestionJobsSortOrderEnum, 0)
	for _, v := range mappingListDataIngestionJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataIngestionJobsSortOrderEnumStringValues Enumerates the set of values in String for ListDataIngestionJobsSortOrderEnum
func GetListDataIngestionJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataIngestionJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataIngestionJobsSortOrderEnum(val string) (ListDataIngestionJobsSortOrderEnum, bool) {
	enum, ok := mappingListDataIngestionJobsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataIngestionJobsSortByEnum Enum with underlying type: string
type ListDataIngestionJobsSortByEnum string

// Set of constants representing the allowable values for ListDataIngestionJobsSortByEnum
const (
	ListDataIngestionJobsSortByTimecreated ListDataIngestionJobsSortByEnum = "timeCreated"
	ListDataIngestionJobsSortByDisplayname ListDataIngestionJobsSortByEnum = "displayName"
)

var mappingListDataIngestionJobsSortByEnum = map[string]ListDataIngestionJobsSortByEnum{
	"timeCreated": ListDataIngestionJobsSortByTimecreated,
	"displayName": ListDataIngestionJobsSortByDisplayname,
}

var mappingListDataIngestionJobsSortByEnumLowerCase = map[string]ListDataIngestionJobsSortByEnum{
	"timecreated": ListDataIngestionJobsSortByTimecreated,
	"displayname": ListDataIngestionJobsSortByDisplayname,
}

// GetListDataIngestionJobsSortByEnumValues Enumerates the set of values for ListDataIngestionJobsSortByEnum
func GetListDataIngestionJobsSortByEnumValues() []ListDataIngestionJobsSortByEnum {
	values := make([]ListDataIngestionJobsSortByEnum, 0)
	for _, v := range mappingListDataIngestionJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataIngestionJobsSortByEnumStringValues Enumerates the set of values in String for ListDataIngestionJobsSortByEnum
func GetListDataIngestionJobsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDataIngestionJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataIngestionJobsSortByEnum(val string) (ListDataIngestionJobsSortByEnum, bool) {
	enum, ok := mappingListDataIngestionJobsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
