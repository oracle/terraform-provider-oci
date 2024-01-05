// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagementStationsRequest wrapper for the ListManagementStations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagementStations.go.html to see an example of how to use ListManagementStationsRequest.
type ListManagementStationsRequest struct {

	// The OCID of the compartment that contains the resources to list.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The current lifecycle state for the object.
	LifecycleState ManagementStationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The OCID of the managed instance for which to list resources.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListManagementStationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListManagementStationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID of the management station.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementStationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementStationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagementStationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagementStationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagementStationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingManagementStationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetManagementStationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementStationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagementStationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagementStationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagementStationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagementStationsResponse wrapper for the ListManagementStations operation
type ListManagementStationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagementStationCollection instances
	ManagementStationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagementStationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagementStationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagementStationsSortOrderEnum Enum with underlying type: string
type ListManagementStationsSortOrderEnum string

// Set of constants representing the allowable values for ListManagementStationsSortOrderEnum
const (
	ListManagementStationsSortOrderAsc  ListManagementStationsSortOrderEnum = "ASC"
	ListManagementStationsSortOrderDesc ListManagementStationsSortOrderEnum = "DESC"
)

var mappingListManagementStationsSortOrderEnum = map[string]ListManagementStationsSortOrderEnum{
	"ASC":  ListManagementStationsSortOrderAsc,
	"DESC": ListManagementStationsSortOrderDesc,
}

var mappingListManagementStationsSortOrderEnumLowerCase = map[string]ListManagementStationsSortOrderEnum{
	"asc":  ListManagementStationsSortOrderAsc,
	"desc": ListManagementStationsSortOrderDesc,
}

// GetListManagementStationsSortOrderEnumValues Enumerates the set of values for ListManagementStationsSortOrderEnum
func GetListManagementStationsSortOrderEnumValues() []ListManagementStationsSortOrderEnum {
	values := make([]ListManagementStationsSortOrderEnum, 0)
	for _, v := range mappingListManagementStationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementStationsSortOrderEnumStringValues Enumerates the set of values in String for ListManagementStationsSortOrderEnum
func GetListManagementStationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagementStationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementStationsSortOrderEnum(val string) (ListManagementStationsSortOrderEnum, bool) {
	enum, ok := mappingListManagementStationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagementStationsSortByEnum Enum with underlying type: string
type ListManagementStationsSortByEnum string

// Set of constants representing the allowable values for ListManagementStationsSortByEnum
const (
	ListManagementStationsSortByTimecreated ListManagementStationsSortByEnum = "timeCreated"
	ListManagementStationsSortByDisplayname ListManagementStationsSortByEnum = "displayName"
)

var mappingListManagementStationsSortByEnum = map[string]ListManagementStationsSortByEnum{
	"timeCreated": ListManagementStationsSortByTimecreated,
	"displayName": ListManagementStationsSortByDisplayname,
}

var mappingListManagementStationsSortByEnumLowerCase = map[string]ListManagementStationsSortByEnum{
	"timecreated": ListManagementStationsSortByTimecreated,
	"displayname": ListManagementStationsSortByDisplayname,
}

// GetListManagementStationsSortByEnumValues Enumerates the set of values for ListManagementStationsSortByEnum
func GetListManagementStationsSortByEnumValues() []ListManagementStationsSortByEnum {
	values := make([]ListManagementStationsSortByEnum, 0)
	for _, v := range mappingListManagementStationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementStationsSortByEnumStringValues Enumerates the set of values in String for ListManagementStationsSortByEnum
func GetListManagementStationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagementStationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementStationsSortByEnum(val string) (ListManagementStationsSortByEnum, bool) {
	enum, ok := mappingListManagementStationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
