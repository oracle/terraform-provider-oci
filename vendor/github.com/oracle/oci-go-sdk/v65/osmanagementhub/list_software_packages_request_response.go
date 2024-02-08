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

// ListSoftwarePackagesRequest wrapper for the ListSoftwarePackages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListSoftwarePackages.go.html to see an example of how to use ListSoftwarePackagesRequest.
type ListSoftwarePackagesRequest struct {

	// The software source OCID.
	SoftwareSourceId *string `mandatory:"true" contributesTo:"path" name:"softwareSourceId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A boolean variable that is used to list only the latest versions of packages, module streams,
	// and stream profiles when set to true. All packages, module streams, and stream profiles are
	// returned when set to false.
	IsLatest *bool `mandatory:"false" contributesTo:"query" name:"isLatest"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListSoftwarePackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListSoftwarePackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSoftwarePackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSoftwarePackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSoftwarePackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSoftwarePackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSoftwarePackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSoftwarePackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSoftwarePackagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSoftwarePackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSoftwarePackagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSoftwarePackagesResponse wrapper for the ListSoftwarePackages operation
type ListSoftwarePackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SoftwarePackageCollection instances
	SoftwarePackageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSoftwarePackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSoftwarePackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSoftwarePackagesSortOrderEnum Enum with underlying type: string
type ListSoftwarePackagesSortOrderEnum string

// Set of constants representing the allowable values for ListSoftwarePackagesSortOrderEnum
const (
	ListSoftwarePackagesSortOrderAsc  ListSoftwarePackagesSortOrderEnum = "ASC"
	ListSoftwarePackagesSortOrderDesc ListSoftwarePackagesSortOrderEnum = "DESC"
)

var mappingListSoftwarePackagesSortOrderEnum = map[string]ListSoftwarePackagesSortOrderEnum{
	"ASC":  ListSoftwarePackagesSortOrderAsc,
	"DESC": ListSoftwarePackagesSortOrderDesc,
}

var mappingListSoftwarePackagesSortOrderEnumLowerCase = map[string]ListSoftwarePackagesSortOrderEnum{
	"asc":  ListSoftwarePackagesSortOrderAsc,
	"desc": ListSoftwarePackagesSortOrderDesc,
}

// GetListSoftwarePackagesSortOrderEnumValues Enumerates the set of values for ListSoftwarePackagesSortOrderEnum
func GetListSoftwarePackagesSortOrderEnumValues() []ListSoftwarePackagesSortOrderEnum {
	values := make([]ListSoftwarePackagesSortOrderEnum, 0)
	for _, v := range mappingListSoftwarePackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwarePackagesSortOrderEnumStringValues Enumerates the set of values in String for ListSoftwarePackagesSortOrderEnum
func GetListSoftwarePackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSoftwarePackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwarePackagesSortOrderEnum(val string) (ListSoftwarePackagesSortOrderEnum, bool) {
	enum, ok := mappingListSoftwarePackagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSoftwarePackagesSortByEnum Enum with underlying type: string
type ListSoftwarePackagesSortByEnum string

// Set of constants representing the allowable values for ListSoftwarePackagesSortByEnum
const (
	ListSoftwarePackagesSortByTimecreated ListSoftwarePackagesSortByEnum = "timeCreated"
	ListSoftwarePackagesSortByDisplayname ListSoftwarePackagesSortByEnum = "displayName"
)

var mappingListSoftwarePackagesSortByEnum = map[string]ListSoftwarePackagesSortByEnum{
	"timeCreated": ListSoftwarePackagesSortByTimecreated,
	"displayName": ListSoftwarePackagesSortByDisplayname,
}

var mappingListSoftwarePackagesSortByEnumLowerCase = map[string]ListSoftwarePackagesSortByEnum{
	"timecreated": ListSoftwarePackagesSortByTimecreated,
	"displayname": ListSoftwarePackagesSortByDisplayname,
}

// GetListSoftwarePackagesSortByEnumValues Enumerates the set of values for ListSoftwarePackagesSortByEnum
func GetListSoftwarePackagesSortByEnumValues() []ListSoftwarePackagesSortByEnum {
	values := make([]ListSoftwarePackagesSortByEnum, 0)
	for _, v := range mappingListSoftwarePackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwarePackagesSortByEnumStringValues Enumerates the set of values in String for ListSoftwarePackagesSortByEnum
func GetListSoftwarePackagesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSoftwarePackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwarePackagesSortByEnum(val string) (ListSoftwarePackagesSortByEnum, bool) {
	enum, ok := mappingListSoftwarePackagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
