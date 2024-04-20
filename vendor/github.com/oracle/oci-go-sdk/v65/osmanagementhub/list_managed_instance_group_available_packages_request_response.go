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

// ListManagedInstanceGroupAvailablePackagesRequest wrapper for the ListManagedInstanceGroupAvailablePackages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceGroupAvailablePackages.go.html to see an example of how to use ListManagedInstanceGroupAvailablePackagesRequest.
type ListManagedInstanceGroupAvailablePackagesRequest struct {

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
	SortOrder ListManagedInstanceGroupAvailablePackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListManagedInstanceGroupAvailablePackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Indicates whether to list only the latest versions of packages, module streams, and stream profiles.
	IsLatest *bool `mandatory:"false" contributesTo:"query" name:"isLatest"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceGroupAvailablePackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceGroupAvailablePackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceGroupAvailablePackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceGroupAvailablePackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceGroupAvailablePackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceGroupAvailablePackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceGroupAvailablePackagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceGroupAvailablePackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceGroupAvailablePackagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceGroupAvailablePackagesResponse wrapper for the ListManagedInstanceGroupAvailablePackages operation
type ListManagedInstanceGroupAvailablePackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceGroupAvailablePackageCollection instances
	ManagedInstanceGroupAvailablePackageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceGroupAvailablePackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceGroupAvailablePackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceGroupAvailablePackagesSortOrderEnum Enum with underlying type: string
type ListManagedInstanceGroupAvailablePackagesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupAvailablePackagesSortOrderEnum
const (
	ListManagedInstanceGroupAvailablePackagesSortOrderAsc  ListManagedInstanceGroupAvailablePackagesSortOrderEnum = "ASC"
	ListManagedInstanceGroupAvailablePackagesSortOrderDesc ListManagedInstanceGroupAvailablePackagesSortOrderEnum = "DESC"
)

var mappingListManagedInstanceGroupAvailablePackagesSortOrderEnum = map[string]ListManagedInstanceGroupAvailablePackagesSortOrderEnum{
	"ASC":  ListManagedInstanceGroupAvailablePackagesSortOrderAsc,
	"DESC": ListManagedInstanceGroupAvailablePackagesSortOrderDesc,
}

var mappingListManagedInstanceGroupAvailablePackagesSortOrderEnumLowerCase = map[string]ListManagedInstanceGroupAvailablePackagesSortOrderEnum{
	"asc":  ListManagedInstanceGroupAvailablePackagesSortOrderAsc,
	"desc": ListManagedInstanceGroupAvailablePackagesSortOrderDesc,
}

// GetListManagedInstanceGroupAvailablePackagesSortOrderEnumValues Enumerates the set of values for ListManagedInstanceGroupAvailablePackagesSortOrderEnum
func GetListManagedInstanceGroupAvailablePackagesSortOrderEnumValues() []ListManagedInstanceGroupAvailablePackagesSortOrderEnum {
	values := make([]ListManagedInstanceGroupAvailablePackagesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceGroupAvailablePackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupAvailablePackagesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupAvailablePackagesSortOrderEnum
func GetListManagedInstanceGroupAvailablePackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceGroupAvailablePackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupAvailablePackagesSortOrderEnum(val string) (ListManagedInstanceGroupAvailablePackagesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupAvailablePackagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceGroupAvailablePackagesSortByEnum Enum with underlying type: string
type ListManagedInstanceGroupAvailablePackagesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupAvailablePackagesSortByEnum
const (
	ListManagedInstanceGroupAvailablePackagesSortByTimecreated ListManagedInstanceGroupAvailablePackagesSortByEnum = "timeCreated"
	ListManagedInstanceGroupAvailablePackagesSortByDisplayname ListManagedInstanceGroupAvailablePackagesSortByEnum = "displayName"
)

var mappingListManagedInstanceGroupAvailablePackagesSortByEnum = map[string]ListManagedInstanceGroupAvailablePackagesSortByEnum{
	"timeCreated": ListManagedInstanceGroupAvailablePackagesSortByTimecreated,
	"displayName": ListManagedInstanceGroupAvailablePackagesSortByDisplayname,
}

var mappingListManagedInstanceGroupAvailablePackagesSortByEnumLowerCase = map[string]ListManagedInstanceGroupAvailablePackagesSortByEnum{
	"timecreated": ListManagedInstanceGroupAvailablePackagesSortByTimecreated,
	"displayname": ListManagedInstanceGroupAvailablePackagesSortByDisplayname,
}

// GetListManagedInstanceGroupAvailablePackagesSortByEnumValues Enumerates the set of values for ListManagedInstanceGroupAvailablePackagesSortByEnum
func GetListManagedInstanceGroupAvailablePackagesSortByEnumValues() []ListManagedInstanceGroupAvailablePackagesSortByEnum {
	values := make([]ListManagedInstanceGroupAvailablePackagesSortByEnum, 0)
	for _, v := range mappingListManagedInstanceGroupAvailablePackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupAvailablePackagesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupAvailablePackagesSortByEnum
func GetListManagedInstanceGroupAvailablePackagesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagedInstanceGroupAvailablePackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupAvailablePackagesSortByEnum(val string) (ListManagedInstanceGroupAvailablePackagesSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupAvailablePackagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
