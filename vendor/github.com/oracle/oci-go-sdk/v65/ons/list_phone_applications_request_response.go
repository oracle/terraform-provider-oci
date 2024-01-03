// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ons

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPhoneApplicationsRequest wrapper for the ListPhoneApplications operation
type ListPhoneApplicationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState PhoneApplicationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique PhoneApplication identifier
	PhoneApplicationId *string `mandatory:"false" contributesTo:"query" name:"phoneApplicationId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use (ascending or descending).
	SortOrder ListPhoneApplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListPhoneApplicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPhoneApplicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPhoneApplicationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPhoneApplicationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPhoneApplicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPhoneApplicationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPhoneApplicationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPhoneApplicationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPhoneApplicationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPhoneApplicationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPhoneApplicationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPhoneApplicationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPhoneApplicationsResponse wrapper for the ListPhoneApplications operation
type ListPhoneApplicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PhoneApplicationCollection instances
	PhoneApplicationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPhoneApplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPhoneApplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPhoneApplicationsSortOrderEnum Enum with underlying type: string
type ListPhoneApplicationsSortOrderEnum string

// Set of constants representing the allowable values for ListPhoneApplicationsSortOrderEnum
const (
	ListPhoneApplicationsSortOrderAsc  ListPhoneApplicationsSortOrderEnum = "ASC"
	ListPhoneApplicationsSortOrderDesc ListPhoneApplicationsSortOrderEnum = "DESC"
)

var mappingListPhoneApplicationsSortOrderEnum = map[string]ListPhoneApplicationsSortOrderEnum{
	"ASC":  ListPhoneApplicationsSortOrderAsc,
	"DESC": ListPhoneApplicationsSortOrderDesc,
}

var mappingListPhoneApplicationsSortOrderEnumLowerCase = map[string]ListPhoneApplicationsSortOrderEnum{
	"asc":  ListPhoneApplicationsSortOrderAsc,
	"desc": ListPhoneApplicationsSortOrderDesc,
}

// GetListPhoneApplicationsSortOrderEnumValues Enumerates the set of values for ListPhoneApplicationsSortOrderEnum
func GetListPhoneApplicationsSortOrderEnumValues() []ListPhoneApplicationsSortOrderEnum {
	values := make([]ListPhoneApplicationsSortOrderEnum, 0)
	for _, v := range mappingListPhoneApplicationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPhoneApplicationsSortOrderEnumStringValues Enumerates the set of values in String for ListPhoneApplicationsSortOrderEnum
func GetListPhoneApplicationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPhoneApplicationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPhoneApplicationsSortOrderEnum(val string) (ListPhoneApplicationsSortOrderEnum, bool) {
	enum, ok := mappingListPhoneApplicationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPhoneApplicationsSortByEnum Enum with underlying type: string
type ListPhoneApplicationsSortByEnum string

// Set of constants representing the allowable values for ListPhoneApplicationsSortByEnum
const (
	ListPhoneApplicationsSortByTimecreated ListPhoneApplicationsSortByEnum = "timeCreated"
	ListPhoneApplicationsSortByDisplayname ListPhoneApplicationsSortByEnum = "displayName"
)

var mappingListPhoneApplicationsSortByEnum = map[string]ListPhoneApplicationsSortByEnum{
	"timeCreated": ListPhoneApplicationsSortByTimecreated,
	"displayName": ListPhoneApplicationsSortByDisplayname,
}

var mappingListPhoneApplicationsSortByEnumLowerCase = map[string]ListPhoneApplicationsSortByEnum{
	"timecreated": ListPhoneApplicationsSortByTimecreated,
	"displayname": ListPhoneApplicationsSortByDisplayname,
}

// GetListPhoneApplicationsSortByEnumValues Enumerates the set of values for ListPhoneApplicationsSortByEnum
func GetListPhoneApplicationsSortByEnumValues() []ListPhoneApplicationsSortByEnum {
	values := make([]ListPhoneApplicationsSortByEnum, 0)
	for _, v := range mappingListPhoneApplicationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPhoneApplicationsSortByEnumStringValues Enumerates the set of values in String for ListPhoneApplicationsSortByEnum
func GetListPhoneApplicationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPhoneApplicationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPhoneApplicationsSortByEnum(val string) (ListPhoneApplicationsSortByEnum, bool) {
	enum, ok := mappingListPhoneApplicationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
