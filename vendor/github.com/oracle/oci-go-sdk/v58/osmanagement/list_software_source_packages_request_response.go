// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListSoftwareSourcePackagesRequest wrapper for the ListSoftwareSourcePackages operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListSoftwareSourcePackages.go.html to see an example of how to use ListSoftwareSourcePackagesRequest.
type ListSoftwareSourcePackagesRequest struct {

	// The OCID of the software source.
	SoftwareSourceId *string `mandatory:"true" contributesTo:"path" name:"softwareSourceId"`

	// The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListSoftwareSourcePackagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListSoftwareSourcePackagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSoftwareSourcePackagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSoftwareSourcePackagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSoftwareSourcePackagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSoftwareSourcePackagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSoftwareSourcePackagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSoftwareSourcePackagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSoftwareSourcePackagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSoftwareSourcePackagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSoftwareSourcePackagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSoftwareSourcePackagesResponse wrapper for the ListSoftwareSourcePackages operation
type ListSoftwareSourcePackagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SoftwarePackageSummary instances
	Items []SoftwarePackageSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSoftwareSourcePackagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSoftwareSourcePackagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSoftwareSourcePackagesSortOrderEnum Enum with underlying type: string
type ListSoftwareSourcePackagesSortOrderEnum string

// Set of constants representing the allowable values for ListSoftwareSourcePackagesSortOrderEnum
const (
	ListSoftwareSourcePackagesSortOrderAsc  ListSoftwareSourcePackagesSortOrderEnum = "ASC"
	ListSoftwareSourcePackagesSortOrderDesc ListSoftwareSourcePackagesSortOrderEnum = "DESC"
)

var mappingListSoftwareSourcePackagesSortOrderEnum = map[string]ListSoftwareSourcePackagesSortOrderEnum{
	"ASC":  ListSoftwareSourcePackagesSortOrderAsc,
	"DESC": ListSoftwareSourcePackagesSortOrderDesc,
}

// GetListSoftwareSourcePackagesSortOrderEnumValues Enumerates the set of values for ListSoftwareSourcePackagesSortOrderEnum
func GetListSoftwareSourcePackagesSortOrderEnumValues() []ListSoftwareSourcePackagesSortOrderEnum {
	values := make([]ListSoftwareSourcePackagesSortOrderEnum, 0)
	for _, v := range mappingListSoftwareSourcePackagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareSourcePackagesSortOrderEnumStringValues Enumerates the set of values in String for ListSoftwareSourcePackagesSortOrderEnum
func GetListSoftwareSourcePackagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSoftwareSourcePackagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareSourcePackagesSortOrderEnum(val string) (ListSoftwareSourcePackagesSortOrderEnum, bool) {
	mappingListSoftwareSourcePackagesSortOrderEnumIgnoreCase := make(map[string]ListSoftwareSourcePackagesSortOrderEnum)
	for k, v := range mappingListSoftwareSourcePackagesSortOrderEnum {
		mappingListSoftwareSourcePackagesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSoftwareSourcePackagesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListSoftwareSourcePackagesSortByEnum Enum with underlying type: string
type ListSoftwareSourcePackagesSortByEnum string

// Set of constants representing the allowable values for ListSoftwareSourcePackagesSortByEnum
const (
	ListSoftwareSourcePackagesSortByTimecreated ListSoftwareSourcePackagesSortByEnum = "TIMECREATED"
	ListSoftwareSourcePackagesSortByDisplayname ListSoftwareSourcePackagesSortByEnum = "DISPLAYNAME"
)

var mappingListSoftwareSourcePackagesSortByEnum = map[string]ListSoftwareSourcePackagesSortByEnum{
	"TIMECREATED": ListSoftwareSourcePackagesSortByTimecreated,
	"DISPLAYNAME": ListSoftwareSourcePackagesSortByDisplayname,
}

// GetListSoftwareSourcePackagesSortByEnumValues Enumerates the set of values for ListSoftwareSourcePackagesSortByEnum
func GetListSoftwareSourcePackagesSortByEnumValues() []ListSoftwareSourcePackagesSortByEnum {
	values := make([]ListSoftwareSourcePackagesSortByEnum, 0)
	for _, v := range mappingListSoftwareSourcePackagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareSourcePackagesSortByEnumStringValues Enumerates the set of values in String for ListSoftwareSourcePackagesSortByEnum
func GetListSoftwareSourcePackagesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListSoftwareSourcePackagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareSourcePackagesSortByEnum(val string) (ListSoftwareSourcePackagesSortByEnum, bool) {
	mappingListSoftwareSourcePackagesSortByEnumIgnoreCase := make(map[string]ListSoftwareSourcePackagesSortByEnum)
	for k, v := range mappingListSoftwareSourcePackagesSortByEnum {
		mappingListSoftwareSourcePackagesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSoftwareSourcePackagesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
