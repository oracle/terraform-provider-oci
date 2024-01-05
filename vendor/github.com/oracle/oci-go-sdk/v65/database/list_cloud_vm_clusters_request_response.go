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

// ListCloudVmClustersRequest wrapper for the ListCloudVmClusters operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListCloudVmClusters.go.html to see an example of how to use ListCloudVmClustersRequest.
type ListCloudVmClustersRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// If provided, filters the results for the specified cloud Exadata infrastructure.
	CloudExadataInfrastructureId *string `mandatory:"false" contributesTo:"query" name:"cloudExadataInfrastructureId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListCloudVmClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCloudVmClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only cloud VM clusters that match the given lifecycle state exactly.
	LifecycleState CloudVmClusterSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudVmClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudVmClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudVmClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudVmClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudVmClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudVmClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudVmClustersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudVmClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudVmClustersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCloudVmClusterSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCloudVmClusterSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudVmClustersResponse wrapper for the ListCloudVmClusters operation
type ListCloudVmClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CloudVmClusterSummary instances
	Items []CloudVmClusterSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudVmClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudVmClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudVmClustersSortByEnum Enum with underlying type: string
type ListCloudVmClustersSortByEnum string

// Set of constants representing the allowable values for ListCloudVmClustersSortByEnum
const (
	ListCloudVmClustersSortByTimecreated ListCloudVmClustersSortByEnum = "TIMECREATED"
	ListCloudVmClustersSortByDisplayname ListCloudVmClustersSortByEnum = "DISPLAYNAME"
)

var mappingListCloudVmClustersSortByEnum = map[string]ListCloudVmClustersSortByEnum{
	"TIMECREATED": ListCloudVmClustersSortByTimecreated,
	"DISPLAYNAME": ListCloudVmClustersSortByDisplayname,
}

var mappingListCloudVmClustersSortByEnumLowerCase = map[string]ListCloudVmClustersSortByEnum{
	"timecreated": ListCloudVmClustersSortByTimecreated,
	"displayname": ListCloudVmClustersSortByDisplayname,
}

// GetListCloudVmClustersSortByEnumValues Enumerates the set of values for ListCloudVmClustersSortByEnum
func GetListCloudVmClustersSortByEnumValues() []ListCloudVmClustersSortByEnum {
	values := make([]ListCloudVmClustersSortByEnum, 0)
	for _, v := range mappingListCloudVmClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudVmClustersSortByEnumStringValues Enumerates the set of values in String for ListCloudVmClustersSortByEnum
func GetListCloudVmClustersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCloudVmClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudVmClustersSortByEnum(val string) (ListCloudVmClustersSortByEnum, bool) {
	enum, ok := mappingListCloudVmClustersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudVmClustersSortOrderEnum Enum with underlying type: string
type ListCloudVmClustersSortOrderEnum string

// Set of constants representing the allowable values for ListCloudVmClustersSortOrderEnum
const (
	ListCloudVmClustersSortOrderAsc  ListCloudVmClustersSortOrderEnum = "ASC"
	ListCloudVmClustersSortOrderDesc ListCloudVmClustersSortOrderEnum = "DESC"
)

var mappingListCloudVmClustersSortOrderEnum = map[string]ListCloudVmClustersSortOrderEnum{
	"ASC":  ListCloudVmClustersSortOrderAsc,
	"DESC": ListCloudVmClustersSortOrderDesc,
}

var mappingListCloudVmClustersSortOrderEnumLowerCase = map[string]ListCloudVmClustersSortOrderEnum{
	"asc":  ListCloudVmClustersSortOrderAsc,
	"desc": ListCloudVmClustersSortOrderDesc,
}

// GetListCloudVmClustersSortOrderEnumValues Enumerates the set of values for ListCloudVmClustersSortOrderEnum
func GetListCloudVmClustersSortOrderEnumValues() []ListCloudVmClustersSortOrderEnum {
	values := make([]ListCloudVmClustersSortOrderEnum, 0)
	for _, v := range mappingListCloudVmClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudVmClustersSortOrderEnumStringValues Enumerates the set of values in String for ListCloudVmClustersSortOrderEnum
func GetListCloudVmClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudVmClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudVmClustersSortOrderEnum(val string) (ListCloudVmClustersSortOrderEnum, bool) {
	enum, ok := mappingListCloudVmClustersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
