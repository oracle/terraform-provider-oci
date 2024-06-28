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

// ListGiVersionMinorVersionsRequest wrapper for the ListGiVersionMinorVersions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListGiVersionMinorVersions.go.html to see an example of how to use ListGiVersionMinorVersionsRequest.
type ListGiVersionMinorVersionsRequest struct {

	// The Oracle Grid Infrastructure major version.
	Version *string `mandatory:"true" contributesTo:"path" name:"version"`

	// The target availability domain. Only passed if the limit is AD-specific.
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// If provided, filters the results to the set of database versions which are supported for the given shape family.
	ShapeFamily ListGiVersionMinorVersionsShapeFamilyEnum `mandatory:"false" contributesTo:"query" name:"shapeFamily" omitEmpty:"true"`

	// If true, returns the Grid Infrastructure versions that can be used for provisioning a cluster
	IsGiVersionForProvisioning *bool `mandatory:"false" contributesTo:"query" name:"isGiVersionForProvisioning"`

	// If provided, filters the results for the given shape.
	Shape *string `mandatory:"false" contributesTo:"query" name:"shape"`

	// Sort by VERSION.  Default order for VERSION is descending.
	SortBy ListGiVersionMinorVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListGiVersionMinorVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGiVersionMinorVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGiVersionMinorVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGiVersionMinorVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGiVersionMinorVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGiVersionMinorVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGiVersionMinorVersionsShapeFamilyEnum(string(request.ShapeFamily)); !ok && request.ShapeFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShapeFamily: %s. Supported values are: %s.", request.ShapeFamily, strings.Join(GetListGiVersionMinorVersionsShapeFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGiVersionMinorVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGiVersionMinorVersionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGiVersionMinorVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGiVersionMinorVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGiVersionMinorVersionsResponse wrapper for the ListGiVersionMinorVersions operation
type ListGiVersionMinorVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []GiMinorVersionSummary instances
	Items []GiMinorVersionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGiVersionMinorVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGiVersionMinorVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGiVersionMinorVersionsShapeFamilyEnum Enum with underlying type: string
type ListGiVersionMinorVersionsShapeFamilyEnum string

// Set of constants representing the allowable values for ListGiVersionMinorVersionsShapeFamilyEnum
const (
	ListGiVersionMinorVersionsShapeFamilySinglenode     ListGiVersionMinorVersionsShapeFamilyEnum = "SINGLENODE"
	ListGiVersionMinorVersionsShapeFamilyYoda           ListGiVersionMinorVersionsShapeFamilyEnum = "YODA"
	ListGiVersionMinorVersionsShapeFamilyVirtualmachine ListGiVersionMinorVersionsShapeFamilyEnum = "VIRTUALMACHINE"
	ListGiVersionMinorVersionsShapeFamilyExadata        ListGiVersionMinorVersionsShapeFamilyEnum = "EXADATA"
	ListGiVersionMinorVersionsShapeFamilyExacc          ListGiVersionMinorVersionsShapeFamilyEnum = "EXACC"
	ListGiVersionMinorVersionsShapeFamilyExadbXs        ListGiVersionMinorVersionsShapeFamilyEnum = "EXADB_XS"
)

var mappingListGiVersionMinorVersionsShapeFamilyEnum = map[string]ListGiVersionMinorVersionsShapeFamilyEnum{
	"SINGLENODE":     ListGiVersionMinorVersionsShapeFamilySinglenode,
	"YODA":           ListGiVersionMinorVersionsShapeFamilyYoda,
	"VIRTUALMACHINE": ListGiVersionMinorVersionsShapeFamilyVirtualmachine,
	"EXADATA":        ListGiVersionMinorVersionsShapeFamilyExadata,
	"EXACC":          ListGiVersionMinorVersionsShapeFamilyExacc,
	"EXADB_XS":       ListGiVersionMinorVersionsShapeFamilyExadbXs,
}

var mappingListGiVersionMinorVersionsShapeFamilyEnumLowerCase = map[string]ListGiVersionMinorVersionsShapeFamilyEnum{
	"singlenode":     ListGiVersionMinorVersionsShapeFamilySinglenode,
	"yoda":           ListGiVersionMinorVersionsShapeFamilyYoda,
	"virtualmachine": ListGiVersionMinorVersionsShapeFamilyVirtualmachine,
	"exadata":        ListGiVersionMinorVersionsShapeFamilyExadata,
	"exacc":          ListGiVersionMinorVersionsShapeFamilyExacc,
	"exadb_xs":       ListGiVersionMinorVersionsShapeFamilyExadbXs,
}

// GetListGiVersionMinorVersionsShapeFamilyEnumValues Enumerates the set of values for ListGiVersionMinorVersionsShapeFamilyEnum
func GetListGiVersionMinorVersionsShapeFamilyEnumValues() []ListGiVersionMinorVersionsShapeFamilyEnum {
	values := make([]ListGiVersionMinorVersionsShapeFamilyEnum, 0)
	for _, v := range mappingListGiVersionMinorVersionsShapeFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetListGiVersionMinorVersionsShapeFamilyEnumStringValues Enumerates the set of values in String for ListGiVersionMinorVersionsShapeFamilyEnum
func GetListGiVersionMinorVersionsShapeFamilyEnumStringValues() []string {
	return []string{
		"SINGLENODE",
		"YODA",
		"VIRTUALMACHINE",
		"EXADATA",
		"EXACC",
		"EXADB_XS",
	}
}

// GetMappingListGiVersionMinorVersionsShapeFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGiVersionMinorVersionsShapeFamilyEnum(val string) (ListGiVersionMinorVersionsShapeFamilyEnum, bool) {
	enum, ok := mappingListGiVersionMinorVersionsShapeFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGiVersionMinorVersionsSortByEnum Enum with underlying type: string
type ListGiVersionMinorVersionsSortByEnum string

// Set of constants representing the allowable values for ListGiVersionMinorVersionsSortByEnum
const (
	ListGiVersionMinorVersionsSortByVersion ListGiVersionMinorVersionsSortByEnum = "VERSION"
)

var mappingListGiVersionMinorVersionsSortByEnum = map[string]ListGiVersionMinorVersionsSortByEnum{
	"VERSION": ListGiVersionMinorVersionsSortByVersion,
}

var mappingListGiVersionMinorVersionsSortByEnumLowerCase = map[string]ListGiVersionMinorVersionsSortByEnum{
	"version": ListGiVersionMinorVersionsSortByVersion,
}

// GetListGiVersionMinorVersionsSortByEnumValues Enumerates the set of values for ListGiVersionMinorVersionsSortByEnum
func GetListGiVersionMinorVersionsSortByEnumValues() []ListGiVersionMinorVersionsSortByEnum {
	values := make([]ListGiVersionMinorVersionsSortByEnum, 0)
	for _, v := range mappingListGiVersionMinorVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGiVersionMinorVersionsSortByEnumStringValues Enumerates the set of values in String for ListGiVersionMinorVersionsSortByEnum
func GetListGiVersionMinorVersionsSortByEnumStringValues() []string {
	return []string{
		"VERSION",
	}
}

// GetMappingListGiVersionMinorVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGiVersionMinorVersionsSortByEnum(val string) (ListGiVersionMinorVersionsSortByEnum, bool) {
	enum, ok := mappingListGiVersionMinorVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGiVersionMinorVersionsSortOrderEnum Enum with underlying type: string
type ListGiVersionMinorVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListGiVersionMinorVersionsSortOrderEnum
const (
	ListGiVersionMinorVersionsSortOrderAsc  ListGiVersionMinorVersionsSortOrderEnum = "ASC"
	ListGiVersionMinorVersionsSortOrderDesc ListGiVersionMinorVersionsSortOrderEnum = "DESC"
)

var mappingListGiVersionMinorVersionsSortOrderEnum = map[string]ListGiVersionMinorVersionsSortOrderEnum{
	"ASC":  ListGiVersionMinorVersionsSortOrderAsc,
	"DESC": ListGiVersionMinorVersionsSortOrderDesc,
}

var mappingListGiVersionMinorVersionsSortOrderEnumLowerCase = map[string]ListGiVersionMinorVersionsSortOrderEnum{
	"asc":  ListGiVersionMinorVersionsSortOrderAsc,
	"desc": ListGiVersionMinorVersionsSortOrderDesc,
}

// GetListGiVersionMinorVersionsSortOrderEnumValues Enumerates the set of values for ListGiVersionMinorVersionsSortOrderEnum
func GetListGiVersionMinorVersionsSortOrderEnumValues() []ListGiVersionMinorVersionsSortOrderEnum {
	values := make([]ListGiVersionMinorVersionsSortOrderEnum, 0)
	for _, v := range mappingListGiVersionMinorVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGiVersionMinorVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListGiVersionMinorVersionsSortOrderEnum
func GetListGiVersionMinorVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGiVersionMinorVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGiVersionMinorVersionsSortOrderEnum(val string) (ListGiVersionMinorVersionsSortOrderEnum, bool) {
	enum, ok := mappingListGiVersionMinorVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
