// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListGiHomesRequest wrapper for the ListGiHomes operation
type ListGiHomesRequest struct {

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The DB system OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). If provided, filters the results to the set of database versions which are supported for the DB system.
	DbSystemId *string `mandatory:"false" contributesTo:"query" name:"dbSystemId"`

	// The Grid Infrastructure version. If provided, filters the results matching the provided grid infrastructure version.
	GiVersion *string `mandatory:"false" contributesTo:"query" name:"giVersion"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for timeCreated is descending.  Default order for displayName is ascending. The displayName sort order is case sensitive.
	SortBy ListGiHomesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListGiHomesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState GiHomeSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGiHomesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGiHomesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGiHomesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGiHomesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGiHomesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGiHomesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGiHomesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGiHomesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGiHomesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGiHomeSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetGiHomeSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGiHomesResponse wrapper for the ListGiHomes operation
type ListGiHomesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GiHomeCollection instances
	GiHomeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGiHomesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGiHomesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGiHomesSortByEnum Enum with underlying type: string
type ListGiHomesSortByEnum string

// Set of constants representing the allowable values for ListGiHomesSortByEnum
const (
	ListGiHomesSortByTimecreated ListGiHomesSortByEnum = "timeCreated"
	ListGiHomesSortByDisplayname ListGiHomesSortByEnum = "displayName"
)

var mappingListGiHomesSortByEnum = map[string]ListGiHomesSortByEnum{
	"timeCreated": ListGiHomesSortByTimecreated,
	"displayName": ListGiHomesSortByDisplayname,
}

var mappingListGiHomesSortByEnumLowerCase = map[string]ListGiHomesSortByEnum{
	"timecreated": ListGiHomesSortByTimecreated,
	"displayname": ListGiHomesSortByDisplayname,
}

// GetListGiHomesSortByEnumValues Enumerates the set of values for ListGiHomesSortByEnum
func GetListGiHomesSortByEnumValues() []ListGiHomesSortByEnum {
	values := make([]ListGiHomesSortByEnum, 0)
	for _, v := range mappingListGiHomesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGiHomesSortByEnumStringValues Enumerates the set of values in String for ListGiHomesSortByEnum
func GetListGiHomesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListGiHomesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGiHomesSortByEnum(val string) (ListGiHomesSortByEnum, bool) {
	enum, ok := mappingListGiHomesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGiHomesSortOrderEnum Enum with underlying type: string
type ListGiHomesSortOrderEnum string

// Set of constants representing the allowable values for ListGiHomesSortOrderEnum
const (
	ListGiHomesSortOrderAsc  ListGiHomesSortOrderEnum = "ASC"
	ListGiHomesSortOrderDesc ListGiHomesSortOrderEnum = "DESC"
)

var mappingListGiHomesSortOrderEnum = map[string]ListGiHomesSortOrderEnum{
	"ASC":  ListGiHomesSortOrderAsc,
	"DESC": ListGiHomesSortOrderDesc,
}

var mappingListGiHomesSortOrderEnumLowerCase = map[string]ListGiHomesSortOrderEnum{
	"asc":  ListGiHomesSortOrderAsc,
	"desc": ListGiHomesSortOrderDesc,
}

// GetListGiHomesSortOrderEnumValues Enumerates the set of values for ListGiHomesSortOrderEnum
func GetListGiHomesSortOrderEnumValues() []ListGiHomesSortOrderEnum {
	values := make([]ListGiHomesSortOrderEnum, 0)
	for _, v := range mappingListGiHomesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGiHomesSortOrderEnumStringValues Enumerates the set of values in String for ListGiHomesSortOrderEnum
func GetListGiHomesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGiHomesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGiHomesSortOrderEnum(val string) (ListGiHomesSortOrderEnum, bool) {
	enum, ok := mappingListGiHomesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
