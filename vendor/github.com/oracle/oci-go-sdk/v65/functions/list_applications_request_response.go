// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListApplicationsRequest wrapper for the ListApplications operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/ListApplications.go.html to see an example of how to use ListApplicationsRequest.
type ListApplicationsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this resource belongs.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return. 1 is the minimum, 50 is the maximum.
	// Default: 10
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token for a list query returned by a previous operation
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only applications that match the lifecycle state in this parameter.
	// Example: `Creating`
	LifecycleState ApplicationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only applications with display names that match the display name string. Matching is exact.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only applications with the specified OCID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Specifies sort order.
	// * **ASC:** Ascending sort order.
	// * **DESC:** Descending sort order.
	SortOrder ListApplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the attribute with which to sort the rules.
	// Default: `displayName`
	// * **timeCreated:** Sorts by timeCreated.
	// * **displayName:** Sorts by displayName.
	// * **id:** Sorts by id.
	SortBy ListApplicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListApplicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApplicationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListApplicationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApplicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListApplicationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApplicationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetApplicationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApplicationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListApplicationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApplicationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListApplicationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListApplicationsResponse wrapper for the ListApplications operation
type ListApplicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ApplicationSummary instances
	Items []ApplicationSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of
	// results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListApplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApplicationsSortOrderEnum Enum with underlying type: string
type ListApplicationsSortOrderEnum string

// Set of constants representing the allowable values for ListApplicationsSortOrderEnum
const (
	ListApplicationsSortOrderAsc  ListApplicationsSortOrderEnum = "ASC"
	ListApplicationsSortOrderDesc ListApplicationsSortOrderEnum = "DESC"
)

var mappingListApplicationsSortOrderEnum = map[string]ListApplicationsSortOrderEnum{
	"ASC":  ListApplicationsSortOrderAsc,
	"DESC": ListApplicationsSortOrderDesc,
}

var mappingListApplicationsSortOrderEnumLowerCase = map[string]ListApplicationsSortOrderEnum{
	"asc":  ListApplicationsSortOrderAsc,
	"desc": ListApplicationsSortOrderDesc,
}

// GetListApplicationsSortOrderEnumValues Enumerates the set of values for ListApplicationsSortOrderEnum
func GetListApplicationsSortOrderEnumValues() []ListApplicationsSortOrderEnum {
	values := make([]ListApplicationsSortOrderEnum, 0)
	for _, v := range mappingListApplicationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplicationsSortOrderEnumStringValues Enumerates the set of values in String for ListApplicationsSortOrderEnum
func GetListApplicationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListApplicationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplicationsSortOrderEnum(val string) (ListApplicationsSortOrderEnum, bool) {
	enum, ok := mappingListApplicationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListApplicationsSortByEnum Enum with underlying type: string
type ListApplicationsSortByEnum string

// Set of constants representing the allowable values for ListApplicationsSortByEnum
const (
	ListApplicationsSortByTimecreated ListApplicationsSortByEnum = "timeCreated"
	ListApplicationsSortById          ListApplicationsSortByEnum = "id"
	ListApplicationsSortByDisplayname ListApplicationsSortByEnum = "displayName"
)

var mappingListApplicationsSortByEnum = map[string]ListApplicationsSortByEnum{
	"timeCreated": ListApplicationsSortByTimecreated,
	"id":          ListApplicationsSortById,
	"displayName": ListApplicationsSortByDisplayname,
}

var mappingListApplicationsSortByEnumLowerCase = map[string]ListApplicationsSortByEnum{
	"timecreated": ListApplicationsSortByTimecreated,
	"id":          ListApplicationsSortById,
	"displayname": ListApplicationsSortByDisplayname,
}

// GetListApplicationsSortByEnumValues Enumerates the set of values for ListApplicationsSortByEnum
func GetListApplicationsSortByEnumValues() []ListApplicationsSortByEnum {
	values := make([]ListApplicationsSortByEnum, 0)
	for _, v := range mappingListApplicationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplicationsSortByEnumStringValues Enumerates the set of values in String for ListApplicationsSortByEnum
func GetListApplicationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"id",
		"displayName",
	}
}

// GetMappingListApplicationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplicationsSortByEnum(val string) (ListApplicationsSortByEnum, bool) {
	enum, ok := mappingListApplicationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
