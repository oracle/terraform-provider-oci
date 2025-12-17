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

// ListBatchJobPoolsRequest wrapper for the ListBatchJobPools operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/batch/ListBatchJobPools.go.html to see an example of how to use ListBatchJobPoolsRequest.
type ListBatchJobPoolsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState BatchJobPoolLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch job pool.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch context.
	BatchContextId *string `mandatory:"false" contributesTo:"query" name:"batchContextId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListBatchJobPoolsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListBatchJobPoolsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBatchJobPoolsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBatchJobPoolsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBatchJobPoolsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBatchJobPoolsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBatchJobPoolsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBatchJobPoolLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBatchJobPoolLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBatchJobPoolsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBatchJobPoolsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBatchJobPoolsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBatchJobPoolsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBatchJobPoolsResponse wrapper for the ListBatchJobPools operation
type ListBatchJobPoolsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BatchJobPoolCollection instances
	BatchJobPoolCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBatchJobPoolsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBatchJobPoolsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBatchJobPoolsSortOrderEnum Enum with underlying type: string
type ListBatchJobPoolsSortOrderEnum string

// Set of constants representing the allowable values for ListBatchJobPoolsSortOrderEnum
const (
	ListBatchJobPoolsSortOrderAsc  ListBatchJobPoolsSortOrderEnum = "ASC"
	ListBatchJobPoolsSortOrderDesc ListBatchJobPoolsSortOrderEnum = "DESC"
)

var mappingListBatchJobPoolsSortOrderEnum = map[string]ListBatchJobPoolsSortOrderEnum{
	"ASC":  ListBatchJobPoolsSortOrderAsc,
	"DESC": ListBatchJobPoolsSortOrderDesc,
}

var mappingListBatchJobPoolsSortOrderEnumLowerCase = map[string]ListBatchJobPoolsSortOrderEnum{
	"asc":  ListBatchJobPoolsSortOrderAsc,
	"desc": ListBatchJobPoolsSortOrderDesc,
}

// GetListBatchJobPoolsSortOrderEnumValues Enumerates the set of values for ListBatchJobPoolsSortOrderEnum
func GetListBatchJobPoolsSortOrderEnumValues() []ListBatchJobPoolsSortOrderEnum {
	values := make([]ListBatchJobPoolsSortOrderEnum, 0)
	for _, v := range mappingListBatchJobPoolsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBatchJobPoolsSortOrderEnumStringValues Enumerates the set of values in String for ListBatchJobPoolsSortOrderEnum
func GetListBatchJobPoolsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBatchJobPoolsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBatchJobPoolsSortOrderEnum(val string) (ListBatchJobPoolsSortOrderEnum, bool) {
	enum, ok := mappingListBatchJobPoolsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBatchJobPoolsSortByEnum Enum with underlying type: string
type ListBatchJobPoolsSortByEnum string

// Set of constants representing the allowable values for ListBatchJobPoolsSortByEnum
const (
	ListBatchJobPoolsSortByTimecreated ListBatchJobPoolsSortByEnum = "timeCreated"
	ListBatchJobPoolsSortByDisplayname ListBatchJobPoolsSortByEnum = "displayName"
)

var mappingListBatchJobPoolsSortByEnum = map[string]ListBatchJobPoolsSortByEnum{
	"timeCreated": ListBatchJobPoolsSortByTimecreated,
	"displayName": ListBatchJobPoolsSortByDisplayname,
}

var mappingListBatchJobPoolsSortByEnumLowerCase = map[string]ListBatchJobPoolsSortByEnum{
	"timecreated": ListBatchJobPoolsSortByTimecreated,
	"displayname": ListBatchJobPoolsSortByDisplayname,
}

// GetListBatchJobPoolsSortByEnumValues Enumerates the set of values for ListBatchJobPoolsSortByEnum
func GetListBatchJobPoolsSortByEnumValues() []ListBatchJobPoolsSortByEnum {
	values := make([]ListBatchJobPoolsSortByEnum, 0)
	for _, v := range mappingListBatchJobPoolsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBatchJobPoolsSortByEnumStringValues Enumerates the set of values in String for ListBatchJobPoolsSortByEnum
func GetListBatchJobPoolsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListBatchJobPoolsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBatchJobPoolsSortByEnum(val string) (ListBatchJobPoolsSortByEnum, bool) {
	enum, ok := mappingListBatchJobPoolsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
