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

// ListExadbVmClustersRequest wrapper for the ListExadbVmClusters operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExadbVmClusters.go.html to see an example of how to use ListExadbVmClustersRequest.
type ListExadbVmClustersRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListExadbVmClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListExadbVmClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only Exadata VM clusters on Exascale Infrastructure that match the given lifecycle state exactly.
	LifecycleState ExadbVmClusterSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only Exadata VM clusters on Exascale Infrastructure that match the given Exascale Database Storage Vault ID.
	ExascaleDbStorageVaultId *string `mandatory:"false" contributesTo:"query" name:"exascaleDbStorageVaultId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExadbVmClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExadbVmClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExadbVmClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExadbVmClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExadbVmClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExadbVmClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExadbVmClustersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExadbVmClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExadbVmClustersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadbVmClusterSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetExadbVmClusterSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExadbVmClustersResponse wrapper for the ListExadbVmClusters operation
type ListExadbVmClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExadbVmClusterSummary instances
	Items []ExadbVmClusterSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExadbVmClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExadbVmClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExadbVmClustersSortByEnum Enum with underlying type: string
type ListExadbVmClustersSortByEnum string

// Set of constants representing the allowable values for ListExadbVmClustersSortByEnum
const (
	ListExadbVmClustersSortByTimecreated ListExadbVmClustersSortByEnum = "TIMECREATED"
	ListExadbVmClustersSortByDisplayname ListExadbVmClustersSortByEnum = "DISPLAYNAME"
)

var mappingListExadbVmClustersSortByEnum = map[string]ListExadbVmClustersSortByEnum{
	"TIMECREATED": ListExadbVmClustersSortByTimecreated,
	"DISPLAYNAME": ListExadbVmClustersSortByDisplayname,
}

var mappingListExadbVmClustersSortByEnumLowerCase = map[string]ListExadbVmClustersSortByEnum{
	"timecreated": ListExadbVmClustersSortByTimecreated,
	"displayname": ListExadbVmClustersSortByDisplayname,
}

// GetListExadbVmClustersSortByEnumValues Enumerates the set of values for ListExadbVmClustersSortByEnum
func GetListExadbVmClustersSortByEnumValues() []ListExadbVmClustersSortByEnum {
	values := make([]ListExadbVmClustersSortByEnum, 0)
	for _, v := range mappingListExadbVmClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExadbVmClustersSortByEnumStringValues Enumerates the set of values in String for ListExadbVmClustersSortByEnum
func GetListExadbVmClustersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListExadbVmClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExadbVmClustersSortByEnum(val string) (ListExadbVmClustersSortByEnum, bool) {
	enum, ok := mappingListExadbVmClustersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExadbVmClustersSortOrderEnum Enum with underlying type: string
type ListExadbVmClustersSortOrderEnum string

// Set of constants representing the allowable values for ListExadbVmClustersSortOrderEnum
const (
	ListExadbVmClustersSortOrderAsc  ListExadbVmClustersSortOrderEnum = "ASC"
	ListExadbVmClustersSortOrderDesc ListExadbVmClustersSortOrderEnum = "DESC"
)

var mappingListExadbVmClustersSortOrderEnum = map[string]ListExadbVmClustersSortOrderEnum{
	"ASC":  ListExadbVmClustersSortOrderAsc,
	"DESC": ListExadbVmClustersSortOrderDesc,
}

var mappingListExadbVmClustersSortOrderEnumLowerCase = map[string]ListExadbVmClustersSortOrderEnum{
	"asc":  ListExadbVmClustersSortOrderAsc,
	"desc": ListExadbVmClustersSortOrderDesc,
}

// GetListExadbVmClustersSortOrderEnumValues Enumerates the set of values for ListExadbVmClustersSortOrderEnum
func GetListExadbVmClustersSortOrderEnumValues() []ListExadbVmClustersSortOrderEnum {
	values := make([]ListExadbVmClustersSortOrderEnum, 0)
	for _, v := range mappingListExadbVmClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExadbVmClustersSortOrderEnumStringValues Enumerates the set of values in String for ListExadbVmClustersSortOrderEnum
func GetListExadbVmClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExadbVmClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExadbVmClustersSortOrderEnum(val string) (ListExadbVmClustersSortOrderEnum, bool) {
	enum, ok := mappingListExadbVmClustersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
