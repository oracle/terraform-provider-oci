// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedInstanceInstalledPackagesRequest wrapper for the ListManagedInstanceInstalledPackages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceInstalledPackages.go.html to see an example of how to use ListManagedInstanceInstalledPackagesRequest.
type ListManagedInstanceInstalledPackagesRequest struct {

	// The OCID of the managed instance.
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

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
	SortOrder ListManagedInstanceInstalledPackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeInstalled is descending. Default order for displayName is ascending.
	SortBy ListManagedInstanceInstalledPackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceInstalledPackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceInstalledPackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceInstalledPackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceInstalledPackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceInstalledPackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceInstalledPackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceInstalledPackagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceInstalledPackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceInstalledPackagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceInstalledPackagesResponse wrapper for the ListManagedInstanceInstalledPackages operation
type ListManagedInstanceInstalledPackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InstalledPackageCollection instances
	InstalledPackageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceInstalledPackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceInstalledPackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceInstalledPackagesSortOrderEnum Enum with underlying type: string
type ListManagedInstanceInstalledPackagesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceInstalledPackagesSortOrderEnum
const (
	ListManagedInstanceInstalledPackagesSortOrderAsc  ListManagedInstanceInstalledPackagesSortOrderEnum = "ASC"
	ListManagedInstanceInstalledPackagesSortOrderDesc ListManagedInstanceInstalledPackagesSortOrderEnum = "DESC"
)

var mappingListManagedInstanceInstalledPackagesSortOrderEnum = map[string]ListManagedInstanceInstalledPackagesSortOrderEnum{
	"ASC":  ListManagedInstanceInstalledPackagesSortOrderAsc,
	"DESC": ListManagedInstanceInstalledPackagesSortOrderDesc,
}

var mappingListManagedInstanceInstalledPackagesSortOrderEnumLowerCase = map[string]ListManagedInstanceInstalledPackagesSortOrderEnum{
	"asc":  ListManagedInstanceInstalledPackagesSortOrderAsc,
	"desc": ListManagedInstanceInstalledPackagesSortOrderDesc,
}

// GetListManagedInstanceInstalledPackagesSortOrderEnumValues Enumerates the set of values for ListManagedInstanceInstalledPackagesSortOrderEnum
func GetListManagedInstanceInstalledPackagesSortOrderEnumValues() []ListManagedInstanceInstalledPackagesSortOrderEnum {
	values := make([]ListManagedInstanceInstalledPackagesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceInstalledPackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceInstalledPackagesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceInstalledPackagesSortOrderEnum
func GetListManagedInstanceInstalledPackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceInstalledPackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceInstalledPackagesSortOrderEnum(val string) (ListManagedInstanceInstalledPackagesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceInstalledPackagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceInstalledPackagesSortByEnum Enum with underlying type: string
type ListManagedInstanceInstalledPackagesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceInstalledPackagesSortByEnum
const (
	ListManagedInstanceInstalledPackagesSortByTimeinstalled ListManagedInstanceInstalledPackagesSortByEnum = "timeInstalled"
	ListManagedInstanceInstalledPackagesSortByTimecreated   ListManagedInstanceInstalledPackagesSortByEnum = "timeCreated"
	ListManagedInstanceInstalledPackagesSortByDisplayname   ListManagedInstanceInstalledPackagesSortByEnum = "displayName"
)

var mappingListManagedInstanceInstalledPackagesSortByEnum = map[string]ListManagedInstanceInstalledPackagesSortByEnum{
	"timeInstalled": ListManagedInstanceInstalledPackagesSortByTimeinstalled,
	"timeCreated":   ListManagedInstanceInstalledPackagesSortByTimecreated,
	"displayName":   ListManagedInstanceInstalledPackagesSortByDisplayname,
}

var mappingListManagedInstanceInstalledPackagesSortByEnumLowerCase = map[string]ListManagedInstanceInstalledPackagesSortByEnum{
	"timeinstalled": ListManagedInstanceInstalledPackagesSortByTimeinstalled,
	"timecreated":   ListManagedInstanceInstalledPackagesSortByTimecreated,
	"displayname":   ListManagedInstanceInstalledPackagesSortByDisplayname,
}

// GetListManagedInstanceInstalledPackagesSortByEnumValues Enumerates the set of values for ListManagedInstanceInstalledPackagesSortByEnum
func GetListManagedInstanceInstalledPackagesSortByEnumValues() []ListManagedInstanceInstalledPackagesSortByEnum {
	values := make([]ListManagedInstanceInstalledPackagesSortByEnum, 0)
	for _, v := range mappingListManagedInstanceInstalledPackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceInstalledPackagesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceInstalledPackagesSortByEnum
func GetListManagedInstanceInstalledPackagesSortByEnumStringValues() []string {
	return []string{
		"timeInstalled",
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagedInstanceInstalledPackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceInstalledPackagesSortByEnum(val string) (ListManagedInstanceInstalledPackagesSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceInstalledPackagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
