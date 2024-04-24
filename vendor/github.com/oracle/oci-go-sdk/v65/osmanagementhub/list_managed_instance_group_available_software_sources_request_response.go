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

// ListManagedInstanceGroupAvailableSoftwareSourcesRequest wrapper for the ListManagedInstanceGroupAvailableSoftwareSources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceGroupAvailableSoftwareSources.go.html to see an example of how to use ListManagedInstanceGroupAvailableSoftwareSourcesRequest.
type ListManagedInstanceGroupAvailableSoftwareSourcesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.
	ManagedInstanceGroupId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceGroupId"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceGroupAvailableSoftwareSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceGroupAvailableSoftwareSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceGroupAvailableSoftwareSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceGroupAvailableSoftwareSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceGroupAvailableSoftwareSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceGroupAvailableSoftwareSourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceGroupAvailableSoftwareSourcesResponse wrapper for the ListManagedInstanceGroupAvailableSoftwareSources operation
type ListManagedInstanceGroupAvailableSoftwareSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AvailableSoftwareSourceCollection instances
	AvailableSoftwareSourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceGroupAvailableSoftwareSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceGroupAvailableSoftwareSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum Enum with underlying type: string
type ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum
const (
	ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderAsc  ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum = "ASC"
	ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderDesc ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum = "DESC"
)

var mappingListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum = map[string]ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum{
	"ASC":  ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderAsc,
	"DESC": ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderDesc,
}

var mappingListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnumLowerCase = map[string]ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum{
	"asc":  ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderAsc,
	"desc": ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderDesc,
}

// GetListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnumValues Enumerates the set of values for ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum
func GetListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnumValues() []ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum {
	values := make([]ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum
func GetListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum(val string) (ListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupAvailableSoftwareSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum Enum with underlying type: string
type ListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum
const (
	ListManagedInstanceGroupAvailableSoftwareSourcesSortByTimecreated ListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum = "timeCreated"
	ListManagedInstanceGroupAvailableSoftwareSourcesSortByDisplayname ListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum = "displayName"
)

var mappingListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum = map[string]ListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum{
	"timeCreated": ListManagedInstanceGroupAvailableSoftwareSourcesSortByTimecreated,
	"displayName": ListManagedInstanceGroupAvailableSoftwareSourcesSortByDisplayname,
}

var mappingListManagedInstanceGroupAvailableSoftwareSourcesSortByEnumLowerCase = map[string]ListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum{
	"timecreated": ListManagedInstanceGroupAvailableSoftwareSourcesSortByTimecreated,
	"displayname": ListManagedInstanceGroupAvailableSoftwareSourcesSortByDisplayname,
}

// GetListManagedInstanceGroupAvailableSoftwareSourcesSortByEnumValues Enumerates the set of values for ListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum
func GetListManagedInstanceGroupAvailableSoftwareSourcesSortByEnumValues() []ListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum {
	values := make([]ListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum, 0)
	for _, v := range mappingListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupAvailableSoftwareSourcesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum
func GetListManagedInstanceGroupAvailableSoftwareSourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum(val string) (ListManagedInstanceGroupAvailableSoftwareSourcesSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupAvailableSoftwareSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
