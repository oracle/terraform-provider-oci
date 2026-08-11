// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ddfs

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInstancesRequest wrapper for the ListInstances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ddfs/ListInstances.go.html to see an example of how to use ListInstancesRequest.
type ListInstancesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState InstanceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Instance.
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
	SortOrder ListInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInstanceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetInstanceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInstancesResponse wrapper for the ListInstances operation
type ListInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InstanceCollection instances
	InstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInstancesSortOrderEnum Enum with underlying type: string
type ListInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListInstancesSortOrderEnum
const (
	ListInstancesSortOrderAsc  ListInstancesSortOrderEnum = "ASC"
	ListInstancesSortOrderDesc ListInstancesSortOrderEnum = "DESC"
)

var mappingListInstancesSortOrderEnum = map[string]ListInstancesSortOrderEnum{
	"ASC":  ListInstancesSortOrderAsc,
	"DESC": ListInstancesSortOrderDesc,
}

var mappingListInstancesSortOrderEnumLowerCase = map[string]ListInstancesSortOrderEnum{
	"asc":  ListInstancesSortOrderAsc,
	"desc": ListInstancesSortOrderDesc,
}

// GetListInstancesSortOrderEnumValues Enumerates the set of values for ListInstancesSortOrderEnum
func GetListInstancesSortOrderEnumValues() []ListInstancesSortOrderEnum {
	values := make([]ListInstancesSortOrderEnum, 0)
	for _, v := range mappingListInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListInstancesSortOrderEnum
func GetListInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstancesSortOrderEnum(val string) (ListInstancesSortOrderEnum, bool) {
	enum, ok := mappingListInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInstancesSortByEnum Enum with underlying type: string
type ListInstancesSortByEnum string

// Set of constants representing the allowable values for ListInstancesSortByEnum
const (
	ListInstancesSortByTimecreated ListInstancesSortByEnum = "timeCreated"
	ListInstancesSortByDisplayname ListInstancesSortByEnum = "displayName"
)

var mappingListInstancesSortByEnum = map[string]ListInstancesSortByEnum{
	"timeCreated": ListInstancesSortByTimecreated,
	"displayName": ListInstancesSortByDisplayname,
}

var mappingListInstancesSortByEnumLowerCase = map[string]ListInstancesSortByEnum{
	"timecreated": ListInstancesSortByTimecreated,
	"displayname": ListInstancesSortByDisplayname,
}

// GetListInstancesSortByEnumValues Enumerates the set of values for ListInstancesSortByEnum
func GetListInstancesSortByEnumValues() []ListInstancesSortByEnum {
	values := make([]ListInstancesSortByEnum, 0)
	for _, v := range mappingListInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstancesSortByEnumStringValues Enumerates the set of values in String for ListInstancesSortByEnum
func GetListInstancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstancesSortByEnum(val string) (ListInstancesSortByEnum, bool) {
	enum, ok := mappingListInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
