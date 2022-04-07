// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDatabaseSoftwareImagesRequest wrapper for the ListDatabaseSoftwareImages operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDatabaseSoftwareImages.go.html to see an example of how to use ListDatabaseSoftwareImagesRequest.
type ListDatabaseSoftwareImagesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListDatabaseSoftwareImagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDatabaseSoftwareImagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState DatabaseSoftwareImageSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given image type exactly.
	ImageType DatabaseSoftwareImageSummaryImageTypeEnum `mandatory:"false" contributesTo:"query" name:"imageType" omitEmpty:"true"`

	// A filter to return only resources that match the given image shape family exactly.
	ImageShapeFamily DatabaseSoftwareImageSummaryImageShapeFamilyEnum `mandatory:"false" contributesTo:"query" name:"imageShapeFamily" omitEmpty:"true"`

	// If provided, filters the results to the set of database versions which are supported for Upgrade.
	IsUpgradeSupported *bool `mandatory:"false" contributesTo:"query" name:"isUpgradeSupported"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseSoftwareImagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseSoftwareImagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseSoftwareImagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseSoftwareImagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseSoftwareImagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseSoftwareImagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseSoftwareImagesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseSoftwareImagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseSoftwareImagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSoftwareImageSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDatabaseSoftwareImageSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSoftwareImageSummaryImageTypeEnum(string(request.ImageType)); !ok && request.ImageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageType: %s. Supported values are: %s.", request.ImageType, strings.Join(GetDatabaseSoftwareImageSummaryImageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSoftwareImageSummaryImageShapeFamilyEnum(string(request.ImageShapeFamily)); !ok && request.ImageShapeFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageShapeFamily: %s. Supported values are: %s.", request.ImageShapeFamily, strings.Join(GetDatabaseSoftwareImageSummaryImageShapeFamilyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseSoftwareImagesResponse wrapper for the ListDatabaseSoftwareImages operation
type ListDatabaseSoftwareImagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DatabaseSoftwareImageSummary instances
	Items []DatabaseSoftwareImageSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseSoftwareImagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseSoftwareImagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseSoftwareImagesSortByEnum Enum with underlying type: string
type ListDatabaseSoftwareImagesSortByEnum string

// Set of constants representing the allowable values for ListDatabaseSoftwareImagesSortByEnum
const (
	ListDatabaseSoftwareImagesSortByTimecreated ListDatabaseSoftwareImagesSortByEnum = "TIMECREATED"
	ListDatabaseSoftwareImagesSortByDisplayname ListDatabaseSoftwareImagesSortByEnum = "DISPLAYNAME"
)

var mappingListDatabaseSoftwareImagesSortByEnum = map[string]ListDatabaseSoftwareImagesSortByEnum{
	"TIMECREATED": ListDatabaseSoftwareImagesSortByTimecreated,
	"DISPLAYNAME": ListDatabaseSoftwareImagesSortByDisplayname,
}

var mappingListDatabaseSoftwareImagesSortByEnumLowerCase = map[string]ListDatabaseSoftwareImagesSortByEnum{
	"timecreated": ListDatabaseSoftwareImagesSortByTimecreated,
	"displayname": ListDatabaseSoftwareImagesSortByDisplayname,
}

// GetListDatabaseSoftwareImagesSortByEnumValues Enumerates the set of values for ListDatabaseSoftwareImagesSortByEnum
func GetListDatabaseSoftwareImagesSortByEnumValues() []ListDatabaseSoftwareImagesSortByEnum {
	values := make([]ListDatabaseSoftwareImagesSortByEnum, 0)
	for _, v := range mappingListDatabaseSoftwareImagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseSoftwareImagesSortByEnumStringValues Enumerates the set of values in String for ListDatabaseSoftwareImagesSortByEnum
func GetListDatabaseSoftwareImagesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListDatabaseSoftwareImagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseSoftwareImagesSortByEnum(val string) (ListDatabaseSoftwareImagesSortByEnum, bool) {
	enum, ok := mappingListDatabaseSoftwareImagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseSoftwareImagesSortOrderEnum Enum with underlying type: string
type ListDatabaseSoftwareImagesSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseSoftwareImagesSortOrderEnum
const (
	ListDatabaseSoftwareImagesSortOrderAsc  ListDatabaseSoftwareImagesSortOrderEnum = "ASC"
	ListDatabaseSoftwareImagesSortOrderDesc ListDatabaseSoftwareImagesSortOrderEnum = "DESC"
)

var mappingListDatabaseSoftwareImagesSortOrderEnum = map[string]ListDatabaseSoftwareImagesSortOrderEnum{
	"ASC":  ListDatabaseSoftwareImagesSortOrderAsc,
	"DESC": ListDatabaseSoftwareImagesSortOrderDesc,
}

var mappingListDatabaseSoftwareImagesSortOrderEnumLowerCase = map[string]ListDatabaseSoftwareImagesSortOrderEnum{
	"asc":  ListDatabaseSoftwareImagesSortOrderAsc,
	"desc": ListDatabaseSoftwareImagesSortOrderDesc,
}

// GetListDatabaseSoftwareImagesSortOrderEnumValues Enumerates the set of values for ListDatabaseSoftwareImagesSortOrderEnum
func GetListDatabaseSoftwareImagesSortOrderEnumValues() []ListDatabaseSoftwareImagesSortOrderEnum {
	values := make([]ListDatabaseSoftwareImagesSortOrderEnum, 0)
	for _, v := range mappingListDatabaseSoftwareImagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseSoftwareImagesSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseSoftwareImagesSortOrderEnum
func GetListDatabaseSoftwareImagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseSoftwareImagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseSoftwareImagesSortOrderEnum(val string) (ListDatabaseSoftwareImagesSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseSoftwareImagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
