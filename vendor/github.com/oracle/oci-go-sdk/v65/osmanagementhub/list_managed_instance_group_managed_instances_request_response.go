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

// ListManagedInstanceGroupManagedInstancesRequest wrapper for the ListManagedInstanceGroupManagedInstances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceGroupManagedInstances.go.html to see an example of how to use ListManagedInstanceGroupManagedInstancesRequest.
type ListManagedInstanceGroupManagedInstancesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.
	ManagedInstanceGroupId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceGroupId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance. This filter returns resources associated with this managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// Indicates whether to include subcompartments in the returned results. Default is false.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListManagedInstanceGroupManagedInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListManagedInstanceGroupManagedInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceGroupManagedInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceGroupManagedInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceGroupManagedInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceGroupManagedInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceGroupManagedInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceGroupManagedInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceGroupManagedInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceGroupManagedInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceGroupManagedInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceGroupManagedInstancesResponse wrapper for the ListManagedInstanceGroupManagedInstances operation
type ListManagedInstanceGroupManagedInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceCollection instances
	ManagedInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceGroupManagedInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceGroupManagedInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceGroupManagedInstancesSortOrderEnum Enum with underlying type: string
type ListManagedInstanceGroupManagedInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupManagedInstancesSortOrderEnum
const (
	ListManagedInstanceGroupManagedInstancesSortOrderAsc  ListManagedInstanceGroupManagedInstancesSortOrderEnum = "ASC"
	ListManagedInstanceGroupManagedInstancesSortOrderDesc ListManagedInstanceGroupManagedInstancesSortOrderEnum = "DESC"
)

var mappingListManagedInstanceGroupManagedInstancesSortOrderEnum = map[string]ListManagedInstanceGroupManagedInstancesSortOrderEnum{
	"ASC":  ListManagedInstanceGroupManagedInstancesSortOrderAsc,
	"DESC": ListManagedInstanceGroupManagedInstancesSortOrderDesc,
}

var mappingListManagedInstanceGroupManagedInstancesSortOrderEnumLowerCase = map[string]ListManagedInstanceGroupManagedInstancesSortOrderEnum{
	"asc":  ListManagedInstanceGroupManagedInstancesSortOrderAsc,
	"desc": ListManagedInstanceGroupManagedInstancesSortOrderDesc,
}

// GetListManagedInstanceGroupManagedInstancesSortOrderEnumValues Enumerates the set of values for ListManagedInstanceGroupManagedInstancesSortOrderEnum
func GetListManagedInstanceGroupManagedInstancesSortOrderEnumValues() []ListManagedInstanceGroupManagedInstancesSortOrderEnum {
	values := make([]ListManagedInstanceGroupManagedInstancesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceGroupManagedInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupManagedInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupManagedInstancesSortOrderEnum
func GetListManagedInstanceGroupManagedInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceGroupManagedInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupManagedInstancesSortOrderEnum(val string) (ListManagedInstanceGroupManagedInstancesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupManagedInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceGroupManagedInstancesSortByEnum Enum with underlying type: string
type ListManagedInstanceGroupManagedInstancesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupManagedInstancesSortByEnum
const (
	ListManagedInstanceGroupManagedInstancesSortByTimecreated ListManagedInstanceGroupManagedInstancesSortByEnum = "timeCreated"
	ListManagedInstanceGroupManagedInstancesSortByDisplayname ListManagedInstanceGroupManagedInstancesSortByEnum = "displayName"
)

var mappingListManagedInstanceGroupManagedInstancesSortByEnum = map[string]ListManagedInstanceGroupManagedInstancesSortByEnum{
	"timeCreated": ListManagedInstanceGroupManagedInstancesSortByTimecreated,
	"displayName": ListManagedInstanceGroupManagedInstancesSortByDisplayname,
}

var mappingListManagedInstanceGroupManagedInstancesSortByEnumLowerCase = map[string]ListManagedInstanceGroupManagedInstancesSortByEnum{
	"timecreated": ListManagedInstanceGroupManagedInstancesSortByTimecreated,
	"displayname": ListManagedInstanceGroupManagedInstancesSortByDisplayname,
}

// GetListManagedInstanceGroupManagedInstancesSortByEnumValues Enumerates the set of values for ListManagedInstanceGroupManagedInstancesSortByEnum
func GetListManagedInstanceGroupManagedInstancesSortByEnumValues() []ListManagedInstanceGroupManagedInstancesSortByEnum {
	values := make([]ListManagedInstanceGroupManagedInstancesSortByEnum, 0)
	for _, v := range mappingListManagedInstanceGroupManagedInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupManagedInstancesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupManagedInstancesSortByEnum
func GetListManagedInstanceGroupManagedInstancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagedInstanceGroupManagedInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupManagedInstancesSortByEnum(val string) (ListManagedInstanceGroupManagedInstancesSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupManagedInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
