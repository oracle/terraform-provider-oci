// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apiaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPrivilegedApiControlsRequest wrapper for the ListPrivilegedApiControls operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/ListPrivilegedApiControls.go.html to see an example of how to use ListPrivilegedApiControlsRequest.
type ListPrivilegedApiControlsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the PrivilegedApiControl.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState PrivilegedApiControlLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only lists of resources that match the entire given service type.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListPrivilegedApiControlsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListPrivilegedApiControlsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPrivilegedApiControlsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPrivilegedApiControlsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPrivilegedApiControlsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPrivilegedApiControlsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPrivilegedApiControlsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrivilegedApiControlLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPrivilegedApiControlLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPrivilegedApiControlsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPrivilegedApiControlsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPrivilegedApiControlsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPrivilegedApiControlsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPrivilegedApiControlsResponse wrapper for the ListPrivilegedApiControls operation
type ListPrivilegedApiControlsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PrivilegedApiControlCollection instances
	PrivilegedApiControlCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPrivilegedApiControlsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPrivilegedApiControlsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPrivilegedApiControlsSortOrderEnum Enum with underlying type: string
type ListPrivilegedApiControlsSortOrderEnum string

// Set of constants representing the allowable values for ListPrivilegedApiControlsSortOrderEnum
const (
	ListPrivilegedApiControlsSortOrderAsc  ListPrivilegedApiControlsSortOrderEnum = "ASC"
	ListPrivilegedApiControlsSortOrderDesc ListPrivilegedApiControlsSortOrderEnum = "DESC"
)

var mappingListPrivilegedApiControlsSortOrderEnum = map[string]ListPrivilegedApiControlsSortOrderEnum{
	"ASC":  ListPrivilegedApiControlsSortOrderAsc,
	"DESC": ListPrivilegedApiControlsSortOrderDesc,
}

var mappingListPrivilegedApiControlsSortOrderEnumLowerCase = map[string]ListPrivilegedApiControlsSortOrderEnum{
	"asc":  ListPrivilegedApiControlsSortOrderAsc,
	"desc": ListPrivilegedApiControlsSortOrderDesc,
}

// GetListPrivilegedApiControlsSortOrderEnumValues Enumerates the set of values for ListPrivilegedApiControlsSortOrderEnum
func GetListPrivilegedApiControlsSortOrderEnumValues() []ListPrivilegedApiControlsSortOrderEnum {
	values := make([]ListPrivilegedApiControlsSortOrderEnum, 0)
	for _, v := range mappingListPrivilegedApiControlsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPrivilegedApiControlsSortOrderEnumStringValues Enumerates the set of values in String for ListPrivilegedApiControlsSortOrderEnum
func GetListPrivilegedApiControlsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPrivilegedApiControlsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPrivilegedApiControlsSortOrderEnum(val string) (ListPrivilegedApiControlsSortOrderEnum, bool) {
	enum, ok := mappingListPrivilegedApiControlsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPrivilegedApiControlsSortByEnum Enum with underlying type: string
type ListPrivilegedApiControlsSortByEnum string

// Set of constants representing the allowable values for ListPrivilegedApiControlsSortByEnum
const (
	ListPrivilegedApiControlsSortByTimecreated ListPrivilegedApiControlsSortByEnum = "timeCreated"
	ListPrivilegedApiControlsSortByDisplayname ListPrivilegedApiControlsSortByEnum = "displayName"
)

var mappingListPrivilegedApiControlsSortByEnum = map[string]ListPrivilegedApiControlsSortByEnum{
	"timeCreated": ListPrivilegedApiControlsSortByTimecreated,
	"displayName": ListPrivilegedApiControlsSortByDisplayname,
}

var mappingListPrivilegedApiControlsSortByEnumLowerCase = map[string]ListPrivilegedApiControlsSortByEnum{
	"timecreated": ListPrivilegedApiControlsSortByTimecreated,
	"displayname": ListPrivilegedApiControlsSortByDisplayname,
}

// GetListPrivilegedApiControlsSortByEnumValues Enumerates the set of values for ListPrivilegedApiControlsSortByEnum
func GetListPrivilegedApiControlsSortByEnumValues() []ListPrivilegedApiControlsSortByEnum {
	values := make([]ListPrivilegedApiControlsSortByEnum, 0)
	for _, v := range mappingListPrivilegedApiControlsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPrivilegedApiControlsSortByEnumStringValues Enumerates the set of values in String for ListPrivilegedApiControlsSortByEnum
func GetListPrivilegedApiControlsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPrivilegedApiControlsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPrivilegedApiControlsSortByEnum(val string) (ListPrivilegedApiControlsSortByEnum, bool) {
	enum, ok := mappingListPrivilegedApiControlsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
