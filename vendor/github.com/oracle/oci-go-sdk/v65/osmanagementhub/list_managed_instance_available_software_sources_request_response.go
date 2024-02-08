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

// ListManagedInstanceAvailableSoftwareSourcesRequest wrapper for the ListManagedInstanceAvailableSoftwareSources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceAvailableSoftwareSources.go.html to see an example of how to use ListManagedInstanceAvailableSoftwareSourcesRequest.
type ListManagedInstanceAvailableSoftwareSourcesRequest struct {

	// The OCID of the managed instance.
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The OCID of the compartment that contains the resources to list.
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
	SortOrder ListManagedInstanceAvailableSoftwareSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListManagedInstanceAvailableSoftwareSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceAvailableSoftwareSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceAvailableSoftwareSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceAvailableSoftwareSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceAvailableSoftwareSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceAvailableSoftwareSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceAvailableSoftwareSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceAvailableSoftwareSourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceAvailableSoftwareSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceAvailableSoftwareSourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceAvailableSoftwareSourcesResponse wrapper for the ListManagedInstanceAvailableSoftwareSources operation
type ListManagedInstanceAvailableSoftwareSourcesResponse struct {

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

func (response ListManagedInstanceAvailableSoftwareSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceAvailableSoftwareSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceAvailableSoftwareSourcesSortOrderEnum Enum with underlying type: string
type ListManagedInstanceAvailableSoftwareSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceAvailableSoftwareSourcesSortOrderEnum
const (
	ListManagedInstanceAvailableSoftwareSourcesSortOrderAsc  ListManagedInstanceAvailableSoftwareSourcesSortOrderEnum = "ASC"
	ListManagedInstanceAvailableSoftwareSourcesSortOrderDesc ListManagedInstanceAvailableSoftwareSourcesSortOrderEnum = "DESC"
)

var mappingListManagedInstanceAvailableSoftwareSourcesSortOrderEnum = map[string]ListManagedInstanceAvailableSoftwareSourcesSortOrderEnum{
	"ASC":  ListManagedInstanceAvailableSoftwareSourcesSortOrderAsc,
	"DESC": ListManagedInstanceAvailableSoftwareSourcesSortOrderDesc,
}

var mappingListManagedInstanceAvailableSoftwareSourcesSortOrderEnumLowerCase = map[string]ListManagedInstanceAvailableSoftwareSourcesSortOrderEnum{
	"asc":  ListManagedInstanceAvailableSoftwareSourcesSortOrderAsc,
	"desc": ListManagedInstanceAvailableSoftwareSourcesSortOrderDesc,
}

// GetListManagedInstanceAvailableSoftwareSourcesSortOrderEnumValues Enumerates the set of values for ListManagedInstanceAvailableSoftwareSourcesSortOrderEnum
func GetListManagedInstanceAvailableSoftwareSourcesSortOrderEnumValues() []ListManagedInstanceAvailableSoftwareSourcesSortOrderEnum {
	values := make([]ListManagedInstanceAvailableSoftwareSourcesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceAvailableSoftwareSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceAvailableSoftwareSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceAvailableSoftwareSourcesSortOrderEnum
func GetListManagedInstanceAvailableSoftwareSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceAvailableSoftwareSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceAvailableSoftwareSourcesSortOrderEnum(val string) (ListManagedInstanceAvailableSoftwareSourcesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceAvailableSoftwareSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceAvailableSoftwareSourcesSortByEnum Enum with underlying type: string
type ListManagedInstanceAvailableSoftwareSourcesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceAvailableSoftwareSourcesSortByEnum
const (
	ListManagedInstanceAvailableSoftwareSourcesSortByTimecreated ListManagedInstanceAvailableSoftwareSourcesSortByEnum = "timeCreated"
	ListManagedInstanceAvailableSoftwareSourcesSortByDisplayname ListManagedInstanceAvailableSoftwareSourcesSortByEnum = "displayName"
)

var mappingListManagedInstanceAvailableSoftwareSourcesSortByEnum = map[string]ListManagedInstanceAvailableSoftwareSourcesSortByEnum{
	"timeCreated": ListManagedInstanceAvailableSoftwareSourcesSortByTimecreated,
	"displayName": ListManagedInstanceAvailableSoftwareSourcesSortByDisplayname,
}

var mappingListManagedInstanceAvailableSoftwareSourcesSortByEnumLowerCase = map[string]ListManagedInstanceAvailableSoftwareSourcesSortByEnum{
	"timecreated": ListManagedInstanceAvailableSoftwareSourcesSortByTimecreated,
	"displayname": ListManagedInstanceAvailableSoftwareSourcesSortByDisplayname,
}

// GetListManagedInstanceAvailableSoftwareSourcesSortByEnumValues Enumerates the set of values for ListManagedInstanceAvailableSoftwareSourcesSortByEnum
func GetListManagedInstanceAvailableSoftwareSourcesSortByEnumValues() []ListManagedInstanceAvailableSoftwareSourcesSortByEnum {
	values := make([]ListManagedInstanceAvailableSoftwareSourcesSortByEnum, 0)
	for _, v := range mappingListManagedInstanceAvailableSoftwareSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceAvailableSoftwareSourcesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceAvailableSoftwareSourcesSortByEnum
func GetListManagedInstanceAvailableSoftwareSourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagedInstanceAvailableSoftwareSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceAvailableSoftwareSourcesSortByEnum(val string) (ListManagedInstanceAvailableSoftwareSourcesSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceAvailableSoftwareSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
