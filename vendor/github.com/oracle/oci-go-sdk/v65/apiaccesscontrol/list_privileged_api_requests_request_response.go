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

// ListPrivilegedApiRequestsRequest wrapper for the ListPrivilegedApiRequests operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/ListPrivilegedApiRequests.go.html to see an example of how to use ListPrivilegedApiRequestsRequest.
type ListPrivilegedApiRequestsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the PrivilegedApiRequest.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource .
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// A filter to return only lists of resources that match the entire given service type.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState PrivilegedApiRequestLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the state. The
	// state value is case-insensitive.
	State PrivilegedApiRequestStateEnum `mandatory:"false" contributesTo:"query" name:"state" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListPrivilegedApiRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListPrivilegedApiRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPrivilegedApiRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPrivilegedApiRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPrivilegedApiRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPrivilegedApiRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPrivilegedApiRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrivilegedApiRequestLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPrivilegedApiRequestLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPrivilegedApiRequestStateEnum(string(request.State)); !ok && request.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", request.State, strings.Join(GetPrivilegedApiRequestStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPrivilegedApiRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPrivilegedApiRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPrivilegedApiRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPrivilegedApiRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPrivilegedApiRequestsResponse wrapper for the ListPrivilegedApiRequests operation
type ListPrivilegedApiRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PrivilegedApiRequestCollection instances
	PrivilegedApiRequestCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPrivilegedApiRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPrivilegedApiRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPrivilegedApiRequestsSortOrderEnum Enum with underlying type: string
type ListPrivilegedApiRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListPrivilegedApiRequestsSortOrderEnum
const (
	ListPrivilegedApiRequestsSortOrderAsc  ListPrivilegedApiRequestsSortOrderEnum = "ASC"
	ListPrivilegedApiRequestsSortOrderDesc ListPrivilegedApiRequestsSortOrderEnum = "DESC"
)

var mappingListPrivilegedApiRequestsSortOrderEnum = map[string]ListPrivilegedApiRequestsSortOrderEnum{
	"ASC":  ListPrivilegedApiRequestsSortOrderAsc,
	"DESC": ListPrivilegedApiRequestsSortOrderDesc,
}

var mappingListPrivilegedApiRequestsSortOrderEnumLowerCase = map[string]ListPrivilegedApiRequestsSortOrderEnum{
	"asc":  ListPrivilegedApiRequestsSortOrderAsc,
	"desc": ListPrivilegedApiRequestsSortOrderDesc,
}

// GetListPrivilegedApiRequestsSortOrderEnumValues Enumerates the set of values for ListPrivilegedApiRequestsSortOrderEnum
func GetListPrivilegedApiRequestsSortOrderEnumValues() []ListPrivilegedApiRequestsSortOrderEnum {
	values := make([]ListPrivilegedApiRequestsSortOrderEnum, 0)
	for _, v := range mappingListPrivilegedApiRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPrivilegedApiRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListPrivilegedApiRequestsSortOrderEnum
func GetListPrivilegedApiRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPrivilegedApiRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPrivilegedApiRequestsSortOrderEnum(val string) (ListPrivilegedApiRequestsSortOrderEnum, bool) {
	enum, ok := mappingListPrivilegedApiRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPrivilegedApiRequestsSortByEnum Enum with underlying type: string
type ListPrivilegedApiRequestsSortByEnum string

// Set of constants representing the allowable values for ListPrivilegedApiRequestsSortByEnum
const (
	ListPrivilegedApiRequestsSortByTimecreated ListPrivilegedApiRequestsSortByEnum = "timeCreated"
	ListPrivilegedApiRequestsSortByDisplayname ListPrivilegedApiRequestsSortByEnum = "displayName"
)

var mappingListPrivilegedApiRequestsSortByEnum = map[string]ListPrivilegedApiRequestsSortByEnum{
	"timeCreated": ListPrivilegedApiRequestsSortByTimecreated,
	"displayName": ListPrivilegedApiRequestsSortByDisplayname,
}

var mappingListPrivilegedApiRequestsSortByEnumLowerCase = map[string]ListPrivilegedApiRequestsSortByEnum{
	"timecreated": ListPrivilegedApiRequestsSortByTimecreated,
	"displayname": ListPrivilegedApiRequestsSortByDisplayname,
}

// GetListPrivilegedApiRequestsSortByEnumValues Enumerates the set of values for ListPrivilegedApiRequestsSortByEnum
func GetListPrivilegedApiRequestsSortByEnumValues() []ListPrivilegedApiRequestsSortByEnum {
	values := make([]ListPrivilegedApiRequestsSortByEnum, 0)
	for _, v := range mappingListPrivilegedApiRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPrivilegedApiRequestsSortByEnumStringValues Enumerates the set of values in String for ListPrivilegedApiRequestsSortByEnum
func GetListPrivilegedApiRequestsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPrivilegedApiRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPrivilegedApiRequestsSortByEnum(val string) (ListPrivilegedApiRequestsSortByEnum, bool) {
	enum, ok := mappingListPrivilegedApiRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
