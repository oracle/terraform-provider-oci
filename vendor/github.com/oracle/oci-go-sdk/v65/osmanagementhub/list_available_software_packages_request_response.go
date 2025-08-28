// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAvailableSoftwarePackagesRequest wrapper for the ListAvailableSoftwarePackages operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListAvailableSoftwarePackages.go.html to see an example of how to use ListAvailableSoftwarePackagesRequest.
type ListAvailableSoftwarePackagesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	SoftwareSourceId *string `mandatory:"true" contributesTo:"path" name:"softwareSourceId"`

	// A filter to return resources that match the given user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// Indicates whether to list only the latest versions of packages, module streams, and stream profiles.
	IsLatest *bool `mandatory:"false" contributesTo:"query" name:"isLatest"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAvailableSoftwarePackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListAvailableSoftwarePackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailableSoftwarePackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailableSoftwarePackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAvailableSoftwarePackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailableSoftwarePackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAvailableSoftwarePackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAvailableSoftwarePackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAvailableSoftwarePackagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAvailableSoftwarePackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAvailableSoftwarePackagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAvailableSoftwarePackagesResponse wrapper for the ListAvailableSoftwarePackages operation
type ListAvailableSoftwarePackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SoftwarePackageCollection instances
	SoftwarePackageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAvailableSoftwarePackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailableSoftwarePackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailableSoftwarePackagesSortOrderEnum Enum with underlying type: string
type ListAvailableSoftwarePackagesSortOrderEnum string

// Set of constants representing the allowable values for ListAvailableSoftwarePackagesSortOrderEnum
const (
	ListAvailableSoftwarePackagesSortOrderAsc  ListAvailableSoftwarePackagesSortOrderEnum = "ASC"
	ListAvailableSoftwarePackagesSortOrderDesc ListAvailableSoftwarePackagesSortOrderEnum = "DESC"
)

var mappingListAvailableSoftwarePackagesSortOrderEnum = map[string]ListAvailableSoftwarePackagesSortOrderEnum{
	"ASC":  ListAvailableSoftwarePackagesSortOrderAsc,
	"DESC": ListAvailableSoftwarePackagesSortOrderDesc,
}

var mappingListAvailableSoftwarePackagesSortOrderEnumLowerCase = map[string]ListAvailableSoftwarePackagesSortOrderEnum{
	"asc":  ListAvailableSoftwarePackagesSortOrderAsc,
	"desc": ListAvailableSoftwarePackagesSortOrderDesc,
}

// GetListAvailableSoftwarePackagesSortOrderEnumValues Enumerates the set of values for ListAvailableSoftwarePackagesSortOrderEnum
func GetListAvailableSoftwarePackagesSortOrderEnumValues() []ListAvailableSoftwarePackagesSortOrderEnum {
	values := make([]ListAvailableSoftwarePackagesSortOrderEnum, 0)
	for _, v := range mappingListAvailableSoftwarePackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableSoftwarePackagesSortOrderEnumStringValues Enumerates the set of values in String for ListAvailableSoftwarePackagesSortOrderEnum
func GetListAvailableSoftwarePackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAvailableSoftwarePackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableSoftwarePackagesSortOrderEnum(val string) (ListAvailableSoftwarePackagesSortOrderEnum, bool) {
	enum, ok := mappingListAvailableSoftwarePackagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAvailableSoftwarePackagesSortByEnum Enum with underlying type: string
type ListAvailableSoftwarePackagesSortByEnum string

// Set of constants representing the allowable values for ListAvailableSoftwarePackagesSortByEnum
const (
	ListAvailableSoftwarePackagesSortByTimecreated ListAvailableSoftwarePackagesSortByEnum = "timeCreated"
	ListAvailableSoftwarePackagesSortByDisplayname ListAvailableSoftwarePackagesSortByEnum = "displayName"
)

var mappingListAvailableSoftwarePackagesSortByEnum = map[string]ListAvailableSoftwarePackagesSortByEnum{
	"timeCreated": ListAvailableSoftwarePackagesSortByTimecreated,
	"displayName": ListAvailableSoftwarePackagesSortByDisplayname,
}

var mappingListAvailableSoftwarePackagesSortByEnumLowerCase = map[string]ListAvailableSoftwarePackagesSortByEnum{
	"timecreated": ListAvailableSoftwarePackagesSortByTimecreated,
	"displayname": ListAvailableSoftwarePackagesSortByDisplayname,
}

// GetListAvailableSoftwarePackagesSortByEnumValues Enumerates the set of values for ListAvailableSoftwarePackagesSortByEnum
func GetListAvailableSoftwarePackagesSortByEnumValues() []ListAvailableSoftwarePackagesSortByEnum {
	values := make([]ListAvailableSoftwarePackagesSortByEnum, 0)
	for _, v := range mappingListAvailableSoftwarePackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableSoftwarePackagesSortByEnumStringValues Enumerates the set of values in String for ListAvailableSoftwarePackagesSortByEnum
func GetListAvailableSoftwarePackagesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAvailableSoftwarePackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableSoftwarePackagesSortByEnum(val string) (ListAvailableSoftwarePackagesSortByEnum, bool) {
	enum, ok := mappingListAvailableSoftwarePackagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
