// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedInstancesInDynamicSetRequest wrapper for the ListManagedInstancesInDynamicSet operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstancesInDynamicSet.go.html to see an example of how to use ListManagedInstancesInDynamicSetRequest.
type ListManagedInstancesInDynamicSetRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dynamic set. This filter returns resources associated with this dynamic set.
	DynamicSetId *string `mandatory:"true" contributesTo:"path" name:"dynamicSetId"`

	// The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Indicates whether to include subcompartments in the returned results. Default is false.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return resources that match the given user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListManagedInstancesInDynamicSetSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListManagedInstancesInDynamicSetSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstancesInDynamicSetRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstancesInDynamicSetRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstancesInDynamicSetRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstancesInDynamicSetRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstancesInDynamicSetRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstancesInDynamicSetSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstancesInDynamicSetSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstancesInDynamicSetSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstancesInDynamicSetSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstancesInDynamicSetResponse wrapper for the ListManagedInstancesInDynamicSet operation
type ListManagedInstancesInDynamicSetResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceCollection instances
	ManagedInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of items in the result. Used for pagination of a list of items.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListManagedInstancesInDynamicSetResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstancesInDynamicSetResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstancesInDynamicSetSortOrderEnum Enum with underlying type: string
type ListManagedInstancesInDynamicSetSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstancesInDynamicSetSortOrderEnum
const (
	ListManagedInstancesInDynamicSetSortOrderAsc  ListManagedInstancesInDynamicSetSortOrderEnum = "ASC"
	ListManagedInstancesInDynamicSetSortOrderDesc ListManagedInstancesInDynamicSetSortOrderEnum = "DESC"
)

var mappingListManagedInstancesInDynamicSetSortOrderEnum = map[string]ListManagedInstancesInDynamicSetSortOrderEnum{
	"ASC":  ListManagedInstancesInDynamicSetSortOrderAsc,
	"DESC": ListManagedInstancesInDynamicSetSortOrderDesc,
}

var mappingListManagedInstancesInDynamicSetSortOrderEnumLowerCase = map[string]ListManagedInstancesInDynamicSetSortOrderEnum{
	"asc":  ListManagedInstancesInDynamicSetSortOrderAsc,
	"desc": ListManagedInstancesInDynamicSetSortOrderDesc,
}

// GetListManagedInstancesInDynamicSetSortOrderEnumValues Enumerates the set of values for ListManagedInstancesInDynamicSetSortOrderEnum
func GetListManagedInstancesInDynamicSetSortOrderEnumValues() []ListManagedInstancesInDynamicSetSortOrderEnum {
	values := make([]ListManagedInstancesInDynamicSetSortOrderEnum, 0)
	for _, v := range mappingListManagedInstancesInDynamicSetSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstancesInDynamicSetSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstancesInDynamicSetSortOrderEnum
func GetListManagedInstancesInDynamicSetSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstancesInDynamicSetSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstancesInDynamicSetSortOrderEnum(val string) (ListManagedInstancesInDynamicSetSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstancesInDynamicSetSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstancesInDynamicSetSortByEnum Enum with underlying type: string
type ListManagedInstancesInDynamicSetSortByEnum string

// Set of constants representing the allowable values for ListManagedInstancesInDynamicSetSortByEnum
const (
	ListManagedInstancesInDynamicSetSortByTimecreated ListManagedInstancesInDynamicSetSortByEnum = "timeCreated"
	ListManagedInstancesInDynamicSetSortByDisplayname ListManagedInstancesInDynamicSetSortByEnum = "displayName"
)

var mappingListManagedInstancesInDynamicSetSortByEnum = map[string]ListManagedInstancesInDynamicSetSortByEnum{
	"timeCreated": ListManagedInstancesInDynamicSetSortByTimecreated,
	"displayName": ListManagedInstancesInDynamicSetSortByDisplayname,
}

var mappingListManagedInstancesInDynamicSetSortByEnumLowerCase = map[string]ListManagedInstancesInDynamicSetSortByEnum{
	"timecreated": ListManagedInstancesInDynamicSetSortByTimecreated,
	"displayname": ListManagedInstancesInDynamicSetSortByDisplayname,
}

// GetListManagedInstancesInDynamicSetSortByEnumValues Enumerates the set of values for ListManagedInstancesInDynamicSetSortByEnum
func GetListManagedInstancesInDynamicSetSortByEnumValues() []ListManagedInstancesInDynamicSetSortByEnum {
	values := make([]ListManagedInstancesInDynamicSetSortByEnum, 0)
	for _, v := range mappingListManagedInstancesInDynamicSetSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstancesInDynamicSetSortByEnumStringValues Enumerates the set of values in String for ListManagedInstancesInDynamicSetSortByEnum
func GetListManagedInstancesInDynamicSetSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagedInstancesInDynamicSetSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstancesInDynamicSetSortByEnum(val string) (ListManagedInstancesInDynamicSetSortByEnum, bool) {
	enum, ok := mappingListManagedInstancesInDynamicSetSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
