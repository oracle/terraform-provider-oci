// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAutonomousDatabaseSoftwareImagesRequest wrapper for the ListAutonomousDatabaseSoftwareImages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousDatabaseSoftwareImages.go.html to see an example of how to use ListAutonomousDatabaseSoftwareImagesRequest.
type ListAutonomousDatabaseSoftwareImagesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given image shape family exactly.
	ImageShapeFamily AutonomousDatabaseSoftwareImageImageShapeFamilyEnum `mandatory:"true" contributesTo:"query" name:"imageShapeFamily" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAutonomousDatabaseSoftwareImagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// parameter according to which Autonomous Database Software Images will be sorted.
	SortBy ListAutonomousDatabaseSoftwareImagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState AutonomousDatabaseSoftwareImageLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutonomousDatabaseSoftwareImagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutonomousDatabaseSoftwareImagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutonomousDatabaseSoftwareImagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutonomousDatabaseSoftwareImagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutonomousDatabaseSoftwareImagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousDatabaseSoftwareImageImageShapeFamilyEnum(string(request.ImageShapeFamily)); !ok && request.ImageShapeFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageShapeFamily: %s. Supported values are: %s.", request.ImageShapeFamily, strings.Join(GetAutonomousDatabaseSoftwareImageImageShapeFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutonomousDatabaseSoftwareImagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAutonomousDatabaseSoftwareImagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutonomousDatabaseSoftwareImagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAutonomousDatabaseSoftwareImagesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseSoftwareImageLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAutonomousDatabaseSoftwareImageLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutonomousDatabaseSoftwareImagesResponse wrapper for the ListAutonomousDatabaseSoftwareImages operation
type ListAutonomousDatabaseSoftwareImagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AutonomousDatabaseSoftwareImageCollection instances
	AutonomousDatabaseSoftwareImageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutonomousDatabaseSoftwareImagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutonomousDatabaseSoftwareImagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutonomousDatabaseSoftwareImagesSortOrderEnum Enum with underlying type: string
type ListAutonomousDatabaseSoftwareImagesSortOrderEnum string

// Set of constants representing the allowable values for ListAutonomousDatabaseSoftwareImagesSortOrderEnum
const (
	ListAutonomousDatabaseSoftwareImagesSortOrderAsc  ListAutonomousDatabaseSoftwareImagesSortOrderEnum = "ASC"
	ListAutonomousDatabaseSoftwareImagesSortOrderDesc ListAutonomousDatabaseSoftwareImagesSortOrderEnum = "DESC"
)

var mappingListAutonomousDatabaseSoftwareImagesSortOrderEnum = map[string]ListAutonomousDatabaseSoftwareImagesSortOrderEnum{
	"ASC":  ListAutonomousDatabaseSoftwareImagesSortOrderAsc,
	"DESC": ListAutonomousDatabaseSoftwareImagesSortOrderDesc,
}

var mappingListAutonomousDatabaseSoftwareImagesSortOrderEnumLowerCase = map[string]ListAutonomousDatabaseSoftwareImagesSortOrderEnum{
	"asc":  ListAutonomousDatabaseSoftwareImagesSortOrderAsc,
	"desc": ListAutonomousDatabaseSoftwareImagesSortOrderDesc,
}

// GetListAutonomousDatabaseSoftwareImagesSortOrderEnumValues Enumerates the set of values for ListAutonomousDatabaseSoftwareImagesSortOrderEnum
func GetListAutonomousDatabaseSoftwareImagesSortOrderEnumValues() []ListAutonomousDatabaseSoftwareImagesSortOrderEnum {
	values := make([]ListAutonomousDatabaseSoftwareImagesSortOrderEnum, 0)
	for _, v := range mappingListAutonomousDatabaseSoftwareImagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousDatabaseSoftwareImagesSortOrderEnumStringValues Enumerates the set of values in String for ListAutonomousDatabaseSoftwareImagesSortOrderEnum
func GetListAutonomousDatabaseSoftwareImagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAutonomousDatabaseSoftwareImagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousDatabaseSoftwareImagesSortOrderEnum(val string) (ListAutonomousDatabaseSoftwareImagesSortOrderEnum, bool) {
	enum, ok := mappingListAutonomousDatabaseSoftwareImagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutonomousDatabaseSoftwareImagesSortByEnum Enum with underlying type: string
type ListAutonomousDatabaseSoftwareImagesSortByEnum string

// Set of constants representing the allowable values for ListAutonomousDatabaseSoftwareImagesSortByEnum
const (
	ListAutonomousDatabaseSoftwareImagesSortByTimecreated ListAutonomousDatabaseSoftwareImagesSortByEnum = "TIMECREATED"
	ListAutonomousDatabaseSoftwareImagesSortByDisplayname ListAutonomousDatabaseSoftwareImagesSortByEnum = "DISPLAYNAME"
)

var mappingListAutonomousDatabaseSoftwareImagesSortByEnum = map[string]ListAutonomousDatabaseSoftwareImagesSortByEnum{
	"TIMECREATED": ListAutonomousDatabaseSoftwareImagesSortByTimecreated,
	"DISPLAYNAME": ListAutonomousDatabaseSoftwareImagesSortByDisplayname,
}

var mappingListAutonomousDatabaseSoftwareImagesSortByEnumLowerCase = map[string]ListAutonomousDatabaseSoftwareImagesSortByEnum{
	"timecreated": ListAutonomousDatabaseSoftwareImagesSortByTimecreated,
	"displayname": ListAutonomousDatabaseSoftwareImagesSortByDisplayname,
}

// GetListAutonomousDatabaseSoftwareImagesSortByEnumValues Enumerates the set of values for ListAutonomousDatabaseSoftwareImagesSortByEnum
func GetListAutonomousDatabaseSoftwareImagesSortByEnumValues() []ListAutonomousDatabaseSoftwareImagesSortByEnum {
	values := make([]ListAutonomousDatabaseSoftwareImagesSortByEnum, 0)
	for _, v := range mappingListAutonomousDatabaseSoftwareImagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousDatabaseSoftwareImagesSortByEnumStringValues Enumerates the set of values in String for ListAutonomousDatabaseSoftwareImagesSortByEnum
func GetListAutonomousDatabaseSoftwareImagesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAutonomousDatabaseSoftwareImagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousDatabaseSoftwareImagesSortByEnum(val string) (ListAutonomousDatabaseSoftwareImagesSortByEnum, bool) {
	enum, ok := mappingListAutonomousDatabaseSoftwareImagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
