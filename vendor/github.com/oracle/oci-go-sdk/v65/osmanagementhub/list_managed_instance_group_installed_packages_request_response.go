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

// ListManagedInstanceGroupInstalledPackagesRequest wrapper for the ListManagedInstanceGroupInstalledPackages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceGroupInstalledPackages.go.html to see an example of how to use ListManagedInstanceGroupInstalledPackagesRequest.
type ListManagedInstanceGroupInstalledPackagesRequest struct {

	// The managed instance group OCID.
	ManagedInstanceGroupId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceGroupId"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The install date after which to list all packages, in ISO 8601 format
	// Example: 2017-07-14T02:40:00.000Z
	TimeInstallDateStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeInstallDateStart"`

	// The install date before which to list all packages, in ISO 8601 format.
	// Example: 2017-07-14T02:40:00.000Z
	TimeInstallDateEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeInstallDateEnd"`

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
	SortOrder ListManagedInstanceGroupInstalledPackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeInstalled is descending. Default order for displayName is ascending.
	SortBy ListManagedInstanceGroupInstalledPackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceGroupInstalledPackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceGroupInstalledPackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceGroupInstalledPackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceGroupInstalledPackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceGroupInstalledPackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceGroupInstalledPackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceGroupInstalledPackagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceGroupInstalledPackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceGroupInstalledPackagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceGroupInstalledPackagesResponse wrapper for the ListManagedInstanceGroupInstalledPackages operation
type ListManagedInstanceGroupInstalledPackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceGroupInstalledPackageCollection instances
	ManagedInstanceGroupInstalledPackageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceGroupInstalledPackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceGroupInstalledPackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceGroupInstalledPackagesSortOrderEnum Enum with underlying type: string
type ListManagedInstanceGroupInstalledPackagesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupInstalledPackagesSortOrderEnum
const (
	ListManagedInstanceGroupInstalledPackagesSortOrderAsc  ListManagedInstanceGroupInstalledPackagesSortOrderEnum = "ASC"
	ListManagedInstanceGroupInstalledPackagesSortOrderDesc ListManagedInstanceGroupInstalledPackagesSortOrderEnum = "DESC"
)

var mappingListManagedInstanceGroupInstalledPackagesSortOrderEnum = map[string]ListManagedInstanceGroupInstalledPackagesSortOrderEnum{
	"ASC":  ListManagedInstanceGroupInstalledPackagesSortOrderAsc,
	"DESC": ListManagedInstanceGroupInstalledPackagesSortOrderDesc,
}

var mappingListManagedInstanceGroupInstalledPackagesSortOrderEnumLowerCase = map[string]ListManagedInstanceGroupInstalledPackagesSortOrderEnum{
	"asc":  ListManagedInstanceGroupInstalledPackagesSortOrderAsc,
	"desc": ListManagedInstanceGroupInstalledPackagesSortOrderDesc,
}

// GetListManagedInstanceGroupInstalledPackagesSortOrderEnumValues Enumerates the set of values for ListManagedInstanceGroupInstalledPackagesSortOrderEnum
func GetListManagedInstanceGroupInstalledPackagesSortOrderEnumValues() []ListManagedInstanceGroupInstalledPackagesSortOrderEnum {
	values := make([]ListManagedInstanceGroupInstalledPackagesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceGroupInstalledPackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupInstalledPackagesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupInstalledPackagesSortOrderEnum
func GetListManagedInstanceGroupInstalledPackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceGroupInstalledPackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupInstalledPackagesSortOrderEnum(val string) (ListManagedInstanceGroupInstalledPackagesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupInstalledPackagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceGroupInstalledPackagesSortByEnum Enum with underlying type: string
type ListManagedInstanceGroupInstalledPackagesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupInstalledPackagesSortByEnum
const (
	ListManagedInstanceGroupInstalledPackagesSortByTimeinstalled ListManagedInstanceGroupInstalledPackagesSortByEnum = "timeInstalled"
	ListManagedInstanceGroupInstalledPackagesSortByTimecreated   ListManagedInstanceGroupInstalledPackagesSortByEnum = "timeCreated"
	ListManagedInstanceGroupInstalledPackagesSortByDisplayname   ListManagedInstanceGroupInstalledPackagesSortByEnum = "displayName"
)

var mappingListManagedInstanceGroupInstalledPackagesSortByEnum = map[string]ListManagedInstanceGroupInstalledPackagesSortByEnum{
	"timeInstalled": ListManagedInstanceGroupInstalledPackagesSortByTimeinstalled,
	"timeCreated":   ListManagedInstanceGroupInstalledPackagesSortByTimecreated,
	"displayName":   ListManagedInstanceGroupInstalledPackagesSortByDisplayname,
}

var mappingListManagedInstanceGroupInstalledPackagesSortByEnumLowerCase = map[string]ListManagedInstanceGroupInstalledPackagesSortByEnum{
	"timeinstalled": ListManagedInstanceGroupInstalledPackagesSortByTimeinstalled,
	"timecreated":   ListManagedInstanceGroupInstalledPackagesSortByTimecreated,
	"displayname":   ListManagedInstanceGroupInstalledPackagesSortByDisplayname,
}

// GetListManagedInstanceGroupInstalledPackagesSortByEnumValues Enumerates the set of values for ListManagedInstanceGroupInstalledPackagesSortByEnum
func GetListManagedInstanceGroupInstalledPackagesSortByEnumValues() []ListManagedInstanceGroupInstalledPackagesSortByEnum {
	values := make([]ListManagedInstanceGroupInstalledPackagesSortByEnum, 0)
	for _, v := range mappingListManagedInstanceGroupInstalledPackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupInstalledPackagesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupInstalledPackagesSortByEnum
func GetListManagedInstanceGroupInstalledPackagesSortByEnumStringValues() []string {
	return []string{
		"timeInstalled",
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagedInstanceGroupInstalledPackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupInstalledPackagesSortByEnum(val string) (ListManagedInstanceGroupInstalledPackagesSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupInstalledPackagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
