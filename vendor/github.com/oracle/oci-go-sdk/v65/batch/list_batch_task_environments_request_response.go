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

// ListBatchTaskEnvironmentsRequest wrapper for the ListBatchTaskEnvironments operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/batch/ListBatchTaskEnvironments.go.html to see an example of how to use ListBatchTaskEnvironmentsRequest.
type ListBatchTaskEnvironmentsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState BatchTaskEnvironmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch task environment.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListBatchTaskEnvironmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListBatchTaskEnvironmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBatchTaskEnvironmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBatchTaskEnvironmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBatchTaskEnvironmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBatchTaskEnvironmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBatchTaskEnvironmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBatchTaskEnvironmentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBatchTaskEnvironmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBatchTaskEnvironmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBatchTaskEnvironmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBatchTaskEnvironmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBatchTaskEnvironmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBatchTaskEnvironmentsResponse wrapper for the ListBatchTaskEnvironments operation
type ListBatchTaskEnvironmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BatchTaskEnvironmentCollection instances
	BatchTaskEnvironmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBatchTaskEnvironmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBatchTaskEnvironmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBatchTaskEnvironmentsSortOrderEnum Enum with underlying type: string
type ListBatchTaskEnvironmentsSortOrderEnum string

// Set of constants representing the allowable values for ListBatchTaskEnvironmentsSortOrderEnum
const (
	ListBatchTaskEnvironmentsSortOrderAsc  ListBatchTaskEnvironmentsSortOrderEnum = "ASC"
	ListBatchTaskEnvironmentsSortOrderDesc ListBatchTaskEnvironmentsSortOrderEnum = "DESC"
)

var mappingListBatchTaskEnvironmentsSortOrderEnum = map[string]ListBatchTaskEnvironmentsSortOrderEnum{
	"ASC":  ListBatchTaskEnvironmentsSortOrderAsc,
	"DESC": ListBatchTaskEnvironmentsSortOrderDesc,
}

var mappingListBatchTaskEnvironmentsSortOrderEnumLowerCase = map[string]ListBatchTaskEnvironmentsSortOrderEnum{
	"asc":  ListBatchTaskEnvironmentsSortOrderAsc,
	"desc": ListBatchTaskEnvironmentsSortOrderDesc,
}

// GetListBatchTaskEnvironmentsSortOrderEnumValues Enumerates the set of values for ListBatchTaskEnvironmentsSortOrderEnum
func GetListBatchTaskEnvironmentsSortOrderEnumValues() []ListBatchTaskEnvironmentsSortOrderEnum {
	values := make([]ListBatchTaskEnvironmentsSortOrderEnum, 0)
	for _, v := range mappingListBatchTaskEnvironmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBatchTaskEnvironmentsSortOrderEnumStringValues Enumerates the set of values in String for ListBatchTaskEnvironmentsSortOrderEnum
func GetListBatchTaskEnvironmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBatchTaskEnvironmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBatchTaskEnvironmentsSortOrderEnum(val string) (ListBatchTaskEnvironmentsSortOrderEnum, bool) {
	enum, ok := mappingListBatchTaskEnvironmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBatchTaskEnvironmentsSortByEnum Enum with underlying type: string
type ListBatchTaskEnvironmentsSortByEnum string

// Set of constants representing the allowable values for ListBatchTaskEnvironmentsSortByEnum
const (
	ListBatchTaskEnvironmentsSortByTimecreated ListBatchTaskEnvironmentsSortByEnum = "timeCreated"
	ListBatchTaskEnvironmentsSortByDisplayname ListBatchTaskEnvironmentsSortByEnum = "displayName"
)

var mappingListBatchTaskEnvironmentsSortByEnum = map[string]ListBatchTaskEnvironmentsSortByEnum{
	"timeCreated": ListBatchTaskEnvironmentsSortByTimecreated,
	"displayName": ListBatchTaskEnvironmentsSortByDisplayname,
}

var mappingListBatchTaskEnvironmentsSortByEnumLowerCase = map[string]ListBatchTaskEnvironmentsSortByEnum{
	"timecreated": ListBatchTaskEnvironmentsSortByTimecreated,
	"displayname": ListBatchTaskEnvironmentsSortByDisplayname,
}

// GetListBatchTaskEnvironmentsSortByEnumValues Enumerates the set of values for ListBatchTaskEnvironmentsSortByEnum
func GetListBatchTaskEnvironmentsSortByEnumValues() []ListBatchTaskEnvironmentsSortByEnum {
	values := make([]ListBatchTaskEnvironmentsSortByEnum, 0)
	for _, v := range mappingListBatchTaskEnvironmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBatchTaskEnvironmentsSortByEnumStringValues Enumerates the set of values in String for ListBatchTaskEnvironmentsSortByEnum
func GetListBatchTaskEnvironmentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListBatchTaskEnvironmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBatchTaskEnvironmentsSortByEnum(val string) (ListBatchTaskEnvironmentsSortByEnum, bool) {
	enum, ok := mappingListBatchTaskEnvironmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
