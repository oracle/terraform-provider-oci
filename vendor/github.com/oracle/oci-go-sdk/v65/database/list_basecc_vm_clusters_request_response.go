// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBaseccVmClustersRequest wrapper for the ListBaseccVmClusters operation
type ListBaseccVmClustersRequest struct {

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// If provided, filters the results for the given BICC Infrastructure.
	BaseInfrastructureId *string `mandatory:"false" contributesTo:"query" name:"baseInfrastructureId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListBaseccVmClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListBaseccVmClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState BaseccVmClusterSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only basecc vmclusters that match the given vmcluster type exactly.
	VmClusterType BaseccVmClusterSummaryVmClusterTypeEnum `mandatory:"false" contributesTo:"query" name:"vmClusterType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBaseccVmClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBaseccVmClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBaseccVmClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBaseccVmClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBaseccVmClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListBaseccVmClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBaseccVmClustersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBaseccVmClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBaseccVmClustersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseccVmClusterSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBaseccVmClusterSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseccVmClusterSummaryVmClusterTypeEnum(string(request.VmClusterType)); !ok && request.VmClusterType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VmClusterType: %s. Supported values are: %s.", request.VmClusterType, strings.Join(GetBaseccVmClusterSummaryVmClusterTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBaseccVmClustersResponse wrapper for the ListBaseccVmClusters operation
type ListBaseccVmClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []BaseccVmClusterSummary instances
	Items []BaseccVmClusterSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBaseccVmClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBaseccVmClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBaseccVmClustersSortByEnum Enum with underlying type: string
type ListBaseccVmClustersSortByEnum string

// Set of constants representing the allowable values for ListBaseccVmClustersSortByEnum
const (
	ListBaseccVmClustersSortByTimecreated ListBaseccVmClustersSortByEnum = "TIMECREATED"
	ListBaseccVmClustersSortByDisplayname ListBaseccVmClustersSortByEnum = "DISPLAYNAME"
)

var mappingListBaseccVmClustersSortByEnum = map[string]ListBaseccVmClustersSortByEnum{
	"TIMECREATED": ListBaseccVmClustersSortByTimecreated,
	"DISPLAYNAME": ListBaseccVmClustersSortByDisplayname,
}

var mappingListBaseccVmClustersSortByEnumLowerCase = map[string]ListBaseccVmClustersSortByEnum{
	"timecreated": ListBaseccVmClustersSortByTimecreated,
	"displayname": ListBaseccVmClustersSortByDisplayname,
}

// GetListBaseccVmClustersSortByEnumValues Enumerates the set of values for ListBaseccVmClustersSortByEnum
func GetListBaseccVmClustersSortByEnumValues() []ListBaseccVmClustersSortByEnum {
	values := make([]ListBaseccVmClustersSortByEnum, 0)
	for _, v := range mappingListBaseccVmClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBaseccVmClustersSortByEnumStringValues Enumerates the set of values in String for ListBaseccVmClustersSortByEnum
func GetListBaseccVmClustersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListBaseccVmClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBaseccVmClustersSortByEnum(val string) (ListBaseccVmClustersSortByEnum, bool) {
	enum, ok := mappingListBaseccVmClustersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBaseccVmClustersSortOrderEnum Enum with underlying type: string
type ListBaseccVmClustersSortOrderEnum string

// Set of constants representing the allowable values for ListBaseccVmClustersSortOrderEnum
const (
	ListBaseccVmClustersSortOrderAsc  ListBaseccVmClustersSortOrderEnum = "ASC"
	ListBaseccVmClustersSortOrderDesc ListBaseccVmClustersSortOrderEnum = "DESC"
)

var mappingListBaseccVmClustersSortOrderEnum = map[string]ListBaseccVmClustersSortOrderEnum{
	"ASC":  ListBaseccVmClustersSortOrderAsc,
	"DESC": ListBaseccVmClustersSortOrderDesc,
}

var mappingListBaseccVmClustersSortOrderEnumLowerCase = map[string]ListBaseccVmClustersSortOrderEnum{
	"asc":  ListBaseccVmClustersSortOrderAsc,
	"desc": ListBaseccVmClustersSortOrderDesc,
}

// GetListBaseccVmClustersSortOrderEnumValues Enumerates the set of values for ListBaseccVmClustersSortOrderEnum
func GetListBaseccVmClustersSortOrderEnumValues() []ListBaseccVmClustersSortOrderEnum {
	values := make([]ListBaseccVmClustersSortOrderEnum, 0)
	for _, v := range mappingListBaseccVmClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBaseccVmClustersSortOrderEnumStringValues Enumerates the set of values in String for ListBaseccVmClustersSortOrderEnum
func GetListBaseccVmClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBaseccVmClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBaseccVmClustersSortOrderEnum(val string) (ListBaseccVmClustersSortOrderEnum, bool) {
	enum, ok := mappingListBaseccVmClustersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
