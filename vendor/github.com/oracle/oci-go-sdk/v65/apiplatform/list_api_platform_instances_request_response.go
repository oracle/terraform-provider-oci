// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apiplatform

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListApiPlatformInstancesRequest wrapper for the ListApiPlatformInstances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiplatform/ListApiPlatformInstances.go.html to see an example of how to use ListApiPlatformInstancesRequest.
type ListApiPlatformInstancesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState ApiPlatformInstanceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given name exactly
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance
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
	SortOrder ListApiPlatformInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `name` is ascending.
	SortBy ListApiPlatformInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListApiPlatformInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApiPlatformInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListApiPlatformInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApiPlatformInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListApiPlatformInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApiPlatformInstanceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetApiPlatformInstanceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApiPlatformInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListApiPlatformInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApiPlatformInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListApiPlatformInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListApiPlatformInstancesResponse wrapper for the ListApiPlatformInstances operation
type ListApiPlatformInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ApiPlatformInstanceCollection instances
	ApiPlatformInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListApiPlatformInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApiPlatformInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApiPlatformInstancesSortOrderEnum Enum with underlying type: string
type ListApiPlatformInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListApiPlatformInstancesSortOrderEnum
const (
	ListApiPlatformInstancesSortOrderAsc  ListApiPlatformInstancesSortOrderEnum = "ASC"
	ListApiPlatformInstancesSortOrderDesc ListApiPlatformInstancesSortOrderEnum = "DESC"
)

var mappingListApiPlatformInstancesSortOrderEnum = map[string]ListApiPlatformInstancesSortOrderEnum{
	"ASC":  ListApiPlatformInstancesSortOrderAsc,
	"DESC": ListApiPlatformInstancesSortOrderDesc,
}

var mappingListApiPlatformInstancesSortOrderEnumLowerCase = map[string]ListApiPlatformInstancesSortOrderEnum{
	"asc":  ListApiPlatformInstancesSortOrderAsc,
	"desc": ListApiPlatformInstancesSortOrderDesc,
}

// GetListApiPlatformInstancesSortOrderEnumValues Enumerates the set of values for ListApiPlatformInstancesSortOrderEnum
func GetListApiPlatformInstancesSortOrderEnumValues() []ListApiPlatformInstancesSortOrderEnum {
	values := make([]ListApiPlatformInstancesSortOrderEnum, 0)
	for _, v := range mappingListApiPlatformInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListApiPlatformInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListApiPlatformInstancesSortOrderEnum
func GetListApiPlatformInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListApiPlatformInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApiPlatformInstancesSortOrderEnum(val string) (ListApiPlatformInstancesSortOrderEnum, bool) {
	enum, ok := mappingListApiPlatformInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListApiPlatformInstancesSortByEnum Enum with underlying type: string
type ListApiPlatformInstancesSortByEnum string

// Set of constants representing the allowable values for ListApiPlatformInstancesSortByEnum
const (
	ListApiPlatformInstancesSortByName        ListApiPlatformInstancesSortByEnum = "name"
	ListApiPlatformInstancesSortByTimecreated ListApiPlatformInstancesSortByEnum = "timeCreated"
)

var mappingListApiPlatformInstancesSortByEnum = map[string]ListApiPlatformInstancesSortByEnum{
	"name":        ListApiPlatformInstancesSortByName,
	"timeCreated": ListApiPlatformInstancesSortByTimecreated,
}

var mappingListApiPlatformInstancesSortByEnumLowerCase = map[string]ListApiPlatformInstancesSortByEnum{
	"name":        ListApiPlatformInstancesSortByName,
	"timecreated": ListApiPlatformInstancesSortByTimecreated,
}

// GetListApiPlatformInstancesSortByEnumValues Enumerates the set of values for ListApiPlatformInstancesSortByEnum
func GetListApiPlatformInstancesSortByEnumValues() []ListApiPlatformInstancesSortByEnum {
	values := make([]ListApiPlatformInstancesSortByEnum, 0)
	for _, v := range mappingListApiPlatformInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListApiPlatformInstancesSortByEnumStringValues Enumerates the set of values in String for ListApiPlatformInstancesSortByEnum
func GetListApiPlatformInstancesSortByEnumStringValues() []string {
	return []string{
		"name",
		"timeCreated",
	}
}

// GetMappingListApiPlatformInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApiPlatformInstancesSortByEnum(val string) (ListApiPlatformInstancesSortByEnum, bool) {
	enum, ok := mappingListApiPlatformInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
