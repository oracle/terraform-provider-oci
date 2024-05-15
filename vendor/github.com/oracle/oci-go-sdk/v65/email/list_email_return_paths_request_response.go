// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEmailReturnPathsRequest wrapper for the ListEmailReturnPaths operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/email/ListEmailReturnPaths.go.html to see an example of how to use ListEmailReturnPathsRequest.
type ListEmailReturnPathsRequest struct {

	// The request ID for tracing from the system
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID for the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Email Domain to which this Email Return Path belongs.
	ParentResourceId *string `mandatory:"false" contributesTo:"query" name:"parentResourceId"`

	// A filter to only return resources that match the given id exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. `1` is the minimum, `1000` is the maximum. For important details about
	// how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending or descending order.
	SortOrder ListEmailReturnPathsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filter returned list by specified lifecycle state. This parameter is case-insensitive.
	LifecycleState EmailReturnPathLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies the attribute with which to sort the return paths.
	// Default: `TIMECREATED`
	// * **TIMECREATED:** Sorts by timeCreated.
	// * **NAME:** Sorts by name.
	// * **ID:** Sorts by id.
	SortBy ListEmailReturnPathsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEmailReturnPathsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEmailReturnPathsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEmailReturnPathsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEmailReturnPathsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEmailReturnPathsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEmailReturnPathsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEmailReturnPathsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmailReturnPathLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetEmailReturnPathLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEmailReturnPathsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEmailReturnPathsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEmailReturnPathsResponse wrapper for the ListEmailReturnPaths operation
type ListEmailReturnPathsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EmailReturnPathCollection instances
	EmailReturnPathCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListEmailReturnPathsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEmailReturnPathsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEmailReturnPathsSortOrderEnum Enum with underlying type: string
type ListEmailReturnPathsSortOrderEnum string

// Set of constants representing the allowable values for ListEmailReturnPathsSortOrderEnum
const (
	ListEmailReturnPathsSortOrderAsc  ListEmailReturnPathsSortOrderEnum = "ASC"
	ListEmailReturnPathsSortOrderDesc ListEmailReturnPathsSortOrderEnum = "DESC"
)

var mappingListEmailReturnPathsSortOrderEnum = map[string]ListEmailReturnPathsSortOrderEnum{
	"ASC":  ListEmailReturnPathsSortOrderAsc,
	"DESC": ListEmailReturnPathsSortOrderDesc,
}

var mappingListEmailReturnPathsSortOrderEnumLowerCase = map[string]ListEmailReturnPathsSortOrderEnum{
	"asc":  ListEmailReturnPathsSortOrderAsc,
	"desc": ListEmailReturnPathsSortOrderDesc,
}

// GetListEmailReturnPathsSortOrderEnumValues Enumerates the set of values for ListEmailReturnPathsSortOrderEnum
func GetListEmailReturnPathsSortOrderEnumValues() []ListEmailReturnPathsSortOrderEnum {
	values := make([]ListEmailReturnPathsSortOrderEnum, 0)
	for _, v := range mappingListEmailReturnPathsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailReturnPathsSortOrderEnumStringValues Enumerates the set of values in String for ListEmailReturnPathsSortOrderEnum
func GetListEmailReturnPathsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEmailReturnPathsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailReturnPathsSortOrderEnum(val string) (ListEmailReturnPathsSortOrderEnum, bool) {
	enum, ok := mappingListEmailReturnPathsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEmailReturnPathsSortByEnum Enum with underlying type: string
type ListEmailReturnPathsSortByEnum string

// Set of constants representing the allowable values for ListEmailReturnPathsSortByEnum
const (
	ListEmailReturnPathsSortByTimecreated ListEmailReturnPathsSortByEnum = "TIMECREATED"
	ListEmailReturnPathsSortById          ListEmailReturnPathsSortByEnum = "ID"
	ListEmailReturnPathsSortByName        ListEmailReturnPathsSortByEnum = "NAME"
)

var mappingListEmailReturnPathsSortByEnum = map[string]ListEmailReturnPathsSortByEnum{
	"TIMECREATED": ListEmailReturnPathsSortByTimecreated,
	"ID":          ListEmailReturnPathsSortById,
	"NAME":        ListEmailReturnPathsSortByName,
}

var mappingListEmailReturnPathsSortByEnumLowerCase = map[string]ListEmailReturnPathsSortByEnum{
	"timecreated": ListEmailReturnPathsSortByTimecreated,
	"id":          ListEmailReturnPathsSortById,
	"name":        ListEmailReturnPathsSortByName,
}

// GetListEmailReturnPathsSortByEnumValues Enumerates the set of values for ListEmailReturnPathsSortByEnum
func GetListEmailReturnPathsSortByEnumValues() []ListEmailReturnPathsSortByEnum {
	values := make([]ListEmailReturnPathsSortByEnum, 0)
	for _, v := range mappingListEmailReturnPathsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailReturnPathsSortByEnumStringValues Enumerates the set of values in String for ListEmailReturnPathsSortByEnum
func GetListEmailReturnPathsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"ID",
		"NAME",
	}
}

// GetMappingListEmailReturnPathsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailReturnPathsSortByEnum(val string) (ListEmailReturnPathsSortByEnum, bool) {
	enum, ok := mappingListEmailReturnPathsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
