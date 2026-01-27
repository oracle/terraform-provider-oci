// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeaiagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListProvisionedCapacitiesRequest wrapper for the ListProvisionedCapacities operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiagent/ListProvisionedCapacities.go.html to see an example of how to use ListProvisionedCapacitiesRequest.
type ListProvisionedCapacitiesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the provisioned capacity.
	ProvisionedCapacityId *string `mandatory:"false" contributesTo:"query" name:"provisionedCapacityId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState ProvisionedCapacityLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListProvisionedCapacitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListProvisionedCapacitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProvisionedCapacitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProvisionedCapacitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProvisionedCapacitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProvisionedCapacitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProvisionedCapacitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProvisionedCapacityLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetProvisionedCapacityLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProvisionedCapacitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProvisionedCapacitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProvisionedCapacitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProvisionedCapacitiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProvisionedCapacitiesResponse wrapper for the ListProvisionedCapacities operation
type ListProvisionedCapacitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProvisionedCapacityCollection instances
	ProvisionedCapacityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProvisionedCapacitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProvisionedCapacitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProvisionedCapacitiesSortOrderEnum Enum with underlying type: string
type ListProvisionedCapacitiesSortOrderEnum string

// Set of constants representing the allowable values for ListProvisionedCapacitiesSortOrderEnum
const (
	ListProvisionedCapacitiesSortOrderAsc  ListProvisionedCapacitiesSortOrderEnum = "ASC"
	ListProvisionedCapacitiesSortOrderDesc ListProvisionedCapacitiesSortOrderEnum = "DESC"
)

var mappingListProvisionedCapacitiesSortOrderEnum = map[string]ListProvisionedCapacitiesSortOrderEnum{
	"ASC":  ListProvisionedCapacitiesSortOrderAsc,
	"DESC": ListProvisionedCapacitiesSortOrderDesc,
}

var mappingListProvisionedCapacitiesSortOrderEnumLowerCase = map[string]ListProvisionedCapacitiesSortOrderEnum{
	"asc":  ListProvisionedCapacitiesSortOrderAsc,
	"desc": ListProvisionedCapacitiesSortOrderDesc,
}

// GetListProvisionedCapacitiesSortOrderEnumValues Enumerates the set of values for ListProvisionedCapacitiesSortOrderEnum
func GetListProvisionedCapacitiesSortOrderEnumValues() []ListProvisionedCapacitiesSortOrderEnum {
	values := make([]ListProvisionedCapacitiesSortOrderEnum, 0)
	for _, v := range mappingListProvisionedCapacitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProvisionedCapacitiesSortOrderEnumStringValues Enumerates the set of values in String for ListProvisionedCapacitiesSortOrderEnum
func GetListProvisionedCapacitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProvisionedCapacitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProvisionedCapacitiesSortOrderEnum(val string) (ListProvisionedCapacitiesSortOrderEnum, bool) {
	enum, ok := mappingListProvisionedCapacitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProvisionedCapacitiesSortByEnum Enum with underlying type: string
type ListProvisionedCapacitiesSortByEnum string

// Set of constants representing the allowable values for ListProvisionedCapacitiesSortByEnum
const (
	ListProvisionedCapacitiesSortByTimecreated ListProvisionedCapacitiesSortByEnum = "timeCreated"
	ListProvisionedCapacitiesSortByDisplayname ListProvisionedCapacitiesSortByEnum = "displayName"
)

var mappingListProvisionedCapacitiesSortByEnum = map[string]ListProvisionedCapacitiesSortByEnum{
	"timeCreated": ListProvisionedCapacitiesSortByTimecreated,
	"displayName": ListProvisionedCapacitiesSortByDisplayname,
}

var mappingListProvisionedCapacitiesSortByEnumLowerCase = map[string]ListProvisionedCapacitiesSortByEnum{
	"timecreated": ListProvisionedCapacitiesSortByTimecreated,
	"displayname": ListProvisionedCapacitiesSortByDisplayname,
}

// GetListProvisionedCapacitiesSortByEnumValues Enumerates the set of values for ListProvisionedCapacitiesSortByEnum
func GetListProvisionedCapacitiesSortByEnumValues() []ListProvisionedCapacitiesSortByEnum {
	values := make([]ListProvisionedCapacitiesSortByEnum, 0)
	for _, v := range mappingListProvisionedCapacitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProvisionedCapacitiesSortByEnumStringValues Enumerates the set of values in String for ListProvisionedCapacitiesSortByEnum
func GetListProvisionedCapacitiesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListProvisionedCapacitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProvisionedCapacitiesSortByEnum(val string) (ListProvisionedCapacitiesSortByEnum, bool) {
	enum, ok := mappingListProvisionedCapacitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
