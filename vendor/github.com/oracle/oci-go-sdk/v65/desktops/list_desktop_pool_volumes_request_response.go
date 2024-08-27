// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package desktops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDesktopPoolVolumesRequest wrapper for the ListDesktopPoolVolumes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/desktops/ListDesktopPoolVolumes.go.html to see an example of how to use ListDesktopPoolVolumesRequest.
type ListDesktopPoolVolumesRequest struct {

	// The OCID of the desktop pool.
	DesktopPoolId *string `mandatory:"true" contributesTo:"path" name:"desktopPoolId"`

	// The OCID of the compartment of the desktop pool.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain.
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// A filter to return only results with the given displayName.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only results with the given OCID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only results with the given lifecycleState.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`

	// The maximum number of results to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A field to sort by.
	SortBy ListDesktopPoolVolumesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A field to indicate the sort order.
	SortOrder ListDesktopPoolVolumesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The unique identifier of the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDesktopPoolVolumesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDesktopPoolVolumesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDesktopPoolVolumesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDesktopPoolVolumesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDesktopPoolVolumesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDesktopPoolVolumesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDesktopPoolVolumesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDesktopPoolVolumesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDesktopPoolVolumesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDesktopPoolVolumesResponse wrapper for the ListDesktopPoolVolumes operation
type ListDesktopPoolVolumesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DesktopPoolVolumeCollection instances
	DesktopPoolVolumeCollection `presentIn:"body"`

	// The unique identifier of the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDesktopPoolVolumesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDesktopPoolVolumesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDesktopPoolVolumesSortByEnum Enum with underlying type: string
type ListDesktopPoolVolumesSortByEnum string

// Set of constants representing the allowable values for ListDesktopPoolVolumesSortByEnum
const (
	ListDesktopPoolVolumesSortByTimecreated ListDesktopPoolVolumesSortByEnum = "TIMECREATED"
	ListDesktopPoolVolumesSortByDisplayname ListDesktopPoolVolumesSortByEnum = "DISPLAYNAME"
)

var mappingListDesktopPoolVolumesSortByEnum = map[string]ListDesktopPoolVolumesSortByEnum{
	"TIMECREATED": ListDesktopPoolVolumesSortByTimecreated,
	"DISPLAYNAME": ListDesktopPoolVolumesSortByDisplayname,
}

var mappingListDesktopPoolVolumesSortByEnumLowerCase = map[string]ListDesktopPoolVolumesSortByEnum{
	"timecreated": ListDesktopPoolVolumesSortByTimecreated,
	"displayname": ListDesktopPoolVolumesSortByDisplayname,
}

// GetListDesktopPoolVolumesSortByEnumValues Enumerates the set of values for ListDesktopPoolVolumesSortByEnum
func GetListDesktopPoolVolumesSortByEnumValues() []ListDesktopPoolVolumesSortByEnum {
	values := make([]ListDesktopPoolVolumesSortByEnum, 0)
	for _, v := range mappingListDesktopPoolVolumesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDesktopPoolVolumesSortByEnumStringValues Enumerates the set of values in String for ListDesktopPoolVolumesSortByEnum
func GetListDesktopPoolVolumesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListDesktopPoolVolumesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDesktopPoolVolumesSortByEnum(val string) (ListDesktopPoolVolumesSortByEnum, bool) {
	enum, ok := mappingListDesktopPoolVolumesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDesktopPoolVolumesSortOrderEnum Enum with underlying type: string
type ListDesktopPoolVolumesSortOrderEnum string

// Set of constants representing the allowable values for ListDesktopPoolVolumesSortOrderEnum
const (
	ListDesktopPoolVolumesSortOrderAsc  ListDesktopPoolVolumesSortOrderEnum = "ASC"
	ListDesktopPoolVolumesSortOrderDesc ListDesktopPoolVolumesSortOrderEnum = "DESC"
)

var mappingListDesktopPoolVolumesSortOrderEnum = map[string]ListDesktopPoolVolumesSortOrderEnum{
	"ASC":  ListDesktopPoolVolumesSortOrderAsc,
	"DESC": ListDesktopPoolVolumesSortOrderDesc,
}

var mappingListDesktopPoolVolumesSortOrderEnumLowerCase = map[string]ListDesktopPoolVolumesSortOrderEnum{
	"asc":  ListDesktopPoolVolumesSortOrderAsc,
	"desc": ListDesktopPoolVolumesSortOrderDesc,
}

// GetListDesktopPoolVolumesSortOrderEnumValues Enumerates the set of values for ListDesktopPoolVolumesSortOrderEnum
func GetListDesktopPoolVolumesSortOrderEnumValues() []ListDesktopPoolVolumesSortOrderEnum {
	values := make([]ListDesktopPoolVolumesSortOrderEnum, 0)
	for _, v := range mappingListDesktopPoolVolumesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDesktopPoolVolumesSortOrderEnumStringValues Enumerates the set of values in String for ListDesktopPoolVolumesSortOrderEnum
func GetListDesktopPoolVolumesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDesktopPoolVolumesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDesktopPoolVolumesSortOrderEnum(val string) (ListDesktopPoolVolumesSortOrderEnum, bool) {
	enum, ok := mappingListDesktopPoolVolumesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
