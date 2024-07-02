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

// ListExascaleDbStorageVaultsRequest wrapper for the ListExascaleDbStorageVaults operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExascaleDbStorageVaults.go.html to see an example of how to use ListExascaleDbStorageVaultsRequest.
type ListExascaleDbStorageVaultsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListExascaleDbStorageVaultsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListExascaleDbStorageVaultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only Exadata Database Storage Vaults that match the given lifecycle state exactly.
	LifecycleState ExascaleDbStorageVaultLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExascaleDbStorageVaultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExascaleDbStorageVaultsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExascaleDbStorageVaultsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExascaleDbStorageVaultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExascaleDbStorageVaultsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExascaleDbStorageVaultsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExascaleDbStorageVaultsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExascaleDbStorageVaultsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExascaleDbStorageVaultsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExascaleDbStorageVaultLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetExascaleDbStorageVaultLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExascaleDbStorageVaultsResponse wrapper for the ListExascaleDbStorageVaults operation
type ListExascaleDbStorageVaultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExascaleDbStorageVaultSummary instances
	Items []ExascaleDbStorageVaultSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExascaleDbStorageVaultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExascaleDbStorageVaultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExascaleDbStorageVaultsSortByEnum Enum with underlying type: string
type ListExascaleDbStorageVaultsSortByEnum string

// Set of constants representing the allowable values for ListExascaleDbStorageVaultsSortByEnum
const (
	ListExascaleDbStorageVaultsSortByTimecreated ListExascaleDbStorageVaultsSortByEnum = "TIMECREATED"
	ListExascaleDbStorageVaultsSortByDisplayname ListExascaleDbStorageVaultsSortByEnum = "DISPLAYNAME"
)

var mappingListExascaleDbStorageVaultsSortByEnum = map[string]ListExascaleDbStorageVaultsSortByEnum{
	"TIMECREATED": ListExascaleDbStorageVaultsSortByTimecreated,
	"DISPLAYNAME": ListExascaleDbStorageVaultsSortByDisplayname,
}

var mappingListExascaleDbStorageVaultsSortByEnumLowerCase = map[string]ListExascaleDbStorageVaultsSortByEnum{
	"timecreated": ListExascaleDbStorageVaultsSortByTimecreated,
	"displayname": ListExascaleDbStorageVaultsSortByDisplayname,
}

// GetListExascaleDbStorageVaultsSortByEnumValues Enumerates the set of values for ListExascaleDbStorageVaultsSortByEnum
func GetListExascaleDbStorageVaultsSortByEnumValues() []ListExascaleDbStorageVaultsSortByEnum {
	values := make([]ListExascaleDbStorageVaultsSortByEnum, 0)
	for _, v := range mappingListExascaleDbStorageVaultsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExascaleDbStorageVaultsSortByEnumStringValues Enumerates the set of values in String for ListExascaleDbStorageVaultsSortByEnum
func GetListExascaleDbStorageVaultsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListExascaleDbStorageVaultsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExascaleDbStorageVaultsSortByEnum(val string) (ListExascaleDbStorageVaultsSortByEnum, bool) {
	enum, ok := mappingListExascaleDbStorageVaultsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExascaleDbStorageVaultsSortOrderEnum Enum with underlying type: string
type ListExascaleDbStorageVaultsSortOrderEnum string

// Set of constants representing the allowable values for ListExascaleDbStorageVaultsSortOrderEnum
const (
	ListExascaleDbStorageVaultsSortOrderAsc  ListExascaleDbStorageVaultsSortOrderEnum = "ASC"
	ListExascaleDbStorageVaultsSortOrderDesc ListExascaleDbStorageVaultsSortOrderEnum = "DESC"
)

var mappingListExascaleDbStorageVaultsSortOrderEnum = map[string]ListExascaleDbStorageVaultsSortOrderEnum{
	"ASC":  ListExascaleDbStorageVaultsSortOrderAsc,
	"DESC": ListExascaleDbStorageVaultsSortOrderDesc,
}

var mappingListExascaleDbStorageVaultsSortOrderEnumLowerCase = map[string]ListExascaleDbStorageVaultsSortOrderEnum{
	"asc":  ListExascaleDbStorageVaultsSortOrderAsc,
	"desc": ListExascaleDbStorageVaultsSortOrderDesc,
}

// GetListExascaleDbStorageVaultsSortOrderEnumValues Enumerates the set of values for ListExascaleDbStorageVaultsSortOrderEnum
func GetListExascaleDbStorageVaultsSortOrderEnumValues() []ListExascaleDbStorageVaultsSortOrderEnum {
	values := make([]ListExascaleDbStorageVaultsSortOrderEnum, 0)
	for _, v := range mappingListExascaleDbStorageVaultsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExascaleDbStorageVaultsSortOrderEnumStringValues Enumerates the set of values in String for ListExascaleDbStorageVaultsSortOrderEnum
func GetListExascaleDbStorageVaultsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExascaleDbStorageVaultsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExascaleDbStorageVaultsSortOrderEnum(val string) (ListExascaleDbStorageVaultsSortOrderEnum, bool) {
	enum, ok := mappingListExascaleDbStorageVaultsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
