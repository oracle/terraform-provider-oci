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

// ListVmClusterNetworksRequest wrapper for the ListVmClusterNetworks operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListVmClusterNetworks.go.html to see an example of how to use ListVmClusterNetworksRequest.
type ListVmClusterNetworksRequest struct {

	// The Exadata infrastructure OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ExadataInfrastructureId *string `mandatory:"true" contributesTo:"path" name:"exadataInfrastructureId"`

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListVmClusterNetworksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListVmClusterNetworksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState VmClusterNetworkSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVmClusterNetworksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVmClusterNetworksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVmClusterNetworksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVmClusterNetworksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVmClusterNetworksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListVmClusterNetworksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVmClusterNetworksSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVmClusterNetworksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVmClusterNetworksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVmClusterNetworkSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetVmClusterNetworkSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVmClusterNetworksResponse wrapper for the ListVmClusterNetworks operation
type ListVmClusterNetworksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []VmClusterNetworkSummary instances
	Items []VmClusterNetworkSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVmClusterNetworksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVmClusterNetworksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVmClusterNetworksSortByEnum Enum with underlying type: string
type ListVmClusterNetworksSortByEnum string

// Set of constants representing the allowable values for ListVmClusterNetworksSortByEnum
const (
	ListVmClusterNetworksSortByTimecreated ListVmClusterNetworksSortByEnum = "TIMECREATED"
	ListVmClusterNetworksSortByDisplayname ListVmClusterNetworksSortByEnum = "DISPLAYNAME"
)

var mappingListVmClusterNetworksSortByEnum = map[string]ListVmClusterNetworksSortByEnum{
	"TIMECREATED": ListVmClusterNetworksSortByTimecreated,
	"DISPLAYNAME": ListVmClusterNetworksSortByDisplayname,
}

var mappingListVmClusterNetworksSortByEnumLowerCase = map[string]ListVmClusterNetworksSortByEnum{
	"timecreated": ListVmClusterNetworksSortByTimecreated,
	"displayname": ListVmClusterNetworksSortByDisplayname,
}

// GetListVmClusterNetworksSortByEnumValues Enumerates the set of values for ListVmClusterNetworksSortByEnum
func GetListVmClusterNetworksSortByEnumValues() []ListVmClusterNetworksSortByEnum {
	values := make([]ListVmClusterNetworksSortByEnum, 0)
	for _, v := range mappingListVmClusterNetworksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVmClusterNetworksSortByEnumStringValues Enumerates the set of values in String for ListVmClusterNetworksSortByEnum
func GetListVmClusterNetworksSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListVmClusterNetworksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVmClusterNetworksSortByEnum(val string) (ListVmClusterNetworksSortByEnum, bool) {
	enum, ok := mappingListVmClusterNetworksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVmClusterNetworksSortOrderEnum Enum with underlying type: string
type ListVmClusterNetworksSortOrderEnum string

// Set of constants representing the allowable values for ListVmClusterNetworksSortOrderEnum
const (
	ListVmClusterNetworksSortOrderAsc  ListVmClusterNetworksSortOrderEnum = "ASC"
	ListVmClusterNetworksSortOrderDesc ListVmClusterNetworksSortOrderEnum = "DESC"
)

var mappingListVmClusterNetworksSortOrderEnum = map[string]ListVmClusterNetworksSortOrderEnum{
	"ASC":  ListVmClusterNetworksSortOrderAsc,
	"DESC": ListVmClusterNetworksSortOrderDesc,
}

var mappingListVmClusterNetworksSortOrderEnumLowerCase = map[string]ListVmClusterNetworksSortOrderEnum{
	"asc":  ListVmClusterNetworksSortOrderAsc,
	"desc": ListVmClusterNetworksSortOrderDesc,
}

// GetListVmClusterNetworksSortOrderEnumValues Enumerates the set of values for ListVmClusterNetworksSortOrderEnum
func GetListVmClusterNetworksSortOrderEnumValues() []ListVmClusterNetworksSortOrderEnum {
	values := make([]ListVmClusterNetworksSortOrderEnum, 0)
	for _, v := range mappingListVmClusterNetworksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVmClusterNetworksSortOrderEnumStringValues Enumerates the set of values in String for ListVmClusterNetworksSortOrderEnum
func GetListVmClusterNetworksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVmClusterNetworksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVmClusterNetworksSortOrderEnum(val string) (ListVmClusterNetworksSortOrderEnum, bool) {
	enum, ok := mappingListVmClusterNetworksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
